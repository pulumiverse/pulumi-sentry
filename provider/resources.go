// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sentry

import (
	"fmt"
	"path/filepath"

	"github.com/jianyuan/terraform-provider-sentry/sentry"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumiverse/pulumi-sentry/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "sentry"
	// modules:
	mainMod = "index" // the sentry module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(sentry.NewProvider(version.Version)())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "sentry",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Sentry",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/pulumiverse/pulumi-sentry/main/sentry.svg",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Description:       "A Pulumi package for creating and managing Sentry resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "sentry", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/pulumiverse",
		Repository: "https://github.com/pulumiverse/pulumi-sentry",
		// The GitHub Org for the provider - defaults to `terraform-providers`
		GitHubOrg: "jianyuan",
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SENTRY_TOKEN"},
				},
				Secret: boolRef(true),
			},
			"base_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SENTRY_BASE_URL"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"sentry_key": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryKey"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"dsn_secret": {Secret: boolRef(true)},
					"secret":     {Secret: boolRef(true)},
				},
			},
			"sentry_organization":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryOrganization")},
			"sentry_organization_code_mapping":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryOrganizationCodeMapping")},
			"sentry_organization_member":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryOrganizationMember")},
			"sentry_organization_repository_github": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryOrganizationRepositoryGithub")},
			"sentry_project":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryProject")},
			"sentry_plugin":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryPlugin")},
			"sentry_rule":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryRule")},
			"sentry_team":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryTeam")},
			"sentry_dashboard":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryDashboard")},
			"sentry_issue_alert":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryIssueAlert")},
			"sentry_metric_alert":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SentryMetricAlert")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"sentry_key": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSentryKey"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"dsn_secret": {Secret: boolRef(true)},
					"secret":     {Secret: boolRef(true)},
				},
			},
			"sentry_organization":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSentryOrganization")},
			"sentry_organization_integration": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSentryOrganizationIntegration")},
			"sentry_team":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSentryTeam")},
			"sentry_dashboard":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSentryDashboard")},
			"sentry_issue_alert":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSentryIssueAlert")},
			"sentry_metric_alert":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSentryMetricAlert")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/sentry",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PackageName:          "pulumiverse_sentry",
			PyProject:            struct{ Enabled bool }{true},
			RespectSchemaVersion: true,
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RootNamespace:        "Pulumiverse",
			RespectSchemaVersion: true,
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
