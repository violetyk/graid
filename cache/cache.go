package cache

// client-specified self pattern

type CacheEngine interface {
	WriteCache()
	ReadCache()
	DeleteCache()
}

type Cache struct {
}

func (self *Cache) Write(engine CacheEngine) {
	engine.WriteCache()

}

// func main() {
// engine := &FileEngine{}
// engine.Write(engine)
// }
