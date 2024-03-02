import pulumi
import pulumi_bytebase as bytebase

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
