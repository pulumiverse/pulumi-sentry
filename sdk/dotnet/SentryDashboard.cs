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
    /// Sentry Dashboard resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Sentry = Pulumiverse.Sentry;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Sentry.SentryDashboard("main", new()
    ///     {
    ///         Organization = data.Sentry_organization.Main.Id,
    ///         Title = "Test dashboard",
    ///         Widgets = new[]
    ///         {
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Number of Errors",
    ///                 DisplayType = "big_number",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Conditions = "!event.type:transaction",
    ///                         OrderBy = "count()",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 0,
    ///                     Y = 0,
    ///                     W = 1,
    ///                     H = 1,
    ///                     MinH = 1,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Number of Issues",
    ///                 DisplayType = "big_number",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "count_unique(issue)",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count_unique(issue)",
    ///                         },
    ///                         Conditions = "!event.type:transaction",
    ///                         OrderBy = "count_unique(issue)",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 1,
    ///                     Y = 0,
    ///                     W = 1,
    ///                     H = 1,
    ///                     MinH = 1,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Events",
    ///                 DisplayType = "line",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Name = "Events",
    ///                         Fields = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Conditions = "!event.type:transaction",
    ///                         OrderBy = "count()",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 2,
    ///                     Y = 0,
    ///                     W = 4,
    ///                     H = 2,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Affected Users",
    ///                 DisplayType = "line",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Name = "Known Users",
    ///                         Fields = new[]
    ///                         {
    ///                             "count_unique(user)",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count_unique(user)",
    ///                         },
    ///                         Conditions = "has:user.email !event.type:transaction",
    ///                         OrderBy = "count_unique(user)",
    ///                     },
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Name = "Anonymous Users",
    ///                         Fields = new[]
    ///                         {
    ///                             "count_unique(user)",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count_unique(user)",
    ///                         },
    ///                         Conditions = "!has:user.email !event.type:transaction",
    ///                         OrderBy = "count_unique(user)",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 1,
    ///                     Y = 2,
    ///                     W = 1,
    ///                     H = 2,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Handled vs. Unhandled",
    ///                 DisplayType = "line",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Name = "Handled",
    ///                         Fields = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Conditions = "error.handled:true",
    ///                         OrderBy = "count()",
    ///                     },
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Name = "Unhandled",
    ///                         Fields = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Conditions = "error.handled:false",
    ///                         OrderBy = "count()",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 0,
    ///                     Y = 2,
    ///                     W = 1,
    ///                     H = 2,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Errors by Country",
    ///                 DisplayType = "table",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "geo.country_code",
    ///                             "geo.region",
    ///                             "count()",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Conditions = "!event.type:transaction has:geo.country_code",
    ///                         OrderBy = "count()",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 4,
    ///                     Y = 6,
    ///                     W = 2,
    ///                     H = 4,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "High Throughput Transactions",
    ///                 DisplayType = "table",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "count()",
    ///                             "transaction",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Columns = new[]
    ///                         {
    ///                             "transaction",
    ///                         },
    ///                         Conditions = "!event.type:error",
    ///                         OrderBy = "-count()",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 0,
    ///                     Y = 6,
    ///                     W = 2,
    ///                     H = 4,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Errors by Browser",
    ///                 DisplayType = "table",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "browser.name",
    ///                             "count()",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Columns = new[]
    ///                         {
    ///                             "browser.name",
    ///                         },
    ///                         Conditions = "!event.type:transaction has:browser.name",
    ///                         OrderBy = "-count()",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 5,
    ///                     Y = 2,
    ///                     W = 1,
    ///                     H = 4,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Overall User Misery",
    ///                 DisplayType = "big_number",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "user_misery(300)",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "user_misery(300)",
    ///                         },
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 0,
    ///                     Y = 1,
    ///                     W = 1,
    ///                     H = 1,
    ///                     MinH = 1,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Overall Apdex",
    ///                 DisplayType = "big_number",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "apdex(300)",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "apdex(300)",
    ///                         },
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 1,
    ///                     Y = 1,
    ///                     W = 1,
    ///                     H = 1,
    ///                     MinH = 1,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "High Throughput Transactions",
    ///                 DisplayType = "top_n",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "transaction",
    ///                             "count()",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Columns = new[]
    ///                         {
    ///                             "transaction",
    ///                         },
    ///                         Conditions = "!event.type:error",
    ///                         OrderBy = "-count()",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 0,
    ///                     Y = 4,
    ///                     W = 2,
    ///                     H = 2,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Issues Assigned to Me or My Teams",
    ///                 DisplayType = "table",
    ///                 Interval = "5m",
    ///                 WidgetType = "issue",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "assignee",
    ///                             "issue",
    ///                             "title",
    ///                         },
    ///                         Columns = new[]
    ///                         {
    ///                             "assignee",
    ///                             "issue",
    ///                             "title",
    ///                         },
    ///                         Conditions = "assigned_or_suggested:me is:unresolved",
    ///                         OrderBy = "priority",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 2,
    ///                     Y = 2,
    ///                     W = 2,
    ///                     H = 4,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Transactions Ordered by Misery",
    ///                 DisplayType = "table",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "transaction",
    ///                             "user_misery(300)",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "user_misery(300)",
    ///                         },
    ///                         Columns = new[]
    ///                         {
    ///                             "transaction",
    ///                         },
    ///                         OrderBy = "-user_misery(300)",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 2,
    ///                     Y = 6,
    ///                     W = 2,
    ///                     H = 4,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///             new Sentry.Inputs.SentryDashboardWidgetArgs
    ///             {
    ///                 Title = "Errors by Browser Over Time",
    ///                 DisplayType = "top_n",
    ///                 Interval = "5m",
    ///                 WidgetType = "discover",
    ///                 Queries = new[]
    ///                 {
    ///                     new Sentry.Inputs.SentryDashboardWidgetQueryArgs
    ///                     {
    ///                         Fields = new[]
    ///                         {
    ///                             "browser.name",
    ///                             "count()",
    ///                         },
    ///                         Aggregates = new[]
    ///                         {
    ///                             "count()",
    ///                         },
    ///                         Columns = new[]
    ///                         {
    ///                             "browser.name",
    ///                         },
    ///                         Conditions = "event.type:error has:browser.name",
    ///                         OrderBy = "-count()",
    ///                     },
    ///                 },
    ///                 Layout = new Sentry.Inputs.SentryDashboardWidgetLayoutArgs
    ///                 {
    ///                     X = 4,
    ///                     Y = 2,
    ///                     W = 1,
    ///                     H = 4,
    ///                     MinH = 2,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import using the dashboard id from the URL:
    /// 
    /// https://sentry.io/dashboard/[dashboard-id]
    /// 
    /// ```sh
    /// $ pulumi import sentry:index/sentryDashboard:SentryDashboard default org-slug/dashboard-id
    /// ```
    /// </summary>
    [SentryResourceType("sentry:index/sentryDashboard:SentryDashboard")]
    public partial class SentryDashboard : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The internal ID for this dashboard.
        /// </summary>
        [Output("internalId")]
        public Output<string> InternalId { get; private set; } = null!;

        /// <summary>
        /// The slug of the organization the dashboard belongs to.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        /// <summary>
        /// Dashboard title.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Dashboard widgets.
        /// </summary>
        [Output("widgets")]
        public Output<ImmutableArray<Outputs.SentryDashboardWidget>> Widgets { get; private set; } = null!;


        /// <summary>
        /// Create a SentryDashboard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SentryDashboard(string name, SentryDashboardArgs args, CustomResourceOptions? options = null)
            : base("sentry:index/sentryDashboard:SentryDashboard", name, args ?? new SentryDashboardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SentryDashboard(string name, Input<string> id, SentryDashboardState? state = null, CustomResourceOptions? options = null)
            : base("sentry:index/sentryDashboard:SentryDashboard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SentryDashboard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SentryDashboard Get(string name, Input<string> id, SentryDashboardState? state = null, CustomResourceOptions? options = null)
        {
            return new SentryDashboard(name, id, state, options);
        }
    }

    public sealed class SentryDashboardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The slug of the organization the dashboard belongs to.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// Dashboard title.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        [Input("widgets")]
        private InputList<Inputs.SentryDashboardWidgetArgs>? _widgets;

        /// <summary>
        /// Dashboard widgets.
        /// </summary>
        public InputList<Inputs.SentryDashboardWidgetArgs> Widgets
        {
            get => _widgets ?? (_widgets = new InputList<Inputs.SentryDashboardWidgetArgs>());
            set => _widgets = value;
        }

        public SentryDashboardArgs()
        {
        }
        public static new SentryDashboardArgs Empty => new SentryDashboardArgs();
    }

    public sealed class SentryDashboardState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The internal ID for this dashboard.
        /// </summary>
        [Input("internalId")]
        public Input<string>? InternalId { get; set; }

        /// <summary>
        /// The slug of the organization the dashboard belongs to.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// Dashboard title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("widgets")]
        private InputList<Inputs.SentryDashboardWidgetGetArgs>? _widgets;

        /// <summary>
        /// Dashboard widgets.
        /// </summary>
        public InputList<Inputs.SentryDashboardWidgetGetArgs> Widgets
        {
            get => _widgets ?? (_widgets = new InputList<Inputs.SentryDashboardWidgetGetArgs>());
            set => _widgets = value;
        }

        public SentryDashboardState()
        {
        }
        public static new SentryDashboardState Empty => new SentryDashboardState();
    }
}
