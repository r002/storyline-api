package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Runtime string  `yaml:"runtime"`
	Main    string  `yaml:"main"`
	EnvVars EnvVars `yaml:"env_variables"`
}

type EnvVars struct {
	Env               string `yaml:"STUDYDASH_ENV"`
	KeyGhWebhook      string `yaml:"STUDYDASH_KEY_GH_WEBHOOK"`
	KeyGhToken        string `yaml:"STUDYDASH_KEY_GH_TOKEN"`
	GhRepoEndpoint    string `yaml:"STUDYDASH_GH_REPO_ENDPOINT"`
	FirestoreEndpoint string `yaml:"STUDYDASH_FIRESTORE_ENDPOINT"`
}

var config *Config

func GetEnvVars() EnvVars {
	if config == nil {
		// If an ENV `*.yaml` is passed in, always honor that first
		if len(os.Args) == 2 {
			loadConfigFromFile(os.Args[1])
		} else {
			_, ok := os.LookupEnv("STUDYDASH_ENV")
			if ok {
				log.Println(">> Loading config from env vars")
				config = new(Config)
				(*config).EnvVars.Env = os.Getenv("STUDYDASH_ENV")
				(*config).EnvVars.KeyGhWebhook = os.Getenv("STUDYDASH_KEY_GH_WEBHOOK")
				(*config).EnvVars.KeyGhToken = os.Getenv("STUDYDASH_KEY_GH_TOKEN")
				(*config).EnvVars.GhRepoEndpoint = os.Getenv("STUDYDASH_GH_REPO_ENDPOINT")
				(*config).EnvVars.FirestoreEndpoint = os.Getenv("STUDYDASH_FIRESTORE_ENDPOINT")
			} else {
				loadConfigFromFile("./app.qa.yaml")
			}
		}
	}
	return config.EnvVars
}

func loadConfigFromFile(configFile string) {
	log.Printf(">> Loading config from file: %q,", configFile)
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf(">> Error reading config: %q", configFile)
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Printf(">> Error in file %q: %v", configFile, err)
	}
}
