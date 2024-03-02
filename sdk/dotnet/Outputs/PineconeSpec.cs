// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PineconeDatabase.Pinecone.Outputs
{

    [OutputType]
    public sealed class PineconeSpec
    {
        /// <summary>
        /// Configuration needed to deploy a pod index.
        /// </summary>
        public readonly Outputs.PineconePodSpec? Pod;
        /// <summary>
        /// Configuration needed to deploy a serverless index.
        /// </summary>
        public readonly Outputs.PineconeServerlessSpec? Serverless;

        [OutputConstructor]
        private PineconeSpec(
            Outputs.PineconePodSpec? pod,

            Outputs.PineconeServerlessSpec? serverless)
        {
            Pod = pod;
            Serverless = serverless;
        }
    }
}
