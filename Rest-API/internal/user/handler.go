package user

import (
	"RestAPI/internal/handlers"
	"RestAPI/pkg/datasource"
	"RestAPI/pkg/logging"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// подсказка для себя, т.е. check на соотвктсвие структуры handler интерфейсу Handler
var _ handlers.Handler = &handler{}

type handler struct {
	logger *logging.Logger
	source *datasource.Source
}

func NewHandler(logger *logging.Logger, source *datasource.Source) handlers.Handler {
	return &handler{
		logger: logger,
		source: source,
	}
}

const (
	rootURL  = "/"
	userURL1 = "/first"
	userURL2 = "/second"
	userURL3 = "/summa"
)

// Handlers
func (h *handler) GetIndex(w http.ResponseWriter, r *http.Request) {

	value, err := json.Marshal(h.source)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSON в качестве Content-Type.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(value)

	h.logger.Debug("JSON data transmitted")
}

func (h *handler) SetValues(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		Summa int `json:"summa"`
	}

	// JSON в качестве Content-Type.
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var src datasource.Source

	if err := dec.Decode(&src); err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.source.SetFirst(src.First)
	h.source.SetSecond(src.Second)

	// Ответ {"Summa" : число}
	resp := &Response{h.source.Summa}
	value, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) //200
	w.Header().Set("Content-Type", "application/json")
	w.Write(value)

	h.logger.Debug("data was successfully updated")
}

func (h *handler) GetFirst(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "First is: %d\n", h.source.First)

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("значение first"))
}

func (h *handler) GetSecond(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Second is: %d\n", h.source.Second)

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("значение second"))
}

func (h *handler) GetResult(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Result is: %d\n", h.source.Summa)

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("значение result"))
}

func (h *handler) Register(router *httprouter.Router) {

	// оригинальный -> router.GET("/", h.GetIndex)
	// *** вариант совместимости с дефолтным http-роутером
	router.HandlerFunc(http.MethodGet, rootURL, h.GetIndex)  // Показать JSON всего выражения
	router.HandlerFunc(http.MethodPut, rootURL, h.SetValues) // Установить first и second используя JSON

	// оригинальный -> router.GET(userURL1, h.GetFirst)
	router.HandlerFunc(http.MethodGet, userURL1, h.GetFirst)

	// оригинальный -> router.GET(userURL2, h.GetSecond)
	router.HandlerFunc(http.MethodGet, userURL2, h.GetSecond)

	// оригинальный -> router.GET(userURL3, h.GetResult) // результат
	router.HandlerFunc(http.MethodGet, userURL3, h.GetResult)
}
