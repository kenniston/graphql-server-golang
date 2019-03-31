package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command {
	Use: "run",
	Aliases: []string{"r"},
	Short: "Starts the GraphQL server, HTTP REST Microservices or Kafka based Microservices",
	Long: `The GraphQL server and the Services API server. The GraphQL 
server provides the query endpoints and the Services 
API provide the endpoints to save and update data.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Start a server based on flags

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	//web.AddFlags(serveCmd.Flags())

	err := viper.GetViper().BindPFlags(runCmd.Flags())
	if err != nil {
		panic(err)
	}
}
