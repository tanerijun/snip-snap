package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tanerijun/snip-snap/internal/models"
)

// Handler dependencies
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	dsn := flag.String("dsn", "", "PostgreSQL data source name (required)")

	flag.Parse()
	if *dsn == "" {
		flag.Usage()
		os.Exit(1)
	}

	db, err := openDb(*dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		snippets: &models.SnippetModel{DB: db},
	}

	mux := app.routes(*staticDir)

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Server started on %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDb(connStr string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
