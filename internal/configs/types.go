package config

type (
	Config struct {
		Service Service `mapstructure:"Service"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}
)
