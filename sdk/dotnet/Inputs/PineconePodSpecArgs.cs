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

    public sealed class bytebasePodSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The environment where the index is hosted.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// Configuration for the behavior of bytebase's internal metadata index.
        /// </summary>
        [Input("metaDataConfig")]
        public Input<Inputs.MetaDataConfigArgs>? MetaDataConfig { get; set; }

        /// <summary>
        /// The type of pod to use. One of `s1`, `p1`, or `p2` appended with `.` and one of `x1`, `x2`, `x4`, or `x8`.
        /// </summary>
        [Input("podType", required: true)]
        public Input<string> PodType { get; set; } = null!;

        /// <summary>
        /// The number of pods to be used in the index. This should be equal to `shards` x `replicas`.
        /// </summary>
        [Input("pods")]
        public Input<int>? Pods { get; set; }

        /// <summary>
        /// The number of replicas. Replicas duplicate your index. They provide higher availability and throughput. Replicas can be scaled up or down as your needs change.
        /// </summary>
        [Input("replicas", required: true)]
        public Input<int> Replicas { get; set; } = null!;

        /// <summary>
        /// The number of shards. Shards split your data across multiple pods so you can fit more data into an index.
        /// </summary>
        [Input("shards")]
        public Input<int>? Shards { get; set; }

        /// <summary>
        /// The name of the collection to be used as the source for the index.
        /// </summary>
        [Input("sourceCollection")]
        public Input<string>? SourceCollection { get; set; }

        public bytebasePodSpecArgs()
        {
        }
        public static new bytebasePodSpecArgs Empty => new bytebasePodSpecArgs();
    }
}
