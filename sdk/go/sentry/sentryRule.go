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

// > **WARNING:** This resource is deprecated and will be removed in the next major version. Use the `SentryIssueAlert` resource instead.
type SentryRule struct {
	pulumi.CustomResourceState

	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch pulumi.StringOutput `pulumi:"actionMatch"`
	// List of actions.
	Actions pulumi.StringMapArrayOutput `pulumi:"actions"`
	// List of conditions.
	Conditions pulumi.StringMapArrayOutput `pulumi:"conditions"`
	// Perform issue alert in a specific environment.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch pulumi.StringOutput `pulumi:"filterMatch"`
	// List of filters.
	Filters pulumi.StringMapArrayOutput `pulumi:"filters"`
	// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
	Frequency pulumi.IntOutput `pulumi:"frequency"`
	// The internal ID for this issue alert.
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	// The issue alert name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The slug of the organization the issue alert belongs to.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The slug of the project to create the issue alert for.
	Project pulumi.StringOutput `pulumi:"project"`
	// Use `project` (singular) instead.
	//
	// Deprecated: Use `project` (singular) instead.
	Projects pulumi.StringArrayOutput `pulumi:"projects"`
}

// NewSentryRule registers a new resource with the given unique name, arguments, and options.
func NewSentryRule(ctx *pulumi.Context,
	name string, args *SentryRuleArgs, opts ...pulumi.ResourceOption) (*SentryRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionMatch == nil {
		return nil, errors.New("invalid value for required argument 'ActionMatch'")
	}
	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Conditions == nil {
		return nil, errors.New("invalid value for required argument 'Conditions'")
	}
	if args.FilterMatch == nil {
		return nil, errors.New("invalid value for required argument 'FilterMatch'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SentryRule
	err := ctx.RegisterResource("sentry:index/sentryRule:SentryRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSentryRule gets an existing SentryRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSentryRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SentryRuleState, opts ...pulumi.ResourceOption) (*SentryRule, error) {
	var resource SentryRule
	err := ctx.ReadResource("sentry:index/sentryRule:SentryRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SentryRule resources.
type sentryRuleState struct {
	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch *string `pulumi:"actionMatch"`
	// List of actions.
	Actions []map[string]string `pulumi:"actions"`
	// List of conditions.
	Conditions []map[string]string `pulumi:"conditions"`
	// Perform issue alert in a specific environment.
	Environment *string `pulumi:"environment"`
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch *string `pulumi:"filterMatch"`
	// List of filters.
	Filters []map[string]string `pulumi:"filters"`
	// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
	Frequency *int `pulumi:"frequency"`
	// The internal ID for this issue alert.
	InternalId *string `pulumi:"internalId"`
	// The issue alert name.
	Name *string `pulumi:"name"`
	// The slug of the organization the issue alert belongs to.
	Organization *string `pulumi:"organization"`
	// The slug of the project to create the issue alert for.
	Project *string `pulumi:"project"`
	// Use `project` (singular) instead.
	//
	// Deprecated: Use `project` (singular) instead.
	Projects []string `pulumi:"projects"`
}

type SentryRuleState struct {
	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch pulumi.StringPtrInput
	// List of actions.
	Actions pulumi.StringMapArrayInput
	// List of conditions.
	Conditions pulumi.StringMapArrayInput
	// Perform issue alert in a specific environment.
	Environment pulumi.StringPtrInput
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch pulumi.StringPtrInput
	// List of filters.
	Filters pulumi.StringMapArrayInput
	// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
	Frequency pulumi.IntPtrInput
	// The internal ID for this issue alert.
	InternalId pulumi.StringPtrInput
	// The issue alert name.
	Name pulumi.StringPtrInput
	// The slug of the organization the issue alert belongs to.
	Organization pulumi.StringPtrInput
	// The slug of the project to create the issue alert for.
	Project pulumi.StringPtrInput
	// Use `project` (singular) instead.
	//
	// Deprecated: Use `project` (singular) instead.
	Projects pulumi.StringArrayInput
}

func (SentryRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*sentryRuleState)(nil)).Elem()
}

type sentryRuleArgs struct {
	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch string `pulumi:"actionMatch"`
	// List of actions.
	Actions []map[string]string `pulumi:"actions"`
	// List of conditions.
	Conditions []map[string]string `pulumi:"conditions"`
	// Perform issue alert in a specific environment.
	Environment *string `pulumi:"environment"`
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch string `pulumi:"filterMatch"`
	// List of filters.
	Filters []map[string]string `pulumi:"filters"`
	// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
	Frequency int `pulumi:"frequency"`
	// The issue alert name.
	Name *string `pulumi:"name"`
	// The slug of the organization the issue alert belongs to.
	Organization string `pulumi:"organization"`
	// The slug of the project to create the issue alert for.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a SentryRule resource.
type SentryRuleArgs struct {
	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch pulumi.StringInput
	// List of actions.
	Actions pulumi.StringMapArrayInput
	// List of conditions.
	Conditions pulumi.StringMapArrayInput
	// Perform issue alert in a specific environment.
	Environment pulumi.StringPtrInput
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch pulumi.StringInput
	// List of filters.
	Filters pulumi.StringMapArrayInput
	// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
	Frequency pulumi.IntInput
	// The issue alert name.
	Name pulumi.StringPtrInput
	// The slug of the organization the issue alert belongs to.
	Organization pulumi.StringInput
	// The slug of the project to create the issue alert for.
	Project pulumi.StringInput
}

func (SentryRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sentryRuleArgs)(nil)).Elem()
}

type SentryRuleInput interface {
	pulumi.Input

	ToSentryRuleOutput() SentryRuleOutput
	ToSentryRuleOutputWithContext(ctx context.Context) SentryRuleOutput
}

func (*SentryRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SentryRule)(nil)).Elem()
}

