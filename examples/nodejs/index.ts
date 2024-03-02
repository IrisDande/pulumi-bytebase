import * as pulumi from "@pulumi/pulumi";
import * as bytebase from "@pulumi/bytebase";

const myRandomResource = new bytebase.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
