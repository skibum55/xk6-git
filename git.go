package git

import (
    "context"
    "time"

     git "github.com/go-git/go-git/v5"

    "go.k6.io/k6/js/common"
    "go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/git".
func init() {
    modules.Register("k6/x/git", new(Git))
}

// is the k6 extension for a Git client.
type Git struct{}

// Client is the Redis client wrapper.
type Client struct {
    client *Git.Client
}

// XClient represents the Client constructor (i.e. `new git.Client()`) and
// returns a new Git client object.
func (r *Redis) XClient(ctxPtr *context.Context, opts *redis.Options) interface{} {
    rt := common.GetRuntime(*ctxPtr)
    return common.Bind(rt, &Client{client: redis.NewClient(opts)}, ctxPtr)
}

// Set the given key with the given value and expiration time.
func (c *Client) Set(key, value string, exp time.Duration) {
    c.client.Set(c.client.Context(), key, value, exp)
}

// Get returns the value for the given key.
func (c *Client) Get(key string) (string, error) {
    res, err := c.client.Get(c.client.Context(), key).Result()
    if err != nil {
        return "", err
    }
    return res, nil
}
