// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixSpokeVpc struct {
	pulumi.CustomResourceState

	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Cloud instance ID.
	CloudInstanceId pulumi.StringOutput `pulumi:"cloudInstanceId"`
	// Type of cloud service provider.
	CloudType pulumi.IntOutput `pulumi:"cloudType"`
	// Specify whether enabling NAT feature on the gateway or not.
	EnableNat pulumi.StringPtrOutput `pulumi:"enableNat"`
	// Name of the gateway which is going to be created.
	GwName pulumi.StringOutput `pulumi:"gwName"`
	// HA Gateway Size.
	HaGwSize pulumi.StringPtrOutput `pulumi:"haGwSize"`
	// HA Subnet. Required if enabling HA for AWS/Azure.
	HaSubnet pulumi.StringPtrOutput `pulumi:"haSubnet"`
	// HA Zone. Required if enabling HA for GCP.
	HaZone pulumi.StringPtrOutput `pulumi:"haZone"`
	// Set to 'enabled' if this feature is desired.
	SingleAzHa pulumi.StringPtrOutput `pulumi:"singleAzHa"`
	// Public Subnet Info.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Instance tag of cloud provider.
	TagLists pulumi.StringArrayOutput `pulumi:"tagLists"`
	// Specify the transit Gateway.
	TransitGw pulumi.StringPtrOutput `pulumi:"transitGw"`
	// VPC-ID/VNet-Name of cloud provider.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Region of cloud provider.
	VpcReg pulumi.StringOutput `pulumi:"vpcReg"`
	// Size of the gateway instance.
	VpcSize pulumi.StringOutput `pulumi:"vpcSize"`
}

