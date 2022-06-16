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
	sentry "github.com/pulumiverse/pulumi-sentry/provider"
	"github.com/pulumiverse/pulumi-sentry/provider/pkg/version"
)

func main() {
	// Modify the path to point to the new provider
	tfgen.Main("sentry", version.Version, sentry.Provider())

	fixSchemaFile()
}

func fixSchemaFile() {
	fileContents, err := ioutil.ReadFile("./provider/cmd/pulumi-resource-sentry/schema.json")
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

	info.PluginName = "sentry"

	serializedInfo, err := json.Marshal(info)
	if err != nil {
		log.Fatalf("cannot serialize nodejs.NodePackageInfo: %v", err)
	}

	packageSpec.Language["nodejs"] = serializedInfo

	fileContents, err = json.MarshalIndent(packageSpec, "", "    ")
	if err != nil {
		log.Fatalf("cannot serialize schema.PackageSpec: %v", err)
	}

	err = ioutil.WriteFile("./provider/cmd/pulumi-resource-sentry/schema.json", fileContents, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
