package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	aliens     int
	iterations int
	verbose    bool

	rootCmd = &cobra.Command{
		Use:     "",
		Short:   "Run the simulation of the invasion.",
		Example: "map.txt --aliens 100 --iterations 10000",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(fmt.Sprintf("Running simulation with %d aliens and %d maximum iterations to run on map %s.", aliens, iterations, args[0]))
			fmt.Println("Simulation complete.")
			return nil
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&aliens, "aliens", "a", 10, "Number of aliens to invade.")
	rootCmd.PersistentFlags().IntVarP(&iterations, "iterations", "i", 1000, "The maximum number of iterations to run.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbosity level, print all aliens steps.")

	_ = viper.BindPFlag("aliens", rootCmd.PersistentFlags().Lookup("aliens"))
	_ = viper.BindPFlag("iterations", rootCmd.PersistentFlags().Lookup("iterations"))
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.SetDefault("aliens", 10)
	viper.SetDefault("iterations", 10000)
	viper.SetDefault("verbose", false)
}
