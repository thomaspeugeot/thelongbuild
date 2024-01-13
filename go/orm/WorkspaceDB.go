// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/thomaspeugeot/thelongbuild/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Workspace_sql sql.NullBool
var dummy_Workspace_time time.Duration
var dummy_Workspace_sort sort.Float64Slice

// WorkspaceAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model workspaceAPI
type WorkspaceAPI struct {
	gorm.Model

	models.Workspace_WOP

	// encoding of pointers
	WorkspacePointersEncoding WorkspacePointersEncoding
}

// WorkspacePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type WorkspacePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field SelectedDiagram is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	SelectedDiagramID sql.NullInt64

	// field A is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	AID sql.NullInt64

	// field B is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	BID sql.NullInt64

	// field C is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	CID sql.NullInt64

	// field D is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	DID sql.NullInt64

	// field E is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	EID sql.NullInt64
}

// WorkspaceDB describes a workspace in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model workspaceDB
type WorkspaceDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field workspaceDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	WorkspacePointersEncoding
}

// WorkspaceDBs arrays workspaceDBs
// swagger:response workspaceDBsResponse
type WorkspaceDBs []WorkspaceDB

// WorkspaceDBResponse provides response
// swagger:response workspaceDBResponse
type WorkspaceDBResponse struct {
	WorkspaceDB
}

// WorkspaceWOP is a Workspace without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type WorkspaceWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Workspace_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoWorkspaceStruct struct {
	// stores WorkspaceDB according to their gorm ID
	Map_WorkspaceDBID_WorkspaceDB map[uint]*WorkspaceDB

	// stores WorkspaceDB ID according to Workspace address
	Map_WorkspacePtr_WorkspaceDBID map[*models.Workspace]uint

	// stores Workspace according to their gorm ID
	Map_WorkspaceDBID_WorkspacePtr map[uint]*models.Workspace

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoWorkspace *BackRepoWorkspaceStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoWorkspace.stage
	return
}

func (backRepoWorkspace *BackRepoWorkspaceStruct) GetDB() *gorm.DB {
	return backRepoWorkspace.db
}

// GetWorkspaceDBFromWorkspacePtr is a handy function to access the back repo instance from the stage instance
func (backRepoWorkspace *BackRepoWorkspaceStruct) GetWorkspaceDBFromWorkspacePtr(workspace *models.Workspace) (workspaceDB *WorkspaceDB) {
	id := backRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID[workspace]
	workspaceDB = backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB[id]
	return
}

// BackRepoWorkspace.CommitPhaseOne commits all staged instances of Workspace to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoWorkspace *BackRepoWorkspaceStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for workspace := range stage.Workspaces {
		backRepoWorkspace.CommitPhaseOneInstance(workspace)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, workspace := range backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr {
		if _, ok := stage.Workspaces[workspace]; !ok {
			backRepoWorkspace.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoWorkspace.CommitDeleteInstance commits deletion of Workspace to the BackRepo
func (backRepoWorkspace *BackRepoWorkspaceStruct) CommitDeleteInstance(id uint) (Error error) {

	workspace := backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[id]

	// workspace is not staged anymore, remove workspaceDB
	workspaceDB := backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB[id]
	query := backRepoWorkspace.db.Unscoped().Delete(&workspaceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID, workspace)
	delete(backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr, id)
	delete(backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB, id)

	return
}

// BackRepoWorkspace.CommitPhaseOneInstance commits workspace staged instances of Workspace to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoWorkspace *BackRepoWorkspaceStruct) CommitPhaseOneInstance(workspace *models.Workspace) (Error error) {

	// check if the workspace is not commited yet
	if _, ok := backRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID[workspace]; ok {
		return
	}

	// initiate workspace
	var workspaceDB WorkspaceDB
	workspaceDB.CopyBasicFieldsFromWorkspace(workspace)

	query := backRepoWorkspace.db.Create(&workspaceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID[workspace] = workspaceDB.ID
	backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[workspaceDB.ID] = workspace
	backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB[workspaceDB.ID] = &workspaceDB

	return
}

// BackRepoWorkspace.CommitPhaseTwo commits all staged instances of Workspace to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWorkspace *BackRepoWorkspaceStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, workspace := range backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr {
		backRepoWorkspace.CommitPhaseTwoInstance(backRepo, idx, workspace)
	}

	return
}

