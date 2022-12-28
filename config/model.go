package config

var (
	env = map[string]string{
		"local": "config.local.yml",
	}
)

type Config struct {
	Application Application `yaml:"app"`
	Database    Database    `yaml:"database"`
	SecretKey   string      `yaml:"secretKey"`
}

type Application struct {
	Env  string `yaml:"env"`
	Name string `yaml:"name"`
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Database struct {
	DBName            string `yaml:"dbName"`
	Host              string `yaml:"host"`
	Password          string `yaml:"password"`
	Port              int    `yaml:"port"`
	Schema            string `yaml:"schema"`
	User              string `yaml:"user"`
	Debug             bool   `yaml:"debug"`
	SetMaxIdleCons    int    `yaml:"setMaxIdleCons"`
	SetMaxOpenCons    int    `yaml:"setMaxOpenCons"`
	SetConMaxIdleTime int    `yaml:"setConMaxIdleTime"`
	SetConMaxLifetime int    `yaml:"setConMaxLifeTime"`
}
