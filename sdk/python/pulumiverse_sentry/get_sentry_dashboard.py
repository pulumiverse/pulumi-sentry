# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetSentryDashboardResult',
    'AwaitableGetSentryDashboardResult',
    'get_sentry_dashboard',
    'get_sentry_dashboard_output',
]

@pulumi.output_type
class GetSentryDashboardResult:
    """
    A collection of values returned by getSentryDashboard.
    """
    def __init__(__self__, id=None, internal_id=None, organization=None, title=None, widgets=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal_id and not isinstance(internal_id, str):
            raise TypeError("Expected argument 'internal_id' to be a str")
        pulumi.set(__self__, "internal_id", internal_id)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if widgets and not isinstance(widgets, list):
            raise TypeError("Expected argument 'widgets' to be a list")
        pulumi.set(__self__, "widgets", widgets)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="internalId")
    def internal_id(self) -> str:
        """
        The internal ID for this dashboard.
        """
        return pulumi.get(self, "internal_id")

    @property
    @pulumi.getter
    def organization(self) -> str:
        """
        The slug of the organization the dashboard belongs to.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        Dashboard title.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def widgets(self) -> Sequence['outputs.GetSentryDashboardWidgetResult']:
        """
        Dashboard widgets.
        """
        return pulumi.get(self, "widgets")


class AwaitableGetSentryDashboardResult(GetSentryDashboardResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSentryDashboardResult(
            id=self.id,
            internal_id=self.internal_id,
            organization=self.organization,
            title=self.title,
            widgets=self.widgets)


def get_sentry_dashboard(internal_id: Optional[str] = None,
                         organization: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSentryDashboardResult:
    """
    Use this data source to access information about an existing resource.

    :param str internal_id: The internal ID for this dashboard.
    :param str organization: The slug of the organization the dashboard belongs to.
    """
    __args__ = dict()
    __args__['internalId'] = internal_id
    __args__['organization'] = organization
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sentry:index/getSentryDashboard:getSentryDashboard', __args__, opts=opts, typ=GetSentryDashboardResult).value

    return AwaitableGetSentryDashboardResult(
        id=pulumi.get(__ret__, 'id'),
        internal_id=pulumi.get(__ret__, 'internal_id'),
        organization=pulumi.get(__ret__, 'organization'),
        title=pulumi.get(__ret__, 'title'),
        widgets=pulumi.get(__ret__, 'widgets'))


@_utilities.lift_output_func(get_sentry_dashboard)
def get_sentry_dashboard_output(internal_id: Optional[pulumi.Input[str]] = None,
                                organization: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSentryDashboardResult]:
    """
    Use this data source to access information about an existing resource.

    :param str internal_id: The internal ID for this dashboard.
    :param str organization: The slug of the organization the dashboard belongs to.
    """
    ...
