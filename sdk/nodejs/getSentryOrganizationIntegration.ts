// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Sentry Organization Integration data source.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumi/sentry";
 *
 * const github = sentry.getSentryOrganizationIntegration({
 *     name: "my-github-organization",
 *     organization: "my-organization",
 *     providerKey: "github",
 * });
 * const slack = sentry.getSentryOrganizationIntegration({
 *     name: "Slack Workspace",
 *     organization: "my-organization",
 *     providerKey: "slack",
 * });
 * ```
 */
export function getSentryOrganizationIntegration(args: GetSentryOrganizationIntegrationArgs, opts?: pulumi.InvokeOptions): Promise<GetSentryOrganizationIntegrationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sentry:index/getSentryOrganizationIntegration:getSentryOrganizationIntegration", {
        "name": args.name,
        "organization": args.organization,
        "providerKey": args.providerKey,
    }, opts);
}

/**
 * A collection of arguments for invoking getSentryOrganizationIntegration.
 */
export interface GetSentryOrganizationIntegrationArgs {
    /**
     * The name of the organization integration.
     */
    name: string;
    /**
     * The slug of the organization the integration belongs to.
     */
    organization: string;
    /**
     * The key of the organization integration provider.
     */
    providerKey: string;
}

/**
 * A collection of values returned by getSentryOrganizationIntegration.
 */
export interface GetSentryOrganizationIntegrationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The internal ID for this organization integration.
     */
    readonly internalId: string;
    /**
     * The name of the organization integration.
     */
    readonly name: string;
    /**
     * The slug of the organization the integration belongs to.
     */
    readonly organization: string;
    /**
     * The key of the organization integration provider.
     */
    readonly providerKey: string;
}
/**
 * Sentry Organization Integration data source.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumi/sentry";
 *
 * const github = sentry.getSentryOrganizationIntegration({
 *     name: "my-github-organization",
 *     organization: "my-organization",
 *     providerKey: "github",
 * });
 * const slack = sentry.getSentryOrganizationIntegration({
 *     name: "Slack Workspace",
 *     organization: "my-organization",
 *     providerKey: "slack",
 * });
 * ```
 */
export function getSentryOrganizationIntegrationOutput(args: GetSentryOrganizationIntegrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSentryOrganizationIntegrationResult> {
    return pulumi.output(args).apply((a: any) => getSentryOrganizationIntegration(a, opts))
}

/**
 * A collection of arguments for invoking getSentryOrganizationIntegration.
 */
export interface GetSentryOrganizationIntegrationOutputArgs {
    /**
     * The name of the organization integration.
     */
    name: pulumi.Input<string>;
    /**
     * The slug of the organization the integration belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * The key of the organization integration provider.
     */
    providerKey: pulumi.Input<string>;
}
