// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Sentry Issue Alert data source. As the object structure of `conditions`, `filters`, and `actions` are undocumented, a tip is to set up an Issue Alert via the Web UI, and use this data source to copy its object structure to your resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumi/sentry";
 * import * as sentry from "@pulumiverse/sentry";
 *
 * // Retrieve an Issue Alert
 * // URL format: https://sentry.io/organizations/[organization]/alerts/rules/[project]/[internal_id]/details/
 * const original = sentry.getSentryIssueAlert({
 *     organization: "my-organization",
 *     project: "my-project",
 *     internalId: "42",
 * });
 * // Create a copy of an Issue Alert
 * const copy = new sentry.SentryIssueAlert("copy", {
 *     organization: original.then(original => original.organization),
 *     project: original.then(original => original.project),
 *     name: original.then(original => `${original.name}-copy`),
 *     actionMatch: original.then(original => original.actionMatch),
 *     filterMatch: original.then(original => original.filterMatch),
 *     frequency: original.then(original => original.frequency),
 *     conditions: original.then(original => original.conditions),
 *     filters: original.then(original => original.filters),
 *     actions: original.then(original => original.actions),
 * });
 * ```
 */
export function getSentryIssueAlert(args: GetSentryIssueAlertArgs, opts?: pulumi.InvokeOptions): Promise<GetSentryIssueAlertResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sentry:index/getSentryIssueAlert:getSentryIssueAlert", {
        "internalId": args.internalId,
        "organization": args.organization,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getSentryIssueAlert.
 */
export interface GetSentryIssueAlertArgs {
    /**
     * The internal ID for this issue alert.
     */
    internalId: string;
    /**
     * The slug of the organization the issue alert belongs to.
     */
    organization: string;
    /**
     * The slug of the project the issue alert belongs to.
     */
    project: string;
}

/**
 * A collection of values returned by getSentryIssueAlert.
 */
export interface GetSentryIssueAlertResult {
    /**
     * Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
     */
    readonly actionMatch: string;
    /**
     * List of actions.
     */
    readonly actions: {[key: string]: string}[];
    /**
     * List of conditions.
     */
    readonly conditions: {[key: string]: string}[];
    /**
     * Perform issue alert in a specific environment.
     */
    readonly environment: string;
    /**
     * Trigger actions if `all`, `any`, or `none` of the specified filters match.
     */
    readonly filterMatch: string;
    /**
     * List of filters.
     */
    readonly filters: {[key: string]: string}[];
    /**
     * Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
     */
    readonly frequency: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The internal ID for this issue alert.
     */
    readonly internalId: string;
    /**
     * The issue alert name.
     */
    readonly name: string;
    /**
     * The slug of the organization the issue alert belongs to.
     */
    readonly organization: string;
    /**
     * The slug of the project the issue alert belongs to.
     */
    readonly project: string;
}
/**
 * Sentry Issue Alert data source. As the object structure of `conditions`, `filters`, and `actions` are undocumented, a tip is to set up an Issue Alert via the Web UI, and use this data source to copy its object structure to your resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumi/sentry";
 * import * as sentry from "@pulumiverse/sentry";
 *
 * // Retrieve an Issue Alert
 * // URL format: https://sentry.io/organizations/[organization]/alerts/rules/[project]/[internal_id]/details/
 * const original = sentry.getSentryIssueAlert({
 *     organization: "my-organization",
 *     project: "my-project",
 *     internalId: "42",
 * });
 * // Create a copy of an Issue Alert
 * const copy = new sentry.SentryIssueAlert("copy", {
 *     organization: original.then(original => original.organization),
 *     project: original.then(original => original.project),
 *     name: original.then(original => `${original.name}-copy`),
 *     actionMatch: original.then(original => original.actionMatch),
 *     filterMatch: original.then(original => original.filterMatch),
 *     frequency: original.then(original => original.frequency),
 *     conditions: original.then(original => original.conditions),
 *     filters: original.then(original => original.filters),
 *     actions: original.then(original => original.actions),
 * });
 * ```
 */
export function getSentryIssueAlertOutput(args: GetSentryIssueAlertOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSentryIssueAlertResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("sentry:index/getSentryIssueAlert:getSentryIssueAlert", {
        "internalId": args.internalId,
        "organization": args.organization,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getSentryIssueAlert.
 */
export interface GetSentryIssueAlertOutputArgs {
    /**
     * The internal ID for this issue alert.
     */
    internalId: pulumi.Input<string>;
    /**
     * The slug of the organization the issue alert belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * The slug of the project the issue alert belongs to.
     */
    project: pulumi.Input<string>;
}
