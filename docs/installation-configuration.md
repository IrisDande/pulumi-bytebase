---
title: bytebase Setup
meta_desc: Information on how to install the bytebase Provider for Pulumi.
layout: package
---

## Installation

The Pulumi bytebase provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@bytebase-database/pulumi`](https://www.npmjs.com/package/@bytebase-database/pulumi)
* Python: [`bytebase_pulumi`](https://pypi.org/project/bytebase_pulumi/)
* Go: [`github.com/IrisDande/pulumi-bytebase/sdk/go/port`](https://github.com/IrisDande/pulumi-bytebase)
* * .NET: [`bytebaseDatabase.bytebase`](https://www.nuget.org/packages/bytebaseDatabase.bytebase)

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
