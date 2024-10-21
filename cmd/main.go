package main

import (
	_ "github.com/denisenkom/go-mssqldb" // MSSQL driver
	"github.com/godevopsdev/dvps/command/azure/keyvault"
	dvpssql "github.com/godevopsdev/dvps/command/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfgFile string

func main() {
	// red := ansi.ColorFunc("red")
	// blue := ansi.ColorFunc("blue")
	// Define root command
	var rootCmd = &cobra.Command{
		Use:   "dvps",
		Short: "A CLI tool to manage databases in your DevOps pipelines",
		Long:  `dvps is a Command-Line Interface (CLI) tool. Suitable for use in DevOps pipelines like Azure DevOps or others.`,
	}
	// Add flags for command-line arguments
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is dvps.yml)")

	// Bind the config flag to Viper
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	// Add subcommands
	rootCmd.AddCommand(dvpssql.DatabaseCmd)
	rootCmd.AddCommand(keyvault.AzureKeyCmd)

	dvpssql.DatabaseCmd.AddCommand(dvpssql.ListDbCmd)
	dvpssql.DatabaseCmd.AddCommand(dvpssql.ConnectCmd)
	dvpssql.DatabaseCmd.AddCommand(dvpssql.ApplyScriptCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
