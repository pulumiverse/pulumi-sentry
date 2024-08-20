// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * import using the organization, project slugs and rule id from the URL:
 *
 * https://sentry.io/organizations/[org-slug]/alerts/rules/[project-slug]/[rule-id]/details/
 *
 * ```sh
 * $ pulumi import sentry:index/sentryIssueAlert:SentryIssueAlert default org-slug/project-slug/rule-id
 * ```
 */
export class SentryIssueAlert extends pulumi.CustomResource {
    /**
     * Get an existing SentryIssueAlert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SentryIssueAlertState, opts?: pulumi.CustomResourceOptions): SentryIssueAlert {
        return new SentryIssueAlert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sentry:index/sentryIssueAlert:SentryIssueAlert';

    /**
     * Returns true if the given object is an instance of SentryIssueAlert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SentryIssueAlert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SentryIssueAlert.__pulumiType;
    }

    /**
     * Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
     */
    public readonly actionMatch!: pulumi.Output<string>;
    /**
     * List of actions. In JSON string format.
     */
    public readonly actions!: pulumi.Output<string>;
    /**
     * List of conditions. In JSON string format.
     */
    public readonly conditions!: pulumi.Output<string>;
    /**
     * Perform issue alert in a specific environment.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * A string determining which filters need to be true before any actions take place. Required when a value is provided for `filters`.
     */
    public readonly filterMatch!: pulumi.Output<string | undefined>;
    /**
     * A list of filters that determine if a rule fires after the necessary conditions have been met. In JSON string format.
     */
    public readonly filters!: pulumi.Output<string | undefined>;
    /**
     * Perform actions at most once every `X` minutes for this issue.
     */
    public readonly frequency!: pulumi.Output<number>;
    /**
     * The issue alert name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * The ID of the team or user that owns the rule.
     */
    public readonly owner!: pulumi.Output<string | undefined>;
    /**
     * The slug of the project the resource belongs to.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a SentryIssueAlert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SentryIssueAlertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SentryIssueAlertArgs | SentryIssueAlertState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SentryIssueAlertState | undefined;
            resourceInputs["actionMatch"] = state ? state.actionMatch : undefined;
            resourceInputs["actions"] = state ? state.actions : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["filterMatch"] = state ? state.filterMatch : undefined;
            resourceInputs["filters"] = state ? state.filters : undefined;
            resourceInputs["frequency"] = state ? state.frequency : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as SentryIssueAlertArgs | undefined;
            if ((!args || args.actionMatch === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actionMatch'");
            }
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.conditions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'conditions'");
            }
            if ((!args || args.frequency === undefined) && !opts.urn) {
                throw new Error("Missing required property 'frequency'");
            }
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["actionMatch"] = args ? args.actionMatch : undefined;
            resourceInputs["actions"] = args ? args.actions : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["filterMatch"] = args ? args.filterMatch : undefined;
            resourceInputs["filters"] = args ? args.filters : undefined;
            resourceInputs["frequency"] = args ? args.frequency : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SentryIssueAlert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SentryIssueAlert resources.
 */
export interface SentryIssueAlertState {
    /**
     * Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
     */
    actionMatch?: pulumi.Input<string>;
    /**
     * List of actions. In JSON string format.
     */
    actions?: pulumi.Input<string>;
    /**
     * List of conditions. In JSON string format.
     */
    conditions?: pulumi.Input<string>;
    /**
     * Perform issue alert in a specific environment.
     */
    environment?: pulumi.Input<string>;
    /**
     * A string determining which filters need to be true before any actions take place. Required when a value is provided for `filters`.
     */
    filterMatch?: pulumi.Input<string>;
    /**
     * A list of filters that determine if a rule fires after the necessary conditions have been met. In JSON string format.
     */
    filters?: pulumi.Input<string>;
    /**
     * Perform actions at most once every `X` minutes for this issue.
     */
    frequency?: pulumi.Input<number>;
    /**
     * The issue alert name.
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization?: pulumi.Input<string>;
    /**
     * The ID of the team or user that owns the rule.
     */
    owner?: pulumi.Input<string>;
    /**
     * The slug of the project the resource belongs to.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SentryIssueAlert resource.
 */
export interface SentryIssueAlertArgs {
    /**
     * Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
     */
    actionMatch: pulumi.Input<string>;
    /**
     * List of actions. In JSON string format.
     */
    actions: pulumi.Input<string>;
    /**
     * List of conditions. In JSON string format.
     */
    conditions: pulumi.Input<string>;
    /**
     * Perform issue alert in a specific environment.
     */
    environment?: pulumi.Input<string>;
    /**
     * A string determining which filters need to be true before any actions take place. Required when a value is provided for `filters`.
     */
    filterMatch?: pulumi.Input<string>;
    /**
     * A list of filters that determine if a rule fires after the necessary conditions have been met. In JSON string format.
     */
    filters?: pulumi.Input<string>;
    /**
     * Perform actions at most once every `X` minutes for this issue.
     */
    frequency: pulumi.Input<number>;
    /**
     * The issue alert name.
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the organization the resource belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * The ID of the team or user that owns the rule.
     */
    owner?: pulumi.Input<string>;
    /**
     * The slug of the project the resource belongs to.
     */
    project: pulumi.Input<string>;
}
