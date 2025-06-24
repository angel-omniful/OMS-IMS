package main

import (
	
	"log"
	"time"
	"github.com/omniful/go_commons/http"
	"github.com/angel-omniful/ims/routes"
	"github.com/angel-omniful/ims/myContext"
    "github.com/angel-omniful/ims/myDb"
	//"github.com/omniful/go_commons/db/sql/migration"
	"github.com/omniful/go_commons/config"
	

	
)



func main() {
    
	ctx:= myContext.GetContext()

	log.Println("Connecting to the server...")

	port:= config.GetString(ctx, "server.port")
	port = ":" + port
	log.Println("Connecting to the server...")
	server:= http.InitializeServer(
	port,                // listen address
	10 * time.Second,       // read timeout
	10 * time.Second,       // write timeout
	70 * time.Second,       // idle timeout
	false,
	// Add custom middleware here...
	)
	log.Println("Registering routes...")
	routes.RegisterAllRoutes(server)
	log.Println("Routes registered successfully!")
	// Start the server
	


	//1.context setup and config loading happening in myContext package
	//2.server setup and initialization happening in server package
	//3.database connection and redis connection happening in myDb package
	// Initialize the database connection
	log.Println("Connecting to the database...")
	db:=myDb.GetDb()
	if db==nil {
		log.Panic("failed to connect")
	}
	log.Println("Database connected successfully")
	// ctx:=myContext.GetContext()
    // //  // Build the database URL
	// log.Println("Starting procedure for migrations...")
	// myHost := config.GetString(ctx, "postgres.master.host")
	// myPort := config.GetString(ctx, "postgres.master.port")
	// myUsername := config.GetString(ctx, "postgres.master.username")
	// myPassword := config.GetString(ctx, "postgres.master.password")
	// myDbname := config.GetString(ctx, "postgres.master.dbname")
    
	// dbURL := migration.BuildSQLDBURL(
    //     myHost,
    //     myPort,
    //     myDbname,
    //     myUsername,
    //     myPassword,
    // )
    // log.Println("Database URL:", dbURL)
    // // Path to migration files
    // migrationPath :="file://C:/Users/angel/OneDrive/Desktop/OMNIFUL-PROJECT/ims/migrations"
    
    // // Initialize the migrator
    // migrator, err := migration.InitializeMigrate(migrationPath, dbURL)
    // if err != nil {
    //     log.Fatalf("Failed to initialize migrator: %v", err)
    // }
    
    // // Apply all pending migrations
    // err = migrator.Up()
    // if err != nil {
    //     log.Fatalf("Failed to apply migrations: %v", err)
    // }
	// log.Println("Migrations applied successfully")	
   if err := server.StartServer("my-service"); err != nil {
		log.Println("server not connected")
     }else{
		log.Println("server connected successfully")
	}

	
}

