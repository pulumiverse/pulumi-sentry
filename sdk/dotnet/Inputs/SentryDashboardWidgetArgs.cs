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

    public sealed class SentryDashboardWidgetArgs : global::Pulumi.ResourceArgs
    {
        [Input("displayType", required: true)]
        public Input<string> DisplayType { get; set; } = null!;

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("interval")]
        public Input<string>? Interval { get; set; }

        [Input("layout", required: true)]
        public Input<Inputs.SentryDashboardWidgetLayoutArgs> Layout { get; set; } = null!;

        [Input("limit")]
        public Input<int>? Limit { get; set; }

        [Input("queries", required: true)]
        private InputList<Inputs.SentryDashboardWidgetQueryArgs>? _queries;
        public InputList<Inputs.SentryDashboardWidgetQueryArgs> Queries
        {
            get => _queries ?? (_queries = new InputList<Inputs.SentryDashboardWidgetQueryArgs>());
            set => _queries = value;
        }

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        [Input("widgetType")]
        public Input<string>? WidgetType { get; set; }

        public SentryDashboardWidgetArgs()
        {
        }
        public static new SentryDashboardWidgetArgs Empty => new SentryDashboardWidgetArgs();
    }
}
