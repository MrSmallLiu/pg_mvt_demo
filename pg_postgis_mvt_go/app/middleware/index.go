package middleware

import "sync"

var once sync.Once

func init() {
	once.Do(func() {
		initCORS()
	})
}