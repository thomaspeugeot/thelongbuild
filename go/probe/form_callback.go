// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
	"github.com/thomaspeugeot/thelongbuild/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__AWitchFormCallback(
	awitch *models.AWitch,
	probe *Probe,
	formGroup *table.FormGroup,
) (awitchFormCallback *AWitchFormCallback) {
	awitchFormCallback = new(AWitchFormCallback)
	awitchFormCallback.probe = probe
	awitchFormCallback.awitch = awitch
	awitchFormCallback.formGroup = formGroup

	awitchFormCallback.CreationMode = (awitch == nil)

	return
}

type AWitchFormCallback struct {
	awitch *models.AWitch

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (awitchFormCallback *AWitchFormCallback) OnSave() {

	log.Println("AWitchFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	awitchFormCallback.probe.formStage.Checkout()

	if awitchFormCallback.awitch == nil {
		awitchFormCallback.awitch = new(models.AWitch).Stage(awitchFormCallback.probe.stageOfInterest)
	}
	awitch_ := awitchFormCallback.awitch
	_ = awitch_

	for _, formDiv := range awitchFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(awitch_.Name), formDiv)
		case "EvolutionDirection":
			FormDivSelectFieldToField(&(awitch_.EvolutionDirection), awitchFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(awitch_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(awitch_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(awitch_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(awitch_.Height), formDiv)
		case "FillColor":
			FormDivBasicFieldToField(&(awitch_.FillColor), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(awitch_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(awitch_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(awitch_.StrokeWidth), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(awitch_.RX), formDiv)
		case "SirRobin:Witches":
			// we need to retrieve the field owner before the change
			var pastSirRobinOwner *models.SirRobin
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "Witches"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				awitchFormCallback.probe.stageOfInterest,
				awitchFormCallback.probe.backRepoOfInterest,
				awitch_,
				&rf)

			if reverseFieldOwner != nil {
				pastSirRobinOwner = reverseFieldOwner.(*models.SirRobin)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSirRobinOwner != nil {
					idx := slices.Index(pastSirRobinOwner.Witches, awitch_)
					pastSirRobinOwner.Witches = slices.Delete(pastSirRobinOwner.Witches, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sirrobin := range *models.GetGongstructInstancesSet[models.SirRobin](awitchFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sirrobin.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSirRobinOwner := _sirrobin // we have a match
						if pastSirRobinOwner != nil {
							if newSirRobinOwner != pastSirRobinOwner {
								idx := slices.Index(pastSirRobinOwner.Witches, awitch_)
								pastSirRobinOwner.Witches = slices.Delete(pastSirRobinOwner.Witches, idx, idx+1)
								newSirRobinOwner.Witches = append(newSirRobinOwner.Witches, awitch_)
							}
						} else {
							newSirRobinOwner.Witches = append(newSirRobinOwner.Witches, awitch_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if awitchFormCallback.formGroup.HasSuppressButtonBeenPressed {
		awitch_.Unstage(awitchFormCallback.probe.stageOfInterest)
	}

	awitchFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AWitch](
		awitchFormCallback.probe,
	)
	awitchFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if awitchFormCallback.CreationMode || awitchFormCallback.formGroup.HasSuppressButtonBeenPressed {
		awitchFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(awitchFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AWitchFormCallback(
			nil,
			awitchFormCallback.probe,
			newFormGroup,
		)
		awitch := new(models.AWitch)
		FillUpForm(awitch, newFormGroup, awitchFormCallback.probe)
		awitchFormCallback.probe.formStage.Commit()
	}

	fillUpTree(awitchFormCallback.probe)
}
func __gong__New__BlackKnightShapeFormCallback(
	blackknightshape *models.BlackKnightShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (blackknightshapeFormCallback *BlackKnightShapeFormCallback) {
	blackknightshapeFormCallback = new(BlackKnightShapeFormCallback)
	blackknightshapeFormCallback.probe = probe
	blackknightshapeFormCallback.blackknightshape = blackknightshape
	blackknightshapeFormCallback.formGroup = formGroup

	blackknightshapeFormCallback.CreationMode = (blackknightshape == nil)

	return
}

type BlackKnightShapeFormCallback struct {
	blackknightshape *models.BlackKnightShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (blackknightshapeFormCallback *BlackKnightShapeFormCallback) OnSave() {

	log.Println("BlackKnightShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	blackknightshapeFormCallback.probe.formStage.Checkout()

	if blackknightshapeFormCallback.blackknightshape == nil {
		blackknightshapeFormCallback.blackknightshape = new(models.BlackKnightShape).Stage(blackknightshapeFormCallback.probe.stageOfInterest)
	}
	blackknightshape_ := blackknightshapeFormCallback.blackknightshape
	_ = blackknightshape_

	for _, formDiv := range blackknightshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(blackknightshape_.Name), formDiv)
		case "BringYourDead":
			FormDivSelectFieldToField(&(blackknightshape_.BringYourDead), blackknightshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(blackknightshape_.Direction), formDiv)
		case "X":
			FormDivBasicFieldToField(&(blackknightshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(blackknightshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(blackknightshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(blackknightshape_.Height), formDiv)
		case "FillColor":
			FormDivBasicFieldToField(&(blackknightshape_.FillColor), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(blackknightshape_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(blackknightshape_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(blackknightshape_.StrokeWidth), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(blackknightshape_.RX), formDiv)
		case "SirRobin:BlackKnightShapes":
			// we need to retrieve the field owner before the change
			var pastSirRobinOwner *models.SirRobin
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "BlackKnightShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				blackknightshapeFormCallback.probe.stageOfInterest,
				blackknightshapeFormCallback.probe.backRepoOfInterest,
				blackknightshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastSirRobinOwner = reverseFieldOwner.(*models.SirRobin)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSirRobinOwner != nil {
					idx := slices.Index(pastSirRobinOwner.BlackKnightShapes, blackknightshape_)
					pastSirRobinOwner.BlackKnightShapes = slices.Delete(pastSirRobinOwner.BlackKnightShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sirrobin := range *models.GetGongstructInstancesSet[models.SirRobin](blackknightshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sirrobin.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSirRobinOwner := _sirrobin // we have a match
						if pastSirRobinOwner != nil {
							if newSirRobinOwner != pastSirRobinOwner {
								idx := slices.Index(pastSirRobinOwner.BlackKnightShapes, blackknightshape_)
								pastSirRobinOwner.BlackKnightShapes = slices.Delete(pastSirRobinOwner.BlackKnightShapes, idx, idx+1)
								newSirRobinOwner.BlackKnightShapes = append(newSirRobinOwner.BlackKnightShapes, blackknightshape_)
							}
						} else {
							newSirRobinOwner.BlackKnightShapes = append(newSirRobinOwner.BlackKnightShapes, blackknightshape_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if blackknightshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		blackknightshape_.Unstage(blackknightshapeFormCallback.probe.stageOfInterest)
	}

	blackknightshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.BlackKnightShape](
		blackknightshapeFormCallback.probe,
	)
	blackknightshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if blackknightshapeFormCallback.CreationMode || blackknightshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		blackknightshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(blackknightshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BlackKnightShapeFormCallback(
			nil,
			blackknightshapeFormCallback.probe,
			newFormGroup,
		)
		blackknightshape := new(models.BlackKnightShape)
		FillUpForm(blackknightshape, newFormGroup, blackknightshapeFormCallback.probe)
		blackknightshapeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(blackknightshapeFormCallback.probe)
}
func __gong__New__BringYourDeadFormCallback(
	bringyourdead *models.BringYourDead,
	probe *Probe,
	formGroup *table.FormGroup,
) (bringyourdeadFormCallback *BringYourDeadFormCallback) {
	bringyourdeadFormCallback = new(BringYourDeadFormCallback)
	bringyourdeadFormCallback.probe = probe
	bringyourdeadFormCallback.bringyourdead = bringyourdead
	bringyourdeadFormCallback.formGroup = formGroup

	bringyourdeadFormCallback.CreationMode = (bringyourdead == nil)

	return
}

type BringYourDeadFormCallback struct {
	bringyourdead *models.BringYourDead

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bringyourdeadFormCallback *BringYourDeadFormCallback) OnSave() {

	log.Println("BringYourDeadFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bringyourdeadFormCallback.probe.formStage.Checkout()

	if bringyourdeadFormCallback.bringyourdead == nil {
		bringyourdeadFormCallback.bringyourdead = new(models.BringYourDead).Stage(bringyourdeadFormCallback.probe.stageOfInterest)
	}
	bringyourdead_ := bringyourdeadFormCallback.bringyourdead
	_ = bringyourdead_

	for _, formDiv := range bringyourdeadFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bringyourdead_.Name), formDiv)
		case "Tag":
			FormDivBasicFieldToField(&(bringyourdead_.Tag), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(bringyourdead_.Description), formDiv)
		case "WhatIsYourPreferedColor:BringYourDeadarameters":
			// we need to retrieve the field owner before the change
			var pastWhatIsYourPreferedColorOwner *models.WhatIsYourPreferedColor
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "BringYourDeadarameters"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				bringyourdeadFormCallback.probe.stageOfInterest,
				bringyourdeadFormCallback.probe.backRepoOfInterest,
				bringyourdead_,
				&rf)

			if reverseFieldOwner != nil {
				pastWhatIsYourPreferedColorOwner = reverseFieldOwner.(*models.WhatIsYourPreferedColor)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastWhatIsYourPreferedColorOwner != nil {
					idx := slices.Index(pastWhatIsYourPreferedColorOwner.BringYourDeadarameters, bringyourdead_)
					pastWhatIsYourPreferedColorOwner.BringYourDeadarameters = slices.Delete(pastWhatIsYourPreferedColorOwner.BringYourDeadarameters, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _whatisyourpreferedcolor := range *models.GetGongstructInstancesSet[models.WhatIsYourPreferedColor](bringyourdeadFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _whatisyourpreferedcolor.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newWhatIsYourPreferedColorOwner := _whatisyourpreferedcolor // we have a match
						if pastWhatIsYourPreferedColorOwner != nil {
							if newWhatIsYourPreferedColorOwner != pastWhatIsYourPreferedColorOwner {
								idx := slices.Index(pastWhatIsYourPreferedColorOwner.BringYourDeadarameters, bringyourdead_)
								pastWhatIsYourPreferedColorOwner.BringYourDeadarameters = slices.Delete(pastWhatIsYourPreferedColorOwner.BringYourDeadarameters, idx, idx+1)
								newWhatIsYourPreferedColorOwner.BringYourDeadarameters = append(newWhatIsYourPreferedColorOwner.BringYourDeadarameters, bringyourdead_)
							}
						} else {
							newWhatIsYourPreferedColorOwner.BringYourDeadarameters = append(newWhatIsYourPreferedColorOwner.BringYourDeadarameters, bringyourdead_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bringyourdeadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bringyourdead_.Unstage(bringyourdeadFormCallback.probe.stageOfInterest)
	}

	bringyourdeadFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.BringYourDead](
		bringyourdeadFormCallback.probe,
	)
	bringyourdeadFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bringyourdeadFormCallback.CreationMode || bringyourdeadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bringyourdeadFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bringyourdeadFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BringYourDeadFormCallback(
			nil,
			bringyourdeadFormCallback.probe,
			newFormGroup,
		)
		bringyourdead := new(models.BringYourDead)
		FillUpForm(bringyourdead, newFormGroup, bringyourdeadFormCallback.probe)
		bringyourdeadFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bringyourdeadFormCallback.probe)
}
func __gong__New__DocumentFormCallback(
	document *models.Document,
	probe *Probe,
	formGroup *table.FormGroup,
) (documentFormCallback *DocumentFormCallback) {
	documentFormCallback = new(DocumentFormCallback)
	documentFormCallback.probe = probe
	documentFormCallback.document = document
	documentFormCallback.formGroup = formGroup

	documentFormCallback.CreationMode = (document == nil)

	return
}

type DocumentFormCallback struct {
	document *models.Document

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (documentFormCallback *DocumentFormCallback) OnSave() {

	log.Println("DocumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentFormCallback.probe.formStage.Checkout()

	if documentFormCallback.document == nil {
		documentFormCallback.document = new(models.Document).Stage(documentFormCallback.probe.stageOfInterest)
	}
	document_ := documentFormCallback.document
	_ = document_

	for _, formDiv := range documentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(document_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if documentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		document_.Unstage(documentFormCallback.probe.stageOfInterest)
	}

	documentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Document](
		documentFormCallback.probe,
	)
	documentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if documentFormCallback.CreationMode || documentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(documentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocumentFormCallback(
			nil,
			documentFormCallback.probe,
			newFormGroup,
		)
		document := new(models.Document)
		FillUpForm(document, newFormGroup, documentFormCallback.probe)
		documentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(documentFormCallback.probe)
}
func __gong__New__DocumentUseFormCallback(
	documentuse *models.DocumentUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (documentuseFormCallback *DocumentUseFormCallback) {
	documentuseFormCallback = new(DocumentUseFormCallback)
	documentuseFormCallback.probe = probe
	documentuseFormCallback.documentuse = documentuse
	documentuseFormCallback.formGroup = formGroup

	documentuseFormCallback.CreationMode = (documentuse == nil)

	return
}

type DocumentUseFormCallback struct {
	documentuse *models.DocumentUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (documentuseFormCallback *DocumentUseFormCallback) OnSave() {

	log.Println("DocumentUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentuseFormCallback.probe.formStage.Checkout()

	if documentuseFormCallback.documentuse == nil {
		documentuseFormCallback.documentuse = new(models.DocumentUse).Stage(documentuseFormCallback.probe.stageOfInterest)
	}
	documentuse_ := documentuseFormCallback.documentuse
	_ = documentuse_

	for _, formDiv := range documentuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(documentuse_.Name), formDiv)
		case "Document":
			FormDivSelectFieldToField(&(documentuse_.Document), documentuseFormCallback.probe.stageOfInterest, formDiv)
		case "Lancelot:DocumentUse":
			// we need to retrieve the field owner before the change
			var pastLancelotOwner *models.Lancelot
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lancelot"
			rf.Fieldname = "DocumentUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				documentuseFormCallback.probe.stageOfInterest,
				documentuseFormCallback.probe.backRepoOfInterest,
				documentuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastLancelotOwner = reverseFieldOwner.(*models.Lancelot)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLancelotOwner != nil {
					idx := slices.Index(pastLancelotOwner.DocumentUse, documentuse_)
					pastLancelotOwner.DocumentUse = slices.Delete(pastLancelotOwner.DocumentUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _lancelot := range *models.GetGongstructInstancesSet[models.Lancelot](documentuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _lancelot.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLancelotOwner := _lancelot // we have a match
						if pastLancelotOwner != nil {
							if newLancelotOwner != pastLancelotOwner {
								idx := slices.Index(pastLancelotOwner.DocumentUse, documentuse_)
								pastLancelotOwner.DocumentUse = slices.Delete(pastLancelotOwner.DocumentUse, idx, idx+1)
								newLancelotOwner.DocumentUse = append(newLancelotOwner.DocumentUse, documentuse_)
							}
						} else {
							newLancelotOwner.DocumentUse = append(newLancelotOwner.DocumentUse, documentuse_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if documentuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentuse_.Unstage(documentuseFormCallback.probe.stageOfInterest)
	}

	documentuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DocumentUse](
		documentuseFormCallback.probe,
	)
	documentuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if documentuseFormCallback.CreationMode || documentuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(documentuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocumentUseFormCallback(
			nil,
			documentuseFormCallback.probe,
			newFormGroup,
		)
		documentuse := new(models.DocumentUse)
		FillUpForm(documentuse, newFormGroup, documentuseFormCallback.probe)
		documentuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(documentuseFormCallback.probe)
}
func __gong__New__GalahadThePureFormCallback(
	galahadthepure *models.GalahadThePure,
	probe *Probe,
	formGroup *table.FormGroup,
) (galahadthepureFormCallback *GalahadThePureFormCallback) {
	galahadthepureFormCallback = new(GalahadThePureFormCallback)
	galahadthepureFormCallback.probe = probe
	galahadthepureFormCallback.galahadthepure = galahadthepure
	galahadthepureFormCallback.formGroup = formGroup

	galahadthepureFormCallback.CreationMode = (galahadthepure == nil)

	return
}

type GalahadThePureFormCallback struct {
	galahadthepure *models.GalahadThePure

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (galahadthepureFormCallback *GalahadThePureFormCallback) OnSave() {

	log.Println("GalahadThePureFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	galahadthepureFormCallback.probe.formStage.Checkout()

	if galahadthepureFormCallback.galahadthepure == nil {
		galahadthepureFormCallback.galahadthepure = new(models.GalahadThePure).Stage(galahadthepureFormCallback.probe.stageOfInterest)
	}
	galahadthepure_ := galahadthepureFormCallback.galahadthepure
	_ = galahadthepure_

	for _, formDiv := range galahadthepureFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(galahadthepure_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(galahadthepure_.Description), formDiv)
		case "WhatIsYourPreferedColor:Galahard":
			// we need to retrieve the field owner before the change
			var pastWhatIsYourPreferedColorOwner *models.WhatIsYourPreferedColor
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "Galahard"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				galahadthepureFormCallback.probe.stageOfInterest,
				galahadthepureFormCallback.probe.backRepoOfInterest,
				galahadthepure_,
				&rf)

			if reverseFieldOwner != nil {
				pastWhatIsYourPreferedColorOwner = reverseFieldOwner.(*models.WhatIsYourPreferedColor)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastWhatIsYourPreferedColorOwner != nil {
					idx := slices.Index(pastWhatIsYourPreferedColorOwner.Galahard, galahadthepure_)
					pastWhatIsYourPreferedColorOwner.Galahard = slices.Delete(pastWhatIsYourPreferedColorOwner.Galahard, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _whatisyourpreferedcolor := range *models.GetGongstructInstancesSet[models.WhatIsYourPreferedColor](galahadthepureFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _whatisyourpreferedcolor.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newWhatIsYourPreferedColorOwner := _whatisyourpreferedcolor // we have a match
						if pastWhatIsYourPreferedColorOwner != nil {
							if newWhatIsYourPreferedColorOwner != pastWhatIsYourPreferedColorOwner {
								idx := slices.Index(pastWhatIsYourPreferedColorOwner.Galahard, galahadthepure_)
								pastWhatIsYourPreferedColorOwner.Galahard = slices.Delete(pastWhatIsYourPreferedColorOwner.Galahard, idx, idx+1)
								newWhatIsYourPreferedColorOwner.Galahard = append(newWhatIsYourPreferedColorOwner.Galahard, galahadthepure_)
							}
						} else {
							newWhatIsYourPreferedColorOwner.Galahard = append(newWhatIsYourPreferedColorOwner.Galahard, galahadthepure_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if galahadthepureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		galahadthepure_.Unstage(galahadthepureFormCallback.probe.stageOfInterest)
	}

	galahadthepureFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GalahadThePure](
		galahadthepureFormCallback.probe,
	)
	galahadthepureFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if galahadthepureFormCallback.CreationMode || galahadthepureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		galahadthepureFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(galahadthepureFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GalahadThePureFormCallback(
			nil,
			galahadthepureFormCallback.probe,
			newFormGroup,
		)
		galahadthepure := new(models.GalahadThePure)
		FillUpForm(galahadthepure, newFormGroup, galahadthepureFormCallback.probe)
		galahadthepureFormCallback.probe.formStage.Commit()
	}

	fillUpTree(galahadthepureFormCallback.probe)
}
func __gong__New__GeoObjectFormCallback(
	geoobject *models.GeoObject,
	probe *Probe,
	formGroup *table.FormGroup,
) (geoobjectFormCallback *GeoObjectFormCallback) {
	geoobjectFormCallback = new(GeoObjectFormCallback)
	geoobjectFormCallback.probe = probe
	geoobjectFormCallback.geoobject = geoobject
	geoobjectFormCallback.formGroup = formGroup

	geoobjectFormCallback.CreationMode = (geoobject == nil)

	return
}

type GeoObjectFormCallback struct {
	geoobject *models.GeoObject

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (geoobjectFormCallback *GeoObjectFormCallback) OnSave() {

	log.Println("GeoObjectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	geoobjectFormCallback.probe.formStage.Checkout()

	if geoobjectFormCallback.geoobject == nil {
		geoobjectFormCallback.geoobject = new(models.GeoObject).Stage(geoobjectFormCallback.probe.stageOfInterest)
	}
	geoobject_ := geoobjectFormCallback.geoobject
	_ = geoobject_

	for _, formDiv := range geoobjectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(geoobject_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if geoobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		geoobject_.Unstage(geoobjectFormCallback.probe.stageOfInterest)
	}

	geoobjectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GeoObject](
		geoobjectFormCallback.probe,
	)
	geoobjectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if geoobjectFormCallback.CreationMode || geoobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		geoobjectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(geoobjectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GeoObjectFormCallback(
			nil,
			geoobjectFormCallback.probe,
			newFormGroup,
		)
		geoobject := new(models.GeoObject)
		FillUpForm(geoobject, newFormGroup, geoobjectFormCallback.probe)
		geoobjectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(geoobjectFormCallback.probe)
}
func __gong__New__GeoObjectUseFormCallback(
	geoobjectuse *models.GeoObjectUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (geoobjectuseFormCallback *GeoObjectUseFormCallback) {
	geoobjectuseFormCallback = new(GeoObjectUseFormCallback)
	geoobjectuseFormCallback.probe = probe
	geoobjectuseFormCallback.geoobjectuse = geoobjectuse
	geoobjectuseFormCallback.formGroup = formGroup

	geoobjectuseFormCallback.CreationMode = (geoobjectuse == nil)

	return
}

type GeoObjectUseFormCallback struct {
	geoobjectuse *models.GeoObjectUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (geoobjectuseFormCallback *GeoObjectUseFormCallback) OnSave() {

	log.Println("GeoObjectUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	geoobjectuseFormCallback.probe.formStage.Checkout()

	if geoobjectuseFormCallback.geoobjectuse == nil {
		geoobjectuseFormCallback.geoobjectuse = new(models.GeoObjectUse).Stage(geoobjectuseFormCallback.probe.stageOfInterest)
	}
	geoobjectuse_ := geoobjectuseFormCallback.geoobjectuse
	_ = geoobjectuse_

	for _, formDiv := range geoobjectuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(geoobjectuse_.Name), formDiv)
		case "GeoObject":
			FormDivSelectFieldToField(&(geoobjectuse_.GeoObject), geoobjectuseFormCallback.probe.stageOfInterest, formDiv)
		case "Document:GeoObjectUse":
			// we need to retrieve the field owner before the change
			var pastDocumentOwner *models.Document
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Document"
			rf.Fieldname = "GeoObjectUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				geoobjectuseFormCallback.probe.stageOfInterest,
				geoobjectuseFormCallback.probe.backRepoOfInterest,
				geoobjectuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastDocumentOwner = reverseFieldOwner.(*models.Document)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDocumentOwner != nil {
					idx := slices.Index(pastDocumentOwner.GeoObjectUse, geoobjectuse_)
					pastDocumentOwner.GeoObjectUse = slices.Delete(pastDocumentOwner.GeoObjectUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _document := range *models.GetGongstructInstancesSet[models.Document](geoobjectuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _document.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDocumentOwner := _document // we have a match
						if pastDocumentOwner != nil {
							if newDocumentOwner != pastDocumentOwner {
								idx := slices.Index(pastDocumentOwner.GeoObjectUse, geoobjectuse_)
								pastDocumentOwner.GeoObjectUse = slices.Delete(pastDocumentOwner.GeoObjectUse, idx, idx+1)
								newDocumentOwner.GeoObjectUse = append(newDocumentOwner.GeoObjectUse, geoobjectuse_)
							}
						} else {
							newDocumentOwner.GeoObjectUse = append(newDocumentOwner.GeoObjectUse, geoobjectuse_)
						}
					}
				}
			}
		case "Lancelot:GeoObjectUse":
			// we need to retrieve the field owner before the change
			var pastLancelotOwner *models.Lancelot
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lancelot"
			rf.Fieldname = "GeoObjectUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				geoobjectuseFormCallback.probe.stageOfInterest,
				geoobjectuseFormCallback.probe.backRepoOfInterest,
				geoobjectuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastLancelotOwner = reverseFieldOwner.(*models.Lancelot)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLancelotOwner != nil {
					idx := slices.Index(pastLancelotOwner.GeoObjectUse, geoobjectuse_)
					pastLancelotOwner.GeoObjectUse = slices.Delete(pastLancelotOwner.GeoObjectUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _lancelot := range *models.GetGongstructInstancesSet[models.Lancelot](geoobjectuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _lancelot.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLancelotOwner := _lancelot // we have a match
						if pastLancelotOwner != nil {
							if newLancelotOwner != pastLancelotOwner {
								idx := slices.Index(pastLancelotOwner.GeoObjectUse, geoobjectuse_)
								pastLancelotOwner.GeoObjectUse = slices.Delete(pastLancelotOwner.GeoObjectUse, idx, idx+1)
								newLancelotOwner.GeoObjectUse = append(newLancelotOwner.GeoObjectUse, geoobjectuse_)
							}
						} else {
							newLancelotOwner.GeoObjectUse = append(newLancelotOwner.GeoObjectUse, geoobjectuse_)
						}
					}
				}
			}
		case "TheBridge:GeoObjectUse":
			// we need to retrieve the field owner before the change
			var pastTheBridgeOwner *models.TheBridge
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TheBridge"
			rf.Fieldname = "GeoObjectUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				geoobjectuseFormCallback.probe.stageOfInterest,
				geoobjectuseFormCallback.probe.backRepoOfInterest,
				geoobjectuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastTheBridgeOwner = reverseFieldOwner.(*models.TheBridge)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTheBridgeOwner != nil {
					idx := slices.Index(pastTheBridgeOwner.GeoObjectUse, geoobjectuse_)
					pastTheBridgeOwner.GeoObjectUse = slices.Delete(pastTheBridgeOwner.GeoObjectUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _thebridge := range *models.GetGongstructInstancesSet[models.TheBridge](geoobjectuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _thebridge.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTheBridgeOwner := _thebridge // we have a match
						if pastTheBridgeOwner != nil {
							if newTheBridgeOwner != pastTheBridgeOwner {
								idx := slices.Index(pastTheBridgeOwner.GeoObjectUse, geoobjectuse_)
								pastTheBridgeOwner.GeoObjectUse = slices.Delete(pastTheBridgeOwner.GeoObjectUse, idx, idx+1)
								newTheBridgeOwner.GeoObjectUse = append(newTheBridgeOwner.GeoObjectUse, geoobjectuse_)
							}
						} else {
							newTheBridgeOwner.GeoObjectUse = append(newTheBridgeOwner.GeoObjectUse, geoobjectuse_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if geoobjectuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		geoobjectuse_.Unstage(geoobjectuseFormCallback.probe.stageOfInterest)
	}

	geoobjectuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GeoObjectUse](
		geoobjectuseFormCallback.probe,
	)
	geoobjectuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if geoobjectuseFormCallback.CreationMode || geoobjectuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		geoobjectuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(geoobjectuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GeoObjectUseFormCallback(
			nil,
			geoobjectuseFormCallback.probe,
			newFormGroup,
		)
		geoobjectuse := new(models.GeoObjectUse)
		FillUpForm(geoobjectuse, newFormGroup, geoobjectuseFormCallback.probe)
		geoobjectuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(geoobjectuseFormCallback.probe)
}
func __gong__New__GroupFormCallback(
	group *models.Group,
	probe *Probe,
	formGroup *table.FormGroup,
) (groupFormCallback *GroupFormCallback) {
	groupFormCallback = new(GroupFormCallback)
	groupFormCallback.probe = probe
	groupFormCallback.group = group
	groupFormCallback.formGroup = formGroup

	groupFormCallback.CreationMode = (group == nil)

	return
}

type GroupFormCallback struct {
	group *models.Group

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (groupFormCallback *GroupFormCallback) OnSave() {

	log.Println("GroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupFormCallback.probe.formStage.Checkout()

	if groupFormCallback.group == nil {
		groupFormCallback.group = new(models.Group).Stage(groupFormCallback.probe.stageOfInterest)
	}
	group_ := groupFormCallback.group
	_ = group_

	for _, formDiv := range groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(group_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_.Unstage(groupFormCallback.probe.stageOfInterest)
	}

	groupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Group](
		groupFormCallback.probe,
	)
	groupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if groupFormCallback.CreationMode || groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			groupFormCallback.probe,
			newFormGroup,
		)
		group := new(models.Group)
		FillUpForm(group, newFormGroup, groupFormCallback.probe)
		groupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(groupFormCallback.probe)
}
func __gong__New__GroupUseFormCallback(
	groupuse *models.GroupUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (groupuseFormCallback *GroupUseFormCallback) {
	groupuseFormCallback = new(GroupUseFormCallback)
	groupuseFormCallback.probe = probe
	groupuseFormCallback.groupuse = groupuse
	groupuseFormCallback.formGroup = formGroup

	groupuseFormCallback.CreationMode = (groupuse == nil)

	return
}

type GroupUseFormCallback struct {
	groupuse *models.GroupUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (groupuseFormCallback *GroupUseFormCallback) OnSave() {

	log.Println("GroupUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupuseFormCallback.probe.formStage.Checkout()

	if groupuseFormCallback.groupuse == nil {
		groupuseFormCallback.groupuse = new(models.GroupUse).Stage(groupuseFormCallback.probe.stageOfInterest)
	}
	groupuse_ := groupuseFormCallback.groupuse
	_ = groupuse_

	for _, formDiv := range groupuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(groupuse_.Name), formDiv)
		case "Group":
			FormDivSelectFieldToField(&(groupuse_.Group), groupuseFormCallback.probe.stageOfInterest, formDiv)
		case "Lancelot:GroupUse":
			// we need to retrieve the field owner before the change
			var pastLancelotOwner *models.Lancelot
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Lancelot"
			rf.Fieldname = "GroupUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupuseFormCallback.probe.stageOfInterest,
				groupuseFormCallback.probe.backRepoOfInterest,
				groupuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastLancelotOwner = reverseFieldOwner.(*models.Lancelot)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLancelotOwner != nil {
					idx := slices.Index(pastLancelotOwner.GroupUse, groupuse_)
					pastLancelotOwner.GroupUse = slices.Delete(pastLancelotOwner.GroupUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _lancelot := range *models.GetGongstructInstancesSet[models.Lancelot](groupuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _lancelot.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLancelotOwner := _lancelot // we have a match
						if pastLancelotOwner != nil {
							if newLancelotOwner != pastLancelotOwner {
								idx := slices.Index(pastLancelotOwner.GroupUse, groupuse_)
								pastLancelotOwner.GroupUse = slices.Delete(pastLancelotOwner.GroupUse, idx, idx+1)
								newLancelotOwner.GroupUse = append(newLancelotOwner.GroupUse, groupuse_)
							}
						} else {
							newLancelotOwner.GroupUse = append(newLancelotOwner.GroupUse, groupuse_)
						}
					}
				}
			}
		case "Repository:GroupUse":
			// we need to retrieve the field owner before the change
			var pastRepositoryOwner *models.Repository
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Repository"
			rf.Fieldname = "GroupUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupuseFormCallback.probe.stageOfInterest,
				groupuseFormCallback.probe.backRepoOfInterest,
				groupuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastRepositoryOwner = reverseFieldOwner.(*models.Repository)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRepositoryOwner != nil {
					idx := slices.Index(pastRepositoryOwner.GroupUse, groupuse_)
					pastRepositoryOwner.GroupUse = slices.Delete(pastRepositoryOwner.GroupUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _repository := range *models.GetGongstructInstancesSet[models.Repository](groupuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _repository.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRepositoryOwner := _repository // we have a match
						if pastRepositoryOwner != nil {
							if newRepositoryOwner != pastRepositoryOwner {
								idx := slices.Index(pastRepositoryOwner.GroupUse, groupuse_)
								pastRepositoryOwner.GroupUse = slices.Delete(pastRepositoryOwner.GroupUse, idx, idx+1)
								newRepositoryOwner.GroupUse = append(newRepositoryOwner.GroupUse, groupuse_)
							}
						} else {
							newRepositoryOwner.GroupUse = append(newRepositoryOwner.GroupUse, groupuse_)
						}
					}
				}
			}
		case "TheBridge:GroupUse":
			// we need to retrieve the field owner before the change
			var pastTheBridgeOwner *models.TheBridge
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TheBridge"
			rf.Fieldname = "GroupUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupuseFormCallback.probe.stageOfInterest,
				groupuseFormCallback.probe.backRepoOfInterest,
				groupuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastTheBridgeOwner = reverseFieldOwner.(*models.TheBridge)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTheBridgeOwner != nil {
					idx := slices.Index(pastTheBridgeOwner.GroupUse, groupuse_)
					pastTheBridgeOwner.GroupUse = slices.Delete(pastTheBridgeOwner.GroupUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _thebridge := range *models.GetGongstructInstancesSet[models.TheBridge](groupuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _thebridge.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTheBridgeOwner := _thebridge // we have a match
						if pastTheBridgeOwner != nil {
							if newTheBridgeOwner != pastTheBridgeOwner {
								idx := slices.Index(pastTheBridgeOwner.GroupUse, groupuse_)
								pastTheBridgeOwner.GroupUse = slices.Delete(pastTheBridgeOwner.GroupUse, idx, idx+1)
								newTheBridgeOwner.GroupUse = append(newTheBridgeOwner.GroupUse, groupuse_)
							}
						} else {
							newTheBridgeOwner.GroupUse = append(newTheBridgeOwner.GroupUse, groupuse_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if groupuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupuse_.Unstage(groupuseFormCallback.probe.stageOfInterest)
	}

	groupuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GroupUse](
		groupuseFormCallback.probe,
	)
	groupuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if groupuseFormCallback.CreationMode || groupuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(groupuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupUseFormCallback(
			nil,
			groupuseFormCallback.probe,
			newFormGroup,
		)
		groupuse := new(models.GroupUse)
		FillUpForm(groupuse, newFormGroup, groupuseFormCallback.probe)
		groupuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(groupuseFormCallback.probe)
}
func __gong__New__KingArthurFormCallback(
	kingarthur *models.KingArthur,
	probe *Probe,
	formGroup *table.FormGroup,
) (kingarthurFormCallback *KingArthurFormCallback) {
	kingarthurFormCallback = new(KingArthurFormCallback)
	kingarthurFormCallback.probe = probe
	kingarthurFormCallback.kingarthur = kingarthur
	kingarthurFormCallback.formGroup = formGroup

	kingarthurFormCallback.CreationMode = (kingarthur == nil)

	return
}

type KingArthurFormCallback struct {
	kingarthur *models.KingArthur

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (kingarthurFormCallback *KingArthurFormCallback) OnSave() {

	log.Println("KingArthurFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	kingarthurFormCallback.probe.formStage.Checkout()

	if kingarthurFormCallback.kingarthur == nil {
		kingarthurFormCallback.kingarthur = new(models.KingArthur).Stage(kingarthurFormCallback.probe.stageOfInterest)
	}
	kingarthur_ := kingarthurFormCallback.kingarthur
	_ = kingarthur_

	for _, formDiv := range kingarthurFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(kingarthur_.Name), formDiv)
		case "IsWithProbaility":
			FormDivBasicFieldToField(&(kingarthur_.IsWithProbaility), formDiv)
		case "Probability":
			FormDivEnumStringFieldToField(&(kingarthur_.Probability), formDiv)
		case "WhatIsYourPreferedColor:KingArthurs":
			// we need to retrieve the field owner before the change
			var pastWhatIsYourPreferedColorOwner *models.WhatIsYourPreferedColor
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "KingArthurs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				kingarthurFormCallback.probe.stageOfInterest,
				kingarthurFormCallback.probe.backRepoOfInterest,
				kingarthur_,
				&rf)

			if reverseFieldOwner != nil {
				pastWhatIsYourPreferedColorOwner = reverseFieldOwner.(*models.WhatIsYourPreferedColor)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastWhatIsYourPreferedColorOwner != nil {
					idx := slices.Index(pastWhatIsYourPreferedColorOwner.KingArthurs, kingarthur_)
					pastWhatIsYourPreferedColorOwner.KingArthurs = slices.Delete(pastWhatIsYourPreferedColorOwner.KingArthurs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _whatisyourpreferedcolor := range *models.GetGongstructInstancesSet[models.WhatIsYourPreferedColor](kingarthurFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _whatisyourpreferedcolor.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newWhatIsYourPreferedColorOwner := _whatisyourpreferedcolor // we have a match
						if pastWhatIsYourPreferedColorOwner != nil {
							if newWhatIsYourPreferedColorOwner != pastWhatIsYourPreferedColorOwner {
								idx := slices.Index(pastWhatIsYourPreferedColorOwner.KingArthurs, kingarthur_)
								pastWhatIsYourPreferedColorOwner.KingArthurs = slices.Delete(pastWhatIsYourPreferedColorOwner.KingArthurs, idx, idx+1)
								newWhatIsYourPreferedColorOwner.KingArthurs = append(newWhatIsYourPreferedColorOwner.KingArthurs, kingarthur_)
							}
						} else {
							newWhatIsYourPreferedColorOwner.KingArthurs = append(newWhatIsYourPreferedColorOwner.KingArthurs, kingarthur_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if kingarthurFormCallback.formGroup.HasSuppressButtonBeenPressed {
		kingarthur_.Unstage(kingarthurFormCallback.probe.stageOfInterest)
	}

	kingarthurFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.KingArthur](
		kingarthurFormCallback.probe,
	)
	kingarthurFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if kingarthurFormCallback.CreationMode || kingarthurFormCallback.formGroup.HasSuppressButtonBeenPressed {
		kingarthurFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(kingarthurFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__KingArthurFormCallback(
			nil,
			kingarthurFormCallback.probe,
			newFormGroup,
		)
		kingarthur := new(models.KingArthur)
		FillUpForm(kingarthur, newFormGroup, kingarthurFormCallback.probe)
		kingarthurFormCallback.probe.formStage.Commit()
	}

	fillUpTree(kingarthurFormCallback.probe)
}
func __gong__New__KingArthurShapeFormCallback(
	kingarthurshape *models.KingArthurShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (kingarthurshapeFormCallback *KingArthurShapeFormCallback) {
	kingarthurshapeFormCallback = new(KingArthurShapeFormCallback)
	kingarthurshapeFormCallback.probe = probe
	kingarthurshapeFormCallback.kingarthurshape = kingarthurshape
	kingarthurshapeFormCallback.formGroup = formGroup

	kingarthurshapeFormCallback.CreationMode = (kingarthurshape == nil)

	return
}

type KingArthurShapeFormCallback struct {
	kingarthurshape *models.KingArthurShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (kingarthurshapeFormCallback *KingArthurShapeFormCallback) OnSave() {

	log.Println("KingArthurShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	kingarthurshapeFormCallback.probe.formStage.Checkout()

	if kingarthurshapeFormCallback.kingarthurshape == nil {
		kingarthurshapeFormCallback.kingarthurshape = new(models.KingArthurShape).Stage(kingarthurshapeFormCallback.probe.stageOfInterest)
	}
	kingarthurshape_ := kingarthurshapeFormCallback.kingarthurshape
	_ = kingarthurshape_

	for _, formDiv := range kingarthurshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(kingarthurshape_.Name), formDiv)
		case "ActorState":
			FormDivSelectFieldToField(&(kingarthurshape_.ActorState), kingarthurshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(kingarthurshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(kingarthurshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(kingarthurshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(kingarthurshape_.Height), formDiv)
		case "FillColor":
			FormDivBasicFieldToField(&(kingarthurshape_.FillColor), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(kingarthurshape_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(kingarthurshape_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(kingarthurshape_.StrokeWidth), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(kingarthurshape_.RX), formDiv)
		case "SirRobin:Arthurs":
			// we need to retrieve the field owner before the change
			var pastSirRobinOwner *models.SirRobin
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "Arthurs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				kingarthurshapeFormCallback.probe.stageOfInterest,
				kingarthurshapeFormCallback.probe.backRepoOfInterest,
				kingarthurshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastSirRobinOwner = reverseFieldOwner.(*models.SirRobin)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSirRobinOwner != nil {
					idx := slices.Index(pastSirRobinOwner.Arthurs, kingarthurshape_)
					pastSirRobinOwner.Arthurs = slices.Delete(pastSirRobinOwner.Arthurs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sirrobin := range *models.GetGongstructInstancesSet[models.SirRobin](kingarthurshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sirrobin.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSirRobinOwner := _sirrobin // we have a match
						if pastSirRobinOwner != nil {
							if newSirRobinOwner != pastSirRobinOwner {
								idx := slices.Index(pastSirRobinOwner.Arthurs, kingarthurshape_)
								pastSirRobinOwner.Arthurs = slices.Delete(pastSirRobinOwner.Arthurs, idx, idx+1)
								newSirRobinOwner.Arthurs = append(newSirRobinOwner.Arthurs, kingarthurshape_)
							}
						} else {
							newSirRobinOwner.Arthurs = append(newSirRobinOwner.Arthurs, kingarthurshape_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if kingarthurshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		kingarthurshape_.Unstage(kingarthurshapeFormCallback.probe.stageOfInterest)
	}

	kingarthurshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.KingArthurShape](
		kingarthurshapeFormCallback.probe,
	)
	kingarthurshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if kingarthurshapeFormCallback.CreationMode || kingarthurshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		kingarthurshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(kingarthurshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__KingArthurShapeFormCallback(
			nil,
			kingarthurshapeFormCallback.probe,
			newFormGroup,
		)
		kingarthurshape := new(models.KingArthurShape)
		FillUpForm(kingarthurshape, newFormGroup, kingarthurshapeFormCallback.probe)
		kingarthurshapeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(kingarthurshapeFormCallback.probe)
}
func __gong__New__KnightWhoSayNiFormCallback(
	knightwhosayni *models.KnightWhoSayNi,
	probe *Probe,
	formGroup *table.FormGroup,
) (knightwhosayniFormCallback *KnightWhoSayNiFormCallback) {
	knightwhosayniFormCallback = new(KnightWhoSayNiFormCallback)
	knightwhosayniFormCallback.probe = probe
	knightwhosayniFormCallback.knightwhosayni = knightwhosayni
	knightwhosayniFormCallback.formGroup = formGroup

	knightwhosayniFormCallback.CreationMode = (knightwhosayni == nil)

	return
}

type KnightWhoSayNiFormCallback struct {
	knightwhosayni *models.KnightWhoSayNi

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (knightwhosayniFormCallback *KnightWhoSayNiFormCallback) OnSave() {

	log.Println("KnightWhoSayNiFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	knightwhosayniFormCallback.probe.formStage.Checkout()

	if knightwhosayniFormCallback.knightwhosayni == nil {
		knightwhosayniFormCallback.knightwhosayni = new(models.KnightWhoSayNi).Stage(knightwhosayniFormCallback.probe.stageOfInterest)
	}
	knightwhosayni_ := knightwhosayniFormCallback.knightwhosayni
	_ = knightwhosayni_

	for _, formDiv := range knightwhosayniFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(knightwhosayni_.Name), formDiv)
		case "Parameter":
			FormDivSelectFieldToField(&(knightwhosayni_.Parameter), knightwhosayniFormCallback.probe.stageOfInterest, formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(knightwhosayni_.Direction), formDiv)
		case "X":
			FormDivBasicFieldToField(&(knightwhosayni_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(knightwhosayni_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(knightwhosayni_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(knightwhosayni_.Height), formDiv)
		case "FillColor":
			FormDivBasicFieldToField(&(knightwhosayni_.FillColor), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(knightwhosayni_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(knightwhosayni_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(knightwhosayni_.StrokeWidth), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(knightwhosayni_.RX), formDiv)
		case "LancelotAgregation:ParameterUse":
			// we need to retrieve the field owner before the change
			var pastLancelotAgregationOwner *models.LancelotAgregation
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LancelotAgregation"
			rf.Fieldname = "ParameterUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				knightwhosayniFormCallback.probe.stageOfInterest,
				knightwhosayniFormCallback.probe.backRepoOfInterest,
				knightwhosayni_,
				&rf)

			if reverseFieldOwner != nil {
				pastLancelotAgregationOwner = reverseFieldOwner.(*models.LancelotAgregation)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLancelotAgregationOwner != nil {
					idx := slices.Index(pastLancelotAgregationOwner.ParameterUse, knightwhosayni_)
					pastLancelotAgregationOwner.ParameterUse = slices.Delete(pastLancelotAgregationOwner.ParameterUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _lancelotagregation := range *models.GetGongstructInstancesSet[models.LancelotAgregation](knightwhosayniFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _lancelotagregation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLancelotAgregationOwner := _lancelotagregation // we have a match
						if pastLancelotAgregationOwner != nil {
							if newLancelotAgregationOwner != pastLancelotAgregationOwner {
								idx := slices.Index(pastLancelotAgregationOwner.ParameterUse, knightwhosayni_)
								pastLancelotAgregationOwner.ParameterUse = slices.Delete(pastLancelotAgregationOwner.ParameterUse, idx, idx+1)
								newLancelotAgregationOwner.ParameterUse = append(newLancelotAgregationOwner.ParameterUse, knightwhosayni_)
							}
						} else {
							newLancelotAgregationOwner.ParameterUse = append(newLancelotAgregationOwner.ParameterUse, knightwhosayni_)
						}
					}
				}
			}
		case "LancelotCategory:ParameterUse":
			// we need to retrieve the field owner before the change
			var pastLancelotCategoryOwner *models.LancelotCategory
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LancelotCategory"
			rf.Fieldname = "ParameterUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				knightwhosayniFormCallback.probe.stageOfInterest,
				knightwhosayniFormCallback.probe.backRepoOfInterest,
				knightwhosayni_,
				&rf)

			if reverseFieldOwner != nil {
				pastLancelotCategoryOwner = reverseFieldOwner.(*models.LancelotCategory)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLancelotCategoryOwner != nil {
					idx := slices.Index(pastLancelotCategoryOwner.ParameterUse, knightwhosayni_)
					pastLancelotCategoryOwner.ParameterUse = slices.Delete(pastLancelotCategoryOwner.ParameterUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _lancelotcategory := range *models.GetGongstructInstancesSet[models.LancelotCategory](knightwhosayniFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _lancelotcategory.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLancelotCategoryOwner := _lancelotcategory // we have a match
						if pastLancelotCategoryOwner != nil {
							if newLancelotCategoryOwner != pastLancelotCategoryOwner {
								idx := slices.Index(pastLancelotCategoryOwner.ParameterUse, knightwhosayni_)
								pastLancelotCategoryOwner.ParameterUse = slices.Delete(pastLancelotCategoryOwner.ParameterUse, idx, idx+1)
								newLancelotCategoryOwner.ParameterUse = append(newLancelotCategoryOwner.ParameterUse, knightwhosayni_)
							}
						} else {
							newLancelotCategoryOwner.ParameterUse = append(newLancelotCategoryOwner.ParameterUse, knightwhosayni_)
						}
					}
				}
			}
		case "Repository:ParameterUse":
			// we need to retrieve the field owner before the change
			var pastRepositoryOwner *models.Repository
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Repository"
			rf.Fieldname = "ParameterUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				knightwhosayniFormCallback.probe.stageOfInterest,
				knightwhosayniFormCallback.probe.backRepoOfInterest,
				knightwhosayni_,
				&rf)

			if reverseFieldOwner != nil {
				pastRepositoryOwner = reverseFieldOwner.(*models.Repository)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRepositoryOwner != nil {
					idx := slices.Index(pastRepositoryOwner.ParameterUse, knightwhosayni_)
					pastRepositoryOwner.ParameterUse = slices.Delete(pastRepositoryOwner.ParameterUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _repository := range *models.GetGongstructInstancesSet[models.Repository](knightwhosayniFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _repository.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRepositoryOwner := _repository // we have a match
						if pastRepositoryOwner != nil {
							if newRepositoryOwner != pastRepositoryOwner {
								idx := slices.Index(pastRepositoryOwner.ParameterUse, knightwhosayni_)
								pastRepositoryOwner.ParameterUse = slices.Delete(pastRepositoryOwner.ParameterUse, idx, idx+1)
								newRepositoryOwner.ParameterUse = append(newRepositoryOwner.ParameterUse, knightwhosayni_)
							}
						} else {
							newRepositoryOwner.ParameterUse = append(newRepositoryOwner.ParameterUse, knightwhosayni_)
						}
					}
				}
			}
		case "SirRobin:KnightWhoSayNis":
			// we need to retrieve the field owner before the change
			var pastSirRobinOwner *models.SirRobin
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "KnightWhoSayNis"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				knightwhosayniFormCallback.probe.stageOfInterest,
				knightwhosayniFormCallback.probe.backRepoOfInterest,
				knightwhosayni_,
				&rf)

			if reverseFieldOwner != nil {
				pastSirRobinOwner = reverseFieldOwner.(*models.SirRobin)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSirRobinOwner != nil {
					idx := slices.Index(pastSirRobinOwner.KnightWhoSayNis, knightwhosayni_)
					pastSirRobinOwner.KnightWhoSayNis = slices.Delete(pastSirRobinOwner.KnightWhoSayNis, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sirrobin := range *models.GetGongstructInstancesSet[models.SirRobin](knightwhosayniFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sirrobin.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSirRobinOwner := _sirrobin // we have a match
						if pastSirRobinOwner != nil {
							if newSirRobinOwner != pastSirRobinOwner {
								idx := slices.Index(pastSirRobinOwner.KnightWhoSayNis, knightwhosayni_)
								pastSirRobinOwner.KnightWhoSayNis = slices.Delete(pastSirRobinOwner.KnightWhoSayNis, idx, idx+1)
								newSirRobinOwner.KnightWhoSayNis = append(newSirRobinOwner.KnightWhoSayNis, knightwhosayni_)
							}
						} else {
							newSirRobinOwner.KnightWhoSayNis = append(newSirRobinOwner.KnightWhoSayNis, knightwhosayni_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if knightwhosayniFormCallback.formGroup.HasSuppressButtonBeenPressed {
		knightwhosayni_.Unstage(knightwhosayniFormCallback.probe.stageOfInterest)
	}

	knightwhosayniFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.KnightWhoSayNi](
		knightwhosayniFormCallback.probe,
	)
	knightwhosayniFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if knightwhosayniFormCallback.CreationMode || knightwhosayniFormCallback.formGroup.HasSuppressButtonBeenPressed {
		knightwhosayniFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(knightwhosayniFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__KnightWhoSayNiFormCallback(
			nil,
			knightwhosayniFormCallback.probe,
			newFormGroup,
		)
		knightwhosayni := new(models.KnightWhoSayNi)
		FillUpForm(knightwhosayni, newFormGroup, knightwhosayniFormCallback.probe)
		knightwhosayniFormCallback.probe.formStage.Commit()
	}

	fillUpTree(knightwhosayniFormCallback.probe)
}
func __gong__New__LancelotFormCallback(
	lancelot *models.Lancelot,
	probe *Probe,
	formGroup *table.FormGroup,
) (lancelotFormCallback *LancelotFormCallback) {
	lancelotFormCallback = new(LancelotFormCallback)
	lancelotFormCallback.probe = probe
	lancelotFormCallback.lancelot = lancelot
	lancelotFormCallback.formGroup = formGroup

	lancelotFormCallback.CreationMode = (lancelot == nil)

	return
}

type LancelotFormCallback struct {
	lancelot *models.Lancelot

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lancelotFormCallback *LancelotFormCallback) OnSave() {

	log.Println("LancelotFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lancelotFormCallback.probe.formStage.Checkout()

	if lancelotFormCallback.lancelot == nil {
		lancelotFormCallback.lancelot = new(models.Lancelot).Stage(lancelotFormCallback.probe.stageOfInterest)
	}
	lancelot_ := lancelotFormCallback.lancelot
	_ = lancelot_

	for _, formDiv := range lancelotFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lancelot_.Name), formDiv)
		case "Tag":
			FormDivBasicFieldToField(&(lancelot_.Tag), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(lancelot_.Description), formDiv)
		case "BringYourDead:Lancelots":
			// we need to retrieve the field owner before the change
			var pastBringYourDeadOwner *models.BringYourDead
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BringYourDead"
			rf.Fieldname = "Lancelots"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				lancelotFormCallback.probe.stageOfInterest,
				lancelotFormCallback.probe.backRepoOfInterest,
				lancelot_,
				&rf)

			if reverseFieldOwner != nil {
				pastBringYourDeadOwner = reverseFieldOwner.(*models.BringYourDead)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBringYourDeadOwner != nil {
					idx := slices.Index(pastBringYourDeadOwner.Lancelots, lancelot_)
					pastBringYourDeadOwner.Lancelots = slices.Delete(pastBringYourDeadOwner.Lancelots, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _bringyourdead := range *models.GetGongstructInstancesSet[models.BringYourDead](lancelotFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _bringyourdead.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBringYourDeadOwner := _bringyourdead // we have a match
						if pastBringYourDeadOwner != nil {
							if newBringYourDeadOwner != pastBringYourDeadOwner {
								idx := slices.Index(pastBringYourDeadOwner.Lancelots, lancelot_)
								pastBringYourDeadOwner.Lancelots = slices.Delete(pastBringYourDeadOwner.Lancelots, idx, idx+1)
								newBringYourDeadOwner.Lancelots = append(newBringYourDeadOwner.Lancelots, lancelot_)
							}
						} else {
							newBringYourDeadOwner.Lancelots = append(newBringYourDeadOwner.Lancelots, lancelot_)
						}
					}
				}
			}
		case "WhatIsYourPreferedColor:Lancelots":
			// we need to retrieve the field owner before the change
			var pastWhatIsYourPreferedColorOwner *models.WhatIsYourPreferedColor
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "Lancelots"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				lancelotFormCallback.probe.stageOfInterest,
				lancelotFormCallback.probe.backRepoOfInterest,
				lancelot_,
				&rf)

			if reverseFieldOwner != nil {
				pastWhatIsYourPreferedColorOwner = reverseFieldOwner.(*models.WhatIsYourPreferedColor)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastWhatIsYourPreferedColorOwner != nil {
					idx := slices.Index(pastWhatIsYourPreferedColorOwner.Lancelots, lancelot_)
					pastWhatIsYourPreferedColorOwner.Lancelots = slices.Delete(pastWhatIsYourPreferedColorOwner.Lancelots, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _whatisyourpreferedcolor := range *models.GetGongstructInstancesSet[models.WhatIsYourPreferedColor](lancelotFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _whatisyourpreferedcolor.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newWhatIsYourPreferedColorOwner := _whatisyourpreferedcolor // we have a match
						if pastWhatIsYourPreferedColorOwner != nil {
							if newWhatIsYourPreferedColorOwner != pastWhatIsYourPreferedColorOwner {
								idx := slices.Index(pastWhatIsYourPreferedColorOwner.Lancelots, lancelot_)
								pastWhatIsYourPreferedColorOwner.Lancelots = slices.Delete(pastWhatIsYourPreferedColorOwner.Lancelots, idx, idx+1)
								newWhatIsYourPreferedColorOwner.Lancelots = append(newWhatIsYourPreferedColorOwner.Lancelots, lancelot_)
							}
						} else {
							newWhatIsYourPreferedColorOwner.Lancelots = append(newWhatIsYourPreferedColorOwner.Lancelots, lancelot_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if lancelotFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelot_.Unstage(lancelotFormCallback.probe.stageOfInterest)
	}

	lancelotFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Lancelot](
		lancelotFormCallback.probe,
	)
	lancelotFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lancelotFormCallback.CreationMode || lancelotFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lancelotFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LancelotFormCallback(
			nil,
			lancelotFormCallback.probe,
			newFormGroup,
		)
		lancelot := new(models.Lancelot)
		FillUpForm(lancelot, newFormGroup, lancelotFormCallback.probe)
		lancelotFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lancelotFormCallback.probe)
}
func __gong__New__LancelotAgregationFormCallback(
	lancelotagregation *models.LancelotAgregation,
	probe *Probe,
	formGroup *table.FormGroup,
) (lancelotagregationFormCallback *LancelotAgregationFormCallback) {
	lancelotagregationFormCallback = new(LancelotAgregationFormCallback)
	lancelotagregationFormCallback.probe = probe
	lancelotagregationFormCallback.lancelotagregation = lancelotagregation
	lancelotagregationFormCallback.formGroup = formGroup

	lancelotagregationFormCallback.CreationMode = (lancelotagregation == nil)

	return
}

type LancelotAgregationFormCallback struct {
	lancelotagregation *models.LancelotAgregation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lancelotagregationFormCallback *LancelotAgregationFormCallback) OnSave() {

	log.Println("LancelotAgregationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lancelotagregationFormCallback.probe.formStage.Checkout()

	if lancelotagregationFormCallback.lancelotagregation == nil {
		lancelotagregationFormCallback.lancelotagregation = new(models.LancelotAgregation).Stage(lancelotagregationFormCallback.probe.stageOfInterest)
	}
	lancelotagregation_ := lancelotagregationFormCallback.lancelotagregation
	_ = lancelotagregation_

	for _, formDiv := range lancelotagregationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lancelotagregation_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if lancelotagregationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotagregation_.Unstage(lancelotagregationFormCallback.probe.stageOfInterest)
	}

	lancelotagregationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LancelotAgregation](
		lancelotagregationFormCallback.probe,
	)
	lancelotagregationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lancelotagregationFormCallback.CreationMode || lancelotagregationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotagregationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lancelotagregationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LancelotAgregationFormCallback(
			nil,
			lancelotagregationFormCallback.probe,
			newFormGroup,
		)
		lancelotagregation := new(models.LancelotAgregation)
		FillUpForm(lancelotagregation, newFormGroup, lancelotagregationFormCallback.probe)
		lancelotagregationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lancelotagregationFormCallback.probe)
}
func __gong__New__LancelotAgregationUseFormCallback(
	lancelotagregationuse *models.LancelotAgregationUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (lancelotagregationuseFormCallback *LancelotAgregationUseFormCallback) {
	lancelotagregationuseFormCallback = new(LancelotAgregationUseFormCallback)
	lancelotagregationuseFormCallback.probe = probe
	lancelotagregationuseFormCallback.lancelotagregationuse = lancelotagregationuse
	lancelotagregationuseFormCallback.formGroup = formGroup

	lancelotagregationuseFormCallback.CreationMode = (lancelotagregationuse == nil)

	return
}

type LancelotAgregationUseFormCallback struct {
	lancelotagregationuse *models.LancelotAgregationUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lancelotagregationuseFormCallback *LancelotAgregationUseFormCallback) OnSave() {

	log.Println("LancelotAgregationUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lancelotagregationuseFormCallback.probe.formStage.Checkout()

	if lancelotagregationuseFormCallback.lancelotagregationuse == nil {
		lancelotagregationuseFormCallback.lancelotagregationuse = new(models.LancelotAgregationUse).Stage(lancelotagregationuseFormCallback.probe.stageOfInterest)
	}
	lancelotagregationuse_ := lancelotagregationuseFormCallback.lancelotagregationuse
	_ = lancelotagregationuse_

	for _, formDiv := range lancelotagregationuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lancelotagregationuse_.Name), formDiv)
		case "ParameterAgregation":
			FormDivSelectFieldToField(&(lancelotagregationuse_.ParameterAgregation), lancelotagregationuseFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if lancelotagregationuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotagregationuse_.Unstage(lancelotagregationuseFormCallback.probe.stageOfInterest)
	}

	lancelotagregationuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LancelotAgregationUse](
		lancelotagregationuseFormCallback.probe,
	)
	lancelotagregationuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lancelotagregationuseFormCallback.CreationMode || lancelotagregationuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotagregationuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lancelotagregationuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LancelotAgregationUseFormCallback(
			nil,
			lancelotagregationuseFormCallback.probe,
			newFormGroup,
		)
		lancelotagregationuse := new(models.LancelotAgregationUse)
		FillUpForm(lancelotagregationuse, newFormGroup, lancelotagregationuseFormCallback.probe)
		lancelotagregationuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lancelotagregationuseFormCallback.probe)
}
func __gong__New__LancelotCategoryFormCallback(
	lancelotcategory *models.LancelotCategory,
	probe *Probe,
	formGroup *table.FormGroup,
) (lancelotcategoryFormCallback *LancelotCategoryFormCallback) {
	lancelotcategoryFormCallback = new(LancelotCategoryFormCallback)
	lancelotcategoryFormCallback.probe = probe
	lancelotcategoryFormCallback.lancelotcategory = lancelotcategory
	lancelotcategoryFormCallback.formGroup = formGroup

	lancelotcategoryFormCallback.CreationMode = (lancelotcategory == nil)

	return
}

type LancelotCategoryFormCallback struct {
	lancelotcategory *models.LancelotCategory

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lancelotcategoryFormCallback *LancelotCategoryFormCallback) OnSave() {

	log.Println("LancelotCategoryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lancelotcategoryFormCallback.probe.formStage.Checkout()

	if lancelotcategoryFormCallback.lancelotcategory == nil {
		lancelotcategoryFormCallback.lancelotcategory = new(models.LancelotCategory).Stage(lancelotcategoryFormCallback.probe.stageOfInterest)
	}
	lancelotcategory_ := lancelotcategoryFormCallback.lancelotcategory
	_ = lancelotcategory_

	for _, formDiv := range lancelotcategoryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lancelotcategory_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if lancelotcategoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotcategory_.Unstage(lancelotcategoryFormCallback.probe.stageOfInterest)
	}

	lancelotcategoryFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LancelotCategory](
		lancelotcategoryFormCallback.probe,
	)
	lancelotcategoryFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lancelotcategoryFormCallback.CreationMode || lancelotcategoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotcategoryFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lancelotcategoryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LancelotCategoryFormCallback(
			nil,
			lancelotcategoryFormCallback.probe,
			newFormGroup,
		)
		lancelotcategory := new(models.LancelotCategory)
		FillUpForm(lancelotcategory, newFormGroup, lancelotcategoryFormCallback.probe)
		lancelotcategoryFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lancelotcategoryFormCallback.probe)
}
func __gong__New__LancelotCategoryUseFormCallback(
	lancelotcategoryuse *models.LancelotCategoryUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (lancelotcategoryuseFormCallback *LancelotCategoryUseFormCallback) {
	lancelotcategoryuseFormCallback = new(LancelotCategoryUseFormCallback)
	lancelotcategoryuseFormCallback.probe = probe
	lancelotcategoryuseFormCallback.lancelotcategoryuse = lancelotcategoryuse
	lancelotcategoryuseFormCallback.formGroup = formGroup

	lancelotcategoryuseFormCallback.CreationMode = (lancelotcategoryuse == nil)

	return
}

type LancelotCategoryUseFormCallback struct {
	lancelotcategoryuse *models.LancelotCategoryUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lancelotcategoryuseFormCallback *LancelotCategoryUseFormCallback) OnSave() {

	log.Println("LancelotCategoryUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lancelotcategoryuseFormCallback.probe.formStage.Checkout()

	if lancelotcategoryuseFormCallback.lancelotcategoryuse == nil {
		lancelotcategoryuseFormCallback.lancelotcategoryuse = new(models.LancelotCategoryUse).Stage(lancelotcategoryuseFormCallback.probe.stageOfInterest)
	}
	lancelotcategoryuse_ := lancelotcategoryuseFormCallback.lancelotcategoryuse
	_ = lancelotcategoryuse_

	for _, formDiv := range lancelotcategoryuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(lancelotcategoryuse_.Name), formDiv)
		case "ParameterCategory":
			FormDivSelectFieldToField(&(lancelotcategoryuse_.ParameterCategory), lancelotcategoryuseFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if lancelotcategoryuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotcategoryuse_.Unstage(lancelotcategoryuseFormCallback.probe.stageOfInterest)
	}

	lancelotcategoryuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LancelotCategoryUse](
		lancelotcategoryuseFormCallback.probe,
	)
	lancelotcategoryuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lancelotcategoryuseFormCallback.CreationMode || lancelotcategoryuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lancelotcategoryuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lancelotcategoryuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LancelotCategoryUseFormCallback(
			nil,
			lancelotcategoryuseFormCallback.probe,
			newFormGroup,
		)
		lancelotcategoryuse := new(models.LancelotCategoryUse)
		FillUpForm(lancelotcategoryuse, newFormGroup, lancelotcategoryuseFormCallback.probe)
		lancelotcategoryuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lancelotcategoryuseFormCallback.probe)
}
func __gong__New__MapObjectFormCallback(
	mapobject *models.MapObject,
	probe *Probe,
	formGroup *table.FormGroup,
) (mapobjectFormCallback *MapObjectFormCallback) {
	mapobjectFormCallback = new(MapObjectFormCallback)
	mapobjectFormCallback.probe = probe
	mapobjectFormCallback.mapobject = mapobject
	mapobjectFormCallback.formGroup = formGroup

	mapobjectFormCallback.CreationMode = (mapobject == nil)

	return
}

type MapObjectFormCallback struct {
	mapobject *models.MapObject

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (mapobjectFormCallback *MapObjectFormCallback) OnSave() {

	log.Println("MapObjectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mapobjectFormCallback.probe.formStage.Checkout()

	if mapobjectFormCallback.mapobject == nil {
		mapobjectFormCallback.mapobject = new(models.MapObject).Stage(mapobjectFormCallback.probe.stageOfInterest)
	}
	mapobject_ := mapobjectFormCallback.mapobject
	_ = mapobject_

	for _, formDiv := range mapobjectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mapobject_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if mapobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapobject_.Unstage(mapobjectFormCallback.probe.stageOfInterest)
	}

	mapobjectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MapObject](
		mapobjectFormCallback.probe,
	)
	mapobjectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if mapobjectFormCallback.CreationMode || mapobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapobjectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(mapobjectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MapObjectFormCallback(
			nil,
			mapobjectFormCallback.probe,
			newFormGroup,
		)
		mapobject := new(models.MapObject)
		FillUpForm(mapobject, newFormGroup, mapobjectFormCallback.probe)
		mapobjectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(mapobjectFormCallback.probe)
}
func __gong__New__MapObjectUseFormCallback(
	mapobjectuse *models.MapObjectUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (mapobjectuseFormCallback *MapObjectUseFormCallback) {
	mapobjectuseFormCallback = new(MapObjectUseFormCallback)
	mapobjectuseFormCallback.probe = probe
	mapobjectuseFormCallback.mapobjectuse = mapobjectuse
	mapobjectuseFormCallback.formGroup = formGroup

	mapobjectuseFormCallback.CreationMode = (mapobjectuse == nil)

	return
}

type MapObjectUseFormCallback struct {
	mapobjectuse *models.MapObjectUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (mapobjectuseFormCallback *MapObjectUseFormCallback) OnSave() {

	log.Println("MapObjectUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mapobjectuseFormCallback.probe.formStage.Checkout()

	if mapobjectuseFormCallback.mapobjectuse == nil {
		mapobjectuseFormCallback.mapobjectuse = new(models.MapObjectUse).Stage(mapobjectuseFormCallback.probe.stageOfInterest)
	}
	mapobjectuse_ := mapobjectuseFormCallback.mapobjectuse
	_ = mapobjectuse_

	for _, formDiv := range mapobjectuseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mapobjectuse_.Name), formDiv)
		case "Map":
			FormDivSelectFieldToField(&(mapobjectuse_.Map), mapobjectuseFormCallback.probe.stageOfInterest, formDiv)
		case "TheBridge:MapUse":
			// we need to retrieve the field owner before the change
			var pastTheBridgeOwner *models.TheBridge
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TheBridge"
			rf.Fieldname = "MapUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				mapobjectuseFormCallback.probe.stageOfInterest,
				mapobjectuseFormCallback.probe.backRepoOfInterest,
				mapobjectuse_,
				&rf)

			if reverseFieldOwner != nil {
				pastTheBridgeOwner = reverseFieldOwner.(*models.TheBridge)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTheBridgeOwner != nil {
					idx := slices.Index(pastTheBridgeOwner.MapUse, mapobjectuse_)
					pastTheBridgeOwner.MapUse = slices.Delete(pastTheBridgeOwner.MapUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _thebridge := range *models.GetGongstructInstancesSet[models.TheBridge](mapobjectuseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _thebridge.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTheBridgeOwner := _thebridge // we have a match
						if pastTheBridgeOwner != nil {
							if newTheBridgeOwner != pastTheBridgeOwner {
								idx := slices.Index(pastTheBridgeOwner.MapUse, mapobjectuse_)
								pastTheBridgeOwner.MapUse = slices.Delete(pastTheBridgeOwner.MapUse, idx, idx+1)
								newTheBridgeOwner.MapUse = append(newTheBridgeOwner.MapUse, mapobjectuse_)
							}
						} else {
							newTheBridgeOwner.MapUse = append(newTheBridgeOwner.MapUse, mapobjectuse_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if mapobjectuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapobjectuse_.Unstage(mapobjectuseFormCallback.probe.stageOfInterest)
	}

	mapobjectuseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MapObjectUse](
		mapobjectuseFormCallback.probe,
	)
	mapobjectuseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if mapobjectuseFormCallback.CreationMode || mapobjectuseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mapobjectuseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(mapobjectuseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MapObjectUseFormCallback(
			nil,
			mapobjectuseFormCallback.probe,
			newFormGroup,
		)
		mapobjectuse := new(models.MapObjectUse)
		FillUpForm(mapobjectuse, newFormGroup, mapobjectuseFormCallback.probe)
		mapobjectuseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(mapobjectuseFormCallback.probe)
}
func __gong__New__RepositoryFormCallback(
	repository *models.Repository,
	probe *Probe,
	formGroup *table.FormGroup,
) (repositoryFormCallback *RepositoryFormCallback) {
	repositoryFormCallback = new(RepositoryFormCallback)
	repositoryFormCallback.probe = probe
	repositoryFormCallback.repository = repository
	repositoryFormCallback.formGroup = formGroup

	repositoryFormCallback.CreationMode = (repository == nil)

	return
}

type RepositoryFormCallback struct {
	repository *models.Repository

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (repositoryFormCallback *RepositoryFormCallback) OnSave() {

	log.Println("RepositoryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	repositoryFormCallback.probe.formStage.Checkout()

	if repositoryFormCallback.repository == nil {
		repositoryFormCallback.repository = new(models.Repository).Stage(repositoryFormCallback.probe.stageOfInterest)
	}
	repository_ := repositoryFormCallback.repository
	_ = repository_

	for _, formDiv := range repositoryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(repository_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if repositoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		repository_.Unstage(repositoryFormCallback.probe.stageOfInterest)
	}

	repositoryFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Repository](
		repositoryFormCallback.probe,
	)
	repositoryFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if repositoryFormCallback.CreationMode || repositoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		repositoryFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(repositoryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RepositoryFormCallback(
			nil,
			repositoryFormCallback.probe,
			newFormGroup,
		)
		repository := new(models.Repository)
		FillUpForm(repository, newFormGroup, repositoryFormCallback.probe)
		repositoryFormCallback.probe.formStage.Commit()
	}

	fillUpTree(repositoryFormCallback.probe)
}
func __gong__New__SirRobinFormCallback(
	sirrobin *models.SirRobin,
	probe *Probe,
	formGroup *table.FormGroup,
) (sirrobinFormCallback *SirRobinFormCallback) {
	sirrobinFormCallback = new(SirRobinFormCallback)
	sirrobinFormCallback.probe = probe
	sirrobinFormCallback.sirrobin = sirrobin
	sirrobinFormCallback.formGroup = formGroup

	sirrobinFormCallback.CreationMode = (sirrobin == nil)

	return
}

type SirRobinFormCallback struct {
	sirrobin *models.SirRobin

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sirrobinFormCallback *SirRobinFormCallback) OnSave() {

	log.Println("SirRobinFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sirrobinFormCallback.probe.formStage.Checkout()

	if sirrobinFormCallback.sirrobin == nil {
		sirrobinFormCallback.sirrobin = new(models.SirRobin).Stage(sirrobinFormCallback.probe.stageOfInterest)
	}
	sirrobin_ := sirrobinFormCallback.sirrobin
	_ = sirrobin_

	for _, formDiv := range sirrobinFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sirrobin_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(sirrobin_.Description), formDiv)
		case "AxisOrign_X":
			FormDivBasicFieldToField(&(sirrobin_.AxisOrign_X), formDiv)
		case "AxisOrign_Y":
			FormDivBasicFieldToField(&(sirrobin_.AxisOrign_Y), formDiv)
		case "VerticalAxis_Top_Y":
			FormDivBasicFieldToField(&(sirrobin_.VerticalAxis_Top_Y), formDiv)
		case "VerticalAxis_Bottom_Y":
			FormDivBasicFieldToField(&(sirrobin_.VerticalAxis_Bottom_Y), formDiv)
		case "VerticalAxis_StrokeWidth":
			FormDivBasicFieldToField(&(sirrobin_.VerticalAxis_StrokeWidth), formDiv)
		case "HorizontalAxis_Right_X":
			FormDivBasicFieldToField(&(sirrobin_.HorizontalAxis_Right_X), formDiv)
		case "Start":
			FormDivBasicFieldToField(&(sirrobin_.Start), formDiv)
		case "End":
			FormDivBasicFieldToField(&(sirrobin_.End), formDiv)
		case "WhatIsYourPreferedColor:Diagrams":
			// we need to retrieve the field owner before the change
			var pastWhatIsYourPreferedColorOwner *models.WhatIsYourPreferedColor
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "Diagrams"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sirrobinFormCallback.probe.stageOfInterest,
				sirrobinFormCallback.probe.backRepoOfInterest,
				sirrobin_,
				&rf)

			if reverseFieldOwner != nil {
				pastWhatIsYourPreferedColorOwner = reverseFieldOwner.(*models.WhatIsYourPreferedColor)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastWhatIsYourPreferedColorOwner != nil {
					idx := slices.Index(pastWhatIsYourPreferedColorOwner.Diagrams, sirrobin_)
					pastWhatIsYourPreferedColorOwner.Diagrams = slices.Delete(pastWhatIsYourPreferedColorOwner.Diagrams, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _whatisyourpreferedcolor := range *models.GetGongstructInstancesSet[models.WhatIsYourPreferedColor](sirrobinFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _whatisyourpreferedcolor.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newWhatIsYourPreferedColorOwner := _whatisyourpreferedcolor // we have a match
						if pastWhatIsYourPreferedColorOwner != nil {
							if newWhatIsYourPreferedColorOwner != pastWhatIsYourPreferedColorOwner {
								idx := slices.Index(pastWhatIsYourPreferedColorOwner.Diagrams, sirrobin_)
								pastWhatIsYourPreferedColorOwner.Diagrams = slices.Delete(pastWhatIsYourPreferedColorOwner.Diagrams, idx, idx+1)
								newWhatIsYourPreferedColorOwner.Diagrams = append(newWhatIsYourPreferedColorOwner.Diagrams, sirrobin_)
							}
						} else {
							newWhatIsYourPreferedColorOwner.Diagrams = append(newWhatIsYourPreferedColorOwner.Diagrams, sirrobin_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if sirrobinFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sirrobin_.Unstage(sirrobinFormCallback.probe.stageOfInterest)
	}

	sirrobinFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SirRobin](
		sirrobinFormCallback.probe,
	)
	sirrobinFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sirrobinFormCallback.CreationMode || sirrobinFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sirrobinFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(sirrobinFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SirRobinFormCallback(
			nil,
			sirrobinFormCallback.probe,
			newFormGroup,
		)
		sirrobin := new(models.SirRobin)
		FillUpForm(sirrobin, newFormGroup, sirrobinFormCallback.probe)
		sirrobinFormCallback.probe.formStage.Commit()
	}

	fillUpTree(sirrobinFormCallback.probe)
}
func __gong__New__TheBridgeFormCallback(
	thebridge *models.TheBridge,
	probe *Probe,
	formGroup *table.FormGroup,
) (thebridgeFormCallback *TheBridgeFormCallback) {
	thebridgeFormCallback = new(TheBridgeFormCallback)
	thebridgeFormCallback.probe = probe
	thebridgeFormCallback.thebridge = thebridge
	thebridgeFormCallback.formGroup = formGroup

	thebridgeFormCallback.CreationMode = (thebridge == nil)

	return
}

type TheBridgeFormCallback struct {
	thebridge *models.TheBridge

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (thebridgeFormCallback *TheBridgeFormCallback) OnSave() {

	log.Println("TheBridgeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	thebridgeFormCallback.probe.formStage.Checkout()

	if thebridgeFormCallback.thebridge == nil {
		thebridgeFormCallback.thebridge = new(models.TheBridge).Stage(thebridgeFormCallback.probe.stageOfInterest)
	}
	thebridge_ := thebridgeFormCallback.thebridge
	_ = thebridge_

	for _, formDiv := range thebridgeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(thebridge_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(thebridge_.Description), formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&(thebridge_.IsNodeExpanded), formDiv)
		}
	}

	// manage the suppress operation
	if thebridgeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		thebridge_.Unstage(thebridgeFormCallback.probe.stageOfInterest)
	}

	thebridgeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TheBridge](
		thebridgeFormCallback.probe,
	)
	thebridgeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if thebridgeFormCallback.CreationMode || thebridgeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		thebridgeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(thebridgeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TheBridgeFormCallback(
			nil,
			thebridgeFormCallback.probe,
			newFormGroup,
		)
		thebridge := new(models.TheBridge)
		FillUpForm(thebridge, newFormGroup, thebridgeFormCallback.probe)
		thebridgeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(thebridgeFormCallback.probe)
}
func __gong__New__TheNuteShapeFormCallback(
	thenuteshape *models.TheNuteShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (thenuteshapeFormCallback *TheNuteShapeFormCallback) {
	thenuteshapeFormCallback = new(TheNuteShapeFormCallback)
	thenuteshapeFormCallback.probe = probe
	thenuteshapeFormCallback.thenuteshape = thenuteshape
	thenuteshapeFormCallback.formGroup = formGroup

	thenuteshapeFormCallback.CreationMode = (thenuteshape == nil)

	return
}

type TheNuteShapeFormCallback struct {
	thenuteshape *models.TheNuteShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (thenuteshapeFormCallback *TheNuteShapeFormCallback) OnSave() {

	log.Println("TheNuteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	thenuteshapeFormCallback.probe.formStage.Checkout()

	if thenuteshapeFormCallback.thenuteshape == nil {
		thenuteshapeFormCallback.thenuteshape = new(models.TheNuteShape).Stage(thenuteshapeFormCallback.probe.stageOfInterest)
	}
	thenuteshape_ := thenuteshapeFormCallback.thenuteshape
	_ = thenuteshape_

	for _, formDiv := range thenuteshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(thenuteshape_.Name), formDiv)
		case "ActorStateTransition":
			FormDivSelectFieldToField(&(thenuteshape_.ActorStateTransition), thenuteshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Start":
			FormDivSelectFieldToField(&(thenuteshape_.Start), thenuteshapeFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(thenuteshape_.End), thenuteshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(thenuteshape_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(thenuteshape_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(thenuteshape_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(thenuteshape_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(thenuteshape_.CornerOffsetRatio), formDiv)
		case "SirRobin:TheNuteShapes":
			// we need to retrieve the field owner before the change
			var pastSirRobinOwner *models.SirRobin
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SirRobin"
			rf.Fieldname = "TheNuteShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				thenuteshapeFormCallback.probe.stageOfInterest,
				thenuteshapeFormCallback.probe.backRepoOfInterest,
				thenuteshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastSirRobinOwner = reverseFieldOwner.(*models.SirRobin)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSirRobinOwner != nil {
					idx := slices.Index(pastSirRobinOwner.TheNuteShapes, thenuteshape_)
					pastSirRobinOwner.TheNuteShapes = slices.Delete(pastSirRobinOwner.TheNuteShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sirrobin := range *models.GetGongstructInstancesSet[models.SirRobin](thenuteshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sirrobin.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSirRobinOwner := _sirrobin // we have a match
						if pastSirRobinOwner != nil {
							if newSirRobinOwner != pastSirRobinOwner {
								idx := slices.Index(pastSirRobinOwner.TheNuteShapes, thenuteshape_)
								pastSirRobinOwner.TheNuteShapes = slices.Delete(pastSirRobinOwner.TheNuteShapes, idx, idx+1)
								newSirRobinOwner.TheNuteShapes = append(newSirRobinOwner.TheNuteShapes, thenuteshape_)
							}
						} else {
							newSirRobinOwner.TheNuteShapes = append(newSirRobinOwner.TheNuteShapes, thenuteshape_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if thenuteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		thenuteshape_.Unstage(thenuteshapeFormCallback.probe.stageOfInterest)
	}

	thenuteshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TheNuteShape](
		thenuteshapeFormCallback.probe,
	)
	thenuteshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if thenuteshapeFormCallback.CreationMode || thenuteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		thenuteshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(thenuteshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TheNuteShapeFormCallback(
			nil,
			thenuteshapeFormCallback.probe,
			newFormGroup,
		)
		thenuteshape := new(models.TheNuteShape)
		FillUpForm(thenuteshape, newFormGroup, thenuteshapeFormCallback.probe)
		thenuteshapeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(thenuteshapeFormCallback.probe)
}
func __gong__New__TheNuteTransitionFormCallback(
	thenutetransition *models.TheNuteTransition,
	probe *Probe,
	formGroup *table.FormGroup,
) (thenutetransitionFormCallback *TheNuteTransitionFormCallback) {
	thenutetransitionFormCallback = new(TheNuteTransitionFormCallback)
	thenutetransitionFormCallback.probe = probe
	thenutetransitionFormCallback.thenutetransition = thenutetransition
	thenutetransitionFormCallback.formGroup = formGroup

	thenutetransitionFormCallback.CreationMode = (thenutetransition == nil)

	return
}

type TheNuteTransitionFormCallback struct {
	thenutetransition *models.TheNuteTransition

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (thenutetransitionFormCallback *TheNuteTransitionFormCallback) OnSave() {

	log.Println("TheNuteTransitionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	thenutetransitionFormCallback.probe.formStage.Checkout()

	if thenutetransitionFormCallback.thenutetransition == nil {
		thenutetransitionFormCallback.thenutetransition = new(models.TheNuteTransition).Stage(thenutetransitionFormCallback.probe.stageOfInterest)
	}
	thenutetransition_ := thenutetransitionFormCallback.thenutetransition
	_ = thenutetransition_

	for _, formDiv := range thenutetransitionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(thenutetransition_.Name), formDiv)
		case "StartState":
			FormDivSelectFieldToField(&(thenutetransition_.StartState), thenutetransitionFormCallback.probe.stageOfInterest, formDiv)
		case "EndState":
			FormDivSelectFieldToField(&(thenutetransition_.EndState), thenutetransitionFormCallback.probe.stageOfInterest, formDiv)
		case "WhatIsYourPreferedColor:Nutes":
			// we need to retrieve the field owner before the change
			var pastWhatIsYourPreferedColorOwner *models.WhatIsYourPreferedColor
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "WhatIsYourPreferedColor"
			rf.Fieldname = "Nutes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				thenutetransitionFormCallback.probe.stageOfInterest,
				thenutetransitionFormCallback.probe.backRepoOfInterest,
				thenutetransition_,
				&rf)

			if reverseFieldOwner != nil {
				pastWhatIsYourPreferedColorOwner = reverseFieldOwner.(*models.WhatIsYourPreferedColor)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastWhatIsYourPreferedColorOwner != nil {
					idx := slices.Index(pastWhatIsYourPreferedColorOwner.Nutes, thenutetransition_)
					pastWhatIsYourPreferedColorOwner.Nutes = slices.Delete(pastWhatIsYourPreferedColorOwner.Nutes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _whatisyourpreferedcolor := range *models.GetGongstructInstancesSet[models.WhatIsYourPreferedColor](thenutetransitionFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _whatisyourpreferedcolor.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newWhatIsYourPreferedColorOwner := _whatisyourpreferedcolor // we have a match
						if pastWhatIsYourPreferedColorOwner != nil {
							if newWhatIsYourPreferedColorOwner != pastWhatIsYourPreferedColorOwner {
								idx := slices.Index(pastWhatIsYourPreferedColorOwner.Nutes, thenutetransition_)
								pastWhatIsYourPreferedColorOwner.Nutes = slices.Delete(pastWhatIsYourPreferedColorOwner.Nutes, idx, idx+1)
								newWhatIsYourPreferedColorOwner.Nutes = append(newWhatIsYourPreferedColorOwner.Nutes, thenutetransition_)
							}
						} else {
							newWhatIsYourPreferedColorOwner.Nutes = append(newWhatIsYourPreferedColorOwner.Nutes, thenutetransition_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if thenutetransitionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		thenutetransition_.Unstage(thenutetransitionFormCallback.probe.stageOfInterest)
	}

	thenutetransitionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TheNuteTransition](
		thenutetransitionFormCallback.probe,
	)
	thenutetransitionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if thenutetransitionFormCallback.CreationMode || thenutetransitionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		thenutetransitionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(thenutetransitionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TheNuteTransitionFormCallback(
			nil,
			thenutetransitionFormCallback.probe,
			newFormGroup,
		)
		thenutetransition := new(models.TheNuteTransition)
		FillUpForm(thenutetransition, newFormGroup, thenutetransitionFormCallback.probe)
		thenutetransitionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(thenutetransitionFormCallback.probe)
}
func __gong__New__UserFormCallback(
	user *models.User,
	probe *Probe,
	formGroup *table.FormGroup,
) (userFormCallback *UserFormCallback) {
	userFormCallback = new(UserFormCallback)
	userFormCallback.probe = probe
	userFormCallback.user = user
	userFormCallback.formGroup = formGroup

	userFormCallback.CreationMode = (user == nil)

	return
}

type UserFormCallback struct {
	user *models.User

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (userFormCallback *UserFormCallback) OnSave() {

	log.Println("UserFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	userFormCallback.probe.formStage.Checkout()

	if userFormCallback.user == nil {
		userFormCallback.user = new(models.User).Stage(userFormCallback.probe.stageOfInterest)
	}
	user_ := userFormCallback.user
	_ = user_

	for _, formDiv := range userFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(user_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if userFormCallback.formGroup.HasSuppressButtonBeenPressed {
		user_.Unstage(userFormCallback.probe.stageOfInterest)
	}

	userFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.User](
		userFormCallback.probe,
	)
	userFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if userFormCallback.CreationMode || userFormCallback.formGroup.HasSuppressButtonBeenPressed {
		userFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(userFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UserFormCallback(
			nil,
			userFormCallback.probe,
			newFormGroup,
		)
		user := new(models.User)
		FillUpForm(user, newFormGroup, userFormCallback.probe)
		userFormCallback.probe.formStage.Commit()
	}

	fillUpTree(userFormCallback.probe)
}
func __gong__New__UserUseFormCallback(
	useruse *models.UserUse,
	probe *Probe,
	formGroup *table.FormGroup,
) (useruseFormCallback *UserUseFormCallback) {
	useruseFormCallback = new(UserUseFormCallback)
	useruseFormCallback.probe = probe
	useruseFormCallback.useruse = useruse
	useruseFormCallback.formGroup = formGroup

	useruseFormCallback.CreationMode = (useruse == nil)

	return
}

type UserUseFormCallback struct {
	useruse *models.UserUse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (useruseFormCallback *UserUseFormCallback) OnSave() {

	log.Println("UserUseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	useruseFormCallback.probe.formStage.Checkout()

	if useruseFormCallback.useruse == nil {
		useruseFormCallback.useruse = new(models.UserUse).Stage(useruseFormCallback.probe.stageOfInterest)
	}
	useruse_ := useruseFormCallback.useruse
	_ = useruse_

	for _, formDiv := range useruseFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(useruse_.Name), formDiv)
		case "User":
			FormDivSelectFieldToField(&(useruse_.User), useruseFormCallback.probe.stageOfInterest, formDiv)
		case "Group:UserUse":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "UserUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				useruseFormCallback.probe.stageOfInterest,
				useruseFormCallback.probe.backRepoOfInterest,
				useruse_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.UserUse, useruse_)
					pastGroupOwner.UserUse = slices.Delete(pastGroupOwner.UserUse, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](useruseFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.UserUse, useruse_)
								pastGroupOwner.UserUse = slices.Delete(pastGroupOwner.UserUse, idx, idx+1)
								newGroupOwner.UserUse = append(newGroupOwner.UserUse, useruse_)
							}
						} else {
							newGroupOwner.UserUse = append(newGroupOwner.UserUse, useruse_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if useruseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		useruse_.Unstage(useruseFormCallback.probe.stageOfInterest)
	}

	useruseFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.UserUse](
		useruseFormCallback.probe,
	)
	useruseFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if useruseFormCallback.CreationMode || useruseFormCallback.formGroup.HasSuppressButtonBeenPressed {
		useruseFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(useruseFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UserUseFormCallback(
			nil,
			useruseFormCallback.probe,
			newFormGroup,
		)
		useruse := new(models.UserUse)
		FillUpForm(useruse, newFormGroup, useruseFormCallback.probe)
		useruseFormCallback.probe.formStage.Commit()
	}

	fillUpTree(useruseFormCallback.probe)
}
func __gong__New__WhatIsYourPreferedColorFormCallback(
	whatisyourpreferedcolor *models.WhatIsYourPreferedColor,
	probe *Probe,
	formGroup *table.FormGroup,
) (whatisyourpreferedcolorFormCallback *WhatIsYourPreferedColorFormCallback) {
	whatisyourpreferedcolorFormCallback = new(WhatIsYourPreferedColorFormCallback)
	whatisyourpreferedcolorFormCallback.probe = probe
	whatisyourpreferedcolorFormCallback.whatisyourpreferedcolor = whatisyourpreferedcolor
	whatisyourpreferedcolorFormCallback.formGroup = formGroup

	whatisyourpreferedcolorFormCallback.CreationMode = (whatisyourpreferedcolor == nil)

	return
}

type WhatIsYourPreferedColorFormCallback struct {
	whatisyourpreferedcolor *models.WhatIsYourPreferedColor

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (whatisyourpreferedcolorFormCallback *WhatIsYourPreferedColorFormCallback) OnSave() {

	log.Println("WhatIsYourPreferedColorFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	whatisyourpreferedcolorFormCallback.probe.formStage.Checkout()

	if whatisyourpreferedcolorFormCallback.whatisyourpreferedcolor == nil {
		whatisyourpreferedcolorFormCallback.whatisyourpreferedcolor = new(models.WhatIsYourPreferedColor).Stage(whatisyourpreferedcolorFormCallback.probe.stageOfInterest)
	}
	whatisyourpreferedcolor_ := whatisyourpreferedcolorFormCallback.whatisyourpreferedcolor
	_ = whatisyourpreferedcolor_

	for _, formDiv := range whatisyourpreferedcolorFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.Description), formDiv)
		case "DiagramsNodeFolded":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.DiagramsNodeFolded), formDiv)
		case "KingArthurNodeFolded":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.KingArthurNodeFolded), formDiv)
		case "RRRR":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.RRRR), formDiv)
		case "IIUU":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.IIUU), formDiv)
		case "LancelotsodeFolded":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.LancelotsodeFolded), formDiv)
		case "RRRRT":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.RRRRT), formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&(whatisyourpreferedcolor_.IsNodeExpanded), formDiv)
		case "TheBridge:WhatIsYourPreferedColor":
			// we need to retrieve the field owner before the change
			var pastTheBridgeOwner *models.TheBridge
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TheBridge"
			rf.Fieldname = "WhatIsYourPreferedColor"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				whatisyourpreferedcolorFormCallback.probe.stageOfInterest,
				whatisyourpreferedcolorFormCallback.probe.backRepoOfInterest,
				whatisyourpreferedcolor_,
				&rf)

			if reverseFieldOwner != nil {
				pastTheBridgeOwner = reverseFieldOwner.(*models.TheBridge)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTheBridgeOwner != nil {
					idx := slices.Index(pastTheBridgeOwner.WhatIsYourPreferedColor, whatisyourpreferedcolor_)
					pastTheBridgeOwner.WhatIsYourPreferedColor = slices.Delete(pastTheBridgeOwner.WhatIsYourPreferedColor, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _thebridge := range *models.GetGongstructInstancesSet[models.TheBridge](whatisyourpreferedcolorFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _thebridge.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTheBridgeOwner := _thebridge // we have a match
						if pastTheBridgeOwner != nil {
							if newTheBridgeOwner != pastTheBridgeOwner {
								idx := slices.Index(pastTheBridgeOwner.WhatIsYourPreferedColor, whatisyourpreferedcolor_)
								pastTheBridgeOwner.WhatIsYourPreferedColor = slices.Delete(pastTheBridgeOwner.WhatIsYourPreferedColor, idx, idx+1)
								newTheBridgeOwner.WhatIsYourPreferedColor = append(newTheBridgeOwner.WhatIsYourPreferedColor, whatisyourpreferedcolor_)
							}
						} else {
							newTheBridgeOwner.WhatIsYourPreferedColor = append(newTheBridgeOwner.WhatIsYourPreferedColor, whatisyourpreferedcolor_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if whatisyourpreferedcolorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		whatisyourpreferedcolor_.Unstage(whatisyourpreferedcolorFormCallback.probe.stageOfInterest)
	}

	whatisyourpreferedcolorFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.WhatIsYourPreferedColor](
		whatisyourpreferedcolorFormCallback.probe,
	)
	whatisyourpreferedcolorFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if whatisyourpreferedcolorFormCallback.CreationMode || whatisyourpreferedcolorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		whatisyourpreferedcolorFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(whatisyourpreferedcolorFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WhatIsYourPreferedColorFormCallback(
			nil,
			whatisyourpreferedcolorFormCallback.probe,
			newFormGroup,
		)
		whatisyourpreferedcolor := new(models.WhatIsYourPreferedColor)
		FillUpForm(whatisyourpreferedcolor, newFormGroup, whatisyourpreferedcolorFormCallback.probe)
		whatisyourpreferedcolorFormCallback.probe.formStage.Commit()
	}

	fillUpTree(whatisyourpreferedcolorFormCallback.probe)
}
func __gong__New__WorkspaceFormCallback(
	workspace *models.Workspace,
	probe *Probe,
	formGroup *table.FormGroup,
) (workspaceFormCallback *WorkspaceFormCallback) {
	workspaceFormCallback = new(WorkspaceFormCallback)
	workspaceFormCallback.probe = probe
	workspaceFormCallback.workspace = workspace
	workspaceFormCallback.formGroup = formGroup

	workspaceFormCallback.CreationMode = (workspace == nil)

	return
}

type WorkspaceFormCallback struct {
	workspace *models.Workspace

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (workspaceFormCallback *WorkspaceFormCallback) OnSave() {

	log.Println("WorkspaceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	workspaceFormCallback.probe.formStage.Checkout()

	if workspaceFormCallback.workspace == nil {
		workspaceFormCallback.workspace = new(models.Workspace).Stage(workspaceFormCallback.probe.stageOfInterest)
	}
	workspace_ := workspaceFormCallback.workspace
	_ = workspace_

	for _, formDiv := range workspaceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(workspace_.Name), formDiv)
		case "SelectedDiagram":
			FormDivSelectFieldToField(&(workspace_.SelectedDiagram), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "A":
			FormDivSelectFieldToField(&(workspace_.A), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "B":
			FormDivSelectFieldToField(&(workspace_.B), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "C":
			FormDivSelectFieldToField(&(workspace_.C), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "D":
			FormDivSelectFieldToField(&(workspace_.D), workspaceFormCallback.probe.stageOfInterest, formDiv)
		case "E":
			FormDivSelectFieldToField(&(workspace_.E), workspaceFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if workspaceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		workspace_.Unstage(workspaceFormCallback.probe.stageOfInterest)
	}

	workspaceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Workspace](
		workspaceFormCallback.probe,
	)
	workspaceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if workspaceFormCallback.CreationMode || workspaceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		workspaceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(workspaceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WorkspaceFormCallback(
			nil,
			workspaceFormCallback.probe,
			newFormGroup,
		)
		workspace := new(models.Workspace)
		FillUpForm(workspace, newFormGroup, workspaceFormCallback.probe)
		workspaceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(workspaceFormCallback.probe)
}
