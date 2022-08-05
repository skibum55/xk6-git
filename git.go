package git

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
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
func (*GIT) PlainClone(directory string, url string, privateKeyFile string) error {
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
