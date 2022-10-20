// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixRbacGroupUserAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixRbacGroupUserAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixRbacGroupUserAttachmentState, opts?: pulumi.CustomResourceOptions): AviatrixRbacGroupUserAttachment {
        return new AviatrixRbacGroupUserAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixRbacGroupUserAttachment:AviatrixRbacGroupUserAttachment';

    /**
     * Returns true if the given object is an instance of AviatrixRbacGroupUserAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixRbacGroupUserAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixRbacGroupUserAttachment.__pulumiType;
    }

    /**
     * RBAC permission group name.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * Account user name.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a AviatrixRbacGroupUserAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixRbacGroupUserAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixRbacGroupUserAttachmentArgs | AviatrixRbacGroupUserAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixRbacGroupUserAttachmentState | undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as AviatrixRbacGroupUserAttachmentArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixRbacGroupUserAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixRbacGroupUserAttachment resources.
 */
export interface AviatrixRbacGroupUserAttachmentState {
    /**
     * RBAC permission group name.
     */
    groupName?: pulumi.Input<string>;
    /**
     * Account user name.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixRbacGroupUserAttachment resource.
 */
export interface AviatrixRbacGroupUserAttachmentArgs {
    /**
     * RBAC permission group name.
     */
    groupName: pulumi.Input<string>;
    /**
     * Account user name.
     */
    userName: pulumi.Input<string>;
}