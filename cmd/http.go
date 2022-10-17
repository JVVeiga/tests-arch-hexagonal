/*
Copyright © 2022 João Veiga
*/
package cmd

import (
	"fmt"
	server2 "github.com/jvveiga/tests-arch-hexagonal/adapters/web/server"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start webserver",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebserver()
		server.Service = &productService
		fmt.Println("Webserver has been started")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
