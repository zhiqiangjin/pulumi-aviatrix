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
    public sealed class GetAviatrixNetworkDomainsNetworkDomainResult
    {
        public readonly string Account;
        public readonly string CloudType;
        public readonly bool EgressInspection;
        public readonly string EgressInspectionName;
        public readonly string InspectionPolicy;
        public readonly bool IntraDomainInspection;
        public readonly string IntraDomainInspectionName;
        public readonly string Name;
        public readonly string Region;
        public readonly string RouteTableId;
        public readonly string TgwName;
        public readonly string Type;

        [OutputConstructor]
        private GetAviatrixNetworkDomainsNetworkDomainResult(
            string account,

            string cloudType,

            bool egressInspection,

            string egressInspectionName,

            string inspectionPolicy,

            bool intraDomainInspection,

            string intraDomainInspectionName,

            string name,

            string region,

            string routeTableId,

            string tgwName,

            string type)
        {
            Account = account;
            CloudType = cloudType;
            EgressInspection = egressInspection;
            EgressInspectionName = egressInspectionName;
            InspectionPolicy = inspectionPolicy;
            IntraDomainInspection = intraDomainInspection;
            IntraDomainInspectionName = intraDomainInspectionName;
            Name = name;
            Region = region;
            RouteTableId = routeTableId;
            TgwName = tgwName;
            Type = type;
        }
    }
}
