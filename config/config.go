package config

type configuration struct {
	DatabaseName     string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
}

func GetConfiguration() *configuration {
	return &configuration{
		DatabaseName:     `postgres`,
		DatabaseHost:     `localhost:54322`,
		DatabaseUser:     `postgres`,
		DatabasePassword: `password`,
	}
}
