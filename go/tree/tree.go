package tree

import (
	"log"
	"slices"

	table "github.com/fullstack-lang/gongtable/go/models"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_stack "github.com/fullstack-lang/gongtree/go/stack"

	"github.com/fullstack-lang/maticons/maticons"
	"github.com/jinzhu/copier"

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

	buttonImplCategoryAddInstance = new(ButtonImplCategoryAddInstance[T])
	buttonImplCategoryAddInstance.category = category
	buttonImplCategoryAddInstance.treeWs = treeWs

	return
}

// new instance button has be pressed
func (buttonImplCategoryAddInstance *ButtonImplCategoryAddInstance[T]) ButtonUpdated(
	gongtreeStage *tree.StageStruct,
	stageButton, front *tree.Button) {

	scenario := buttonImplCategoryAddInstance.category.scenario
	stage := buttonImplCategoryAddInstance.treeWs.WeberStack.Stage
	probe := buttonImplCategoryAddInstance.treeWs.WeberStack.Probe

	// prepare the form for the newly created instance
	formStage := probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	var instance T
	switch any(instance).(type) {
	case *models.KingArthur:
		actorState := new(models.KingArthur).Stage(stage)
		scenario.KingArthurs = append(scenario.KingArthurs, actorState)
		stage.Commit()
		thelongbuild_probe.FillUpNamedFormFromGongstruct(
			actorState, probe, formStage, models.ModelForm.ToString())

	case *models.TheNuteTransition:

	case *models.Lancelot:
		parameter := new(models.Lancelot).Stage(stage)
		scenario.Lancelots = append(scenario.Lancelots, parameter)
		stage.Commit()
		thelongbuild_probe.FillUpNamedFormFromGongstruct(
			parameter, probe, formStage, models.ModelForm.ToString())

	case *models.BringYourDead:
		scenarioparameter := new(models.BringYourDead).Stage(stage)
		scenario.BringYourDeadarameters = append(scenario.BringYourDeadarameters, scenarioparameter)
		stage.Commit()
		thelongbuild_probe.FillUpNamedFormFromGongstruct(
			scenarioparameter, probe, formStage, models.ModelForm.ToString())

	case *models.GalahadThePure:
	case *models.SirRobin:
		diagram := new(models.SirRobin).Stage(stage)
		scenario.Diagrams = append(scenario.Diagrams, diagram)
		stage.Commit()
		thelongbuild_probe.FillUpNamedFormFromGongstruct(
			diagram, probe, formStage, models.ModelForm.ToString())
	}

	formStage.Commit()
	buttonImplCategoryAddInstance.treeWs.GenerateTree()
}

type ButtonImplDiagram struct {
	diagramNode *tree.Node

	// type of button
	Icon maticons.ButtonType
}

func NewButtonImplDiagram(classdiagramNode *tree.Node, icon maticons.ButtonType) (buttonImplDiagram *ButtonImplDiagram) {

	buttonImplDiagram = new(ButtonImplDiagram)
	buttonImplDiagram.diagramNode = classdiagramNode
	buttonImplDiagram.Icon = icon

	return
}

func (buttonImplDiagram *ButtonImplDiagram) ButtonUpdated(
	gongtreeStage *tree.StageStruct,
	stageButton, front *tree.Button) {

	diagramNodeImpl, ok := buttonImplDiagram.diagramNode.Impl.(*NodeImplDiagram)

	if !ok {
		log.Fatalln("not a diagram node")
	}

	switch buttonImplDiagram.Icon {
	case maticons.BUTTON_draw:
		diagramNodeImpl.diagram.IsInDrawMode = true
	case maticons.BUTTON_edit_off:
		diagramNodeImpl.diagram.IsInDrawMode = false
	case maticons.BUTTON_save:
		// the whole model is saved to the backend
		diagramNodeImpl.treeWs.WeberStack.Stage.CommitWithSuspendedCallbacks()
	}

	// reset the diagram
	diagramNodeImpl.treeWs.SVGGenerator.GenerateSVG(diagramNodeImpl.diagram)

	diagramNodeImpl.treeWs.ComputeNodesConf()
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

	buttonImplParameterFlip = new(ButtonImplParameterFlip)
	buttonImplParameterFlip.treeWs = treeWs
	buttonImplParameterFlip.diagramNode = parameterNode
	buttonImplParameterFlip.shapeWithDirection = shapeWithDirection

	return
}

func (buttonImplParameterFlip *ButtonImplParameterFlip) ButtonUpdated(
	gongtreeStage *tree.StageStruct,
	stageButton, front *tree.Button) {

	if buttonImplParameterFlip.shapeWithDirection.GetDirection() == models.DIRECTION_DOWN {
		buttonImplParameterFlip.shapeWithDirection.SetDirection(models.DIRECTION_UP)
	} else {
		buttonImplParameterFlip.shapeWithDirection.SetDirection(models.DIRECTION_DOWN)
	}

	buttonImplParameterFlip.treeWs.WeberStack.Stage.Commit()
	selectedDiagram := models.GetWorkspace(buttonImplParameterFlip.treeWs.WeberStack.Stage).SelectedDiagram
	buttonImplParameterFlip.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
}

