// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc")]
    public partial class AviatrixSpokeVpc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Cloud instance ID.
        /// </summary>
        [Output("cloudInstanceId")]
        public Output<string> CloudInstanceId { get; private set; } = null!;

        /// <summary>
        /// Type of cloud service provider.
        /// </summary>
        [Output("cloudType")]
        public Output<int> CloudType { get; private set; } = null!;

        /// <summary>
        /// Specify whether enabling NAT feature on the gateway or not.
        /// </summary>
        [Output("enableNat")]
        public Output<string?> EnableNat { get; private set; } = null!;

        /// <summary>
        /// Name of the gateway which is going to be created.
        /// </summary>
        [Output("gwName")]
        public Output<string> GwName { get; private set; } = null!;

        /// <summary>
        /// HA Gateway Size.
        /// </summary>
        [Output("haGwSize")]
        public Output<string?> HaGwSize { get; private set; } = null!;

        /// <summary>
        /// HA Subnet. Required if enabling HA for AWS/Azure.
        /// </summary>
        [Output("haSubnet")]
        public Output<string?> HaSubnet { get; private set; } = null!;

        /// <summary>
        /// HA Zone. Required if enabling HA for GCP.
        /// </summary>
        [Output("haZone")]
        public Output<string?> HaZone { get; private set; } = null!;

        /// <summary>
        /// Set to 'enabled' if this feature is desired.
        /// </summary>
        [Output("singleAzHa")]
        public Output<string?> SingleAzHa { get; private set; } = null!;

        /// <summary>
        /// Public Subnet Info.
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Instance tag of cloud provider.
        /// </summary>
        [Output("tagLists")]
        public Output<ImmutableArray<string>> TagLists { get; private set; } = null!;

        /// <summary>
        /// Specify the transit Gateway.
        /// </summary>
        [Output("transitGw")]
        public Output<string?> TransitGw { get; private set; } = null!;

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// Region of cloud provider.
        /// </summary>
        [Output("vpcReg")]
        public Output<string> VpcReg { get; private set; } = null!;

        /// <summary>
        /// Size of the gateway instance.
        /// </summary>
        [Output("vpcSize")]
        public Output<string> VpcSize { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixSpokeVpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixSpokeVpc(string name, AviatrixSpokeVpcArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc", name, args ?? new AviatrixSpokeVpcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixSpokeVpc(string name, Input<string> id, AviatrixSpokeVpcState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixSpokeVpc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixSpokeVpc Get(string name, Input<string> id, AviatrixSpokeVpcState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixSpokeVpc(name, id, state, options);
        }
    }

    public sealed class AviatrixSpokeVpcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Type of cloud service provider.
        /// </summary>
        [Input("cloudType", required: true)]
        public Input<int> CloudType { get; set; } = null!;

        /// <summary>
        /// Specify whether enabling NAT feature on the gateway or not.
        /// </summary>
        [Input("enableNat")]
        public Input<string>? EnableNat { get; set; }

        /// <summary>
        /// Name of the gateway which is going to be created.
        /// </summary>
        [Input("gwName", required: true)]
        public Input<string> GwName { get; set; } = null!;

        /// <summary>
        /// HA Gateway Size.
        /// </summary>
        [Input("haGwSize")]
        public Input<string>? HaGwSize { get; set; }

        /// <summary>
        /// HA Subnet. Required if enabling HA for AWS/Azure.
        /// </summary>
        [Input("haSubnet")]
        public Input<string>? HaSubnet { get; set; }

        /// <summary>
        /// HA Zone. Required if enabling HA for GCP.
        /// </summary>
        [Input("haZone")]
        public Input<string>? HaZone { get; set; }

        /// <summary>
        /// Set to 'enabled' if this feature is desired.
        /// </summary>
        [Input("singleAzHa")]
        public Input<string>? SingleAzHa { get; set; }

        /// <summary>
        /// Public Subnet Info.
        /// </summary>
        [Input("subnet", required: true)]
        public Input<string> Subnet { get; set; } = null!;

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// Instance tag of cloud provider.
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        /// <summary>
        /// Specify the transit Gateway.
        /// </summary>
        [Input("transitGw")]
        public Input<string>? TransitGw { get; set; }

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// Region of cloud provider.
        /// </summary>
        [Input("vpcReg", required: true)]
        public Input<string> VpcReg { get; set; } = null!;

        /// <summary>
        /// Size of the gateway instance.
        /// </summary>
        [Input("vpcSize", required: true)]
        public Input<string> VpcSize { get; set; } = null!;

        public AviatrixSpokeVpcArgs()
        {
        }
        public static new AviatrixSpokeVpcArgs Empty => new AviatrixSpokeVpcArgs();
    }

    public sealed class AviatrixSpokeVpcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of a Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Cloud instance ID.
        /// </summary>
        [Input("cloudInstanceId")]
        public Input<string>? CloudInstanceId { get; set; }

        /// <summary>
        /// Type of cloud service provider.
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Specify whether enabling NAT feature on the gateway or not.
        /// </summary>
        [Input("enableNat")]
        public Input<string>? EnableNat { get; set; }

        /// <summary>
        /// Name of the gateway which is going to be created.
        /// </summary>
        [Input("gwName")]
        public Input<string>? GwName { get; set; }

        /// <summary>
        /// HA Gateway Size.
        /// </summary>
        [Input("haGwSize")]
        public Input<string>? HaGwSize { get; set; }

        /// <summary>
        /// HA Subnet. Required if enabling HA for AWS/Azure.
        /// </summary>
        [Input("haSubnet")]
        public Input<string>? HaSubnet { get; set; }

        /// <summary>
        /// HA Zone. Required if enabling HA for GCP.
        /// </summary>
        [Input("haZone")]
        public Input<string>? HaZone { get; set; }

        /// <summary>
        /// Set to 'enabled' if this feature is desired.
        /// </summary>
        [Input("singleAzHa")]
        public Input<string>? SingleAzHa { get; set; }

        /// <summary>
        /// Public Subnet Info.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("tagLists")]
        private InputList<string>? _tagLists;

        /// <summary>
        /// Instance tag of cloud provider.
        /// </summary>
        public InputList<string> TagLists
        {
            get => _tagLists ?? (_tagLists = new InputList<string>());
            set => _tagLists = value;
        }

        /// <summary>
        /// Specify the transit Gateway.
        /// </summary>
        [Input("transitGw")]
        public Input<string>? TransitGw { get; set; }

        /// <summary>
        /// VPC-ID/VNet-Name of cloud provider.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Region of cloud provider.
        /// </summary>
        [Input("vpcReg")]
        public Input<string>? VpcReg { get; set; }

        /// <summary>
        /// Size of the gateway instance.
        /// </summary>
        [Input("vpcSize")]
        public Input<string>? VpcSize { get; set; }

        public AviatrixSpokeVpcState()
        {
        }
        public static new AviatrixSpokeVpcState Empty => new AviatrixSpokeVpcState();
    }
}