// NewAviatrixSpokeVpc registers a new resource with the given unique name, arguments, and options.
func NewAviatrixSpokeVpc(ctx *pulumi.Context,
	name string, args *AviatrixSpokeVpcArgs, opts ...pulumi.ResourceOption) (*AviatrixSpokeVpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.CloudType == nil {
		return nil, errors.New("invalid value for required argument 'CloudType'")
	}
	if args.GwName == nil {
		return nil, errors.New("invalid value for required argument 'GwName'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VpcReg == nil {
		return nil, errors.New("invalid value for required argument 'VpcReg'")
	}
	if args.VpcSize == nil {
		return nil, errors.New("invalid value for required argument 'VpcSize'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixSpokeVpc
	err := ctx.RegisterResource("aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixSpokeVpc gets an existing AviatrixSpokeVpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixSpokeVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixSpokeVpcState, opts ...pulumi.ResourceOption) (*AviatrixSpokeVpc, error) {
	var resource AviatrixSpokeVpc
	err := ctx.ReadResource("aviatrix:index/aviatrixSpokeVpc:AviatrixSpokeVpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixSpokeVpc resources.
type aviatrixSpokeVpcState struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName *string `pulumi:"accountName"`
	// Cloud instance ID.
	CloudInstanceId *string `pulumi:"cloudInstanceId"`
	// Type of cloud service provider.
	CloudType *int `pulumi:"cloudType"`
	// Specify whether enabling NAT feature on the gateway or not.
	EnableNat *string `pulumi:"enableNat"`
	// Name of the gateway which is going to be created.
	GwName *string `pulumi:"gwName"`
	// HA Gateway Size.
	HaGwSize *string `pulumi:"haGwSize"`
	// HA Subnet. Required if enabling HA for AWS/Azure.
	HaSubnet *string `pulumi:"haSubnet"`
	// HA Zone. Required if enabling HA for GCP.
	HaZone *string `pulumi:"haZone"`
	// Set to 'enabled' if this feature is desired.
	SingleAzHa *string `pulumi:"singleAzHa"`
	// Public Subnet Info.
	Subnet *string `pulumi:"subnet"`
	// Instance tag of cloud provider.
	TagLists []string `pulumi:"tagLists"`
	// Specify the transit Gateway.
	TransitGw *string `pulumi:"transitGw"`
	// VPC-ID/VNet-Name of cloud provider.
	VpcId *string `pulumi:"vpcId"`
	// Region of cloud provider.
	VpcReg *string `pulumi:"vpcReg"`
	// Size of the gateway instance.
	VpcSize *string `pulumi:"vpcSize"`
}

type AviatrixSpokeVpcState struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringPtrInput
	// Cloud instance ID.
	CloudInstanceId pulumi.StringPtrInput
	// Type of cloud service provider.
	CloudType pulumi.IntPtrInput
	// Specify whether enabling NAT feature on the gateway or not.
	EnableNat pulumi.StringPtrInput
	// Name of the gateway which is going to be created.
	GwName pulumi.StringPtrInput
	// HA Gateway Size.
	HaGwSize pulumi.StringPtrInput
	// HA Subnet. Required if enabling HA for AWS/Azure.
	HaSubnet pulumi.StringPtrInput
	// HA Zone. Required if enabling HA for GCP.
	HaZone pulumi.StringPtrInput
	// Set to 'enabled' if this feature is desired.
	SingleAzHa pulumi.StringPtrInput
	// Public Subnet Info.
	Subnet pulumi.StringPtrInput
	// Instance tag of cloud provider.
	TagLists pulumi.StringArrayInput
	// Specify the transit Gateway.
	TransitGw pulumi.StringPtrInput
	// VPC-ID/VNet-Name of cloud provider.
	VpcId pulumi.StringPtrInput
	// Region of cloud provider.
	VpcReg pulumi.StringPtrInput
	// Size of the gateway instance.
	VpcSize pulumi.StringPtrInput
}

func (AviatrixSpokeVpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixSpokeVpcState)(nil)).Elem()
}

type aviatrixSpokeVpcArgs struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName string `pulumi:"accountName"`
	// Type of cloud service provider.
	CloudType int `pulumi:"cloudType"`
	// Specify whether enabling NAT feature on the gateway or not.
	EnableNat *string `pulumi:"enableNat"`
	// Name of the gateway which is going to be created.
	GwName string `pulumi:"gwName"`
	// HA Gateway Size.
	HaGwSize *string `pulumi:"haGwSize"`
	// HA Subnet. Required if enabling HA for AWS/Azure.
	HaSubnet *string `pulumi:"haSubnet"`
	// HA Zone. Required if enabling HA for GCP.
	HaZone *string `pulumi:"haZone"`
	// Set to 'enabled' if this feature is desired.
	SingleAzHa *string `pulumi:"singleAzHa"`
	// Public Subnet Info.
	Subnet string `pulumi:"subnet"`
	// Instance tag of cloud provider.
	TagLists []string `pulumi:"tagLists"`
	// Specify the transit Gateway.
	TransitGw *string `pulumi:"transitGw"`
	// VPC-ID/VNet-Name of cloud provider.
	VpcId string `pulumi:"vpcId"`
	// Region of cloud provider.
	VpcReg string `pulumi:"vpcReg"`
	// Size of the gateway instance.
	VpcSize string `pulumi:"vpcSize"`
}

// The set of arguments for constructing a AviatrixSpokeVpc resource.
type AviatrixSpokeVpcArgs struct {
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	AccountName pulumi.StringInput
	// Type of cloud service provider.
	CloudType pulumi.IntInput
	// Specify whether enabling NAT feature on the gateway or not.
	EnableNat pulumi.StringPtrInput
	// Name of the gateway which is going to be created.
	GwName pulumi.StringInput
	// HA Gateway Size.
	HaGwSize pulumi.StringPtrInput
	// HA Subnet. Required if enabling HA for AWS/Azure.
	HaSubnet pulumi.StringPtrInput
	// HA Zone. Required if enabling HA for GCP.
	HaZone pulumi.StringPtrInput
	// Set to 'enabled' if this feature is desired.
	SingleAzHa pulumi.StringPtrInput
	// Public Subnet Info.
	Subnet pulumi.StringInput
	// Instance tag of cloud provider.
	TagLists pulumi.StringArrayInput
	// Specify the transit Gateway.
	TransitGw pulumi.StringPtrInput
	// VPC-ID/VNet-Name of cloud provider.
	VpcId pulumi.StringInput
	// Region of cloud provider.
	VpcReg pulumi.StringInput
	// Size of the gateway instance.
	VpcSize pulumi.StringInput
}

func (AviatrixSpokeVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixSpokeVpcArgs)(nil)).Elem()
}

