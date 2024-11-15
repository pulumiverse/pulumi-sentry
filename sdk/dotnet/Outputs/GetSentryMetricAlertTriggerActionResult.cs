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
    public sealed class GetSentryMetricAlertTriggerActionResult
    {
        public readonly string Id;
        public readonly int IntegrationId;
        public readonly string TargetIdentifier;
        public readonly string TargetType;
        public readonly string Type;

        [OutputConstructor]
        private GetSentryMetricAlertTriggerActionResult(
            string id,

            int integrationId,

            string targetIdentifier,

            string targetType,

            string type)
        {
            Id = id;
            IntegrationId = integrationId;
            TargetIdentifier = targetIdentifier;
            TargetType = targetType;
            Type = type;
        }
    }
}
