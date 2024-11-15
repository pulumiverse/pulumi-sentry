// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-sentry/sdk/go/sentry/internal"
)

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
//			// Create an organization member
//			_, err := sentry.NewSentryOrganizationMember(ctx, "john_doe", &sentry.SentryOrganizationMemberArgs{
//				Organization: pulumi.String("my-organization"),
//				Email:        pulumi.String("test@example.com"),
//				Role:         pulumi.String("member"),
//				Teams: pulumi.StringArray{
//					pulumi.String("my-team"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// import using the organization, membership id from the URL:
//
// https://sentry.io/settings/[org-slug]/members/[member-id]/
//
// ```sh
// $ pulumi import sentry:index/sentryOrganizationMember:SentryOrganizationMember john_doe org-slug/member-id
// ```
type SentryOrganizationMember struct {
	pulumi.CustomResourceState

	// The email of the organization member.
	Email pulumi.StringOutput `pulumi:"email"`
	// The invite has expired.
	Expired pulumi.BoolOutput `pulumi:"expired"`
	// The internal ID for this organization membership.
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	// The slug of the organization the user should be invited to.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The invite is pending.
	Pending pulumi.BoolOutput `pulumi:"pending"`
	// This is the role of the organization member.
	Role pulumi.StringOutput `pulumi:"role"`
	// The teams the organization member should be added to.
	Teams pulumi.StringArrayOutput `pulumi:"teams"`
}

// NewSentryOrganizationMember registers a new resource with the given unique name, arguments, and options.
func NewSentryOrganizationMember(ctx *pulumi.Context,
	name string, args *SentryOrganizationMemberArgs, opts ...pulumi.ResourceOption) (*SentryOrganizationMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SentryOrganizationMember
	err := ctx.RegisterResource("sentry:index/sentryOrganizationMember:SentryOrganizationMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSentryOrganizationMember gets an existing SentryOrganizationMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSentryOrganizationMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SentryOrganizationMemberState, opts ...pulumi.ResourceOption) (*SentryOrganizationMember, error) {
	var resource SentryOrganizationMember
	err := ctx.ReadResource("sentry:index/sentryOrganizationMember:SentryOrganizationMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SentryOrganizationMember resources.
type sentryOrganizationMemberState struct {
	// The email of the organization member.
	Email *string `pulumi:"email"`
	// The invite has expired.
	Expired *bool `pulumi:"expired"`
	// The internal ID for this organization membership.
	InternalId *string `pulumi:"internalId"`
	// The slug of the organization the user should be invited to.
	Organization *string `pulumi:"organization"`
	// The invite is pending.
	Pending *bool `pulumi:"pending"`
	// This is the role of the organization member.
	Role *string `pulumi:"role"`
	// The teams the organization member should be added to.
	Teams []string `pulumi:"teams"`
}

type SentryOrganizationMemberState struct {
	// The email of the organization member.
	Email pulumi.StringPtrInput
	// The invite has expired.
	Expired pulumi.BoolPtrInput
	// The internal ID for this organization membership.
	InternalId pulumi.StringPtrInput
	// The slug of the organization the user should be invited to.
	Organization pulumi.StringPtrInput
	// The invite is pending.
	Pending pulumi.BoolPtrInput
	// This is the role of the organization member.
	Role pulumi.StringPtrInput
	// The teams the organization member should be added to.
	Teams pulumi.StringArrayInput
}

func (SentryOrganizationMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*sentryOrganizationMemberState)(nil)).Elem()
}

type sentryOrganizationMemberArgs struct {
	// The email of the organization member.
	Email string `pulumi:"email"`
	// The slug of the organization the user should be invited to.
	Organization string `pulumi:"organization"`
	// This is the role of the organization member.
	Role string `pulumi:"role"`
	// The teams the organization member should be added to.
	Teams []string `pulumi:"teams"`
}

// The set of arguments for constructing a SentryOrganizationMember resource.
type SentryOrganizationMemberArgs struct {
	// The email of the organization member.
	Email pulumi.StringInput
	// The slug of the organization the user should be invited to.
	Organization pulumi.StringInput
	// This is the role of the organization member.
	Role pulumi.StringInput
	// The teams the organization member should be added to.
	Teams pulumi.StringArrayInput
}

func (SentryOrganizationMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sentryOrganizationMemberArgs)(nil)).Elem()
}

type SentryOrganizationMemberInput interface {
	pulumi.Input

	ToSentryOrganizationMemberOutput() SentryOrganizationMemberOutput
	ToSentryOrganizationMemberOutputWithContext(ctx context.Context) SentryOrganizationMemberOutput
}

func (*SentryOrganizationMember) ElementType() reflect.Type {
	return reflect.TypeOf((**SentryOrganizationMember)(nil)).Elem()
}

func (i *SentryOrganizationMember) ToSentryOrganizationMemberOutput() SentryOrganizationMemberOutput {
	return i.ToSentryOrganizationMemberOutputWithContext(context.Background())
}

func (i *SentryOrganizationMember) ToSentryOrganizationMemberOutputWithContext(ctx context.Context) SentryOrganizationMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryOrganizationMemberOutput)
}

// SentryOrganizationMemberArrayInput is an input type that accepts SentryOrganizationMemberArray and SentryOrganizationMemberArrayOutput values.
// You can construct a concrete instance of `SentryOrganizationMemberArrayInput` via:
//
//	SentryOrganizationMemberArray{ SentryOrganizationMemberArgs{...} }
type SentryOrganizationMemberArrayInput interface {
	pulumi.Input

	ToSentryOrganizationMemberArrayOutput() SentryOrganizationMemberArrayOutput
	ToSentryOrganizationMemberArrayOutputWithContext(context.Context) SentryOrganizationMemberArrayOutput
}

