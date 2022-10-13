// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixAwsTgwIntraDomainInspection struct {
	pulumi.CustomResourceState

	// Firewall domain name.
	FirewallDomainName pulumi.StringOutput `pulumi:"firewallDomainName"`
	// Route domain name.
	RouteDomainName pulumi.StringOutput `pulumi:"routeDomainName"`
	// AWS TGW name.
	TgwName pulumi.StringOutput `pulumi:"tgwName"`
}

// NewAviatrixAwsTgwIntraDomainInspection registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAwsTgwIntraDomainInspection(ctx *pulumi.Context,
	name string, args *AviatrixAwsTgwIntraDomainInspectionArgs, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwIntraDomainInspection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallDomainName == nil {
		return nil, errors.New("invalid value for required argument 'FirewallDomainName'")
	}
	if args.RouteDomainName == nil {
		return nil, errors.New("invalid value for required argument 'RouteDomainName'")
	}
	if args.TgwName == nil {
		return nil, errors.New("invalid value for required argument 'TgwName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAwsTgwIntraDomainInspection
	err := ctx.RegisterResource("aviatrix:index/aviatrixAwsTgwIntraDomainInspection:AviatrixAwsTgwIntraDomainInspection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAwsTgwIntraDomainInspection gets an existing AviatrixAwsTgwIntraDomainInspection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAwsTgwIntraDomainInspection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAwsTgwIntraDomainInspectionState, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwIntraDomainInspection, error) {
	var resource AviatrixAwsTgwIntraDomainInspection
	err := ctx.ReadResource("aviatrix:index/aviatrixAwsTgwIntraDomainInspection:AviatrixAwsTgwIntraDomainInspection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAwsTgwIntraDomainInspection resources.
type aviatrixAwsTgwIntraDomainInspectionState struct {
	// Firewall domain name.
	FirewallDomainName *string `pulumi:"firewallDomainName"`
	// Route domain name.
	RouteDomainName *string `pulumi:"routeDomainName"`
	// AWS TGW name.
	TgwName *string `pulumi:"tgwName"`
}

type AviatrixAwsTgwIntraDomainInspectionState struct {
	// Firewall domain name.
	FirewallDomainName pulumi.StringPtrInput
	// Route domain name.
	RouteDomainName pulumi.StringPtrInput
	// AWS TGW name.
	TgwName pulumi.StringPtrInput
}

func (AviatrixAwsTgwIntraDomainInspectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwIntraDomainInspectionState)(nil)).Elem()
}

type aviatrixAwsTgwIntraDomainInspectionArgs struct {
	// Firewall domain name.
	FirewallDomainName string `pulumi:"firewallDomainName"`
	// Route domain name.
	RouteDomainName string `pulumi:"routeDomainName"`
	// AWS TGW name.
	TgwName string `pulumi:"tgwName"`
}

// The set of arguments for constructing a AviatrixAwsTgwIntraDomainInspection resource.
type AviatrixAwsTgwIntraDomainInspectionArgs struct {
	// Firewall domain name.
	FirewallDomainName pulumi.StringInput
	// Route domain name.
	RouteDomainName pulumi.StringInput
	// AWS TGW name.
	TgwName pulumi.StringInput
}

func (AviatrixAwsTgwIntraDomainInspectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwIntraDomainInspectionArgs)(nil)).Elem()
}

type AviatrixAwsTgwIntraDomainInspectionInput interface {
	pulumi.Input

	ToAviatrixAwsTgwIntraDomainInspectionOutput() AviatrixAwsTgwIntraDomainInspectionOutput
	ToAviatrixAwsTgwIntraDomainInspectionOutputWithContext(ctx context.Context) AviatrixAwsTgwIntraDomainInspectionOutput
}

func (*AviatrixAwsTgwIntraDomainInspection) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwIntraDomainInspection)(nil)).Elem()
}

func (i *AviatrixAwsTgwIntraDomainInspection) ToAviatrixAwsTgwIntraDomainInspectionOutput() AviatrixAwsTgwIntraDomainInspectionOutput {
	return i.ToAviatrixAwsTgwIntraDomainInspectionOutputWithContext(context.Background())
}

func (i *AviatrixAwsTgwIntraDomainInspection) ToAviatrixAwsTgwIntraDomainInspectionOutputWithContext(ctx context.Context) AviatrixAwsTgwIntraDomainInspectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwIntraDomainInspectionOutput)
}

// AviatrixAwsTgwIntraDomainInspectionArrayInput is an input type that accepts AviatrixAwsTgwIntraDomainInspectionArray and AviatrixAwsTgwIntraDomainInspectionArrayOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwIntraDomainInspectionArrayInput` via:
//
//	AviatrixAwsTgwIntraDomainInspectionArray{ AviatrixAwsTgwIntraDomainInspectionArgs{...} }
type AviatrixAwsTgwIntraDomainInspectionArrayInput interface {
	pulumi.Input

	ToAviatrixAwsTgwIntraDomainInspectionArrayOutput() AviatrixAwsTgwIntraDomainInspectionArrayOutput
	ToAviatrixAwsTgwIntraDomainInspectionArrayOutputWithContext(context.Context) AviatrixAwsTgwIntraDomainInspectionArrayOutput
}

type AviatrixAwsTgwIntraDomainInspectionArray []AviatrixAwsTgwIntraDomainInspectionInput

func (AviatrixAwsTgwIntraDomainInspectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwIntraDomainInspection)(nil)).Elem()
}

