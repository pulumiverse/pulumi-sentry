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
    /// Sentry Project Spike Protection resource. This resource is used to create and manage spike protection for a project.
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
    ///     var defaultSentryProject = new Sentry.SentryProject("defaultSentryProject", new()
    ///     {
    ///         Organization = "my-organization",
    ///         Teams = new[]
    ///         {
    ///             "my-first-team",
    ///             "my-second-team",
    ///         },
    ///         Platform = "javascript",
    ///     });
    /// 
    ///     // Enable spike protection for the project
    ///     var defaultSentryProjectSpikeProtection = new Sentry.SentryProjectSpikeProtection("defaultSentryProjectSpikeProtection", new()
    ///     {
    ///         Organization = defaultSentryProject.Organization,
    ///         Project = defaultSentryProject.Id,
    ///         Enabled = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SentryResourceType("sentry:index/sentryProjectSpikeProtection:SentryProjectSpikeProtection")]
    public partial class SentryProjectSpikeProtection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The slug of the organization the project belongs to.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        /// <summary>
        /// The slug of the project to enable or disable spike protection for.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a SentryProjectSpikeProtection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SentryProjectSpikeProtection(string name, SentryProjectSpikeProtectionArgs args, CustomResourceOptions? options = null)
            : base("sentry:index/sentryProjectSpikeProtection:SentryProjectSpikeProtection", name, args ?? new SentryProjectSpikeProtectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SentryProjectSpikeProtection(string name, Input<string> id, SentryProjectSpikeProtectionState? state = null, CustomResourceOptions? options = null)
            : base("sentry:index/sentryProjectSpikeProtection:SentryProjectSpikeProtection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SentryProjectSpikeProtection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SentryProjectSpikeProtection Get(string name, Input<string> id, SentryProjectSpikeProtectionState? state = null, CustomResourceOptions? options = null)
        {
            return new SentryProjectSpikeProtection(name, id, state, options);
        }
    }

    public sealed class SentryProjectSpikeProtectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The slug of the organization the project belongs to.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// The slug of the project to enable or disable spike protection for.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public SentryProjectSpikeProtectionArgs()
        {
        }
        public static new SentryProjectSpikeProtectionArgs Empty => new SentryProjectSpikeProtectionArgs();
    }

    public sealed class SentryProjectSpikeProtectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The slug of the organization the project belongs to.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// The slug of the project to enable or disable spike protection for.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public SentryProjectSpikeProtectionState()
        {
        }
        public static new SentryProjectSpikeProtectionState Empty => new SentryProjectSpikeProtectionState();
    }
}
