/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure camera values",
	Long: `Set values for camera location, description, ip address, port #, and delimiter.
	These need to be set properly to make the rest of the tool work.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		fmt.Printf("location before: %v\n", Camera.Location)
		Camera.Config()
		fmt.Printf("location after: %v\n", Camera.Location)
		cameraJSON, errJson := json.Marshal(Camera)
		if errJson != nil {
			fmt.Printf("Error converting struct to json. Error: %v", errJson)
		}
		err := os.WriteFile("iv3-cli.json", cameraJSON, 0644)
		if err != nil {
			fmt.Printf("Error writing to iv3-cli.json. Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
