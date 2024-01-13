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
var dummy_MapObjectUse_sql sql.NullBool
var dummy_MapObjectUse_time time.Duration
var dummy_MapObjectUse_sort sort.Float64Slice

// MapObjectUseAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model mapobjectuseAPI
type MapObjectUseAPI struct {
	gorm.Model

	models.MapObjectUse_WOP

	// encoding of pointers
	MapObjectUsePointersEncoding MapObjectUsePointersEncoding
}

// MapObjectUsePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type MapObjectUsePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Map is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	MapID sql.NullInt64
}

// MapObjectUseDB describes a mapobjectuse in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model mapobjectuseDB
type MapObjectUseDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field mapobjectuseDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	MapObjectUsePointersEncoding
}

// MapObjectUseDBs arrays mapobjectuseDBs
// swagger:response mapobjectuseDBsResponse
type MapObjectUseDBs []MapObjectUseDB

// MapObjectUseDBResponse provides response
// swagger:response mapobjectuseDBResponse
type MapObjectUseDBResponse struct {
	MapObjectUseDB
}

// MapObjectUseWOP is a MapObjectUse without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type MapObjectUseWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var MapObjectUse_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoMapObjectUseStruct struct {
	// stores MapObjectUseDB according to their gorm ID
	Map_MapObjectUseDBID_MapObjectUseDB map[uint]*MapObjectUseDB

	// stores MapObjectUseDB ID according to MapObjectUse address
	Map_MapObjectUsePtr_MapObjectUseDBID map[*models.MapObjectUse]uint

	// stores MapObjectUse according to their gorm ID
	Map_MapObjectUseDBID_MapObjectUsePtr map[uint]*models.MapObjectUse

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoMapObjectUse.stage
	return
}

func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) GetDB() *gorm.DB {
	return backRepoMapObjectUse.db
}

// GetMapObjectUseDBFromMapObjectUsePtr is a handy function to access the back repo instance from the stage instance
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) GetMapObjectUseDBFromMapObjectUsePtr(mapobjectuse *models.MapObjectUse) (mapobjectuseDB *MapObjectUseDB) {
	id := backRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID[mapobjectuse]
	mapobjectuseDB = backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB[id]
	return
}

// BackRepoMapObjectUse.CommitPhaseOne commits all staged instances of MapObjectUse to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for mapobjectuse := range stage.MapObjectUses {
		backRepoMapObjectUse.CommitPhaseOneInstance(mapobjectuse)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, mapobjectuse := range backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr {
		if _, ok := stage.MapObjectUses[mapobjectuse]; !ok {
			backRepoMapObjectUse.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoMapObjectUse.CommitDeleteInstance commits deletion of MapObjectUse to the BackRepo
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CommitDeleteInstance(id uint) (Error error) {

	mapobjectuse := backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[id]

	// mapobjectuse is not staged anymore, remove mapobjectuseDB
	mapobjectuseDB := backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB[id]
	query := backRepoMapObjectUse.db.Unscoped().Delete(&mapobjectuseDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID, mapobjectuse)
	delete(backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr, id)
	delete(backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB, id)

	return
}

// BackRepoMapObjectUse.CommitPhaseOneInstance commits mapobjectuse staged instances of MapObjectUse to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CommitPhaseOneInstance(mapobjectuse *models.MapObjectUse) (Error error) {

	// check if the mapobjectuse is not commited yet
	if _, ok := backRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID[mapobjectuse]; ok {
		return
	}

	// initiate mapobjectuse
	var mapobjectuseDB MapObjectUseDB
	mapobjectuseDB.CopyBasicFieldsFromMapObjectUse(mapobjectuse)

	query := backRepoMapObjectUse.db.Create(&mapobjectuseDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID[mapobjectuse] = mapobjectuseDB.ID
	backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[mapobjectuseDB.ID] = mapobjectuse
	backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB[mapobjectuseDB.ID] = &mapobjectuseDB

	return
}

// BackRepoMapObjectUse.CommitPhaseTwo commits all staged instances of MapObjectUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, mapobjectuse := range backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr {
		backRepoMapObjectUse.CommitPhaseTwoInstance(backRepo, idx, mapobjectuse)
	}

	return
}

