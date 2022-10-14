// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixControllerPrivateModeConfig struct {
	pulumi.CustomResourceState

	// Copilot instance ID to associate with the Controller for Private Mode.
	CopilotInstanceId pulumi.StringPtrOutput `pulumi:"copilotInstanceId"`
	// Whether to enable Private Mode on the Controller.
	EnablePrivateMode pulumi.BoolOutput `pulumi:"enablePrivateMode"`
	// Set of proxies.
	Proxies pulumi.StringArrayOutput `pulumi:"proxies"`
}

// NewAviatrixControllerPrivateModeConfig registers a new resource with the given unique name, arguments, and options.
func NewAviatrixControllerPrivateModeConfig(ctx *pulumi.Context,
	name string, args *AviatrixControllerPrivateModeConfigArgs, opts ...pulumi.ResourceOption) (*AviatrixControllerPrivateModeConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnablePrivateMode == nil {
		return nil, errors.New("invalid value for required argument 'EnablePrivateMode'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixControllerPrivateModeConfig
	err := ctx.RegisterResource("aviatrix:index/aviatrixControllerPrivateModeConfig:AviatrixControllerPrivateModeConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixControllerPrivateModeConfig gets an existing AviatrixControllerPrivateModeConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixControllerPrivateModeConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixControllerPrivateModeConfigState, opts ...pulumi.ResourceOption) (*AviatrixControllerPrivateModeConfig, error) {
	var resource AviatrixControllerPrivateModeConfig
	err := ctx.ReadResource("aviatrix:index/aviatrixControllerPrivateModeConfig:AviatrixControllerPrivateModeConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixControllerPrivateModeConfig resources.
type aviatrixControllerPrivateModeConfigState struct {
	// Copilot instance ID to associate with the Controller for Private Mode.
	CopilotInstanceId *string `pulumi:"copilotInstanceId"`
	// Whether to enable Private Mode on the Controller.
	EnablePrivateMode *bool `pulumi:"enablePrivateMode"`
	// Set of proxies.
	Proxies []string `pulumi:"proxies"`
}

type AviatrixControllerPrivateModeConfigState struct {
	// Copilot instance ID to associate with the Controller for Private Mode.
	CopilotInstanceId pulumi.StringPtrInput
	// Whether to enable Private Mode on the Controller.
	EnablePrivateMode pulumi.BoolPtrInput
	// Set of proxies.
	Proxies pulumi.StringArrayInput
}

func (AviatrixControllerPrivateModeConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixControllerPrivateModeConfigState)(nil)).Elem()
}

type aviatrixControllerPrivateModeConfigArgs struct {
	// Copilot instance ID to associate with the Controller for Private Mode.
	CopilotInstanceId *string `pulumi:"copilotInstanceId"`
	// Whether to enable Private Mode on the Controller.
	EnablePrivateMode bool `pulumi:"enablePrivateMode"`
	// Set of proxies.
	Proxies []string `pulumi:"proxies"`
}

// The set of arguments for constructing a AviatrixControllerPrivateModeConfig resource.
type AviatrixControllerPrivateModeConfigArgs struct {
	// Copilot instance ID to associate with the Controller for Private Mode.
	CopilotInstanceId pulumi.StringPtrInput
	// Whether to enable Private Mode on the Controller.
	EnablePrivateMode pulumi.BoolInput
	// Set of proxies.
	Proxies pulumi.StringArrayInput
}

func (AviatrixControllerPrivateModeConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixControllerPrivateModeConfigArgs)(nil)).Elem()
}

type AviatrixControllerPrivateModeConfigInput interface {
	pulumi.Input

	ToAviatrixControllerPrivateModeConfigOutput() AviatrixControllerPrivateModeConfigOutput
	ToAviatrixControllerPrivateModeConfigOutputWithContext(ctx context.Context) AviatrixControllerPrivateModeConfigOutput
}

func (*AviatrixControllerPrivateModeConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixControllerPrivateModeConfig)(nil)).Elem()
}

