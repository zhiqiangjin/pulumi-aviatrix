// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixSegmentationSecurityDomain struct {
	pulumi.CustomResourceState

	// Security domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
}

// NewAviatrixSegmentationSecurityDomain registers a new resource with the given unique name, arguments, and options.
func NewAviatrixSegmentationSecurityDomain(ctx *pulumi.Context,
	name string, args *AviatrixSegmentationSecurityDomainArgs, opts ...pulumi.ResourceOption) (*AviatrixSegmentationSecurityDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixSegmentationSecurityDomain
	err := ctx.RegisterResource("aviatrix:index/aviatrixSegmentationSecurityDomain:AviatrixSegmentationSecurityDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixSegmentationSecurityDomain gets an existing AviatrixSegmentationSecurityDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixSegmentationSecurityDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixSegmentationSecurityDomainState, opts ...pulumi.ResourceOption) (*AviatrixSegmentationSecurityDomain, error) {
	var resource AviatrixSegmentationSecurityDomain
	err := ctx.ReadResource("aviatrix:index/aviatrixSegmentationSecurityDomain:AviatrixSegmentationSecurityDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixSegmentationSecurityDomain resources.
type aviatrixSegmentationSecurityDomainState struct {
	// Security domain name.
	DomainName *string `pulumi:"domainName"`
}

type AviatrixSegmentationSecurityDomainState struct {
	// Security domain name.
	DomainName pulumi.StringPtrInput
}

func (AviatrixSegmentationSecurityDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixSegmentationSecurityDomainState)(nil)).Elem()
}

type aviatrixSegmentationSecurityDomainArgs struct {
	// Security domain name.
	DomainName string `pulumi:"domainName"`
}

// The set of arguments for constructing a AviatrixSegmentationSecurityDomain resource.
type AviatrixSegmentationSecurityDomainArgs struct {
	// Security domain name.
	DomainName pulumi.StringInput
}

func (AviatrixSegmentationSecurityDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixSegmentationSecurityDomainArgs)(nil)).Elem()
}

type AviatrixSegmentationSecurityDomainInput interface {
	pulumi.Input

	ToAviatrixSegmentationSecurityDomainOutput() AviatrixSegmentationSecurityDomainOutput
	ToAviatrixSegmentationSecurityDomainOutputWithContext(ctx context.Context) AviatrixSegmentationSecurityDomainOutput
}

func (*AviatrixSegmentationSecurityDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixSegmentationSecurityDomain)(nil)).Elem()
}

func (i *AviatrixSegmentationSecurityDomain) ToAviatrixSegmentationSecurityDomainOutput() AviatrixSegmentationSecurityDomainOutput {
	return i.ToAviatrixSegmentationSecurityDomainOutputWithContext(context.Background())
}

func (i *AviatrixSegmentationSecurityDomain) ToAviatrixSegmentationSecurityDomainOutputWithContext(ctx context.Context) AviatrixSegmentationSecurityDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSegmentationSecurityDomainOutput)
}

// AviatrixSegmentationSecurityDomainArrayInput is an input type that accepts AviatrixSegmentationSecurityDomainArray and AviatrixSegmentationSecurityDomainArrayOutput values.
// You can construct a concrete instance of `AviatrixSegmentationSecurityDomainArrayInput` via:
//
//	AviatrixSegmentationSecurityDomainArray{ AviatrixSegmentationSecurityDomainArgs{...} }
type AviatrixSegmentationSecurityDomainArrayInput interface {
	pulumi.Input

	ToAviatrixSegmentationSecurityDomainArrayOutput() AviatrixSegmentationSecurityDomainArrayOutput
	ToAviatrixSegmentationSecurityDomainArrayOutputWithContext(context.Context) AviatrixSegmentationSecurityDomainArrayOutput
}

type AviatrixSegmentationSecurityDomainArray []AviatrixSegmentationSecurityDomainInput

func (AviatrixSegmentationSecurityDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixSegmentationSecurityDomain)(nil)).Elem()
}

func (i AviatrixSegmentationSecurityDomainArray) ToAviatrixSegmentationSecurityDomainArrayOutput() AviatrixSegmentationSecurityDomainArrayOutput {
	return i.ToAviatrixSegmentationSecurityDomainArrayOutputWithContext(context.Background())
}

func (i AviatrixSegmentationSecurityDomainArray) ToAviatrixSegmentationSecurityDomainArrayOutputWithContext(ctx context.Context) AviatrixSegmentationSecurityDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSegmentationSecurityDomainArrayOutput)
}

