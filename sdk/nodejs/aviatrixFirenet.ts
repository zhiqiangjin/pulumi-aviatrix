// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class AviatrixFirenet extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixFirenet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixFirenetState, opts?: pulumi.CustomResourceOptions): AviatrixFirenet {
        return new AviatrixFirenet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixFirenet:AviatrixFirenet';

    /**
     * Returns true if the given object is an instance of AviatrixFirenet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixFirenet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixFirenet.__pulumiType;
    }

    /**
     * Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as
     * of provider version R2.19.2+.
     */
    public readonly eastWestInspectionExcludedCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Enable/Disable egress through firewall.
     */
    public readonly egressEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of egress static cidrs.
     */
    public readonly egressStaticCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * List of firewall instances to be associated with fireNet.
     *
     * @deprecated Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
     */
    public readonly firewallInstanceAssociations!: pulumi.Output<outputs.AviatrixFirenetFirewallInstanceAssociation[] | undefined>;
    /**
     * Hashing algorithm to load balance traffic across the firewall.
     */
    public readonly hashingAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * Enable/Disable traffic inspection.
     */
    public readonly inspectionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enable Keep Alive via Firewall LAN Interface.
     */
    public readonly keepAliveViaLanInterfaceEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enable this to manage firewall_instance_associations in-line. If this is false, associations must be managed via
     * standalone aviatrix_firewall_instance_association resources. Type: boolean, Default: true, Valid values: true/false.
     */
    public readonly manageFirewallInstanceAssociation!: pulumi.Output<boolean | undefined>;
    /**
     * Enable TGW segmentation for egress.
     */
    public readonly tgwSegmentationForEgressEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * VPC ID.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a AviatrixFirenet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixFirenetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixFirenetArgs | AviatrixFirenetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixFirenetState | undefined;
            resourceInputs["eastWestInspectionExcludedCidrs"] = state ? state.eastWestInspectionExcludedCidrs : undefined;
            resourceInputs["egressEnabled"] = state ? state.egressEnabled : undefined;
            resourceInputs["egressStaticCidrs"] = state ? state.egressStaticCidrs : undefined;
            resourceInputs["firewallInstanceAssociations"] = state ? state.firewallInstanceAssociations : undefined;
            resourceInputs["hashingAlgorithm"] = state ? state.hashingAlgorithm : undefined;
            resourceInputs["inspectionEnabled"] = state ? state.inspectionEnabled : undefined;
            resourceInputs["keepAliveViaLanInterfaceEnabled"] = state ? state.keepAliveViaLanInterfaceEnabled : undefined;
            resourceInputs["manageFirewallInstanceAssociation"] = state ? state.manageFirewallInstanceAssociation : undefined;
            resourceInputs["tgwSegmentationForEgressEnabled"] = state ? state.tgwSegmentationForEgressEnabled : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as AviatrixFirenetArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["eastWestInspectionExcludedCidrs"] = args ? args.eastWestInspectionExcludedCidrs : undefined;
            resourceInputs["egressEnabled"] = args ? args.egressEnabled : undefined;
            resourceInputs["egressStaticCidrs"] = args ? args.egressStaticCidrs : undefined;
            resourceInputs["firewallInstanceAssociations"] = args ? args.firewallInstanceAssociations : undefined;
            resourceInputs["hashingAlgorithm"] = args ? args.hashingAlgorithm : undefined;
            resourceInputs["inspectionEnabled"] = args ? args.inspectionEnabled : undefined;
            resourceInputs["keepAliveViaLanInterfaceEnabled"] = args ? args.keepAliveViaLanInterfaceEnabled : undefined;
            resourceInputs["manageFirewallInstanceAssociation"] = args ? args.manageFirewallInstanceAssociation : undefined;
            resourceInputs["tgwSegmentationForEgressEnabled"] = args ? args.tgwSegmentationForEgressEnabled : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixFirenet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixFirenet resources.
 */
export interface AviatrixFirenetState {
    /**
     * Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as
     * of provider version R2.19.2+.
     */
    eastWestInspectionExcludedCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable/Disable egress through firewall.
     */
    egressEnabled?: pulumi.Input<boolean>;
    /**
     * List of egress static cidrs.
     */
    egressStaticCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of firewall instances to be associated with fireNet.
     *
     * @deprecated Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
     */
    firewallInstanceAssociations?: pulumi.Input<pulumi.Input<inputs.AviatrixFirenetFirewallInstanceAssociation>[]>;
    /**
     * Hashing algorithm to load balance traffic across the firewall.
     */
    hashingAlgorithm?: pulumi.Input<string>;
    /**
     * Enable/Disable traffic inspection.
     */
    inspectionEnabled?: pulumi.Input<boolean>;
    /**
     * Enable Keep Alive via Firewall LAN Interface.
     */
    keepAliveViaLanInterfaceEnabled?: pulumi.Input<boolean>;
    /**
     * Enable this to manage firewall_instance_associations in-line. If this is false, associations must be managed via
     * standalone aviatrix_firewall_instance_association resources. Type: boolean, Default: true, Valid values: true/false.
     */
    manageFirewallInstanceAssociation?: pulumi.Input<boolean>;
    /**
     * Enable TGW segmentation for egress.
     */
    tgwSegmentationForEgressEnabled?: pulumi.Input<boolean>;
    /**
     * VPC ID.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixFirenet resource.
 */
export interface AviatrixFirenetArgs {
    /**
     * Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as
     * of provider version R2.19.2+.
     */
    eastWestInspectionExcludedCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable/Disable egress through firewall.
     */
    egressEnabled?: pulumi.Input<boolean>;
    /**
     * List of egress static cidrs.
     */
    egressStaticCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of firewall instances to be associated with fireNet.
     *
     * @deprecated Please set `manage_firewall_instance_association` to false, and use the standalone aviatrix_firewall_instance_association resource instead.
     */
    firewallInstanceAssociations?: pulumi.Input<pulumi.Input<inputs.AviatrixFirenetFirewallInstanceAssociation>[]>;
    /**
     * Hashing algorithm to load balance traffic across the firewall.
     */
    hashingAlgorithm?: pulumi.Input<string>;
    /**
     * Enable/Disable traffic inspection.
     */
    inspectionEnabled?: pulumi.Input<boolean>;
    /**
     * Enable Keep Alive via Firewall LAN Interface.
     */
    keepAliveViaLanInterfaceEnabled?: pulumi.Input<boolean>;
    /**
     * Enable this to manage firewall_instance_associations in-line. If this is false, associations must be managed via
     * standalone aviatrix_firewall_instance_association resources. Type: boolean, Default: true, Valid values: true/false.
     */
    manageFirewallInstanceAssociation?: pulumi.Input<boolean>;
    /**
     * Enable TGW segmentation for egress.
     */
    tgwSegmentationForEgressEnabled?: pulumi.Input<boolean>;
    /**
     * VPC ID.
     */
    vpcId: pulumi.Input<string>;
}
