// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Inputs
{

    public sealed class AviatrixFirewallPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dstIp", required: true)]
        public Input<string> DstIp { get; set; } = null!;

        [Input("logEnabled")]
        public Input<bool>? LogEnabled { get; set; }

        [Input("port", required: true)]
        public Input<string> Port { get; set; } = null!;

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("srcIp", required: true)]
        public Input<string> SrcIp { get; set; } = null!;

        public AviatrixFirewallPolicyArgs()
        {
        }
        public static new AviatrixFirewallPolicyArgs Empty => new AviatrixFirewallPolicyArgs();
    }
}
