import git from 'k6/x/git';

export default function () {
  let r = (Math.random() + 1).toString(36).substring(7);
  var folder = "myfolder" + r
  console.log(folder)
  // console.log("Debug test")
  //  pick one or both of the methods below depending on your use case - shallow clones
   console.log(git.plainCloneHTTP(folder,"https://github.com/dgzlopes/xk6-redis.git","myapikey",true,1))
  //console.log(git.plainCloneSSH(folder,"ssh://git@0.0.0.0:22222/xk6-git","/home/codespace/.ssh/id_rsa",false,1))
}
