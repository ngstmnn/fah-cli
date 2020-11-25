package cmd

import (
	"github.com/ngstmnn/fah-cli/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	cfgFile    string
	connection pkg.Connection
	rootCmd    = &cobra.Command{
		Use:   "fah-cli",
		Short: "A CLI to easily access a FAH-Client through the network.",
		Long:  "Control your Folding At Home client using a cli.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			port := viper.GetInt("port")
			hostname := viper.GetString("hostname")

			if con, err := pkg.Open(hostname, port); err != nil {
				return err
			} else {
				connection = con
			}

			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return connection.Close()
		},
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
