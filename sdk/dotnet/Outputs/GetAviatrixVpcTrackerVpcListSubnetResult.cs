// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix.Outputs
{

    [OutputType]
    public sealed class GetAviatrixVpcTrackerVpcListSubnetResult
    {
        public readonly string Cidr;
        public readonly string GwIp;
        public readonly string Name;
        public readonly string Region;

        [OutputConstructor]
        private GetAviatrixVpcTrackerVpcListSubnetResult(
            string cidr,

            string gwIp,

            string name,

            string region)
        {
            Cidr = cidr;
            GwIp = gwIp;
            Name = name;
            Region = region;
        }
    }
}
