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
    'GetSentryAllProjectsResult',
    'AwaitableGetSentryAllProjectsResult',
    'get_sentry_all_projects',
    'get_sentry_all_projects_output',
]

@pulumi.output_type
class GetSentryAllProjectsResult:
    """
    A collection of values returned by getSentryAllProjects.
    """
    def __init__(__self__, id=None, organization=None, project_slugs=None, projects=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if project_slugs and not isinstance(project_slugs, list):
            raise TypeError("Expected argument 'project_slugs' to be a list")
        pulumi.set(__self__, "project_slugs", project_slugs)
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        pulumi.set(__self__, "projects", projects)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def organization(self) -> str:
        """
        The slug of the organization the resource belongs to.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="projectSlugs")
    def project_slugs(self) -> Sequence[str]:
        """
        The slugs of the projects.
        """
        return pulumi.get(self, "project_slugs")

    @property
    @pulumi.getter
    def projects(self) -> Sequence['outputs.GetSentryAllProjectsProjectResult']:
        """
        The list of projects.
        """
        return pulumi.get(self, "projects")


class AwaitableGetSentryAllProjectsResult(GetSentryAllProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSentryAllProjectsResult(
            id=self.id,
            organization=self.organization,
            project_slugs=self.project_slugs,
            projects=self.projects)


def get_sentry_all_projects(organization: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSentryAllProjectsResult:
    """
    Return a list of projects available to the authenticated session.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_sentry as sentry

    default = sentry.get_sentry_all_projects(organization="my-organization")
    ```


    :param str organization: The slug of the organization the resource belongs to.
    """
    __args__ = dict()
    __args__['organization'] = organization
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sentry:index/getSentryAllProjects:getSentryAllProjects', __args__, opts=opts, typ=GetSentryAllProjectsResult).value

    return AwaitableGetSentryAllProjectsResult(
        id=pulumi.get(__ret__, 'id'),
        organization=pulumi.get(__ret__, 'organization'),
        project_slugs=pulumi.get(__ret__, 'project_slugs'),
        projects=pulumi.get(__ret__, 'projects'))


@_utilities.lift_output_func(get_sentry_all_projects)
def get_sentry_all_projects_output(organization: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSentryAllProjectsResult]:
    """
    Return a list of projects available to the authenticated session.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_sentry as sentry

    default = sentry.get_sentry_all_projects(organization="my-organization")
    ```


    :param str organization: The slug of the organization the resource belongs to.
    """
    ...
