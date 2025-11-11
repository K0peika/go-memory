package cache

type Cache struct {
	cache map[string]any
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]any),
	}
}
func (c *Cache) Set(key string, value any) {
	c.cache[key] = value
}

func (c *Cache) Get(key string) any {
	if _, ok := c.cache[key]; ok {
		return c.cache[key]
	}
	return "key NOT found"
}

func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}
