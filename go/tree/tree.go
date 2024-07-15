package tree

import (
	"slices"

	table "github.com/fullstack-lang/gongtable/go/models"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gongtree_stack "github.com/fullstack-lang/gongtree/go/stack"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/thomaspeugeot/thelongbuild/go/icons"
	"github.com/thomaspeugeot/thelongbuild/go/models"
	thelongbuild_stack "github.com/thomaspeugeot/thelongbuild/go/stack"

	thelongbuild_probe "github.com/thomaspeugeot/thelongbuild/go/probe"
)

type ButtonImplCategoryAddInstance[T ModelObject] struct {
	diagramNode *tree.Node
	treeWs      *TreeWs
	category    *Category[T]
}

func NewButtonImplCategoryAddInstance[T ModelObject](
	category *Category[T],
	treeWs *TreeWs,
) (buttonImplCategoryAddInstance *ButtonImplCategoryAddInstance[T]) {
	return
}

// new instance button has be pressed
func (buttonImplCategoryAddInstance *ButtonImplCategoryAddInstance[T]) ButtonUpdated(
	gongtreeStage *tree.StageStruct,
	stageButton, front *tree.Button) {
}

type ButtonImplDiagram struct {
	diagramNode *tree.Node

	// type of button
	Icon maticons.ButtonType
}

func NewButtonImplDiagram(classdiagramNode *tree.Node, icon maticons.ButtonType) (buttonImplDiagram *ButtonImplDiagram) {
	return
}

func (buttonImplDiagram *ButtonImplDiagram) ButtonUpdated(
	gongtreeStage *tree.StageStruct,
	stageButton, front *tree.Button) {
}

type ButtonImplParameterFlip struct {
	treeWs             *TreeWs
	diagramNode        *tree.Node
	shapeWithDirection models.ShapeWithDirection
}

func NewButtonImplParameterFlip(
	treeWs *TreeWs,
	parameterNode *tree.Node,
	shapeWithDirection models.ShapeWithDirection) (buttonImplParameterFlip *ButtonImplParameterFlip) {
	return
}

func (buttonImplParameterFlip *ButtonImplParameterFlip) ButtonUpdated(
	gongtreeStage *tree.StageStruct,
	stageButton, front *tree.Button) {
}

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
	*models.KingArthur | *models.GalahadThePure | *models.Lancelot | *models.BringYourDead | *models.TheNuteTransition | *models.SirRobin
	GetName() string
}

func NewNodeImpl[T2 models.GenericNode[T], T models.Gongstruct](
	treeWs *TreeWs,
	instance *T,
) (nodeCallback *NodeImpl[T2, T]) {
	nodeCallback = new(NodeImpl[T2, T])
	nodeCallback.treeWs = treeWs
	nodeCallback.instance = instance

	nodeCallback.FillUpForm = thelongbuild_probe.FillUpNamedFormFromGongstruct[T]

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
	WeberStack *thelongbuild_stack.Stack

	TreeStack *gongtree_stack.Stack

	NodeTree *tree.Tree

	// AllDiagramNodes is a pratical way to access all diagrams nodes
	// tree
	AllDiagramNodes []*tree.Node
}

func NewTreeWs(
	Weber *thelongbuild_stack.Stack,
	Tree *gongtree_stack.Stack,
) (treeWs *TreeWs) {
	treeWs = new(TreeWs)

	treeWs.WeberStack = Weber
	treeWs.TreeStack = Tree

	return
}

// ComputeNodesConf parses all nodes and:
// - for diagram nodes if the diagram is not in edit mode, display an edit button, else display a save button
func (treeWs *TreeWs) ComputeNodesConf() {}

