package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/ochk/terraform-provider-ochk/ochk"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ochk.Provider,
	})
}
