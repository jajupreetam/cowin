package config

import (
	"fmt"
	slog "github.com/go-eden/slf4go"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

func InitProfiles() {
	activeProfiles := GetProfiles()
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./resources")
	viper.AddConfigPath("/deployments")
	viper.AddConfigPath("/deployments/config")
	if err := viper.ReadInConfig(); err != nil {
		slog.Fatalf("Error reading config, %s", err)
	} else {
		fmt.Println("Read config")
	}
	for _, profile := range activeProfiles {
		var profilePrefix = "config-"
		var profileName = profilePrefix + profile
		viper.SetConfigName(profileName)
		err := viper.MergeInConfig()
		if err != nil {
			slog.Warnf("Error merging config named %s, %s", profileName, err)
		} else {
			fmt.Println("Read " + profileName)
		}
		if profile == "oc" {
			viper.SetConfigName("application-oc")
			err := viper.MergeInConfig()
			if err != nil {
				slog.Warnf("Error merging config named %s, %s", "application-oc", err)
			} else {
				fmt.Println("Read application-oc")
			}
		}
	}
	c := viper.AllSettings()
	bs, _ := yaml.Marshal(c)
	s := string(bs)
	fmt.Println("Merged Configuration:")
	fmt.Println(s)
}

func GetProfiles() []string {
	var activeProfiles = []string{}
	var profiles = strings.Split(getEnv("PROFILE", ""), ",")
	for _, profile := range profiles {
		if len(profile) > 0 {
			activeProfiles = append(activeProfiles, profile)
		}
	}
	fmt.Printf("Running profiles %s\n", activeProfiles)
	return activeProfiles
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
