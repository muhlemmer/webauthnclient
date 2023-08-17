/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/muhlemmer/webauthnclient/client"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a WebAuthN client",
	Long: `Initialize a WebAuthN. The an RP, credentials and authenticator are
	created and stored to a JSON state file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.NewClient(rpName, rpDomain, rpOrigin).Store(stateFile)
	},
}

var (
	rpName   string
	rpDomain string
	rpOrigin string
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&rpName, "name", "ZITADEL", "Name of the WebAuthN RP")
	initCmd.Flags().StringVar(&rpDomain, "domain", "localhost", "Domain of the WebAuthN RP")
	initCmd.Flags().StringVar(&rpOrigin, "origin", "http://localhost:9000", "Origin of the WebAuthN RP")
}
