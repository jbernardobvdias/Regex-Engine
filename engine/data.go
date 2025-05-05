package engine

type Data struct {
	Input   string
	Pattern string
}

type TreeNode struct {
	Children []*TreeNode
	Info     string
}
