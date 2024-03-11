package database

import "sync"

var instance map[string]interface{}
var once sync.Once

func GetInstance() map[string]interface{} {
	once.Do(func() {
		instance = make(map[string]interface{})
	})

	return instance
}
