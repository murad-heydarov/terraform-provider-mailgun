package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/murad-heydarov/terraform-mailgun-provider/mailgun"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mailgun.Provider})
}
