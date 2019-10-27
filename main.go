package main

import (
	"net/http"
	"net/url"

	"github.com/riari/moni/command"
	"github.com/riari/moni/config"
	"github.com/riari/moni/monzo"
	"github.com/spf13/cobra"
)

func main() {
	conf := config.Initialise()

	client := monzo.Client{
		Client:      http.DefaultClient,
		BaseURL:     &url.URL{Scheme: "https", Host: "api.monzo.com"},
		AccessToken: conf.GetString("access_token"),
	}

	client.Auth = &monzo.AuthService{Client: &client}

	var rootCmd = &cobra.Command{Use: "moni"}
	rootCmd.AddCommand(command.Account())
	rootCmd.AddCommand(command.Auth(client))
	rootCmd.AddCommand(command.Config(conf))

	rootCmd.Execute()
}
