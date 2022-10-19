/*
## Задача № 1
Написать API для указанных маршрутов(endpoints)
"/first"  // Случайное число
"/second" // Случайное число
"/summa"  // Сумма двух случайных чисел

результат вернуть в виде JSON

"math/rand"
number := rand.Intn(100)

## Задача № 2

сайт: postman.com
Установить программу и создать учетную запись.

*/

package main

import (
	"RestAPI/internal/config"
	"RestAPI/internal/user"
	"RestAPI/pkg/datasource"
	"RestAPI/pkg/logging"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	// Использован httprouter Джулиана Шмидта
	// https://github.com/julienschmidt/httprouter
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

	logger.Info("generating data for an application")
	src := datasource.GetInstance()

	//log.Println("register user handler")
	logger.Info("register user handler")
	handler := user.NewHandler(logger, src)
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
