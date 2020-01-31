package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"ochk.pl/terraform_provider/ochk"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ochk.Provider,
	})
}