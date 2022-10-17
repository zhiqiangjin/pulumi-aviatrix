// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixGatewayCertificateConfig struct {
	pulumi.CustomResourceState

	// CA Certificate.
	CaCertificate pulumi.StringOutput `pulumi:"caCertificate"`
	// CA Private Key.
	CaPrivateKey pulumi.StringOutput `pulumi:"caPrivateKey"`
}

// NewAviatrixGatewayCertificateConfig registers a new resource with the given unique name, arguments, and options.
func NewAviatrixGatewayCertificateConfig(ctx *pulumi.Context,
	name string, args *AviatrixGatewayCertificateConfigArgs, opts ...pulumi.ResourceOption) (*AviatrixGatewayCertificateConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaCertificate == nil {
		return nil, errors.New("invalid value for required argument 'CaCertificate'")
	}
	if args.CaPrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'CaPrivateKey'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixGatewayCertificateConfig
	err := ctx.RegisterResource("aviatrix:index/aviatrixGatewayCertificateConfig:AviatrixGatewayCertificateConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixGatewayCertificateConfig gets an existing AviatrixGatewayCertificateConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixGatewayCertificateConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixGatewayCertificateConfigState, opts ...pulumi.ResourceOption) (*AviatrixGatewayCertificateConfig, error) {
	var resource AviatrixGatewayCertificateConfig
	err := ctx.ReadResource("aviatrix:index/aviatrixGatewayCertificateConfig:AviatrixGatewayCertificateConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixGatewayCertificateConfig resources.
type aviatrixGatewayCertificateConfigState struct {
	// CA Certificate.
	CaCertificate *string `pulumi:"caCertificate"`
	// CA Private Key.
	CaPrivateKey *string `pulumi:"caPrivateKey"`
}

type AviatrixGatewayCertificateConfigState struct {
	// CA Certificate.
	CaCertificate pulumi.StringPtrInput
	// CA Private Key.
	CaPrivateKey pulumi.StringPtrInput
}

func (AviatrixGatewayCertificateConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixGatewayCertificateConfigState)(nil)).Elem()
}

type aviatrixGatewayCertificateConfigArgs struct {
	// CA Certificate.
	CaCertificate string `pulumi:"caCertificate"`
	// CA Private Key.
	CaPrivateKey string `pulumi:"caPrivateKey"`
}

// The set of arguments for constructing a AviatrixGatewayCertificateConfig resource.
type AviatrixGatewayCertificateConfigArgs struct {
	// CA Certificate.
	CaCertificate pulumi.StringInput
	// CA Private Key.
	CaPrivateKey pulumi.StringInput
}

func (AviatrixGatewayCertificateConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixGatewayCertificateConfigArgs)(nil)).Elem()
}

type AviatrixGatewayCertificateConfigInput interface {
	pulumi.Input

	ToAviatrixGatewayCertificateConfigOutput() AviatrixGatewayCertificateConfigOutput
	ToAviatrixGatewayCertificateConfigOutputWithContext(ctx context.Context) AviatrixGatewayCertificateConfigOutput
}

func (*AviatrixGatewayCertificateConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixGatewayCertificateConfig)(nil)).Elem()
}

func (i *AviatrixGatewayCertificateConfig) ToAviatrixGatewayCertificateConfigOutput() AviatrixGatewayCertificateConfigOutput {
	return i.ToAviatrixGatewayCertificateConfigOutputWithContext(context.Background())
}

func (i *AviatrixGatewayCertificateConfig) ToAviatrixGatewayCertificateConfigOutputWithContext(ctx context.Context) AviatrixGatewayCertificateConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixGatewayCertificateConfigOutput)
}

// AviatrixGatewayCertificateConfigArrayInput is an input type that accepts AviatrixGatewayCertificateConfigArray and AviatrixGatewayCertificateConfigArrayOutput values.
// You can construct a concrete instance of `AviatrixGatewayCertificateConfigArrayInput` via:
//
//	AviatrixGatewayCertificateConfigArray{ AviatrixGatewayCertificateConfigArgs{...} }
type AviatrixGatewayCertificateConfigArrayInput interface {
	pulumi.Input

	ToAviatrixGatewayCertificateConfigArrayOutput() AviatrixGatewayCertificateConfigArrayOutput
	ToAviatrixGatewayCertificateConfigArrayOutputWithContext(context.Context) AviatrixGatewayCertificateConfigArrayOutput
}

type AviatrixGatewayCertificateConfigArray []AviatrixGatewayCertificateConfigInput

func (AviatrixGatewayCertificateConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixGatewayCertificateConfig)(nil)).Elem()
}

func (i AviatrixGatewayCertificateConfigArray) ToAviatrixGatewayCertificateConfigArrayOutput() AviatrixGatewayCertificateConfigArrayOutput {
	return i.ToAviatrixGatewayCertificateConfigArrayOutputWithContext(context.Background())
}

func (i AviatrixGatewayCertificateConfigArray) ToAviatrixGatewayCertificateConfigArrayOutputWithContext(ctx context.Context) AviatrixGatewayCertificateConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixGatewayCertificateConfigArrayOutput)
}

