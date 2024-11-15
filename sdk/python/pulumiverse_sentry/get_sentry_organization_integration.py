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

__all__ = [
    'GetSentryOrganizationIntegrationResult',
    'AwaitableGetSentryOrganizationIntegrationResult',
    'get_sentry_organization_integration',
    'get_sentry_organization_integration_output',
]

@pulumi.output_type
class GetSentryOrganizationIntegrationResult:
    """
    A collection of values returned by getSentryOrganizationIntegration.
    """
    def __init__(__self__, id=None, internal_id=None, name=None, organization=None, provider_key=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal_id and not isinstance(internal_id, str):
            raise TypeError("Expected argument 'internal_id' to be a str")
        pulumi.set(__self__, "internal_id", internal_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if provider_key and not isinstance(provider_key, str):
            raise TypeError("Expected argument 'provider_key' to be a str")
        pulumi.set(__self__, "provider_key", provider_key)

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
        The internal ID for this organization integration.
        """
        return pulumi.get(self, "internal_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the organization integration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> str:
        """
        The slug of the organization the integration belongs to.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="providerKey")
    def provider_key(self) -> str:
        """
        The key of the organization integration provider.
        """
        return pulumi.get(self, "provider_key")


class AwaitableGetSentryOrganizationIntegrationResult(GetSentryOrganizationIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSentryOrganizationIntegrationResult(
            id=self.id,
            internal_id=self.internal_id,
            name=self.name,
            organization=self.organization,
            provider_key=self.provider_key)


def get_sentry_organization_integration(name: Optional[str] = None,
                                        organization: Optional[str] = None,
                                        provider_key: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSentryOrganizationIntegrationResult:
    """
    Sentry Organization Integration data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_sentry as sentry

    # Retrieve a Github organization integration
    github = sentry.get_sentry_organization_integration(organization="my-organization",
        provider_key="github",
        name="my-github-organization")
    # Retrieve a Slack integration
    slack = sentry.get_sentry_organization_integration(organization="my-organization",
        provider_key="slack",
        name="Slack Workspace")
    ```


    :param str name: The name of the organization integration.
    :param str organization: The slug of the organization the integration belongs to.
    :param str provider_key: The key of the organization integration provider.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['organization'] = organization
    __args__['providerKey'] = provider_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sentry:index/getSentryOrganizationIntegration:getSentryOrganizationIntegration', __args__, opts=opts, typ=GetSentryOrganizationIntegrationResult).value

    return AwaitableGetSentryOrganizationIntegrationResult(
        id=pulumi.get(__ret__, 'id'),
        internal_id=pulumi.get(__ret__, 'internal_id'),
        name=pulumi.get(__ret__, 'name'),
        organization=pulumi.get(__ret__, 'organization'),
        provider_key=pulumi.get(__ret__, 'provider_key'))
def get_sentry_organization_integration_output(name: Optional[pulumi.Input[str]] = None,
                                               organization: Optional[pulumi.Input[str]] = None,
                                               provider_key: Optional[pulumi.Input[str]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSentryOrganizationIntegrationResult]:
    """
    Sentry Organization Integration data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_sentry as sentry

    # Retrieve a Github organization integration
    github = sentry.get_sentry_organization_integration(organization="my-organization",
        provider_key="github",
        name="my-github-organization")
    # Retrieve a Slack integration
    slack = sentry.get_sentry_organization_integration(organization="my-organization",
        provider_key="slack",
        name="Slack Workspace")
    ```


    :param str name: The name of the organization integration.
    :param str organization: The slug of the organization the integration belongs to.
    :param str provider_key: The key of the organization integration provider.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['organization'] = organization
    __args__['providerKey'] = provider_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('sentry:index/getSentryOrganizationIntegration:getSentryOrganizationIntegration', __args__, opts=opts, typ=GetSentryOrganizationIntegrationResult)
    return __ret__.apply(lambda __response__: GetSentryOrganizationIntegrationResult(
        id=pulumi.get(__response__, 'id'),
        internal_id=pulumi.get(__response__, 'internal_id'),
        name=pulumi.get(__response__, 'name'),
        organization=pulumi.get(__response__, 'organization'),
        provider_key=pulumi.get(__response__, 'provider_key')))
