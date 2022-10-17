// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixRbacGroupAccessAccountAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixRbacGroupAccessAccountAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixRbacGroupAccessAccountAttachmentState, opts?: pulumi.CustomResourceOptions): AviatrixRbacGroupAccessAccountAttachment {
        return new AviatrixRbacGroupAccessAccountAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixRbacGroupAccessAccountAttachment:AviatrixRbacGroupAccessAccountAttachment';

    /**
     * Returns true if the given object is an instance of AviatrixRbacGroupAccessAccountAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixRbacGroupAccessAccountAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixRbacGroupAccessAccountAttachment.__pulumiType;
    }

    /**
     * Access account name.
     */
    public readonly accessAccountName!: pulumi.Output<string>;
    /**
     * RBAC permission group name.
     */
    public readonly groupName!: pulumi.Output<string>;

    /**
     * Create a AviatrixRbacGroupAccessAccountAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixRbacGroupAccessAccountAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixRbacGroupAccessAccountAttachmentArgs | AviatrixRbacGroupAccessAccountAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixRbacGroupAccessAccountAttachmentState | undefined;
            resourceInputs["accessAccountName"] = state ? state.accessAccountName : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
        } else {
            const args = argsOrState as AviatrixRbacGroupAccessAccountAttachmentArgs | undefined;
            if ((!args || args.accessAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessAccountName'");
            }
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            resourceInputs["accessAccountName"] = args ? args.accessAccountName : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixRbacGroupAccessAccountAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixRbacGroupAccessAccountAttachment resources.
 */
export interface AviatrixRbacGroupAccessAccountAttachmentState {
    /**
     * Access account name.
     */
    accessAccountName?: pulumi.Input<string>;
    /**
     * RBAC permission group name.
     */
    groupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixRbacGroupAccessAccountAttachment resource.
 */
export interface AviatrixRbacGroupAccessAccountAttachmentArgs {
    /**
     * Access account name.
     */
    accessAccountName: pulumi.Input<string>;
    /**
     * RBAC permission group name.
     */
    groupName: pulumi.Input<string>;
}