// BackRepoWorkspace.CommitPhaseTwoInstance commits {{structname }} of models.Workspace to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWorkspace *BackRepoWorkspaceStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, workspace *models.Workspace) (Error error) {

	// fetch matching workspaceDB
	if workspaceDB, ok := backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB[idx]; ok {

		workspaceDB.CopyBasicFieldsFromWorkspace(workspace)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value workspace.SelectedDiagram translates to updating the workspace.SelectedDiagramID
		workspaceDB.SelectedDiagramID.Valid = true // allow for a 0 value (nil association)
		if workspace.SelectedDiagram != nil {
			if SelectedDiagramId, ok := backRepo.BackRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[workspace.SelectedDiagram]; ok {
				workspaceDB.SelectedDiagramID.Int64 = int64(SelectedDiagramId)
				workspaceDB.SelectedDiagramID.Valid = true
			}
		} else {
			workspaceDB.SelectedDiagramID.Int64 = 0
			workspaceDB.SelectedDiagramID.Valid = true
		}

		// commit pointer value workspace.A translates to updating the workspace.AID
		workspaceDB.AID.Valid = true // allow for a 0 value (nil association)
		if workspace.A != nil {
			if AId, ok := backRepo.BackRepoAWitch.Map_AWitchPtr_AWitchDBID[workspace.A]; ok {
				workspaceDB.AID.Int64 = int64(AId)
				workspaceDB.AID.Valid = true
			}
		} else {
			workspaceDB.AID.Int64 = 0
			workspaceDB.AID.Valid = true
		}

		// commit pointer value workspace.B translates to updating the workspace.BID
		workspaceDB.BID.Valid = true // allow for a 0 value (nil association)
		if workspace.B != nil {
			if BId, ok := backRepo.BackRepoKnightWhoSayNi.Map_KnightWhoSayNiPtr_KnightWhoSayNiDBID[workspace.B]; ok {
				workspaceDB.BID.Int64 = int64(BId)
				workspaceDB.BID.Valid = true
			}
		} else {
			workspaceDB.BID.Int64 = 0
			workspaceDB.BID.Valid = true
		}

		// commit pointer value workspace.C translates to updating the workspace.CID
		workspaceDB.CID.Valid = true // allow for a 0 value (nil association)
		if workspace.C != nil {
			if CId, ok := backRepo.BackRepoBlackKnightShape.Map_BlackKnightShapePtr_BlackKnightShapeDBID[workspace.C]; ok {
				workspaceDB.CID.Int64 = int64(CId)
				workspaceDB.CID.Valid = true
			}
		} else {
			workspaceDB.CID.Int64 = 0
			workspaceDB.CID.Valid = true
		}

		// commit pointer value workspace.D translates to updating the workspace.DID
		workspaceDB.DID.Valid = true // allow for a 0 value (nil association)
		if workspace.D != nil {
			if DId, ok := backRepo.BackRepoKingArthurShape.Map_KingArthurShapePtr_KingArthurShapeDBID[workspace.D]; ok {
				workspaceDB.DID.Int64 = int64(DId)
				workspaceDB.DID.Valid = true
			}
		} else {
			workspaceDB.DID.Int64 = 0
			workspaceDB.DID.Valid = true
		}

		// commit pointer value workspace.E translates to updating the workspace.EID
		workspaceDB.EID.Valid = true // allow for a 0 value (nil association)
		if workspace.E != nil {
			if EId, ok := backRepo.BackRepoTheNuteShape.Map_TheNuteShapePtr_TheNuteShapeDBID[workspace.E]; ok {
				workspaceDB.EID.Int64 = int64(EId)
				workspaceDB.EID.Valid = true
			}
		} else {
			workspaceDB.EID.Int64 = 0
			workspaceDB.EID.Valid = true
		}

		query := backRepoWorkspace.db.Save(&workspaceDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Workspace intance %s", workspace.Name))
		return err
	}

	return
}

