// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Return a client key bound to a project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumiverse/sentry";
 *
 * // Create a key
 * const _default = new sentry.SentryKey("default", {
 *     organization: "my-organization",
 *     project: "web-app",
 * });
 * ```
 *
 * ## Import
 *
 * import using the organization, project slugs and key id from the URL:
 *
 * https://sentry.io/settings/[org-slug]/projects/[project-slug]/keys/[key-id]/
 *
 * ```sh
 * $ pulumi import sentry:index/sentryKey:SentryKey default org-slug/project-slug/key-id
 * ```
 */
export class SentryKey extends pulumi.CustomResource {
    /**
     * Get an existing SentryKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SentryKeyState, opts?: pulumi.CustomResourceOptions): SentryKey {
        return new SentryKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sentry:index/sentryKey:SentryKey';

    /**
     * Returns true if the given object is an instance of SentryKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SentryKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SentryKey.__pulumiType;
    }

    /**
     * Security header endpoint for features like CSP and Expect-CT reports.
     */
    public /*out*/ readonly dsnCsp!: pulumi.Output<string>;
    /**
     * The DSN tells the SDK where to send the events to.
     */
    public /*out*/ readonly dsnPublic!: pulumi.Output<string>;
    /**
     * Deprecated DSN includes a secret which is no longer required by newer SDK versions. If you are unsure which to use, follow installation instructions for your language.
     */
    public /*out*/ readonly dsnSecret!: pulumi.Output<string>;
    /**
     * The name of the client key.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * The slug of the project the resource belongs to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The ID of the project that the key belongs to.
     */
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    /**
     * The public key.
     */
    public /*out*/ readonly public!: pulumi.Output<string>;
    /**
     * Number of events that can be reported within the rate limit window.
     */
    public readonly rateLimitCount!: pulumi.Output<number | undefined>;
    /**
     * Length of time that will be considered when checking the rate limit.
     */
    public readonly rateLimitWindow!: pulumi.Output<number | undefined>;
    /**
     * The secret key.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;

    /**
     * Create a SentryKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SentryKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SentryKeyArgs | SentryKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SentryKeyState | undefined;
            resourceInputs["dsnCsp"] = state ? state.dsnCsp : undefined;
            resourceInputs["dsnPublic"] = state ? state.dsnPublic : undefined;
            resourceInputs["dsnSecret"] = state ? state.dsnSecret : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["public"] = state ? state.public : undefined;
            resourceInputs["rateLimitCount"] = state ? state.rateLimitCount : undefined;
            resourceInputs["rateLimitWindow"] = state ? state.rateLimitWindow : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
        } else {
            const args = argsOrState as SentryKeyArgs | undefined;
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["rateLimitCount"] = args ? args.rateLimitCount : undefined;
            resourceInputs["rateLimitWindow"] = args ? args.rateLimitWindow : undefined;
            resourceInputs["dsnCsp"] = undefined /*out*/;
            resourceInputs["dsnPublic"] = undefined /*out*/;
            resourceInputs["dsnSecret"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["public"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["dsnSecret", "secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SentryKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SentryKey resources.
 */
export interface SentryKeyState {
    /**
     * Security header endpoint for features like CSP and Expect-CT reports.
     */
    dsnCsp?: pulumi.Input<string>;
    /**
     * The DSN tells the SDK where to send the events to.
     */
    dsnPublic?: pulumi.Input<string>;
    /**
     * Deprecated DSN includes a secret which is no longer required by newer SDK versions. If you are unsure which to use, follow installation instructions for your language.
     */
    dsnSecret?: pulumi.Input<string>;
    /**
     * The name of the client key.
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization?: pulumi.Input<string>;
    /**
     * The slug of the project the resource belongs to.
     */
    project?: pulumi.Input<string>;
    /**
     * The ID of the project that the key belongs to.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The public key.
     */
    public?: pulumi.Input<string>;
    /**
     * Number of events that can be reported within the rate limit window.
     */
    rateLimitCount?: pulumi.Input<number>;
    /**
     * Length of time that will be considered when checking the rate limit.
     */
    rateLimitWindow?: pulumi.Input<number>;
    /**
     * The secret key.
     */
    secret?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SentryKey resource.
 */
export interface SentryKeyArgs {
    /**
     * The name of the client key.
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * The slug of the project the resource belongs to.
     */
    project: pulumi.Input<string>;
    /**
     * Number of events that can be reported within the rate limit window.
     */
    rateLimitCount?: pulumi.Input<number>;
    /**
     * Length of time that will be considered when checking the rate limit.
     */
    rateLimitWindow?: pulumi.Input<number>;
}
