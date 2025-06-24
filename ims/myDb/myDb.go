package myDb

import(
	"log"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/db/sql/postgres"
	
	"github.com/omniful/go_commons/redis"
	"github.com/angel-omniful/ims/myContext"

)
var db *postgres.DbCluster
var cache *redis.Client


func init(){
	ctx := myContext.GetContext()
	myHost := config.GetString(ctx, "postgres.master.host")
	myPort := config.GetString(ctx, "postgres.master.port")
	myUsername := config.GetString(ctx, "postgres.master.username")
	myPassword := config.GetString(ctx, "postgres.master.password")
	myDbname := config.GetString(ctx, "postgres.master.dbname")
	maxOpenConns := config.GetInt(ctx, "postgres.master.max_open_conns")
	maxIdleConns := config.GetInt(ctx, "postgres.master.max_idle_conns")
	connMaxLifetime := config.GetDuration(ctx, "postgres.master.conn_max_lifetime")
	debugMode := config.GetBool(ctx, "postgres.master.debug_mode")

	masterConfig := postgres.DBConfig{
		Host:               myHost,
		Port:               myPort,
		Username:           myUsername,
		Password:           myPassword,
		Dbname:             myDbname,
		MaxOpenConnections: maxOpenConns,
		MaxIdleConnections: maxIdleConns,
		ConnMaxLifetime:  connMaxLifetime,
		DebugMode:          debugMode,
	}
	// Initialize slavesConfig as an empty slice
	// they can be added later if needed
	slavesConfig := make([]postgres.DBConfig, 0) // read replicas

	db = postgres.InitializeDBInstance(masterConfig, &slavesConfig)

	//db is a cluster here
	log.Println("Connecting to the database...")
	if db == nil {
		log.Panic("failed to connect to the database")
	} else {
		log.Println("Database connected successfully!")
	}

	log.Println("Connecting to the redis_cache...")
	config := &redis.Config{
    Hosts: []string{"localhost:6379"},
    PoolSize: 50,
    MinIdleConn: 10,
	}

	cache = redis.NewClient(config)
	if cache == nil {
		log.Panic("failed to connect to the redis cache")
	} else {
		log.Println("Redis cache connected successfully!")
	}
}

func GetDb() *postgres.DbCluster {
	return db
}

func GetCache() *redis.Client {
	return cache
}		

	