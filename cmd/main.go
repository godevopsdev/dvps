package main

import (
	"database/sql"
	"fmt"
	"github.com/godevopsdev/dvps/command/crypto"
	dvpssql "github.com/godevopsdev/dvps/command/sql"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb" // MSSQL driver
	_ "github.com/lib/pq"                // PostgreSQL driver
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func main() {
	// red := ansi.ColorFunc("red")
	// blue := ansi.ColorFunc("blue")
	// Define root command
	var rootCmd = &cobra.Command{
		Use:   "dvps",
		Short: "A CLI tool to manage databases in your DevOps pipelines",
		Long: `dvps is a Command-Line Interface (CLI) tool designed for managing
databases such and MSSQL and PostgreSQL. It allows you to connect, run SQL queries, 
and manage migrations. Suitable for use in DevOps pipelines like Azure DevOps.`,
	}
	// Add flags for command-line arguments
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is dvps.yml)")

	// Add flags for database configuration
	rootCmd.PersistentFlags().String("dbtype", "", "The type of database (e.g., sql, postgres)")
	rootCmd.PersistentFlags().String("server", "", "The database server")
	rootCmd.PersistentFlags().Int("port", 0, "The database port")
	rootCmd.PersistentFlags().String("name", "", "The database name")
	rootCmd.PersistentFlags().String("option", "", "Additional connection options")

	// Bind the Cobra flags to Viper (so they can be overridden by flags or config file)
	viper.BindPFlag("database.dbtype", rootCmd.PersistentFlags().Lookup("dbtype"))
	viper.BindPFlag("database.server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("database.port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("database.name", rootCmd.PersistentFlags().Lookup("database"))
	viper.BindPFlag("database.option", rootCmd.PersistentFlags().Lookup("option"))

	// Add subcommands
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(applyScriptCmd)
	rootCmd.AddCommand(azureKeyCmd)
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		// If a config file is provided, use it
		viper.SetConfigFile(cfgFile)
	} else {
		// Default to dvps.yml if no config file is provided
		viper.SetConfigName("dvps")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	yellow := ansi.ColorFunc("yellow+b+h")
	// Environment variable overrides (optional but useful for debugging)
	fmt.Println(yellow("Current Configuration:"))
	fmt.Println(yellow("Database Type:"), viper.GetString("database.dbtype"))
	fmt.Println(yellow("Connection String:"), replacePwd(dvpssql.ConstructConnString(viper.GetString("database.dbtype"))))
}

// constructConnString replaces placeholders in the connection string with environment variables
func replacePwd(connString string) string {
	// Replace ${DB_USERNAME} and ${DB_PASSWORD} with actual environment variable values
	connString = strings.ReplaceAll(connString, os.Getenv("DB_PASSWORD"), "********")
	return connString
}

// Subcommand: Connect to the database
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the database and verify the connection",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the configuration
		initConfig()

		// Get database type and connection string from Viper
		dbType := viper.GetString("database.dbtype")
		connString := dvpssql.ConstructConnString(dbType)

		// Connect to the database
		db, err := openDB(dbType, connString)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
		defer db.Close()

		// Test the connection
		if err := db.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}
		green := ansi.ColorFunc("green")
		fmt.Printf(green("Connected to %s database successfully!\n"), dbType)
	},
}

// Subcommand: Connect to the database
var applyScriptCmd = &cobra.Command{
	Use:   "applySql [folder]",
	Short: "Apply database scripts from folder location",
	Args:  cobra.ExactArgs(1),
	Run:   applyCmd,
}

// Subcommand: Connect to the database
var azureKeyCmd = &cobra.Command{
	Use:   "azureKey [keyname]",
	Short: "Generate Key Pair that can be load in Azure Key Vault ",
	Args:  cobra.ExactArgs(1),
	Run:   azureKey,
}

func azureKey(cmd *cobra.Command, args []string) {
	keyname := args[0]
	crypto.GenerateRSA2048(keyname)
}

func applyCmd(cmd *cobra.Command, args []string) {

	folder := args[0]
	initConfig()

	// Get database type and connection string from Viper
	dbType := viper.GetString("database.dbtype")
	connString := dvpssql.ConstructConnString(dbType)

	// Connect to the database
	db, err := openDB(dbType, connString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	if err := dvpssql.ReadAndExecuteSQLFiles(db, folder); err != nil {
		log.Fatalf("Error executing file: %v", err)
	}

	green := ansi.ColorFunc("green")
	fmt.Printf(green("Connected to %s database successfully!\n"), dbType)
}

func openDB(dbType, connString string) (*sql.DB, error) {
	switch dbType {
	case "postgres":
		return sql.Open("postgres", connString)
	case "sql":
		return sql.Open("sqlserver", connString)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
