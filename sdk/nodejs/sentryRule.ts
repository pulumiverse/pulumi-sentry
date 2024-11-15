// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > **WARNING:** This resource is deprecated and will be removed in the next major version. Use the `sentry.SentryIssueAlert` resource instead.
 */
export class SentryRule extends pulumi.CustomResource {
    /**
     * Get an existing SentryRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SentryRuleState, opts?: pulumi.CustomResourceOptions): SentryRule {
        return new SentryRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sentry:index/sentryRule:SentryRule';

    /**
     * Returns true if the given object is an instance of SentryRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SentryRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SentryRule.__pulumiType;
    }

    /**
     * Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
     */
    public readonly actionMatch!: pulumi.Output<string>;
    /**
     * List of actions.
     */
    public readonly actions!: pulumi.Output<{[key: string]: string}[]>;
    /**
     * List of conditions.
     */
    public readonly conditions!: pulumi.Output<{[key: string]: string}[]>;
    /**
     * Perform issue alert in a specific environment.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * Trigger actions if `all`, `any`, or `none` of the specified filters match.
     */
    public readonly filterMatch!: pulumi.Output<string>;
    /**
     * List of filters.
     */
    public readonly filters!: pulumi.Output<{[key: string]: string}[] | undefined>;
    /**
     * Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
     */
    public readonly frequency!: pulumi.Output<number>;
    /**
     * The internal ID for this issue alert.
     */
    public /*out*/ readonly internalId!: pulumi.Output<string>;
    /**
     * The issue alert name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The slug of the organization the issue alert belongs to.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * The slug of the project to create the issue alert for.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Use `project` (singular) instead.
     *
     * @deprecated Use `project` (singular) instead.
     */
    public /*out*/ readonly projects!: pulumi.Output<string[]>;

    /**
     * Create a SentryRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SentryRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SentryRuleArgs | SentryRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SentryRuleState | undefined;
            resourceInputs["actionMatch"] = state ? state.actionMatch : undefined;
            resourceInputs["actions"] = state ? state.actions : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["filterMatch"] = state ? state.filterMatch : undefined;
            resourceInputs["filters"] = state ? state.filters : undefined;
            resourceInputs["frequency"] = state ? state.frequency : undefined;
            resourceInputs["internalId"] = state ? state.internalId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projects"] = state ? state.projects : undefined;
        } else {
            const args = argsOrState as SentryRuleArgs | undefined;
            if ((!args || args.actionMatch === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actionMatch'");
            }
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.conditions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'conditions'");
            }
            if ((!args || args.filterMatch === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filterMatch'");
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
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["internalId"] = undefined /*out*/;
            resourceInputs["projects"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SentryRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SentryRule resources.
 */
export interface SentryRuleState {
    /**
     * Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
     */
    actionMatch?: pulumi.Input<string>;
    /**
     * List of actions.
     */
    actions?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * List of conditions.
     */
    conditions?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * Perform issue alert in a specific environment.
     */
    environment?: pulumi.Input<string>;
    /**
     * Trigger actions if `all`, `any`, or `none` of the specified filters match.
     */
    filterMatch?: pulumi.Input<string>;
    /**
     * List of filters.
     */
    filters?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
     */
    frequency?: pulumi.Input<number>;
    /**
     * The internal ID for this issue alert.
     */
    internalId?: pulumi.Input<string>;
    /**
     * The issue alert name.
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the organization the issue alert belongs to.
     */
    organization?: pulumi.Input<string>;
    /**
     * The slug of the project to create the issue alert for.
     */
    project?: pulumi.Input<string>;
    /**
     * Use `project` (singular) instead.
     *
     * @deprecated Use `project` (singular) instead.
     */
    projects?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SentryRule resource.
 */
export interface SentryRuleArgs {
    /**
     * Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
     */
    actionMatch: pulumi.Input<string>;
    /**
     * List of actions.
     */
    actions: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * List of conditions.
     */
    conditions: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * Perform issue alert in a specific environment.
     */
    environment?: pulumi.Input<string>;
    /**
     * Trigger actions if `all`, `any`, or `none` of the specified filters match.
     */
    filterMatch: pulumi.Input<string>;
    /**
     * List of filters.
     */
    filters?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
     */
    frequency: pulumi.Input<number>;
    /**
     * The issue alert name.
     */
    name?: pulumi.Input<string>;
    /**
     * The slug of the organization the issue alert belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * The slug of the project to create the issue alert for.
     */
    project: pulumi.Input<string>;
}
