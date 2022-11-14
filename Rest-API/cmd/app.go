// *** REST API
// *** user-service

package main

import (
	"RestAPI/internal/config"
	"RestAPI/internal/user"
	"RestAPI/internal/user/db"
	mongodb "RestAPI/pkg/client"
	"RestAPI/pkg/logging"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

// main
func main() {

	// логгер
	logger := logging.GetLogger()

	//fmt.Println(params.First(), params.Second(), params.Result())

	//log.Println("create router")
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	cfgMongo := cfg.MongoDB
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username,
		cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Collection, logger)

	// *** START testing the database layer:
	// -> RestAPI/internal/user/db
	// -> RestAPI/pkg/client
	// -> RestAPI/pkg/logging

	// first user
	user1 := user.User{
		ID:           "",
		Email:        "bob_mail@mail.ru",
		Username:     "Boris",
		PasswordHash: "12345",
	}
	user1ID, err := storage.Create(context.Background(), user1)
	if err != nil {
		panic(err)
	}
	logger.Info(user1ID)
	// second user
	user2 := user.User{
		ID:           "",
		Email:        "second_mail@mail.ru",
		Username:     "Second",
		PasswordHash: "23456",
	}
	user2ID, err := storage.Create(context.Background(), user2)
	if err != nil {
		panic(err)
	}
	logger.Info(user2ID)
	// third user
	user3 := user.User{
		ID:           "",
		Email:        "third_mail@mail.ru",
		Username:     "Third",
		PasswordHash: "34567",
	}
	user3ID, err := storage.Create(context.Background(), user3)
	if err != nil {
		panic(err)
	}
	logger.Info(user3ID)
	// Testing method «FindOne»
	user2Found, err := storage.FindOne(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(user2Found)
	// Testing method «Update»
	user2Found.Email = "newEMail@here.ok"
	err = storage.Update(context.Background(), user2Found)
	if err != nil {
		panic(err)
	}
	// Testing method «Delete»
	err = storage.Delete(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}
	_, err = storage.FindOne(context.Background(), user2ID)
	if err != nil {
		fmt.Println(err)
	}
	users, err := storage.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// *** END testing the database layer

	//log.Println("register user handler")
	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {

	logger := logging.GetLogger()
	//log.Println("start application ")
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)

		logger.Infof("server is listening unix socket %s", socketPath)

	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		//log.Println("server is listening port 8080")
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//log.Fatalln(server.Serve(listener))
	logger.Fatal(server.Serve(listener))
}
