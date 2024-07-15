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

type NodeImplActorStateTransition struct {
	treeWs               *TreeWs
	actorStateTransition *models.TheNuteTransition

	// IsInDrawMode is true if the actorState is being edited
	IsInDrawMode bool
}

func (nodeImplActorStateTransition *NodeImplActorStateTransition) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
}

func NewNodeImplDiagram(
	treeWs *TreeWs,
	diagram *models.SirRobin,
	scenario *models.WhatIsYourPreferedColor,
) (nodeImplDiagram *NodeImplDiagram) {
	nodeImplDiagram = new(NodeImplDiagram)
	nodeImplDiagram.treeWs = treeWs
	nodeImplDiagram.diagram = diagram
	return
}

type NodeImplDiagram struct {
	treeWs  *TreeWs
	diagram *models.SirRobin
}

func (nodeImplDiagram *NodeImplDiagram) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
}

func NewNodeImplEvolutionDirection(
	treeWs *TreeWs,
	evolutionDirection *models.GalahadThePure,
) (nodeImplEvolutionDirection *NodeImplEvolutionDirection) {
	nodeImplEvolutionDirection = new(NodeImplEvolutionDirection)
	nodeImplEvolutionDirection.treeWs = treeWs
	nodeImplEvolutionDirection.evolutionDirection = evolutionDirection
	return
}

type NodeImplEvolutionDirection struct {
	treeWs             *TreeWs
	evolutionDirection *models.GalahadThePure

	// IsInDrawMode is true if the evolutionDirection is being edited
	IsInDrawMode bool
}

func (nodeImplEvolutionDirection *NodeImplEvolutionDirection) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
}

func NewNodeImplParameter(
	treeWs *TreeWs,
	parameter *models.Lancelot,
) (nodeImplParameter *NodeImplParameter) {
	nodeImplParameter = new(NodeImplParameter)
	nodeImplParameter.treeWs = treeWs
	nodeImplParameter.parameter = parameter
	return
}

type NodeImplParameter struct {
	treeWs    *TreeWs
	parameter *models.Lancelot
}

func (nodeImplParameter *NodeImplParameter) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
}

func NewNodeImplScenario(
	treeWs *TreeWs,
	scenario *models.WhatIsYourPreferedColor,
) (nodeImplScenario *NodeImplScenario) {
	nodeImplScenario = new(NodeImplScenario)
	nodeImplScenario.treeWs = treeWs
	nodeImplScenario.scenario = scenario
	return
}

type NodeImplScenario struct {
	treeWs   *TreeWs
	scenario *models.WhatIsYourPreferedColor

	// IsInDrawMode is true if the scenario is being edited
	IsInDrawMode bool
}

func (nodeImplScenario *NodeImplScenario) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
}

func NewNodeImplScenarioParameter(
	treeWs *TreeWs,
	scenarioparameter *models.BringYourDead,
) (nodeImplScenarioParameter *NodeImplScenarioParameter) {
	nodeImplScenarioParameter = new(NodeImplScenarioParameter)
	nodeImplScenarioParameter.treeWs = treeWs
	nodeImplScenarioParameter.scenarioParameter = scenarioparameter
	return
}

type NodeImplScenarioParameter struct {
	treeWs            *TreeWs
	scenarioParameter *models.BringYourDead
}

func (nodeImplScenarioParameter *NodeImplScenarioParameter) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
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
