// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAviatrixGatewayImage(args: GetAviatrixGatewayImageArgs, opts?: pulumi.InvokeOptions): Promise<GetAviatrixGatewayImageResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aviatrix:index/getAviatrixGatewayImage:getAviatrixGatewayImage", {
        "cloudType": args.cloudType,
        "softwareVersion": args.softwareVersion,
    }, opts);
}

/**
 * A collection of arguments for invoking getAviatrixGatewayImage.
 */
export interface GetAviatrixGatewayImageArgs {
    cloudType: number;
    softwareVersion: string;
}

/**
 * A collection of values returned by getAviatrixGatewayImage.
 */
export interface GetAviatrixGatewayImageResult {
    readonly cloudType: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageVersion: string;
    readonly softwareVersion: string;
}

export function getAviatrixGatewayImageOutput(args: GetAviatrixGatewayImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAviatrixGatewayImageResult> {
    return pulumi.output(args).apply(a => getAviatrixGatewayImage(a, opts))
}

/**
 * A collection of arguments for invoking getAviatrixGatewayImage.
 */
export interface GetAviatrixGatewayImageOutputArgs {
    cloudType: pulumi.Input<number>;
    softwareVersion: pulumi.Input<string>;
}
