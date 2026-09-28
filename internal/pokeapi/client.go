package pokeapi

import (
	"time"

	"github.com/dillon-goknit/pokedex/internal/pokecache"
)

type Client struct {
	cache *pokecache.Cache
}

func NewClient(interval time.Duration) *Client {
	return &Client{
		cache: pokecache.NewCache(interval),
	}
}
