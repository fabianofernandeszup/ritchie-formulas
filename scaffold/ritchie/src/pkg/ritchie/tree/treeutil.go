package tree

//"usage": "port-forward",
//"help": "Create k8s port forward",
//"formula": {
//"path": "kubernetes/core",
//"bin": "kubernetes-${so}",
//"config": "port-config.json",
//"repoUrl": "http://ritchie-cli-bucket152849730126474.s3-website-sa-east-1.amazonaws.com/formulas/kubernetes/core"
//},
//"parent": "root_k8s_create"

type Formula struct {
	Path string
	Bin string
	Config string
	RepoUrl string
}

type Command struct {
	Usage string
	Help string
	Formula Formula
	Parent string
}

type Tree struct {
	Commands []Command
	Version string
}
