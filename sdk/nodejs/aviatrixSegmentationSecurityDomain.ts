// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixSegmentationSecurityDomain extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixSegmentationSecurityDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixSegmentationSecurityDomainState, opts?: pulumi.CustomResourceOptions): AviatrixSegmentationSecurityDomain {
        return new AviatrixSegmentationSecurityDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixSegmentationSecurityDomain:AviatrixSegmentationSecurityDomain';

    /**
     * Returns true if the given object is an instance of AviatrixSegmentationSecurityDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixSegmentationSecurityDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixSegmentationSecurityDomain.__pulumiType;
    }

    /**
     * Security domain name.
     */
    public readonly domainName!: pulumi.Output<string>;

    /**
     * Create a AviatrixSegmentationSecurityDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixSegmentationSecurityDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixSegmentationSecurityDomainArgs | AviatrixSegmentationSecurityDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixSegmentationSecurityDomainState | undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
        } else {
            const args = argsOrState as AviatrixSegmentationSecurityDomainArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            resourceInputs["domainName"] = args ? args.domainName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixSegmentationSecurityDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixSegmentationSecurityDomain resources.
 */
export interface AviatrixSegmentationSecurityDomainState {
    /**
     * Security domain name.
     */
    domainName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixSegmentationSecurityDomain resource.
 */
export interface AviatrixSegmentationSecurityDomainArgs {
    /**
     * Security domain name.
     */
    domainName: pulumi.Input<string>;
}
