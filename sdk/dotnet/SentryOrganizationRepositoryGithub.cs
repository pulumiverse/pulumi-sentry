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
    /// Sentry Github Organization Repository resource.
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
    ///     // Retrieve the Github organization integration
    ///     var github = Sentry.GetSentryOrganizationIntegration.Invoke(new()
    ///     {
    ///         Organization = "my-organization",
    ///         ProviderKey = "github",
    ///         Name = "my-github-organization",
    ///     });
    /// 
    ///     var @this = new Sentry.SentryOrganizationRepositoryGithub("this", new()
    ///     {
    ///         Organization = "my-organization",
    ///         IntegrationId = github.Apply(getSentryOrganizationIntegrationResult =&gt; getSentryOrganizationIntegrationResult.InternalId),
    ///         Identifier = "my-github-organization/my-github-repo",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import using the organization slug from the URL:
    /// 
    /// https://sentry.io/organizations/[org-slug]/
    /// 
    /// [github-org] and [github-repo] are the slugs to your repo
    /// 
    /// ```sh
    /// $ pulumi import sentry:index/sentryOrganizationRepositoryGithub:SentryOrganizationRepositoryGithub this org-slug/github-org/github-repo
    /// ```
    /// </summary>
    [SentryResourceType("sentry:index/sentryOrganizationRepositoryGithub:SentryOrganizationRepositoryGithub")]
    public partial class SentryOrganizationRepositoryGithub : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The repo identifier. For Github it is {github*org}/{github*repo}.
        /// </summary>
        [Output("identifier")]
        public Output<string> Identifier { get; private set; } = null!;

        /// <summary>
        /// The organization integration ID for Github.
        /// </summary>
        [Output("integrationId")]
        public Output<string> IntegrationId { get; private set; } = null!;

        /// <summary>
        /// The internal ID for this organization repository.
        /// </summary>
        [Output("internalId")]
        public Output<string> InternalId { get; private set; } = null!;

        /// <summary>
        /// The slug of the Sentry organization this resource belongs to.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;


        /// <summary>
        /// Create a SentryOrganizationRepositoryGithub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SentryOrganizationRepositoryGithub(string name, SentryOrganizationRepositoryGithubArgs args, CustomResourceOptions? options = null)
            : base("sentry:index/sentryOrganizationRepositoryGithub:SentryOrganizationRepositoryGithub", name, args ?? new SentryOrganizationRepositoryGithubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SentryOrganizationRepositoryGithub(string name, Input<string> id, SentryOrganizationRepositoryGithubState? state = null, CustomResourceOptions? options = null)
            : base("sentry:index/sentryOrganizationRepositoryGithub:SentryOrganizationRepositoryGithub", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SentryOrganizationRepositoryGithub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SentryOrganizationRepositoryGithub Get(string name, Input<string> id, SentryOrganizationRepositoryGithubState? state = null, CustomResourceOptions? options = null)
        {
            return new SentryOrganizationRepositoryGithub(name, id, state, options);
        }
    }

    public sealed class SentryOrganizationRepositoryGithubArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repo identifier. For Github it is {github*org}/{github*repo}.
        /// </summary>
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        /// <summary>
        /// The organization integration ID for Github.
        /// </summary>
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        /// <summary>
        /// The slug of the Sentry organization this resource belongs to.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        public SentryOrganizationRepositoryGithubArgs()
        {
        }
        public static new SentryOrganizationRepositoryGithubArgs Empty => new SentryOrganizationRepositoryGithubArgs();
    }

    public sealed class SentryOrganizationRepositoryGithubState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repo identifier. For Github it is {github*org}/{github*repo}.
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        /// <summary>
        /// The organization integration ID for Github.
        /// </summary>
        [Input("integrationId")]
        public Input<string>? IntegrationId { get; set; }

        /// <summary>
        /// The internal ID for this organization repository.
        /// </summary>
        [Input("internalId")]
        public Input<string>? InternalId { get; set; }

        /// <summary>
        /// The slug of the Sentry organization this resource belongs to.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        public SentryOrganizationRepositoryGithubState()
        {
        }
        public static new SentryOrganizationRepositoryGithubState Empty => new SentryOrganizationRepositoryGithubState();
    }
}