// AviatrixSegmentationSecurityDomainMapInput is an input type that accepts AviatrixSegmentationSecurityDomainMap and AviatrixSegmentationSecurityDomainMapOutput values.
// You can construct a concrete instance of `AviatrixSegmentationSecurityDomainMapInput` via:
//
//	AviatrixSegmentationSecurityDomainMap{ "key": AviatrixSegmentationSecurityDomainArgs{...} }
type AviatrixSegmentationSecurityDomainMapInput interface {
	pulumi.Input

	ToAviatrixSegmentationSecurityDomainMapOutput() AviatrixSegmentationSecurityDomainMapOutput
	ToAviatrixSegmentationSecurityDomainMapOutputWithContext(context.Context) AviatrixSegmentationSecurityDomainMapOutput
}

type AviatrixSegmentationSecurityDomainMap map[string]AviatrixSegmentationSecurityDomainInput

func (AviatrixSegmentationSecurityDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixSegmentationSecurityDomain)(nil)).Elem()
}

func (i AviatrixSegmentationSecurityDomainMap) ToAviatrixSegmentationSecurityDomainMapOutput() AviatrixSegmentationSecurityDomainMapOutput {
	return i.ToAviatrixSegmentationSecurityDomainMapOutputWithContext(context.Background())
}

func (i AviatrixSegmentationSecurityDomainMap) ToAviatrixSegmentationSecurityDomainMapOutputWithContext(ctx context.Context) AviatrixSegmentationSecurityDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixSegmentationSecurityDomainMapOutput)
}

type AviatrixSegmentationSecurityDomainOutput struct{ *pulumi.OutputState }

func (AviatrixSegmentationSecurityDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixSegmentationSecurityDomain)(nil)).Elem()
}

func (o AviatrixSegmentationSecurityDomainOutput) ToAviatrixSegmentationSecurityDomainOutput() AviatrixSegmentationSecurityDomainOutput {
	return o
}

func (o AviatrixSegmentationSecurityDomainOutput) ToAviatrixSegmentationSecurityDomainOutputWithContext(ctx context.Context) AviatrixSegmentationSecurityDomainOutput {
	return o
}

// Security domain name.
func (o AviatrixSegmentationSecurityDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixSegmentationSecurityDomain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

type AviatrixSegmentationSecurityDomainArrayOutput struct{ *pulumi.OutputState }

func (AviatrixSegmentationSecurityDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixSegmentationSecurityDomain)(nil)).Elem()
}

func (o AviatrixSegmentationSecurityDomainArrayOutput) ToAviatrixSegmentationSecurityDomainArrayOutput() AviatrixSegmentationSecurityDomainArrayOutput {
	return o
}

func (o AviatrixSegmentationSecurityDomainArrayOutput) ToAviatrixSegmentationSecurityDomainArrayOutputWithContext(ctx context.Context) AviatrixSegmentationSecurityDomainArrayOutput {
	return o
}

func (o AviatrixSegmentationSecurityDomainArrayOutput) Index(i pulumi.IntInput) AviatrixSegmentationSecurityDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixSegmentationSecurityDomain {
		return vs[0].([]*AviatrixSegmentationSecurityDomain)[vs[1].(int)]
	}).(AviatrixSegmentationSecurityDomainOutput)
}

type AviatrixSegmentationSecurityDomainMapOutput struct{ *pulumi.OutputState }

func (AviatrixSegmentationSecurityDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixSegmentationSecurityDomain)(nil)).Elem()
}

func (o AviatrixSegmentationSecurityDomainMapOutput) ToAviatrixSegmentationSecurityDomainMapOutput() AviatrixSegmentationSecurityDomainMapOutput {
	return o
}

func (o AviatrixSegmentationSecurityDomainMapOutput) ToAviatrixSegmentationSecurityDomainMapOutputWithContext(ctx context.Context) AviatrixSegmentationSecurityDomainMapOutput {
	return o
}

func (o AviatrixSegmentationSecurityDomainMapOutput) MapIndex(k pulumi.StringInput) AviatrixSegmentationSecurityDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixSegmentationSecurityDomain {
		return vs[0].(map[string]*AviatrixSegmentationSecurityDomain)[vs[1].(string)]
	}).(AviatrixSegmentationSecurityDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSegmentationSecurityDomainInput)(nil)).Elem(), &AviatrixSegmentationSecurityDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSegmentationSecurityDomainArrayInput)(nil)).Elem(), AviatrixSegmentationSecurityDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixSegmentationSecurityDomainMapInput)(nil)).Elem(), AviatrixSegmentationSecurityDomainMap{})
	pulumi.RegisterOutputType(AviatrixSegmentationSecurityDomainOutput{})
	pulumi.RegisterOutputType(AviatrixSegmentationSecurityDomainArrayOutput{})
	pulumi.RegisterOutputType(AviatrixSegmentationSecurityDomainMapOutput{})
}
