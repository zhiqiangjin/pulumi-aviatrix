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
    public sealed class GetAviatrixVpcTrackerVpcListResult
    {
        public readonly string AccountName;
        public readonly string Cidr;
        public readonly int CloudType;
        public readonly int InstanceCount;
        public readonly string Name;
        public readonly string Region;
        public readonly ImmutableArray<Outputs.GetAviatrixVpcTrackerVpcListSubnetResult> Subnets;
        public readonly string VpcId;

        [OutputConstructor]
        private GetAviatrixVpcTrackerVpcListResult(
            string accountName,

            string cidr,

            int cloudType,

            int instanceCount,

            string name,

            string region,

            ImmutableArray<Outputs.GetAviatrixVpcTrackerVpcListSubnetResult> subnets,

            string vpcId)
        {
            AccountName = accountName;
            Cidr = cidr;
            CloudType = cloudType;
            InstanceCount = instanceCount;
            Name = name;
            Region = region;
            Subnets = subnets;
            VpcId = vpcId;
        }
    }
}
