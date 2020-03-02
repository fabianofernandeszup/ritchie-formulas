package tree

type Formula struct {
	Path string
	Bin string
	Config *string
	RepoUrl string
}

type Command struct {
	Usage string
	Help string
	Formula *Formula
	Parent string
}

type Tree struct {
	Commands []Command
	Version string
}
