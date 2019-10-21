package main

import (
	"net/http"
	"net/url"

	"github.com/riari/moni/account"
	"github.com/riari/moni/config"
	"github.com/riari/moni/monzo"
	"github.com/riari/moni/ping"
	"github.com/spf13/cobra"
)

func main() {
	conf := config.Initialise()

	client := monzo.Client{
		Client:      http.DefaultClient,
		BaseURL:     &url.URL{Scheme: "https", Host: "api.monzo.com"},
		AccessToken: conf.GetString("access_token"),
	}

	var rootCmd = &cobra.Command{Use: "moni"}
	rootCmd.AddCommand(account.Command())
	rootCmd.AddCommand(config.Command(conf))
	rootCmd.AddCommand(ping.Command(client))

	rootCmd.Execute()
}
