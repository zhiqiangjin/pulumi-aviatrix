// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AviatrixAwsTgwVpnConn struct {
	pulumi.CustomResourceState

	// Unique name of the connection.
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
	// static VPN connection. Default value: 'dynamic'.
	ConnectionType pulumi.StringPtrOutput `pulumi:"connectionType"`
	// Enable Global Acceleration.
	EnableGlobalAcceleration pulumi.BoolPtrOutput `pulumi:"enableGlobalAcceleration"`
	// Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
	EnableLearnedCidrsApproval pulumi.BoolPtrOutput `pulumi:"enableLearnedCidrsApproval"`
	// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun1 pulumi.StringPtrOutput `pulumi:"insideIpCidrTun1"`
	// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun2 pulumi.StringPtrOutput `pulumi:"insideIpCidrTun2"`
	// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun1 pulumi.StringPtrOutput `pulumi:"preSharedKeyTun1"`
	// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun2 pulumi.StringPtrOutput `pulumi:"preSharedKeyTun2"`
	// Public IP address. Example: '40.0.0.0'.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
	RemoteAsNumber pulumi.StringPtrOutput `pulumi:"remoteAsNumber"`
	// Remote CIDRs joined as a string with ','. Required for a static VPN connection.
	RemoteCidr pulumi.StringPtrOutput `pulumi:"remoteCidr"`
	// The name of a route domain, to which the vpn will be attached.
	RouteDomainName pulumi.StringOutput `pulumi:"routeDomainName"`
	// This parameter represents the name of an AWS TGW.
	TgwName pulumi.StringOutput `pulumi:"tgwName"`
	// ID of the vpn connection.
	VpnId pulumi.StringOutput `pulumi:"vpnId"`
	// VPN tunnel data.
	VpnTunnelDatas AviatrixAwsTgwVpnConnVpnTunnelDataArrayOutput `pulumi:"vpnTunnelDatas"`
}

