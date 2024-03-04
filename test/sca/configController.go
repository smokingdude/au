package main

import (
	"fmt"

	"github.com/spf13/viper"
)

const DefaultNugetBaseUrl = "https://api.nuget.org/v3-flatcontainer/"
const DefaultNpmRegistryUrl = "https://registry.npmjs.com/"

type HttpClientConfig struct {
	Proxy    string
	Insecure bool
}

type AnalyzersConfig struct {
	Java  JavaAnalyzerConfig
	Npm   NpmAnalyzerConfig
	Nuget NugetAnalyzerConfig
	Pip   PipAnalyzerConfig
	Go    GoAnalyzerConfig
}

type JavaAnalyzerConfig struct {
	Proxy      JavaProxyAnalyzerConfig
	Registries []string
}

type NpmAnalyzerConfig struct {
	Registry string
}

type NugetAnalyzerConfig struct {
	Registry string
}

type PipAnalyzerConfig struct {
	Indexes []string
	Proxy   string
}

type GoAnalyzerConfig struct {
	HttpProxy string
	Proxy     string
	Private   string
}

type JavaProxyAnalyzerConfig struct {
	Login    string
	Password string
	Address  string
	Port     string
}

type VulnDBConfig struct {
	Baseurl string
}

type Configurations struct {
	Source     string
	Logdir     string
	Report     string
	HttpClient HttpClientConfig
	Analyzers  AnalyzersConfig
	VulnDB     VulnDBConfig
}

var config Configurations

/*
	A function to read and get YAML config file fields.

\param[in] name - YAML filename
\return An error (nil if OK and any other if not)
*/
func loadConfig(name string) error {
	fmt.Printf("Using config file %s", name)
	viper.AddConfigPath(".")
	viper.SetConfigName("bom.json")
	viper.SetConfigType("yaml")
	viper.SetDefault("Source", "/code")
	viper.SetDefault("Logdir", "/var/log/recdep")
	viper.SetDefault("Report", "/result/result.json")
	viper.SetDefault("HttpClient.Insecure", false)
	viper.SetDefault("Analyzers.Npm.Registry", DefaultNpmRegistryUrl)
	viper.SetDefault("Analyzers.Nuget.Registry", DefaultNugetBaseUrl)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	fmt.Println(config)

	return nil

}

func isProxySet() bool {
	return len(config.HttpClient.Proxy) > 0
}

func isGoHttpProxySet() bool {
	return len(config.Analyzers.Go.HttpProxy) > 0
}

func isGoProxySet() bool {
	return len(config.Analyzers.Go.Proxy) > 0
}

func isGoPrivateSet() bool {
	return len(config.Analyzers.Go.Private) > 0
}

func isPipIndexSet() bool {
	return len(config.Analyzers.Pip.Indexes) > 0
}

func isPipProxySet() bool {
	return len(config.Analyzers.Pip.Proxy) > 0
}

func isNpmRegistrySet() bool {
	return len(config.Analyzers.Npm.Registry) > 0
}
