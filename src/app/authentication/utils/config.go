package utils

import (
	"authentication/constants"
	"authentication/models"
	"log"

	"github.com/spf13/viper"
)

// Takes yaml filePath and return struct 
func ParseYAML() models.DbDetails {
    viper.SetConfigFile("resources/application.yml")

    // Read the configuration file into Viper
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("%s: %s",constants.ErrReadConfigFile, err)
    }

    // Unmarshal the configuration into a struct
    var dbDetails models.DbDetails
    if err := viper.UnmarshalKey("database", &dbDetails); err != nil {
        log.Fatalf("%s: %s",constants.ErrDecodeConfig, err)
    }

    // Return the database details
    return dbDetails
}
