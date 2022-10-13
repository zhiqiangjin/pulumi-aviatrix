// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain")]
    public partial class AviatrixAwsTgwNetworkDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set to true if the network domain is an aviatrix firewall domain.
        /// </summary>
        [Output("aviatrixFirewall")]
        public Output<bool?> AviatrixFirewall { get; private set; } = null!;

        /// <summary>
        /// Network domain name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set to true if the network domain is a native egress domain.
        /// </summary>
        [Output("nativeEgress")]
        public Output<bool?> NativeEgress { get; private set; } = null!;

        /// <summary>
        /// Set to true if the network domain is a native firewall domain.
        /// </summary>
        [Output("nativeFirewall")]
        public Output<bool?> NativeFirewall { get; private set; } = null!;

        /// <summary>
        /// AWS TGW name.
        /// </summary>
        [Output("tgwName")]
        public Output<string> TgwName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsTgwNetworkDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsTgwNetworkDomain(string name, AviatrixAwsTgwNetworkDomainArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain", name, args ?? new AviatrixAwsTgwNetworkDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsTgwNetworkDomain(string name, Input<string> id, AviatrixAwsTgwNetworkDomainState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/astipkovits/pulumi-aviatrix/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixAwsTgwNetworkDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsTgwNetworkDomain Get(string name, Input<string> id, AviatrixAwsTgwNetworkDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsTgwNetworkDomain(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsTgwNetworkDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true if the network domain is an aviatrix firewall domain.
        /// </summary>
        [Input("aviatrixFirewall")]
        public Input<bool>? AviatrixFirewall { get; set; }

        /// <summary>
        /// Network domain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to true if the network domain is a native egress domain.
        /// </summary>
        [Input("nativeEgress")]
        public Input<bool>? NativeEgress { get; set; }

        /// <summary>
        /// Set to true if the network domain is a native firewall domain.
        /// </summary>
        [Input("nativeFirewall")]
        public Input<bool>? NativeFirewall { get; set; }

        /// <summary>
        /// AWS TGW name.
        /// </summary>
        [Input("tgwName", required: true)]
        public Input<string> TgwName { get; set; } = null!;

        public AviatrixAwsTgwNetworkDomainArgs()
        {
        }
        public static new AviatrixAwsTgwNetworkDomainArgs Empty => new AviatrixAwsTgwNetworkDomainArgs();
    }

    public sealed class AviatrixAwsTgwNetworkDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true if the network domain is an aviatrix firewall domain.
        /// </summary>
        [Input("aviatrixFirewall")]
        public Input<bool>? AviatrixFirewall { get; set; }

        /// <summary>
        /// Network domain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to true if the network domain is a native egress domain.
        /// </summary>
        [Input("nativeEgress")]
        public Input<bool>? NativeEgress { get; set; }

        /// <summary>
        /// Set to true if the network domain is a native firewall domain.
        /// </summary>
        [Input("nativeFirewall")]
        public Input<bool>? NativeFirewall { get; set; }

        /// <summary>
        /// AWS TGW name.
        /// </summary>
        [Input("tgwName")]
        public Input<string>? TgwName { get; set; }

        public AviatrixAwsTgwNetworkDomainState()
        {
        }
        public static new AviatrixAwsTgwNetworkDomainState Empty => new AviatrixAwsTgwNetworkDomainState();
    }
}