func NewCategory[T ModelObject](
	instances *[]T,
	nodeFolded *bool, // pointer to the persistant boolean
	scenario *models.WhatIsYourPreferedColor,
) (category *Category[T]) {
	category = new(Category[T])
	category.instances = instances
	category.nodeFolded = nodeFolded
	category.scenario = scenario
	return category
}

// Category wraps instances of model objects
// in order to have genericity in tree operations
type Category[T ModelObject] struct {
	scenario   *models.WhatIsYourPreferedColor
	instances  *[]T
	nodeFolded *bool // pointer to the persistant boolean
}

func (category *Category[T]) GetNodeName() string {

	var instance T
	switch any(instance).(type) {
	case *models.KingArthur:
		return "Actor States"
	case *models.TheNuteTransition:
		return "Actor State Transitions"
	case *models.Lancelot:
		return "Parameters"
	case *models.BringYourDead:
		return "Scenario Parameters"
	case *models.GalahadThePure:
		return "Evolution Directions"
	case *models.SirRobin:
		return "Diagrams"
	default:
		log.Fatalln("no name for category")
	}
	return ""
}

func (category *Category[T]) GetIcon() *tree.SVGIcon {

	var instance T
	switch any(instance).(type) {
	case *models.KingArthur:
		return icons.Actor_stateIcon
	case *models.TheNuteTransition:
		return icons.ActorStateTransitionIcon
	case *models.Lancelot:
		return icons.Arrow_up_and_arrow_downIcon
	case *models.BringYourDead:
		return icons.Arrow_up_and_arrow_downIcon
	case *models.GalahadThePure:
		return icons.DirectionEvolutionIcon
	case *models.SirRobin:
		return icons.DiagramTimeEvolutionIcon
	default:
		log.Fatalln("no icon for category")
	}
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

	log.Println("Node clicked", old.GetName())

	// in case this is just the unfolding / folding of node, do nothing
	if old.IsExpanded != updatedNode.IsExpanded {
		old.IsExpanded = updatedNode.IsExpanded

		nodeCallback.instance.SetIsNodeExpanded(updatedNode.IsExpanded)

		nodeCallback.treeWs.WeberStack.Stage.Commit()
		return
	}

	formStage := nodeCallback.treeWs.WeberStack.Probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	nodeCallback.FillUpForm(
		nodeCallback.instance, nodeCallback.treeWs.WeberStack.Probe,
		formStage,
		models.ModelForm.ToString())

	switch inst := any(nodeCallback.instance).(type) {
	case *models.SirRobin:
		diagram := inst
		nodeCallback.treeWs.SVGGenerator.GenerateSVG(diagram)
	}
	// OnAfterUpdate( genericNodeCallback.instance, stage, old, new)
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

	log.Println("NodeActorState clicked", stagedNode.GetName())

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		stagedNode.Commit(nodeImplActorState.treeWs.TreeStack.Stage)

		// add the actor state to the diagam
		actorStateShape :=
			(&models.KingArthurShape{
				Name:       nodeImplActorState.actorState.GetName(),
				ActorState: nodeImplActorState.actorState,
			}).
				Stage(nodeImplActorState.treeWs.WeberStack.Stage)

		workspace := models.GetWorkspace(nodeImplActorState.treeWs.WeberStack.Stage)

		selectedDiagram := workspace.SelectedDiagram
		selectedDiagram.Arthurs = append(selectedDiagram.Arthurs, actorStateShape)

		// set up default values

		defaultShape := workspace.D
		if defaultShape == nil {
			log.Fatalln("No default ActorState Shape")
		}

		copier.CopyWithOption(actorStateShape, defaultShape, copier.Option{IgnoreEmpty: true, DeepCopy: false})
		actorStateShape.Name = nodeImplActorState.actorState.GetName()
		actorStateShape.ActorState = nodeImplActorState.actorState

		nodeImplActorState.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(nodeImplActorState.treeWs.TreeStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplActorState.treeWs.WeberStack.Stage).SelectedDiagram
		for idx, actorStateShape := range selectedDiagram.Arthurs {
			if actorStateShape.ActorState == nodeImplActorState.actorState {
				selectedDiagram.Arthurs = slices.Delete(selectedDiagram.Arthurs, idx, idx+1)
				actorStateShape.Unstage(nodeImplActorState.treeWs.WeberStack.Stage)

				// one have to remove "from / to" transtions shapes
				var actorStateShapeTransitionShapeForRemoval []*models.TheNuteShape
				for _, actorStateShapeTransition := range selectedDiagram.TheNuteShapes {
					if actorStateShapeTransition.ActorStateTransition.StartState == actorStateShape.ActorState ||
						actorStateShapeTransition.ActorStateTransition.EndState == actorStateShape.ActorState {
						actorStateShapeTransition.Unstage(nodeImplActorState.treeWs.WeberStack.Stage)
						actorStateShapeTransitionShapeForRemoval = append(actorStateShapeTransitionShapeForRemoval, actorStateShapeTransition)
					}
				}
				for _, actorStateShapeTransitionShape := range actorStateShapeTransitionShapeForRemoval {
					idx := slices.Index(selectedDiagram.TheNuteShapes, actorStateShapeTransitionShape)
					selectedDiagram.TheNuteShapes =
						slices.Delete(selectedDiagram.TheNuteShapes, idx, idx+1)
				}

				continue
			}
		}

		// slices.Delete(selectedDiagram.ActorStateShapes, )
		nodeImplActorState.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// in any cases, have the form editor set up with the instance
	formStage := nodeImplActorState.treeWs.WeberStack.Probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	// if the node is checked, fill up the form for the shape
	if frontNode.IsChecked {
		selectedDiagram := models.GetWorkspace(nodeImplActorState.treeWs.WeberStack.Stage).SelectedDiagram
		for _, actorstateShape := range selectedDiagram.Arthurs {
			if actorstateShape.ActorState == nodeImplActorState.actorState {
				thelongbuild_probe.FillUpNamedFormFromGongstruct(
					actorstateShape,
					nodeImplActorState.treeWs.WeberStack.Probe,
					formStage,
					models.ShapeForm.ToString())
				continue
			}
		}
	}

	// fill up the form for the model object
	thelongbuild_probe.FillUpNamedFormFromGongstruct(
		nodeImplActorState.actorState,
		nodeImplActorState.treeWs.WeberStack.Probe,
		formStage,
		models.ModelForm.ToString())

	// recompute nodes conf
	nodeImplActorState.treeWs.ComputeNodesConf()
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

	log.Println("NodeActorStateTransition clicked", stagedNode.GetName())

	// in case this is just the unfolding / folding of node, do nothing
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		stagedNode.Commit(nodeImplActorStateTransition.treeWs.TreeStack.Stage)

		// add the actor state to the diagam
		actorStateTransitionShape :=
			(&models.TheNuteShape{
				Name:                 nodeImplActorStateTransition.actorStateTransition.GetName(),
				ActorStateTransition: nodeImplActorStateTransition.actorStateTransition,
			}).
				Stage(nodeImplActorStateTransition.treeWs.WeberStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplActorStateTransition.treeWs.WeberStack.Stage).SelectedDiagram
		selectedDiagram.TheNuteShapes = append(selectedDiagram.TheNuteShapes, actorStateTransitionShape)

		actorStateTransitionShape.Name = nodeImplActorStateTransition.actorStateTransition.GetName()
		actorStateTransitionShape.ActorStateTransition = nodeImplActorStateTransition.actorStateTransition
		actorStateTransitionShape.StartOrientation = models.ORIENTATION_HORIZONTAL
		actorStateTransitionShape.EndOrientation = models.ORIENTATION_HORIZONTAL
		actorStateTransitionShape.StartRatio = 0.5
		actorStateTransitionShape.EndRatio = 0.5
		actorStateTransitionShape.CornerOffsetRatio = 2

		// find the start and end shape
		for _, actorStateShape := range selectedDiagram.Arthurs {
			if actorStateShape.ActorState == nodeImplActorStateTransition.actorStateTransition.StartState {
				actorStateTransitionShape.Start = actorStateShape
			}
			if actorStateShape.ActorState == nodeImplActorStateTransition.actorStateTransition.EndState {
				actorStateTransitionShape.End = actorStateShape
			}
		}

		nodeImplActorStateTransition.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(nodeImplActorStateTransition.treeWs.TreeStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplActorStateTransition.treeWs.WeberStack.Stage).SelectedDiagram
		for idx, actorStateTransitionShape := range selectedDiagram.TheNuteShapes {
			if actorStateTransitionShape.ActorStateTransition == nodeImplActorStateTransition.actorStateTransition {
				selectedDiagram.TheNuteShapes = slices.Delete(selectedDiagram.TheNuteShapes, idx, idx+1)
				actorStateTransitionShape.Unstage(nodeImplActorStateTransition.treeWs.WeberStack.Stage)

				// one have to remove "from / to" transtions shapes
				var actorStateShapeTransitionShapeForRemoval []*models.TheNuteShape
				for _, actorStateShapeTransition := range selectedDiagram.TheNuteShapes {
					if actorStateShapeTransition.ActorStateTransition.StartState == actorStateTransitionShape.ActorStateTransition.StartState ||
						actorStateShapeTransition.ActorStateTransition.EndState == actorStateTransitionShape.ActorStateTransition.EndState {
						actorStateShapeTransition.Unstage(nodeImplActorStateTransition.treeWs.WeberStack.Stage)
						actorStateShapeTransitionShapeForRemoval = append(actorStateShapeTransitionShapeForRemoval, actorStateShapeTransition)
					}
				}
				for _, actorStateShapeTransitionShape := range actorStateShapeTransitionShapeForRemoval {
					idx := slices.Index(selectedDiagram.TheNuteShapes, actorStateShapeTransitionShape)
					selectedDiagram.TheNuteShapes =
						slices.Delete(selectedDiagram.TheNuteShapes, idx, idx+1)
				}

				continue
			}
		}

		// slices.Delete(selectedDiagram.ActorStateTransitionShapes, )
		nodeImplActorStateTransition.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// in any cases, have the form editor set up with the instance
	formStage := nodeImplActorStateTransition.treeWs.WeberStack.Probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	// if the node is checked, fill up the form for the shape
	if frontNode.IsChecked {
		selectedDiagram := models.GetWorkspace(nodeImplActorStateTransition.treeWs.WeberStack.Stage).SelectedDiagram
		for _, actorstatetransitionShape := range selectedDiagram.TheNuteShapes {
			if actorstatetransitionShape.ActorStateTransition == nodeImplActorStateTransition.actorStateTransition {
				thelongbuild_probe.FillUpNamedFormFromGongstruct(
					actorstatetransitionShape,
					nodeImplActorStateTransition.treeWs.WeberStack.Probe,
					formStage,
					models.ShapeForm.ToString())
				continue
			}
		}
	}

	// fill up the form for the model object
	thelongbuild_probe.FillUpNamedFormFromGongstruct(
		nodeImplActorStateTransition.actorStateTransition,
		nodeImplActorStateTransition.treeWs.WeberStack.Probe,
		formStage,
		models.ModelForm.ToString())

	// recompute nodes conf
	nodeImplActorStateTransition.treeWs.ComputeNodesConf()
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

	log.Println("NodeCategory clicked", stagedNode.GetName())

	// in case this is just the unfolding / folding of node, do nothing
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded

		nodeImplCategory.category.SetIsNodeFolded(frontNode.IsExpanded)
		nodeImplCategory.treeWs.WeberStack.Stage.Commit()
		return
	}

	// recompute nodes conf
	nodeImplCategory.treeWs.ComputeNodesConf()
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

	log.Println("Node clicked", stagedNode.GetName())

	// in case this is just the unfolding / folding of node, do nothing
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		stagedNode.Commit(nodeImplDiagram.treeWs.TreeStack.Stage)

		// fetch the workbench and set up the selected diagram
		workspace := models.GetWorkspace(nodeImplDiagram.treeWs.WeberStack.Stage)
		workspace.SelectedDiagram = nodeImplDiagram.diagram

		for _, otherDiagramNode := range nodeImplDiagram.treeWs.AllDiagramNodes {
			if otherDiagramNode == stagedNode {
				continue
			}

			// uncheck the other node
			if otherDiagramNode.IsChecked {
				// log.Println("Node " + node.Name + " is checked and should be unchecked")
				otherDiagramNode.IsChecked = false
				otherDiagramNode.Commit(nodeImplDiagram.treeWs.TreeStack.Stage)
			}
		}

		// commit models stage here to encode the workspace selected diagram/scenario
		// change, this will generate a new svg drawing
		nodeImplDiagram.treeWs.WeberStack.Stage.Commit()
	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(nodeImplDiagram.treeWs.TreeStack.Stage)
	}

	formStage := nodeImplDiagram.treeWs.WeberStack.Probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	thelongbuild_probe.FillUpNamedFormFromGongstruct(
		nodeImplDiagram.diagram, nodeImplDiagram.treeWs.WeberStack.Probe,
		formStage,
		models.ModelForm.ToString())

	// recompute nodes conf
	nodeImplDiagram.treeWs.ComputeNodesConf()
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

	log.Println("NodeEvolutionDirection clicked", stagedNode.GetName())

	// in case this is just the unfolding / folding of node, do nothing
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		stagedNode.Commit(nodeImplEvolutionDirection.treeWs.TreeStack.Stage)

		// add shape to the diagam
		evolutionDirectionShape :=
			(&models.AWitch{
				Name: nodeImplEvolutionDirection.evolutionDirection.GetName(),
			}).
				Stage(nodeImplEvolutionDirection.treeWs.WeberStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplEvolutionDirection.treeWs.WeberStack.Stage).SelectedDiagram
		selectedDiagram.Witches = append(selectedDiagram.Witches, evolutionDirectionShape)

		// set up default values
		workspace := models.GetWorkspace(nodeImplEvolutionDirection.treeWs.WeberStack.Stage)
		defaultShape := workspace.A
		if defaultShape == nil {
			log.Fatalln("No default Evolution Direction Shape")
		}

		copier.CopyWithOption(evolutionDirectionShape, defaultShape, copier.Option{IgnoreEmpty: true, DeepCopy: false})
		evolutionDirectionShape.Name = nodeImplEvolutionDirection.evolutionDirection.GetName()
		evolutionDirectionShape.EvolutionDirection = nodeImplEvolutionDirection.evolutionDirection

		nodeImplEvolutionDirection.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(nodeImplEvolutionDirection.treeWs.TreeStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplEvolutionDirection.treeWs.WeberStack.Stage).SelectedDiagram
		for idx, evolutionDirectionShape := range selectedDiagram.Witches {
			if evolutionDirectionShape.EvolutionDirection == nodeImplEvolutionDirection.evolutionDirection {
				selectedDiagram.Witches = slices.Delete(selectedDiagram.Witches, idx, idx+1)
				evolutionDirectionShape.Unstage(nodeImplEvolutionDirection.treeWs.WeberStack.Stage)
				continue
			}
		}

		// slices.Delete(selectedDiagram.EvolutionDirectionShapes, )
		nodeImplEvolutionDirection.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// in any cases, have the form editor set up with the instance
	formStage := nodeImplEvolutionDirection.treeWs.WeberStack.Probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	// if the node is checked, fill up the form for the shape
	if frontNode.IsChecked {
		selectedDiagram := models.GetWorkspace(nodeImplEvolutionDirection.treeWs.WeberStack.Stage).SelectedDiagram
		for _, evolutionDirectionShape := range selectedDiagram.Witches {
			if evolutionDirectionShape.EvolutionDirection == nodeImplEvolutionDirection.evolutionDirection {
				thelongbuild_probe.FillUpNamedFormFromGongstruct(
					evolutionDirectionShape,
					nodeImplEvolutionDirection.treeWs.WeberStack.Probe,
					formStage,
					models.ShapeForm.ToString())
				continue
			}
		}
	}

	// fill up the form for the model object
	thelongbuild_probe.FillUpNamedFormFromGongstruct(
		nodeImplEvolutionDirection.evolutionDirection,
		nodeImplEvolutionDirection.treeWs.WeberStack.Probe,
		formStage,
		models.ModelForm.ToString())

	// recompute nodes conf
	nodeImplEvolutionDirection.treeWs.ComputeNodesConf()
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

	log.Println("NodeParameter clicked", stagedNode.GetName())

	// in case this is just the unfolding / folding of node, do nothing
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		stagedNode.Commit(nodeImplParameter.treeWs.TreeStack.Stage)

		// add shape to the diagam
		parameterShape :=
			(&models.KnightWhoSayNi{
				Name: nodeImplParameter.parameter.GetName(),
			}).
				Stage(nodeImplParameter.treeWs.WeberStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplParameter.treeWs.WeberStack.Stage).SelectedDiagram
		selectedDiagram.KnightWhoSayNis = append(selectedDiagram.KnightWhoSayNis, parameterShape)

		// set up default values
		workspace := models.GetWorkspace(nodeImplParameter.treeWs.WeberStack.Stage)
		defaultShape := workspace.B
		if defaultShape == nil {
			log.Fatalln("No default Parameter Shape")
		}

		copier.CopyWithOption(parameterShape, defaultShape, copier.Option{IgnoreEmpty: true, DeepCopy: false})
		parameterShape.Name = nodeImplParameter.parameter.GetName()
		parameterShape.Parameter = nodeImplParameter.parameter

		nodeImplParameter.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(nodeImplParameter.treeWs.TreeStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplParameter.treeWs.WeberStack.Stage).SelectedDiagram
		for idx, parameterShape := range selectedDiagram.KnightWhoSayNis {
			if parameterShape.Parameter == nodeImplParameter.parameter {
				selectedDiagram.KnightWhoSayNis = slices.Delete(selectedDiagram.KnightWhoSayNis, idx, idx+1)
				parameterShape.Unstage(nodeImplParameter.treeWs.WeberStack.Stage)
				continue
			}
		}

		// slices.Delete(selectedDiagram.ParameterShapes, )
		nodeImplParameter.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// in any cases, have the form editor set up with the instance
	formStage := nodeImplParameter.treeWs.WeberStack.Probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	// if the node is checked, fill up the form for the shape
	if frontNode.IsChecked {
		selectedDiagram := models.GetWorkspace(nodeImplParameter.treeWs.WeberStack.Stage).SelectedDiagram
		for _, parameterShape := range selectedDiagram.KnightWhoSayNis {
			if parameterShape.Parameter == nodeImplParameter.parameter {
				thelongbuild_probe.FillUpNamedFormFromGongstruct(
					parameterShape,
					nodeImplParameter.treeWs.WeberStack.Probe,
					formStage,
					models.ShapeForm.ToString())
				continue
			}
		}
	}

	// fill up the form for the model object
	thelongbuild_probe.FillUpNamedFormFromGongstruct(
		nodeImplParameter.parameter,
		nodeImplParameter.treeWs.WeberStack.Probe,
		formStage,
		models.ModelForm.ToString())

	// recompute nodes conf
	nodeImplParameter.treeWs.ComputeNodesConf()
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

	log.Println("NodeScenario clicked", stagedNode.GetName())

	// in case this is just the unfolding / folding of node, do nothing
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded

		nodeImplScenario.scenario.IsNodeExpanded = frontNode.IsExpanded
		nodeImplScenario.scenario.Commit(nodeImplScenario.treeWs.WeberStack.Stage)
		return
	}

	formStage := nodeImplScenario.treeWs.WeberStack.Probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	// in any cases, have the form editor set up with the instance
	thelongbuild_probe.FillUpNamedFormFromGongstruct(
		nodeImplScenario.scenario,
		nodeImplScenario.treeWs.WeberStack.Probe,
		formStage,
		models.ModelForm.ToString())

	// recompute nodes conf
	nodeImplScenario.treeWs.ComputeNodesConf()
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

	log.Println("NodeScenarioParameter clicked", stagedNode.GetName())

	// in case this is just the unfolding / folding of node, do nothing
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		stagedNode.Commit(nodeImplScenarioParameter.treeWs.TreeStack.Stage)

		// add shape to the diagam
		scenarioparameterShape :=
			(&models.BlackKnightShape{
				Name: nodeImplScenarioParameter.scenarioParameter.GetName(),
			}).
				Stage(nodeImplScenarioParameter.treeWs.WeberStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplScenarioParameter.treeWs.WeberStack.Stage).SelectedDiagram
		selectedDiagram.BlackKnightShapes = append(selectedDiagram.BlackKnightShapes, scenarioparameterShape)

		// set up default values
		workspace := models.GetWorkspace(nodeImplScenarioParameter.treeWs.WeberStack.Stage)
		defaultShape := workspace.C
		if defaultShape == nil {
			log.Fatalln("No default ScenarioParameter Shape")
		}

		copier.CopyWithOption(scenarioparameterShape, defaultShape, copier.Option{IgnoreEmpty: true, DeepCopy: false})
		scenarioparameterShape.Name = nodeImplScenarioParameter.scenarioParameter.GetName()
		scenarioparameterShape.BringYourDead = nodeImplScenarioParameter.scenarioParameter

		nodeImplScenarioParameter.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// node was checked and user wants to uncheck it. This is not possible
	// from a application logic point of view
	// on need to commit the staged node for the front to reconstruct
	// the node as checked and overides the unchecking action
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stagedNode.Commit(nodeImplScenarioParameter.treeWs.TreeStack.Stage)

		selectedDiagram := models.GetWorkspace(nodeImplScenarioParameter.treeWs.WeberStack.Stage).SelectedDiagram
		for idx, scenarioparameterShape := range selectedDiagram.BlackKnightShapes {
			if scenarioparameterShape.BringYourDead == nodeImplScenarioParameter.scenarioParameter {
				selectedDiagram.BlackKnightShapes = slices.Delete(selectedDiagram.BlackKnightShapes, idx, idx+1)
				scenarioparameterShape.Unstage(nodeImplScenarioParameter.treeWs.WeberStack.Stage)
				continue
			}
		}

		// slices.Delete(selectedDiagram.ScenarioParameterShapes, )
		nodeImplScenarioParameter.treeWs.SVGGenerator.GenerateSVG(selectedDiagram)
	}

	// in any cases, have the form editor set up with the instance
	formStage := nodeImplScenarioParameter.treeWs.WeberStack.Probe.GetFormStage()
	formStage.Reset()
	formStage.Commit()

	// if the node is checked, fill up the form for the shape
	if frontNode.IsChecked {
		selectedDiagram := models.GetWorkspace(nodeImplScenarioParameter.treeWs.WeberStack.Stage).SelectedDiagram
		for _, scenarioparameterShape := range selectedDiagram.BlackKnightShapes {
			if scenarioparameterShape.BringYourDead == nodeImplScenarioParameter.scenarioParameter {
				thelongbuild_probe.FillUpNamedFormFromGongstruct(
					scenarioparameterShape,
					nodeImplScenarioParameter.treeWs.WeberStack.Probe,
					formStage,
					models.ShapeForm.ToString())
				continue
			}
		}
	}

	// fill up the form for the model object
	thelongbuild_probe.FillUpNamedFormFromGongstruct(
		nodeImplScenarioParameter.scenarioParameter,
		nodeImplScenarioParameter.treeWs.WeberStack.Probe,
		formStage,
		models.ModelForm.ToString())

	// recompute nodes conf
	nodeImplScenarioParameter.treeWs.ComputeNodesConf()
}

