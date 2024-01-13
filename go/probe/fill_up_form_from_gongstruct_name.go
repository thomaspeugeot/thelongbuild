// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "AWitch":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "AWitch Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AWitchFormCallback(
			nil,
			probe,
			formGroup,
		)
		awitch := new(models.AWitch)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(awitch, formGroup, probe)
	case "BlackKnightShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "BlackKnightShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BlackKnightShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		blackknightshape := new(models.BlackKnightShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(blackknightshape, formGroup, probe)
	case "BringYourDead":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "BringYourDead Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BringYourDeadFormCallback(
			nil,
			probe,
			formGroup,
		)
		bringyourdead := new(models.BringYourDead)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bringyourdead, formGroup, probe)
	case "Document":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Document Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		document := new(models.Document)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(document, formGroup, probe)
	case "DocumentUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DocumentUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		documentuse := new(models.DocumentUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(documentuse, formGroup, probe)
	case "GalahadThePure":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GalahadThePure Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GalahadThePureFormCallback(
			nil,
			probe,
			formGroup,
		)
		galahadthepure := new(models.GalahadThePure)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(galahadthepure, formGroup, probe)
	case "GeoObject":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GeoObject Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GeoObjectFormCallback(
			nil,
			probe,
			formGroup,
		)
		geoobject := new(models.GeoObject)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(geoobject, formGroup, probe)
	case "GeoObjectUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GeoObjectUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GeoObjectUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		geoobjectuse := new(models.GeoObjectUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(geoobjectuse, formGroup, probe)
	case "Group":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		group := new(models.Group)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group, formGroup, probe)
	case "GroupUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "GroupUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		groupuse := new(models.GroupUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(groupuse, formGroup, probe)
	case "KingArthur":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "KingArthur Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KingArthurFormCallback(
			nil,
			probe,
			formGroup,
		)
		kingarthur := new(models.KingArthur)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(kingarthur, formGroup, probe)
	case "KingArthurShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "KingArthurShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KingArthurShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		kingarthurshape := new(models.KingArthurShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(kingarthurshape, formGroup, probe)
	case "KnightWhoSayNi":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "KnightWhoSayNi Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KnightWhoSayNiFormCallback(
			nil,
			probe,
			formGroup,
		)
		knightwhosayni := new(models.KnightWhoSayNi)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(knightwhosayni, formGroup, probe)
	case "Lancelot":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Lancelot Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotFormCallback(
			nil,
			probe,
			formGroup,
		)
		lancelot := new(models.Lancelot)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lancelot, formGroup, probe)
	case "LancelotAgregation":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "LancelotAgregation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotAgregationFormCallback(
			nil,
			probe,
			formGroup,
		)
		lancelotagregation := new(models.LancelotAgregation)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lancelotagregation, formGroup, probe)
	case "LancelotAgregationUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "LancelotAgregationUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotAgregationUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		lancelotagregationuse := new(models.LancelotAgregationUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lancelotagregationuse, formGroup, probe)
	case "LancelotCategory":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "LancelotCategory Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotCategoryFormCallback(
			nil,
			probe,
			formGroup,
		)
		lancelotcategory := new(models.LancelotCategory)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lancelotcategory, formGroup, probe)
	case "LancelotCategoryUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "LancelotCategoryUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LancelotCategoryUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		lancelotcategoryuse := new(models.LancelotCategoryUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lancelotcategoryuse, formGroup, probe)
	case "MapObject":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "MapObject Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapObjectFormCallback(
			nil,
			probe,
			formGroup,
		)
		mapobject := new(models.MapObject)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mapobject, formGroup, probe)
	case "MapObjectUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "MapObjectUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapObjectUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		mapobjectuse := new(models.MapObjectUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mapobjectuse, formGroup, probe)
	case "Repository":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Repository Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RepositoryFormCallback(
			nil,
			probe,
			formGroup,
		)
		repository := new(models.Repository)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(repository, formGroup, probe)
	case "SirRobin":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SirRobin Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SirRobinFormCallback(
			nil,
			probe,
			formGroup,
		)
		sirrobin := new(models.SirRobin)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sirrobin, formGroup, probe)
	case "TheBridge":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "TheBridge Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TheBridgeFormCallback(
			nil,
			probe,
			formGroup,
		)
		thebridge := new(models.TheBridge)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(thebridge, formGroup, probe)
	case "TheNuteShape":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "TheNuteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TheNuteShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		thenuteshape := new(models.TheNuteShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(thenuteshape, formGroup, probe)
	case "TheNuteTransition":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "TheNuteTransition Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TheNuteTransitionFormCallback(
			nil,
			probe,
			formGroup,
		)
		thenutetransition := new(models.TheNuteTransition)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(thenutetransition, formGroup, probe)
	case "User":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "User Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserFormCallback(
			nil,
			probe,
			formGroup,
		)
		user := new(models.User)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(user, formGroup, probe)
	case "UserUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "UserUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UserUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		useruse := new(models.UserUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(useruse, formGroup, probe)
	case "WhatIsYourPreferedColor":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "WhatIsYourPreferedColor Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WhatIsYourPreferedColorFormCallback(
			nil,
			probe,
			formGroup,
		)
		whatisyourpreferedcolor := new(models.WhatIsYourPreferedColor)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(whatisyourpreferedcolor, formGroup, probe)
	case "Workspace":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Workspace Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WorkspaceFormCallback(
			nil,
			probe,
			formGroup,
		)
		workspace := new(models.Workspace)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(workspace, formGroup, probe)
	}
	formStage.Commit()
}
