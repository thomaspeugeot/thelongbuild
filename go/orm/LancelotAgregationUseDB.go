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
var dummy_LancelotAgregationUse_sql sql.NullBool
var dummy_LancelotAgregationUse_time time.Duration
var dummy_LancelotAgregationUse_sort sort.Float64Slice

// LancelotAgregationUseAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model lancelotagregationuseAPI
type LancelotAgregationUseAPI struct {
	gorm.Model

	models.LancelotAgregationUse_WOP

	// encoding of pointers
	LancelotAgregationUsePointersEncoding LancelotAgregationUsePointersEncoding
}

// LancelotAgregationUsePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LancelotAgregationUsePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ParameterAgregation is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ParameterAgregationID sql.NullInt64
}

// LancelotAgregationUseDB describes a lancelotagregationuse in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model lancelotagregationuseDB
type LancelotAgregationUseDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field lancelotagregationuseDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	LancelotAgregationUsePointersEncoding
}

// LancelotAgregationUseDBs arrays lancelotagregationuseDBs
// swagger:response lancelotagregationuseDBsResponse
type LancelotAgregationUseDBs []LancelotAgregationUseDB

// LancelotAgregationUseDBResponse provides response
// swagger:response lancelotagregationuseDBResponse
type LancelotAgregationUseDBResponse struct {
	LancelotAgregationUseDB
}

// LancelotAgregationUseWOP is a LancelotAgregationUse without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LancelotAgregationUseWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var LancelotAgregationUse_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoLancelotAgregationUseStruct struct {
	// stores LancelotAgregationUseDB according to their gorm ID
	Map_LancelotAgregationUseDBID_LancelotAgregationUseDB map[uint]*LancelotAgregationUseDB

	// stores LancelotAgregationUseDB ID according to LancelotAgregationUse address
	Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID map[*models.LancelotAgregationUse]uint

	// stores LancelotAgregationUse according to their gorm ID
	Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr map[uint]*models.LancelotAgregationUse

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoLancelotAgregationUse.stage
	return
}

func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) GetDB() *gorm.DB {
	return backRepoLancelotAgregationUse.db
}

// GetLancelotAgregationUseDBFromLancelotAgregationUsePtr is a handy function to access the back repo instance from the stage instance
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) GetLancelotAgregationUseDBFromLancelotAgregationUsePtr(lancelotagregationuse *models.LancelotAgregationUse) (lancelotagregationuseDB *LancelotAgregationUseDB) {
	id := backRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID[lancelotagregationuse]
	lancelotagregationuseDB = backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB[id]
	return
}

// BackRepoLancelotAgregationUse.CommitPhaseOne commits all staged instances of LancelotAgregationUse to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for lancelotagregationuse := range stage.LancelotAgregationUses {
		backRepoLancelotAgregationUse.CommitPhaseOneInstance(lancelotagregationuse)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, lancelotagregationuse := range backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr {
		if _, ok := stage.LancelotAgregationUses[lancelotagregationuse]; !ok {
			backRepoLancelotAgregationUse.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLancelotAgregationUse.CommitDeleteInstance commits deletion of LancelotAgregationUse to the BackRepo
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CommitDeleteInstance(id uint) (Error error) {

	lancelotagregationuse := backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[id]

	// lancelotagregationuse is not staged anymore, remove lancelotagregationuseDB
	lancelotagregationuseDB := backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB[id]
	query := backRepoLancelotAgregationUse.db.Unscoped().Delete(&lancelotagregationuseDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID, lancelotagregationuse)
	delete(backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr, id)
	delete(backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB, id)

	return
}

// BackRepoLancelotAgregationUse.CommitPhaseOneInstance commits lancelotagregationuse staged instances of LancelotAgregationUse to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CommitPhaseOneInstance(lancelotagregationuse *models.LancelotAgregationUse) (Error error) {

	// check if the lancelotagregationuse is not commited yet
	if _, ok := backRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID[lancelotagregationuse]; ok {
		return
	}

	// initiate lancelotagregationuse
	var lancelotagregationuseDB LancelotAgregationUseDB
	lancelotagregationuseDB.CopyBasicFieldsFromLancelotAgregationUse(lancelotagregationuse)

	query := backRepoLancelotAgregationUse.db.Create(&lancelotagregationuseDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID[lancelotagregationuse] = lancelotagregationuseDB.ID
	backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[lancelotagregationuseDB.ID] = lancelotagregationuse
	backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB[lancelotagregationuseDB.ID] = &lancelotagregationuseDB

	return
}

// BackRepoLancelotAgregationUse.CommitPhaseTwo commits all staged instances of LancelotAgregationUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, lancelotagregationuse := range backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr {
		backRepoLancelotAgregationUse.CommitPhaseTwoInstance(backRepo, idx, lancelotagregationuse)
	}

	return
}

