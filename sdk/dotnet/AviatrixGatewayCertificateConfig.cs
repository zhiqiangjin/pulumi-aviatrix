// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixGatewayCertificateConfig:AviatrixGatewayCertificateConfig")]
    public partial class AviatrixGatewayCertificateConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// CA Certificate.
        /// </summary>
        [Output("caCertificate")]
        public Output<string> CaCertificate { get; private set; } = null!;

        /// <summary>
        /// CA Private Key.
        /// </summary>
        [Output("caPrivateKey")]
        public Output<string> CaPrivateKey { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixGatewayCertificateConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixGatewayCertificateConfig(string name, AviatrixGatewayCertificateConfigArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixGatewayCertificateConfig:AviatrixGatewayCertificateConfig", name, args ?? new AviatrixGatewayCertificateConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixGatewayCertificateConfig(string name, Input<string> id, AviatrixGatewayCertificateConfigState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixGatewayCertificateConfig:AviatrixGatewayCertificateConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixGatewayCertificateConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixGatewayCertificateConfig Get(string name, Input<string> id, AviatrixGatewayCertificateConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixGatewayCertificateConfig(name, id, state, options);
        }
    }

    public sealed class AviatrixGatewayCertificateConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CA Certificate.
        /// </summary>
        [Input("caCertificate", required: true)]
        public Input<string> CaCertificate { get; set; } = null!;

        /// <summary>
        /// CA Private Key.
        /// </summary>
        [Input("caPrivateKey", required: true)]
        public Input<string> CaPrivateKey { get; set; } = null!;

        public AviatrixGatewayCertificateConfigArgs()
        {
        }
        public static new AviatrixGatewayCertificateConfigArgs Empty => new AviatrixGatewayCertificateConfigArgs();
    }

    public sealed class AviatrixGatewayCertificateConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CA Certificate.
        /// </summary>
        [Input("caCertificate")]
        public Input<string>? CaCertificate { get; set; }

        /// <summary>
        /// CA Private Key.
        /// </summary>
        [Input("caPrivateKey")]
        public Input<string>? CaPrivateKey { get; set; }

        public AviatrixGatewayCertificateConfigState()
        {
        }
        public static new AviatrixGatewayCertificateConfigState Empty => new AviatrixGatewayCertificateConfigState();
    }
}
