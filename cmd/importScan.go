/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xNaaro/defectdojo-go/importScan"
)

var file_name string

// importScanCmd represents the importScan command
var importScanCmd = &cobra.Command{
	Use:     "importScan",
	Aliases: []string{"imp"},
	Short:   "Import results",
	Long:    `Import scan results to Defect Dojo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("importScan called")
		var test2 = 1
		res := importScan.CreateImport(test2)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(importScanCmd)

	importScanCmd.PersistentFlags().String("file_name", "results.json", "File name or absolute path to upload.")
	viper.BindPFlag("file_name", importScanCmd.PersistentFlags().Lookup("file_name"))

	importScanCmd.PersistentFlags().String("scan_type", "Bandit Scan", "Scan type, one of the supported by DefectDojo. Case and space sensitive.")
	viper.BindPFlag("scan_type", importScanCmd.PersistentFlags().Lookup("scan_type"))

	importScanCmd.PersistentFlags().String("product_name", "", "Product name to upload report.")
	viper.BindPFlag("product_name", importScanCmd.PersistentFlags().Lookup("product_name"))

	importScanCmd.PersistentFlags().String("engagement_name", "", "Engangement name to upload report.")
	viper.BindPFlag("engagement_name", importScanCmd.PersistentFlags().Lookup("engagement_name"))

	importScanCmd.PersistentFlags().String("minimum_severity", "Info", "Minimum severity.")
	viper.BindPFlag("minimum_severity", importScanCmd.PersistentFlags().Lookup("minimum_severity"))

	importScanCmd.PersistentFlags().String("active", "true", "Active status.")
	viper.BindPFlag("active", importScanCmd.PersistentFlags().Lookup("active"))

	importScanCmd.PersistentFlags().String("verified", "true", "Verified scan.")
	viper.BindPFlag("verified", importScanCmd.PersistentFlags().Lookup("verified"))

	importScanCmd.PersistentFlags().String("close_old_findings", "false", "Close old findings.")
	viper.BindPFlag("close_old_findings", importScanCmd.PersistentFlags().Lookup("close_old_findings"))

	importScanCmd.PersistentFlags().String("push_to_jira", "false", "Push to Jira.")
	viper.BindPFlag("push_to_jira", importScanCmd.PersistentFlags().Lookup("push_to_jira"))

	importScanCmd.PersistentFlags().String("scan_date", "2025-03-18", "Scan date.")
	viper.BindPFlag("scan_date", importScanCmd.PersistentFlags().Lookup("scan_date"))

	importScanCmd.PersistentFlags().String("check_list", "true", "Check list.")
	viper.BindPFlag("check_list", importScanCmd.PersistentFlags().Lookup("check_list"))

	importScanCmd.PersistentFlags().String("status", "Not Started", "Status.")
	viper.BindPFlag("status", importScanCmd.PersistentFlags().Lookup("close_oldstatus_findings"))
}