func (i *AviatrixControllerPrivateModeConfig) ToAviatrixControllerPrivateModeConfigOutput() AviatrixControllerPrivateModeConfigOutput {
	return i.ToAviatrixControllerPrivateModeConfigOutputWithContext(context.Background())
}

func (i *AviatrixControllerPrivateModeConfig) ToAviatrixControllerPrivateModeConfigOutputWithContext(ctx context.Context) AviatrixControllerPrivateModeConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerPrivateModeConfigOutput)
}

// AviatrixControllerPrivateModeConfigArrayInput is an input type that accepts AviatrixControllerPrivateModeConfigArray and AviatrixControllerPrivateModeConfigArrayOutput values.
// You can construct a concrete instance of `AviatrixControllerPrivateModeConfigArrayInput` via:
//
//	AviatrixControllerPrivateModeConfigArray{ AviatrixControllerPrivateModeConfigArgs{...} }
type AviatrixControllerPrivateModeConfigArrayInput interface {
	pulumi.Input

	ToAviatrixControllerPrivateModeConfigArrayOutput() AviatrixControllerPrivateModeConfigArrayOutput
	ToAviatrixControllerPrivateModeConfigArrayOutputWithContext(context.Context) AviatrixControllerPrivateModeConfigArrayOutput
}

type AviatrixControllerPrivateModeConfigArray []AviatrixControllerPrivateModeConfigInput

func (AviatrixControllerPrivateModeConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixControllerPrivateModeConfig)(nil)).Elem()
}

func (i AviatrixControllerPrivateModeConfigArray) ToAviatrixControllerPrivateModeConfigArrayOutput() AviatrixControllerPrivateModeConfigArrayOutput {
	return i.ToAviatrixControllerPrivateModeConfigArrayOutputWithContext(context.Background())
}

func (i AviatrixControllerPrivateModeConfigArray) ToAviatrixControllerPrivateModeConfigArrayOutputWithContext(ctx context.Context) AviatrixControllerPrivateModeConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerPrivateModeConfigArrayOutput)
}

// AviatrixControllerPrivateModeConfigMapInput is an input type that accepts AviatrixControllerPrivateModeConfigMap and AviatrixControllerPrivateModeConfigMapOutput values.
// You can construct a concrete instance of `AviatrixControllerPrivateModeConfigMapInput` via:
//
//	AviatrixControllerPrivateModeConfigMap{ "key": AviatrixControllerPrivateModeConfigArgs{...} }
type AviatrixControllerPrivateModeConfigMapInput interface {
	pulumi.Input

	ToAviatrixControllerPrivateModeConfigMapOutput() AviatrixControllerPrivateModeConfigMapOutput
	ToAviatrixControllerPrivateModeConfigMapOutputWithContext(context.Context) AviatrixControllerPrivateModeConfigMapOutput
}

type AviatrixControllerPrivateModeConfigMap map[string]AviatrixControllerPrivateModeConfigInput

func (AviatrixControllerPrivateModeConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixControllerPrivateModeConfig)(nil)).Elem()
}

func (i AviatrixControllerPrivateModeConfigMap) ToAviatrixControllerPrivateModeConfigMapOutput() AviatrixControllerPrivateModeConfigMapOutput {
	return i.ToAviatrixControllerPrivateModeConfigMapOutputWithContext(context.Background())
}

func (i AviatrixControllerPrivateModeConfigMap) ToAviatrixControllerPrivateModeConfigMapOutputWithContext(ctx context.Context) AviatrixControllerPrivateModeConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixControllerPrivateModeConfigMapOutput)
}

type AviatrixControllerPrivateModeConfigOutput struct{ *pulumi.OutputState }

func (AviatrixControllerPrivateModeConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixControllerPrivateModeConfig)(nil)).Elem()
}

