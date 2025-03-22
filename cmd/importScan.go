/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xNaaro/defectdojo-go/importScan"
)

// importScanCmd represents the importScan command
var importScanCmd = &cobra.Command{
	Use:     "importScan",
	Aliases: []string{"imp"},
	Short:   "Import results",
	Long:    `Import scan results to Defect Dojo.`,
	Run: func(cmd *cobra.Command, args []string) {
		product_name, _ := cmd.Flags().GetString("product_name")
		engagement_name, _ := cmd.Flags().GetString("engagement_name")

		file_name, _ := cmd.Flags().GetString("file_name")
		scan_type, _ := cmd.Flags().GetString("scan_type")
		minimum_severity, _ := cmd.Flags().GetString("minimum_severity")
		active, _ := cmd.Flags().GetString("active")
		verified, _ := cmd.Flags().GetString("verified")
		close_old_findings, _ := cmd.Flags().GetString("close_old_findings")
		push_to_jira, _ := cmd.Flags().GetString("push_to_jira")
		scan_date, _ := cmd.Flags().GetString("scan_date")
		check_list, _ := cmd.Flags().GetString("check_list")
		status, _ := cmd.Flags().GetString("status")
		importScan.CreateImport(product_name, engagement_name, file_name, scan_type, minimum_severity, active, verified, close_old_findings, push_to_jira, scan_date, check_list, status)
	},
}

func init() {
	const YYYYMMDD = "2006-01-02"
	currentTime := time.Now()
	var default_date = currentTime.Format(YYYYMMDD)
	rootCmd.AddCommand(importScanCmd)

	importScanCmd.PersistentFlags().String("file_name", "results.json", "File name or path to upload.")
	viper.BindPFlag("file_name", importScanCmd.PersistentFlags().Lookup("file_name"))
	importScanCmd.MarkPersistentFlagRequired("file_name")

	importScanCmd.PersistentFlags().String("scan_type", "Bandit Scan", "Scan type, one of the supported by DefectDojo. Case and space sensitive.")
	viper.BindPFlag("scan_type", importScanCmd.PersistentFlags().Lookup("scan_type"))
	importScanCmd.MarkPersistentFlagRequired("scan_type")

	importScanCmd.PersistentFlags().String("product_name", "", "Product name to upload report.")
	viper.BindPFlag("product_name", importScanCmd.PersistentFlags().Lookup("product_name"))
	importScanCmd.MarkPersistentFlagRequired("product_name")

	importScanCmd.PersistentFlags().String("engagement_name", "", "Engangement name to upload report.")
	viper.BindPFlag("engagement_name", importScanCmd.PersistentFlags().Lookup("engagement_name"))
	importScanCmd.MarkPersistentFlagRequired("engagement_name")

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

	importScanCmd.PersistentFlags().String("scan_date", default_date, "Scan date.")
	viper.BindPFlag("scan_date", importScanCmd.PersistentFlags().Lookup("scan_date"))

	importScanCmd.PersistentFlags().String("check_list", "true", "Check list.")
	viper.BindPFlag("check_list", importScanCmd.PersistentFlags().Lookup("check_list"))

	importScanCmd.PersistentFlags().String("status", "Not Started", "Status.")
	viper.BindPFlag("status", importScanCmd.PersistentFlags().Lookup("close_oldstatus_findings"))
}
