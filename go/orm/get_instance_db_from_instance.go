// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/thelongbuild/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | AWitchDB | BlackKnightShapeDB | BringYourDeadDB | DocumentDB | DocumentUseDB | GalahadThePureDB | GeoObjectDB | GeoObjectUseDB | GroupDB | GroupUseDB | KingArthurDB | KingArthurShapeDB | KnightWhoSayNiDB | LancelotDB | LancelotAgregationDB | LancelotAgregationUseDB | LancelotCategoryDB | LancelotCategoryUseDB | MapObjectDB | MapObjectUseDB | RepositoryDB | SirRobinDB | TheBridgeDB | TheNuteShapeDB | TheNuteTransitionDB | UserDB | UserUseDB | WhatIsYourPreferedColorDB | WorkspaceDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AWitch:
		awitchInstance := any(concreteInstance).(*models.AWitch)
		ret2 := backRepo.BackRepoAWitch.GetAWitchDBFromAWitchPtr(awitchInstance)
		ret = any(ret2).(*T2)
	case *models.BlackKnightShape:
		blackknightshapeInstance := any(concreteInstance).(*models.BlackKnightShape)
		ret2 := backRepo.BackRepoBlackKnightShape.GetBlackKnightShapeDBFromBlackKnightShapePtr(blackknightshapeInstance)
		ret = any(ret2).(*T2)
	case *models.BringYourDead:
		bringyourdeadInstance := any(concreteInstance).(*models.BringYourDead)
		ret2 := backRepo.BackRepoBringYourDead.GetBringYourDeadDBFromBringYourDeadPtr(bringyourdeadInstance)
		ret = any(ret2).(*T2)
	case *models.Document:
		documentInstance := any(concreteInstance).(*models.Document)
		ret2 := backRepo.BackRepoDocument.GetDocumentDBFromDocumentPtr(documentInstance)
		ret = any(ret2).(*T2)
	case *models.DocumentUse:
		documentuseInstance := any(concreteInstance).(*models.DocumentUse)
		ret2 := backRepo.BackRepoDocumentUse.GetDocumentUseDBFromDocumentUsePtr(documentuseInstance)
		ret = any(ret2).(*T2)
	case *models.GalahadThePure:
		galahadthepureInstance := any(concreteInstance).(*models.GalahadThePure)
		ret2 := backRepo.BackRepoGalahadThePure.GetGalahadThePureDBFromGalahadThePurePtr(galahadthepureInstance)
		ret = any(ret2).(*T2)
	case *models.GeoObject:
		geoobjectInstance := any(concreteInstance).(*models.GeoObject)
		ret2 := backRepo.BackRepoGeoObject.GetGeoObjectDBFromGeoObjectPtr(geoobjectInstance)
		ret = any(ret2).(*T2)
	case *models.GeoObjectUse:
		geoobjectuseInstance := any(concreteInstance).(*models.GeoObjectUse)
		ret2 := backRepo.BackRepoGeoObjectUse.GetGeoObjectUseDBFromGeoObjectUsePtr(geoobjectuseInstance)
		ret = any(ret2).(*T2)
	case *models.Group:
		groupInstance := any(concreteInstance).(*models.Group)
		ret2 := backRepo.BackRepoGroup.GetGroupDBFromGroupPtr(groupInstance)
		ret = any(ret2).(*T2)
	case *models.GroupUse:
		groupuseInstance := any(concreteInstance).(*models.GroupUse)
		ret2 := backRepo.BackRepoGroupUse.GetGroupUseDBFromGroupUsePtr(groupuseInstance)
		ret = any(ret2).(*T2)
	case *models.KingArthur:
		kingarthurInstance := any(concreteInstance).(*models.KingArthur)
		ret2 := backRepo.BackRepoKingArthur.GetKingArthurDBFromKingArthurPtr(kingarthurInstance)
		ret = any(ret2).(*T2)
	case *models.KingArthurShape:
		kingarthurshapeInstance := any(concreteInstance).(*models.KingArthurShape)
		ret2 := backRepo.BackRepoKingArthurShape.GetKingArthurShapeDBFromKingArthurShapePtr(kingarthurshapeInstance)
		ret = any(ret2).(*T2)
	case *models.KnightWhoSayNi:
		knightwhosayniInstance := any(concreteInstance).(*models.KnightWhoSayNi)
		ret2 := backRepo.BackRepoKnightWhoSayNi.GetKnightWhoSayNiDBFromKnightWhoSayNiPtr(knightwhosayniInstance)
		ret = any(ret2).(*T2)
	case *models.Lancelot:
		lancelotInstance := any(concreteInstance).(*models.Lancelot)
		ret2 := backRepo.BackRepoLancelot.GetLancelotDBFromLancelotPtr(lancelotInstance)
		ret = any(ret2).(*T2)
	case *models.LancelotAgregation:
		lancelotagregationInstance := any(concreteInstance).(*models.LancelotAgregation)
		ret2 := backRepo.BackRepoLancelotAgregation.GetLancelotAgregationDBFromLancelotAgregationPtr(lancelotagregationInstance)
		ret = any(ret2).(*T2)
	case *models.LancelotAgregationUse:
		lancelotagregationuseInstance := any(concreteInstance).(*models.LancelotAgregationUse)
		ret2 := backRepo.BackRepoLancelotAgregationUse.GetLancelotAgregationUseDBFromLancelotAgregationUsePtr(lancelotagregationuseInstance)
		ret = any(ret2).(*T2)
	case *models.LancelotCategory:
		lancelotcategoryInstance := any(concreteInstance).(*models.LancelotCategory)
		ret2 := backRepo.BackRepoLancelotCategory.GetLancelotCategoryDBFromLancelotCategoryPtr(lancelotcategoryInstance)
		ret = any(ret2).(*T2)
	case *models.LancelotCategoryUse:
		lancelotcategoryuseInstance := any(concreteInstance).(*models.LancelotCategoryUse)
		ret2 := backRepo.BackRepoLancelotCategoryUse.GetLancelotCategoryUseDBFromLancelotCategoryUsePtr(lancelotcategoryuseInstance)
		ret = any(ret2).(*T2)
	case *models.MapObject:
		mapobjectInstance := any(concreteInstance).(*models.MapObject)
		ret2 := backRepo.BackRepoMapObject.GetMapObjectDBFromMapObjectPtr(mapobjectInstance)
		ret = any(ret2).(*T2)
	case *models.MapObjectUse:
		mapobjectuseInstance := any(concreteInstance).(*models.MapObjectUse)
		ret2 := backRepo.BackRepoMapObjectUse.GetMapObjectUseDBFromMapObjectUsePtr(mapobjectuseInstance)
		ret = any(ret2).(*T2)
	case *models.Repository:
		repositoryInstance := any(concreteInstance).(*models.Repository)
		ret2 := backRepo.BackRepoRepository.GetRepositoryDBFromRepositoryPtr(repositoryInstance)
		ret = any(ret2).(*T2)
	case *models.SirRobin:
		sirrobinInstance := any(concreteInstance).(*models.SirRobin)
		ret2 := backRepo.BackRepoSirRobin.GetSirRobinDBFromSirRobinPtr(sirrobinInstance)
		ret = any(ret2).(*T2)
	case *models.TheBridge:
		thebridgeInstance := any(concreteInstance).(*models.TheBridge)
		ret2 := backRepo.BackRepoTheBridge.GetTheBridgeDBFromTheBridgePtr(thebridgeInstance)
		ret = any(ret2).(*T2)
	case *models.TheNuteShape:
		thenuteshapeInstance := any(concreteInstance).(*models.TheNuteShape)
		ret2 := backRepo.BackRepoTheNuteShape.GetTheNuteShapeDBFromTheNuteShapePtr(thenuteshapeInstance)
		ret = any(ret2).(*T2)
	case *models.TheNuteTransition:
		thenutetransitionInstance := any(concreteInstance).(*models.TheNuteTransition)
		ret2 := backRepo.BackRepoTheNuteTransition.GetTheNuteTransitionDBFromTheNuteTransitionPtr(thenutetransitionInstance)
		ret = any(ret2).(*T2)
	case *models.User:
		userInstance := any(concreteInstance).(*models.User)
		ret2 := backRepo.BackRepoUser.GetUserDBFromUserPtr(userInstance)
		ret = any(ret2).(*T2)
	case *models.UserUse:
		useruseInstance := any(concreteInstance).(*models.UserUse)
		ret2 := backRepo.BackRepoUserUse.GetUserUseDBFromUserUsePtr(useruseInstance)
		ret = any(ret2).(*T2)
	case *models.WhatIsYourPreferedColor:
		whatisyourpreferedcolorInstance := any(concreteInstance).(*models.WhatIsYourPreferedColor)
		ret2 := backRepo.BackRepoWhatIsYourPreferedColor.GetWhatIsYourPreferedColorDBFromWhatIsYourPreferedColorPtr(whatisyourpreferedcolorInstance)
		ret = any(ret2).(*T2)
	case *models.Workspace:
		workspaceInstance := any(concreteInstance).(*models.Workspace)
		ret2 := backRepo.BackRepoWorkspace.GetWorkspaceDBFromWorkspacePtr(workspaceInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AWitch:
		tmp := GetInstanceDBFromInstance[models.AWitch, AWitchDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BlackKnightShape:
		tmp := GetInstanceDBFromInstance[models.BlackKnightShape, BlackKnightShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BringYourDead:
		tmp := GetInstanceDBFromInstance[models.BringYourDead, BringYourDeadDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Document:
		tmp := GetInstanceDBFromInstance[models.Document, DocumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DocumentUse:
		tmp := GetInstanceDBFromInstance[models.DocumentUse, DocumentUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GalahadThePure:
		tmp := GetInstanceDBFromInstance[models.GalahadThePure, GalahadThePureDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GeoObject:
		tmp := GetInstanceDBFromInstance[models.GeoObject, GeoObjectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GeoObjectUse:
		tmp := GetInstanceDBFromInstance[models.GeoObjectUse, GeoObjectUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GroupUse:
		tmp := GetInstanceDBFromInstance[models.GroupUse, GroupUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.KingArthur:
		tmp := GetInstanceDBFromInstance[models.KingArthur, KingArthurDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.KingArthurShape:
		tmp := GetInstanceDBFromInstance[models.KingArthurShape, KingArthurShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.KnightWhoSayNi:
		tmp := GetInstanceDBFromInstance[models.KnightWhoSayNi, KnightWhoSayNiDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lancelot:
		tmp := GetInstanceDBFromInstance[models.Lancelot, LancelotDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LancelotAgregation:
		tmp := GetInstanceDBFromInstance[models.LancelotAgregation, LancelotAgregationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LancelotAgregationUse:
		tmp := GetInstanceDBFromInstance[models.LancelotAgregationUse, LancelotAgregationUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LancelotCategory:
		tmp := GetInstanceDBFromInstance[models.LancelotCategory, LancelotCategoryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LancelotCategoryUse:
		tmp := GetInstanceDBFromInstance[models.LancelotCategoryUse, LancelotCategoryUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MapObject:
		tmp := GetInstanceDBFromInstance[models.MapObject, MapObjectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MapObjectUse:
		tmp := GetInstanceDBFromInstance[models.MapObjectUse, MapObjectUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Repository:
		tmp := GetInstanceDBFromInstance[models.Repository, RepositoryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SirRobin:
		tmp := GetInstanceDBFromInstance[models.SirRobin, SirRobinDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TheBridge:
		tmp := GetInstanceDBFromInstance[models.TheBridge, TheBridgeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TheNuteShape:
		tmp := GetInstanceDBFromInstance[models.TheNuteShape, TheNuteShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TheNuteTransition:
		tmp := GetInstanceDBFromInstance[models.TheNuteTransition, TheNuteTransitionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.User:
		tmp := GetInstanceDBFromInstance[models.User, UserDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.UserUse:
		tmp := GetInstanceDBFromInstance[models.UserUse, UserUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.WhatIsYourPreferedColor:
		tmp := GetInstanceDBFromInstance[models.WhatIsYourPreferedColor, WhatIsYourPreferedColorDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Workspace:
		tmp := GetInstanceDBFromInstance[models.Workspace, WorkspaceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AWitch:
		tmp := GetInstanceDBFromInstance[models.AWitch, AWitchDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BlackKnightShape:
		tmp := GetInstanceDBFromInstance[models.BlackKnightShape, BlackKnightShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BringYourDead:
		tmp := GetInstanceDBFromInstance[models.BringYourDead, BringYourDeadDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Document:
		tmp := GetInstanceDBFromInstance[models.Document, DocumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DocumentUse:
		tmp := GetInstanceDBFromInstance[models.DocumentUse, DocumentUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GalahadThePure:
		tmp := GetInstanceDBFromInstance[models.GalahadThePure, GalahadThePureDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GeoObject:
		tmp := GetInstanceDBFromInstance[models.GeoObject, GeoObjectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GeoObjectUse:
		tmp := GetInstanceDBFromInstance[models.GeoObjectUse, GeoObjectUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.GroupUse:
		tmp := GetInstanceDBFromInstance[models.GroupUse, GroupUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.KingArthur:
		tmp := GetInstanceDBFromInstance[models.KingArthur, KingArthurDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.KingArthurShape:
		tmp := GetInstanceDBFromInstance[models.KingArthurShape, KingArthurShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.KnightWhoSayNi:
		tmp := GetInstanceDBFromInstance[models.KnightWhoSayNi, KnightWhoSayNiDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lancelot:
		tmp := GetInstanceDBFromInstance[models.Lancelot, LancelotDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LancelotAgregation:
		tmp := GetInstanceDBFromInstance[models.LancelotAgregation, LancelotAgregationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LancelotAgregationUse:
		tmp := GetInstanceDBFromInstance[models.LancelotAgregationUse, LancelotAgregationUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LancelotCategory:
		tmp := GetInstanceDBFromInstance[models.LancelotCategory, LancelotCategoryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LancelotCategoryUse:
		tmp := GetInstanceDBFromInstance[models.LancelotCategoryUse, LancelotCategoryUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MapObject:
		tmp := GetInstanceDBFromInstance[models.MapObject, MapObjectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MapObjectUse:
		tmp := GetInstanceDBFromInstance[models.MapObjectUse, MapObjectUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Repository:
		tmp := GetInstanceDBFromInstance[models.Repository, RepositoryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SirRobin:
		tmp := GetInstanceDBFromInstance[models.SirRobin, SirRobinDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TheBridge:
		tmp := GetInstanceDBFromInstance[models.TheBridge, TheBridgeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TheNuteShape:
		tmp := GetInstanceDBFromInstance[models.TheNuteShape, TheNuteShapeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TheNuteTransition:
		tmp := GetInstanceDBFromInstance[models.TheNuteTransition, TheNuteTransitionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.User:
		tmp := GetInstanceDBFromInstance[models.User, UserDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.UserUse:
		tmp := GetInstanceDBFromInstance[models.UserUse, UserUseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.WhatIsYourPreferedColor:
		tmp := GetInstanceDBFromInstance[models.WhatIsYourPreferedColor, WhatIsYourPreferedColorDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Workspace:
		tmp := GetInstanceDBFromInstance[models.Workspace, WorkspaceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
