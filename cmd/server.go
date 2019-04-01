package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"graphql-server/web/graphql"
	"log"
)

var runCmd = &cobra.Command {
	Use: "run",
	Aliases: []string{"r"},
	Short: "Starts the GraphQL server, HTTP REST Microservices or Kafka based Microservices",
	Long: `The GraphQL server and the Services API server. The GraphQL 
server provides the query endpoints and the Services 
API provide the endpoints to save and update data.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		serverType := cmd.Flag("server-type").Value.String()

		var builder graphql.IBuilder

		if serverType == "SERVICE" {
			log.Println("Starting server with Services...")
			builder = new(graphql.ServiceServerBuilder)
		} else if serverType == "KAFKA" {
			log.Println("Starting server with Kafka...")
			builder = new(graphql.KafkaServerBuilder)
		} else {
			log.Fatal("invalid server type")
		}

		manager := graphql.NewServerManager(builder)
		err := manager.CreateServer()

		if err != nil {
			log.Fatal(err)
		}

		server := builder.GetResult()

		return server.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("server-type", "t", viper.GetString("server-type"), "Configure server type as SERVICE or KAFKA")
	runCmd.Flags().StringP("server-port", "p", viper.GetString("server-port"), "Configure server port")

	err := viper.GetViper().BindPFlags(runCmd.Flags())
	if err != nil {
		panic(err)
	}
}
