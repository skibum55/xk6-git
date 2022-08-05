import git from 'k6/x/git';

const client = git.newClient();

export function setup() {
  git.set(client,"snake","camel",0)
  git.set(client,"foo",100,10)
}

export default function () {
  console.log(git.get(client,"snake"))
  console.log(gitt.get(client,"foo"))
  if (git.do(client,"PING","bzzz") == "bzzz"){
    console.log("PONG!")
  }
}

export function teardown () {
  git.del(client,"foo")
}
