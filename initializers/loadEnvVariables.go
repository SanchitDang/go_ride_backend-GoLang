package initializers

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func LoadEnvVariables(){
		// Set the file name of the configuration file
		viper.SetConfigFile("config.yaml")

		// Read the configuration file
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
}