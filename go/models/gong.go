// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	AWitchs           map[*AWitch]any
	AWitchs_mapString map[string]*AWitch

	// insertion point for slice of pointers maps

	OnAfterAWitchCreateCallback OnAfterCreateInterface[AWitch]
	OnAfterAWitchUpdateCallback OnAfterUpdateInterface[AWitch]
	OnAfterAWitchDeleteCallback OnAfterDeleteInterface[AWitch]
	OnAfterAWitchReadCallback   OnAfterReadInterface[AWitch]

	BlackKnightShapes           map[*BlackKnightShape]any
	BlackKnightShapes_mapString map[string]*BlackKnightShape

	// insertion point for slice of pointers maps

	OnAfterBlackKnightShapeCreateCallback OnAfterCreateInterface[BlackKnightShape]
	OnAfterBlackKnightShapeUpdateCallback OnAfterUpdateInterface[BlackKnightShape]
	OnAfterBlackKnightShapeDeleteCallback OnAfterDeleteInterface[BlackKnightShape]
	OnAfterBlackKnightShapeReadCallback   OnAfterReadInterface[BlackKnightShape]

	BringYourDeads           map[*BringYourDead]any
	BringYourDeads_mapString map[string]*BringYourDead

	// insertion point for slice of pointers maps
	BringYourDead_Lancelots_reverseMap map[*Lancelot]*BringYourDead

	OnAfterBringYourDeadCreateCallback OnAfterCreateInterface[BringYourDead]
	OnAfterBringYourDeadUpdateCallback OnAfterUpdateInterface[BringYourDead]
	OnAfterBringYourDeadDeleteCallback OnAfterDeleteInterface[BringYourDead]
	OnAfterBringYourDeadReadCallback   OnAfterReadInterface[BringYourDead]

	Documents           map[*Document]any
	Documents_mapString map[string]*Document

	// insertion point for slice of pointers maps
	Document_GeoObjectUse_reverseMap map[*GeoObjectUse]*Document

	OnAfterDocumentCreateCallback OnAfterCreateInterface[Document]
	OnAfterDocumentUpdateCallback OnAfterUpdateInterface[Document]
	OnAfterDocumentDeleteCallback OnAfterDeleteInterface[Document]
	OnAfterDocumentReadCallback   OnAfterReadInterface[Document]

	DocumentUses           map[*DocumentUse]any
	DocumentUses_mapString map[string]*DocumentUse

	// insertion point for slice of pointers maps

	OnAfterDocumentUseCreateCallback OnAfterCreateInterface[DocumentUse]
	OnAfterDocumentUseUpdateCallback OnAfterUpdateInterface[DocumentUse]
	OnAfterDocumentUseDeleteCallback OnAfterDeleteInterface[DocumentUse]
	OnAfterDocumentUseReadCallback   OnAfterReadInterface[DocumentUse]

	GalahadThePures           map[*GalahadThePure]any
	GalahadThePures_mapString map[string]*GalahadThePure

	// insertion point for slice of pointers maps

	OnAfterGalahadThePureCreateCallback OnAfterCreateInterface[GalahadThePure]
	OnAfterGalahadThePureUpdateCallback OnAfterUpdateInterface[GalahadThePure]
	OnAfterGalahadThePureDeleteCallback OnAfterDeleteInterface[GalahadThePure]
	OnAfterGalahadThePureReadCallback   OnAfterReadInterface[GalahadThePure]

	GeoObjects           map[*GeoObject]any
	GeoObjects_mapString map[string]*GeoObject

	// insertion point for slice of pointers maps

	OnAfterGeoObjectCreateCallback OnAfterCreateInterface[GeoObject]
	OnAfterGeoObjectUpdateCallback OnAfterUpdateInterface[GeoObject]
	OnAfterGeoObjectDeleteCallback OnAfterDeleteInterface[GeoObject]
	OnAfterGeoObjectReadCallback   OnAfterReadInterface[GeoObject]

	GeoObjectUses           map[*GeoObjectUse]any
	GeoObjectUses_mapString map[string]*GeoObjectUse

	// insertion point for slice of pointers maps

	OnAfterGeoObjectUseCreateCallback OnAfterCreateInterface[GeoObjectUse]
	OnAfterGeoObjectUseUpdateCallback OnAfterUpdateInterface[GeoObjectUse]
	OnAfterGeoObjectUseDeleteCallback OnAfterDeleteInterface[GeoObjectUse]
	OnAfterGeoObjectUseReadCallback   OnAfterReadInterface[GeoObjectUse]

	Groups           map[*Group]any
	Groups_mapString map[string]*Group

	// insertion point for slice of pointers maps
	Group_UserUse_reverseMap map[*UserUse]*Group

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

	GroupUses           map[*GroupUse]any
	GroupUses_mapString map[string]*GroupUse

	// insertion point for slice of pointers maps

	OnAfterGroupUseCreateCallback OnAfterCreateInterface[GroupUse]
	OnAfterGroupUseUpdateCallback OnAfterUpdateInterface[GroupUse]
	OnAfterGroupUseDeleteCallback OnAfterDeleteInterface[GroupUse]
	OnAfterGroupUseReadCallback   OnAfterReadInterface[GroupUse]

	KingArthurs           map[*KingArthur]any
	KingArthurs_mapString map[string]*KingArthur

	// insertion point for slice of pointers maps

	OnAfterKingArthurCreateCallback OnAfterCreateInterface[KingArthur]
	OnAfterKingArthurUpdateCallback OnAfterUpdateInterface[KingArthur]
	OnAfterKingArthurDeleteCallback OnAfterDeleteInterface[KingArthur]
	OnAfterKingArthurReadCallback   OnAfterReadInterface[KingArthur]

	KingArthurShapes           map[*KingArthurShape]any
	KingArthurShapes_mapString map[string]*KingArthurShape

	// insertion point for slice of pointers maps

	OnAfterKingArthurShapeCreateCallback OnAfterCreateInterface[KingArthurShape]
	OnAfterKingArthurShapeUpdateCallback OnAfterUpdateInterface[KingArthurShape]
	OnAfterKingArthurShapeDeleteCallback OnAfterDeleteInterface[KingArthurShape]
	OnAfterKingArthurShapeReadCallback   OnAfterReadInterface[KingArthurShape]

	KnightWhoSayNis           map[*KnightWhoSayNi]any
	KnightWhoSayNis_mapString map[string]*KnightWhoSayNi

	// insertion point for slice of pointers maps

	OnAfterKnightWhoSayNiCreateCallback OnAfterCreateInterface[KnightWhoSayNi]
	OnAfterKnightWhoSayNiUpdateCallback OnAfterUpdateInterface[KnightWhoSayNi]
	OnAfterKnightWhoSayNiDeleteCallback OnAfterDeleteInterface[KnightWhoSayNi]
	OnAfterKnightWhoSayNiReadCallback   OnAfterReadInterface[KnightWhoSayNi]

	Lancelots           map[*Lancelot]any
	Lancelots_mapString map[string]*Lancelot

	// insertion point for slice of pointers maps
	Lancelot_GroupUse_reverseMap     map[*GroupUse]*Lancelot
	Lancelot_DocumentUse_reverseMap  map[*DocumentUse]*Lancelot
	Lancelot_GeoObjectUse_reverseMap map[*GeoObjectUse]*Lancelot

	OnAfterLancelotCreateCallback OnAfterCreateInterface[Lancelot]
	OnAfterLancelotUpdateCallback OnAfterUpdateInterface[Lancelot]
	OnAfterLancelotDeleteCallback OnAfterDeleteInterface[Lancelot]
	OnAfterLancelotReadCallback   OnAfterReadInterface[Lancelot]

	LancelotAgregations           map[*LancelotAgregation]any
	LancelotAgregations_mapString map[string]*LancelotAgregation

	// insertion point for slice of pointers maps
	LancelotAgregation_ParameterUse_reverseMap map[*KnightWhoSayNi]*LancelotAgregation

	OnAfterLancelotAgregationCreateCallback OnAfterCreateInterface[LancelotAgregation]
	OnAfterLancelotAgregationUpdateCallback OnAfterUpdateInterface[LancelotAgregation]
	OnAfterLancelotAgregationDeleteCallback OnAfterDeleteInterface[LancelotAgregation]
	OnAfterLancelotAgregationReadCallback   OnAfterReadInterface[LancelotAgregation]

	LancelotAgregationUses           map[*LancelotAgregationUse]any
	LancelotAgregationUses_mapString map[string]*LancelotAgregationUse

	// insertion point for slice of pointers maps

	OnAfterLancelotAgregationUseCreateCallback OnAfterCreateInterface[LancelotAgregationUse]
	OnAfterLancelotAgregationUseUpdateCallback OnAfterUpdateInterface[LancelotAgregationUse]
	OnAfterLancelotAgregationUseDeleteCallback OnAfterDeleteInterface[LancelotAgregationUse]
	OnAfterLancelotAgregationUseReadCallback   OnAfterReadInterface[LancelotAgregationUse]

	LancelotCategorys           map[*LancelotCategory]any
	LancelotCategorys_mapString map[string]*LancelotCategory

	// insertion point for slice of pointers maps
	LancelotCategory_ParameterUse_reverseMap map[*KnightWhoSayNi]*LancelotCategory

	OnAfterLancelotCategoryCreateCallback OnAfterCreateInterface[LancelotCategory]
	OnAfterLancelotCategoryUpdateCallback OnAfterUpdateInterface[LancelotCategory]
	OnAfterLancelotCategoryDeleteCallback OnAfterDeleteInterface[LancelotCategory]
	OnAfterLancelotCategoryReadCallback   OnAfterReadInterface[LancelotCategory]

	LancelotCategoryUses           map[*LancelotCategoryUse]any
	LancelotCategoryUses_mapString map[string]*LancelotCategoryUse

	// insertion point for slice of pointers maps

	OnAfterLancelotCategoryUseCreateCallback OnAfterCreateInterface[LancelotCategoryUse]
	OnAfterLancelotCategoryUseUpdateCallback OnAfterUpdateInterface[LancelotCategoryUse]
	OnAfterLancelotCategoryUseDeleteCallback OnAfterDeleteInterface[LancelotCategoryUse]
	OnAfterLancelotCategoryUseReadCallback   OnAfterReadInterface[LancelotCategoryUse]

	MapObjects           map[*MapObject]any
	MapObjects_mapString map[string]*MapObject

	// insertion point for slice of pointers maps

	OnAfterMapObjectCreateCallback OnAfterCreateInterface[MapObject]
	OnAfterMapObjectUpdateCallback OnAfterUpdateInterface[MapObject]
	OnAfterMapObjectDeleteCallback OnAfterDeleteInterface[MapObject]
	OnAfterMapObjectReadCallback   OnAfterReadInterface[MapObject]

	MapObjectUses           map[*MapObjectUse]any
	MapObjectUses_mapString map[string]*MapObjectUse

	// insertion point for slice of pointers maps

	OnAfterMapObjectUseCreateCallback OnAfterCreateInterface[MapObjectUse]
	OnAfterMapObjectUseUpdateCallback OnAfterUpdateInterface[MapObjectUse]
	OnAfterMapObjectUseDeleteCallback OnAfterDeleteInterface[MapObjectUse]
	OnAfterMapObjectUseReadCallback   OnAfterReadInterface[MapObjectUse]

	Repositorys           map[*Repository]any
	Repositorys_mapString map[string]*Repository

	// insertion point for slice of pointers maps
	Repository_ParameterUse_reverseMap map[*KnightWhoSayNi]*Repository
	Repository_GroupUse_reverseMap     map[*GroupUse]*Repository

	OnAfterRepositoryCreateCallback OnAfterCreateInterface[Repository]
	OnAfterRepositoryUpdateCallback OnAfterUpdateInterface[Repository]
	OnAfterRepositoryDeleteCallback OnAfterDeleteInterface[Repository]
	OnAfterRepositoryReadCallback   OnAfterReadInterface[Repository]

	SirRobins           map[*SirRobin]any
	SirRobins_mapString map[string]*SirRobin

	// insertion point for slice of pointers maps
	SirRobin_Witches_reverseMap           map[*AWitch]*SirRobin
	SirRobin_Arthurs_reverseMap           map[*KingArthurShape]*SirRobin
	SirRobin_KnightWhoSayNis_reverseMap   map[*KnightWhoSayNi]*SirRobin
	SirRobin_BlackKnightShapes_reverseMap map[*BlackKnightShape]*SirRobin
	SirRobin_TheNuteShapes_reverseMap     map[*TheNuteShape]*SirRobin

	OnAfterSirRobinCreateCallback OnAfterCreateInterface[SirRobin]
	OnAfterSirRobinUpdateCallback OnAfterUpdateInterface[SirRobin]
	OnAfterSirRobinDeleteCallback OnAfterDeleteInterface[SirRobin]
	OnAfterSirRobinReadCallback   OnAfterReadInterface[SirRobin]

	TheBridges           map[*TheBridge]any
	TheBridges_mapString map[string]*TheBridge

	// insertion point for slice of pointers maps
	TheBridge_WhatIsYourPreferedColor_reverseMap map[*WhatIsYourPreferedColor]*TheBridge
	TheBridge_GroupUse_reverseMap                map[*GroupUse]*TheBridge
	TheBridge_GeoObjectUse_reverseMap            map[*GeoObjectUse]*TheBridge
	TheBridge_MapUse_reverseMap                  map[*MapObjectUse]*TheBridge

	OnAfterTheBridgeCreateCallback OnAfterCreateInterface[TheBridge]
	OnAfterTheBridgeUpdateCallback OnAfterUpdateInterface[TheBridge]
	OnAfterTheBridgeDeleteCallback OnAfterDeleteInterface[TheBridge]
	OnAfterTheBridgeReadCallback   OnAfterReadInterface[TheBridge]

	TheNuteShapes           map[*TheNuteShape]any
	TheNuteShapes_mapString map[string]*TheNuteShape

	// insertion point for slice of pointers maps

	OnAfterTheNuteShapeCreateCallback OnAfterCreateInterface[TheNuteShape]
	OnAfterTheNuteShapeUpdateCallback OnAfterUpdateInterface[TheNuteShape]
	OnAfterTheNuteShapeDeleteCallback OnAfterDeleteInterface[TheNuteShape]
	OnAfterTheNuteShapeReadCallback   OnAfterReadInterface[TheNuteShape]

	TheNuteTransitions           map[*TheNuteTransition]any
	TheNuteTransitions_mapString map[string]*TheNuteTransition

	// insertion point for slice of pointers maps

	OnAfterTheNuteTransitionCreateCallback OnAfterCreateInterface[TheNuteTransition]
	OnAfterTheNuteTransitionUpdateCallback OnAfterUpdateInterface[TheNuteTransition]
	OnAfterTheNuteTransitionDeleteCallback OnAfterDeleteInterface[TheNuteTransition]
	OnAfterTheNuteTransitionReadCallback   OnAfterReadInterface[TheNuteTransition]

	Users           map[*User]any
	Users_mapString map[string]*User

	// insertion point for slice of pointers maps

	OnAfterUserCreateCallback OnAfterCreateInterface[User]
	OnAfterUserUpdateCallback OnAfterUpdateInterface[User]
	OnAfterUserDeleteCallback OnAfterDeleteInterface[User]
	OnAfterUserReadCallback   OnAfterReadInterface[User]

	UserUses           map[*UserUse]any
	UserUses_mapString map[string]*UserUse

	// insertion point for slice of pointers maps

	OnAfterUserUseCreateCallback OnAfterCreateInterface[UserUse]
	OnAfterUserUseUpdateCallback OnAfterUpdateInterface[UserUse]
	OnAfterUserUseDeleteCallback OnAfterDeleteInterface[UserUse]
	OnAfterUserUseReadCallback   OnAfterReadInterface[UserUse]

	WhatIsYourPreferedColors           map[*WhatIsYourPreferedColor]any
	WhatIsYourPreferedColors_mapString map[string]*WhatIsYourPreferedColor

	// insertion point for slice of pointers maps
	WhatIsYourPreferedColor_Diagrams_reverseMap               map[*SirRobin]*WhatIsYourPreferedColor
	WhatIsYourPreferedColor_KingArthurs_reverseMap            map[*KingArthur]*WhatIsYourPreferedColor
	WhatIsYourPreferedColor_Nutes_reverseMap                  map[*TheNuteTransition]*WhatIsYourPreferedColor
	WhatIsYourPreferedColor_Galahard_reverseMap               map[*GalahadThePure]*WhatIsYourPreferedColor
	WhatIsYourPreferedColor_Lancelots_reverseMap              map[*Lancelot]*WhatIsYourPreferedColor
	WhatIsYourPreferedColor_BringYourDeadarameters_reverseMap map[*BringYourDead]*WhatIsYourPreferedColor

	OnAfterWhatIsYourPreferedColorCreateCallback OnAfterCreateInterface[WhatIsYourPreferedColor]
	OnAfterWhatIsYourPreferedColorUpdateCallback OnAfterUpdateInterface[WhatIsYourPreferedColor]
	OnAfterWhatIsYourPreferedColorDeleteCallback OnAfterDeleteInterface[WhatIsYourPreferedColor]
	OnAfterWhatIsYourPreferedColorReadCallback   OnAfterReadInterface[WhatIsYourPreferedColor]

	Workspaces           map[*Workspace]any
	Workspaces_mapString map[string]*Workspace

	// insertion point for slice of pointers maps

	OnAfterWorkspaceCreateCallback OnAfterCreateInterface[Workspace]
	OnAfterWorkspaceUpdateCallback OnAfterUpdateInterface[Workspace]
	OnAfterWorkspaceDeleteCallback OnAfterDeleteInterface[Workspace]
	OnAfterWorkspaceReadCallback   OnAfterReadInterface[Workspace]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/thomaspeugeot/thelongbuild/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitAWitch(awitch *AWitch)
	CheckoutAWitch(awitch *AWitch)
	CommitBlackKnightShape(blackknightshape *BlackKnightShape)
	CheckoutBlackKnightShape(blackknightshape *BlackKnightShape)
	CommitBringYourDead(bringyourdead *BringYourDead)
	CheckoutBringYourDead(bringyourdead *BringYourDead)
	CommitDocument(document *Document)
	CheckoutDocument(document *Document)
	CommitDocumentUse(documentuse *DocumentUse)
	CheckoutDocumentUse(documentuse *DocumentUse)
	CommitGalahadThePure(galahadthepure *GalahadThePure)
	CheckoutGalahadThePure(galahadthepure *GalahadThePure)
	CommitGeoObject(geoobject *GeoObject)
	CheckoutGeoObject(geoobject *GeoObject)
	CommitGeoObjectUse(geoobjectuse *GeoObjectUse)
	CheckoutGeoObjectUse(geoobjectuse *GeoObjectUse)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
	CommitGroupUse(groupuse *GroupUse)
	CheckoutGroupUse(groupuse *GroupUse)
	CommitKingArthur(kingarthur *KingArthur)
	CheckoutKingArthur(kingarthur *KingArthur)
	CommitKingArthurShape(kingarthurshape *KingArthurShape)
	CheckoutKingArthurShape(kingarthurshape *KingArthurShape)
	CommitKnightWhoSayNi(knightwhosayni *KnightWhoSayNi)
	CheckoutKnightWhoSayNi(knightwhosayni *KnightWhoSayNi)
	CommitLancelot(lancelot *Lancelot)
	CheckoutLancelot(lancelot *Lancelot)
	CommitLancelotAgregation(lancelotagregation *LancelotAgregation)
	CheckoutLancelotAgregation(lancelotagregation *LancelotAgregation)
	CommitLancelotAgregationUse(lancelotagregationuse *LancelotAgregationUse)
	CheckoutLancelotAgregationUse(lancelotagregationuse *LancelotAgregationUse)
	CommitLancelotCategory(lancelotcategory *LancelotCategory)
	CheckoutLancelotCategory(lancelotcategory *LancelotCategory)
	CommitLancelotCategoryUse(lancelotcategoryuse *LancelotCategoryUse)
	CheckoutLancelotCategoryUse(lancelotcategoryuse *LancelotCategoryUse)
	CommitMapObject(mapobject *MapObject)
	CheckoutMapObject(mapobject *MapObject)
	CommitMapObjectUse(mapobjectuse *MapObjectUse)
	CheckoutMapObjectUse(mapobjectuse *MapObjectUse)
	CommitRepository(repository *Repository)
	CheckoutRepository(repository *Repository)
	CommitSirRobin(sirrobin *SirRobin)
	CheckoutSirRobin(sirrobin *SirRobin)
	CommitTheBridge(thebridge *TheBridge)
	CheckoutTheBridge(thebridge *TheBridge)
	CommitTheNuteShape(thenuteshape *TheNuteShape)
	CheckoutTheNuteShape(thenuteshape *TheNuteShape)
	CommitTheNuteTransition(thenutetransition *TheNuteTransition)
	CheckoutTheNuteTransition(thenutetransition *TheNuteTransition)
	CommitUser(user *User)
	CheckoutUser(user *User)
	CommitUserUse(useruse *UserUse)
	CheckoutUserUse(useruse *UserUse)
	CommitWhatIsYourPreferedColor(whatisyourpreferedcolor *WhatIsYourPreferedColor)
	CheckoutWhatIsYourPreferedColor(whatisyourpreferedcolor *WhatIsYourPreferedColor)
	CommitWorkspace(workspace *Workspace)
	CheckoutWorkspace(workspace *Workspace)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		AWitchs:           make(map[*AWitch]any),
		AWitchs_mapString: make(map[string]*AWitch),

		BlackKnightShapes:           make(map[*BlackKnightShape]any),
		BlackKnightShapes_mapString: make(map[string]*BlackKnightShape),

		BringYourDeads:           make(map[*BringYourDead]any),
		BringYourDeads_mapString: make(map[string]*BringYourDead),

		Documents:           make(map[*Document]any),
		Documents_mapString: make(map[string]*Document),

		DocumentUses:           make(map[*DocumentUse]any),
		DocumentUses_mapString: make(map[string]*DocumentUse),

		GalahadThePures:           make(map[*GalahadThePure]any),
		GalahadThePures_mapString: make(map[string]*GalahadThePure),

		GeoObjects:           make(map[*GeoObject]any),
		GeoObjects_mapString: make(map[string]*GeoObject),

		GeoObjectUses:           make(map[*GeoObjectUse]any),
		GeoObjectUses_mapString: make(map[string]*GeoObjectUse),

		Groups:           make(map[*Group]any),
		Groups_mapString: make(map[string]*Group),

		GroupUses:           make(map[*GroupUse]any),
		GroupUses_mapString: make(map[string]*GroupUse),

		KingArthurs:           make(map[*KingArthur]any),
		KingArthurs_mapString: make(map[string]*KingArthur),

		KingArthurShapes:           make(map[*KingArthurShape]any),
		KingArthurShapes_mapString: make(map[string]*KingArthurShape),

		KnightWhoSayNis:           make(map[*KnightWhoSayNi]any),
		KnightWhoSayNis_mapString: make(map[string]*KnightWhoSayNi),

		Lancelots:           make(map[*Lancelot]any),
		Lancelots_mapString: make(map[string]*Lancelot),

		LancelotAgregations:           make(map[*LancelotAgregation]any),
		LancelotAgregations_mapString: make(map[string]*LancelotAgregation),

		LancelotAgregationUses:           make(map[*LancelotAgregationUse]any),
		LancelotAgregationUses_mapString: make(map[string]*LancelotAgregationUse),

		LancelotCategorys:           make(map[*LancelotCategory]any),
		LancelotCategorys_mapString: make(map[string]*LancelotCategory),

		LancelotCategoryUses:           make(map[*LancelotCategoryUse]any),
		LancelotCategoryUses_mapString: make(map[string]*LancelotCategoryUse),

		MapObjects:           make(map[*MapObject]any),
		MapObjects_mapString: make(map[string]*MapObject),

		MapObjectUses:           make(map[*MapObjectUse]any),
		MapObjectUses_mapString: make(map[string]*MapObjectUse),

		Repositorys:           make(map[*Repository]any),
		Repositorys_mapString: make(map[string]*Repository),

		SirRobins:           make(map[*SirRobin]any),
		SirRobins_mapString: make(map[string]*SirRobin),

		TheBridges:           make(map[*TheBridge]any),
		TheBridges_mapString: make(map[string]*TheBridge),

		TheNuteShapes:           make(map[*TheNuteShape]any),
		TheNuteShapes_mapString: make(map[string]*TheNuteShape),

		TheNuteTransitions:           make(map[*TheNuteTransition]any),
		TheNuteTransitions_mapString: make(map[string]*TheNuteTransition),

		Users:           make(map[*User]any),
		Users_mapString: make(map[string]*User),

		UserUses:           make(map[*UserUse]any),
		UserUses_mapString: make(map[string]*UserUse),

		WhatIsYourPreferedColors:           make(map[*WhatIsYourPreferedColor]any),
		WhatIsYourPreferedColors_mapString: make(map[string]*WhatIsYourPreferedColor),

		Workspaces:           make(map[*Workspace]any),
		Workspaces_mapString: make(map[string]*Workspace),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AWitch"] = len(stage.AWitchs)
	stage.Map_GongStructName_InstancesNb["BlackKnightShape"] = len(stage.BlackKnightShapes)
	stage.Map_GongStructName_InstancesNb["BringYourDead"] = len(stage.BringYourDeads)
	stage.Map_GongStructName_InstancesNb["Document"] = len(stage.Documents)
	stage.Map_GongStructName_InstancesNb["DocumentUse"] = len(stage.DocumentUses)
	stage.Map_GongStructName_InstancesNb["GalahadThePure"] = len(stage.GalahadThePures)
	stage.Map_GongStructName_InstancesNb["GeoObject"] = len(stage.GeoObjects)
	stage.Map_GongStructName_InstancesNb["GeoObjectUse"] = len(stage.GeoObjectUses)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["GroupUse"] = len(stage.GroupUses)
	stage.Map_GongStructName_InstancesNb["KingArthur"] = len(stage.KingArthurs)
	stage.Map_GongStructName_InstancesNb["KingArthurShape"] = len(stage.KingArthurShapes)
	stage.Map_GongStructName_InstancesNb["KnightWhoSayNi"] = len(stage.KnightWhoSayNis)
	stage.Map_GongStructName_InstancesNb["Lancelot"] = len(stage.Lancelots)
	stage.Map_GongStructName_InstancesNb["LancelotAgregation"] = len(stage.LancelotAgregations)
	stage.Map_GongStructName_InstancesNb["LancelotAgregationUse"] = len(stage.LancelotAgregationUses)
	stage.Map_GongStructName_InstancesNb["LancelotCategory"] = len(stage.LancelotCategorys)
	stage.Map_GongStructName_InstancesNb["LancelotCategoryUse"] = len(stage.LancelotCategoryUses)
	stage.Map_GongStructName_InstancesNb["MapObject"] = len(stage.MapObjects)
	stage.Map_GongStructName_InstancesNb["MapObjectUse"] = len(stage.MapObjectUses)
	stage.Map_GongStructName_InstancesNb["Repository"] = len(stage.Repositorys)
	stage.Map_GongStructName_InstancesNb["SirRobin"] = len(stage.SirRobins)
	stage.Map_GongStructName_InstancesNb["TheBridge"] = len(stage.TheBridges)
	stage.Map_GongStructName_InstancesNb["TheNuteShape"] = len(stage.TheNuteShapes)
	stage.Map_GongStructName_InstancesNb["TheNuteTransition"] = len(stage.TheNuteTransitions)
	stage.Map_GongStructName_InstancesNb["User"] = len(stage.Users)
	stage.Map_GongStructName_InstancesNb["UserUse"] = len(stage.UserUses)
	stage.Map_GongStructName_InstancesNb["WhatIsYourPreferedColor"] = len(stage.WhatIsYourPreferedColors)
	stage.Map_GongStructName_InstancesNb["Workspace"] = len(stage.Workspaces)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AWitch"] = len(stage.AWitchs)
	stage.Map_GongStructName_InstancesNb["BlackKnightShape"] = len(stage.BlackKnightShapes)
	stage.Map_GongStructName_InstancesNb["BringYourDead"] = len(stage.BringYourDeads)
	stage.Map_GongStructName_InstancesNb["Document"] = len(stage.Documents)
	stage.Map_GongStructName_InstancesNb["DocumentUse"] = len(stage.DocumentUses)
	stage.Map_GongStructName_InstancesNb["GalahadThePure"] = len(stage.GalahadThePures)
	stage.Map_GongStructName_InstancesNb["GeoObject"] = len(stage.GeoObjects)
	stage.Map_GongStructName_InstancesNb["GeoObjectUse"] = len(stage.GeoObjectUses)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["GroupUse"] = len(stage.GroupUses)
	stage.Map_GongStructName_InstancesNb["KingArthur"] = len(stage.KingArthurs)
	stage.Map_GongStructName_InstancesNb["KingArthurShape"] = len(stage.KingArthurShapes)
	stage.Map_GongStructName_InstancesNb["KnightWhoSayNi"] = len(stage.KnightWhoSayNis)
	stage.Map_GongStructName_InstancesNb["Lancelot"] = len(stage.Lancelots)
	stage.Map_GongStructName_InstancesNb["LancelotAgregation"] = len(stage.LancelotAgregations)
	stage.Map_GongStructName_InstancesNb["LancelotAgregationUse"] = len(stage.LancelotAgregationUses)
	stage.Map_GongStructName_InstancesNb["LancelotCategory"] = len(stage.LancelotCategorys)
	stage.Map_GongStructName_InstancesNb["LancelotCategoryUse"] = len(stage.LancelotCategoryUses)
	stage.Map_GongStructName_InstancesNb["MapObject"] = len(stage.MapObjects)
	stage.Map_GongStructName_InstancesNb["MapObjectUse"] = len(stage.MapObjectUses)
	stage.Map_GongStructName_InstancesNb["Repository"] = len(stage.Repositorys)
	stage.Map_GongStructName_InstancesNb["SirRobin"] = len(stage.SirRobins)
	stage.Map_GongStructName_InstancesNb["TheBridge"] = len(stage.TheBridges)
	stage.Map_GongStructName_InstancesNb["TheNuteShape"] = len(stage.TheNuteShapes)
	stage.Map_GongStructName_InstancesNb["TheNuteTransition"] = len(stage.TheNuteTransitions)
	stage.Map_GongStructName_InstancesNb["User"] = len(stage.Users)
	stage.Map_GongStructName_InstancesNb["UserUse"] = len(stage.UserUses)
	stage.Map_GongStructName_InstancesNb["WhatIsYourPreferedColor"] = len(stage.WhatIsYourPreferedColors)
	stage.Map_GongStructName_InstancesNb["Workspace"] = len(stage.Workspaces)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts awitch to the model stage
