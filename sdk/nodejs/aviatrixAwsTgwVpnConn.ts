// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class AviatrixAwsTgwVpnConn extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAwsTgwVpnConn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAwsTgwVpnConnState, opts?: pulumi.CustomResourceOptions): AviatrixAwsTgwVpnConn {
        return new AviatrixAwsTgwVpnConn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAwsTgwVpnConn:AviatrixAwsTgwVpnConn';

    /**
     * Returns true if the given object is an instance of AviatrixAwsTgwVpnConn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAwsTgwVpnConn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAwsTgwVpnConn.__pulumiType;
    }

    /**
     * Unique name of the connection.
     */
    public readonly connectionName!: pulumi.Output<string>;
    /**
     * Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
     * static VPN connection. Default value: 'dynamic'.
     */
    public readonly connectionType!: pulumi.Output<string | undefined>;
    /**
     * Enable Global Acceleration.
     */
    public readonly enableGlobalAcceleration!: pulumi.Output<boolean | undefined>;
    /**
     * Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
     */
    public readonly enableLearnedCidrsApproval!: pulumi.Output<boolean | undefined>;
    /**
     * Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
     */
    public readonly insideIpCidrTun1!: pulumi.Output<string | undefined>;
    /**
     * Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
     */
    public readonly insideIpCidrTun2!: pulumi.Output<string | undefined>;
    /**
     * Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
     */
    public readonly preSharedKeyTun1!: pulumi.Output<string | undefined>;
    /**
     * Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
     */
    public readonly preSharedKeyTun2!: pulumi.Output<string | undefined>;
    /**
     * Public IP address. Example: '40.0.0.0'.
     */
    public readonly publicIp!: pulumi.Output<string>;
    /**
     * AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
     */
    public readonly remoteAsNumber!: pulumi.Output<string | undefined>;
    /**
     * Remote CIDRs joined as a string with ','. Required for a static VPN connection.
     */
    public readonly remoteCidr!: pulumi.Output<string | undefined>;
    /**
     * The name of a route domain, to which the vpn will be attached.
     */
    public readonly routeDomainName!: pulumi.Output<string>;
    /**
     * This parameter represents the name of an AWS TGW.
     */
    public readonly tgwName!: pulumi.Output<string>;
    /**
     * ID of the vpn connection.
     */
    public /*out*/ readonly vpnId!: pulumi.Output<string>;
    /**
     * VPN tunnel data.
     */
    public /*out*/ readonly vpnTunnelDatas!: pulumi.Output<outputs.AviatrixAwsTgwVpnConnVpnTunnelData[]>;

    /**
     * Create a AviatrixAwsTgwVpnConn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAwsTgwVpnConnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAwsTgwVpnConnArgs | AviatrixAwsTgwVpnConnState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAwsTgwVpnConnState | undefined;
            resourceInputs["connectionName"] = state ? state.connectionName : undefined;
            resourceInputs["connectionType"] = state ? state.connectionType : undefined;
            resourceInputs["enableGlobalAcceleration"] = state ? state.enableGlobalAcceleration : undefined;
            resourceInputs["enableLearnedCidrsApproval"] = state ? state.enableLearnedCidrsApproval : undefined;
            resourceInputs["insideIpCidrTun1"] = state ? state.insideIpCidrTun1 : undefined;
            resourceInputs["insideIpCidrTun2"] = state ? state.insideIpCidrTun2 : undefined;
            resourceInputs["preSharedKeyTun1"] = state ? state.preSharedKeyTun1 : undefined;
            resourceInputs["preSharedKeyTun2"] = state ? state.preSharedKeyTun2 : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["remoteAsNumber"] = state ? state.remoteAsNumber : undefined;
            resourceInputs["remoteCidr"] = state ? state.remoteCidr : undefined;
            resourceInputs["routeDomainName"] = state ? state.routeDomainName : undefined;
            resourceInputs["tgwName"] = state ? state.tgwName : undefined;
            resourceInputs["vpnId"] = state ? state.vpnId : undefined;
            resourceInputs["vpnTunnelDatas"] = state ? state.vpnTunnelDatas : undefined;
        } else {
            const args = argsOrState as AviatrixAwsTgwVpnConnArgs | undefined;
            if ((!args || args.connectionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionName'");
            }
            if ((!args || args.publicIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicIp'");
            }
            if ((!args || args.routeDomainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeDomainName'");
            }
            if ((!args || args.tgwName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tgwName'");
            }
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["connectionType"] = args ? args.connectionType : undefined;
            resourceInputs["enableGlobalAcceleration"] = args ? args.enableGlobalAcceleration : undefined;
            resourceInputs["enableLearnedCidrsApproval"] = args ? args.enableLearnedCidrsApproval : undefined;
            resourceInputs["insideIpCidrTun1"] = args ? args.insideIpCidrTun1 : undefined;
            resourceInputs["insideIpCidrTun2"] = args ? args.insideIpCidrTun2 : undefined;
            resourceInputs["preSharedKeyTun1"] = args ? args.preSharedKeyTun1 : undefined;
            resourceInputs["preSharedKeyTun2"] = args ? args.preSharedKeyTun2 : undefined;
            resourceInputs["publicIp"] = args ? args.publicIp : undefined;
            resourceInputs["remoteAsNumber"] = args ? args.remoteAsNumber : undefined;
            resourceInputs["remoteCidr"] = args ? args.remoteCidr : undefined;
            resourceInputs["routeDomainName"] = args ? args.routeDomainName : undefined;
            resourceInputs["tgwName"] = args ? args.tgwName : undefined;
            resourceInputs["vpnId"] = undefined /*out*/;
            resourceInputs["vpnTunnelDatas"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAwsTgwVpnConn.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAwsTgwVpnConn resources.
 */
export interface AviatrixAwsTgwVpnConnState {
    /**
     * Unique name of the connection.
     */
    connectionName?: pulumi.Input<string>;
    /**
     * Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
     * static VPN connection. Default value: 'dynamic'.
     */
    connectionType?: pulumi.Input<string>;
    /**
     * Enable Global Acceleration.
     */
    enableGlobalAcceleration?: pulumi.Input<boolean>;
    /**
     * Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
     */
    enableLearnedCidrsApproval?: pulumi.Input<boolean>;
    /**
     * Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
     */
    insideIpCidrTun1?: pulumi.Input<string>;
    /**
     * Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
     */
    insideIpCidrTun2?: pulumi.Input<string>;
    /**
     * Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
     */
    preSharedKeyTun1?: pulumi.Input<string>;
    /**
     * Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
     */
    preSharedKeyTun2?: pulumi.Input<string>;
    /**
     * Public IP address. Example: '40.0.0.0'.
     */
    publicIp?: pulumi.Input<string>;
    /**
     * AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
     */
    remoteAsNumber?: pulumi.Input<string>;
    /**
     * Remote CIDRs joined as a string with ','. Required for a static VPN connection.
     */
    remoteCidr?: pulumi.Input<string>;
    /**
     * The name of a route domain, to which the vpn will be attached.
     */
    routeDomainName?: pulumi.Input<string>;
    /**
     * This parameter represents the name of an AWS TGW.
     */
    tgwName?: pulumi.Input<string>;
    /**
     * ID of the vpn connection.
     */
    vpnId?: pulumi.Input<string>;
    /**
     * VPN tunnel data.
     */
    vpnTunnelDatas?: pulumi.Input<pulumi.Input<inputs.AviatrixAwsTgwVpnConnVpnTunnelData>[]>;
}

/**
 * The set of arguments for constructing a AviatrixAwsTgwVpnConn resource.
 */
export interface AviatrixAwsTgwVpnConnArgs {
    /**
     * Unique name of the connection.
     */
    connectionName: pulumi.Input<string>;
    /**
     * Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
     * static VPN connection. Default value: 'dynamic'.
     */
    connectionType?: pulumi.Input<string>;
    /**
     * Enable Global Acceleration.
     */
    enableGlobalAcceleration?: pulumi.Input<boolean>;
    /**
     * Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
     */
    enableLearnedCidrsApproval?: pulumi.Input<boolean>;
    /**
     * Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
     */
    insideIpCidrTun1?: pulumi.Input<string>;
    /**
     * Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
     */
    insideIpCidrTun2?: pulumi.Input<string>;
    /**
     * Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
     */
    preSharedKeyTun1?: pulumi.Input<string>;
    /**
     * Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
     */
    preSharedKeyTun2?: pulumi.Input<string>;
    /**
     * Public IP address. Example: '40.0.0.0'.
     */
    publicIp: pulumi.Input<string>;
    /**
     * AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
     */
    remoteAsNumber?: pulumi.Input<string>;
    /**
     * Remote CIDRs joined as a string with ','. Required for a static VPN connection.
     */
    remoteCidr?: pulumi.Input<string>;
    /**
     * The name of a route domain, to which the vpn will be attached.
     */
    routeDomainName: pulumi.Input<string>;
    /**
     * This parameter represents the name of an AWS TGW.
     */
    tgwName: pulumi.Input<string>;
}