// NewAviatrixAwsTgwVpnConn registers a new resource with the given unique name, arguments, and options.
func NewAviatrixAwsTgwVpnConn(ctx *pulumi.Context,
	name string, args *AviatrixAwsTgwVpnConnArgs, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwVpnConn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionName'")
	}
	if args.PublicIp == nil {
		return nil, errors.New("invalid value for required argument 'PublicIp'")
	}
	if args.RouteDomainName == nil {
		return nil, errors.New("invalid value for required argument 'RouteDomainName'")
	}
	if args.TgwName == nil {
		return nil, errors.New("invalid value for required argument 'TgwName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AviatrixAwsTgwVpnConn
	err := ctx.RegisterResource("aviatrix:index/aviatrixAwsTgwVpnConn:AviatrixAwsTgwVpnConn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAviatrixAwsTgwVpnConn gets an existing AviatrixAwsTgwVpnConn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAviatrixAwsTgwVpnConn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AviatrixAwsTgwVpnConnState, opts ...pulumi.ResourceOption) (*AviatrixAwsTgwVpnConn, error) {
	var resource AviatrixAwsTgwVpnConn
	err := ctx.ReadResource("aviatrix:index/aviatrixAwsTgwVpnConn:AviatrixAwsTgwVpnConn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AviatrixAwsTgwVpnConn resources.
type aviatrixAwsTgwVpnConnState struct {
	// Unique name of the connection.
	ConnectionName *string `pulumi:"connectionName"`
	// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
	// static VPN connection. Default value: 'dynamic'.
	ConnectionType *string `pulumi:"connectionType"`
	// Enable Global Acceleration.
	EnableGlobalAcceleration *bool `pulumi:"enableGlobalAcceleration"`
	// Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
	EnableLearnedCidrsApproval *bool `pulumi:"enableLearnedCidrsApproval"`
	// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun1 *string `pulumi:"insideIpCidrTun1"`
	// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun2 *string `pulumi:"insideIpCidrTun2"`
	// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun1 *string `pulumi:"preSharedKeyTun1"`
	// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun2 *string `pulumi:"preSharedKeyTun2"`
	// Public IP address. Example: '40.0.0.0'.
	PublicIp *string `pulumi:"publicIp"`
	// AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
	RemoteAsNumber *string `pulumi:"remoteAsNumber"`
	// Remote CIDRs joined as a string with ','. Required for a static VPN connection.
	RemoteCidr *string `pulumi:"remoteCidr"`
	// The name of a route domain, to which the vpn will be attached.
	RouteDomainName *string `pulumi:"routeDomainName"`
	// This parameter represents the name of an AWS TGW.
	TgwName *string `pulumi:"tgwName"`
	// ID of the vpn connection.
	VpnId *string `pulumi:"vpnId"`
	// VPN tunnel data.
	VpnTunnelDatas []AviatrixAwsTgwVpnConnVpnTunnelData `pulumi:"vpnTunnelDatas"`
}

type AviatrixAwsTgwVpnConnState struct {
	// Unique name of the connection.
	ConnectionName pulumi.StringPtrInput
	// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
	// static VPN connection. Default value: 'dynamic'.
	ConnectionType pulumi.StringPtrInput
	// Enable Global Acceleration.
	EnableGlobalAcceleration pulumi.BoolPtrInput
	// Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
	EnableLearnedCidrsApproval pulumi.BoolPtrInput
	// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun1 pulumi.StringPtrInput
	// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun2 pulumi.StringPtrInput
	// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun1 pulumi.StringPtrInput
	// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun2 pulumi.StringPtrInput
	// Public IP address. Example: '40.0.0.0'.
	PublicIp pulumi.StringPtrInput
	// AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
	RemoteAsNumber pulumi.StringPtrInput
	// Remote CIDRs joined as a string with ','. Required for a static VPN connection.
	RemoteCidr pulumi.StringPtrInput
	// The name of a route domain, to which the vpn will be attached.
	RouteDomainName pulumi.StringPtrInput
	// This parameter represents the name of an AWS TGW.
	TgwName pulumi.StringPtrInput
	// ID of the vpn connection.
	VpnId pulumi.StringPtrInput
	// VPN tunnel data.
	VpnTunnelDatas AviatrixAwsTgwVpnConnVpnTunnelDataArrayInput
}

func (AviatrixAwsTgwVpnConnState) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwVpnConnState)(nil)).Elem()
}

type aviatrixAwsTgwVpnConnArgs struct {
	// Unique name of the connection.
	ConnectionName string `pulumi:"connectionName"`
	// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
	// static VPN connection. Default value: 'dynamic'.
	ConnectionType *string `pulumi:"connectionType"`
	// Enable Global Acceleration.
	EnableGlobalAcceleration *bool `pulumi:"enableGlobalAcceleration"`
	// Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
	EnableLearnedCidrsApproval *bool `pulumi:"enableLearnedCidrsApproval"`
	// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun1 *string `pulumi:"insideIpCidrTun1"`
	// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun2 *string `pulumi:"insideIpCidrTun2"`
	// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun1 *string `pulumi:"preSharedKeyTun1"`
	// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun2 *string `pulumi:"preSharedKeyTun2"`
	// Public IP address. Example: '40.0.0.0'.
	PublicIp string `pulumi:"publicIp"`
	// AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
	RemoteAsNumber *string `pulumi:"remoteAsNumber"`
	// Remote CIDRs joined as a string with ','. Required for a static VPN connection.
	RemoteCidr *string `pulumi:"remoteCidr"`
	// The name of a route domain, to which the vpn will be attached.
	RouteDomainName string `pulumi:"routeDomainName"`
	// This parameter represents the name of an AWS TGW.
	TgwName string `pulumi:"tgwName"`
}

// The set of arguments for constructing a AviatrixAwsTgwVpnConn resource.
type AviatrixAwsTgwVpnConnArgs struct {
	// Unique name of the connection.
	ConnectionName pulumi.StringInput
	// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
	// static VPN connection. Default value: 'dynamic'.
	ConnectionType pulumi.StringPtrInput
	// Enable Global Acceleration.
	EnableGlobalAcceleration pulumi.BoolPtrInput
	// Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
	EnableLearnedCidrsApproval pulumi.BoolPtrInput
	// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun1 pulumi.StringPtrInput
	// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
	InsideIpCidrTun2 pulumi.StringPtrInput
	// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun1 pulumi.StringPtrInput
	// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
	PreSharedKeyTun2 pulumi.StringPtrInput
	// Public IP address. Example: '40.0.0.0'.
	PublicIp pulumi.StringInput
	// AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
	RemoteAsNumber pulumi.StringPtrInput
	// Remote CIDRs joined as a string with ','. Required for a static VPN connection.
	RemoteCidr pulumi.StringPtrInput
	// The name of a route domain, to which the vpn will be attached.
	RouteDomainName pulumi.StringInput
	// This parameter represents the name of an AWS TGW.
	TgwName pulumi.StringInput
}

func (AviatrixAwsTgwVpnConnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aviatrixAwsTgwVpnConnArgs)(nil)).Elem()
}

type AviatrixAwsTgwVpnConnInput interface {
	pulumi.Input

	ToAviatrixAwsTgwVpnConnOutput() AviatrixAwsTgwVpnConnOutput
	ToAviatrixAwsTgwVpnConnOutputWithContext(ctx context.Context) AviatrixAwsTgwVpnConnOutput
}

func (*AviatrixAwsTgwVpnConn) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwVpnConn)(nil)).Elem()
}

func (i *AviatrixAwsTgwVpnConn) ToAviatrixAwsTgwVpnConnOutput() AviatrixAwsTgwVpnConnOutput {
	return i.ToAviatrixAwsTgwVpnConnOutputWithContext(context.Background())
}

