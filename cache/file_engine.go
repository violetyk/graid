package cache

// embed Cache
type FileEngine struct {
	*Cache
}

// implements CacheEngine
func (self *Cache) WriteCache() {

}
