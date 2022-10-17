// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AviatrixSamlEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing AviatrixSamlEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AviatrixSamlEndpointState, opts?: pulumi.CustomResourceOptions): AviatrixSamlEndpoint {
        return new AviatrixSamlEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aviatrix:index/aviatrixSamlEndpoint:AviatrixSamlEndpoint';

    /**
     * Returns true if the given object is an instance of AviatrixSamlEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AviatrixSamlEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AviatrixSamlEndpoint.__pulumiType;
    }

    /**
     * Access type.
     */
    public readonly accessSetBy!: pulumi.Output<string | undefined>;
    /**
     * Switch to differentiate if it is for controller login.
     */
    public readonly controllerLogin!: pulumi.Output<boolean | undefined>;
    /**
     * Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
     */
    public readonly customEntityId!: pulumi.Output<string | undefined>;
    /**
     * Custom SAML Request Template.
     */
    public readonly customSamlRequestTemplate!: pulumi.Output<string | undefined>;
    /**
     * SAML Endpoint Name.
     */
    public readonly endpointName!: pulumi.Output<string>;
    /**
     * IDP Metadata.
     */
    public readonly idpMetadata!: pulumi.Output<string | undefined>;
    /**
     * Type of IDP Metadata.
     */
    public readonly idpMetadataType!: pulumi.Output<string>;
    /**
     * IDP Metadata.
     */
    public readonly idpMetadataUrl!: pulumi.Output<string | undefined>;
    /**
     * List of RBAC groups.
     */
    public readonly rbacGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Whether to sign SAML AuthnRequests
     */
    public readonly signAuthnRequests!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AviatrixSamlEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AviatrixSamlEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AviatrixSamlEndpointArgs | AviatrixSamlEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AviatrixSamlEndpointState | undefined;
            resourceInputs["accessSetBy"] = state ? state.accessSetBy : undefined;
            resourceInputs["controllerLogin"] = state ? state.controllerLogin : undefined;
            resourceInputs["customEntityId"] = state ? state.customEntityId : undefined;
            resourceInputs["customSamlRequestTemplate"] = state ? state.customSamlRequestTemplate : undefined;
            resourceInputs["endpointName"] = state ? state.endpointName : undefined;
            resourceInputs["idpMetadata"] = state ? state.idpMetadata : undefined;
            resourceInputs["idpMetadataType"] = state ? state.idpMetadataType : undefined;
            resourceInputs["idpMetadataUrl"] = state ? state.idpMetadataUrl : undefined;
            resourceInputs["rbacGroups"] = state ? state.rbacGroups : undefined;
            resourceInputs["signAuthnRequests"] = state ? state.signAuthnRequests : undefined;
        } else {
            const args = argsOrState as AviatrixSamlEndpointArgs | undefined;
            if ((!args || args.endpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointName'");
            }
            if ((!args || args.idpMetadataType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idpMetadataType'");
            }
            resourceInputs["accessSetBy"] = args ? args.accessSetBy : undefined;
            resourceInputs["controllerLogin"] = args ? args.controllerLogin : undefined;
            resourceInputs["customEntityId"] = args ? args.customEntityId : undefined;
            resourceInputs["customSamlRequestTemplate"] = args ? args.customSamlRequestTemplate : undefined;
            resourceInputs["endpointName"] = args ? args.endpointName : undefined;
            resourceInputs["idpMetadata"] = args ? args.idpMetadata : undefined;
            resourceInputs["idpMetadataType"] = args ? args.idpMetadataType : undefined;
            resourceInputs["idpMetadataUrl"] = args ? args.idpMetadataUrl : undefined;
            resourceInputs["rbacGroups"] = args ? args.rbacGroups : undefined;
            resourceInputs["signAuthnRequests"] = args ? args.signAuthnRequests : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AviatrixSamlEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AviatrixSamlEndpoint resources.
 */
export interface AviatrixSamlEndpointState {
    /**
     * Access type.
     */
    accessSetBy?: pulumi.Input<string>;
    /**
     * Switch to differentiate if it is for controller login.
     */
    controllerLogin?: pulumi.Input<boolean>;
    /**
     * Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
     */
    customEntityId?: pulumi.Input<string>;
    /**
     * Custom SAML Request Template.
     */
    customSamlRequestTemplate?: pulumi.Input<string>;
    /**
     * SAML Endpoint Name.
     */
    endpointName?: pulumi.Input<string>;
    /**
     * IDP Metadata.
     */
    idpMetadata?: pulumi.Input<string>;
    /**
     * Type of IDP Metadata.
     */
    idpMetadataType?: pulumi.Input<string>;
    /**
     * IDP Metadata.
     */
    idpMetadataUrl?: pulumi.Input<string>;
    /**
     * List of RBAC groups.
     */
    rbacGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to sign SAML AuthnRequests
     */
    signAuthnRequests?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AviatrixSamlEndpoint resource.
 */
export interface AviatrixSamlEndpointArgs {
    /**
     * Access type.
     */
    accessSetBy?: pulumi.Input<string>;
    /**
     * Switch to differentiate if it is for controller login.
     */
    controllerLogin?: pulumi.Input<boolean>;
    /**
     * Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname'.
     */
    customEntityId?: pulumi.Input<string>;
    /**
     * Custom SAML Request Template.
     */
    customSamlRequestTemplate?: pulumi.Input<string>;
    /**
     * SAML Endpoint Name.
     */
    endpointName: pulumi.Input<string>;
    /**
     * IDP Metadata.
     */
    idpMetadata?: pulumi.Input<string>;
    /**
     * Type of IDP Metadata.
     */
    idpMetadataType: pulumi.Input<string>;
    /**
     * IDP Metadata.
     */
    idpMetadataUrl?: pulumi.Input<string>;
    /**
     * List of RBAC groups.
     */
    rbacGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to sign SAML AuthnRequests
     */
    signAuthnRequests?: pulumi.Input<boolean>;
}
