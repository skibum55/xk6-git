import git from 'k6/x/git';

export function setup() {}

export default function () {
  // console.log(git.plainclone("",""))
  console.log(git.plainClone("","ssh://git@localhost:2222/xk6-git","/home/codespace/.ssh/gitlab_rsa"))
}

// export function teardown () {
//   git.del(client,"foo")
// }
