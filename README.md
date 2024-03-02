# bytebase Pulumi Provider

<img src="img/bytebase.svg" width="50%">

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/bytebase-io/pulumi-bytebase)

This Pulumi bytebase Provider enables you to manage your [bytebase](https://www.bytebase.io/) collections and indexes using any language of Pulumi Infrastructure as Code.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @bytebase-database/pulumi
```

or `yarn`:

```bash
yarn add @bytebase-database/bytebase
```

### Python

To use from Python, install using `pip`:

```bash
pip install bytebase_pulumi
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/IrisDande/pulumi-bytebase/sdk
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package bytebaseDatabase.bytebase
```

## Configuration

The following configuration points are available for the `bytebase` provider:

- `bytebase:APIKey` - This is the bytebase API key. (environment: `bytebase_API_KEY`)
