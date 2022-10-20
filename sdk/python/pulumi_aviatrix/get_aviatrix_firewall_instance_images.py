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

__all__ = [
    'GetAviatrixFirewallInstanceImagesResult',
    'AwaitableGetAviatrixFirewallInstanceImagesResult',
    'get_aviatrix_firewall_instance_images',
    'get_aviatrix_firewall_instance_images_output',
]

@pulumi.output_type
class GetAviatrixFirewallInstanceImagesResult:
    """
    A collection of values returned by getAviatrixFirewallInstanceImages.
    """
    def __init__(__self__, firewall_images=None, id=None, vpc_id=None):
        if firewall_images and not isinstance(firewall_images, list):
            raise TypeError("Expected argument 'firewall_images' to be a list")
        pulumi.set(__self__, "firewall_images", firewall_images)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="firewallImages")
    def firewall_images(self) -> Sequence['outputs.GetAviatrixFirewallInstanceImagesFirewallImageResult']:
        return pulumi.get(self, "firewall_images")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetAviatrixFirewallInstanceImagesResult(GetAviatrixFirewallInstanceImagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAviatrixFirewallInstanceImagesResult(
            firewall_images=self.firewall_images,
            id=self.id,
            vpc_id=self.vpc_id)


def get_aviatrix_firewall_instance_images(vpc_id: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAviatrixFirewallInstanceImagesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aviatrix:index/getAviatrixFirewallInstanceImages:getAviatrixFirewallInstanceImages', __args__, opts=opts, typ=GetAviatrixFirewallInstanceImagesResult).value

    return AwaitableGetAviatrixFirewallInstanceImagesResult(
        firewall_images=__ret__.firewall_images,
        id=__ret__.id,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_aviatrix_firewall_instance_images)
def get_aviatrix_firewall_instance_images_output(vpc_id: Optional[pulumi.Input[str]] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAviatrixFirewallInstanceImagesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...