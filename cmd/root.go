/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/apoorvajagtap/tempConverter/pkg"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tempin",
	Short: "A CLI used to convert given temperature",
	Long: `tempin could be used to convert a given temperature  into the specified unit, 
i.e; from Farenhite to Celcius, or Celcius to Kelvin or vice-versa.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		temperature, _ := strconv.ParseFloat(args[0], 64)
		unit := strings.ToLower(args[1])

		kel, _ := cmd.Flags().GetBool("kelvin")
		cel, _ := cmd.Flags().GetBool("celcius")
		far, _ := cmd.Flags().GetBool("farenhiet")

		if kel {
			pkg.ToKelvin(temperature, unit)
		}

		if cel {
			pkg.ToCelcius(temperature, unit)
		}

		if far {
			pkg.ToFarenhiet(temperature, unit)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.temperatureIn.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("celcius", "c", false, "Convert to celcius")
	rootCmd.Flags().BoolP("kelvin", "k", false, "Convert to kelvin")
	rootCmd.Flags().BoolP("farenhiet", "f", false, "Convert to farenhiet")
}