func (awitch *AWitch) Stage(stage *StageStruct) *AWitch {
	stage.AWitchs[awitch] = __member
	stage.AWitchs_mapString[awitch.Name] = awitch

	return awitch
}

// Unstage removes awitch off the model stage
func (awitch *AWitch) Unstage(stage *StageStruct) *AWitch {
	delete(stage.AWitchs, awitch)
	delete(stage.AWitchs_mapString, awitch.Name)
	return awitch
}

// UnstageVoid removes awitch off the model stage
func (awitch *AWitch) UnstageVoid(stage *StageStruct) {
	delete(stage.AWitchs, awitch)
	delete(stage.AWitchs_mapString, awitch.Name)
}

// commit awitch to the back repo (if it is already staged)
func (awitch *AWitch) Commit(stage *StageStruct) *AWitch {
	if _, ok := stage.AWitchs[awitch]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAWitch(awitch)
		}
	}
	return awitch
}

func (awitch *AWitch) CommitVoid(stage *StageStruct) {
	awitch.Commit(stage)
}

// Checkout awitch to the back repo (if it is already staged)
func (awitch *AWitch) Checkout(stage *StageStruct) *AWitch {
	if _, ok := stage.AWitchs[awitch]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAWitch(awitch)
		}
	}
	return awitch
}

// for satisfaction of GongStruct interface
func (awitch *AWitch) GetName() (res string) {
	return awitch.Name
}

// Stage puts blackknightshape to the model stage
func (blackknightshape *BlackKnightShape) Stage(stage *StageStruct) *BlackKnightShape {
	stage.BlackKnightShapes[blackknightshape] = __member
	stage.BlackKnightShapes_mapString[blackknightshape.Name] = blackknightshape

	return blackknightshape
}

// Unstage removes blackknightshape off the model stage
func (blackknightshape *BlackKnightShape) Unstage(stage *StageStruct) *BlackKnightShape {
	delete(stage.BlackKnightShapes, blackknightshape)
	delete(stage.BlackKnightShapes_mapString, blackknightshape.Name)
	return blackknightshape
}

// UnstageVoid removes blackknightshape off the model stage
func (blackknightshape *BlackKnightShape) UnstageVoid(stage *StageStruct) {
	delete(stage.BlackKnightShapes, blackknightshape)
	delete(stage.BlackKnightShapes_mapString, blackknightshape.Name)
}

// commit blackknightshape to the back repo (if it is already staged)
func (blackknightshape *BlackKnightShape) Commit(stage *StageStruct) *BlackKnightShape {
	if _, ok := stage.BlackKnightShapes[blackknightshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBlackKnightShape(blackknightshape)
		}
	}
	return blackknightshape
}

func (blackknightshape *BlackKnightShape) CommitVoid(stage *StageStruct) {
	blackknightshape.Commit(stage)
}

// Checkout blackknightshape to the back repo (if it is already staged)
func (blackknightshape *BlackKnightShape) Checkout(stage *StageStruct) *BlackKnightShape {
	if _, ok := stage.BlackKnightShapes[blackknightshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBlackKnightShape(blackknightshape)
		}
	}
	return blackknightshape
}

// for satisfaction of GongStruct interface
func (blackknightshape *BlackKnightShape) GetName() (res string) {
	return blackknightshape.Name
}

// Stage puts bringyourdead to the model stage
func (bringyourdead *BringYourDead) Stage(stage *StageStruct) *BringYourDead {
	stage.BringYourDeads[bringyourdead] = __member
	stage.BringYourDeads_mapString[bringyourdead.Name] = bringyourdead

	return bringyourdead
}

// Unstage removes bringyourdead off the model stage
func (bringyourdead *BringYourDead) Unstage(stage *StageStruct) *BringYourDead {
	delete(stage.BringYourDeads, bringyourdead)
	delete(stage.BringYourDeads_mapString, bringyourdead.Name)
	return bringyourdead
}

// UnstageVoid removes bringyourdead off the model stage
func (bringyourdead *BringYourDead) UnstageVoid(stage *StageStruct) {
	delete(stage.BringYourDeads, bringyourdead)
	delete(stage.BringYourDeads_mapString, bringyourdead.Name)
}

// commit bringyourdead to the back repo (if it is already staged)
func (bringyourdead *BringYourDead) Commit(stage *StageStruct) *BringYourDead {
	if _, ok := stage.BringYourDeads[bringyourdead]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBringYourDead(bringyourdead)
		}
	}
	return bringyourdead
}

func (bringyourdead *BringYourDead) CommitVoid(stage *StageStruct) {
	bringyourdead.Commit(stage)
}

// Checkout bringyourdead to the back repo (if it is already staged)
func (bringyourdead *BringYourDead) Checkout(stage *StageStruct) *BringYourDead {
	if _, ok := stage.BringYourDeads[bringyourdead]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBringYourDead(bringyourdead)
		}
	}
	return bringyourdead
}

// for satisfaction of GongStruct interface
func (bringyourdead *BringYourDead) GetName() (res string) {
	return bringyourdead.Name
}

// Stage puts document to the model stage
func (document *Document) Stage(stage *StageStruct) *Document {
	stage.Documents[document] = __member
	stage.Documents_mapString[document.Name] = document

	return document
}

// Unstage removes document off the model stage
func (document *Document) Unstage(stage *StageStruct) *Document {
	delete(stage.Documents, document)
	delete(stage.Documents_mapString, document.Name)
	return document
}

// UnstageVoid removes document off the model stage
func (document *Document) UnstageVoid(stage *StageStruct) {
	delete(stage.Documents, document)
	delete(stage.Documents_mapString, document.Name)
}

// commit document to the back repo (if it is already staged)
func (document *Document) Commit(stage *StageStruct) *Document {
	if _, ok := stage.Documents[document]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocument(document)
		}
	}
	return document
}

func (document *Document) CommitVoid(stage *StageStruct) {
	document.Commit(stage)
}

// Checkout document to the back repo (if it is already staged)
func (document *Document) Checkout(stage *StageStruct) *Document {
	if _, ok := stage.Documents[document]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocument(document)
		}
	}
	return document
}

// for satisfaction of GongStruct interface
func (document *Document) GetName() (res string) {
	return document.Name
}

// Stage puts documentuse to the model stage
func (documentuse *DocumentUse) Stage(stage *StageStruct) *DocumentUse {
	stage.DocumentUses[documentuse] = __member
	stage.DocumentUses_mapString[documentuse.Name] = documentuse

	return documentuse
}

// Unstage removes documentuse off the model stage
func (documentuse *DocumentUse) Unstage(stage *StageStruct) *DocumentUse {
	delete(stage.DocumentUses, documentuse)
	delete(stage.DocumentUses_mapString, documentuse.Name)
	return documentuse
}

// UnstageVoid removes documentuse off the model stage
func (documentuse *DocumentUse) UnstageVoid(stage *StageStruct) {
	delete(stage.DocumentUses, documentuse)
	delete(stage.DocumentUses_mapString, documentuse.Name)
}

// commit documentuse to the back repo (if it is already staged)
func (documentuse *DocumentUse) Commit(stage *StageStruct) *DocumentUse {
	if _, ok := stage.DocumentUses[documentuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocumentUse(documentuse)
		}
	}
	return documentuse
}

func (documentuse *DocumentUse) CommitVoid(stage *StageStruct) {
	documentuse.Commit(stage)
}

// Checkout documentuse to the back repo (if it is already staged)
func (documentuse *DocumentUse) Checkout(stage *StageStruct) *DocumentUse {
	if _, ok := stage.DocumentUses[documentuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocumentUse(documentuse)
		}
	}
	return documentuse
}

// for satisfaction of GongStruct interface
func (documentuse *DocumentUse) GetName() (res string) {
	return documentuse.Name
}

// Stage puts galahadthepure to the model stage
func (galahadthepure *GalahadThePure) Stage(stage *StageStruct) *GalahadThePure {
	stage.GalahadThePures[galahadthepure] = __member
	stage.GalahadThePures_mapString[galahadthepure.Name] = galahadthepure

	return galahadthepure
}

// Unstage removes galahadthepure off the model stage
func (galahadthepure *GalahadThePure) Unstage(stage *StageStruct) *GalahadThePure {
	delete(stage.GalahadThePures, galahadthepure)
	delete(stage.GalahadThePures_mapString, galahadthepure.Name)
	return galahadthepure
}

// UnstageVoid removes galahadthepure off the model stage
func (galahadthepure *GalahadThePure) UnstageVoid(stage *StageStruct) {
	delete(stage.GalahadThePures, galahadthepure)
	delete(stage.GalahadThePures_mapString, galahadthepure.Name)
}

// commit galahadthepure to the back repo (if it is already staged)
func (galahadthepure *GalahadThePure) Commit(stage *StageStruct) *GalahadThePure {
	if _, ok := stage.GalahadThePures[galahadthepure]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGalahadThePure(galahadthepure)
		}
	}
	return galahadthepure
}

func (galahadthepure *GalahadThePure) CommitVoid(stage *StageStruct) {
	galahadthepure.Commit(stage)
}

// Checkout galahadthepure to the back repo (if it is already staged)
func (galahadthepure *GalahadThePure) Checkout(stage *StageStruct) *GalahadThePure {
	if _, ok := stage.GalahadThePures[galahadthepure]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGalahadThePure(galahadthepure)
		}
	}
	return galahadthepure
}

// for satisfaction of GongStruct interface
func (galahadthepure *GalahadThePure) GetName() (res string) {
	return galahadthepure.Name
}

// Stage puts geoobject to the model stage
func (geoobject *GeoObject) Stage(stage *StageStruct) *GeoObject {
	stage.GeoObjects[geoobject] = __member
	stage.GeoObjects_mapString[geoobject.Name] = geoobject

	return geoobject
}

// Unstage removes geoobject off the model stage
func (geoobject *GeoObject) Unstage(stage *StageStruct) *GeoObject {
	delete(stage.GeoObjects, geoobject)
	delete(stage.GeoObjects_mapString, geoobject.Name)
	return geoobject
}

// UnstageVoid removes geoobject off the model stage
func (geoobject *GeoObject) UnstageVoid(stage *StageStruct) {
	delete(stage.GeoObjects, geoobject)
	delete(stage.GeoObjects_mapString, geoobject.Name)
}

// commit geoobject to the back repo (if it is already staged)
func (geoobject *GeoObject) Commit(stage *StageStruct) *GeoObject {
	if _, ok := stage.GeoObjects[geoobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGeoObject(geoobject)
		}
	}
	return geoobject
}

func (geoobject *GeoObject) CommitVoid(stage *StageStruct) {
	geoobject.Commit(stage)
}

// Checkout geoobject to the back repo (if it is already staged)
func (geoobject *GeoObject) Checkout(stage *StageStruct) *GeoObject {
	if _, ok := stage.GeoObjects[geoobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGeoObject(geoobject)
		}
	}
	return geoobject
}

// for satisfaction of GongStruct interface
func (geoobject *GeoObject) GetName() (res string) {
	return geoobject.Name
}

// Stage puts geoobjectuse to the model stage
func (geoobjectuse *GeoObjectUse) Stage(stage *StageStruct) *GeoObjectUse {
	stage.GeoObjectUses[geoobjectuse] = __member
	stage.GeoObjectUses_mapString[geoobjectuse.Name] = geoobjectuse

	return geoobjectuse
}

// Unstage removes geoobjectuse off the model stage
func (geoobjectuse *GeoObjectUse) Unstage(stage *StageStruct) *GeoObjectUse {
	delete(stage.GeoObjectUses, geoobjectuse)
	delete(stage.GeoObjectUses_mapString, geoobjectuse.Name)
	return geoobjectuse
}

// UnstageVoid removes geoobjectuse off the model stage
func (geoobjectuse *GeoObjectUse) UnstageVoid(stage *StageStruct) {
	delete(stage.GeoObjectUses, geoobjectuse)
	delete(stage.GeoObjectUses_mapString, geoobjectuse.Name)
}

