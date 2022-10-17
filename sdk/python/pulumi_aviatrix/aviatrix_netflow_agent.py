# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixNetflowAgentArgs', 'AviatrixNetflowAgent']

@pulumi.input_type
class AviatrixNetflowAgentArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[int],
                 server_ip: pulumi.Input[str],
                 excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a AviatrixNetflowAgent resource.
        :param pulumi.Input[int] port: Netflow server port.
        :param pulumi.Input[str] server_ip: Netflow server IP address.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_gateways: List of excluded gateways.
        :param pulumi.Input[int] version: Netflow version.
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "server_ip", server_ip)
        if excluded_gateways is not None:
            pulumi.set(__self__, "excluded_gateways", excluded_gateways)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        Netflow server port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> pulumi.Input[str]:
        """
        Netflow server IP address.
        """
        return pulumi.get(self, "server_ip")

    @server_ip.setter
    def server_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_ip", value)

    @property
    @pulumi.getter(name="excludedGateways")
    def excluded_gateways(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of excluded gateways.
        """
        return pulumi.get(self, "excluded_gateways")

    @excluded_gateways.setter
    def excluded_gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_gateways", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        Netflow version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _AviatrixNetflowAgentState:
    def __init__(__self__, *,
                 excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AviatrixNetflowAgent resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_gateways: List of excluded gateways.
        :param pulumi.Input[int] port: Netflow server port.
        :param pulumi.Input[str] server_ip: Netflow server IP address.
        :param pulumi.Input[str] status: Enabled or not.
        :param pulumi.Input[int] version: Netflow version.
        """
        if excluded_gateways is not None:
            pulumi.set(__self__, "excluded_gateways", excluded_gateways)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_ip is not None:
            pulumi.set(__self__, "server_ip", server_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="excludedGateways")
    def excluded_gateways(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of excluded gateways.
        """
        return pulumi.get(self, "excluded_gateways")

    @excluded_gateways.setter
    def excluded_gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_gateways", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Netflow server port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Netflow server IP address.
        """
        return pulumi.get(self, "server_ip")

    @server_ip.setter
    def server_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_ip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enabled or not.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        Netflow version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class AviatrixNetflowAgent(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a AviatrixNetflowAgent resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_gateways: List of excluded gateways.
        :param pulumi.Input[int] port: Netflow server port.
        :param pulumi.Input[str] server_ip: Netflow server IP address.
        :param pulumi.Input[int] version: Netflow version.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixNetflowAgentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixNetflowAgent resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixNetflowAgentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixNetflowAgentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_ip: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixNetflowAgentArgs.__new__(AviatrixNetflowAgentArgs)

            __props__.__dict__["excluded_gateways"] = excluded_gateways
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if server_ip is None and not opts.urn:
                raise TypeError("Missing required property 'server_ip'")
            __props__.__dict__["server_ip"] = server_ip
            __props__.__dict__["version"] = version
            __props__.__dict__["status"] = None
        super(AviatrixNetflowAgent, __self__).__init__(
            'aviatrix:index/aviatrixNetflowAgent:AviatrixNetflowAgent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            excluded_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            port: Optional[pulumi.Input[int]] = None,
            server_ip: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'AviatrixNetflowAgent':
        """
        Get an existing AviatrixNetflowAgent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_gateways: List of excluded gateways.
        :param pulumi.Input[int] port: Netflow server port.
        :param pulumi.Input[str] server_ip: Netflow server IP address.
        :param pulumi.Input[str] status: Enabled or not.
        :param pulumi.Input[int] version: Netflow version.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixNetflowAgentState.__new__(_AviatrixNetflowAgentState)

        __props__.__dict__["excluded_gateways"] = excluded_gateways
        __props__.__dict__["port"] = port
        __props__.__dict__["server_ip"] = server_ip
        __props__.__dict__["status"] = status
        __props__.__dict__["version"] = version
        return AviatrixNetflowAgent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="excludedGateways")
    def excluded_gateways(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of excluded gateways.
        """
        return pulumi.get(self, "excluded_gateways")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Netflow server port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="serverIp")
    def server_ip(self) -> pulumi.Output[str]:
        """
        Netflow server IP address.
        """
        return pulumi.get(self, "server_ip")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enabled or not.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[int]]:
        """
        Netflow version.
        """
        return pulumi.get(self, "version")

