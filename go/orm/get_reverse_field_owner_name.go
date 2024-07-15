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
		switch reverseField.GongstructName {
		// insertion point
		case "SirRobin":
			switch reverseField.Fieldname {
			case "Witches":
				if _sirrobin, ok := stage.SirRobin_Witches_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		}

	case *models.BlackKnightShape:
		switch reverseField.GongstructName {
		// insertion point
		case "SirRobin":
			switch reverseField.Fieldname {
			case "BlackKnightShapes":
				if _sirrobin, ok := stage.SirRobin_BlackKnightShapes_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		}

	case *models.BringYourDead:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "BringYourDeadarameters":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_BringYourDeadarameters_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		}

	case *models.Document:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DocumentUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Lancelot":
			switch reverseField.Fieldname {
			case "DocumentUse":
				if _lancelot, ok := stage.Lancelot_DocumentUse_reverseMap[inst]; ok {
					res = _lancelot.Name
				}
			}
		}

	case *models.GalahadThePure:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Galahard":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_Galahard_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		}

	case *models.GeoObject:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GeoObjectUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Document":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				if _document, ok := stage.Document_GeoObjectUse_reverseMap[inst]; ok {
					res = _document.Name
				}
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				if _lancelot, ok := stage.Lancelot_GeoObjectUse_reverseMap[inst]; ok {
					res = _lancelot.Name
				}
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				if _thebridge, ok := stage.TheBridge_GeoObjectUse_reverseMap[inst]; ok {
					res = _thebridge.Name
				}
			}
		}

	case *models.Group:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GroupUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Lancelot":
			switch reverseField.Fieldname {
			case "GroupUse":
				if _lancelot, ok := stage.Lancelot_GroupUse_reverseMap[inst]; ok {
					res = _lancelot.Name
				}
			}
		case "Repository":
			switch reverseField.Fieldname {
			case "GroupUse":
				if _repository, ok := stage.Repository_GroupUse_reverseMap[inst]; ok {
					res = _repository.Name
				}
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "GroupUse":
				if _thebridge, ok := stage.TheBridge_GroupUse_reverseMap[inst]; ok {
					res = _thebridge.Name
				}
			}
		}

	case *models.KingArthur:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "KingArthurs":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_KingArthurs_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		}

	case *models.KingArthurShape:
		switch reverseField.GongstructName {
		// insertion point
		case "SirRobin":
			switch reverseField.Fieldname {
			case "Arthurs":
				if _sirrobin, ok := stage.SirRobin_Arthurs_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		}

	case *models.KnightWhoSayNi:
		switch reverseField.GongstructName {
		// insertion point
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			case "ParameterUse":
				if _lancelotagregation, ok := stage.LancelotAgregation_ParameterUse_reverseMap[inst]; ok {
					res = _lancelotagregation.Name
				}
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			case "ParameterUse":
				if _lancelotcategory, ok := stage.LancelotCategory_ParameterUse_reverseMap[inst]; ok {
					res = _lancelotcategory.Name
				}
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
		}

	case *models.Lancelot:
		switch reverseField.GongstructName {
		// insertion point
		case "BringYourDead":
			switch reverseField.Fieldname {
			case "Lancelots":
				if _bringyourdead, ok := stage.BringYourDead_Lancelots_reverseMap[inst]; ok {
					res = _bringyourdead.Name
				}
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Lancelots":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_Lancelots_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		}

	case *models.LancelotAgregation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LancelotAgregationUse:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LancelotCategory:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LancelotCategoryUse:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MapObject:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MapObjectUse:
		switch reverseField.GongstructName {
		// insertion point
		case "TheBridge":
			switch reverseField.Fieldname {
			case "MapUse":
				if _thebridge, ok := stage.TheBridge_MapUse_reverseMap[inst]; ok {
					res = _thebridge.Name
				}
			}
		}

	case *models.Repository:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SirRobin:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Diagrams":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_Diagrams_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		}

	case *models.TheBridge:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TheNuteShape:
		switch reverseField.GongstructName {
		// insertion point
		case "SirRobin":
			switch reverseField.Fieldname {
			case "TheNuteShapes":
				if _sirrobin, ok := stage.SirRobin_TheNuteShapes_reverseMap[inst]; ok {
					res = _sirrobin.Name
				}
			}
		}

	case *models.TheNuteTransition:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Nutes":
				if _whatisyourpreferedcolor, ok := stage.WhatIsYourPreferedColor_Nutes_reverseMap[inst]; ok {
					res = _whatisyourpreferedcolor.Name
				}
			}
		}

	case *models.User:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.UserUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "UserUse":
				if _group, ok := stage.Group_UserUse_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		}

	case *models.WhatIsYourPreferedColor:
		switch reverseField.GongstructName {
		// insertion point
		case "TheBridge":
			switch reverseField.Fieldname {
			case "WhatIsYourPreferedColor":
				if _thebridge, ok := stage.TheBridge_WhatIsYourPreferedColor_reverseMap[inst]; ok {
					res = _thebridge.Name
				}
			}
		}

	case *models.Workspace:
		switch reverseField.GongstructName {
		// insertion point
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
		switch reverseField.GongstructName {
		// insertion point
		case "SirRobin":
			switch reverseField.Fieldname {
			case "Witches":
				res = stage.SirRobin_Witches_reverseMap[inst]
			}
		}

	case *models.BlackKnightShape:
		switch reverseField.GongstructName {
		// insertion point
		case "SirRobin":
			switch reverseField.Fieldname {
			case "BlackKnightShapes":
				res = stage.SirRobin_BlackKnightShapes_reverseMap[inst]
			}
		}

	case *models.BringYourDead:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "BringYourDeadarameters":
				res = stage.WhatIsYourPreferedColor_BringYourDeadarameters_reverseMap[inst]
			}
		}

	case *models.Document:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DocumentUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Lancelot":
			switch reverseField.Fieldname {
			case "DocumentUse":
				res = stage.Lancelot_DocumentUse_reverseMap[inst]
			}
		}

	case *models.GalahadThePure:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Galahard":
				res = stage.WhatIsYourPreferedColor_Galahard_reverseMap[inst]
			}
		}

	case *models.GeoObject:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GeoObjectUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Document":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				res = stage.Document_GeoObjectUse_reverseMap[inst]
			}
		case "Lancelot":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				res = stage.Lancelot_GeoObjectUse_reverseMap[inst]
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "GeoObjectUse":
				res = stage.TheBridge_GeoObjectUse_reverseMap[inst]
			}
		}

	case *models.Group:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.GroupUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Lancelot":
			switch reverseField.Fieldname {
			case "GroupUse":
				res = stage.Lancelot_GroupUse_reverseMap[inst]
			}
		case "Repository":
			switch reverseField.Fieldname {
			case "GroupUse":
				res = stage.Repository_GroupUse_reverseMap[inst]
			}
		case "TheBridge":
			switch reverseField.Fieldname {
			case "GroupUse":
				res = stage.TheBridge_GroupUse_reverseMap[inst]
			}
		}

	case *models.KingArthur:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "KingArthurs":
				res = stage.WhatIsYourPreferedColor_KingArthurs_reverseMap[inst]
			}
		}

	case *models.KingArthurShape:
		switch reverseField.GongstructName {
		// insertion point
		case "SirRobin":
			switch reverseField.Fieldname {
			case "Arthurs":
				res = stage.SirRobin_Arthurs_reverseMap[inst]
			}
		}

	case *models.KnightWhoSayNi:
		switch reverseField.GongstructName {
		// insertion point
		case "LancelotAgregation":
			switch reverseField.Fieldname {
			case "ParameterUse":
				res = stage.LancelotAgregation_ParameterUse_reverseMap[inst]
			}
		case "LancelotCategory":
			switch reverseField.Fieldname {
			case "ParameterUse":
				res = stage.LancelotCategory_ParameterUse_reverseMap[inst]
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
		}

	case *models.Lancelot:
		switch reverseField.GongstructName {
		// insertion point
		case "BringYourDead":
			switch reverseField.Fieldname {
			case "Lancelots":
				res = stage.BringYourDead_Lancelots_reverseMap[inst]
			}
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Lancelots":
				res = stage.WhatIsYourPreferedColor_Lancelots_reverseMap[inst]
			}
		}

	case *models.LancelotAgregation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LancelotAgregationUse:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LancelotCategory:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.LancelotCategoryUse:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MapObject:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MapObjectUse:
		switch reverseField.GongstructName {
		// insertion point
		case "TheBridge":
			switch reverseField.Fieldname {
			case "MapUse":
				res = stage.TheBridge_MapUse_reverseMap[inst]
			}
		}

	case *models.Repository:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SirRobin:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Diagrams":
				res = stage.WhatIsYourPreferedColor_Diagrams_reverseMap[inst]
			}
		}

	case *models.TheBridge:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TheNuteShape:
		switch reverseField.GongstructName {
		// insertion point
		case "SirRobin":
			switch reverseField.Fieldname {
			case "TheNuteShapes":
				res = stage.SirRobin_TheNuteShapes_reverseMap[inst]
			}
		}

	case *models.TheNuteTransition:
		switch reverseField.GongstructName {
		// insertion point
		case "WhatIsYourPreferedColor":
			switch reverseField.Fieldname {
			case "Nutes":
				res = stage.WhatIsYourPreferedColor_Nutes_reverseMap[inst]
			}
		}

	case *models.User:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.UserUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "UserUse":
				res = stage.Group_UserUse_reverseMap[inst]
			}
		}

	case *models.WhatIsYourPreferedColor:
		switch reverseField.GongstructName {
		// insertion point
		case "TheBridge":
			switch reverseField.Fieldname {
			case "WhatIsYourPreferedColor":
				res = stage.TheBridge_WhatIsYourPreferedColor_reverseMap[inst]
			}
		}

	case *models.Workspace:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
