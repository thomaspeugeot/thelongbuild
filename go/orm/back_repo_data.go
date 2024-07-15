// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AWitchAPIs []*AWitchAPI

	BlackKnightShapeAPIs []*BlackKnightShapeAPI

	BringYourDeadAPIs []*BringYourDeadAPI

	DocumentAPIs []*DocumentAPI

	DocumentUseAPIs []*DocumentUseAPI

	GalahadThePureAPIs []*GalahadThePureAPI

	GeoObjectAPIs []*GeoObjectAPI

	GeoObjectUseAPIs []*GeoObjectUseAPI

	GroupAPIs []*GroupAPI

	GroupUseAPIs []*GroupUseAPI

	KingArthurAPIs []*KingArthurAPI

	KingArthurShapeAPIs []*KingArthurShapeAPI

	KnightWhoSayNiAPIs []*KnightWhoSayNiAPI

	LancelotAPIs []*LancelotAPI

	LancelotAgregationAPIs []*LancelotAgregationAPI

	LancelotAgregationUseAPIs []*LancelotAgregationUseAPI

	LancelotCategoryAPIs []*LancelotCategoryAPI

	LancelotCategoryUseAPIs []*LancelotCategoryUseAPI

	MapObjectAPIs []*MapObjectAPI

	MapObjectUseAPIs []*MapObjectUseAPI

	RepositoryAPIs []*RepositoryAPI

	SirRobinAPIs []*SirRobinAPI

	TheBridgeAPIs []*TheBridgeAPI

	TheNuteShapeAPIs []*TheNuteShapeAPI

	TheNuteTransitionAPIs []*TheNuteTransitionAPI

	UserAPIs []*UserAPI

	UserUseAPIs []*UserUseAPI

	WhatIsYourPreferedColorAPIs []*WhatIsYourPreferedColorAPI

	WorkspaceAPIs []*WorkspaceAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, awitchDB := range backRepo.BackRepoAWitch.Map_AWitchDBID_AWitchDB {

		var awitchAPI AWitchAPI
		awitchAPI.ID = awitchDB.ID
		awitchAPI.AWitchPointersEncoding = awitchDB.AWitchPointersEncoding
		awitchDB.CopyBasicFieldsToAWitch_WOP(&awitchAPI.AWitch_WOP)

		backRepoData.AWitchAPIs = append(backRepoData.AWitchAPIs, &awitchAPI)
	}

	for _, blackknightshapeDB := range backRepo.BackRepoBlackKnightShape.Map_BlackKnightShapeDBID_BlackKnightShapeDB {

		var blackknightshapeAPI BlackKnightShapeAPI
		blackknightshapeAPI.ID = blackknightshapeDB.ID
		blackknightshapeAPI.BlackKnightShapePointersEncoding = blackknightshapeDB.BlackKnightShapePointersEncoding
		blackknightshapeDB.CopyBasicFieldsToBlackKnightShape_WOP(&blackknightshapeAPI.BlackKnightShape_WOP)

		backRepoData.BlackKnightShapeAPIs = append(backRepoData.BlackKnightShapeAPIs, &blackknightshapeAPI)
	}

	for _, bringyourdeadDB := range backRepo.BackRepoBringYourDead.Map_BringYourDeadDBID_BringYourDeadDB {

		var bringyourdeadAPI BringYourDeadAPI
		bringyourdeadAPI.ID = bringyourdeadDB.ID
		bringyourdeadAPI.BringYourDeadPointersEncoding = bringyourdeadDB.BringYourDeadPointersEncoding
		bringyourdeadDB.CopyBasicFieldsToBringYourDead_WOP(&bringyourdeadAPI.BringYourDead_WOP)

		backRepoData.BringYourDeadAPIs = append(backRepoData.BringYourDeadAPIs, &bringyourdeadAPI)
	}

	for _, documentDB := range backRepo.BackRepoDocument.Map_DocumentDBID_DocumentDB {

		var documentAPI DocumentAPI
		documentAPI.ID = documentDB.ID
		documentAPI.DocumentPointersEncoding = documentDB.DocumentPointersEncoding
		documentDB.CopyBasicFieldsToDocument_WOP(&documentAPI.Document_WOP)

		backRepoData.DocumentAPIs = append(backRepoData.DocumentAPIs, &documentAPI)
	}

	for _, documentuseDB := range backRepo.BackRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB {

		var documentuseAPI DocumentUseAPI
		documentuseAPI.ID = documentuseDB.ID
		documentuseAPI.DocumentUsePointersEncoding = documentuseDB.DocumentUsePointersEncoding
		documentuseDB.CopyBasicFieldsToDocumentUse_WOP(&documentuseAPI.DocumentUse_WOP)

		backRepoData.DocumentUseAPIs = append(backRepoData.DocumentUseAPIs, &documentuseAPI)
	}

	for _, galahadthepureDB := range backRepo.BackRepoGalahadThePure.Map_GalahadThePureDBID_GalahadThePureDB {

		var galahadthepureAPI GalahadThePureAPI
		galahadthepureAPI.ID = galahadthepureDB.ID
		galahadthepureAPI.GalahadThePurePointersEncoding = galahadthepureDB.GalahadThePurePointersEncoding
		galahadthepureDB.CopyBasicFieldsToGalahadThePure_WOP(&galahadthepureAPI.GalahadThePure_WOP)

		backRepoData.GalahadThePureAPIs = append(backRepoData.GalahadThePureAPIs, &galahadthepureAPI)
	}

	for _, geoobjectDB := range backRepo.BackRepoGeoObject.Map_GeoObjectDBID_GeoObjectDB {

		var geoobjectAPI GeoObjectAPI
		geoobjectAPI.ID = geoobjectDB.ID
		geoobjectAPI.GeoObjectPointersEncoding = geoobjectDB.GeoObjectPointersEncoding
		geoobjectDB.CopyBasicFieldsToGeoObject_WOP(&geoobjectAPI.GeoObject_WOP)

		backRepoData.GeoObjectAPIs = append(backRepoData.GeoObjectAPIs, &geoobjectAPI)
	}

	for _, geoobjectuseDB := range backRepo.BackRepoGeoObjectUse.Map_GeoObjectUseDBID_GeoObjectUseDB {

		var geoobjectuseAPI GeoObjectUseAPI
		geoobjectuseAPI.ID = geoobjectuseDB.ID
		geoobjectuseAPI.GeoObjectUsePointersEncoding = geoobjectuseDB.GeoObjectUsePointersEncoding
		geoobjectuseDB.CopyBasicFieldsToGeoObjectUse_WOP(&geoobjectuseAPI.GeoObjectUse_WOP)

		backRepoData.GeoObjectUseAPIs = append(backRepoData.GeoObjectUseAPIs, &geoobjectuseAPI)
	}

	for _, groupDB := range backRepo.BackRepoGroup.Map_GroupDBID_GroupDB {

		var groupAPI GroupAPI
		groupAPI.ID = groupDB.ID
		groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
		groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)

		backRepoData.GroupAPIs = append(backRepoData.GroupAPIs, &groupAPI)
	}

	for _, groupuseDB := range backRepo.BackRepoGroupUse.Map_GroupUseDBID_GroupUseDB {

		var groupuseAPI GroupUseAPI
		groupuseAPI.ID = groupuseDB.ID
		groupuseAPI.GroupUsePointersEncoding = groupuseDB.GroupUsePointersEncoding
		groupuseDB.CopyBasicFieldsToGroupUse_WOP(&groupuseAPI.GroupUse_WOP)

		backRepoData.GroupUseAPIs = append(backRepoData.GroupUseAPIs, &groupuseAPI)
	}

	for _, kingarthurDB := range backRepo.BackRepoKingArthur.Map_KingArthurDBID_KingArthurDB {

		var kingarthurAPI KingArthurAPI
		kingarthurAPI.ID = kingarthurDB.ID
		kingarthurAPI.KingArthurPointersEncoding = kingarthurDB.KingArthurPointersEncoding
		kingarthurDB.CopyBasicFieldsToKingArthur_WOP(&kingarthurAPI.KingArthur_WOP)

		backRepoData.KingArthurAPIs = append(backRepoData.KingArthurAPIs, &kingarthurAPI)
	}

	for _, kingarthurshapeDB := range backRepo.BackRepoKingArthurShape.Map_KingArthurShapeDBID_KingArthurShapeDB {

		var kingarthurshapeAPI KingArthurShapeAPI
		kingarthurshapeAPI.ID = kingarthurshapeDB.ID
		kingarthurshapeAPI.KingArthurShapePointersEncoding = kingarthurshapeDB.KingArthurShapePointersEncoding
		kingarthurshapeDB.CopyBasicFieldsToKingArthurShape_WOP(&kingarthurshapeAPI.KingArthurShape_WOP)

		backRepoData.KingArthurShapeAPIs = append(backRepoData.KingArthurShapeAPIs, &kingarthurshapeAPI)
	}

	for _, knightwhosayniDB := range backRepo.BackRepoKnightWhoSayNi.Map_KnightWhoSayNiDBID_KnightWhoSayNiDB {

		var knightwhosayniAPI KnightWhoSayNiAPI
		knightwhosayniAPI.ID = knightwhosayniDB.ID
		knightwhosayniAPI.KnightWhoSayNiPointersEncoding = knightwhosayniDB.KnightWhoSayNiPointersEncoding
		knightwhosayniDB.CopyBasicFieldsToKnightWhoSayNi_WOP(&knightwhosayniAPI.KnightWhoSayNi_WOP)

		backRepoData.KnightWhoSayNiAPIs = append(backRepoData.KnightWhoSayNiAPIs, &knightwhosayniAPI)
	}

	for _, lancelotDB := range backRepo.BackRepoLancelot.Map_LancelotDBID_LancelotDB {

		var lancelotAPI LancelotAPI
		lancelotAPI.ID = lancelotDB.ID
		lancelotAPI.LancelotPointersEncoding = lancelotDB.LancelotPointersEncoding
		lancelotDB.CopyBasicFieldsToLancelot_WOP(&lancelotAPI.Lancelot_WOP)

		backRepoData.LancelotAPIs = append(backRepoData.LancelotAPIs, &lancelotAPI)
	}

	for _, lancelotagregationDB := range backRepo.BackRepoLancelotAgregation.Map_LancelotAgregationDBID_LancelotAgregationDB {

		var lancelotagregationAPI LancelotAgregationAPI
		lancelotagregationAPI.ID = lancelotagregationDB.ID
		lancelotagregationAPI.LancelotAgregationPointersEncoding = lancelotagregationDB.LancelotAgregationPointersEncoding
		lancelotagregationDB.CopyBasicFieldsToLancelotAgregation_WOP(&lancelotagregationAPI.LancelotAgregation_WOP)

		backRepoData.LancelotAgregationAPIs = append(backRepoData.LancelotAgregationAPIs, &lancelotagregationAPI)
	}

	for _, lancelotagregationuseDB := range backRepo.BackRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB {

		var lancelotagregationuseAPI LancelotAgregationUseAPI
		lancelotagregationuseAPI.ID = lancelotagregationuseDB.ID
		lancelotagregationuseAPI.LancelotAgregationUsePointersEncoding = lancelotagregationuseDB.LancelotAgregationUsePointersEncoding
		lancelotagregationuseDB.CopyBasicFieldsToLancelotAgregationUse_WOP(&lancelotagregationuseAPI.LancelotAgregationUse_WOP)

		backRepoData.LancelotAgregationUseAPIs = append(backRepoData.LancelotAgregationUseAPIs, &lancelotagregationuseAPI)
	}

	for _, lancelotcategoryDB := range backRepo.BackRepoLancelotCategory.Map_LancelotCategoryDBID_LancelotCategoryDB {

		var lancelotcategoryAPI LancelotCategoryAPI
		lancelotcategoryAPI.ID = lancelotcategoryDB.ID
		lancelotcategoryAPI.LancelotCategoryPointersEncoding = lancelotcategoryDB.LancelotCategoryPointersEncoding
		lancelotcategoryDB.CopyBasicFieldsToLancelotCategory_WOP(&lancelotcategoryAPI.LancelotCategory_WOP)

		backRepoData.LancelotCategoryAPIs = append(backRepoData.LancelotCategoryAPIs, &lancelotcategoryAPI)
	}

	for _, lancelotcategoryuseDB := range backRepo.BackRepoLancelotCategoryUse.Map_LancelotCategoryUseDBID_LancelotCategoryUseDB {

		var lancelotcategoryuseAPI LancelotCategoryUseAPI
		lancelotcategoryuseAPI.ID = lancelotcategoryuseDB.ID
		lancelotcategoryuseAPI.LancelotCategoryUsePointersEncoding = lancelotcategoryuseDB.LancelotCategoryUsePointersEncoding
		lancelotcategoryuseDB.CopyBasicFieldsToLancelotCategoryUse_WOP(&lancelotcategoryuseAPI.LancelotCategoryUse_WOP)

		backRepoData.LancelotCategoryUseAPIs = append(backRepoData.LancelotCategoryUseAPIs, &lancelotcategoryuseAPI)
	}

	for _, mapobjectDB := range backRepo.BackRepoMapObject.Map_MapObjectDBID_MapObjectDB {

		var mapobjectAPI MapObjectAPI
		mapobjectAPI.ID = mapobjectDB.ID
		mapobjectAPI.MapObjectPointersEncoding = mapobjectDB.MapObjectPointersEncoding
		mapobjectDB.CopyBasicFieldsToMapObject_WOP(&mapobjectAPI.MapObject_WOP)

		backRepoData.MapObjectAPIs = append(backRepoData.MapObjectAPIs, &mapobjectAPI)
	}

	for _, mapobjectuseDB := range backRepo.BackRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB {

		var mapobjectuseAPI MapObjectUseAPI
		mapobjectuseAPI.ID = mapobjectuseDB.ID
		mapobjectuseAPI.MapObjectUsePointersEncoding = mapobjectuseDB.MapObjectUsePointersEncoding
		mapobjectuseDB.CopyBasicFieldsToMapObjectUse_WOP(&mapobjectuseAPI.MapObjectUse_WOP)

		backRepoData.MapObjectUseAPIs = append(backRepoData.MapObjectUseAPIs, &mapobjectuseAPI)
	}

	for _, repositoryDB := range backRepo.BackRepoRepository.Map_RepositoryDBID_RepositoryDB {

		var repositoryAPI RepositoryAPI
		repositoryAPI.ID = repositoryDB.ID
		repositoryAPI.RepositoryPointersEncoding = repositoryDB.RepositoryPointersEncoding
		repositoryDB.CopyBasicFieldsToRepository_WOP(&repositoryAPI.Repository_WOP)

		backRepoData.RepositoryAPIs = append(backRepoData.RepositoryAPIs, &repositoryAPI)
	}

	for _, sirrobinDB := range backRepo.BackRepoSirRobin.Map_SirRobinDBID_SirRobinDB {

		var sirrobinAPI SirRobinAPI
		sirrobinAPI.ID = sirrobinDB.ID
		sirrobinAPI.SirRobinPointersEncoding = sirrobinDB.SirRobinPointersEncoding
		sirrobinDB.CopyBasicFieldsToSirRobin_WOP(&sirrobinAPI.SirRobin_WOP)

		backRepoData.SirRobinAPIs = append(backRepoData.SirRobinAPIs, &sirrobinAPI)
	}

	for _, thebridgeDB := range backRepo.BackRepoTheBridge.Map_TheBridgeDBID_TheBridgeDB {

		var thebridgeAPI TheBridgeAPI
		thebridgeAPI.ID = thebridgeDB.ID
		thebridgeAPI.TheBridgePointersEncoding = thebridgeDB.TheBridgePointersEncoding
		thebridgeDB.CopyBasicFieldsToTheBridge_WOP(&thebridgeAPI.TheBridge_WOP)

		backRepoData.TheBridgeAPIs = append(backRepoData.TheBridgeAPIs, &thebridgeAPI)
	}

	for _, thenuteshapeDB := range backRepo.BackRepoTheNuteShape.Map_TheNuteShapeDBID_TheNuteShapeDB {

		var thenuteshapeAPI TheNuteShapeAPI
		thenuteshapeAPI.ID = thenuteshapeDB.ID
		thenuteshapeAPI.TheNuteShapePointersEncoding = thenuteshapeDB.TheNuteShapePointersEncoding
		thenuteshapeDB.CopyBasicFieldsToTheNuteShape_WOP(&thenuteshapeAPI.TheNuteShape_WOP)

		backRepoData.TheNuteShapeAPIs = append(backRepoData.TheNuteShapeAPIs, &thenuteshapeAPI)
	}

	for _, thenutetransitionDB := range backRepo.BackRepoTheNuteTransition.Map_TheNuteTransitionDBID_TheNuteTransitionDB {

		var thenutetransitionAPI TheNuteTransitionAPI
		thenutetransitionAPI.ID = thenutetransitionDB.ID
		thenutetransitionAPI.TheNuteTransitionPointersEncoding = thenutetransitionDB.TheNuteTransitionPointersEncoding
		thenutetransitionDB.CopyBasicFieldsToTheNuteTransition_WOP(&thenutetransitionAPI.TheNuteTransition_WOP)

		backRepoData.TheNuteTransitionAPIs = append(backRepoData.TheNuteTransitionAPIs, &thenutetransitionAPI)
	}

	for _, userDB := range backRepo.BackRepoUser.Map_UserDBID_UserDB {

		var userAPI UserAPI
		userAPI.ID = userDB.ID
		userAPI.UserPointersEncoding = userDB.UserPointersEncoding
		userDB.CopyBasicFieldsToUser_WOP(&userAPI.User_WOP)

		backRepoData.UserAPIs = append(backRepoData.UserAPIs, &userAPI)
	}

	for _, useruseDB := range backRepo.BackRepoUserUse.Map_UserUseDBID_UserUseDB {

		var useruseAPI UserUseAPI
		useruseAPI.ID = useruseDB.ID
		useruseAPI.UserUsePointersEncoding = useruseDB.UserUsePointersEncoding
		useruseDB.CopyBasicFieldsToUserUse_WOP(&useruseAPI.UserUse_WOP)

		backRepoData.UserUseAPIs = append(backRepoData.UserUseAPIs, &useruseAPI)
	}

	for _, whatisyourpreferedcolorDB := range backRepo.BackRepoWhatIsYourPreferedColor.Map_WhatIsYourPreferedColorDBID_WhatIsYourPreferedColorDB {

		var whatisyourpreferedcolorAPI WhatIsYourPreferedColorAPI
		whatisyourpreferedcolorAPI.ID = whatisyourpreferedcolorDB.ID
		whatisyourpreferedcolorAPI.WhatIsYourPreferedColorPointersEncoding = whatisyourpreferedcolorDB.WhatIsYourPreferedColorPointersEncoding
		whatisyourpreferedcolorDB.CopyBasicFieldsToWhatIsYourPreferedColor_WOP(&whatisyourpreferedcolorAPI.WhatIsYourPreferedColor_WOP)

		backRepoData.WhatIsYourPreferedColorAPIs = append(backRepoData.WhatIsYourPreferedColorAPIs, &whatisyourpreferedcolorAPI)
	}

	for _, workspaceDB := range backRepo.BackRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB {

		var workspaceAPI WorkspaceAPI
		workspaceAPI.ID = workspaceDB.ID
		workspaceAPI.WorkspacePointersEncoding = workspaceDB.WorkspacePointersEncoding
		workspaceDB.CopyBasicFieldsToWorkspace_WOP(&workspaceAPI.Workspace_WOP)

		backRepoData.WorkspaceAPIs = append(backRepoData.WorkspaceAPIs, &workspaceAPI)
	}

}
