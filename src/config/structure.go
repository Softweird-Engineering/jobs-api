package config

type Configuration struct {
	Gin     Gin
	Server  ServerConfiguration
	Log     LogConfiguration
	Mongodb MongodbConfiguration
}

type LogConfiguration struct {
	Filename string
}

type ServerConfiguration struct {
	Host     string
	Port     string
	BasePath string
}

type Gin struct {
	Mode string
}

type MongodbConfiguration struct {
	Dsn      string
	Timeouts TimeoutsConfiguration
	Schema   string
}

type TimeoutsConfiguration struct {
	Connection int
	Ping       int
}
