package main

import (
    "context"
    "database/sql"
    "log"

    _ "github.com/lib/pq"
    "github.com/jemaimedamine22-png/school-api/api"
    "github.com/jemaimedamine22-png/school-api/db"
    "github.com/jemaimedamine22-png/school-api/util"
)

func main() {
    
    tp, err := api.InitTracer("otel-collector:4317")
    if err != nil {
        log.Fatalf("Failed to initialize tracer: %v", err)
    }
    defer func() {
        if err := tp.Shutdown(context.Background()); err != nil {
            log.Printf("Error shutting down tracer provider: %v", err)
        }
    }()

    
    config, err := util.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

    
    conn, err := sql.Open(config.DBDriver, config.DBSource)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    
    store := db.NewStore(conn)
    
    log.Println("connected to database successfully!", store)

    
    server := api.NewServer(store)

    err = server.Start(config.ServerAddress)
    if err != nil {
        log.Fatal("cannot start server:", err)
    }
}