// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Inputs
{

    public sealed class AviatrixMicrosegPolicyListPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("dstAppDomains", required: true)]
        private InputList<string>? _dstAppDomains;
        public InputList<string> DstAppDomains
        {
            get => _dstAppDomains ?? (_dstAppDomains = new InputList<string>());
            set => _dstAppDomains = value;
        }

        [Input("logging")]
        public Input<bool>? Logging { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("portRanges")]
        private InputList<Inputs.AviatrixMicrosegPolicyListPolicyPortRangeGetArgs>? _portRanges;
        public InputList<Inputs.AviatrixMicrosegPolicyListPolicyPortRangeGetArgs> PortRanges
        {
            get => _portRanges ?? (_portRanges = new InputList<Inputs.AviatrixMicrosegPolicyListPolicyPortRangeGetArgs>());
            set => _portRanges = value;
        }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("srcAppDomains", required: true)]
        private InputList<string>? _srcAppDomains;
        public InputList<string> SrcAppDomains
        {
            get => _srcAppDomains ?? (_srcAppDomains = new InputList<string>());
            set => _srcAppDomains = value;
        }

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        [Input("watch")]
        public Input<bool>? Watch { get; set; }

        public AviatrixMicrosegPolicyListPolicyGetArgs()
        {
        }
        public static new AviatrixMicrosegPolicyListPolicyGetArgs Empty => new AviatrixMicrosegPolicyListPolicyGetArgs();
    }
}
