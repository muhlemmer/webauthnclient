/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/muhlemmer/webauthnclient/client"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register this client with a WebAuthN server",
	Long: `Register this client with a WebAuthN server using
	attestation options.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.LoadClient(stateFile)
		if err != nil {
			return err
		}
		resp, err := c.CreateAttestationResponse(args[0])
		if err != nil {
			return err
		}
		fmt.Println(resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
