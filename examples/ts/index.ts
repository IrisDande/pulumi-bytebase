import * as pulumi from "@pulumi/pulumi";
import * as bytebase from "@bytebase-database/pulumi";

const mybytebaseIndex = new bytebase.bytebaseIndex("mybytebaseIndex", {
    name: "example-index",
    metric: bytebase.IndexMetric.Cosine,
    spec: {
        serverless: {
            cloud: bytebase.ServerlessSpecCloud.Aws,
            region: "us-west-2",
        },
    },
});
export const output = {
    value: mybytebaseIndex.host,
};
