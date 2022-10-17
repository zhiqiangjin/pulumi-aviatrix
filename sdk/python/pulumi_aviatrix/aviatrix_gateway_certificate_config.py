# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AviatrixGatewayCertificateConfigArgs', 'AviatrixGatewayCertificateConfig']

@pulumi.input_type
class AviatrixGatewayCertificateConfigArgs:
    def __init__(__self__, *,
                 ca_certificate: pulumi.Input[str],
                 ca_private_key: pulumi.Input[str]):
        """
        The set of arguments for constructing a AviatrixGatewayCertificateConfig resource.
        :param pulumi.Input[str] ca_certificate: CA Certificate.
        :param pulumi.Input[str] ca_private_key: CA Private Key.
        """
        pulumi.set(__self__, "ca_certificate", ca_certificate)
        pulumi.set(__self__, "ca_private_key", ca_private_key)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> pulumi.Input[str]:
        """
        CA Certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_certificate", value)

    @property
    @pulumi.getter(name="caPrivateKey")
    def ca_private_key(self) -> pulumi.Input[str]:
        """
        CA Private Key.
        """
        return pulumi.get(self, "ca_private_key")

    @ca_private_key.setter
    def ca_private_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_private_key", value)


@pulumi.input_type
class _AviatrixGatewayCertificateConfigState:
    def __init__(__self__, *,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 ca_private_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AviatrixGatewayCertificateConfig resources.
        :param pulumi.Input[str] ca_certificate: CA Certificate.
        :param pulumi.Input[str] ca_private_key: CA Private Key.
        """
        if ca_certificate is not None:
            pulumi.set(__self__, "ca_certificate", ca_certificate)
        if ca_private_key is not None:
            pulumi.set(__self__, "ca_private_key", ca_private_key)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        CA Certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @ca_certificate.setter
    def ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate", value)

    @property
    @pulumi.getter(name="caPrivateKey")
    def ca_private_key(self) -> Optional[pulumi.Input[str]]:
        """
        CA Private Key.
        """
        return pulumi.get(self, "ca_private_key")

    @ca_private_key.setter
    def ca_private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_private_key", value)


class AviatrixGatewayCertificateConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 ca_private_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AviatrixGatewayCertificateConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_certificate: CA Certificate.
        :param pulumi.Input[str] ca_private_key: CA Private Key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AviatrixGatewayCertificateConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AviatrixGatewayCertificateConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AviatrixGatewayCertificateConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AviatrixGatewayCertificateConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_certificate: Optional[pulumi.Input[str]] = None,
                 ca_private_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AviatrixGatewayCertificateConfigArgs.__new__(AviatrixGatewayCertificateConfigArgs)

            if ca_certificate is None and not opts.urn:
                raise TypeError("Missing required property 'ca_certificate'")
            __props__.__dict__["ca_certificate"] = ca_certificate
            if ca_private_key is None and not opts.urn:
                raise TypeError("Missing required property 'ca_private_key'")
            __props__.__dict__["ca_private_key"] = ca_private_key
        super(AviatrixGatewayCertificateConfig, __self__).__init__(
            'aviatrix:index/aviatrixGatewayCertificateConfig:AviatrixGatewayCertificateConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ca_certificate: Optional[pulumi.Input[str]] = None,
            ca_private_key: Optional[pulumi.Input[str]] = None) -> 'AviatrixGatewayCertificateConfig':
        """
        Get an existing AviatrixGatewayCertificateConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_certificate: CA Certificate.
        :param pulumi.Input[str] ca_private_key: CA Private Key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AviatrixGatewayCertificateConfigState.__new__(_AviatrixGatewayCertificateConfigState)

        __props__.__dict__["ca_certificate"] = ca_certificate
        __props__.__dict__["ca_private_key"] = ca_private_key
        return AviatrixGatewayCertificateConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="caCertificate")
    def ca_certificate(self) -> pulumi.Output[str]:
        """
        CA Certificate.
        """
        return pulumi.get(self, "ca_certificate")

    @property
    @pulumi.getter(name="caPrivateKey")
    def ca_private_key(self) -> pulumi.Output[str]:
        """
        CA Private Key.
        """
        return pulumi.get(self, "ca_private_key")

