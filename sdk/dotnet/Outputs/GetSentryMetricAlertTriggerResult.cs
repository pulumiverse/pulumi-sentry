// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Sentry.Outputs
{

    [OutputType]
    public sealed class GetSentryMetricAlertTriggerResult
    {
        public readonly ImmutableArray<Outputs.GetSentryMetricAlertTriggerActionResult> Actions;
        public readonly double AlertThreshold;
        public readonly string Id;
        public readonly string Label;
        public readonly double ResolveThreshold;
        public readonly int ThresholdType;

        [OutputConstructor]
        private GetSentryMetricAlertTriggerResult(
            ImmutableArray<Outputs.GetSentryMetricAlertTriggerActionResult> actions,

            double alertThreshold,

            string id,

            string label,

            double resolveThreshold,

            int thresholdType)
        {
            Actions = actions;
            AlertThreshold = alertThreshold;
            Id = id;
            Label = label;
            ResolveThreshold = resolveThreshold;
            ThresholdType = thresholdType;
        }
    }
}
