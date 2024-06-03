// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Enable spike protection for all projects in an organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumi/sentry";
 * import * as sentry from "@pulumiverse/sentry";
 *
 * // Enable spike protection for several projects in a Sentry organization.
 * const web_app = new sentry.SentryProject("web-app", {
 *     organization: "my-organization",
 *     teams: ["my-first-team"],
 *     slug: "web-app",
 *     platform: "go",
 * });
 * const mobile_app = new sentry.SentryProject("mobile-app", {
 *     organization: "my-organization",
 *     teams: ["my-second-team"],
 *     slug: "mobile-app",
 *     platform: "android",
 * });
 * const mainSentryAllProjectsSpikeProtection = new sentry.SentryAllProjectsSpikeProtection("mainSentryAllProjectsSpikeProtection", {
 *     organization: "my-organization",
 *     projects: [
 *         web_app.id,
 *         mobile_app.id,
 *     ],
 *     enabled: true,
 * });
 * const all = sentry.getSentryAllProjects({
 *     organization: "my-organization",
 * });
 * const mainIndex_sentryAllProjectsSpikeProtectionSentryAllProjectsSpikeProtection = new sentry.SentryAllProjectsSpikeProtection("mainIndex/sentryAllProjectsSpikeProtectionSentryAllProjectsSpikeProtection", {
 *     organization: all.then(all => all.organization),
 *     projects: all.then(all => all.projectSlugs),
 *     enabled: true,
 * });
 * ```
 */
export class SentryAllProjectsSpikeProtection extends pulumi.CustomResource {
    /**
     * Get an existing SentryAllProjectsSpikeProtection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SentryAllProjectsSpikeProtectionState, opts?: pulumi.CustomResourceOptions): SentryAllProjectsSpikeProtection {
        return new SentryAllProjectsSpikeProtection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sentry:index/sentryAllProjectsSpikeProtection:SentryAllProjectsSpikeProtection';

    /**
     * Returns true if the given object is an instance of SentryAllProjectsSpikeProtection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SentryAllProjectsSpikeProtection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SentryAllProjectsSpikeProtection.__pulumiType;
    }

    /**
     * Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off for all projects.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The slug of the organization the resource belongs to.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * The slugs of the projects to enable or disable spike protection for.
     */
    public readonly projects!: pulumi.Output<string[]>;

    /**
     * Create a SentryAllProjectsSpikeProtection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SentryAllProjectsSpikeProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SentryAllProjectsSpikeProtectionArgs | SentryAllProjectsSpikeProtectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SentryAllProjectsSpikeProtectionState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["projects"] = state ? state.projects : undefined;
        } else {
            const args = argsOrState as SentryAllProjectsSpikeProtectionArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.projects === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projects'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["projects"] = args ? args.projects : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SentryAllProjectsSpikeProtection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SentryAllProjectsSpikeProtection resources.
 */
export interface SentryAllProjectsSpikeProtectionState {
    /**
     * Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off for all projects.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization?: pulumi.Input<string>;
    /**
     * The slugs of the projects to enable or disable spike protection for.
     */
    projects?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SentryAllProjectsSpikeProtection resource.
 */
export interface SentryAllProjectsSpikeProtectionArgs {
    /**
     * Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off for all projects.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * The slugs of the projects to enable or disable spike protection for.
     */
    projects: pulumi.Input<pulumi.Input<string>[]>;
}
