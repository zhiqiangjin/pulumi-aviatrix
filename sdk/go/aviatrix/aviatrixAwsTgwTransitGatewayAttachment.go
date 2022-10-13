// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixAwsTgwTransitGatewayAttachment struct {
	pulumi.CustomResourceState

	// Region of cloud provider.
	Region pulumi.StringOutput `pulumi:"region"`
	// Name of the AWS TGW.
	TgwName pulumi.StringOutput `pulumi:"tgwName"`
	// Name of the transit gateway to be attached to tgw.
	TransitGatewayName pulumi.StringOutput `pulumi:"transitGatewayName"`
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	VpcAccountName pulumi.StringOutput `pulumi:"vpcAccountName"`
	// This parameter represents the ID of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewAviatrixAwsTgwTransitGatewayAttachment registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAwsTgwTransitGatewayAttachment(ctx *pulumi.Context,
	name string, args *AviatrixAwsTgwTransitGatewayAttachmentArgs, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwTransitGatewayAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.TgwName == nil {
		return nil, errors.New("invalid value for required argument 'TgwName'")
	}
	if args.TransitGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayName'")
	}
	if args.VpcAccountName == nil {
		return nil, errors.New("invalid value for required argument 'VpcAccountName'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAwsTgwTransitGatewayAttachment
	err := ctx.RegisterResource("aviatrix:index/aviatrixAwsTgwTransitGatewayAttachment:AviatrixAwsTgwTransitGatewayAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAwsTgwTransitGatewayAttachment gets an existing AviatrixAwsTgwTransitGatewayAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAwsTgwTransitGatewayAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAwsTgwTransitGatewayAttachmentState, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwTransitGatewayAttachment, error) {
	var resource AviatrixAwsTgwTransitGatewayAttachment
	err := ctx.ReadResource("aviatrix:index/aviatrixAwsTgwTransitGatewayAttachment:AviatrixAwsTgwTransitGatewayAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAwsTgwTransitGatewayAttachment resources.
type aviatrixAwsTgwTransitGatewayAttachmentState struct {
	// Region of cloud provider.
	Region *string `pulumi:"region"`
	// Name of the AWS TGW.
	TgwName *string `pulumi:"tgwName"`
	// Name of the transit gateway to be attached to tgw.
	TransitGatewayName *string `pulumi:"transitGatewayName"`
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	VpcAccountName *string `pulumi:"vpcAccountName"`
	// This parameter represents the ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

type AviatrixAwsTgwTransitGatewayAttachmentState struct {
	// Region of cloud provider.
	Region pulumi.StringPtrInput
	// Name of the AWS TGW.
	TgwName pulumi.StringPtrInput
	// Name of the transit gateway to be attached to tgw.
	TransitGatewayName pulumi.StringPtrInput
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	VpcAccountName pulumi.StringPtrInput
	// This parameter represents the ID of the VPC.
	VpcId pulumi.StringPtrInput
}

func (AviatrixAwsTgwTransitGatewayAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwTransitGatewayAttachmentState)(nil)).Elem()
}

type aviatrixAwsTgwTransitGatewayAttachmentArgs struct {
	// Region of cloud provider.
	Region string `pulumi:"region"`
	// Name of the AWS TGW.
	TgwName string `pulumi:"tgwName"`
	// Name of the transit gateway to be attached to tgw.
	TransitGatewayName string `pulumi:"transitGatewayName"`
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	VpcAccountName string `pulumi:"vpcAccountName"`
	// This parameter represents the ID of the VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a AviatrixAwsTgwTransitGatewayAttachment resource.
type AviatrixAwsTgwTransitGatewayAttachmentArgs struct {
	// Region of cloud provider.
	Region pulumi.StringInput
	// Name of the AWS TGW.
	TgwName pulumi.StringInput
	// Name of the transit gateway to be attached to tgw.
	TransitGatewayName pulumi.StringInput
	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	VpcAccountName pulumi.StringInput
	// This parameter represents the ID of the VPC.
	VpcId pulumi.StringInput
}

func (AviatrixAwsTgwTransitGatewayAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwTransitGatewayAttachmentArgs)(nil)).Elem()
}

type AviatrixAwsTgwTransitGatewayAttachmentInput interface {
	pulumi.Input

	ToAviatrixAwsTgwTransitGatewayAttachmentOutput() AviatrixAwsTgwTransitGatewayAttachmentOutput
	ToAviatrixAwsTgwTransitGatewayAttachmentOutputWithContext(ctx context.Context) AviatrixAwsTgwTransitGatewayAttachmentOutput
}

func (*AviatrixAwsTgwTransitGatewayAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwTransitGatewayAttachment)(nil)).Elem()
}

