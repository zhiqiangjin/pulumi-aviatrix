// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aviatrix
{
    [AviatrixResourceType("aviatrix:index/aviatrixFirewallInstance:AviatrixFirewallInstance")]
    public partial class AviatrixFirewallInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Availability domain for OCI.
        /// </summary>
        [Output("availabilityDomain")]
        public Output<string> AvailabilityDomain { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
        /// </summary>
        [Output("bootstrapBucketName")]
        public Output<string?> BootstrapBucketName { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
        /// only.
        /// </summary>
        [Output("bootstrapStorageName")]
        public Output<string?> BootstrapStorageName { get; private set; } = null!;

        /// <summary>
        /// Cloud Type
        /// </summary>
        [Output("cloudType")]
        public Output<int> CloudType { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Output("containerFolder")]
        public Output<string?> ContainerFolder { get; private set; } = null!;

        /// <summary>
        /// ID of Egress Interface created.
        /// </summary>
        [Output("egressInterface")]
        public Output<string> EgressInterface { get; private set; } = null!;

        /// <summary>
        /// Egress Interface Subnet.
        /// </summary>
        [Output("egressSubnet")]
        public Output<string> EgressSubnet { get; private set; } = null!;

        /// <summary>
        /// Egress VPC ID. Required for GCP.
        /// </summary>
        [Output("egressVpcId")]
        public Output<string?> EgressVpcId { get; private set; } = null!;

        /// <summary>
        /// Fault domain for OCI.
        /// </summary>
        [Output("faultDomain")]
        public Output<string> FaultDomain { get; private set; } = null!;

        /// <summary>
        /// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Output("fileShareFolder")]
        public Output<string?> FileShareFolder { get; private set; } = null!;

        /// <summary>
        /// Name of the primary FireNet gateway.
        /// </summary>
        [Output("firenetGwName")]
        public Output<string?> FirenetGwName { get; private set; } = null!;

        /// <summary>
        /// One of the AWS AMIs from Palo Alto Networks.
        /// </summary>
        [Output("firewallImage")]
        public Output<string> FirewallImage { get; private set; } = null!;

        /// <summary>
        /// Firewall image ID.
        /// </summary>
        [Output("firewallImageId")]
        public Output<string> FirewallImageId { get; private set; } = null!;

        /// <summary>
        /// Version of firewall image.
        /// </summary>
        [Output("firewallImageVersion")]
        public Output<string> FirewallImageVersion { get; private set; } = null!;

        /// <summary>
        /// Name of the firewall instance to be created.
        /// </summary>
        [Output("firewallName")]
        public Output<string> FirewallName { get; private set; } = null!;

        /// <summary>
        /// Instance size of the firewall.
        /// </summary>
        [Output("firewallSize")]
        public Output<string> FirewallSize { get; private set; } = null!;

        /// <summary>
        /// GCP VPC ID
        /// </summary>
        [Output("gcpVpcId")]
        public Output<string> GcpVpcId { get; private set; } = null!;

        /// <summary>
        /// Advanced option. IAM role. Only available for AWS.
        /// </summary>
        [Output("iamRole")]
        public Output<string?> IamRole { get; private set; } = null!;

        /// <summary>
        /// ID of the firewall instance created.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
        /// </summary>
        [Output("keyName")]
        public Output<string?> KeyName { get; private set; } = null!;

        /// <summary>
        /// ID of Lan Interface created.
        /// </summary>
        [Output("lanInterface")]
        public Output<string> LanInterface { get; private set; } = null!;

        /// <summary>
        /// ID of Management Interface created.
        /// </summary>
        [Output("managementInterface")]
        public Output<string> ManagementInterface { get; private set; } = null!;

        /// <summary>
        /// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
        /// Fortinet series.
        /// </summary>
        [Output("managementSubnet")]
        public Output<string?> ManagementSubnet { get; private set; } = null!;

        /// <summary>
        /// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
        /// series.
        /// </summary>
        [Output("managementVpcId")]
        public Output<string?> ManagementVpcId { get; private set; } = null!;

        /// <summary>
        /// Authentication method. Applicable to Azure deployment only.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Management Public IP.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Output("sasUrlConfig")]
        public Output<string?> SasUrlConfig { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Output("sasUrlLicense")]
        public Output<string?> SasUrlLicense { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Output("shareDirectory")]
        public Output<string?> ShareDirectory { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
        /// </summary>
        [Output("sicKey")]
        public Output<string?> SicKey { get; private set; } = null!;

        /// <summary>
        /// Authentication method. Applicable to Azure deployment only.
        /// </summary>
        [Output("sshPublicKey")]
        public Output<string?> SshPublicKey { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Output("storageAccessKey")]
        public Output<string?> StorageAccessKey { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the firewall instance.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        /// <summary>
        /// Applicable to Azure deployment only. 'admin' as a username is not accepted.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// ID of the Security VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// Availability Zone. Only available for AWS, GCP and Azure.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a AviatrixFirewallInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AviatrixFirewallInstance(string name, AviatrixFirewallInstanceArgs args, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirewallInstance:AviatrixFirewallInstance", name, args ?? new AviatrixFirewallInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AviatrixFirewallInstance(string name, Input<string> id, AviatrixFirewallInstanceState? state = null, CustomResourceOptions? options = null)
            : base("aviatrix:index/aviatrixFirewallInstance:AviatrixFirewallInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AviatrixFirewallInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AviatrixFirewallInstance Get(string name, Input<string> id, AviatrixFirewallInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new AviatrixFirewallInstance(name, id, state, options);
        }
    }

    public sealed class AviatrixFirewallInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Availability domain for OCI.
        /// </summary>
        [Input("availabilityDomain")]
        public Input<string>? AvailabilityDomain { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
        /// </summary>
        [Input("bootstrapBucketName")]
        public Input<string>? BootstrapBucketName { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
        /// only.
        /// </summary>
        [Input("bootstrapStorageName")]
        public Input<string>? BootstrapStorageName { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Input("containerFolder")]
        public Input<string>? ContainerFolder { get; set; }

        /// <summary>
        /// Egress Interface Subnet.
        /// </summary>
        [Input("egressSubnet", required: true)]
        public Input<string> EgressSubnet { get; set; } = null!;

        /// <summary>
        /// Egress VPC ID. Required for GCP.
        /// </summary>
        [Input("egressVpcId")]
        public Input<string>? EgressVpcId { get; set; }

        /// <summary>
        /// Fault domain for OCI.
        /// </summary>
        [Input("faultDomain")]
        public Input<string>? FaultDomain { get; set; }

        /// <summary>
        /// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Input("fileShareFolder")]
        public Input<string>? FileShareFolder { get; set; }

        /// <summary>
        /// Name of the primary FireNet gateway.
        /// </summary>
        [Input("firenetGwName")]
        public Input<string>? FirenetGwName { get; set; }

        /// <summary>
        /// One of the AWS AMIs from Palo Alto Networks.
        /// </summary>
        [Input("firewallImage", required: true)]
        public Input<string> FirewallImage { get; set; } = null!;

        /// <summary>
        /// Firewall image ID.
        /// </summary>
        [Input("firewallImageId")]
        public Input<string>? FirewallImageId { get; set; }

        /// <summary>
        /// Version of firewall image.
        /// </summary>
        [Input("firewallImageVersion")]
        public Input<string>? FirewallImageVersion { get; set; }

        /// <summary>
        /// Name of the firewall instance to be created.
        /// </summary>
        [Input("firewallName", required: true)]
        public Input<string> FirewallName { get; set; } = null!;

        /// <summary>
        /// Instance size of the firewall.
        /// </summary>
        [Input("firewallSize", required: true)]
        public Input<string> FirewallSize { get; set; } = null!;

        /// <summary>
        /// Advanced option. IAM role. Only available for AWS.
        /// </summary>
        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        /// <summary>
        /// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
        /// Fortinet series.
        /// </summary>
        [Input("managementSubnet")]
        public Input<string>? ManagementSubnet { get; set; }

        /// <summary>
        /// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
        /// series.
        /// </summary>
        [Input("managementVpcId")]
        public Input<string>? ManagementVpcId { get; set; }

        /// <summary>
        /// Authentication method. Applicable to Azure deployment only.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Input("sasUrlConfig")]
        public Input<string>? SasUrlConfig { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Input("sasUrlLicense")]
        public Input<string>? SasUrlLicense { get; set; }

        /// <summary>
        /// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Input("shareDirectory")]
        public Input<string>? ShareDirectory { get; set; }

        /// <summary>
        /// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
        /// </summary>
        [Input("sicKey")]
        public Input<string>? SicKey { get; set; }

        /// <summary>
        /// Authentication method. Applicable to Azure deployment only.
        /// </summary>
        [Input("sshPublicKey")]
        public Input<string>? SshPublicKey { get; set; }

        /// <summary>
        /// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Input("storageAccessKey")]
        public Input<string>? StorageAccessKey { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the firewall instance.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// Applicable to Azure deployment only. 'admin' as a username is not accepted.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// ID of the Security VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// Availability Zone. Only available for AWS, GCP and Azure.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AviatrixFirewallInstanceArgs()
        {
        }
        public static new AviatrixFirewallInstanceArgs Empty => new AviatrixFirewallInstanceArgs();
    }

    public sealed class AviatrixFirewallInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Availability domain for OCI.
        /// </summary>
        [Input("availabilityDomain")]
        public Input<string>? AvailabilityDomain { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
        /// </summary>
        [Input("bootstrapBucketName")]
        public Input<string>? BootstrapBucketName { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
        /// only.
        /// </summary>
        [Input("bootstrapStorageName")]
        public Input<string>? BootstrapStorageName { get; set; }

        /// <summary>
        /// Cloud Type
        /// </summary>
        [Input("cloudType")]
        public Input<int>? CloudType { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Input("containerFolder")]
        public Input<string>? ContainerFolder { get; set; }

        /// <summary>
        /// ID of Egress Interface created.
        /// </summary>
        [Input("egressInterface")]
        public Input<string>? EgressInterface { get; set; }

        /// <summary>
        /// Egress Interface Subnet.
        /// </summary>
        [Input("egressSubnet")]
        public Input<string>? EgressSubnet { get; set; }

        /// <summary>
        /// Egress VPC ID. Required for GCP.
        /// </summary>
        [Input("egressVpcId")]
        public Input<string>? EgressVpcId { get; set; }

        /// <summary>
        /// Fault domain for OCI.
        /// </summary>
        [Input("faultDomain")]
        public Input<string>? FaultDomain { get; set; }

        /// <summary>
        /// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Input("fileShareFolder")]
        public Input<string>? FileShareFolder { get; set; }

        /// <summary>
        /// Name of the primary FireNet gateway.
        /// </summary>
        [Input("firenetGwName")]
        public Input<string>? FirenetGwName { get; set; }

        /// <summary>
        /// One of the AWS AMIs from Palo Alto Networks.
        /// </summary>
        [Input("firewallImage")]
        public Input<string>? FirewallImage { get; set; }

        /// <summary>
        /// Firewall image ID.
        /// </summary>
        [Input("firewallImageId")]
        public Input<string>? FirewallImageId { get; set; }

        /// <summary>
        /// Version of firewall image.
        /// </summary>
        [Input("firewallImageVersion")]
        public Input<string>? FirewallImageVersion { get; set; }

        /// <summary>
        /// Name of the firewall instance to be created.
        /// </summary>
        [Input("firewallName")]
        public Input<string>? FirewallName { get; set; }

        /// <summary>
        /// Instance size of the firewall.
        /// </summary>
        [Input("firewallSize")]
        public Input<string>? FirewallSize { get; set; }

        /// <summary>
        /// GCP VPC ID
        /// </summary>
        [Input("gcpVpcId")]
        public Input<string>? GcpVpcId { get; set; }

        /// <summary>
        /// Advanced option. IAM role. Only available for AWS.
        /// </summary>
        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        /// <summary>
        /// ID of the firewall instance created.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// ID of Lan Interface created.
        /// </summary>
        [Input("lanInterface")]
        public Input<string>? LanInterface { get; set; }

        /// <summary>
        /// ID of Management Interface created.
        /// </summary>
        [Input("managementInterface")]
        public Input<string>? ManagementInterface { get; set; }

        /// <summary>
        /// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
        /// Fortinet series.
        /// </summary>
        [Input("managementSubnet")]
        public Input<string>? ManagementSubnet { get; set; }

        /// <summary>
        /// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
        /// series.
        /// </summary>
        [Input("managementVpcId")]
        public Input<string>? ManagementVpcId { get; set; }

        /// <summary>
        /// Authentication method. Applicable to Azure deployment only.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Management Public IP.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Input("sasUrlConfig")]
        public Input<string>? SasUrlConfig { get; set; }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
        /// </summary>
        [Input("sasUrlLicense")]
        public Input<string>? SasUrlLicense { get; set; }

        /// <summary>
        /// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Input("shareDirectory")]
        public Input<string>? ShareDirectory { get; set; }

        /// <summary>
        /// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
        /// </summary>
        [Input("sicKey")]
        public Input<string>? SicKey { get; set; }

        /// <summary>
        /// Authentication method. Applicable to Azure deployment only.
        /// </summary>
        [Input("sshPublicKey")]
        public Input<string>? SshPublicKey { get; set; }

        /// <summary>
        /// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
        /// </summary>
        [Input("storageAccessKey")]
        public Input<string>? StorageAccessKey { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the firewall instance.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// Applicable to Azure deployment only. 'admin' as a username is not accepted.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// ID of the Security VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Availability Zone. Only available for AWS, GCP and Azure.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AviatrixFirewallInstanceState()
        {
        }
        public static new AviatrixFirewallInstanceState Empty => new AviatrixFirewallInstanceState();
    }
}
