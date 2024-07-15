package tree

import (
	"log"

	table "github.com/fullstack-lang/gongtable/go/models"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gongtree_stack "github.com/fullstack-lang/gongtree/go/stack"

	"github.com/thomaspeugeot/thelongbuild/go/models"

	thelongbuild_probe "github.com/thomaspeugeot/thelongbuild/go/probe"
)

type ModelObject interface {
	GetName() string
}

func NewNodeImpl[T2 models.GenericNode[T], T models.Gongstruct](
	treeWs *TreeWs,
	instance *T,
) (nodeCallback *NodeImpl[T2, T]) {
	nodeCallback = new(NodeImpl[T2, T])
	nodeCallback.treeWs = treeWs
	nodeCallback.instance = instance

	// the following line is the real culprit
	nodeCallback.FillUpForm = thelongbuild_probe.FillUpNamedFormFromGongstruct[T]

	log.Println("fdsqfsdfd")
	return
}

type NodeImpl[T2 models.GenericNode[T], T models.Gongstruct] struct {
	treeWs   *TreeWs
	instance T2

	FillUpForm func(*T, *thelongbuild_probe.Probe, *table.StageStruct, string)
}

func (nodeCallback *NodeImpl[T2, T]) OnAfterUpdate(stage *tree.StageStruct, old, updatedNode *tree.Node) {
}

// TreeWs, for tree workspace holds the supporting data for performing operation on the weber tree
type TreeWs struct {
	TreeStack *gongtree_stack.Stack

	NodeTree *tree.Tree

	// AllDiagramNodes is a pratical way to access all diagrams nodes
	// tree
	AllDiagramNodes []*tree.Node
}

func (treeWs *TreeWs) GenerateTree() {

	bridge := new(models.TheBridge)
	analysisNode := new(tree.Node).Stage(treeWs.TreeStack.Stage)
	analysisNode.Impl = NewNodeImpl(treeWs, bridge)
}
