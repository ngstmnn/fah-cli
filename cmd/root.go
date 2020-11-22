package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "fah-cli",
		Short: "A CLI to easily access a FAH-Client through the network.",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to the configuration file")

	rootCmd.PersistentFlags().IntP("port", "p", 36330, "Port the server listens on")
	_ = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

	rootCmd.PersistentFlags().StringP("hostname", "s", "localhost", "Hostname of the server to connect to")
	_ = viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("hostname"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("fah-cli")
	}

	viper.SetEnvPrefix("FAH_CLI")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
