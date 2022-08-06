import git from 'k6/x/git';

export default function () {
  // console.log("Debug test")
  // console.log(git.plainClone("","","somewhere"))
  console.log(git.plainClone("","ssh://git@0.0.0.0:22222/xk6-git","/home/codespace/.ssh/id_rsa"))
}