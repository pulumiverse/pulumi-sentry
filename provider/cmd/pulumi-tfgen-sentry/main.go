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

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	sentry "github.com/pulumiverse/pulumi-sentry/provider"
	"github.com/pulumiverse/pulumi-sentry/provider/pkg/version"
)

var pluginName = "sentry"

func main() {
	// Modify the path to point to the new provider
	tfgen.Main("sentry", version.Version, sentry.Provider())

	fixSchemaFile()
	fixPackageJson()
}

func fixSchemaFile() {
	filePath := "./provider/cmd/pulumi-resource-sentry/schema.json"
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var packageSpec schema.PackageSpec
	err = json.Unmarshal(fileContents, &packageSpec)
	if err != nil {
		log.Fatalf("cannot deserialize schema.PackageSpec: %v", err)
	}

	var info nodejs.NodePackageInfo
	err = json.Unmarshal(packageSpec.Language["nodejs"], &info)
	if err != nil {
		log.Fatalf("cannot deserialize nodejs.NodePackageInfo: %v", err)
	}

	info.PluginName = pluginName

	serializedInfo, err := json.Marshal(info)
	if err != nil {
		log.Fatalf("cannot serialize nodejs.NodePackageInfo: %v", err)
	}

	packageSpec.Language["nodejs"] = serializedInfo

	fileContents, err = json.MarshalIndent(packageSpec, "", "    ")
	if err != nil {
		log.Fatalf("cannot serialize schema.PackageSpec: %v", err)
	}

	err = ioutil.WriteFile(filePath, fileContents, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func fixPackageJson() {
	// stolen from https://github.com/pulumi/pulumi/blob/v3.34.1/pkg/codegen/nodejs/gen.go#L2093
	// should be kept in sync or hopefully deleted in the future
	type npmPackage struct {
		Name             string                  `json:"name"`
		Version          string                  `json:"version"`
		Description      string                  `json:"description,omitempty"`
		Keywords         []string                `json:"keywords,omitempty"`
		Homepage         string                  `json:"homepage,omitempty"`
		Repository       string                  `json:"repository,omitempty"`
		License          string                  `json:"license,omitempty"`
		Scripts          map[string]string       `json:"scripts,omitempty"`
		Dependencies     map[string]string       `json:"dependencies,omitempty"`
		DevDependencies  map[string]string       `json:"devDependencies,omitempty"`
		PeerDependencies map[string]string       `json:"peerDependencies,omitempty"`
		Resolutions      map[string]string       `json:"resolutions,omitempty"`
		Pulumi           plugin.PulumiPluginJSON `json:"pulumi,omitempty"`
	}

	filePath := "sdk/nodejs/package.json"
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var packageJson npmPackage
	err = json.Unmarshal(fileContents, &packageJson)
	if err != nil {
		log.Fatalf("cannot deserialize npmPackage: %v", err)
	}

	packageJson.Pulumi.Name = pluginName

	fileContents, err = json.MarshalIndent(packageJson, "", "    ")
	if err != nil {
		log.Fatalf("cannot serialize schema.packageJson: %v", err)
	}

	err = ioutil.WriteFile(filePath, fileContents, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
