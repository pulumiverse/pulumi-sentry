---
title: Sentry
meta_desc: Provides an overview of the Sentry Provider for Pulumi.
layout: overview
---

The Sentry provider for Pulumi can be used to provision anything in [Sentry](https://sentry.io).
The Sentry provider must be configured with credentials to deploy and update resources in Sentry.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as sentry from "@pulumiverse/sentry";
const db = new sentry.Database("example", {
    cloudProvider: "azure",
    keyspace: "default",
    regions: ["westus2"],
    name: "example-db"
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumiverse_sentry as sentry

db = sentry.Database("example",
    cloud_provider="azure",
    keyspace="default",
    regions=["westus2"],
    name="example-db"
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	sentry "github.com/pulumiverse/pulumi-sentry/sdk/go/sentry"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		db, err := sentry.NewDatabase(ctx, "example", &sentry.DatabaseArgs{
            CloudProvider: pulumi.String("azure"),
            Keyspace: pulumi.String("default"),
            Regions: pulumi.StringArray{
                pulumi.String("westus2")
            },
            Name: pulumi.String("example-db"),
		})
		if err != nil {
			return fmt.Errorf("error creating instance server: %v", err)
		}

		ctx.Export("dbId", db.Id)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using Pulumi;
using Pulumiverse.sentry;

class sentryDb : Stack
{
    public sentryDb()
    {
        var db = new Database("example", new DatabaseArgs{
            CloudProvider: "azure",
            Keyspace: "default",
            Regions: new[] {"westus2"},
            Name: "example-db"
        });
    }
}
```

{{% /choosable %}}

{{< /chooser >}}
