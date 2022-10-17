// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Outputs
{

    [OutputType]
    public sealed class AviatrixGatewayDnatDnatPolicy
    {
        public readonly bool? ApplyRouteEntry;
        public readonly string? Connection;
        public readonly string? DnatIps;
        public readonly string? DnatPort;
        public readonly string? DstCidr;
        public readonly string? DstPort;
        public readonly string? ExcludeRtb;
        public readonly string? Interface;
        public readonly string? Mark;
        public readonly string? Protocol;
        public readonly string? SrcCidr;
        public readonly string? SrcPort;

        [OutputConstructor]
        private AviatrixGatewayDnatDnatPolicy(
            bool? applyRouteEntry,

            string? connection,

            string? dnatIps,

            string? dnatPort,

            string? dstCidr,

            string? dstPort,

            string? excludeRtb,

            string? @interface,

            string? mark,

            string? protocol,

            string? srcCidr,

            string? srcPort)
        {
            ApplyRouteEntry = applyRouteEntry;
            Connection = connection;
            DnatIps = dnatIps;
            DnatPort = dnatPort;
            DstCidr = dstCidr;
            DstPort = dstPort;
            ExcludeRtb = excludeRtb;
            Interface = @interface;
            Mark = mark;
            Protocol = protocol;
            SrcCidr = srcCidr;
            SrcPort = srcPort;
        }
    }
}
