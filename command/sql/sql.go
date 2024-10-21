package sql

import (
	"database/sql"
	"fmt"
	"github.com/godevopsdev/dvps/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Database struct {
	DbType string `mapstructure:"dbtype"`
	Server string `mapstructure:"server"`
	Port   int    `mapstructure:"port"`
	Name   string `mapstructure:"name"`
	Option string `mapstructure:"option"`
	Folder string `mapstructure:"folder"`
}

// Config struct to hold the list of databases
type Config struct {
	Databases []Database `mapstructure:"databases"`
}

// DatabaseCmd is the parent command for all database-related operations.
var DatabaseCmd = &cobra.Command{
	Use:   "dbs",
	Short: "Apply command to all databases from config file",
	Long: `dbs is designed for managing databases such and MSSQL and PostgreSQL. 
It allows you to connect, run SQL queries, and manage migrations.`,
}

// ConnectCmd Connect to the database
var ListDbCmd = &cobra.Command{
	Use:   "list",
	Short: "List all databases from config file",
	Run:   listDbCmd,
}

// ConnectCmd Connect to the database
var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the database and verify the connection",
	Run:   connectCmd,
}

// ApplyScriptCmd  Apply all script to databases
var ApplyScriptCmd = &cobra.Command{
	Use:   "applySql",
	Short: "Apply database scripts from all folders location in config file",
	Run:   applyScriptCmd,
}

func connectCmd(cmd *cobra.Command, args []string) {
	// Initialize the configuration
	config := initConfig()
	for _, cfgDb := range config.Databases {
		connString := constructConnString(cfgDb)
		// Connect to the database
		db, err := openDB(cfgDb.DbType, connString)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
		// Test the connection
		if err := db.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}
		db.Close()
		fmt.Println(color.Green("Connected to %s database successfully. Server:%s Db:%s", cfgDb.DbType, cfgDb.Server, cfgDb.Name))
	}
	printDbCount(config.Databases)
}

func applyScriptCmd(cmd *cobra.Command, args []string) {
	config := initConfig()
	for _, cfgDb := range config.Databases {
		connString := constructConnString(cfgDb)
		// Connect to the database
		db, err := openDB(cfgDb.DbType, connString)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
		// Test the connection
		if err := db.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}
		if err := readAndExecuteSQLFiles(db, cfgDb.Folder); err != nil {
			log.Fatalf("Error executing file: %v", err)
		}
		db.Close()
		fmt.Println(color.Green("Scripts run to %s database successfully. Server:%s Db:%s", cfgDb.DbType, cfgDb.Server, cfgDb.Name))
	}
	printDbCount(config.Databases)
}

func listDbCmd(cmd *cobra.Command, args []string) {
	config := initConfig()
	printConfigDatabases(config.Databases)
	printDbCount(config.Databases)
}

// replacePwd replaces password in the connection string with *****
func replacePwd(connString string) string {
	// Replace ${DB_USERNAME} and ${DB_PASSWORD} with actual environment variable values
	connString = strings.ReplaceAll(connString, os.Getenv("DB_PASSWORD"), "********")
	return connString
}

func openDB(dbType string, connString string) (*sql.DB, error) {
	switch dbType {
	case "postgres":
		return sql.Open("postgres", connString)
	case "sql":
		return sql.Open("sqlserver", connString)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

func initConfig() Config {
	var config Config
	cfgFile := viper.GetString("config")
	if cfgFile != "" {
		// If a config file is provided, use it
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")
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
	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	// Sort the databases by the Folder field
	sort.Slice(config.Databases, func(i, j int) bool {
		return config.Databases[i].Folder < config.Databases[j].Folder
	})
	return config
}

func printConfigDatabases(dbs []Database) {
	for i, db := range dbs {
		fmt.Println(color.Yellow("Database: %d", i+1))
		fmt.Println(color.Yellow("  Type: %s", db.DbType))
		fmt.Println(color.Yellow("  Server: %s", db.Server))
		fmt.Println(color.Yellow("  Port: %d", db.Port))
		fmt.Println(color.Yellow("  Name: %s", db.Name))
		fmt.Println(color.Yellow("  Option: %s", db.Option))
		fmt.Println(color.Green("  Folder: %s", db.Folder))
	}
}

func printDbCount(dbs []Database) {
	nbDb := len(dbs)
	if nbDb == 0 {
		fmt.Println(color.Red("No Databases"))
	}
	if nbDb == 1 {
		fmt.Println(color.Green("1 Database"))
	}
	if nbDb > 1 {
		fmt.Println(color.Green("%d Databases", nbDb))
	}
}

// constructConnString replaces placeholders in the connection string with environment variables, if type is not recognize we use a postgres db type
func constructConnString(db Database) string {
	// Replace ${DB_USERNAME} and ${DB_PASSWORD} with actual environment variable values
	switch db.DbType {
	case "postgres":
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), db.Server, db.Port, db.Name, db.Option)
	case "sql":
		return fmt.Sprintf("server=%s,%d;user id=%s;password=%s;database=%s", db.Server, db.Port, os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), db.Name)
	default:
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), db.Server, db.Port, db.Name, db.Option)
	}
}

// readAndExecuteSQLFiles Read and execute all .sql files in lexicographical order from a folder
func readAndExecuteSQLFiles(db *sql.DB, folder string) error {
	// Open the directory
	files, err := os.ReadDir(folder)
	if err != nil {
		return fmt.Errorf("failed to read folder: %v", err)
	}

	var sqlFiles []string

	// Filter only .sql files and collect their paths
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, filepath.Join(folder, file.Name()))
		}
	}

	// Sort files lexicographically by filename
	sort.Strings(sqlFiles)

	// Execute each .sql file's content in sorted order
	for _, sqlFile := range sqlFiles {
		content, err := os.ReadFile(sqlFile)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", sqlFile, err)
		}

		fmt.Printf("Executing SQL from file: %s\n", sqlFile)

		// Execute the SQL content
		_, execErr := db.Exec(string(content))
		if execErr != nil {
			return fmt.Errorf("failed to execute SQL from file %s: %v", sqlFile, execErr)
		}

		fmt.Printf("Successfully executed SQL from file: %s\n", sqlFile)
	}

	return nil
}
