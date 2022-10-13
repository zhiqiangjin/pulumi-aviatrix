// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixAwsPeer:AviatrixAwsPeer")]
    public partial class AviatrixAwsPeer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName1")]
        public Output<string> AccountName1 { get; private set; } = null!;

        /// <summary>
        /// This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
        /// </summary>
        [Output("accountName2")]
        public Output<string> AccountName2 { get; private set; } = null!;

        /// <summary>
        /// List of route table ID of vpc_id1.
        /// </summary>
        [Output("rtbList1Outputs")]
        public Output<ImmutableArray<string>> RtbList1Outputs { get; private set; } = null!;

        /// <summary>
        /// List of Route table ID.
        /// </summary>
        [Output("rtbList1s")]
        public Output<ImmutableArray<string>> RtbList1s { get; private set; } = null!;

        /// <summary>
        /// List of route table ID of vpc_id2.
        /// </summary>
        [Output("rtbList2Outputs")]
        public Output<ImmutableArray<string>> RtbList2Outputs { get; private set; } = null!;

        /// <summary>
        /// List of Route table ID.
        /// </summary>
        [Output("rtbList2s")]
        public Output<ImmutableArray<string>> RtbList2s { get; private set; } = null!;

        /// <summary>
        /// VPC-ID of AWS cloud.
        /// </summary>
        [Output("vpcId1")]
        public Output<string> VpcId1 { get; private set; } = null!;

        /// <summary>
        /// VPC-ID of AWS cloud.
        /// </summary>
        [Output("vpcId2")]
        public Output<string> VpcId2 { get; private set; } = null!;

        /// <summary>
        /// Region of AWS cloud.
        /// </summary>
        [Output("vpcReg1")]
        public Output<string> VpcReg1 { get; private set; } = null!;

        /// <summary>
        /// Region of AWS cloud.
        /// </summary>
        [Output("vpcReg2")]
        public Output<string> VpcReg2 { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixAwsPeer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixAwsPeer(string name, AviatrixAwsPeerArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsPeer:AviatrixAwsPeer", name, args ?? new AviatrixAwsPeerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixAwsPeer(string name, Input<string> id, AviatrixAwsPeerState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixAwsPeer:AviatrixAwsPeer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixAwsPeer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixAwsPeer Get(string name, Input<string> id, AviatrixAwsPeerState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixAwsPeer(name, id, state, options);
        }
    }

    public sealed class AviatrixAwsPeerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName1", required: true)]
        public Input<string> AccountName1 { get; set; } = null!;

        /// <summary>
        /// This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName2", required: true)]
        public Input<string> AccountName2 { get; set; } = null!;

        [Input("rtbList1s")]
        private InputList<string>? _rtbList1s;

        /// <summary>
        /// List of Route table ID.
        /// </summary>
        public InputList<string> RtbList1s
        {
            get => _rtbList1s ?? (_rtbList1s = new InputList<string>());
            set => _rtbList1s = value;
        }

        [Input("rtbList2s")]
        private InputList<string>? _rtbList2s;

        /// <summary>
        /// List of Route table ID.
        /// </summary>
        public InputList<string> RtbList2s
        {
            get => _rtbList2s ?? (_rtbList2s = new InputList<string>());
            set => _rtbList2s = value;
        }

        /// <summary>
        /// VPC-ID of AWS cloud.
        /// </summary>
        [Input("vpcId1", required: true)]
        public Input<string> VpcId1 { get; set; } = null!;

        /// <summary>
        /// VPC-ID of AWS cloud.
        /// </summary>
        [Input("vpcId2", required: true)]
        public Input<string> VpcId2 { get; set; } = null!;

        /// <summary>
        /// Region of AWS cloud.
        /// </summary>
        [Input("vpcReg1", required: true)]
        public Input<string> VpcReg1 { get; set; } = null!;

        /// <summary>
        /// Region of AWS cloud.
        /// </summary>
        [Input("vpcReg2", required: true)]
        public Input<string> VpcReg2 { get; set; } = null!;

        public AviatrixAwsPeerArgs()
        {
        }
        public static new AviatrixAwsPeerArgs Empty => new AviatrixAwsPeerArgs();
    }

    public sealed class AviatrixAwsPeerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName1")]
        public Input<string>? AccountName1 { get; set; }

        /// <summary>
        /// This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.
        /// </summary>
        [Input("accountName2")]
        public Input<string>? AccountName2 { get; set; }

        [Input("rtbList1Outputs")]
        private InputList<string>? _rtbList1Outputs;

        /// <summary>
        /// List of route table ID of vpc_id1.
        /// </summary>
        public InputList<string> RtbList1Outputs
        {
            get => _rtbList1Outputs ?? (_rtbList1Outputs = new InputList<string>());
            set => _rtbList1Outputs = value;
        }

        [Input("rtbList1s")]
        private InputList<string>? _rtbList1s;

        /// <summary>
        /// List of Route table ID.
        /// </summary>
        public InputList<string> RtbList1s
        {
            get => _rtbList1s ?? (_rtbList1s = new InputList<string>());
            set => _rtbList1s = value;
        }

        [Input("rtbList2Outputs")]
        private InputList<string>? _rtbList2Outputs;

        /// <summary>
        /// List of route table ID of vpc_id2.
        /// </summary>
        public InputList<string> RtbList2Outputs
        {
            get => _rtbList2Outputs ?? (_rtbList2Outputs = new InputList<string>());
            set => _rtbList2Outputs = value;
        }

        [Input("rtbList2s")]
        private InputList<string>? _rtbList2s;

        /// <summary>
        /// List of Route table ID.
        /// </summary>
        public InputList<string> RtbList2s
        {
            get => _rtbList2s ?? (_rtbList2s = new InputList<string>());
            set => _rtbList2s = value;
        }

        /// <summary>
        /// VPC-ID of AWS cloud.
        /// </summary>
        [Input("vpcId1")]
        public Input<string>? VpcId1 { get; set; }

        /// <summary>
        /// VPC-ID of AWS cloud.
        /// </summary>
        [Input("vpcId2")]
        public Input<string>? VpcId2 { get; set; }

        /// <summary>
        /// Region of AWS cloud.
        /// </summary>
        [Input("vpcReg1")]
        public Input<string>? VpcReg1 { get; set; }

        /// <summary>
        /// Region of AWS cloud.
        /// </summary>
        [Input("vpcReg2")]
        public Input<string>? VpcReg2 { get; set; }

        public AviatrixAwsPeerState()
        {
        }
        public static new AviatrixAwsPeerState Empty => new AviatrixAwsPeerState();
    }
}
