// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage a PagerDuty service integration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumi/sentry";
 * import * as sentry from "@pulumiverse/sentry";
 *
 * const pagerduty = sentry.getSentryOrganizationIntegration({
 *     organization: "my-organization",
 *     providerKey: "pagerduty",
 *     name: "my-pagerduty-organization",
 * });
 * // Associate a PagerDuty service and integration key with a Sentry PagerDuty integration
 * const test = new sentry.SentryIntegrationPagerDuty("test", {
 *     organization: "my-organization",
 *     integrationId: pagerduty.then(pagerduty => pagerduty.id),
 *     service: "my-pagerduty-service",
 *     integrationKey: "my-pagerduty-integration-key",
 * });
 * ```
 *
 * ## Import
 *
 * import using the organization slug from the URL:
 *
 * https://sentry.io/api/0/organizations/[org-slug]/integrations/
 *
 * [integration-id] is the top-level `id` of the PagerDuty organization integration
 *
 * [service-id] is the `id` of the service_table record to import under the configData property
 *
 * ```sh
 * $ pulumi import sentry:index/sentryIntegrationPagerDuty:SentryIntegrationPagerDuty default org-slug/integration-id/service-id
 * ```
 */
export class SentryIntegrationPagerDuty extends pulumi.CustomResource {
    /**
     * Get an existing SentryIntegrationPagerDuty resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SentryIntegrationPagerDutyState, opts?: pulumi.CustomResourceOptions): SentryIntegrationPagerDuty {
        return new SentryIntegrationPagerDuty(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sentry:index/sentryIntegrationPagerDuty:SentryIntegrationPagerDuty';

    /**
     * Returns true if the given object is an instance of SentryIntegrationPagerDuty.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SentryIntegrationPagerDuty {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SentryIntegrationPagerDuty.__pulumiType;
    }

    /**
     * The ID of the PagerDuty integration. Source from the URL `https://<organization>.sentry.io/settings/integrations/pagerduty/<integration-id>/` or use the `sentry.getSentryOrganizationIntegration` data source.
     */
    public readonly integrationId!: pulumi.Output<string>;
    /**
     * The integration key of the PagerDuty service.
     */
    public readonly integrationKey!: pulumi.Output<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * The name of the PagerDuty service.
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a SentryIntegrationPagerDuty resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SentryIntegrationPagerDutyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SentryIntegrationPagerDutyArgs | SentryIntegrationPagerDutyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SentryIntegrationPagerDutyState | undefined;
            resourceInputs["integrationId"] = state ? state.integrationId : undefined;
            resourceInputs["integrationKey"] = state ? state.integrationKey : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as SentryIntegrationPagerDutyArgs | undefined;
            if ((!args || args.integrationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationId'");
            }
            if ((!args || args.integrationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationKey'");
            }
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            resourceInputs["integrationId"] = args ? args.integrationId : undefined;
            resourceInputs["integrationKey"] = args ? args.integrationKey : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SentryIntegrationPagerDuty.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SentryIntegrationPagerDuty resources.
 */
export interface SentryIntegrationPagerDutyState {
    /**
     * The ID of the PagerDuty integration. Source from the URL `https://<organization>.sentry.io/settings/integrations/pagerduty/<integration-id>/` or use the `sentry.getSentryOrganizationIntegration` data source.
     */
    integrationId?: pulumi.Input<string>;
    /**
     * The integration key of the PagerDuty service.
     */
    integrationKey?: pulumi.Input<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization?: pulumi.Input<string>;
    /**
     * The name of the PagerDuty service.
     */
    service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SentryIntegrationPagerDuty resource.
 */
export interface SentryIntegrationPagerDutyArgs {
    /**
     * The ID of the PagerDuty integration. Source from the URL `https://<organization>.sentry.io/settings/integrations/pagerduty/<integration-id>/` or use the `sentry.getSentryOrganizationIntegration` data source.
     */
    integrationId: pulumi.Input<string>;
    /**
     * The integration key of the PagerDuty service.
     */
    integrationKey: pulumi.Input<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * The name of the PagerDuty service.
     */
    service: pulumi.Input<string>;
}
