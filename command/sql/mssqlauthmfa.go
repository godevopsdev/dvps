package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

func ConnectMfa() {
	// Azure SQL Server details
	server := "sitsqlleoint16.database.windows.net"
	database := "sitsqdleoint16"

	cred, err := azidentity.NewInteractiveBrowserCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get a token using the Device Code Flow
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	token, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{"https://database.windows.net/.default"},
	})
	if err != nil {
		log.Fatalf("Failed to acquire token: %v", err)
	}

	// Connection string for Azure SQL with Access Token
	connString := fmt.Sprintf("sqlserver://%s?database=%s&access_token=%s", server, database, token.Token)

	// Open a connection to the database
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	defer db.Close()

	// Ping the database to verify connection
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Successfully connected to the Azure SQL Database!")
}
