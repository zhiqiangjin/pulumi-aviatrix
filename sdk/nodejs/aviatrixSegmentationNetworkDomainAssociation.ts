// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixSegmentationNetworkDomainAssociation extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixSegmentationNetworkDomainAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixSegmentationNetworkDomainAssociationState, opts?: pulumi.CustomResourceOptions): AviatrixSegmentationNetworkDomainAssociation {
        return new AviatrixSegmentationNetworkDomainAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixSegmentationNetworkDomainAssociation:AviatrixSegmentationNetworkDomainAssociation';

    /**
     * Returns true if the given object is an instance of AviatrixSegmentationNetworkDomainAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixSegmentationNetworkDomainAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixSegmentationNetworkDomainAssociation.__pulumiType;
    }

    /**
     * Attachment name, either Spoke or Edge.
     */
    public readonly attachmentName!: pulumi.Output<string>;
    /**
     * Network Domain name.
     */
    public readonly networkDomainName!: pulumi.Output<string>;
    /**
     * Transit Gateway name.
     */
    public readonly transitGatewayName!: pulumi.Output<string>;

    /**
     * Create a AviatrixSegmentationNetworkDomainAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixSegmentationNetworkDomainAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixSegmentationNetworkDomainAssociationArgs | AviatrixSegmentationNetworkDomainAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixSegmentationNetworkDomainAssociationState | undefined;
            resourceInputs["attachmentName"] = state ? state.attachmentName : undefined;
            resourceInputs["networkDomainName"] = state ? state.networkDomainName : undefined;
            resourceInputs["transitGatewayName"] = state ? state.transitGatewayName : undefined;
        } else {
            const args = argsOrState as AviatrixSegmentationNetworkDomainAssociationArgs | undefined;
            if ((!args || args.attachmentName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attachmentName'");
            }
            if ((!args || args.networkDomainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkDomainName'");
            }
            if ((!args || args.transitGatewayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayName'");
            }
            resourceInputs["attachmentName"] = args ? args.attachmentName : undefined;
            resourceInputs["networkDomainName"] = args ? args.networkDomainName : undefined;
            resourceInputs["transitGatewayName"] = args ? args.transitGatewayName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixSegmentationNetworkDomainAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixSegmentationNetworkDomainAssociation resources.
 */
export interface AviatrixSegmentationNetworkDomainAssociationState {
    /**
     * Attachment name, either Spoke or Edge.
     */
    attachmentName?: pulumi.Input<string>;
    /**
     * Network Domain name.
     */
    networkDomainName?: pulumi.Input<string>;
    /**
     * Transit Gateway name.
     */
    transitGatewayName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixSegmentationNetworkDomainAssociation resource.
 */
export interface AviatrixSegmentationNetworkDomainAssociationArgs {
    /**
     * Attachment name, either Spoke or Edge.
     */
    attachmentName: pulumi.Input<string>;
    /**
     * Network Domain name.
     */
    networkDomainName: pulumi.Input<string>;
    /**
     * Transit Gateway name.
     */
    transitGatewayName: pulumi.Input<string>;
}
