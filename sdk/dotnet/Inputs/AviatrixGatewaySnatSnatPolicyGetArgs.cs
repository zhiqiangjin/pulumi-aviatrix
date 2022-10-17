// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Inputs
{

    public sealed class AviatrixGatewaySnatSnatPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("applyRouteEntry")]
        public Input<bool>? ApplyRouteEntry { get; set; }

        [Input("connection")]
        public Input<string>? Connection { get; set; }

        [Input("dstCidr")]
        public Input<string>? DstCidr { get; set; }

        [Input("dstPort")]
        public Input<string>? DstPort { get; set; }

        [Input("excludeRtb")]
        public Input<string>? ExcludeRtb { get; set; }

        [Input("interface")]
        public Input<string>? Interface { get; set; }

        [Input("mark")]
        public Input<string>? Mark { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("snatIps")]
        public Input<string>? SnatIps { get; set; }

        [Input("snatPort")]
        public Input<string>? SnatPort { get; set; }

        [Input("srcCidr")]
        public Input<string>? SrcCidr { get; set; }

        [Input("srcPort")]
        public Input<string>? SrcPort { get; set; }

        public AviatrixGatewaySnatSnatPolicyGetArgs()
        {
        }
        public static new AviatrixGatewaySnatSnatPolicyGetArgs Empty => new AviatrixGatewaySnatSnatPolicyGetArgs();
    }
}
