# Sentry Resource Provider

The Sentry Resource Provider lets you manage [Sentry](https://sentry.io/) resources.

It is based on the sentry terraform provider https://registry.terraform.io/providers/jianyuan/sentry/latest/docs

## Installing

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
