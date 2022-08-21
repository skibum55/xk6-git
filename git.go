package git

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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

// Clone gets a git repository using ssh and a private key
// exported function for javascript - uppercase first letter
func (*GIT) PlainCloneSSH(directory string, url string, privateKeyFile string, skiptls bool, depth int) error {
	if len(directory) == 0 {
		directory = "~"
	}
	if len(url) == 0 {
		url = "ssh://git@localhost/test_repo.git"
	}

	if len(privateKeyFile) == 0 {
		home := os.Getenv("HOME")
		privateKeyFile = path.Join(home, ".ssh/id_rsa")
	}

	var password string
	publicKeys, err1 := ssh.NewPublicKeysFromFile("git", privateKeyFile, password)
	if err1 != nil {
		return err1
	}
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:  url,
		Auth: publicKeys,
		// RemoteName:        "",
		// ReferenceName:     "",
		// SingleBranch:      false,
		// NoCheckout:        false,
		Depth: depth,
		// RecurseSubmodules: 0,
		// Progress:          nil,
		// Tags:              0,
		InsecureSkipTLS: skiptls,
		// CABundle:          []byte{},
	})
	if err != nil {
		return err
	} else {
		return err
	}
}

// Clone gets a git repository using ssh and a private key
// exported function for javascript - uppercase first letter
func (*GIT) PlainCloneHTTP(directory string, url string, token string, skiptls bool, depth int) error {
	if len(directory) == 0 {
		directory = "~"
	}
	if len(url) == 0 {
		url = "ssh://git@localhost/test_repo.git"
	}

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: "xk6-git", // yes, this can be anything except an empty string
			Password: token,
		},
		URL:             url,
		Progress:        os.Stdout,
		Depth:           depth,
		InsecureSkipTLS: skiptls,
	})
	if err != nil {
		return err
	} else {
		return err
	}
}