type AviatrixSpokeVpcInput interface {
	pulumi.Input

	ToAviatrixSpokeVpcOutput() AviatrixSpokeVpcOutput
	ToAviatrixSpokeVpcOutputWithContext(ctx context.Context) AviatrixSpokeVpcOutput
}

func (*AviatrixSpokeVpc) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixSpokeVpc)(nil)).Elem()
}

func (i *AviatrixSpokeVpc) ToAviatrixSpokeVpcOutput() AviatrixSpokeVpcOutput {
	return i.ToAviatrixSpokeVpcOutputWithContext(context.Background())
}

func (i *AviatrixSpokeVpc) ToAviatrixSpokeVpcOutputWithContext(ctx context.Context) AviatrixSpokeVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSpokeVpcOutput)
}

// AviatrixSpokeVpcArrayInput is an input type that accepts AviatrixSpokeVpcArray and AviatrixSpokeVpcArrayOutput values.
// You can construct a concrete instance of `AviatrixSpokeVpcArrayInput` via:
//
//	AviatrixSpokeVpcArray{ AviatrixSpokeVpcArgs{...} }
type AviatrixSpokeVpcArrayInput interface {
	pulumi.Input

	ToAviatrixSpokeVpcArrayOutput() AviatrixSpokeVpcArrayOutput
	ToAviatrixSpokeVpcArrayOutputWithContext(context.Context) AviatrixSpokeVpcArrayOutput
}

type AviatrixSpokeVpcArray []AviatrixSpokeVpcInput

func (AviatrixSpokeVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixSpokeVpc)(nil)).Elem()
}

func (i AviatrixSpokeVpcArray) ToAviatrixSpokeVpcArrayOutput() AviatrixSpokeVpcArrayOutput {
	return i.ToAviatrixSpokeVpcArrayOutputWithContext(context.Background())
}

func (i AviatrixSpokeVpcArray) ToAviatrixSpokeVpcArrayOutputWithContext(ctx context.Context) AviatrixSpokeVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSpokeVpcArrayOutput)
}

// AviatrixSpokeVpcMapInput is an input type that accepts AviatrixSpokeVpcMap and AviatrixSpokeVpcMapOutput values.
// You can construct a concrete instance of `AviatrixSpokeVpcMapInput` via:
//
//	AviatrixSpokeVpcMap{ "key": AviatrixSpokeVpcArgs{...} }
type AviatrixSpokeVpcMapInput interface {
	pulumi.Input

	ToAviatrixSpokeVpcMapOutput() AviatrixSpokeVpcMapOutput
	ToAviatrixSpokeVpcMapOutputWithContext(context.Context) AviatrixSpokeVpcMapOutput
}

type AviatrixSpokeVpcMap map[string]AviatrixSpokeVpcInput

func (AviatrixSpokeVpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixSpokeVpc)(nil)).Elem()
}

func (i AviatrixSpokeVpcMap) ToAviatrixSpokeVpcMapOutput() AviatrixSpokeVpcMapOutput {
	return i.ToAviatrixSpokeVpcMapOutputWithContext(context.Background())
}

func (i AviatrixSpokeVpcMap) ToAviatrixSpokeVpcMapOutputWithContext(ctx context.Context) AviatrixSpokeVpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSpokeVpcMapOutput)
}

type AviatrixSpokeVpcOutput struct{ *pulumi.OutputState }

func (AviatrixSpokeVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixSpokeVpc)(nil)).Elem()
}

func (o AviatrixSpokeVpcOutput) ToAviatrixSpokeVpcOutput() AviatrixSpokeVpcOutput {
	return o
}

func (o AviatrixSpokeVpcOutput) ToAviatrixSpokeVpcOutputWithContext(ctx context.Context) AviatrixSpokeVpcOutput {
	return o
}

// This parameter represents the name of a Cloud-Account in Aviatrix controller.
func (o AviatrixSpokeVpcOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// Cloud instance ID.
func (o AviatrixSpokeVpcOutput) CloudInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringOutput { return v.CloudInstanceId }).(pulumi.StringOutput)
}

// Type of cloud service provider.
func (o AviatrixSpokeVpcOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.IntOutput { return v.CloudType }).(pulumi.IntOutput)
}

