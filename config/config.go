package config

type Configuration struct {
	ServiceHost   string
	ServicePort   int
	MongoDB       string
	MongoUser     string
	MongoPassword string
	MongoURI      string
}

func LoadCfg() (*Configuration, error) {
	return &Configuration{
		ServiceHost:   "localhost:",
		ServicePort:   1997,
		MongoDB:       "mongodb",
		MongoUser:     "testUser",
		MongoPassword: "testPassword",
		MongoURI:      "mongodb://localhost:27017/mongodb",
	}, nil
}
