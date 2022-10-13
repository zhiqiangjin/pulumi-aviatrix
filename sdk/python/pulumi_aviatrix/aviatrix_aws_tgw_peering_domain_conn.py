# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixAwsTgwPeeringDomainConnArgs', 'AviatrixAwsTgwPeeringDomainConn']

@pulumi.input_type
class AviatrixAwsTgwPeeringDomainConnArgs:
    def __init__(__self__, *,
                 domain_name1: pulumi.Input[str],
                 domain_name2: pulumi.Input[str],
                 tgw_name1: pulumi.Input[str],
                 tgw_name2: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixAwsTgwPeeringDomainConn resource.
        :param pulumi.Input[str] domain_name1: The name of the source domain to make a connection.
        :param pulumi.Input[str] domain_name2: The name of the destination domain to make a connection.
        :param pulumi.Input[str] tgw_name1: The AWS tgw name of the source domain to make a connection.
        :param pulumi.Input[str] tgw_name2: The AWS tgw name of the destination domain to make a connection.
        """
        pulumi.set(__self__, "domain_name1", domain_name1)
        pulumi.set(__self__, "domain_name2", domain_name2)
        pulumi.set(__self__, "tgw_name1", tgw_name1)
        pulumi.set(__self__, "tgw_name2", tgw_name2)

    @property
    @pulumi.getter(name="domainName1")
    def domain_name1(self) -> pulumi.Input[str]:
        """
        The name of the source domain to make a connection.
        """
        return pulumi.get(self, "domain_name1")

    @domain_name1.setter
    def domain_name1(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name1", value)

    @property
    @pulumi.getter(name="domainName2")
    def domain_name2(self) -> pulumi.Input[str]:
        """
        The name of the destination domain to make a connection.
        """
        return pulumi.get(self, "domain_name2")

    @domain_name2.setter
    def domain_name2(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name2", value)

    @property
    @pulumi.getter(name="tgwName1")
    def tgw_name1(self) -> pulumi.Input[str]:
        """
        The AWS tgw name of the source domain to make a connection.
        """
        return pulumi.get(self, "tgw_name1")

    @tgw_name1.setter
    def tgw_name1(self, value: pulumi.Input[str]):
        pulumi.set(self, "tgw_name1", value)

    @property
    @pulumi.getter(name="tgwName2")
    def tgw_name2(self) -> pulumi.Input[str]:
        """
        The AWS tgw name of the destination domain to make a connection.
        """
        return pulumi.get(self, "tgw_name2")

    @tgw_name2.setter
    def tgw_name2(self, value: pulumi.Input[str]):
        pulumi.set(self, "tgw_name2", value)


@pulumi.input_type
class _AviatrixAwsTgwPeeringDomainConnState:
    def __init__(__self__, *,
                 domain_name1: Optional[pulumi.Input[str]] = None,
                 domain_name2: Optional[pulumi.Input[str]] = None,
                 tgw_name1: Optional[pulumi.Input[str]] = None,
                 tgw_name2: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixAwsTgwPeeringDomainConn resources.
        :param pulumi.Input[str] domain_name1: The name of the source domain to make a connection.
        :param pulumi.Input[str] domain_name2: The name of the destination domain to make a connection.
        :param pulumi.Input[str] tgw_name1: The AWS tgw name of the source domain to make a connection.
        :param pulumi.Input[str] tgw_name2: The AWS tgw name of the destination domain to make a connection.
        """
        if domain_name1 is not None:
            pulumi.set(__self__, "domain_name1", domain_name1)
        if domain_name2 is not None:
            pulumi.set(__self__, "domain_name2", domain_name2)
        if tgw_name1 is not None:
            pulumi.set(__self__, "tgw_name1", tgw_name1)
        if tgw_name2 is not None:
            pulumi.set(__self__, "tgw_name2", tgw_name2)

    @property
    @pulumi.getter(name="domainName1")
    def domain_name1(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the source domain to make a connection.
        """
        return pulumi.get(self, "domain_name1")

    @domain_name1.setter
    def domain_name1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name1", value)

    @property
    @pulumi.getter(name="domainName2")
    def domain_name2(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the destination domain to make a connection.
        """
        return pulumi.get(self, "domain_name2")

    @domain_name2.setter
    def domain_name2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name2", value)

    @property
    @pulumi.getter(name="tgwName1")
    def tgw_name1(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS tgw name of the source domain to make a connection.
        """
        return pulumi.get(self, "tgw_name1")

    @tgw_name1.setter
    def tgw_name1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tgw_name1", value)

    @property
    @pulumi.getter(name="tgwName2")
    def tgw_name2(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS tgw name of the destination domain to make a connection.
        """
        return pulumi.get(self, "tgw_name2")

    @tgw_name2.setter
    def tgw_name2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tgw_name2", value)


class AviatrixAwsTgwPeeringDomainConn(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name1: Optional[pulumi.Input[str]] = None,
                 domain_name2: Optional[pulumi.Input[str]] = None,
                 tgw_name1: Optional[pulumi.Input[str]] = None,
                 tgw_name2: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixAwsTgwPeeringDomainConn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name1: The name of the source domain to make a connection.
        :param pulumi.Input[str] domain_name2: The name of the destination domain to make a connection.
        :param pulumi.Input[str] tgw_name1: The AWS tgw name of the source domain to make a connection.
        :param pulumi.Input[str] tgw_name2: The AWS tgw name of the destination domain to make a connection.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixAwsTgwPeeringDomainConnArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixAwsTgwPeeringDomainConn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixAwsTgwPeeringDomainConnArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixAwsTgwPeeringDomainConnArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name1: Optional[pulumi.Input[str]] = None,
                 domain_name2: Optional[pulumi.Input[str]] = None,
                 tgw_name1: Optional[pulumi.Input[str]] = None,
                 tgw_name2: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixAwsTgwPeeringDomainConnArgs.__new__(AviatrixAwsTgwPeeringDomainConnArgs)

            if domain_name1 is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name1'")
            __props__.__dict__["domain_name1"] = domain_name1
            if domain_name2 is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name2'")
            __props__.__dict__["domain_name2"] = domain_name2
            if tgw_name1 is None and not opts.urn:
                raise TypeError("Missing required property 'tgw_name1'")
            __props__.__dict__["tgw_name1"] = tgw_name1
            if tgw_name2 is None and not opts.urn:
                raise TypeError("Missing required property 'tgw_name2'")
            __props__.__dict__["tgw_name2"] = tgw_name2
        super(AviatrixAwsTgwPeeringDomainConn, __self__).__init__(
            'aviatrix:index/aviatrixAwsTgwPeeringDomainConn:AviatrixAwsTgwPeeringDomainConn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_name1: Optional[pulumi.Input[str]] = None,
            domain_name2: Optional[pulumi.Input[str]] = None,
            tgw_name1: Optional[pulumi.Input[str]] = None,
            tgw_name2: Optional[pulumi.Input[str]] = None) -> 'AviatrixAwsTgwPeeringDomainConn':
        """
        Get an existing AviatrixAwsTgwPeeringDomainConn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name1: The name of the source domain to make a connection.
        :param pulumi.Input[str] domain_name2: The name of the destination domain to make a connection.
        :param pulumi.Input[str] tgw_name1: The AWS tgw name of the source domain to make a connection.
        :param pulumi.Input[str] tgw_name2: The AWS tgw name of the destination domain to make a connection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixAwsTgwPeeringDomainConnState.__new__(_AviatrixAwsTgwPeeringDomainConnState)

        __props__.__dict__["domain_name1"] = domain_name1
        __props__.__dict__["domain_name2"] = domain_name2
        __props__.__dict__["tgw_name1"] = tgw_name1
        __props__.__dict__["tgw_name2"] = tgw_name2
        return AviatrixAwsTgwPeeringDomainConn(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainName1")
    def domain_name1(self) -> pulumi.Output[str]:
        """
        The name of the source domain to make a connection.
        """
        return pulumi.get(self, "domain_name1")

    @property
    @pulumi.getter(name="domainName2")
    def domain_name2(self) -> pulumi.Output[str]:
        """
        The name of the destination domain to make a connection.
        """
        return pulumi.get(self, "domain_name2")

    @property
    @pulumi.getter(name="tgwName1")
    def tgw_name1(self) -> pulumi.Output[str]:
        """
        The AWS tgw name of the source domain to make a connection.
        """
        return pulumi.get(self, "tgw_name1")

    @property
    @pulumi.getter(name="tgwName2")
    def tgw_name2(self) -> pulumi.Output[str]:
        """
        The AWS tgw name of the destination domain to make a connection.
        """
        return pulumi.get(self, "tgw_name2")

