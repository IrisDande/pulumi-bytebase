// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace bytebaseDatabase.bytebase.Outputs
{

    [OutputType]
    public sealed class bytebaseServerlessSpec
    {
        /// <summary>
        /// The public cloud where you would like your index hosted.
        /// </summary>
        public readonly bytebaseDatabase.bytebase.ServerlessSpecCloud Cloud;
        /// <summary>
        /// The region where you would like your index to be created. Different cloud providers have different regions available.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private bytebaseServerlessSpec(
            bytebaseDatabase.bytebase.ServerlessSpecCloud cloud,

            string region)
        {
            Cloud = cloud;
            Region = region;
        }
    }
}
