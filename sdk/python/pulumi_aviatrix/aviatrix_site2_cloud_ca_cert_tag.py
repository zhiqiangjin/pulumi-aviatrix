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

__all__ = ['AviatrixSite2CloudCaCertTagArgs', 'AviatrixSite2CloudCaCertTag']

@pulumi.input_type
class AviatrixSite2CloudCaCertTagArgs:
    def __init__(__self__, *,
                 ca_certificates: pulumi.Input[Sequence[pulumi.Input['AviatrixSite2CloudCaCertTagCaCertificateArgs']]],
                 tag_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixSite2CloudCaCertTag resource.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixSite2CloudCaCertTagCaCertificateArgs']]] ca_certificates: A set of CA certificates.
        :param pulumi.Input[str] tag_name: Unique name of the ca cert tag.
        """
        pulumi.set(__self__, "ca_certificates", ca_certificates)
        pulumi.set(__self__, "tag_name", tag_name)

    @property
    @pulumi.getter(name="caCertificates")
    def ca_certificates(self) -> pulumi.Input[Sequence[pulumi.Input['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]:
        """
        A set of CA certificates.
        """
        return pulumi.get(self, "ca_certificates")

    @ca_certificates.setter
    def ca_certificates(self, value: pulumi.Input[Sequence[pulumi.Input['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]):
        pulumi.set(self, "ca_certificates", value)

    @property
    @pulumi.getter(name="tagName")
    def tag_name(self) -> pulumi.Input[str]:
        """
        Unique name of the ca cert tag.
        """
        return pulumi.get(self, "tag_name")

    @tag_name.setter
    def tag_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag_name", value)


@pulumi.input_type
class _AviatrixSite2CloudCaCertTagState:
    def __init__(__self__, *,
                 ca_certificates: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]] = None,
                 tag_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixSite2CloudCaCertTag resources.
        :param pulumi.Input[Sequence[pulumi.Input['AviatrixSite2CloudCaCertTagCaCertificateArgs']]] ca_certificates: A set of CA certificates.
        :param pulumi.Input[str] tag_name: Unique name of the ca cert tag.
        """
        if ca_certificates is not None:
            pulumi.set(__self__, "ca_certificates", ca_certificates)
        if tag_name is not None:
            pulumi.set(__self__, "tag_name", tag_name)

    @property
    @pulumi.getter(name="caCertificates")
    def ca_certificates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]]:
        """
        A set of CA certificates.
        """
        return pulumi.get(self, "ca_certificates")

    @ca_certificates.setter
    def ca_certificates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]]):
        pulumi.set(self, "ca_certificates", value)

    @property
    @pulumi.getter(name="tagName")
    def tag_name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the ca cert tag.
        """
        return pulumi.get(self, "tag_name")

    @tag_name.setter
    def tag_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_name", value)


class AviatrixSite2CloudCaCertTag(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]]] = None,
                 tag_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixSite2CloudCaCertTag resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]] ca_certificates: A set of CA certificates.
        :param pulumi.Input[str] tag_name: Unique name of the ca cert tag.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixSite2CloudCaCertTagArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixSite2CloudCaCertTag resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixSite2CloudCaCertTagArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixSite2CloudCaCertTagArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]]] = None,
                 tag_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixSite2CloudCaCertTagArgs.__new__(AviatrixSite2CloudCaCertTagArgs)

            if ca_certificates is None and not opts.urn:
                raise TypeError("Missing required property 'ca_certificates'")
            __props__.__dict__["ca_certificates"] = ca_certificates
            if tag_name is None and not opts.urn:
                raise TypeError("Missing required property 'tag_name'")
            __props__.__dict__["tag_name"] = tag_name
        super(AviatrixSite2CloudCaCertTag, __self__).__init__(
            'aviatrix:index/aviatrixSite2CloudCaCertTag:AviatrixSite2CloudCaCertTag',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ca_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]]] = None,
            tag_name: Optional[pulumi.Input[str]] = None) -> 'AviatrixSite2CloudCaCertTag':
        """
        Get an existing AviatrixSite2CloudCaCertTag resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AviatrixSite2CloudCaCertTagCaCertificateArgs']]]] ca_certificates: A set of CA certificates.
        :param pulumi.Input[str] tag_name: Unique name of the ca cert tag.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixSite2CloudCaCertTagState.__new__(_AviatrixSite2CloudCaCertTagState)

        __props__.__dict__["ca_certificates"] = ca_certificates
        __props__.__dict__["tag_name"] = tag_name
        return AviatrixSite2CloudCaCertTag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="caCertificates")
    def ca_certificates(self) -> pulumi.Output[Sequence['outputs.AviatrixSite2CloudCaCertTagCaCertificate']]:
        """
        A set of CA certificates.
        """
        return pulumi.get(self, "ca_certificates")

    @property
    @pulumi.getter(name="tagName")
    def tag_name(self) -> pulumi.Output[str]:
        """
        Unique name of the ca cert tag.
        """
        return pulumi.get(self, "tag_name")
