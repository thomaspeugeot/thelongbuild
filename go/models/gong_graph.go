// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AWitch:
		ok = stage.IsStagedAWitch(target)

	case *BlackKnightShape:
		ok = stage.IsStagedBlackKnightShape(target)

	case *BringYourDead:
		ok = stage.IsStagedBringYourDead(target)

	case *Document:
		ok = stage.IsStagedDocument(target)

	case *DocumentUse:
		ok = stage.IsStagedDocumentUse(target)

	case *GalahadThePure:
		ok = stage.IsStagedGalahadThePure(target)

	case *GeoObject:
		ok = stage.IsStagedGeoObject(target)

	case *GeoObjectUse:
		ok = stage.IsStagedGeoObjectUse(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *GroupUse:
		ok = stage.IsStagedGroupUse(target)

	case *KingArthur:
		ok = stage.IsStagedKingArthur(target)

	case *KingArthurShape:
		ok = stage.IsStagedKingArthurShape(target)

	case *KnightWhoSayNi:
		ok = stage.IsStagedKnightWhoSayNi(target)

	case *Lancelot:
		ok = stage.IsStagedLancelot(target)

	case *LancelotAgregation:
		ok = stage.IsStagedLancelotAgregation(target)

	case *LancelotAgregationUse:
		ok = stage.IsStagedLancelotAgregationUse(target)

	case *LancelotCategory:
		ok = stage.IsStagedLancelotCategory(target)

	case *LancelotCategoryUse:
		ok = stage.IsStagedLancelotCategoryUse(target)

	case *MapObject:
		ok = stage.IsStagedMapObject(target)

	case *MapObjectUse:
		ok = stage.IsStagedMapObjectUse(target)

	case *Repository:
		ok = stage.IsStagedRepository(target)

	case *SirRobin:
		ok = stage.IsStagedSirRobin(target)

	case *TheBridge:
		ok = stage.IsStagedTheBridge(target)

	case *TheNuteShape:
		ok = stage.IsStagedTheNuteShape(target)

	case *TheNuteTransition:
		ok = stage.IsStagedTheNuteTransition(target)

	case *User:
		ok = stage.IsStagedUser(target)

	case *UserUse:
		ok = stage.IsStagedUserUse(target)

	case *WhatIsYourPreferedColor:
		ok = stage.IsStagedWhatIsYourPreferedColor(target)

	case *Workspace:
		ok = stage.IsStagedWorkspace(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
	func (stage *StageStruct) IsStagedAWitch(awitch *AWitch) (ok bool) {

		_, ok = stage.AWitchs[awitch]
	
		return
	}

	func (stage *StageStruct) IsStagedBlackKnightShape(blackknightshape *BlackKnightShape) (ok bool) {

		_, ok = stage.BlackKnightShapes[blackknightshape]
	
		return
	}

	func (stage *StageStruct) IsStagedBringYourDead(bringyourdead *BringYourDead) (ok bool) {

		_, ok = stage.BringYourDeads[bringyourdead]
	
		return
	}

	func (stage *StageStruct) IsStagedDocument(document *Document) (ok bool) {

		_, ok = stage.Documents[document]
	
		return
	}

	func (stage *StageStruct) IsStagedDocumentUse(documentuse *DocumentUse) (ok bool) {

		_, ok = stage.DocumentUses[documentuse]
	
		return
	}

	func (stage *StageStruct) IsStagedGalahadThePure(galahadthepure *GalahadThePure) (ok bool) {

		_, ok = stage.GalahadThePures[galahadthepure]
	
		return
	}

	func (stage *StageStruct) IsStagedGeoObject(geoobject *GeoObject) (ok bool) {

		_, ok = stage.GeoObjects[geoobject]
	
		return
	}

	func (stage *StageStruct) IsStagedGeoObjectUse(geoobjectuse *GeoObjectUse) (ok bool) {

		_, ok = stage.GeoObjectUses[geoobjectuse]
	
		return
	}

	func (stage *StageStruct) IsStagedGroup(group *Group) (ok bool) {

		_, ok = stage.Groups[group]
	
		return
	}

	func (stage *StageStruct) IsStagedGroupUse(groupuse *GroupUse) (ok bool) {

		_, ok = stage.GroupUses[groupuse]
	
		return
	}

	func (stage *StageStruct) IsStagedKingArthur(kingarthur *KingArthur) (ok bool) {

		_, ok = stage.KingArthurs[kingarthur]
	
		return
	}

	func (stage *StageStruct) IsStagedKingArthurShape(kingarthurshape *KingArthurShape) (ok bool) {

		_, ok = stage.KingArthurShapes[kingarthurshape]
	
		return
	}

	func (stage *StageStruct) IsStagedKnightWhoSayNi(knightwhosayni *KnightWhoSayNi) (ok bool) {

		_, ok = stage.KnightWhoSayNis[knightwhosayni]
	
		return
	}

	func (stage *StageStruct) IsStagedLancelot(lancelot *Lancelot) (ok bool) {

		_, ok = stage.Lancelots[lancelot]
	
		return
	}

	func (stage *StageStruct) IsStagedLancelotAgregation(lancelotagregation *LancelotAgregation) (ok bool) {

		_, ok = stage.LancelotAgregations[lancelotagregation]
	
		return
	}

	func (stage *StageStruct) IsStagedLancelotAgregationUse(lancelotagregationuse *LancelotAgregationUse) (ok bool) {

		_, ok = stage.LancelotAgregationUses[lancelotagregationuse]
	
		return
	}

	func (stage *StageStruct) IsStagedLancelotCategory(lancelotcategory *LancelotCategory) (ok bool) {

		_, ok = stage.LancelotCategorys[lancelotcategory]
	
		return
	}

	func (stage *StageStruct) IsStagedLancelotCategoryUse(lancelotcategoryuse *LancelotCategoryUse) (ok bool) {

		_, ok = stage.LancelotCategoryUses[lancelotcategoryuse]
	
		return
	}

	func (stage *StageStruct) IsStagedMapObject(mapobject *MapObject) (ok bool) {

		_, ok = stage.MapObjects[mapobject]
	
		return
	}

	func (stage *StageStruct) IsStagedMapObjectUse(mapobjectuse *MapObjectUse) (ok bool) {

		_, ok = stage.MapObjectUses[mapobjectuse]
	
		return
	}

	func (stage *StageStruct) IsStagedRepository(repository *Repository) (ok bool) {

		_, ok = stage.Repositorys[repository]
	
		return
	}

	func (stage *StageStruct) IsStagedSirRobin(sirrobin *SirRobin) (ok bool) {

		_, ok = stage.SirRobins[sirrobin]
	
		return
	}

	func (stage *StageStruct) IsStagedTheBridge(thebridge *TheBridge) (ok bool) {

		_, ok = stage.TheBridges[thebridge]
	
		return
	}

	func (stage *StageStruct) IsStagedTheNuteShape(thenuteshape *TheNuteShape) (ok bool) {

		_, ok = stage.TheNuteShapes[thenuteshape]
	
		return
	}

	func (stage *StageStruct) IsStagedTheNuteTransition(thenutetransition *TheNuteTransition) (ok bool) {

		_, ok = stage.TheNuteTransitions[thenutetransition]
	
		return
	}

	func (stage *StageStruct) IsStagedUser(user *User) (ok bool) {

		_, ok = stage.Users[user]
	
		return
	}

	func (stage *StageStruct) IsStagedUserUse(useruse *UserUse) (ok bool) {

		_, ok = stage.UserUses[useruse]
	
		return
	}

	func (stage *StageStruct) IsStagedWhatIsYourPreferedColor(whatisyourpreferedcolor *WhatIsYourPreferedColor) (ok bool) {

		_, ok = stage.WhatIsYourPreferedColors[whatisyourpreferedcolor]
	
		return
	}

	func (stage *StageStruct) IsStagedWorkspace(workspace *Workspace) (ok bool) {

		_, ok = stage.Workspaces[workspace]
	
		return
	}


// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AWitch:
		stage.StageBranchAWitch(target)

	case *BlackKnightShape:
		stage.StageBranchBlackKnightShape(target)

	case *BringYourDead:
		stage.StageBranchBringYourDead(target)

	case *Document:
		stage.StageBranchDocument(target)

	case *DocumentUse:
		stage.StageBranchDocumentUse(target)

	case *GalahadThePure:
		stage.StageBranchGalahadThePure(target)

	case *GeoObject:
		stage.StageBranchGeoObject(target)

	case *GeoObjectUse:
		stage.StageBranchGeoObjectUse(target)

	case *Group:
		stage.StageBranchGroup(target)

	case *GroupUse:
		stage.StageBranchGroupUse(target)

	case *KingArthur:
		stage.StageBranchKingArthur(target)

	case *KingArthurShape:
		stage.StageBranchKingArthurShape(target)

	case *KnightWhoSayNi:
		stage.StageBranchKnightWhoSayNi(target)

	case *Lancelot:
		stage.StageBranchLancelot(target)

	case *LancelotAgregation:
		stage.StageBranchLancelotAgregation(target)

	case *LancelotAgregationUse:
		stage.StageBranchLancelotAgregationUse(target)

	case *LancelotCategory:
		stage.StageBranchLancelotCategory(target)

	case *LancelotCategoryUse:
		stage.StageBranchLancelotCategoryUse(target)

	case *MapObject:
		stage.StageBranchMapObject(target)

	case *MapObjectUse:
		stage.StageBranchMapObjectUse(target)

	case *Repository:
		stage.StageBranchRepository(target)

	case *SirRobin:
		stage.StageBranchSirRobin(target)

	case *TheBridge:
		stage.StageBranchTheBridge(target)

	case *TheNuteShape:
		stage.StageBranchTheNuteShape(target)

	case *TheNuteTransition:
		stage.StageBranchTheNuteTransition(target)

	case *User:
		stage.StageBranchUser(target)

	case *UserUse:
		stage.StageBranchUserUse(target)

	case *WhatIsYourPreferedColor:
		stage.StageBranchWhatIsYourPreferedColor(target)

	case *Workspace:
		stage.StageBranchWorkspace(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAWitch(awitch *AWitch) {

	// check if instance is already staged
	if IsStaged(stage, awitch) {
		return
	}

	awitch.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if awitch.EvolutionDirection != nil {
		StageBranch(stage, awitch.EvolutionDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBlackKnightShape(blackknightshape *BlackKnightShape) {

	// check if instance is already staged
	if IsStaged(stage, blackknightshape) {
		return
	}

	blackknightshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if blackknightshape.BringYourDead != nil {
		StageBranch(stage, blackknightshape.BringYourDead)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBringYourDead(bringyourdead *BringYourDead) {

	// check if instance is already staged
	if IsStaged(stage, bringyourdead) {
		return
	}

	bringyourdead.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lancelot := range bringyourdead.Lancelots {
		StageBranch(stage, _lancelot)
	}

}

func (stage *StageStruct) StageBranchDocument(document *Document) {

	// check if instance is already staged
	if IsStaged(stage, document) {
		return
	}

	document.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geoobjectuse := range document.GeoObjectUse {
		StageBranch(stage, _geoobjectuse)
	}

}

func (stage *StageStruct) StageBranchDocumentUse(documentuse *DocumentUse) {

	// check if instance is already staged
	if IsStaged(stage, documentuse) {
		return
	}

	documentuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if documentuse.Document != nil {
		StageBranch(stage, documentuse.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGalahadThePure(galahadthepure *GalahadThePure) {

	// check if instance is already staged
	if IsStaged(stage, galahadthepure) {
		return
	}

	galahadthepure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGeoObject(geoobject *GeoObject) {

	// check if instance is already staged
	if IsStaged(stage, geoobject) {
		return
	}

	geoobject.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGeoObjectUse(geoobjectuse *GeoObjectUse) {

	// check if instance is already staged
	if IsStaged(stage, geoobjectuse) {
		return
	}

	geoobjectuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geoobjectuse.GeoObject != nil {
		StageBranch(stage, geoobjectuse.GeoObject)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGroup(group *Group) {

	// check if instance is already staged
	if IsStaged(stage, group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _useruse := range group.UserUse {
		StageBranch(stage, _useruse)
	}

}

func (stage *StageStruct) StageBranchGroupUse(groupuse *GroupUse) {

	// check if instance is already staged
	if IsStaged(stage, groupuse) {
		return
	}

	groupuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if groupuse.Group != nil {
		StageBranch(stage, groupuse.Group)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchKingArthur(kingarthur *KingArthur) {

	// check if instance is already staged
	if IsStaged(stage, kingarthur) {
		return
	}

	kingarthur.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchKingArthurShape(kingarthurshape *KingArthurShape) {

	// check if instance is already staged
	if IsStaged(stage, kingarthurshape) {
		return
	}

	kingarthurshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if kingarthurshape.ActorState != nil {
		StageBranch(stage, kingarthurshape.ActorState)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchKnightWhoSayNi(knightwhosayni *KnightWhoSayNi) {

	// check if instance is already staged
	if IsStaged(stage, knightwhosayni) {
		return
	}

	knightwhosayni.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if knightwhosayni.Parameter != nil {
		StageBranch(stage, knightwhosayni.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLancelot(lancelot *Lancelot) {

	// check if instance is already staged
	if IsStaged(stage, lancelot) {
		return
	}

	lancelot.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _groupuse := range lancelot.GroupUse {
		StageBranch(stage, _groupuse)
	}
	for _, _documentuse := range lancelot.DocumentUse {
		StageBranch(stage, _documentuse)
	}
	for _, _geoobjectuse := range lancelot.GeoObjectUse {
		StageBranch(stage, _geoobjectuse)
	}

}

func (stage *StageStruct) StageBranchLancelotAgregation(lancelotagregation *LancelotAgregation) {

	// check if instance is already staged
	if IsStaged(stage, lancelotagregation) {
		return
	}

	lancelotagregation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _knightwhosayni := range lancelotagregation.ParameterUse {
		StageBranch(stage, _knightwhosayni)
	}

}

func (stage *StageStruct) StageBranchLancelotAgregationUse(lancelotagregationuse *LancelotAgregationUse) {

	// check if instance is already staged
	if IsStaged(stage, lancelotagregationuse) {
		return
	}

	lancelotagregationuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lancelotagregationuse.ParameterAgregation != nil {
		StageBranch(stage, lancelotagregationuse.ParameterAgregation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLancelotCategory(lancelotcategory *LancelotCategory) {

	// check if instance is already staged
	if IsStaged(stage, lancelotcategory) {
		return
	}

	lancelotcategory.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _knightwhosayni := range lancelotcategory.ParameterUse {
		StageBranch(stage, _knightwhosayni)
	}

}

func (stage *StageStruct) StageBranchLancelotCategoryUse(lancelotcategoryuse *LancelotCategoryUse) {

	// check if instance is already staged
	if IsStaged(stage, lancelotcategoryuse) {
		return
	}

	lancelotcategoryuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lancelotcategoryuse.ParameterCategory != nil {
		StageBranch(stage, lancelotcategoryuse.ParameterCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMapObject(mapobject *MapObject) {

	// check if instance is already staged
	if IsStaged(stage, mapobject) {
		return
	}

	mapobject.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMapObjectUse(mapobjectuse *MapObjectUse) {

	// check if instance is already staged
	if IsStaged(stage, mapobjectuse) {
		return
	}

	mapobjectuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mapobjectuse.Map != nil {
		StageBranch(stage, mapobjectuse.Map)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRepository(repository *Repository) {

	// check if instance is already staged
	if IsStaged(stage, repository) {
		return
	}

	repository.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _knightwhosayni := range repository.ParameterUse {
		StageBranch(stage, _knightwhosayni)
	}
	for _, _groupuse := range repository.GroupUse {
		StageBranch(stage, _groupuse)
	}

}

func (stage *StageStruct) StageBranchSirRobin(sirrobin *SirRobin) {

	// check if instance is already staged
	if IsStaged(stage, sirrobin) {
		return
	}

	sirrobin.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _awitch := range sirrobin.Witches {
		StageBranch(stage, _awitch)
	}
	for _, _kingarthurshape := range sirrobin.Arthurs {
		StageBranch(stage, _kingarthurshape)
	}
	for _, _knightwhosayni := range sirrobin.KnightWhoSayNis {
		StageBranch(stage, _knightwhosayni)
	}
	for _, _blackknightshape := range sirrobin.BlackKnightShapes {
		StageBranch(stage, _blackknightshape)
	}
	for _, _thenuteshape := range sirrobin.TheNuteShapes {
		StageBranch(stage, _thenuteshape)
	}

}

func (stage *StageStruct) StageBranchTheBridge(thebridge *TheBridge) {

	// check if instance is already staged
	if IsStaged(stage, thebridge) {
		return
	}

	thebridge.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _whatisyourpreferedcolor := range thebridge.WhatIsYourPreferedColor {
		StageBranch(stage, _whatisyourpreferedcolor)
	}
	for _, _groupuse := range thebridge.GroupUse {
		StageBranch(stage, _groupuse)
	}
	for _, _geoobjectuse := range thebridge.GeoObjectUse {
		StageBranch(stage, _geoobjectuse)
	}
	for _, _mapobjectuse := range thebridge.MapUse {
		StageBranch(stage, _mapobjectuse)
	}

}

func (stage *StageStruct) StageBranchTheNuteShape(thenuteshape *TheNuteShape) {

	// check if instance is already staged
	if IsStaged(stage, thenuteshape) {
		return
	}

	thenuteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if thenuteshape.ActorStateTransition != nil {
		StageBranch(stage, thenuteshape.ActorStateTransition)
	}
	if thenuteshape.Start != nil {
		StageBranch(stage, thenuteshape.Start)
	}
	if thenuteshape.End != nil {
		StageBranch(stage, thenuteshape.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTheNuteTransition(thenutetransition *TheNuteTransition) {

	// check if instance is already staged
	if IsStaged(stage, thenutetransition) {
		return
	}

	thenutetransition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if thenutetransition.StartState != nil {
		StageBranch(stage, thenutetransition.StartState)
	}
	if thenutetransition.EndState != nil {
		StageBranch(stage, thenutetransition.EndState)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchUser(user *User) {

	// check if instance is already staged
	if IsStaged(stage, user) {
		return
	}

	user.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchUserUse(useruse *UserUse) {

	// check if instance is already staged
	if IsStaged(stage, useruse) {
		return
	}

	useruse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if useruse.User != nil {
		StageBranch(stage, useruse.User)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchWhatIsYourPreferedColor(whatisyourpreferedcolor *WhatIsYourPreferedColor) {

	// check if instance is already staged
	if IsStaged(stage, whatisyourpreferedcolor) {
		return
	}

	whatisyourpreferedcolor.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sirrobin := range whatisyourpreferedcolor.Diagrams {
		StageBranch(stage, _sirrobin)
	}
	for _, _kingarthur := range whatisyourpreferedcolor.KingArthurs {
		StageBranch(stage, _kingarthur)
	}
	for _, _thenutetransition := range whatisyourpreferedcolor.Nutes {
		StageBranch(stage, _thenutetransition)
	}
	for _, _galahadthepure := range whatisyourpreferedcolor.Galahard {
		StageBranch(stage, _galahadthepure)
	}
	for _, _lancelot := range whatisyourpreferedcolor.Lancelots {
		StageBranch(stage, _lancelot)
	}
	for _, _bringyourdead := range whatisyourpreferedcolor.BringYourDeadarameters {
		StageBranch(stage, _bringyourdead)
	}

}

func (stage *StageStruct) StageBranchWorkspace(workspace *Workspace) {

	// check if instance is already staged
	if IsStaged(stage, workspace) {
		return
	}

	workspace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if workspace.SelectedDiagram != nil {
		StageBranch(stage, workspace.SelectedDiagram)
	}
	if workspace.A != nil {
		StageBranch(stage, workspace.A)
	}
	if workspace.B != nil {
		StageBranch(stage, workspace.B)
	}
	if workspace.C != nil {
		StageBranch(stage, workspace.C)
	}
	if workspace.D != nil {
		StageBranch(stage, workspace.D)
	}
	if workspace.E != nil {
		StageBranch(stage, workspace.E)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}


// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *AWitch:
		stage.UnstageBranchAWitch(target)

	case *BlackKnightShape:
		stage.UnstageBranchBlackKnightShape(target)

	case *BringYourDead:
		stage.UnstageBranchBringYourDead(target)

	case *Document:
		stage.UnstageBranchDocument(target)

	case *DocumentUse:
		stage.UnstageBranchDocumentUse(target)

	case *GalahadThePure:
		stage.UnstageBranchGalahadThePure(target)

	case *GeoObject:
		stage.UnstageBranchGeoObject(target)

	case *GeoObjectUse:
		stage.UnstageBranchGeoObjectUse(target)

	case *Group:
		stage.UnstageBranchGroup(target)

	case *GroupUse:
		stage.UnstageBranchGroupUse(target)

	case *KingArthur:
		stage.UnstageBranchKingArthur(target)

	case *KingArthurShape:
		stage.UnstageBranchKingArthurShape(target)

	case *KnightWhoSayNi:
		stage.UnstageBranchKnightWhoSayNi(target)

	case *Lancelot:
		stage.UnstageBranchLancelot(target)

	case *LancelotAgregation:
		stage.UnstageBranchLancelotAgregation(target)

	case *LancelotAgregationUse:
		stage.UnstageBranchLancelotAgregationUse(target)

	case *LancelotCategory:
		stage.UnstageBranchLancelotCategory(target)

	case *LancelotCategoryUse:
		stage.UnstageBranchLancelotCategoryUse(target)

	case *MapObject:
		stage.UnstageBranchMapObject(target)

	case *MapObjectUse:
		stage.UnstageBranchMapObjectUse(target)

	case *Repository:
		stage.UnstageBranchRepository(target)

	case *SirRobin:
		stage.UnstageBranchSirRobin(target)

	case *TheBridge:
		stage.UnstageBranchTheBridge(target)

	case *TheNuteShape:
		stage.UnstageBranchTheNuteShape(target)

	case *TheNuteTransition:
		stage.UnstageBranchTheNuteTransition(target)

	case *User:
		stage.UnstageBranchUser(target)

	case *UserUse:
		stage.UnstageBranchUserUse(target)

	case *WhatIsYourPreferedColor:
		stage.UnstageBranchWhatIsYourPreferedColor(target)

	case *Workspace:
		stage.UnstageBranchWorkspace(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAWitch(awitch *AWitch) {

	// check if instance is already staged
	if ! IsStaged(stage, awitch) {
		return
	}

	awitch.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if awitch.EvolutionDirection != nil {
		UnstageBranch(stage, awitch.EvolutionDirection)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBlackKnightShape(blackknightshape *BlackKnightShape) {

	// check if instance is already staged
	if ! IsStaged(stage, blackknightshape) {
		return
	}

	blackknightshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if blackknightshape.BringYourDead != nil {
		UnstageBranch(stage, blackknightshape.BringYourDead)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBringYourDead(bringyourdead *BringYourDead) {

	// check if instance is already staged
	if ! IsStaged(stage, bringyourdead) {
		return
	}

	bringyourdead.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lancelot := range bringyourdead.Lancelots {
		UnstageBranch(stage, _lancelot)
	}

}

func (stage *StageStruct) UnstageBranchDocument(document *Document) {

	// check if instance is already staged
	if ! IsStaged(stage, document) {
		return
	}

	document.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geoobjectuse := range document.GeoObjectUse {
		UnstageBranch(stage, _geoobjectuse)
	}

}

func (stage *StageStruct) UnstageBranchDocumentUse(documentuse *DocumentUse) {

	// check if instance is already staged
	if ! IsStaged(stage, documentuse) {
		return
	}

	documentuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if documentuse.Document != nil {
		UnstageBranch(stage, documentuse.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGalahadThePure(galahadthepure *GalahadThePure) {

	// check if instance is already staged
	if ! IsStaged(stage, galahadthepure) {
		return
	}

	galahadthepure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGeoObject(geoobject *GeoObject) {

	// check if instance is already staged
	if ! IsStaged(stage, geoobject) {
		return
	}

	geoobject.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGeoObjectUse(geoobjectuse *GeoObjectUse) {

	// check if instance is already staged
	if ! IsStaged(stage, geoobjectuse) {
		return
	}

	geoobjectuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geoobjectuse.GeoObject != nil {
		UnstageBranch(stage, geoobjectuse.GeoObject)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGroup(group *Group) {

	// check if instance is already staged
	if ! IsStaged(stage, group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _useruse := range group.UserUse {
		UnstageBranch(stage, _useruse)
	}

}

func (stage *StageStruct) UnstageBranchGroupUse(groupuse *GroupUse) {

	// check if instance is already staged
	if ! IsStaged(stage, groupuse) {
		return
	}

	groupuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if groupuse.Group != nil {
		UnstageBranch(stage, groupuse.Group)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchKingArthur(kingarthur *KingArthur) {

	// check if instance is already staged
	if ! IsStaged(stage, kingarthur) {
		return
	}

	kingarthur.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchKingArthurShape(kingarthurshape *KingArthurShape) {

	// check if instance is already staged
	if ! IsStaged(stage, kingarthurshape) {
		return
	}

	kingarthurshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if kingarthurshape.ActorState != nil {
		UnstageBranch(stage, kingarthurshape.ActorState)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchKnightWhoSayNi(knightwhosayni *KnightWhoSayNi) {

	// check if instance is already staged
	if ! IsStaged(stage, knightwhosayni) {
		return
	}

	knightwhosayni.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if knightwhosayni.Parameter != nil {
		UnstageBranch(stage, knightwhosayni.Parameter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLancelot(lancelot *Lancelot) {

	// check if instance is already staged
	if ! IsStaged(stage, lancelot) {
		return
	}

	lancelot.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _groupuse := range lancelot.GroupUse {
		UnstageBranch(stage, _groupuse)
	}
	for _, _documentuse := range lancelot.DocumentUse {
		UnstageBranch(stage, _documentuse)
	}
	for _, _geoobjectuse := range lancelot.GeoObjectUse {
		UnstageBranch(stage, _geoobjectuse)
	}

}

func (stage *StageStruct) UnstageBranchLancelotAgregation(lancelotagregation *LancelotAgregation) {

	// check if instance is already staged
	if ! IsStaged(stage, lancelotagregation) {
		return
	}

	lancelotagregation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _knightwhosayni := range lancelotagregation.ParameterUse {
		UnstageBranch(stage, _knightwhosayni)
	}

}

func (stage *StageStruct) UnstageBranchLancelotAgregationUse(lancelotagregationuse *LancelotAgregationUse) {

	// check if instance is already staged
	if ! IsStaged(stage, lancelotagregationuse) {
		return
	}

	lancelotagregationuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lancelotagregationuse.ParameterAgregation != nil {
		UnstageBranch(stage, lancelotagregationuse.ParameterAgregation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLancelotCategory(lancelotcategory *LancelotCategory) {

	// check if instance is already staged
	if ! IsStaged(stage, lancelotcategory) {
		return
	}

	lancelotcategory.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _knightwhosayni := range lancelotcategory.ParameterUse {
		UnstageBranch(stage, _knightwhosayni)
	}

}

func (stage *StageStruct) UnstageBranchLancelotCategoryUse(lancelotcategoryuse *LancelotCategoryUse) {

	// check if instance is already staged
	if ! IsStaged(stage, lancelotcategoryuse) {
		return
	}

	lancelotcategoryuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lancelotcategoryuse.ParameterCategory != nil {
		UnstageBranch(stage, lancelotcategoryuse.ParameterCategory)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMapObject(mapobject *MapObject) {

	// check if instance is already staged
	if ! IsStaged(stage, mapobject) {
		return
	}

	mapobject.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMapObjectUse(mapobjectuse *MapObjectUse) {

	// check if instance is already staged
	if ! IsStaged(stage, mapobjectuse) {
		return
	}

	mapobjectuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mapobjectuse.Map != nil {
		UnstageBranch(stage, mapobjectuse.Map)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRepository(repository *Repository) {

	// check if instance is already staged
	if ! IsStaged(stage, repository) {
		return
	}

	repository.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _knightwhosayni := range repository.ParameterUse {
		UnstageBranch(stage, _knightwhosayni)
	}
	for _, _groupuse := range repository.GroupUse {
		UnstageBranch(stage, _groupuse)
	}

}

func (stage *StageStruct) UnstageBranchSirRobin(sirrobin *SirRobin) {

	// check if instance is already staged
	if ! IsStaged(stage, sirrobin) {
		return
	}

	sirrobin.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _awitch := range sirrobin.Witches {
		UnstageBranch(stage, _awitch)
	}
	for _, _kingarthurshape := range sirrobin.Arthurs {
		UnstageBranch(stage, _kingarthurshape)
	}
	for _, _knightwhosayni := range sirrobin.KnightWhoSayNis {
		UnstageBranch(stage, _knightwhosayni)
	}
	for _, _blackknightshape := range sirrobin.BlackKnightShapes {
		UnstageBranch(stage, _blackknightshape)
	}
	for _, _thenuteshape := range sirrobin.TheNuteShapes {
		UnstageBranch(stage, _thenuteshape)
	}

}

func (stage *StageStruct) UnstageBranchTheBridge(thebridge *TheBridge) {

	// check if instance is already staged
	if ! IsStaged(stage, thebridge) {
		return
	}

	thebridge.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _whatisyourpreferedcolor := range thebridge.WhatIsYourPreferedColor {
		UnstageBranch(stage, _whatisyourpreferedcolor)
	}
	for _, _groupuse := range thebridge.GroupUse {
		UnstageBranch(stage, _groupuse)
	}
	for _, _geoobjectuse := range thebridge.GeoObjectUse {
		UnstageBranch(stage, _geoobjectuse)
	}
	for _, _mapobjectuse := range thebridge.MapUse {
		UnstageBranch(stage, _mapobjectuse)
	}

}

func (stage *StageStruct) UnstageBranchTheNuteShape(thenuteshape *TheNuteShape) {

	// check if instance is already staged
	if ! IsStaged(stage, thenuteshape) {
		return
	}

	thenuteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if thenuteshape.ActorStateTransition != nil {
		UnstageBranch(stage, thenuteshape.ActorStateTransition)
	}
	if thenuteshape.Start != nil {
		UnstageBranch(stage, thenuteshape.Start)
	}
	if thenuteshape.End != nil {
		UnstageBranch(stage, thenuteshape.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTheNuteTransition(thenutetransition *TheNuteTransition) {

	// check if instance is already staged
	if ! IsStaged(stage, thenutetransition) {
		return
	}

	thenutetransition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if thenutetransition.StartState != nil {
		UnstageBranch(stage, thenutetransition.StartState)
	}
	if thenutetransition.EndState != nil {
		UnstageBranch(stage, thenutetransition.EndState)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchUser(user *User) {

	// check if instance is already staged
	if ! IsStaged(stage, user) {
		return
	}

	user.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchUserUse(useruse *UserUse) {

	// check if instance is already staged
	if ! IsStaged(stage, useruse) {
		return
	}

	useruse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if useruse.User != nil {
		UnstageBranch(stage, useruse.User)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchWhatIsYourPreferedColor(whatisyourpreferedcolor *WhatIsYourPreferedColor) {

	// check if instance is already staged
	if ! IsStaged(stage, whatisyourpreferedcolor) {
		return
	}

	whatisyourpreferedcolor.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sirrobin := range whatisyourpreferedcolor.Diagrams {
		UnstageBranch(stage, _sirrobin)
	}
	for _, _kingarthur := range whatisyourpreferedcolor.KingArthurs {
		UnstageBranch(stage, _kingarthur)
	}
	for _, _thenutetransition := range whatisyourpreferedcolor.Nutes {
		UnstageBranch(stage, _thenutetransition)
	}
	for _, _galahadthepure := range whatisyourpreferedcolor.Galahard {
		UnstageBranch(stage, _galahadthepure)
	}
	for _, _lancelot := range whatisyourpreferedcolor.Lancelots {
		UnstageBranch(stage, _lancelot)
	}
	for _, _bringyourdead := range whatisyourpreferedcolor.BringYourDeadarameters {
		UnstageBranch(stage, _bringyourdead)
	}

}

func (stage *StageStruct) UnstageBranchWorkspace(workspace *Workspace) {

	// check if instance is already staged
	if ! IsStaged(stage, workspace) {
		return
	}

	workspace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if workspace.SelectedDiagram != nil {
		UnstageBranch(stage, workspace.SelectedDiagram)
	}
	if workspace.A != nil {
		UnstageBranch(stage, workspace.A)
	}
	if workspace.B != nil {
		UnstageBranch(stage, workspace.B)
	}
	if workspace.C != nil {
		UnstageBranch(stage, workspace.C)
	}
	if workspace.D != nil {
		UnstageBranch(stage, workspace.D)
	}
	if workspace.E != nil {
		UnstageBranch(stage, workspace.E)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

