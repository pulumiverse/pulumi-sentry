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
from ._inputs import *

__all__ = ['SentryDashboardArgs', 'SentryDashboard']

@pulumi.input_type
class SentryDashboardArgs:
    def __init__(__self__, *,
                 organization: pulumi.Input[str],
                 title: pulumi.Input[str],
                 widgets: Optional[pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetArgs']]]] = None):
        """
        The set of arguments for constructing a SentryDashboard resource.
        :param pulumi.Input[str] organization: The slug of the organization the dashboard belongs to.
        :param pulumi.Input[str] title: Dashboard title.
        :param pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetArgs']]] widgets: Dashboard widgets.
        """
        pulumi.set(__self__, "organization", organization)
        pulumi.set(__self__, "title", title)
        if widgets is not None:
            pulumi.set(__self__, "widgets", widgets)

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Input[str]:
        """
        The slug of the organization the dashboard belongs to.
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Dashboard title.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def widgets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetArgs']]]]:
        """
        Dashboard widgets.
        """
        return pulumi.get(self, "widgets")

    @widgets.setter
    def widgets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetArgs']]]]):
        pulumi.set(self, "widgets", value)


@pulumi.input_type
class _SentryDashboardState:
    def __init__(__self__, *,
                 internal_id: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 widgets: Optional[pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetArgs']]]] = None):
        """
        Input properties used for looking up and filtering SentryDashboard resources.
        :param pulumi.Input[str] internal_id: The internal ID for this dashboard.
        :param pulumi.Input[str] organization: The slug of the organization the dashboard belongs to.
        :param pulumi.Input[str] title: Dashboard title.
        :param pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetArgs']]] widgets: Dashboard widgets.
        """
        if internal_id is not None:
            pulumi.set(__self__, "internal_id", internal_id)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if widgets is not None:
            pulumi.set(__self__, "widgets", widgets)

    @property
    @pulumi.getter(name="internalId")
    def internal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The internal ID for this dashboard.
        """
        return pulumi.get(self, "internal_id")

    @internal_id.setter
    def internal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_id", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[str]]:
        """
        The slug of the organization the dashboard belongs to.
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Dashboard title.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def widgets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetArgs']]]]:
        """
        Dashboard widgets.
        """
        return pulumi.get(self, "widgets")

    @widgets.setter
    def widgets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SentryDashboardWidgetArgs']]]]):
        pulumi.set(self, "widgets", value)


class SentryDashboard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 widgets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SentryDashboardWidgetArgs', 'SentryDashboardWidgetArgsDict']]]]] = None,
                 __props__=None):
        """
        Sentry Dashboard resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_sentry as sentry

        main = sentry.SentryDashboard("main",
            organization=main_sentry_organization["id"],
            title="Test dashboard",
            widgets=[
                {
                    "title": "Number of Errors",
                    "display_type": "big_number",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["count()"],
                        "aggregates": ["count()"],
                        "conditions": "!event.type:transaction",
                        "order_by": "count()",
                    }],
                    "layout": {
                        "x": 0,
                        "y": 0,
                        "w": 1,
                        "h": 1,
                        "min_h": 1,
                    },
                },
                {
                    "title": "Number of Issues",
                    "display_type": "big_number",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["count_unique(issue)"],
                        "aggregates": ["count_unique(issue)"],
                        "conditions": "!event.type:transaction",
                        "order_by": "count_unique(issue)",
                    }],
                    "layout": {
                        "x": 1,
                        "y": 0,
                        "w": 1,
                        "h": 1,
                        "min_h": 1,
                    },
                },
                {
                    "title": "Events",
                    "display_type": "line",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "name": "Events",
                        "fields": ["count()"],
                        "aggregates": ["count()"],
                        "conditions": "!event.type:transaction",
                        "order_by": "count()",
                    }],
                    "layout": {
                        "x": 2,
                        "y": 0,
                        "w": 4,
                        "h": 2,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Affected Users",
                    "display_type": "line",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [
                        {
                            "name": "Known Users",
                            "fields": ["count_unique(user)"],
                            "aggregates": ["count_unique(user)"],
                            "conditions": "has:user.email !event.type:transaction",
                            "order_by": "count_unique(user)",
                        },
                        {
                            "name": "Anonymous Users",
                            "fields": ["count_unique(user)"],
                            "aggregates": ["count_unique(user)"],
                            "conditions": "!has:user.email !event.type:transaction",
                            "order_by": "count_unique(user)",
                        },
                    ],
                    "layout": {
                        "x": 1,
                        "y": 2,
                        "w": 1,
                        "h": 2,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Handled vs. Unhandled",
                    "display_type": "line",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [
                        {
                            "name": "Handled",
                            "fields": ["count()"],
                            "aggregates": ["count()"],
                            "conditions": "error.handled:true",
                            "order_by": "count()",
                        },
                        {
                            "name": "Unhandled",
                            "fields": ["count()"],
                            "aggregates": ["count()"],
                            "conditions": "error.handled:false",
                            "order_by": "count()",
                        },
                    ],
                    "layout": {
                        "x": 0,
                        "y": 2,
                        "w": 1,
                        "h": 2,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Errors by Country",
                    "display_type": "world_map",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["count()"],
                        "aggregates": ["count()"],
                        "conditions": "!event.type:transaction has:geo.country_code",
                        "order_by": "count()",
                    }],
                    "layout": {
                        "x": 4,
                        "y": 6,
                        "w": 2,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "High Throughput Transactions",
                    "display_type": "table",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "count()",
                            "transaction",
                        ],
                        "aggregates": ["count()"],
                        "columns": ["transaction"],
                        "conditions": "!event.type:error",
                        "order_by": "-count()",
                    }],
                    "layout": {
                        "x": 0,
                        "y": 6,
                        "w": 2,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Errors by Browser",
                    "display_type": "table",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "browser.name",
                            "count()",
                        ],
                        "aggregates": ["count()"],
                        "columns": ["browser.name"],
                        "conditions": "!event.type:transaction has:browser.name",
                        "order_by": "-count()",
                    }],
                    "layout": {
                        "x": 5,
                        "y": 2,
                        "w": 1,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Overall User Misery",
                    "display_type": "big_number",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["user_misery(300)"],
                        "aggregates": ["user_misery(300)"],
                    }],
                    "layout": {
                        "x": 0,
                        "y": 1,
                        "w": 1,
                        "h": 1,
                        "min_h": 1,
                    },
                },
                {
                    "title": "Overall Apdex",
                    "display_type": "big_number",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["apdex(300)"],
                        "aggregates": ["apdex(300)"],
                    }],
                    "layout": {
                        "x": 1,
                        "y": 1,
                        "w": 1,
                        "h": 1,
                        "min_h": 1,
                    },
                },
                {
                    "title": "High Throughput Transactions",
                    "display_type": "top_n",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "transaction",
                            "count()",
                        ],
                        "aggregates": ["count()"],
                        "columns": ["transaction"],
                        "conditions": "!event.type:error",
                        "order_by": "-count()",
                    }],
                    "layout": {
                        "x": 0,
                        "y": 4,
                        "w": 2,
                        "h": 2,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Issues Assigned to Me or My Teams",
                    "display_type": "table",
                    "interval": "5m",
                    "widget_type": "issue",
                    "queries": [{
                        "fields": [
                            "assignee",
                            "issue",
                            "title",
                        ],
                        "columns": [
                            "assignee",
                            "issue",
                            "title",
                        ],
                        "conditions": "assigned_or_suggested:me is:unresolved",
                        "order_by": "priority",
                    }],
                    "layout": {
                        "x": 2,
                        "y": 2,
                        "w": 2,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Transactions Ordered by Misery",
                    "display_type": "table",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "transaction",
                            "user_misery(300)",
                        ],
                        "aggregates": ["user_misery(300)"],
                        "columns": ["transaction"],
                        "order_by": "-user_misery(300)",
                    }],
                    "layout": {
                        "x": 2,
                        "y": 6,
                        "w": 2,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Errors by Browser Over Time",
                    "display_type": "top_n",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "browser.name",
                            "count()",
                        ],
                        "aggregates": ["count()"],
                        "columns": ["browser.name"],
                        "conditions": "event.type:error has:browser.name",
                        "order_by": "-count()",
                    }],
                    "layout": {
                        "x": 4,
                        "y": 2,
                        "w": 1,
                        "h": 4,
                        "min_h": 2,
                    },
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] organization: The slug of the organization the dashboard belongs to.
        :param pulumi.Input[str] title: Dashboard title.
        :param pulumi.Input[Sequence[pulumi.Input[Union['SentryDashboardWidgetArgs', 'SentryDashboardWidgetArgsDict']]]] widgets: Dashboard widgets.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SentryDashboardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Sentry Dashboard resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_sentry as sentry

        main = sentry.SentryDashboard("main",
            organization=main_sentry_organization["id"],
            title="Test dashboard",
            widgets=[
                {
                    "title": "Number of Errors",
                    "display_type": "big_number",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["count()"],
                        "aggregates": ["count()"],
                        "conditions": "!event.type:transaction",
                        "order_by": "count()",
                    }],
                    "layout": {
                        "x": 0,
                        "y": 0,
                        "w": 1,
                        "h": 1,
                        "min_h": 1,
                    },
                },
                {
                    "title": "Number of Issues",
                    "display_type": "big_number",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["count_unique(issue)"],
                        "aggregates": ["count_unique(issue)"],
                        "conditions": "!event.type:transaction",
                        "order_by": "count_unique(issue)",
                    }],
                    "layout": {
                        "x": 1,
                        "y": 0,
                        "w": 1,
                        "h": 1,
                        "min_h": 1,
                    },
                },
                {
                    "title": "Events",
                    "display_type": "line",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "name": "Events",
                        "fields": ["count()"],
                        "aggregates": ["count()"],
                        "conditions": "!event.type:transaction",
                        "order_by": "count()",
                    }],
                    "layout": {
                        "x": 2,
                        "y": 0,
                        "w": 4,
                        "h": 2,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Affected Users",
                    "display_type": "line",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [
                        {
                            "name": "Known Users",
                            "fields": ["count_unique(user)"],
                            "aggregates": ["count_unique(user)"],
                            "conditions": "has:user.email !event.type:transaction",
                            "order_by": "count_unique(user)",
                        },
                        {
                            "name": "Anonymous Users",
                            "fields": ["count_unique(user)"],
                            "aggregates": ["count_unique(user)"],
                            "conditions": "!has:user.email !event.type:transaction",
                            "order_by": "count_unique(user)",
                        },
                    ],
                    "layout": {
                        "x": 1,
                        "y": 2,
                        "w": 1,
                        "h": 2,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Handled vs. Unhandled",
                    "display_type": "line",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [
                        {
                            "name": "Handled",
                            "fields": ["count()"],
                            "aggregates": ["count()"],
                            "conditions": "error.handled:true",
                            "order_by": "count()",
                        },
                        {
                            "name": "Unhandled",
                            "fields": ["count()"],
                            "aggregates": ["count()"],
                            "conditions": "error.handled:false",
                            "order_by": "count()",
                        },
                    ],
                    "layout": {
                        "x": 0,
                        "y": 2,
                        "w": 1,
                        "h": 2,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Errors by Country",
                    "display_type": "world_map",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["count()"],
                        "aggregates": ["count()"],
                        "conditions": "!event.type:transaction has:geo.country_code",
                        "order_by": "count()",
                    }],
                    "layout": {
                        "x": 4,
                        "y": 6,
                        "w": 2,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "High Throughput Transactions",
                    "display_type": "table",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "count()",
                            "transaction",
                        ],
                        "aggregates": ["count()"],
                        "columns": ["transaction"],
                        "conditions": "!event.type:error",
                        "order_by": "-count()",
                    }],
                    "layout": {
                        "x": 0,
                        "y": 6,
                        "w": 2,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Errors by Browser",
                    "display_type": "table",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "browser.name",
                            "count()",
                        ],
                        "aggregates": ["count()"],
                        "columns": ["browser.name"],
                        "conditions": "!event.type:transaction has:browser.name",
                        "order_by": "-count()",
                    }],
                    "layout": {
                        "x": 5,
                        "y": 2,
                        "w": 1,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Overall User Misery",
                    "display_type": "big_number",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["user_misery(300)"],
                        "aggregates": ["user_misery(300)"],
                    }],
                    "layout": {
                        "x": 0,
                        "y": 1,
                        "w": 1,
                        "h": 1,
                        "min_h": 1,
                    },
                },
                {
                    "title": "Overall Apdex",
                    "display_type": "big_number",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": ["apdex(300)"],
                        "aggregates": ["apdex(300)"],
                    }],
                    "layout": {
                        "x": 1,
                        "y": 1,
                        "w": 1,
                        "h": 1,
                        "min_h": 1,
                    },
                },
                {
                    "title": "High Throughput Transactions",
                    "display_type": "top_n",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "transaction",
                            "count()",
                        ],
                        "aggregates": ["count()"],
                        "columns": ["transaction"],
                        "conditions": "!event.type:error",
                        "order_by": "-count()",
                    }],
                    "layout": {
                        "x": 0,
                        "y": 4,
                        "w": 2,
                        "h": 2,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Issues Assigned to Me or My Teams",
                    "display_type": "table",
                    "interval": "5m",
                    "widget_type": "issue",
                    "queries": [{
                        "fields": [
                            "assignee",
                            "issue",
                            "title",
                        ],
                        "columns": [
                            "assignee",
                            "issue",
                            "title",
                        ],
                        "conditions": "assigned_or_suggested:me is:unresolved",
                        "order_by": "priority",
                    }],
                    "layout": {
                        "x": 2,
                        "y": 2,
                        "w": 2,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Transactions Ordered by Misery",
                    "display_type": "table",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "transaction",
                            "user_misery(300)",
                        ],
                        "aggregates": ["user_misery(300)"],
                        "columns": ["transaction"],
                        "order_by": "-user_misery(300)",
                    }],
                    "layout": {
                        "x": 2,
                        "y": 6,
                        "w": 2,
                        "h": 4,
                        "min_h": 2,
                    },
                },
                {
                    "title": "Errors by Browser Over Time",
                    "display_type": "top_n",
                    "interval": "5m",
                    "widget_type": "discover",
                    "queries": [{
                        "fields": [
                            "browser.name",
                            "count()",
                        ],
                        "aggregates": ["count()"],
                        "columns": ["browser.name"],
                        "conditions": "event.type:error has:browser.name",
                        "order_by": "-count()",
                    }],
                    "layout": {
                        "x": 4,
                        "y": 2,
                        "w": 1,
                        "h": 4,
                        "min_h": 2,
                    },
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param SentryDashboardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SentryDashboardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 widgets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SentryDashboardWidgetArgs', 'SentryDashboardWidgetArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SentryDashboardArgs.__new__(SentryDashboardArgs)

            if organization is None and not opts.urn:
                raise TypeError("Missing required property 'organization'")
            __props__.__dict__["organization"] = organization
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            __props__.__dict__["widgets"] = widgets
            __props__.__dict__["internal_id"] = None
        super(SentryDashboard, __self__).__init__(
            'sentry:index/sentryDashboard:SentryDashboard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            internal_id: Optional[pulumi.Input[str]] = None,
            organization: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None,
            widgets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SentryDashboardWidgetArgs', 'SentryDashboardWidgetArgsDict']]]]] = None) -> 'SentryDashboard':
        """
        Get an existing SentryDashboard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] internal_id: The internal ID for this dashboard.
        :param pulumi.Input[str] organization: The slug of the organization the dashboard belongs to.
        :param pulumi.Input[str] title: Dashboard title.
        :param pulumi.Input[Sequence[pulumi.Input[Union['SentryDashboardWidgetArgs', 'SentryDashboardWidgetArgsDict']]]] widgets: Dashboard widgets.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SentryDashboardState.__new__(_SentryDashboardState)

        __props__.__dict__["internal_id"] = internal_id
        __props__.__dict__["organization"] = organization
        __props__.__dict__["title"] = title
        __props__.__dict__["widgets"] = widgets
        return SentryDashboard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="internalId")
    def internal_id(self) -> pulumi.Output[str]:
        """
        The internal ID for this dashboard.
        """
        return pulumi.get(self, "internal_id")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[str]:
        """
        The slug of the organization the dashboard belongs to.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        Dashboard title.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def widgets(self) -> pulumi.Output[Optional[Sequence['outputs.SentryDashboardWidget']]]:
        """
        Dashboard widgets.
        """
        return pulumi.get(self, "widgets")