// commit geoobjectuse to the back repo (if it is already staged)
func (geoobjectuse *GeoObjectUse) Commit(stage *StageStruct) *GeoObjectUse {
	if _, ok := stage.GeoObjectUses[geoobjectuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGeoObjectUse(geoobjectuse)
		}
	}
	return geoobjectuse
}

func (geoobjectuse *GeoObjectUse) CommitVoid(stage *StageStruct) {
	geoobjectuse.Commit(stage)
}

// Checkout geoobjectuse to the back repo (if it is already staged)
func (geoobjectuse *GeoObjectUse) Checkout(stage *StageStruct) *GeoObjectUse {
	if _, ok := stage.GeoObjectUses[geoobjectuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGeoObjectUse(geoobjectuse)
		}
	}
	return geoobjectuse
}

// for satisfaction of GongStruct interface
func (geoobjectuse *GeoObjectUse) GetName() (res string) {
	return geoobjectuse.Name
}

// Stage puts group to the model stage
func (group *Group) Stage(stage *StageStruct) *Group {
	stage.Groups[group] = __member
	stage.Groups_mapString[group.Name] = group

	return group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *StageStruct) *Group {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *StageStruct) {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
}

// commit group to the back repo (if it is already staged)
func (group *Group) Commit(stage *StageStruct) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroup(group)
		}
	}
	return group
}

func (group *Group) CommitVoid(stage *StageStruct) {
	group.Commit(stage)
}

// Checkout group to the back repo (if it is already staged)
func (group *Group) Checkout(stage *StageStruct) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGroup(group)
		}
	}
	return group
}

// for satisfaction of GongStruct interface
func (group *Group) GetName() (res string) {
	return group.Name
}

// Stage puts groupuse to the model stage
func (groupuse *GroupUse) Stage(stage *StageStruct) *GroupUse {
	stage.GroupUses[groupuse] = __member
	stage.GroupUses_mapString[groupuse.Name] = groupuse

	return groupuse
}

// Unstage removes groupuse off the model stage
func (groupuse *GroupUse) Unstage(stage *StageStruct) *GroupUse {
	delete(stage.GroupUses, groupuse)
	delete(stage.GroupUses_mapString, groupuse.Name)
	return groupuse
}

// UnstageVoid removes groupuse off the model stage
func (groupuse *GroupUse) UnstageVoid(stage *StageStruct) {
	delete(stage.GroupUses, groupuse)
	delete(stage.GroupUses_mapString, groupuse.Name)
}

// commit groupuse to the back repo (if it is already staged)
func (groupuse *GroupUse) Commit(stage *StageStruct) *GroupUse {
	if _, ok := stage.GroupUses[groupuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroupUse(groupuse)
		}
	}
	return groupuse
}

func (groupuse *GroupUse) CommitVoid(stage *StageStruct) {
	groupuse.Commit(stage)
}

// Checkout groupuse to the back repo (if it is already staged)
func (groupuse *GroupUse) Checkout(stage *StageStruct) *GroupUse {
	if _, ok := stage.GroupUses[groupuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGroupUse(groupuse)
		}
	}
	return groupuse
}

// for satisfaction of GongStruct interface
func (groupuse *GroupUse) GetName() (res string) {
	return groupuse.Name
}

// Stage puts kingarthur to the model stage
func (kingarthur *KingArthur) Stage(stage *StageStruct) *KingArthur {
	stage.KingArthurs[kingarthur] = __member
	stage.KingArthurs_mapString[kingarthur.Name] = kingarthur

	return kingarthur
}

// Unstage removes kingarthur off the model stage
func (kingarthur *KingArthur) Unstage(stage *StageStruct) *KingArthur {
	delete(stage.KingArthurs, kingarthur)
	delete(stage.KingArthurs_mapString, kingarthur.Name)
	return kingarthur
}

// UnstageVoid removes kingarthur off the model stage
func (kingarthur *KingArthur) UnstageVoid(stage *StageStruct) {
	delete(stage.KingArthurs, kingarthur)
	delete(stage.KingArthurs_mapString, kingarthur.Name)
}

// commit kingarthur to the back repo (if it is already staged)
func (kingarthur *KingArthur) Commit(stage *StageStruct) *KingArthur {
	if _, ok := stage.KingArthurs[kingarthur]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitKingArthur(kingarthur)
		}
	}
	return kingarthur
}

func (kingarthur *KingArthur) CommitVoid(stage *StageStruct) {
	kingarthur.Commit(stage)
}

// Checkout kingarthur to the back repo (if it is already staged)
func (kingarthur *KingArthur) Checkout(stage *StageStruct) *KingArthur {
	if _, ok := stage.KingArthurs[kingarthur]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutKingArthur(kingarthur)
		}
	}
	return kingarthur
}

// for satisfaction of GongStruct interface
func (kingarthur *KingArthur) GetName() (res string) {
	return kingarthur.Name
}

// Stage puts kingarthurshape to the model stage
func (kingarthurshape *KingArthurShape) Stage(stage *StageStruct) *KingArthurShape {
	stage.KingArthurShapes[kingarthurshape] = __member
	stage.KingArthurShapes_mapString[kingarthurshape.Name] = kingarthurshape

	return kingarthurshape
}

// Unstage removes kingarthurshape off the model stage
func (kingarthurshape *KingArthurShape) Unstage(stage *StageStruct) *KingArthurShape {
	delete(stage.KingArthurShapes, kingarthurshape)
	delete(stage.KingArthurShapes_mapString, kingarthurshape.Name)
	return kingarthurshape
}

// UnstageVoid removes kingarthurshape off the model stage
func (kingarthurshape *KingArthurShape) UnstageVoid(stage *StageStruct) {
	delete(stage.KingArthurShapes, kingarthurshape)
	delete(stage.KingArthurShapes_mapString, kingarthurshape.Name)
}

// commit kingarthurshape to the back repo (if it is already staged)
func (kingarthurshape *KingArthurShape) Commit(stage *StageStruct) *KingArthurShape {
	if _, ok := stage.KingArthurShapes[kingarthurshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitKingArthurShape(kingarthurshape)
		}
	}
	return kingarthurshape
}

func (kingarthurshape *KingArthurShape) CommitVoid(stage *StageStruct) {
	kingarthurshape.Commit(stage)
}

// Checkout kingarthurshape to the back repo (if it is already staged)
func (kingarthurshape *KingArthurShape) Checkout(stage *StageStruct) *KingArthurShape {
	if _, ok := stage.KingArthurShapes[kingarthurshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutKingArthurShape(kingarthurshape)
		}
	}
	return kingarthurshape
}

// for satisfaction of GongStruct interface
func (kingarthurshape *KingArthurShape) GetName() (res string) {
	return kingarthurshape.Name
}

// Stage puts knightwhosayni to the model stage
func (knightwhosayni *KnightWhoSayNi) Stage(stage *StageStruct) *KnightWhoSayNi {
	stage.KnightWhoSayNis[knightwhosayni] = __member
	stage.KnightWhoSayNis_mapString[knightwhosayni.Name] = knightwhosayni

	return knightwhosayni
}

// Unstage removes knightwhosayni off the model stage
func (knightwhosayni *KnightWhoSayNi) Unstage(stage *StageStruct) *KnightWhoSayNi {
	delete(stage.KnightWhoSayNis, knightwhosayni)
	delete(stage.KnightWhoSayNis_mapString, knightwhosayni.Name)
	return knightwhosayni
}

// UnstageVoid removes knightwhosayni off the model stage
func (knightwhosayni *KnightWhoSayNi) UnstageVoid(stage *StageStruct) {
	delete(stage.KnightWhoSayNis, knightwhosayni)
	delete(stage.KnightWhoSayNis_mapString, knightwhosayni.Name)
}

// commit knightwhosayni to the back repo (if it is already staged)
func (knightwhosayni *KnightWhoSayNi) Commit(stage *StageStruct) *KnightWhoSayNi {
	if _, ok := stage.KnightWhoSayNis[knightwhosayni]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitKnightWhoSayNi(knightwhosayni)
		}
	}
	return knightwhosayni
}

func (knightwhosayni *KnightWhoSayNi) CommitVoid(stage *StageStruct) {
	knightwhosayni.Commit(stage)
}

// Checkout knightwhosayni to the back repo (if it is already staged)
func (knightwhosayni *KnightWhoSayNi) Checkout(stage *StageStruct) *KnightWhoSayNi {
	if _, ok := stage.KnightWhoSayNis[knightwhosayni]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutKnightWhoSayNi(knightwhosayni)
		}
	}
	return knightwhosayni
}

// for satisfaction of GongStruct interface
func (knightwhosayni *KnightWhoSayNi) GetName() (res string) {
	return knightwhosayni.Name
}

// Stage puts lancelot to the model stage
func (lancelot *Lancelot) Stage(stage *StageStruct) *Lancelot {
	stage.Lancelots[lancelot] = __member
	stage.Lancelots_mapString[lancelot.Name] = lancelot

	return lancelot
}

// Unstage removes lancelot off the model stage
func (lancelot *Lancelot) Unstage(stage *StageStruct) *Lancelot {
	delete(stage.Lancelots, lancelot)
	delete(stage.Lancelots_mapString, lancelot.Name)
	return lancelot
}

// UnstageVoid removes lancelot off the model stage
func (lancelot *Lancelot) UnstageVoid(stage *StageStruct) {
	delete(stage.Lancelots, lancelot)
	delete(stage.Lancelots_mapString, lancelot.Name)
}

// commit lancelot to the back repo (if it is already staged)
func (lancelot *Lancelot) Commit(stage *StageStruct) *Lancelot {
	if _, ok := stage.Lancelots[lancelot]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLancelot(lancelot)
		}
	}
	return lancelot
}

func (lancelot *Lancelot) CommitVoid(stage *StageStruct) {
	lancelot.Commit(stage)
}

// Checkout lancelot to the back repo (if it is already staged)
func (lancelot *Lancelot) Checkout(stage *StageStruct) *Lancelot {
	if _, ok := stage.Lancelots[lancelot]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLancelot(lancelot)
		}
	}
	return lancelot
}

// for satisfaction of GongStruct interface
func (lancelot *Lancelot) GetName() (res string) {
	return lancelot.Name
}

// Stage puts lancelotagregation to the model stage
func (lancelotagregation *LancelotAgregation) Stage(stage *StageStruct) *LancelotAgregation {
	stage.LancelotAgregations[lancelotagregation] = __member
	stage.LancelotAgregations_mapString[lancelotagregation.Name] = lancelotagregation

	return lancelotagregation
}

// Unstage removes lancelotagregation off the model stage
func (lancelotagregation *LancelotAgregation) Unstage(stage *StageStruct) *LancelotAgregation {
	delete(stage.LancelotAgregations, lancelotagregation)
	delete(stage.LancelotAgregations_mapString, lancelotagregation.Name)
	return lancelotagregation
}

// UnstageVoid removes lancelotagregation off the model stage
func (lancelotagregation *LancelotAgregation) UnstageVoid(stage *StageStruct) {
	delete(stage.LancelotAgregations, lancelotagregation)
	delete(stage.LancelotAgregations_mapString, lancelotagregation.Name)
}

// commit lancelotagregation to the back repo (if it is already staged)
func (lancelotagregation *LancelotAgregation) Commit(stage *StageStruct) *LancelotAgregation {
	if _, ok := stage.LancelotAgregations[lancelotagregation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLancelotAgregation(lancelotagregation)
		}
	}
	return lancelotagregation
}

func (lancelotagregation *LancelotAgregation) CommitVoid(stage *StageStruct) {
	lancelotagregation.Commit(stage)
}

// Checkout lancelotagregation to the back repo (if it is already staged)
func (lancelotagregation *LancelotAgregation) Checkout(stage *StageStruct) *LancelotAgregation {
	if _, ok := stage.LancelotAgregations[lancelotagregation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLancelotAgregation(lancelotagregation)
		}
	}
	return lancelotagregation
}

// for satisfaction of GongStruct interface
func (lancelotagregation *LancelotAgregation) GetName() (res string) {
	return lancelotagregation.Name
}

// Stage puts lancelotagregationuse to the model stage
func (lancelotagregationuse *LancelotAgregationUse) Stage(stage *StageStruct) *LancelotAgregationUse {
	stage.LancelotAgregationUses[lancelotagregationuse] = __member
	stage.LancelotAgregationUses_mapString[lancelotagregationuse.Name] = lancelotagregationuse

	return lancelotagregationuse
}

// Unstage removes lancelotagregationuse off the model stage
func (lancelotagregationuse *LancelotAgregationUse) Unstage(stage *StageStruct) *LancelotAgregationUse {
	delete(stage.LancelotAgregationUses, lancelotagregationuse)
	delete(stage.LancelotAgregationUses_mapString, lancelotagregationuse.Name)
	return lancelotagregationuse
}

// UnstageVoid removes lancelotagregationuse off the model stage
func (lancelotagregationuse *LancelotAgregationUse) UnstageVoid(stage *StageStruct) {
	delete(stage.LancelotAgregationUses, lancelotagregationuse)
	delete(stage.LancelotAgregationUses_mapString, lancelotagregationuse.Name)
}

// commit lancelotagregationuse to the back repo (if it is already staged)
func (lancelotagregationuse *LancelotAgregationUse) Commit(stage *StageStruct) *LancelotAgregationUse {
	if _, ok := stage.LancelotAgregationUses[lancelotagregationuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLancelotAgregationUse(lancelotagregationuse)
		}
	}
	return lancelotagregationuse
}

func (lancelotagregationuse *LancelotAgregationUse) CommitVoid(stage *StageStruct) {
	lancelotagregationuse.Commit(stage)
}

// Checkout lancelotagregationuse to the back repo (if it is already staged)
func (lancelotagregationuse *LancelotAgregationUse) Checkout(stage *StageStruct) *LancelotAgregationUse {
	if _, ok := stage.LancelotAgregationUses[lancelotagregationuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLancelotAgregationUse(lancelotagregationuse)
		}
	}
	return lancelotagregationuse
}

// for satisfaction of GongStruct interface
func (lancelotagregationuse *LancelotAgregationUse) GetName() (res string) {
	return lancelotagregationuse.Name
}

// Stage puts lancelotcategory to the model stage
func (lancelotcategory *LancelotCategory) Stage(stage *StageStruct) *LancelotCategory {
	stage.LancelotCategorys[lancelotcategory] = __member
	stage.LancelotCategorys_mapString[lancelotcategory.Name] = lancelotcategory

	return lancelotcategory
}

// Unstage removes lancelotcategory off the model stage
func (lancelotcategory *LancelotCategory) Unstage(stage *StageStruct) *LancelotCategory {
	delete(stage.LancelotCategorys, lancelotcategory)
	delete(stage.LancelotCategorys_mapString, lancelotcategory.Name)
	return lancelotcategory
}

// UnstageVoid removes lancelotcategory off the model stage
func (lancelotcategory *LancelotCategory) UnstageVoid(stage *StageStruct) {
	delete(stage.LancelotCategorys, lancelotcategory)
	delete(stage.LancelotCategorys_mapString, lancelotcategory.Name)
}

// commit lancelotcategory to the back repo (if it is already staged)
func (lancelotcategory *LancelotCategory) Commit(stage *StageStruct) *LancelotCategory {
	if _, ok := stage.LancelotCategorys[lancelotcategory]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLancelotCategory(lancelotcategory)
		}
	}
	return lancelotcategory
}

func (lancelotcategory *LancelotCategory) CommitVoid(stage *StageStruct) {
	lancelotcategory.Commit(stage)
}

// Checkout lancelotcategory to the back repo (if it is already staged)
func (lancelotcategory *LancelotCategory) Checkout(stage *StageStruct) *LancelotCategory {
	if _, ok := stage.LancelotCategorys[lancelotcategory]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLancelotCategory(lancelotcategory)
		}
	}
	return lancelotcategory
}

// for satisfaction of GongStruct interface
func (lancelotcategory *LancelotCategory) GetName() (res string) {
	return lancelotcategory.Name
}

// Stage puts lancelotcategoryuse to the model stage
func (lancelotcategoryuse *LancelotCategoryUse) Stage(stage *StageStruct) *LancelotCategoryUse {
	stage.LancelotCategoryUses[lancelotcategoryuse] = __member
	stage.LancelotCategoryUses_mapString[lancelotcategoryuse.Name] = lancelotcategoryuse

	return lancelotcategoryuse
}

// Unstage removes lancelotcategoryuse off the model stage
func (lancelotcategoryuse *LancelotCategoryUse) Unstage(stage *StageStruct) *LancelotCategoryUse {
	delete(stage.LancelotCategoryUses, lancelotcategoryuse)
	delete(stage.LancelotCategoryUses_mapString, lancelotcategoryuse.Name)
	return lancelotcategoryuse
}

// UnstageVoid removes lancelotcategoryuse off the model stage
func (lancelotcategoryuse *LancelotCategoryUse) UnstageVoid(stage *StageStruct) {
	delete(stage.LancelotCategoryUses, lancelotcategoryuse)
	delete(stage.LancelotCategoryUses_mapString, lancelotcategoryuse.Name)
}

// commit lancelotcategoryuse to the back repo (if it is already staged)
func (lancelotcategoryuse *LancelotCategoryUse) Commit(stage *StageStruct) *LancelotCategoryUse {
	if _, ok := stage.LancelotCategoryUses[lancelotcategoryuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLancelotCategoryUse(lancelotcategoryuse)
		}
	}
	return lancelotcategoryuse
}

func (lancelotcategoryuse *LancelotCategoryUse) CommitVoid(stage *StageStruct) {
	lancelotcategoryuse.Commit(stage)
}

// Checkout lancelotcategoryuse to the back repo (if it is already staged)
func (lancelotcategoryuse *LancelotCategoryUse) Checkout(stage *StageStruct) *LancelotCategoryUse {
	if _, ok := stage.LancelotCategoryUses[lancelotcategoryuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLancelotCategoryUse(lancelotcategoryuse)
		}
	}
	return lancelotcategoryuse
}

// for satisfaction of GongStruct interface
func (lancelotcategoryuse *LancelotCategoryUse) GetName() (res string) {
	return lancelotcategoryuse.Name
}

// Stage puts mapobject to the model stage
func (mapobject *MapObject) Stage(stage *StageStruct) *MapObject {
	stage.MapObjects[mapobject] = __member
	stage.MapObjects_mapString[mapobject.Name] = mapobject

	return mapobject
}

// Unstage removes mapobject off the model stage
func (mapobject *MapObject) Unstage(stage *StageStruct) *MapObject {
	delete(stage.MapObjects, mapobject)
	delete(stage.MapObjects_mapString, mapobject.Name)
	return mapobject
}

// UnstageVoid removes mapobject off the model stage
func (mapobject *MapObject) UnstageVoid(stage *StageStruct) {
	delete(stage.MapObjects, mapobject)
	delete(stage.MapObjects_mapString, mapobject.Name)
}

// commit mapobject to the back repo (if it is already staged)
func (mapobject *MapObject) Commit(stage *StageStruct) *MapObject {
	if _, ok := stage.MapObjects[mapobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMapObject(mapobject)
		}
	}
	return mapobject
}

func (mapobject *MapObject) CommitVoid(stage *StageStruct) {
	mapobject.Commit(stage)
}

// Checkout mapobject to the back repo (if it is already staged)
func (mapobject *MapObject) Checkout(stage *StageStruct) *MapObject {
	if _, ok := stage.MapObjects[mapobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMapObject(mapobject)
		}
	}
	return mapobject
}

// for satisfaction of GongStruct interface
func (mapobject *MapObject) GetName() (res string) {
	return mapobject.Name
}

// Stage puts mapobjectuse to the model stage
func (mapobjectuse *MapObjectUse) Stage(stage *StageStruct) *MapObjectUse {
	stage.MapObjectUses[mapobjectuse] = __member
	stage.MapObjectUses_mapString[mapobjectuse.Name] = mapobjectuse

	return mapobjectuse
}

// Unstage removes mapobjectuse off the model stage
func (mapobjectuse *MapObjectUse) Unstage(stage *StageStruct) *MapObjectUse {
	delete(stage.MapObjectUses, mapobjectuse)
	delete(stage.MapObjectUses_mapString, mapobjectuse.Name)
	return mapobjectuse
}

// UnstageVoid removes mapobjectuse off the model stage
func (mapobjectuse *MapObjectUse) UnstageVoid(stage *StageStruct) {
	delete(stage.MapObjectUses, mapobjectuse)
	delete(stage.MapObjectUses_mapString, mapobjectuse.Name)
}

// commit mapobjectuse to the back repo (if it is already staged)
func (mapobjectuse *MapObjectUse) Commit(stage *StageStruct) *MapObjectUse {
	if _, ok := stage.MapObjectUses[mapobjectuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMapObjectUse(mapobjectuse)
		}
	}
	return mapobjectuse
}

func (mapobjectuse *MapObjectUse) CommitVoid(stage *StageStruct) {
	mapobjectuse.Commit(stage)
}

// Checkout mapobjectuse to the back repo (if it is already staged)
func (mapobjectuse *MapObjectUse) Checkout(stage *StageStruct) *MapObjectUse {
	if _, ok := stage.MapObjectUses[mapobjectuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMapObjectUse(mapobjectuse)
		}
	}
	return mapobjectuse
}

