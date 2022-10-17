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

__all__ = ['AviatrixFirewallArgs', 'AviatrixFirewall']

@pulumi.input_type
class AviatrixFirewallArgs:
    def __init__(__self__, *,
                 gw_name: pulumi.Input[str],
                 base_log_enabled: Optional[pulumi.Input[bool]] = None,
                 base_policy: Optional[pulumi.Input[str]] = None,
                 manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]] = None):
        """
        The set of arguments for constructing a AviatrixFirewall resource.
        :param pulumi.Input[str] gw_name: The name of gateway.
        :param pulumi.Input[bool] base_log_enabled: Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        :param pulumi.Input[str] base_policy: New base policy.
        :param pulumi.Input[bool] manage_firewall_policies: Enable to manage firewall policies via in-line rules. If false, policies must be managed using
               `aviatrix_firewall_policy` resources.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]] policies: New access policy for the gateway.
        """
        pulumi.set(__self__, "gw_name", gw_name)
        if base_log_enabled is not None:
            pulumi.set(__self__, "base_log_enabled", base_log_enabled)
        if base_policy is not None:
            pulumi.set(__self__, "base_policy", base_policy)
        if manage_firewall_policies is not None:
            pulumi.set(__self__, "manage_firewall_policies", manage_firewall_policies)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> pulumi.Input[str]:
        """
        The name of gateway.
        """
        return pulumi.get(self, "gw_name")

    @gw_name.setter
    def gw_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "gw_name", value)

    @property
    @pulumi.getter(name="baseLogEnabled")
    def base_log_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        """
        return pulumi.get(self, "base_log_enabled")

    @base_log_enabled.setter
    def base_log_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "base_log_enabled", value)

    @property
    @pulumi.getter(name="basePolicy")
    def base_policy(self) -> Optional[pulumi.Input[str]]:
        """
        New base policy.
        """
        return pulumi.get(self, "base_policy")

    @base_policy.setter
    def base_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_policy", value)

    @property
    @pulumi.getter(name="manageFirewallPolicies")
    def manage_firewall_policies(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable to manage firewall policies via in-line rules. If false, policies must be managed using
        `aviatrix_firewall_policy` resources.
        """
        return pulumi.get(self, "manage_firewall_policies")

    @manage_firewall_policies.setter
    def manage_firewall_policies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manage_firewall_policies", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]]:
        """
        New access policy for the gateway.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class _AviatrixFirewallState:
    def __init__(__self__, *,
                 base_log_enabled: Optional[pulumi.Input[bool]] = None,
                 base_policy: Optional[pulumi.Input[str]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]] = None):
        """
        Input properties used for looking up and filtering AviatrixFirewall resources.
        :param pulumi.Input[bool] base_log_enabled: Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        :param pulumi.Input[str] base_policy: New base policy.
        :param pulumi.Input[str] gw_name: The name of gateway.
        :param pulumi.Input[bool] manage_firewall_policies: Enable to manage firewall policies via in-line rules. If false, policies must be managed using
               `aviatrix_firewall_policy` resources.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]] policies: New access policy for the gateway.
        """
        if base_log_enabled is not None:
            pulumi.set(__self__, "base_log_enabled", base_log_enabled)
        if base_policy is not None:
            pulumi.set(__self__, "base_policy", base_policy)
        if gw_name is not None:
            pulumi.set(__self__, "gw_name", gw_name)
        if manage_firewall_policies is not None:
            pulumi.set(__self__, "manage_firewall_policies", manage_firewall_policies)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="baseLogEnabled")
    def base_log_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        """
        return pulumi.get(self, "base_log_enabled")

    @base_log_enabled.setter
    def base_log_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "base_log_enabled", value)

    @property
    @pulumi.getter(name="basePolicy")
    def base_policy(self) -> Optional[pulumi.Input[str]]:
        """
        New base policy.
        """
        return pulumi.get(self, "base_policy")

    @base_policy.setter
    def base_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_policy", value)

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of gateway.
        """
        return pulumi.get(self, "gw_name")

    @gw_name.setter
    def gw_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gw_name", value)

    @property
    @pulumi.getter(name="manageFirewallPolicies")
    def manage_firewall_policies(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable to manage firewall policies via in-line rules. If false, policies must be managed using
        `aviatrix_firewall_policy` resources.
        """
        return pulumi.get(self, "manage_firewall_policies")

    @manage_firewall_policies.setter
    def manage_firewall_policies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manage_firewall_policies", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]]:
        """
        New access policy for the gateway.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixFirewallPolicyArgs']]]]):
        pulumi.set(self, "policies", value)


class AviatrixFirewall(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_log_enabled: Optional[pulumi.Input[bool]] = None,
                 base_policy: Optional[pulumi.Input[str]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]]] = None,
                 __props__=None):
        """
        Create a AviatrixFirewall resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] base_log_enabled: Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        :param pulumi.Input[str] base_policy: New base policy.
        :param pulumi.Input[str] gw_name: The name of gateway.
        :param pulumi.Input[bool] manage_firewall_policies: Enable to manage firewall policies via in-line rules. If false, policies must be managed using
               `aviatrix_firewall_policy` resources.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]] policies: New access policy for the gateway.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixFirewallArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixFirewall resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixFirewallArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixFirewallArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_log_enabled: Optional[pulumi.Input[bool]] = None,
                 base_policy: Optional[pulumi.Input[str]] = None,
                 gw_name: Optional[pulumi.Input[str]] = None,
                 manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixFirewallArgs.__new__(AviatrixFirewallArgs)

            __props__.__dict__["base_log_enabled"] = base_log_enabled
            __props__.__dict__["base_policy"] = base_policy
            if gw_name is None and not opts.urn:
                raise TypeError("Missing required property 'gw_name'")
            __props__.__dict__["gw_name"] = gw_name
            __props__.__dict__["manage_firewall_policies"] = manage_firewall_policies
            __props__.__dict__["policies"] = policies
        super(AviatrixFirewall, __self__).__init__(
            'aviatrix:index/aviatrixFirewall:AviatrixFirewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            base_log_enabled: Optional[pulumi.Input[bool]] = None,
            base_policy: Optional[pulumi.Input[str]] = None,
            gw_name: Optional[pulumi.Input[str]] = None,
            manage_firewall_policies: Optional[pulumi.Input[bool]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]]] = None) -> 'AviatrixFirewall':
        """
        Get an existing AviatrixFirewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] base_log_enabled: Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        :param pulumi.Input[str] base_policy: New base policy.
        :param pulumi.Input[str] gw_name: The name of gateway.
        :param pulumi.Input[bool] manage_firewall_policies: Enable to manage firewall policies via in-line rules. If false, policies must be managed using
               `aviatrix_firewall_policy` resources.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixFirewallPolicyArgs']]]] policies: New access policy for the gateway.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixFirewallState.__new__(_AviatrixFirewallState)

        __props__.__dict__["base_log_enabled"] = base_log_enabled
        __props__.__dict__["base_policy"] = base_policy
        __props__.__dict__["gw_name"] = gw_name
        __props__.__dict__["manage_firewall_policies"] = manage_firewall_policies
        __props__.__dict__["policies"] = policies
        return AviatrixFirewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="baseLogEnabled")
    def base_log_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether enable logging or not. Valid values: true, false. Default value: false.
        """
        return pulumi.get(self, "base_log_enabled")

    @property
    @pulumi.getter(name="basePolicy")
    def base_policy(self) -> pulumi.Output[Optional[str]]:
        """
        New base policy.
        """
        return pulumi.get(self, "base_policy")

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> pulumi.Output[str]:
        """
        The name of gateway.
        """
        return pulumi.get(self, "gw_name")

    @property
    @pulumi.getter(name="manageFirewallPolicies")
    def manage_firewall_policies(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable to manage firewall policies via in-line rules. If false, policies must be managed using
        `aviatrix_firewall_policy` resources.
        """
        return pulumi.get(self, "manage_firewall_policies")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence['outputs.AviatrixFirewallPolicy']]]:
        """
        New access policy for the gateway.
        """
        return pulumi.get(self, "policies")

