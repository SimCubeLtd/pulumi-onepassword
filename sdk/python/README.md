# 1Password Provider
[1Password Website](1password.com)

Based on the [1Password terraform provider](https://github.com/1Password/terraform-provider-onepassword)

&nbsp;

The 1Password resource provider for Pulumi lets you use 1Password resources in your Infrastructure as Code deployments.


## Installing

This package is available in C#, TypeScript, Python and Go

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Onepassword
```

### Node.js (Java/TypeScript)

```bash
npm install @SimCubeLtd/pulumi-onepassword
```

or `yarn`:

```bash
yarn add @SimCubeLtd/pulumi-onepassword
```

### Python

To use from Python, install using `pip`:

```bash
pip install simcubeltd_pulumi_onepassword
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/SimCubeLtd/pulumi-onepassword/sdk/go/onepassword
```

## Configuration

The following configuration entries are available:

| **Key**           | **Value**                                            |
|-------------------|:-----------------------------------------------------|
| onepassword:token | The access token for your 1Password Connect server   |
| onepassword:url   | URL where your 1Password Connect API is being served |
