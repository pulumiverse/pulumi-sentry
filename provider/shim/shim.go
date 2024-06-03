package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	pf "github.com/jianyuan/terraform-provider-sentry/internal/provider"
	sdk "github.com/jianyuan/terraform-provider-sentry/sentry"
)

func SDKProvider(version string) *schema.Provider {
	return sdk.NewProvider(version)()
}

func PFProvider(version string) provider.Provider {
	return pf.New(version)()
}
