// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-sentry/sdk/go/sentry/internal"
)

// Sentry Issue Alert resource. Note that there's no public documentation for the values of conditions, filters, and actions. You can either inspect the request payload sent when creating or editing an issue alert on Sentry or inspect [Sentry's rules registry in the source code](https://github.com/getsentry/sentry/tree/master/src/sentry/rules). Since v0.11.2, you should also omit the name property of each condition, filter, and action.
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
//			slack, err := sentry.GetSentryOrganizationIntegration(ctx, &sentry.GetSentryOrganizationIntegrationArgs{
//				Organization: sentry_project.Test.Organization,
//				ProviderKey:  "slack",
//				Name:         "Slack Workspace",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = sentry.NewSentryIssueAlert(ctx, "main", &sentry.SentryIssueAlertArgs{
//				Organization: pulumi.Any(sentry_project.Main.Organization),
//				Project:      pulumi.Any(sentry_project.Main.Id),
//				ActionMatch:  pulumi.String("any"),
//				FilterMatch:  pulumi.String("any"),
//				Frequency:    pulumi.Int(30),
//				Conditions: pulumi.AnyMapArray{
//					pulumi.AnyMap{
//						"id": pulumi.Any("sentry.rules.conditions.first_seen_event.FirstSeenEventCondition"),
//					},
//					pulumi.AnyMap{
//						"id": pulumi.Any("sentry.rules.conditions.regression_event.RegressionEventCondition"),
//					},
//					pulumi.AnyMap{
//						"id":             pulumi.Any("sentry.rules.conditions.event_frequency.EventFrequencyCondition"),
//						"value":          pulumi.Any(100),
//						"comparisonType": pulumi.Any("count"),
//						"interval":       pulumi.Any("1h"),
//					},
//					pulumi.AnyMap{
//						"id":             pulumi.Any("sentry.rules.conditions.event_frequency.EventUniqueUserFrequencyCondition"),
//						"value":          pulumi.Any(100),
//						"comparisonType": pulumi.Any("count"),
//						"interval":       pulumi.Any("1h"),
//					},
//					pulumi.AnyMap{
//						"id":             pulumi.Any("sentry.rules.conditions.event_frequency.EventFrequencyPercentCondition"),
//						"value":          pulumi.Any("50.0"),
//						"comparisonType": pulumi.Any("count"),
//						"interval":       pulumi.Any("1h"),
//					},
//				},
//				Filters: pulumi.AnyMapArray{
//					pulumi.AnyMap{
//						"id":              pulumi.Any("sentry.rules.filters.age_comparison.AgeComparisonFilter"),
//						"value":           pulumi.Any(10),
//						"time":            pulumi.Any("minute"),
//						"comparison_type": pulumi.Any("older"),
//					},
//					pulumi.AnyMap{
//						"id":    pulumi.Any("sentry.rules.filters.issue_occurrences.IssueOccurrencesFilter"),
//						"value": pulumi.Any(10),
//					},
//					pulumi.AnyMap{
//						"id":               pulumi.Any("sentry.rules.filters.assigned_to.AssignedToFilter"),
//						"targetType":       pulumi.Any("Team"),
//						"targetIdentifier": pulumi.Any(sentry_team.Main.Team_id),
//					},
//					pulumi.AnyMap{
//						"id": pulumi.Any("sentry.rules.filters.latest_release.LatestReleaseFilter"),
//					},
//					pulumi.AnyMap{
//						"id":        pulumi.Any("sentry.rules.filters.event_attribute.EventAttributeFilter"),
//						"attribute": pulumi.Any("message"),
//						"match":     pulumi.Any("co"),
//						"value":     pulumi.Any("test"),
//					},
//					pulumi.AnyMap{
//						"id":    pulumi.Any("sentry.rules.filters.tagged_event.TaggedEventFilter"),
//						"key":   pulumi.Any("test"),
//						"match": pulumi.Any("co"),
//						"value": pulumi.Any("test"),
//					},
//					pulumi.AnyMap{
//						"id":    pulumi.Any("sentry.rules.filters.level.LevelFilter"),
//						"match": pulumi.Any("eq"),
//						"level": pulumi.Any("50"),
//					},
//				},
//				Actions: pulumi.AnyMapArray{
//					pulumi.AnyMap{
//						"id":               pulumi.Any("sentry.mail.actions.NotifyEmailAction"),
//						"targetType":       pulumi.Any("IssueOwners"),
//						"targetIdentifier": pulumi.Any(""),
//					},
//					pulumi.AnyMap{
//						"id":               pulumi.Any("sentry.mail.actions.NotifyEmailAction"),
//						"targetType":       pulumi.Any("Team"),
//						"targetIdentifier": pulumi.Any(sentry_team.Main.Team_id),
//					},
//					pulumi.AnyMap{
//						"id": pulumi.Any("sentry.rules.actions.notify_event.NotifyEventAction"),
//					},
//					pulumi.AnyMap{
//						"id":        pulumi.Any("sentry.integrations.slack.notify_action.SlackNotifyServiceAction"),
//						"channel":   pulumi.Any("#general"),
//						"workspace": *pulumi.String(slack.InternalId),
//					},
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
// import using the organization, project slugs and rule id from the URLhttps://sentry.io/organizations/[org-slug]/alerts/rules/[project-slug]/[rule-id]/details/
//
// ```sh
//
//	$ pulumi import sentry:index/sentryIssueAlert:SentryIssueAlert default org-slug/project-slug/rule-id
//
// ```
type SentryIssueAlert struct {
	pulumi.CustomResourceState

	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch pulumi.StringOutput `pulumi:"actionMatch"`
	// List of actions.
	Actions pulumi.MapArrayOutput `pulumi:"actions"`
	// List of conditions.
	Conditions pulumi.MapArrayOutput `pulumi:"conditions"`
	// Perform issue alert in a specific environment.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch pulumi.StringOutput `pulumi:"filterMatch"`
	// List of filters.
	Filters pulumi.MapArrayOutput `pulumi:"filters"`
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

// NewSentryIssueAlert registers a new resource with the given unique name, arguments, and options.
func NewSentryIssueAlert(ctx *pulumi.Context,
	name string, args *SentryIssueAlertArgs, opts ...pulumi.ResourceOption) (*SentryIssueAlert, error) {
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
	var resource SentryIssueAlert
	err := ctx.RegisterResource("sentry:index/sentryIssueAlert:SentryIssueAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSentryIssueAlert gets an existing SentryIssueAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSentryIssueAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SentryIssueAlertState, opts ...pulumi.ResourceOption) (*SentryIssueAlert, error) {
	var resource SentryIssueAlert
	err := ctx.ReadResource("sentry:index/sentryIssueAlert:SentryIssueAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SentryIssueAlert resources.
type sentryIssueAlertState struct {
	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch *string `pulumi:"actionMatch"`
	// List of actions.
	Actions []map[string]interface{} `pulumi:"actions"`
	// List of conditions.
	Conditions []map[string]interface{} `pulumi:"conditions"`
	// Perform issue alert in a specific environment.
	Environment *string `pulumi:"environment"`
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch *string `pulumi:"filterMatch"`
	// List of filters.
	Filters []map[string]interface{} `pulumi:"filters"`
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

type SentryIssueAlertState struct {
	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch pulumi.StringPtrInput
	// List of actions.
	Actions pulumi.MapArrayInput
	// List of conditions.
	Conditions pulumi.MapArrayInput
	// Perform issue alert in a specific environment.
	Environment pulumi.StringPtrInput
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch pulumi.StringPtrInput
	// List of filters.
	Filters pulumi.MapArrayInput
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

func (SentryIssueAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*sentryIssueAlertState)(nil)).Elem()
}

type sentryIssueAlertArgs struct {
	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch string `pulumi:"actionMatch"`
	// List of actions.
	Actions []map[string]interface{} `pulumi:"actions"`
	// List of conditions.
	Conditions []map[string]interface{} `pulumi:"conditions"`
	// Perform issue alert in a specific environment.
	Environment *string `pulumi:"environment"`
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch string `pulumi:"filterMatch"`
	// List of filters.
	Filters []map[string]interface{} `pulumi:"filters"`
	// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
	Frequency int `pulumi:"frequency"`
	// The issue alert name.
	Name *string `pulumi:"name"`
	// The slug of the organization the issue alert belongs to.
	Organization string `pulumi:"organization"`
	// The slug of the project to create the issue alert for.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a SentryIssueAlert resource.
type SentryIssueAlertArgs struct {
	// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
	ActionMatch pulumi.StringInput
	// List of actions.
	Actions pulumi.MapArrayInput
	// List of conditions.
	Conditions pulumi.MapArrayInput
	// Perform issue alert in a specific environment.
	Environment pulumi.StringPtrInput
	// Trigger actions if `all`, `any`, or `none` of the specified filters match.
	FilterMatch pulumi.StringInput
	// List of filters.
	Filters pulumi.MapArrayInput
	// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
	Frequency pulumi.IntInput
	// The issue alert name.
	Name pulumi.StringPtrInput
	// The slug of the organization the issue alert belongs to.
	Organization pulumi.StringInput
	// The slug of the project to create the issue alert for.
	Project pulumi.StringInput
}

func (SentryIssueAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sentryIssueAlertArgs)(nil)).Elem()
}

type SentryIssueAlertInput interface {
	pulumi.Input

	ToSentryIssueAlertOutput() SentryIssueAlertOutput
	ToSentryIssueAlertOutputWithContext(ctx context.Context) SentryIssueAlertOutput
}

func (*SentryIssueAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**SentryIssueAlert)(nil)).Elem()
}

func (i *SentryIssueAlert) ToSentryIssueAlertOutput() SentryIssueAlertOutput {
	return i.ToSentryIssueAlertOutputWithContext(context.Background())
}

func (i *SentryIssueAlert) ToSentryIssueAlertOutputWithContext(ctx context.Context) SentryIssueAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryIssueAlertOutput)
}

