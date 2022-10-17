// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class AviatrixGatewaySnat extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixGatewaySnat resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixGatewaySnatState, opts?: pulumi.CustomResourceOptions): AviatrixGatewaySnat {
        return new AviatrixGatewaySnat(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixGatewaySnat:AviatrixGatewaySnat';

    /**
     * Returns true if the given object is an instance of AviatrixGatewaySnat.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixGatewaySnat {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixGatewaySnat.__pulumiType;
    }

    /**
     * Computed attribute to store the previous connection policy.
     */
    public /*out*/ readonly connectionPolicies!: pulumi.Output<outputs.AviatrixGatewaySnatConnectionPolicy[]>;
    /**
     * Name of the gateway.
     */
    public readonly gwName!: pulumi.Output<string>;
    /**
     * Computed attribute to store the previous interface policy.
     */
    public /*out*/ readonly interfacePolicies!: pulumi.Output<outputs.AviatrixGatewaySnatInterfacePolicy[]>;
    /**
     * Nat mode. Currently only supports 'customized_snat'.
     */
    public readonly snatMode!: pulumi.Output<string | undefined>;
    /**
     * Policy rules applied for 'snat_mode'' of 'customized_snat'.'
     */
    public readonly snatPolicies!: pulumi.Output<outputs.AviatrixGatewaySnatSnatPolicy[] | undefined>;
    /**
     * Whether to sync the policies to the HA gateway.
     */
    public readonly syncToHa!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AviatrixGatewaySnat resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixGatewaySnatArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixGatewaySnatArgs | AviatrixGatewaySnatState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixGatewaySnatState | undefined;
            resourceInputs["connectionPolicies"] = state ? state.connectionPolicies : undefined;
            resourceInputs["gwName"] = state ? state.gwName : undefined;
            resourceInputs["interfacePolicies"] = state ? state.interfacePolicies : undefined;
            resourceInputs["snatMode"] = state ? state.snatMode : undefined;
            resourceInputs["snatPolicies"] = state ? state.snatPolicies : undefined;
            resourceInputs["syncToHa"] = state ? state.syncToHa : undefined;
        } else {
            const args = argsOrState as AviatrixGatewaySnatArgs | undefined;
            if ((!args || args.gwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gwName'");
            }
            resourceInputs["gwName"] = args ? args.gwName : undefined;
            resourceInputs["snatMode"] = args ? args.snatMode : undefined;
            resourceInputs["snatPolicies"] = args ? args.snatPolicies : undefined;
            resourceInputs["syncToHa"] = args ? args.syncToHa : undefined;
            resourceInputs["connectionPolicies"] = undefined /*out*/;
            resourceInputs["interfacePolicies"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixGatewaySnat.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixGatewaySnat resources.
 */
export interface AviatrixGatewaySnatState {
    /**
     * Computed attribute to store the previous connection policy.
     */
    connectionPolicies?: pulumi.Input<pulumi.Input<inputs.AviatrixGatewaySnatConnectionPolicy>[]>;
    /**
     * Name of the gateway.
     */
    gwName?: pulumi.Input<string>;
    /**
     * Computed attribute to store the previous interface policy.
     */
    interfacePolicies?: pulumi.Input<pulumi.Input<inputs.AviatrixGatewaySnatInterfacePolicy>[]>;
    /**
     * Nat mode. Currently only supports 'customized_snat'.
     */
    snatMode?: pulumi.Input<string>;
    /**
     * Policy rules applied for 'snat_mode'' of 'customized_snat'.'
     */
    snatPolicies?: pulumi.Input<pulumi.Input<inputs.AviatrixGatewaySnatSnatPolicy>[]>;
    /**
     * Whether to sync the policies to the HA gateway.
     */
    syncToHa?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AviatrixGatewaySnat resource.
 */
export interface AviatrixGatewaySnatArgs {
    /**
     * Name of the gateway.
     */
    gwName: pulumi.Input<string>;
    /**
     * Nat mode. Currently only supports 'customized_snat'.
     */
    snatMode?: pulumi.Input<string>;
    /**
     * Policy rules applied for 'snat_mode'' of 'customized_snat'.'
     */
    snatPolicies?: pulumi.Input<pulumi.Input<inputs.AviatrixGatewaySnatSnatPolicy>[]>;
    /**
     * Whether to sync the policies to the HA gateway.
     */
    syncToHa?: pulumi.Input<boolean>;
}
