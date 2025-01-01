package pokeapi

import (
	"net/http"
	"time"

	"github.com/MechamJonathan/pokedexcli/internal/pokecache"
)

// client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// new client
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
