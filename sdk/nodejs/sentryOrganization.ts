// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Sentry Organization resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumiverse/sentry";
 *
 * // Create an organization
 * const _default = new sentry.SentryOrganization("default", {
 *     name: "My Organization",
 *     slug: "my-organization",
 *     agreeTerms: true,
 * });
 * ```
 *
 * ## Import
 *
 * import using the organization slug from the URL:
 *
 * https://sentry.io/organizations/[org-slug]/issues/
 *
 * ```sh
 * $ pulumi import sentry:index/sentryOrganization:SentryOrganization default org-slug
 * ```
 */
export class SentryOrganization extends pulumi.CustomResource {
    /**
     * Get an existing SentryOrganization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SentryOrganizationState, opts?: pulumi.CustomResourceOptions): SentryOrganization {
        return new SentryOrganization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sentry:index/sentryOrganization:SentryOrganization';

    /**
     * Returns true if the given object is an instance of SentryOrganization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SentryOrganization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SentryOrganization.__pulumiType;
    }

    /**
     * You agree to the applicable terms of service and privacy policy.
     */
    public readonly agreeTerms!: pulumi.Output<boolean>;
    /**
     * The internal ID for this organization.
     */
    public /*out*/ readonly internalId!: pulumi.Output<string>;
    /**
     * The human readable name for the organization.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The unique URL slug for this organization.
     */
    public readonly slug!: pulumi.Output<string>;

    /**
     * Create a SentryOrganization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SentryOrganizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SentryOrganizationArgs | SentryOrganizationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SentryOrganizationState | undefined;
            resourceInputs["agreeTerms"] = state ? state.agreeTerms : undefined;
            resourceInputs["internalId"] = state ? state.internalId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as SentryOrganizationArgs | undefined;
            if ((!args || args.agreeTerms === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agreeTerms'");
            }
            resourceInputs["agreeTerms"] = args ? args.agreeTerms : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["slug"] = args ? args.slug : undefined;
            resourceInputs["internalId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SentryOrganization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SentryOrganization resources.
 */
export interface SentryOrganizationState {
    /**
     * You agree to the applicable terms of service and privacy policy.
     */
    agreeTerms?: pulumi.Input<boolean>;
    /**
     * The internal ID for this organization.
     */
    internalId?: pulumi.Input<string>;
    /**
     * The human readable name for the organization.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique URL slug for this organization.
     */
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SentryOrganization resource.
 */
export interface SentryOrganizationArgs {
    /**
     * You agree to the applicable terms of service and privacy policy.
     */
    agreeTerms: pulumi.Input<boolean>;
    /**
     * The human readable name for the organization.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique URL slug for this organization.
     */
    slug?: pulumi.Input<string>;
}
