package main

import (
	"RestAPI/internal/config"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"os"
	"testing"
)

// Глобальная функция тестирования для «package main»
func TestMain(m *testing.M) {

	fmt.Println("setup package main ...")
	res := m.Run()
	fmt.Println("... tear-down package main")

	os.Exit(res)
}

func TestServerAPI(t *testing.T) {

	// Setup: GET JSON from "http:/localhost:8080/"
	cfg := config.GetConfig()

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/users", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		t.Fatal(err)
	}

	// JSON в качестве Content-Type.
	contentType := resp.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		t.Fatal(err)
	}

	if mediatype != "application/json" {
		t.Errorf("expect application/json Content-Type: %v", mediatype)
		return // *** А нужен ли тут "return"?
	}

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()

	/* TODO some kind of checks
	var req datasource.Source
	if err := dec.Decode(&req); err != nil {
		t.Fatal(err)
	}

	// *** вложенность в тестах
	// ReatAPI/Result
	t.Run("Result", func(t *testing.T) {
		if req.First+req.Second != req.Summa {
			t.Errorf("calculation error: %d + %d != %d", req.First, req.Second, req.Summa)
		} else {
			fmt.Printf("First: %d, Second: %d, Result: %d\n", req.First, req.Second, req.Summa)
		}
	})
	*/
}
