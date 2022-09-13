package config

type Config struct {
	Port int16
}

func NewConfig() Config {
	var port int16 = 8008
	return Config{Port: port}
}
