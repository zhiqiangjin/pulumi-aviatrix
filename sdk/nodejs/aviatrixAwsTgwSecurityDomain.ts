// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixAwsTgwSecurityDomain extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAwsTgwSecurityDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAwsTgwSecurityDomainState, opts?: pulumi.CustomResourceOptions): AviatrixAwsTgwSecurityDomain {
        return new AviatrixAwsTgwSecurityDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAwsTgwSecurityDomain:AviatrixAwsTgwSecurityDomain';

    /**
     * Returns true if the given object is an instance of AviatrixAwsTgwSecurityDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAwsTgwSecurityDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAwsTgwSecurityDomain.__pulumiType;
    }

    /**
     * Set to true if the security domain is an aviatrix firewall domain.
     */
    public readonly aviatrixFirewall!: pulumi.Output<boolean | undefined>;
    /**
     * Security domain name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Set to true if the security domain is a native egress domain.
     */
    public readonly nativeEgress!: pulumi.Output<boolean | undefined>;
    /**
     * Set to true if the security domain is a native firewall domain.
     */
    public readonly nativeFirewall!: pulumi.Output<boolean | undefined>;
    /**
     * AWS TGW name.
     */
    public readonly tgwName!: pulumi.Output<string>;

    /**
     * Create a AviatrixAwsTgwSecurityDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAwsTgwSecurityDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAwsTgwSecurityDomainArgs | AviatrixAwsTgwSecurityDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAwsTgwSecurityDomainState | undefined;
            resourceInputs["aviatrixFirewall"] = state ? state.aviatrixFirewall : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nativeEgress"] = state ? state.nativeEgress : undefined;
            resourceInputs["nativeFirewall"] = state ? state.nativeFirewall : undefined;
            resourceInputs["tgwName"] = state ? state.tgwName : undefined;
        } else {
            const args = argsOrState as AviatrixAwsTgwSecurityDomainArgs | undefined;
            if ((!args || args.tgwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tgwName'");
            }
            resourceInputs["aviatrixFirewall"] = args ? args.aviatrixFirewall : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nativeEgress"] = args ? args.nativeEgress : undefined;
            resourceInputs["nativeFirewall"] = args ? args.nativeFirewall : undefined;
            resourceInputs["tgwName"] = args ? args.tgwName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAwsTgwSecurityDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAwsTgwSecurityDomain resources.
 */
export interface AviatrixAwsTgwSecurityDomainState {
    /**
     * Set to true if the security domain is an aviatrix firewall domain.
     */
    aviatrixFirewall?: pulumi.Input<boolean>;
    /**
     * Security domain name.
     */
    name?: pulumi.Input<string>;
    /**
     * Set to true if the security domain is a native egress domain.
     */
    nativeEgress?: pulumi.Input<boolean>;
    /**
     * Set to true if the security domain is a native firewall domain.
     */
    nativeFirewall?: pulumi.Input<boolean>;
    /**
     * AWS TGW name.
     */
    tgwName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAwsTgwSecurityDomain resource.
 */
export interface AviatrixAwsTgwSecurityDomainArgs {
    /**
     * Set to true if the security domain is an aviatrix firewall domain.
     */
    aviatrixFirewall?: pulumi.Input<boolean>;
    /**
     * Security domain name.
     */
    name?: pulumi.Input<string>;
    /**
     * Set to true if the security domain is a native egress domain.
     */
    nativeEgress?: pulumi.Input<boolean>;
    /**
     * Set to true if the security domain is a native firewall domain.
     */
    nativeFirewall?: pulumi.Input<boolean>;
    /**
     * AWS TGW name.
     */
    tgwName: pulumi.Input<string>;
}