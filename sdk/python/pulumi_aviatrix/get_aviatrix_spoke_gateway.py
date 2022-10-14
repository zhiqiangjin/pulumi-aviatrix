# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAviatrixSpokeGatewayResult',
    'AwaitableGetAviatrixSpokeGatewayResult',
    'get_aviatrix_spoke_gateway',
    'get_aviatrix_spoke_gateway_output',
]

@pulumi.output_type
class GetAviatrixSpokeGatewayResult:
    """
    A collection of values returned by getAviatrixSpokeGateway.
    """
    def __init__(__self__, account_name=None, allocate_new_eip=None, approved_learned_cidrs=None, availability_domain=None, azure_eip_name_resource_group=None, bgp_ecmp=None, bgp_hold_time=None, bgp_polling_time=None, cloud_instance_id=None, cloud_type=None, customized_spoke_vpc_routes=None, disable_route_propagation=None, eip=None, enable_active_standby=None, enable_active_standby_preemptive=None, enable_auto_advertise_s2c_cidrs=None, enable_bgp=None, enable_encrypt_volume=None, enable_jumbo_frame=None, enable_learned_cidrs_approval=None, enable_monitor_gateway_subnets=None, enable_private_oob=None, enable_private_vpc_default_route=None, enable_skip_public_route_table_update=None, enable_spot_instance=None, enable_vpc_dns_server=None, fault_domain=None, filtered_spoke_vpc_routes=None, gw_name=None, gw_size=None, ha_availability_domain=None, ha_azure_eip_name_resource_group=None, ha_cloud_instance_id=None, ha_eip=None, ha_fault_domain=None, ha_gw_name=None, ha_gw_size=None, ha_image_version=None, ha_insane_mode_az=None, ha_oob_availability_zone=None, ha_oob_management_subnet=None, ha_private_ip=None, ha_public_ip=None, ha_security_group_id=None, ha_software_version=None, ha_subnet=None, ha_zone=None, id=None, image_version=None, included_advertised_spoke_routes=None, insane_mode=None, insane_mode_az=None, learned_cidrs_approval_mode=None, local_as_number=None, monitor_exclude_lists=None, oob_availability_zone=None, oob_management_subnet=None, prepend_as_paths=None, private_ip=None, public_ip=None, security_group_id=None, single_az_ha=None, single_ip_snat=None, software_version=None, spoke_bgp_manual_advertise_cidrs=None, spot_price=None, subnet=None, tag_lists=None, tags=None, transit_gw=None, tunnel_detection_time=None, vpc_id=None, vpc_reg=None, zone=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        pulumi.set(__self__, "account_name", account_name)
        if allocate_new_eip and not isinstance(allocate_new_eip, bool):
            raise TypeError("Expected argument 'allocate_new_eip' to be a bool")
        pulumi.set(__self__, "allocate_new_eip", allocate_new_eip)
        if approved_learned_cidrs and not isinstance(approved_learned_cidrs, list):
            raise TypeError("Expected argument 'approved_learned_cidrs' to be a list")
        pulumi.set(__self__, "approved_learned_cidrs", approved_learned_cidrs)
        if availability_domain and not isinstance(availability_domain, str):
            raise TypeError("Expected argument 'availability_domain' to be a str")
        pulumi.set(__self__, "availability_domain", availability_domain)
        if azure_eip_name_resource_group and not isinstance(azure_eip_name_resource_group, str):
            raise TypeError("Expected argument 'azure_eip_name_resource_group' to be a str")
        pulumi.set(__self__, "azure_eip_name_resource_group", azure_eip_name_resource_group)
        if bgp_ecmp and not isinstance(bgp_ecmp, bool):
            raise TypeError("Expected argument 'bgp_ecmp' to be a bool")
        pulumi.set(__self__, "bgp_ecmp", bgp_ecmp)
        if bgp_hold_time and not isinstance(bgp_hold_time, int):
            raise TypeError("Expected argument 'bgp_hold_time' to be a int")
        pulumi.set(__self__, "bgp_hold_time", bgp_hold_time)
        if bgp_polling_time and not isinstance(bgp_polling_time, int):
            raise TypeError("Expected argument 'bgp_polling_time' to be a int")
        pulumi.set(__self__, "bgp_polling_time", bgp_polling_time)
        if cloud_instance_id and not isinstance(cloud_instance_id, str):
            raise TypeError("Expected argument 'cloud_instance_id' to be a str")
        pulumi.set(__self__, "cloud_instance_id", cloud_instance_id)
        if cloud_type and not isinstance(cloud_type, int):
            raise TypeError("Expected argument 'cloud_type' to be a int")
        pulumi.set(__self__, "cloud_type", cloud_type)
        if customized_spoke_vpc_routes and not isinstance(customized_spoke_vpc_routes, str):
            raise TypeError("Expected argument 'customized_spoke_vpc_routes' to be a str")
        pulumi.set(__self__, "customized_spoke_vpc_routes", customized_spoke_vpc_routes)
        if disable_route_propagation and not isinstance(disable_route_propagation, bool):
            raise TypeError("Expected argument 'disable_route_propagation' to be a bool")
        pulumi.set(__self__, "disable_route_propagation", disable_route_propagation)
        if eip and not isinstance(eip, str):
            raise TypeError("Expected argument 'eip' to be a str")
        pulumi.set(__self__, "eip", eip)
        if enable_active_standby and not isinstance(enable_active_standby, bool):
            raise TypeError("Expected argument 'enable_active_standby' to be a bool")
        pulumi.set(__self__, "enable_active_standby", enable_active_standby)
        if enable_active_standby_preemptive and not isinstance(enable_active_standby_preemptive, bool):
            raise TypeError("Expected argument 'enable_active_standby_preemptive' to be a bool")
        pulumi.set(__self__, "enable_active_standby_preemptive", enable_active_standby_preemptive)
        if enable_auto_advertise_s2c_cidrs and not isinstance(enable_auto_advertise_s2c_cidrs, bool):
            raise TypeError("Expected argument 'enable_auto_advertise_s2c_cidrs' to be a bool")
        pulumi.set(__self__, "enable_auto_advertise_s2c_cidrs", enable_auto_advertise_s2c_cidrs)
        if enable_bgp and not isinstance(enable_bgp, bool):
            raise TypeError("Expected argument 'enable_bgp' to be a bool")
        pulumi.set(__self__, "enable_bgp", enable_bgp)
        if enable_encrypt_volume and not isinstance(enable_encrypt_volume, bool):
            raise TypeError("Expected argument 'enable_encrypt_volume' to be a bool")
        pulumi.set(__self__, "enable_encrypt_volume", enable_encrypt_volume)
        if enable_jumbo_frame and not isinstance(enable_jumbo_frame, bool):
            raise TypeError("Expected argument 'enable_jumbo_frame' to be a bool")
        pulumi.set(__self__, "enable_jumbo_frame", enable_jumbo_frame)
        if enable_learned_cidrs_approval and not isinstance(enable_learned_cidrs_approval, bool):
            raise TypeError("Expected argument 'enable_learned_cidrs_approval' to be a bool")
        pulumi.set(__self__, "enable_learned_cidrs_approval", enable_learned_cidrs_approval)
        if enable_monitor_gateway_subnets and not isinstance(enable_monitor_gateway_subnets, bool):
            raise TypeError("Expected argument 'enable_monitor_gateway_subnets' to be a bool")
        pulumi.set(__self__, "enable_monitor_gateway_subnets", enable_monitor_gateway_subnets)
        if enable_private_oob and not isinstance(enable_private_oob, bool):
            raise TypeError("Expected argument 'enable_private_oob' to be a bool")
        pulumi.set(__self__, "enable_private_oob", enable_private_oob)
        if enable_private_vpc_default_route and not isinstance(enable_private_vpc_default_route, bool):
            raise TypeError("Expected argument 'enable_private_vpc_default_route' to be a bool")
        pulumi.set(__self__, "enable_private_vpc_default_route", enable_private_vpc_default_route)
        if enable_skip_public_route_table_update and not isinstance(enable_skip_public_route_table_update, bool):
            raise TypeError("Expected argument 'enable_skip_public_route_table_update' to be a bool")
        pulumi.set(__self__, "enable_skip_public_route_table_update", enable_skip_public_route_table_update)
        if enable_spot_instance and not isinstance(enable_spot_instance, bool):
            raise TypeError("Expected argument 'enable_spot_instance' to be a bool")
        pulumi.set(__self__, "enable_spot_instance", enable_spot_instance)
        if enable_vpc_dns_server and not isinstance(enable_vpc_dns_server, bool):
            raise TypeError("Expected argument 'enable_vpc_dns_server' to be a bool")
        pulumi.set(__self__, "enable_vpc_dns_server", enable_vpc_dns_server)
        if fault_domain and not isinstance(fault_domain, str):
            raise TypeError("Expected argument 'fault_domain' to be a str")
        pulumi.set(__self__, "fault_domain", fault_domain)
        if filtered_spoke_vpc_routes and not isinstance(filtered_spoke_vpc_routes, str):
            raise TypeError("Expected argument 'filtered_spoke_vpc_routes' to be a str")
        pulumi.set(__self__, "filtered_spoke_vpc_routes", filtered_spoke_vpc_routes)
        if gw_name and not isinstance(gw_name, str):
            raise TypeError("Expected argument 'gw_name' to be a str")
        pulumi.set(__self__, "gw_name", gw_name)
        if gw_size and not isinstance(gw_size, str):
            raise TypeError("Expected argument 'gw_size' to be a str")
        pulumi.set(__self__, "gw_size", gw_size)
        if ha_availability_domain and not isinstance(ha_availability_domain, str):
            raise TypeError("Expected argument 'ha_availability_domain' to be a str")
        pulumi.set(__self__, "ha_availability_domain", ha_availability_domain)
        if ha_azure_eip_name_resource_group and not isinstance(ha_azure_eip_name_resource_group, str):
            raise TypeError("Expected argument 'ha_azure_eip_name_resource_group' to be a str")
        pulumi.set(__self__, "ha_azure_eip_name_resource_group", ha_azure_eip_name_resource_group)
        if ha_cloud_instance_id and not isinstance(ha_cloud_instance_id, str):
            raise TypeError("Expected argument 'ha_cloud_instance_id' to be a str")
        pulumi.set(__self__, "ha_cloud_instance_id", ha_cloud_instance_id)
        if ha_eip and not isinstance(ha_eip, str):
            raise TypeError("Expected argument 'ha_eip' to be a str")
        pulumi.set(__self__, "ha_eip", ha_eip)
        if ha_fault_domain and not isinstance(ha_fault_domain, str):
            raise TypeError("Expected argument 'ha_fault_domain' to be a str")
        pulumi.set(__self__, "ha_fault_domain", ha_fault_domain)
        if ha_gw_name and not isinstance(ha_gw_name, str):
            raise TypeError("Expected argument 'ha_gw_name' to be a str")
        pulumi.set(__self__, "ha_gw_name", ha_gw_name)
        if ha_gw_size and not isinstance(ha_gw_size, str):
            raise TypeError("Expected argument 'ha_gw_size' to be a str")
        pulumi.set(__self__, "ha_gw_size", ha_gw_size)
        if ha_image_version and not isinstance(ha_image_version, str):
            raise TypeError("Expected argument 'ha_image_version' to be a str")
        pulumi.set(__self__, "ha_image_version", ha_image_version)
        if ha_insane_mode_az and not isinstance(ha_insane_mode_az, str):
            raise TypeError("Expected argument 'ha_insane_mode_az' to be a str")
        pulumi.set(__self__, "ha_insane_mode_az", ha_insane_mode_az)
        if ha_oob_availability_zone and not isinstance(ha_oob_availability_zone, str):
            raise TypeError("Expected argument 'ha_oob_availability_zone' to be a str")
        pulumi.set(__self__, "ha_oob_availability_zone", ha_oob_availability_zone)
        if ha_oob_management_subnet and not isinstance(ha_oob_management_subnet, str):
            raise TypeError("Expected argument 'ha_oob_management_subnet' to be a str")
        pulumi.set(__self__, "ha_oob_management_subnet", ha_oob_management_subnet)
        if ha_private_ip and not isinstance(ha_private_ip, str):
            raise TypeError("Expected argument 'ha_private_ip' to be a str")
        pulumi.set(__self__, "ha_private_ip", ha_private_ip)
        if ha_public_ip and not isinstance(ha_public_ip, str):
            raise TypeError("Expected argument 'ha_public_ip' to be a str")
        pulumi.set(__self__, "ha_public_ip", ha_public_ip)
        if ha_security_group_id and not isinstance(ha_security_group_id, str):
            raise TypeError("Expected argument 'ha_security_group_id' to be a str")
        pulumi.set(__self__, "ha_security_group_id", ha_security_group_id)
        if ha_software_version and not isinstance(ha_software_version, str):
            raise TypeError("Expected argument 'ha_software_version' to be a str")
        pulumi.set(__self__, "ha_software_version", ha_software_version)
        if ha_subnet and not isinstance(ha_subnet, str):
            raise TypeError("Expected argument 'ha_subnet' to be a str")
        pulumi.set(__self__, "ha_subnet", ha_subnet)
        if ha_zone and not isinstance(ha_zone, str):
            raise TypeError("Expected argument 'ha_zone' to be a str")
        pulumi.set(__self__, "ha_zone", ha_zone)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_version and not isinstance(image_version, str):
            raise TypeError("Expected argument 'image_version' to be a str")
        pulumi.set(__self__, "image_version", image_version)
        if included_advertised_spoke_routes and not isinstance(included_advertised_spoke_routes, str):
            raise TypeError("Expected argument 'included_advertised_spoke_routes' to be a str")
        pulumi.set(__self__, "included_advertised_spoke_routes", included_advertised_spoke_routes)
        if insane_mode and not isinstance(insane_mode, bool):
            raise TypeError("Expected argument 'insane_mode' to be a bool")
        pulumi.set(__self__, "insane_mode", insane_mode)
        if insane_mode_az and not isinstance(insane_mode_az, str):
            raise TypeError("Expected argument 'insane_mode_az' to be a str")
        pulumi.set(__self__, "insane_mode_az", insane_mode_az)
        if learned_cidrs_approval_mode and not isinstance(learned_cidrs_approval_mode, str):
            raise TypeError("Expected argument 'learned_cidrs_approval_mode' to be a str")
        pulumi.set(__self__, "learned_cidrs_approval_mode", learned_cidrs_approval_mode)
        if local_as_number and not isinstance(local_as_number, str):
            raise TypeError("Expected argument 'local_as_number' to be a str")
        pulumi.set(__self__, "local_as_number", local_as_number)
        if monitor_exclude_lists and not isinstance(monitor_exclude_lists, list):
            raise TypeError("Expected argument 'monitor_exclude_lists' to be a list")
        pulumi.set(__self__, "monitor_exclude_lists", monitor_exclude_lists)
        if oob_availability_zone and not isinstance(oob_availability_zone, str):
            raise TypeError("Expected argument 'oob_availability_zone' to be a str")
        pulumi.set(__self__, "oob_availability_zone", oob_availability_zone)
        if oob_management_subnet and not isinstance(oob_management_subnet, str):
            raise TypeError("Expected argument 'oob_management_subnet' to be a str")
        pulumi.set(__self__, "oob_management_subnet", oob_management_subnet)
        if prepend_as_paths and not isinstance(prepend_as_paths, list):
            raise TypeError("Expected argument 'prepend_as_paths' to be a list")
        pulumi.set(__self__, "prepend_as_paths", prepend_as_paths)
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        pulumi.set(__self__, "private_ip", private_ip)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if single_az_ha and not isinstance(single_az_ha, bool):
            raise TypeError("Expected argument 'single_az_ha' to be a bool")
        pulumi.set(__self__, "single_az_ha", single_az_ha)
        if single_ip_snat and not isinstance(single_ip_snat, bool):
            raise TypeError("Expected argument 'single_ip_snat' to be a bool")
        pulumi.set(__self__, "single_ip_snat", single_ip_snat)
        if software_version and not isinstance(software_version, str):
            raise TypeError("Expected argument 'software_version' to be a str")
        pulumi.set(__self__, "software_version", software_version)
        if spoke_bgp_manual_advertise_cidrs and not isinstance(spoke_bgp_manual_advertise_cidrs, list):
            raise TypeError("Expected argument 'spoke_bgp_manual_advertise_cidrs' to be a list")
        pulumi.set(__self__, "spoke_bgp_manual_advertise_cidrs", spoke_bgp_manual_advertise_cidrs)
        if spot_price and not isinstance(spot_price, str):
            raise TypeError("Expected argument 'spot_price' to be a str")
        pulumi.set(__self__, "spot_price", spot_price)
        if subnet and not isinstance(subnet, str):
            raise TypeError("Expected argument 'subnet' to be a str")
        pulumi.set(__self__, "subnet", subnet)
        if tag_lists and not isinstance(tag_lists, list):
            raise TypeError("Expected argument 'tag_lists' to be a list")
        pulumi.set(__self__, "tag_lists", tag_lists)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if transit_gw and not isinstance(transit_gw, str):
            raise TypeError("Expected argument 'transit_gw' to be a str")
        pulumi.set(__self__, "transit_gw", transit_gw)
        if tunnel_detection_time and not isinstance(tunnel_detection_time, int):
            raise TypeError("Expected argument 'tunnel_detection_time' to be a int")
        pulumi.set(__self__, "tunnel_detection_time", tunnel_detection_time)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_reg and not isinstance(vpc_reg, str):
            raise TypeError("Expected argument 'vpc_reg' to be a str")
        pulumi.set(__self__, "vpc_reg", vpc_reg)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="allocateNewEip")
    def allocate_new_eip(self) -> bool:
        return pulumi.get(self, "allocate_new_eip")

    @property
    @pulumi.getter(name="approvedLearnedCidrs")
    def approved_learned_cidrs(self) -> Sequence[str]:
        return pulumi.get(self, "approved_learned_cidrs")

    @property
    @pulumi.getter(name="availabilityDomain")
    def availability_domain(self) -> str:
        return pulumi.get(self, "availability_domain")

    @property
    @pulumi.getter(name="azureEipNameResourceGroup")
    def azure_eip_name_resource_group(self) -> str:
        return pulumi.get(self, "azure_eip_name_resource_group")

    @property
    @pulumi.getter(name="bgpEcmp")
    def bgp_ecmp(self) -> bool:
        return pulumi.get(self, "bgp_ecmp")

    @property
    @pulumi.getter(name="bgpHoldTime")
    def bgp_hold_time(self) -> int:
        return pulumi.get(self, "bgp_hold_time")

    @property
    @pulumi.getter(name="bgpPollingTime")
    def bgp_polling_time(self) -> int:
        return pulumi.get(self, "bgp_polling_time")

    @property
    @pulumi.getter(name="cloudInstanceId")
    def cloud_instance_id(self) -> str:
        return pulumi.get(self, "cloud_instance_id")

    @property
    @pulumi.getter(name="cloudType")
    def cloud_type(self) -> int:
        return pulumi.get(self, "cloud_type")

    @property
    @pulumi.getter(name="customizedSpokeVpcRoutes")
    def customized_spoke_vpc_routes(self) -> str:
        return pulumi.get(self, "customized_spoke_vpc_routes")

    @property
    @pulumi.getter(name="disableRoutePropagation")
    def disable_route_propagation(self) -> bool:
        return pulumi.get(self, "disable_route_propagation")

    @property
    @pulumi.getter
    def eip(self) -> str:
        return pulumi.get(self, "eip")

    @property
    @pulumi.getter(name="enableActiveStandby")
    def enable_active_standby(self) -> bool:
        return pulumi.get(self, "enable_active_standby")

    @property
    @pulumi.getter(name="enableActiveStandbyPreemptive")
    def enable_active_standby_preemptive(self) -> bool:
        return pulumi.get(self, "enable_active_standby_preemptive")

    @property
    @pulumi.getter(name="enableAutoAdvertiseS2cCidrs")
    def enable_auto_advertise_s2c_cidrs(self) -> bool:
        return pulumi.get(self, "enable_auto_advertise_s2c_cidrs")

    @property
    @pulumi.getter(name="enableBgp")
    def enable_bgp(self) -> bool:
        return pulumi.get(self, "enable_bgp")

    @property
    @pulumi.getter(name="enableEncryptVolume")
    def enable_encrypt_volume(self) -> bool:
        return pulumi.get(self, "enable_encrypt_volume")

    @property
    @pulumi.getter(name="enableJumboFrame")
    def enable_jumbo_frame(self) -> bool:
        return pulumi.get(self, "enable_jumbo_frame")

    @property
    @pulumi.getter(name="enableLearnedCidrsApproval")
    def enable_learned_cidrs_approval(self) -> bool:
        return pulumi.get(self, "enable_learned_cidrs_approval")

    @property
    @pulumi.getter(name="enableMonitorGatewaySubnets")
    def enable_monitor_gateway_subnets(self) -> bool:
        return pulumi.get(self, "enable_monitor_gateway_subnets")

    @property
    @pulumi.getter(name="enablePrivateOob")
    def enable_private_oob(self) -> bool:
        return pulumi.get(self, "enable_private_oob")

    @property
    @pulumi.getter(name="enablePrivateVpcDefaultRoute")
    def enable_private_vpc_default_route(self) -> bool:
        return pulumi.get(self, "enable_private_vpc_default_route")

    @property
    @pulumi.getter(name="enableSkipPublicRouteTableUpdate")
    def enable_skip_public_route_table_update(self) -> bool:
        return pulumi.get(self, "enable_skip_public_route_table_update")

    @property
    @pulumi.getter(name="enableSpotInstance")
    def enable_spot_instance(self) -> bool:
        return pulumi.get(self, "enable_spot_instance")

    @property
    @pulumi.getter(name="enableVpcDnsServer")
    def enable_vpc_dns_server(self) -> bool:
        return pulumi.get(self, "enable_vpc_dns_server")

    @property
    @pulumi.getter(name="faultDomain")
    def fault_domain(self) -> str:
        return pulumi.get(self, "fault_domain")

    @property
    @pulumi.getter(name="filteredSpokeVpcRoutes")
    def filtered_spoke_vpc_routes(self) -> str:
        return pulumi.get(self, "filtered_spoke_vpc_routes")

    @property
    @pulumi.getter(name="gwName")
    def gw_name(self) -> str:
        return pulumi.get(self, "gw_name")

    @property
    @pulumi.getter(name="gwSize")
    def gw_size(self) -> str:
        return pulumi.get(self, "gw_size")

    @property
    @pulumi.getter(name="haAvailabilityDomain")
    def ha_availability_domain(self) -> str:
        return pulumi.get(self, "ha_availability_domain")

    @property
    @pulumi.getter(name="haAzureEipNameResourceGroup")
    def ha_azure_eip_name_resource_group(self) -> str:
        return pulumi.get(self, "ha_azure_eip_name_resource_group")

    @property
    @pulumi.getter(name="haCloudInstanceId")
    def ha_cloud_instance_id(self) -> str:
        return pulumi.get(self, "ha_cloud_instance_id")

    @property
    @pulumi.getter(name="haEip")
    def ha_eip(self) -> str:
        return pulumi.get(self, "ha_eip")

    @property
    @pulumi.getter(name="haFaultDomain")
    def ha_fault_domain(self) -> str:
        return pulumi.get(self, "ha_fault_domain")

    @property
    @pulumi.getter(name="haGwName")
    def ha_gw_name(self) -> str:
        return pulumi.get(self, "ha_gw_name")

    @property
    @pulumi.getter(name="haGwSize")
    def ha_gw_size(self) -> str:
        return pulumi.get(self, "ha_gw_size")

    @property
    @pulumi.getter(name="haImageVersion")
    def ha_image_version(self) -> str:
        return pulumi.get(self, "ha_image_version")

    @property
    @pulumi.getter(name="haInsaneModeAz")
    def ha_insane_mode_az(self) -> str:
        return pulumi.get(self, "ha_insane_mode_az")

    @property
    @pulumi.getter(name="haOobAvailabilityZone")
    def ha_oob_availability_zone(self) -> str:
        return pulumi.get(self, "ha_oob_availability_zone")

    @property
    @pulumi.getter(name="haOobManagementSubnet")
    def ha_oob_management_subnet(self) -> str:
        return pulumi.get(self, "ha_oob_management_subnet")

    @property
    @pulumi.getter(name="haPrivateIp")
    def ha_private_ip(self) -> str:
        return pulumi.get(self, "ha_private_ip")

    @property
    @pulumi.getter(name="haPublicIp")
    def ha_public_ip(self) -> str:
        return pulumi.get(self, "ha_public_ip")

    @property
    @pulumi.getter(name="haSecurityGroupId")
    def ha_security_group_id(self) -> str:
        return pulumi.get(self, "ha_security_group_id")

    @property
    @pulumi.getter(name="haSoftwareVersion")
    def ha_software_version(self) -> str:
        return pulumi.get(self, "ha_software_version")

    @property
    @pulumi.getter(name="haSubnet")
    def ha_subnet(self) -> str:
        return pulumi.get(self, "ha_subnet")

    @property
    @pulumi.getter(name="haZone")
    def ha_zone(self) -> str:
        return pulumi.get(self, "ha_zone")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageVersion")
    def image_version(self) -> str:
        return pulumi.get(self, "image_version")

    @property
    @pulumi.getter(name="includedAdvertisedSpokeRoutes")
    def included_advertised_spoke_routes(self) -> str:
        return pulumi.get(self, "included_advertised_spoke_routes")

    @property
    @pulumi.getter(name="insaneMode")
    def insane_mode(self) -> bool:
        return pulumi.get(self, "insane_mode")

    @property
    @pulumi.getter(name="insaneModeAz")
    def insane_mode_az(self) -> str:
        return pulumi.get(self, "insane_mode_az")

    @property
    @pulumi.getter(name="learnedCidrsApprovalMode")
    def learned_cidrs_approval_mode(self) -> str:
        return pulumi.get(self, "learned_cidrs_approval_mode")

    @property
    @pulumi.getter(name="localAsNumber")
    def local_as_number(self) -> str:
        return pulumi.get(self, "local_as_number")

    @property
    @pulumi.getter(name="monitorExcludeLists")
    def monitor_exclude_lists(self) -> Sequence[str]:
        return pulumi.get(self, "monitor_exclude_lists")

    @property
    @pulumi.getter(name="oobAvailabilityZone")
    def oob_availability_zone(self) -> str:
        return pulumi.get(self, "oob_availability_zone")

    @property
    @pulumi.getter(name="oobManagementSubnet")
    def oob_management_subnet(self) -> str:
        return pulumi.get(self, "oob_management_subnet")

    @property
    @pulumi.getter(name="prependAsPaths")
    def prepend_as_paths(self) -> Sequence[str]:
        return pulumi.get(self, "prepend_as_paths")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="singleAzHa")
    def single_az_ha(self) -> bool:
        return pulumi.get(self, "single_az_ha")

    @property
    @pulumi.getter(name="singleIpSnat")
    def single_ip_snat(self) -> bool:
        return pulumi.get(self, "single_ip_snat")

    @property
    @pulumi.getter(name="softwareVersion")
    def software_version(self) -> str:
        return pulumi.get(self, "software_version")

    @property
    @pulumi.getter(name="spokeBgpManualAdvertiseCidrs")
    def spoke_bgp_manual_advertise_cidrs(self) -> Sequence[str]:
        return pulumi.get(self, "spoke_bgp_manual_advertise_cidrs")

    @property
    @pulumi.getter(name="spotPrice")
    def spot_price(self) -> str:
        return pulumi.get(self, "spot_price")

    @property
    @pulumi.getter
    def subnet(self) -> str:
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="tagLists")
    def tag_lists(self) -> Sequence[str]:
        return pulumi.get(self, "tag_lists")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitGw")
    def transit_gw(self) -> str:
        return pulumi.get(self, "transit_gw")

    @property
    @pulumi.getter(name="tunnelDetectionTime")
    def tunnel_detection_time(self) -> int:
        return pulumi.get(self, "tunnel_detection_time")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcReg")
    def vpc_reg(self) -> str:
        return pulumi.get(self, "vpc_reg")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetAviatrixSpokeGatewayResult(GetAviatrixSpokeGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAviatrixSpokeGatewayResult(
            account_name=self.account_name,
            allocate_new_eip=self.allocate_new_eip,
            approved_learned_cidrs=self.approved_learned_cidrs,
            availability_domain=self.availability_domain,
            azure_eip_name_resource_group=self.azure_eip_name_resource_group,
            bgp_ecmp=self.bgp_ecmp,
            bgp_hold_time=self.bgp_hold_time,
            bgp_polling_time=self.bgp_polling_time,
            cloud_instance_id=self.cloud_instance_id,
            cloud_type=self.cloud_type,
            customized_spoke_vpc_routes=self.customized_spoke_vpc_routes,
            disable_route_propagation=self.disable_route_propagation,
            eip=self.eip,
            enable_active_standby=self.enable_active_standby,
            enable_active_standby_preemptive=self.enable_active_standby_preemptive,
            enable_auto_advertise_s2c_cidrs=self.enable_auto_advertise_s2c_cidrs,
            enable_bgp=self.enable_bgp,
            enable_encrypt_volume=self.enable_encrypt_volume,
            enable_jumbo_frame=self.enable_jumbo_frame,
            enable_learned_cidrs_approval=self.enable_learned_cidrs_approval,
            enable_monitor_gateway_subnets=self.enable_monitor_gateway_subnets,
            enable_private_oob=self.enable_private_oob,
            enable_private_vpc_default_route=self.enable_private_vpc_default_route,
            enable_skip_public_route_table_update=self.enable_skip_public_route_table_update,
            enable_spot_instance=self.enable_spot_instance,
            enable_vpc_dns_server=self.enable_vpc_dns_server,
            fault_domain=self.fault_domain,
            filtered_spoke_vpc_routes=self.filtered_spoke_vpc_routes,
            gw_name=self.gw_name,
            gw_size=self.gw_size,
            ha_availability_domain=self.ha_availability_domain,
            ha_azure_eip_name_resource_group=self.ha_azure_eip_name_resource_group,
            ha_cloud_instance_id=self.ha_cloud_instance_id,
            ha_eip=self.ha_eip,
            ha_fault_domain=self.ha_fault_domain,
            ha_gw_name=self.ha_gw_name,
            ha_gw_size=self.ha_gw_size,
            ha_image_version=self.ha_image_version,
            ha_insane_mode_az=self.ha_insane_mode_az,
            ha_oob_availability_zone=self.ha_oob_availability_zone,
            ha_oob_management_subnet=self.ha_oob_management_subnet,
            ha_private_ip=self.ha_private_ip,
            ha_public_ip=self.ha_public_ip,
            ha_security_group_id=self.ha_security_group_id,
            ha_software_version=self.ha_software_version,
            ha_subnet=self.ha_subnet,
            ha_zone=self.ha_zone,
            id=self.id,
            image_version=self.image_version,
            included_advertised_spoke_routes=self.included_advertised_spoke_routes,
            insane_mode=self.insane_mode,
            insane_mode_az=self.insane_mode_az,
            learned_cidrs_approval_mode=self.learned_cidrs_approval_mode,
            local_as_number=self.local_as_number,
            monitor_exclude_lists=self.monitor_exclude_lists,
            oob_availability_zone=self.oob_availability_zone,
            oob_management_subnet=self.oob_management_subnet,
            prepend_as_paths=self.prepend_as_paths,
            private_ip=self.private_ip,
            public_ip=self.public_ip,
            security_group_id=self.security_group_id,
            single_az_ha=self.single_az_ha,
            single_ip_snat=self.single_ip_snat,
            software_version=self.software_version,
            spoke_bgp_manual_advertise_cidrs=self.spoke_bgp_manual_advertise_cidrs,
            spot_price=self.spot_price,
            subnet=self.subnet,
            tag_lists=self.tag_lists,
            tags=self.tags,
            transit_gw=self.transit_gw,
            tunnel_detection_time=self.tunnel_detection_time,
            vpc_id=self.vpc_id,
            vpc_reg=self.vpc_reg,
            zone=self.zone)


def get_aviatrix_spoke_gateway(gw_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAviatrixSpokeGatewayResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['gwName'] = gw_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aviatrix:index/getAviatrixSpokeGateway:getAviatrixSpokeGateway', __args__, opts=opts, typ=GetAviatrixSpokeGatewayResult).value

    return AwaitableGetAviatrixSpokeGatewayResult(
        account_name=__ret__.account_name,
        allocate_new_eip=__ret__.allocate_new_eip,
        approved_learned_cidrs=__ret__.approved_learned_cidrs,
        availability_domain=__ret__.availability_domain,
        azure_eip_name_resource_group=__ret__.azure_eip_name_resource_group,
        bgp_ecmp=__ret__.bgp_ecmp,
        bgp_hold_time=__ret__.bgp_hold_time,
        bgp_polling_time=__ret__.bgp_polling_time,
        cloud_instance_id=__ret__.cloud_instance_id,
        cloud_type=__ret__.cloud_type,
        customized_spoke_vpc_routes=__ret__.customized_spoke_vpc_routes,
        disable_route_propagation=__ret__.disable_route_propagation,
        eip=__ret__.eip,
        enable_active_standby=__ret__.enable_active_standby,
        enable_active_standby_preemptive=__ret__.enable_active_standby_preemptive,
        enable_auto_advertise_s2c_cidrs=__ret__.enable_auto_advertise_s2c_cidrs,
        enable_bgp=__ret__.enable_bgp,
        enable_encrypt_volume=__ret__.enable_encrypt_volume,
        enable_jumbo_frame=__ret__.enable_jumbo_frame,
        enable_learned_cidrs_approval=__ret__.enable_learned_cidrs_approval,
        enable_monitor_gateway_subnets=__ret__.enable_monitor_gateway_subnets,
        enable_private_oob=__ret__.enable_private_oob,
        enable_private_vpc_default_route=__ret__.enable_private_vpc_default_route,
        enable_skip_public_route_table_update=__ret__.enable_skip_public_route_table_update,
        enable_spot_instance=__ret__.enable_spot_instance,
        enable_vpc_dns_server=__ret__.enable_vpc_dns_server,
        fault_domain=__ret__.fault_domain,
        filtered_spoke_vpc_routes=__ret__.filtered_spoke_vpc_routes,
        gw_name=__ret__.gw_name,
        gw_size=__ret__.gw_size,
        ha_availability_domain=__ret__.ha_availability_domain,
        ha_azure_eip_name_resource_group=__ret__.ha_azure_eip_name_resource_group,
        ha_cloud_instance_id=__ret__.ha_cloud_instance_id,
        ha_eip=__ret__.ha_eip,
        ha_fault_domain=__ret__.ha_fault_domain,
        ha_gw_name=__ret__.ha_gw_name,
        ha_gw_size=__ret__.ha_gw_size,
        ha_image_version=__ret__.ha_image_version,
        ha_insane_mode_az=__ret__.ha_insane_mode_az,
        ha_oob_availability_zone=__ret__.ha_oob_availability_zone,
        ha_oob_management_subnet=__ret__.ha_oob_management_subnet,
        ha_private_ip=__ret__.ha_private_ip,
        ha_public_ip=__ret__.ha_public_ip,
        ha_security_group_id=__ret__.ha_security_group_id,
        ha_software_version=__ret__.ha_software_version,
        ha_subnet=__ret__.ha_subnet,
        ha_zone=__ret__.ha_zone,
        id=__ret__.id,
        image_version=__ret__.image_version,
        included_advertised_spoke_routes=__ret__.included_advertised_spoke_routes,
        insane_mode=__ret__.insane_mode,
        insane_mode_az=__ret__.insane_mode_az,
        learned_cidrs_approval_mode=__ret__.learned_cidrs_approval_mode,
        local_as_number=__ret__.local_as_number,
        monitor_exclude_lists=__ret__.monitor_exclude_lists,
        oob_availability_zone=__ret__.oob_availability_zone,
        oob_management_subnet=__ret__.oob_management_subnet,
        prepend_as_paths=__ret__.prepend_as_paths,
        private_ip=__ret__.private_ip,
        public_ip=__ret__.public_ip,
        security_group_id=__ret__.security_group_id,
        single_az_ha=__ret__.single_az_ha,
        single_ip_snat=__ret__.single_ip_snat,
        software_version=__ret__.software_version,
        spoke_bgp_manual_advertise_cidrs=__ret__.spoke_bgp_manual_advertise_cidrs,
        spot_price=__ret__.spot_price,
        subnet=__ret__.subnet,
        tag_lists=__ret__.tag_lists,
        tags=__ret__.tags,
        transit_gw=__ret__.transit_gw,
        tunnel_detection_time=__ret__.tunnel_detection_time,
        vpc_id=__ret__.vpc_id,
        vpc_reg=__ret__.vpc_reg,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_aviatrix_spoke_gateway)
def get_aviatrix_spoke_gateway_output(gw_name: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAviatrixSpokeGatewayResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...