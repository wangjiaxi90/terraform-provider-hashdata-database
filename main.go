package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/wangjiaxi90/terraform-provider-hashdata-database/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

//var (
//	// these will be set by the goreleaser configuration
//	// to appropriate values for the compiled binary
//	version string = ""
//
//	// goreleaser can also pass the specific commit if you want
//	// commit  string = ""
//)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})

	//var debugMode bool
	//
	//flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	//flag.Parse()
	//
	//opts := &plugin.ServeOpts{ProviderFunc: provider.Provider()}
	//
	//if debugMode {
	//	// TODO: update this string with the full name of your provider as used in your configs
	//	err := plugin.Debug(context.Background(), "registry.terraform.io/hashicorp/scaffolding", opts)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//	return
	//}
	//
	//plugin.Serve(opts)
}
