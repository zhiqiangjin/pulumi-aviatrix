// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixAwsTgwDirectconnect extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAwsTgwDirectconnect resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAwsTgwDirectconnectState, opts?: pulumi.CustomResourceOptions): AviatrixAwsTgwDirectconnect {
        return new AviatrixAwsTgwDirectconnect(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAwsTgwDirectconnect:AviatrixAwsTgwDirectconnect';

    /**
     * Returns true if the given object is an instance of AviatrixAwsTgwDirectconnect.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAwsTgwDirectconnect {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAwsTgwDirectconnect.__pulumiType;
    }

    /**
     * Public IP address. Example: '40.0.0.0'.
     */
    public readonly allowedPrefix!: pulumi.Output<string>;
    /**
     * This parameter represents the name of an Account in Aviatrix controller.
     */
    public readonly directconnectAccountName!: pulumi.Output<string>;
    /**
     * This parameter represents the name of a Direct Connect Gateway ID.
     */
    public readonly dxGatewayId!: pulumi.Output<string>;
    /**
     * Switch to enable/disable encrypted transit approval for direct connection. Valid values: true, false.
     */
    public readonly enableLearnedCidrsApproval!: pulumi.Output<boolean | undefined>;
    /**
     * The name of an Aviatrix network domain, to which the direct connect gateway will be attached.
     */
    public readonly networkDomainName!: pulumi.Output<string | undefined>;
    /**
     * The name of an Aviatrix security domain, to which the direct connect gateway will be attached.
     *
     * @deprecated Please use network_domain_name instead.
     */
    public readonly securityDomainName!: pulumi.Output<string | undefined>;
    /**
     * This parameter represents the name of an AWS TGW.
     */
    public readonly tgwName!: pulumi.Output<string>;

    /**
     * Create a AviatrixAwsTgwDirectconnect resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAwsTgwDirectconnectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAwsTgwDirectconnectArgs | AviatrixAwsTgwDirectconnectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAwsTgwDirectconnectState | undefined;
            resourceInputs["allowedPrefix"] = state ? state.allowedPrefix : undefined;
            resourceInputs["directconnectAccountName"] = state ? state.directconnectAccountName : undefined;
            resourceInputs["dxGatewayId"] = state ? state.dxGatewayId : undefined;
            resourceInputs["enableLearnedCidrsApproval"] = state ? state.enableLearnedCidrsApproval : undefined;
            resourceInputs["networkDomainName"] = state ? state.networkDomainName : undefined;
            resourceInputs["securityDomainName"] = state ? state.securityDomainName : undefined;
            resourceInputs["tgwName"] = state ? state.tgwName : undefined;
        } else {
            const args = argsOrState as AviatrixAwsTgwDirectconnectArgs | undefined;
            if ((!args || args.allowedPrefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedPrefix'");
            }
            if ((!args || args.directconnectAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'directconnectAccountName'");
            }
            if ((!args || args.dxGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dxGatewayId'");
            }
            if ((!args || args.tgwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tgwName'");
            }
            resourceInputs["allowedPrefix"] = args ? args.allowedPrefix : undefined;
            resourceInputs["directconnectAccountName"] = args ? args.directconnectAccountName : undefined;
            resourceInputs["dxGatewayId"] = args ? args.dxGatewayId : undefined;
            resourceInputs["enableLearnedCidrsApproval"] = args ? args.enableLearnedCidrsApproval : undefined;
            resourceInputs["networkDomainName"] = args ? args.networkDomainName : undefined;
            resourceInputs["securityDomainName"] = args ? args.securityDomainName : undefined;
            resourceInputs["tgwName"] = args ? args.tgwName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAwsTgwDirectconnect.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAwsTgwDirectconnect resources.
 */
export interface AviatrixAwsTgwDirectconnectState {
    /**
     * Public IP address. Example: '40.0.0.0'.
     */
    allowedPrefix?: pulumi.Input<string>;
    /**
     * This parameter represents the name of an Account in Aviatrix controller.
     */
    directconnectAccountName?: pulumi.Input<string>;
    /**
     * This parameter represents the name of a Direct Connect Gateway ID.
     */
    dxGatewayId?: pulumi.Input<string>;
    /**
     * Switch to enable/disable encrypted transit approval for direct connection. Valid values: true, false.
     */
    enableLearnedCidrsApproval?: pulumi.Input<boolean>;
    /**
     * The name of an Aviatrix network domain, to which the direct connect gateway will be attached.
     */
    networkDomainName?: pulumi.Input<string>;
    /**
     * The name of an Aviatrix security domain, to which the direct connect gateway will be attached.
     *
     * @deprecated Please use network_domain_name instead.
     */
    securityDomainName?: pulumi.Input<string>;
    /**
     * This parameter represents the name of an AWS TGW.
     */
    tgwName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAwsTgwDirectconnect resource.
 */
export interface AviatrixAwsTgwDirectconnectArgs {
    /**
     * Public IP address. Example: '40.0.0.0'.
     */
    allowedPrefix: pulumi.Input<string>;
    /**
     * This parameter represents the name of an Account in Aviatrix controller.
     */
    directconnectAccountName: pulumi.Input<string>;
    /**
     * This parameter represents the name of a Direct Connect Gateway ID.
     */
    dxGatewayId: pulumi.Input<string>;
    /**
     * Switch to enable/disable encrypted transit approval for direct connection. Valid values: true, false.
     */
    enableLearnedCidrsApproval?: pulumi.Input<boolean>;
    /**
     * The name of an Aviatrix network domain, to which the direct connect gateway will be attached.
     */
    networkDomainName?: pulumi.Input<string>;
    /**
     * The name of an Aviatrix security domain, to which the direct connect gateway will be attached.
     *
     * @deprecated Please use network_domain_name instead.
     */
    securityDomainName?: pulumi.Input<string>;
    /**
     * This parameter represents the name of an AWS TGW.
     */
    tgwName: pulumi.Input<string>;
}