package session

import (
	"os"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestNewSession(t *testing.T) {
	os.Setenv("MEMECACHE_PORT", "localhost:11211")
	os.Setenv("MEMECACHE_NAME", "Test")
	os.Setenv("SESSION_SECRET", "secret")

	session := NewSession()

	name := session.Name
	if name != "Test" {
		t.Errorf("Session failed to set name to Test. It was %s instead", name)
	}

	item := &memcache.Item{Key: "foo", Value: []byte("bar")}
	session.Memcache.Client.Set(item)

	savedItem, err := session.Memcache.Client.Get(item.Key)
	if savedItem.Value == nil || err != nil {
		t.Errorf("Session failed to store item %v. Instead found %v", item, savedItem)
	}

	err = session.Memcache.Client.DeleteAll()
	if err != nil {
		t.Error("Session failed to delete item")
	}
}
