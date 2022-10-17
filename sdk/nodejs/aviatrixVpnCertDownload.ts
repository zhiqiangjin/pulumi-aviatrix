// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixVpnCertDownload extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixVpnCertDownload resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixVpnCertDownloadState, opts?: pulumi.CustomResourceOptions): AviatrixVpnCertDownload {
        return new AviatrixVpnCertDownload(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixVpnCertDownload:AviatrixVpnCertDownload';

    /**
     * Returns true if the given object is an instance of AviatrixVpnCertDownload.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixVpnCertDownload {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixVpnCertDownload.__pulumiType;
    }

    /**
     * Whether the VPN Certificate download is enabled. Supported Values: "true", "false"
     */
    public readonly downloadEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is
     * supported. Example: ["saml_endpoint_1"].
     */
    public readonly samlEndpoints!: pulumi.Output<string[] | undefined>;

    /**
     * Create a AviatrixVpnCertDownload resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AviatrixVpnCertDownloadArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixVpnCertDownloadArgs | AviatrixVpnCertDownloadState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixVpnCertDownloadState | undefined;
            resourceInputs["downloadEnabled"] = state ? state.downloadEnabled : undefined;
            resourceInputs["samlEndpoints"] = state ? state.samlEndpoints : undefined;
        } else {
            const args = argsOrState as AviatrixVpnCertDownloadArgs | undefined;
            resourceInputs["downloadEnabled"] = args ? args.downloadEnabled : undefined;
            resourceInputs["samlEndpoints"] = args ? args.samlEndpoints : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixVpnCertDownload.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixVpnCertDownload resources.
 */
export interface AviatrixVpnCertDownloadState {
    /**
     * Whether the VPN Certificate download is enabled. Supported Values: "true", "false"
     */
    downloadEnabled?: pulumi.Input<boolean>;
    /**
     * List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is
     * supported. Example: ["saml_endpoint_1"].
     */
    samlEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AviatrixVpnCertDownload resource.
 */
export interface AviatrixVpnCertDownloadArgs {
    /**
     * Whether the VPN Certificate download is enabled. Supported Values: "true", "false"
     */
    downloadEnabled?: pulumi.Input<boolean>;
    /**
     * List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is
     * supported. Example: ["saml_endpoint_1"].
     */
    samlEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
}