func (i *AviatrixAwsTgwVpnConn) ToAviatrixAwsTgwVpnConnOutputWithContext(ctx context.Context) AviatrixAwsTgwVpnConnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwVpnConnOutput)
}

// AviatrixAwsTgwVpnConnArrayInput is an input type that accepts AviatrixAwsTgwVpnConnArray and AviatrixAwsTgwVpnConnArrayOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwVpnConnArrayInput` via:
//
//	AviatrixAwsTgwVpnConnArray{ AviatrixAwsTgwVpnConnArgs{...} }
type AviatrixAwsTgwVpnConnArrayInput interface {
	pulumi.Input

	ToAviatrixAwsTgwVpnConnArrayOutput() AviatrixAwsTgwVpnConnArrayOutput
	ToAviatrixAwsTgwVpnConnArrayOutputWithContext(context.Context) AviatrixAwsTgwVpnConnArrayOutput
}

type AviatrixAwsTgwVpnConnArray []AviatrixAwsTgwVpnConnInput

func (AviatrixAwsTgwVpnConnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwVpnConn)(nil)).Elem()
}

func (i AviatrixAwsTgwVpnConnArray) ToAviatrixAwsTgwVpnConnArrayOutput() AviatrixAwsTgwVpnConnArrayOutput {
	return i.ToAviatrixAwsTgwVpnConnArrayOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwVpnConnArray) ToAviatrixAwsTgwVpnConnArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwVpnConnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwVpnConnArrayOutput)
}

// AviatrixAwsTgwVpnConnMapInput is an input type that accepts AviatrixAwsTgwVpnConnMap and AviatrixAwsTgwVpnConnMapOutput values.
// You can construct a concrete instance of `AviatrixAwsTgwVpnConnMapInput` via:
//
//	AviatrixAwsTgwVpnConnMap{ "key": AviatrixAwsTgwVpnConnArgs{...} }
type AviatrixAwsTgwVpnConnMapInput interface {
	pulumi.Input

	ToAviatrixAwsTgwVpnConnMapOutput() AviatrixAwsTgwVpnConnMapOutput
	ToAviatrixAwsTgwVpnConnMapOutputWithContext(context.Context) AviatrixAwsTgwVpnConnMapOutput
}

type AviatrixAwsTgwVpnConnMap map[string]AviatrixAwsTgwVpnConnInput

func (AviatrixAwsTgwVpnConnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwVpnConn)(nil)).Elem()
}

func (i AviatrixAwsTgwVpnConnMap) ToAviatrixAwsTgwVpnConnMapOutput() AviatrixAwsTgwVpnConnMapOutput {
	return i.ToAviatrixAwsTgwVpnConnMapOutputWithContext(context.Background())
}

func (i AviatrixAwsTgwVpnConnMap) ToAviatrixAwsTgwVpnConnMapOutputWithContext(ctx context.Context) AviatrixAwsTgwVpnConnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AviatrixAwsTgwVpnConnMapOutput)
}

type AviatrixAwsTgwVpnConnOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwVpnConnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AviatrixAwsTgwVpnConn)(nil)).Elem()
}

func (o AviatrixAwsTgwVpnConnOutput) ToAviatrixAwsTgwVpnConnOutput() AviatrixAwsTgwVpnConnOutput {
	return o
}

func (o AviatrixAwsTgwVpnConnOutput) ToAviatrixAwsTgwVpnConnOutputWithContext(ctx context.Context) AviatrixAwsTgwVpnConnOutput {
	return o
}

// Unique name of the connection.
func (o AviatrixAwsTgwVpnConnOutput) ConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringOutput { return v.ConnectionName }).(pulumi.StringOutput)
}

// Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a
// static VPN connection. Default value: 'dynamic'.
func (o AviatrixAwsTgwVpnConnOutput) ConnectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringPtrOutput { return v.ConnectionType }).(pulumi.StringPtrOutput)
}

// Enable Global Acceleration.
func (o AviatrixAwsTgwVpnConnOutput) EnableGlobalAcceleration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.BoolPtrOutput { return v.EnableGlobalAcceleration }).(pulumi.BoolPtrOutput)
}

// Switch to enable/disable encrypted transit approval for vpn connection. Valid values: true, false.
func (o AviatrixAwsTgwVpnConnOutput) EnableLearnedCidrsApproval() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.BoolPtrOutput { return v.EnableLearnedCidrsApproval }).(pulumi.BoolPtrOutput)
}

// Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.
func (o AviatrixAwsTgwVpnConnOutput) InsideIpCidrTun1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringPtrOutput { return v.InsideIpCidrTun1 }).(pulumi.StringPtrOutput)
}

// Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.
func (o AviatrixAwsTgwVpnConnOutput) InsideIpCidrTun2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringPtrOutput { return v.InsideIpCidrTun2 }).(pulumi.StringPtrOutput)
}

// Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
func (o AviatrixAwsTgwVpnConnOutput) PreSharedKeyTun1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringPtrOutput { return v.PreSharedKeyTun1 }).(pulumi.StringPtrOutput)
}

// Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric, underscore(_) and dot(.). It cannot start with 0
func (o AviatrixAwsTgwVpnConnOutput) PreSharedKeyTun2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringPtrOutput { return v.PreSharedKeyTun2 }).(pulumi.StringPtrOutput)
}

// Public IP address. Example: '40.0.0.0'.
func (o AviatrixAwsTgwVpnConnOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

// AWS side as a number. Integer between 1-4294967294. Example: '12'. Required for a dynamic VPN connection.
func (o AviatrixAwsTgwVpnConnOutput) RemoteAsNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringPtrOutput { return v.RemoteAsNumber }).(pulumi.StringPtrOutput)
}

// Remote CIDRs joined as a string with ','. Required for a static VPN connection.
func (o AviatrixAwsTgwVpnConnOutput) RemoteCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringPtrOutput { return v.RemoteCidr }).(pulumi.StringPtrOutput)
}

// The name of a route domain, to which the vpn will be attached.
func (o AviatrixAwsTgwVpnConnOutput) RouteDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringOutput { return v.RouteDomainName }).(pulumi.StringOutput)
}

// This parameter represents the name of an AWS TGW.
func (o AviatrixAwsTgwVpnConnOutput) TgwName() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringOutput { return v.TgwName }).(pulumi.StringOutput)
}

// ID of the vpn connection.
func (o AviatrixAwsTgwVpnConnOutput) VpnId() pulumi.StringOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) pulumi.StringOutput { return v.VpnId }).(pulumi.StringOutput)
}

// VPN tunnel data.
func (o AviatrixAwsTgwVpnConnOutput) VpnTunnelDatas() AviatrixAwsTgwVpnConnVpnTunnelDataArrayOutput {
	return o.ApplyT(func(v *AviatrixAwsTgwVpnConn) AviatrixAwsTgwVpnConnVpnTunnelDataArrayOutput { return v.VpnTunnelDatas }).(AviatrixAwsTgwVpnConnVpnTunnelDataArrayOutput)
}

type AviatrixAwsTgwVpnConnArrayOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwVpnConnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AviatrixAwsTgwVpnConn)(nil)).Elem()
}

func (o AviatrixAwsTgwVpnConnArrayOutput) ToAviatrixAwsTgwVpnConnArrayOutput() AviatrixAwsTgwVpnConnArrayOutput {
	return o
}

func (o AviatrixAwsTgwVpnConnArrayOutput) ToAviatrixAwsTgwVpnConnArrayOutputWithContext(ctx context.Context) AviatrixAwsTgwVpnConnArrayOutput {
	return o
}

func (o AviatrixAwsTgwVpnConnArrayOutput) Index(i pulumi.IntInput) AviatrixAwsTgwVpnConnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AviatrixAwsTgwVpnConn {
		return vs[0].([]*AviatrixAwsTgwVpnConn)[vs[1].(int)]
	}).(AviatrixAwsTgwVpnConnOutput)
}

type AviatrixAwsTgwVpnConnMapOutput struct{ *pulumi.OutputState }

func (AviatrixAwsTgwVpnConnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AviatrixAwsTgwVpnConn)(nil)).Elem()
}

func (o AviatrixAwsTgwVpnConnMapOutput) ToAviatrixAwsTgwVpnConnMapOutput() AviatrixAwsTgwVpnConnMapOutput {
	return o
}

func (o AviatrixAwsTgwVpnConnMapOutput) ToAviatrixAwsTgwVpnConnMapOutputWithContext(ctx context.Context) AviatrixAwsTgwVpnConnMapOutput {
	return o
}

func (o AviatrixAwsTgwVpnConnMapOutput) MapIndex(k pulumi.StringInput) AviatrixAwsTgwVpnConnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AviatrixAwsTgwVpnConn {
		return vs[0].(map[string]*AviatrixAwsTgwVpnConn)[vs[1].(string)]
	}).(AviatrixAwsTgwVpnConnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwVpnConnInput)(nil)).Elem(), &AviatrixAwsTgwVpnConn{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwVpnConnArrayInput)(nil)).Elem(), AviatrixAwsTgwVpnConnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AviatrixAwsTgwVpnConnMapInput)(nil)).Elem(), AviatrixAwsTgwVpnConnMap{})
	pulumi.RegisterOutputType(AviatrixAwsTgwVpnConnOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwVpnConnArrayOutput{})
	pulumi.RegisterOutputType(AviatrixAwsTgwVpnConnMapOutput{})
}
