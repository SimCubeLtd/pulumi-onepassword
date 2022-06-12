// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package onepassword

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/SimCubeLtd/pulumi-onepassword/provider/pkg/version"
	"github.com/1Password/terraform-provider-onepassword/onepassword"
)

const (
	mainPkg = "onepassword"
	mainMod = "index"
)

func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(onepassword.Provider())

	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "onepassword",
		DisplayName: "1Password",
		Publisher: "SimCubeLtd",
		LogoURL: "https://avatars.githubusercontent.com/u/38230737?s=200&v=4",
		PluginDownloadURL: "https://github.com/SimCubeLtd/pulumi-onepassword/releases/download/v${VERSION}",
		Description:       "A Pulumi package for creating and managing onepassword resources.",
		Keywords:   []string{"pulumi", "onepassword", "category/cloud"},
		License:    "MIT",
		Homepage:   "https://github.com/SimCubeLtd/pulumi-onepassword",
		Repository: "https://github.com/SimCubeLtd/pulumi-onepassword",
		GitHubOrg: "1Password",
		Config:    map[string]*tfbridge.SchemaInfo{},
		Resources:            map[string]*tfbridge.ResourceInfo{
		    "onepassword_item": {
		        Tok: tfbridge.MakeResource(mainPkg, mainMod, "Item"),
		        Fields: map[string]*tfbridge.SchemaInfo{
		            "password": {
		                Secret: tfbridge.True(),
		            },
		        },
		    },
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"onepassword_item": {
			    Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetItem"),
			},
			"onepassword_vault": {
  			    Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetVault"),
            },
		},
		JavaScript: &tfbridge.JavaScriptInfo{
		    PackageName: "@simcubeltd/pulumi-onepassword",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
		    PackageName: "simcubeltd_pulumi_onepassword",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/SimCubeLtd/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}