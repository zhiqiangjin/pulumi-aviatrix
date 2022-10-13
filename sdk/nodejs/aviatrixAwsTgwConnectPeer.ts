// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixAwsTgwConnectPeer extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAwsTgwConnectPeer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAwsTgwConnectPeerState, opts?: pulumi.CustomResourceOptions): AviatrixAwsTgwConnectPeer {
        return new AviatrixAwsTgwConnectPeer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAwsTgwConnectPeer:AviatrixAwsTgwConnectPeer';

    /**
     * Returns true if the given object is an instance of AviatrixAwsTgwConnectPeer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAwsTgwConnectPeer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAwsTgwConnectPeer.__pulumiType;
    }

    /**
     * Set of BGP Inside CIDR Blocks.
     */
    public readonly bgpInsideCidrs!: pulumi.Output<string[]>;
    /**
     * Connect Attachment ID.
     */
    public readonly connectAttachmentId!: pulumi.Output<string>;
    /**
     * Connect Peer ID.
     */
    public /*out*/ readonly connectPeerId!: pulumi.Output<string>;
    /**
     * Connect Peer Name.
     */
    public readonly connectPeerName!: pulumi.Output<string>;
    /**
     * AWS TGW Connect connection name.
     */
    public readonly connectionName!: pulumi.Output<string>;
    /**
     * Peer AS Number.
     */
    public readonly peerAsNumber!: pulumi.Output<string>;
    /**
     * Peer GRE IP Address.
     */
    public readonly peerGreAddress!: pulumi.Output<string>;
    /**
     * AWS TGW GRE IP Address.
     */
    public readonly tgwGreAddress!: pulumi.Output<string | undefined>;
    /**
     * AWS TGW Name.
     */
    public readonly tgwName!: pulumi.Output<string>;

    /**
     * Create a AviatrixAwsTgwConnectPeer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAwsTgwConnectPeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAwsTgwConnectPeerArgs | AviatrixAwsTgwConnectPeerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAwsTgwConnectPeerState | undefined;
            resourceInputs["bgpInsideCidrs"] = state ? state.bgpInsideCidrs : undefined;
            resourceInputs["connectAttachmentId"] = state ? state.connectAttachmentId : undefined;
            resourceInputs["connectPeerId"] = state ? state.connectPeerId : undefined;
            resourceInputs["connectPeerName"] = state ? state.connectPeerName : undefined;
            resourceInputs["connectionName"] = state ? state.connectionName : undefined;
            resourceInputs["peerAsNumber"] = state ? state.peerAsNumber : undefined;
            resourceInputs["peerGreAddress"] = state ? state.peerGreAddress : undefined;
            resourceInputs["tgwGreAddress"] = state ? state.tgwGreAddress : undefined;
            resourceInputs["tgwName"] = state ? state.tgwName : undefined;
        } else {
            const args = argsOrState as AviatrixAwsTgwConnectPeerArgs | undefined;
            if ((!args || args.bgpInsideCidrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpInsideCidrs'");
            }
            if ((!args || args.connectAttachmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectAttachmentId'");
            }
            if ((!args || args.connectPeerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectPeerName'");
            }
            if ((!args || args.connectionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionName'");
            }
            if ((!args || args.peerAsNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerAsNumber'");
            }
            if ((!args || args.peerGreAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerGreAddress'");
            }
            if ((!args || args.tgwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tgwName'");
            }
            resourceInputs["bgpInsideCidrs"] = args ? args.bgpInsideCidrs : undefined;
            resourceInputs["connectAttachmentId"] = args ? args.connectAttachmentId : undefined;
            resourceInputs["connectPeerName"] = args ? args.connectPeerName : undefined;
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["peerAsNumber"] = args ? args.peerAsNumber : undefined;
            resourceInputs["peerGreAddress"] = args ? args.peerGreAddress : undefined;
            resourceInputs["tgwGreAddress"] = args ? args.tgwGreAddress : undefined;
            resourceInputs["tgwName"] = args ? args.tgwName : undefined;
            resourceInputs["connectPeerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAwsTgwConnectPeer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAwsTgwConnectPeer resources.
 */
export interface AviatrixAwsTgwConnectPeerState {
    /**
     * Set of BGP Inside CIDR Blocks.
     */
    bgpInsideCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Connect Attachment ID.
     */
    connectAttachmentId?: pulumi.Input<string>;
    /**
     * Connect Peer ID.
     */
    connectPeerId?: pulumi.Input<string>;
    /**
     * Connect Peer Name.
     */
    connectPeerName?: pulumi.Input<string>;
    /**
     * AWS TGW Connect connection name.
     */
    connectionName?: pulumi.Input<string>;
    /**
     * Peer AS Number.
     */
    peerAsNumber?: pulumi.Input<string>;
    /**
     * Peer GRE IP Address.
     */
    peerGreAddress?: pulumi.Input<string>;
    /**
     * AWS TGW GRE IP Address.
     */
    tgwGreAddress?: pulumi.Input<string>;
    /**
     * AWS TGW Name.
     */
    tgwName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAwsTgwConnectPeer resource.
 */
export interface AviatrixAwsTgwConnectPeerArgs {
    /**
     * Set of BGP Inside CIDR Blocks.
     */
    bgpInsideCidrs: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Connect Attachment ID.
     */
    connectAttachmentId: pulumi.Input<string>;
    /**
     * Connect Peer Name.
     */
    connectPeerName: pulumi.Input<string>;
    /**
     * AWS TGW Connect connection name.
     */
    connectionName: pulumi.Input<string>;
    /**
     * Peer AS Number.
     */
    peerAsNumber: pulumi.Input<string>;
    /**
     * Peer GRE IP Address.
     */
    peerGreAddress: pulumi.Input<string>;
    /**
     * AWS TGW GRE IP Address.
     */
    tgwGreAddress?: pulumi.Input<string>;
    /**
     * AWS TGW Name.
     */
    tgwName: pulumi.Input<string>;
}
