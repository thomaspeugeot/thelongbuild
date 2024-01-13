// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_{{databaseName}} models.StageStruct
var ___dummy__Time_{{databaseName}} time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_{{databaseName}} map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["{{databaseName}}"] = {{databaseName}}Injection
// }

// {{databaseName}}Injection will stage objects of database "{{databaseName}}"
func {{databaseName}}Injection(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}

`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField 

	// insertion initialization of objects to stage
	map_AWitch_Identifiers := make(map[*AWitch]string)
	_ = map_AWitch_Identifiers

	awitchOrdered := []*AWitch{}
	for awitch := range stage.AWitchs {
		awitchOrdered = append(awitchOrdered, awitch)
	}
	sort.Slice(awitchOrdered[:], func(i, j int) bool {
		return awitchOrdered[i].Name < awitchOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of AWitch"
	for idx, awitch := range awitchOrdered {

		id = generatesIdentifier("AWitch", idx, awitch.Name)
		map_AWitch_Identifiers[awitch] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AWitch")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", awitch.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// AWitch values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(awitch.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", awitch.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", awitch.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", awitch.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", awitch.Height))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(awitch.FillColor))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", awitch.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(awitch.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", awitch.StrokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", awitch.RX))
		initializerStatements += setValueField

	}

	map_BlackKnightShape_Identifiers := make(map[*BlackKnightShape]string)
	_ = map_BlackKnightShape_Identifiers

	blackknightshapeOrdered := []*BlackKnightShape{}
	for blackknightshape := range stage.BlackKnightShapes {
		blackknightshapeOrdered = append(blackknightshapeOrdered, blackknightshape)
	}
	sort.Slice(blackknightshapeOrdered[:], func(i, j int) bool {
		return blackknightshapeOrdered[i].Name < blackknightshapeOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of BlackKnightShape"
	for idx, blackknightshape := range blackknightshapeOrdered {

		id = generatesIdentifier("BlackKnightShape", idx, blackknightshape.Name)
		map_BlackKnightShape_Identifiers[blackknightshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BlackKnightShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", blackknightshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// BlackKnightShape values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(blackknightshape.Name))
		initializerStatements += setValueField

		if blackknightshape.Direction != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Direction")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+blackknightshape.Direction.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", blackknightshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", blackknightshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", blackknightshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", blackknightshape.Height))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(blackknightshape.FillColor))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", blackknightshape.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(blackknightshape.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", blackknightshape.StrokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", blackknightshape.RX))
		initializerStatements += setValueField

	}

	map_BringYourDead_Identifiers := make(map[*BringYourDead]string)
	_ = map_BringYourDead_Identifiers

	bringyourdeadOrdered := []*BringYourDead{}
	for bringyourdead := range stage.BringYourDeads {
		bringyourdeadOrdered = append(bringyourdeadOrdered, bringyourdead)
	}
	sort.Slice(bringyourdeadOrdered[:], func(i, j int) bool {
		return bringyourdeadOrdered[i].Name < bringyourdeadOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of BringYourDead"
	for idx, bringyourdead := range bringyourdeadOrdered {

		id = generatesIdentifier("BringYourDead", idx, bringyourdead.Name)
		map_BringYourDead_Identifiers[bringyourdead] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BringYourDead")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bringyourdead.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// BringYourDead values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bringyourdead.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Tag")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bringyourdead.Tag))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bringyourdead.Description))
		initializerStatements += setValueField

	}

	map_Document_Identifiers := make(map[*Document]string)
	_ = map_Document_Identifiers

	documentOrdered := []*Document{}
	for document := range stage.Documents {
		documentOrdered = append(documentOrdered, document)
	}
	sort.Slice(documentOrdered[:], func(i, j int) bool {
		return documentOrdered[i].Name < documentOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of Document"
	for idx, document := range documentOrdered {

		id = generatesIdentifier("Document", idx, document.Name)
		map_Document_Identifiers[document] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Document")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", document.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// Document values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(document.Name))
		initializerStatements += setValueField

	}

	map_DocumentUse_Identifiers := make(map[*DocumentUse]string)
	_ = map_DocumentUse_Identifiers

	documentuseOrdered := []*DocumentUse{}
	for documentuse := range stage.DocumentUses {
		documentuseOrdered = append(documentuseOrdered, documentuse)
	}
	sort.Slice(documentuseOrdered[:], func(i, j int) bool {
		return documentuseOrdered[i].Name < documentuseOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of DocumentUse"
	for idx, documentuse := range documentuseOrdered {

		id = generatesIdentifier("DocumentUse", idx, documentuse.Name)
		map_DocumentUse_Identifiers[documentuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DocumentUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", documentuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// DocumentUse values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentuse.Name))
		initializerStatements += setValueField

	}

	map_GalahadThePure_Identifiers := make(map[*GalahadThePure]string)
	_ = map_GalahadThePure_Identifiers

	galahadthepureOrdered := []*GalahadThePure{}
	for galahadthepure := range stage.GalahadThePures {
		galahadthepureOrdered = append(galahadthepureOrdered, galahadthepure)
	}
	sort.Slice(galahadthepureOrdered[:], func(i, j int) bool {
		return galahadthepureOrdered[i].Name < galahadthepureOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of GalahadThePure"
	for idx, galahadthepure := range galahadthepureOrdered {

		id = generatesIdentifier("GalahadThePure", idx, galahadthepure.Name)
		map_GalahadThePure_Identifiers[galahadthepure] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GalahadThePure")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", galahadthepure.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// GalahadThePure values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(galahadthepure.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(galahadthepure.Description))
		initializerStatements += setValueField

	}

	map_GeoObject_Identifiers := make(map[*GeoObject]string)
	_ = map_GeoObject_Identifiers

	geoobjectOrdered := []*GeoObject{}
	for geoobject := range stage.GeoObjects {
		geoobjectOrdered = append(geoobjectOrdered, geoobject)
	}
	sort.Slice(geoobjectOrdered[:], func(i, j int) bool {
		return geoobjectOrdered[i].Name < geoobjectOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of GeoObject"
	for idx, geoobject := range geoobjectOrdered {

		id = generatesIdentifier("GeoObject", idx, geoobject.Name)
		map_GeoObject_Identifiers[geoobject] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GeoObject")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", geoobject.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// GeoObject values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(geoobject.Name))
		initializerStatements += setValueField

	}

	map_GeoObjectUse_Identifiers := make(map[*GeoObjectUse]string)
	_ = map_GeoObjectUse_Identifiers

	geoobjectuseOrdered := []*GeoObjectUse{}
	for geoobjectuse := range stage.GeoObjectUses {
		geoobjectuseOrdered = append(geoobjectuseOrdered, geoobjectuse)
	}
	sort.Slice(geoobjectuseOrdered[:], func(i, j int) bool {
		return geoobjectuseOrdered[i].Name < geoobjectuseOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of GeoObjectUse"
	for idx, geoobjectuse := range geoobjectuseOrdered {

		id = generatesIdentifier("GeoObjectUse", idx, geoobjectuse.Name)
		map_GeoObjectUse_Identifiers[geoobjectuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GeoObjectUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", geoobjectuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// GeoObjectUse values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(geoobjectuse.Name))
		initializerStatements += setValueField

	}

	map_Group_Identifiers := make(map[*Group]string)
	_ = map_Group_Identifiers

	groupOrdered := []*Group{}
	for group := range stage.Groups {
		groupOrdered = append(groupOrdered, group)
	}
	sort.Slice(groupOrdered[:], func(i, j int) bool {
		return groupOrdered[i].Name < groupOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of Group"
	for idx, group := range groupOrdered {

		id = generatesIdentifier("Group", idx, group.Name)
		map_Group_Identifiers[group] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// Group values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.Name))
		initializerStatements += setValueField

	}

	map_GroupUse_Identifiers := make(map[*GroupUse]string)
	_ = map_GroupUse_Identifiers

	groupuseOrdered := []*GroupUse{}
	for groupuse := range stage.GroupUses {
		groupuseOrdered = append(groupuseOrdered, groupuse)
	}
	sort.Slice(groupuseOrdered[:], func(i, j int) bool {
		return groupuseOrdered[i].Name < groupuseOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of GroupUse"
	for idx, groupuse := range groupuseOrdered {

		id = generatesIdentifier("GroupUse", idx, groupuse.Name)
		map_GroupUse_Identifiers[groupuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GroupUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", groupuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// GroupUse values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(groupuse.Name))
		initializerStatements += setValueField

	}

	map_KingArthur_Identifiers := make(map[*KingArthur]string)
	_ = map_KingArthur_Identifiers

	kingarthurOrdered := []*KingArthur{}
	for kingarthur := range stage.KingArthurs {
		kingarthurOrdered = append(kingarthurOrdered, kingarthur)
	}
	sort.Slice(kingarthurOrdered[:], func(i, j int) bool {
		return kingarthurOrdered[i].Name < kingarthurOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of KingArthur"
	for idx, kingarthur := range kingarthurOrdered {

		id = generatesIdentifier("KingArthur", idx, kingarthur.Name)
		map_KingArthur_Identifiers[kingarthur] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "KingArthur")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", kingarthur.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// KingArthur values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kingarthur.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsWithProbaility")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", kingarthur.IsWithProbaility))
		initializerStatements += setValueField

		if kingarthur.Probability != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Probability")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+kingarthur.Probability.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_KingArthurShape_Identifiers := make(map[*KingArthurShape]string)
	_ = map_KingArthurShape_Identifiers

	kingarthurshapeOrdered := []*KingArthurShape{}
	for kingarthurshape := range stage.KingArthurShapes {
		kingarthurshapeOrdered = append(kingarthurshapeOrdered, kingarthurshape)
	}
	sort.Slice(kingarthurshapeOrdered[:], func(i, j int) bool {
		return kingarthurshapeOrdered[i].Name < kingarthurshapeOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of KingArthurShape"
	for idx, kingarthurshape := range kingarthurshapeOrdered {

		id = generatesIdentifier("KingArthurShape", idx, kingarthurshape.Name)
		map_KingArthurShape_Identifiers[kingarthurshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "KingArthurShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", kingarthurshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// KingArthurShape values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kingarthurshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", kingarthurshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", kingarthurshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", kingarthurshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", kingarthurshape.Height))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kingarthurshape.FillColor))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", kingarthurshape.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kingarthurshape.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", kingarthurshape.StrokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", kingarthurshape.RX))
		initializerStatements += setValueField

	}

	map_KnightWhoSayNi_Identifiers := make(map[*KnightWhoSayNi]string)
	_ = map_KnightWhoSayNi_Identifiers

	knightwhosayniOrdered := []*KnightWhoSayNi{}
	for knightwhosayni := range stage.KnightWhoSayNis {
		knightwhosayniOrdered = append(knightwhosayniOrdered, knightwhosayni)
	}
	sort.Slice(knightwhosayniOrdered[:], func(i, j int) bool {
		return knightwhosayniOrdered[i].Name < knightwhosayniOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of KnightWhoSayNi"
	for idx, knightwhosayni := range knightwhosayniOrdered {

		id = generatesIdentifier("KnightWhoSayNi", idx, knightwhosayni.Name)
		map_KnightWhoSayNi_Identifiers[knightwhosayni] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "KnightWhoSayNi")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", knightwhosayni.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// KnightWhoSayNi values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(knightwhosayni.Name))
		initializerStatements += setValueField

		if knightwhosayni.Direction != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Direction")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+knightwhosayni.Direction.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", knightwhosayni.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", knightwhosayni.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", knightwhosayni.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", knightwhosayni.Height))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(knightwhosayni.FillColor))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", knightwhosayni.FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(knightwhosayni.Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", knightwhosayni.StrokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", knightwhosayni.RX))
		initializerStatements += setValueField

	}

	map_Lancelot_Identifiers := make(map[*Lancelot]string)
	_ = map_Lancelot_Identifiers

	lancelotOrdered := []*Lancelot{}
	for lancelot := range stage.Lancelots {
		lancelotOrdered = append(lancelotOrdered, lancelot)
	}
	sort.Slice(lancelotOrdered[:], func(i, j int) bool {
		return lancelotOrdered[i].Name < lancelotOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of Lancelot"
	for idx, lancelot := range lancelotOrdered {

		id = generatesIdentifier("Lancelot", idx, lancelot.Name)
		map_Lancelot_Identifiers[lancelot] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lancelot")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lancelot.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// Lancelot values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lancelot.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Tag")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lancelot.Tag))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lancelot.Description))
		initializerStatements += setValueField

	}

	map_LancelotAgregation_Identifiers := make(map[*LancelotAgregation]string)
	_ = map_LancelotAgregation_Identifiers

	lancelotagregationOrdered := []*LancelotAgregation{}
	for lancelotagregation := range stage.LancelotAgregations {
		lancelotagregationOrdered = append(lancelotagregationOrdered, lancelotagregation)
	}
	sort.Slice(lancelotagregationOrdered[:], func(i, j int) bool {
		return lancelotagregationOrdered[i].Name < lancelotagregationOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of LancelotAgregation"
	for idx, lancelotagregation := range lancelotagregationOrdered {

		id = generatesIdentifier("LancelotAgregation", idx, lancelotagregation.Name)
		map_LancelotAgregation_Identifiers[lancelotagregation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LancelotAgregation")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lancelotagregation.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// LancelotAgregation values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lancelotagregation.Name))
		initializerStatements += setValueField

	}

	map_LancelotAgregationUse_Identifiers := make(map[*LancelotAgregationUse]string)
	_ = map_LancelotAgregationUse_Identifiers

	lancelotagregationuseOrdered := []*LancelotAgregationUse{}
	for lancelotagregationuse := range stage.LancelotAgregationUses {
		lancelotagregationuseOrdered = append(lancelotagregationuseOrdered, lancelotagregationuse)
	}
	sort.Slice(lancelotagregationuseOrdered[:], func(i, j int) bool {
		return lancelotagregationuseOrdered[i].Name < lancelotagregationuseOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of LancelotAgregationUse"
	for idx, lancelotagregationuse := range lancelotagregationuseOrdered {

		id = generatesIdentifier("LancelotAgregationUse", idx, lancelotagregationuse.Name)
		map_LancelotAgregationUse_Identifiers[lancelotagregationuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LancelotAgregationUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lancelotagregationuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// LancelotAgregationUse values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lancelotagregationuse.Name))
		initializerStatements += setValueField

	}

	map_LancelotCategory_Identifiers := make(map[*LancelotCategory]string)
	_ = map_LancelotCategory_Identifiers

	lancelotcategoryOrdered := []*LancelotCategory{}
	for lancelotcategory := range stage.LancelotCategorys {
		lancelotcategoryOrdered = append(lancelotcategoryOrdered, lancelotcategory)
	}
	sort.Slice(lancelotcategoryOrdered[:], func(i, j int) bool {
		return lancelotcategoryOrdered[i].Name < lancelotcategoryOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of LancelotCategory"
	for idx, lancelotcategory := range lancelotcategoryOrdered {

		id = generatesIdentifier("LancelotCategory", idx, lancelotcategory.Name)
		map_LancelotCategory_Identifiers[lancelotcategory] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LancelotCategory")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lancelotcategory.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// LancelotCategory values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lancelotcategory.Name))
		initializerStatements += setValueField

	}

	map_LancelotCategoryUse_Identifiers := make(map[*LancelotCategoryUse]string)
	_ = map_LancelotCategoryUse_Identifiers

	lancelotcategoryuseOrdered := []*LancelotCategoryUse{}
	for lancelotcategoryuse := range stage.LancelotCategoryUses {
		lancelotcategoryuseOrdered = append(lancelotcategoryuseOrdered, lancelotcategoryuse)
	}
	sort.Slice(lancelotcategoryuseOrdered[:], func(i, j int) bool {
		return lancelotcategoryuseOrdered[i].Name < lancelotcategoryuseOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of LancelotCategoryUse"
	for idx, lancelotcategoryuse := range lancelotcategoryuseOrdered {

		id = generatesIdentifier("LancelotCategoryUse", idx, lancelotcategoryuse.Name)
		map_LancelotCategoryUse_Identifiers[lancelotcategoryuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LancelotCategoryUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lancelotcategoryuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// LancelotCategoryUse values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lancelotcategoryuse.Name))
		initializerStatements += setValueField

	}

	map_MapObject_Identifiers := make(map[*MapObject]string)
	_ = map_MapObject_Identifiers

	mapobjectOrdered := []*MapObject{}
	for mapobject := range stage.MapObjects {
		mapobjectOrdered = append(mapobjectOrdered, mapobject)
	}
	sort.Slice(mapobjectOrdered[:], func(i, j int) bool {
		return mapobjectOrdered[i].Name < mapobjectOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of MapObject"
	for idx, mapobject := range mapobjectOrdered {

		id = generatesIdentifier("MapObject", idx, mapobject.Name)
		map_MapObject_Identifiers[mapobject] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MapObject")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", mapobject.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// MapObject values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mapobject.Name))
		initializerStatements += setValueField

	}

	map_MapObjectUse_Identifiers := make(map[*MapObjectUse]string)
	_ = map_MapObjectUse_Identifiers

	mapobjectuseOrdered := []*MapObjectUse{}
	for mapobjectuse := range stage.MapObjectUses {
		mapobjectuseOrdered = append(mapobjectuseOrdered, mapobjectuse)
	}
	sort.Slice(mapobjectuseOrdered[:], func(i, j int) bool {
		return mapobjectuseOrdered[i].Name < mapobjectuseOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of MapObjectUse"
	for idx, mapobjectuse := range mapobjectuseOrdered {

		id = generatesIdentifier("MapObjectUse", idx, mapobjectuse.Name)
		map_MapObjectUse_Identifiers[mapobjectuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MapObjectUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", mapobjectuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// MapObjectUse values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mapobjectuse.Name))
		initializerStatements += setValueField

	}

	map_Repository_Identifiers := make(map[*Repository]string)
	_ = map_Repository_Identifiers

	repositoryOrdered := []*Repository{}
	for repository := range stage.Repositorys {
		repositoryOrdered = append(repositoryOrdered, repository)
	}
	sort.Slice(repositoryOrdered[:], func(i, j int) bool {
		return repositoryOrdered[i].Name < repositoryOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of Repository"
	for idx, repository := range repositoryOrdered {

		id = generatesIdentifier("Repository", idx, repository.Name)
		map_Repository_Identifiers[repository] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Repository")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", repository.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// Repository values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(repository.Name))
		initializerStatements += setValueField

	}

	map_SirRobin_Identifiers := make(map[*SirRobin]string)
	_ = map_SirRobin_Identifiers

	sirrobinOrdered := []*SirRobin{}
	for sirrobin := range stage.SirRobins {
		sirrobinOrdered = append(sirrobinOrdered, sirrobin)
	}
	sort.Slice(sirrobinOrdered[:], func(i, j int) bool {
		return sirrobinOrdered[i].Name < sirrobinOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of SirRobin"
	for idx, sirrobin := range sirrobinOrdered {

		id = generatesIdentifier("SirRobin", idx, sirrobin.Name)
		map_SirRobin_Identifiers[sirrobin] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SirRobin")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", sirrobin.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// SirRobin values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sirrobin.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sirrobin.Description))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AxisOrign_X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", sirrobin.AxisOrign_X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AxisOrign_Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", sirrobin.AxisOrign_Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "VerticalAxis_Top_Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", sirrobin.VerticalAxis_Top_Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "VerticalAxis_Bottom_Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", sirrobin.VerticalAxis_Bottom_Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "VerticalAxis_StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", sirrobin.VerticalAxis_StrokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HorizontalAxis_Right_X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", sirrobin.HorizontalAxis_Right_X))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Start")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", sirrobin.Start.String())
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "End")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", sirrobin.End.String())
		initializerStatements += setValueField

	}

	map_TheBridge_Identifiers := make(map[*TheBridge]string)
	_ = map_TheBridge_Identifiers

	thebridgeOrdered := []*TheBridge{}
	for thebridge := range stage.TheBridges {
		thebridgeOrdered = append(thebridgeOrdered, thebridge)
	}
	sort.Slice(thebridgeOrdered[:], func(i, j int) bool {
		return thebridgeOrdered[i].Name < thebridgeOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of TheBridge"
	for idx, thebridge := range thebridgeOrdered {

		id = generatesIdentifier("TheBridge", idx, thebridge.Name)
		map_TheBridge_Identifiers[thebridge] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TheBridge")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", thebridge.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// TheBridge values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(thebridge.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(thebridge.Description))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", thebridge.IsNodeExpanded))
		initializerStatements += setValueField

	}

	map_TheNuteShape_Identifiers := make(map[*TheNuteShape]string)
	_ = map_TheNuteShape_Identifiers

	thenuteshapeOrdered := []*TheNuteShape{}
	for thenuteshape := range stage.TheNuteShapes {
		thenuteshapeOrdered = append(thenuteshapeOrdered, thenuteshape)
	}
	sort.Slice(thenuteshapeOrdered[:], func(i, j int) bool {
		return thenuteshapeOrdered[i].Name < thenuteshapeOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of TheNuteShape"
	for idx, thenuteshape := range thenuteshapeOrdered {

		id = generatesIdentifier("TheNuteShape", idx, thenuteshape.Name)
		map_TheNuteShape_Identifiers[thenuteshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TheNuteShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", thenuteshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// TheNuteShape values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(thenuteshape.Name))
		initializerStatements += setValueField

		if thenuteshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+thenuteshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", thenuteshape.StartRatio))
		initializerStatements += setValueField

		if thenuteshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+thenuteshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", thenuteshape.EndRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", thenuteshape.CornerOffsetRatio))
		initializerStatements += setValueField

	}

	map_TheNuteTransition_Identifiers := make(map[*TheNuteTransition]string)
	_ = map_TheNuteTransition_Identifiers

	thenutetransitionOrdered := []*TheNuteTransition{}
	for thenutetransition := range stage.TheNuteTransitions {
		thenutetransitionOrdered = append(thenutetransitionOrdered, thenutetransition)
	}
	sort.Slice(thenutetransitionOrdered[:], func(i, j int) bool {
		return thenutetransitionOrdered[i].Name < thenutetransitionOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of TheNuteTransition"
	for idx, thenutetransition := range thenutetransitionOrdered {

		id = generatesIdentifier("TheNuteTransition", idx, thenutetransition.Name)
		map_TheNuteTransition_Identifiers[thenutetransition] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TheNuteTransition")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", thenutetransition.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// TheNuteTransition values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(thenutetransition.Name))
		initializerStatements += setValueField

	}

	map_User_Identifiers := make(map[*User]string)
	_ = map_User_Identifiers

	userOrdered := []*User{}
	for user := range stage.Users {
		userOrdered = append(userOrdered, user)
	}
	sort.Slice(userOrdered[:], func(i, j int) bool {
		return userOrdered[i].Name < userOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of User"
	for idx, user := range userOrdered {

		id = generatesIdentifier("User", idx, user.Name)
		map_User_Identifiers[user] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "User")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", user.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// User values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(user.Name))
		initializerStatements += setValueField

	}

	map_UserUse_Identifiers := make(map[*UserUse]string)
	_ = map_UserUse_Identifiers

	useruseOrdered := []*UserUse{}
	for useruse := range stage.UserUses {
		useruseOrdered = append(useruseOrdered, useruse)
	}
	sort.Slice(useruseOrdered[:], func(i, j int) bool {
		return useruseOrdered[i].Name < useruseOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of UserUse"
	for idx, useruse := range useruseOrdered {

		id = generatesIdentifier("UserUse", idx, useruse.Name)
		map_UserUse_Identifiers[useruse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "UserUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", useruse.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// UserUse values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(useruse.Name))
		initializerStatements += setValueField

	}

	map_WhatIsYourPreferedColor_Identifiers := make(map[*WhatIsYourPreferedColor]string)
	_ = map_WhatIsYourPreferedColor_Identifiers

	whatisyourpreferedcolorOrdered := []*WhatIsYourPreferedColor{}
	for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
		whatisyourpreferedcolorOrdered = append(whatisyourpreferedcolorOrdered, whatisyourpreferedcolor)
	}
	sort.Slice(whatisyourpreferedcolorOrdered[:], func(i, j int) bool {
		return whatisyourpreferedcolorOrdered[i].Name < whatisyourpreferedcolorOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of WhatIsYourPreferedColor"
	for idx, whatisyourpreferedcolor := range whatisyourpreferedcolorOrdered {

		id = generatesIdentifier("WhatIsYourPreferedColor", idx, whatisyourpreferedcolor.Name)
		map_WhatIsYourPreferedColor_Identifiers[whatisyourpreferedcolor] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "WhatIsYourPreferedColor")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", whatisyourpreferedcolor.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// WhatIsYourPreferedColor values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(whatisyourpreferedcolor.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(whatisyourpreferedcolor.Description))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DiagramsNodeFolded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", whatisyourpreferedcolor.DiagramsNodeFolded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "KingArthurNodeFolded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", whatisyourpreferedcolor.KingArthurNodeFolded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RRRR")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", whatisyourpreferedcolor.RRRR))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IIUU")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", whatisyourpreferedcolor.IIUU))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LancelotsodeFolded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", whatisyourpreferedcolor.LancelotsodeFolded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RRRRT")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", whatisyourpreferedcolor.RRRRT))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", whatisyourpreferedcolor.IsNodeExpanded))
		initializerStatements += setValueField

	}

	map_Workspace_Identifiers := make(map[*Workspace]string)
	_ = map_Workspace_Identifiers

	workspaceOrdered := []*Workspace{}
	for workspace := range stage.Workspaces {
		workspaceOrdered = append(workspaceOrdered, workspace)
	}
	sort.Slice(workspaceOrdered[:], func(i, j int) bool {
		return workspaceOrdered[i].Name < workspaceOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of Workspace"
	for idx, workspace := range workspaceOrdered {

		id = generatesIdentifier("Workspace", idx, workspace.Name)
		map_Workspace_Identifiers[workspace] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Workspace")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", workspace.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// Workspace values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(workspace.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, awitch := range awitchOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AWitch", idx, awitch.Name)
		map_AWitch_Identifiers[awitch] = id

		// Initialisation of values
		if awitch.EvolutionDirection != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EvolutionDirection")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GalahadThePure_Identifiers[awitch.EvolutionDirection])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, blackknightshape := range blackknightshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("BlackKnightShape", idx, blackknightshape.Name)
		map_BlackKnightShape_Identifiers[blackknightshape] = id

		// Initialisation of values
		if blackknightshape.BringYourDead != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "BringYourDead")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BringYourDead_Identifiers[blackknightshape.BringYourDead])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, bringyourdead := range bringyourdeadOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("BringYourDead", idx, bringyourdead.Name)
		map_BringYourDead_Identifiers[bringyourdead] = id

		// Initialisation of values
		for _, _lancelot := range bringyourdead.Lancelots {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lancelots")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lancelot_Identifiers[_lancelot])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, document := range documentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Document", idx, document.Name)
		map_Document_Identifiers[document] = id

		// Initialisation of values
		for _, _geoobjectuse := range document.GeoObjectUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GeoObjectUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GeoObjectUse_Identifiers[_geoobjectuse])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, documentuse := range documentuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DocumentUse", idx, documentuse.Name)
		map_DocumentUse_Identifiers[documentuse] = id

		// Initialisation of values
		if documentuse.Document != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Document")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Document_Identifiers[documentuse.Document])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, galahadthepure := range galahadthepureOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GalahadThePure", idx, galahadthepure.Name)
		map_GalahadThePure_Identifiers[galahadthepure] = id

		// Initialisation of values
	}

	for idx, geoobject := range geoobjectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GeoObject", idx, geoobject.Name)
		map_GeoObject_Identifiers[geoobject] = id

		// Initialisation of values
	}

	for idx, geoobjectuse := range geoobjectuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GeoObjectUse", idx, geoobjectuse.Name)
		map_GeoObjectUse_Identifiers[geoobjectuse] = id

		// Initialisation of values
		if geoobjectuse.GeoObject != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GeoObject")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GeoObject_Identifiers[geoobjectuse.GeoObject])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, group := range groupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Group", idx, group.Name)
		map_Group_Identifiers[group] = id

		// Initialisation of values
		for _, _useruse := range group.UserUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "UserUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_UserUse_Identifiers[_useruse])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, groupuse := range groupuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GroupUse", idx, groupuse.Name)
		map_GroupUse_Identifiers[groupuse] = id

		// Initialisation of values
		if groupuse.Group != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Group")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[groupuse.Group])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, kingarthur := range kingarthurOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("KingArthur", idx, kingarthur.Name)
		map_KingArthur_Identifiers[kingarthur] = id

		// Initialisation of values
	}

	for idx, kingarthurshape := range kingarthurshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("KingArthurShape", idx, kingarthurshape.Name)
		map_KingArthurShape_Identifiers[kingarthurshape] = id

		// Initialisation of values
		if kingarthurshape.ActorState != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ActorState")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KingArthur_Identifiers[kingarthurshape.ActorState])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, knightwhosayni := range knightwhosayniOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("KnightWhoSayNi", idx, knightwhosayni.Name)
		map_KnightWhoSayNi_Identifiers[knightwhosayni] = id

		// Initialisation of values
		if knightwhosayni.Parameter != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Parameter")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lancelot_Identifiers[knightwhosayni.Parameter])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, lancelot := range lancelotOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Lancelot", idx, lancelot.Name)
		map_Lancelot_Identifiers[lancelot] = id

		// Initialisation of values
		for _, _groupuse := range lancelot.GroupUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GroupUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GroupUse_Identifiers[_groupuse])
			pointersInitializesStatements += setPointerField
		}

		for _, _documentuse := range lancelot.DocumentUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DocumentUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DocumentUse_Identifiers[_documentuse])
			pointersInitializesStatements += setPointerField
		}

		for _, _geoobjectuse := range lancelot.GeoObjectUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GeoObjectUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GeoObjectUse_Identifiers[_geoobjectuse])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, lancelotagregation := range lancelotagregationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LancelotAgregation", idx, lancelotagregation.Name)
		map_LancelotAgregation_Identifiers[lancelotagregation] = id

		// Initialisation of values
		for _, _knightwhosayni := range lancelotagregation.ParameterUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParameterUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KnightWhoSayNi_Identifiers[_knightwhosayni])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, lancelotagregationuse := range lancelotagregationuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LancelotAgregationUse", idx, lancelotagregationuse.Name)
		map_LancelotAgregationUse_Identifiers[lancelotagregationuse] = id

		// Initialisation of values
		if lancelotagregationuse.ParameterAgregation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParameterAgregation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LancelotAgregation_Identifiers[lancelotagregationuse.ParameterAgregation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, lancelotcategory := range lancelotcategoryOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LancelotCategory", idx, lancelotcategory.Name)
		map_LancelotCategory_Identifiers[lancelotcategory] = id

		// Initialisation of values
		for _, _knightwhosayni := range lancelotcategory.ParameterUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParameterUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KnightWhoSayNi_Identifiers[_knightwhosayni])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, lancelotcategoryuse := range lancelotcategoryuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LancelotCategoryUse", idx, lancelotcategoryuse.Name)
		map_LancelotCategoryUse_Identifiers[lancelotcategoryuse] = id

		// Initialisation of values
		if lancelotcategoryuse.ParameterCategory != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParameterCategory")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LancelotCategory_Identifiers[lancelotcategoryuse.ParameterCategory])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, mapobject := range mapobjectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MapObject", idx, mapobject.Name)
		map_MapObject_Identifiers[mapobject] = id

		// Initialisation of values
	}

	for idx, mapobjectuse := range mapobjectuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MapObjectUse", idx, mapobjectuse.Name)
		map_MapObjectUse_Identifiers[mapobjectuse] = id

		// Initialisation of values
		if mapobjectuse.Map != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Map")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MapObject_Identifiers[mapobjectuse.Map])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, repository := range repositoryOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Repository", idx, repository.Name)
		map_Repository_Identifiers[repository] = id

		// Initialisation of values
		for _, _knightwhosayni := range repository.ParameterUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParameterUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KnightWhoSayNi_Identifiers[_knightwhosayni])
			pointersInitializesStatements += setPointerField
		}

		for _, _groupuse := range repository.GroupUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GroupUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GroupUse_Identifiers[_groupuse])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, sirrobin := range sirrobinOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SirRobin", idx, sirrobin.Name)
		map_SirRobin_Identifiers[sirrobin] = id

		// Initialisation of values
		for _, _awitch := range sirrobin.Witches {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Witches")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AWitch_Identifiers[_awitch])
			pointersInitializesStatements += setPointerField
		}

		for _, _kingarthurshape := range sirrobin.Arthurs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Arthurs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KingArthurShape_Identifiers[_kingarthurshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _knightwhosayni := range sirrobin.KnightWhoSayNis {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "KnightWhoSayNis")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KnightWhoSayNi_Identifiers[_knightwhosayni])
			pointersInitializesStatements += setPointerField
		}

		for _, _blackknightshape := range sirrobin.BlackKnightShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "BlackKnightShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BlackKnightShape_Identifiers[_blackknightshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _thenuteshape := range sirrobin.TheNuteShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TheNuteShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TheNuteShape_Identifiers[_thenuteshape])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, thebridge := range thebridgeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TheBridge", idx, thebridge.Name)
		map_TheBridge_Identifiers[thebridge] = id

		// Initialisation of values
		for _, _whatisyourpreferedcolor := range thebridge.WhatIsYourPreferedColor {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "WhatIsYourPreferedColor")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_WhatIsYourPreferedColor_Identifiers[_whatisyourpreferedcolor])
			pointersInitializesStatements += setPointerField
		}

		for _, _groupuse := range thebridge.GroupUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GroupUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GroupUse_Identifiers[_groupuse])
			pointersInitializesStatements += setPointerField
		}

		for _, _geoobjectuse := range thebridge.GeoObjectUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GeoObjectUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GeoObjectUse_Identifiers[_geoobjectuse])
			pointersInitializesStatements += setPointerField
		}

		for _, _mapobjectuse := range thebridge.MapUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MapUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MapObjectUse_Identifiers[_mapobjectuse])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, thenuteshape := range thenuteshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TheNuteShape", idx, thenuteshape.Name)
		map_TheNuteShape_Identifiers[thenuteshape] = id

		// Initialisation of values
		if thenuteshape.ActorStateTransition != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ActorStateTransition")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TheNuteTransition_Identifiers[thenuteshape.ActorStateTransition])
			pointersInitializesStatements += setPointerField
		}

		if thenuteshape.Start != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Start")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KingArthurShape_Identifiers[thenuteshape.Start])
			pointersInitializesStatements += setPointerField
		}

		if thenuteshape.End != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "End")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KingArthurShape_Identifiers[thenuteshape.End])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, thenutetransition := range thenutetransitionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TheNuteTransition", idx, thenutetransition.Name)
		map_TheNuteTransition_Identifiers[thenutetransition] = id

		// Initialisation of values
		if thenutetransition.StartState != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "StartState")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KingArthur_Identifiers[thenutetransition.StartState])
			pointersInitializesStatements += setPointerField
		}

		if thenutetransition.EndState != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EndState")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KingArthur_Identifiers[thenutetransition.EndState])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, user := range userOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("User", idx, user.Name)
		map_User_Identifiers[user] = id

		// Initialisation of values
	}

	for idx, useruse := range useruseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("UserUse", idx, useruse.Name)
		map_UserUse_Identifiers[useruse] = id

		// Initialisation of values
		if useruse.User != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "User")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_User_Identifiers[useruse.User])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, whatisyourpreferedcolor := range whatisyourpreferedcolorOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("WhatIsYourPreferedColor", idx, whatisyourpreferedcolor.Name)
		map_WhatIsYourPreferedColor_Identifiers[whatisyourpreferedcolor] = id

		// Initialisation of values
		for _, _sirrobin := range whatisyourpreferedcolor.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SirRobin_Identifiers[_sirrobin])
			pointersInitializesStatements += setPointerField
		}

		for _, _kingarthur := range whatisyourpreferedcolor.KingArthurs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "KingArthurs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KingArthur_Identifiers[_kingarthur])
			pointersInitializesStatements += setPointerField
		}

		for _, _thenutetransition := range whatisyourpreferedcolor.Nutes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Nutes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TheNuteTransition_Identifiers[_thenutetransition])
			pointersInitializesStatements += setPointerField
		}

		for _, _galahadthepure := range whatisyourpreferedcolor.Galahard {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Galahard")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GalahadThePure_Identifiers[_galahadthepure])
			pointersInitializesStatements += setPointerField
		}

		for _, _lancelot := range whatisyourpreferedcolor.Lancelots {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lancelots")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lancelot_Identifiers[_lancelot])
			pointersInitializesStatements += setPointerField
		}

		for _, _bringyourdead := range whatisyourpreferedcolor.BringYourDeadarameters {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "BringYourDeadarameters")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BringYourDead_Identifiers[_bringyourdead])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, workspace := range workspaceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Workspace", idx, workspace.Name)
		map_Workspace_Identifiers[workspace] = id

		// Initialisation of values
		if workspace.SelectedDiagram != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SelectedDiagram")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SirRobin_Identifiers[workspace.SelectedDiagram])
			pointersInitializesStatements += setPointerField
		}

		if workspace.A != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "A")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AWitch_Identifiers[workspace.A])
			pointersInitializesStatements += setPointerField
		}

		if workspace.B != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "B")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KnightWhoSayNi_Identifiers[workspace.B])
			pointersInitializesStatements += setPointerField
		}

		if workspace.C != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "C")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BlackKnightShape_Identifiers[workspace.C])
			pointersInitializesStatements += setPointerField
		}

		if workspace.D != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "D")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_KingArthurShape_Identifiers[workspace.D])
			pointersInitializesStatements += setPointerField
		}

		if workspace.E != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "E")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TheNuteShape_Identifiers[workspace.E])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar ___dummy__%s_%s %s.StageStruct",
				stage.MetaPackageImportAlias,
				strings.ReplaceAll(filepath.Base(name), ".go", ""),
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
