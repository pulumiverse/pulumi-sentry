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
    /// Sentry Issue Alert resource. Note that there's no public documentation for the values of conditions, filters, and actions. You can either inspect the request payload sent when creating or editing an issue alert on Sentry or inspect [Sentry's rules registry in the source code](https://github.com/getsentry/sentry/tree/master/src/sentry/rules). Since v0.11.2, you should also omit the name property of each condition, filter, and action.
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
    ///     // Retrieve a Slack integration
    ///     var slack = Sentry.GetSentryOrganizationIntegration.Invoke(new()
    ///     {
    ///         Organization = test.Organization,
    ///         ProviderKey = "slack",
    ///         Name = "Slack Workspace",
    ///     });
    /// 
    ///     var main = new Sentry.SentryIssueAlert("main", new()
    ///     {
    ///         Organization = mainSentryProject.Organization,
    ///         Project = mainSentryProject.Id,
    ///         Name = "My issue alert",
    ///         ActionMatch = "any",
    ///         FilterMatch = "any",
    ///         Frequency = 30,
    ///         Conditions = new[]
    ///         {
    ///             
    ///             {
    ///                 { "id", "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.conditions.regression_event.RegressionEventCondition" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.conditions.event_frequency.EventFrequencyCondition" },
    ///                 { "value", "100" },
    ///                 { "comparisonType", "count" },
    ///                 { "interval", "1h" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.conditions.event_frequency.EventUniqueUserFrequencyCondition" },
    ///                 { "value", "100" },
    ///                 { "comparisonType", "count" },
    ///                 { "interval", "1h" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.conditions.event_frequency.EventFrequencyPercentCondition" },
    ///                 { "value", "50.0" },
    ///                 { "comparisonType", "count" },
    ///                 { "interval", "1h" },
    ///             },
    ///         },
    ///         Filters = new[]
    ///         {
    ///             
    ///             {
    ///                 { "id", "sentry.rules.filters.age_comparison.AgeComparisonFilter" },
    ///                 { "value", "10" },
    ///                 { "time", "minute" },
    ///                 { "comparison_type", "older" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.filters.issue_occurrences.IssueOccurrencesFilter" },
    ///                 { "value", "10" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.filters.assigned_to.AssignedToFilter" },
    ///                 { "targetType", "Team" },
    ///                 { "targetIdentifier", mainSentryTeam.TeamId },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.filters.latest_release.LatestReleaseFilter" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.filters.event_attribute.EventAttributeFilter" },
    ///                 { "attribute", "message" },
    ///                 { "match", "co" },
    ///                 { "value", "test" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.filters.tagged_event.TaggedEventFilter" },
    ///                 { "key", "test" },
    ///                 { "match", "co" },
    ///                 { "value", "test" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.filters.level.LevelFilter" },
    ///                 { "match", "eq" },
    ///                 { "level", "50" },
    ///             },
    ///         },
    ///         Actions = new[]
    ///         {
    ///             
    ///             {
    ///                 { "id", "sentry.mail.actions.NotifyEmailAction" },
    ///                 { "targetType", "IssueOwners" },
    ///                 { "targetIdentifier", "" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.mail.actions.NotifyEmailAction" },
    ///                 { "targetType", "Team" },
    ///                 { "targetIdentifier", mainSentryTeam.TeamId },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.rules.actions.notify_event.NotifyEventAction" },
    ///             },
    ///             
    ///             {
    ///                 { "id", "sentry.integrations.slack.notify_action.SlackNotifyServiceAction" },
    ///                 { "channel", "#general" },
    ///                 { "workspace", slack.Apply(getSentryOrganizationIntegrationResult =&gt; getSentryOrganizationIntegrationResult.InternalId) },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import using the organization, project slugs and rule id from the URL:
    /// 
    /// https://sentry.io/organizations/[org-slug]/alerts/rules/[project-slug]/[rule-id]/details/
    /// 
    /// ```sh
    /// $ pulumi import sentry:index/sentryIssueAlert:SentryIssueAlert default org-slug/project-slug/rule-id
    /// ```
    /// </summary>
    [SentryResourceType("sentry:index/sentryIssueAlert:SentryIssueAlert")]
    public partial class SentryIssueAlert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
        /// </summary>
        [Output("actionMatch")]
        public Output<string> ActionMatch { get; private set; } = null!;

        /// <summary>
        /// List of actions.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<ImmutableDictionary<string, string>>> Actions { get; private set; } = null!;

        /// <summary>
        /// List of conditions.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<ImmutableDictionary<string, string>>> Conditions { get; private set; } = null!;

        /// <summary>
        /// Perform issue alert in a specific environment.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// Trigger actions if `all`, `any`, or `none` of the specified filters match.
        /// </summary>
        [Output("filterMatch")]
        public Output<string> FilterMatch { get; private set; } = null!;

        /// <summary>
        /// List of filters.
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<ImmutableDictionary<string, string>>> Filters { get; private set; } = null!;

        /// <summary>
        /// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
        /// </summary>
        [Output("frequency")]
        public Output<int> Frequency { get; private set; } = null!;

        /// <summary>
        /// The internal ID for this issue alert.
        /// </summary>
        [Output("internalId")]
        public Output<string> InternalId { get; private set; } = null!;

        /// <summary>
        /// The issue alert name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The slug of the organization the issue alert belongs to.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        /// <summary>
        /// The slug of the project to create the issue alert for.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Use `project` (singular) instead.
        /// </summary>
        [Output("projects")]
        public Output<ImmutableArray<string>> Projects { get; private set; } = null!;


        /// <summary>
        /// Create a SentryIssueAlert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SentryIssueAlert(string name, SentryIssueAlertArgs args, CustomResourceOptions? options = null)
            : base("sentry:index/sentryIssueAlert:SentryIssueAlert", name, args ?? new SentryIssueAlertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SentryIssueAlert(string name, Input<string> id, SentryIssueAlertState? state = null, CustomResourceOptions? options = null)
            : base("sentry:index/sentryIssueAlert:SentryIssueAlert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SentryIssueAlert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SentryIssueAlert Get(string name, Input<string> id, SentryIssueAlertState? state = null, CustomResourceOptions? options = null)
        {
            return new SentryIssueAlert(name, id, state, options);
        }
    }

    public sealed class SentryIssueAlertArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
        /// </summary>
        [Input("actionMatch", required: true)]
        public Input<string> ActionMatch { get; set; } = null!;

        [Input("actions", required: true)]
        private InputList<ImmutableDictionary<string, string>>? _actions;

        /// <summary>
        /// List of actions.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Actions
        {
            get => _actions ?? (_actions = new InputList<ImmutableDictionary<string, string>>());
            set => _actions = value;
        }

        [Input("conditions", required: true)]
        private InputList<ImmutableDictionary<string, string>>? _conditions;

        /// <summary>
        /// List of conditions.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<ImmutableDictionary<string, string>>());
            set => _conditions = value;
        }

        /// <summary>
        /// Perform issue alert in a specific environment.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// Trigger actions if `all`, `any`, or `none` of the specified filters match.
        /// </summary>
        [Input("filterMatch", required: true)]
        public Input<string> FilterMatch { get; set; } = null!;

        [Input("filters")]
        private InputList<ImmutableDictionary<string, string>>? _filters;

        /// <summary>
        /// List of filters.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Filters
        {
            get => _filters ?? (_filters = new InputList<ImmutableDictionary<string, string>>());
            set => _filters = value;
        }

        /// <summary>
        /// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
        /// </summary>
        [Input("frequency", required: true)]
        public Input<int> Frequency { get; set; } = null!;

        /// <summary>
        /// The issue alert name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The slug of the organization the issue alert belongs to.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// The slug of the project to create the issue alert for.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public SentryIssueAlertArgs()
        {
        }
        public static new SentryIssueAlertArgs Empty => new SentryIssueAlertArgs();
    }

    public sealed class SentryIssueAlertState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
        /// </summary>
        [Input("actionMatch")]
        public Input<string>? ActionMatch { get; set; }

        [Input("actions")]
        private InputList<ImmutableDictionary<string, string>>? _actions;

        /// <summary>
        /// List of actions.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Actions
        {
            get => _actions ?? (_actions = new InputList<ImmutableDictionary<string, string>>());
            set => _actions = value;
        }

        [Input("conditions")]
        private InputList<ImmutableDictionary<string, string>>? _conditions;

        /// <summary>
        /// List of conditions.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<ImmutableDictionary<string, string>>());
            set => _conditions = value;
        }

        /// <summary>
        /// Perform issue alert in a specific environment.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// Trigger actions if `all`, `any`, or `none` of the specified filters match.
        /// </summary>
        [Input("filterMatch")]
        public Input<string>? FilterMatch { get; set; }

        [Input("filters")]
        private InputList<ImmutableDictionary<string, string>>? _filters;

        /// <summary>
        /// List of filters.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Filters
        {
            get => _filters ?? (_filters = new InputList<ImmutableDictionary<string, string>>());
            set => _filters = value;
        }

        /// <summary>
        /// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
        /// </summary>
        [Input("frequency")]
        public Input<int>? Frequency { get; set; }

        /// <summary>
        /// The internal ID for this issue alert.
        /// </summary>
        [Input("internalId")]
        public Input<string>? InternalId { get; set; }

        /// <summary>
        /// The issue alert name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The slug of the organization the issue alert belongs to.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// The slug of the project to create the issue alert for.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("projects")]
        private InputList<string>? _projects;

        /// <summary>
        /// Use `project` (singular) instead.
        /// </summary>
        [Obsolete(@"Use `project` (singular) instead.")]
        public InputList<string> Projects
        {
            get => _projects ?? (_projects = new InputList<string>());
            set => _projects = value;
        }

        public SentryIssueAlertState()
        {
        }
        public static new SentryIssueAlertState Empty => new SentryIssueAlertState();
    }
}
