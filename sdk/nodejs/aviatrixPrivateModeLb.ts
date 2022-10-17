// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class AviatrixPrivateModeLb extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixPrivateModeLb resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixPrivateModeLbState, opts?: pulumi.CustomResourceOptions): AviatrixPrivateModeLb {
        return new AviatrixPrivateModeLb(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixPrivateModeLb:AviatrixPrivateModeLb';

    /**
     * Returns true if the given object is an instance of AviatrixPrivateModeLb.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixPrivateModeLb {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixPrivateModeLb.__pulumiType;
    }

    /**
     * Name of the access account.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * Type of load balancer to create. Must be one of controller or multicloud.
     */
    public readonly lbType!: pulumi.Output<string>;
    /**
     * VPC ID of multicloud access VPC to connect to. Required when lb_type is multicloud.
     */
    public readonly multicloudAccessVpcId!: pulumi.Output<string | undefined>;
    /**
     * List of multicloud proxies.
     */
    public readonly proxies!: pulumi.Output<outputs.AviatrixPrivateModeLbProxy[] | undefined>;
    /**
     * Name of the VPC region.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * ID of the VPC for the load balancer.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a AviatrixPrivateModeLb resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixPrivateModeLbArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixPrivateModeLbArgs | AviatrixPrivateModeLbState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixPrivateModeLbState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["lbType"] = state ? state.lbType : undefined;
            resourceInputs["multicloudAccessVpcId"] = state ? state.multicloudAccessVpcId : undefined;
            resourceInputs["proxies"] = state ? state.proxies : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as AviatrixPrivateModeLbArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.lbType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lbType'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["lbType"] = args ? args.lbType : undefined;
            resourceInputs["multicloudAccessVpcId"] = args ? args.multicloudAccessVpcId : undefined;
            resourceInputs["proxies"] = args ? args.proxies : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixPrivateModeLb.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixPrivateModeLb resources.
 */
export interface AviatrixPrivateModeLbState {
    /**
     * Name of the access account.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Type of load balancer to create. Must be one of controller or multicloud.
     */
    lbType?: pulumi.Input<string>;
    /**
     * VPC ID of multicloud access VPC to connect to. Required when lb_type is multicloud.
     */
    multicloudAccessVpcId?: pulumi.Input<string>;
    /**
     * List of multicloud proxies.
     */
    proxies?: pulumi.Input<pulumi.Input<inputs.AviatrixPrivateModeLbProxy>[]>;
    /**
     * Name of the VPC region.
     */
    region?: pulumi.Input<string>;
    /**
     * ID of the VPC for the load balancer.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixPrivateModeLb resource.
 */
export interface AviatrixPrivateModeLbArgs {
    /**
     * Name of the access account.
     */
    accountName: pulumi.Input<string>;
    /**
     * Type of load balancer to create. Must be one of controller or multicloud.
     */
    lbType: pulumi.Input<string>;
    /**
     * VPC ID of multicloud access VPC to connect to. Required when lb_type is multicloud.
     */
    multicloudAccessVpcId?: pulumi.Input<string>;
    /**
     * List of multicloud proxies.
     */
    proxies?: pulumi.Input<pulumi.Input<inputs.AviatrixPrivateModeLbProxy>[]>;
    /**
     * Name of the VPC region.
     */
    region: pulumi.Input<string>;
    /**
     * ID of the VPC for the load balancer.
     */
    vpcId: pulumi.Input<string>;
}
