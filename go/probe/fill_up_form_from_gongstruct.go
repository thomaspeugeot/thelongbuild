// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

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
		formGroup.OnSave = __gong__New__AWitchFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BlackKnightShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "BlackKnightShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BlackKnightShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.BringYourDead:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "BringYourDead Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BringYourDeadFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Document:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Document Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DocumentUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DocumentUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GalahadThePure:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GalahadThePure Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GalahadThePureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GeoObject:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GeoObject Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GeoObjectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GeoObjectUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GeoObjectUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GeoObjectUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GroupUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GroupUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.KingArthur:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "KingArthur Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KingArthurFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.KingArthurShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "KingArthurShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KingArthurShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.KnightWhoSayNi:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "KnightWhoSayNi Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KnightWhoSayNiFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lancelot:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Lancelot Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LancelotAgregation:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LancelotAgregation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotAgregationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LancelotAgregationUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LancelotAgregationUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotAgregationUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LancelotCategory:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LancelotCategory Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotCategoryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LancelotCategoryUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LancelotCategoryUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotCategoryUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MapObject:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MapObject Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapObjectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MapObjectUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MapObjectUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapObjectUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Repository:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Repository Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RepositoryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SirRobin:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SirRobin Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SirRobinFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TheBridge:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TheBridge Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TheBridgeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TheNuteShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TheNuteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TheNuteShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TheNuteTransition:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TheNuteTransition Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TheNuteTransitionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.User:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "User Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.UserUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "UserUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.WhatIsYourPreferedColor:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "WhatIsYourPreferedColor Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WhatIsYourPreferedColorFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Workspace:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Workspace Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WorkspaceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
