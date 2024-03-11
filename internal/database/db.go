package database

import "sync"

//nolint:gochecknoglobals  // because this should be singletone
var (
	instance map[string]interface{}
	once     sync.Once
)

func GetInstance() map[string]interface{} {
	once.Do(func() {
		instance = make(map[string]interface{})
	})

	return instance
}
