// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixAccountUser extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixAccountUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixAccountUserState, opts?: pulumi.CustomResourceOptions): AviatrixAccountUser {
        return new AviatrixAccountUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixAccountUser:AviatrixAccountUser';

    /**
     * Returns true if the given object is an instance of AviatrixAccountUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixAccountUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixAccountUser.__pulumiType;
    }

    /**
     * Email of address of account user to be created.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Login password for the account user to be created.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or
     * underscores. 1 to 80 in length. No spaces are allowed.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a AviatrixAccountUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixAccountUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixAccountUserArgs | AviatrixAccountUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixAccountUserState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as AviatrixAccountUserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixAccountUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixAccountUser resources.
 */
export interface AviatrixAccountUserState {
    /**
     * Email of address of account user to be created.
     */
    email?: pulumi.Input<string>;
    /**
     * Login password for the account user to be created.
     */
    password?: pulumi.Input<string>;
    /**
     * Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or
     * underscores. 1 to 80 in length. No spaces are allowed.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AviatrixAccountUser resource.
 */
export interface AviatrixAccountUserArgs {
    /**
     * Email of address of account user to be created.
     */
    email: pulumi.Input<string>;
    /**
     * Login password for the account user to be created.
     */
    password: pulumi.Input<string>;
    /**
     * Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or
     * underscores. 1 to 80 in length. No spaces are allowed.
     */
    username: pulumi.Input<string>;
}