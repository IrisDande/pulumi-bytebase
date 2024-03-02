import pulumi
import pulumi_bytebase as bytebase

my_random_resource = bytebase.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
