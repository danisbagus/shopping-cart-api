package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"

	//driver
	_ "github.com/lib/pq"
)

var usageCommands = `
Run database migrations & seeder

Usage:
    [command]

Available Migration Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version

`

func main() {
	godotenv.Load()
	var rootCmd = &cobra.Command{
		Use:   "migrate",
		Short: "psql Migration Service",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}

			host := os.Getenv("DB_HOST")
			username := os.Getenv("DB_USERNAME")
			password := os.Getenv("DB_PASSWORD")
			name := os.Getenv("DB_NAME")
			port := os.Getenv("DB_PORT")
			SSLMode := os.Getenv("DB_SSL_MODE")

			dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, name, SSLMode)
			goose.SetDialect("postgres")
			db, err := sql.Open("postgres", dataSourceName)
			if err != nil {
				log.Fatalf("failed connect to psql: %v", err)
			}

			appPath, _ := os.Getwd()
			dir := appPath + "/app/migration/psql/migrations"
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
			command := args[0]
			cmdArgs := args[1:]
			if err := goose.Run(command, db, dir, cmdArgs...); err != nil {
				log.Fatalf("goose run: %v", err)
			}
		},
	}

	rootCmd.SetUsageTemplate(usageCommands)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