// BackRepoLancelotAgregationUse.CommitPhaseTwoInstance commits {{structname }} of models.LancelotAgregationUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, lancelotagregationuse *models.LancelotAgregationUse) (Error error) {

	// fetch matching lancelotagregationuseDB
	if lancelotagregationuseDB, ok := backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB[idx]; ok {

		lancelotagregationuseDB.CopyBasicFieldsFromLancelotAgregationUse(lancelotagregationuse)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value lancelotagregationuse.ParameterAgregation translates to updating the lancelotagregationuse.ParameterAgregationID
		lancelotagregationuseDB.ParameterAgregationID.Valid = true // allow for a 0 value (nil association)
		if lancelotagregationuse.ParameterAgregation != nil {
			if ParameterAgregationId, ok := backRepo.BackRepoLancelotAgregation.Map_LancelotAgregationPtr_LancelotAgregationDBID[lancelotagregationuse.ParameterAgregation]; ok {
				lancelotagregationuseDB.ParameterAgregationID.Int64 = int64(ParameterAgregationId)
				lancelotagregationuseDB.ParameterAgregationID.Valid = true
			}
		} else {
			lancelotagregationuseDB.ParameterAgregationID.Int64 = 0
			lancelotagregationuseDB.ParameterAgregationID.Valid = true
		}

		query := backRepoLancelotAgregationUse.db.Save(&lancelotagregationuseDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown LancelotAgregationUse intance %s", lancelotagregationuse.Name))
		return err
	}

	return
}

// BackRepoLancelotAgregationUse.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CheckoutPhaseOne() (Error error) {

	lancelotagregationuseDBArray := make([]LancelotAgregationUseDB, 0)
	query := backRepoLancelotAgregationUse.db.Find(&lancelotagregationuseDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	lancelotagregationuseInstancesToBeRemovedFromTheStage := make(map[*models.LancelotAgregationUse]any)
	for key, value := range backRepoLancelotAgregationUse.stage.LancelotAgregationUses {
		lancelotagregationuseInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, lancelotagregationuseDB := range lancelotagregationuseDBArray {
		backRepoLancelotAgregationUse.CheckoutPhaseOneInstance(&lancelotagregationuseDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		lancelotagregationuse, ok := backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[lancelotagregationuseDB.ID]
		if ok {
			delete(lancelotagregationuseInstancesToBeRemovedFromTheStage, lancelotagregationuse)
		}
	}

	// remove from stage and back repo's 3 maps all lancelotagregationuses that are not in the checkout
	for lancelotagregationuse := range lancelotagregationuseInstancesToBeRemovedFromTheStage {
		lancelotagregationuse.Unstage(backRepoLancelotAgregationUse.GetStage())

		// remove instance from the back repo 3 maps
		lancelotagregationuseID := backRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID[lancelotagregationuse]
		delete(backRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID, lancelotagregationuse)
		delete(backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB, lancelotagregationuseID)
		delete(backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr, lancelotagregationuseID)
	}

	return
}

// CheckoutPhaseOneInstance takes a lancelotagregationuseDB that has been found in the DB, updates the backRepo and stages the
// models version of the lancelotagregationuseDB
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CheckoutPhaseOneInstance(lancelotagregationuseDB *LancelotAgregationUseDB) (Error error) {

	lancelotagregationuse, ok := backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[lancelotagregationuseDB.ID]
	if !ok {
		lancelotagregationuse = new(models.LancelotAgregationUse)

		backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[lancelotagregationuseDB.ID] = lancelotagregationuse
		backRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID[lancelotagregationuse] = lancelotagregationuseDB.ID

		// append model store with the new element
		lancelotagregationuse.Name = lancelotagregationuseDB.Name_Data.String
		lancelotagregationuse.Stage(backRepoLancelotAgregationUse.GetStage())
	}
	lancelotagregationuseDB.CopyBasicFieldsToLancelotAgregationUse(lancelotagregationuse)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	lancelotagregationuse.Stage(backRepoLancelotAgregationUse.GetStage())

	// preserve pointer to lancelotagregationuseDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LancelotAgregationUseDBID_LancelotAgregationUseDB)[lancelotagregationuseDB hold variable pointers
	lancelotagregationuseDB_Data := *lancelotagregationuseDB
	preservedPtrToLancelotAgregationUse := &lancelotagregationuseDB_Data
	backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB[lancelotagregationuseDB.ID] = preservedPtrToLancelotAgregationUse

	return
}

// BackRepoLancelotAgregationUse.CheckoutPhaseTwo Checkouts all staged instances of LancelotAgregationUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, lancelotagregationuseDB := range backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB {
		backRepoLancelotAgregationUse.CheckoutPhaseTwoInstance(backRepo, lancelotagregationuseDB)
	}
	return
}

