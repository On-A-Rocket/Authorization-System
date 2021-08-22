package config

type Interface interface {
	Auth() AuthInterface
	Database() DatabaseInterface
	Redis() RedisInterface
}

type Config struct {
	auth     AuthInterface
	database DatabaseInterface
	redis    RedisInterface
}

func Initialize() Interface {
	return &Config{
		auth:     newAuthConfig(),
		database: newDatabaseConfig(),
		redis:    newRedisConfig(),
	}
}

func (config *Config) Auth() AuthInterface {
	return config.auth
}

func (config *Config) Database() DatabaseInterface {
	return config.database
}

func (config *Config) Redis() RedisInterface {
	return config.redis
}
