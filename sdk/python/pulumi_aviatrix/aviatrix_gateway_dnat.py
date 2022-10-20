# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AviatrixGatewayDnatArgs', 'AviatrixGatewayDnat']

@pulumi.input_type
class AviatrixGatewayDnatArgs:
    def __init__(__self__, *,
                 dnat_policies: pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatDnatPolicyArgs']]],
                 gw_name: pulumi.Input[str],
                 sync_to_ha: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AviatrixGatewayDnat resource.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatDnatPolicyArgs']]] dnat_policies: Policy rule to be applied to gateway.
        :param pulumi.Input[str] gw_name: Name of the gateway.
        :param pulumi.Input[bool] sync_to_ha: Whether to sync the policies to the HA gateway.
        """
        pulumi.set(__self__, "dnat_policies", dnat_policies)
        pulumi.set(__self__, "gw_name", gw_name)
        if sync_to_ha is not None:
            pulumi.set(__self__, "sync_to_ha", sync_to_ha)

    @property
    @pulumi.getter(name="dnatPolicies")
    def dnat_policies(self) -> pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatDnatPolicyArgs']]]:
        """
        Policy rule to be applied to gateway.
        """
        return pulumi.get(self, "dnat_policies")

    @dnat_policies.setter
    def dnat_policies(self, value: pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatDnatPolicyArgs']]]):
        pulumi.set(self, "dnat_policies", value)

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> pulumi.Input[str]:
        """
        Name of the gateway.
        """
        return pulumi.get(self, "gw_name")

    @gw_name.setter
    def gw_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "gw_name", value)

    @property
    @pulumi.getter(name="syncToHa")
    def sync_to_ha(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to sync the policies to the HA gateway.
        """
        return pulumi.get(self, "sync_to_ha")

    @sync_to_ha.setter
    def sync_to_ha(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sync_to_ha", value)


@pulumi.input_type
class _AviatrixGatewayDnatState:
    def __init__(__self__, *,
                 connection_policies: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatConnectionPolicyArgs']]]] = None,
                 dnat_policies: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatDnatPolicyArgs']]]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 interface_policies: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatInterfacePolicyArgs']]]] = None,
                 sync_to_ha: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AviatrixGatewayDnat resources.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatConnectionPolicyArgs']]] connection_policies: Computed attribute to store the previous connection policy.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatDnatPolicyArgs']]] dnat_policies: Policy rule to be applied to gateway.
        :param pulumi.Input[str] gw_name: Name of the gateway.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatInterfacePolicyArgs']]] interface_policies: Computed attribute to store the previous interface policy.
        :param pulumi.Input[bool] sync_to_ha: Whether to sync the policies to the HA gateway.
        """
        if connection_policies is not None:
            pulumi.set(__self__, "connection_policies", connection_policies)
        if dnat_policies is not None:
            pulumi.set(__self__, "dnat_policies", dnat_policies)
        if gw_name is not None:
            pulumi.set(__self__, "gw_name", gw_name)
        if interface_policies is not None:
            pulumi.set(__self__, "interface_policies", interface_policies)
        if sync_to_ha is not None:
            pulumi.set(__self__, "sync_to_ha", sync_to_ha)

    @property
    @pulumi.getter(name="connectionPolicies")
    def connection_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatConnectionPolicyArgs']]]]:
        """
        Computed attribute to store the previous connection policy.
        """
        return pulumi.get(self, "connection_policies")

    @connection_policies.setter
    def connection_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatConnectionPolicyArgs']]]]):
        pulumi.set(self, "connection_policies", value)

    @property
    @pulumi.getter(name="dnatPolicies")
    def dnat_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatDnatPolicyArgs']]]]:
        """
        Policy rule to be applied to gateway.
        """
        return pulumi.get(self, "dnat_policies")

    @dnat_policies.setter
    def dnat_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatDnatPolicyArgs']]]]):
        pulumi.set(self, "dnat_policies", value)

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the gateway.
        """
        return pulumi.get(self, "gw_name")

    @gw_name.setter
    def gw_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gw_name", value)

    @property
    @pulumi.getter(name="interfacePolicies")
    def interface_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatInterfacePolicyArgs']]]]:
        """
        Computed attribute to store the previous interface policy.
        """
        return pulumi.get(self, "interface_policies")

    @interface_policies.setter
    def interface_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixGatewayDnatInterfacePolicyArgs']]]]):
        pulumi.set(self, "interface_policies", value)

    @property
    @pulumi.getter(name="syncToHa")
    def sync_to_ha(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to sync the policies to the HA gateway.
        """
        return pulumi.get(self, "sync_to_ha")

    @sync_to_ha.setter
    def sync_to_ha(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sync_to_ha", value)


class AviatrixGatewayDnat(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dnat_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatDnatPolicyArgs']]]]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 sync_to_ha: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a AviatrixGatewayDnat resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatDnatPolicyArgs']]]] dnat_policies: Policy rule to be applied to gateway.
        :param pulumi.Input[str] gw_name: Name of the gateway.
        :param pulumi.Input[bool] sync_to_ha: Whether to sync the policies to the HA gateway.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixGatewayDnatArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixGatewayDnat resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixGatewayDnatArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixGatewayDnatArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dnat_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatDnatPolicyArgs']]]]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 sync_to_ha: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixGatewayDnatArgs.__new__(AviatrixGatewayDnatArgs)

            if dnat_policies is None and not opts.urn:
                raise TypeError("Missing required property 'dnat_policies'")
            __props__.__dict__["dnat_policies"] = dnat_policies
            if gw_name is None and not opts.urn:
                raise TypeError("Missing required property 'gw_name'")
            __props__.__dict__["gw_name"] = gw_name
            __props__.__dict__["sync_to_ha"] = sync_to_ha
            __props__.__dict__["connection_policies"] = None
            __props__.__dict__["interface_policies"] = None
        super(AviatrixGatewayDnat, __self__).__init__(
            'aviatrix:index/aviatrixGatewayDnat:AviatrixGatewayDnat',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatConnectionPolicyArgs']]]]] = None,
            dnat_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatDnatPolicyArgs']]]]] = None,
            gw_name: Optional[pulumi.Input[str]] = None,
            interface_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatInterfacePolicyArgs']]]]] = None,
            sync_to_ha: Optional[pulumi.Input[bool]] = None) -> 'AviatrixGatewayDnat':
        """
        Get an existing AviatrixGatewayDnat resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatConnectionPolicyArgs']]]] connection_policies: Computed attribute to store the previous connection policy.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatDnatPolicyArgs']]]] dnat_policies: Policy rule to be applied to gateway.
        :param pulumi.Input[str] gw_name: Name of the gateway.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixGatewayDnatInterfacePolicyArgs']]]] interface_policies: Computed attribute to store the previous interface policy.
        :param pulumi.Input[bool] sync_to_ha: Whether to sync the policies to the HA gateway.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixGatewayDnatState.__new__(_AviatrixGatewayDnatState)

        __props__.__dict__["connection_policies"] = connection_policies
        __props__.__dict__["dnat_policies"] = dnat_policies
        __props__.__dict__["gw_name"] = gw_name
        __props__.__dict__["interface_policies"] = interface_policies
        __props__.__dict__["sync_to_ha"] = sync_to_ha
        return AviatrixGatewayDnat(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionPolicies")
    def connection_policies(self) -> pulumi.Output[Sequence['outputs.AviatrixGatewayDnatConnectionPolicy']]:
        """
        Computed attribute to store the previous connection policy.
        """
        return pulumi.get(self, "connection_policies")

    @property
    @pulumi.getter(name="dnatPolicies")
    def dnat_policies(self) -> pulumi.Output[Sequence['outputs.AviatrixGatewayDnatDnatPolicy']]:
        """
        Policy rule to be applied to gateway.
        """
        return pulumi.get(self, "dnat_policies")

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> pulumi.Output[str]:
        """
        Name of the gateway.
        """
        return pulumi.get(self, "gw_name")

    @property
    @pulumi.getter(name="interfacePolicies")
    def interface_policies(self) -> pulumi.Output[Sequence['outputs.AviatrixGatewayDnatInterfacePolicy']]:
        """
        Computed attribute to store the previous interface policy.
        """
        return pulumi.get(self, "interface_policies")

    @property
    @pulumi.getter(name="syncToHa")
    def sync_to_ha(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to sync the policies to the HA gateway.
        """
        return pulumi.get(self, "sync_to_ha")