// Specify whether enabling NAT feature on the gateway or not.
func (o AviatrixSpokeVpcOutput) EnableNat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringPtrOutput { return v.EnableNat }).(pulumi.StringPtrOutput)
}

// Name of the gateway which is going to be created.
func (o AviatrixSpokeVpcOutput) GwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringOutput { return v.GwName }).(pulumi.StringOutput)
}

// HA Gateway Size.
func (o AviatrixSpokeVpcOutput) HaGwSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringPtrOutput { return v.HaGwSize }).(pulumi.StringPtrOutput)
}

// HA Subnet. Required if enabling HA for AWS/Azure.
func (o AviatrixSpokeVpcOutput) HaSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringPtrOutput { return v.HaSubnet }).(pulumi.StringPtrOutput)
}

// HA Zone. Required if enabling HA for GCP.
func (o AviatrixSpokeVpcOutput) HaZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringPtrOutput { return v.HaZone }).(pulumi.StringPtrOutput)
}

// Set to 'enabled' if this feature is desired.
func (o AviatrixSpokeVpcOutput) SingleAzHa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringPtrOutput { return v.SingleAzHa }).(pulumi.StringPtrOutput)
}

// Public Subnet Info.
func (o AviatrixSpokeVpcOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Instance tag of cloud provider.
func (o AviatrixSpokeVpcOutput) TagLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringArrayOutput { return v.TagLists }).(pulumi.StringArrayOutput)
}

// Specify the transit Gateway.
func (o AviatrixSpokeVpcOutput) TransitGw() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringPtrOutput { return v.TransitGw }).(pulumi.StringPtrOutput)
}

// VPC-ID/VNet-Name of cloud provider.
func (o AviatrixSpokeVpcOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Region of cloud provider.
func (o AviatrixSpokeVpcOutput) VpcReg() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringOutput { return v.VpcReg }).(pulumi.StringOutput)
}

// Size of the gateway instance.
func (o AviatrixSpokeVpcOutput) VpcSize() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSpokeVpc) pulumi.StringOutput { return v.VpcSize }).(pulumi.StringOutput)
}

type AviatrixSpokeVpcArrayOutput struct{ *pulumi.OutputState }

func (AviatrixSpokeVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixSpokeVpc)(nil)).Elem()
}

func (o AviatrixSpokeVpcArrayOutput) ToAviatrixSpokeVpcArrayOutput() AviatrixSpokeVpcArrayOutput {
	return o
}

func (o AviatrixSpokeVpcArrayOutput) ToAviatrixSpokeVpcArrayOutputWithContext(ctx context.Context) AviatrixSpokeVpcArrayOutput {
	return o
}

func (o AviatrixSpokeVpcArrayOutput) Index(i pulumi.IntInput) AviatrixSpokeVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixSpokeVpc {
		return vs[0].([]*AviatrixSpokeVpc)[vs[1].(int)]
	}).(AviatrixSpokeVpcOutput)
}

type AviatrixSpokeVpcMapOutput struct{ *pulumi.OutputState }

func (AviatrixSpokeVpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixSpokeVpc)(nil)).Elem()
}

func (o AviatrixSpokeVpcMapOutput) ToAviatrixSpokeVpcMapOutput() AviatrixSpokeVpcMapOutput {
	return o
}

func (o AviatrixSpokeVpcMapOutput) ToAviatrixSpokeVpcMapOutputWithContext(ctx context.Context) AviatrixSpokeVpcMapOutput {
	return o
}

func (o AviatrixSpokeVpcMapOutput) MapIndex(k pulumi.StringInput) AviatrixSpokeVpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixSpokeVpc {
		return vs[0].(map[string]*AviatrixSpokeVpc)[vs[1].(string)]
	}).(AviatrixSpokeVpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSpokeVpcInput)(nil)).Elem(), &AviatrixSpokeVpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSpokeVpcArrayInput)(nil)).Elem(), AviatrixSpokeVpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSpokeVpcMapInput)(nil)).Elem(), AviatrixSpokeVpcMap{})
	pulumi.RegisterOutputType(AviatrixSpokeVpcOutput{})
	pulumi.RegisterOutputType(AviatrixSpokeVpcArrayOutput{})
	pulumi.RegisterOutputType(AviatrixSpokeVpcMapOutput{})
}