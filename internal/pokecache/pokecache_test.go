package pokecache

import (
	"fmt"
	"testing"
	"time"

	"github.com/lukasz0707/pokedexcli/internal/utility"
	"github.com/stretchr/testify/assert"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val utility.LocationAreaResp
	}{
		{
			key: "https://example.com",
			val: utility.LocationAreaResp{
				Count: 2,
			},
		},
		{
			key: "https://example.com/path",
			val: utility.LocationAreaResp{
				Count: 10,
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			assert.Equal(t, val, c.val)
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", utility.LocationAreaResp{Count: 5})

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
