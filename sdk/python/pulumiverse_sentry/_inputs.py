# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'SentryDashboardWidgetArgs',
    'SentryDashboardWidgetLayoutArgs',
    'SentryDashboardWidgetQueryArgs',
    'SentryMetricAlertTriggerArgs',
    'SentryMetricAlertTriggerActionArgs',
    'SentryProjectSymbolSourceLayoutArgs',
]

@pulumi.input_type
class SentryDashboardWidgetArgs:
    def __init__(__self__, *,
                 display_type: pulumi.Input[str],
                 layout: pulumi.Input['SentryDashboardWidgetLayoutArgs'],
                 queries: pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetQueryArgs']]],
                 title: pulumi.Input[str],
                 id: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[str]] = None,
                 limit: Optional[pulumi.Input[int]] = None,
                 widget_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        """
        pulumi.set(__self__, "display_type", display_type)
        pulumi.set(__self__, "layout", layout)
        pulumi.set(__self__, "queries", queries)
        pulumi.set(__self__, "title", title)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if limit is not None:
            pulumi.set(__self__, "limit", limit)
        if widget_type is not None:
            pulumi.set(__self__, "widget_type", widget_type)

    @property
    @pulumi.getter(name="displayType")
    def display_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "display_type")

    @display_type.setter
    def display_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_type", value)

    @property
    @pulumi.getter
    def layout(self) -> pulumi.Input['SentryDashboardWidgetLayoutArgs']:
        return pulumi.get(self, "layout")

    @layout.setter
    def layout(self, value: pulumi.Input['SentryDashboardWidgetLayoutArgs']):
        pulumi.set(self, "layout", value)

    @property
    @pulumi.getter
    def queries(self) -> pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetQueryArgs']]]:
        return pulumi.get(self, "queries")

    @queries.setter
    def queries(self, value: pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetQueryArgs']]]):
        pulumi.set(self, "queries", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def limit(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "limit")

    @limit.setter
    def limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "limit", value)

    @property
    @pulumi.getter(name="widgetType")
    def widget_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "widget_type")

    @widget_type.setter
    def widget_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "widget_type", value)


@pulumi.input_type
class SentryDashboardWidgetLayoutArgs:
    def __init__(__self__, *,
                 h: pulumi.Input[int],
                 min_h: pulumi.Input[int],
                 w: pulumi.Input[int],
                 x: pulumi.Input[int],
                 y: pulumi.Input[int]):
        pulumi.set(__self__, "h", h)
        pulumi.set(__self__, "min_h", min_h)
        pulumi.set(__self__, "w", w)
        pulumi.set(__self__, "x", x)
        pulumi.set(__self__, "y", y)

    @property
    @pulumi.getter
    def h(self) -> pulumi.Input[int]:
        return pulumi.get(self, "h")

    @h.setter
    def h(self, value: pulumi.Input[int]):
        pulumi.set(self, "h", value)

    @property
    @pulumi.getter(name="minH")
    def min_h(self) -> pulumi.Input[int]:
        return pulumi.get(self, "min_h")

    @min_h.setter
    def min_h(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_h", value)

    @property
    @pulumi.getter
    def w(self) -> pulumi.Input[int]:
        return pulumi.get(self, "w")

    @w.setter
    def w(self, value: pulumi.Input[int]):
        pulumi.set(self, "w", value)

    @property
    @pulumi.getter
    def x(self) -> pulumi.Input[int]:
        return pulumi.get(self, "x")

    @x.setter
    def x(self, value: pulumi.Input[int]):
        pulumi.set(self, "x", value)

    @property
    @pulumi.getter
    def y(self) -> pulumi.Input[int]:
        return pulumi.get(self, "y")

    @y.setter
    def y(self, value: pulumi.Input[int]):
        pulumi.set(self, "y", value)


@pulumi.input_type
class SentryDashboardWidgetQueryArgs:
    def __init__(__self__, *,
                 aggregates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 columns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 conditions: Optional[pulumi.Input[str]] = None,
                 field_aliases: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order_by: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        """
        if aggregates is not None:
            pulumi.set(__self__, "aggregates", aggregates)
        if columns is not None:
            pulumi.set(__self__, "columns", columns)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if field_aliases is not None:
            pulumi.set(__self__, "field_aliases", field_aliases)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if order_by is not None:
            pulumi.set(__self__, "order_by", order_by)

    @property
    @pulumi.getter
    def aggregates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "aggregates")

    @aggregates.setter
    def aggregates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "aggregates", value)

    @property
    @pulumi.getter
    def columns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "columns")

    @columns.setter
    def columns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "columns", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="fieldAliases")
    def field_aliases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "field_aliases")

    @field_aliases.setter
    def field_aliases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "field_aliases", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "order_by")

    @order_by.setter
    def order_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "order_by", value)


