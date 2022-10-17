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
    public sealed class AviatrixGatewaySnatSnatPolicy
    {
        public readonly bool? ApplyRouteEntry;
        public readonly string? Connection;
        public readonly string? DstCidr;
        public readonly string? DstPort;
        public readonly string? ExcludeRtb;
        public readonly string? Interface;
        public readonly string? Mark;
        public readonly string? Protocol;
        public readonly string? SnatIps;
        public readonly string? SnatPort;
        public readonly string? SrcCidr;
        public readonly string? SrcPort;

        [OutputConstructor]
        private AviatrixGatewaySnatSnatPolicy(
            bool? applyRouteEntry,

            string? connection,

            string? dstCidr,

            string? dstPort,

            string? excludeRtb,

            string? @interface,

            string? mark,

            string? protocol,

            string? snatIps,

            string? snatPort,

            string? srcCidr,

            string? srcPort)
        {
            ApplyRouteEntry = applyRouteEntry;
            Connection = connection;
            DstCidr = dstCidr;
            DstPort = dstPort;
            ExcludeRtb = excludeRtb;
            Interface = @interface;
            Mark = mark;
            Protocol = protocol;
            SnatIps = snatIps;
            SnatPort = snatPort;
            SrcCidr = srcCidr;
            SrcPort = srcPort;
        }
    }
}
