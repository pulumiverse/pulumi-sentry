// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Sentry Plugin resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sentry from "@pulumiverse/sentry";
 *
 * // Create a plugin
 * const _default = new sentry.SentryPlugin("default", {
 *     organization: "my-organization",
 *     project: "web-app",
 *     plugin: "slack",
 *     config: {
 *         webhook: "slack://webhook",
 *     },
 * });
 * ```
 */
export class SentryPlugin extends pulumi.CustomResource {
    /**
     * Get an existing SentryPlugin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SentryPluginState, opts?: pulumi.CustomResourceOptions): SentryPlugin {
        return new SentryPlugin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sentry:index/sentryPlugin:SentryPlugin';

    /**
     * Returns true if the given object is an instance of SentryPlugin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SentryPlugin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SentryPlugin.__pulumiType;
    }

    /**
     * Plugin config.
     */
    public readonly config!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The slug of the organization the project belongs to.
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * Plugin ID.
     */
    public readonly plugin!: pulumi.Output<string>;
    /**
     * The slug of the project to create the plugin for.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a SentryPlugin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SentryPluginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SentryPluginArgs | SentryPluginState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SentryPluginState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["plugin"] = state ? state.plugin : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as SentryPluginArgs | undefined;
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.plugin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plugin'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["plugin"] = args ? args.plugin : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SentryPlugin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SentryPlugin resources.
 */
export interface SentryPluginState {
    /**
     * Plugin config.
     */
    config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The slug of the organization the project belongs to.
     */
    organization?: pulumi.Input<string>;
    /**
     * Plugin ID.
     */
    plugin?: pulumi.Input<string>;
    /**
     * The slug of the project to create the plugin for.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SentryPlugin resource.
 */
export interface SentryPluginArgs {
    /**
     * Plugin config.
     */
    config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The slug of the organization the project belongs to.
     */
    organization: pulumi.Input<string>;
    /**
     * Plugin ID.
     */
    plugin: pulumi.Input<string>;
    /**
     * The slug of the project to create the plugin for.
     */
    project: pulumi.Input<string>;
}
