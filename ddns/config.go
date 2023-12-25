package ddns

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Config FullConfig

type Authorize struct {
	ApiKey string `yaml:"apiKey"`
	Email  string `yaml:"email"`
}

type DNS struct {
	Domain        string `yaml:"domain"`
	RecordName    string `yaml:"recordName"`
	Endpoint      string `yaml:"endpoint"`
	ScheduledTask int64  `yaml:"scheduledTask"`
}

type Cloudflare struct {
	Authorize `yaml:"authorize"`
	DNS       `yaml:"dns"`
}

type Telegram struct {
	BotToken string `yaml:"botToken"`
	ChatID   int64  `yaml:"chatID"`
}

type Notifications struct {
	Telegram `yaml:"telegram"`
}

type FullConfig struct {
	Cloudflare    `yaml:"cloudflare"`
	Notifications `yaml:"notifications"`
}

func LoadConfig(filePath string) error {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(f, &Config)
}
