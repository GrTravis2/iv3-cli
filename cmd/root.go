/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/GrTravis2/iv3"
	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
)

/* var Config struct {
	camera iv3.Camera
} */

var Camera iv3.Camera

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iv3-cli",
	Short: "a cli tool to rapidly configure and command iv3 cameras",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	//cobra.OnInitialize(initConfig)
	//Dump viper and just read in a global var struct - should be ok for now...
	data, err := os.ReadFile("iv3-cli.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err2 := yaml.Unmarshal([]byte(data), &Camera)
	if err2 != nil {
		log.Fatalf("error: %v", err2)
	}
	fmt.Printf("Camera in same scope: %+v\n", Camera)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.iv3-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
