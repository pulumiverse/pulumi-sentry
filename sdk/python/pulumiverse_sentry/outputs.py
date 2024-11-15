# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'SentryDashboardWidget',
    'SentryDashboardWidgetLayout',
    'SentryDashboardWidgetQuery',
    'SentryMetricAlertTrigger',
    'SentryMetricAlertTriggerAction',
    'GetSentryDashboardWidgetResult',
    'GetSentryDashboardWidgetLayoutResult',
    'GetSentryDashboardWidgetQueryResult',
    'GetSentryMetricAlertTriggerResult',
    'GetSentryMetricAlertTriggerActionResult',
]

@pulumi.output_type
class SentryDashboardWidget(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "displayType":
            suggest = "display_type"
        elif key == "widgetType":
            suggest = "widget_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SentryDashboardWidget. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SentryDashboardWidget.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SentryDashboardWidget.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 display_type: str,
                 layout: 'outputs.SentryDashboardWidgetLayout',
                 queries: Sequence['outputs.SentryDashboardWidgetQuery'],
                 title: str,
                 id: Optional[str] = None,
                 interval: Optional[str] = None,
                 limit: Optional[int] = None,
                 widget_type: Optional[str] = None):
        """
        :param str id: The ID of this resource.
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
    def display_type(self) -> str:
        return pulumi.get(self, "display_type")

    @property
    @pulumi.getter
    def layout(self) -> 'outputs.SentryDashboardWidgetLayout':
        return pulumi.get(self, "layout")

    @property
    @pulumi.getter
    def queries(self) -> Sequence['outputs.SentryDashboardWidgetQuery']:
        return pulumi.get(self, "queries")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interval(self) -> Optional[str]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def limit(self) -> Optional[int]:
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter(name="widgetType")
    def widget_type(self) -> Optional[str]:
        return pulumi.get(self, "widget_type")


@pulumi.output_type
class SentryDashboardWidgetLayout(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "minH":
            suggest = "min_h"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SentryDashboardWidgetLayout. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SentryDashboardWidgetLayout.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SentryDashboardWidgetLayout.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 h: int,
                 min_h: int,
                 w: int,
                 x: int,
                 y: int):
        pulumi.set(__self__, "h", h)
        pulumi.set(__self__, "min_h", min_h)
        pulumi.set(__self__, "w", w)
        pulumi.set(__self__, "x", x)
        pulumi.set(__self__, "y", y)

    @property
    @pulumi.getter
    def h(self) -> int:
        return pulumi.get(self, "h")

    @property
    @pulumi.getter(name="minH")
    def min_h(self) -> int:
        return pulumi.get(self, "min_h")

    @property
    @pulumi.getter
    def w(self) -> int:
        return pulumi.get(self, "w")

    @property
    @pulumi.getter
    def x(self) -> int:
        return pulumi.get(self, "x")

    @property
    @pulumi.getter
    def y(self) -> int:
        return pulumi.get(self, "y")


@pulumi.output_type
class SentryDashboardWidgetQuery(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fieldAliases":
            suggest = "field_aliases"
        elif key == "orderBy":
            suggest = "order_by"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SentryDashboardWidgetQuery. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SentryDashboardWidgetQuery.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SentryDashboardWidgetQuery.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 aggregates: Optional[Sequence[str]] = None,
                 columns: Optional[Sequence[str]] = None,
                 conditions: Optional[str] = None,
                 field_aliases: Optional[Sequence[str]] = None,
                 fields: Optional[Sequence[str]] = None,
                 id: Optional[str] = None,
                 name: Optional[str] = None,
                 order_by: Optional[str] = None):
        """
        :param str id: The ID of this resource.
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
    def aggregates(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "aggregates")

    @property
    @pulumi.getter
    def columns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "columns")

    @property
    @pulumi.getter
    def conditions(self) -> Optional[str]:
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="fieldAliases")
    def field_aliases(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "field_aliases")

    @property
    @pulumi.getter
    def fields(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")


@pulumi.output_type
class SentryMetricAlertTrigger(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "alertThreshold":
            suggest = "alert_threshold"
        elif key == "thresholdType":
            suggest = "threshold_type"
        elif key == "resolveThreshold":
            suggest = "resolve_threshold"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SentryMetricAlertTrigger. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SentryMetricAlertTrigger.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SentryMetricAlertTrigger.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 alert_threshold: float,
                 label: str,
                 threshold_type: int,
                 actions: Optional[Sequence['outputs.SentryMetricAlertTriggerAction']] = None,
                 id: Optional[str] = None,
                 resolve_threshold: Optional[float] = None):
        """
        :param str id: The ID of this resource.
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
    def alert_threshold(self) -> float:
        return pulumi.get(self, "alert_threshold")

    @property
    @pulumi.getter
    def label(self) -> str:
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="thresholdType")
    def threshold_type(self) -> int:
        return pulumi.get(self, "threshold_type")

    @property
    @pulumi.getter
    def actions(self) -> Optional[Sequence['outputs.SentryMetricAlertTriggerAction']]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resolveThreshold")
    def resolve_threshold(self) -> Optional[float]:
        return pulumi.get(self, "resolve_threshold")


@pulumi.output_type
class SentryMetricAlertTriggerAction(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "targetType":
            suggest = "target_type"
        elif key == "integrationId":
            suggest = "integration_id"
        elif key == "targetIdentifier":
            suggest = "target_identifier"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SentryMetricAlertTriggerAction. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SentryMetricAlertTriggerAction.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SentryMetricAlertTriggerAction.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 target_type: str,
                 type: str,
                 id: Optional[str] = None,
                 integration_id: Optional[int] = None,
                 target_identifier: Optional[str] = None):
        """
        :param str id: The ID of this resource.
        """
        pulumi.set(__self__, "target_type", target_type)
        pulumi.set(__self__, "type", type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if integration_id is not None:
            pulumi.set(__self__, "integration_id", integration_id)
        if target_identifier is not None:
            pulumi.set(__self__, "target_identifier", target_identifier)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> str:
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> Optional[int]:
        return pulumi.get(self, "integration_id")

    @property
    @pulumi.getter(name="targetIdentifier")
    def target_identifier(self) -> Optional[str]:
        return pulumi.get(self, "target_identifier")


@pulumi.output_type
class GetSentryDashboardWidgetResult(dict):
    def __init__(__self__, *,
                 display_type: str,
                 id: str,
                 interval: str,
                 layouts: Sequence['outputs.GetSentryDashboardWidgetLayoutResult'],
                 limit: int,
                 queries: Sequence['outputs.GetSentryDashboardWidgetQueryResult'],
                 title: str,
                 widget_type: str):
        pulumi.set(__self__, "display_type", display_type)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "layouts", layouts)
        pulumi.set(__self__, "limit", limit)
        pulumi.set(__self__, "queries", queries)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "widget_type", widget_type)

    @property
    @pulumi.getter(name="displayType")
    def display_type(self) -> str:
        return pulumi.get(self, "display_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interval(self) -> str:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def layouts(self) -> Sequence['outputs.GetSentryDashboardWidgetLayoutResult']:
        return pulumi.get(self, "layouts")

    @property
    @pulumi.getter
    def limit(self) -> int:
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter
    def queries(self) -> Sequence['outputs.GetSentryDashboardWidgetQueryResult']:
        return pulumi.get(self, "queries")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="widgetType")
    def widget_type(self) -> str:
        return pulumi.get(self, "widget_type")


@pulumi.output_type
class GetSentryDashboardWidgetLayoutResult(dict):
    def __init__(__self__, *,
                 h: int,
                 min_h: int,
                 w: int,
                 x: int,
                 y: int):
        pulumi.set(__self__, "h", h)
        pulumi.set(__self__, "min_h", min_h)
        pulumi.set(__self__, "w", w)
        pulumi.set(__self__, "x", x)
        pulumi.set(__self__, "y", y)

    @property
    @pulumi.getter
    def h(self) -> int:
        return pulumi.get(self, "h")

    @property
    @pulumi.getter(name="minH")
    def min_h(self) -> int:
        return pulumi.get(self, "min_h")

    @property
    @pulumi.getter
    def w(self) -> int:
        return pulumi.get(self, "w")

    @property
    @pulumi.getter
    def x(self) -> int:
        return pulumi.get(self, "x")

    @property
    @pulumi.getter
    def y(self) -> int:
        return pulumi.get(self, "y")


@pulumi.output_type
class GetSentryDashboardWidgetQueryResult(dict):
    def __init__(__self__, *,
                 aggregates: Sequence[str],
                 columns: Sequence[str],
                 conditions: str,
                 field_aliases: Sequence[str],
                 fields: Sequence[str],
                 id: str,
                 name: str,
                 order_by: str):
        pulumi.set(__self__, "aggregates", aggregates)
        pulumi.set(__self__, "columns", columns)
        pulumi.set(__self__, "conditions", conditions)
        pulumi.set(__self__, "field_aliases", field_aliases)
        pulumi.set(__self__, "fields", fields)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "order_by", order_by)

    @property
    @pulumi.getter
    def aggregates(self) -> Sequence[str]:
        return pulumi.get(self, "aggregates")

    @property
    @pulumi.getter
    def columns(self) -> Sequence[str]:
        return pulumi.get(self, "columns")

    @property
    @pulumi.getter
    def conditions(self) -> str:
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="fieldAliases")
    def field_aliases(self) -> Sequence[str]:
        return pulumi.get(self, "field_aliases")

    @property
    @pulumi.getter
    def fields(self) -> Sequence[str]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> str:
        return pulumi.get(self, "order_by")


@pulumi.output_type
class GetSentryMetricAlertTriggerResult(dict):
    def __init__(__self__, *,
                 actions: Sequence['outputs.GetSentryMetricAlertTriggerActionResult'],
                 alert_threshold: float,
                 id: str,
                 label: str,
                 resolve_threshold: float,
                 threshold_type: int):
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "alert_threshold", alert_threshold)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "resolve_threshold", resolve_threshold)
        pulumi.set(__self__, "threshold_type", threshold_type)

    @property
    @pulumi.getter
    def actions(self) -> Sequence['outputs.GetSentryMetricAlertTriggerActionResult']:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="alertThreshold")
    def alert_threshold(self) -> float:
        return pulumi.get(self, "alert_threshold")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def label(self) -> str:
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="resolveThreshold")
    def resolve_threshold(self) -> float:
        return pulumi.get(self, "resolve_threshold")

    @property
    @pulumi.getter(name="thresholdType")
    def threshold_type(self) -> int:
        return pulumi.get(self, "threshold_type")


@pulumi.output_type
class GetSentryMetricAlertTriggerActionResult(dict):
    def __init__(__self__, *,
                 id: str,
                 integration_id: int,
                 target_identifier: str,
                 target_type: str,
                 type: str):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "integration_id", integration_id)
        pulumi.set(__self__, "target_identifier", target_identifier)
        pulumi.set(__self__, "target_type", target_type)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="integrationId")
    def integration_id(self) -> int:
        return pulumi.get(self, "integration_id")

    @property
    @pulumi.getter(name="targetIdentifier")
    def target_identifier(self) -> str:
        return pulumi.get(self, "target_identifier")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> str:
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


