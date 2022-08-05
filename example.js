import git from 'k6/x/git';

export function setup() {}

export default function () {
  // console.log(git.plainclone("",""))
  console.log(git.plainClone("","ssh://git@mlocalhost:2222/xk6-git"))
}

// export function teardown () {
//   git.del(client,"foo")
// }
