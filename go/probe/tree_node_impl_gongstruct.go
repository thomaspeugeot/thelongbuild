// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe      *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "AWitch" {
		fillUpTable[models.AWitch](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "BlackKnightShape" {
		fillUpTable[models.BlackKnightShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "BringYourDead" {
		fillUpTable[models.BringYourDead](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Document" {
		fillUpTable[models.Document](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DocumentUse" {
		fillUpTable[models.DocumentUse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GalahadThePure" {
		fillUpTable[models.GalahadThePure](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GeoObject" {
		fillUpTable[models.GeoObject](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GeoObjectUse" {
		fillUpTable[models.GeoObjectUse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Group" {
		fillUpTable[models.Group](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GroupUse" {
		fillUpTable[models.GroupUse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "KingArthur" {
		fillUpTable[models.KingArthur](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "KingArthurShape" {
		fillUpTable[models.KingArthurShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "KnightWhoSayNi" {
		fillUpTable[models.KnightWhoSayNi](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Lancelot" {
		fillUpTable[models.Lancelot](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LancelotAgregation" {
		fillUpTable[models.LancelotAgregation](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LancelotAgregationUse" {
		fillUpTable[models.LancelotAgregationUse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LancelotCategory" {
		fillUpTable[models.LancelotCategory](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LancelotCategoryUse" {
		fillUpTable[models.LancelotCategoryUse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MapObject" {
		fillUpTable[models.MapObject](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MapObjectUse" {
		fillUpTable[models.MapObjectUse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Repository" {
		fillUpTable[models.Repository](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SirRobin" {
		fillUpTable[models.SirRobin](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TheBridge" {
		fillUpTable[models.TheBridge](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TheNuteShape" {
		fillUpTable[models.TheNuteShape](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TheNuteTransition" {
		fillUpTable[models.TheNuteTransition](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "User" {
		fillUpTable[models.User](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "UserUse" {
		fillUpTable[models.UserUse](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "WhatIsYourPreferedColor" {
		fillUpTable[models.WhatIsYourPreferedColor](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Workspace" {
		fillUpTable[models.Workspace](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