// TreeWs, for tree workspace holds the supporting data for performing operation on the weber tree
type TreeWs struct {
	WeberStack *thelongbuild_stack.Stack

	TreeStack *gongtree_stack.Stack

	NodeTree *tree.Tree

	// AllDiagramNodes is a pratical way to access all diagrams nodes
	// tree
	AllDiagramNodes []*tree.Node

	// SVGGenerator is the callback for generating the SVG
	SVGGenerator GenerateSVGInterface
}

type GenerateSVGInterface interface {
	GenerateSVG(diagram *models.SirRobin)
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

func (treeWs *TreeWs) computeNodeConfOfScenarioCateforyObjects(scenarioNodeChildren *tree.Node) {

	scenarioNodeChildren.Buttons = scenarioNodeChildren.Buttons[:0]

	addButton := (&tree.Button{
		Name: "Add instance for: " + scenarioNodeChildren.Name + string(gongtree_buttons.BUTTON_add),
		Icon: string(gongtree_buttons.BUTTON_add)}).Stage(treeWs.TreeStack.Stage)

	scenarioNodeChildren.Buttons = append(scenarioNodeChildren.Buttons, addButton)

	switch impl := scenarioNodeChildren.Impl.(type) {
	case *NodeImplCategory[*models.KingArthur]:
		addButton.Impl = NewButtonImplCategoryAddInstance(impl.category, treeWs)
	case *NodeImplCategory[*models.Lancelot]:
		addButton.Impl = NewButtonImplCategoryAddInstance(impl.category, treeWs)
	case *NodeImplCategory[*models.BringYourDead]:
		addButton.Impl = NewButtonImplCategoryAddInstance(impl.category, treeWs)
	case *NodeImplCategory[*models.GalahadThePure]:
		addButton.Impl = NewButtonImplCategoryAddInstance(impl.category, treeWs)
	case *NodeImplCategory[*models.SirRobin]:
		addButton.Impl = NewButtonImplCategoryAddInstance(impl.category, treeWs)
	}
}

func (treeWs *TreeWs) computeNodeConfOfScenarioObjects(scenarioNodeChildren *tree.Node,
	isScenarioSelected bool, inModificationMode bool,
	map_actorStateInDiagram map[*models.KingArthur]any,
	map_actorStateTransitionInDiagram map[*models.TheNuteTransition]any,
	map_evolutionDirectionInDiagram map[*models.GalahadThePure]any,
	map_parameterInDiagram map[*models.Lancelot]*models.KnightWhoSayNi,
	map_scenarioParameterInDiagram map[*models.BringYourDead]*models.BlackKnightShape) {

	//
	// First parse all instances nodes and set the enability of
	// the check button
	//
	switch scenarioNodeChildren.Impl.(type) {
	case *NodeImplActorState,
		*NodeImplActorStateTransition,
		*NodeImplEvolutionDirection,
		*NodeImplParameter,
		*NodeImplScenarioParameter:

		// node is disabled if model element shape is not in the scenario or
		// if the diagram is not in modification
		scenarioNodeChildren.IsCheckboxDisabled = !isScenarioSelected || !inModificationMode

		// log.Println(
		// 	"computeNodeConfOfScenarioObjects, IsCheckboxDisabled", scenarioNodeChildren.IsCheckboxDisabled,
		// 	"computeNodeConfOfScenarioObjects", scenarioNodeChildren.Name)

		scenarioNodeChildren.IsChecked = false
	}

	switch impl := scenarioNodeChildren.Impl.(type) {

	//
	// Nodes for instances below scenario
	//
	case *NodeImplActorState:
		if _, ok := map_actorStateInDiagram[impl.actorState]; ok {
			scenarioNodeChildren.IsChecked = true
		}
	case *NodeImplActorStateTransition:
		if _, ok := map_actorStateTransitionInDiagram[impl.actorStateTransition]; ok {
			scenarioNodeChildren.IsChecked = true
		}
	case *NodeImplEvolutionDirection:
		if _, ok := map_evolutionDirectionInDiagram[impl.evolutionDirection]; ok {
			scenarioNodeChildren.IsChecked = true
		}
	case *NodeImplParameter:
		if parameterShape, ok := map_parameterInDiagram[impl.parameter]; ok {
			scenarioNodeChildren.IsChecked = true

			scenarioNodeChildren.Buttons = scenarioNodeChildren.Buttons[:0]

			// add button for flip influence arrows directions  if in modification mode
			if isScenarioSelected && inModificationMode {
				newButton := new(tree.Button).Stage(treeWs.TreeStack.Stage)
				newButton.Name = "Flip Direction for: " + scenarioNodeChildren.Name

				newButton.SVGIcon = icons.Flip_directionIcon.Stage(treeWs.TreeStack.Stage)

				scenarioNodeChildren.Buttons = append(scenarioNodeChildren.Buttons, newButton)

				newButton.Impl = NewButtonImplParameterFlip(treeWs, scenarioNodeChildren, parameterShape)
			}
		}
	case *NodeImplScenarioParameter:
		if scenarioParameterShape, ok := map_scenarioParameterInDiagram[impl.scenarioParameter]; ok {
			scenarioNodeChildren.IsChecked = true

			scenarioNodeChildren.Buttons = scenarioNodeChildren.Buttons[:0]

			// add button for flip influence arrows directions  if in modification mode
			if isScenarioSelected && inModificationMode {
				newButton := new(tree.Button).Stage(treeWs.TreeStack.Stage)
				newButton.Name = "Flip Direction for: " + scenarioNodeChildren.Name

				newButton.SVGIcon = icons.Flip_directionIcon.Stage(treeWs.TreeStack.Stage)

				scenarioNodeChildren.Buttons = append(scenarioNodeChildren.Buttons, newButton)

				newButton.Impl = NewButtonImplParameterFlip(treeWs, scenarioNodeChildren, scenarioParameterShape)
			}
		}
	}
}

// addNewButtonToDiagramInstance add a button if conditionForButton is true
func addNewButtonToDiagramInstance(
	gongtreeStage *tree.StageStruct,
	diagramNode *tree.Node,
	icon maticons.ButtonType, conditionForButton bool) {

	if !conditionForButton {
		return
	}

	diagramNodeImpl := diagramNode.Impl.(*NodeImplDiagram)

	newButton := (&tree.Button{
		Name: diagramNodeImpl.diagram.Name + " " + string(icon),
		Icon: string(icon)}).Stage(gongtreeStage)

	newButton.Impl = NewButtonImplDiagram(diagramNode, icon)

	diagramNode.Buttons = append(diagramNode.Buttons, newButton)
}

// ComputeNodesConf parses all nodes and:
// - for diagram nodes if the diagram is not in edit mode, display an edit button, else display a save button
func (treeWs *TreeWs) ComputeNodesConf() {

	//
	// COMPUTE EACH DIAGRAM NODE STATUS
	//

	// compute wether one of the diagrams is in draw/edit mode
	// if so, all diagram check need to be disabled
	var inModificationMode bool

	for _, diagramNode := range treeWs.AllDiagramNodes {
		diagramNodeImpl, ok := diagramNode.Impl.(*NodeImplDiagram)
		if !ok {
			log.Fatalln("not a good interface")
		}
		// log.Println("ComputeNodesConf: diagram name", diagramNodeImpl.diagram.Name, "IsInDrawMode", diagramNodeImpl.diagram.IsInDrawMode)
		if diagramNodeImpl.diagram.IsInDrawMode {
			inModificationMode = true
		}
	}

	var selectedDiagram *models.SirRobin

	// pass for adding the buttons on the diagram that
	// is checked
	for _, diagramNode := range treeWs.AllDiagramNodes {

		// remove all buttons of the node
		for _, _button := range diagramNode.Buttons {
			_button.Unstage(treeWs.TreeStack.Stage)
		}
		diagramNode.Buttons = make([]*tree.Button, 0)

		diagramNodeImpl, ok := diagramNode.Impl.(*NodeImplDiagram)

		if !ok {
			log.Fatalln("not a diagram node")
		}

		// compute the condition on the node diagrams that will determine wether one
		// button or another is necessary
		addAnEditButton := diagramNode.IsChecked && !diagramNodeImpl.diagram.IsInDrawMode
		addNewButtonToDiagramInstance(treeWs.TreeStack.Stage, diagramNode, maticons.BUTTON_draw, addAnEditButton)

		addNewButtonToDiagramInstance(treeWs.TreeStack.Stage, diagramNode, maticons.BUTTON_edit_off, diagramNodeImpl.diagram.IsInDrawMode)
		addNewButtonToDiagramInstance(treeWs.TreeStack.Stage, diagramNode, maticons.BUTTON_save, diagramNodeImpl.diagram.IsInDrawMode)

		// disable checkbox is one diagram is currently being edited
		diagramNode.IsCheckboxDisabled = inModificationMode

		if diagramNode.IsChecked {
			selectedDiagram = diagramNodeImpl.diagram
		}
	}

	//
	// Make make of model elements whose shape is in the diagram
	// (in order to compute the check status)
	//
	map_actorStateInDiagram := make(map[*models.KingArthur]any, 0)
	if selectedDiagram != nil {
		for _, actorStateShape := range selectedDiagram.Arthurs {
			map_actorStateInDiagram[actorStateShape.ActorState] = true
		}
	}
	map_actorStateTransitionInDiagram := make(map[*models.TheNuteTransition]any, 0)
	if selectedDiagram != nil {
		for _, actorStateTransitionShape := range selectedDiagram.TheNuteShapes {
			map_actorStateTransitionInDiagram[actorStateTransitionShape.ActorStateTransition] = true
		}
	}
	map_evolutionDirectionInDiagram := make(map[*models.GalahadThePure]any, 0)
	if selectedDiagram != nil {
		for _, evolutionDirectionShape := range selectedDiagram.Witches {
			map_evolutionDirectionInDiagram[evolutionDirectionShape.EvolutionDirection] = true
		}
	}
	map_parameterInDiagram := make(map[*models.Lancelot]*models.KnightWhoSayNi, 0)
	if selectedDiagram != nil {
		for _, parameterShape := range selectedDiagram.KnightWhoSayNis {
			map_parameterInDiagram[parameterShape.Parameter] = parameterShape
		}
	}
	map_scenarioParameterInDiagram := make(map[*models.BringYourDead]*models.BlackKnightShape, 0)
	if selectedDiagram != nil {
		for _, scenarioparameterShape := range selectedDiagram.BlackKnightShapes {
			map_scenarioParameterInDiagram[scenarioparameterShape.BringYourDead] = scenarioparameterShape
		}
	}

	selectedScenario := models.GetWorkspace(treeWs.WeberStack.Stage).GetPreferedColor(treeWs.WeberStack.Stage)
	for _, analysisNode := range treeWs.NodeTree.RootNodes {
		for _, scenarioNode := range analysisNode.Children {

			nodeImplScenario, ok := scenarioNode.Impl.(*NodeImplScenario)
			if !ok {
				continue
			}
			isScenarioSelected := selectedScenario != nil && nodeImplScenario.scenario == selectedScenario
			for _, scenarioNodeChildren := range scenarioNode.Children {

				treeWs.computeNodeConfOfScenarioCateforyObjects(scenarioNodeChildren)

				for _, scenarioNodeGrandChildren := range scenarioNodeChildren.Children {
					treeWs.computeNodeConfOfScenarioObjects(scenarioNodeGrandChildren, isScenarioSelected, inModificationMode,
						map_actorStateInDiagram,
						map_actorStateTransitionInDiagram,
						map_evolutionDirectionInDiagram,
						map_parameterInDiagram,
						map_scenarioParameterInDiagram)
				}
			}
		}
	}

	//
	// COMMIT CHANGES
	//
	treeWs.TreeStack.Stage.Commit()
}

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
