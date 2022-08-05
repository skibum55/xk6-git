package git

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"

	"go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/git".
func init() {
	modules.Register("k6/x/git", new(GIT))
}

// GIT is the k6 extension for a Git client.
type GIT struct{}

// Repository is the GIT object we'll start with.
// type Repository struct {
// 	Storer storage.Storer
// }

// Clone gets a git repository
func (*GIT) PlainClone(directory string, url string) error {
	if directory == "" {
		directory = "~"
	}
	if url == "" {
		url = "ssh://git@localhost/test_repo.git"
	}

	var privateKeyFile = "~/.ssh/gitlab_rsa"
	var password string
	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, password)
	if err != nil {
		return err
	}
	_, err = git.PlainClone(directory, false, &git.CloneOptions{
		URL:             url,
		Auth:            publicKeys,
		InsecureSkipTLS: true,
	})
	if err != nil {
		return err
	} else {
		return err
	}
}

//CheckIfError(err)

// 	// ... retrieving the branch being pointed by HEAD
// 	ref, err := r.Head()
// 	CheckIfError(err)
// 	// ... retrieving the commit object
// 	commit, err := r.CommitObject(ref.Hash())
// 	CheckIfError(err)

// 	fmt.Println(commit)
// func (*GIT) Clone(storer *git.storer, fs, &git.CloneOptions) *git.Clone {
// 	URL: "https://github.com/git-fixtures/basic.git",
// }

// // Set adds a key/value
// func (*REDIS) Set(client *redis.Client, key string, value interface{}, expiration time.Duration) {
// 	// TODO: Make expiration configurable. Or document somewhere the unit.
// 	err := client.Set(context.Background(), key, value, expiration*time.Second).Err()
// 	if err != nil {
// 		ReportError(err, "Failed to set the specified key/value pair")
// 	}
// }
// 	_, err := os.Stat(privateKeyFile)
// 	if err != nil {
// 		Warning("read file %s failed %s\n", privateKeyFile, err.Error())
// 		return
// 	}

// // XClient represents the Client constructor (i.e. `new git.Client()`) and
// // returns a new Git client object.
// func (r *Redis) XClient(ctxPtr *context.Context, opts *redis.Options) interface{} {
//     rt := common.GetRuntime(*ctxPtr)
//     return common.Bind(rt, &Client{client: redis.NewClient(opts)}, ctxPtr)
// }

// // Set the given key with the given value and expiration time.
// func (c *Client) Set(key, value string, exp time.Duration) {
//     c.client.Set(c.client.Context(), key, value, exp)
// }

// // Get returns the value for the given key.
// func (c *Client) Get(key string) (string, error) {
//     res, err := c.client.Get(c.client.Context(), key).Result()
//     if err != nil {
//         return "", err
//     }
//     return res, nil
// }
