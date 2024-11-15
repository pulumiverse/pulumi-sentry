// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-sentry/sdk/go/sentry/internal"
)

// ## Example Usage
func LookupSentryMetricAlert(ctx *pulumi.Context, args *LookupSentryMetricAlertArgs, opts ...pulumi.InvokeOption) (*LookupSentryMetricAlertResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSentryMetricAlertResult
	err := ctx.Invoke("sentry:index/getSentryMetricAlert:getSentryMetricAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSentryMetricAlert.
type LookupSentryMetricAlertArgs struct {
	// The internal ID for this metric alert.
	InternalId string `pulumi:"internalId"`
	// The slug of the organization the metric alert belongs to.
	Organization string `pulumi:"organization"`
	// The slug of the project the metric alert belongs to.
	Project string `pulumi:"project"`
}

// A collection of values returned by getSentryMetricAlert.
type LookupSentryMetricAlertResult struct {
	Aggregate   string `pulumi:"aggregate"`
	Dataset     string `pulumi:"dataset"`
	Environment string `pulumi:"environment"`
	// The events type of dataset.
	EventTypes []string `pulumi:"eventTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The internal ID for this metric alert.
	InternalId string `pulumi:"internalId"`
	// The metric alert name.
	Name string `pulumi:"name"`
	// The slug of the organization the metric alert belongs to.
	Organization string `pulumi:"organization"`
	Owner        string `pulumi:"owner"`
	// The slug of the project the metric alert belongs to.
	Project          string                        `pulumi:"project"`
	Query            string                        `pulumi:"query"`
	ResolveThreshold float64                       `pulumi:"resolveThreshold"`
	ThresholdType    int                           `pulumi:"thresholdType"`
	TimeWindow       float64                       `pulumi:"timeWindow"`
	Triggers         []GetSentryMetricAlertTrigger `pulumi:"triggers"`
}

func LookupSentryMetricAlertOutput(ctx *pulumi.Context, args LookupSentryMetricAlertOutputArgs, opts ...pulumi.InvokeOption) LookupSentryMetricAlertResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSentryMetricAlertResultOutput, error) {
			args := v.(LookupSentryMetricAlertArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupSentryMetricAlertResult
			secret, err := ctx.InvokePackageRaw("sentry:index/getSentryMetricAlert:getSentryMetricAlert", args, &rv, "", opts...)
			if err != nil {
				return LookupSentryMetricAlertResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupSentryMetricAlertResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupSentryMetricAlertResultOutput), nil
			}
			return output, nil
		}).(LookupSentryMetricAlertResultOutput)
}

// A collection of arguments for invoking getSentryMetricAlert.
type LookupSentryMetricAlertOutputArgs struct {
	// The internal ID for this metric alert.
	InternalId pulumi.StringInput `pulumi:"internalId"`
	// The slug of the organization the metric alert belongs to.
	Organization pulumi.StringInput `pulumi:"organization"`
	// The slug of the project the metric alert belongs to.
	Project pulumi.StringInput `pulumi:"project"`
}

func (LookupSentryMetricAlertOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSentryMetricAlertArgs)(nil)).Elem()
}

// A collection of values returned by getSentryMetricAlert.
type LookupSentryMetricAlertResultOutput struct{ *pulumi.OutputState }

func (LookupSentryMetricAlertResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSentryMetricAlertResult)(nil)).Elem()
}

func (o LookupSentryMetricAlertResultOutput) ToLookupSentryMetricAlertResultOutput() LookupSentryMetricAlertResultOutput {
	return o
}

func (o LookupSentryMetricAlertResultOutput) ToLookupSentryMetricAlertResultOutputWithContext(ctx context.Context) LookupSentryMetricAlertResultOutput {
	return o
}

func (o LookupSentryMetricAlertResultOutput) Aggregate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Aggregate }).(pulumi.StringOutput)
}

func (o LookupSentryMetricAlertResultOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Dataset }).(pulumi.StringOutput)
}

func (o LookupSentryMetricAlertResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Environment }).(pulumi.StringOutput)
}

// The events type of dataset.
func (o LookupSentryMetricAlertResultOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) []string { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSentryMetricAlertResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Id }).(pulumi.StringOutput)
}

// The internal ID for this metric alert.
func (o LookupSentryMetricAlertResultOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.InternalId }).(pulumi.StringOutput)
}

// The metric alert name.
func (o LookupSentryMetricAlertResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Name }).(pulumi.StringOutput)
}

// The slug of the organization the metric alert belongs to.
func (o LookupSentryMetricAlertResultOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Organization }).(pulumi.StringOutput)
}

func (o LookupSentryMetricAlertResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Owner }).(pulumi.StringOutput)
}

// The slug of the project the metric alert belongs to.
func (o LookupSentryMetricAlertResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupSentryMetricAlertResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) string { return v.Query }).(pulumi.StringOutput)
}

func (o LookupSentryMetricAlertResultOutput) ResolveThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) float64 { return v.ResolveThreshold }).(pulumi.Float64Output)
}

func (o LookupSentryMetricAlertResultOutput) ThresholdType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) int { return v.ThresholdType }).(pulumi.IntOutput)
}

func (o LookupSentryMetricAlertResultOutput) TimeWindow() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) float64 { return v.TimeWindow }).(pulumi.Float64Output)
}

func (o LookupSentryMetricAlertResultOutput) Triggers() GetSentryMetricAlertTriggerArrayOutput {
	return o.ApplyT(func(v LookupSentryMetricAlertResult) []GetSentryMetricAlertTrigger { return v.Triggers }).(GetSentryMetricAlertTriggerArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSentryMetricAlertResultOutput{})
}