// AviatrixGatewayCertificateConfigMapInput is an input type that accepts AviatrixGatewayCertificateConfigMap and AviatrixGatewayCertificateConfigMapOutput values.
// You can construct a concrete instance of `AviatrixGatewayCertificateConfigMapInput` via:
//
//	AviatrixGatewayCertificateConfigMap{ "key": AviatrixGatewayCertificateConfigArgs{...} }
type AviatrixGatewayCertificateConfigMapInput interface {
	pulumi.Input

	ToAviatrixGatewayCertificateConfigMapOutput() AviatrixGatewayCertificateConfigMapOutput
	ToAviatrixGatewayCertificateConfigMapOutputWithContext(context.Context) AviatrixGatewayCertificateConfigMapOutput
}

type AviatrixGatewayCertificateConfigMap map[string]AviatrixGatewayCertificateConfigInput

func (AviatrixGatewayCertificateConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixGatewayCertificateConfig)(nil)).Elem()
}

func (i AviatrixGatewayCertificateConfigMap) ToAviatrixGatewayCertificateConfigMapOutput() AviatrixGatewayCertificateConfigMapOutput {
	return i.ToAviatrixGatewayCertificateConfigMapOutputWithContext(context.Background())
}

func (i AviatrixGatewayCertificateConfigMap) ToAviatrixGatewayCertificateConfigMapOutputWithContext(ctx context.Context) AviatrixGatewayCertificateConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixGatewayCertificateConfigMapOutput)
}

type AviatrixGatewayCertificateConfigOutput struct{ *pulumi.OutputState }

func (AviatrixGatewayCertificateConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixGatewayCertificateConfig)(nil)).Elem()
}

func (o AviatrixGatewayCertificateConfigOutput) ToAviatrixGatewayCertificateConfigOutput() AviatrixGatewayCertificateConfigOutput {
	return o
}

func (o AviatrixGatewayCertificateConfigOutput) ToAviatrixGatewayCertificateConfigOutputWithContext(ctx context.Context) AviatrixGatewayCertificateConfigOutput {
	return o
}

// CA Certificate.
func (o AviatrixGatewayCertificateConfigOutput) CaCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixGatewayCertificateConfig) pulumi.StringOutput { return v.CaCertificate }).(pulumi.StringOutput)
}

// CA Private Key.
func (o AviatrixGatewayCertificateConfigOutput) CaPrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixGatewayCertificateConfig) pulumi.StringOutput { return v.CaPrivateKey }).(pulumi.StringOutput)
}

type AviatrixGatewayCertificateConfigArrayOutput struct{ *pulumi.OutputState }

func (AviatrixGatewayCertificateConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixGatewayCertificateConfig)(nil)).Elem()
}

func (o AviatrixGatewayCertificateConfigArrayOutput) ToAviatrixGatewayCertificateConfigArrayOutput() AviatrixGatewayCertificateConfigArrayOutput {
	return o
}

func (o AviatrixGatewayCertificateConfigArrayOutput) ToAviatrixGatewayCertificateConfigArrayOutputWithContext(ctx context.Context) AviatrixGatewayCertificateConfigArrayOutput {
	return o
}

func (o AviatrixGatewayCertificateConfigArrayOutput) Index(i pulumi.IntInput) AviatrixGatewayCertificateConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixGatewayCertificateConfig {
		return vs[0].([]*AviatrixGatewayCertificateConfig)[vs[1].(int)]
	}).(AviatrixGatewayCertificateConfigOutput)
}

type AviatrixGatewayCertificateConfigMapOutput struct{ *pulumi.OutputState }

func (AviatrixGatewayCertificateConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixGatewayCertificateConfig)(nil)).Elem()
}

func (o AviatrixGatewayCertificateConfigMapOutput) ToAviatrixGatewayCertificateConfigMapOutput() AviatrixGatewayCertificateConfigMapOutput {
	return o
}

func (o AviatrixGatewayCertificateConfigMapOutput) ToAviatrixGatewayCertificateConfigMapOutputWithContext(ctx context.Context) AviatrixGatewayCertificateConfigMapOutput {
	return o
}

func (o AviatrixGatewayCertificateConfigMapOutput) MapIndex(k pulumi.StringInput) AviatrixGatewayCertificateConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixGatewayCertificateConfig {
		return vs[0].(map[string]*AviatrixGatewayCertificateConfig)[vs[1].(string)]
	}).(AviatrixGatewayCertificateConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixGatewayCertificateConfigInput)(nil)).Elem(), &AviatrixGatewayCertificateConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixGatewayCertificateConfigArrayInput)(nil)).Elem(), AviatrixGatewayCertificateConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixGatewayCertificateConfigMapInput)(nil)).Elem(), AviatrixGatewayCertificateConfigMap{})
	pulumi.RegisterOutputType(AviatrixGatewayCertificateConfigOutput{})
	pulumi.RegisterOutputType(AviatrixGatewayCertificateConfigArrayOutput{})
	pulumi.RegisterOutputType(AviatrixGatewayCertificateConfigMapOutput{})
}