// for satisfaction of GongStruct interface
func (mapobjectuse *MapObjectUse) GetName() (res string) {
	return mapobjectuse.Name
}

// Stage puts repository to the model stage
func (repository *Repository) Stage(stage *StageStruct) *Repository {
	stage.Repositorys[repository] = __member
	stage.Repositorys_mapString[repository.Name] = repository

	return repository
}

// Unstage removes repository off the model stage
func (repository *Repository) Unstage(stage *StageStruct) *Repository {
	delete(stage.Repositorys, repository)
	delete(stage.Repositorys_mapString, repository.Name)
	return repository
}

// UnstageVoid removes repository off the model stage
func (repository *Repository) UnstageVoid(stage *StageStruct) {
	delete(stage.Repositorys, repository)
	delete(stage.Repositorys_mapString, repository.Name)
}

// commit repository to the back repo (if it is already staged)
func (repository *Repository) Commit(stage *StageStruct) *Repository {
	if _, ok := stage.Repositorys[repository]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRepository(repository)
		}
	}
	return repository
}

func (repository *Repository) CommitVoid(stage *StageStruct) {
	repository.Commit(stage)
}

// Checkout repository to the back repo (if it is already staged)
func (repository *Repository) Checkout(stage *StageStruct) *Repository {
	if _, ok := stage.Repositorys[repository]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRepository(repository)
		}
	}
	return repository
}

// for satisfaction of GongStruct interface
func (repository *Repository) GetName() (res string) {
	return repository.Name
}

// Stage puts sirrobin to the model stage
func (sirrobin *SirRobin) Stage(stage *StageStruct) *SirRobin {
	stage.SirRobins[sirrobin] = __member
	stage.SirRobins_mapString[sirrobin.Name] = sirrobin

	return sirrobin
}

// Unstage removes sirrobin off the model stage
func (sirrobin *SirRobin) Unstage(stage *StageStruct) *SirRobin {
	delete(stage.SirRobins, sirrobin)
	delete(stage.SirRobins_mapString, sirrobin.Name)
	return sirrobin
}

// UnstageVoid removes sirrobin off the model stage
func (sirrobin *SirRobin) UnstageVoid(stage *StageStruct) {
	delete(stage.SirRobins, sirrobin)
	delete(stage.SirRobins_mapString, sirrobin.Name)
}

// commit sirrobin to the back repo (if it is already staged)
func (sirrobin *SirRobin) Commit(stage *StageStruct) *SirRobin {
	if _, ok := stage.SirRobins[sirrobin]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSirRobin(sirrobin)
		}
	}
	return sirrobin
}

func (sirrobin *SirRobin) CommitVoid(stage *StageStruct) {
	sirrobin.Commit(stage)
}

// Checkout sirrobin to the back repo (if it is already staged)
func (sirrobin *SirRobin) Checkout(stage *StageStruct) *SirRobin {
	if _, ok := stage.SirRobins[sirrobin]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSirRobin(sirrobin)
		}
	}
	return sirrobin
}

// for satisfaction of GongStruct interface
func (sirrobin *SirRobin) GetName() (res string) {
	return sirrobin.Name
}

// Stage puts thebridge to the model stage
func (thebridge *TheBridge) Stage(stage *StageStruct) *TheBridge {
	stage.TheBridges[thebridge] = __member
	stage.TheBridges_mapString[thebridge.Name] = thebridge

	return thebridge
}

// Unstage removes thebridge off the model stage
func (thebridge *TheBridge) Unstage(stage *StageStruct) *TheBridge {
	delete(stage.TheBridges, thebridge)
	delete(stage.TheBridges_mapString, thebridge.Name)
	return thebridge
}

// UnstageVoid removes thebridge off the model stage
func (thebridge *TheBridge) UnstageVoid(stage *StageStruct) {
	delete(stage.TheBridges, thebridge)
	delete(stage.TheBridges_mapString, thebridge.Name)
}

// commit thebridge to the back repo (if it is already staged)
func (thebridge *TheBridge) Commit(stage *StageStruct) *TheBridge {
	if _, ok := stage.TheBridges[thebridge]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTheBridge(thebridge)
		}
	}
	return thebridge
}

func (thebridge *TheBridge) CommitVoid(stage *StageStruct) {
	thebridge.Commit(stage)
}

// Checkout thebridge to the back repo (if it is already staged)
func (thebridge *TheBridge) Checkout(stage *StageStruct) *TheBridge {
	if _, ok := stage.TheBridges[thebridge]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTheBridge(thebridge)
		}
	}
	return thebridge
}

// for satisfaction of GongStruct interface
func (thebridge *TheBridge) GetName() (res string) {
	return thebridge.Name
}

// Stage puts thenuteshape to the model stage
func (thenuteshape *TheNuteShape) Stage(stage *StageStruct) *TheNuteShape {
	stage.TheNuteShapes[thenuteshape] = __member
	stage.TheNuteShapes_mapString[thenuteshape.Name] = thenuteshape

	return thenuteshape
}

// Unstage removes thenuteshape off the model stage
func (thenuteshape *TheNuteShape) Unstage(stage *StageStruct) *TheNuteShape {
	delete(stage.TheNuteShapes, thenuteshape)
	delete(stage.TheNuteShapes_mapString, thenuteshape.Name)
	return thenuteshape
}

// UnstageVoid removes thenuteshape off the model stage
func (thenuteshape *TheNuteShape) UnstageVoid(stage *StageStruct) {
	delete(stage.TheNuteShapes, thenuteshape)
	delete(stage.TheNuteShapes_mapString, thenuteshape.Name)
}

// commit thenuteshape to the back repo (if it is already staged)
func (thenuteshape *TheNuteShape) Commit(stage *StageStruct) *TheNuteShape {
	if _, ok := stage.TheNuteShapes[thenuteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTheNuteShape(thenuteshape)
		}
	}
	return thenuteshape
}

func (thenuteshape *TheNuteShape) CommitVoid(stage *StageStruct) {
	thenuteshape.Commit(stage)
}

// Checkout thenuteshape to the back repo (if it is already staged)
func (thenuteshape *TheNuteShape) Checkout(stage *StageStruct) *TheNuteShape {
	if _, ok := stage.TheNuteShapes[thenuteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTheNuteShape(thenuteshape)
		}
	}
	return thenuteshape
}

// for satisfaction of GongStruct interface
func (thenuteshape *TheNuteShape) GetName() (res string) {
	return thenuteshape.Name
}

// Stage puts thenutetransition to the model stage
func (thenutetransition *TheNuteTransition) Stage(stage *StageStruct) *TheNuteTransition {
	stage.TheNuteTransitions[thenutetransition] = __member
	stage.TheNuteTransitions_mapString[thenutetransition.Name] = thenutetransition

	return thenutetransition
}

// Unstage removes thenutetransition off the model stage
func (thenutetransition *TheNuteTransition) Unstage(stage *StageStruct) *TheNuteTransition {
	delete(stage.TheNuteTransitions, thenutetransition)
	delete(stage.TheNuteTransitions_mapString, thenutetransition.Name)
	return thenutetransition
}

// UnstageVoid removes thenutetransition off the model stage
func (thenutetransition *TheNuteTransition) UnstageVoid(stage *StageStruct) {
	delete(stage.TheNuteTransitions, thenutetransition)
	delete(stage.TheNuteTransitions_mapString, thenutetransition.Name)
}

// commit thenutetransition to the back repo (if it is already staged)
func (thenutetransition *TheNuteTransition) Commit(stage *StageStruct) *TheNuteTransition {
	if _, ok := stage.TheNuteTransitions[thenutetransition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTheNuteTransition(thenutetransition)
		}
	}
	return thenutetransition
}

func (thenutetransition *TheNuteTransition) CommitVoid(stage *StageStruct) {
	thenutetransition.Commit(stage)
}

// Checkout thenutetransition to the back repo (if it is already staged)
func (thenutetransition *TheNuteTransition) Checkout(stage *StageStruct) *TheNuteTransition {
	if _, ok := stage.TheNuteTransitions[thenutetransition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTheNuteTransition(thenutetransition)
		}
	}
	return thenutetransition
}

// for satisfaction of GongStruct interface
func (thenutetransition *TheNuteTransition) GetName() (res string) {
	return thenutetransition.Name
}

// Stage puts user to the model stage
func (user *User) Stage(stage *StageStruct) *User {
	stage.Users[user] = __member
	stage.Users_mapString[user.Name] = user

	return user
}

// Unstage removes user off the model stage
func (user *User) Unstage(stage *StageStruct) *User {
	delete(stage.Users, user)
	delete(stage.Users_mapString, user.Name)
	return user
}

// UnstageVoid removes user off the model stage
func (user *User) UnstageVoid(stage *StageStruct) {
	delete(stage.Users, user)
	delete(stage.Users_mapString, user.Name)
}

// commit user to the back repo (if it is already staged)
func (user *User) Commit(stage *StageStruct) *User {
	if _, ok := stage.Users[user]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUser(user)
		}
	}
	return user
}

func (user *User) CommitVoid(stage *StageStruct) {
	user.Commit(stage)
}

// Checkout user to the back repo (if it is already staged)
func (user *User) Checkout(stage *StageStruct) *User {
	if _, ok := stage.Users[user]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUser(user)
		}
	}
	return user
}

// for satisfaction of GongStruct interface
func (user *User) GetName() (res string) {
	return user.Name
}

// Stage puts useruse to the model stage
func (useruse *UserUse) Stage(stage *StageStruct) *UserUse {
	stage.UserUses[useruse] = __member
	stage.UserUses_mapString[useruse.Name] = useruse

	return useruse
}

// Unstage removes useruse off the model stage
func (useruse *UserUse) Unstage(stage *StageStruct) *UserUse {
	delete(stage.UserUses, useruse)
	delete(stage.UserUses_mapString, useruse.Name)
	return useruse
}

// UnstageVoid removes useruse off the model stage
func (useruse *UserUse) UnstageVoid(stage *StageStruct) {
	delete(stage.UserUses, useruse)
	delete(stage.UserUses_mapString, useruse.Name)
}

// commit useruse to the back repo (if it is already staged)
func (useruse *UserUse) Commit(stage *StageStruct) *UserUse {
	if _, ok := stage.UserUses[useruse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUserUse(useruse)
		}
	}
	return useruse
}

func (useruse *UserUse) CommitVoid(stage *StageStruct) {
	useruse.Commit(stage)
}

// Checkout useruse to the back repo (if it is already staged)
func (useruse *UserUse) Checkout(stage *StageStruct) *UserUse {
	if _, ok := stage.UserUses[useruse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUserUse(useruse)
		}
	}
	return useruse
}

// for satisfaction of GongStruct interface
func (useruse *UserUse) GetName() (res string) {
	return useruse.Name
}

// Stage puts whatisyourpreferedcolor to the model stage
func (whatisyourpreferedcolor *WhatIsYourPreferedColor) Stage(stage *StageStruct) *WhatIsYourPreferedColor {
	stage.WhatIsYourPreferedColors[whatisyourpreferedcolor] = __member
	stage.WhatIsYourPreferedColors_mapString[whatisyourpreferedcolor.Name] = whatisyourpreferedcolor

	return whatisyourpreferedcolor
}

// Unstage removes whatisyourpreferedcolor off the model stage
func (whatisyourpreferedcolor *WhatIsYourPreferedColor) Unstage(stage *StageStruct) *WhatIsYourPreferedColor {
	delete(stage.WhatIsYourPreferedColors, whatisyourpreferedcolor)
	delete(stage.WhatIsYourPreferedColors_mapString, whatisyourpreferedcolor.Name)
	return whatisyourpreferedcolor
}

// UnstageVoid removes whatisyourpreferedcolor off the model stage
func (whatisyourpreferedcolor *WhatIsYourPreferedColor) UnstageVoid(stage *StageStruct) {
	delete(stage.WhatIsYourPreferedColors, whatisyourpreferedcolor)
	delete(stage.WhatIsYourPreferedColors_mapString, whatisyourpreferedcolor.Name)
}

// commit whatisyourpreferedcolor to the back repo (if it is already staged)
func (whatisyourpreferedcolor *WhatIsYourPreferedColor) Commit(stage *StageStruct) *WhatIsYourPreferedColor {
	if _, ok := stage.WhatIsYourPreferedColors[whatisyourpreferedcolor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitWhatIsYourPreferedColor(whatisyourpreferedcolor)
		}
	}
	return whatisyourpreferedcolor
}

func (whatisyourpreferedcolor *WhatIsYourPreferedColor) CommitVoid(stage *StageStruct) {
	whatisyourpreferedcolor.Commit(stage)
}

// Checkout whatisyourpreferedcolor to the back repo (if it is already staged)
func (whatisyourpreferedcolor *WhatIsYourPreferedColor) Checkout(stage *StageStruct) *WhatIsYourPreferedColor {
	if _, ok := stage.WhatIsYourPreferedColors[whatisyourpreferedcolor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutWhatIsYourPreferedColor(whatisyourpreferedcolor)
		}
	}
	return whatisyourpreferedcolor
}

// for satisfaction of GongStruct interface
func (whatisyourpreferedcolor *WhatIsYourPreferedColor) GetName() (res string) {
	return whatisyourpreferedcolor.Name
}

// Stage puts workspace to the model stage
func (workspace *Workspace) Stage(stage *StageStruct) *Workspace {
	stage.Workspaces[workspace] = __member
	stage.Workspaces_mapString[workspace.Name] = workspace

	return workspace
}

// Unstage removes workspace off the model stage
func (workspace *Workspace) Unstage(stage *StageStruct) *Workspace {
	delete(stage.Workspaces, workspace)
	delete(stage.Workspaces_mapString, workspace.Name)
	return workspace
}

// UnstageVoid removes workspace off the model stage
func (workspace *Workspace) UnstageVoid(stage *StageStruct) {
	delete(stage.Workspaces, workspace)
	delete(stage.Workspaces_mapString, workspace.Name)
}

// commit workspace to the back repo (if it is already staged)
func (workspace *Workspace) Commit(stage *StageStruct) *Workspace {
	if _, ok := stage.Workspaces[workspace]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitWorkspace(workspace)
		}
	}
	return workspace
}

func (workspace *Workspace) CommitVoid(stage *StageStruct) {
	workspace.Commit(stage)
}

// Checkout workspace to the back repo (if it is already staged)
func (workspace *Workspace) Checkout(stage *StageStruct) *Workspace {
	if _, ok := stage.Workspaces[workspace]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutWorkspace(workspace)
		}
	}
	return workspace
}

