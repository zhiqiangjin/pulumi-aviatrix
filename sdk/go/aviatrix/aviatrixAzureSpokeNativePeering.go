// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixAzureSpokeNativePeering struct {
	pulumi.CustomResourceState

	// An Aviatrix account that corresponds to a subscription in Azure.
	SpokeAccountName pulumi.StringOutput `pulumi:"spokeAccountName"`
	// Spoke VNet region.
	SpokeRegion pulumi.StringOutput `pulumi:"spokeRegion"`
	// Combination of the Spoke VNet name and resource group.
	SpokeVpcId pulumi.StringOutput `pulumi:"spokeVpcId"`
	// Name of an azure transit gateway with transit firenet enabled.
	TransitGatewayName pulumi.StringOutput `pulumi:"transitGatewayName"`
}

// NewAviatrixAzureSpokeNativePeering registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAzureSpokeNativePeering(ctx *pulumi.Context,
	name string, args *AviatrixAzureSpokeNativePeeringArgs, opts ...pulumi.ResourceOption) (*AviatrixAzureSpokeNativePeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SpokeAccountName == nil {
		return nil, errors.New("invalid value for required argument 'SpokeAccountName'")
	}
	if args.SpokeRegion == nil {
		return nil, errors.New("invalid value for required argument 'SpokeRegion'")
	}
	if args.SpokeVpcId == nil {
		return nil, errors.New("invalid value for required argument 'SpokeVpcId'")
	}
	if args.TransitGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAzureSpokeNativePeering
	err := ctx.RegisterResource("aviatrix:index/aviatrixAzureSpokeNativePeering:AviatrixAzureSpokeNativePeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAzureSpokeNativePeering gets an existing AviatrixAzureSpokeNativePeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAzureSpokeNativePeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAzureSpokeNativePeeringState, opts ...pulumi.ResourceOption) (*AviatrixAzureSpokeNativePeering, error) {
	var resource AviatrixAzureSpokeNativePeering
	err := ctx.ReadResource("aviatrix:index/aviatrixAzureSpokeNativePeering:AviatrixAzureSpokeNativePeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAzureSpokeNativePeering resources.
type aviatrixAzureSpokeNativePeeringState struct {
	// An Aviatrix account that corresponds to a subscription in Azure.
	SpokeAccountName *string `pulumi:"spokeAccountName"`
	// Spoke VNet region.
	SpokeRegion *string `pulumi:"spokeRegion"`
	// Combination of the Spoke VNet name and resource group.
	SpokeVpcId *string `pulumi:"spokeVpcId"`
	// Name of an azure transit gateway with transit firenet enabled.
	TransitGatewayName *string `pulumi:"transitGatewayName"`
}

type AviatrixAzureSpokeNativePeeringState struct {
	// An Aviatrix account that corresponds to a subscription in Azure.
	SpokeAccountName pulumi.StringPtrInput
	// Spoke VNet region.
	SpokeRegion pulumi.StringPtrInput
	// Combination of the Spoke VNet name and resource group.
	SpokeVpcId pulumi.StringPtrInput
	// Name of an azure transit gateway with transit firenet enabled.
	TransitGatewayName pulumi.StringPtrInput
}

func (AviatrixAzureSpokeNativePeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAzureSpokeNativePeeringState)(nil)).Elem()
}

type aviatrixAzureSpokeNativePeeringArgs struct {
	// An Aviatrix account that corresponds to a subscription in Azure.
	SpokeAccountName string `pulumi:"spokeAccountName"`
	// Spoke VNet region.
	SpokeRegion string `pulumi:"spokeRegion"`
	// Combination of the Spoke VNet name and resource group.
	SpokeVpcId string `pulumi:"spokeVpcId"`
	// Name of an azure transit gateway with transit firenet enabled.
	TransitGatewayName string `pulumi:"transitGatewayName"`
}

// The set of arguments for constructing a AviatrixAzureSpokeNativePeering resource.
type AviatrixAzureSpokeNativePeeringArgs struct {
	// An Aviatrix account that corresponds to a subscription in Azure.
	SpokeAccountName pulumi.StringInput
	// Spoke VNet region.
	SpokeRegion pulumi.StringInput
	// Combination of the Spoke VNet name and resource group.
	SpokeVpcId pulumi.StringInput
	// Name of an azure transit gateway with transit firenet enabled.
	TransitGatewayName pulumi.StringInput
}

func (AviatrixAzureSpokeNativePeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAzureSpokeNativePeeringArgs)(nil)).Elem()
}