// BackRepoLancelotAgregationUse.CheckoutPhaseTwoInstance Checkouts staged instances of LancelotAgregationUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, lancelotagregationuseDB *LancelotAgregationUseDB) (Error error) {

	lancelotagregationuse := backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[lancelotagregationuseDB.ID]

	lancelotagregationuseDB.DecodePointers(backRepo, lancelotagregationuse)

	return
}

func (lancelotagregationuseDB *LancelotAgregationUseDB) DecodePointers(backRepo *BackRepoStruct, lancelotagregationuse *models.LancelotAgregationUse) {

	// insertion point for checkout of pointer encoding
	// ParameterAgregation field
	lancelotagregationuse.ParameterAgregation = nil
	if lancelotagregationuseDB.ParameterAgregationID.Int64 != 0 {
		lancelotagregationuse.ParameterAgregation = backRepo.BackRepoLancelotAgregation.Map_LancelotAgregationDBID_LancelotAgregationPtr[uint(lancelotagregationuseDB.ParameterAgregationID.Int64)]
	}
	return
}

// CommitLancelotAgregationUse allows commit of a single lancelotagregationuse (if already staged)
func (backRepo *BackRepoStruct) CommitLancelotAgregationUse(lancelotagregationuse *models.LancelotAgregationUse) {
	backRepo.BackRepoLancelotAgregationUse.CommitPhaseOneInstance(lancelotagregationuse)
	if id, ok := backRepo.BackRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID[lancelotagregationuse]; ok {
		backRepo.BackRepoLancelotAgregationUse.CommitPhaseTwoInstance(backRepo, id, lancelotagregationuse)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitLancelotAgregationUse allows checkout of a single lancelotagregationuse (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLancelotAgregationUse(lancelotagregationuse *models.LancelotAgregationUse) {
	// check if the lancelotagregationuse is staged
	if _, ok := backRepo.BackRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID[lancelotagregationuse]; ok {

		if id, ok := backRepo.BackRepoLancelotAgregationUse.Map_LancelotAgregationUsePtr_LancelotAgregationUseDBID[lancelotagregationuse]; ok {
			var lancelotagregationuseDB LancelotAgregationUseDB
			lancelotagregationuseDB.ID = id

			if err := backRepo.BackRepoLancelotAgregationUse.db.First(&lancelotagregationuseDB, id).Error; err != nil {
				log.Fatalln("CheckoutLancelotAgregationUse : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLancelotAgregationUse.CheckoutPhaseOneInstance(&lancelotagregationuseDB)
			backRepo.BackRepoLancelotAgregationUse.CheckoutPhaseTwoInstance(backRepo, &lancelotagregationuseDB)
		}
	}
}

// CopyBasicFieldsFromLancelotAgregationUse
func (lancelotagregationuseDB *LancelotAgregationUseDB) CopyBasicFieldsFromLancelotAgregationUse(lancelotagregationuse *models.LancelotAgregationUse) {
	// insertion point for fields commit

	lancelotagregationuseDB.Name_Data.String = lancelotagregationuse.Name
	lancelotagregationuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromLancelotAgregationUse_WOP
func (lancelotagregationuseDB *LancelotAgregationUseDB) CopyBasicFieldsFromLancelotAgregationUse_WOP(lancelotagregationuse *models.LancelotAgregationUse_WOP) {
	// insertion point for fields commit

	lancelotagregationuseDB.Name_Data.String = lancelotagregationuse.Name
	lancelotagregationuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromLancelotAgregationUseWOP
func (lancelotagregationuseDB *LancelotAgregationUseDB) CopyBasicFieldsFromLancelotAgregationUseWOP(lancelotagregationuse *LancelotAgregationUseWOP) {
	// insertion point for fields commit

	lancelotagregationuseDB.Name_Data.String = lancelotagregationuse.Name
	lancelotagregationuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsToLancelotAgregationUse
func (lancelotagregationuseDB *LancelotAgregationUseDB) CopyBasicFieldsToLancelotAgregationUse(lancelotagregationuse *models.LancelotAgregationUse) {
	// insertion point for checkout of basic fields (back repo to stage)
	lancelotagregationuse.Name = lancelotagregationuseDB.Name_Data.String
}

// CopyBasicFieldsToLancelotAgregationUse_WOP
func (lancelotagregationuseDB *LancelotAgregationUseDB) CopyBasicFieldsToLancelotAgregationUse_WOP(lancelotagregationuse *models.LancelotAgregationUse_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	lancelotagregationuse.Name = lancelotagregationuseDB.Name_Data.String
}

// CopyBasicFieldsToLancelotAgregationUseWOP
func (lancelotagregationuseDB *LancelotAgregationUseDB) CopyBasicFieldsToLancelotAgregationUseWOP(lancelotagregationuse *LancelotAgregationUseWOP) {
	lancelotagregationuse.ID = int(lancelotagregationuseDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	lancelotagregationuse.Name = lancelotagregationuseDB.Name_Data.String
}

// Backup generates a json file from a slice of all LancelotAgregationUseDB instances in the backrepo
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LancelotAgregationUseDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LancelotAgregationUseDB, 0)
	for _, lancelotagregationuseDB := range backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB {
		forBackup = append(forBackup, lancelotagregationuseDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json LancelotAgregationUse ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json LancelotAgregationUse file", err.Error())
	}
}

// Backup generates a json file from a slice of all LancelotAgregationUseDB instances in the backrepo
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LancelotAgregationUseDB, 0)
	for _, lancelotagregationuseDB := range backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB {
		forBackup = append(forBackup, lancelotagregationuseDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("LancelotAgregationUse")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&LancelotAgregationUse_Fields, -1)
	for _, lancelotagregationuseDB := range forBackup {

		var lancelotagregationuseWOP LancelotAgregationUseWOP
		lancelotagregationuseDB.CopyBasicFieldsToLancelotAgregationUseWOP(&lancelotagregationuseWOP)

		row := sh.AddRow()
		row.WriteStruct(&lancelotagregationuseWOP, -1)
	}
}

// RestoreXL from the "LancelotAgregationUse" sheet all LancelotAgregationUseDB instances
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLancelotAgregationUseid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["LancelotAgregationUse"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLancelotAgregationUse.rowVisitorLancelotAgregationUse)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) rowVisitorLancelotAgregationUse(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var lancelotagregationuseWOP LancelotAgregationUseWOP
		row.ReadStruct(&lancelotagregationuseWOP)

		// add the unmarshalled struct to the stage
		lancelotagregationuseDB := new(LancelotAgregationUseDB)
		lancelotagregationuseDB.CopyBasicFieldsFromLancelotAgregationUseWOP(&lancelotagregationuseWOP)

		lancelotagregationuseDB_ID_atBackupTime := lancelotagregationuseDB.ID
		lancelotagregationuseDB.ID = 0
		query := backRepoLancelotAgregationUse.db.Create(lancelotagregationuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB[lancelotagregationuseDB.ID] = lancelotagregationuseDB
		BackRepoLancelotAgregationUseid_atBckpTime_newID[lancelotagregationuseDB_ID_atBackupTime] = lancelotagregationuseDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LancelotAgregationUseDB.json" in dirPath that stores an array
// of LancelotAgregationUseDB and stores it in the database
// the map BackRepoLancelotAgregationUseid_atBckpTime_newID is updated accordingly
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLancelotAgregationUseid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LancelotAgregationUseDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json LancelotAgregationUse file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LancelotAgregationUseDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LancelotAgregationUseDBID_LancelotAgregationUseDB
	for _, lancelotagregationuseDB := range forRestore {

		lancelotagregationuseDB_ID_atBackupTime := lancelotagregationuseDB.ID
		lancelotagregationuseDB.ID = 0
		query := backRepoLancelotAgregationUse.db.Create(lancelotagregationuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB[lancelotagregationuseDB.ID] = lancelotagregationuseDB
		BackRepoLancelotAgregationUseid_atBckpTime_newID[lancelotagregationuseDB_ID_atBackupTime] = lancelotagregationuseDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json LancelotAgregationUse file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<LancelotAgregationUse>id_atBckpTime_newID
// to compute new index
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) RestorePhaseTwo() {

	for _, lancelotagregationuseDB := range backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB {

		// next line of code is to avert unused variable compilation error
		_ = lancelotagregationuseDB

		// insertion point for reindexing pointers encoding
		// reindexing ParameterAgregation field
		if lancelotagregationuseDB.ParameterAgregationID.Int64 != 0 {
			lancelotagregationuseDB.ParameterAgregationID.Int64 = int64(BackRepoLancelotAgregationid_atBckpTime_newID[uint(lancelotagregationuseDB.ParameterAgregationID.Int64)])
			lancelotagregationuseDB.ParameterAgregationID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoLancelotAgregationUse.db.Model(lancelotagregationuseDB).Updates(*lancelotagregationuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoLancelotAgregationUse.ResetReversePointers commits all staged instances of LancelotAgregationUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, lancelotagregationuse := range backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr {
		backRepoLancelotAgregationUse.ResetReversePointersInstance(backRepo, idx, lancelotagregationuse)
	}

	return
}

func (backRepoLancelotAgregationUse *BackRepoLancelotAgregationUseStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, lancelotagregationuse *models.LancelotAgregationUse) (Error error) {

	// fetch matching lancelotagregationuseDB
	if lancelotagregationuseDB, ok := backRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUseDB[idx]; ok {
		_ = lancelotagregationuseDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLancelotAgregationUseid_atBckpTime_newID map[uint]uint