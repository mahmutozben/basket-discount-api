package cmd

import (
	"basket-discount-api/internal/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type RootCmd struct {
	cfgFile      string
	AppConfig    *config.Configuration
	cobraCommand *cobra.Command
	DebugMode    bool
	LogLevel     string
}

var RootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "basket-discount",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications. This application is a tool to generate the needed files to quickly create a Cobra application.`,
	},
	AppConfig: &config.Configuration{},
}

// Execute executes the root command.
func Execute() {
	if err := RootCommand.cobraCommand.Execute(); err != nil {
		fmt.Errorf("Failed to execute root command. %v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCommand.cobraCommand.PersistentFlags().StringVarP(&RootCommand.cfgFile, "config", "c", "config.yaml", "config file (default is $HOME/config.yaml)")
	RootCommand.cobraCommand.PersistentPreRunE = func(cmd *cobra.Command, args []string) (err error) {
		return
	}
}

func initConfig() {

	if RootCommand.cfgFile != "" {
		viper.SetConfigFile(RootCommand.cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("coupon.yaml")
	}

	viper.Set("Verbose", true)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file:%s", viper.ConfigFileUsed())

		err = viper.Unmarshal(RootCommand.AppConfig)
		if err != nil {
			fmt.Errorf("unable to decode into config struct, %s", err.Error())
		}
		fmt.Printf("Config: %v", RootCommand.AppConfig)
	} else {
		RootCommand.cobraCommand.Help()
		fmt.Errorf("Failed to read config file. %v", err.Error())
	}
}