type SentryOrganizationMemberArray []SentryOrganizationMemberInput

func (SentryOrganizationMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SentryOrganizationMember)(nil)).Elem()
}

func (i SentryOrganizationMemberArray) ToSentryOrganizationMemberArrayOutput() SentryOrganizationMemberArrayOutput {
	return i.ToSentryOrganizationMemberArrayOutputWithContext(context.Background())
}

func (i SentryOrganizationMemberArray) ToSentryOrganizationMemberArrayOutputWithContext(ctx context.Context) SentryOrganizationMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryOrganizationMemberArrayOutput)
}

// SentryOrganizationMemberMapInput is an input type that accepts SentryOrganizationMemberMap and SentryOrganizationMemberMapOutput values.
// You can construct a concrete instance of `SentryOrganizationMemberMapInput` via:
//
//	SentryOrganizationMemberMap{ "key": SentryOrganizationMemberArgs{...} }
type SentryOrganizationMemberMapInput interface {
	pulumi.Input

	ToSentryOrganizationMemberMapOutput() SentryOrganizationMemberMapOutput
	ToSentryOrganizationMemberMapOutputWithContext(context.Context) SentryOrganizationMemberMapOutput
}

type SentryOrganizationMemberMap map[string]SentryOrganizationMemberInput

func (SentryOrganizationMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SentryOrganizationMember)(nil)).Elem()
}

func (i SentryOrganizationMemberMap) ToSentryOrganizationMemberMapOutput() SentryOrganizationMemberMapOutput {
	return i.ToSentryOrganizationMemberMapOutputWithContext(context.Background())
}

func (i SentryOrganizationMemberMap) ToSentryOrganizationMemberMapOutputWithContext(ctx context.Context) SentryOrganizationMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryOrganizationMemberMapOutput)
}

type SentryOrganizationMemberOutput struct{ *pulumi.OutputState }

func (SentryOrganizationMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SentryOrganizationMember)(nil)).Elem()
}

func (o SentryOrganizationMemberOutput) ToSentryOrganizationMemberOutput() SentryOrganizationMemberOutput {
	return o
}

func (o SentryOrganizationMemberOutput) ToSentryOrganizationMemberOutputWithContext(ctx context.Context) SentryOrganizationMemberOutput {
	return o
}

// The email of the organization member.
func (o SentryOrganizationMemberOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryOrganizationMember) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The invite has expired.
func (o SentryOrganizationMemberOutput) Expired() pulumi.BoolOutput {
	return o.ApplyT(func(v *SentryOrganizationMember) pulumi.BoolOutput { return v.Expired }).(pulumi.BoolOutput)
}

// The internal ID for this organization membership.
func (o SentryOrganizationMemberOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryOrganizationMember) pulumi.StringOutput { return v.InternalId }).(pulumi.StringOutput)
}

// The slug of the organization the user should be invited to.
func (o SentryOrganizationMemberOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryOrganizationMember) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The invite is pending.
func (o SentryOrganizationMemberOutput) Pending() pulumi.BoolOutput {
	return o.ApplyT(func(v *SentryOrganizationMember) pulumi.BoolOutput { return v.Pending }).(pulumi.BoolOutput)
}

// This is the role of the organization member.
func (o SentryOrganizationMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryOrganizationMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The teams the organization member should be added to.
func (o SentryOrganizationMemberOutput) Teams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SentryOrganizationMember) pulumi.StringArrayOutput { return v.Teams }).(pulumi.StringArrayOutput)
}

type SentryOrganizationMemberArrayOutput struct{ *pulumi.OutputState }

func (SentryOrganizationMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SentryOrganizationMember)(nil)).Elem()
}

func (o SentryOrganizationMemberArrayOutput) ToSentryOrganizationMemberArrayOutput() SentryOrganizationMemberArrayOutput {
	return o
}

func (o SentryOrganizationMemberArrayOutput) ToSentryOrganizationMemberArrayOutputWithContext(ctx context.Context) SentryOrganizationMemberArrayOutput {
	return o
}

func (o SentryOrganizationMemberArrayOutput) Index(i pulumi.IntInput) SentryOrganizationMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SentryOrganizationMember {
		return vs[0].([]*SentryOrganizationMember)[vs[1].(int)]
	}).(SentryOrganizationMemberOutput)
}

type SentryOrganizationMemberMapOutput struct{ *pulumi.OutputState }

func (SentryOrganizationMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SentryOrganizationMember)(nil)).Elem()
}

func (o SentryOrganizationMemberMapOutput) ToSentryOrganizationMemberMapOutput() SentryOrganizationMemberMapOutput {
	return o
}

func (o SentryOrganizationMemberMapOutput) ToSentryOrganizationMemberMapOutputWithContext(ctx context.Context) SentryOrganizationMemberMapOutput {
	return o
}

func (o SentryOrganizationMemberMapOutput) MapIndex(k pulumi.StringInput) SentryOrganizationMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SentryOrganizationMember {
		return vs[0].(map[string]*SentryOrganizationMember)[vs[1].(string)]
	}).(SentryOrganizationMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SentryOrganizationMemberInput)(nil)).Elem(), &SentryOrganizationMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*SentryOrganizationMemberArrayInput)(nil)).Elem(), SentryOrganizationMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SentryOrganizationMemberMapInput)(nil)).Elem(), SentryOrganizationMemberMap{})
	pulumi.RegisterOutputType(SentryOrganizationMemberOutput{})
	pulumi.RegisterOutputType(SentryOrganizationMemberArrayOutput{})
	pulumi.RegisterOutputType(SentryOrganizationMemberMapOutput{})
}
