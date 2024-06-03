// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Sentry
{
    /// <summary>
    /// Manage a PagerDuty service integration.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Sentry = Pulumi.Sentry;
    /// using Sentry = Pulumiverse.Sentry;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pagerduty = Sentry.GetSentryOrganizationIntegration.Invoke(new()
    ///     {
    ///         Organization = "my-organization",
    ///         ProviderKey = "pagerduty",
    ///         Name = "my-pagerduty-organization",
    ///     });
    /// 
    ///     // Associate a PagerDuty service and integration key with a Sentry PagerDuty integration
    ///     var test = new Sentry.SentryIntegrationPagerDuty("test", new()
    ///     {
    ///         Organization = "my-organization",
    ///         IntegrationId = pagerduty.Apply(getSentryOrganizationIntegrationResult =&gt; getSentryOrganizationIntegrationResult.Id),
    ///         Service = "my-pagerduty-service",
    ///         IntegrationKey = "my-pagerduty-integration-key",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import using the organization slug from the URL:
    /// 
    /// https://sentry.io/api/0/organizations/[org-slug]/integrations/
    /// 
    /// [integration-id] is the top-level `id` of the PagerDuty organization integration
    /// 
    /// [service-id] is the `id` of the service_table record to import under the configData property
    /// 
    /// ```sh
    /// $ pulumi import sentry:index/sentryIntegrationPagerDuty:SentryIntegrationPagerDuty default org-slug/integration-id/service-id
    /// ```
    /// </summary>
    [SentryResourceType("sentry:index/sentryIntegrationPagerDuty:SentryIntegrationPagerDuty")]
    public partial class SentryIntegrationPagerDuty : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the PagerDuty integration. Source from the URL `https://&lt;organization&gt;.sentry.io/settings/integrations/pagerduty/&lt;integration-id&gt;/` or use the `sentry.getSentryOrganizationIntegration` data source.
        /// </summary>
        [Output("integrationId")]
        public Output<string> IntegrationId { get; private set; } = null!;

        /// <summary>
        /// The integration key of the PagerDuty service.
        /// </summary>
        [Output("integrationKey")]
        public Output<string> IntegrationKey { get; private set; } = null!;

        /// <summary>
        /// The slug of the organization the resource belongs to.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        /// <summary>
        /// The name of the PagerDuty service.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;


        /// <summary>
        /// Create a SentryIntegrationPagerDuty resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SentryIntegrationPagerDuty(string name, SentryIntegrationPagerDutyArgs args, CustomResourceOptions? options = null)
            : base("sentry:index/sentryIntegrationPagerDuty:SentryIntegrationPagerDuty", name, args ?? new SentryIntegrationPagerDutyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SentryIntegrationPagerDuty(string name, Input<string> id, SentryIntegrationPagerDutyState? state = null, CustomResourceOptions? options = null)
            : base("sentry:index/sentryIntegrationPagerDuty:SentryIntegrationPagerDuty", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SentryIntegrationPagerDuty resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SentryIntegrationPagerDuty Get(string name, Input<string> id, SentryIntegrationPagerDutyState? state = null, CustomResourceOptions? options = null)
        {
            return new SentryIntegrationPagerDuty(name, id, state, options);
        }
    }

    public sealed class SentryIntegrationPagerDutyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the PagerDuty integration. Source from the URL `https://&lt;organization&gt;.sentry.io/settings/integrations/pagerduty/&lt;integration-id&gt;/` or use the `sentry.getSentryOrganizationIntegration` data source.
        /// </summary>
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        /// <summary>
        /// The integration key of the PagerDuty service.
        /// </summary>
        [Input("integrationKey", required: true)]
        public Input<string> IntegrationKey { get; set; } = null!;

        /// <summary>
        /// The slug of the organization the resource belongs to.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// The name of the PagerDuty service.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public SentryIntegrationPagerDutyArgs()
        {
        }
        public static new SentryIntegrationPagerDutyArgs Empty => new SentryIntegrationPagerDutyArgs();
    }

    public sealed class SentryIntegrationPagerDutyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the PagerDuty integration. Source from the URL `https://&lt;organization&gt;.sentry.io/settings/integrations/pagerduty/&lt;integration-id&gt;/` or use the `sentry.getSentryOrganizationIntegration` data source.
        /// </summary>
        [Input("integrationId")]
        public Input<string>? IntegrationId { get; set; }

        /// <summary>
        /// The integration key of the PagerDuty service.
        /// </summary>
        [Input("integrationKey")]
        public Input<string>? IntegrationKey { get; set; }

        /// <summary>
        /// The slug of the organization the resource belongs to.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// The name of the PagerDuty service.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public SentryIntegrationPagerDutyState()
        {
        }
        public static new SentryIntegrationPagerDutyState Empty => new SentryIntegrationPagerDutyState();
    }
}
