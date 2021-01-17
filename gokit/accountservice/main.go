package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go_projects/accountservice/account"
)

func main() {
	var (
		dbName   = flag.String("db.name", "test_docs", "database name")
		dbType   = flag.String("db.type", "couch", "persistence type")
		httpPort = flag.String("http.port", ":8080", "HTTP listen port")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	db, err := account.InitDB(*dbName, *dbType)
	if err != nil {
		logger.Log("Error connecting to database, %s", err)
	}
	repo, _ := account.NewRepository(db)

	var s account.Service
	{
		s = account.NewService(repo)
		s = account.NewLoggingMiddleware(logger)(s)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("listening on port", *httpPort)
		h := account.MakeHTTPHandler(s)
		errs <- http.ListenAndServe(*httpPort, h)
	}()

	level.Error(logger).Log("exit", <-errs)
}