func (i *SentryIssueAlert) ToOutput(ctx context.Context) pulumix.Output[*SentryIssueAlert] {
	return pulumix.Output[*SentryIssueAlert]{
		OutputState: i.ToSentryIssueAlertOutputWithContext(ctx).OutputState,
	}
}

// SentryIssueAlertArrayInput is an input type that accepts SentryIssueAlertArray and SentryIssueAlertArrayOutput values.
// You can construct a concrete instance of `SentryIssueAlertArrayInput` via:
//
//	SentryIssueAlertArray{ SentryIssueAlertArgs{...} }
type SentryIssueAlertArrayInput interface {
	pulumi.Input

	ToSentryIssueAlertArrayOutput() SentryIssueAlertArrayOutput
	ToSentryIssueAlertArrayOutputWithContext(context.Context) SentryIssueAlertArrayOutput
}

type SentryIssueAlertArray []SentryIssueAlertInput

func (SentryIssueAlertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SentryIssueAlert)(nil)).Elem()
}

func (i SentryIssueAlertArray) ToSentryIssueAlertArrayOutput() SentryIssueAlertArrayOutput {
	return i.ToSentryIssueAlertArrayOutputWithContext(context.Background())
}

func (i SentryIssueAlertArray) ToSentryIssueAlertArrayOutputWithContext(ctx context.Context) SentryIssueAlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryIssueAlertArrayOutput)
}

