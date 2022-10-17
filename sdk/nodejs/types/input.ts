// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface AviatrixAppDomainSelector {
    matchExpressions: pulumi.Input<pulumi.Input<inputs.AviatrixAppDomainSelectorMatchExpression>[]>;
}

export interface AviatrixAppDomainSelectorMatchExpression {
    accountId?: pulumi.Input<string>;
    accountName?: pulumi.Input<string>;
    cidr?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    resId?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    type?: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}

export interface AviatrixAwsTgwSecurityDomain {
    /**
     * @deprecated Please set `manage_vpc_attachment` to false, and use the standalone aviatrix_aws_tgw_vpc_attachment resource instead.
     */
    attachedVpcs?: pulumi.Input<pulumi.Input<inputs.AviatrixAwsTgwSecurityDomainAttachedVpc>[]>;
    aviatrixFirewall?: pulumi.Input<boolean>;
    connectedDomains?: pulumi.Input<pulumi.Input<string>[]>;
    nativeEgress?: pulumi.Input<boolean>;
    nativeFirewall?: pulumi.Input<boolean>;
    securityDomainName: pulumi.Input<string>;
}

export interface AviatrixAwsTgwSecurityDomainAttachedVpc {
    customizedRouteAdvertisement?: pulumi.Input<string>;
    customizedRoutes?: pulumi.Input<string>;
    disableLocalRoutePropagation?: pulumi.Input<boolean>;
    routeTables?: pulumi.Input<string>;
    subnets?: pulumi.Input<string>;
    vpcAccountName: pulumi.Input<string>;
    vpcId: pulumi.Input<string>;
    vpcRegion: pulumi.Input<string>;
}

export interface AviatrixAwsTgwVpnConnVpnTunnelData {
    lastStatusChangeTime?: pulumi.Input<string>;
    routeCount?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    statusMessage?: pulumi.Input<string>;
    tgwAsn?: pulumi.Input<string>;
    tunnelName?: pulumi.Input<string>;
    vpnInsideAddress?: pulumi.Input<string>;
    vpnOutsideAddress?: pulumi.Input<string>;
}

export interface AviatrixFirenetFirewallInstanceAssociation {
    attached?: pulumi.Input<boolean>;
    egressInterface?: pulumi.Input<string>;
    firenetGwName: pulumi.Input<string>;
    firewallName?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    lanInterface?: pulumi.Input<string>;
    managementInterface?: pulumi.Input<string>;
    vendorType?: pulumi.Input<string>;
}

export interface AviatrixFirewallPolicy {
    action: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    dstIp: pulumi.Input<string>;
    logEnabled?: pulumi.Input<boolean>;
    port: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    srcIp: pulumi.Input<string>;
}

export interface AviatrixFirewallTagCidrList {
    cidr: pulumi.Input<string>;
    cidrTagName: pulumi.Input<string>;
}

export interface AviatrixFqdnDomainName {
    action?: pulumi.Input<string>;
    fqdn: pulumi.Input<string>;
    port: pulumi.Input<string>;
    proto: pulumi.Input<string>;
}