func (i *SentryRule) ToSentryRuleOutput() SentryRuleOutput {
	return i.ToSentryRuleOutputWithContext(context.Background())
}

func (i *SentryRule) ToSentryRuleOutputWithContext(ctx context.Context) SentryRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryRuleOutput)
}

// SentryRuleArrayInput is an input type that accepts SentryRuleArray and SentryRuleArrayOutput values.
// You can construct a concrete instance of `SentryRuleArrayInput` via:
//
//	SentryRuleArray{ SentryRuleArgs{...} }
type SentryRuleArrayInput interface {
	pulumi.Input

	ToSentryRuleArrayOutput() SentryRuleArrayOutput
	ToSentryRuleArrayOutputWithContext(context.Context) SentryRuleArrayOutput
}

type SentryRuleArray []SentryRuleInput

func (SentryRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SentryRule)(nil)).Elem()
}

func (i SentryRuleArray) ToSentryRuleArrayOutput() SentryRuleArrayOutput {
	return i.ToSentryRuleArrayOutputWithContext(context.Background())
}

func (i SentryRuleArray) ToSentryRuleArrayOutputWithContext(ctx context.Context) SentryRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryRuleArrayOutput)
}

// SentryRuleMapInput is an input type that accepts SentryRuleMap and SentryRuleMapOutput values.
// You can construct a concrete instance of `SentryRuleMapInput` via:
//
//	SentryRuleMap{ "key": SentryRuleArgs{...} }
type SentryRuleMapInput interface {
	pulumi.Input

	ToSentryRuleMapOutput() SentryRuleMapOutput
	ToSentryRuleMapOutputWithContext(context.Context) SentryRuleMapOutput
}

type SentryRuleMap map[string]SentryRuleInput

func (SentryRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SentryRule)(nil)).Elem()
}

func (i SentryRuleMap) ToSentryRuleMapOutput() SentryRuleMapOutput {
	return i.ToSentryRuleMapOutputWithContext(context.Background())
}