func (i SentryIssueAlertArray) ToOutput(ctx context.Context) pulumix.Output[[]*SentryIssueAlert] {
	return pulumix.Output[[]*SentryIssueAlert]{
		OutputState: i.ToSentryIssueAlertArrayOutputWithContext(ctx).OutputState,
	}
}

// SentryIssueAlertMapInput is an input type that accepts SentryIssueAlertMap and SentryIssueAlertMapOutput values.
// You can construct a concrete instance of `SentryIssueAlertMapInput` via:
//
//	SentryIssueAlertMap{ "key": SentryIssueAlertArgs{...} }
type SentryIssueAlertMapInput interface {
	pulumi.Input

	ToSentryIssueAlertMapOutput() SentryIssueAlertMapOutput
	ToSentryIssueAlertMapOutputWithContext(context.Context) SentryIssueAlertMapOutput
}

type SentryIssueAlertMap map[string]SentryIssueAlertInput

func (SentryIssueAlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SentryIssueAlert)(nil)).Elem()
}

func (i SentryIssueAlertMap) ToSentryIssueAlertMapOutput() SentryIssueAlertMapOutput {
	return i.ToSentryIssueAlertMapOutputWithContext(context.Background())
}

func (i SentryIssueAlertMap) ToSentryIssueAlertMapOutputWithContext(ctx context.Context) SentryIssueAlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryIssueAlertMapOutput)
}

func (i SentryIssueAlertMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SentryIssueAlert] {
	return pulumix.Output[map[string]*SentryIssueAlert]{
		OutputState: i.ToSentryIssueAlertMapOutputWithContext(ctx).OutputState,
	}
}

type SentryIssueAlertOutput struct{ *pulumi.OutputState }

func (SentryIssueAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SentryIssueAlert)(nil)).Elem()
}

func (o SentryIssueAlertOutput) ToSentryIssueAlertOutput() SentryIssueAlertOutput {
	return o
}

func (o SentryIssueAlertOutput) ToSentryIssueAlertOutputWithContext(ctx context.Context) SentryIssueAlertOutput {
	return o
}

