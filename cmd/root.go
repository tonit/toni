package cmd

import (
	"fmt"
	"os"
	"strconv"

	execute "github.com/alexellis/go-execute/pkg/v1"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "toni",
		Short: "A CLI for Toni Menzel Consulting Services",
		Long:  `A CLI for Toni Menzel Consulting Services.`,
	}

	initCmd = &cobra.Command{
		Use:   "open",
		Short: "Open Toni's Blog",
		RunE: func(cmd *cobra.Command, args []string) error {
			println("Hello World")
			command := execute.ExecTask{
				Command:     "open",
				Args:        []string{"https://tonimenzel.de"},
				StreamStdio: true,
			}

			var res, err = command.Execute()
			if err != nil {
				fmt.Println("Error..", err)
			}
			fmt.Println("Exit: " + strconv.Itoa(res.ExitCode))
			return nil
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toni.yaml)")
	rootCmd.AddCommand(initCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
