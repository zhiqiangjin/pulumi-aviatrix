// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAviatrixAccount(ctx *pulumi.Context, args *LookupAviatrixAccountArgs, opts ...pulumi.InvokeOption) (*LookupAviatrixAccountResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAviatrixAccountResult
	err := ctx.Invoke("aviatrix:index/getAviatrixAccount:getAviatrixAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAviatrixAccount.
type LookupAviatrixAccountArgs struct {
	AccountName string `pulumi:"accountName"`
}

// A collection of values returned by getAviatrixAccount.
type LookupAviatrixAccountResult struct {
	AccountName              string `pulumi:"accountName"`
	AlicloudAccountId        string `pulumi:"alicloudAccountId"`
	ArmSubscriptionId        string `pulumi:"armSubscriptionId"`
	AwsAccountNumber         string `pulumi:"awsAccountNumber"`
	AwsCaCertPath            string `pulumi:"awsCaCertPath"`
	AwsGatewayRoleApp        string `pulumi:"awsGatewayRoleApp"`
	AwsGatewayRoleEc2        string `pulumi:"awsGatewayRoleEc2"`
	AwsRoleArn               string `pulumi:"awsRoleArn"`
	AwsRoleEc2               string `pulumi:"awsRoleEc2"`
	AwschinaAccountNumber    string `pulumi:"awschinaAccountNumber"`
	AwschinaIam              bool   `pulumi:"awschinaIam"`
	AwschinaRoleApp          string `pulumi:"awschinaRoleApp"`
	AwschinaRoleEc2          string `pulumi:"awschinaRoleEc2"`
	AwsgovAccountNumber      string `pulumi:"awsgovAccountNumber"`
	AwsgovIam                bool   `pulumi:"awsgovIam"`
	AwsgovRoleApp            string `pulumi:"awsgovRoleApp"`
	AwsgovRoleEc2            string `pulumi:"awsgovRoleEc2"`
	AwssAccountNumber        string `pulumi:"awssAccountNumber"`
	AwssCapAccountName       string `pulumi:"awssCapAccountName"`
	AwssCapAgency            string `pulumi:"awssCapAgency"`
	AwssCapCertKeyPath       string `pulumi:"awssCapCertKeyPath"`
	AwssCapCertPath          string `pulumi:"awssCapCertPath"`
	AwssCapRoleName          string `pulumi:"awssCapRoleName"`
	AwssCapUrl               string `pulumi:"awssCapUrl"`
	AwstsAccountNumber       string `pulumi:"awstsAccountNumber"`
	AwstsCapAgency           string `pulumi:"awstsCapAgency"`
	AwstsCapCertKeyPath      string `pulumi:"awstsCapCertKeyPath"`
	AwstsCapCertPath         string `pulumi:"awstsCapCertPath"`
	AwstsCapMission          string `pulumi:"awstsCapMission"`
	AwstsCapRoleName         string `pulumi:"awstsCapRoleName"`
	AwstsCapUrl              string `pulumi:"awstsCapUrl"`
	AzurechinaSubscriptionId string `pulumi:"azurechinaSubscriptionId"`
	AzuregovSubscriptionId   string `pulumi:"azuregovSubscriptionId"`
	CloudType                int    `pulumi:"cloudType"`
	GcloudProjectId          string `pulumi:"gcloudProjectId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func LookupAviatrixAccountOutput(ctx *pulumi.Context, args LookupAviatrixAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAviatrixAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAviatrixAccountResult, error) {
			args := v.(LookupAviatrixAccountArgs)
			r, err := LookupAviatrixAccount(ctx, &args, opts...)
			var s LookupAviatrixAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAviatrixAccountResultOutput)
}

// A collection of arguments for invoking getAviatrixAccount.
type LookupAviatrixAccountOutputArgs struct {
	AccountName pulumi.StringInput `pulumi:"accountName"`
}

func (LookupAviatrixAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAviatrixAccountArgs)(nil)).Elem()
}

// A collection of values returned by getAviatrixAccount.
type LookupAviatrixAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAviatrixAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAviatrixAccountResult)(nil)).Elem()
}

func (o LookupAviatrixAccountResultOutput) ToLookupAviatrixAccountResultOutput() LookupAviatrixAccountResultOutput {
	return o
}

func (o LookupAviatrixAccountResultOutput) ToLookupAviatrixAccountResultOutputWithContext(ctx context.Context) LookupAviatrixAccountResultOutput {
	return o
}

func (o LookupAviatrixAccountResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AlicloudAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AlicloudAccountId }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) ArmSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.ArmSubscriptionId }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsAccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsAccountNumber }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsCaCertPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsCaCertPath }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsGatewayRoleApp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsGatewayRoleApp }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsGatewayRoleEc2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsGatewayRoleEc2 }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsRoleArn }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsRoleEc2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsRoleEc2 }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwschinaAccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwschinaAccountNumber }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwschinaIam() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) bool { return v.AwschinaIam }).(pulumi.BoolOutput)
}

func (o LookupAviatrixAccountResultOutput) AwschinaRoleApp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwschinaRoleApp }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwschinaRoleEc2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwschinaRoleEc2 }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsgovAccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsgovAccountNumber }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsgovIam() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) bool { return v.AwsgovIam }).(pulumi.BoolOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsgovRoleApp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsgovRoleApp }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwsgovRoleEc2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwsgovRoleEc2 }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwssAccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwssAccountNumber }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwssCapAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwssCapAccountName }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwssCapAgency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwssCapAgency }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwssCapCertKeyPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwssCapCertKeyPath }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwssCapCertPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwssCapCertPath }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwssCapRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwssCapRoleName }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwssCapUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwssCapUrl }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwstsAccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwstsAccountNumber }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwstsCapAgency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwstsCapAgency }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwstsCapCertKeyPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwstsCapCertKeyPath }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwstsCapCertPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwstsCapCertPath }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwstsCapMission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwstsCapMission }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwstsCapRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwstsCapRoleName }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AwstsCapUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AwstsCapUrl }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AzurechinaSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AzurechinaSubscriptionId }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) AzuregovSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.AzuregovSubscriptionId }).(pulumi.StringOutput)
}

func (o LookupAviatrixAccountResultOutput) CloudType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) int { return v.CloudType }).(pulumi.IntOutput)
}

func (o LookupAviatrixAccountResultOutput) GcloudProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.GcloudProjectId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAviatrixAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAviatrixAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAviatrixAccountResultOutput{})
}
