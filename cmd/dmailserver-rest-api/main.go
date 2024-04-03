package main

import (
	"errors"
	"log"
	"log/slog"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/pjotrscholtze/dmailserver-rest-api/cmd/dmailserver-rest-api/cnf"
	"github.com/pjotrscholtze/dmailserver-rest-api/cmd/dmailserver-rest-api/controller"
	"github.com/pjotrscholtze/dmailserver-rest-api/cmd/dmailserver-rest-api/repo"
	"github.com/pjotrscholtze/dmailserver-rest-api/models"
	"github.com/pjotrscholtze/dmailserver-rest-api/restapi"
	"github.com/pjotrscholtze/dmailserver-rest-api/restapi/operations"
)

func main() {
	configPath := "../../config/config.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	config, err := cnf.GetConfig(configPath)
	if err != nil {
		panic(err)
	}
	_ = config

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("Starting DMailserver-rest-api")

	slog.Info("Setting command prefix", "commandPrefix", config.ServerConfig.CommandPrefix)
	sr := repo.NewSetupRepo(config.ServerConfig.CommandPrefix)

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewDmailserverRestAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Swagger Petstore - OpenAPI 3.0"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	// Applies when the "x-token" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		if token == config.ServerConfig.APIKey { //"abcdefuvwxyz" {
			prin := models.Principal(token)
			return &prin, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", token)
		return nil, errors.New("incorrect api key auth")
	}
	controller.SetupController(api, sr)
	server.Port = config.ServerConfig.Port
	server.Host = config.ServerConfig.Host
	slog.Info("Setting up port and host for connection listening", "port", server.Port, "host", server.Host)
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
