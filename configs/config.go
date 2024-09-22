package configs

// viper lida com necessidades de configuração pra aplicações go
// incluindo definições default, leitura de arquivos de configuração, etc.
import "github.com/spf13/viper"

// cfg é um ponteiro para a configuração do sistema
var cfg *config

// config armazena a configuração do sistema
type config struct {
	API APIConfig
	DB  DBConfig
}

// DBConfig armazena a configuração do database
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// APIConfig armazena a configuração da API
type APIConfig struct {
	Port string
}

func init() {
	// seta as configurações default da api incluindo ports (5432 é a porta padrão do postgres)
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// se houver um erro ao ler o arquivo de configuração, ele será retornado
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// relembrando: cfg é um ponteiro para a struct config
	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}

	return nil
}
func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
