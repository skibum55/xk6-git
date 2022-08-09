# /bin/bash
set -e

# download the ssh server
curl -sf https://gobinaries.com/vvatanabe/git-ssh-test-server| sh

# run the server

git-ssh-test-server -c config.yaml