// for satisfaction of GongStruct interface
func (workspace *Workspace) GetName() (res string) {
	return workspace.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAWitch(AWitch *AWitch)
	CreateORMBlackKnightShape(BlackKnightShape *BlackKnightShape)
	CreateORMBringYourDead(BringYourDead *BringYourDead)
	CreateORMDocument(Document *Document)
	CreateORMDocumentUse(DocumentUse *DocumentUse)
	CreateORMGalahadThePure(GalahadThePure *GalahadThePure)
	CreateORMGeoObject(GeoObject *GeoObject)
	CreateORMGeoObjectUse(GeoObjectUse *GeoObjectUse)
	CreateORMGroup(Group *Group)
	CreateORMGroupUse(GroupUse *GroupUse)
	CreateORMKingArthur(KingArthur *KingArthur)
	CreateORMKingArthurShape(KingArthurShape *KingArthurShape)
	CreateORMKnightWhoSayNi(KnightWhoSayNi *KnightWhoSayNi)
	CreateORMLancelot(Lancelot *Lancelot)
	CreateORMLancelotAgregation(LancelotAgregation *LancelotAgregation)
	CreateORMLancelotAgregationUse(LancelotAgregationUse *LancelotAgregationUse)
	CreateORMLancelotCategory(LancelotCategory *LancelotCategory)
	CreateORMLancelotCategoryUse(LancelotCategoryUse *LancelotCategoryUse)
	CreateORMMapObject(MapObject *MapObject)
	CreateORMMapObjectUse(MapObjectUse *MapObjectUse)
	CreateORMRepository(Repository *Repository)
	CreateORMSirRobin(SirRobin *SirRobin)
	CreateORMTheBridge(TheBridge *TheBridge)
	CreateORMTheNuteShape(TheNuteShape *TheNuteShape)
	CreateORMTheNuteTransition(TheNuteTransition *TheNuteTransition)
	CreateORMUser(User *User)
	CreateORMUserUse(UserUse *UserUse)
	CreateORMWhatIsYourPreferedColor(WhatIsYourPreferedColor *WhatIsYourPreferedColor)
	CreateORMWorkspace(Workspace *Workspace)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAWitch(AWitch *AWitch)
	DeleteORMBlackKnightShape(BlackKnightShape *BlackKnightShape)
	DeleteORMBringYourDead(BringYourDead *BringYourDead)
	DeleteORMDocument(Document *Document)
	DeleteORMDocumentUse(DocumentUse *DocumentUse)
	DeleteORMGalahadThePure(GalahadThePure *GalahadThePure)
	DeleteORMGeoObject(GeoObject *GeoObject)
	DeleteORMGeoObjectUse(GeoObjectUse *GeoObjectUse)
	DeleteORMGroup(Group *Group)
	DeleteORMGroupUse(GroupUse *GroupUse)
	DeleteORMKingArthur(KingArthur *KingArthur)
	DeleteORMKingArthurShape(KingArthurShape *KingArthurShape)
	DeleteORMKnightWhoSayNi(KnightWhoSayNi *KnightWhoSayNi)
	DeleteORMLancelot(Lancelot *Lancelot)
	DeleteORMLancelotAgregation(LancelotAgregation *LancelotAgregation)
	DeleteORMLancelotAgregationUse(LancelotAgregationUse *LancelotAgregationUse)
	DeleteORMLancelotCategory(LancelotCategory *LancelotCategory)
	DeleteORMLancelotCategoryUse(LancelotCategoryUse *LancelotCategoryUse)
	DeleteORMMapObject(MapObject *MapObject)
	DeleteORMMapObjectUse(MapObjectUse *MapObjectUse)
	DeleteORMRepository(Repository *Repository)
	DeleteORMSirRobin(SirRobin *SirRobin)
	DeleteORMTheBridge(TheBridge *TheBridge)
	DeleteORMTheNuteShape(TheNuteShape *TheNuteShape)
	DeleteORMTheNuteTransition(TheNuteTransition *TheNuteTransition)
	DeleteORMUser(User *User)
	DeleteORMUserUse(UserUse *UserUse)
	DeleteORMWhatIsYourPreferedColor(WhatIsYourPreferedColor *WhatIsYourPreferedColor)
	DeleteORMWorkspace(Workspace *Workspace)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.AWitchs = make(map[*AWitch]any)
	stage.AWitchs_mapString = make(map[string]*AWitch)

	stage.BlackKnightShapes = make(map[*BlackKnightShape]any)
	stage.BlackKnightShapes_mapString = make(map[string]*BlackKnightShape)

	stage.BringYourDeads = make(map[*BringYourDead]any)
	stage.BringYourDeads_mapString = make(map[string]*BringYourDead)

	stage.Documents = make(map[*Document]any)
	stage.Documents_mapString = make(map[string]*Document)

	stage.DocumentUses = make(map[*DocumentUse]any)
	stage.DocumentUses_mapString = make(map[string]*DocumentUse)

	stage.GalahadThePures = make(map[*GalahadThePure]any)
	stage.GalahadThePures_mapString = make(map[string]*GalahadThePure)

	stage.GeoObjects = make(map[*GeoObject]any)
	stage.GeoObjects_mapString = make(map[string]*GeoObject)

	stage.GeoObjectUses = make(map[*GeoObjectUse]any)
	stage.GeoObjectUses_mapString = make(map[string]*GeoObjectUse)

	stage.Groups = make(map[*Group]any)
	stage.Groups_mapString = make(map[string]*Group)

	stage.GroupUses = make(map[*GroupUse]any)
	stage.GroupUses_mapString = make(map[string]*GroupUse)

	stage.KingArthurs = make(map[*KingArthur]any)
	stage.KingArthurs_mapString = make(map[string]*KingArthur)

	stage.KingArthurShapes = make(map[*KingArthurShape]any)
	stage.KingArthurShapes_mapString = make(map[string]*KingArthurShape)

	stage.KnightWhoSayNis = make(map[*KnightWhoSayNi]any)
	stage.KnightWhoSayNis_mapString = make(map[string]*KnightWhoSayNi)

	stage.Lancelots = make(map[*Lancelot]any)
	stage.Lancelots_mapString = make(map[string]*Lancelot)

	stage.LancelotAgregations = make(map[*LancelotAgregation]any)
	stage.LancelotAgregations_mapString = make(map[string]*LancelotAgregation)

	stage.LancelotAgregationUses = make(map[*LancelotAgregationUse]any)
	stage.LancelotAgregationUses_mapString = make(map[string]*LancelotAgregationUse)

	stage.LancelotCategorys = make(map[*LancelotCategory]any)
	stage.LancelotCategorys_mapString = make(map[string]*LancelotCategory)

	stage.LancelotCategoryUses = make(map[*LancelotCategoryUse]any)
	stage.LancelotCategoryUses_mapString = make(map[string]*LancelotCategoryUse)

	stage.MapObjects = make(map[*MapObject]any)
	stage.MapObjects_mapString = make(map[string]*MapObject)

	stage.MapObjectUses = make(map[*MapObjectUse]any)
	stage.MapObjectUses_mapString = make(map[string]*MapObjectUse)

	stage.Repositorys = make(map[*Repository]any)
	stage.Repositorys_mapString = make(map[string]*Repository)

	stage.SirRobins = make(map[*SirRobin]any)
	stage.SirRobins_mapString = make(map[string]*SirRobin)

	stage.TheBridges = make(map[*TheBridge]any)
	stage.TheBridges_mapString = make(map[string]*TheBridge)

	stage.TheNuteShapes = make(map[*TheNuteShape]any)
	stage.TheNuteShapes_mapString = make(map[string]*TheNuteShape)

	stage.TheNuteTransitions = make(map[*TheNuteTransition]any)
	stage.TheNuteTransitions_mapString = make(map[string]*TheNuteTransition)

	stage.Users = make(map[*User]any)
	stage.Users_mapString = make(map[string]*User)

	stage.UserUses = make(map[*UserUse]any)
	stage.UserUses_mapString = make(map[string]*UserUse)

	stage.WhatIsYourPreferedColors = make(map[*WhatIsYourPreferedColor]any)
	stage.WhatIsYourPreferedColors_mapString = make(map[string]*WhatIsYourPreferedColor)

	stage.Workspaces = make(map[*Workspace]any)
	stage.Workspaces_mapString = make(map[string]*Workspace)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.AWitchs = nil
	stage.AWitchs_mapString = nil

	stage.BlackKnightShapes = nil
	stage.BlackKnightShapes_mapString = nil

	stage.BringYourDeads = nil
	stage.BringYourDeads_mapString = nil

	stage.Documents = nil
	stage.Documents_mapString = nil

	stage.DocumentUses = nil
	stage.DocumentUses_mapString = nil

	stage.GalahadThePures = nil
	stage.GalahadThePures_mapString = nil

	stage.GeoObjects = nil
	stage.GeoObjects_mapString = nil

	stage.GeoObjectUses = nil
	stage.GeoObjectUses_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

	stage.GroupUses = nil
	stage.GroupUses_mapString = nil

	stage.KingArthurs = nil
	stage.KingArthurs_mapString = nil

	stage.KingArthurShapes = nil
	stage.KingArthurShapes_mapString = nil

	stage.KnightWhoSayNis = nil
	stage.KnightWhoSayNis_mapString = nil

	stage.Lancelots = nil
	stage.Lancelots_mapString = nil

	stage.LancelotAgregations = nil
	stage.LancelotAgregations_mapString = nil

	stage.LancelotAgregationUses = nil
	stage.LancelotAgregationUses_mapString = nil

	stage.LancelotCategorys = nil
	stage.LancelotCategorys_mapString = nil

	stage.LancelotCategoryUses = nil
	stage.LancelotCategoryUses_mapString = nil

	stage.MapObjects = nil
	stage.MapObjects_mapString = nil

	stage.MapObjectUses = nil
	stage.MapObjectUses_mapString = nil

	stage.Repositorys = nil
	stage.Repositorys_mapString = nil

	stage.SirRobins = nil
	stage.SirRobins_mapString = nil

	stage.TheBridges = nil
	stage.TheBridges_mapString = nil

	stage.TheNuteShapes = nil
	stage.TheNuteShapes_mapString = nil

	stage.TheNuteTransitions = nil
	stage.TheNuteTransitions_mapString = nil

	stage.Users = nil
	stage.Users_mapString = nil

	stage.UserUses = nil
	stage.UserUses_mapString = nil

	stage.WhatIsYourPreferedColors = nil
	stage.WhatIsYourPreferedColors_mapString = nil

	stage.Workspaces = nil
	stage.Workspaces_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for awitch := range stage.AWitchs {
		awitch.Unstage(stage)
	}

	for blackknightshape := range stage.BlackKnightShapes {
		blackknightshape.Unstage(stage)
	}

	for bringyourdead := range stage.BringYourDeads {
		bringyourdead.Unstage(stage)
	}

	for document := range stage.Documents {
		document.Unstage(stage)
	}

	for documentuse := range stage.DocumentUses {
		documentuse.Unstage(stage)
	}

	for galahadthepure := range stage.GalahadThePures {
		galahadthepure.Unstage(stage)
	}

	for geoobject := range stage.GeoObjects {
		geoobject.Unstage(stage)
	}

	for geoobjectuse := range stage.GeoObjectUses {
		geoobjectuse.Unstage(stage)
	}

	for group := range stage.Groups {
		group.Unstage(stage)
	}

	for groupuse := range stage.GroupUses {
		groupuse.Unstage(stage)
	}

	for kingarthur := range stage.KingArthurs {
		kingarthur.Unstage(stage)
	}

	for kingarthurshape := range stage.KingArthurShapes {
		kingarthurshape.Unstage(stage)
	}

	for knightwhosayni := range stage.KnightWhoSayNis {
		knightwhosayni.Unstage(stage)
	}

	for lancelot := range stage.Lancelots {
		lancelot.Unstage(stage)
	}

	for lancelotagregation := range stage.LancelotAgregations {
		lancelotagregation.Unstage(stage)
	}

	for lancelotagregationuse := range stage.LancelotAgregationUses {
		lancelotagregationuse.Unstage(stage)
	}

	for lancelotcategory := range stage.LancelotCategorys {
		lancelotcategory.Unstage(stage)
	}

	for lancelotcategoryuse := range stage.LancelotCategoryUses {
		lancelotcategoryuse.Unstage(stage)
	}

	for mapobject := range stage.MapObjects {
		mapobject.Unstage(stage)
	}

	for mapobjectuse := range stage.MapObjectUses {
		mapobjectuse.Unstage(stage)
	}

	for repository := range stage.Repositorys {
		repository.Unstage(stage)
	}

	for sirrobin := range stage.SirRobins {
		sirrobin.Unstage(stage)
	}

	for thebridge := range stage.TheBridges {
		thebridge.Unstage(stage)
	}

	for thenuteshape := range stage.TheNuteShapes {
		thenuteshape.Unstage(stage)
	}

	for thenutetransition := range stage.TheNuteTransitions {
		thenutetransition.Unstage(stage)
	}

	for user := range stage.Users {
		user.Unstage(stage)
	}

	for useruse := range stage.UserUses {
		useruse.Unstage(stage)
	}

	for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
		whatisyourpreferedcolor.Unstage(stage)
	}

	for workspace := range stage.Workspaces {
		workspace.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	AWitch | BlackKnightShape | BringYourDead | Document | DocumentUse | GalahadThePure | GeoObject | GeoObjectUse | Group | GroupUse | KingArthur | KingArthurShape | KnightWhoSayNi | Lancelot | LancelotAgregation | LancelotAgregationUse | LancelotCategory | LancelotCategoryUse | MapObject | MapObjectUse | Repository | SirRobin | TheBridge | TheNuteShape | TheNuteTransition | User | UserUse | WhatIsYourPreferedColor | Workspace
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*AWitch | *BlackKnightShape | *BringYourDead | *Document | *DocumentUse | *GalahadThePure | *GeoObject | *GeoObjectUse | *Group | *GroupUse | *KingArthur | *KingArthurShape | *KnightWhoSayNi | *Lancelot | *LancelotAgregation | *LancelotAgregationUse | *LancelotCategory | *LancelotCategoryUse | *MapObject | *MapObjectUse | *Repository | *SirRobin | *TheBridge | *TheNuteShape | *TheNuteTransition | *User | *UserUse | *WhatIsYourPreferedColor | *Workspace
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*AWitch]any |
		map[*BlackKnightShape]any |
		map[*BringYourDead]any |
		map[*Document]any |
		map[*DocumentUse]any |
		map[*GalahadThePure]any |
		map[*GeoObject]any |
		map[*GeoObjectUse]any |
		map[*Group]any |
		map[*GroupUse]any |
		map[*KingArthur]any |
		map[*KingArthurShape]any |
		map[*KnightWhoSayNi]any |
		map[*Lancelot]any |
		map[*LancelotAgregation]any |
		map[*LancelotAgregationUse]any |
		map[*LancelotCategory]any |
		map[*LancelotCategoryUse]any |
		map[*MapObject]any |
		map[*MapObjectUse]any |
		map[*Repository]any |
		map[*SirRobin]any |
		map[*TheBridge]any |
		map[*TheNuteShape]any |
		map[*TheNuteTransition]any |
		map[*User]any |
		map[*UserUse]any |
		map[*WhatIsYourPreferedColor]any |
		map[*Workspace]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*AWitch |
		map[string]*BlackKnightShape |
		map[string]*BringYourDead |
		map[string]*Document |
		map[string]*DocumentUse |
		map[string]*GalahadThePure |
		map[string]*GeoObject |
		map[string]*GeoObjectUse |
		map[string]*Group |
		map[string]*GroupUse |
		map[string]*KingArthur |
		map[string]*KingArthurShape |
		map[string]*KnightWhoSayNi |
		map[string]*Lancelot |
		map[string]*LancelotAgregation |
		map[string]*LancelotAgregationUse |
		map[string]*LancelotCategory |
		map[string]*LancelotCategoryUse |
		map[string]*MapObject |
		map[string]*MapObjectUse |
		map[string]*Repository |
		map[string]*SirRobin |
		map[string]*TheBridge |
		map[string]*TheNuteShape |
		map[string]*TheNuteTransition |
		map[string]*User |
		map[string]*UserUse |
		map[string]*WhatIsYourPreferedColor |
		map[string]*Workspace |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*AWitch]any:
		return any(&stage.AWitchs).(*Type)
	case map[*BlackKnightShape]any:
		return any(&stage.BlackKnightShapes).(*Type)
	case map[*BringYourDead]any:
		return any(&stage.BringYourDeads).(*Type)
	case map[*Document]any:
		return any(&stage.Documents).(*Type)
	case map[*DocumentUse]any:
		return any(&stage.DocumentUses).(*Type)
	case map[*GalahadThePure]any:
		return any(&stage.GalahadThePures).(*Type)
	case map[*GeoObject]any:
		return any(&stage.GeoObjects).(*Type)
	case map[*GeoObjectUse]any:
		return any(&stage.GeoObjectUses).(*Type)
	case map[*Group]any:
		return any(&stage.Groups).(*Type)
	case map[*GroupUse]any:
		return any(&stage.GroupUses).(*Type)
	case map[*KingArthur]any:
		return any(&stage.KingArthurs).(*Type)
	case map[*KingArthurShape]any:
		return any(&stage.KingArthurShapes).(*Type)
	case map[*KnightWhoSayNi]any:
		return any(&stage.KnightWhoSayNis).(*Type)
	case map[*Lancelot]any:
		return any(&stage.Lancelots).(*Type)
	case map[*LancelotAgregation]any:
		return any(&stage.LancelotAgregations).(*Type)
	case map[*LancelotAgregationUse]any:
		return any(&stage.LancelotAgregationUses).(*Type)
	case map[*LancelotCategory]any:
		return any(&stage.LancelotCategorys).(*Type)
	case map[*LancelotCategoryUse]any:
		return any(&stage.LancelotCategoryUses).(*Type)
	case map[*MapObject]any:
		return any(&stage.MapObjects).(*Type)
	case map[*MapObjectUse]any:
		return any(&stage.MapObjectUses).(*Type)
	case map[*Repository]any:
		return any(&stage.Repositorys).(*Type)
	case map[*SirRobin]any:
		return any(&stage.SirRobins).(*Type)
	case map[*TheBridge]any:
		return any(&stage.TheBridges).(*Type)
	case map[*TheNuteShape]any:
		return any(&stage.TheNuteShapes).(*Type)
	case map[*TheNuteTransition]any:
		return any(&stage.TheNuteTransitions).(*Type)
	case map[*User]any:
		return any(&stage.Users).(*Type)
	case map[*UserUse]any:
		return any(&stage.UserUses).(*Type)
	case map[*WhatIsYourPreferedColor]any:
		return any(&stage.WhatIsYourPreferedColors).(*Type)
	case map[*Workspace]any:
		return any(&stage.Workspaces).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*AWitch:
		return any(&stage.AWitchs_mapString).(*Type)
	case map[string]*BlackKnightShape:
		return any(&stage.BlackKnightShapes_mapString).(*Type)
	case map[string]*BringYourDead:
		return any(&stage.BringYourDeads_mapString).(*Type)
	case map[string]*Document:
		return any(&stage.Documents_mapString).(*Type)
	case map[string]*DocumentUse:
		return any(&stage.DocumentUses_mapString).(*Type)
	case map[string]*GalahadThePure:
		return any(&stage.GalahadThePures_mapString).(*Type)
	case map[string]*GeoObject:
		return any(&stage.GeoObjects_mapString).(*Type)
	case map[string]*GeoObjectUse:
		return any(&stage.GeoObjectUses_mapString).(*Type)
	case map[string]*Group:
		return any(&stage.Groups_mapString).(*Type)
	case map[string]*GroupUse:
		return any(&stage.GroupUses_mapString).(*Type)
	case map[string]*KingArthur:
		return any(&stage.KingArthurs_mapString).(*Type)
	case map[string]*KingArthurShape:
		return any(&stage.KingArthurShapes_mapString).(*Type)
	case map[string]*KnightWhoSayNi:
		return any(&stage.KnightWhoSayNis_mapString).(*Type)
	case map[string]*Lancelot:
		return any(&stage.Lancelots_mapString).(*Type)
	case map[string]*LancelotAgregation:
		return any(&stage.LancelotAgregations_mapString).(*Type)
	case map[string]*LancelotAgregationUse:
		return any(&stage.LancelotAgregationUses_mapString).(*Type)
	case map[string]*LancelotCategory:
		return any(&stage.LancelotCategorys_mapString).(*Type)
	case map[string]*LancelotCategoryUse:
		return any(&stage.LancelotCategoryUses_mapString).(*Type)
	case map[string]*MapObject:
		return any(&stage.MapObjects_mapString).(*Type)
	case map[string]*MapObjectUse:
		return any(&stage.MapObjectUses_mapString).(*Type)
	case map[string]*Repository:
		return any(&stage.Repositorys_mapString).(*Type)
	case map[string]*SirRobin:
		return any(&stage.SirRobins_mapString).(*Type)
	case map[string]*TheBridge:
		return any(&stage.TheBridges_mapString).(*Type)
	case map[string]*TheNuteShape:
		return any(&stage.TheNuteShapes_mapString).(*Type)
	case map[string]*TheNuteTransition:
		return any(&stage.TheNuteTransitions_mapString).(*Type)
	case map[string]*User:
		return any(&stage.Users_mapString).(*Type)
	case map[string]*UserUse:
		return any(&stage.UserUses_mapString).(*Type)
	case map[string]*WhatIsYourPreferedColor:
		return any(&stage.WhatIsYourPreferedColors_mapString).(*Type)
	case map[string]*Workspace:
		return any(&stage.Workspaces_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AWitch:
		return any(&stage.AWitchs).(*map[*Type]any)
	case BlackKnightShape:
		return any(&stage.BlackKnightShapes).(*map[*Type]any)
	case BringYourDead:
		return any(&stage.BringYourDeads).(*map[*Type]any)
	case Document:
		return any(&stage.Documents).(*map[*Type]any)
	case DocumentUse:
		return any(&stage.DocumentUses).(*map[*Type]any)
	case GalahadThePure:
		return any(&stage.GalahadThePures).(*map[*Type]any)
	case GeoObject:
		return any(&stage.GeoObjects).(*map[*Type]any)
	case GeoObjectUse:
		return any(&stage.GeoObjectUses).(*map[*Type]any)
	case Group:
		return any(&stage.Groups).(*map[*Type]any)
	case GroupUse:
		return any(&stage.GroupUses).(*map[*Type]any)
	case KingArthur:
		return any(&stage.KingArthurs).(*map[*Type]any)
	case KingArthurShape:
		return any(&stage.KingArthurShapes).(*map[*Type]any)
	case KnightWhoSayNi:
		return any(&stage.KnightWhoSayNis).(*map[*Type]any)
	case Lancelot:
		return any(&stage.Lancelots).(*map[*Type]any)
	case LancelotAgregation:
		return any(&stage.LancelotAgregations).(*map[*Type]any)
	case LancelotAgregationUse:
		return any(&stage.LancelotAgregationUses).(*map[*Type]any)
	case LancelotCategory:
		return any(&stage.LancelotCategorys).(*map[*Type]any)
	case LancelotCategoryUse:
		return any(&stage.LancelotCategoryUses).(*map[*Type]any)
	case MapObject:
		return any(&stage.MapObjects).(*map[*Type]any)
	case MapObjectUse:
		return any(&stage.MapObjectUses).(*map[*Type]any)
	case Repository:
		return any(&stage.Repositorys).(*map[*Type]any)
	case SirRobin:
		return any(&stage.SirRobins).(*map[*Type]any)
	case TheBridge:
		return any(&stage.TheBridges).(*map[*Type]any)
	case TheNuteShape:
		return any(&stage.TheNuteShapes).(*map[*Type]any)
	case TheNuteTransition:
		return any(&stage.TheNuteTransitions).(*map[*Type]any)
	case User:
		return any(&stage.Users).(*map[*Type]any)
	case UserUse:
		return any(&stage.UserUses).(*map[*Type]any)
	case WhatIsYourPreferedColor:
		return any(&stage.WhatIsYourPreferedColors).(*map[*Type]any)
	case Workspace:
		return any(&stage.Workspaces).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AWitch:
		return any(&stage.AWitchs).(*map[Type]any)
	case *BlackKnightShape:
		return any(&stage.BlackKnightShapes).(*map[Type]any)
	case *BringYourDead:
		return any(&stage.BringYourDeads).(*map[Type]any)
	case *Document:
		return any(&stage.Documents).(*map[Type]any)
	case *DocumentUse:
		return any(&stage.DocumentUses).(*map[Type]any)
	case *GalahadThePure:
		return any(&stage.GalahadThePures).(*map[Type]any)
	case *GeoObject:
		return any(&stage.GeoObjects).(*map[Type]any)
	case *GeoObjectUse:
		return any(&stage.GeoObjectUses).(*map[Type]any)
	case *Group:
		return any(&stage.Groups).(*map[Type]any)
	case *GroupUse:
		return any(&stage.GroupUses).(*map[Type]any)
	case *KingArthur:
		return any(&stage.KingArthurs).(*map[Type]any)
	case *KingArthurShape:
		return any(&stage.KingArthurShapes).(*map[Type]any)
	case *KnightWhoSayNi:
		return any(&stage.KnightWhoSayNis).(*map[Type]any)
	case *Lancelot:
		return any(&stage.Lancelots).(*map[Type]any)
	case *LancelotAgregation:
		return any(&stage.LancelotAgregations).(*map[Type]any)
	case *LancelotAgregationUse:
		return any(&stage.LancelotAgregationUses).(*map[Type]any)
	case *LancelotCategory:
		return any(&stage.LancelotCategorys).(*map[Type]any)
	case *LancelotCategoryUse:
		return any(&stage.LancelotCategoryUses).(*map[Type]any)
	case *MapObject:
		return any(&stage.MapObjects).(*map[Type]any)
	case *MapObjectUse:
		return any(&stage.MapObjectUses).(*map[Type]any)
	case *Repository:
		return any(&stage.Repositorys).(*map[Type]any)
	case *SirRobin:
		return any(&stage.SirRobins).(*map[Type]any)
	case *TheBridge:
		return any(&stage.TheBridges).(*map[Type]any)
	case *TheNuteShape:
		return any(&stage.TheNuteShapes).(*map[Type]any)
	case *TheNuteTransition:
		return any(&stage.TheNuteTransitions).(*map[Type]any)
	case *User:
		return any(&stage.Users).(*map[Type]any)
	case *UserUse:
		return any(&stage.UserUses).(*map[Type]any)
	case *WhatIsYourPreferedColor:
		return any(&stage.WhatIsYourPreferedColors).(*map[Type]any)
	case *Workspace:
		return any(&stage.Workspaces).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AWitch:
		return any(&stage.AWitchs_mapString).(*map[string]*Type)
	case BlackKnightShape:
		return any(&stage.BlackKnightShapes_mapString).(*map[string]*Type)
	case BringYourDead:
		return any(&stage.BringYourDeads_mapString).(*map[string]*Type)
	case Document:
		return any(&stage.Documents_mapString).(*map[string]*Type)
	case DocumentUse:
		return any(&stage.DocumentUses_mapString).(*map[string]*Type)
	case GalahadThePure:
		return any(&stage.GalahadThePures_mapString).(*map[string]*Type)
	case GeoObject:
		return any(&stage.GeoObjects_mapString).(*map[string]*Type)
	case GeoObjectUse:
		return any(&stage.GeoObjectUses_mapString).(*map[string]*Type)
	case Group:
		return any(&stage.Groups_mapString).(*map[string]*Type)
	case GroupUse:
		return any(&stage.GroupUses_mapString).(*map[string]*Type)
	case KingArthur:
		return any(&stage.KingArthurs_mapString).(*map[string]*Type)
	case KingArthurShape:
		return any(&stage.KingArthurShapes_mapString).(*map[string]*Type)
	case KnightWhoSayNi:
		return any(&stage.KnightWhoSayNis_mapString).(*map[string]*Type)
	case Lancelot:
		return any(&stage.Lancelots_mapString).(*map[string]*Type)
	case LancelotAgregation:
		return any(&stage.LancelotAgregations_mapString).(*map[string]*Type)
	case LancelotAgregationUse:
		return any(&stage.LancelotAgregationUses_mapString).(*map[string]*Type)
	case LancelotCategory:
		return any(&stage.LancelotCategorys_mapString).(*map[string]*Type)
	case LancelotCategoryUse:
		return any(&stage.LancelotCategoryUses_mapString).(*map[string]*Type)
	case MapObject:
		return any(&stage.MapObjects_mapString).(*map[string]*Type)
	case MapObjectUse:
		return any(&stage.MapObjectUses_mapString).(*map[string]*Type)
	case Repository:
		return any(&stage.Repositorys_mapString).(*map[string]*Type)
	case SirRobin:
		return any(&stage.SirRobins_mapString).(*map[string]*Type)
	case TheBridge:
		return any(&stage.TheBridges_mapString).(*map[string]*Type)
	case TheNuteShape:
		return any(&stage.TheNuteShapes_mapString).(*map[string]*Type)
	case TheNuteTransition:
		return any(&stage.TheNuteTransitions_mapString).(*map[string]*Type)
	case User:
		return any(&stage.Users_mapString).(*map[string]*Type)
	case UserUse:
		return any(&stage.UserUses_mapString).(*map[string]*Type)
	case WhatIsYourPreferedColor:
		return any(&stage.WhatIsYourPreferedColors_mapString).(*map[string]*Type)
	case Workspace:
		return any(&stage.Workspaces_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case AWitch:
		return any(&AWitch{
			// Initialisation of associations
			// field is initialized with an instance of GalahadThePure with the name of the field
			EvolutionDirection: &GalahadThePure{Name: "EvolutionDirection"},
		}).(*Type)
	case BlackKnightShape:
		return any(&BlackKnightShape{
			// Initialisation of associations
			// field is initialized with an instance of BringYourDead with the name of the field
			BringYourDead: &BringYourDead{Name: "BringYourDead"},
		}).(*Type)
	case BringYourDead:
		return any(&BringYourDead{
			// Initialisation of associations
			// field is initialized with an instance of Lancelot with the name of the field
			Lancelots: []*Lancelot{{Name: "Lancelots"}},
		}).(*Type)
	case Document:
		return any(&Document{
			// Initialisation of associations
			// field is initialized with an instance of GeoObjectUse with the name of the field
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
		}).(*Type)
	case DocumentUse:
		return any(&DocumentUse{
			// Initialisation of associations
			// field is initialized with an instance of Document with the name of the field
			Document: &Document{Name: "Document"},
		}).(*Type)
	case GalahadThePure:
		return any(&GalahadThePure{
			// Initialisation of associations
		}).(*Type)
	case GeoObject:
		return any(&GeoObject{
			// Initialisation of associations
		}).(*Type)
	case GeoObjectUse:
		return any(&GeoObjectUse{
			// Initialisation of associations
			// field is initialized with an instance of GeoObject with the name of the field
			GeoObject: &GeoObject{Name: "GeoObject"},
		}).(*Type)
	case Group:
		return any(&Group{
			// Initialisation of associations
			// field is initialized with an instance of UserUse with the name of the field
			UserUse: []*UserUse{{Name: "UserUse"}},
		}).(*Type)
	case GroupUse:
		return any(&GroupUse{
			// Initialisation of associations
			// field is initialized with an instance of Group with the name of the field
			Group: &Group{Name: "Group"},
		}).(*Type)
	case KingArthur:
		return any(&KingArthur{
			// Initialisation of associations
		}).(*Type)
	case KingArthurShape:
		return any(&KingArthurShape{
			// Initialisation of associations
			// field is initialized with an instance of KingArthur with the name of the field
			ActorState: &KingArthur{Name: "ActorState"},
		}).(*Type)
	case KnightWhoSayNi:
		return any(&KnightWhoSayNi{
			// Initialisation of associations
			// field is initialized with an instance of Lancelot with the name of the field
			Parameter: &Lancelot{Name: "Parameter"},
		}).(*Type)
	case Lancelot:
		return any(&Lancelot{
			// Initialisation of associations
			// field is initialized with an instance of GroupUse with the name of the field
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
			// field is initialized with an instance of DocumentUse with the name of the field
			DocumentUse: []*DocumentUse{{Name: "DocumentUse"}},
			// field is initialized with an instance of GeoObjectUse with the name of the field
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
		}).(*Type)
	case LancelotAgregation:
		return any(&LancelotAgregation{
			// Initialisation of associations
			// field is initialized with an instance of KnightWhoSayNi with the name of the field
			ParameterUse: []*KnightWhoSayNi{{Name: "ParameterUse"}},
		}).(*Type)
	case LancelotAgregationUse:
		return any(&LancelotAgregationUse{
			// Initialisation of associations
			// field is initialized with an instance of LancelotAgregation with the name of the field
			ParameterAgregation: &LancelotAgregation{Name: "ParameterAgregation"},
		}).(*Type)
	case LancelotCategory:
		return any(&LancelotCategory{
			// Initialisation of associations
			// field is initialized with an instance of KnightWhoSayNi with the name of the field
			ParameterUse: []*KnightWhoSayNi{{Name: "ParameterUse"}},
		}).(*Type)
	case LancelotCategoryUse:
		return any(&LancelotCategoryUse{
			// Initialisation of associations
			// field is initialized with an instance of LancelotCategory with the name of the field
			ParameterCategory: &LancelotCategory{Name: "ParameterCategory"},
		}).(*Type)
	case MapObject:
		return any(&MapObject{
			// Initialisation of associations
		}).(*Type)
	case MapObjectUse:
		return any(&MapObjectUse{
			// Initialisation of associations
			// field is initialized with an instance of MapObject with the name of the field
			Map: &MapObject{Name: "Map"},
		}).(*Type)
	case Repository:
		return any(&Repository{
			// Initialisation of associations
			// field is initialized with an instance of KnightWhoSayNi with the name of the field
			ParameterUse: []*KnightWhoSayNi{{Name: "ParameterUse"}},
			// field is initialized with an instance of GroupUse with the name of the field
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
		}).(*Type)
	case SirRobin:
		return any(&SirRobin{
			// Initialisation of associations
			// field is initialized with an instance of AWitch with the name of the field
			Witches: []*AWitch{{Name: "Witches"}},
			// field is initialized with an instance of KingArthurShape with the name of the field
			Arthurs: []*KingArthurShape{{Name: "Arthurs"}},
			// field is initialized with an instance of KnightWhoSayNi with the name of the field
			KnightWhoSayNis: []*KnightWhoSayNi{{Name: "KnightWhoSayNis"}},
			// field is initialized with an instance of BlackKnightShape with the name of the field
			BlackKnightShapes: []*BlackKnightShape{{Name: "BlackKnightShapes"}},
			// field is initialized with an instance of TheNuteShape with the name of the field
			TheNuteShapes: []*TheNuteShape{{Name: "TheNuteShapes"}},
		}).(*Type)
	case TheBridge:
		return any(&TheBridge{
			// Initialisation of associations
			// field is initialized with an instance of WhatIsYourPreferedColor with the name of the field
			WhatIsYourPreferedColor: []*WhatIsYourPreferedColor{{Name: "WhatIsYourPreferedColor"}},
			// field is initialized with an instance of GroupUse with the name of the field
			GroupUse: []*GroupUse{{Name: "GroupUse"}},
			// field is initialized with an instance of GeoObjectUse with the name of the field
			GeoObjectUse: []*GeoObjectUse{{Name: "GeoObjectUse"}},
			// field is initialized with an instance of MapObjectUse with the name of the field
			MapUse: []*MapObjectUse{{Name: "MapUse"}},
		}).(*Type)
	case TheNuteShape:
		return any(&TheNuteShape{
			// Initialisation of associations
			// field is initialized with an instance of TheNuteTransition with the name of the field
			ActorStateTransition: &TheNuteTransition{Name: "ActorStateTransition"},
			// field is initialized with an instance of KingArthurShape with the name of the field
			Start: &KingArthurShape{Name: "Start"},
			// field is initialized with an instance of KingArthurShape with the name of the field
			End: &KingArthurShape{Name: "End"},
		}).(*Type)
	case TheNuteTransition:
		return any(&TheNuteTransition{
			// Initialisation of associations
			// field is initialized with an instance of KingArthur with the name of the field
			StartState: &KingArthur{Name: "StartState"},
			// field is initialized with an instance of KingArthur with the name of the field
			EndState: &KingArthur{Name: "EndState"},
		}).(*Type)
	case User:
		return any(&User{
			// Initialisation of associations
		}).(*Type)
	case UserUse:
		return any(&UserUse{
			// Initialisation of associations
			// field is initialized with an instance of User with the name of the field
			User: &User{Name: "User"},
		}).(*Type)
	case WhatIsYourPreferedColor:
		return any(&WhatIsYourPreferedColor{
			// Initialisation of associations
			// field is initialized with an instance of SirRobin with the name of the field
			Diagrams: []*SirRobin{{Name: "Diagrams"}},
			// field is initialized with an instance of KingArthur with the name of the field
			KingArthurs: []*KingArthur{{Name: "KingArthurs"}},
			// field is initialized with an instance of TheNuteTransition with the name of the field
			Nutes: []*TheNuteTransition{{Name: "Nutes"}},
			// field is initialized with an instance of GalahadThePure with the name of the field
			Galahard: []*GalahadThePure{{Name: "Galahard"}},
			// field is initialized with an instance of Lancelot with the name of the field
			Lancelots: []*Lancelot{{Name: "Lancelots"}},
			// field is initialized with an instance of BringYourDead with the name of the field
			BringYourDeadarameters: []*BringYourDead{{Name: "BringYourDeadarameters"}},
		}).(*Type)
	case Workspace:
		return any(&Workspace{
			// Initialisation of associations
			// field is initialized with an instance of SirRobin with the name of the field
			SelectedDiagram: &SirRobin{Name: "SelectedDiagram"},
			// field is initialized with an instance of AWitch with the name of the field
			A: &AWitch{Name: "A"},
			// field is initialized with an instance of KnightWhoSayNi with the name of the field
			B: &KnightWhoSayNi{Name: "B"},
			// field is initialized with an instance of BlackKnightShape with the name of the field
			C: &BlackKnightShape{Name: "C"},
			// field is initialized with an instance of KingArthurShape with the name of the field
			D: &KingArthurShape{Name: "D"},
			// field is initialized with an instance of TheNuteShape with the name of the field
			E: &TheNuteShape{Name: "E"},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AWitch
	case AWitch:
		switch fieldname {
		// insertion point for per direct association field
		case "EvolutionDirection":
			res := make(map[*GalahadThePure][]*AWitch)
			for awitch := range stage.AWitchs {
				if awitch.EvolutionDirection != nil {
					galahadthepure_ := awitch.EvolutionDirection
					var awitchs []*AWitch
					_, ok := res[galahadthepure_]
					if ok {
						awitchs = res[galahadthepure_]
					} else {
						awitchs = make([]*AWitch, 0)
					}
					awitchs = append(awitchs, awitch)
					res[galahadthepure_] = awitchs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BlackKnightShape
	case BlackKnightShape:
		switch fieldname {
		// insertion point for per direct association field
		case "BringYourDead":
			res := make(map[*BringYourDead][]*BlackKnightShape)
			for blackknightshape := range stage.BlackKnightShapes {
				if blackknightshape.BringYourDead != nil {
					bringyourdead_ := blackknightshape.BringYourDead
					var blackknightshapes []*BlackKnightShape
					_, ok := res[bringyourdead_]
					if ok {
						blackknightshapes = res[bringyourdead_]
					} else {
						blackknightshapes = make([]*BlackKnightShape, 0)
					}
					blackknightshapes = append(blackknightshapes, blackknightshape)
					res[bringyourdead_] = blackknightshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BringYourDead
	case BringYourDead:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Document
	case Document:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DocumentUse
	case DocumentUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Document":
			res := make(map[*Document][]*DocumentUse)
			for documentuse := range stage.DocumentUses {
				if documentuse.Document != nil {
					document_ := documentuse.Document
					var documentuses []*DocumentUse
					_, ok := res[document_]
					if ok {
						documentuses = res[document_]
					} else {
						documentuses = make([]*DocumentUse, 0)
					}
					documentuses = append(documentuses, documentuse)
					res[document_] = documentuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GalahadThePure
	case GalahadThePure:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoObject
	case GeoObject:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoObjectUse
	case GeoObjectUse:
		switch fieldname {
		// insertion point for per direct association field
		case "GeoObject":
			res := make(map[*GeoObject][]*GeoObjectUse)
			for geoobjectuse := range stage.GeoObjectUses {
				if geoobjectuse.GeoObject != nil {
					geoobject_ := geoobjectuse.GeoObject
					var geoobjectuses []*GeoObjectUse
					_, ok := res[geoobject_]
					if ok {
						geoobjectuses = res[geoobject_]
					} else {
						geoobjectuses = make([]*GeoObjectUse, 0)
					}
					geoobjectuses = append(geoobjectuses, geoobjectuse)
					res[geoobject_] = geoobjectuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupUse
	case GroupUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Group":
			res := make(map[*Group][]*GroupUse)
			for groupuse := range stage.GroupUses {
				if groupuse.Group != nil {
					group_ := groupuse.Group
					var groupuses []*GroupUse
					_, ok := res[group_]
					if ok {
						groupuses = res[group_]
					} else {
						groupuses = make([]*GroupUse, 0)
					}
					groupuses = append(groupuses, groupuse)
					res[group_] = groupuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of KingArthur
	case KingArthur:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of KingArthurShape
	case KingArthurShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ActorState":
			res := make(map[*KingArthur][]*KingArthurShape)
			for kingarthurshape := range stage.KingArthurShapes {
				if kingarthurshape.ActorState != nil {
					kingarthur_ := kingarthurshape.ActorState
					var kingarthurshapes []*KingArthurShape
					_, ok := res[kingarthur_]
					if ok {
						kingarthurshapes = res[kingarthur_]
					} else {
						kingarthurshapes = make([]*KingArthurShape, 0)
					}
					kingarthurshapes = append(kingarthurshapes, kingarthurshape)
					res[kingarthur_] = kingarthurshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of KnightWhoSayNi
	case KnightWhoSayNi:
		switch fieldname {
		// insertion point for per direct association field
		case "Parameter":
			res := make(map[*Lancelot][]*KnightWhoSayNi)
			for knightwhosayni := range stage.KnightWhoSayNis {
				if knightwhosayni.Parameter != nil {
					lancelot_ := knightwhosayni.Parameter
					var knightwhosaynis []*KnightWhoSayNi
					_, ok := res[lancelot_]
					if ok {
						knightwhosaynis = res[lancelot_]
					} else {
						knightwhosaynis = make([]*KnightWhoSayNi, 0)
					}
					knightwhosaynis = append(knightwhosaynis, knightwhosayni)
					res[lancelot_] = knightwhosaynis
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Lancelot
	case Lancelot:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LancelotAgregation
	case LancelotAgregation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LancelotAgregationUse
	case LancelotAgregationUse:
		switch fieldname {
		// insertion point for per direct association field
		case "ParameterAgregation":
			res := make(map[*LancelotAgregation][]*LancelotAgregationUse)
			for lancelotagregationuse := range stage.LancelotAgregationUses {
				if lancelotagregationuse.ParameterAgregation != nil {
					lancelotagregation_ := lancelotagregationuse.ParameterAgregation
					var lancelotagregationuses []*LancelotAgregationUse
					_, ok := res[lancelotagregation_]
					if ok {
						lancelotagregationuses = res[lancelotagregation_]
					} else {
						lancelotagregationuses = make([]*LancelotAgregationUse, 0)
					}
					lancelotagregationuses = append(lancelotagregationuses, lancelotagregationuse)
					res[lancelotagregation_] = lancelotagregationuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LancelotCategory
	case LancelotCategory:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LancelotCategoryUse
	case LancelotCategoryUse:
		switch fieldname {
		// insertion point for per direct association field
		case "ParameterCategory":
			res := make(map[*LancelotCategory][]*LancelotCategoryUse)
			for lancelotcategoryuse := range stage.LancelotCategoryUses {
				if lancelotcategoryuse.ParameterCategory != nil {
					lancelotcategory_ := lancelotcategoryuse.ParameterCategory
					var lancelotcategoryuses []*LancelotCategoryUse
					_, ok := res[lancelotcategory_]
					if ok {
						lancelotcategoryuses = res[lancelotcategory_]
					} else {
						lancelotcategoryuses = make([]*LancelotCategoryUse, 0)
					}
					lancelotcategoryuses = append(lancelotcategoryuses, lancelotcategoryuse)
					res[lancelotcategory_] = lancelotcategoryuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MapObject
	case MapObject:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MapObjectUse
	case MapObjectUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Map":
			res := make(map[*MapObject][]*MapObjectUse)
			for mapobjectuse := range stage.MapObjectUses {
				if mapobjectuse.Map != nil {
					mapobject_ := mapobjectuse.Map
					var mapobjectuses []*MapObjectUse
					_, ok := res[mapobject_]
					if ok {
						mapobjectuses = res[mapobject_]
					} else {
						mapobjectuses = make([]*MapObjectUse, 0)
					}
					mapobjectuses = append(mapobjectuses, mapobjectuse)
					res[mapobject_] = mapobjectuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Repository
	case Repository:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SirRobin
	case SirRobin:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TheBridge
	case TheBridge:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TheNuteShape
	case TheNuteShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ActorStateTransition":
			res := make(map[*TheNuteTransition][]*TheNuteShape)
			for thenuteshape := range stage.TheNuteShapes {
				if thenuteshape.ActorStateTransition != nil {
					thenutetransition_ := thenuteshape.ActorStateTransition
					var thenuteshapes []*TheNuteShape
					_, ok := res[thenutetransition_]
					if ok {
						thenuteshapes = res[thenutetransition_]
					} else {
						thenuteshapes = make([]*TheNuteShape, 0)
					}
					thenuteshapes = append(thenuteshapes, thenuteshape)
					res[thenutetransition_] = thenuteshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Start":
			res := make(map[*KingArthurShape][]*TheNuteShape)
			for thenuteshape := range stage.TheNuteShapes {
				if thenuteshape.Start != nil {
					kingarthurshape_ := thenuteshape.Start
					var thenuteshapes []*TheNuteShape
					_, ok := res[kingarthurshape_]
					if ok {
						thenuteshapes = res[kingarthurshape_]
					} else {
						thenuteshapes = make([]*TheNuteShape, 0)
					}
					thenuteshapes = append(thenuteshapes, thenuteshape)
					res[kingarthurshape_] = thenuteshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*KingArthurShape][]*TheNuteShape)
			for thenuteshape := range stage.TheNuteShapes {
				if thenuteshape.End != nil {
					kingarthurshape_ := thenuteshape.End
					var thenuteshapes []*TheNuteShape
					_, ok := res[kingarthurshape_]
					if ok {
						thenuteshapes = res[kingarthurshape_]
					} else {
						thenuteshapes = make([]*TheNuteShape, 0)
					}
					thenuteshapes = append(thenuteshapes, thenuteshape)
					res[kingarthurshape_] = thenuteshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TheNuteTransition
	case TheNuteTransition:
		switch fieldname {
		// insertion point for per direct association field
		case "StartState":
			res := make(map[*KingArthur][]*TheNuteTransition)
			for thenutetransition := range stage.TheNuteTransitions {
				if thenutetransition.StartState != nil {
					kingarthur_ := thenutetransition.StartState
					var thenutetransitions []*TheNuteTransition
					_, ok := res[kingarthur_]
					if ok {
						thenutetransitions = res[kingarthur_]
					} else {
						thenutetransitions = make([]*TheNuteTransition, 0)
					}
					thenutetransitions = append(thenutetransitions, thenutetransition)
					res[kingarthur_] = thenutetransitions
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndState":
			res := make(map[*KingArthur][]*TheNuteTransition)
			for thenutetransition := range stage.TheNuteTransitions {
				if thenutetransition.EndState != nil {
					kingarthur_ := thenutetransition.EndState
					var thenutetransitions []*TheNuteTransition
					_, ok := res[kingarthur_]
					if ok {
						thenutetransitions = res[kingarthur_]
					} else {
						thenutetransitions = make([]*TheNuteTransition, 0)
					}
					thenutetransitions = append(thenutetransitions, thenutetransition)
					res[kingarthur_] = thenutetransitions
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of User
	case User:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UserUse
	case UserUse:
		switch fieldname {
		// insertion point for per direct association field
		case "User":
			res := make(map[*User][]*UserUse)
			for useruse := range stage.UserUses {
				if useruse.User != nil {
					user_ := useruse.User
					var useruses []*UserUse
					_, ok := res[user_]
					if ok {
						useruses = res[user_]
					} else {
						useruses = make([]*UserUse, 0)
					}
					useruses = append(useruses, useruse)
					res[user_] = useruses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of WhatIsYourPreferedColor
	case WhatIsYourPreferedColor:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Workspace
	case Workspace:
		switch fieldname {
		// insertion point for per direct association field
		case "SelectedDiagram":
			res := make(map[*SirRobin][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.SelectedDiagram != nil {
					sirrobin_ := workspace.SelectedDiagram
					var workspaces []*Workspace
					_, ok := res[sirrobin_]
					if ok {
						workspaces = res[sirrobin_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[sirrobin_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "A":
			res := make(map[*AWitch][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.A != nil {
					awitch_ := workspace.A
					var workspaces []*Workspace
					_, ok := res[awitch_]
					if ok {
						workspaces = res[awitch_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[awitch_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "B":
			res := make(map[*KnightWhoSayNi][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.B != nil {
					knightwhosayni_ := workspace.B
					var workspaces []*Workspace
					_, ok := res[knightwhosayni_]
					if ok {
						workspaces = res[knightwhosayni_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[knightwhosayni_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "C":
			res := make(map[*BlackKnightShape][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.C != nil {
					blackknightshape_ := workspace.C
					var workspaces []*Workspace
					_, ok := res[blackknightshape_]
					if ok {
						workspaces = res[blackknightshape_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[blackknightshape_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "D":
			res := make(map[*KingArthurShape][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.D != nil {
					kingarthurshape_ := workspace.D
					var workspaces []*Workspace
					_, ok := res[kingarthurshape_]
					if ok {
						workspaces = res[kingarthurshape_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[kingarthurshape_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		case "E":
			res := make(map[*TheNuteShape][]*Workspace)
			for workspace := range stage.Workspaces {
				if workspace.E != nil {
					thenuteshape_ := workspace.E
					var workspaces []*Workspace
					_, ok := res[thenuteshape_]
					if ok {
						workspaces = res[thenuteshape_]
					} else {
						workspaces = make([]*Workspace, 0)
					}
					workspaces = append(workspaces, workspace)
					res[thenuteshape_] = workspaces
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AWitch
	case AWitch:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BlackKnightShape
	case BlackKnightShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BringYourDead
	case BringYourDead:
		switch fieldname {
		// insertion point for per direct association field
		case "Lancelots":
			res := make(map[*Lancelot]*BringYourDead)
			for bringyourdead := range stage.BringYourDeads {
				for _, lancelot_ := range bringyourdead.Lancelots {
					res[lancelot_] = bringyourdead
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Document
	case Document:
		switch fieldname {
		// insertion point for per direct association field
		case "GeoObjectUse":
			res := make(map[*GeoObjectUse]*Document)
			for document := range stage.Documents {
				for _, geoobjectuse_ := range document.GeoObjectUse {
					res[geoobjectuse_] = document
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DocumentUse
	case DocumentUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GalahadThePure
	case GalahadThePure:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoObject
	case GeoObject:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoObjectUse
	case GeoObjectUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "UserUse":
			res := make(map[*UserUse]*Group)
			for group := range stage.Groups {
				for _, useruse_ := range group.UserUse {
					res[useruse_] = group
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of GroupUse
	case GroupUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of KingArthur
	case KingArthur:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of KingArthurShape
	case KingArthurShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of KnightWhoSayNi
	case KnightWhoSayNi:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lancelot
	case Lancelot:
		switch fieldname {
		// insertion point for per direct association field
		case "GroupUse":
			res := make(map[*GroupUse]*Lancelot)
			for lancelot := range stage.Lancelots {
				for _, groupuse_ := range lancelot.GroupUse {
					res[groupuse_] = lancelot
				}
			}
			return any(res).(map[*End]*Start)
		case "DocumentUse":
			res := make(map[*DocumentUse]*Lancelot)
			for lancelot := range stage.Lancelots {
				for _, documentuse_ := range lancelot.DocumentUse {
					res[documentuse_] = lancelot
				}
			}
			return any(res).(map[*End]*Start)
		case "GeoObjectUse":
			res := make(map[*GeoObjectUse]*Lancelot)
			for lancelot := range stage.Lancelots {
				for _, geoobjectuse_ := range lancelot.GeoObjectUse {
					res[geoobjectuse_] = lancelot
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of LancelotAgregation
	case LancelotAgregation:
		switch fieldname {
		// insertion point for per direct association field
		case "ParameterUse":
			res := make(map[*KnightWhoSayNi]*LancelotAgregation)
			for lancelotagregation := range stage.LancelotAgregations {
				for _, knightwhosayni_ := range lancelotagregation.ParameterUse {
					res[knightwhosayni_] = lancelotagregation
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of LancelotAgregationUse
	case LancelotAgregationUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LancelotCategory
	case LancelotCategory:
		switch fieldname {
		// insertion point for per direct association field
		case "ParameterUse":
			res := make(map[*KnightWhoSayNi]*LancelotCategory)
			for lancelotcategory := range stage.LancelotCategorys {
				for _, knightwhosayni_ := range lancelotcategory.ParameterUse {
					res[knightwhosayni_] = lancelotcategory
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of LancelotCategoryUse
	case LancelotCategoryUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MapObject
	case MapObject:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MapObjectUse
	case MapObjectUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Repository
	case Repository:
		switch fieldname {
		// insertion point for per direct association field
		case "ParameterUse":
			res := make(map[*KnightWhoSayNi]*Repository)
			for repository := range stage.Repositorys {
				for _, knightwhosayni_ := range repository.ParameterUse {
					res[knightwhosayni_] = repository
				}
			}
			return any(res).(map[*End]*Start)
		case "GroupUse":
			res := make(map[*GroupUse]*Repository)
			for repository := range stage.Repositorys {
				for _, groupuse_ := range repository.GroupUse {
					res[groupuse_] = repository
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SirRobin
	case SirRobin:
		switch fieldname {
		// insertion point for per direct association field
		case "Witches":
			res := make(map[*AWitch]*SirRobin)
			for sirrobin := range stage.SirRobins {
				for _, awitch_ := range sirrobin.Witches {
					res[awitch_] = sirrobin
				}
			}
			return any(res).(map[*End]*Start)
		case "Arthurs":
			res := make(map[*KingArthurShape]*SirRobin)
			for sirrobin := range stage.SirRobins {
				for _, kingarthurshape_ := range sirrobin.Arthurs {
					res[kingarthurshape_] = sirrobin
				}
			}
			return any(res).(map[*End]*Start)
		case "KnightWhoSayNis":
			res := make(map[*KnightWhoSayNi]*SirRobin)
			for sirrobin := range stage.SirRobins {
				for _, knightwhosayni_ := range sirrobin.KnightWhoSayNis {
					res[knightwhosayni_] = sirrobin
				}
			}
			return any(res).(map[*End]*Start)
		case "BlackKnightShapes":
			res := make(map[*BlackKnightShape]*SirRobin)
			for sirrobin := range stage.SirRobins {
				for _, blackknightshape_ := range sirrobin.BlackKnightShapes {
					res[blackknightshape_] = sirrobin
				}
			}
			return any(res).(map[*End]*Start)
		case "TheNuteShapes":
			res := make(map[*TheNuteShape]*SirRobin)
			for sirrobin := range stage.SirRobins {
				for _, thenuteshape_ := range sirrobin.TheNuteShapes {
					res[thenuteshape_] = sirrobin
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of TheBridge
	case TheBridge:
		switch fieldname {
		// insertion point for per direct association field
		case "WhatIsYourPreferedColor":
			res := make(map[*WhatIsYourPreferedColor]*TheBridge)
			for thebridge := range stage.TheBridges {
				for _, whatisyourpreferedcolor_ := range thebridge.WhatIsYourPreferedColor {
					res[whatisyourpreferedcolor_] = thebridge
				}
			}
			return any(res).(map[*End]*Start)
		case "GroupUse":
			res := make(map[*GroupUse]*TheBridge)
			for thebridge := range stage.TheBridges {
				for _, groupuse_ := range thebridge.GroupUse {
					res[groupuse_] = thebridge
				}
			}
			return any(res).(map[*End]*Start)
		case "GeoObjectUse":
			res := make(map[*GeoObjectUse]*TheBridge)
			for thebridge := range stage.TheBridges {
				for _, geoobjectuse_ := range thebridge.GeoObjectUse {
					res[geoobjectuse_] = thebridge
				}
			}
			return any(res).(map[*End]*Start)
		case "MapUse":
			res := make(map[*MapObjectUse]*TheBridge)
			for thebridge := range stage.TheBridges {
				for _, mapobjectuse_ := range thebridge.MapUse {
					res[mapobjectuse_] = thebridge
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of TheNuteShape
	case TheNuteShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TheNuteTransition
	case TheNuteTransition:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of User
	case User:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UserUse
	case UserUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of WhatIsYourPreferedColor
	case WhatIsYourPreferedColor:
		switch fieldname {
		// insertion point for per direct association field
		case "Diagrams":
			res := make(map[*SirRobin]*WhatIsYourPreferedColor)
			for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
				for _, sirrobin_ := range whatisyourpreferedcolor.Diagrams {
					res[sirrobin_] = whatisyourpreferedcolor
				}
			}
			return any(res).(map[*End]*Start)
		case "KingArthurs":
			res := make(map[*KingArthur]*WhatIsYourPreferedColor)
			for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
				for _, kingarthur_ := range whatisyourpreferedcolor.KingArthurs {
					res[kingarthur_] = whatisyourpreferedcolor
				}
			}
			return any(res).(map[*End]*Start)
		case "Nutes":
			res := make(map[*TheNuteTransition]*WhatIsYourPreferedColor)
			for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
				for _, thenutetransition_ := range whatisyourpreferedcolor.Nutes {
					res[thenutetransition_] = whatisyourpreferedcolor
				}
			}
			return any(res).(map[*End]*Start)
		case "Galahard":
			res := make(map[*GalahadThePure]*WhatIsYourPreferedColor)
			for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
				for _, galahadthepure_ := range whatisyourpreferedcolor.Galahard {
					res[galahadthepure_] = whatisyourpreferedcolor
				}
			}
			return any(res).(map[*End]*Start)
		case "Lancelots":
			res := make(map[*Lancelot]*WhatIsYourPreferedColor)
			for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
				for _, lancelot_ := range whatisyourpreferedcolor.Lancelots {
					res[lancelot_] = whatisyourpreferedcolor
				}
			}
			return any(res).(map[*End]*Start)
		case "BringYourDeadarameters":
			res := make(map[*BringYourDead]*WhatIsYourPreferedColor)
			for whatisyourpreferedcolor := range stage.WhatIsYourPreferedColors {
				for _, bringyourdead_ := range whatisyourpreferedcolor.BringYourDeadarameters {
					res[bringyourdead_] = whatisyourpreferedcolor
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Workspace
	case Workspace:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case AWitch:
		res = "AWitch"
	case BlackKnightShape:
		res = "BlackKnightShape"
	case BringYourDead:
		res = "BringYourDead"
	case Document:
		res = "Document"
	case DocumentUse:
		res = "DocumentUse"
	case GalahadThePure:
		res = "GalahadThePure"
	case GeoObject:
		res = "GeoObject"
	case GeoObjectUse:
		res = "GeoObjectUse"
	case Group:
		res = "Group"
	case GroupUse:
		res = "GroupUse"
	case KingArthur:
		res = "KingArthur"
	case KingArthurShape:
		res = "KingArthurShape"
	case KnightWhoSayNi:
		res = "KnightWhoSayNi"
	case Lancelot:
		res = "Lancelot"
	case LancelotAgregation:
		res = "LancelotAgregation"
	case LancelotAgregationUse:
		res = "LancelotAgregationUse"
	case LancelotCategory:
		res = "LancelotCategory"
	case LancelotCategoryUse:
		res = "LancelotCategoryUse"
	case MapObject:
		res = "MapObject"
	case MapObjectUse:
		res = "MapObjectUse"
	case Repository:
		res = "Repository"
	case SirRobin:
		res = "SirRobin"
	case TheBridge:
		res = "TheBridge"
	case TheNuteShape:
		res = "TheNuteShape"
	case TheNuteTransition:
		res = "TheNuteTransition"
	case User:
		res = "User"
	case UserUse:
		res = "UserUse"
	case WhatIsYourPreferedColor:
		res = "WhatIsYourPreferedColor"
	case Workspace:
		res = "Workspace"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AWitch:
		res = "AWitch"
	case *BlackKnightShape:
		res = "BlackKnightShape"
	case *BringYourDead:
		res = "BringYourDead"
	case *Document:
		res = "Document"
	case *DocumentUse:
		res = "DocumentUse"
	case *GalahadThePure:
		res = "GalahadThePure"
	case *GeoObject:
		res = "GeoObject"
	case *GeoObjectUse:
		res = "GeoObjectUse"
	case *Group:
		res = "Group"
	case *GroupUse:
		res = "GroupUse"
	case *KingArthur:
		res = "KingArthur"
	case *KingArthurShape:
		res = "KingArthurShape"
	case *KnightWhoSayNi:
		res = "KnightWhoSayNi"
	case *Lancelot:
		res = "Lancelot"
	case *LancelotAgregation:
		res = "LancelotAgregation"
	case *LancelotAgregationUse:
		res = "LancelotAgregationUse"
	case *LancelotCategory:
		res = "LancelotCategory"
	case *LancelotCategoryUse:
		res = "LancelotCategoryUse"
	case *MapObject:
		res = "MapObject"
	case *MapObjectUse:
		res = "MapObjectUse"
	case *Repository:
		res = "Repository"
	case *SirRobin:
		res = "SirRobin"
	case *TheBridge:
		res = "TheBridge"
	case *TheNuteShape:
		res = "TheNuteShape"
	case *TheNuteTransition:
		res = "TheNuteTransition"
	case *User:
		res = "User"
	case *UserUse:
		res = "UserUse"
	case *WhatIsYourPreferedColor:
		res = "WhatIsYourPreferedColor"
	case *Workspace:
		res = "Workspace"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case AWitch:
		res = []string{"Name", "EvolutionDirection", "X", "Y", "Width", "Height", "FillColor", "FillOpacity", "Stroke", "StrokeWidth", "RX"}
	case BlackKnightShape:
		res = []string{"Name", "BringYourDead", "Direction", "X", "Y", "Width", "Height", "FillColor", "FillOpacity", "Stroke", "StrokeWidth", "RX"}
	case BringYourDead:
		res = []string{"Name", "Tag", "Description", "Lancelots"}
	case Document:
		res = []string{"Name", "GeoObjectUse"}
	case DocumentUse:
		res = []string{"Name", "Document"}
	case GalahadThePure:
		res = []string{"Name", "Description"}
	case GeoObject:
		res = []string{"Name"}
	case GeoObjectUse:
		res = []string{"Name", "GeoObject"}
	case Group:
		res = []string{"Name", "UserUse"}
	case GroupUse:
		res = []string{"Name", "Group"}
	case KingArthur:
		res = []string{"Name", "IsWithProbaility", "Probability"}
	case KingArthurShape:
		res = []string{"Name", "ActorState", "X", "Y", "Width", "Height", "FillColor", "FillOpacity", "Stroke", "StrokeWidth", "RX"}
	case KnightWhoSayNi:
		res = []string{"Name", "Parameter", "Direction", "X", "Y", "Width", "Height", "FillColor", "FillOpacity", "Stroke", "StrokeWidth", "RX"}
	case Lancelot:
		res = []string{"Name", "Tag", "Description", "GroupUse", "DocumentUse", "GeoObjectUse"}
	case LancelotAgregation:
		res = []string{"Name", "ParameterUse"}
	case LancelotAgregationUse:
		res = []string{"Name", "ParameterAgregation"}
	case LancelotCategory:
		res = []string{"Name", "ParameterUse"}
	case LancelotCategoryUse:
		res = []string{"Name", "ParameterCategory"}
	case MapObject:
		res = []string{"Name"}
	case MapObjectUse:
		res = []string{"Name", "Map"}
	case Repository:
		res = []string{"Name", "ParameterUse", "GroupUse"}
	case SirRobin:
		res = []string{"Name", "Description", "Witches", "Arthurs", "KnightWhoSayNis", "BlackKnightShapes", "TheNuteShapes", "AxisOrign_X", "AxisOrign_Y", "VerticalAxis_Top_Y", "VerticalAxis_Bottom_Y", "VerticalAxis_StrokeWidth", "HorizontalAxis_Right_X", "Start", "End"}
	case TheBridge:
		res = []string{"Name", "Description", "WhatIsYourPreferedColor", "GroupUse", "GeoObjectUse", "MapUse", "IsNodeExpanded"}
	case TheNuteShape:
		res = []string{"Name", "ActorStateTransition", "Start", "End", "StartOrientation", "StartRatio", "EndOrientation", "EndRatio", "CornerOffsetRatio"}
	case TheNuteTransition:
		res = []string{"Name", "StartState", "EndState"}
	case User:
		res = []string{"Name"}
	case UserUse:
		res = []string{"Name", "User"}
	case WhatIsYourPreferedColor:
		res = []string{"Name", "Description", "Diagrams", "DiagramsNodeFolded", "KingArthurs", "KingArthurNodeFolded", "Nutes", "RRRR", "Galahard", "IIUU", "Lancelots", "LancelotsodeFolded", "BringYourDeadarameters", "RRRRT", "IsNodeExpanded"}
	case Workspace:
		res = []string{"Name", "SelectedDiagram", "A", "B", "C", "D", "E"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case AWitch:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SirRobin"
		rf.Fieldname = "Witches"
		res = append(res, rf)
	case BlackKnightShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SirRobin"
		rf.Fieldname = "BlackKnightShapes"
		res = append(res, rf)
	case BringYourDead:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "WhatIsYourPreferedColor"
		rf.Fieldname = "BringYourDeadarameters"
		res = append(res, rf)
	case Document:
		var rf ReverseField
		_ = rf
	case DocumentUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Lancelot"
		rf.Fieldname = "DocumentUse"
		res = append(res, rf)
	case GalahadThePure:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "WhatIsYourPreferedColor"
		rf.Fieldname = "Galahard"
		res = append(res, rf)
	case GeoObject:
		var rf ReverseField
		_ = rf
	case GeoObjectUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Document"
		rf.Fieldname = "GeoObjectUse"
		res = append(res, rf)
		rf.GongstructName = "Lancelot"
		rf.Fieldname = "GeoObjectUse"
		res = append(res, rf)
		rf.GongstructName = "TheBridge"
		rf.Fieldname = "GeoObjectUse"
		res = append(res, rf)
	case Group:
		var rf ReverseField
		_ = rf
	case GroupUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Lancelot"
		rf.Fieldname = "GroupUse"
		res = append(res, rf)
		rf.GongstructName = "Repository"
		rf.Fieldname = "GroupUse"
		res = append(res, rf)
		rf.GongstructName = "TheBridge"
		rf.Fieldname = "GroupUse"
		res = append(res, rf)
	case KingArthur:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "WhatIsYourPreferedColor"
		rf.Fieldname = "KingArthurs"
		res = append(res, rf)
	case KingArthurShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SirRobin"
		rf.Fieldname = "Arthurs"
		res = append(res, rf)
	case KnightWhoSayNi:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "LancelotAgregation"
		rf.Fieldname = "ParameterUse"
		res = append(res, rf)
		rf.GongstructName = "LancelotCategory"
		rf.Fieldname = "ParameterUse"
		res = append(res, rf)
		rf.GongstructName = "Repository"
		rf.Fieldname = "ParameterUse"
		res = append(res, rf)
		rf.GongstructName = "SirRobin"
		rf.Fieldname = "KnightWhoSayNis"
		res = append(res, rf)
	case Lancelot:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BringYourDead"
		rf.Fieldname = "Lancelots"
		res = append(res, rf)
		rf.GongstructName = "WhatIsYourPreferedColor"
		rf.Fieldname = "Lancelots"
		res = append(res, rf)
	case LancelotAgregation:
		var rf ReverseField
		_ = rf
	case LancelotAgregationUse:
		var rf ReverseField
		_ = rf
	case LancelotCategory:
		var rf ReverseField
		_ = rf
	case LancelotCategoryUse:
		var rf ReverseField
		_ = rf
	case MapObject:
		var rf ReverseField
		_ = rf
	case MapObjectUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TheBridge"
		rf.Fieldname = "MapUse"
		res = append(res, rf)
	case Repository:
		var rf ReverseField
		_ = rf
	case SirRobin:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "WhatIsYourPreferedColor"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
	case TheBridge:
		var rf ReverseField
		_ = rf
	case TheNuteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SirRobin"
		rf.Fieldname = "TheNuteShapes"
		res = append(res, rf)
	case TheNuteTransition:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "WhatIsYourPreferedColor"
		rf.Fieldname = "Nutes"
		res = append(res, rf)
	case User:
		var rf ReverseField
		_ = rf
	case UserUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "UserUse"
		res = append(res, rf)
	case WhatIsYourPreferedColor:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TheBridge"
		rf.Fieldname = "WhatIsYourPreferedColor"
		res = append(res, rf)
	case Workspace:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AWitch:
		res = []string{"Name", "EvolutionDirection", "X", "Y", "Width", "Height", "FillColor", "FillOpacity", "Stroke", "StrokeWidth", "RX"}
	case *BlackKnightShape:
		res = []string{"Name", "BringYourDead", "Direction", "X", "Y", "Width", "Height", "FillColor", "FillOpacity", "Stroke", "StrokeWidth", "RX"}
	case *BringYourDead:
		res = []string{"Name", "Tag", "Description", "Lancelots"}
	case *Document:
		res = []string{"Name", "GeoObjectUse"}
	case *DocumentUse:
		res = []string{"Name", "Document"}
	case *GalahadThePure:
		res = []string{"Name", "Description"}
	case *GeoObject:
		res = []string{"Name"}
	case *GeoObjectUse:
		res = []string{"Name", "GeoObject"}
	case *Group:
		res = []string{"Name", "UserUse"}
	case *GroupUse:
		res = []string{"Name", "Group"}
	case *KingArthur:
		res = []string{"Name", "IsWithProbaility", "Probability"}
	case *KingArthurShape:
		res = []string{"Name", "ActorState", "X", "Y", "Width", "Height", "FillColor", "FillOpacity", "Stroke", "StrokeWidth", "RX"}
	case *KnightWhoSayNi:
		res = []string{"Name", "Parameter", "Direction", "X", "Y", "Width", "Height", "FillColor", "FillOpacity", "Stroke", "StrokeWidth", "RX"}
	case *Lancelot:
		res = []string{"Name", "Tag", "Description", "GroupUse", "DocumentUse", "GeoObjectUse"}
	case *LancelotAgregation:
		res = []string{"Name", "ParameterUse"}
	case *LancelotAgregationUse:
		res = []string{"Name", "ParameterAgregation"}
	case *LancelotCategory:
		res = []string{"Name", "ParameterUse"}
	case *LancelotCategoryUse:
		res = []string{"Name", "ParameterCategory"}
	case *MapObject:
		res = []string{"Name"}
	case *MapObjectUse:
		res = []string{"Name", "Map"}
	case *Repository:
		res = []string{"Name", "ParameterUse", "GroupUse"}
	case *SirRobin:
		res = []string{"Name", "Description", "Witches", "Arthurs", "KnightWhoSayNis", "BlackKnightShapes", "TheNuteShapes", "AxisOrign_X", "AxisOrign_Y", "VerticalAxis_Top_Y", "VerticalAxis_Bottom_Y", "VerticalAxis_StrokeWidth", "HorizontalAxis_Right_X", "Start", "End"}
	case *TheBridge:
		res = []string{"Name", "Description", "WhatIsYourPreferedColor", "GroupUse", "GeoObjectUse", "MapUse", "IsNodeExpanded"}
	case *TheNuteShape:
		res = []string{"Name", "ActorStateTransition", "Start", "End", "StartOrientation", "StartRatio", "EndOrientation", "EndRatio", "CornerOffsetRatio"}
	case *TheNuteTransition:
		res = []string{"Name", "StartState", "EndState"}
	case *User:
		res = []string{"Name"}
	case *UserUse:
		res = []string{"Name", "User"}
	case *WhatIsYourPreferedColor:
		res = []string{"Name", "Description", "Diagrams", "DiagramsNodeFolded", "KingArthurs", "KingArthurNodeFolded", "Nutes", "RRRR", "Galahard", "IIUU", "Lancelots", "LancelotsodeFolded", "BringYourDeadarameters", "RRRRT", "IsNodeExpanded"}
	case *Workspace:
		res = []string{"Name", "SelectedDiagram", "A", "B", "C", "D", "E"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *AWitch:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EvolutionDirection":
			if inferedInstance.EvolutionDirection != nil {
				res = inferedInstance.EvolutionDirection.Name
			}
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Height":
			res = fmt.Sprintf("%f", inferedInstance.Height)
		case "FillColor":
			res = inferedInstance.FillColor
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "RX":
			res = fmt.Sprintf("%f", inferedInstance.RX)
		}
	case *BlackKnightShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "BringYourDead":
			if inferedInstance.BringYourDead != nil {
				res = inferedInstance.BringYourDead.Name
			}
		case "Direction":
			enum := inferedInstance.Direction
			res = enum.ToCodeString()
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Height":
			res = fmt.Sprintf("%f", inferedInstance.Height)
		case "FillColor":
			res = inferedInstance.FillColor
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "RX":
			res = fmt.Sprintf("%f", inferedInstance.RX)
		}
	case *BringYourDead:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Tag":
			res = inferedInstance.Tag
		case "Description":
			res = inferedInstance.Description
		case "Lancelots":
			for idx, __instance__ := range inferedInstance.Lancelots {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Document:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GeoObjectUse":
			for idx, __instance__ := range inferedInstance.GeoObjectUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *DocumentUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Document":
			if inferedInstance.Document != nil {
				res = inferedInstance.Document.Name
			}
		}
	case *GalahadThePure:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Description":
			res = inferedInstance.Description
		}
	case *GeoObject:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *GeoObjectUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GeoObject":
			if inferedInstance.GeoObject != nil {
				res = inferedInstance.GeoObject.Name
			}
		}
	case *Group:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "UserUse":
			for idx, __instance__ := range inferedInstance.UserUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *GroupUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Group":
			if inferedInstance.Group != nil {
				res = inferedInstance.Group.Name
			}
		}
	case *KingArthur:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsWithProbaility":
			res = fmt.Sprintf("%t", inferedInstance.IsWithProbaility)
		case "Probability":
			enum := inferedInstance.Probability
			res = enum.ToCodeString()
		}
	case *KingArthurShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ActorState":
			if inferedInstance.ActorState != nil {
				res = inferedInstance.ActorState.Name
			}
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Height":
			res = fmt.Sprintf("%f", inferedInstance.Height)
		case "FillColor":
			res = inferedInstance.FillColor
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "RX":
			res = fmt.Sprintf("%f", inferedInstance.RX)
		}
	case *KnightWhoSayNi:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Parameter":
			if inferedInstance.Parameter != nil {
				res = inferedInstance.Parameter.Name
			}
		case "Direction":
			enum := inferedInstance.Direction
			res = enum.ToCodeString()
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Height":
			res = fmt.Sprintf("%f", inferedInstance.Height)
		case "FillColor":
			res = inferedInstance.FillColor
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "RX":
			res = fmt.Sprintf("%f", inferedInstance.RX)
		}
	case *Lancelot:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Tag":
			res = inferedInstance.Tag
		case "Description":
			res = inferedInstance.Description
		case "GroupUse":
			for idx, __instance__ := range inferedInstance.GroupUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DocumentUse":
			for idx, __instance__ := range inferedInstance.DocumentUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GeoObjectUse":
			for idx, __instance__ := range inferedInstance.GeoObjectUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *LancelotAgregation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterUse":
			for idx, __instance__ := range inferedInstance.ParameterUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *LancelotAgregationUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterAgregation":
			if inferedInstance.ParameterAgregation != nil {
				res = inferedInstance.ParameterAgregation.Name
			}
		}
	case *LancelotCategory:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterUse":
			for idx, __instance__ := range inferedInstance.ParameterUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *LancelotCategoryUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterCategory":
			if inferedInstance.ParameterCategory != nil {
				res = inferedInstance.ParameterCategory.Name
			}
		}
	case *MapObject:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *MapObjectUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Map":
			if inferedInstance.Map != nil {
				res = inferedInstance.Map.Name
			}
		}
	case *Repository:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterUse":
			for idx, __instance__ := range inferedInstance.ParameterUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GroupUse":
			for idx, __instance__ := range inferedInstance.GroupUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *SirRobin:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Description":
			res = inferedInstance.Description
		case "Witches":
			for idx, __instance__ := range inferedInstance.Witches {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Arthurs":
			for idx, __instance__ := range inferedInstance.Arthurs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "KnightWhoSayNis":
			for idx, __instance__ := range inferedInstance.KnightWhoSayNis {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "BlackKnightShapes":
			for idx, __instance__ := range inferedInstance.BlackKnightShapes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "TheNuteShapes":
			for idx, __instance__ := range inferedInstance.TheNuteShapes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "AxisOrign_X":
			res = fmt.Sprintf("%f", inferedInstance.AxisOrign_X)
		case "AxisOrign_Y":
			res = fmt.Sprintf("%f", inferedInstance.AxisOrign_Y)
		case "VerticalAxis_Top_Y":
			res = fmt.Sprintf("%f", inferedInstance.VerticalAxis_Top_Y)
		case "VerticalAxis_Bottom_Y":
			res = fmt.Sprintf("%f", inferedInstance.VerticalAxis_Bottom_Y)
		case "VerticalAxis_StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.VerticalAxis_StrokeWidth)
		case "HorizontalAxis_Right_X":
			res = fmt.Sprintf("%f", inferedInstance.HorizontalAxis_Right_X)
		case "Start":
			res = inferedInstance.Start.String()
		case "End":
			res = inferedInstance.End.String()
		}
	case *TheBridge:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Description":
			res = inferedInstance.Description
		case "WhatIsYourPreferedColor":
			for idx, __instance__ := range inferedInstance.WhatIsYourPreferedColor {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GroupUse":
			for idx, __instance__ := range inferedInstance.GroupUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GeoObjectUse":
			for idx, __instance__ := range inferedInstance.GeoObjectUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "MapUse":
			for idx, __instance__ := range inferedInstance.MapUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "IsNodeExpanded":
			res = fmt.Sprintf("%t", inferedInstance.IsNodeExpanded)
		}
	case *TheNuteShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ActorStateTransition":
			if inferedInstance.ActorStateTransition != nil {
				res = inferedInstance.ActorStateTransition.Name
			}
		case "Start":
			if inferedInstance.Start != nil {
				res = inferedInstance.Start.Name
			}
		case "End":
			if inferedInstance.End != nil {
				res = inferedInstance.End.Name
			}
		case "StartOrientation":
			enum := inferedInstance.StartOrientation
			res = enum.ToCodeString()
		case "StartRatio":
			res = fmt.Sprintf("%f", inferedInstance.StartRatio)
		case "EndOrientation":
			enum := inferedInstance.EndOrientation
			res = enum.ToCodeString()
		case "EndRatio":
			res = fmt.Sprintf("%f", inferedInstance.EndRatio)
		case "CornerOffsetRatio":
			res = fmt.Sprintf("%f", inferedInstance.CornerOffsetRatio)
		}
	case *TheNuteTransition:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "StartState":
			if inferedInstance.StartState != nil {
				res = inferedInstance.StartState.Name
			}
		case "EndState":
			if inferedInstance.EndState != nil {
				res = inferedInstance.EndState.Name
			}
		}
	case *User:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *UserUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "User":
			if inferedInstance.User != nil {
				res = inferedInstance.User.Name
			}
		}
	case *WhatIsYourPreferedColor:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Description":
			res = inferedInstance.Description
		case "Diagrams":
			for idx, __instance__ := range inferedInstance.Diagrams {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DiagramsNodeFolded":
			res = fmt.Sprintf("%t", inferedInstance.DiagramsNodeFolded)
		case "KingArthurs":
			for idx, __instance__ := range inferedInstance.KingArthurs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "KingArthurNodeFolded":
			res = fmt.Sprintf("%t", inferedInstance.KingArthurNodeFolded)
		case "Nutes":
			for idx, __instance__ := range inferedInstance.Nutes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "RRRR":
			res = fmt.Sprintf("%t", inferedInstance.RRRR)
		case "Galahard":
			for idx, __instance__ := range inferedInstance.Galahard {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "IIUU":
			res = fmt.Sprintf("%t", inferedInstance.IIUU)
		case "Lancelots":
			for idx, __instance__ := range inferedInstance.Lancelots {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "LancelotsodeFolded":
			res = fmt.Sprintf("%t", inferedInstance.LancelotsodeFolded)
		case "BringYourDeadarameters":
			for idx, __instance__ := range inferedInstance.BringYourDeadarameters {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "RRRRT":
			res = fmt.Sprintf("%t", inferedInstance.RRRRT)
		case "IsNodeExpanded":
			res = fmt.Sprintf("%t", inferedInstance.IsNodeExpanded)
		}
	case *Workspace:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SelectedDiagram":
			if inferedInstance.SelectedDiagram != nil {
				res = inferedInstance.SelectedDiagram.Name
			}
		case "A":
			if inferedInstance.A != nil {
				res = inferedInstance.A.Name
			}
		case "B":
			if inferedInstance.B != nil {
				res = inferedInstance.B.Name
			}
		case "C":
			if inferedInstance.C != nil {
				res = inferedInstance.C.Name
			}
		case "D":
			if inferedInstance.D != nil {
				res = inferedInstance.D.Name
			}
		case "E":
			if inferedInstance.E != nil {
				res = inferedInstance.E.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case AWitch:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EvolutionDirection":
			if inferedInstance.EvolutionDirection != nil {
				res = inferedInstance.EvolutionDirection.Name
			}
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Height":
			res = fmt.Sprintf("%f", inferedInstance.Height)
		case "FillColor":
			res = inferedInstance.FillColor
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "RX":
			res = fmt.Sprintf("%f", inferedInstance.RX)
		}
	case BlackKnightShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "BringYourDead":
			if inferedInstance.BringYourDead != nil {
				res = inferedInstance.BringYourDead.Name
			}
		case "Direction":
			enum := inferedInstance.Direction
			res = enum.ToCodeString()
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Height":
			res = fmt.Sprintf("%f", inferedInstance.Height)
		case "FillColor":
			res = inferedInstance.FillColor
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "RX":
			res = fmt.Sprintf("%f", inferedInstance.RX)
		}
	case BringYourDead:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Tag":
			res = inferedInstance.Tag
		case "Description":
			res = inferedInstance.Description
		case "Lancelots":
			for idx, __instance__ := range inferedInstance.Lancelots {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Document:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GeoObjectUse":
			for idx, __instance__ := range inferedInstance.GeoObjectUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case DocumentUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Document":
			if inferedInstance.Document != nil {
				res = inferedInstance.Document.Name
			}
		}
	case GalahadThePure:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Description":
			res = inferedInstance.Description
		}
	case GeoObject:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case GeoObjectUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GeoObject":
			if inferedInstance.GeoObject != nil {
				res = inferedInstance.GeoObject.Name
			}
		}
	case Group:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "UserUse":
			for idx, __instance__ := range inferedInstance.UserUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case GroupUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Group":
			if inferedInstance.Group != nil {
				res = inferedInstance.Group.Name
			}
		}
	case KingArthur:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IsWithProbaility":
			res = fmt.Sprintf("%t", inferedInstance.IsWithProbaility)
		case "Probability":
			enum := inferedInstance.Probability
			res = enum.ToCodeString()
		}
	case KingArthurShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ActorState":
			if inferedInstance.ActorState != nil {
				res = inferedInstance.ActorState.Name
			}
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Height":
			res = fmt.Sprintf("%f", inferedInstance.Height)
		case "FillColor":
			res = inferedInstance.FillColor
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "RX":
			res = fmt.Sprintf("%f", inferedInstance.RX)
		}
	case KnightWhoSayNi:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Parameter":
			if inferedInstance.Parameter != nil {
				res = inferedInstance.Parameter.Name
			}
		case "Direction":
			enum := inferedInstance.Direction
			res = enum.ToCodeString()
		case "X":
			res = fmt.Sprintf("%f", inferedInstance.X)
		case "Y":
			res = fmt.Sprintf("%f", inferedInstance.Y)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Height":
			res = fmt.Sprintf("%f", inferedInstance.Height)
		case "FillColor":
			res = inferedInstance.FillColor
		case "FillOpacity":
			res = fmt.Sprintf("%f", inferedInstance.FillOpacity)
		case "Stroke":
			res = inferedInstance.Stroke
		case "StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
		case "RX":
			res = fmt.Sprintf("%f", inferedInstance.RX)
		}
	case Lancelot:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Tag":
			res = inferedInstance.Tag
		case "Description":
			res = inferedInstance.Description
		case "GroupUse":
			for idx, __instance__ := range inferedInstance.GroupUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DocumentUse":
			for idx, __instance__ := range inferedInstance.DocumentUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GeoObjectUse":
			for idx, __instance__ := range inferedInstance.GeoObjectUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case LancelotAgregation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterUse":
			for idx, __instance__ := range inferedInstance.ParameterUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case LancelotAgregationUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterAgregation":
			if inferedInstance.ParameterAgregation != nil {
				res = inferedInstance.ParameterAgregation.Name
			}
		}
	case LancelotCategory:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterUse":
			for idx, __instance__ := range inferedInstance.ParameterUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case LancelotCategoryUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterCategory":
			if inferedInstance.ParameterCategory != nil {
				res = inferedInstance.ParameterCategory.Name
			}
		}
	case MapObject:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case MapObjectUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Map":
			if inferedInstance.Map != nil {
				res = inferedInstance.Map.Name
			}
		}
	case Repository:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ParameterUse":
			for idx, __instance__ := range inferedInstance.ParameterUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GroupUse":
			for idx, __instance__ := range inferedInstance.GroupUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case SirRobin:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Description":
			res = inferedInstance.Description
		case "Witches":
			for idx, __instance__ := range inferedInstance.Witches {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Arthurs":
			for idx, __instance__ := range inferedInstance.Arthurs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "KnightWhoSayNis":
			for idx, __instance__ := range inferedInstance.KnightWhoSayNis {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "BlackKnightShapes":
			for idx, __instance__ := range inferedInstance.BlackKnightShapes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "TheNuteShapes":
			for idx, __instance__ := range inferedInstance.TheNuteShapes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "AxisOrign_X":
			res = fmt.Sprintf("%f", inferedInstance.AxisOrign_X)
		case "AxisOrign_Y":
			res = fmt.Sprintf("%f", inferedInstance.AxisOrign_Y)
		case "VerticalAxis_Top_Y":
			res = fmt.Sprintf("%f", inferedInstance.VerticalAxis_Top_Y)
		case "VerticalAxis_Bottom_Y":
			res = fmt.Sprintf("%f", inferedInstance.VerticalAxis_Bottom_Y)
		case "VerticalAxis_StrokeWidth":
			res = fmt.Sprintf("%f", inferedInstance.VerticalAxis_StrokeWidth)
		case "HorizontalAxis_Right_X":
			res = fmt.Sprintf("%f", inferedInstance.HorizontalAxis_Right_X)
		case "Start":
			res = inferedInstance.Start.String()
		case "End":
			res = inferedInstance.End.String()
		}
	case TheBridge:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Description":
			res = inferedInstance.Description
		case "WhatIsYourPreferedColor":
			for idx, __instance__ := range inferedInstance.WhatIsYourPreferedColor {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GroupUse":
			for idx, __instance__ := range inferedInstance.GroupUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GeoObjectUse":
			for idx, __instance__ := range inferedInstance.GeoObjectUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "MapUse":
			for idx, __instance__ := range inferedInstance.MapUse {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "IsNodeExpanded":
			res = fmt.Sprintf("%t", inferedInstance.IsNodeExpanded)
		}
	case TheNuteShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ActorStateTransition":
			if inferedInstance.ActorStateTransition != nil {
				res = inferedInstance.ActorStateTransition.Name
			}
		case "Start":
			if inferedInstance.Start != nil {
				res = inferedInstance.Start.Name
			}
		case "End":
			if inferedInstance.End != nil {
				res = inferedInstance.End.Name
			}
		case "StartOrientation":
			enum := inferedInstance.StartOrientation
			res = enum.ToCodeString()
		case "StartRatio":
			res = fmt.Sprintf("%f", inferedInstance.StartRatio)
		case "EndOrientation":
			enum := inferedInstance.EndOrientation
			res = enum.ToCodeString()
		case "EndRatio":
			res = fmt.Sprintf("%f", inferedInstance.EndRatio)
		case "CornerOffsetRatio":
			res = fmt.Sprintf("%f", inferedInstance.CornerOffsetRatio)
		}
	case TheNuteTransition:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "StartState":
			if inferedInstance.StartState != nil {
				res = inferedInstance.StartState.Name
			}
		case "EndState":
			if inferedInstance.EndState != nil {
				res = inferedInstance.EndState.Name
			}
		}
	case User:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case UserUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "User":
			if inferedInstance.User != nil {
				res = inferedInstance.User.Name
			}
		}
	case WhatIsYourPreferedColor:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Description":
			res = inferedInstance.Description
		case "Diagrams":
			for idx, __instance__ := range inferedInstance.Diagrams {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DiagramsNodeFolded":
			res = fmt.Sprintf("%t", inferedInstance.DiagramsNodeFolded)
		case "KingArthurs":
			for idx, __instance__ := range inferedInstance.KingArthurs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "KingArthurNodeFolded":
			res = fmt.Sprintf("%t", inferedInstance.KingArthurNodeFolded)
		case "Nutes":
			for idx, __instance__ := range inferedInstance.Nutes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "RRRR":
			res = fmt.Sprintf("%t", inferedInstance.RRRR)
		case "Galahard":
			for idx, __instance__ := range inferedInstance.Galahard {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "IIUU":
			res = fmt.Sprintf("%t", inferedInstance.IIUU)
		case "Lancelots":
			for idx, __instance__ := range inferedInstance.Lancelots {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "LancelotsodeFolded":
			res = fmt.Sprintf("%t", inferedInstance.LancelotsodeFolded)
		case "BringYourDeadarameters":
			for idx, __instance__ := range inferedInstance.BringYourDeadarameters {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "RRRRT":
			res = fmt.Sprintf("%t", inferedInstance.RRRRT)
		case "IsNodeExpanded":
			res = fmt.Sprintf("%t", inferedInstance.IsNodeExpanded)
		}
	case Workspace:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SelectedDiagram":
			if inferedInstance.SelectedDiagram != nil {
				res = inferedInstance.SelectedDiagram.Name
			}
		case "A":
			if inferedInstance.A != nil {
				res = inferedInstance.A.Name
			}
		case "B":
			if inferedInstance.B != nil {
				res = inferedInstance.B.Name
			}
		case "C":
			if inferedInstance.C != nil {
				res = inferedInstance.C.Name
			}
		case "D":
			if inferedInstance.D != nil {
				res = inferedInstance.D.Name
			}
		case "E":
			if inferedInstance.E != nil {
				res = inferedInstance.E.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
