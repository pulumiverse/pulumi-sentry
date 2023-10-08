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

// Sentry Team data source.
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
//			_, err := sentry.LookupSentryTeam(ctx, &sentry.LookupSentryTeamArgs{
//				Organization: "my-organization",
//				Slug:         "my-team",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSentryTeam(ctx *pulumi.Context, args *LookupSentryTeamArgs, opts ...pulumi.InvokeOption) (*LookupSentryTeamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSentryTeamResult
	err := ctx.Invoke("sentry:index/getSentryTeam:getSentryTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSentryTeam.
type LookupSentryTeamArgs struct {
	// The slug of the organization the team should be created for.
	Organization string `pulumi:"organization"`
	// The unique URL slug for this team.
	Slug string `pulumi:"slug"`
}

// A collection of values returned by getSentryTeam.
type LookupSentryTeamResult struct {
	HasAccess bool `pulumi:"hasAccess"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The internal ID for this team.
	InternalId string `pulumi:"internalId"`
	IsMember   bool   `pulumi:"isMember"`
	IsPending  bool   `pulumi:"isPending"`
	// The human readable name for this organization.
	Name string `pulumi:"name"`
	// The slug of the organization the team should be created for.
	Organization string `pulumi:"organization"`
	// The unique URL slug for this team.
	Slug string `pulumi:"slug"`
}

func LookupSentryTeamOutput(ctx *pulumi.Context, args LookupSentryTeamOutputArgs, opts ...pulumi.InvokeOption) LookupSentryTeamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSentryTeamResult, error) {
			args := v.(LookupSentryTeamArgs)
			r, err := LookupSentryTeam(ctx, &args, opts...)
			var s LookupSentryTeamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSentryTeamResultOutput)
}

// A collection of arguments for invoking getSentryTeam.
type LookupSentryTeamOutputArgs struct {
	// The slug of the organization the team should be created for.
	Organization pulumi.StringInput `pulumi:"organization"`
	// The unique URL slug for this team.
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (LookupSentryTeamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSentryTeamArgs)(nil)).Elem()
}

// A collection of values returned by getSentryTeam.
type LookupSentryTeamResultOutput struct{ *pulumi.OutputState }

func (LookupSentryTeamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSentryTeamResult)(nil)).Elem()
}

func (o LookupSentryTeamResultOutput) ToLookupSentryTeamResultOutput() LookupSentryTeamResultOutput {
	return o
}

func (o LookupSentryTeamResultOutput) ToLookupSentryTeamResultOutputWithContext(ctx context.Context) LookupSentryTeamResultOutput {
	return o
}

func (o LookupSentryTeamResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSentryTeamResult] {
	return pulumix.Output[LookupSentryTeamResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupSentryTeamResultOutput) HasAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSentryTeamResult) bool { return v.HasAccess }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSentryTeamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryTeamResult) string { return v.Id }).(pulumi.StringOutput)
}

// The internal ID for this team.
func (o LookupSentryTeamResultOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryTeamResult) string { return v.InternalId }).(pulumi.StringOutput)
}

func (o LookupSentryTeamResultOutput) IsMember() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSentryTeamResult) bool { return v.IsMember }).(pulumi.BoolOutput)
}

func (o LookupSentryTeamResultOutput) IsPending() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSentryTeamResult) bool { return v.IsPending }).(pulumi.BoolOutput)
}

// The human readable name for this organization.
func (o LookupSentryTeamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryTeamResult) string { return v.Name }).(pulumi.StringOutput)
}

// The slug of the organization the team should be created for.
func (o LookupSentryTeamResultOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryTeamResult) string { return v.Organization }).(pulumi.StringOutput)
}

// The unique URL slug for this team.
func (o LookupSentryTeamResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryTeamResult) string { return v.Slug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSentryTeamResultOutput{})
}