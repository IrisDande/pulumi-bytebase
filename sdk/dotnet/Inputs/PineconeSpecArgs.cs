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

    public sealed class bytebaseSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration needed to deploy a pod index.
        /// </summary>
        [Input("pod")]
        public Input<Inputs.bytebasePodSpecArgs>? Pod { get; set; }

        /// <summary>
        /// Configuration needed to deploy a serverless index.
        /// </summary>
        [Input("serverless")]
        public Input<Inputs.bytebaseServerlessSpecArgs>? Serverless { get; set; }

        public bytebaseSpecArgs()
        {
        }
        public static new bytebaseSpecArgs Empty => new bytebaseSpecArgs();
    }
}
