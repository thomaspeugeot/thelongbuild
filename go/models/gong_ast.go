// generated code - do not edit
package models

import (
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var dummy_time_import time.Time

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(stage *StageStruct, pathToFile string) error {

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *StageStruct, inFile *ast.File, fset *token.FileSet) error {
	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", inFile.Name)
	}
	if len(inFile.Imports) == 3 {
		stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "Injection")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								stage, &cmap, assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "map_DocLink_Identifier") {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}
							var expressionType GONG__ExpressionType = GONG__STRUCT_INSTANCE
							var docLink GONG__Identifier

							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								expressionType = GONG__FIELD_OR_CONST_VALUE
							}

							var callExpr *ast.CallExpr
							if callExpr, ok = kve.Value.(*ast.CallExpr); ok {

								var se *ast.SelectorExpr
								if se, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(callExpr.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}

								// check the arg type to select wether this is a int or a string enum
								var bl *ast.BasicLit
								if bl, ok = callExpr.Args[0].(*ast.BasicLit); ok {
									switch bl.Kind {
									case token.STRING:
										expressionType = GONG__ENUM_CAST_STRING
									case token.INT:
										expressionType = GONG__ENUM_CAST_INT
									}
								} else {
									log.Fatal("Expression should be a basic lit" +
										fset.Position(se.Pos()).String())
								}

								docLink.Ident = id.Name + "." + se.Sel.Name
								_ = callExpr
							}

							var se2 *ast.SelectorExpr
							switch expressionType {
							case GONG__FIELD_OR_CONST_VALUE:
								if se2, ok = kve.Value.(*ast.SelectorExpr); ok {

									var ident *ast.Ident
									if _, ok = se2.X.(*ast.ParenExpr); ok {
										expressionType = GONG__FIELD_VALUE
										fieldName = se2.Sel.Name
									} else if ident, ok = se2.X.(*ast.Ident); ok {
										expressionType = GONG__IDENTIFIER_CONST
										docLink.Ident = ident.Name + "." + se2.Sel.Name
									} else {
										log.Fatal("Expression should be a selector expression or an ident" +
											fset.Position(kve.Pos()).String())
									}
								} else {

								}
							}

							var pe *ast.ParenExpr
							switch expressionType {
							case GONG__STRUCT_INSTANCE:
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							case GONG__FIELD_VALUE:
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}
							switch expressionType {
							case GONG__FIELD_VALUE, GONG__STRUCT_INSTANCE:
								// expect a Composite Litteral with no Element <type>{}
								var cl *ast.CompositeLit
								if cl, ok = pe.X.(*ast.CompositeLit); !ok {
									log.Fatal("Expression should be a composite lit" +
										fset.Position(pe.Pos()).String())
								}

								var se *ast.SelectorExpr
								if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector" +
										fset.Position(cl.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}
								docLink.Ident = id.Name + "." + se.Sel.Name
							}

							switch expressionType {
							case GONG__FIELD_VALUE:
								docLink.Ident += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink.Ident == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							docLink.Type = expressionType
							stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_AWitch = make(map[string]*AWitch)
var __gong__map_BlackKnightShape = make(map[string]*BlackKnightShape)
var __gong__map_BringYourDead = make(map[string]*BringYourDead)
var __gong__map_Document = make(map[string]*Document)
var __gong__map_DocumentUse = make(map[string]*DocumentUse)
var __gong__map_GalahadThePure = make(map[string]*GalahadThePure)
var __gong__map_GeoObject = make(map[string]*GeoObject)
var __gong__map_GeoObjectUse = make(map[string]*GeoObjectUse)
var __gong__map_Group = make(map[string]*Group)
var __gong__map_GroupUse = make(map[string]*GroupUse)
var __gong__map_KingArthur = make(map[string]*KingArthur)
var __gong__map_KingArthurShape = make(map[string]*KingArthurShape)
var __gong__map_KnightWhoSayNi = make(map[string]*KnightWhoSayNi)
var __gong__map_Lancelot = make(map[string]*Lancelot)
var __gong__map_LancelotAgregation = make(map[string]*LancelotAgregation)
var __gong__map_LancelotAgregationUse = make(map[string]*LancelotAgregationUse)
var __gong__map_LancelotCategory = make(map[string]*LancelotCategory)
var __gong__map_LancelotCategoryUse = make(map[string]*LancelotCategoryUse)
var __gong__map_MapObject = make(map[string]*MapObject)
var __gong__map_MapObjectUse = make(map[string]*MapObjectUse)
var __gong__map_Repository = make(map[string]*Repository)
var __gong__map_SirRobin = make(map[string]*SirRobin)
var __gong__map_TheBridge = make(map[string]*TheBridge)
var __gong__map_TheNuteShape = make(map[string]*TheNuteShape)
var __gong__map_TheNuteTransition = make(map[string]*TheNuteTransition)
var __gong__map_User = make(map[string]*User)
var __gong__map_UserUse = make(map[string]*UserUse)
var __gong__map_WhatIsYourPreferedColor = make(map[string]*WhatIsYourPreferedColor)
var __gong__map_Workspace = make(map[string]*Workspace)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *StageStruct, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed.Ident
						}
					}
				}
			}
		}
	}

	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "AWitch":
										instanceAWitch := (&AWitch{Name: instanceName}).Stage(stage)
										instance = any(instanceAWitch)
										__gong__map_AWitch[identifier] = instanceAWitch
									case "BlackKnightShape":
										instanceBlackKnightShape := (&BlackKnightShape{Name: instanceName}).Stage(stage)
										instance = any(instanceBlackKnightShape)
										__gong__map_BlackKnightShape[identifier] = instanceBlackKnightShape
									case "BringYourDead":
										instanceBringYourDead := (&BringYourDead{Name: instanceName}).Stage(stage)
										instance = any(instanceBringYourDead)
										__gong__map_BringYourDead[identifier] = instanceBringYourDead
									case "Document":
										instanceDocument := (&Document{Name: instanceName}).Stage(stage)
										instance = any(instanceDocument)
										__gong__map_Document[identifier] = instanceDocument
									case "DocumentUse":
										instanceDocumentUse := (&DocumentUse{Name: instanceName}).Stage(stage)
										instance = any(instanceDocumentUse)
										__gong__map_DocumentUse[identifier] = instanceDocumentUse
									case "GalahadThePure":
										instanceGalahadThePure := (&GalahadThePure{Name: instanceName}).Stage(stage)
										instance = any(instanceGalahadThePure)
										__gong__map_GalahadThePure[identifier] = instanceGalahadThePure
									case "GeoObject":
										instanceGeoObject := (&GeoObject{Name: instanceName}).Stage(stage)
										instance = any(instanceGeoObject)
										__gong__map_GeoObject[identifier] = instanceGeoObject
									case "GeoObjectUse":
										instanceGeoObjectUse := (&GeoObjectUse{Name: instanceName}).Stage(stage)
										instance = any(instanceGeoObjectUse)
										__gong__map_GeoObjectUse[identifier] = instanceGeoObjectUse
									case "Group":
										instanceGroup := (&Group{Name: instanceName}).Stage(stage)
										instance = any(instanceGroup)
										__gong__map_Group[identifier] = instanceGroup
									case "GroupUse":
										instanceGroupUse := (&GroupUse{Name: instanceName}).Stage(stage)
										instance = any(instanceGroupUse)
										__gong__map_GroupUse[identifier] = instanceGroupUse
									case "KingArthur":
										instanceKingArthur := (&KingArthur{Name: instanceName}).Stage(stage)
										instance = any(instanceKingArthur)
										__gong__map_KingArthur[identifier] = instanceKingArthur
									case "KingArthurShape":
										instanceKingArthurShape := (&KingArthurShape{Name: instanceName}).Stage(stage)
										instance = any(instanceKingArthurShape)
										__gong__map_KingArthurShape[identifier] = instanceKingArthurShape
									case "KnightWhoSayNi":
										instanceKnightWhoSayNi := (&KnightWhoSayNi{Name: instanceName}).Stage(stage)
										instance = any(instanceKnightWhoSayNi)
										__gong__map_KnightWhoSayNi[identifier] = instanceKnightWhoSayNi
									case "Lancelot":
										instanceLancelot := (&Lancelot{Name: instanceName}).Stage(stage)
										instance = any(instanceLancelot)
										__gong__map_Lancelot[identifier] = instanceLancelot
									case "LancelotAgregation":
										instanceLancelotAgregation := (&LancelotAgregation{Name: instanceName}).Stage(stage)
										instance = any(instanceLancelotAgregation)
										__gong__map_LancelotAgregation[identifier] = instanceLancelotAgregation
									case "LancelotAgregationUse":
										instanceLancelotAgregationUse := (&LancelotAgregationUse{Name: instanceName}).Stage(stage)
										instance = any(instanceLancelotAgregationUse)
										__gong__map_LancelotAgregationUse[identifier] = instanceLancelotAgregationUse
									case "LancelotCategory":
										instanceLancelotCategory := (&LancelotCategory{Name: instanceName}).Stage(stage)
										instance = any(instanceLancelotCategory)
										__gong__map_LancelotCategory[identifier] = instanceLancelotCategory
									case "LancelotCategoryUse":
										instanceLancelotCategoryUse := (&LancelotCategoryUse{Name: instanceName}).Stage(stage)
										instance = any(instanceLancelotCategoryUse)
										__gong__map_LancelotCategoryUse[identifier] = instanceLancelotCategoryUse
									case "MapObject":
										instanceMapObject := (&MapObject{Name: instanceName}).Stage(stage)
										instance = any(instanceMapObject)
										__gong__map_MapObject[identifier] = instanceMapObject
									case "MapObjectUse":
										instanceMapObjectUse := (&MapObjectUse{Name: instanceName}).Stage(stage)
										instance = any(instanceMapObjectUse)
										__gong__map_MapObjectUse[identifier] = instanceMapObjectUse
									case "Repository":
										instanceRepository := (&Repository{Name: instanceName}).Stage(stage)
										instance = any(instanceRepository)
										__gong__map_Repository[identifier] = instanceRepository
									case "SirRobin":
										instanceSirRobin := (&SirRobin{Name: instanceName}).Stage(stage)
										instance = any(instanceSirRobin)
										__gong__map_SirRobin[identifier] = instanceSirRobin
									case "TheBridge":
										instanceTheBridge := (&TheBridge{Name: instanceName}).Stage(stage)
										instance = any(instanceTheBridge)
										__gong__map_TheBridge[identifier] = instanceTheBridge
									case "TheNuteShape":
										instanceTheNuteShape := (&TheNuteShape{Name: instanceName}).Stage(stage)
										instance = any(instanceTheNuteShape)
										__gong__map_TheNuteShape[identifier] = instanceTheNuteShape
									case "TheNuteTransition":
										instanceTheNuteTransition := (&TheNuteTransition{Name: instanceName}).Stage(stage)
										instance = any(instanceTheNuteTransition)
										__gong__map_TheNuteTransition[identifier] = instanceTheNuteTransition
									case "User":
										instanceUser := (&User{Name: instanceName}).Stage(stage)
										instance = any(instanceUser)
										__gong__map_User[identifier] = instanceUser
									case "UserUse":
										instanceUserUse := (&UserUse{Name: instanceName}).Stage(stage)
										instance = any(instanceUserUse)
										__gong__map_UserUse[identifier] = instanceUserUse
									case "WhatIsYourPreferedColor":
										instanceWhatIsYourPreferedColor := (&WhatIsYourPreferedColor{Name: instanceName}).Stage(stage)
										instance = any(instanceWhatIsYourPreferedColor)
										__gong__map_WhatIsYourPreferedColor[identifier] = instanceWhatIsYourPreferedColor
									case "Workspace":
										instanceWorkspace := (&Workspace{Name: instanceName}).Stage(stage)
										instance = any(instanceWorkspace)
										__gong__map_Workspace[identifier] = instanceWorkspace
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "AWitch":
							switch fieldName {
							// insertion point for date assign code
							}
						case "BlackKnightShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "BringYourDead":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Document":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DocumentUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GalahadThePure":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GeoObject":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GeoObjectUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Group":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GroupUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "KingArthur":
							switch fieldName {
							// insertion point for date assign code
							}
						case "KingArthurShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "KnightWhoSayNi":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Lancelot":
							switch fieldName {
							// insertion point for date assign code
							}
						case "LancelotAgregation":
							switch fieldName {
							// insertion point for date assign code
							}
						case "LancelotAgregationUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "LancelotCategory":
							switch fieldName {
							// insertion point for date assign code
							}
						case "LancelotCategoryUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "MapObject":
							switch fieldName {
							// insertion point for date assign code
							}
						case "MapObjectUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Repository":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SirRobin":
							switch fieldName {
							// insertion point for date assign code
							case "Start":
								__gong__map_SirRobin[identifier].Start, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							case "End":
								__gong__map_SirRobin[identifier].End, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "TheBridge":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TheNuteShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TheNuteTransition":
							switch fieldName {
							// insertion point for date assign code
							}
						case "User":
							switch fieldName {
							// insertion point for date assign code
							}
						case "UserUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "WhatIsYourPreferedColor":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Workspace":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "AWitch":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "BlackKnightShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "BringYourDead":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Lancelots":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Lancelot[targetIdentifier]
							__gong__map_BringYourDead[identifier].Lancelots =
								append(__gong__map_BringYourDead[identifier].Lancelots, target)
						}
					case "Document":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GeoObjectUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GeoObjectUse[targetIdentifier]
							__gong__map_Document[identifier].GeoObjectUse =
								append(__gong__map_Document[identifier].GeoObjectUse, target)
						}
					case "DocumentUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GalahadThePure":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GeoObject":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GeoObjectUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Group":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "UserUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_UserUse[targetIdentifier]
							__gong__map_Group[identifier].UserUse =
								append(__gong__map_Group[identifier].UserUse, target)
						}
					case "GroupUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "KingArthur":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "KingArthurShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "KnightWhoSayNi":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Lancelot":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GroupUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GroupUse[targetIdentifier]
							__gong__map_Lancelot[identifier].GroupUse =
								append(__gong__map_Lancelot[identifier].GroupUse, target)
						case "DocumentUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DocumentUse[targetIdentifier]
							__gong__map_Lancelot[identifier].DocumentUse =
								append(__gong__map_Lancelot[identifier].DocumentUse, target)
						case "GeoObjectUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GeoObjectUse[targetIdentifier]
							__gong__map_Lancelot[identifier].GeoObjectUse =
								append(__gong__map_Lancelot[identifier].GeoObjectUse, target)
						}
					case "LancelotAgregation":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ParameterUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_KnightWhoSayNi[targetIdentifier]
							__gong__map_LancelotAgregation[identifier].ParameterUse =
								append(__gong__map_LancelotAgregation[identifier].ParameterUse, target)
						}
					case "LancelotAgregationUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "LancelotCategory":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ParameterUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_KnightWhoSayNi[targetIdentifier]
							__gong__map_LancelotCategory[identifier].ParameterUse =
								append(__gong__map_LancelotCategory[identifier].ParameterUse, target)
						}
					case "LancelotCategoryUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "MapObject":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "MapObjectUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Repository":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "ParameterUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_KnightWhoSayNi[targetIdentifier]
							__gong__map_Repository[identifier].ParameterUse =
								append(__gong__map_Repository[identifier].ParameterUse, target)
						case "GroupUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GroupUse[targetIdentifier]
							__gong__map_Repository[identifier].GroupUse =
								append(__gong__map_Repository[identifier].GroupUse, target)
						}
					case "SirRobin":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Witches":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_AWitch[targetIdentifier]
							__gong__map_SirRobin[identifier].Witches =
								append(__gong__map_SirRobin[identifier].Witches, target)
						case "Arthurs":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_KingArthurShape[targetIdentifier]
							__gong__map_SirRobin[identifier].Arthurs =
								append(__gong__map_SirRobin[identifier].Arthurs, target)
						case "KnightWhoSayNis":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_KnightWhoSayNi[targetIdentifier]
							__gong__map_SirRobin[identifier].KnightWhoSayNis =
								append(__gong__map_SirRobin[identifier].KnightWhoSayNis, target)
						case "BlackKnightShapes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_BlackKnightShape[targetIdentifier]
							__gong__map_SirRobin[identifier].BlackKnightShapes =
								append(__gong__map_SirRobin[identifier].BlackKnightShapes, target)
						case "TheNuteShapes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_TheNuteShape[targetIdentifier]
							__gong__map_SirRobin[identifier].TheNuteShapes =
								append(__gong__map_SirRobin[identifier].TheNuteShapes, target)
						}
					case "TheBridge":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "WhatIsYourPreferedColor":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_WhatIsYourPreferedColor[targetIdentifier]
							__gong__map_TheBridge[identifier].WhatIsYourPreferedColor =
								append(__gong__map_TheBridge[identifier].WhatIsYourPreferedColor, target)
						case "GroupUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GroupUse[targetIdentifier]
							__gong__map_TheBridge[identifier].GroupUse =
								append(__gong__map_TheBridge[identifier].GroupUse, target)
						case "GeoObjectUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GeoObjectUse[targetIdentifier]
							__gong__map_TheBridge[identifier].GeoObjectUse =
								append(__gong__map_TheBridge[identifier].GeoObjectUse, target)
						case "MapUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_MapObjectUse[targetIdentifier]
							__gong__map_TheBridge[identifier].MapUse =
								append(__gong__map_TheBridge[identifier].MapUse, target)
						}
					case "TheNuteShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "TheNuteTransition":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "User":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "UserUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "WhatIsYourPreferedColor":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Diagrams":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SirRobin[targetIdentifier]
							__gong__map_WhatIsYourPreferedColor[identifier].Diagrams =
								append(__gong__map_WhatIsYourPreferedColor[identifier].Diagrams, target)
						case "KingArthurs":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_KingArthur[targetIdentifier]
							__gong__map_WhatIsYourPreferedColor[identifier].KingArthurs =
								append(__gong__map_WhatIsYourPreferedColor[identifier].KingArthurs, target)
						case "Nutes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_TheNuteTransition[targetIdentifier]
							__gong__map_WhatIsYourPreferedColor[identifier].Nutes =
								append(__gong__map_WhatIsYourPreferedColor[identifier].Nutes, target)
						case "Galahard":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GalahadThePure[targetIdentifier]
							__gong__map_WhatIsYourPreferedColor[identifier].Galahard =
								append(__gong__map_WhatIsYourPreferedColor[identifier].Galahard, target)
						case "Lancelots":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Lancelot[targetIdentifier]
							__gong__map_WhatIsYourPreferedColor[identifier].Lancelots =
								append(__gong__map_WhatIsYourPreferedColor[identifier].Lancelots, target)
						case "BringYourDeadarameters":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_BringYourDead[targetIdentifier]
							__gong__map_WhatIsYourPreferedColor[identifier].BringYourDeadarameters =
								append(__gong__map_WhatIsYourPreferedColor[identifier].BringYourDeadarameters, target)
						}
					case "Workspace":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used

			if bl, ok := expr.(*ast.BasicLit); ok {
				// expression is  for instance ... = 18.000
				basicLit = bl
			} else if ue, ok := expr.(*ast.UnaryExpr); ok {
				// expression is  for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				basicLit = ue.X.(*ast.BasicLit)
				exprSign = -1
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "AWitch":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AWitch[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AWitch[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AWitch[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AWitch[identifier].Width = exprSign * fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AWitch[identifier].Height = exprSign * fielValue
				case "FillColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AWitch[identifier].FillColor = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AWitch[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AWitch[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AWitch[identifier].StrokeWidth = exprSign * fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AWitch[identifier].RX = exprSign * fielValue
				}
			case "BlackKnightShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_BlackKnightShape[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BlackKnightShape[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BlackKnightShape[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BlackKnightShape[identifier].Width = exprSign * fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BlackKnightShape[identifier].Height = exprSign * fielValue
				case "FillColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_BlackKnightShape[identifier].FillColor = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BlackKnightShape[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_BlackKnightShape[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BlackKnightShape[identifier].StrokeWidth = exprSign * fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_BlackKnightShape[identifier].RX = exprSign * fielValue
				}
			case "BringYourDead":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_BringYourDead[identifier].Name = fielValue
				case "Tag":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_BringYourDead[identifier].Tag = fielValue
				case "Description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_BringYourDead[identifier].Description = fielValue
				}
			case "Document":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Document[identifier].Name = fielValue
				}
			case "DocumentUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DocumentUse[identifier].Name = fielValue
				}
			case "GalahadThePure":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GalahadThePure[identifier].Name = fielValue
				case "Description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GalahadThePure[identifier].Description = fielValue
				}
			case "GeoObject":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GeoObject[identifier].Name = fielValue
				}
			case "GeoObjectUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GeoObjectUse[identifier].Name = fielValue
				}
			case "Group":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].Name = fielValue
				}
			case "GroupUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GroupUse[identifier].Name = fielValue
				}
			case "KingArthur":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_KingArthur[identifier].Name = fielValue
				}
			case "KingArthurShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_KingArthurShape[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KingArthurShape[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KingArthurShape[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KingArthurShape[identifier].Width = exprSign * fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KingArthurShape[identifier].Height = exprSign * fielValue
				case "FillColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_KingArthurShape[identifier].FillColor = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KingArthurShape[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_KingArthurShape[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KingArthurShape[identifier].StrokeWidth = exprSign * fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KingArthurShape[identifier].RX = exprSign * fielValue
				}
			case "KnightWhoSayNi":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_KnightWhoSayNi[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KnightWhoSayNi[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KnightWhoSayNi[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KnightWhoSayNi[identifier].Width = exprSign * fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KnightWhoSayNi[identifier].Height = exprSign * fielValue
				case "FillColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_KnightWhoSayNi[identifier].FillColor = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KnightWhoSayNi[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_KnightWhoSayNi[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KnightWhoSayNi[identifier].StrokeWidth = exprSign * fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KnightWhoSayNi[identifier].RX = exprSign * fielValue
				}
			case "Lancelot":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Lancelot[identifier].Name = fielValue
				case "Tag":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Lancelot[identifier].Tag = fielValue
				case "Description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Lancelot[identifier].Description = fielValue
				}
			case "LancelotAgregation":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LancelotAgregation[identifier].Name = fielValue
				}
			case "LancelotAgregationUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LancelotAgregationUse[identifier].Name = fielValue
				}
			case "LancelotCategory":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LancelotCategory[identifier].Name = fielValue
				}
			case "LancelotCategoryUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LancelotCategoryUse[identifier].Name = fielValue
				}
			case "MapObject":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MapObject[identifier].Name = fielValue
				}
			case "MapObjectUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MapObjectUse[identifier].Name = fielValue
				}
			case "Repository":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Repository[identifier].Name = fielValue
				}
			case "SirRobin":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SirRobin[identifier].Name = fielValue
				case "Description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SirRobin[identifier].Description = fielValue
				case "AxisOrign_X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SirRobin[identifier].AxisOrign_X = exprSign * fielValue
				case "AxisOrign_Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SirRobin[identifier].AxisOrign_Y = exprSign * fielValue
				case "VerticalAxis_Top_Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SirRobin[identifier].VerticalAxis_Top_Y = exprSign * fielValue
				case "VerticalAxis_Bottom_Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SirRobin[identifier].VerticalAxis_Bottom_Y = exprSign * fielValue
				case "VerticalAxis_StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SirRobin[identifier].VerticalAxis_StrokeWidth = exprSign * fielValue
				case "HorizontalAxis_Right_X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SirRobin[identifier].HorizontalAxis_Right_X = exprSign * fielValue
				}
			case "TheBridge":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TheBridge[identifier].Name = fielValue
				case "Description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TheBridge[identifier].Description = fielValue
				}
			case "TheNuteShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TheNuteShape[identifier].Name = fielValue
				case "StartRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TheNuteShape[identifier].StartRatio = exprSign * fielValue
				case "EndRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TheNuteShape[identifier].EndRatio = exprSign * fielValue
				case "CornerOffsetRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TheNuteShape[identifier].CornerOffsetRatio = exprSign * fielValue
				}
			case "TheNuteTransition":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TheNuteTransition[identifier].Name = fielValue
				}
			case "User":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_User[identifier].Name = fielValue
				}
			case "UserUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_UserUse[identifier].Name = fielValue
				}
			case "WhatIsYourPreferedColor":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_WhatIsYourPreferedColor[identifier].Name = fielValue
				case "Description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_WhatIsYourPreferedColor[identifier].Description = fielValue
				}
			case "Workspace":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Workspace[identifier].Name = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "AWitch":
				switch fieldName {
				// insertion point for field dependant code
				case "EvolutionDirection":
					targetIdentifier := ident.Name
					__gong__map_AWitch[identifier].EvolutionDirection = __gong__map_GalahadThePure[targetIdentifier]
				}
			case "BlackKnightShape":
				switch fieldName {
				// insertion point for field dependant code
				case "BringYourDead":
					targetIdentifier := ident.Name
					__gong__map_BlackKnightShape[identifier].BringYourDead = __gong__map_BringYourDead[targetIdentifier]
				}
			case "BringYourDead":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Document":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DocumentUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Document":
					targetIdentifier := ident.Name
					__gong__map_DocumentUse[identifier].Document = __gong__map_Document[targetIdentifier]
				}
			case "GalahadThePure":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GeoObject":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GeoObjectUse":
				switch fieldName {
				// insertion point for field dependant code
				case "GeoObject":
					targetIdentifier := ident.Name
					__gong__map_GeoObjectUse[identifier].GeoObject = __gong__map_GeoObject[targetIdentifier]
				}
			case "Group":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GroupUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Group":
					targetIdentifier := ident.Name
					__gong__map_GroupUse[identifier].Group = __gong__map_Group[targetIdentifier]
				}
			case "KingArthur":
				switch fieldName {
				// insertion point for field dependant code
				case "IsWithProbaility":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_KingArthur[identifier].IsWithProbaility = fielValue
				}
			case "KingArthurShape":
				switch fieldName {
				// insertion point for field dependant code
				case "ActorState":
					targetIdentifier := ident.Name
					__gong__map_KingArthurShape[identifier].ActorState = __gong__map_KingArthur[targetIdentifier]
				}
			case "KnightWhoSayNi":
				switch fieldName {
				// insertion point for field dependant code
				case "Parameter":
					targetIdentifier := ident.Name
					__gong__map_KnightWhoSayNi[identifier].Parameter = __gong__map_Lancelot[targetIdentifier]
				}
			case "Lancelot":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "LancelotAgregation":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "LancelotAgregationUse":
				switch fieldName {
				// insertion point for field dependant code
				case "ParameterAgregation":
					targetIdentifier := ident.Name
					__gong__map_LancelotAgregationUse[identifier].ParameterAgregation = __gong__map_LancelotAgregation[targetIdentifier]
				}
			case "LancelotCategory":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "LancelotCategoryUse":
				switch fieldName {
				// insertion point for field dependant code
				case "ParameterCategory":
					targetIdentifier := ident.Name
					__gong__map_LancelotCategoryUse[identifier].ParameterCategory = __gong__map_LancelotCategory[targetIdentifier]
				}
			case "MapObject":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "MapObjectUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Map":
					targetIdentifier := ident.Name
					__gong__map_MapObjectUse[identifier].Map = __gong__map_MapObject[targetIdentifier]
				}
			case "Repository":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SirRobin":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "TheBridge":
				switch fieldName {
				// insertion point for field dependant code
				case "IsNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TheBridge[identifier].IsNodeExpanded = fielValue
				}
			case "TheNuteShape":
				switch fieldName {
				// insertion point for field dependant code
				case "ActorStateTransition":
					targetIdentifier := ident.Name
					__gong__map_TheNuteShape[identifier].ActorStateTransition = __gong__map_TheNuteTransition[targetIdentifier]
				case "Start":
					targetIdentifier := ident.Name
					__gong__map_TheNuteShape[identifier].Start = __gong__map_KingArthurShape[targetIdentifier]
				case "End":
					targetIdentifier := ident.Name
					__gong__map_TheNuteShape[identifier].End = __gong__map_KingArthurShape[targetIdentifier]
				}
			case "TheNuteTransition":
				switch fieldName {
				// insertion point for field dependant code
				case "StartState":
					targetIdentifier := ident.Name
					__gong__map_TheNuteTransition[identifier].StartState = __gong__map_KingArthur[targetIdentifier]
				case "EndState":
					targetIdentifier := ident.Name
					__gong__map_TheNuteTransition[identifier].EndState = __gong__map_KingArthur[targetIdentifier]
				}
			case "User":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "UserUse":
				switch fieldName {
				// insertion point for field dependant code
				case "User":
					targetIdentifier := ident.Name
					__gong__map_UserUse[identifier].User = __gong__map_User[targetIdentifier]
				}
			case "WhatIsYourPreferedColor":
				switch fieldName {
				// insertion point for field dependant code
				case "DiagramsNodeFolded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_WhatIsYourPreferedColor[identifier].DiagramsNodeFolded = fielValue
				case "KingArthurNodeFolded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_WhatIsYourPreferedColor[identifier].KingArthurNodeFolded = fielValue
				case "RRRR":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_WhatIsYourPreferedColor[identifier].RRRR = fielValue
				case "IIUU":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_WhatIsYourPreferedColor[identifier].IIUU = fielValue
				case "LancelotsodeFolded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_WhatIsYourPreferedColor[identifier].LancelotsodeFolded = fielValue
				case "RRRRT":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_WhatIsYourPreferedColor[identifier].RRRRT = fielValue
				case "IsNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_WhatIsYourPreferedColor[identifier].IsNodeExpanded = fielValue
				}
			case "Workspace":
				switch fieldName {
				// insertion point for field dependant code
				case "SelectedDiagram":
					targetIdentifier := ident.Name
					__gong__map_Workspace[identifier].SelectedDiagram = __gong__map_SirRobin[targetIdentifier]
				case "A":
					targetIdentifier := ident.Name
					__gong__map_Workspace[identifier].A = __gong__map_AWitch[targetIdentifier]
				case "B":
					targetIdentifier := ident.Name
					__gong__map_Workspace[identifier].B = __gong__map_KnightWhoSayNi[targetIdentifier]
				case "C":
					targetIdentifier := ident.Name
					__gong__map_Workspace[identifier].C = __gong__map_BlackKnightShape[targetIdentifier]
				case "D":
					targetIdentifier := ident.Name
					__gong__map_Workspace[identifier].D = __gong__map_KingArthurShape[targetIdentifier]
				case "E":
					targetIdentifier := ident.Name
					__gong__map_Workspace[identifier].E = __gong__map_TheNuteShape[targetIdentifier]
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "AWitch":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "BlackKnightShape":
					switch fieldName {
					// insertion point for enum assign code
					case "Direction":
						var val DirectionType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_BlackKnightShape[identifier].Direction = DirectionType(val)
					}
				case "BringYourDead":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Document":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DocumentUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GalahadThePure":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GeoObject":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GeoObjectUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Group":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GroupUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "KingArthur":
					switch fieldName {
					// insertion point for enum assign code
					case "Probability":
						var val ProbabilityEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_KingArthur[identifier].Probability = ProbabilityEnum(val)
					}
				case "KingArthurShape":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "KnightWhoSayNi":
					switch fieldName {
					// insertion point for enum assign code
					case "Direction":
						var val DirectionType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_KnightWhoSayNi[identifier].Direction = DirectionType(val)
					}
				case "Lancelot":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "LancelotAgregation":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "LancelotAgregationUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "LancelotCategory":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "LancelotCategoryUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "MapObject":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "MapObjectUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Repository":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SirRobin":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TheBridge":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TheNuteShape":
					switch fieldName {
					// insertion point for enum assign code
					case "StartOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_TheNuteShape[identifier].StartOrientation = OrientationType(val)
					case "EndOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_TheNuteShape[identifier].EndOrientation = OrientationType(val)
					}
				case "TheNuteTransition":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "User":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "UserUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "WhatIsYourPreferedColor":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Workspace":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
