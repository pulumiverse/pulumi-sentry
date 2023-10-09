// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-sentry/sdk/go/sentry/internal"
)

// Sentry Organization Integration data source.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-sentry/sdk/go/sentry"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sentry.GetSentryOrganizationIntegration(ctx, &sentry.GetSentryOrganizationIntegrationArgs{
//				Name:         "my-github-organization",
//				Organization: "my-organization",
//				ProviderKey:  "github",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = sentry.GetSentryOrganizationIntegration(ctx, &sentry.GetSentryOrganizationIntegrationArgs{
//				Name:         "Slack Workspace",
//				Organization: "my-organization",
//				ProviderKey:  "slack",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSentryOrganizationIntegration(ctx *pulumi.Context, args *GetSentryOrganizationIntegrationArgs, opts ...pulumi.InvokeOption) (*GetSentryOrganizationIntegrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSentryOrganizationIntegrationResult
	err := ctx.Invoke("sentry:index/getSentryOrganizationIntegration:getSentryOrganizationIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSentryOrganizationIntegration.
type GetSentryOrganizationIntegrationArgs struct {
	// The name of the organization integration.
	Name string `pulumi:"name"`
	// The slug of the organization the integration belongs to.
	Organization string `pulumi:"organization"`
	// The key of the organization integration provider.
	ProviderKey string `pulumi:"providerKey"`
}

// A collection of values returned by getSentryOrganizationIntegration.
type GetSentryOrganizationIntegrationResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The internal ID for this organization integration.
	InternalId string `pulumi:"internalId"`
	// The name of the organization integration.
	Name string `pulumi:"name"`
	// The slug of the organization the integration belongs to.
	Organization string `pulumi:"organization"`
	// The key of the organization integration provider.
	ProviderKey string `pulumi:"providerKey"`
}

func GetSentryOrganizationIntegrationOutput(ctx *pulumi.Context, args GetSentryOrganizationIntegrationOutputArgs, opts ...pulumi.InvokeOption) GetSentryOrganizationIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSentryOrganizationIntegrationResult, error) {
			args := v.(GetSentryOrganizationIntegrationArgs)
			r, err := GetSentryOrganizationIntegration(ctx, &args, opts...)
			var s GetSentryOrganizationIntegrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSentryOrganizationIntegrationResultOutput)
}

// A collection of arguments for invoking getSentryOrganizationIntegration.
type GetSentryOrganizationIntegrationOutputArgs struct {
	// The name of the organization integration.
	Name pulumi.StringInput `pulumi:"name"`
	// The slug of the organization the integration belongs to.
	Organization pulumi.StringInput `pulumi:"organization"`
	// The key of the organization integration provider.
	ProviderKey pulumi.StringInput `pulumi:"providerKey"`
}

func (GetSentryOrganizationIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSentryOrganizationIntegrationArgs)(nil)).Elem()
}

// A collection of values returned by getSentryOrganizationIntegration.
type GetSentryOrganizationIntegrationResultOutput struct{ *pulumi.OutputState }

func (GetSentryOrganizationIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSentryOrganizationIntegrationResult)(nil)).Elem()
}

func (o GetSentryOrganizationIntegrationResultOutput) ToGetSentryOrganizationIntegrationResultOutput() GetSentryOrganizationIntegrationResultOutput {
	return o
}

func (o GetSentryOrganizationIntegrationResultOutput) ToGetSentryOrganizationIntegrationResultOutputWithContext(ctx context.Context) GetSentryOrganizationIntegrationResultOutput {
	return o
}

func (o GetSentryOrganizationIntegrationResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSentryOrganizationIntegrationResult] {
	return pulumix.Output[GetSentryOrganizationIntegrationResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetSentryOrganizationIntegrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSentryOrganizationIntegrationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The internal ID for this organization integration.
func (o GetSentryOrganizationIntegrationResultOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSentryOrganizationIntegrationResult) string { return v.InternalId }).(pulumi.StringOutput)
}

// The name of the organization integration.
func (o GetSentryOrganizationIntegrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSentryOrganizationIntegrationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The slug of the organization the integration belongs to.
func (o GetSentryOrganizationIntegrationResultOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v GetSentryOrganizationIntegrationResult) string { return v.Organization }).(pulumi.StringOutput)
}

// The key of the organization integration provider.
func (o GetSentryOrganizationIntegrationResultOutput) ProviderKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetSentryOrganizationIntegrationResult) string { return v.ProviderKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSentryOrganizationIntegrationResultOutput{})
}
