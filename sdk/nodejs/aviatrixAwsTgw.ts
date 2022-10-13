// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class AviatrixAwsTgw extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAwsTgw resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAwsTgwState, opts?: pulumi.CustomResourceOptions): AviatrixAwsTgw {
        return new AviatrixAwsTgw(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw';

    /**
     * Returns true if the given object is an instance of AviatrixAwsTgw.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAwsTgw {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAwsTgw.__pulumiType;
    }

    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * A list of Names of Aviatrix Transit Gateway to attach to one of the three default domains.
     *
     * @deprecated Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
     */
    public readonly attachedAviatrixTransitGateways!: pulumi.Output<string[] | undefined>;
    /**
     * BGP Local ASN (Autonomous System Number), Integer between 1-4294967294.
     */
    public readonly awsSideAsNumber!: pulumi.Output<string>;
    /**
     * TGW CIDRs.
     */
    public readonly cidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWS GOV (256). Default value: 1.
     */
    public readonly cloudType!: pulumi.Output<number | undefined>;
    /**
     * Enable Multicast.
     */
    public readonly enableMulticast!: pulumi.Output<boolean | undefined>;
    /**
     * Inspection mode. Valid values: 'Domain-based' and 'Connection-based'.
     */
    public readonly inspectionMode!: pulumi.Output<string | undefined>;
    /**
     * This parameter is a switch used to determine whether or not to manage security domains to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, security domains must be managed using the
     * aviatrix_aws_tgw_security_domain resource. Valid values: true, false. Default value: true.
     */
    public readonly manageSecurityDomain!: pulumi.Output<boolean | undefined>;
    /**
     * This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, attachment of transit gateways must be done using the
     * aviatrix_aws_tgw_transit_gateway_attachment resource. Valid values: true, false. Default value: true.
     */
    public readonly manageTransitGatewayAttachment!: pulumi.Output<boolean | undefined>;
    /**
     * This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, attachment of VPCs must be done using the
     * aviatrix_aws_tgw_vpc_attachment resource. Valid values: true, false. Default value: true.
     */
    public readonly manageVpcAttachment!: pulumi.Output<boolean | undefined>;
    /**
     * Region of cloud provider.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Security Domains to create together with AWS TGW's creation.
     *
     * @deprecated Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
     */
    public readonly securityDomains!: pulumi.Output<outputs.AviatrixAwsTgwSecurityDomain[] | undefined>;
    /**
     * TGW ID.
     */
    public /*out*/ readonly tgwId!: pulumi.Output<string>;
    /**
     * Name of the AWS TGW which is going to be created.
     */
    public readonly tgwName!: pulumi.Output<string>;

    /**
     * Create a AviatrixAwsTgw resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAwsTgwArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAwsTgwArgs | AviatrixAwsTgwState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAwsTgwState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["attachedAviatrixTransitGateways"] = state ? state.attachedAviatrixTransitGateways : undefined;
            resourceInputs["awsSideAsNumber"] = state ? state.awsSideAsNumber : undefined;
            resourceInputs["cidrs"] = state ? state.cidrs : undefined;
            resourceInputs["cloudType"] = state ? state.cloudType : undefined;
            resourceInputs["enableMulticast"] = state ? state.enableMulticast : undefined;
            resourceInputs["inspectionMode"] = state ? state.inspectionMode : undefined;
            resourceInputs["manageSecurityDomain"] = state ? state.manageSecurityDomain : undefined;
            resourceInputs["manageTransitGatewayAttachment"] = state ? state.manageTransitGatewayAttachment : undefined;
            resourceInputs["manageVpcAttachment"] = state ? state.manageVpcAttachment : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityDomains"] = state ? state.securityDomains : undefined;
            resourceInputs["tgwId"] = state ? state.tgwId : undefined;
            resourceInputs["tgwName"] = state ? state.tgwName : undefined;
        } else {
            const args = argsOrState as AviatrixAwsTgwArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.awsSideAsNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsSideAsNumber'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.tgwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tgwName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["attachedAviatrixTransitGateways"] = args ? args.attachedAviatrixTransitGateways : undefined;
            resourceInputs["awsSideAsNumber"] = args ? args.awsSideAsNumber : undefined;
            resourceInputs["cidrs"] = args ? args.cidrs : undefined;
            resourceInputs["cloudType"] = args ? args.cloudType : undefined;
            resourceInputs["enableMulticast"] = args ? args.enableMulticast : undefined;
            resourceInputs["inspectionMode"] = args ? args.inspectionMode : undefined;
            resourceInputs["manageSecurityDomain"] = args ? args.manageSecurityDomain : undefined;
            resourceInputs["manageTransitGatewayAttachment"] = args ? args.manageTransitGatewayAttachment : undefined;
            resourceInputs["manageVpcAttachment"] = args ? args.manageVpcAttachment : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityDomains"] = args ? args.securityDomains : undefined;
            resourceInputs["tgwName"] = args ? args.tgwName : undefined;
            resourceInputs["tgwId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAwsTgw.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAwsTgw resources.
 */
export interface AviatrixAwsTgwState {
    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    accountName?: pulumi.Input<string>;
    /**
     * A list of Names of Aviatrix Transit Gateway to attach to one of the three default domains.
     *
     * @deprecated Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
     */
    attachedAviatrixTransitGateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * BGP Local ASN (Autonomous System Number), Integer between 1-4294967294.
     */
    awsSideAsNumber?: pulumi.Input<string>;
    /**
     * TGW CIDRs.
     */
    cidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWS GOV (256). Default value: 1.
     */
    cloudType?: pulumi.Input<number>;
    /**
     * Enable Multicast.
     */
    enableMulticast?: pulumi.Input<boolean>;
    /**
     * Inspection mode. Valid values: 'Domain-based' and 'Connection-based'.
     */
    inspectionMode?: pulumi.Input<string>;
    /**
     * This parameter is a switch used to determine whether or not to manage security domains to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, security domains must be managed using the
     * aviatrix_aws_tgw_security_domain resource. Valid values: true, false. Default value: true.
     */
    manageSecurityDomain?: pulumi.Input<boolean>;
    /**
     * This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, attachment of transit gateways must be done using the
     * aviatrix_aws_tgw_transit_gateway_attachment resource. Valid values: true, false. Default value: true.
     */
    manageTransitGatewayAttachment?: pulumi.Input<boolean>;
    /**
     * This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, attachment of VPCs must be done using the
     * aviatrix_aws_tgw_vpc_attachment resource. Valid values: true, false. Default value: true.
     */
    manageVpcAttachment?: pulumi.Input<boolean>;
    /**
     * Region of cloud provider.
     */
    region?: pulumi.Input<string>;
    /**
     * Security Domains to create together with AWS TGW's creation.
     *
     * @deprecated Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
     */
    securityDomains?: pulumi.Input<pulumi.Input<inputs.AviatrixAwsTgwSecurityDomain>[]>;
    /**
     * TGW ID.
     */
    tgwId?: pulumi.Input<string>;
    /**
     * Name of the AWS TGW which is going to be created.
     */
    tgwName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAwsTgw resource.
 */
export interface AviatrixAwsTgwArgs {
    /**
     * This parameter represents the name of a Cloud-Account in Aviatrix controller.
     */
    accountName: pulumi.Input<string>;
    /**
     * A list of Names of Aviatrix Transit Gateway to attach to one of the three default domains.
     *
     * @deprecated Please set `manage_transit_gateway_attachment` to false, and use the standalone aviatrix_aws_tgw_transit_gateway_attachment resource instead.
     */
    attachedAviatrixTransitGateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * BGP Local ASN (Autonomous System Number), Integer between 1-4294967294.
     */
    awsSideAsNumber: pulumi.Input<string>;
    /**
     * TGW CIDRs.
     */
    cidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWS GOV (256). Default value: 1.
     */
    cloudType?: pulumi.Input<number>;
    /**
     * Enable Multicast.
     */
    enableMulticast?: pulumi.Input<boolean>;
    /**
     * Inspection mode. Valid values: 'Domain-based' and 'Connection-based'.
     */
    inspectionMode?: pulumi.Input<string>;
    /**
     * This parameter is a switch used to determine whether or not to manage security domains to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, security domains must be managed using the
     * aviatrix_aws_tgw_security_domain resource. Valid values: true, false. Default value: true.
     */
    manageSecurityDomain?: pulumi.Input<boolean>;
    /**
     * This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, attachment of transit gateways must be done using the
     * aviatrix_aws_tgw_transit_gateway_attachment resource. Valid values: true, false. Default value: true.
     */
    manageTransitGatewayAttachment?: pulumi.Input<boolean>;
    /**
     * This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the
     * aviatrix_aws_tgw resource. If this is set to false, attachment of VPCs must be done using the
     * aviatrix_aws_tgw_vpc_attachment resource. Valid values: true, false. Default value: true.
     */
    manageVpcAttachment?: pulumi.Input<boolean>;
    /**
     * Region of cloud provider.
     */
    region: pulumi.Input<string>;
    /**
     * Security Domains to create together with AWS TGW's creation.
     *
     * @deprecated Please set `manage_security_domain` to false, and use the standalone aviatrix_aws_tgw_network_domain resource instead.
     */
    securityDomains?: pulumi.Input<pulumi.Input<inputs.AviatrixAwsTgwSecurityDomain>[]>;
    /**
     * Name of the AWS TGW which is going to be created.
     */
    tgwName: pulumi.Input<string>;
}