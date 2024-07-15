// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	log.Println("")

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AWitch:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AWitch Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.BlackKnightShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "BlackKnightShape Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.BringYourDead:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "BringYourDead Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.Document:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Document Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.DocumentUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DocumentUse Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.GalahadThePure:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GalahadThePure Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.GeoObject:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GeoObject Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.GeoObjectUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GeoObjectUse Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.Group:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Group Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.GroupUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GroupUse Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.KingArthur:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "KingArthur Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.KingArthurShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "KingArthurShape Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.KnightWhoSayNi:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "KnightWhoSayNi Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.Lancelot:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Lancelot Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.LancelotAgregation:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LancelotAgregation Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.LancelotAgregationUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LancelotAgregationUse Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.LancelotCategory:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LancelotCategory Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.LancelotCategoryUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LancelotCategoryUse Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.MapObject:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MapObject Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.MapObjectUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MapObjectUse Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.Repository:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Repository Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.SirRobin:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SirRobin Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.TheBridge:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TheBridge Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.TheNuteShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TheNuteShape Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.TheNuteTransition:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TheNuteTransition Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.User:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "User Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.UserUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "UserUse Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.WhatIsYourPreferedColor:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "WhatIsYourPreferedColor Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	case *models.Workspace:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Workspace Form",
		}).Stage(formStage)

		formGroup.HasSuppressButton = true

	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
