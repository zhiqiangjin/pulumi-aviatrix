// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixAwsTgwPeering extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAwsTgwPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAwsTgwPeeringState, opts?: pulumi.CustomResourceOptions): AviatrixAwsTgwPeering {
        return new AviatrixAwsTgwPeering(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAwsTgwPeering:AviatrixAwsTgwPeering';

    /**
     * Returns true if the given object is an instance of AviatrixAwsTgwPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAwsTgwPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAwsTgwPeering.__pulumiType;
    }

    /**
     * Name of the first AWS tgw to make a peer pair.
     */
    public readonly tgwName1!: pulumi.Output<string>;
    /**
     * Name of the second AWS tgw to make a peer pair.
     */
    public readonly tgwName2!: pulumi.Output<string>;

    /**
     * Create a AviatrixAwsTgwPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAwsTgwPeeringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAwsTgwPeeringArgs | AviatrixAwsTgwPeeringState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAwsTgwPeeringState | undefined;
            resourceInputs["tgwName1"] = state ? state.tgwName1 : undefined;
            resourceInputs["tgwName2"] = state ? state.tgwName2 : undefined;
        } else {
            const args = argsOrState as AviatrixAwsTgwPeeringArgs | undefined;
            if ((!args || args.tgwName1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tgwName1'");
            }
            if ((!args || args.tgwName2 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tgwName2'");
            }
            resourceInputs["tgwName1"] = args ? args.tgwName1 : undefined;
            resourceInputs["tgwName2"] = args ? args.tgwName2 : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAwsTgwPeering.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAwsTgwPeering resources.
 */
export interface AviatrixAwsTgwPeeringState {
    /**
     * Name of the first AWS tgw to make a peer pair.
     */
    tgwName1?: pulumi.Input<string>;
    /**
     * Name of the second AWS tgw to make a peer pair.
     */
    tgwName2?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAwsTgwPeering resource.
 */
export interface AviatrixAwsTgwPeeringArgs {
    /**
     * Name of the first AWS tgw to make a peer pair.
     */
    tgwName1: pulumi.Input<string>;
    /**
     * Name of the second AWS tgw to make a peer pair.
     */
    tgwName2: pulumi.Input<string>;
}