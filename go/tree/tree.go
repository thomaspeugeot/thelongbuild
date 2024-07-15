package tree

import (
	"log"

	table "github.com/fullstack-lang/gongtable/go/models"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gongtree_stack "github.com/fullstack-lang/gongtree/go/stack"

	"github.com/thomaspeugeot/thelongbuild/go/models"

	thelongbuild_probe "github.com/thomaspeugeot/thelongbuild/go/probe"
)

func NewCategory[T ModelObject](
	instances *[]T,
	nodeFolded *bool, // pointer to the persistant boolean
	scenario *models.WhatIsYourPreferedColor,
) (category *Category[T]) {
	return nil
}

// Category wraps instances of model objects
// in order to have genericity in tree operations
type Category[T ModelObject] struct {
	scenario   *models.WhatIsYourPreferedColor
	instances  *[]T
	nodeFolded *bool // pointer to the persistant boolean
}

func (category *Category[T]) GetNodeName() string {
	return ""
}

func (category *Category[T]) GetIcon() *tree.SVGIcon {
	return nil
}

func (category *Category[T]) GetIsNodeFolded() bool {
	return *category.nodeFolded
}

func (category *Category[T]) SetIsNodeFolded(val bool) {
	*category.nodeFolded = val
}

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

func NewNodeImplActorState(
	treeWs *TreeWs,
	actorState *models.KingArthur,
) (nodeImplActorState *NodeImplActorState) {
	nodeImplActorState = new(NodeImplActorState)
	nodeImplActorState.treeWs = treeWs
	nodeImplActorState.actorState = actorState
	return
}

type NodeImplActorState struct {
	treeWs     *TreeWs
	actorState *models.KingArthur

	// IsInDrawMode is true if the actorState is being edited
	IsInDrawMode bool
}

func (nodeImplActorState *NodeImplActorState) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
}

func NewNodeImplActorStateTransition(
	treeWs *TreeWs,
	actorStateTransition *models.TheNuteTransition,
) (nodeImplActorStateTransition *NodeImplActorStateTransition) {
	nodeImplActorStateTransition = new(NodeImplActorStateTransition)
	nodeImplActorStateTransition.treeWs = treeWs
	nodeImplActorStateTransition.actorStateTransition = actorStateTransition
	return
}

type NodeImplActorStateTransition struct {
	treeWs               *TreeWs
	actorStateTransition *models.TheNuteTransition

	// IsInDrawMode is true if the actorState is being edited
	IsInDrawMode bool
}

func (nodeImplActorStateTransition *NodeImplActorStateTransition) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
}

func NewNodeImplCategory[T ModelObject](
	treeWs *TreeWs,
	category *Category[T],
) (nodeImplCategory *NodeImplCategory[T]) {
	nodeImplCategory = new(NodeImplCategory[T])
	nodeImplCategory.treeWs = treeWs
	nodeImplCategory.category = category
	return
}

type NodeImplCategory[T ModelObject] struct {
	treeWs   *TreeWs
	category *Category[T]
	scenario *models.WhatIsYourPreferedColor

	// IsInDrawMode is true if the category is being edited
	IsInDrawMode bool
}

func (nodeImplCategory *NodeImplCategory[T]) OnAfterUpdate(stage *tree.StageStruct, stagedNode, frontNode *tree.Node) {
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
