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
    /// Sentry Project resource.
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
    ///     // Create a project
    ///     var @default = new Sentry.SentryProject("default", new()
    ///     {
    ///         Organization = "my-organization",
    ///         Teams = new[]
    ///         {
    ///             "my-first-team",
    ///             "my-second-team",
    ///         },
    ///         Name = "Web App",
    ///         Slug = "web-app",
    ///         Platform = "javascript",
    ///         ResolveAge = 720,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import using the organization and team slugs from the URL:
    /// 
    /// https://sentry.io/settings/[org-slug]/projects/[project-slug]/
    /// 
    /// ```sh
    /// $ pulumi import sentry:index/sentryProject:SentryProject default org-slug/project-slug
    /// ```
    /// </summary>
    [SentryResourceType("sentry:index/sentryProject:SentryProject")]
    public partial class SentryProject : global::Pulumi.CustomResource
    {
        [Output("color")]
        public Output<string> Color { get; private set; } = null!;

        /// <summary>
        /// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
        /// </summary>
        [Output("digestsMaxDelay")]
        public Output<int> DigestsMaxDelay { get; private set; } = null!;

        /// <summary>
        /// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
        /// </summary>
        [Output("digestsMinDelay")]
        public Output<int> DigestsMinDelay { get; private set; } = null!;

        [Output("features")]
        public Output<ImmutableArray<string>> Features { get; private set; } = null!;

        /// <summary>
        /// The internal ID for this project.
        /// </summary>
        [Output("internalId")]
        public Output<string> InternalId { get; private set; } = null!;

        [Output("isBookmarked")]
        public Output<bool> IsBookmarked { get; private set; } = null!;

        [Output("isPublic")]
        public Output<bool> IsPublic { get; private set; } = null!;

        /// <summary>
        /// The name for the project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The slug of the organization the project belongs to.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        /// <summary>
        /// The optional platform for this project.
        /// </summary>
        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        /// <summary>
        /// Use `internal_id` instead.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Hours in which an issue is automatically resolve if not seen after this amount of time.
        /// </summary>
        [Output("resolveAge")]
        public Output<int> ResolveAge { get; private set; } = null!;

        /// <summary>
        /// The optional slug for this project.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The slug of the team to create the project for. **Deprecated** Use `teams` instead.
        /// </summary>
        [Output("team")]
        public Output<string?> Team { get; private set; } = null!;

        /// <summary>
        /// The slugs of the teams to create the project for.
        /// </summary>
        [Output("teams")]
        public Output<ImmutableArray<string>> Teams { get; private set; } = null!;


        /// <summary>
        /// Create a SentryProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SentryProject(string name, SentryProjectArgs args, CustomResourceOptions? options = null)
            : base("sentry:index/sentryProject:SentryProject", name, args ?? new SentryProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SentryProject(string name, Input<string> id, SentryProjectState? state = null, CustomResourceOptions? options = null)
            : base("sentry:index/sentryProject:SentryProject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SentryProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SentryProject Get(string name, Input<string> id, SentryProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new SentryProject(name, id, state, options);
        }
    }

    public sealed class SentryProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
        /// </summary>
        [Input("digestsMaxDelay")]
        public Input<int>? DigestsMaxDelay { get; set; }

        /// <summary>
        /// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
        /// </summary>
        [Input("digestsMinDelay")]
        public Input<int>? DigestsMinDelay { get; set; }

        /// <summary>
        /// The name for the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The slug of the organization the project belongs to.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// The optional platform for this project.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// Hours in which an issue is automatically resolve if not seen after this amount of time.
        /// </summary>
        [Input("resolveAge")]
        public Input<int>? ResolveAge { get; set; }

        /// <summary>
        /// The optional slug for this project.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// The slug of the team to create the project for. **Deprecated** Use `teams` instead.
        /// </summary>
        [Input("team")]
        public Input<string>? Team { get; set; }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// The slugs of the teams to create the project for.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        public SentryProjectArgs()
        {
        }
        public static new SentryProjectArgs Empty => new SentryProjectArgs();
    }

    public sealed class SentryProjectState : global::Pulumi.ResourceArgs
    {
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
        /// </summary>
        [Input("digestsMaxDelay")]
        public Input<int>? DigestsMaxDelay { get; set; }

        /// <summary>
        /// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
        /// </summary>
        [Input("digestsMinDelay")]
        public Input<int>? DigestsMinDelay { get; set; }

        [Input("features")]
        private InputList<string>? _features;
        public InputList<string> Features
        {
            get => _features ?? (_features = new InputList<string>());
            set => _features = value;
        }

        /// <summary>
        /// The internal ID for this project.
        /// </summary>
        [Input("internalId")]
        public Input<string>? InternalId { get; set; }

        [Input("isBookmarked")]
        public Input<bool>? IsBookmarked { get; set; }

        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// The name for the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The slug of the organization the project belongs to.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// The optional platform for this project.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// Use `internal_id` instead.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Hours in which an issue is automatically resolve if not seen after this amount of time.
        /// </summary>
        [Input("resolveAge")]
        public Input<int>? ResolveAge { get; set; }

        /// <summary>
        /// The optional slug for this project.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The slug of the team to create the project for. **Deprecated** Use `teams` instead.
        /// </summary>
        [Input("team")]
        public Input<string>? Team { get; set; }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// The slugs of the teams to create the project for.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        public SentryProjectState()
        {
        }
        public static new SentryProjectState Empty => new SentryProjectState();
    }
}
