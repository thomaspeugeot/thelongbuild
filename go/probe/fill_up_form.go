// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
	"github.com/thomaspeugeot/thelongbuild/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.AWitch:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("EvolutionDirection", instanceWithInferedType.EvolutionDirection, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillColor", instanceWithInferedType.FillColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "Witches"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SirRobin),
					"Witches",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SirRobin, *models.AWitch](
					nil,
					"Witches",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.BlackKnightShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("BringYourDead", instanceWithInferedType.BringYourDead, formGroup, probe)
		EnumTypeStringToForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillColor", instanceWithInferedType.FillColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "BlackKnightShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SirRobin),
					"BlackKnightShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SirRobin, *models.BlackKnightShape](
					nil,
					"BlackKnightShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.BringYourDead:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tag", instanceWithInferedType.Tag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 500, true, 500)
		AssociationSliceToForm("Lancelots", instanceWithInferedType, &instanceWithInferedType.Lancelots, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "BringYourDeadarameters"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.WhatIsYourPreferedColor),
					"BringYourDeadarameters",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.WhatIsYourPreferedColor, *models.BringYourDead](
					nil,
					"BringYourDeadarameters",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Document:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GeoObjectUse", instanceWithInferedType, &instanceWithInferedType.GeoObjectUse, formGroup, probe)

	case *models.DocumentUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Document", instanceWithInferedType.Document, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lancelot"
			rf.Fieldname = "DocumentUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Lancelot),
					"DocumentUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Lancelot, *models.DocumentUse](
					nil,
					"DocumentUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GalahadThePure:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "Galahard"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.WhatIsYourPreferedColor),
					"Galahard",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.WhatIsYourPreferedColor, *models.GalahadThePure](
					nil,
					"Galahard",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GeoObject:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.GeoObjectUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("GeoObject", instanceWithInferedType.GeoObject, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Document"
			rf.Fieldname = "GeoObjectUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Document),
					"GeoObjectUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Document, *models.GeoObjectUse](
					nil,
					"GeoObjectUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lancelot"
			rf.Fieldname = "GeoObjectUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Lancelot),
					"GeoObjectUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Lancelot, *models.GeoObjectUse](
					nil,
					"GeoObjectUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TheBridge"
			rf.Fieldname = "GeoObjectUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TheBridge),
					"GeoObjectUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.TheBridge, *models.GeoObjectUse](
					nil,
					"GeoObjectUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("UserUse", instanceWithInferedType, &instanceWithInferedType.UserUse, formGroup, probe)

	case *models.GroupUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Group", instanceWithInferedType.Group, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lancelot"
			rf.Fieldname = "GroupUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Lancelot),
					"GroupUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Lancelot, *models.GroupUse](
					nil,
					"GroupUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Repository"
			rf.Fieldname = "GroupUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Repository),
					"GroupUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Repository, *models.GroupUse](
					nil,
					"GroupUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TheBridge"
			rf.Fieldname = "GroupUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TheBridge),
					"GroupUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.TheBridge, *models.GroupUse](
					nil,
					"GroupUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.KingArthur:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("IsWithProbaility", instanceWithInferedType.IsWithProbaility, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Probability", instanceWithInferedType.Probability, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "KingArthurs"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.WhatIsYourPreferedColor),
					"KingArthurs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.WhatIsYourPreferedColor, *models.KingArthur](
					nil,
					"KingArthurs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.KingArthurShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ActorState", instanceWithInferedType.ActorState, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillColor", instanceWithInferedType.FillColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "Arthurs"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SirRobin),
					"Arthurs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SirRobin, *models.KingArthurShape](
					nil,
					"Arthurs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.KnightWhoSayNi:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Parameter", instanceWithInferedType.Parameter, formGroup, probe)
		EnumTypeStringToForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillColor", instanceWithInferedType.FillColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RX", instanceWithInferedType.RX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LancelotAgregation"
			rf.Fieldname = "ParameterUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.LancelotAgregation),
					"ParameterUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.LancelotAgregation, *models.KnightWhoSayNi](
					nil,
					"ParameterUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LancelotCategory"
			rf.Fieldname = "ParameterUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.LancelotCategory),
					"ParameterUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.LancelotCategory, *models.KnightWhoSayNi](
					nil,
					"ParameterUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Repository"
			rf.Fieldname = "ParameterUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Repository),
					"ParameterUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Repository, *models.KnightWhoSayNi](
					nil,
					"ParameterUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "KnightWhoSayNis"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SirRobin),
					"KnightWhoSayNis",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SirRobin, *models.KnightWhoSayNi](
					nil,
					"KnightWhoSayNis",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Lancelot:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tag", instanceWithInferedType.Tag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 500, true, 500)
		AssociationSliceToForm("GroupUse", instanceWithInferedType, &instanceWithInferedType.GroupUse, formGroup, probe)
		AssociationSliceToForm("DocumentUse", instanceWithInferedType, &instanceWithInferedType.DocumentUse, formGroup, probe)
		AssociationSliceToForm("GeoObjectUse", instanceWithInferedType, &instanceWithInferedType.GeoObjectUse, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BringYourDead"
			rf.Fieldname = "Lancelots"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.BringYourDead),
					"Lancelots",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.BringYourDead, *models.Lancelot](
					nil,
					"Lancelots",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "Lancelots"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.WhatIsYourPreferedColor),
					"Lancelots",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.WhatIsYourPreferedColor, *models.Lancelot](
					nil,
					"Lancelots",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.LancelotAgregation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ParameterUse", instanceWithInferedType, &instanceWithInferedType.ParameterUse, formGroup, probe)

	case *models.LancelotAgregationUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ParameterAgregation", instanceWithInferedType.ParameterAgregation, formGroup, probe)

	case *models.LancelotCategory:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ParameterUse", instanceWithInferedType, &instanceWithInferedType.ParameterUse, formGroup, probe)

	case *models.LancelotCategoryUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ParameterCategory", instanceWithInferedType.ParameterCategory, formGroup, probe)

	case *models.MapObject:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.MapObjectUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Map", instanceWithInferedType.Map, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TheBridge"
			rf.Fieldname = "MapUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TheBridge),
					"MapUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.TheBridge, *models.MapObjectUse](
					nil,
					"MapUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Repository:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ParameterUse", instanceWithInferedType, &instanceWithInferedType.ParameterUse, formGroup, probe)
		AssociationSliceToForm("GroupUse", instanceWithInferedType, &instanceWithInferedType.GroupUse, formGroup, probe)

	case *models.SirRobin:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("Witches", instanceWithInferedType, &instanceWithInferedType.Witches, formGroup, probe)
		AssociationSliceToForm("Arthurs", instanceWithInferedType, &instanceWithInferedType.Arthurs, formGroup, probe)
		AssociationSliceToForm("KnightWhoSayNis", instanceWithInferedType, &instanceWithInferedType.KnightWhoSayNis, formGroup, probe)
		AssociationSliceToForm("BlackKnightShapes", instanceWithInferedType, &instanceWithInferedType.BlackKnightShapes, formGroup, probe)
		AssociationSliceToForm("TheNuteShapes", instanceWithInferedType, &instanceWithInferedType.TheNuteShapes, formGroup, probe)
		BasicFieldtoForm("AxisOrign_X", instanceWithInferedType.AxisOrign_X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AxisOrign_Y", instanceWithInferedType.AxisOrign_Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("VerticalAxis_Top_Y", instanceWithInferedType.VerticalAxis_Top_Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("VerticalAxis_Bottom_Y", instanceWithInferedType.VerticalAxis_Bottom_Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("VerticalAxis_StrokeWidth", instanceWithInferedType.VerticalAxis_StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HorizontalAxis_Right_X", instanceWithInferedType.HorizontalAxis_Right_X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Start", instanceWithInferedType.Start, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End", instanceWithInferedType.End, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "Diagrams"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.WhatIsYourPreferedColor),
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.WhatIsYourPreferedColor, *models.SirRobin](
					nil,
					"Diagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.TheBridge:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 500, true, 500)
		AssociationSliceToForm("WhatIsYourPreferedColor", instanceWithInferedType, &instanceWithInferedType.WhatIsYourPreferedColor, formGroup, probe)
		AssociationSliceToForm("GroupUse", instanceWithInferedType, &instanceWithInferedType.GroupUse, formGroup, probe)
		AssociationSliceToForm("GeoObjectUse", instanceWithInferedType, &instanceWithInferedType.GeoObjectUse, formGroup, probe)
		AssociationSliceToForm("MapUse", instanceWithInferedType, &instanceWithInferedType.MapUse, formGroup, probe)
		BasicFieldtoForm("IsNodeExpanded", instanceWithInferedType.IsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.TheNuteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ActorStateTransition", instanceWithInferedType.ActorStateTransition, formGroup, probe)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "TheNuteShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SirRobin),
					"TheNuteShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SirRobin, *models.TheNuteShape](
					nil,
					"TheNuteShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.TheNuteTransition:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("StartState", instanceWithInferedType.StartState, formGroup, probe)
		AssociationFieldToForm("EndState", instanceWithInferedType.EndState, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "Nutes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.WhatIsYourPreferedColor),
					"Nutes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.WhatIsYourPreferedColor, *models.TheNuteTransition](
					nil,
					"Nutes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.User:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.UserUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("User", instanceWithInferedType.User, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "UserUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"UserUse",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.UserUse](
					nil,
					"UserUse",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.WhatIsYourPreferedColor:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		BasicFieldtoForm("DiagramsNodeFolded", instanceWithInferedType.DiagramsNodeFolded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("KingArthurs", instanceWithInferedType, &instanceWithInferedType.KingArthurs, formGroup, probe)
		BasicFieldtoForm("KingArthurNodeFolded", instanceWithInferedType.KingArthurNodeFolded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Nutes", instanceWithInferedType, &instanceWithInferedType.Nutes, formGroup, probe)
		BasicFieldtoForm("RRRR", instanceWithInferedType.RRRR, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Galahard", instanceWithInferedType, &instanceWithInferedType.Galahard, formGroup, probe)
		BasicFieldtoForm("IIUU", instanceWithInferedType.IIUU, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Lancelots", instanceWithInferedType, &instanceWithInferedType.Lancelots, formGroup, probe)
		BasicFieldtoForm("LancelotsodeFolded", instanceWithInferedType.LancelotsodeFolded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("BringYourDeadarameters", instanceWithInferedType, &instanceWithInferedType.BringYourDeadarameters, formGroup, probe)
		BasicFieldtoForm("RRRRT", instanceWithInferedType.RRRRT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsNodeExpanded", instanceWithInferedType.IsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TheBridge"
			rf.Fieldname = "WhatIsYourPreferedColor"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TheBridge),
					"WhatIsYourPreferedColor",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.TheBridge, *models.WhatIsYourPreferedColor](
					nil,
					"WhatIsYourPreferedColor",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Workspace:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SelectedDiagram", instanceWithInferedType.SelectedDiagram, formGroup, probe)
		AssociationFieldToForm("A", instanceWithInferedType.A, formGroup, probe)
		AssociationFieldToForm("B", instanceWithInferedType.B, formGroup, probe)
		AssociationFieldToForm("C", instanceWithInferedType.C, formGroup, probe)
		AssociationFieldToForm("D", instanceWithInferedType.D, formGroup, probe)
		AssociationFieldToForm("E", instanceWithInferedType.E, formGroup, probe)

	default:
		_ = instanceWithInferedType
	}
}