// BackRepoWorkspace.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoWorkspace *BackRepoWorkspaceStruct) CheckoutPhaseOne() (Error error) {

	workspaceDBArray := make([]WorkspaceDB, 0)
	query := backRepoWorkspace.db.Find(&workspaceDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	workspaceInstancesToBeRemovedFromTheStage := make(map[*models.Workspace]any)
	for key, value := range backRepoWorkspace.stage.Workspaces {
		workspaceInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, workspaceDB := range workspaceDBArray {
		backRepoWorkspace.CheckoutPhaseOneInstance(&workspaceDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		workspace, ok := backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[workspaceDB.ID]
		if ok {
			delete(workspaceInstancesToBeRemovedFromTheStage, workspace)
		}
	}

	// remove from stage and back repo's 3 maps all workspaces that are not in the checkout
	for workspace := range workspaceInstancesToBeRemovedFromTheStage {
		workspace.Unstage(backRepoWorkspace.GetStage())

		// remove instance from the back repo 3 maps
		workspaceID := backRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID[workspace]
		delete(backRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID, workspace)
		delete(backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB, workspaceID)
		delete(backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr, workspaceID)
	}

	return
}

// CheckoutPhaseOneInstance takes a workspaceDB that has been found in the DB, updates the backRepo and stages the
// models version of the workspaceDB
func (backRepoWorkspace *BackRepoWorkspaceStruct) CheckoutPhaseOneInstance(workspaceDB *WorkspaceDB) (Error error) {

	workspace, ok := backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[workspaceDB.ID]
	if !ok {
		workspace = new(models.Workspace)

		backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[workspaceDB.ID] = workspace
		backRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID[workspace] = workspaceDB.ID

		// append model store with the new element
		workspace.Name = workspaceDB.Name_Data.String
		workspace.Stage(backRepoWorkspace.GetStage())
	}
	workspaceDB.CopyBasicFieldsToWorkspace(workspace)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	workspace.Stage(backRepoWorkspace.GetStage())

	// preserve pointer to workspaceDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_WorkspaceDBID_WorkspaceDB)[workspaceDB hold variable pointers
	workspaceDB_Data := *workspaceDB
	preservedPtrToWorkspace := &workspaceDB_Data
	backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB[workspaceDB.ID] = preservedPtrToWorkspace

	return
}

// BackRepoWorkspace.CheckoutPhaseTwo Checkouts all staged instances of Workspace to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWorkspace *BackRepoWorkspaceStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, workspaceDB := range backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB {
		backRepoWorkspace.CheckoutPhaseTwoInstance(backRepo, workspaceDB)
	}
	return
}

// BackRepoWorkspace.CheckoutPhaseTwoInstance Checkouts staged instances of Workspace to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWorkspace *BackRepoWorkspaceStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, workspaceDB *WorkspaceDB) (Error error) {

	workspace := backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[workspaceDB.ID]

	workspaceDB.DecodePointers(backRepo, workspace)

	return
}

func (workspaceDB *WorkspaceDB) DecodePointers(backRepo *BackRepoStruct, workspace *models.Workspace) {

	// insertion point for checkout of pointer encoding
	// SelectedDiagram field
	workspace.SelectedDiagram = nil
	if workspaceDB.SelectedDiagramID.Int64 != 0 {
		workspace.SelectedDiagram = backRepo.BackRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[uint(workspaceDB.SelectedDiagramID.Int64)]
	}
	// A field
	workspace.A = nil
	if workspaceDB.AID.Int64 != 0 {
		workspace.A = backRepo.BackRepoAWitch.Map_AWitchDBID_AWitchPtr[uint(workspaceDB.AID.Int64)]
	}
	// B field
	workspace.B = nil
	if workspaceDB.BID.Int64 != 0 {
		workspace.B = backRepo.BackRepoKnightWhoSayNi.Map_KnightWhoSayNiDBID_KnightWhoSayNiPtr[uint(workspaceDB.BID.Int64)]
	}
	// C field
	workspace.C = nil
	if workspaceDB.CID.Int64 != 0 {
		workspace.C = backRepo.BackRepoBlackKnightShape.Map_BlackKnightShapeDBID_BlackKnightShapePtr[uint(workspaceDB.CID.Int64)]
	}
	// D field
	workspace.D = nil
	if workspaceDB.DID.Int64 != 0 {
		workspace.D = backRepo.BackRepoKingArthurShape.Map_KingArthurShapeDBID_KingArthurShapePtr[uint(workspaceDB.DID.Int64)]
	}
	// E field
	workspace.E = nil
	if workspaceDB.EID.Int64 != 0 {
		workspace.E = backRepo.BackRepoTheNuteShape.Map_TheNuteShapeDBID_TheNuteShapePtr[uint(workspaceDB.EID.Int64)]
	}
	return
}

