package connect

import "sync"

var manager sync.Map

func Add(key string, ctx ConnContext) {
	manager.Store(key, ctx)
}

func Get(key string) (*ConnContext, bool) {
	value, ok := manager.Load(key)
	if ok {
		ctx := value.(ConnContext)
		return &ctx, true
	}
	return nil, false
}

func Delete(key string) {
	manager.Delete(key)
}

func Range(f func(key string, ctx ConnContext) bool) {
	manager.Range(func(key, value interface{}) bool {
		f(key.(string), value.(ConnContext))
		return true
	})
}