type AviatrixAzureSpokeNativePeeringInput interface {
	pulumi.Input

	ToAviatrixAzureSpokeNativePeeringOutput() AviatrixAzureSpokeNativePeeringOutput
	ToAviatrixAzureSpokeNativePeeringOutputWithContext(ctx context.Context) AviatrixAzureSpokeNativePeeringOutput
}

func (*AviatrixAzureSpokeNativePeering) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAzureSpokeNativePeering)(nil)).Elem()
}

func (i *AviatrixAzureSpokeNativePeering) ToAviatrixAzureSpokeNativePeeringOutput() AviatrixAzureSpokeNativePeeringOutput {
	return i.ToAviatrixAzureSpokeNativePeeringOutputWithContext(context.Background())
}

func (i *AviatrixAzureSpokeNativePeering) ToAviatrixAzureSpokeNativePeeringOutputWithContext(ctx context.Context) AviatrixAzureSpokeNativePeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAzureSpokeNativePeeringOutput)
}

// AviatrixAzureSpokeNativePeeringArrayInput is an input type that accepts AviatrixAzureSpokeNativePeeringArray and AviatrixAzureSpokeNativePeeringArrayOutput values.
// You can construct a concrete instance of `AviatrixAzureSpokeNativePeeringArrayInput` via:
//
//	AviatrixAzureSpokeNativePeeringArray{ AviatrixAzureSpokeNativePeeringArgs{...} }
type AviatrixAzureSpokeNativePeeringArrayInput interface {
	pulumi.Input

	ToAviatrixAzureSpokeNativePeeringArrayOutput() AviatrixAzureSpokeNativePeeringArrayOutput
	ToAviatrixAzureSpokeNativePeeringArrayOutputWithContext(context.Context) AviatrixAzureSpokeNativePeeringArrayOutput
}

type AviatrixAzureSpokeNativePeeringArray []AviatrixAzureSpokeNativePeeringInput

func (AviatrixAzureSpokeNativePeeringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAzureSpokeNativePeering)(nil)).Elem()
}

func (i AviatrixAzureSpokeNativePeeringArray) ToAviatrixAzureSpokeNativePeeringArrayOutput() AviatrixAzureSpokeNativePeeringArrayOutput {
	return i.ToAviatrixAzureSpokeNativePeeringArrayOutputWithContext(context.Background())
}

func (i AviatrixAzureSpokeNativePeeringArray) ToAviatrixAzureSpokeNativePeeringArrayOutputWithContext(ctx context.Context) AviatrixAzureSpokeNativePeeringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAzureSpokeNativePeeringArrayOutput)
}

// AviatrixAzureSpokeNativePeeringMapInput is an input type that accepts AviatrixAzureSpokeNativePeeringMap and AviatrixAzureSpokeNativePeeringMapOutput values.
// You can construct a concrete instance of `AviatrixAzureSpokeNativePeeringMapInput` via:
//
//	AviatrixAzureSpokeNativePeeringMap{ "key": AviatrixAzureSpokeNativePeeringArgs{...} }
type AviatrixAzureSpokeNativePeeringMapInput interface {
	pulumi.Input

	ToAviatrixAzureSpokeNativePeeringMapOutput() AviatrixAzureSpokeNativePeeringMapOutput
	ToAviatrixAzureSpokeNativePeeringMapOutputWithContext(context.Context) AviatrixAzureSpokeNativePeeringMapOutput
}

type AviatrixAzureSpokeNativePeeringMap map[string]AviatrixAzureSpokeNativePeeringInput

func (AviatrixAzureSpokeNativePeeringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAzureSpokeNativePeering)(nil)).Elem()
}

func (i AviatrixAzureSpokeNativePeeringMap) ToAviatrixAzureSpokeNativePeeringMapOutput() AviatrixAzureSpokeNativePeeringMapOutput {
	return i.ToAviatrixAzureSpokeNativePeeringMapOutputWithContext(context.Background())
}

func (i AviatrixAzureSpokeNativePeeringMap) ToAviatrixAzureSpokeNativePeeringMapOutputWithContext(ctx context.Context) AviatrixAzureSpokeNativePeeringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAzureSpokeNativePeeringMapOutput)
}

