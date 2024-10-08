package sql

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ConstructConnString replaces placeholders in the connection string with environment variables
func ConstructConnString(dbType string) string {
	// Replace ${DB_USERNAME} and ${DB_PASSWORD} with actual environment variable values
	switch dbType {
	case "postgres":
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), viper.GetString("database.server"), viper.GetInt("database.port"), viper.GetString("database.name"), viper.GetString("database.option"))
	case "sql":
		return fmt.Sprintf("server=%s,%d;user id=%s;password=%s;database=%s", viper.GetString("database.server"), viper.GetInt("database.port"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), viper.GetString("database.name"))
	case "master":
		return fmt.Sprintf("server=%s,%d;user id=%s;password=%s;database=%s", viper.GetString("database.server"), viper.GetInt("database.port"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), "master")
	default:
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), viper.GetString("database.server"), viper.GetInt("database.port"), viper.GetString("database.name"), viper.GetString("database.option"))
	}
}

// ReadAndExecuteSQLFiles Read and execute all .sql files in lexicographical order from a folder
func ReadAndExecuteSQLFiles(db *sql.DB, folder string) error {
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
