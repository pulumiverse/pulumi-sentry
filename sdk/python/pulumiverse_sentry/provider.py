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

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 base_url: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] base_url: The target Sentry Base API URL in the format `https://[hostname]/api/`. The default value is `https://sentry.io/api/`.
               The value must be provided when working with Sentry On-Premise. The value can be sourced from the `SENTRY_BASE_URL`
               environment variable.
        :param pulumi.Input[str] token: The authentication token used to connect to Sentry. The value can be sourced from the `SENTRY_AUTH_TOKEN` environment
               variable.
        """
        if base_url is None:
            base_url = _utilities.get_env('SENTRY_BASE_URL')
        if base_url is not None:
            pulumi.set(__self__, "base_url", base_url)
        if token is None:
            token = _utilities.get_env('SENTRY_TOKEN')
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> Optional[pulumi.Input[str]]:
        """
        The target Sentry Base API URL in the format `https://[hostname]/api/`. The default value is `https://sentry.io/api/`.
        The value must be provided when working with Sentry On-Premise. The value can be sourced from the `SENTRY_BASE_URL`
        environment variable.
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_url", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication token used to connect to Sentry. The value can be sourced from the `SENTRY_AUTH_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the sentry package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] base_url: The target Sentry Base API URL in the format `https://[hostname]/api/`. The default value is `https://sentry.io/api/`.
               The value must be provided when working with Sentry On-Premise. The value can be sourced from the `SENTRY_BASE_URL`
               environment variable.
        :param pulumi.Input[str] token: The authentication token used to connect to Sentry. The value can be sourced from the `SENTRY_AUTH_TOKEN` environment
               variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the sentry package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if base_url is None:
                base_url = _utilities.get_env('SENTRY_BASE_URL')
            __props__.__dict__["base_url"] = base_url
            if token is None:
                token = _utilities.get_env('SENTRY_TOKEN')
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'sentry',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> pulumi.Output[Optional[str]]:
        """
        The target Sentry Base API URL in the format `https://[hostname]/api/`. The default value is `https://sentry.io/api/`.
        The value must be provided when working with Sentry On-Premise. The value can be sourced from the `SENTRY_BASE_URL`
        environment variable.
        """
        return pulumi.get(self, "base_url")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional[str]]:
        """
        The authentication token used to connect to Sentry. The value can be sourced from the `SENTRY_AUTH_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "token")

