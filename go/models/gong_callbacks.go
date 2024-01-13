// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AWitch:
		if stage.OnAfterAWitchCreateCallback != nil {
			stage.OnAfterAWitchCreateCallback.OnAfterCreate(stage, target)
		}
	case *BlackKnightShape:
		if stage.OnAfterBlackKnightShapeCreateCallback != nil {
			stage.OnAfterBlackKnightShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *BringYourDead:
		if stage.OnAfterBringYourDeadCreateCallback != nil {
			stage.OnAfterBringYourDeadCreateCallback.OnAfterCreate(stage, target)
		}
	case *Document:
		if stage.OnAfterDocumentCreateCallback != nil {
			stage.OnAfterDocumentCreateCallback.OnAfterCreate(stage, target)
		}
	case *DocumentUse:
		if stage.OnAfterDocumentUseCreateCallback != nil {
			stage.OnAfterDocumentUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *GalahadThePure:
		if stage.OnAfterGalahadThePureCreateCallback != nil {
			stage.OnAfterGalahadThePureCreateCallback.OnAfterCreate(stage, target)
		}
	case *GeoObject:
		if stage.OnAfterGeoObjectCreateCallback != nil {
			stage.OnAfterGeoObjectCreateCallback.OnAfterCreate(stage, target)
		}
	case *GeoObjectUse:
		if stage.OnAfterGeoObjectUseCreateCallback != nil {
			stage.OnAfterGeoObjectUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupCreateCallback != nil {
			stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *GroupUse:
		if stage.OnAfterGroupUseCreateCallback != nil {
			stage.OnAfterGroupUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *KingArthur:
		if stage.OnAfterKingArthurCreateCallback != nil {
			stage.OnAfterKingArthurCreateCallback.OnAfterCreate(stage, target)
		}
	case *KingArthurShape:
		if stage.OnAfterKingArthurShapeCreateCallback != nil {
			stage.OnAfterKingArthurShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *KnightWhoSayNi:
		if stage.OnAfterKnightWhoSayNiCreateCallback != nil {
			stage.OnAfterKnightWhoSayNiCreateCallback.OnAfterCreate(stage, target)
		}
	case *Lancelot:
		if stage.OnAfterLancelotCreateCallback != nil {
			stage.OnAfterLancelotCreateCallback.OnAfterCreate(stage, target)
		}
	case *LancelotAgregation:
		if stage.OnAfterLancelotAgregationCreateCallback != nil {
			stage.OnAfterLancelotAgregationCreateCallback.OnAfterCreate(stage, target)
		}
	case *LancelotAgregationUse:
		if stage.OnAfterLancelotAgregationUseCreateCallback != nil {
			stage.OnAfterLancelotAgregationUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *LancelotCategory:
		if stage.OnAfterLancelotCategoryCreateCallback != nil {
			stage.OnAfterLancelotCategoryCreateCallback.OnAfterCreate(stage, target)
		}
	case *LancelotCategoryUse:
		if stage.OnAfterLancelotCategoryUseCreateCallback != nil {
			stage.OnAfterLancelotCategoryUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *MapObject:
		if stage.OnAfterMapObjectCreateCallback != nil {
			stage.OnAfterMapObjectCreateCallback.OnAfterCreate(stage, target)
		}
	case *MapObjectUse:
		if stage.OnAfterMapObjectUseCreateCallback != nil {
			stage.OnAfterMapObjectUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Repository:
		if stage.OnAfterRepositoryCreateCallback != nil {
			stage.OnAfterRepositoryCreateCallback.OnAfterCreate(stage, target)
		}
	case *SirRobin:
		if stage.OnAfterSirRobinCreateCallback != nil {
			stage.OnAfterSirRobinCreateCallback.OnAfterCreate(stage, target)
		}
	case *TheBridge:
		if stage.OnAfterTheBridgeCreateCallback != nil {
			stage.OnAfterTheBridgeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TheNuteShape:
		if stage.OnAfterTheNuteShapeCreateCallback != nil {
			stage.OnAfterTheNuteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TheNuteTransition:
		if stage.OnAfterTheNuteTransitionCreateCallback != nil {
			stage.OnAfterTheNuteTransitionCreateCallback.OnAfterCreate(stage, target)
		}
	case *User:
		if stage.OnAfterUserCreateCallback != nil {
			stage.OnAfterUserCreateCallback.OnAfterCreate(stage, target)
		}
	case *UserUse:
		if stage.OnAfterUserUseCreateCallback != nil {
			stage.OnAfterUserUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *WhatIsYourPreferedColor:
		if stage.OnAfterWhatIsYourPreferedColorCreateCallback != nil {
			stage.OnAfterWhatIsYourPreferedColorCreateCallback.OnAfterCreate(stage, target)
		}
	case *Workspace:
		if stage.OnAfterWorkspaceCreateCallback != nil {
			stage.OnAfterWorkspaceCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AWitch:
		newTarget := any(new).(*AWitch)
		if stage.OnAfterAWitchUpdateCallback != nil {
			stage.OnAfterAWitchUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BlackKnightShape:
		newTarget := any(new).(*BlackKnightShape)
		if stage.OnAfterBlackKnightShapeUpdateCallback != nil {
			stage.OnAfterBlackKnightShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BringYourDead:
		newTarget := any(new).(*BringYourDead)
		if stage.OnAfterBringYourDeadUpdateCallback != nil {
			stage.OnAfterBringYourDeadUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Document:
		newTarget := any(new).(*Document)
		if stage.OnAfterDocumentUpdateCallback != nil {
			stage.OnAfterDocumentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DocumentUse:
		newTarget := any(new).(*DocumentUse)
		if stage.OnAfterDocumentUseUpdateCallback != nil {
			stage.OnAfterDocumentUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GalahadThePure:
		newTarget := any(new).(*GalahadThePure)
		if stage.OnAfterGalahadThePureUpdateCallback != nil {
			stage.OnAfterGalahadThePureUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GeoObject:
		newTarget := any(new).(*GeoObject)
		if stage.OnAfterGeoObjectUpdateCallback != nil {
			stage.OnAfterGeoObjectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GeoObjectUse:
		newTarget := any(new).(*GeoObjectUse)
		if stage.OnAfterGeoObjectUseUpdateCallback != nil {
			stage.OnAfterGeoObjectUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group:
		newTarget := any(new).(*Group)
		if stage.OnAfterGroupUpdateCallback != nil {
			stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GroupUse:
		newTarget := any(new).(*GroupUse)
		if stage.OnAfterGroupUseUpdateCallback != nil {
			stage.OnAfterGroupUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *KingArthur:
		newTarget := any(new).(*KingArthur)
		if stage.OnAfterKingArthurUpdateCallback != nil {
			stage.OnAfterKingArthurUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *KingArthurShape:
		newTarget := any(new).(*KingArthurShape)
		if stage.OnAfterKingArthurShapeUpdateCallback != nil {
			stage.OnAfterKingArthurShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *KnightWhoSayNi:
		newTarget := any(new).(*KnightWhoSayNi)
		if stage.OnAfterKnightWhoSayNiUpdateCallback != nil {
			stage.OnAfterKnightWhoSayNiUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Lancelot:
		newTarget := any(new).(*Lancelot)
		if stage.OnAfterLancelotUpdateCallback != nil {
			stage.OnAfterLancelotUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LancelotAgregation:
		newTarget := any(new).(*LancelotAgregation)
		if stage.OnAfterLancelotAgregationUpdateCallback != nil {
			stage.OnAfterLancelotAgregationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LancelotAgregationUse:
		newTarget := any(new).(*LancelotAgregationUse)
		if stage.OnAfterLancelotAgregationUseUpdateCallback != nil {
			stage.OnAfterLancelotAgregationUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LancelotCategory:
		newTarget := any(new).(*LancelotCategory)
		if stage.OnAfterLancelotCategoryUpdateCallback != nil {
			stage.OnAfterLancelotCategoryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LancelotCategoryUse:
		newTarget := any(new).(*LancelotCategoryUse)
		if stage.OnAfterLancelotCategoryUseUpdateCallback != nil {
			stage.OnAfterLancelotCategoryUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MapObject:
		newTarget := any(new).(*MapObject)
		if stage.OnAfterMapObjectUpdateCallback != nil {
			stage.OnAfterMapObjectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MapObjectUse:
		newTarget := any(new).(*MapObjectUse)
		if stage.OnAfterMapObjectUseUpdateCallback != nil {
			stage.OnAfterMapObjectUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Repository:
		newTarget := any(new).(*Repository)
		if stage.OnAfterRepositoryUpdateCallback != nil {
			stage.OnAfterRepositoryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SirRobin:
		newTarget := any(new).(*SirRobin)
		if stage.OnAfterSirRobinUpdateCallback != nil {
			stage.OnAfterSirRobinUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TheBridge:
		newTarget := any(new).(*TheBridge)
		if stage.OnAfterTheBridgeUpdateCallback != nil {
			stage.OnAfterTheBridgeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TheNuteShape:
		newTarget := any(new).(*TheNuteShape)
		if stage.OnAfterTheNuteShapeUpdateCallback != nil {
			stage.OnAfterTheNuteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TheNuteTransition:
		newTarget := any(new).(*TheNuteTransition)
		if stage.OnAfterTheNuteTransitionUpdateCallback != nil {
			stage.OnAfterTheNuteTransitionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *User:
		newTarget := any(new).(*User)
		if stage.OnAfterUserUpdateCallback != nil {
			stage.OnAfterUserUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *UserUse:
		newTarget := any(new).(*UserUse)
		if stage.OnAfterUserUseUpdateCallback != nil {
			stage.OnAfterUserUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *WhatIsYourPreferedColor:
		newTarget := any(new).(*WhatIsYourPreferedColor)
		if stage.OnAfterWhatIsYourPreferedColorUpdateCallback != nil {
			stage.OnAfterWhatIsYourPreferedColorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Workspace:
		newTarget := any(new).(*Workspace)
		if stage.OnAfterWorkspaceUpdateCallback != nil {
			stage.OnAfterWorkspaceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AWitch:
		if stage.OnAfterAWitchDeleteCallback != nil {
			staged := any(staged).(*AWitch)
			stage.OnAfterAWitchDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BlackKnightShape:
		if stage.OnAfterBlackKnightShapeDeleteCallback != nil {
			staged := any(staged).(*BlackKnightShape)
			stage.OnAfterBlackKnightShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BringYourDead:
		if stage.OnAfterBringYourDeadDeleteCallback != nil {
			staged := any(staged).(*BringYourDead)
			stage.OnAfterBringYourDeadDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Document:
		if stage.OnAfterDocumentDeleteCallback != nil {
			staged := any(staged).(*Document)
			stage.OnAfterDocumentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DocumentUse:
		if stage.OnAfterDocumentUseDeleteCallback != nil {
			staged := any(staged).(*DocumentUse)
			stage.OnAfterDocumentUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GalahadThePure:
		if stage.OnAfterGalahadThePureDeleteCallback != nil {
			staged := any(staged).(*GalahadThePure)
			stage.OnAfterGalahadThePureDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GeoObject:
		if stage.OnAfterGeoObjectDeleteCallback != nil {
			staged := any(staged).(*GeoObject)
			stage.OnAfterGeoObjectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GeoObjectUse:
		if stage.OnAfterGeoObjectUseDeleteCallback != nil {
			staged := any(staged).(*GeoObjectUse)
			stage.OnAfterGeoObjectUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group:
		if stage.OnAfterGroupDeleteCallback != nil {
			staged := any(staged).(*Group)
			stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GroupUse:
		if stage.OnAfterGroupUseDeleteCallback != nil {
			staged := any(staged).(*GroupUse)
			stage.OnAfterGroupUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *KingArthur:
		if stage.OnAfterKingArthurDeleteCallback != nil {
			staged := any(staged).(*KingArthur)
			stage.OnAfterKingArthurDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *KingArthurShape:
		if stage.OnAfterKingArthurShapeDeleteCallback != nil {
			staged := any(staged).(*KingArthurShape)
			stage.OnAfterKingArthurShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *KnightWhoSayNi:
		if stage.OnAfterKnightWhoSayNiDeleteCallback != nil {
			staged := any(staged).(*KnightWhoSayNi)
			stage.OnAfterKnightWhoSayNiDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Lancelot:
		if stage.OnAfterLancelotDeleteCallback != nil {
			staged := any(staged).(*Lancelot)
			stage.OnAfterLancelotDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LancelotAgregation:
		if stage.OnAfterLancelotAgregationDeleteCallback != nil {
			staged := any(staged).(*LancelotAgregation)
			stage.OnAfterLancelotAgregationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LancelotAgregationUse:
		if stage.OnAfterLancelotAgregationUseDeleteCallback != nil {
			staged := any(staged).(*LancelotAgregationUse)
			stage.OnAfterLancelotAgregationUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LancelotCategory:
		if stage.OnAfterLancelotCategoryDeleteCallback != nil {
			staged := any(staged).(*LancelotCategory)
			stage.OnAfterLancelotCategoryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LancelotCategoryUse:
		if stage.OnAfterLancelotCategoryUseDeleteCallback != nil {
			staged := any(staged).(*LancelotCategoryUse)
			stage.OnAfterLancelotCategoryUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MapObject:
		if stage.OnAfterMapObjectDeleteCallback != nil {
			staged := any(staged).(*MapObject)
			stage.OnAfterMapObjectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MapObjectUse:
		if stage.OnAfterMapObjectUseDeleteCallback != nil {
			staged := any(staged).(*MapObjectUse)
			stage.OnAfterMapObjectUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Repository:
		if stage.OnAfterRepositoryDeleteCallback != nil {
			staged := any(staged).(*Repository)
			stage.OnAfterRepositoryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SirRobin:
		if stage.OnAfterSirRobinDeleteCallback != nil {
			staged := any(staged).(*SirRobin)
			stage.OnAfterSirRobinDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TheBridge:
		if stage.OnAfterTheBridgeDeleteCallback != nil {
			staged := any(staged).(*TheBridge)
			stage.OnAfterTheBridgeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TheNuteShape:
		if stage.OnAfterTheNuteShapeDeleteCallback != nil {
			staged := any(staged).(*TheNuteShape)
			stage.OnAfterTheNuteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TheNuteTransition:
		if stage.OnAfterTheNuteTransitionDeleteCallback != nil {
			staged := any(staged).(*TheNuteTransition)
			stage.OnAfterTheNuteTransitionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *User:
		if stage.OnAfterUserDeleteCallback != nil {
			staged := any(staged).(*User)
			stage.OnAfterUserDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *UserUse:
		if stage.OnAfterUserUseDeleteCallback != nil {
			staged := any(staged).(*UserUse)
			stage.OnAfterUserUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *WhatIsYourPreferedColor:
		if stage.OnAfterWhatIsYourPreferedColorDeleteCallback != nil {
			staged := any(staged).(*WhatIsYourPreferedColor)
			stage.OnAfterWhatIsYourPreferedColorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Workspace:
		if stage.OnAfterWorkspaceDeleteCallback != nil {
			staged := any(staged).(*Workspace)
			stage.OnAfterWorkspaceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AWitch:
		if stage.OnAfterAWitchReadCallback != nil {
			stage.OnAfterAWitchReadCallback.OnAfterRead(stage, target)
		}
	case *BlackKnightShape:
		if stage.OnAfterBlackKnightShapeReadCallback != nil {
			stage.OnAfterBlackKnightShapeReadCallback.OnAfterRead(stage, target)
		}
	case *BringYourDead:
		if stage.OnAfterBringYourDeadReadCallback != nil {
			stage.OnAfterBringYourDeadReadCallback.OnAfterRead(stage, target)
		}
	case *Document:
		if stage.OnAfterDocumentReadCallback != nil {
			stage.OnAfterDocumentReadCallback.OnAfterRead(stage, target)
		}
	case *DocumentUse:
		if stage.OnAfterDocumentUseReadCallback != nil {
			stage.OnAfterDocumentUseReadCallback.OnAfterRead(stage, target)
		}
	case *GalahadThePure:
		if stage.OnAfterGalahadThePureReadCallback != nil {
			stage.OnAfterGalahadThePureReadCallback.OnAfterRead(stage, target)
		}
	case *GeoObject:
		if stage.OnAfterGeoObjectReadCallback != nil {
			stage.OnAfterGeoObjectReadCallback.OnAfterRead(stage, target)
		}
	case *GeoObjectUse:
		if stage.OnAfterGeoObjectUseReadCallback != nil {
			stage.OnAfterGeoObjectUseReadCallback.OnAfterRead(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupReadCallback != nil {
			stage.OnAfterGroupReadCallback.OnAfterRead(stage, target)
		}
	case *GroupUse:
		if stage.OnAfterGroupUseReadCallback != nil {
			stage.OnAfterGroupUseReadCallback.OnAfterRead(stage, target)
		}
	case *KingArthur:
		if stage.OnAfterKingArthurReadCallback != nil {
			stage.OnAfterKingArthurReadCallback.OnAfterRead(stage, target)
		}
	case *KingArthurShape:
		if stage.OnAfterKingArthurShapeReadCallback != nil {
			stage.OnAfterKingArthurShapeReadCallback.OnAfterRead(stage, target)
		}
	case *KnightWhoSayNi:
		if stage.OnAfterKnightWhoSayNiReadCallback != nil {
			stage.OnAfterKnightWhoSayNiReadCallback.OnAfterRead(stage, target)
		}
	case *Lancelot:
		if stage.OnAfterLancelotReadCallback != nil {
			stage.OnAfterLancelotReadCallback.OnAfterRead(stage, target)
		}
	case *LancelotAgregation:
		if stage.OnAfterLancelotAgregationReadCallback != nil {
			stage.OnAfterLancelotAgregationReadCallback.OnAfterRead(stage, target)
		}
	case *LancelotAgregationUse:
		if stage.OnAfterLancelotAgregationUseReadCallback != nil {
			stage.OnAfterLancelotAgregationUseReadCallback.OnAfterRead(stage, target)
		}
	case *LancelotCategory:
		if stage.OnAfterLancelotCategoryReadCallback != nil {
			stage.OnAfterLancelotCategoryReadCallback.OnAfterRead(stage, target)
		}
	case *LancelotCategoryUse:
		if stage.OnAfterLancelotCategoryUseReadCallback != nil {
			stage.OnAfterLancelotCategoryUseReadCallback.OnAfterRead(stage, target)
		}
	case *MapObject:
		if stage.OnAfterMapObjectReadCallback != nil {
			stage.OnAfterMapObjectReadCallback.OnAfterRead(stage, target)
		}
	case *MapObjectUse:
		if stage.OnAfterMapObjectUseReadCallback != nil {
			stage.OnAfterMapObjectUseReadCallback.OnAfterRead(stage, target)
		}
	case *Repository:
		if stage.OnAfterRepositoryReadCallback != nil {
			stage.OnAfterRepositoryReadCallback.OnAfterRead(stage, target)
		}
	case *SirRobin:
		if stage.OnAfterSirRobinReadCallback != nil {
			stage.OnAfterSirRobinReadCallback.OnAfterRead(stage, target)
		}
	case *TheBridge:
		if stage.OnAfterTheBridgeReadCallback != nil {
			stage.OnAfterTheBridgeReadCallback.OnAfterRead(stage, target)
		}
	case *TheNuteShape:
		if stage.OnAfterTheNuteShapeReadCallback != nil {
			stage.OnAfterTheNuteShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TheNuteTransition:
		if stage.OnAfterTheNuteTransitionReadCallback != nil {
			stage.OnAfterTheNuteTransitionReadCallback.OnAfterRead(stage, target)
		}
	case *User:
		if stage.OnAfterUserReadCallback != nil {
			stage.OnAfterUserReadCallback.OnAfterRead(stage, target)
		}
	case *UserUse:
		if stage.OnAfterUserUseReadCallback != nil {
			stage.OnAfterUserUseReadCallback.OnAfterRead(stage, target)
		}
	case *WhatIsYourPreferedColor:
		if stage.OnAfterWhatIsYourPreferedColorReadCallback != nil {
			stage.OnAfterWhatIsYourPreferedColorReadCallback.OnAfterRead(stage, target)
		}
	case *Workspace:
		if stage.OnAfterWorkspaceReadCallback != nil {
			stage.OnAfterWorkspaceReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AWitch:
		stage.OnAfterAWitchUpdateCallback = any(callback).(OnAfterUpdateInterface[AWitch])
	
	case *BlackKnightShape:
		stage.OnAfterBlackKnightShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[BlackKnightShape])
	
	case *BringYourDead:
		stage.OnAfterBringYourDeadUpdateCallback = any(callback).(OnAfterUpdateInterface[BringYourDead])
	
	case *Document:
		stage.OnAfterDocumentUpdateCallback = any(callback).(OnAfterUpdateInterface[Document])
	
	case *DocumentUse:
		stage.OnAfterDocumentUseUpdateCallback = any(callback).(OnAfterUpdateInterface[DocumentUse])
	
	case *GalahadThePure:
		stage.OnAfterGalahadThePureUpdateCallback = any(callback).(OnAfterUpdateInterface[GalahadThePure])
	
	case *GeoObject:
		stage.OnAfterGeoObjectUpdateCallback = any(callback).(OnAfterUpdateInterface[GeoObject])
	
	case *GeoObjectUse:
		stage.OnAfterGeoObjectUseUpdateCallback = any(callback).(OnAfterUpdateInterface[GeoObjectUse])
	
	case *Group:
		stage.OnAfterGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[Group])
	
	case *GroupUse:
		stage.OnAfterGroupUseUpdateCallback = any(callback).(OnAfterUpdateInterface[GroupUse])
	
	case *KingArthur:
		stage.OnAfterKingArthurUpdateCallback = any(callback).(OnAfterUpdateInterface[KingArthur])
	
	case *KingArthurShape:
		stage.OnAfterKingArthurShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[KingArthurShape])
	
	case *KnightWhoSayNi:
		stage.OnAfterKnightWhoSayNiUpdateCallback = any(callback).(OnAfterUpdateInterface[KnightWhoSayNi])
	
	case *Lancelot:
		stage.OnAfterLancelotUpdateCallback = any(callback).(OnAfterUpdateInterface[Lancelot])
	
	case *LancelotAgregation:
		stage.OnAfterLancelotAgregationUpdateCallback = any(callback).(OnAfterUpdateInterface[LancelotAgregation])
	
	case *LancelotAgregationUse:
		stage.OnAfterLancelotAgregationUseUpdateCallback = any(callback).(OnAfterUpdateInterface[LancelotAgregationUse])
	
	case *LancelotCategory:
		stage.OnAfterLancelotCategoryUpdateCallback = any(callback).(OnAfterUpdateInterface[LancelotCategory])
	
	case *LancelotCategoryUse:
		stage.OnAfterLancelotCategoryUseUpdateCallback = any(callback).(OnAfterUpdateInterface[LancelotCategoryUse])
	
	case *MapObject:
		stage.OnAfterMapObjectUpdateCallback = any(callback).(OnAfterUpdateInterface[MapObject])
	
	case *MapObjectUse:
		stage.OnAfterMapObjectUseUpdateCallback = any(callback).(OnAfterUpdateInterface[MapObjectUse])
	
	case *Repository:
		stage.OnAfterRepositoryUpdateCallback = any(callback).(OnAfterUpdateInterface[Repository])
	
	case *SirRobin:
		stage.OnAfterSirRobinUpdateCallback = any(callback).(OnAfterUpdateInterface[SirRobin])
	
	case *TheBridge:
		stage.OnAfterTheBridgeUpdateCallback = any(callback).(OnAfterUpdateInterface[TheBridge])
	
	case *TheNuteShape:
		stage.OnAfterTheNuteShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TheNuteShape])
	
	case *TheNuteTransition:
		stage.OnAfterTheNuteTransitionUpdateCallback = any(callback).(OnAfterUpdateInterface[TheNuteTransition])
	
	case *User:
		stage.OnAfterUserUpdateCallback = any(callback).(OnAfterUpdateInterface[User])
	
	case *UserUse:
		stage.OnAfterUserUseUpdateCallback = any(callback).(OnAfterUpdateInterface[UserUse])
	
	case *WhatIsYourPreferedColor:
		stage.OnAfterWhatIsYourPreferedColorUpdateCallback = any(callback).(OnAfterUpdateInterface[WhatIsYourPreferedColor])
	
	case *Workspace:
		stage.OnAfterWorkspaceUpdateCallback = any(callback).(OnAfterUpdateInterface[Workspace])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AWitch:
		stage.OnAfterAWitchCreateCallback = any(callback).(OnAfterCreateInterface[AWitch])
	
	case *BlackKnightShape:
		stage.OnAfterBlackKnightShapeCreateCallback = any(callback).(OnAfterCreateInterface[BlackKnightShape])
	
	case *BringYourDead:
		stage.OnAfterBringYourDeadCreateCallback = any(callback).(OnAfterCreateInterface[BringYourDead])
	
	case *Document:
		stage.OnAfterDocumentCreateCallback = any(callback).(OnAfterCreateInterface[Document])
	
	case *DocumentUse:
		stage.OnAfterDocumentUseCreateCallback = any(callback).(OnAfterCreateInterface[DocumentUse])
	
	case *GalahadThePure:
		stage.OnAfterGalahadThePureCreateCallback = any(callback).(OnAfterCreateInterface[GalahadThePure])
	
	case *GeoObject:
		stage.OnAfterGeoObjectCreateCallback = any(callback).(OnAfterCreateInterface[GeoObject])
	
	case *GeoObjectUse:
		stage.OnAfterGeoObjectUseCreateCallback = any(callback).(OnAfterCreateInterface[GeoObjectUse])
	
	case *Group:
		stage.OnAfterGroupCreateCallback = any(callback).(OnAfterCreateInterface[Group])
	
	case *GroupUse:
		stage.OnAfterGroupUseCreateCallback = any(callback).(OnAfterCreateInterface[GroupUse])
	
	case *KingArthur:
		stage.OnAfterKingArthurCreateCallback = any(callback).(OnAfterCreateInterface[KingArthur])
	
	case *KingArthurShape:
		stage.OnAfterKingArthurShapeCreateCallback = any(callback).(OnAfterCreateInterface[KingArthurShape])
	
	case *KnightWhoSayNi:
		stage.OnAfterKnightWhoSayNiCreateCallback = any(callback).(OnAfterCreateInterface[KnightWhoSayNi])
	
	case *Lancelot:
		stage.OnAfterLancelotCreateCallback = any(callback).(OnAfterCreateInterface[Lancelot])
	
	case *LancelotAgregation:
		stage.OnAfterLancelotAgregationCreateCallback = any(callback).(OnAfterCreateInterface[LancelotAgregation])
	
	case *LancelotAgregationUse:
		stage.OnAfterLancelotAgregationUseCreateCallback = any(callback).(OnAfterCreateInterface[LancelotAgregationUse])
	
	case *LancelotCategory:
		stage.OnAfterLancelotCategoryCreateCallback = any(callback).(OnAfterCreateInterface[LancelotCategory])
	
	case *LancelotCategoryUse:
		stage.OnAfterLancelotCategoryUseCreateCallback = any(callback).(OnAfterCreateInterface[LancelotCategoryUse])
	
	case *MapObject:
		stage.OnAfterMapObjectCreateCallback = any(callback).(OnAfterCreateInterface[MapObject])
	
	case *MapObjectUse:
		stage.OnAfterMapObjectUseCreateCallback = any(callback).(OnAfterCreateInterface[MapObjectUse])
	
	case *Repository:
		stage.OnAfterRepositoryCreateCallback = any(callback).(OnAfterCreateInterface[Repository])
	
	case *SirRobin:
		stage.OnAfterSirRobinCreateCallback = any(callback).(OnAfterCreateInterface[SirRobin])
	
	case *TheBridge:
		stage.OnAfterTheBridgeCreateCallback = any(callback).(OnAfterCreateInterface[TheBridge])
	
	case *TheNuteShape:
		stage.OnAfterTheNuteShapeCreateCallback = any(callback).(OnAfterCreateInterface[TheNuteShape])
	
	case *TheNuteTransition:
		stage.OnAfterTheNuteTransitionCreateCallback = any(callback).(OnAfterCreateInterface[TheNuteTransition])
	
	case *User:
		stage.OnAfterUserCreateCallback = any(callback).(OnAfterCreateInterface[User])
	
	case *UserUse:
		stage.OnAfterUserUseCreateCallback = any(callback).(OnAfterCreateInterface[UserUse])
	
	case *WhatIsYourPreferedColor:
		stage.OnAfterWhatIsYourPreferedColorCreateCallback = any(callback).(OnAfterCreateInterface[WhatIsYourPreferedColor])
	
	case *Workspace:
		stage.OnAfterWorkspaceCreateCallback = any(callback).(OnAfterCreateInterface[Workspace])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AWitch:
		stage.OnAfterAWitchDeleteCallback = any(callback).(OnAfterDeleteInterface[AWitch])
	
	case *BlackKnightShape:
		stage.OnAfterBlackKnightShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[BlackKnightShape])
	
	case *BringYourDead:
		stage.OnAfterBringYourDeadDeleteCallback = any(callback).(OnAfterDeleteInterface[BringYourDead])
	
	case *Document:
		stage.OnAfterDocumentDeleteCallback = any(callback).(OnAfterDeleteInterface[Document])
	
	case *DocumentUse:
		stage.OnAfterDocumentUseDeleteCallback = any(callback).(OnAfterDeleteInterface[DocumentUse])
	
	case *GalahadThePure:
		stage.OnAfterGalahadThePureDeleteCallback = any(callback).(OnAfterDeleteInterface[GalahadThePure])
	
	case *GeoObject:
		stage.OnAfterGeoObjectDeleteCallback = any(callback).(OnAfterDeleteInterface[GeoObject])
	
	case *GeoObjectUse:
		stage.OnAfterGeoObjectUseDeleteCallback = any(callback).(OnAfterDeleteInterface[GeoObjectUse])
	
	case *Group:
		stage.OnAfterGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[Group])
	
	case *GroupUse:
		stage.OnAfterGroupUseDeleteCallback = any(callback).(OnAfterDeleteInterface[GroupUse])
	
	case *KingArthur:
		stage.OnAfterKingArthurDeleteCallback = any(callback).(OnAfterDeleteInterface[KingArthur])
	
	case *KingArthurShape:
		stage.OnAfterKingArthurShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[KingArthurShape])
	
	case *KnightWhoSayNi:
		stage.OnAfterKnightWhoSayNiDeleteCallback = any(callback).(OnAfterDeleteInterface[KnightWhoSayNi])
	
	case *Lancelot:
		stage.OnAfterLancelotDeleteCallback = any(callback).(OnAfterDeleteInterface[Lancelot])
	
	case *LancelotAgregation:
		stage.OnAfterLancelotAgregationDeleteCallback = any(callback).(OnAfterDeleteInterface[LancelotAgregation])
	
	case *LancelotAgregationUse:
		stage.OnAfterLancelotAgregationUseDeleteCallback = any(callback).(OnAfterDeleteInterface[LancelotAgregationUse])
	
	case *LancelotCategory:
		stage.OnAfterLancelotCategoryDeleteCallback = any(callback).(OnAfterDeleteInterface[LancelotCategory])
	
	case *LancelotCategoryUse:
		stage.OnAfterLancelotCategoryUseDeleteCallback = any(callback).(OnAfterDeleteInterface[LancelotCategoryUse])
	
	case *MapObject:
		stage.OnAfterMapObjectDeleteCallback = any(callback).(OnAfterDeleteInterface[MapObject])
	
	case *MapObjectUse:
		stage.OnAfterMapObjectUseDeleteCallback = any(callback).(OnAfterDeleteInterface[MapObjectUse])
	
	case *Repository:
		stage.OnAfterRepositoryDeleteCallback = any(callback).(OnAfterDeleteInterface[Repository])
	
	case *SirRobin:
		stage.OnAfterSirRobinDeleteCallback = any(callback).(OnAfterDeleteInterface[SirRobin])
	
	case *TheBridge:
		stage.OnAfterTheBridgeDeleteCallback = any(callback).(OnAfterDeleteInterface[TheBridge])
	
	case *TheNuteShape:
		stage.OnAfterTheNuteShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TheNuteShape])
	
	case *TheNuteTransition:
		stage.OnAfterTheNuteTransitionDeleteCallback = any(callback).(OnAfterDeleteInterface[TheNuteTransition])
	
	case *User:
		stage.OnAfterUserDeleteCallback = any(callback).(OnAfterDeleteInterface[User])
	
	case *UserUse:
		stage.OnAfterUserUseDeleteCallback = any(callback).(OnAfterDeleteInterface[UserUse])
	
	case *WhatIsYourPreferedColor:
		stage.OnAfterWhatIsYourPreferedColorDeleteCallback = any(callback).(OnAfterDeleteInterface[WhatIsYourPreferedColor])
	
	case *Workspace:
		stage.OnAfterWorkspaceDeleteCallback = any(callback).(OnAfterDeleteInterface[Workspace])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AWitch:
		stage.OnAfterAWitchReadCallback = any(callback).(OnAfterReadInterface[AWitch])
	
	case *BlackKnightShape:
		stage.OnAfterBlackKnightShapeReadCallback = any(callback).(OnAfterReadInterface[BlackKnightShape])
	
	case *BringYourDead:
		stage.OnAfterBringYourDeadReadCallback = any(callback).(OnAfterReadInterface[BringYourDead])
	
	case *Document:
		stage.OnAfterDocumentReadCallback = any(callback).(OnAfterReadInterface[Document])
	
	case *DocumentUse:
		stage.OnAfterDocumentUseReadCallback = any(callback).(OnAfterReadInterface[DocumentUse])
	
	case *GalahadThePure:
		stage.OnAfterGalahadThePureReadCallback = any(callback).(OnAfterReadInterface[GalahadThePure])
	
	case *GeoObject:
		stage.OnAfterGeoObjectReadCallback = any(callback).(OnAfterReadInterface[GeoObject])
	
	case *GeoObjectUse:
		stage.OnAfterGeoObjectUseReadCallback = any(callback).(OnAfterReadInterface[GeoObjectUse])
	
	case *Group:
		stage.OnAfterGroupReadCallback = any(callback).(OnAfterReadInterface[Group])
	
	case *GroupUse:
		stage.OnAfterGroupUseReadCallback = any(callback).(OnAfterReadInterface[GroupUse])
	
	case *KingArthur:
		stage.OnAfterKingArthurReadCallback = any(callback).(OnAfterReadInterface[KingArthur])
	
	case *KingArthurShape:
		stage.OnAfterKingArthurShapeReadCallback = any(callback).(OnAfterReadInterface[KingArthurShape])
	
	case *KnightWhoSayNi:
		stage.OnAfterKnightWhoSayNiReadCallback = any(callback).(OnAfterReadInterface[KnightWhoSayNi])
	
	case *Lancelot:
		stage.OnAfterLancelotReadCallback = any(callback).(OnAfterReadInterface[Lancelot])
	
	case *LancelotAgregation:
		stage.OnAfterLancelotAgregationReadCallback = any(callback).(OnAfterReadInterface[LancelotAgregation])
	
	case *LancelotAgregationUse:
		stage.OnAfterLancelotAgregationUseReadCallback = any(callback).(OnAfterReadInterface[LancelotAgregationUse])
	
	case *LancelotCategory:
		stage.OnAfterLancelotCategoryReadCallback = any(callback).(OnAfterReadInterface[LancelotCategory])
	
	case *LancelotCategoryUse:
		stage.OnAfterLancelotCategoryUseReadCallback = any(callback).(OnAfterReadInterface[LancelotCategoryUse])
	
	case *MapObject:
		stage.OnAfterMapObjectReadCallback = any(callback).(OnAfterReadInterface[MapObject])
	
	case *MapObjectUse:
		stage.OnAfterMapObjectUseReadCallback = any(callback).(OnAfterReadInterface[MapObjectUse])
	
	case *Repository:
		stage.OnAfterRepositoryReadCallback = any(callback).(OnAfterReadInterface[Repository])
	
	case *SirRobin:
		stage.OnAfterSirRobinReadCallback = any(callback).(OnAfterReadInterface[SirRobin])
	
	case *TheBridge:
		stage.OnAfterTheBridgeReadCallback = any(callback).(OnAfterReadInterface[TheBridge])
	
	case *TheNuteShape:
		stage.OnAfterTheNuteShapeReadCallback = any(callback).(OnAfterReadInterface[TheNuteShape])
	
	case *TheNuteTransition:
		stage.OnAfterTheNuteTransitionReadCallback = any(callback).(OnAfterReadInterface[TheNuteTransition])
	
	case *User:
		stage.OnAfterUserReadCallback = any(callback).(OnAfterReadInterface[User])
	
	case *UserUse:
		stage.OnAfterUserUseReadCallback = any(callback).(OnAfterReadInterface[UserUse])
	
	case *WhatIsYourPreferedColor:
		stage.OnAfterWhatIsYourPreferedColorReadCallback = any(callback).(OnAfterReadInterface[WhatIsYourPreferedColor])
	
	case *Workspace:
		stage.OnAfterWorkspaceReadCallback = any(callback).(OnAfterReadInterface[Workspace])
	
	}
}
