// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/thelongbuild/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.AWitch:
		tmp := GetInstanceDBFromInstance[models.AWitch, AWitchDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "Witches":
				if _sirrobin, ok := stage.SirRobin_Witches_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.BlackKnightShape:
		tmp := GetInstanceDBFromInstance[models.BlackKnightShape, BlackKnightShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "BlackKnightShapes":
				if _sirrobin, ok := stage.SirRobin_BlackKnightShapes_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.BringYourDead:
		tmp := GetInstanceDBFromInstance[models.BringYourDead, BringYourDeadDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "BringYourDeadarameters":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_BringYourDeadarameters_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Document:
		tmp := GetInstanceDBFromInstance[models.Document, DocumentDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.DocumentUse:
		tmp := GetInstanceDBFromInstance[models.DocumentUse, DocumentUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			case "DocumentUse":
				if _lancelot, ok := stage.Lancelot_DocumentUse_reverseMap[inst]; ok {
					res = _lancelot.Name
				}
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.GalahadThePure:
		tmp := GetInstanceDBFromInstance[models.GalahadThePure, GalahadThePureDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Galahard":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_Galahard_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.GeoObject:
		tmp := GetInstanceDBFromInstance[models.GeoObject, GeoObjectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.GeoObjectUse:
		tmp := GetInstanceDBFromInstance[models.GeoObjectUse, GeoObjectUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				if _document, ok := stage.Document_GeoObjectUse_reverseMap[inst]; ok {
					res = _document.Name
				}
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				if _lancelot, ok := stage.Lancelot_GeoObjectUse_reverseMap[inst]; ok {
					res = _lancelot.Name
				}
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				if _thebridge, ok := stage.TheBridge_GeoObjectUse_reverseMap[inst]; ok {
					res = _thebridge.Name
				}
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.GroupUse:
		tmp := GetInstanceDBFromInstance[models.GroupUse, GroupUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			case "GroupUse":
				if _lancelot, ok := stage.Lancelot_GroupUse_reverseMap[inst]; ok {
					res = _lancelot.Name
				}
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			case "GroupUse":
				if _repository, ok := stage.Repository_GroupUse_reverseMap[inst]; ok {
					res = _repository.Name
				}
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "GroupUse":
				if _thebridge, ok := stage.TheBridge_GroupUse_reverseMap[inst]; ok {
					res = _thebridge.Name
				}
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.KingArthur:
		tmp := GetInstanceDBFromInstance[models.KingArthur, KingArthurDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "KingArthurs":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_KingArthurs_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.KingArthurShape:
		tmp := GetInstanceDBFromInstance[models.KingArthurShape, KingArthurShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "Arthurs":
				if _sirrobin, ok := stage.SirRobin_Arthurs_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.KnightWhoSayNi:
		tmp := GetInstanceDBFromInstance[models.KnightWhoSayNi, KnightWhoSayNiDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			case "ParameterUse":
				if _lancelotagregation, ok := stage.LancelotAgregation_ParameterUse_reverseMap[inst]; ok {
					res = _lancelotagregation.Name
				}
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			case "ParameterUse":
				if _lancelotcategory, ok := stage.LancelotCategory_ParameterUse_reverseMap[inst]; ok {
					res = _lancelotcategory.Name
				}
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			case "ParameterUse":
				if _repository, ok := stage.Repository_ParameterUse_reverseMap[inst]; ok {
					res = _repository.Name
				}
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "KnightWhoSayNis":
				if _sirrobin, ok := stage.SirRobin_KnightWhoSayNis_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Lancelot:
		tmp := GetInstanceDBFromInstance[models.Lancelot, LancelotDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			case "Lancelots":
				if _bringyourdead, ok := stage.BringYourDead_Lancelots_reverseMap[inst]; ok {
					res = _bringyourdead.Name
				}
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Lancelots":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_Lancelots_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.LancelotAgregation:
		tmp := GetInstanceDBFromInstance[models.LancelotAgregation, LancelotAgregationDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.LancelotAgregationUse:
		tmp := GetInstanceDBFromInstance[models.LancelotAgregationUse, LancelotAgregationUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.LancelotCategory:
		tmp := GetInstanceDBFromInstance[models.LancelotCategory, LancelotCategoryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.LancelotCategoryUse:
		tmp := GetInstanceDBFromInstance[models.LancelotCategoryUse, LancelotCategoryUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.MapObject:
		tmp := GetInstanceDBFromInstance[models.MapObject, MapObjectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.MapObjectUse:
		tmp := GetInstanceDBFromInstance[models.MapObjectUse, MapObjectUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "MapUse":
				if _thebridge, ok := stage.TheBridge_MapUse_reverseMap[inst]; ok {
					res = _thebridge.Name
				}
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Repository:
		tmp := GetInstanceDBFromInstance[models.Repository, RepositoryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.SirRobin:
		tmp := GetInstanceDBFromInstance[models.SirRobin, SirRobinDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Diagrams":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_Diagrams_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.TheBridge:
		tmp := GetInstanceDBFromInstance[models.TheBridge, TheBridgeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.TheNuteShape:
		tmp := GetInstanceDBFromInstance[models.TheNuteShape, TheNuteShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "TheNuteShapes":
				if _sirrobin, ok := stage.SirRobin_TheNuteShapes_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.TheNuteTransition:
		tmp := GetInstanceDBFromInstance[models.TheNuteTransition, TheNuteTransitionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Nutes":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_Nutes_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.User:
		tmp := GetInstanceDBFromInstance[models.User, UserDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.UserUse:
		tmp := GetInstanceDBFromInstance[models.UserUse, UserUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			case "UserUse":
				if _group, ok := stage.Group_UserUse_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.WhatIsYourPreferedColor:
		tmp := GetInstanceDBFromInstance[models.WhatIsYourPreferedColor, WhatIsYourPreferedColorDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "WhatIsYourPreferedColor":
				if _thebridge, ok := stage.TheBridge_WhatIsYourPreferedColor_reverseMap[inst]; ok {
					res = _thebridge.Name
				}
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Workspace:
		tmp := GetInstanceDBFromInstance[models.Workspace, WorkspaceDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.AWitch:
		tmp := GetInstanceDBFromInstance[models.AWitch, AWitchDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "Witches":
				res = stage.SirRobin_Witches_reverseMap[inst]
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.BlackKnightShape:
		tmp := GetInstanceDBFromInstance[models.BlackKnightShape, BlackKnightShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "BlackKnightShapes":
				res = stage.SirRobin_BlackKnightShapes_reverseMap[inst]
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.BringYourDead:
		tmp := GetInstanceDBFromInstance[models.BringYourDead, BringYourDeadDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "BringYourDeadarameters":
				res = stage.WhatIsYourPreferedColor_BringYourDeadarameters_reverseMap[inst]
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Document:
		tmp := GetInstanceDBFromInstance[models.Document, DocumentDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.DocumentUse:
		tmp := GetInstanceDBFromInstance[models.DocumentUse, DocumentUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			case "DocumentUse":
				res = stage.Lancelot_DocumentUse_reverseMap[inst]
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.GalahadThePure:
		tmp := GetInstanceDBFromInstance[models.GalahadThePure, GalahadThePureDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Galahard":
				res = stage.WhatIsYourPreferedColor_Galahard_reverseMap[inst]
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.GeoObject:
		tmp := GetInstanceDBFromInstance[models.GeoObject, GeoObjectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.GeoObjectUse:
		tmp := GetInstanceDBFromInstance[models.GeoObjectUse, GeoObjectUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				res = stage.Document_GeoObjectUse_reverseMap[inst]
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				res = stage.Lancelot_GeoObjectUse_reverseMap[inst]
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				res = stage.TheBridge_GeoObjectUse_reverseMap[inst]
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.GroupUse:
		tmp := GetInstanceDBFromInstance[models.GroupUse, GroupUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			case "GroupUse":
				res = stage.Lancelot_GroupUse_reverseMap[inst]
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			case "GroupUse":
				res = stage.Repository_GroupUse_reverseMap[inst]
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "GroupUse":
				res = stage.TheBridge_GroupUse_reverseMap[inst]
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.KingArthur:
		tmp := GetInstanceDBFromInstance[models.KingArthur, KingArthurDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "KingArthurs":
				res = stage.WhatIsYourPreferedColor_KingArthurs_reverseMap[inst]
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.KingArthurShape:
		tmp := GetInstanceDBFromInstance[models.KingArthurShape, KingArthurShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "Arthurs":
				res = stage.SirRobin_Arthurs_reverseMap[inst]
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.KnightWhoSayNi:
		tmp := GetInstanceDBFromInstance[models.KnightWhoSayNi, KnightWhoSayNiDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			case "ParameterUse":
				res = stage.LancelotAgregation_ParameterUse_reverseMap[inst]
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			case "ParameterUse":
				res = stage.LancelotCategory_ParameterUse_reverseMap[inst]
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			case "ParameterUse":
				res = stage.Repository_ParameterUse_reverseMap[inst]
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "KnightWhoSayNis":
				res = stage.SirRobin_KnightWhoSayNis_reverseMap[inst]
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Lancelot:
		tmp := GetInstanceDBFromInstance[models.Lancelot, LancelotDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			case "Lancelots":
				res = stage.BringYourDead_Lancelots_reverseMap[inst]
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Lancelots":
				res = stage.WhatIsYourPreferedColor_Lancelots_reverseMap[inst]
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.LancelotAgregation:
		tmp := GetInstanceDBFromInstance[models.LancelotAgregation, LancelotAgregationDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.LancelotAgregationUse:
		tmp := GetInstanceDBFromInstance[models.LancelotAgregationUse, LancelotAgregationUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.LancelotCategory:
		tmp := GetInstanceDBFromInstance[models.LancelotCategory, LancelotCategoryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.LancelotCategoryUse:
		tmp := GetInstanceDBFromInstance[models.LancelotCategoryUse, LancelotCategoryUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.MapObject:
		tmp := GetInstanceDBFromInstance[models.MapObject, MapObjectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.MapObjectUse:
		tmp := GetInstanceDBFromInstance[models.MapObjectUse, MapObjectUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "MapUse":
				res = stage.TheBridge_MapUse_reverseMap[inst]
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Repository:
		tmp := GetInstanceDBFromInstance[models.Repository, RepositoryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.SirRobin:
		tmp := GetInstanceDBFromInstance[models.SirRobin, SirRobinDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Diagrams":
				res = stage.WhatIsYourPreferedColor_Diagrams_reverseMap[inst]
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.TheBridge:
		tmp := GetInstanceDBFromInstance[models.TheBridge, TheBridgeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.TheNuteShape:
		tmp := GetInstanceDBFromInstance[models.TheNuteShape, TheNuteShapeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			case "TheNuteShapes":
				res = stage.SirRobin_TheNuteShapes_reverseMap[inst]
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.TheNuteTransition:
		tmp := GetInstanceDBFromInstance[models.TheNuteTransition, TheNuteTransitionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Nutes":
				res = stage.WhatIsYourPreferedColor_Nutes_reverseMap[inst]
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.User:
		tmp := GetInstanceDBFromInstance[models.User, UserDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.UserUse:
		tmp := GetInstanceDBFromInstance[models.UserUse, UserUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			case "UserUse":
				res = stage.Group_UserUse_reverseMap[inst]
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.WhatIsYourPreferedColor:
		tmp := GetInstanceDBFromInstance[models.WhatIsYourPreferedColor, WhatIsYourPreferedColorDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "WhatIsYourPreferedColor":
				res = stage.TheBridge_WhatIsYourPreferedColor_reverseMap[inst]
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	case *models.Workspace:
		tmp := GetInstanceDBFromInstance[models.Workspace, WorkspaceDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "AWitch":
			switch reverseField.Fieldname {
			}
		case "BlackKnightShape":
			switch reverseField.Fieldname {
			}
		case "BringYourDead":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "DocumentUse":
			switch reverseField.Fieldname {
			}
		case "GalahadThePure":
			switch reverseField.Fieldname {
			}
		case "GeoObject":
			switch reverseField.Fieldname {
			}
		case "GeoObjectUse":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "GroupUse":
			switch reverseField.Fieldname {
			}
		case "KingArthur":
			switch reverseField.Fieldname {
			}
		case "KingArthurShape":
			switch reverseField.Fieldname {
			}
		case "KnightWhoSayNi":
			switch reverseField.Fieldname {
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			}
		case "LancelotAgregationUse":
			switch reverseField.Fieldname {
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			}
		case "LancelotCategoryUse":
			switch reverseField.Fieldname {
			}
		case "MapObject":
			switch reverseField.Fieldname {
			}
		case "MapObjectUse":
			switch reverseField.Fieldname {
			}
		case "Repository":
			switch reverseField.Fieldname {
			}
		case "SirRobin":
			switch reverseField.Fieldname {
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			}
		case "TheNuteShape":
			switch reverseField.Fieldname {
			}
		case "TheNuteTransition":
			switch reverseField.Fieldname {
			}
		case "User":
			switch reverseField.Fieldname {
			}
		case "UserUse":
			switch reverseField.Fieldname {
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			}
		case "Workspace":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