func (i SentryRuleMap) ToSentryRuleMapOutputWithContext(ctx context.Context) SentryRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryRuleMapOutput)
}

type SentryRuleOutput struct{ *pulumi.OutputState }

func (SentryRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SentryRule)(nil)).Elem()
}

func (o SentryRuleOutput) ToSentryRuleOutput() SentryRuleOutput {
	return o
}

func (o SentryRuleOutput) ToSentryRuleOutputWithContext(ctx context.Context) SentryRuleOutput {
	return o
}

// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
func (o SentryRuleOutput) ActionMatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringOutput { return v.ActionMatch }).(pulumi.StringOutput)
}

// List of actions.
func (o SentryRuleOutput) Actions() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringMapArrayOutput { return v.Actions }).(pulumi.StringMapArrayOutput)
}

// List of conditions.
func (o SentryRuleOutput) Conditions() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringMapArrayOutput { return v.Conditions }).(pulumi.StringMapArrayOutput)
}

// Perform issue alert in a specific environment.
func (o SentryRuleOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Trigger actions if `all`, `any`, or `none` of the specified filters match.
func (o SentryRuleOutput) FilterMatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringOutput { return v.FilterMatch }).(pulumi.StringOutput)
}

// List of filters.
func (o SentryRuleOutput) Filters() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringMapArrayOutput { return v.Filters }).(pulumi.StringMapArrayOutput)
}

// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
func (o SentryRuleOutput) Frequency() pulumi.IntOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.IntOutput { return v.Frequency }).(pulumi.IntOutput)
}

// The internal ID for this issue alert.
func (o SentryRuleOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringOutput { return v.InternalId }).(pulumi.StringOutput)
}

// The issue alert name.
func (o SentryRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The slug of the organization the issue alert belongs to.
func (o SentryRuleOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The slug of the project to create the issue alert for.
func (o SentryRuleOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Use `project` (singular) instead.
//
// Deprecated: Use `project` (singular) instead.
func (o SentryRuleOutput) Projects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SentryRule) pulumi.StringArrayOutput { return v.Projects }).(pulumi.StringArrayOutput)
}

type SentryRuleArrayOutput struct{ *pulumi.OutputState }

func (SentryRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SentryRule)(nil)).Elem()
}

func (o SentryRuleArrayOutput) ToSentryRuleArrayOutput() SentryRuleArrayOutput {
	return o
}

func (o SentryRuleArrayOutput) ToSentryRuleArrayOutputWithContext(ctx context.Context) SentryRuleArrayOutput {
	return o
}

func (o SentryRuleArrayOutput) Index(i pulumi.IntInput) SentryRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SentryRule {
		return vs[0].([]*SentryRule)[vs[1].(int)]
	}).(SentryRuleOutput)
}

type SentryRuleMapOutput struct{ *pulumi.OutputState }

func (SentryRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SentryRule)(nil)).Elem()
}

func (o SentryRuleMapOutput) ToSentryRuleMapOutput() SentryRuleMapOutput {
	return o
}

func (o SentryRuleMapOutput) ToSentryRuleMapOutputWithContext(ctx context.Context) SentryRuleMapOutput {
	return o
}

func (o SentryRuleMapOutput) MapIndex(k pulumi.StringInput) SentryRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SentryRule {
		return vs[0].(map[string]*SentryRule)[vs[1].(string)]
	}).(SentryRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SentryRuleInput)(nil)).Elem(), &SentryRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SentryRuleArrayInput)(nil)).Elem(), SentryRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SentryRuleMapInput)(nil)).Elem(), SentryRuleMap{})
	pulumi.RegisterOutputType(SentryRuleOutput{})
	pulumi.RegisterOutputType(SentryRuleArrayOutput{})
	pulumi.RegisterOutputType(SentryRuleMapOutput{})
}
