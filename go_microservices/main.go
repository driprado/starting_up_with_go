package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"github.com/driprado/starting_up_with_go/go_microservices/account"
	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//"user:password@localhost:5432/dbname?sslmode=disable"
const dbsource = "postgress:postgres@localhost:5432/gokitexample?sslmode=disable" // URI to be used to interact with postgres db

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address") // setup the port
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)                                                                        // formats logger
		logger = log.NewSyncLogger(logger)                                                                             // syncs logger
		logger = log.With(logger, "service", "account", "time:", log.DefaultTimestampUTC, "caller", log.DefaultCaller) // Add lables to logger
	}

	level.Info(logger).Log("msg", "service started")     // Will log when service starts
	defer level.Info(logger).Log("msg", "service ended") // Will log when go code is done

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	flag.Parse()                // parse cli flag "-1" if we get one
	ctx := context.Background() // creates non nil empty context to be passed to microservices
	var srv account.Service     // type Service defined in account
	{
		repository := account.NewRepo(db, logger)
		srv = account.NewService(repository, logger)
	}

	errs := make(chan error) // create errors channel
	go func() {
		c := make(chan os.Signal, 1)                      // create channel to see if os sends signal
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) // notify for this type of signals
		errs <- fmt.Errorf("%s", <-c)                     // send to err channel create above
	}()

	endpoints := account.MakeEndpoints(srv) // defined in `6_endpoint.go` takes a service returns 2 endpoints

	go func() { // go routine to execute the server
		fmt.Println("listen on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs) // drain error channel into logger

}