export interface AviatrixFqdnGwFilterTagList {
    gwName: pulumi.Input<string>;
    sourceIpLists?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface AviatrixGatewayDnatConnectionPolicy {
    applyRouteEntry?: pulumi.Input<boolean>;
    connection?: pulumi.Input<string>;
    dnatIps?: pulumi.Input<string>;
    dnatPort?: pulumi.Input<string>;
    dstCidr?: pulumi.Input<string>;
    dstPort?: pulumi.Input<string>;
    excludeRtb?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    mark?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    srcCidr?: pulumi.Input<string>;
    srcPort?: pulumi.Input<string>;
}

export interface AviatrixGatewayDnatDnatPolicy {
    applyRouteEntry?: pulumi.Input<boolean>;
    connection?: pulumi.Input<string>;
    dnatIps?: pulumi.Input<string>;
    dnatPort?: pulumi.Input<string>;
    dstCidr?: pulumi.Input<string>;
    dstPort?: pulumi.Input<string>;
    excludeRtb?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    mark?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    srcCidr?: pulumi.Input<string>;
    srcPort?: pulumi.Input<string>;
}

export interface AviatrixGatewayDnatInterfacePolicy {
    applyRouteEntry?: pulumi.Input<boolean>;
    connection?: pulumi.Input<string>;
    dnatIps?: pulumi.Input<string>;
    dnatPort?: pulumi.Input<string>;
    dstCidr?: pulumi.Input<string>;
    dstPort?: pulumi.Input<string>;
    excludeRtb?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    mark?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    srcCidr?: pulumi.Input<string>;
    srcPort?: pulumi.Input<string>;
}

export interface AviatrixGatewaySnatConnectionPolicy {
    applyRouteEntry?: pulumi.Input<boolean>;
    connection?: pulumi.Input<string>;
    dstCidr?: pulumi.Input<string>;
    dstPort?: pulumi.Input<string>;
    excludeRtb?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    mark?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    snatIps?: pulumi.Input<string>;
    snatPort?: pulumi.Input<string>;
    srcCidr?: pulumi.Input<string>;
    srcPort?: pulumi.Input<string>;
}

export interface AviatrixGatewaySnatInterfacePolicy {
    applyRouteEntry?: pulumi.Input<boolean>;
    connection?: pulumi.Input<string>;
    dstCidr?: pulumi.Input<string>;
    dstPort?: pulumi.Input<string>;
    excludeRtb?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    mark?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    snatIps?: pulumi.Input<string>;
    snatPort?: pulumi.Input<string>;
    srcCidr?: pulumi.Input<string>;
    srcPort?: pulumi.Input<string>;
}

export interface AviatrixGatewaySnatSnatPolicy {
    applyRouteEntry?: pulumi.Input<boolean>;
    connection?: pulumi.Input<string>;
    dstCidr?: pulumi.Input<string>;
    dstPort?: pulumi.Input<string>;
    excludeRtb?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    mark?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    snatIps?: pulumi.Input<string>;
    snatPort?: pulumi.Input<string>;
    srcCidr?: pulumi.Input<string>;
    srcPort?: pulumi.Input<string>;
}

export interface AviatrixMicrosegPolicyListPolicy {
    action: pulumi.Input<string>;
    dstAppDomains: pulumi.Input<pulumi.Input<string>[]>;
    logging?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    portRanges?: pulumi.Input<pulumi.Input<inputs.AviatrixMicrosegPolicyListPolicyPortRange>[]>;
    priority?: pulumi.Input<number>;
    protocol: pulumi.Input<string>;
    srcAppDomains: pulumi.Input<pulumi.Input<string>[]>;
    uuid?: pulumi.Input<string>;
    watch?: pulumi.Input<boolean>;
}

export interface AviatrixMicrosegPolicyListPolicyPortRange {
    hi?: pulumi.Input<number>;
    lo: pulumi.Input<number>;
}

export interface AviatrixPrivateModeLbProxy {
    instanceId: pulumi.Input<string>;
    vpcId: pulumi.Input<string>;
}

export interface AviatrixSite2CloudCaCertTagCaCertificate {
    certContent: pulumi.Input<string>;
    commonName?: pulumi.Input<string>;
    expirationTime?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    issuerName?: pulumi.Input<string>;
    uniqueSerial?: pulumi.Input<string>;
}

export interface AviatrixTransitGatewayBgpLanInterface {
    subnet: pulumi.Input<string>;
    vpcId: pulumi.Input<string>;
}

export interface AviatrixTransitGatewayHaBgpLanInterface {
    subnet: pulumi.Input<string>;
    vpcId: pulumi.Input<string>;
}

export interface AviatrixVpcPrivateSubnet {
    cidr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
}

export interface AviatrixVpcPublicSubnet {
    cidr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
}

export interface AviatrixVpcSubnet {
    cidr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
}

export interface AviatrixVpnProfilePolicy {
    action: pulumi.Input<string>;
    port: pulumi.Input<string>;
    proto: pulumi.Input<string>;
    target: pulumi.Input<string>;
}

export interface GetAviatrixFirenetFirewallInstanceAssociation {
    attached?: boolean;
    egressInterface?: string;
    firenetGwName?: string;
    firewallName?: string;
    instanceId?: string;
    lanInterface?: string;
    managementInterface?: string;
    vendorType?: string;
}

export interface GetAviatrixFirenetFirewallInstanceAssociationArgs {
    attached?: pulumi.Input<boolean>;
    egressInterface?: pulumi.Input<string>;
    firenetGwName?: pulumi.Input<string>;
    firewallName?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    lanInterface?: pulumi.Input<string>;
    managementInterface?: pulumi.Input<string>;
    vendorType?: pulumi.Input<string>;
}

export interface ProviderIgnoreTags {
    keyPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    keys?: pulumi.Input<pulumi.Input<string>[]>;
}
export namespace config {
}
