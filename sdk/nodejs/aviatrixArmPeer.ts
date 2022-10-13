// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixArmPeer extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixArmPeer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixArmPeerState, opts?: pulumi.CustomResourceOptions): AviatrixArmPeer {
        return new AviatrixArmPeer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixArmPeer:AviatrixArmPeer';

    /**
     * Returns true if the given object is an instance of AviatrixArmPeer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixArmPeer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixArmPeer.__pulumiType;
    }

    /**
     * This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
     */
    public readonly accountName1!: pulumi.Output<string>;
    /**
     * This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
     */
    public readonly accountName2!: pulumi.Output<string>;
    /**
     * List of VNet CIDR of vnet_name_resource_group1.
     */
    public /*out*/ readonly vnetCidr1s!: pulumi.Output<string[]>;
    /**
     * List of VNet CIDR of vnet_name_resource_group2.
     */
    public /*out*/ readonly vnetCidr2s!: pulumi.Output<string[]>;
    /**
     * VNet-Name of Azure cloud.
     */
    public readonly vnetNameResourceGroup1!: pulumi.Output<string>;
    /**
     * VNet-Name of Azure cloud.
     */
    public readonly vnetNameResourceGroup2!: pulumi.Output<string>;
    /**
     * Region of Azure cloud.
     */
    public readonly vnetReg1!: pulumi.Output<string>;
    /**
     * Region of Azure cloud.
     */
    public readonly vnetReg2!: pulumi.Output<string>;

    /**
     * Create a AviatrixArmPeer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixArmPeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixArmPeerArgs | AviatrixArmPeerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixArmPeerState | undefined;
            resourceInputs["accountName1"] = state ? state.accountName1 : undefined;
            resourceInputs["accountName2"] = state ? state.accountName2 : undefined;
            resourceInputs["vnetCidr1s"] = state ? state.vnetCidr1s : undefined;
            resourceInputs["vnetCidr2s"] = state ? state.vnetCidr2s : undefined;
            resourceInputs["vnetNameResourceGroup1"] = state ? state.vnetNameResourceGroup1 : undefined;
            resourceInputs["vnetNameResourceGroup2"] = state ? state.vnetNameResourceGroup2 : undefined;
            resourceInputs["vnetReg1"] = state ? state.vnetReg1 : undefined;
            resourceInputs["vnetReg2"] = state ? state.vnetReg2 : undefined;
        } else {
            const args = argsOrState as AviatrixArmPeerArgs | undefined;
            if ((!args || args.accountName1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName1'");
            }
            if ((!args || args.accountName2 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName2'");
            }
            if ((!args || args.vnetNameResourceGroup1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vnetNameResourceGroup1'");
            }
            if ((!args || args.vnetNameResourceGroup2 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vnetNameResourceGroup2'");
            }
            if ((!args || args.vnetReg1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vnetReg1'");
            }
            if ((!args || args.vnetReg2 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vnetReg2'");
            }
            resourceInputs["accountName1"] = args ? args.accountName1 : undefined;
            resourceInputs["accountName2"] = args ? args.accountName2 : undefined;
            resourceInputs["vnetNameResourceGroup1"] = args ? args.vnetNameResourceGroup1 : undefined;
            resourceInputs["vnetNameResourceGroup2"] = args ? args.vnetNameResourceGroup2 : undefined;
            resourceInputs["vnetReg1"] = args ? args.vnetReg1 : undefined;
            resourceInputs["vnetReg2"] = args ? args.vnetReg2 : undefined;
            resourceInputs["vnetCidr1s"] = undefined /*out*/;
            resourceInputs["vnetCidr2s"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixArmPeer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixArmPeer resources.
 */
export interface AviatrixArmPeerState {
    /**
     * This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
     */
    accountName1?: pulumi.Input<string>;
    /**
     * This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
     */
    accountName2?: pulumi.Input<string>;
    /**
     * List of VNet CIDR of vnet_name_resource_group1.
     */
    vnetCidr1s?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of VNet CIDR of vnet_name_resource_group2.
     */
    vnetCidr2s?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VNet-Name of Azure cloud.
     */
    vnetNameResourceGroup1?: pulumi.Input<string>;
    /**
     * VNet-Name of Azure cloud.
     */
    vnetNameResourceGroup2?: pulumi.Input<string>;
    /**
     * Region of Azure cloud.
     */
    vnetReg1?: pulumi.Input<string>;
    /**
     * Region of Azure cloud.
     */
    vnetReg2?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixArmPeer resource.
 */
export interface AviatrixArmPeerArgs {
    /**
     * This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
     */
    accountName1: pulumi.Input<string>;
    /**
     * This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
     */
    accountName2: pulumi.Input<string>;
    /**
     * VNet-Name of Azure cloud.
     */
    vnetNameResourceGroup1: pulumi.Input<string>;
    /**
     * VNet-Name of Azure cloud.
     */
    vnetNameResourceGroup2: pulumi.Input<string>;
    /**
     * Region of Azure cloud.
     */
    vnetReg1: pulumi.Input<string>;
    /**
     * Region of Azure cloud.
     */
    vnetReg2: pulumi.Input<string>;
}