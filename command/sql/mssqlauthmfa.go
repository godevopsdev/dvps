package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

func ConnectMfa() {
	// Azure SQL Server details
	server := "sitsqlleoint16.database.windows.net"
	database := "sitsqdleoint16"

	ibc := azidentity.InteractiveBrowserCredentialOptions{
		TenantID: "8287db1c-e074-412b-9c78-e1d70e3cc068",
	}
	cred, err := azidentity.NewInteractiveBrowserCredential(&ibc)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	tokenProvider := func() (string, error) {
		token, err := cred.GetToken(context.TODO(), policy.TokenRequestOptions{
			Scopes: []string{"https://database.windows.net//.default"},
		})
		return token.Token, err
	}

	// Connection string for Azure SQL with Access Token
	connString := fmt.Sprintf("sqlserver://%s?database=%s&fedauth=ActiveDirectoryInteractive", server, database)
	connector, err := mssql.NewAccessTokenConnector(connString, tokenProvider)
	if err != nil {
		log.Fatal("Connector creation failed:", err.Error())
	}
	// Open a connection to the database
	db := sql.OpenDB(connector)
	// Ping the database to verify connection
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Println("Successfully connected to the Azure SQL Database!")
	row := db.QueryRow("select 1, 'abc'")
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)
	err = db.Close()
	if err != nil {
		log.Fatal("error closed:", err.Error())
	}

}