func (o SentryIssueAlertOutput) ToOutput(ctx context.Context) pulumix.Output[*SentryIssueAlert] {
	return pulumix.Output[*SentryIssueAlert]{
		OutputState: o.OutputState,
	}
}

// Trigger actions when an event is captured by Sentry and `any` or `all` of the specified conditions happen.
func (o SentryIssueAlertOutput) ActionMatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.StringOutput { return v.ActionMatch }).(pulumi.StringOutput)
}

// List of actions.
func (o SentryIssueAlertOutput) Actions() pulumi.MapArrayOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.MapArrayOutput { return v.Actions }).(pulumi.MapArrayOutput)
}

// List of conditions.
func (o SentryIssueAlertOutput) Conditions() pulumi.MapArrayOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.MapArrayOutput { return v.Conditions }).(pulumi.MapArrayOutput)
}

// Perform issue alert in a specific environment.
func (o SentryIssueAlertOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Trigger actions if `all`, `any`, or `none` of the specified filters match.
func (o SentryIssueAlertOutput) FilterMatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.StringOutput { return v.FilterMatch }).(pulumi.StringOutput)
}

// List of filters.
func (o SentryIssueAlertOutput) Filters() pulumi.MapArrayOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.MapArrayOutput { return v.Filters }).(pulumi.MapArrayOutput)
}

// Perform actions at most once every `X` minutes for this issue. Defaults to `30`.
func (o SentryIssueAlertOutput) Frequency() pulumi.IntOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.IntOutput { return v.Frequency }).(pulumi.IntOutput)
}

// The internal ID for this issue alert.
func (o SentryIssueAlertOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.StringOutput { return v.InternalId }).(pulumi.StringOutput)
}

// The issue alert name.
func (o SentryIssueAlertOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The slug of the organization the issue alert belongs to.
func (o SentryIssueAlertOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The slug of the project to create the issue alert for.
func (o SentryIssueAlertOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Use `project` (singular) instead.
//
// Deprecated: Use `project` (singular) instead.
func (o SentryIssueAlertOutput) Projects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SentryIssueAlert) pulumi.StringArrayOutput { return v.Projects }).(pulumi.StringArrayOutput)
}

type SentryIssueAlertArrayOutput struct{ *pulumi.OutputState }

func (SentryIssueAlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SentryIssueAlert)(nil)).Elem()
}

func (o SentryIssueAlertArrayOutput) ToSentryIssueAlertArrayOutput() SentryIssueAlertArrayOutput {
	return o
}

func (o SentryIssueAlertArrayOutput) ToSentryIssueAlertArrayOutputWithContext(ctx context.Context) SentryIssueAlertArrayOutput {
	return o
}

func (o SentryIssueAlertArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SentryIssueAlert] {
	return pulumix.Output[[]*SentryIssueAlert]{
		OutputState: o.OutputState,
	}
}

func (o SentryIssueAlertArrayOutput) Index(i pulumi.IntInput) SentryIssueAlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SentryIssueAlert {
		return vs[0].([]*SentryIssueAlert)[vs[1].(int)]
	}).(SentryIssueAlertOutput)
}

type SentryIssueAlertMapOutput struct{ *pulumi.OutputState }

func (SentryIssueAlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SentryIssueAlert)(nil)).Elem()
}

func (o SentryIssueAlertMapOutput) ToSentryIssueAlertMapOutput() SentryIssueAlertMapOutput {
	return o
}

func (o SentryIssueAlertMapOutput) ToSentryIssueAlertMapOutputWithContext(ctx context.Context) SentryIssueAlertMapOutput {
	return o
}

func (o SentryIssueAlertMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SentryIssueAlert] {
	return pulumix.Output[map[string]*SentryIssueAlert]{
		OutputState: o.OutputState,
	}
}

func (o SentryIssueAlertMapOutput) MapIndex(k pulumi.StringInput) SentryIssueAlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SentryIssueAlert {
		return vs[0].(map[string]*SentryIssueAlert)[vs[1].(string)]
	}).(SentryIssueAlertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SentryIssueAlertInput)(nil)).Elem(), &SentryIssueAlert{})
	pulumi.RegisterInputType(reflect.TypeOf((*SentryIssueAlertArrayInput)(nil)).Elem(), SentryIssueAlertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SentryIssueAlertMapInput)(nil)).Elem(), SentryIssueAlertMap{})
	pulumi.RegisterOutputType(SentryIssueAlertOutput{})
	pulumi.RegisterOutputType(SentryIssueAlertArrayOutput{})
	pulumi.RegisterOutputType(SentryIssueAlertMapOutput{})
}
