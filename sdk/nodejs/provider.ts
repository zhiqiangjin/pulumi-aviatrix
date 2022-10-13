// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The provider type for the aviatrix package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'aviatrix';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    public readonly controllerIp!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string>;
    public readonly pathToCaCertificate!: pulumi.Output<string | undefined>;
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.controllerIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'controllerIp'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["controllerIp"] = args ? args.controllerIp : undefined;
            resourceInputs["ignoreTags"] = pulumi.output(args ? args.ignoreTags : undefined).apply(JSON.stringify);
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["pathToCaCertificate"] = args ? args.pathToCaCertificate : undefined;
            resourceInputs["skipVersionValidation"] = pulumi.output(args ? args.skipVersionValidation : undefined).apply(JSON.stringify);
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["verifySslCertificate"] = pulumi.output(args ? args.verifySslCertificate : undefined).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    controllerIp: pulumi.Input<string>;
    /**
     * Configuration block with settings to ignore tags across all resources.
     */
    ignoreTags?: pulumi.Input<inputs.ProviderIgnoreTags>;
    password: pulumi.Input<string>;
    pathToCaCertificate?: pulumi.Input<string>;
    skipVersionValidation?: pulumi.Input<boolean>;
    username: pulumi.Input<string>;
    verifySslCertificate?: pulumi.Input<boolean>;
}
