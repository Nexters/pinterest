package config

import (
	"github.com/rs/zerolog/log"

	env "github.com/Netflix/go-env"
)

// Settings 환경변수
type Settings struct {
	// App 앱 환경변수
	App struct {
		Port int `env:"PORT" json:"port"`
	}

	Database struct {
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		Port     int    `env:"DATABASE_PORT"`
		Name     string `env:"DATABASE_NAME"`
		URL      string `env:"DATABASE_URL"`
	}

	Extras env.EnvSet `json:"-"`
}

// NewSettings 생성자
func NewSettings() *Settings {
	var settings Settings

	extras, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Fatal().Err(err)
	}

	settings.Extras = extras
	return &settings
}
