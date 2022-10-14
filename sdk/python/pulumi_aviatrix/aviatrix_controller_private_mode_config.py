# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixControllerPrivateModeConfigArgs', 'AviatrixControllerPrivateModeConfig']

@pulumi.input_type
class AviatrixControllerPrivateModeConfigArgs:
    def __init__(__self__, *,
                 enable_private_mode: pulumi.Input[bool],
                 copilot_instance_id: Optional[pulumi.Input[str]] = None,
                 proxies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AviatrixControllerPrivateModeConfig resource.
        :param pulumi.Input[bool] enable_private_mode: Whether to enable Private Mode on the Controller.
        :param pulumi.Input[str] copilot_instance_id: Copilot instance ID to associate with the Controller for Private Mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] proxies: Set of proxies.
        """
        pulumi.set(__self__, "enable_private_mode", enable_private_mode)
        if copilot_instance_id is not None:
            pulumi.set(__self__, "copilot_instance_id", copilot_instance_id)
        if proxies is not None:
            pulumi.set(__self__, "proxies", proxies)

    @property
    @pulumi.getter(name="enablePrivateMode")
    def enable_private_mode(self) -> pulumi.Input[bool]:
        """
        Whether to enable Private Mode on the Controller.
        """
        return pulumi.get(self, "enable_private_mode")

    @enable_private_mode.setter
    def enable_private_mode(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable_private_mode", value)

    @property
    @pulumi.getter(name="copilotInstanceId")
    def copilot_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Copilot instance ID to associate with the Controller for Private Mode.
        """
        return pulumi.get(self, "copilot_instance_id")

    @copilot_instance_id.setter
    def copilot_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "copilot_instance_id", value)

    @property
    @pulumi.getter
    def proxies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of proxies.
        """
        return pulumi.get(self, "proxies")

    @proxies.setter
    def proxies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "proxies", value)


@pulumi.input_type
class _AviatrixControllerPrivateModeConfigState:
    def __init__(__self__, *,
                 copilot_instance_id: Optional[pulumi.Input[str]] = None,
                 enable_private_mode: Optional[pulumi.Input[bool]] = None,
                 proxies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AviatrixControllerPrivateModeConfig resources.
        :param pulumi.Input[str] copilot_instance_id: Copilot instance ID to associate with the Controller for Private Mode.
        :param pulumi.Input[bool] enable_private_mode: Whether to enable Private Mode on the Controller.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] proxies: Set of proxies.
        """
        if copilot_instance_id is not None:
            pulumi.set(__self__, "copilot_instance_id", copilot_instance_id)
        if enable_private_mode is not None:
            pulumi.set(__self__, "enable_private_mode", enable_private_mode)
        if proxies is not None:
            pulumi.set(__self__, "proxies", proxies)

    @property
    @pulumi.getter(name="copilotInstanceId")
    def copilot_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Copilot instance ID to associate with the Controller for Private Mode.
        """
        return pulumi.get(self, "copilot_instance_id")

    @copilot_instance_id.setter
    def copilot_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "copilot_instance_id", value)

    @property
    @pulumi.getter(name="enablePrivateMode")
    def enable_private_mode(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable Private Mode on the Controller.
        """
        return pulumi.get(self, "enable_private_mode")

    @enable_private_mode.setter
    def enable_private_mode(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_private_mode", value)

    @property
    @pulumi.getter
    def proxies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of proxies.
        """
        return pulumi.get(self, "proxies")

    @proxies.setter
    def proxies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "proxies", value)


class AviatrixControllerPrivateModeConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 copilot_instance_id: Optional[pulumi.Input[str]] = None,
                 enable_private_mode: Optional[pulumi.Input[bool]] = None,
                 proxies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a AviatrixControllerPrivateModeConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] copilot_instance_id: Copilot instance ID to associate with the Controller for Private Mode.
        :param pulumi.Input[bool] enable_private_mode: Whether to enable Private Mode on the Controller.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] proxies: Set of proxies.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixControllerPrivateModeConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixControllerPrivateModeConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixControllerPrivateModeConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixControllerPrivateModeConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 copilot_instance_id: Optional[pulumi.Input[str]] = None,
                 enable_private_mode: Optional[pulumi.Input[bool]] = None,
                 proxies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixControllerPrivateModeConfigArgs.__new__(AviatrixControllerPrivateModeConfigArgs)

            __props__.__dict__["copilot_instance_id"] = copilot_instance_id
            if enable_private_mode is None and not opts.urn:
                raise TypeError("Missing required property 'enable_private_mode'")
            __props__.__dict__["enable_private_mode"] = enable_private_mode
            __props__.__dict__["proxies"] = proxies
        super(AviatrixControllerPrivateModeConfig, __self__).__init__(
            'aviatrix:index/aviatrixControllerPrivateModeConfig:AviatrixControllerPrivateModeConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            copilot_instance_id: Optional[pulumi.Input[str]] = None,
            enable_private_mode: Optional[pulumi.Input[bool]] = None,
            proxies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AviatrixControllerPrivateModeConfig':
        """
        Get an existing AviatrixControllerPrivateModeConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] copilot_instance_id: Copilot instance ID to associate with the Controller for Private Mode.
        :param pulumi.Input[bool] enable_private_mode: Whether to enable Private Mode on the Controller.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] proxies: Set of proxies.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixControllerPrivateModeConfigState.__new__(_AviatrixControllerPrivateModeConfigState)

        __props__.__dict__["copilot_instance_id"] = copilot_instance_id
        __props__.__dict__["enable_private_mode"] = enable_private_mode
        __props__.__dict__["proxies"] = proxies
        return AviatrixControllerPrivateModeConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="copilotInstanceId")
    def copilot_instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        Copilot instance ID to associate with the Controller for Private Mode.
        """
        return pulumi.get(self, "copilot_instance_id")

    @property
    @pulumi.getter(name="enablePrivateMode")
    def enable_private_mode(self) -> pulumi.Output[bool]:
        """
        Whether to enable Private Mode on the Controller.
        """
        return pulumi.get(self, "enable_private_mode")

    @property
    @pulumi.getter
    def proxies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Set of proxies.
        """
        return pulumi.get(self, "proxies")
