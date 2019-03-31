package graphql

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type ServerType int
const (
	ServiceServerType 	ServerType = 0
	KafkaServerType   	ServerType = 1
)

var (
	// ErrInvalidServerType is returned when a invalid server type is given to
	// Server Manager.
	ErrInvalidServerType = errors.New("invalid server type")
)

// Server Manager (aka Server Director) manage the creation of GraphQL Server.
// The server created by the Service Manager's builder can be linked to
// Services or to the Kafka
//
//   Usage:
//
//      builder := new graphql.Builder
//      manager := NewServerManager(builder)
//
//   Create the service with Microservices
//
//		manager.CreateServer(ServiceServerType)
//		server := builder.GetResult()
//
//   Or, create the server with Kafka
//
//		manager.CreateServer(KafkaServerType)
//		server := builder.GetResult()
//
type ServerManager struct {
	// Server builder used to create a new GraphQL Server
	builder Builder
}

// NewServerManager creates a new ServerManager with a concrete Server Builder
func NewServerManager(builder Builder) ServerManager {
	return ServerManager{builder}
}

// Construct builds the product from a series of steps.
func (d *ServerManager) CreateServer(serverType ServerType) (err error) {
	config := viper.GetViper()
	if serverType == ServiceServerType || serverType == KafkaServerType {
		d.builder.Build(config)
	} else {
		return ErrInvalidServerType
	}

	return nil
}

// Builder is a interface for building
//
type Builder interface {
	// Build creates a new server with viper configs
	Build(v *viper.Viper)

	// GetResult return a configured server
	GetResult() Server
}

// QLServerBuilder create a new server using Services endpoints
// to resolver GraphQL Queries and Mutations
//
type ServiceServerBuilder struct {
	// If true indicates that server has been completely created and configured
	built bool

	// Server pointer to a configured server
	server *Server
}

// CreateServer builds the GraphQL Server from a series of steps.
func (s *ServiceServerBuilder) Build(v *viper.Viper) {
	if s.built { return }

	// Builds the server

	s.built = true
}

// GetResult returns the Server which has been build during the Build step.
func (s *ServiceServerBuilder) GetResult() *Server {
	return s.server
}

// KafkaServerBuilder create a new server using Kafka Topics
// to resolver GraphQL Queries and Mutations
//
type KafkaServerBuilder struct {
	// If true indicates that server has been completely created and configured
	built bool

	// Server pointer to a configured server
	server *Server
}

// CreateServer builds the GraphQL Server from a series of steps.
func (s *KafkaServerBuilder) Build(v *viper.Viper) {
	if s.built { return }

	// Builds the server

	s.built = true
}

// GetResult returns the Server which has been build during the Build step.
func (s *KafkaServerBuilder) GetResult() *Server {
	return s.server
}

// Server is a interface for a GraphQL Server
//
//
type Server interface {
	Run() error
}

// ServiceServer is a server which uses Microservices endpoints
// to resolver GraphQL Queries and Mutations.
//
// This server has a set of Microservices endpoints which are
// called through HTTP protocol to resolve GraphQL queries and
// mutations.
//
// The Microservices on this server mode communicate directly
// with each other.
//
type ServiceServer struct {
	Service
}

func (s *ServiceServer) Run() error {
	fmt.Println("GraphQL Server Running with Microservices...")
	return nil
}

type Service struct {
	Endpoint string
	Port string
}

// KafkaServer is a server which uses Kafka Topics to resolver
// GraphQL Queries and Mutations.
//
// This server has a set of Kafka Topics to resolve GraphQL queries
// and mutations.
//
// The Microservices on this server mode communicate with each other
// through Kafka.
//
type KafkaServer struct {

}

func (s *KafkaServer) Run() error {
	fmt.Println("GraphQL Server Running with Kafka...")
	return nil
}