@pulumi.input_type
class SentryMetricAlertTriggerArgs:
    def __init__(__self__, *,
                 alert_threshold: pulumi.Input[float],
                 label: pulumi.Input[str],
                 threshold_type: pulumi.Input[int],
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input['SentryMetricAlertTriggerActionArgs']]]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 resolve_threshold: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        """
        pulumi.set(__self__, "alert_threshold", alert_threshold)
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "threshold_type", threshold_type)
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if resolve_threshold is not None:
            pulumi.set(__self__, "resolve_threshold", resolve_threshold)

    @property
    @pulumi.getter(name="alertThreshold")
    def alert_threshold(self) -> pulumi.Input[float]:
        return pulumi.get(self, "alert_threshold")

    @alert_threshold.setter
    def alert_threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "alert_threshold", value)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[str]:
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter(name="thresholdType")
    def threshold_type(self) -> pulumi.Input[int]:
        return pulumi.get(self, "threshold_type")

    @threshold_type.setter
    def threshold_type(self, value: pulumi.Input[int]):
        pulumi.set(self, "threshold_type", value)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SentryMetricAlertTriggerActionArgs']]]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SentryMetricAlertTriggerActionArgs']]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="resolveThreshold")
    def resolve_threshold(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "resolve_threshold")

    @resolve_threshold.setter
    def resolve_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "resolve_threshold", value)


@pulumi.input_type
class SentryMetricAlertTriggerActionArgs:
    def __init__(__self__, *,
                 target_type: pulumi.Input[str],
                 type: pulumi.Input[str],
                 id: Optional[pulumi.Input[str]] = None,
                 input_channel_id: Optional[pulumi.Input[str]] = None,
                 integration_id: Optional[pulumi.Input[int]] = None,
                 target_identifier: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        :param pulumi.Input[str] input_channel_id: Slack channel ID to avoid rate-limiting, see [here](https://docs.sentry.io/product/integrations/notification-incidents/slack/#rate-limiting-error)
        """
        pulumi.set(__self__, "target_type", target_type)
        pulumi.set(__self__, "type", type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if input_channel_id is not None:
            pulumi.set(__self__, "input_channel_id", input_channel_id)
        if integration_id is not None:
            pulumi.set(__self__, "integration_id", integration_id)
        if target_identifier is not None:
            pulumi.set(__self__, "target_identifier", target_identifier)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_type", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="inputChannelId")
    def input_channel_id(self) -> Optional[pulumi.Input[str]]:
        """
        Slack channel ID to avoid rate-limiting, see [here](https://docs.sentry.io/product/integrations/notification-incidents/slack/#rate-limiting-error)
        """
        return pulumi.get(self, "input_channel_id")

    @input_channel_id.setter
    def input_channel_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_channel_id", value)

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "integration_id")

    @integration_id.setter
    def integration_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "integration_id", value)

    @property
    @pulumi.getter(name="targetIdentifier")
    def target_identifier(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target_identifier")

    @target_identifier.setter
    def target_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_identifier", value)


@pulumi.input_type
class SentryProjectSymbolSourceLayoutArgs:
    def __init__(__self__, *,
                 casing: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] casing: The casing of the symbol source layout. The layout of the folder structure. The options are: `default` - Default (mixed case), `uppercase` - Uppercase, `lowercase` - Lowercase.
        :param pulumi.Input[str] type: The layout of the folder structure. The options are: `native` - Platform-Specific (SymStore / GDB / LLVM), `symstore` - Microsoft SymStore, `symstore_index2` - Microsoft SymStore (with index2.txt), `ssqp` - Microsoft SSQP, `unified` - Unified Symbol Server Layout, `debuginfod` - debuginfod.
        """
        pulumi.set(__self__, "casing", casing)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def casing(self) -> pulumi.Input[str]:
        """
        The casing of the symbol source layout. The layout of the folder structure. The options are: `default` - Default (mixed case), `uppercase` - Uppercase, `lowercase` - Lowercase.
        """
        return pulumi.get(self, "casing")

    @casing.setter
    def casing(self, value: pulumi.Input[str]):
        pulumi.set(self, "casing", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The layout of the folder structure. The options are: `native` - Platform-Specific (SymStore / GDB / LLVM), `symstore` - Microsoft SymStore, `symstore_index2` - Microsoft SymStore (with index2.txt), `ssqp` - Microsoft SSQP, `unified` - Unified Symbol Server Layout, `debuginfod` - debuginfod.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


