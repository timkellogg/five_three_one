package session

import (
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/bradleypeabody/gorilla-sessions-memcache"
)

// Session - in memory storage
type Session struct {
	Name     string
	Memcache *gsm.MemcacheStore
}

// NewSession - create memecache instance
func NewSession() *Session {
	s := &Session{}

	port := os.Getenv("MEMECACHE_PORT")
	name := os.Getenv("MEMECACHE_NAME")
	token := os.Getenv("MEMECACHE_SECRET")

	client := memcache.New(port)

	s.Memcache = gsm.NewMemcacheStore(client, name, []byte(token))
	s.Name = name

	return s
}
