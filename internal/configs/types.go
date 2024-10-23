package config

type (
	Config struct {
		Service  Service  `mapstructure:"Service"`
		Database Database `mapstructure:"Database"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}
)
