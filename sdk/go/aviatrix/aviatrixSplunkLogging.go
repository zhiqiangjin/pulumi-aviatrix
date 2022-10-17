// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixSplunkLogging struct {
	pulumi.CustomResourceState

	// Custom configuration.
	CustomInputConfig pulumi.StringPtrOutput `pulumi:"customInputConfig"`
	// Configuration file. Use the filebase64 function to read from a file.
	CustomOutputConfigFile pulumi.StringPtrOutput `pulumi:"customOutputConfigFile"`
	// List of excluded gateways.
	ExcludedGateways pulumi.StringArrayOutput `pulumi:"excludedGateways"`
	// Port number.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Server IP.
	Server pulumi.StringPtrOutput `pulumi:"server"`
	// Enabled or not.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAviatrixSplunkLogging registers a new resource with the given unique name, arguments, and options.
func NewAviatrixSplunkLogging(ctx *pulumi.Context,
	name string, args *AviatrixSplunkLoggingArgs, opts ...pulumi.ResourceOption) (*AviatrixSplunkLogging, error) {
	if args == nil {
		args = &AviatrixSplunkLoggingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixSplunkLogging
	err := ctx.RegisterResource("aviatrix:index/aviatrixSplunkLogging:AviatrixSplunkLogging", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixSplunkLogging gets an existing AviatrixSplunkLogging resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixSplunkLogging(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixSplunkLoggingState, opts ...pulumi.ResourceOption) (*AviatrixSplunkLogging, error) {
	var resource AviatrixSplunkLogging
	err := ctx.ReadResource("aviatrix:index/aviatrixSplunkLogging:AviatrixSplunkLogging", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixSplunkLogging resources.
type aviatrixSplunkLoggingState struct {
	// Custom configuration.
	CustomInputConfig *string `pulumi:"customInputConfig"`
	// Configuration file. Use the filebase64 function to read from a file.
	CustomOutputConfigFile *string `pulumi:"customOutputConfigFile"`
	// List of excluded gateways.
	ExcludedGateways []string `pulumi:"excludedGateways"`
	// Port number.
	Port *int `pulumi:"port"`
	// Server IP.
	Server *string `pulumi:"server"`
	// Enabled or not.
	Status *string `pulumi:"status"`
}

type AviatrixSplunkLoggingState struct {
	// Custom configuration.
	CustomInputConfig pulumi.StringPtrInput
	// Configuration file. Use the filebase64 function to read from a file.
	CustomOutputConfigFile pulumi.StringPtrInput
	// List of excluded gateways.
	ExcludedGateways pulumi.StringArrayInput
	// Port number.
	Port pulumi.IntPtrInput
	// Server IP.
	Server pulumi.StringPtrInput
	// Enabled or not.
	Status pulumi.StringPtrInput
}

func (AviatrixSplunkLoggingState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixSplunkLoggingState)(nil)).Elem()
}

type aviatrixSplunkLoggingArgs struct {
	// Custom configuration.
	CustomInputConfig *string `pulumi:"customInputConfig"`
	// Configuration file. Use the filebase64 function to read from a file.
	CustomOutputConfigFile *string `pulumi:"customOutputConfigFile"`
	// List of excluded gateways.
	ExcludedGateways []string `pulumi:"excludedGateways"`
	// Port number.
	Port *int `pulumi:"port"`
	// Server IP.
	Server *string `pulumi:"server"`
}

// The set of arguments for constructing a AviatrixSplunkLogging resource.
type AviatrixSplunkLoggingArgs struct {
	// Custom configuration.
	CustomInputConfig pulumi.StringPtrInput
	// Configuration file. Use the filebase64 function to read from a file.
	CustomOutputConfigFile pulumi.StringPtrInput
	// List of excluded gateways.
	ExcludedGateways pulumi.StringArrayInput
	// Port number.
	Port pulumi.IntPtrInput
	// Server IP.
	Server pulumi.StringPtrInput
}

func (AviatrixSplunkLoggingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixSplunkLoggingArgs)(nil)).Elem()
}

type AviatrixSplunkLoggingInput interface {
	pulumi.Input

	ToAviatrixSplunkLoggingOutput() AviatrixSplunkLoggingOutput
	ToAviatrixSplunkLoggingOutputWithContext(ctx context.Context) AviatrixSplunkLoggingOutput
}

func (*AviatrixSplunkLogging) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixSplunkLogging)(nil)).Elem()
}

func (i *AviatrixSplunkLogging) ToAviatrixSplunkLoggingOutput() AviatrixSplunkLoggingOutput {
	return i.ToAviatrixSplunkLoggingOutputWithContext(context.Background())
}

func (i *AviatrixSplunkLogging) ToAviatrixSplunkLoggingOutputWithContext(ctx context.Context) AviatrixSplunkLoggingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSplunkLoggingOutput)
}

// AviatrixSplunkLoggingArrayInput is an input type that accepts AviatrixSplunkLoggingArray and AviatrixSplunkLoggingArrayOutput values.
// You can construct a concrete instance of `AviatrixSplunkLoggingArrayInput` via:
//
//	AviatrixSplunkLoggingArray{ AviatrixSplunkLoggingArgs{...} }
type AviatrixSplunkLoggingArrayInput interface {
	pulumi.Input

	ToAviatrixSplunkLoggingArrayOutput() AviatrixSplunkLoggingArrayOutput
	ToAviatrixSplunkLoggingArrayOutputWithContext(context.Context) AviatrixSplunkLoggingArrayOutput
}

type AviatrixSplunkLoggingArray []AviatrixSplunkLoggingInput

func (AviatrixSplunkLoggingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixSplunkLogging)(nil)).Elem()
}

