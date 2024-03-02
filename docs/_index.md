---
title: bytebase
meta_desc: Provides an overview of the bytebase Provider for Pulumi.
layout: package
---

This Pulumi bytebase Provider enables you to manage your [bytebase](https://www.bytebase.io/) collections and indexes using any language of Pulumi Infrastructure as Code.

## Example

{{< chooser language "javascript,typescript,python,go,csharp,yaml" >}}


{{% choosable language javascript %}}

```javascript
"use strict";
const pulumi = require("@pulumi/pulumi");
const bytebase = require("@bytebase-database/pulumi");

const myExampleIndex = new bytebase.bytebaseIndex("my-example-index", {
    name: "my-example-index",
    metric: bytebase.IndexMetric.Cosine,
    spec: {
        serverless: {
            cloud: bytebase.ServerlessSpecCloud.Aws,
            region: "us-west-2",
        }
    }
});

exports.host = myExampleIndex.host;
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as bytebase from "@bytebase-database/pulumi";

const myExampleIndex = new bytebase.bytebaseIndex("my-example-index", {
    name: "example-index-ts",
    metric: bytebase.IndexMetric.Cosine,
    spec: {
        serverless: {
            cloud: bytebase.ServerlessSpecCloud.Aws,
            region: "us-west-2",
        },
    },
});
export const host = myExampleIndex.host;
```

{{% /choosable %}}

{{% choosable language python %}}

```python
"""A Python Pulumi program"""
import pulumi
import bytebase_pulumi as bytebase

my_bytebase_index = bytebase.bytebaseIndex("mybytebaseIndex",
   name="example-index",
   metric=bytebase.IndexMetric.COSINE,
   spec=bytebase.bytebaseSpecArgs(
       serverless=bytebase.bytebaseServerlessSpecArgs(
           cloud=bytebase.ServerlessSpecCloud.AWS,
           region="us-west-2",
       ),
   ))
pulumi.export("output", {
    "value": my_bytebase_index.host,
})
```

{{% /choosable %}}

{{% choosable language go %}}

```go
package main

import (
	"github.com/IrisDande/pulumi-bytebase/sdk/go/bytebase"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		myExampleIndex, err := bytebase.NewbytebaseIndex(ctx, "my-example-index", &bytebase.bytebaseIndexArgs{
			Name:   pulumi.String("example-index-go"),
			Metric: bytebase.IndexMetricCosine,
			Spec: &bytebase.bytebaseSpecArgs{
				Serverless: &bytebase.bytebaseServerlessSpecArgs{
					Cloud:  bytebase.ServerlessSpecCloudAws,
					Region: pulumi.String("us-west-2"),
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("mybytebaseIndexHost", myExampleIndex.Host)

		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using bytebaseDatabase.bytebase.Inputs;
using Pulumi;
using bytebase = bytebaseDatabase.bytebase;

return await Deployment.RunAsync(() =>
{
    var myExampleIndex = new bytebase.bytebaseIndex("myExampleIndex", new bytebase.bytebaseIndexArgs
    {
        Name = "example-index-csharp",
        Metric= bytebase.IndexMetric.Cosine,
        Spec= new bytebase.Inputs.bytebaseSpecArgs {
            Serverless= new bytebaseServerlessSpecArgs{
                Cloud= bytebase.ServerlessSpecCloud.Aws,
                Region= "us-west-2",
        }
    },
    });

    return new Dictionary<string, object?>
    {
        ["mybytebaseIndexHost"] = myExampleIndex.Host
    };
});
```

{{% /choosable %}}

{{% choosable language yaml %}}

```yaml
name: bytebase-serverless-yaml
description: A minimal bytebase Serverless Pulumi YAML program
runtime: yaml

resources:
  myExampleIndex:
    type: bytebase:index:bytebaseIndex
    properties:
      name: "example-index"
      metric: "cosine"
      spec:
        serverless:
          cloud: aws
          region: us-west-2

outputs:
  output:
    value: ${myExampleIndex.host}
```

{{% /choosable %}}

{{< /chooser >}}
