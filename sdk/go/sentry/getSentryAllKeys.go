// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-sentry/sdk/go/sentry/internal"
)

// List a Project's Client Keys.
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
//			_, err := sentry.GetSentryAllKeys(ctx, &sentry.GetSentryAllKeysArgs{
//				Organization: "my-organization",
//				Project:      "web-app",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSentryAllKeys(ctx *pulumi.Context, args *GetSentryAllKeysArgs, opts ...pulumi.InvokeOption) (*GetSentryAllKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSentryAllKeysResult
	err := ctx.Invoke("sentry:index/getSentryAllKeys:getSentryAllKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSentryAllKeys.
type GetSentryAllKeysArgs struct {
	// Filter client keys by `active` or `inactive`. Defaults to returning all keys if not specified.
	FilterStatus *string `pulumi:"filterStatus"`
	// The slug of the organization the resource belongs to.
	Organization string `pulumi:"organization"`
	// The slug of the project the resource belongs to.
	Project string `pulumi:"project"`
}

// A collection of values returned by getSentryAllKeys.
type GetSentryAllKeysResult struct {
	// Filter client keys by `active` or `inactive`. Defaults to returning all keys if not specified.
	FilterStatus *string `pulumi:"filterStatus"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of client keys.
	Keys []GetSentryAllKeysKey `pulumi:"keys"`
	// The slug of the organization the resource belongs to.
	Organization string `pulumi:"organization"`
	// The slug of the project the resource belongs to.
	Project string `pulumi:"project"`
}

func GetSentryAllKeysOutput(ctx *pulumi.Context, args GetSentryAllKeysOutputArgs, opts ...pulumi.InvokeOption) GetSentryAllKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSentryAllKeysResult, error) {
			args := v.(GetSentryAllKeysArgs)
			r, err := GetSentryAllKeys(ctx, &args, opts...)
			var s GetSentryAllKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSentryAllKeysResultOutput)
}

// A collection of arguments for invoking getSentryAllKeys.
type GetSentryAllKeysOutputArgs struct {
	// Filter client keys by `active` or `inactive`. Defaults to returning all keys if not specified.
	FilterStatus pulumi.StringPtrInput `pulumi:"filterStatus"`
	// The slug of the organization the resource belongs to.
	Organization pulumi.StringInput `pulumi:"organization"`
	// The slug of the project the resource belongs to.
	Project pulumi.StringInput `pulumi:"project"`
}

func (GetSentryAllKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSentryAllKeysArgs)(nil)).Elem()
}

// A collection of values returned by getSentryAllKeys.
type GetSentryAllKeysResultOutput struct{ *pulumi.OutputState }

func (GetSentryAllKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSentryAllKeysResult)(nil)).Elem()
}

func (o GetSentryAllKeysResultOutput) ToGetSentryAllKeysResultOutput() GetSentryAllKeysResultOutput {
	return o
}

func (o GetSentryAllKeysResultOutput) ToGetSentryAllKeysResultOutputWithContext(ctx context.Context) GetSentryAllKeysResultOutput {
	return o
}

// Filter client keys by `active` or `inactive`. Defaults to returning all keys if not specified.
func (o GetSentryAllKeysResultOutput) FilterStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSentryAllKeysResult) *string { return v.FilterStatus }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSentryAllKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSentryAllKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of client keys.
func (o GetSentryAllKeysResultOutput) Keys() GetSentryAllKeysKeyArrayOutput {
	return o.ApplyT(func(v GetSentryAllKeysResult) []GetSentryAllKeysKey { return v.Keys }).(GetSentryAllKeysKeyArrayOutput)
}

// The slug of the organization the resource belongs to.
func (o GetSentryAllKeysResultOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v GetSentryAllKeysResult) string { return v.Organization }).(pulumi.StringOutput)
}

// The slug of the project the resource belongs to.
func (o GetSentryAllKeysResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetSentryAllKeysResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSentryAllKeysResultOutput{})
}