func (i *AviatrixAwsTgwTransitGatewayAttachment) ToAviatrixAwsTgwTransitGatewayAttachmentOutput() AviatrixAwsTgwTransitGatewayAttachmentOutput {
	return i.ToAviatrixAwsTgwTransitGatewayAttachmentOutputWithContext(context.Background())
}

func (i *AviatrixAwsTgwTransitGatewayAttachment) ToAviatrixAwsTgwTransitGatewayAttachmentOutputWithContext(ctx context.Context) AviatrixAwsTgwTransitGatewayAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwTransitGatewayAttachmentOutput)
}

// AviatrixAwsTgwTransitGatewayAttachmentArrayInput is an input type that accepts AviatrixAwsTgwTransitGatewayAttachmentArray and AviatrixAwsTgwTransitGatewayAttachmentArrayOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwTransitGatewayAttachmentArrayInput` via:
//
//	AviatrixAwsTgwTransitGatewayAttachmentArray{ AviatrixAwsTgwTransitGatewayAttachmentArgs{...} }
type AviatrixAwsTgwTransitGatewayAttachmentArrayInput interface {
	pulumi.Input

	ToAviatrixAwsTgwTransitGatewayAttachmentArrayOutput() AviatrixAwsTgwTransitGatewayAttachmentArrayOutput
	ToAviatrixAwsTgwTransitGatewayAttachmentArrayOutputWithContext(context.Context) AviatrixAwsTgwTransitGatewayAttachmentArrayOutput
}

type AviatrixAwsTgwTransitGatewayAttachmentArray []AviatrixAwsTgwTransitGatewayAttachmentInput

func (AviatrixAwsTgwTransitGatewayAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwTransitGatewayAttachment)(nil)).Elem()
}

func (i AviatrixAwsTgwTransitGatewayAttachmentArray) ToAviatrixAwsTgwTransitGatewayAttachmentArrayOutput() AviatrixAwsTgwTransitGatewayAttachmentArrayOutput {
	return i.ToAviatrixAwsTgwTransitGatewayAttachmentArrayOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwTransitGatewayAttachmentArray) ToAviatrixAwsTgwTransitGatewayAttachmentArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwTransitGatewayAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwTransitGatewayAttachmentArrayOutput)
}

// AviatrixAwsTgwTransitGatewayAttachmentMapInput is an input type that accepts AviatrixAwsTgwTransitGatewayAttachmentMap and AviatrixAwsTgwTransitGatewayAttachmentMapOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwTransitGatewayAttachmentMapInput` via:
//
//	AviatrixAwsTgwTransitGatewayAttachmentMap{ "key": AviatrixAwsTgwTransitGatewayAttachmentArgs{...} }
type AviatrixAwsTgwTransitGatewayAttachmentMapInput interface {
	pulumi.Input

	ToAviatrixAwsTgwTransitGatewayAttachmentMapOutput() AviatrixAwsTgwTransitGatewayAttachmentMapOutput
	ToAviatrixAwsTgwTransitGatewayAttachmentMapOutputWithContext(context.Context) AviatrixAwsTgwTransitGatewayAttachmentMapOutput
}

type AviatrixAwsTgwTransitGatewayAttachmentMap map[string]AviatrixAwsTgwTransitGatewayAttachmentInput

func (AviatrixAwsTgwTransitGatewayAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwTransitGatewayAttachment)(nil)).Elem()
}

func (i AviatrixAwsTgwTransitGatewayAttachmentMap) ToAviatrixAwsTgwTransitGatewayAttachmentMapOutput() AviatrixAwsTgwTransitGatewayAttachmentMapOutput {
	return i.ToAviatrixAwsTgwTransitGatewayAttachmentMapOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwTransitGatewayAttachmentMap) ToAviatrixAwsTgwTransitGatewayAttachmentMapOutputWithContext(ctx context.Context) AviatrixAwsTgwTransitGatewayAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwTransitGatewayAttachmentMapOutput)
}