func (o AviatrixControllerPrivateModeConfigOutput) ToAviatrixControllerPrivateModeConfigOutput() AviatrixControllerPrivateModeConfigOutput {
	return o
}

func (o AviatrixControllerPrivateModeConfigOutput) ToAviatrixControllerPrivateModeConfigOutputWithContext(ctx context.Context) AviatrixControllerPrivateModeConfigOutput {
	return o
}

// Copilot instance ID to associate with the Controller for Private Mode.
func (o AviatrixControllerPrivateModeConfigOutput) CopilotInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixControllerPrivateModeConfig) pulumi.StringPtrOutput { return v.CopilotInstanceId }).(pulumi.StringPtrOutput)
}

// Whether to enable Private Mode on the Controller.
func (o AviatrixControllerPrivateModeConfigOutput) EnablePrivateMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *AviatrixControllerPrivateModeConfig) pulumi.BoolOutput { return v.EnablePrivateMode }).(pulumi.BoolOutput)
}

// Set of proxies.
func (o AviatrixControllerPrivateModeConfigOutput) Proxies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AviatrixControllerPrivateModeConfig) pulumi.StringArrayOutput { return v.Proxies }).(pulumi.StringArrayOutput)
}

type AviatrixControllerPrivateModeConfigArrayOutput struct{ *pulumi.OutputState }

func (AviatrixControllerPrivateModeConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixControllerPrivateModeConfig)(nil)).Elem()
}

func (o AviatrixControllerPrivateModeConfigArrayOutput) ToAviatrixControllerPrivateModeConfigArrayOutput() AviatrixControllerPrivateModeConfigArrayOutput {
	return o
}

func (o AviatrixControllerPrivateModeConfigArrayOutput) ToAviatrixControllerPrivateModeConfigArrayOutputWithContext(ctx context.Context) AviatrixControllerPrivateModeConfigArrayOutput {
	return o
}

func (o AviatrixControllerPrivateModeConfigArrayOutput) Index(i pulumi.IntInput) AviatrixControllerPrivateModeConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixControllerPrivateModeConfig {
		return vs[0].([]*AviatrixControllerPrivateModeConfig)[vs[1].(int)]
	}).(AviatrixControllerPrivateModeConfigOutput)
}

type AviatrixControllerPrivateModeConfigMapOutput struct{ *pulumi.OutputState }

func (AviatrixControllerPrivateModeConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixControllerPrivateModeConfig)(nil)).Elem()
}

func (o AviatrixControllerPrivateModeConfigMapOutput) ToAviatrixControllerPrivateModeConfigMapOutput() AviatrixControllerPrivateModeConfigMapOutput {
	return o
}

func (o AviatrixControllerPrivateModeConfigMapOutput) ToAviatrixControllerPrivateModeConfigMapOutputWithContext(ctx context.Context) AviatrixControllerPrivateModeConfigMapOutput {
	return o
}

func (o AviatrixControllerPrivateModeConfigMapOutput) MapIndex(k pulumi.StringInput) AviatrixControllerPrivateModeConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixControllerPrivateModeConfig {
		return vs[0].(map[string]*AviatrixControllerPrivateModeConfig)[vs[1].(string)]
	}).(AviatrixControllerPrivateModeConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerPrivateModeConfigInput)(nil)).Elem(), &AviatrixControllerPrivateModeConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerPrivateModeConfigArrayInput)(nil)).Elem(), AviatrixControllerPrivateModeConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixControllerPrivateModeConfigMapInput)(nil)).Elem(), AviatrixControllerPrivateModeConfigMap{})
	pulumi.RegisterOutputType(AviatrixControllerPrivateModeConfigOutput{})
	pulumi.RegisterOutputType(AviatrixControllerPrivateModeConfigArrayOutput{})
	pulumi.RegisterOutputType(AviatrixControllerPrivateModeConfigMapOutput{})
}