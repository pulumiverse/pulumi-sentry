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
    public static class GetSentryTeam
    {
        /// <summary>
        /// Sentry Team data source.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Sentry = Pulumi.Sentry;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Retrieve a team
        ///     var @default = Sentry.GetSentryTeam.Invoke(new()
        ///     {
        ///         Organization = "my-organization",
        ///         Slug = "my-team",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSentryTeamResult> InvokeAsync(GetSentryTeamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSentryTeamResult>("sentry:index/getSentryTeam:getSentryTeam", args ?? new GetSentryTeamArgs(), options.WithDefaults());

        /// <summary>
        /// Sentry Team data source.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Sentry = Pulumi.Sentry;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     // Retrieve a team
        ///     var @default = Sentry.GetSentryTeam.Invoke(new()
        ///     {
        ///         Organization = "my-organization",
        ///         Slug = "my-team",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSentryTeamResult> Invoke(GetSentryTeamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSentryTeamResult>("sentry:index/getSentryTeam:getSentryTeam", args ?? new GetSentryTeamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSentryTeamArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The slug of the organization the team should be created for.
        /// </summary>
        [Input("organization", required: true)]
        public string Organization { get; set; } = null!;

        /// <summary>
        /// The unique URL slug for this team.
        /// </summary>
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        public GetSentryTeamArgs()
        {
        }
        public static new GetSentryTeamArgs Empty => new GetSentryTeamArgs();
    }

    public sealed class GetSentryTeamInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The slug of the organization the team should be created for.
        /// </summary>
        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        /// <summary>
        /// The unique URL slug for this team.
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        public GetSentryTeamInvokeArgs()
        {
        }
        public static new GetSentryTeamInvokeArgs Empty => new GetSentryTeamInvokeArgs();
    }


    [OutputType]
    public sealed class GetSentryTeamResult
    {
        public readonly bool HasAccess;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The internal ID for this team.
        /// </summary>
        public readonly string InternalId;
        public readonly bool IsMember;
        public readonly bool IsPending;
        /// <summary>
        /// The human readable name for this organization.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The slug of the organization the team should be created for.
        /// </summary>
        public readonly string Organization;
        /// <summary>
        /// The unique URL slug for this team.
        /// </summary>
        public readonly string Slug;

        [OutputConstructor]
        private GetSentryTeamResult(
            bool hasAccess,

            string id,

            string internalId,

            bool isMember,

            bool isPending,

            string name,

            string organization,

            string slug)
        {
            HasAccess = hasAccess;
            Id = id;
            InternalId = internalId;
            IsMember = isMember;
            IsPending = isPending;
            Name = name;
            Organization = organization;
            Slug = slug;
        }
    }
}