func (i AviatrixAwsTgwIntraDomainInspectionArray) ToAviatrixAwsTgwIntraDomainInspectionArrayOutput() AviatrixAwsTgwIntraDomainInspectionArrayOutput {
	return i.ToAviatrixAwsTgwIntraDomainInspectionArrayOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwIntraDomainInspectionArray) ToAviatrixAwsTgwIntraDomainInspectionArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwIntraDomainInspectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwIntraDomainInspectionArrayOutput)
}

// AviatrixAwsTgwIntraDomainInspectionMapInput is an input type that accepts AviatrixAwsTgwIntraDomainInspectionMap and AviatrixAwsTgwIntraDomainInspectionMapOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwIntraDomainInspectionMapInput` via:
//
//	AviatrixAwsTgwIntraDomainInspectionMap{ "key": AviatrixAwsTgwIntraDomainInspectionArgs{...} }
type AviatrixAwsTgwIntraDomainInspectionMapInput interface {
	pulumi.Input

	ToAviatrixAwsTgwIntraDomainInspectionMapOutput() AviatrixAwsTgwIntraDomainInspectionMapOutput
	ToAviatrixAwsTgwIntraDomainInspectionMapOutputWithContext(context.Context) AviatrixAwsTgwIntraDomainInspectionMapOutput
}

type AviatrixAwsTgwIntraDomainInspectionMap map[string]AviatrixAwsTgwIntraDomainInspectionInput

func (AviatrixAwsTgwIntraDomainInspectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwIntraDomainInspection)(nil)).Elem()
}

func (i AviatrixAwsTgwIntraDomainInspectionMap) ToAviatrixAwsTgwIntraDomainInspectionMapOutput() AviatrixAwsTgwIntraDomainInspectionMapOutput {
	return i.ToAviatrixAwsTgwIntraDomainInspectionMapOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwIntraDomainInspectionMap) ToAviatrixAwsTgwIntraDomainInspectionMapOutputWithContext(ctx context.Context) AviatrixAwsTgwIntraDomainInspectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwIntraDomainInspectionMapOutput)
}

type AviatrixAwsTgwIntraDomainInspectionOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwIntraDomainInspectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwIntraDomainInspection)(nil)).Elem()
}

func (o AviatrixAwsTgwIntraDomainInspectionOutput) ToAviatrixAwsTgwIntraDomainInspectionOutput() AviatrixAwsTgwIntraDomainInspectionOutput {
	return o
}

func (o AviatrixAwsTgwIntraDomainInspectionOutput) ToAviatrixAwsTgwIntraDomainInspectionOutputWithContext(ctx context.Context) AviatrixAwsTgwIntraDomainInspectionOutput {
	return o
}

// Firewall domain name.
func (o AviatrixAwsTgwIntraDomainInspectionOutput) FirewallDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwIntraDomainInspection) pulumi.StringOutput { return v.FirewallDomainName }).(pulumi.StringOutput)
}

// Route domain name.
func (o AviatrixAwsTgwIntraDomainInspectionOutput) RouteDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwIntraDomainInspection) pulumi.StringOutput { return v.RouteDomainName }).(pulumi.StringOutput)
}

// AWS TGW name.
func (o AviatrixAwsTgwIntraDomainInspectionOutput) TgwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwIntraDomainInspection) pulumi.StringOutput { return v.TgwName }).(pulumi.StringOutput)
}

type AviatrixAwsTgwIntraDomainInspectionArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwIntraDomainInspectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwIntraDomainInspection)(nil)).Elem()
}

func (o AviatrixAwsTgwIntraDomainInspectionArrayOutput) ToAviatrixAwsTgwIntraDomainInspectionArrayOutput() AviatrixAwsTgwIntraDomainInspectionArrayOutput {
	return o
}

func (o AviatrixAwsTgwIntraDomainInspectionArrayOutput) ToAviatrixAwsTgwIntraDomainInspectionArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwIntraDomainInspectionArrayOutput {
	return o
}

func (o AviatrixAwsTgwIntraDomainInspectionArrayOutput) Index(i pulumi.IntInput) AviatrixAwsTgwIntraDomainInspectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAwsTgwIntraDomainInspection {
		return vs[0].([]*AviatrixAwsTgwIntraDomainInspection)[vs[1].(int)]
	}).(AviatrixAwsTgwIntraDomainInspectionOutput)
}

type AviatrixAwsTgwIntraDomainInspectionMapOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwIntraDomainInspectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwIntraDomainInspection)(nil)).Elem()
}

func (o AviatrixAwsTgwIntraDomainInspectionMapOutput) ToAviatrixAwsTgwIntraDomainInspectionMapOutput() AviatrixAwsTgwIntraDomainInspectionMapOutput {
	return o
}

func (o AviatrixAwsTgwIntraDomainInspectionMapOutput) ToAviatrixAwsTgwIntraDomainInspectionMapOutputWithContext(ctx context.Context) AviatrixAwsTgwIntraDomainInspectionMapOutput {
	return o
}

func (o AviatrixAwsTgwIntraDomainInspectionMapOutput) MapIndex(k pulumi.StringInput) AviatrixAwsTgwIntraDomainInspectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAwsTgwIntraDomainInspection {
		return vs[0].(map[string]*AviatrixAwsTgwIntraDomainInspection)[vs[1].(string)]
	}).(AviatrixAwsTgwIntraDomainInspectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwIntraDomainInspectionInput)(nil)).Elem(), &AviatrixAwsTgwIntraDomainInspection{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwIntraDomainInspectionArrayInput)(nil)).Elem(), AviatrixAwsTgwIntraDomainInspectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwIntraDomainInspectionMapInput)(nil)).Elem(), AviatrixAwsTgwIntraDomainInspectionMap{})
	pulumi.RegisterOutputType(AviatrixAwsTgwIntraDomainInspectionOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwIntraDomainInspectionArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwIntraDomainInspectionMapOutput{})
}