// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *AWitch:
		// insertion point per field

	case *BlackKnightShape:
		// insertion point per field

	case *BringYourDead:
		// insertion point per field
		if fieldName == "Lancelots" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*BringYourDead) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*BringYourDead)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Lancelots).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Lancelots = _inferedTypeInstance.Lancelots[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Lancelots =
								append(_inferedTypeInstance.Lancelots, any(fieldInstance).(*Lancelot))
						}
					}
				}
			}
		}

	case *Document:
		// insertion point per field
		if fieldName == "GeoObjectUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Document) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Document)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GeoObjectUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GeoObjectUse = _inferedTypeInstance.GeoObjectUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GeoObjectUse =
								append(_inferedTypeInstance.GeoObjectUse, any(fieldInstance).(*GeoObjectUse))
						}
					}
				}
			}
		}

	case *DocumentUse:
		// insertion point per field

	case *GalahadThePure:
		// insertion point per field

	case *GeoObject:
		// insertion point per field

	case *GeoObjectUse:
		// insertion point per field

	case *Group:
		// insertion point per field
		if fieldName == "UserUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Group) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Group)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.UserUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.UserUse = _inferedTypeInstance.UserUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.UserUse =
								append(_inferedTypeInstance.UserUse, any(fieldInstance).(*UserUse))
						}
					}
				}
			}
		}

	case *GroupUse:
		// insertion point per field

	case *KingArthur:
		// insertion point per field

	case *KingArthurShape:
		// insertion point per field

	case *KnightWhoSayNi:
		// insertion point per field

	case *Lancelot:
		// insertion point per field
		if fieldName == "GroupUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Lancelot) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Lancelot)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GroupUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GroupUse = _inferedTypeInstance.GroupUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GroupUse =
								append(_inferedTypeInstance.GroupUse, any(fieldInstance).(*GroupUse))
						}
					}
				}
			}
		}
		if fieldName == "DocumentUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Lancelot) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Lancelot)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DocumentUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DocumentUse = _inferedTypeInstance.DocumentUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DocumentUse =
								append(_inferedTypeInstance.DocumentUse, any(fieldInstance).(*DocumentUse))
						}
					}
				}
			}
		}
		if fieldName == "GeoObjectUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Lancelot) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Lancelot)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GeoObjectUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GeoObjectUse = _inferedTypeInstance.GeoObjectUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GeoObjectUse =
								append(_inferedTypeInstance.GeoObjectUse, any(fieldInstance).(*GeoObjectUse))
						}
					}
				}
			}
		}

	case *LancelotAgregation:
		// insertion point per field
		if fieldName == "ParameterUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*LancelotAgregation) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*LancelotAgregation)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ParameterUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ParameterUse = _inferedTypeInstance.ParameterUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ParameterUse =
								append(_inferedTypeInstance.ParameterUse, any(fieldInstance).(*KnightWhoSayNi))
						}
					}
				}
			}
		}

	case *LancelotAgregationUse:
		// insertion point per field

	case *LancelotCategory:
		// insertion point per field
		if fieldName == "ParameterUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*LancelotCategory) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*LancelotCategory)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ParameterUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ParameterUse = _inferedTypeInstance.ParameterUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ParameterUse =
								append(_inferedTypeInstance.ParameterUse, any(fieldInstance).(*KnightWhoSayNi))
						}
					}
				}
			}
		}

	case *LancelotCategoryUse:
		// insertion point per field

	case *MapObject:
		// insertion point per field

	case *MapObjectUse:
		// insertion point per field

	case *Repository:
		// insertion point per field
		if fieldName == "ParameterUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Repository) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Repository)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ParameterUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ParameterUse = _inferedTypeInstance.ParameterUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ParameterUse =
								append(_inferedTypeInstance.ParameterUse, any(fieldInstance).(*KnightWhoSayNi))
						}
					}
				}
			}
		}
		if fieldName == "GroupUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Repository) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Repository)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GroupUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GroupUse = _inferedTypeInstance.GroupUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GroupUse =
								append(_inferedTypeInstance.GroupUse, any(fieldInstance).(*GroupUse))
						}
					}
				}
			}
		}

	case *SirRobin:
		// insertion point per field
		if fieldName == "Witches" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SirRobin) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SirRobin)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Witches).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Witches = _inferedTypeInstance.Witches[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Witches =
								append(_inferedTypeInstance.Witches, any(fieldInstance).(*AWitch))
						}
					}
				}
			}
		}
		if fieldName == "Arthurs" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SirRobin) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SirRobin)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Arthurs).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Arthurs = _inferedTypeInstance.Arthurs[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Arthurs =
								append(_inferedTypeInstance.Arthurs, any(fieldInstance).(*KingArthurShape))
						}
					}
				}
			}
		}
		if fieldName == "KnightWhoSayNis" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SirRobin) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SirRobin)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.KnightWhoSayNis).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.KnightWhoSayNis = _inferedTypeInstance.KnightWhoSayNis[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.KnightWhoSayNis =
								append(_inferedTypeInstance.KnightWhoSayNis, any(fieldInstance).(*KnightWhoSayNi))
						}
					}
				}
			}
		}
		if fieldName == "BlackKnightShapes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SirRobin) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SirRobin)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.BlackKnightShapes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.BlackKnightShapes = _inferedTypeInstance.BlackKnightShapes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.BlackKnightShapes =
								append(_inferedTypeInstance.BlackKnightShapes, any(fieldInstance).(*BlackKnightShape))
						}
					}
				}
			}
		}
		if fieldName == "TheNuteShapes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SirRobin) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SirRobin)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TheNuteShapes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TheNuteShapes = _inferedTypeInstance.TheNuteShapes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TheNuteShapes =
								append(_inferedTypeInstance.TheNuteShapes, any(fieldInstance).(*TheNuteShape))
						}
					}
				}
			}
		}

	case *TheBridge:
		// insertion point per field
		if fieldName == "WhatIsYourPreferedColor" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*TheBridge) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*TheBridge)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.WhatIsYourPreferedColor).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.WhatIsYourPreferedColor = _inferedTypeInstance.WhatIsYourPreferedColor[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.WhatIsYourPreferedColor =
								append(_inferedTypeInstance.WhatIsYourPreferedColor, any(fieldInstance).(*WhatIsYourPreferedColor))
						}
					}
				}
			}
		}
		if fieldName == "GroupUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*TheBridge) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*TheBridge)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GroupUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GroupUse = _inferedTypeInstance.GroupUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GroupUse =
								append(_inferedTypeInstance.GroupUse, any(fieldInstance).(*GroupUse))
						}
					}
				}
			}
		}
		if fieldName == "GeoObjectUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*TheBridge) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*TheBridge)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GeoObjectUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GeoObjectUse = _inferedTypeInstance.GeoObjectUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GeoObjectUse =
								append(_inferedTypeInstance.GeoObjectUse, any(fieldInstance).(*GeoObjectUse))
						}
					}
				}
			}
		}
		if fieldName == "MapUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*TheBridge) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*TheBridge)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.MapUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.MapUse = _inferedTypeInstance.MapUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.MapUse =
								append(_inferedTypeInstance.MapUse, any(fieldInstance).(*MapObjectUse))
						}
					}
				}
			}
		}

	case *TheNuteShape:
		// insertion point per field

	case *TheNuteTransition:
		// insertion point per field

	case *User:
		// insertion point per field

	case *UserUse:
		// insertion point per field

	case *WhatIsYourPreferedColor:
		// insertion point per field
		if fieldName == "Diagrams" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*WhatIsYourPreferedColor) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*WhatIsYourPreferedColor)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Diagrams).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Diagrams = _inferedTypeInstance.Diagrams[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Diagrams =
								append(_inferedTypeInstance.Diagrams, any(fieldInstance).(*SirRobin))
						}
					}
				}
			}
		}
		if fieldName == "KingArthurs" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*WhatIsYourPreferedColor) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*WhatIsYourPreferedColor)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.KingArthurs).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.KingArthurs = _inferedTypeInstance.KingArthurs[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.KingArthurs =
								append(_inferedTypeInstance.KingArthurs, any(fieldInstance).(*KingArthur))
						}
					}
				}
			}
		}
		if fieldName == "Nutes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*WhatIsYourPreferedColor) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*WhatIsYourPreferedColor)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Nutes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Nutes = _inferedTypeInstance.Nutes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Nutes =
								append(_inferedTypeInstance.Nutes, any(fieldInstance).(*TheNuteTransition))
						}
					}
				}
			}
		}
		if fieldName == "Galahard" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*WhatIsYourPreferedColor) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*WhatIsYourPreferedColor)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Galahard).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Galahard = _inferedTypeInstance.Galahard[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Galahard =
								append(_inferedTypeInstance.Galahard, any(fieldInstance).(*GalahadThePure))
						}
					}
				}
			}
		}
		if fieldName == "Lancelots" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*WhatIsYourPreferedColor) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*WhatIsYourPreferedColor)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Lancelots).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Lancelots = _inferedTypeInstance.Lancelots[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Lancelots =
								append(_inferedTypeInstance.Lancelots, any(fieldInstance).(*Lancelot))
						}
					}
				}
			}
		}
		if fieldName == "BringYourDeadarameters" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*WhatIsYourPreferedColor) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*WhatIsYourPreferedColor)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.BringYourDeadarameters).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.BringYourDeadarameters = _inferedTypeInstance.BringYourDeadarameters[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.BringYourDeadarameters =
								append(_inferedTypeInstance.BringYourDeadarameters, any(fieldInstance).(*BringYourDead))
						}
					}
				}
			}
		}

	case *Workspace:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AWitch
	// insertion point per field

	// Compute reverse map for named struct BlackKnightShape
	// insertion point per field

	// Compute reverse map for named struct BringYourDead
	// insertion point per field
	clear(stage.BringYourDead_Lancelots_reverseMap)
	stage.BringYourDead_Lancelots_reverseMap = make(map[*Lancelot]*BringYourDead)
	for bringyourdead := range stage.BringYourDeads {
		_ = bringyourdead
		for _, _lancelot := range bringyourdead.Lancelots {
			stage.BringYourDead_Lancelots_reverseMap[_lancelot] = bringyourdead
		}
	}

	// Compute reverse map for named struct Document
	// insertion point per field
	clear(stage.Document_GeoObjectUse_reverseMap)
	stage.Document_GeoObjectUse_reverseMap = make(map[*GeoObjectUse]*Document)
	for document := range stage.Documents {
		_ = document
		for _, _geoobjectuse := range document.GeoObjectUse {
			stage.Document_GeoObjectUse_reverseMap[_geoobjectuse] = document
		}
	}

	// Compute reverse map for named struct DocumentUse
	// insertion point per field

	// Compute reverse map for named struct GalahadThePure
	// insertion point per field

	// Compute reverse map for named struct GeoObject
	// insertion point per field

	// Compute reverse map for named struct GeoObjectUse
	// insertion point per field

	// Compute reverse map for named struct Group
	// insertion point per field
	clear(stage.Group_UserUse_reverseMap)
	stage.Group_UserUse_reverseMap = make(map[*UserUse]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _useruse := range group.UserUse {
			stage.Group_UserUse_reverseMap[_useruse] = group
		}
	}

	// Compute reverse map for named struct GroupUse
	// insertion point per field

	// Compute reverse map for named struct KingArthur
	// insertion point per field

	// Compute reverse map for named struct KingArthurShape
	// insertion point per field

	// Compute reverse map for named struct KnightWhoSayNi
	// insertion point per field

	// Compute reverse map for named struct Lancelot
	// insertion point per field
	clear(stage.Lancelot_GroupUse_reverseMap)
	stage.Lancelot_GroupUse_reverseMap = make(map[*GroupUse]*Lancelot)
	for lancelot := range stage.Lancelots {
		_ = lancelot
		for _, _groupuse := range lancelot.GroupUse {
			stage.Lancelot_GroupUse_reverseMap[_groupuse] = lancelot
		}
	}
	clear(stage.Lancelot_DocumentUse_reverseMap)
	stage.Lancelot_DocumentUse_reverseMap = make(map[*DocumentUse]*Lancelot)
	for lancelot := range stage.Lancelots {
		_ = lancelot
		for _, _documentuse := range lancelot.DocumentUse {
			stage.Lancelot_DocumentUse_reverseMap[_documentuse] = lancelot
		}
	}
	clear(stage.Lancelot_GeoObjectUse_reverseMap)
	stage.Lancelot_GeoObjectUse_reverseMap = make(map[*GeoObjectUse]*Lancelot)
	for lancelot := range stage.Lancelots {
		_ = lancelot
		for _, _geoobjectuse := range lancelot.GeoObjectUse {
			stage.Lancelot_GeoObjectUse_reverseMap[_geoobjectuse] = lancelot
		}
	}

	// Compute reverse map for named struct LancelotAgregation
	// insertion point per field
	clear(stage.LancelotAgregation_ParameterUse_reverseMap)
	stage.LancelotAgregation_ParameterUse_reverseMap = make(map[*KnightWhoSayNi]*LancelotAgregation)
	for lancelotagregation := range stage.LancelotAgregations {
		_ = lancelotagregation
		for _, _knightwhosayni := range lancelotagregation.ParameterUse {
			stage.LancelotAgregation_ParameterUse_reverseMap[_knightwhosayni] = lancelotagregation
		}
	}

	// Compute reverse map for named struct LancelotAgregationUse
	// insertion point per field

	// Compute reverse map for named struct LancelotCategory
	// insertion point per field
	clear(stage.LancelotCategory_ParameterUse_reverseMap)
	stage.LancelotCategory_ParameterUse_reverseMap = make(map[*KnightWhoSayNi]*LancelotCategory)
	for lancelotcategory := range stage.LancelotCategorys {
		_ = lancelotcategory
		for _, _knightwhosayni := range lancelotcategory.ParameterUse {
			stage.LancelotCategory_ParameterUse_reverseMap[_knightwhosayni] = lancelotcategory
		}
	}

	// Compute reverse map for named struct LancelotCategoryUse
	// insertion point per field

	// Compute reverse map for named struct MapObject
	// insertion point per field

	// Compute reverse map for named struct MapObjectUse
	// insertion point per field

	// Compute reverse map for named struct Repository
	// insertion point per field
	clear(stage.Repository_ParameterUse_reverseMap)
	stage.Repository_ParameterUse_reverseMap = make(map[*KnightWhoSayNi]*Repository)
	for repository := range stage.Repositorys {
		_ = repository
		for _, _knightwhosayni := range repository.ParameterUse {
			stage.Repository_ParameterUse_reverseMap[_knightwhosayni] = repository
		}
	}
	clear(stage.Repository_GroupUse_reverseMap)
	stage.Repository_GroupUse_reverseMap = make(map[*GroupUse]*Repository)
	for repository := range stage.Repositorys {
		_ = repository
		for _, _groupuse := range repository.GroupUse {
			stage.Repository_GroupUse_reverseMap[_groupuse] = repository
		}
	}

	// Compute reverse map for named struct SirRobin
	// insertion point per field
	clear(stage.SirRobin_Witches_reverseMap)
	stage.SirRobin_Witches_reverseMap = make(map[*AWitch]*SirRobin)
	for sirrobin := range stage.SirRobins {
		_ = sirrobin
		for _, _awitch := range sirrobin.Witches {
			stage.SirRobin_Witches_reverseMap[_awitch] = sirrobin
		}
	}
	clear(stage.SirRobin_Arthurs_reverseMap)
	stage.SirRobin_Arthurs_reverseMap = make(map[*KingArthurShape]*SirRobin)
	for sirrobin := range stage.SirRobins {
		_ = sirrobin
		for _, _kingarthurshape := range sirrobin.Arthurs {
			stage.SirRobin_Arthurs_reverseMap[_kingarthurshape] = sirrobin
		}
	}
	clear(stage.SirRobin_KnightWhoSayNis_reverseMap)
	stage.SirRobin_KnightWhoSayNis_reverseMap = make(map[*KnightWhoSayNi]*SirRobin)
	for sirrobin := range stage.SirRobins {
		_ = sirrobin
		for _, _knightwhosayni := range sirrobin.KnightWhoSayNis {
			stage.SirRobin_KnightWhoSayNis_reverseMap[_knightwhosayni] = sirrobin
		}
	}
	clear(stage.SirRobin_BlackKnightShapes_reverseMap)
	stage.SirRobin_BlackKnightShapes_reverseMap = make(map[*BlackKnightShape]*SirRobin)
	for sirrobin := range stage.SirRobins {
		_ = sirrobin
		for _, _blackknightshape := range sirrobin.BlackKnightShapes {
			stage.SirRobin_BlackKnightShapes_reverseMap[_blackknightshape] = sirrobin
		}
	}
	clear(stage.SirRobin_TheNuteShapes_reverseMap)
	stage.SirRobin_TheNuteShapes_reverseMap = make(map[*TheNuteShape]*SirRobin)
	for sirrobin := range stage.SirRobins {
		_ = sirrobin
		for _, _thenuteshape := range sirrobin.TheNuteShapes {
			stage.SirRobin_TheNuteShapes_reverseMap[_thenuteshape] = sirrobin
		}
	}

	// Compute reverse map for named struct TheBridge
	// insertion point per field
	clear(stage.TheBridge_WhatIsYourPreferedColor_reverseMap)
	stage.TheBridge_WhatIsYourPreferedColor_reverseMap = make(map[*WhatIsYourPreferedColor]*TheBridge)
	for thebridge := range stage.TheBridges {
		_ = thebridge
		for _, _whatisyourpreferedcolor := range thebridge.WhatIsYourPreferedColor {
			stage.TheBridge_WhatIsYourPreferedColor_reverseMap[_whatisyourpreferedcolor] = thebridge
		}
	}
	clear(stage.TheBridge_GroupUse_reverseMap)
	stage.TheBridge_GroupUse_reverseMap = make(map[*GroupUse]*TheBridge)
	for thebridge := range stage.TheBridges {
		_ = thebridge
		for _, _groupuse := range thebridge.GroupUse {
			stage.TheBridge_GroupUse_reverseMap[_groupuse] = thebridge
		}
	}
	clear(stage.TheBridge_GeoObjectUse_reverseMap)
	stage.TheBridge_GeoObjectUse_reverseMap = make(map[*GeoObjectUse]*TheBridge)
	for thebridge := range stage.TheBridges {
		_ = thebridge
		for _, _geoobjectuse := range thebridge.GeoObjectUse {
			stage.TheBridge_GeoObjectUse_reverseMap[_geoobjectuse] = thebridge
		}
	}
	clear(stage.TheBridge_MapUse_reverseMap)
	stage.TheBridge_MapUse_reverseMap = make(map[*MapObjectUse]*TheBridge)
	for thebridge := range stage.TheBridges {
		_ = thebridge
		for _, _mapobjectuse := range thebridge.MapUse {
			stage.TheBridge_MapUse_reverseMap[_mapobjectuse] = thebridge
		}
	}

	// Compute reverse map for named struct TheNuteShape
	// insertion point per field

	// Compute reverse map for named struct TheNuteTransition
	// insertion point per field

	// Compute reverse map for named struct User
	// insertion point per field

	// Compute reverse map for named struct UserUse
	// insertion point per field

	// Compute reverse map for named struct WhatIsYourPreferedColor
	// insertion point per field
	clear(stage.WhatIsYourPreferedColor_Diagrams_reverseMap)
	stage.WhatIsYourPreferedColor_Diagrams_reverseMap = make(map[*SirRobin]*WhatIsYourPreferedColor)
	for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
		_ = whatisyourpreferedcolor
		for _, _sirrobin := range whatisyourpreferedcolor.Diagrams {
			stage.WhatIsYourPreferedColor_Diagrams_reverseMap[_sirrobin] = whatisyourpreferedcolor
		}
	}
	clear(stage.WhatIsYourPreferedColor_KingArthurs_reverseMap)
	stage.WhatIsYourPreferedColor_KingArthurs_reverseMap = make(map[*KingArthur]*WhatIsYourPreferedColor)
	for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
		_ = whatisyourpreferedcolor
		for _, _kingarthur := range whatisyourpreferedcolor.KingArthurs {
			stage.WhatIsYourPreferedColor_KingArthurs_reverseMap[_kingarthur] = whatisyourpreferedcolor
		}
	}
	clear(stage.WhatIsYourPreferedColor_Nutes_reverseMap)
	stage.WhatIsYourPreferedColor_Nutes_reverseMap = make(map[*TheNuteTransition]*WhatIsYourPreferedColor)
	for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
		_ = whatisyourpreferedcolor
		for _, _thenutetransition := range whatisyourpreferedcolor.Nutes {
			stage.WhatIsYourPreferedColor_Nutes_reverseMap[_thenutetransition] = whatisyourpreferedcolor
		}
	}
	clear(stage.WhatIsYourPreferedColor_Galahard_reverseMap)
	stage.WhatIsYourPreferedColor_Galahard_reverseMap = make(map[*GalahadThePure]*WhatIsYourPreferedColor)
	for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
		_ = whatisyourpreferedcolor
		for _, _galahadthepure := range whatisyourpreferedcolor.Galahard {
			stage.WhatIsYourPreferedColor_Galahard_reverseMap[_galahadthepure] = whatisyourpreferedcolor
		}
	}
	clear(stage.WhatIsYourPreferedColor_Lancelots_reverseMap)
	stage.WhatIsYourPreferedColor_Lancelots_reverseMap = make(map[*Lancelot]*WhatIsYourPreferedColor)
	for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
		_ = whatisyourpreferedcolor
		for _, _lancelot := range whatisyourpreferedcolor.Lancelots {
			stage.WhatIsYourPreferedColor_Lancelots_reverseMap[_lancelot] = whatisyourpreferedcolor
		}
	}
	clear(stage.WhatIsYourPreferedColor_BringYourDeadarameters_reverseMap)
	stage.WhatIsYourPreferedColor_BringYourDeadarameters_reverseMap = make(map[*BringYourDead]*WhatIsYourPreferedColor)
	for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
		_ = whatisyourpreferedcolor
		for _, _bringyourdead := range whatisyourpreferedcolor.BringYourDeadarameters {
			stage.WhatIsYourPreferedColor_BringYourDeadarameters_reverseMap[_bringyourdead] = whatisyourpreferedcolor
		}
	}

	// Compute reverse map for named struct Workspace
	// insertion point per field

}