func (i AviatrixSplunkLoggingArray) ToAviatrixSplunkLoggingArrayOutput() AviatrixSplunkLoggingArrayOutput {
	return i.ToAviatrixSplunkLoggingArrayOutputWithContext(context.Background())
}

func (i AviatrixSplunkLoggingArray) ToAviatrixSplunkLoggingArrayOutputWithContext(ctx context.Context) AviatrixSplunkLoggingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSplunkLoggingArrayOutput)
}

// AviatrixSplunkLoggingMapInput is an input type that accepts AviatrixSplunkLoggingMap and AviatrixSplunkLoggingMapOutput values.
// You can construct a concrete instance of `AviatrixSplunkLoggingMapInput` via:
//
//	AviatrixSplunkLoggingMap{ "key": AviatrixSplunkLoggingArgs{...} }
type AviatrixSplunkLoggingMapInput interface {
	pulumi.Input

	ToAviatrixSplunkLoggingMapOutput() AviatrixSplunkLoggingMapOutput
	ToAviatrixSplunkLoggingMapOutputWithContext(context.Context) AviatrixSplunkLoggingMapOutput
}

type AviatrixSplunkLoggingMap map[string]AviatrixSplunkLoggingInput

func (AviatrixSplunkLoggingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixSplunkLogging)(nil)).Elem()
}

func (i AviatrixSplunkLoggingMap) ToAviatrixSplunkLoggingMapOutput() AviatrixSplunkLoggingMapOutput {
	return i.ToAviatrixSplunkLoggingMapOutputWithContext(context.Background())
}

func (i AviatrixSplunkLoggingMap) ToAviatrixSplunkLoggingMapOutputWithContext(ctx context.Context) AviatrixSplunkLoggingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSplunkLoggingMapOutput)
}

type AviatrixSplunkLoggingOutput struct{ *pulumi.OutputState }

func (AviatrixSplunkLoggingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixSplunkLogging)(nil)).Elem()
}

func (o AviatrixSplunkLoggingOutput) ToAviatrixSplunkLoggingOutput() AviatrixSplunkLoggingOutput {
	return o
}

func (o AviatrixSplunkLoggingOutput) ToAviatrixSplunkLoggingOutputWithContext(ctx context.Context) AviatrixSplunkLoggingOutput {
	return o
}

// Custom configuration.
func (o AviatrixSplunkLoggingOutput) CustomInputConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSplunkLogging) pulumi.StringPtrOutput { return v.CustomInputConfig }).(pulumi.StringPtrOutput)
}

// Configuration file. Use the filebase64 function to read from a file.
func (o AviatrixSplunkLoggingOutput) CustomOutputConfigFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSplunkLogging) pulumi.StringPtrOutput { return v.CustomOutputConfigFile }).(pulumi.StringPtrOutput)
}

// List of excluded gateways.
func (o AviatrixSplunkLoggingOutput) ExcludedGateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixSplunkLogging) pulumi.StringArrayOutput { return v.ExcludedGateways }).(pulumi.StringArrayOutput)
}

// Port number.
func (o AviatrixSplunkLoggingOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AviatrixSplunkLogging) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// Server IP.
func (o AviatrixSplunkLoggingOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixSplunkLogging) pulumi.StringPtrOutput { return v.Server }).(pulumi.StringPtrOutput)
}

// Enabled or not.
func (o AviatrixSplunkLoggingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSplunkLogging) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AviatrixSplunkLoggingArrayOutput struct{ *pulumi.OutputState }

func (AviatrixSplunkLoggingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixSplunkLogging)(nil)).Elem()
}

func (o AviatrixSplunkLoggingArrayOutput) ToAviatrixSplunkLoggingArrayOutput() AviatrixSplunkLoggingArrayOutput {
	return o
}

func (o AviatrixSplunkLoggingArrayOutput) ToAviatrixSplunkLoggingArrayOutputWithContext(ctx context.Context) AviatrixSplunkLoggingArrayOutput {
	return o
}

func (o AviatrixSplunkLoggingArrayOutput) Index(i pulumi.IntInput) AviatrixSplunkLoggingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixSplunkLogging {
		return vs[0].([]*AviatrixSplunkLogging)[vs[1].(int)]
	}).(AviatrixSplunkLoggingOutput)
}

type AviatrixSplunkLoggingMapOutput struct{ *pulumi.OutputState }

func (AviatrixSplunkLoggingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixSplunkLogging)(nil)).Elem()
}

func (o AviatrixSplunkLoggingMapOutput) ToAviatrixSplunkLoggingMapOutput() AviatrixSplunkLoggingMapOutput {
	return o
}

func (o AviatrixSplunkLoggingMapOutput) ToAviatrixSplunkLoggingMapOutputWithContext(ctx context.Context) AviatrixSplunkLoggingMapOutput {
	return o
}

func (o AviatrixSplunkLoggingMapOutput) MapIndex(k pulumi.StringInput) AviatrixSplunkLoggingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixSplunkLogging {
		return vs[0].(map[string]*AviatrixSplunkLogging)[vs[1].(string)]
	}).(AviatrixSplunkLoggingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSplunkLoggingInput)(nil)).Elem(), &AviatrixSplunkLogging{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSplunkLoggingArrayInput)(nil)).Elem(), AviatrixSplunkLoggingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSplunkLoggingMapInput)(nil)).Elem(), AviatrixSplunkLoggingMap{})
	pulumi.RegisterOutputType(AviatrixSplunkLoggingOutput{})
	pulumi.RegisterOutputType(AviatrixSplunkLoggingArrayOutput{})
	pulumi.RegisterOutputType(AviatrixSplunkLoggingMapOutput{})
}