// CommitWorkspace allows commit of a single workspace (if already staged)
func (backRepo *BackRepoStruct) CommitWorkspace(workspace *models.Workspace) {
	backRepo.BackRepoWorkspace.CommitPhaseOneInstance(workspace)
	if id, ok := backRepo.BackRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID[workspace]; ok {
		backRepo.BackRepoWorkspace.CommitPhaseTwoInstance(backRepo, id, workspace)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitWorkspace allows checkout of a single workspace (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutWorkspace(workspace *models.Workspace) {
	// check if the workspace is staged
	if _, ok := backRepo.BackRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID[workspace]; ok {

		if id, ok := backRepo.BackRepoWorkspace.Map_WorkspacePtr_WorkspaceDBID[workspace]; ok {
			var workspaceDB WorkspaceDB
			workspaceDB.ID = id

			if err := backRepo.BackRepoWorkspace.db.First(&workspaceDB, id).Error; err != nil {
				log.Fatalln("CheckoutWorkspace : Problem with getting object with id:", id)
			}
			backRepo.BackRepoWorkspace.CheckoutPhaseOneInstance(&workspaceDB)
			backRepo.BackRepoWorkspace.CheckoutPhaseTwoInstance(backRepo, &workspaceDB)
		}
	}
}

// CopyBasicFieldsFromWorkspace
func (workspaceDB *WorkspaceDB) CopyBasicFieldsFromWorkspace(workspace *models.Workspace) {
	// insertion point for fields commit

	workspaceDB.Name_Data.String = workspace.Name
	workspaceDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromWorkspace_WOP
func (workspaceDB *WorkspaceDB) CopyBasicFieldsFromWorkspace_WOP(workspace *models.Workspace_WOP) {
	// insertion point for fields commit

	workspaceDB.Name_Data.String = workspace.Name
	workspaceDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromWorkspaceWOP
func (workspaceDB *WorkspaceDB) CopyBasicFieldsFromWorkspaceWOP(workspace *WorkspaceWOP) {
	// insertion point for fields commit

	workspaceDB.Name_Data.String = workspace.Name
	workspaceDB.Name_Data.Valid = true
}

// CopyBasicFieldsToWorkspace
func (workspaceDB *WorkspaceDB) CopyBasicFieldsToWorkspace(workspace *models.Workspace) {
	// insertion point for checkout of basic fields (back repo to stage)
	workspace.Name = workspaceDB.Name_Data.String
}

// CopyBasicFieldsToWorkspace_WOP
func (workspaceDB *WorkspaceDB) CopyBasicFieldsToWorkspace_WOP(workspace *models.Workspace_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	workspace.Name = workspaceDB.Name_Data.String
}

// CopyBasicFieldsToWorkspaceWOP
func (workspaceDB *WorkspaceDB) CopyBasicFieldsToWorkspaceWOP(workspace *WorkspaceWOP) {
	workspace.ID = int(workspaceDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	workspace.Name = workspaceDB.Name_Data.String
}

// Backup generates a json file from a slice of all WorkspaceDB instances in the backrepo
func (backRepoWorkspace *BackRepoWorkspaceStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "WorkspaceDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*WorkspaceDB, 0)
	for _, workspaceDB := range backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB {
		forBackup = append(forBackup, workspaceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Workspace ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Workspace file", err.Error())
	}
}

// Backup generates a json file from a slice of all WorkspaceDB instances in the backrepo
func (backRepoWorkspace *BackRepoWorkspaceStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*WorkspaceDB, 0)
	for _, workspaceDB := range backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB {
		forBackup = append(forBackup, workspaceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Workspace")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Workspace_Fields, -1)
	for _, workspaceDB := range forBackup {

		var workspaceWOP WorkspaceWOP
		workspaceDB.CopyBasicFieldsToWorkspaceWOP(&workspaceWOP)

		row := sh.AddRow()
		row.WriteStruct(&workspaceWOP, -1)
	}
}

// RestoreXL from the "Workspace" sheet all WorkspaceDB instances
func (backRepoWorkspace *BackRepoWorkspaceStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoWorkspaceid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Workspace"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoWorkspace.rowVisitorWorkspace)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoWorkspace *BackRepoWorkspaceStruct) rowVisitorWorkspace(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var workspaceWOP WorkspaceWOP
		row.ReadStruct(&workspaceWOP)

		// add the unmarshalled struct to the stage
		workspaceDB := new(WorkspaceDB)
		workspaceDB.CopyBasicFieldsFromWorkspaceWOP(&workspaceWOP)

		workspaceDB_ID_atBackupTime := workspaceDB.ID
		workspaceDB.ID = 0
		query := backRepoWorkspace.db.Create(workspaceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB[workspaceDB.ID] = workspaceDB
		BackRepoWorkspaceid_atBckpTime_newID[workspaceDB_ID_atBackupTime] = workspaceDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "WorkspaceDB.json" in dirPath that stores an array
// of WorkspaceDB and stores it in the database
// the map BackRepoWorkspaceid_atBckpTime_newID is updated accordingly
func (backRepoWorkspace *BackRepoWorkspaceStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoWorkspaceid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "WorkspaceDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Workspace file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*WorkspaceDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_WorkspaceDBID_WorkspaceDB
	for _, workspaceDB := range forRestore {

		workspaceDB_ID_atBackupTime := workspaceDB.ID
		workspaceDB.ID = 0
		query := backRepoWorkspace.db.Create(workspaceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB[workspaceDB.ID] = workspaceDB
		BackRepoWorkspaceid_atBckpTime_newID[workspaceDB_ID_atBackupTime] = workspaceDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Workspace file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Workspace>id_atBckpTime_newID
// to compute new index
func (backRepoWorkspace *BackRepoWorkspaceStruct) RestorePhaseTwo() {

	for _, workspaceDB := range backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB {

		// next line of code is to avert unused variable compilation error
		_ = workspaceDB

		// insertion point for reindexing pointers encoding
		// reindexing SelectedDiagram field
		if workspaceDB.SelectedDiagramID.Int64 != 0 {
			workspaceDB.SelectedDiagramID.Int64 = int64(BackRepoSirRobinid_atBckpTime_newID[uint(workspaceDB.SelectedDiagramID.Int64)])
			workspaceDB.SelectedDiagramID.Valid = true
		}

		// reindexing A field
		if workspaceDB.AID.Int64 != 0 {
			workspaceDB.AID.Int64 = int64(BackRepoAWitchid_atBckpTime_newID[uint(workspaceDB.AID.Int64)])
			workspaceDB.AID.Valid = true
		}

		// reindexing B field
		if workspaceDB.BID.Int64 != 0 {
			workspaceDB.BID.Int64 = int64(BackRepoKnightWhoSayNiid_atBckpTime_newID[uint(workspaceDB.BID.Int64)])
			workspaceDB.BID.Valid = true
		}

		// reindexing C field
		if workspaceDB.CID.Int64 != 0 {
			workspaceDB.CID.Int64 = int64(BackRepoBlackKnightShapeid_atBckpTime_newID[uint(workspaceDB.CID.Int64)])
			workspaceDB.CID.Valid = true
		}

		// reindexing D field
		if workspaceDB.DID.Int64 != 0 {
			workspaceDB.DID.Int64 = int64(BackRepoKingArthurShapeid_atBckpTime_newID[uint(workspaceDB.DID.Int64)])
			workspaceDB.DID.Valid = true
		}

		// reindexing E field
		if workspaceDB.EID.Int64 != 0 {
			workspaceDB.EID.Int64 = int64(BackRepoTheNuteShapeid_atBckpTime_newID[uint(workspaceDB.EID.Int64)])
			workspaceDB.EID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoWorkspace.db.Model(workspaceDB).Updates(*workspaceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoWorkspace.ResetReversePointers commits all staged instances of Workspace to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoWorkspace *BackRepoWorkspaceStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, workspace := range backRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr {
		backRepoWorkspace.ResetReversePointersInstance(backRepo, idx, workspace)
	}

	return
}

func (backRepoWorkspace *BackRepoWorkspaceStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, workspace *models.Workspace) (Error error) {

	// fetch matching workspaceDB
	if workspaceDB, ok := backRepoWorkspace.Map_WorkspaceDBID_WorkspaceDB[idx]; ok {
		_ = workspaceDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoWorkspaceid_atBckpTime_newID map[uint]uint
