package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TechBowl-japan/go-stations/db"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/handler/router"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func realMain() error {
	// config values
	const (
		defaultPort     = ":8080"
		defaultDBPath   = ".sqlite3/todo.db"
		defaultUserID   = "admin"
		defaultPassword = "password"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	userID := os.Getenv("BASIC_AUTH_USER_ID")
	if userID == "" {
		userID = defaultUserID
	}

	password := os.Getenv("BASIC_AUTH_PASSWORD")
	if password == "" {
		password = defaultPassword
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	defer todoDB.Close()

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	// 呼び出し順は、Recovery -> BoxOSInfo -> Logging(before) -> handler -> Logging(after) -> Recovery(defer)
	mux := middleware.Recovery(middleware.BoxOSInfo(middleware.Logging(router.NewRouter(todoDB, userID, password))))

	// TODO: サーバーをlistenする
	if err := http.ListenAndServe(defaultPort, mux); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
