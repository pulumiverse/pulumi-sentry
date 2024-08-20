// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage an Opsgenie team integration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumi/sentry";
 * import * as sentry from "@pulumiverse/sentry";
 *
 * const opsgenie = sentry.getSentryOrganizationIntegration({
 *     organization: "my-organization",
 *     providerKey: "opsgenie",
 *     name: "my-pagerduty-organization",
 * });
 * // Associate a Opsgenie service and integration key with a Sentry Opsgenie integration
 * const test = new sentry.SentryIntegrationOpsgenie("test", {
 *     organization: "my-organization",
 *     integrationId: opsgenie.then(opsgenie => opsgenie.id),
 *     team: "my-opsgenie-team",
 *     integrationKey: "c6100908-5c5d-4905-8436-2448fad41bee",
 * });
 * ```
 *
 * ## Import
 *
 * import using the organization slug from the URL:
 *
 * https://sentry.io/api/0/organizations/[org-slug]/integrations/
 *
 * [integration-id] is the top-level `id` of the Opsgenie organization integration
 *
 * [service-id] is the `id` of the service_table record to import under the configData property
 *
 * ```sh
 * $ pulumi import sentry:index/sentryIntegrationOpsgenie:SentryIntegrationOpsgenie default org-slug/integration-id/service-id
 * ```
 */
export class SentryIntegrationOpsgenie extends pulumi.CustomResource {
    /**
     * Get an existing SentryIntegrationOpsgenie resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SentryIntegrationOpsgenieState, opts?: pulumi.CustomResourceOptions): SentryIntegrationOpsgenie {
        return new SentryIntegrationOpsgenie(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sentry:index/sentryIntegrationOpsgenie:SentryIntegrationOpsgenie';

    /**
     * Returns true if the given object is an instance of SentryIntegrationOpsgenie.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SentryIntegrationOpsgenie {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SentryIntegrationOpsgenie.__pulumiType;
    }

    /**
     * The ID of the Opsgenie integration. Source from the URL `https://<organization>.sentry.io/settings/integrations/opsgenie/<integration-id>/` or use the `sentry.getSentryOrganizationIntegration` data source.
     */
    public readonly integrationId!: pulumi.Output<string>;
    /**
     * The integration key of the Opsgenie service.
     */
    public readonly integrationKey!: pulumi.Output<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * The name of the Opsgenie team. In Sentry, this is called Label.
     */
    public readonly team!: pulumi.Output<string>;

    /**
     * Create a SentryIntegrationOpsgenie resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SentryIntegrationOpsgenieArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SentryIntegrationOpsgenieArgs | SentryIntegrationOpsgenieState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SentryIntegrationOpsgenieState | undefined;
            resourceInputs["integrationId"] = state ? state.integrationId : undefined;
            resourceInputs["integrationKey"] = state ? state.integrationKey : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["team"] = state ? state.team : undefined;
        } else {
            const args = argsOrState as SentryIntegrationOpsgenieArgs | undefined;
            if ((!args || args.integrationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationId'");
            }
            if ((!args || args.integrationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationKey'");
            }
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.team === undefined) && !opts.urn) {
                throw new Error("Missing required property 'team'");
            }
            resourceInputs["integrationId"] = args ? args.integrationId : undefined;
            resourceInputs["integrationKey"] = args ? args.integrationKey : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["team"] = args ? args.team : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SentryIntegrationOpsgenie.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SentryIntegrationOpsgenie resources.
 */
export interface SentryIntegrationOpsgenieState {
    /**
     * The ID of the Opsgenie integration. Source from the URL `https://<organization>.sentry.io/settings/integrations/opsgenie/<integration-id>/` or use the `sentry.getSentryOrganizationIntegration` data source.
     */
    integrationId?: pulumi.Input<string>;
    /**
     * The integration key of the Opsgenie service.
     */
    integrationKey?: pulumi.Input<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization?: pulumi.Input<string>;
    /**
     * The name of the Opsgenie team. In Sentry, this is called Label.
     */
    team?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SentryIntegrationOpsgenie resource.
 */
export interface SentryIntegrationOpsgenieArgs {
    /**
     * The ID of the Opsgenie integration. Source from the URL `https://<organization>.sentry.io/settings/integrations/opsgenie/<integration-id>/` or use the `sentry.getSentryOrganizationIntegration` data source.
     */
    integrationId: pulumi.Input<string>;
    /**
     * The integration key of the Opsgenie service.
     */
    integrationKey: pulumi.Input<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * The name of the Opsgenie team. In Sentry, this is called Label.
     */
    team: pulumi.Input<string>;
}
