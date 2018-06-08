package connect

import "sync"

var manager sync.Map

func Add(key interface{}, ctx ConnContext) {
	manager.Store(key, ctx)
}

func Get(key interface{}) (*ConnContext, bool) {
	value, ok := manager.Load(key)
	if ok {
		ctx := value.(ConnContext)
		return &ctx, true
	}
	return nil, false
}

func Delete(key interface{}) {
	manager.Delete(key)
}

func Range(f func(key interface{}, ctx ConnContext) bool) {
	manager.Range(func(key, value interface{}) bool {
		f(key, value.(ConnContext))
		return true
	})
}
