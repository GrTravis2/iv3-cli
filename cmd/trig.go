/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/GrTravis2/iv3"
	"github.com/spf13/cobra"
)

// trigCmd represents the trig command
var trigCmd = &cobra.Command{
	Use:   "trig",
	Short: "Trigger loaded camera and read results",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trig called")
		result := iv3.TriggerStatusResult(Camera)
		fmt.Printf("Camera result #: %v\nCamera overall status pass: %v", result.ResultNumber, result.TotalPassResult) //will work after next iv3 release
	},
}

func init() {
	rootCmd.AddCommand(trigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