func (treeWs *TreeWs) GenerateTree() {

	// reset all diagram nodes
	treeWs.AllDiagramNodes = make([]*tree.Node, 0)

	treeWs.TreeStack.Stage.Reset()

	// create new tree
	treeWs.NodeTree = new(tree.Tree).Stage(treeWs.TreeStack.Stage)
	treeWs.NodeTree.Name = string(models.Sidebar)

	list := models.GetGongstrucsSorted[*models.TheBridge](treeWs.WeberStack.Stage)

	// reverse map for transition to actor state shapes
	assoc1 := models.GetAssociationName[models.TheNuteTransition]().StartState.Name
	map_Transition_StartActorStateShape := models.GetPointerReverseMap[
		models.KingArthurShape,
		models.TheNuteTransition,
	](assoc1, treeWs.WeberStack.Stage)
	_ = map_Transition_StartActorStateShape
	assoc2 := models.GetAssociationName[models.TheNuteTransition]().EndState.Name
	map_Transition_EndActorStateShape := models.GetPointerReverseMap[
		models.KingArthurShape,
		models.TheNuteTransition,
	](assoc2, treeWs.WeberStack.Stage)
	_ = map_Transition_EndActorStateShape

	for _, analysis := range list {
		analysisNode := new(tree.Node).Stage(treeWs.TreeStack.Stage)
		analysisNode.Name = analysis.Name
		analysisNode.IsWithPreceedingIcon = true
		analysisNode.PreceedingIcon = string(maticons.BUTTON_library_books)
		analysisNode.IsExpanded = analysis.GetIsNodeExpanded()

		analysisNode.IsNodeClickable = true
		analysisNode.Impl = NewNodeImpl(treeWs, analysis)

		treeWs.NodeTree.RootNodes = append(treeWs.NodeTree.RootNodes, analysisNode)

		slices.SortFunc(analysis.WhatIsYourPreferedColor, models.CompareGongstructByName)
		for _, scenario := range analysis.WhatIsYourPreferedColor {
			scenarioNode := new(tree.Node).Stage(treeWs.TreeStack.Stage)
			scenarioNode.Name = scenario.Name
			scenarioNode.IsExpanded = true
			scenarioNode.IsWithPreceedingIcon = false
			// scenarioNode.PreceedingIcon = string(maticons.BUTTON_theater_comedy)
			scenarioNode.PreceedingSVGIcon = icons.ScenarioIcon.Stage(treeWs.TreeStack.Stage)

			scenarioNode.IsNodeClickable = true
			scenarioNode.Impl = NewNodeImplScenario(treeWs, scenario)

			analysisNode.Children = append(analysisNode.Children, scenarioNode)

			generateTreeForCategory(treeWs, scenarioNode,
				NewCategory(&scenario.Diagrams,
					&scenario.DiagramsNodeFolded,
					scenario,
				))

			generateTreeForCategory(treeWs, scenarioNode,
				NewCategory(&scenario.Galahard,
					&scenario.IIUU,
					scenario,
				))

			generateTreeForCategory(treeWs, scenarioNode,
				NewCategory(&scenario.Lancelots,
					&scenario.LancelotsodeFolded,
					scenario,
				))

			generateTreeForCategory(treeWs, scenarioNode,
				NewCategory(&scenario.BringYourDeadarameters,
					&scenario.RRRRT,
					scenario,
				))

			generateTreeForCategory(treeWs, scenarioNode,
				NewCategory(&scenario.KingArthurs,
					&scenario.KingArthurNodeFolded,
					scenario,
				))

			generateTreeForCategory(treeWs, scenarioNode,
				NewCategory(&scenario.Nutes,
					&scenario.RRRR,
					scenario,
				))

		}
	}
	treeWs.ComputeNodesConf()

	// add a node for the Add Analysis
	analysisNode := new(tree.Node).Stage(treeWs.TreeStack.Stage)
	analysisNode.Name = "New Analysis"
	analysisNode.FontStyle = tree.ITALIC
	analysisNode.IsWithPreceedingIcon = true
	analysisNode.PreceedingIcon = string(maticons.BUTTON_library_books)
	treeWs.NodeTree.RootNodes = append(treeWs.NodeTree.RootNodes, analysisNode)

	treeWs.TreeStack.Stage.Commit()
	treeWs.TreeStack.Probe.Refresh()
}

func generateTreeForCategory[T ModelObject](
	treeWs *TreeWs,
	scenarioNode *tree.Node,
	category *Category[T],
) {

	categorysNode := new(tree.Node).Stage(treeWs.TreeStack.Stage)
	categorysNode.Name = category.GetNodeName()
	categorysNode.IsExpanded = *category.nodeFolded
	categorysNode.IsWithPreceedingIcon = false
	categorysNode.PreceedingSVGIcon = category.GetIcon().Stage(treeWs.TreeStack.Stage)
	categorysNode.Impl = NewNodeImplCategory(treeWs, category)
	scenarioNode.Children = append(scenarioNode.Children, categorysNode)

	for _, instance := range *category.instances {

		instanceNode := new(tree.Node).Stage(treeWs.TreeStack.Stage)
		instanceNode.Name = instance.GetName()
		instanceNode.IsExpanded = true
		instanceNode.IsWithPreceedingIcon = false
		instanceNode.PreceedingSVGIcon = category.GetIcon().Stage(treeWs.TreeStack.Stage)

		categorysNode.Children = append(categorysNode.Children, instanceNode)

		instanceNode.IsNodeClickable = true

		switch _inst := any(instance).(type) {
		case *models.KingArthur:
			instanceNode.Impl = NewNodeImplActorState(treeWs, _inst)
		case *models.Lancelot:
			instanceNode.Impl = NewNodeImplParameter(treeWs, _inst)
		case *models.BringYourDead:
			instanceNode.Impl = NewNodeImplScenarioParameter(treeWs, _inst)
		case *models.TheNuteTransition:
			instanceNode.Impl = NewNodeImplActorStateTransition(treeWs, _inst)
		case *models.GalahadThePure:
			instanceNode.Impl = NewNodeImplEvolutionDirection(treeWs, _inst)
		case *models.SirRobin:
			selectedDiagram := models.GetWorkspace(treeWs.WeberStack.Stage).SelectedDiagram
			if selectedDiagram == _inst {
				instanceNode.IsChecked = true
			}
			treeWs.AllDiagramNodes = append(treeWs.AllDiagramNodes, instanceNode)

			instanceNode.Impl = NewNodeImplDiagram(treeWs, _inst,
				models.GetWorkspace(treeWs.WeberStack.Stage).
					GetPreferedColor(treeWs.WeberStack.Stage))
		}

		instanceNode.HasCheckboxButton = true
	}
}
