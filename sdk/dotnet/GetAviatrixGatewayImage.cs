// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    public static class GetAviatrixGatewayImage
    {
        public static Task<GetAviatrixGatewayImageResult> InvokeAsync(GetAviatrixGatewayImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAviatrixGatewayImageResult>("aviatrix:index/getAviatrixGatewayImage:getAviatrixGatewayImage", args ?? new GetAviatrixGatewayImageArgs(), options.WithDefaults());

        public static Output<GetAviatrixGatewayImageResult> Invoke(GetAviatrixGatewayImageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAviatrixGatewayImageResult>("aviatrix:index/getAviatrixGatewayImage:getAviatrixGatewayImage", args ?? new GetAviatrixGatewayImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAviatrixGatewayImageArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloudType", required: true)]
        public int CloudType { get; set; }

        [Input("softwareVersion", required: true)]
        public string SoftwareVersion { get; set; } = null!;

        public GetAviatrixGatewayImageArgs()
        {
        }
        public static new GetAviatrixGatewayImageArgs Empty => new GetAviatrixGatewayImageArgs();
    }

    public sealed class GetAviatrixGatewayImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloudType", required: true)]
        public Input<int> CloudType { get; set; } = null!;

        [Input("softwareVersion", required: true)]
        public Input<string> SoftwareVersion { get; set; } = null!;

        public GetAviatrixGatewayImageInvokeArgs()
        {
        }
        public static new GetAviatrixGatewayImageInvokeArgs Empty => new GetAviatrixGatewayImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetAviatrixGatewayImageResult
    {
        public readonly int CloudType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ImageVersion;
        public readonly string SoftwareVersion;

        [OutputConstructor]
        private GetAviatrixGatewayImageResult(
            int cloudType,

            string id,

            string imageVersion,

            string softwareVersion)
        {
            CloudType = cloudType;
            Id = id;
            ImageVersion = imageVersion;
            SoftwareVersion = softwareVersion;
        }
    }
}