type AviatrixAwsTgwTransitGatewayAttachmentOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwTransitGatewayAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwTransitGatewayAttachment)(nil)).Elem()
}

func (o AviatrixAwsTgwTransitGatewayAttachmentOutput) ToAviatrixAwsTgwTransitGatewayAttachmentOutput() AviatrixAwsTgwTransitGatewayAttachmentOutput {
	return o
}

func (o AviatrixAwsTgwTransitGatewayAttachmentOutput) ToAviatrixAwsTgwTransitGatewayAttachmentOutputWithContext(ctx context.Context) AviatrixAwsTgwTransitGatewayAttachmentOutput {
	return o
}

// Region of cloud provider.
func (o AviatrixAwsTgwTransitGatewayAttachmentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwTransitGatewayAttachment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Name of the AWS TGW.
func (o AviatrixAwsTgwTransitGatewayAttachmentOutput) TgwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwTransitGatewayAttachment) pulumi.StringOutput { return v.TgwName }).(pulumi.StringOutput)
}

// Name of the transit gateway to be attached to tgw.
func (o AviatrixAwsTgwTransitGatewayAttachmentOutput) TransitGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwTransitGatewayAttachment) pulumi.StringOutput { return v.TransitGatewayName }).(pulumi.StringOutput)
}

// This parameter represents the name of a Cloud-Account in Aviatrix controller.
func (o AviatrixAwsTgwTransitGatewayAttachmentOutput) VpcAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwTransitGatewayAttachment) pulumi.StringOutput { return v.VpcAccountName }).(pulumi.StringOutput)
}

// This parameter represents the ID of the VPC.
func (o AviatrixAwsTgwTransitGatewayAttachmentOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwTransitGatewayAttachment) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type AviatrixAwsTgwTransitGatewayAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwTransitGatewayAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwTransitGatewayAttachment)(nil)).Elem()
}

func (o AviatrixAwsTgwTransitGatewayAttachmentArrayOutput) ToAviatrixAwsTgwTransitGatewayAttachmentArrayOutput() AviatrixAwsTgwTransitGatewayAttachmentArrayOutput {
	return o
}

func (o AviatrixAwsTgwTransitGatewayAttachmentArrayOutput) ToAviatrixAwsTgwTransitGatewayAttachmentArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwTransitGatewayAttachmentArrayOutput {
	return o
}

func (o AviatrixAwsTgwTransitGatewayAttachmentArrayOutput) Index(i pulumi.IntInput) AviatrixAwsTgwTransitGatewayAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAwsTgwTransitGatewayAttachment {
		return vs[0].([]*AviatrixAwsTgwTransitGatewayAttachment)[vs[1].(int)]
	}).(AviatrixAwsTgwTransitGatewayAttachmentOutput)
}

type AviatrixAwsTgwTransitGatewayAttachmentMapOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwTransitGatewayAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwTransitGatewayAttachment)(nil)).Elem()
}

func (o AviatrixAwsTgwTransitGatewayAttachmentMapOutput) ToAviatrixAwsTgwTransitGatewayAttachmentMapOutput() AviatrixAwsTgwTransitGatewayAttachmentMapOutput {
	return o
}

func (o AviatrixAwsTgwTransitGatewayAttachmentMapOutput) ToAviatrixAwsTgwTransitGatewayAttachmentMapOutputWithContext(ctx context.Context) AviatrixAwsTgwTransitGatewayAttachmentMapOutput {
	return o
}

func (o AviatrixAwsTgwTransitGatewayAttachmentMapOutput) MapIndex(k pulumi.StringInput) AviatrixAwsTgwTransitGatewayAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAwsTgwTransitGatewayAttachment {
		return vs[0].(map[string]*AviatrixAwsTgwTransitGatewayAttachment)[vs[1].(string)]
	}).(AviatrixAwsTgwTransitGatewayAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwTransitGatewayAttachmentInput)(nil)).Elem(), &AviatrixAwsTgwTransitGatewayAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwTransitGatewayAttachmentArrayInput)(nil)).Elem(), AviatrixAwsTgwTransitGatewayAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwTransitGatewayAttachmentMapInput)(nil)).Elem(), AviatrixAwsTgwTransitGatewayAttachmentMap{})
	pulumi.RegisterOutputType(AviatrixAwsTgwTransitGatewayAttachmentOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwTransitGatewayAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwTransitGatewayAttachmentMapOutput{})
}
