// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixFirewallInstance struct {
	pulumi.CustomResourceState

	// Availability domain for OCI.
	AvailabilityDomain pulumi.StringOutput `pulumi:"availabilityDomain"`
	// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
	BootstrapBucketName pulumi.StringPtrOutput `pulumi:"bootstrapBucketName"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
	// only.
	BootstrapStorageName pulumi.StringPtrOutput `pulumi:"bootstrapStorageName"`
	// Cloud Type
	CloudType pulumi.IntOutput `pulumi:"cloudType"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	ContainerFolder pulumi.StringPtrOutput `pulumi:"containerFolder"`
	// ID of Egress Interface created.
	EgressInterface pulumi.StringOutput `pulumi:"egressInterface"`
	// Egress Interface Subnet.
	EgressSubnet pulumi.StringOutput `pulumi:"egressSubnet"`
	// Egress VPC ID. Required for GCP.
	EgressVpcId pulumi.StringPtrOutput `pulumi:"egressVpcId"`
	// Fault domain for OCI.
	FaultDomain pulumi.StringOutput `pulumi:"faultDomain"`
	// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	FileShareFolder pulumi.StringPtrOutput `pulumi:"fileShareFolder"`
	// Name of the primary FireNet gateway.
	FirenetGwName pulumi.StringPtrOutput `pulumi:"firenetGwName"`
	// One of the AWS AMIs from Palo Alto Networks.
	FirewallImage pulumi.StringOutput `pulumi:"firewallImage"`
	// Firewall image ID.
	FirewallImageId pulumi.StringOutput `pulumi:"firewallImageId"`
	// Version of firewall image.
	FirewallImageVersion pulumi.StringOutput `pulumi:"firewallImageVersion"`
	// Name of the firewall instance to be created.
	FirewallName pulumi.StringOutput `pulumi:"firewallName"`
	// Instance size of the firewall.
	FirewallSize pulumi.StringOutput `pulumi:"firewallSize"`
	// GCP VPC ID
	GcpVpcId pulumi.StringOutput `pulumi:"gcpVpcId"`
	// Advanced option. IAM role. Only available for AWS.
	IamRole pulumi.StringPtrOutput `pulumi:"iamRole"`
	// ID of the firewall instance created.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// ID of Lan Interface created.
	LanInterface pulumi.StringOutput `pulumi:"lanInterface"`
	// ID of Management Interface created.
	ManagementInterface pulumi.StringOutput `pulumi:"managementInterface"`
	// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
	// Fortinet series.
	ManagementSubnet pulumi.StringPtrOutput `pulumi:"managementSubnet"`
	// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
	// series.
	ManagementVpcId pulumi.StringPtrOutput `pulumi:"managementVpcId"`
	// Authentication method. Applicable to Azure deployment only.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Management Public IP.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlConfig pulumi.StringPtrOutput `pulumi:"sasUrlConfig"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlLicense pulumi.StringPtrOutput `pulumi:"sasUrlLicense"`
	// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	ShareDirectory pulumi.StringPtrOutput `pulumi:"shareDirectory"`
	// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
	SicKey pulumi.StringPtrOutput `pulumi:"sicKey"`
	// Authentication method. Applicable to Azure deployment only.
	SshPublicKey pulumi.StringPtrOutput `pulumi:"sshPublicKey"`
	// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	StorageAccessKey pulumi.StringPtrOutput `pulumi:"storageAccessKey"`
	// A map of tags to assign to the firewall instance.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// Applicable to Azure deployment only. 'admin' as a username is not accepted.
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// ID of the Security VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Availability Zone. Only available for AWS, GCP and Azure.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewAviatrixFirewallInstance registers a new resource with the given unique name, arguments, and options.
func NewAviatrixFirewallInstance(ctx *pulumi.Context,
	name string, args *AviatrixFirewallInstanceArgs, opts ...pulumi.ResourceOption) (*AviatrixFirewallInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EgressSubnet == nil {
		return nil, errors.New("invalid value for required argument 'EgressSubnet'")
	}
	if args.FirewallImage == nil {
		return nil, errors.New("invalid value for required argument 'FirewallImage'")
	}
	if args.FirewallName == nil {
		return nil, errors.New("invalid value for required argument 'FirewallName'")
	}
	if args.FirewallSize == nil {
		return nil, errors.New("invalid value for required argument 'FirewallSize'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixFirewallInstance
	err := ctx.RegisterResource("aviatrix:index/aviatrixFirewallInstance:AviatrixFirewallInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixFirewallInstance gets an existing AviatrixFirewallInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixFirewallInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixFirewallInstanceState, opts ...pulumi.ResourceOption) (*AviatrixFirewallInstance, error) {
	var resource AviatrixFirewallInstance
	err := ctx.ReadResource("aviatrix:index/aviatrixFirewallInstance:AviatrixFirewallInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixFirewallInstance resources.
type aviatrixFirewallInstanceState struct {
	// Availability domain for OCI.
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
	BootstrapBucketName *string `pulumi:"bootstrapBucketName"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
	// only.
	BootstrapStorageName *string `pulumi:"bootstrapStorageName"`
	// Cloud Type
	CloudType *int `pulumi:"cloudType"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	ContainerFolder *string `pulumi:"containerFolder"`
	// ID of Egress Interface created.
	EgressInterface *string `pulumi:"egressInterface"`
	// Egress Interface Subnet.
	EgressSubnet *string `pulumi:"egressSubnet"`
	// Egress VPC ID. Required for GCP.
	EgressVpcId *string `pulumi:"egressVpcId"`
	// Fault domain for OCI.
	FaultDomain *string `pulumi:"faultDomain"`
	// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	FileShareFolder *string `pulumi:"fileShareFolder"`
	// Name of the primary FireNet gateway.
	FirenetGwName *string `pulumi:"firenetGwName"`
	// One of the AWS AMIs from Palo Alto Networks.
	FirewallImage *string `pulumi:"firewallImage"`
	// Firewall image ID.
	FirewallImageId *string `pulumi:"firewallImageId"`
	// Version of firewall image.
	FirewallImageVersion *string `pulumi:"firewallImageVersion"`
	// Name of the firewall instance to be created.
	FirewallName *string `pulumi:"firewallName"`
	// Instance size of the firewall.
	FirewallSize *string `pulumi:"firewallSize"`
	// GCP VPC ID
	GcpVpcId *string `pulumi:"gcpVpcId"`
	// Advanced option. IAM role. Only available for AWS.
	IamRole *string `pulumi:"iamRole"`
	// ID of the firewall instance created.
	InstanceId *string `pulumi:"instanceId"`
	// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
	KeyName *string `pulumi:"keyName"`
	// ID of Lan Interface created.
	LanInterface *string `pulumi:"lanInterface"`
	// ID of Management Interface created.
	ManagementInterface *string `pulumi:"managementInterface"`
	// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
	// Fortinet series.
	ManagementSubnet *string `pulumi:"managementSubnet"`
	// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
	// series.
	ManagementVpcId *string `pulumi:"managementVpcId"`
	// Authentication method. Applicable to Azure deployment only.
	Password *string `pulumi:"password"`
	// Management Public IP.
	PublicIp *string `pulumi:"publicIp"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlConfig *string `pulumi:"sasUrlConfig"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlLicense *string `pulumi:"sasUrlLicense"`
	// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	ShareDirectory *string `pulumi:"shareDirectory"`
	// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
	SicKey *string `pulumi:"sicKey"`
	// Authentication method. Applicable to Azure deployment only.
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	StorageAccessKey *string `pulumi:"storageAccessKey"`
	// A map of tags to assign to the firewall instance.
	Tags map[string]string `pulumi:"tags"`
	// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
	UserData *string `pulumi:"userData"`
	// Applicable to Azure deployment only. 'admin' as a username is not accepted.
	Username *string `pulumi:"username"`
	// ID of the Security VPC.
	VpcId *string `pulumi:"vpcId"`
	// Availability Zone. Only available for AWS, GCP and Azure.
	Zone *string `pulumi:"zone"`
}

type AviatrixFirewallInstanceState struct {
	// Availability domain for OCI.
	AvailabilityDomain pulumi.StringPtrInput
	// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
	BootstrapBucketName pulumi.StringPtrInput
	// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
	// only.
	BootstrapStorageName pulumi.StringPtrInput
	// Cloud Type
	CloudType pulumi.IntPtrInput
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	ContainerFolder pulumi.StringPtrInput
	// ID of Egress Interface created.
	EgressInterface pulumi.StringPtrInput
	// Egress Interface Subnet.
	EgressSubnet pulumi.StringPtrInput
	// Egress VPC ID. Required for GCP.
	EgressVpcId pulumi.StringPtrInput
	// Fault domain for OCI.
	FaultDomain pulumi.StringPtrInput
	// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	FileShareFolder pulumi.StringPtrInput
	// Name of the primary FireNet gateway.
	FirenetGwName pulumi.StringPtrInput
	// One of the AWS AMIs from Palo Alto Networks.
	FirewallImage pulumi.StringPtrInput
	// Firewall image ID.
	FirewallImageId pulumi.StringPtrInput
	// Version of firewall image.
	FirewallImageVersion pulumi.StringPtrInput
	// Name of the firewall instance to be created.
	FirewallName pulumi.StringPtrInput
	// Instance size of the firewall.
	FirewallSize pulumi.StringPtrInput
	// GCP VPC ID
	GcpVpcId pulumi.StringPtrInput
	// Advanced option. IAM role. Only available for AWS.
	IamRole pulumi.StringPtrInput
	// ID of the firewall instance created.
	InstanceId pulumi.StringPtrInput
	// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
	KeyName pulumi.StringPtrInput
	// ID of Lan Interface created.
	LanInterface pulumi.StringPtrInput
	// ID of Management Interface created.
	ManagementInterface pulumi.StringPtrInput
	// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
	// Fortinet series.
	ManagementSubnet pulumi.StringPtrInput
	// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
	// series.
	ManagementVpcId pulumi.StringPtrInput
	// Authentication method. Applicable to Azure deployment only.
	Password pulumi.StringPtrInput
	// Management Public IP.
	PublicIp pulumi.StringPtrInput
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlConfig pulumi.StringPtrInput
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlLicense pulumi.StringPtrInput
	// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	ShareDirectory pulumi.StringPtrInput
	// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
	SicKey pulumi.StringPtrInput
	// Authentication method. Applicable to Azure deployment only.
	SshPublicKey pulumi.StringPtrInput
	// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	StorageAccessKey pulumi.StringPtrInput
	// A map of tags to assign to the firewall instance.
	Tags pulumi.StringMapInput
	// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
	UserData pulumi.StringPtrInput
	// Applicable to Azure deployment only. 'admin' as a username is not accepted.
	Username pulumi.StringPtrInput
	// ID of the Security VPC.
	VpcId pulumi.StringPtrInput
	// Availability Zone. Only available for AWS, GCP and Azure.
	Zone pulumi.StringPtrInput
}

func (AviatrixFirewallInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixFirewallInstanceState)(nil)).Elem()
}

type aviatrixFirewallInstanceArgs struct {
	// Availability domain for OCI.
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
	BootstrapBucketName *string `pulumi:"bootstrapBucketName"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
	// only.
	BootstrapStorageName *string `pulumi:"bootstrapStorageName"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	ContainerFolder *string `pulumi:"containerFolder"`
	// Egress Interface Subnet.
	EgressSubnet string `pulumi:"egressSubnet"`
	// Egress VPC ID. Required for GCP.
	EgressVpcId *string `pulumi:"egressVpcId"`
	// Fault domain for OCI.
	FaultDomain *string `pulumi:"faultDomain"`
	// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	FileShareFolder *string `pulumi:"fileShareFolder"`
	// Name of the primary FireNet gateway.
	FirenetGwName *string `pulumi:"firenetGwName"`
	// One of the AWS AMIs from Palo Alto Networks.
	FirewallImage string `pulumi:"firewallImage"`
	// Firewall image ID.
	FirewallImageId *string `pulumi:"firewallImageId"`
	// Version of firewall image.
	FirewallImageVersion *string `pulumi:"firewallImageVersion"`
	// Name of the firewall instance to be created.
	FirewallName string `pulumi:"firewallName"`
	// Instance size of the firewall.
	FirewallSize string `pulumi:"firewallSize"`
	// Advanced option. IAM role. Only available for AWS.
	IamRole *string `pulumi:"iamRole"`
	// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
	KeyName *string `pulumi:"keyName"`
	// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
	// Fortinet series.
	ManagementSubnet *string `pulumi:"managementSubnet"`
	// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
	// series.
	ManagementVpcId *string `pulumi:"managementVpcId"`
	// Authentication method. Applicable to Azure deployment only.
	Password *string `pulumi:"password"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlConfig *string `pulumi:"sasUrlConfig"`
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlLicense *string `pulumi:"sasUrlLicense"`
	// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	ShareDirectory *string `pulumi:"shareDirectory"`
	// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
	SicKey *string `pulumi:"sicKey"`
	// Authentication method. Applicable to Azure deployment only.
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	StorageAccessKey *string `pulumi:"storageAccessKey"`
	// A map of tags to assign to the firewall instance.
	Tags map[string]string `pulumi:"tags"`
	// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
	UserData *string `pulumi:"userData"`
	// Applicable to Azure deployment only. 'admin' as a username is not accepted.
	Username *string `pulumi:"username"`
	// ID of the Security VPC.
	VpcId string `pulumi:"vpcId"`
	// Availability Zone. Only available for AWS, GCP and Azure.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a AviatrixFirewallInstance resource.
type AviatrixFirewallInstanceArgs struct {
	// Availability domain for OCI.
	AvailabilityDomain pulumi.StringPtrInput
	// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
	BootstrapBucketName pulumi.StringPtrInput
	// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
	// only.
	BootstrapStorageName pulumi.StringPtrInput
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	ContainerFolder pulumi.StringPtrInput
	// Egress Interface Subnet.
	EgressSubnet pulumi.StringInput
	// Egress VPC ID. Required for GCP.
	EgressVpcId pulumi.StringPtrInput
	// Fault domain for OCI.
	FaultDomain pulumi.StringPtrInput
	// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	FileShareFolder pulumi.StringPtrInput
	// Name of the primary FireNet gateway.
	FirenetGwName pulumi.StringPtrInput
	// One of the AWS AMIs from Palo Alto Networks.
	FirewallImage pulumi.StringInput
	// Firewall image ID.
	FirewallImageId pulumi.StringPtrInput
	// Version of firewall image.
	FirewallImageVersion pulumi.StringPtrInput
	// Name of the firewall instance to be created.
	FirewallName pulumi.StringInput
	// Instance size of the firewall.
	FirewallSize pulumi.StringInput
	// Advanced option. IAM role. Only available for AWS.
	IamRole pulumi.StringPtrInput
	// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
	KeyName pulumi.StringPtrInput
	// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
	// Fortinet series.
	ManagementSubnet pulumi.StringPtrInput
	// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
	// series.
	ManagementVpcId pulumi.StringPtrInput
	// Authentication method. Applicable to Azure deployment only.
	Password pulumi.StringPtrInput
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlConfig pulumi.StringPtrInput
	// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
	SasUrlLicense pulumi.StringPtrInput
	// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	ShareDirectory pulumi.StringPtrInput
	// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
	SicKey pulumi.StringPtrInput
	// Authentication method. Applicable to Azure deployment only.
	SshPublicKey pulumi.StringPtrInput
	// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
	StorageAccessKey pulumi.StringPtrInput
	// A map of tags to assign to the firewall instance.
	Tags pulumi.StringMapInput
	// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
	UserData pulumi.StringPtrInput
	// Applicable to Azure deployment only. 'admin' as a username is not accepted.
	Username pulumi.StringPtrInput
	// ID of the Security VPC.
	VpcId pulumi.StringInput
	// Availability Zone. Only available for AWS, GCP and Azure.
	Zone pulumi.StringPtrInput
}

func (AviatrixFirewallInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixFirewallInstanceArgs)(nil)).Elem()
}

type AviatrixFirewallInstanceInput interface {
	pulumi.Input

	ToAviatrixFirewallInstanceOutput() AviatrixFirewallInstanceOutput
	ToAviatrixFirewallInstanceOutputWithContext(ctx context.Context) AviatrixFirewallInstanceOutput
}

func (*AviatrixFirewallInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixFirewallInstance)(nil)).Elem()
}

func (i *AviatrixFirewallInstance) ToAviatrixFirewallInstanceOutput() AviatrixFirewallInstanceOutput {
	return i.ToAviatrixFirewallInstanceOutputWithContext(context.Background())
}

func (i *AviatrixFirewallInstance) ToAviatrixFirewallInstanceOutputWithContext(ctx context.Context) AviatrixFirewallInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirewallInstanceOutput)
}

// AviatrixFirewallInstanceArrayInput is an input type that accepts AviatrixFirewallInstanceArray and AviatrixFirewallInstanceArrayOutput values.
// You can construct a concrete instance of `AviatrixFirewallInstanceArrayInput` via:
//
//	AviatrixFirewallInstanceArray{ AviatrixFirewallInstanceArgs{...} }
type AviatrixFirewallInstanceArrayInput interface {
	pulumi.Input

	ToAviatrixFirewallInstanceArrayOutput() AviatrixFirewallInstanceArrayOutput
	ToAviatrixFirewallInstanceArrayOutputWithContext(context.Context) AviatrixFirewallInstanceArrayOutput
}

type AviatrixFirewallInstanceArray []AviatrixFirewallInstanceInput

func (AviatrixFirewallInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixFirewallInstance)(nil)).Elem()
}

func (i AviatrixFirewallInstanceArray) ToAviatrixFirewallInstanceArrayOutput() AviatrixFirewallInstanceArrayOutput {
	return i.ToAviatrixFirewallInstanceArrayOutputWithContext(context.Background())
}

func (i AviatrixFirewallInstanceArray) ToAviatrixFirewallInstanceArrayOutputWithContext(ctx context.Context) AviatrixFirewallInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirewallInstanceArrayOutput)
}

// AviatrixFirewallInstanceMapInput is an input type that accepts AviatrixFirewallInstanceMap and AviatrixFirewallInstanceMapOutput values.
// You can construct a concrete instance of `AviatrixFirewallInstanceMapInput` via:
//
//	AviatrixFirewallInstanceMap{ "key": AviatrixFirewallInstanceArgs{...} }
type AviatrixFirewallInstanceMapInput interface {
	pulumi.Input

	ToAviatrixFirewallInstanceMapOutput() AviatrixFirewallInstanceMapOutput
	ToAviatrixFirewallInstanceMapOutputWithContext(context.Context) AviatrixFirewallInstanceMapOutput
}

type AviatrixFirewallInstanceMap map[string]AviatrixFirewallInstanceInput

func (AviatrixFirewallInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixFirewallInstance)(nil)).Elem()
}

func (i AviatrixFirewallInstanceMap) ToAviatrixFirewallInstanceMapOutput() AviatrixFirewallInstanceMapOutput {
	return i.ToAviatrixFirewallInstanceMapOutputWithContext(context.Background())
}

func (i AviatrixFirewallInstanceMap) ToAviatrixFirewallInstanceMapOutputWithContext(ctx context.Context) AviatrixFirewallInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixFirewallInstanceMapOutput)
}

type AviatrixFirewallInstanceOutput struct{ *pulumi.OutputState }

func (AviatrixFirewallInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixFirewallInstance)(nil)).Elem()
}

func (o AviatrixFirewallInstanceOutput) ToAviatrixFirewallInstanceOutput() AviatrixFirewallInstanceOutput {
	return o
}

func (o AviatrixFirewallInstanceOutput) ToAviatrixFirewallInstanceOutputWithContext(ctx context.Context) AviatrixFirewallInstanceOutput {
	return o
}

// Availability domain for OCI.
func (o AviatrixFirewallInstanceOutput) AvailabilityDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.AvailabilityDomain }).(pulumi.StringOutput)
}

// Advanced option. Bootstrap bucket name. Only available for AWS and GCP.
func (o AviatrixFirewallInstanceOutput) BootstrapBucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.BootstrapBucketName }).(pulumi.StringPtrOutput)
}

// Advanced option. Bootstrap storage name. Applicable to Azure and Palo Alto Networks VM-Series/Fortinet Series deployment
// only.
func (o AviatrixFirewallInstanceOutput) BootstrapStorageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.BootstrapStorageName }).(pulumi.StringPtrOutput)
}

// Cloud Type
func (o AviatrixFirewallInstanceOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.IntOutput { return v.CloudType }).(pulumi.IntOutput)
}

// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
func (o AviatrixFirewallInstanceOutput) ContainerFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.ContainerFolder }).(pulumi.StringPtrOutput)
}

// ID of Egress Interface created.
func (o AviatrixFirewallInstanceOutput) EgressInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.EgressInterface }).(pulumi.StringOutput)
}

// Egress Interface Subnet.
func (o AviatrixFirewallInstanceOutput) EgressSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.EgressSubnet }).(pulumi.StringOutput)
}

// Egress VPC ID. Required for GCP.
func (o AviatrixFirewallInstanceOutput) EgressVpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.EgressVpcId }).(pulumi.StringPtrOutput)
}

// Fault domain for OCI.
func (o AviatrixFirewallInstanceOutput) FaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.FaultDomain }).(pulumi.StringOutput)
}

// Advanced option. File share folder. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
func (o AviatrixFirewallInstanceOutput) FileShareFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.FileShareFolder }).(pulumi.StringPtrOutput)
}

// Name of the primary FireNet gateway.
func (o AviatrixFirewallInstanceOutput) FirenetGwName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.FirenetGwName }).(pulumi.StringPtrOutput)
}

// One of the AWS AMIs from Palo Alto Networks.
func (o AviatrixFirewallInstanceOutput) FirewallImage() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.FirewallImage }).(pulumi.StringOutput)
}

// Firewall image ID.
func (o AviatrixFirewallInstanceOutput) FirewallImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.FirewallImageId }).(pulumi.StringOutput)
}

// Version of firewall image.
func (o AviatrixFirewallInstanceOutput) FirewallImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.FirewallImageVersion }).(pulumi.StringOutput)
}

// Name of the firewall instance to be created.
func (o AviatrixFirewallInstanceOutput) FirewallName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.FirewallName }).(pulumi.StringOutput)
}

// Instance size of the firewall.
func (o AviatrixFirewallInstanceOutput) FirewallSize() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.FirewallSize }).(pulumi.StringOutput)
}

// GCP VPC ID
func (o AviatrixFirewallInstanceOutput) GcpVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.GcpVpcId }).(pulumi.StringOutput)
}

// Advanced option. IAM role. Only available for AWS.
func (o AviatrixFirewallInstanceOutput) IamRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.IamRole }).(pulumi.StringPtrOutput)
}

// ID of the firewall instance created.
func (o AviatrixFirewallInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Applicable to AWS deployment only. AWS Key Pair name. If not provided, a Key Pair will be generated.
func (o AviatrixFirewallInstanceOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.KeyName }).(pulumi.StringPtrOutput)
}

// ID of Lan Interface created.
func (o AviatrixFirewallInstanceOutput) LanInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.LanInterface }).(pulumi.StringOutput)
}

// ID of Management Interface created.
func (o AviatrixFirewallInstanceOutput) ManagementInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.ManagementInterface }).(pulumi.StringOutput)
}

// Management Interface Subnet. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or
// Fortinet series.
func (o AviatrixFirewallInstanceOutput) ManagementSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.ManagementSubnet }).(pulumi.StringPtrOutput)
}

// Management VPC ID. Required for GCP Palo Alto Networks VM-Series. Required to be empty for GCP Check Point or Fortinet
// series.
func (o AviatrixFirewallInstanceOutput) ManagementVpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.ManagementVpcId }).(pulumi.StringPtrOutput)
}

// Authentication method. Applicable to Azure deployment only.
func (o AviatrixFirewallInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Management Public IP.
func (o AviatrixFirewallInstanceOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
func (o AviatrixFirewallInstanceOutput) SasUrlConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.SasUrlConfig }).(pulumi.StringPtrOutput)
}

// Advanced option. Bootstrap storage name. Applicable to Azure and Fortinet Series deployment only.
func (o AviatrixFirewallInstanceOutput) SasUrlLicense() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.SasUrlLicense }).(pulumi.StringPtrOutput)
}

// Advanced option. Share directory. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
func (o AviatrixFirewallInstanceOutput) ShareDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.ShareDirectory }).(pulumi.StringPtrOutput)
}

// Advanced option. Bic key. Applicable to Azure and Check Point Series deployment only.
func (o AviatrixFirewallInstanceOutput) SicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.SicKey }).(pulumi.StringPtrOutput)
}

// Authentication method. Applicable to Azure deployment only.
func (o AviatrixFirewallInstanceOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.SshPublicKey }).(pulumi.StringPtrOutput)
}

// Advanced option. Storage access key. Applicable to Azure and Palo Alto Networks VM-Series deployment only.
func (o AviatrixFirewallInstanceOutput) StorageAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.StorageAccessKey }).(pulumi.StringPtrOutput)
}

// A map of tags to assign to the firewall instance.
func (o AviatrixFirewallInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Advanced option. Bootstrap storage name. Applicable to Check Point Series and Fortinet Series deployment only.
func (o AviatrixFirewallInstanceOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// Applicable to Azure deployment only. 'admin' as a username is not accepted.
func (o AviatrixFirewallInstanceOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

// ID of the Security VPC.
func (o AviatrixFirewallInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Availability Zone. Only available for AWS, GCP and Azure.
func (o AviatrixFirewallInstanceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixFirewallInstance) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type AviatrixFirewallInstanceArrayOutput struct{ *pulumi.OutputState }

func (AviatrixFirewallInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixFirewallInstance)(nil)).Elem()
}

func (o AviatrixFirewallInstanceArrayOutput) ToAviatrixFirewallInstanceArrayOutput() AviatrixFirewallInstanceArrayOutput {
	return o
}

func (o AviatrixFirewallInstanceArrayOutput) ToAviatrixFirewallInstanceArrayOutputWithContext(ctx context.Context) AviatrixFirewallInstanceArrayOutput {
	return o
}

func (o AviatrixFirewallInstanceArrayOutput) Index(i pulumi.IntInput) AviatrixFirewallInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixFirewallInstance {
		return vs[0].([]*AviatrixFirewallInstance)[vs[1].(int)]
	}).(AviatrixFirewallInstanceOutput)
}

type AviatrixFirewallInstanceMapOutput struct{ *pulumi.OutputState }

func (AviatrixFirewallInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixFirewallInstance)(nil)).Elem()
}

func (o AviatrixFirewallInstanceMapOutput) ToAviatrixFirewallInstanceMapOutput() AviatrixFirewallInstanceMapOutput {
	return o
}

func (o AviatrixFirewallInstanceMapOutput) ToAviatrixFirewallInstanceMapOutputWithContext(ctx context.Context) AviatrixFirewallInstanceMapOutput {
	return o
}

func (o AviatrixFirewallInstanceMapOutput) MapIndex(k pulumi.StringInput) AviatrixFirewallInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixFirewallInstance {
		return vs[0].(map[string]*AviatrixFirewallInstance)[vs[1].(string)]
	}).(AviatrixFirewallInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirewallInstanceInput)(nil)).Elem(), &AviatrixFirewallInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirewallInstanceArrayInput)(nil)).Elem(), AviatrixFirewallInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixFirewallInstanceMapInput)(nil)).Elem(), AviatrixFirewallInstanceMap{})
	pulumi.RegisterOutputType(AviatrixFirewallInstanceOutput{})
	pulumi.RegisterOutputType(AviatrixFirewallInstanceArrayOutput{})
	pulumi.RegisterOutputType(AviatrixFirewallInstanceMapOutput{})
}