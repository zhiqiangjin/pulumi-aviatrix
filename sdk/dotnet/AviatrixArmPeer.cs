// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixArmPeer:AviatrixArmPeer")]
    public partial class AviatrixArmPeer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName1")]
        public Output<string> AccountName1 { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName2")]
        public Output<string> AccountName2 { get; private set; } = null!;

        /// <summary>
        /// List of VNet CIDR of vnet_name_resource_group1.
        /// </summary>
        [Output("vnetCidr1s")]
        public Output<ImmutableArray<string>> VnetCidr1s { get; private set; } = null!;

        /// <summary>
        /// List of VNet CIDR of vnet_name_resource_group2.
        /// </summary>
        [Output("vnetCidr2s")]
        public Output<ImmutableArray<string>> VnetCidr2s { get; private set; } = null!;

        /// <summary>
        /// VNet-Name of Azure cloud.
        /// </summary>
        [Output("vnetNameResourceGroup1")]
        public Output<string> VnetNameResourceGroup1 { get; private set; } = null!;

        /// <summary>
        /// VNet-Name of Azure cloud.
        /// </summary>
        [Output("vnetNameResourceGroup2")]
        public Output<string> VnetNameResourceGroup2 { get; private set; } = null!;

        /// <summary>
        /// Region of Azure cloud.
        /// </summary>
        [Output("vnetReg1")]
        public Output<string> VnetReg1 { get; private set; } = null!;

        /// <summary>
        /// Region of Azure cloud.
        /// </summary>
        [Output("vnetReg2")]
        public Output<string> VnetReg2 { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixArmPeer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixArmPeer(string name, AviatrixArmPeerArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixArmPeer:AviatrixArmPeer", name, args ?? new AviatrixArmPeerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixArmPeer(string name, Input<string> id, AviatrixArmPeerState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixArmPeer:AviatrixArmPeer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixArmPeer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixArmPeer Get(string name, Input<string> id, AviatrixArmPeerState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixArmPeer(name, id, state, options);
        }
    }

    public sealed class AviatrixArmPeerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName1", required: true)]
        public Input<string> AccountName1 { get; set; } = null!;

        /// <summary>
        /// This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName2", required: true)]
        public Input<string> AccountName2 { get; set; } = null!;

        /// <summary>
        /// VNet-Name of Azure cloud.
        /// </summary>
        [Input("vnetNameResourceGroup1", required: true)]
        public Input<string> VnetNameResourceGroup1 { get; set; } = null!;

        /// <summary>
        /// VNet-Name of Azure cloud.
        /// </summary>
        [Input("vnetNameResourceGroup2", required: true)]
        public Input<string> VnetNameResourceGroup2 { get; set; } = null!;

        /// <summary>
        /// Region of Azure cloud.
        /// </summary>
        [Input("vnetReg1", required: true)]
        public Input<string> VnetReg1 { get; set; } = null!;

        /// <summary>
        /// Region of Azure cloud.
        /// </summary>
        [Input("vnetReg2", required: true)]
        public Input<string> VnetReg2 { get; set; } = null!;

        public AviatrixArmPeerArgs()
        {
        }
        public static new AviatrixArmPeerArgs Empty => new AviatrixArmPeerArgs();
    }

    public sealed class AviatrixArmPeerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName1")]
        public Input<string>? AccountName1 { get; set; }

        /// <summary>
        /// This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName2")]
        public Input<string>? AccountName2 { get; set; }

        [Input("vnetCidr1s")]
        private InputList<string>? _vnetCidr1s;

        /// <summary>
        /// List of VNet CIDR of vnet_name_resource_group1.
        /// </summary>
        public InputList<string> VnetCidr1s
        {
            get => _vnetCidr1s ?? (_vnetCidr1s = new InputList<string>());
            set => _vnetCidr1s = value;
        }

        [Input("vnetCidr2s")]
        private InputList<string>? _vnetCidr2s;

        /// <summary>
        /// List of VNet CIDR of vnet_name_resource_group2.
        /// </summary>
        public InputList<string> VnetCidr2s
        {
            get => _vnetCidr2s ?? (_vnetCidr2s = new InputList<string>());
            set => _vnetCidr2s = value;
        }

        /// <summary>
        /// VNet-Name of Azure cloud.
        /// </summary>
        [Input("vnetNameResourceGroup1")]
        public Input<string>? VnetNameResourceGroup1 { get; set; }

        /// <summary>
        /// VNet-Name of Azure cloud.
        /// </summary>
        [Input("vnetNameResourceGroup2")]
        public Input<string>? VnetNameResourceGroup2 { get; set; }

        /// <summary>
        /// Region of Azure cloud.
        /// </summary>
        [Input("vnetReg1")]
        public Input<string>? VnetReg1 { get; set; }

        /// <summary>
        /// Region of Azure cloud.
        /// </summary>
        [Input("vnetReg2")]
        public Input<string>? VnetReg2 { get; set; }

        public AviatrixArmPeerState()
        {
        }
        public static new AviatrixArmPeerState Empty => new AviatrixArmPeerState();
    }
}
