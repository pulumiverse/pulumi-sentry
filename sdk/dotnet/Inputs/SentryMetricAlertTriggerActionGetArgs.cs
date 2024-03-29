// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Sentry.Inputs
{

    public sealed class SentryMetricAlertTriggerActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("integrationId")]
        public Input<int>? IntegrationId { get; set; }

        [Input("targetIdentifier")]
        public Input<string>? TargetIdentifier { get; set; }

        [Input("targetType", required: true)]
        public Input<string> TargetType { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SentryMetricAlertTriggerActionGetArgs()
        {
        }
        public static new SentryMetricAlertTriggerActionGetArgs Empty => new SentryMetricAlertTriggerActionGetArgs();
    }
}