// BackRepoMapObjectUse.CommitPhaseTwoInstance commits {{structname }} of models.MapObjectUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, mapobjectuse *models.MapObjectUse) (Error error) {

	// fetch matching mapobjectuseDB
	if mapobjectuseDB, ok := backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB[idx]; ok {

		mapobjectuseDB.CopyBasicFieldsFromMapObjectUse(mapobjectuse)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value mapobjectuse.Map translates to updating the mapobjectuse.MapID
		mapobjectuseDB.MapID.Valid = true // allow for a 0 value (nil association)
		if mapobjectuse.Map != nil {
			if MapId, ok := backRepo.BackRepoMapObject.Map_MapObjectPtr_MapObjectDBID[mapobjectuse.Map]; ok {
				mapobjectuseDB.MapID.Int64 = int64(MapId)
				mapobjectuseDB.MapID.Valid = true
			}
		} else {
			mapobjectuseDB.MapID.Int64 = 0
			mapobjectuseDB.MapID.Valid = true
		}

		query := backRepoMapObjectUse.db.Save(&mapobjectuseDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown MapObjectUse intance %s", mapobjectuse.Name))
		return err
	}

	return
}

// BackRepoMapObjectUse.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CheckoutPhaseOne() (Error error) {

	mapobjectuseDBArray := make([]MapObjectUseDB, 0)
	query := backRepoMapObjectUse.db.Find(&mapobjectuseDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	mapobjectuseInstancesToBeRemovedFromTheStage := make(map[*models.MapObjectUse]any)
	for key, value := range backRepoMapObjectUse.stage.MapObjectUses {
		mapobjectuseInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, mapobjectuseDB := range mapobjectuseDBArray {
		backRepoMapObjectUse.CheckoutPhaseOneInstance(&mapobjectuseDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		mapobjectuse, ok := backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[mapobjectuseDB.ID]
		if ok {
			delete(mapobjectuseInstancesToBeRemovedFromTheStage, mapobjectuse)
		}
	}

	// remove from stage and back repo's 3 maps all mapobjectuses that are not in the checkout
	for mapobjectuse := range mapobjectuseInstancesToBeRemovedFromTheStage {
		mapobjectuse.Unstage(backRepoMapObjectUse.GetStage())

		// remove instance from the back repo 3 maps
		mapobjectuseID := backRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID[mapobjectuse]
		delete(backRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID, mapobjectuse)
		delete(backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB, mapobjectuseID)
		delete(backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr, mapobjectuseID)
	}

	return
}

// CheckoutPhaseOneInstance takes a mapobjectuseDB that has been found in the DB, updates the backRepo and stages the
// models version of the mapobjectuseDB
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CheckoutPhaseOneInstance(mapobjectuseDB *MapObjectUseDB) (Error error) {

	mapobjectuse, ok := backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[mapobjectuseDB.ID]
	if !ok {
		mapobjectuse = new(models.MapObjectUse)

		backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[mapobjectuseDB.ID] = mapobjectuse
		backRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID[mapobjectuse] = mapobjectuseDB.ID

		// append model store with the new element
		mapobjectuse.Name = mapobjectuseDB.Name_Data.String
		mapobjectuse.Stage(backRepoMapObjectUse.GetStage())
	}
	mapobjectuseDB.CopyBasicFieldsToMapObjectUse(mapobjectuse)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	mapobjectuse.Stage(backRepoMapObjectUse.GetStage())

	// preserve pointer to mapobjectuseDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_MapObjectUseDBID_MapObjectUseDB)[mapobjectuseDB hold variable pointers
	mapobjectuseDB_Data := *mapobjectuseDB
	preservedPtrToMapObjectUse := &mapobjectuseDB_Data
	backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB[mapobjectuseDB.ID] = preservedPtrToMapObjectUse

	return
}

// BackRepoMapObjectUse.CheckoutPhaseTwo Checkouts all staged instances of MapObjectUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, mapobjectuseDB := range backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB {
		backRepoMapObjectUse.CheckoutPhaseTwoInstance(backRepo, mapobjectuseDB)
	}
	return
}

