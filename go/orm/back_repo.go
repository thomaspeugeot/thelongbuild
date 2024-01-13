// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/thomaspeugeot/thelongbuild/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAWitch BackRepoAWitchStruct

	BackRepoBlackKnightShape BackRepoBlackKnightShapeStruct

	BackRepoBringYourDead BackRepoBringYourDeadStruct

	BackRepoDocument BackRepoDocumentStruct

	BackRepoDocumentUse BackRepoDocumentUseStruct

	BackRepoGalahadThePure BackRepoGalahadThePureStruct

	BackRepoGeoObject BackRepoGeoObjectStruct

	BackRepoGeoObjectUse BackRepoGeoObjectUseStruct

	BackRepoGroup BackRepoGroupStruct

	BackRepoGroupUse BackRepoGroupUseStruct

	BackRepoKingArthur BackRepoKingArthurStruct

	BackRepoKingArthurShape BackRepoKingArthurShapeStruct

	BackRepoKnightWhoSayNi BackRepoKnightWhoSayNiStruct

	BackRepoLancelot BackRepoLancelotStruct

	BackRepoLancelotAgregation BackRepoLancelotAgregationStruct

	BackRepoLancelotAgregationUse BackRepoLancelotAgregationUseStruct

	BackRepoLancelotCategory BackRepoLancelotCategoryStruct

	BackRepoLancelotCategoryUse BackRepoLancelotCategoryUseStruct

	BackRepoMapObject BackRepoMapObjectStruct

	BackRepoMapObjectUse BackRepoMapObjectUseStruct

	BackRepoRepository BackRepoRepositoryStruct

	BackRepoSirRobin BackRepoSirRobinStruct

	BackRepoTheBridge BackRepoTheBridgeStruct

	BackRepoTheNuteShape BackRepoTheNuteShapeStruct

	BackRepoTheNuteTransition BackRepoTheNuteTransitionStruct

	BackRepoUser BackRepoUserStruct

	BackRepoUserUse BackRepoUserUseStruct

	BackRepoWhatIsYourPreferedColor BackRepoWhatIsYourPreferedColorStruct

	BackRepoWorkspace BackRepoWorkspaceStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&AWitchDB{},
		&BlackKnightShapeDB{},
		&BringYourDeadDB{},
		&DocumentDB{},
		&DocumentUseDB{},
		&GalahadThePureDB{},
		&GeoObjectDB{},
		&GeoObjectUseDB{},
		&GroupDB{},
		&GroupUseDB{},
		&KingArthurDB{},
		&KingArthurShapeDB{},
		&KnightWhoSayNiDB{},
		&LancelotDB{},
		&LancelotAgregationDB{},
		&LancelotAgregationUseDB{},
		&LancelotCategoryDB{},
		&LancelotCategoryUseDB{},
		&MapObjectDB{},
		&MapObjectUseDB{},
		&RepositoryDB{},
		&SirRobinDB{},
		&TheBridgeDB{},
		&TheNuteShapeDB{},
		&TheNuteTransitionDB{},
		&UserDB{},
		&UserUseDB{},
		&WhatIsYourPreferedColorDB{},
		&WorkspaceDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAWitch = BackRepoAWitchStruct{
		Map_AWitchDBID_AWitchPtr: make(map[uint]*models.AWitch, 0),
		Map_AWitchDBID_AWitchDB:  make(map[uint]*AWitchDB, 0),
		Map_AWitchPtr_AWitchDBID: make(map[*models.AWitch]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBlackKnightShape = BackRepoBlackKnightShapeStruct{
		Map_BlackKnightShapeDBID_BlackKnightShapePtr: make(map[uint]*models.BlackKnightShape, 0),
		Map_BlackKnightShapeDBID_BlackKnightShapeDB:  make(map[uint]*BlackKnightShapeDB, 0),
		Map_BlackKnightShapePtr_BlackKnightShapeDBID: make(map[*models.BlackKnightShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBringYourDead = BackRepoBringYourDeadStruct{
		Map_BringYourDeadDBID_BringYourDeadPtr: make(map[uint]*models.BringYourDead, 0),
		Map_BringYourDeadDBID_BringYourDeadDB:  make(map[uint]*BringYourDeadDB, 0),
		Map_BringYourDeadPtr_BringYourDeadDBID: make(map[*models.BringYourDead]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDocument = BackRepoDocumentStruct{
		Map_DocumentDBID_DocumentPtr: make(map[uint]*models.Document, 0),
		Map_DocumentDBID_DocumentDB:  make(map[uint]*DocumentDB, 0),
		Map_DocumentPtr_DocumentDBID: make(map[*models.Document]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDocumentUse = BackRepoDocumentUseStruct{
		Map_DocumentUseDBID_DocumentUsePtr: make(map[uint]*models.DocumentUse, 0),
		Map_DocumentUseDBID_DocumentUseDB:  make(map[uint]*DocumentUseDB, 0),
		Map_DocumentUsePtr_DocumentUseDBID: make(map[*models.DocumentUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGalahadThePure = BackRepoGalahadThePureStruct{
		Map_GalahadThePureDBID_GalahadThePurePtr: make(map[uint]*models.GalahadThePure, 0),
		Map_GalahadThePureDBID_GalahadThePureDB:  make(map[uint]*GalahadThePureDB, 0),
		Map_GalahadThePurePtr_GalahadThePureDBID: make(map[*models.GalahadThePure]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGeoObject = BackRepoGeoObjectStruct{
		Map_GeoObjectDBID_GeoObjectPtr: make(map[uint]*models.GeoObject, 0),
		Map_GeoObjectDBID_GeoObjectDB:  make(map[uint]*GeoObjectDB, 0),
		Map_GeoObjectPtr_GeoObjectDBID: make(map[*models.GeoObject]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGeoObjectUse = BackRepoGeoObjectUseStruct{
		Map_GeoObjectUseDBID_GeoObjectUsePtr: make(map[uint]*models.GeoObjectUse, 0),
		Map_GeoObjectUseDBID_GeoObjectUseDB:  make(map[uint]*GeoObjectUseDB, 0),
		Map_GeoObjectUsePtr_GeoObjectUseDBID: make(map[*models.GeoObjectUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGroup = BackRepoGroupStruct{
		Map_GroupDBID_GroupPtr: make(map[uint]*models.Group, 0),
		Map_GroupDBID_GroupDB:  make(map[uint]*GroupDB, 0),
		Map_GroupPtr_GroupDBID: make(map[*models.Group]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGroupUse = BackRepoGroupUseStruct{
		Map_GroupUseDBID_GroupUsePtr: make(map[uint]*models.GroupUse, 0),
		Map_GroupUseDBID_GroupUseDB:  make(map[uint]*GroupUseDB, 0),
		Map_GroupUsePtr_GroupUseDBID: make(map[*models.GroupUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKingArthur = BackRepoKingArthurStruct{
		Map_KingArthurDBID_KingArthurPtr: make(map[uint]*models.KingArthur, 0),
		Map_KingArthurDBID_KingArthurDB:  make(map[uint]*KingArthurDB, 0),
		Map_KingArthurPtr_KingArthurDBID: make(map[*models.KingArthur]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKingArthurShape = BackRepoKingArthurShapeStruct{
		Map_KingArthurShapeDBID_KingArthurShapePtr: make(map[uint]*models.KingArthurShape, 0),
		Map_KingArthurShapeDBID_KingArthurShapeDB:  make(map[uint]*KingArthurShapeDB, 0),
		Map_KingArthurShapePtr_KingArthurShapeDBID: make(map[*models.KingArthurShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKnightWhoSayNi = BackRepoKnightWhoSayNiStruct{
		Map_KnightWhoSayNiDBID_KnightWhoSayNiPtr: make(map[uint]*models.KnightWhoSayNi, 0),
		Map_KnightWhoSayNiDBID_KnightWhoSayNiDB:  make(map[uint]*KnightWhoSayNiDB, 0),
		Map_KnightWhoSayNiPtr_KnightWhoSayNiDBID: make(map[*models.KnightWhoSayNi]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLancelot = BackRepoLancelotStruct{
		Map_LancelotDBID_LancelotPtr: make(map[uint]*models.Lancelot, 0),
		Map_LancelotDBID_LancelotDB:  make(map[uint]*LancelotDB, 0),
		Map_LancelotPtr_LancelotDBID: make(map[*models.Lancelot]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLancelotAgregation = BackRepoLancelotAgregationStruct{
		Map_LancelotAgregationDBID_LancelotAgregationPtr: make(map[uint]*models.LancelotAgregation, 0),
		Map_LancelotAgregationDBID_LancelotAgregationDB:  make(map[uint]*LancelotAgregationDB, 0),
		Map_LancelotAgregationPtr_LancelotAgregationDBID: make(map[*models.LancelotAgregation]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLancelotAgregationUse = BackRepoLancelotAgregationUseStruct{
		Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr: make(map[uint]*models.LancelotAgregationUse, 0),
		Map_LancelotAgregationUseDBID_LancelotAgregationUseDB:  make(map[uint]*LancelotAgregationUseDB, 0),
		Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID: make(map[*models.LancelotAgregationUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLancelotCategory = BackRepoLancelotCategoryStruct{
		Map_LancelotCategoryDBID_LancelotCategoryPtr: make(map[uint]*models.LancelotCategory, 0),
		Map_LancelotCategoryDBID_LancelotCategoryDB:  make(map[uint]*LancelotCategoryDB, 0),
		Map_LancelotCategoryPtr_LancelotCategoryDBID: make(map[*models.LancelotCategory]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLancelotCategoryUse = BackRepoLancelotCategoryUseStruct{
		Map_LancelotCategoryUseDBID_LancelotCategoryUsePtr: make(map[uint]*models.LancelotCategoryUse, 0),
		Map_LancelotCategoryUseDBID_LancelotCategoryUseDB:  make(map[uint]*LancelotCategoryUseDB, 0),
		Map_LancelotCategoryUsePtr_LancelotCategoryUseDBID: make(map[*models.LancelotCategoryUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMapObject = BackRepoMapObjectStruct{
		Map_MapObjectDBID_MapObjectPtr: make(map[uint]*models.MapObject, 0),
		Map_MapObjectDBID_MapObjectDB:  make(map[uint]*MapObjectDB, 0),
		Map_MapObjectPtr_MapObjectDBID: make(map[*models.MapObject]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMapObjectUse = BackRepoMapObjectUseStruct{
		Map_MapObjectUseDBID_MapObjectUsePtr: make(map[uint]*models.MapObjectUse, 0),
		Map_MapObjectUseDBID_MapObjectUseDB:  make(map[uint]*MapObjectUseDB, 0),
		Map_MapObjectUsePtr_MapObjectUseDBID: make(map[*models.MapObjectUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRepository = BackRepoRepositoryStruct{
		Map_RepositoryDBID_RepositoryPtr: make(map[uint]*models.Repository, 0),
		Map_RepositoryDBID_RepositoryDB:  make(map[uint]*RepositoryDB, 0),
		Map_RepositoryPtr_RepositoryDBID: make(map[*models.Repository]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSirRobin = BackRepoSirRobinStruct{
		Map_SirRobinDBID_SirRobinPtr: make(map[uint]*models.SirRobin, 0),
		Map_SirRobinDBID_SirRobinDB:  make(map[uint]*SirRobinDB, 0),
		Map_SirRobinPtr_SirRobinDBID: make(map[*models.SirRobin]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTheBridge = BackRepoTheBridgeStruct{
		Map_TheBridgeDBID_TheBridgePtr: make(map[uint]*models.TheBridge, 0),
		Map_TheBridgeDBID_TheBridgeDB:  make(map[uint]*TheBridgeDB, 0),
		Map_TheBridgePtr_TheBridgeDBID: make(map[*models.TheBridge]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTheNuteShape = BackRepoTheNuteShapeStruct{
		Map_TheNuteShapeDBID_TheNuteShapePtr: make(map[uint]*models.TheNuteShape, 0),
		Map_TheNuteShapeDBID_TheNuteShapeDB:  make(map[uint]*TheNuteShapeDB, 0),
		Map_TheNuteShapePtr_TheNuteShapeDBID: make(map[*models.TheNuteShape]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTheNuteTransition = BackRepoTheNuteTransitionStruct{
		Map_TheNuteTransitionDBID_TheNuteTransitionPtr: make(map[uint]*models.TheNuteTransition, 0),
		Map_TheNuteTransitionDBID_TheNuteTransitionDB:  make(map[uint]*TheNuteTransitionDB, 0),
		Map_TheNuteTransitionPtr_TheNuteTransitionDBID: make(map[*models.TheNuteTransition]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoUser = BackRepoUserStruct{
		Map_UserDBID_UserPtr: make(map[uint]*models.User, 0),
		Map_UserDBID_UserDB:  make(map[uint]*UserDB, 0),
		Map_UserPtr_UserDBID: make(map[*models.User]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoUserUse = BackRepoUserUseStruct{
		Map_UserUseDBID_UserUsePtr: make(map[uint]*models.UserUse, 0),
		Map_UserUseDBID_UserUseDB:  make(map[uint]*UserUseDB, 0),
		Map_UserUsePtr_UserUseDBID: make(map[*models.UserUse]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoWhatIsYourPreferedColor = BackRepoWhatIsYourPreferedColorStruct{
		Map_WhatIsYourPreferedColorDBID_WhatIsYourPreferedColorPtr: make(map[uint]*models.WhatIsYourPreferedColor, 0),
		Map_WhatIsYourPreferedColorDBID_WhatIsYourPreferedColorDB:  make(map[uint]*WhatIsYourPreferedColorDB, 0),
		Map_WhatIsYourPreferedColorPtr_WhatIsYourPreferedColorDBID: make(map[*models.WhatIsYourPreferedColor]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoWorkspace = BackRepoWorkspaceStruct{
		Map_WorkspaceDBID_WorkspacePtr: make(map[uint]*models.Workspace, 0),
		Map_WorkspaceDBID_WorkspaceDB:  make(map[uint]*WorkspaceDB, 0),
		Map_WorkspacePtr_WorkspaceDBID: make(map[*models.Workspace]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAWitch.CommitPhaseOne(stage)
	backRepo.BackRepoBlackKnightShape.CommitPhaseOne(stage)
	backRepo.BackRepoBringYourDead.CommitPhaseOne(stage)
	backRepo.BackRepoDocument.CommitPhaseOne(stage)
	backRepo.BackRepoDocumentUse.CommitPhaseOne(stage)
	backRepo.BackRepoGalahadThePure.CommitPhaseOne(stage)
	backRepo.BackRepoGeoObject.CommitPhaseOne(stage)
	backRepo.BackRepoGeoObjectUse.CommitPhaseOne(stage)
	backRepo.BackRepoGroup.CommitPhaseOne(stage)
	backRepo.BackRepoGroupUse.CommitPhaseOne(stage)
	backRepo.BackRepoKingArthur.CommitPhaseOne(stage)
	backRepo.BackRepoKingArthurShape.CommitPhaseOne(stage)
	backRepo.BackRepoKnightWhoSayNi.CommitPhaseOne(stage)
	backRepo.BackRepoLancelot.CommitPhaseOne(stage)
	backRepo.BackRepoLancelotAgregation.CommitPhaseOne(stage)
	backRepo.BackRepoLancelotAgregationUse.CommitPhaseOne(stage)
	backRepo.BackRepoLancelotCategory.CommitPhaseOne(stage)
	backRepo.BackRepoLancelotCategoryUse.CommitPhaseOne(stage)
	backRepo.BackRepoMapObject.CommitPhaseOne(stage)
	backRepo.BackRepoMapObjectUse.CommitPhaseOne(stage)
	backRepo.BackRepoRepository.CommitPhaseOne(stage)
	backRepo.BackRepoSirRobin.CommitPhaseOne(stage)
	backRepo.BackRepoTheBridge.CommitPhaseOne(stage)
	backRepo.BackRepoTheNuteShape.CommitPhaseOne(stage)
	backRepo.BackRepoTheNuteTransition.CommitPhaseOne(stage)
	backRepo.BackRepoUser.CommitPhaseOne(stage)
	backRepo.BackRepoUserUse.CommitPhaseOne(stage)
	backRepo.BackRepoWhatIsYourPreferedColor.CommitPhaseOne(stage)
	backRepo.BackRepoWorkspace.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAWitch.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBlackKnightShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBringYourDead.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDocument.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDocumentUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGalahadThePure.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGeoObject.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGeoObjectUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGroup.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGroupUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKingArthur.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKingArthurShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKnightWhoSayNi.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLancelot.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLancelotAgregation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLancelotAgregationUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLancelotCategory.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLancelotCategoryUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMapObject.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMapObjectUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRepository.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSirRobin.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTheBridge.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTheNuteShape.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTheNuteTransition.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUser.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUserUse.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWhatIsYourPreferedColor.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWorkspace.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAWitch.CheckoutPhaseOne()
	backRepo.BackRepoBlackKnightShape.CheckoutPhaseOne()
	backRepo.BackRepoBringYourDead.CheckoutPhaseOne()
	backRepo.BackRepoDocument.CheckoutPhaseOne()
	backRepo.BackRepoDocumentUse.CheckoutPhaseOne()
	backRepo.BackRepoGalahadThePure.CheckoutPhaseOne()
	backRepo.BackRepoGeoObject.CheckoutPhaseOne()
	backRepo.BackRepoGeoObjectUse.CheckoutPhaseOne()
	backRepo.BackRepoGroup.CheckoutPhaseOne()
	backRepo.BackRepoGroupUse.CheckoutPhaseOne()
	backRepo.BackRepoKingArthur.CheckoutPhaseOne()
	backRepo.BackRepoKingArthurShape.CheckoutPhaseOne()
	backRepo.BackRepoKnightWhoSayNi.CheckoutPhaseOne()
	backRepo.BackRepoLancelot.CheckoutPhaseOne()
	backRepo.BackRepoLancelotAgregation.CheckoutPhaseOne()
	backRepo.BackRepoLancelotAgregationUse.CheckoutPhaseOne()
	backRepo.BackRepoLancelotCategory.CheckoutPhaseOne()
	backRepo.BackRepoLancelotCategoryUse.CheckoutPhaseOne()
	backRepo.BackRepoMapObject.CheckoutPhaseOne()
	backRepo.BackRepoMapObjectUse.CheckoutPhaseOne()
	backRepo.BackRepoRepository.CheckoutPhaseOne()
	backRepo.BackRepoSirRobin.CheckoutPhaseOne()
	backRepo.BackRepoTheBridge.CheckoutPhaseOne()
	backRepo.BackRepoTheNuteShape.CheckoutPhaseOne()
	backRepo.BackRepoTheNuteTransition.CheckoutPhaseOne()
	backRepo.BackRepoUser.CheckoutPhaseOne()
	backRepo.BackRepoUserUse.CheckoutPhaseOne()
	backRepo.BackRepoWhatIsYourPreferedColor.CheckoutPhaseOne()
	backRepo.BackRepoWorkspace.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAWitch.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBlackKnightShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBringYourDead.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDocument.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDocumentUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGalahadThePure.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGeoObject.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGeoObjectUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGroup.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGroupUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKingArthur.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKingArthurShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKnightWhoSayNi.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLancelot.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLancelotAgregation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLancelotAgregationUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLancelotCategory.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLancelotCategoryUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMapObject.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMapObjectUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRepository.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSirRobin.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTheBridge.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTheNuteShape.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTheNuteTransition.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUser.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUserUse.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWhatIsYourPreferedColor.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWorkspace.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAWitch.Backup(dirPath)
	backRepo.BackRepoBlackKnightShape.Backup(dirPath)
	backRepo.BackRepoBringYourDead.Backup(dirPath)
	backRepo.BackRepoDocument.Backup(dirPath)
	backRepo.BackRepoDocumentUse.Backup(dirPath)
	backRepo.BackRepoGalahadThePure.Backup(dirPath)
	backRepo.BackRepoGeoObject.Backup(dirPath)
	backRepo.BackRepoGeoObjectUse.Backup(dirPath)
	backRepo.BackRepoGroup.Backup(dirPath)
	backRepo.BackRepoGroupUse.Backup(dirPath)
	backRepo.BackRepoKingArthur.Backup(dirPath)
	backRepo.BackRepoKingArthurShape.Backup(dirPath)
	backRepo.BackRepoKnightWhoSayNi.Backup(dirPath)
	backRepo.BackRepoLancelot.Backup(dirPath)
	backRepo.BackRepoLancelotAgregation.Backup(dirPath)
	backRepo.BackRepoLancelotAgregationUse.Backup(dirPath)
	backRepo.BackRepoLancelotCategory.Backup(dirPath)
	backRepo.BackRepoLancelotCategoryUse.Backup(dirPath)
	backRepo.BackRepoMapObject.Backup(dirPath)
	backRepo.BackRepoMapObjectUse.Backup(dirPath)
	backRepo.BackRepoRepository.Backup(dirPath)
	backRepo.BackRepoSirRobin.Backup(dirPath)
	backRepo.BackRepoTheBridge.Backup(dirPath)
	backRepo.BackRepoTheNuteShape.Backup(dirPath)
	backRepo.BackRepoTheNuteTransition.Backup(dirPath)
	backRepo.BackRepoUser.Backup(dirPath)
	backRepo.BackRepoUserUse.Backup(dirPath)
	backRepo.BackRepoWhatIsYourPreferedColor.Backup(dirPath)
	backRepo.BackRepoWorkspace.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAWitch.BackupXL(file)
	backRepo.BackRepoBlackKnightShape.BackupXL(file)
	backRepo.BackRepoBringYourDead.BackupXL(file)
	backRepo.BackRepoDocument.BackupXL(file)
	backRepo.BackRepoDocumentUse.BackupXL(file)
	backRepo.BackRepoGalahadThePure.BackupXL(file)
	backRepo.BackRepoGeoObject.BackupXL(file)
	backRepo.BackRepoGeoObjectUse.BackupXL(file)
	backRepo.BackRepoGroup.BackupXL(file)
	backRepo.BackRepoGroupUse.BackupXL(file)
	backRepo.BackRepoKingArthur.BackupXL(file)
	backRepo.BackRepoKingArthurShape.BackupXL(file)
	backRepo.BackRepoKnightWhoSayNi.BackupXL(file)
	backRepo.BackRepoLancelot.BackupXL(file)
	backRepo.BackRepoLancelotAgregation.BackupXL(file)
	backRepo.BackRepoLancelotAgregationUse.BackupXL(file)
	backRepo.BackRepoLancelotCategory.BackupXL(file)
	backRepo.BackRepoLancelotCategoryUse.BackupXL(file)
	backRepo.BackRepoMapObject.BackupXL(file)
	backRepo.BackRepoMapObjectUse.BackupXL(file)
	backRepo.BackRepoRepository.BackupXL(file)
	backRepo.BackRepoSirRobin.BackupXL(file)
	backRepo.BackRepoTheBridge.BackupXL(file)
	backRepo.BackRepoTheNuteShape.BackupXL(file)
	backRepo.BackRepoTheNuteTransition.BackupXL(file)
	backRepo.BackRepoUser.BackupXL(file)
	backRepo.BackRepoUserUse.BackupXL(file)
	backRepo.BackRepoWhatIsYourPreferedColor.BackupXL(file)
	backRepo.BackRepoWorkspace.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAWitch.RestorePhaseOne(dirPath)
	backRepo.BackRepoBlackKnightShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoBringYourDead.RestorePhaseOne(dirPath)
	backRepo.BackRepoDocument.RestorePhaseOne(dirPath)
	backRepo.BackRepoDocumentUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoGalahadThePure.RestorePhaseOne(dirPath)
	backRepo.BackRepoGeoObject.RestorePhaseOne(dirPath)
	backRepo.BackRepoGeoObjectUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoGroup.RestorePhaseOne(dirPath)
	backRepo.BackRepoGroupUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoKingArthur.RestorePhaseOne(dirPath)
	backRepo.BackRepoKingArthurShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoKnightWhoSayNi.RestorePhaseOne(dirPath)
	backRepo.BackRepoLancelot.RestorePhaseOne(dirPath)
	backRepo.BackRepoLancelotAgregation.RestorePhaseOne(dirPath)
	backRepo.BackRepoLancelotAgregationUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoLancelotCategory.RestorePhaseOne(dirPath)
	backRepo.BackRepoLancelotCategoryUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoMapObject.RestorePhaseOne(dirPath)
	backRepo.BackRepoMapObjectUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoRepository.RestorePhaseOne(dirPath)
	backRepo.BackRepoSirRobin.RestorePhaseOne(dirPath)
	backRepo.BackRepoTheBridge.RestorePhaseOne(dirPath)
	backRepo.BackRepoTheNuteShape.RestorePhaseOne(dirPath)
	backRepo.BackRepoTheNuteTransition.RestorePhaseOne(dirPath)
	backRepo.BackRepoUser.RestorePhaseOne(dirPath)
	backRepo.BackRepoUserUse.RestorePhaseOne(dirPath)
	backRepo.BackRepoWhatIsYourPreferedColor.RestorePhaseOne(dirPath)
	backRepo.BackRepoWorkspace.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAWitch.RestorePhaseTwo()
	backRepo.BackRepoBlackKnightShape.RestorePhaseTwo()
	backRepo.BackRepoBringYourDead.RestorePhaseTwo()
	backRepo.BackRepoDocument.RestorePhaseTwo()
	backRepo.BackRepoDocumentUse.RestorePhaseTwo()
	backRepo.BackRepoGalahadThePure.RestorePhaseTwo()
	backRepo.BackRepoGeoObject.RestorePhaseTwo()
	backRepo.BackRepoGeoObjectUse.RestorePhaseTwo()
	backRepo.BackRepoGroup.RestorePhaseTwo()
	backRepo.BackRepoGroupUse.RestorePhaseTwo()
	backRepo.BackRepoKingArthur.RestorePhaseTwo()
	backRepo.BackRepoKingArthurShape.RestorePhaseTwo()
	backRepo.BackRepoKnightWhoSayNi.RestorePhaseTwo()
	backRepo.BackRepoLancelot.RestorePhaseTwo()
	backRepo.BackRepoLancelotAgregation.RestorePhaseTwo()
	backRepo.BackRepoLancelotAgregationUse.RestorePhaseTwo()
	backRepo.BackRepoLancelotCategory.RestorePhaseTwo()
	backRepo.BackRepoLancelotCategoryUse.RestorePhaseTwo()
	backRepo.BackRepoMapObject.RestorePhaseTwo()
	backRepo.BackRepoMapObjectUse.RestorePhaseTwo()
	backRepo.BackRepoRepository.RestorePhaseTwo()
	backRepo.BackRepoSirRobin.RestorePhaseTwo()
	backRepo.BackRepoTheBridge.RestorePhaseTwo()
	backRepo.BackRepoTheNuteShape.RestorePhaseTwo()
	backRepo.BackRepoTheNuteTransition.RestorePhaseTwo()
	backRepo.BackRepoUser.RestorePhaseTwo()
	backRepo.BackRepoUserUse.RestorePhaseTwo()
	backRepo.BackRepoWhatIsYourPreferedColor.RestorePhaseTwo()
	backRepo.BackRepoWorkspace.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAWitch.RestoreXLPhaseOne(file)
	backRepo.BackRepoBlackKnightShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoBringYourDead.RestoreXLPhaseOne(file)
	backRepo.BackRepoDocument.RestoreXLPhaseOne(file)
	backRepo.BackRepoDocumentUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoGalahadThePure.RestoreXLPhaseOne(file)
	backRepo.BackRepoGeoObject.RestoreXLPhaseOne(file)
	backRepo.BackRepoGeoObjectUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoGroup.RestoreXLPhaseOne(file)
	backRepo.BackRepoGroupUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoKingArthur.RestoreXLPhaseOne(file)
	backRepo.BackRepoKingArthurShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoKnightWhoSayNi.RestoreXLPhaseOne(file)
	backRepo.BackRepoLancelot.RestoreXLPhaseOne(file)
	backRepo.BackRepoLancelotAgregation.RestoreXLPhaseOne(file)
	backRepo.BackRepoLancelotAgregationUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoLancelotCategory.RestoreXLPhaseOne(file)
	backRepo.BackRepoLancelotCategoryUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoMapObject.RestoreXLPhaseOne(file)
	backRepo.BackRepoMapObjectUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoRepository.RestoreXLPhaseOne(file)
	backRepo.BackRepoSirRobin.RestoreXLPhaseOne(file)
	backRepo.BackRepoTheBridge.RestoreXLPhaseOne(file)
	backRepo.BackRepoTheNuteShape.RestoreXLPhaseOne(file)
	backRepo.BackRepoTheNuteTransition.RestoreXLPhaseOne(file)
	backRepo.BackRepoUser.RestoreXLPhaseOne(file)
	backRepo.BackRepoUserUse.RestoreXLPhaseOne(file)
	backRepo.BackRepoWhatIsYourPreferedColor.RestoreXLPhaseOne(file)
	backRepo.BackRepoWorkspace.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}
