package main

import (
	"context"
	"flag"
	"log"
	"terraform-provider-convcase/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var version = ""

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "hashicorp.com/kasaikou/convcase",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.NewProvider(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
