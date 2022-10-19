package datasource

import (
	"math/rand"
	"sync"
	"time"
)

// Это надо (по-хорошему), конечно, обернуть в интерфейс и добавить еще методы чтения
// Но для учебы сгодится...
type Source struct {
	sync.Mutex

	First  int `json:"first"`
	Second int `json:"second"`
	Summa  int `json:"summa"`
}

var instance *Source
var once sync.Once

func (src *Source) SetFirst(iValue int) {
	src.Lock()
	defer src.Unlock()

	src.First = iValue
	src.Summa = src.First + src.Second
}

func (src *Source) SetSecond(iValue int) {
	src.Lock()
	defer src.Unlock()

	src.Second = iValue
	src.Summa = src.First + src.Second
}

// Source реализован как Singleton
func GetInstance() *Source {

	once.Do(func() {

		rand.Seed(time.Now().UnixNano())
		instance = &Source{
			First:  rand.Intn(100),
			Second: rand.Intn(100),
		}
		instance.Summa = instance.First + instance.Second
	})

	return instance
}