// BackRepoMapObjectUse.CheckoutPhaseTwoInstance Checkouts staged instances of MapObjectUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, mapobjectuseDB *MapObjectUseDB) (Error error) {

	mapobjectuse := backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[mapobjectuseDB.ID]

	mapobjectuseDB.DecodePointers(backRepo, mapobjectuse)

	return
}

func (mapobjectuseDB *MapObjectUseDB) DecodePointers(backRepo *BackRepoStruct, mapobjectuse *models.MapObjectUse) {

	// insertion point for checkout of pointer encoding
	// Map field
	mapobjectuse.Map = nil
	if mapobjectuseDB.MapID.Int64 != 0 {
		mapobjectuse.Map = backRepo.BackRepoMapObject.Map_MapObjectDBID_MapObjectPtr[uint(mapobjectuseDB.MapID.Int64)]
	}
	return
}

// CommitMapObjectUse allows commit of a single mapobjectuse (if already staged)
func (backRepo *BackRepoStruct) CommitMapObjectUse(mapobjectuse *models.MapObjectUse) {
	backRepo.BackRepoMapObjectUse.CommitPhaseOneInstance(mapobjectuse)
	if id, ok := backRepo.BackRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID[mapobjectuse]; ok {
		backRepo.BackRepoMapObjectUse.CommitPhaseTwoInstance(backRepo, id, mapobjectuse)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitMapObjectUse allows checkout of a single mapobjectuse (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutMapObjectUse(mapobjectuse *models.MapObjectUse) {
	// check if the mapobjectuse is staged
	if _, ok := backRepo.BackRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID[mapobjectuse]; ok {

		if id, ok := backRepo.BackRepoMapObjectUse.Map_MapObjectUsePtr_MapObjectUseDBID[mapobjectuse]; ok {
			var mapobjectuseDB MapObjectUseDB
			mapobjectuseDB.ID = id

			if err := backRepo.BackRepoMapObjectUse.db.First(&mapobjectuseDB, id).Error; err != nil {
				log.Fatalln("CheckoutMapObjectUse : Problem with getting object with id:", id)
			}
			backRepo.BackRepoMapObjectUse.CheckoutPhaseOneInstance(&mapobjectuseDB)
			backRepo.BackRepoMapObjectUse.CheckoutPhaseTwoInstance(backRepo, &mapobjectuseDB)
		}
	}
}

// CopyBasicFieldsFromMapObjectUse
func (mapobjectuseDB *MapObjectUseDB) CopyBasicFieldsFromMapObjectUse(mapobjectuse *models.MapObjectUse) {
	// insertion point for fields commit

	mapobjectuseDB.Name_Data.String = mapobjectuse.Name
	mapobjectuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromMapObjectUse_WOP
func (mapobjectuseDB *MapObjectUseDB) CopyBasicFieldsFromMapObjectUse_WOP(mapobjectuse *models.MapObjectUse_WOP) {
	// insertion point for fields commit

	mapobjectuseDB.Name_Data.String = mapobjectuse.Name
	mapobjectuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromMapObjectUseWOP
func (mapobjectuseDB *MapObjectUseDB) CopyBasicFieldsFromMapObjectUseWOP(mapobjectuse *MapObjectUseWOP) {
	// insertion point for fields commit

	mapobjectuseDB.Name_Data.String = mapobjectuse.Name
	mapobjectuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsToMapObjectUse
func (mapobjectuseDB *MapObjectUseDB) CopyBasicFieldsToMapObjectUse(mapobjectuse *models.MapObjectUse) {
	// insertion point for checkout of basic fields (back repo to stage)
	mapobjectuse.Name = mapobjectuseDB.Name_Data.String
}

// CopyBasicFieldsToMapObjectUse_WOP
func (mapobjectuseDB *MapObjectUseDB) CopyBasicFieldsToMapObjectUse_WOP(mapobjectuse *models.MapObjectUse_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	mapobjectuse.Name = mapobjectuseDB.Name_Data.String
}

// CopyBasicFieldsToMapObjectUseWOP
func (mapobjectuseDB *MapObjectUseDB) CopyBasicFieldsToMapObjectUseWOP(mapobjectuse *MapObjectUseWOP) {
	mapobjectuse.ID = int(mapobjectuseDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	mapobjectuse.Name = mapobjectuseDB.Name_Data.String
}

// Backup generates a json file from a slice of all MapObjectUseDB instances in the backrepo
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "MapObjectUseDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MapObjectUseDB, 0)
	for _, mapobjectuseDB := range backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB {
		forBackup = append(forBackup, mapobjectuseDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json MapObjectUse ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json MapObjectUse file", err.Error())
	}
}

// Backup generates a json file from a slice of all MapObjectUseDB instances in the backrepo
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MapObjectUseDB, 0)
	for _, mapobjectuseDB := range backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB {
		forBackup = append(forBackup, mapobjectuseDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("MapObjectUse")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&MapObjectUse_Fields, -1)
	for _, mapobjectuseDB := range forBackup {

		var mapobjectuseWOP MapObjectUseWOP
		mapobjectuseDB.CopyBasicFieldsToMapObjectUseWOP(&mapobjectuseWOP)

		row := sh.AddRow()
		row.WriteStruct(&mapobjectuseWOP, -1)
	}
}

// RestoreXL from the "MapObjectUse" sheet all MapObjectUseDB instances
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoMapObjectUseid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["MapObjectUse"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoMapObjectUse.rowVisitorMapObjectUse)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) rowVisitorMapObjectUse(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var mapobjectuseWOP MapObjectUseWOP
		row.ReadStruct(&mapobjectuseWOP)

		// add the unmarshalled struct to the stage
		mapobjectuseDB := new(MapObjectUseDB)
		mapobjectuseDB.CopyBasicFieldsFromMapObjectUseWOP(&mapobjectuseWOP)

		mapobjectuseDB_ID_atBackupTime := mapobjectuseDB.ID
		mapobjectuseDB.ID = 0
		query := backRepoMapObjectUse.db.Create(mapobjectuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB[mapobjectuseDB.ID] = mapobjectuseDB
		BackRepoMapObjectUseid_atBckpTime_newID[mapobjectuseDB_ID_atBackupTime] = mapobjectuseDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "MapObjectUseDB.json" in dirPath that stores an array
// of MapObjectUseDB and stores it in the database
// the map BackRepoMapObjectUseid_atBckpTime_newID is updated accordingly
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoMapObjectUseid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "MapObjectUseDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json MapObjectUse file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*MapObjectUseDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_MapObjectUseDBID_MapObjectUseDB
	for _, mapobjectuseDB := range forRestore {

		mapobjectuseDB_ID_atBackupTime := mapobjectuseDB.ID
		mapobjectuseDB.ID = 0
		query := backRepoMapObjectUse.db.Create(mapobjectuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB[mapobjectuseDB.ID] = mapobjectuseDB
		BackRepoMapObjectUseid_atBckpTime_newID[mapobjectuseDB_ID_atBackupTime] = mapobjectuseDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json MapObjectUse file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<MapObjectUse>id_atBckpTime_newID
// to compute new index
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) RestorePhaseTwo() {

	for _, mapobjectuseDB := range backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB {

		// next line of code is to avert unused variable compilation error
		_ = mapobjectuseDB

		// insertion point for reindexing pointers encoding
		// reindexing Map field
		if mapobjectuseDB.MapID.Int64 != 0 {
			mapobjectuseDB.MapID.Int64 = int64(BackRepoMapObjectid_atBckpTime_newID[uint(mapobjectuseDB.MapID.Int64)])
			mapobjectuseDB.MapID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoMapObjectUse.db.Model(mapobjectuseDB).Updates(*mapobjectuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoMapObjectUse.ResetReversePointers commits all staged instances of MapObjectUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, mapobjectuse := range backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr {
		backRepoMapObjectUse.ResetReversePointersInstance(backRepo, idx, mapobjectuse)
	}

	return
}

func (backRepoMapObjectUse *BackRepoMapObjectUseStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, mapobjectuse *models.MapObjectUse) (Error error) {

	// fetch matching mapobjectuseDB
	if mapobjectuseDB, ok := backRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUseDB[idx]; ok {
		_ = mapobjectuseDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoMapObjectUseid_atBckpTime_newID map[uint]uint