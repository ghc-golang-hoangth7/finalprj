package common

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const ConfigPath = "../../../config/.env"

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	DbName           string
	PlanesHost       string
	PlanesPort       int
	FlightsHost      string
	FlightsPort      int
	GraphQLHost      string
	GraphQLPort      int
}

func LoadConfig() (*Config, error) {
	if os.Getenv("ENV_LOADED") != "true" {
		err := godotenv.Load(os.Args[1])
		if err != nil {
			return nil, fmt.Errorf("failed to load .env file: %w", err)
		}
	}

	config := &Config{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		DbName:           os.Getenv("DB_NAME"),
		PlanesHost:       os.Getenv("PLANES_SERVICE_HOST"),
		FlightsHost:      os.Getenv("FLIGHTS_SERVICE_HOST"),
		GraphQLHost:      os.Getenv("GRAPHQL_HOST"),
	}

	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid POSTGRES_PORT value: %w", err)
	}
	config.PostgresPort = postgresPort

	planesPort, err := strconv.Atoi(os.Getenv("PLANES_SERVICE_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid PLANES_SERVICE_PORT value: %w", err)
	}
	config.PlanesPort = planesPort

	flightsPort, err := strconv.Atoi(os.Getenv("FLIGHTS_SERVICE_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid FLIGHTS_SERVICE_PORT value: %w", err)
	}
	config.FlightsPort = flightsPort

	graphqlPort, err := strconv.Atoi(os.Getenv("GRAPHQL_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid GRAPHQL_PORT value: %w", err)
	}
	config.GraphQLPort = graphqlPort

	return config, nil
}

func (c *Config) GetGraphQLHAddr() string {
	return fmt.Sprintf("%v:%v", c.GraphQLHost, c.GraphQLPort)
}
func (c *Config) GetPlanesAddr() string {
	return fmt.Sprintf("%v:%v", c.PlanesHost, c.PlanesPort)
}
func (c *Config) GetFlightsAddr() string {
	return fmt.Sprintf("%v:%v", c.FlightsHost, c.FlightsPort)
}
