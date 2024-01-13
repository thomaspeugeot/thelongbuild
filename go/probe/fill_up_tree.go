package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
)

func fillUpTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(probe.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "AWitch":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AWitch](probe.stageOfInterest)
			for _awitch := range set {
				nodeInstance := (&tree.Node{Name: _awitch.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_awitch, "AWitch", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "BlackKnightShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.BlackKnightShape](probe.stageOfInterest)
			for _blackknightshape := range set {
				nodeInstance := (&tree.Node{Name: _blackknightshape.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_blackknightshape, "BlackKnightShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "BringYourDead":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.BringYourDead](probe.stageOfInterest)
			for _bringyourdead := range set {
				nodeInstance := (&tree.Node{Name: _bringyourdead.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_bringyourdead, "BringYourDead", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Document":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Document](probe.stageOfInterest)
			for _document := range set {
				nodeInstance := (&tree.Node{Name: _document.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_document, "Document", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DocumentUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DocumentUse](probe.stageOfInterest)
			for _documentuse := range set {
				nodeInstance := (&tree.Node{Name: _documentuse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_documentuse, "DocumentUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GalahadThePure":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GalahadThePure](probe.stageOfInterest)
			for _galahadthepure := range set {
				nodeInstance := (&tree.Node{Name: _galahadthepure.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_galahadthepure, "GalahadThePure", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GeoObject":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GeoObject](probe.stageOfInterest)
			for _geoobject := range set {
				nodeInstance := (&tree.Node{Name: _geoobject.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_geoobject, "GeoObject", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GeoObjectUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GeoObjectUse](probe.stageOfInterest)
			for _geoobjectuse := range set {
				nodeInstance := (&tree.Node{Name: _geoobjectuse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_geoobjectuse, "GeoObjectUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Group":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Group](probe.stageOfInterest)
			for _group := range set {
				nodeInstance := (&tree.Node{Name: _group.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_group, "Group", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GroupUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GroupUse](probe.stageOfInterest)
			for _groupuse := range set {
				nodeInstance := (&tree.Node{Name: _groupuse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_groupuse, "GroupUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "KingArthur":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.KingArthur](probe.stageOfInterest)
			for _kingarthur := range set {
				nodeInstance := (&tree.Node{Name: _kingarthur.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_kingarthur, "KingArthur", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "KingArthurShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.KingArthurShape](probe.stageOfInterest)
			for _kingarthurshape := range set {
				nodeInstance := (&tree.Node{Name: _kingarthurshape.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_kingarthurshape, "KingArthurShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "KnightWhoSayNi":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.KnightWhoSayNi](probe.stageOfInterest)
			for _knightwhosayni := range set {
				nodeInstance := (&tree.Node{Name: _knightwhosayni.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_knightwhosayni, "KnightWhoSayNi", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Lancelot":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Lancelot](probe.stageOfInterest)
			for _lancelot := range set {
				nodeInstance := (&tree.Node{Name: _lancelot.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lancelot, "Lancelot", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LancelotAgregation":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LancelotAgregation](probe.stageOfInterest)
			for _lancelotagregation := range set {
				nodeInstance := (&tree.Node{Name: _lancelotagregation.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lancelotagregation, "LancelotAgregation", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LancelotAgregationUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LancelotAgregationUse](probe.stageOfInterest)
			for _lancelotagregationuse := range set {
				nodeInstance := (&tree.Node{Name: _lancelotagregationuse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lancelotagregationuse, "LancelotAgregationUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LancelotCategory":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LancelotCategory](probe.stageOfInterest)
			for _lancelotcategory := range set {
				nodeInstance := (&tree.Node{Name: _lancelotcategory.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lancelotcategory, "LancelotCategory", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LancelotCategoryUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LancelotCategoryUse](probe.stageOfInterest)
			for _lancelotcategoryuse := range set {
				nodeInstance := (&tree.Node{Name: _lancelotcategoryuse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_lancelotcategoryuse, "LancelotCategoryUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MapObject":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MapObject](probe.stageOfInterest)
			for _mapobject := range set {
				nodeInstance := (&tree.Node{Name: _mapobject.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_mapobject, "MapObject", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MapObjectUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MapObjectUse](probe.stageOfInterest)
			for _mapobjectuse := range set {
				nodeInstance := (&tree.Node{Name: _mapobjectuse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_mapobjectuse, "MapObjectUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Repository":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Repository](probe.stageOfInterest)
			for _repository := range set {
				nodeInstance := (&tree.Node{Name: _repository.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_repository, "Repository", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SirRobin":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SirRobin](probe.stageOfInterest)
			for _sirrobin := range set {
				nodeInstance := (&tree.Node{Name: _sirrobin.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_sirrobin, "SirRobin", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TheBridge":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TheBridge](probe.stageOfInterest)
			for _thebridge := range set {
				nodeInstance := (&tree.Node{Name: _thebridge.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_thebridge, "TheBridge", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TheNuteShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TheNuteShape](probe.stageOfInterest)
			for _thenuteshape := range set {
				nodeInstance := (&tree.Node{Name: _thenuteshape.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_thenuteshape, "TheNuteShape", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TheNuteTransition":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TheNuteTransition](probe.stageOfInterest)
			for _thenutetransition := range set {
				nodeInstance := (&tree.Node{Name: _thenutetransition.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_thenutetransition, "TheNuteTransition", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "User":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.User](probe.stageOfInterest)
			for _user := range set {
				nodeInstance := (&tree.Node{Name: _user.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_user, "User", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "UserUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.UserUse](probe.stageOfInterest)
			for _useruse := range set {
				nodeInstance := (&tree.Node{Name: _useruse.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_useruse, "UserUse", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "WhatIsYourPreferedColor":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.WhatIsYourPreferedColor](probe.stageOfInterest)
			for _whatisyourpreferedcolor := range set {
				nodeInstance := (&tree.Node{Name: _whatisyourpreferedcolor.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_whatisyourpreferedcolor, "WhatIsYourPreferedColor", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Workspace":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Workspace](probe.stageOfInterest)
			for _workspace := range set {
				nodeInstance := (&tree.Node{Name: _workspace.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_workspace, "Workspace", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
