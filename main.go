package main

import (
	"flag"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "registry.terraform.io/ochk/ochk",
		ProviderFunc: ochk.Provider,
	}

	plugin.Serve(opts)

}
