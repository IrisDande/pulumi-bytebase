// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace bytebaseDatabase.bytebase.Inputs
{

    public sealed class bytebaseServerlessSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The public cloud where you would like your index hosted.
        /// </summary>
        [Input("cloud", required: true)]
        public Input<bytebaseDatabase.bytebase.ServerlessSpecCloud> Cloud { get; set; } = null!;

        /// <summary>
        /// The region where you would like your index to be created. Different cloud providers have different regions available.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public bytebaseServerlessSpecArgs()
        {
        }
        public static new bytebaseServerlessSpecArgs Empty => new bytebaseServerlessSpecArgs();
    }
}
