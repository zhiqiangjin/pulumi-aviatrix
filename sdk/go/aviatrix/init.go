// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aviatrix:index/aviatrixAccount:AviatrixAccount":
		r = &AviatrixAccount{}
	case "aviatrix:index/aviatrixAccountUser:AviatrixAccountUser":
		r = &AviatrixAccountUser{}
	case "aviatrix:index/aviatrixAppDomain:AviatrixAppDomain":
		r = &AviatrixAppDomain{}
	case "aviatrix:index/aviatrixArmPeer:AviatrixArmPeer":
		r = &AviatrixArmPeer{}
	case "aviatrix:index/aviatrixAwsGuardDuty:AviatrixAwsGuardDuty":
		r = &AviatrixAwsGuardDuty{}
	case "aviatrix:index/aviatrixAwsPeer:AviatrixAwsPeer":
		r = &AviatrixAwsPeer{}
	case "aviatrix:index/aviatrixAwsTgw:AviatrixAwsTgw":
		r = &AviatrixAwsTgw{}
	case "aviatrix:index/aviatrixAwsTgwConnect:AviatrixAwsTgwConnect":
		r = &AviatrixAwsTgwConnect{}
	case "aviatrix:index/aviatrixAwsTgwConnectPeer:AviatrixAwsTgwConnectPeer":
		r = &AviatrixAwsTgwConnectPeer{}
	case "aviatrix:index/aviatrixAwsTgwDirectconnect:AviatrixAwsTgwDirectconnect":
		r = &AviatrixAwsTgwDirectconnect{}
	case "aviatrix:index/aviatrixAwsTgwIntraDomainInspection:AviatrixAwsTgwIntraDomainInspection":
		r = &AviatrixAwsTgwIntraDomainInspection{}
	case "aviatrix:index/aviatrixAwsTgwNetworkDomain:AviatrixAwsTgwNetworkDomain":
		r = &AviatrixAwsTgwNetworkDomain{}
	case "aviatrix:index/aviatrixAwsTgwPeering:AviatrixAwsTgwPeering":
		r = &AviatrixAwsTgwPeering{}
	case "aviatrix:index/aviatrixAwsTgwPeeringDomainConn:AviatrixAwsTgwPeeringDomainConn":
		r = &AviatrixAwsTgwPeeringDomainConn{}
	case "aviatrix:index/aviatrixAwsTgwSecurityDomain:AviatrixAwsTgwSecurityDomain":
		r = &AviatrixAwsTgwSecurityDomain{}
	case "aviatrix:index/aviatrixAwsTgwSecurityDomainConn:AviatrixAwsTgwSecurityDomainConn":
		r = &AviatrixAwsTgwSecurityDomainConn{}
	case "aviatrix:index/aviatrixAwsTgwTransitGatewayAttachment:AviatrixAwsTgwTransitGatewayAttachment":
		r = &AviatrixAwsTgwTransitGatewayAttachment{}
	case "aviatrix:index/aviatrixAwsTgwVpcAttachment:AviatrixAwsTgwVpcAttachment":
		r = &AviatrixAwsTgwVpcAttachment{}
	case "aviatrix:index/aviatrixAwsTgwVpnConn:AviatrixAwsTgwVpnConn":
		r = &AviatrixAwsTgwVpnConn{}
	case "aviatrix:index/aviatrixAzurePeer:AviatrixAzurePeer":
		r = &AviatrixAzurePeer{}
	case "aviatrix:index/aviatrixVpc:AviatrixVpc":
		r = &AviatrixVpc{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:aviatrix" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAccountUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAppDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixArmPeer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsGuardDuty",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsPeer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgw",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwConnect",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwConnectPeer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwDirectconnect",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwIntraDomainInspection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwNetworkDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwPeering",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwPeeringDomainConn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwSecurityDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwSecurityDomainConn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwTransitGatewayAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwVpcAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAwsTgwVpnConn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixAzurePeer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aviatrix",
		"index/aviatrixVpc",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"aviatrix",
		&pkg{version},
	)
}