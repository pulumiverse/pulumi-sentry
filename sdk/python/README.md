# Sentry Resource Provider

**This plugin requires pulumi cli v3.35.3 or later**

The Sentry Resource Provider lets you manage [Sentry](https://sentry.io/) resources.

It is based on the sentry terraform provider https://registry.terraform.io/providers/jianyuan/sentry/latest/docs

## Installing

If you get an error like this:

```
error: could not load plugin for sentry provider 'urn:pulumi:my-stack::my-project::pulumi:providers:sentry::my-provider': no resource plugin 'pulumi-resource-sentry' found in the workspace at version vX.Y.Z or on your $PATH, install the plugin using `pulumi plugin install resource sentry vX.Y.Z`
```

You can use this command to install the missing version:

```
pulumi plugin install resource sentry vX.Y.Z --server github://api.github.com/pulumiverse
```

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/sentry
```

or `yarn`:

```bash
yarn add @pulumiverse/sentry
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumiverse_sentry
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-sentry/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumiverse.Sentry
```

## Configuration

The following configuration points are available for the `sentry` provider:

-   `sentry:token` (Required) This is the Sentry authentication token. The value can be sourced from the SENTRY_TOKEN environment variable.
-   `sentry:baseUrl` (Optional) This is the target Sentry base API endpoint. The default value is https://sentry.io/api/. The value must be provided when working with Sentry On-Premise. The value can be sourced from the SENTRY_BASE_URL environment variable.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/sentry/api-docs/).
