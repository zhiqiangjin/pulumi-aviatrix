// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixGeoVpn:AviatrixGeoVpn")]
    public partial class AviatrixGeoVpn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
        /// </summary>
        [Output("cloudType")]
        public Output<int> CloudType { get; private set; } = null!;

        /// <summary>
        /// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// List of ELB names to attach to this Geo VPN name.
        /// </summary>
        [Output("elbDnsNames")]
        public Output<ImmutableArray<string>> ElbDnsNames { get; private set; } = null!;

        /// <summary>
        /// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixGeoVpn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixGeoVpn(string name, AviatrixGeoVpnArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixGeoVpn:AviatrixGeoVpn", name, args ?? new AviatrixGeoVpnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixGeoVpn(string name, Input<string> id, AviatrixGeoVpnState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixGeoVpn:AviatrixGeoVpn", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/astipkovits/pulumi-aviatrix/raw/main/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AviatrixGeoVpn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixGeoVpn Get(string name, Input<string> id, AviatrixGeoVpnState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixGeoVpn(name, id, state, options);
        }
    }

    public sealed class AviatrixGeoVpnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
        /// </summary>
        [Input("cloudType", required: true)]
        public Input<int> CloudType { get; set; } = null!;

        /// <summary>
        /// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        [Input("elbDnsNames", required: true)]
        private InputList<string>? _elbDnsNames;

        /// <summary>
        /// List of ELB names to attach to this Geo VPN name.
        /// </summary>
        public InputList<string> ElbDnsNames
        {
            get => _elbDnsNames ?? (_elbDnsNames = new InputList<string>());
            set => _elbDnsNames = value;
        }

        /// <summary>
        /// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public AviatrixGeoVpnArgs()
        {
        }
        public static new AviatrixGeoVpnArgs Empty => new AviatrixGeoVpnArgs();
    }

    public sealed class AviatrixGeoVpnState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("elbDnsNames")]
        private InputList<string>? _elbDnsNames;

        /// <summary>
        /// List of ELB names to attach to this Geo VPN name.
        /// </summary>
        public InputList<string> ElbDnsNames
        {
            get => _elbDnsNames ?? (_elbDnsNames = new InputList<string>());
            set => _elbDnsNames = value;
        }

        /// <summary>
        /// The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public AviatrixGeoVpnState()
        {
        }
        public static new AviatrixGeoVpnState Empty => new AviatrixGeoVpnState();
    }
}