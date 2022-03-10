package config

type Configuration struct {
    Gin         Gin
	Server      ServerConfiguration
	Log         LogConfiguration
}

type LogConfiguration struct {
    Filename    string
}

type ServerConfiguration struct {
	Host        string
	Port        string
	BasePath    string
}

type Gin struct {
    Mode        string
}