type AviatrixAzureSpokeNativePeeringOutput struct{ *pulumi.OutputState }

func (AviatrixAzureSpokeNativePeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAzureSpokeNativePeering)(nil)).Elem()
}

func (o AviatrixAzureSpokeNativePeeringOutput) ToAviatrixAzureSpokeNativePeeringOutput() AviatrixAzureSpokeNativePeeringOutput {
	return o
}

func (o AviatrixAzureSpokeNativePeeringOutput) ToAviatrixAzureSpokeNativePeeringOutputWithContext(ctx context.Context) AviatrixAzureSpokeNativePeeringOutput {
	return o
}

// An Aviatrix account that corresponds to a subscription in Azure.
func (o AviatrixAzureSpokeNativePeeringOutput) SpokeAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzureSpokeNativePeering) pulumi.StringOutput { return v.SpokeAccountName }).(pulumi.StringOutput)
}

// Spoke VNet region.
func (o AviatrixAzureSpokeNativePeeringOutput) SpokeRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzureSpokeNativePeering) pulumi.StringOutput { return v.SpokeRegion }).(pulumi.StringOutput)
}

// Combination of the Spoke VNet name and resource group.
func (o AviatrixAzureSpokeNativePeeringOutput) SpokeVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzureSpokeNativePeering) pulumi.StringOutput { return v.SpokeVpcId }).(pulumi.StringOutput)
}

// Name of an azure transit gateway with transit firenet enabled.
func (o AviatrixAzureSpokeNativePeeringOutput) TransitGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAzureSpokeNativePeering) pulumi.StringOutput { return v.TransitGatewayName }).(pulumi.StringOutput)
}

type AviatrixAzureSpokeNativePeeringArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAzureSpokeNativePeeringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAzureSpokeNativePeering)(nil)).Elem()
}

func (o AviatrixAzureSpokeNativePeeringArrayOutput) ToAviatrixAzureSpokeNativePeeringArrayOutput() AviatrixAzureSpokeNativePeeringArrayOutput {
	return o
}

func (o AviatrixAzureSpokeNativePeeringArrayOutput) ToAviatrixAzureSpokeNativePeeringArrayOutputWithContext(ctx context.Context) AviatrixAzureSpokeNativePeeringArrayOutput {
	return o
}

func (o AviatrixAzureSpokeNativePeeringArrayOutput) Index(i pulumi.IntInput) AviatrixAzureSpokeNativePeeringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAzureSpokeNativePeering {
		return vs[0].([]*AviatrixAzureSpokeNativePeering)[vs[1].(int)]
	}).(AviatrixAzureSpokeNativePeeringOutput)
}

type AviatrixAzureSpokeNativePeeringMapOutput struct{ *pulumi.OutputState }

func (AviatrixAzureSpokeNativePeeringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAzureSpokeNativePeering)(nil)).Elem()
}

func (o AviatrixAzureSpokeNativePeeringMapOutput) ToAviatrixAzureSpokeNativePeeringMapOutput() AviatrixAzureSpokeNativePeeringMapOutput {
	return o
}

func (o AviatrixAzureSpokeNativePeeringMapOutput) ToAviatrixAzureSpokeNativePeeringMapOutputWithContext(ctx context.Context) AviatrixAzureSpokeNativePeeringMapOutput {
	return o
}

func (o AviatrixAzureSpokeNativePeeringMapOutput) MapIndex(k pulumi.StringInput) AviatrixAzureSpokeNativePeeringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAzureSpokeNativePeering {
		return vs[0].(map[string]*AviatrixAzureSpokeNativePeering)[vs[1].(string)]
	}).(AviatrixAzureSpokeNativePeeringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAzureSpokeNativePeeringInput)(nil)).Elem(), &AviatrixAzureSpokeNativePeering{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAzureSpokeNativePeeringArrayInput)(nil)).Elem(), AviatrixAzureSpokeNativePeeringArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAzureSpokeNativePeeringMapInput)(nil)).Elem(), AviatrixAzureSpokeNativePeeringMap{})
	pulumi.RegisterOutputType(AviatrixAzureSpokeNativePeeringOutput{})
	pulumi.RegisterOutputType(AviatrixAzureSpokeNativePeeringArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAzureSpokeNativePeeringMapOutput{})
}