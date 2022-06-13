#!/bin/bash

make tfgen
make build_sdks
make install_nodejs_sdk

rm -f ~/.pulumi/bin/pulumi-resource-sentry
cp ./bin/pulumi-resource-sentry ~/.pulumi/bin/pulumi-resource-sentry

cd examples/my-example/ts
yarn link "@pulumiverse/sentry"
