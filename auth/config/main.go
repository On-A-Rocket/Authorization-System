package config

type Interface interface {
	Database() DatabaseConfigInterface
}

type Config struct {
	database DatabaseConfigInterface
}

func Initialize() Interface {
	return &Config{
		database: newDatabaseConfig(),
	}
}

func (config *Config) Database() DatabaseConfigInterface {
	return config.database
}
