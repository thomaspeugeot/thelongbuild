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
var dummy_Lancelot_sql sql.NullBool
var dummy_Lancelot_time time.Duration
var dummy_Lancelot_sort sort.Float64Slice

// LancelotAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model lancelotAPI
type LancelotAPI struct {
	gorm.Model

	models.Lancelot_WOP

	// encoding of pointers
	LancelotPointersEncoding LancelotPointersEncoding
}

// LancelotPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LancelotPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field GroupUse is a slice of pointers to another Struct (optional or 0..1)
	GroupUse IntSlice `gorm:"type:TEXT"`

	// field DocumentUse is a slice of pointers to another Struct (optional or 0..1)
	DocumentUse IntSlice `gorm:"type:TEXT"`

	// field GeoObjectUse is a slice of pointers to another Struct (optional or 0..1)
	GeoObjectUse IntSlice `gorm:"type:TEXT"`
}

// LancelotDB describes a lancelot in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model lancelotDB
type LancelotDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field lancelotDB.Name
	Name_Data sql.NullString

	// Declation for basic field lancelotDB.Tag
	Tag_Data sql.NullString

	// Declation for basic field lancelotDB.Description
	Description_Data sql.NullString
	// encoding of pointers
	LancelotPointersEncoding
}

// LancelotDBs arrays lancelotDBs
// swagger:response lancelotDBsResponse
type LancelotDBs []LancelotDB

// LancelotDBResponse provides response
// swagger:response lancelotDBResponse
type LancelotDBResponse struct {
	LancelotDB
}

// LancelotWOP is a Lancelot without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LancelotWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Tag string `xlsx:"2"`

	Description string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var Lancelot_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Tag",
	"Description",
}

type BackRepoLancelotStruct struct {
	// stores LancelotDB according to their gorm ID
	Map_LancelotDBID_LancelotDB map[uint]*LancelotDB

	// stores LancelotDB ID according to Lancelot address
	Map_LancelotPtr_LancelotDBID map[*models.Lancelot]uint

	// stores Lancelot according to their gorm ID
	Map_LancelotDBID_LancelotPtr map[uint]*models.Lancelot

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoLancelot *BackRepoLancelotStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoLancelot.stage
	return
}

func (backRepoLancelot *BackRepoLancelotStruct) GetDB() *gorm.DB {
	return backRepoLancelot.db
}

// GetLancelotDBFromLancelotPtr is a handy function to access the back repo instance from the stage instance
func (backRepoLancelot *BackRepoLancelotStruct) GetLancelotDBFromLancelotPtr(lancelot *models.Lancelot) (lancelotDB *LancelotDB) {
	id := backRepoLancelot.Map_LancelotPtr_LancelotDBID[lancelot]
	lancelotDB = backRepoLancelot.Map_LancelotDBID_LancelotDB[id]
	return
}

// BackRepoLancelot.CommitPhaseOne commits all staged instances of Lancelot to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLancelot *BackRepoLancelotStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for lancelot := range stage.Lancelots {
		backRepoLancelot.CommitPhaseOneInstance(lancelot)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, lancelot := range backRepoLancelot.Map_LancelotDBID_LancelotPtr {
		if _, ok := stage.Lancelots[lancelot]; !ok {
			backRepoLancelot.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLancelot.CommitDeleteInstance commits deletion of Lancelot to the BackRepo
func (backRepoLancelot *BackRepoLancelotStruct) CommitDeleteInstance(id uint) (Error error) {

	lancelot := backRepoLancelot.Map_LancelotDBID_LancelotPtr[id]

	// lancelot is not staged anymore, remove lancelotDB
	lancelotDB := backRepoLancelot.Map_LancelotDBID_LancelotDB[id]
	query := backRepoLancelot.db.Unscoped().Delete(&lancelotDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoLancelot.Map_LancelotPtr_LancelotDBID, lancelot)
	delete(backRepoLancelot.Map_LancelotDBID_LancelotPtr, id)
	delete(backRepoLancelot.Map_LancelotDBID_LancelotDB, id)

	return
}

// BackRepoLancelot.CommitPhaseOneInstance commits lancelot staged instances of Lancelot to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLancelot *BackRepoLancelotStruct) CommitPhaseOneInstance(lancelot *models.Lancelot) (Error error) {

	// check if the lancelot is not commited yet
	if _, ok := backRepoLancelot.Map_LancelotPtr_LancelotDBID[lancelot]; ok {
		return
	}

	// initiate lancelot
	var lancelotDB LancelotDB
	lancelotDB.CopyBasicFieldsFromLancelot(lancelot)

	query := backRepoLancelot.db.Create(&lancelotDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoLancelot.Map_LancelotPtr_LancelotDBID[lancelot] = lancelotDB.ID
	backRepoLancelot.Map_LancelotDBID_LancelotPtr[lancelotDB.ID] = lancelot
	backRepoLancelot.Map_LancelotDBID_LancelotDB[lancelotDB.ID] = &lancelotDB

	return
}

// BackRepoLancelot.CommitPhaseTwo commits all staged instances of Lancelot to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelot *BackRepoLancelotStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, lancelot := range backRepoLancelot.Map_LancelotDBID_LancelotPtr {
		backRepoLancelot.CommitPhaseTwoInstance(backRepo, idx, lancelot)
	}

	return
}

// BackRepoLancelot.CommitPhaseTwoInstance commits {{structname }} of models.Lancelot to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelot *BackRepoLancelotStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, lancelot *models.Lancelot) (Error error) {

	// fetch matching lancelotDB
	if lancelotDB, ok := backRepoLancelot.Map_LancelotDBID_LancelotDB[idx]; ok {

		lancelotDB.CopyBasicFieldsFromLancelot(lancelot)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		lancelotDB.LancelotPointersEncoding.GroupUse = make([]int, 0)
		// 2. encode
		for _, groupuseAssocEnd := range lancelot.GroupUse {
			groupuseAssocEnd_DB :=
				backRepo.BackRepoGroupUse.GetGroupUseDBFromGroupUsePtr(groupuseAssocEnd)

			// the stage might be inconsistant, meaning that the groupuseAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if groupuseAssocEnd_DB == nil {
				continue
			}

			lancelotDB.LancelotPointersEncoding.GroupUse =
				append(lancelotDB.LancelotPointersEncoding.GroupUse, int(groupuseAssocEnd_DB.ID))
		}

		// 1. reset
		lancelotDB.LancelotPointersEncoding.DocumentUse = make([]int, 0)
		// 2. encode
		for _, documentuseAssocEnd := range lancelot.DocumentUse {
			documentuseAssocEnd_DB :=
				backRepo.BackRepoDocumentUse.GetDocumentUseDBFromDocumentUsePtr(documentuseAssocEnd)

			// the stage might be inconsistant, meaning that the documentuseAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if documentuseAssocEnd_DB == nil {
				continue
			}

			lancelotDB.LancelotPointersEncoding.DocumentUse =
				append(lancelotDB.LancelotPointersEncoding.DocumentUse, int(documentuseAssocEnd_DB.ID))
		}

		// 1. reset
		lancelotDB.LancelotPointersEncoding.GeoObjectUse = make([]int, 0)
		// 2. encode
		for _, geoobjectuseAssocEnd := range lancelot.GeoObjectUse {
			geoobjectuseAssocEnd_DB :=
				backRepo.BackRepoGeoObjectUse.GetGeoObjectUseDBFromGeoObjectUsePtr(geoobjectuseAssocEnd)

			// the stage might be inconsistant, meaning that the geoobjectuseAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if geoobjectuseAssocEnd_DB == nil {
				continue
			}

			lancelotDB.LancelotPointersEncoding.GeoObjectUse =
				append(lancelotDB.LancelotPointersEncoding.GeoObjectUse, int(geoobjectuseAssocEnd_DB.ID))
		}

		query := backRepoLancelot.db.Save(&lancelotDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Lancelot intance %s", lancelot.Name))
		return err
	}

	return
}

// BackRepoLancelot.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoLancelot *BackRepoLancelotStruct) CheckoutPhaseOne() (Error error) {

	lancelotDBArray := make([]LancelotDB, 0)
	query := backRepoLancelot.db.Find(&lancelotDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	lancelotInstancesToBeRemovedFromTheStage := make(map[*models.Lancelot]any)
	for key, value := range backRepoLancelot.stage.Lancelots {
		lancelotInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, lancelotDB := range lancelotDBArray {
		backRepoLancelot.CheckoutPhaseOneInstance(&lancelotDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		lancelot, ok := backRepoLancelot.Map_LancelotDBID_LancelotPtr[lancelotDB.ID]
		if ok {
			delete(lancelotInstancesToBeRemovedFromTheStage, lancelot)
		}
	}

	// remove from stage and back repo's 3 maps all lancelots that are not in the checkout
	for lancelot := range lancelotInstancesToBeRemovedFromTheStage {
		lancelot.Unstage(backRepoLancelot.GetStage())

		// remove instance from the back repo 3 maps
		lancelotID := backRepoLancelot.Map_LancelotPtr_LancelotDBID[lancelot]
		delete(backRepoLancelot.Map_LancelotPtr_LancelotDBID, lancelot)
		delete(backRepoLancelot.Map_LancelotDBID_LancelotDB, lancelotID)
		delete(backRepoLancelot.Map_LancelotDBID_LancelotPtr, lancelotID)
	}

	return
}

// CheckoutPhaseOneInstance takes a lancelotDB that has been found in the DB, updates the backRepo and stages the
// models version of the lancelotDB
func (backRepoLancelot *BackRepoLancelotStruct) CheckoutPhaseOneInstance(lancelotDB *LancelotDB) (Error error) {

	lancelot, ok := backRepoLancelot.Map_LancelotDBID_LancelotPtr[lancelotDB.ID]
	if !ok {
		lancelot = new(models.Lancelot)

		backRepoLancelot.Map_LancelotDBID_LancelotPtr[lancelotDB.ID] = lancelot
		backRepoLancelot.Map_LancelotPtr_LancelotDBID[lancelot] = lancelotDB.ID

		// append model store with the new element
		lancelot.Name = lancelotDB.Name_Data.String
		lancelot.Stage(backRepoLancelot.GetStage())
	}
	lancelotDB.CopyBasicFieldsToLancelot(lancelot)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	lancelot.Stage(backRepoLancelot.GetStage())

	// preserve pointer to lancelotDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LancelotDBID_LancelotDB)[lancelotDB hold variable pointers
	lancelotDB_Data := *lancelotDB
	preservedPtrToLancelot := &lancelotDB_Data
	backRepoLancelot.Map_LancelotDBID_LancelotDB[lancelotDB.ID] = preservedPtrToLancelot

	return
}

// BackRepoLancelot.CheckoutPhaseTwo Checkouts all staged instances of Lancelot to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelot *BackRepoLancelotStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, lancelotDB := range backRepoLancelot.Map_LancelotDBID_LancelotDB {
		backRepoLancelot.CheckoutPhaseTwoInstance(backRepo, lancelotDB)
	}
	return
}

// BackRepoLancelot.CheckoutPhaseTwoInstance Checkouts staged instances of Lancelot to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelot *BackRepoLancelotStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, lancelotDB *LancelotDB) (Error error) {

	lancelot := backRepoLancelot.Map_LancelotDBID_LancelotPtr[lancelotDB.ID]

	lancelotDB.DecodePointers(backRepo, lancelot)

	return
}

func (lancelotDB *LancelotDB) DecodePointers(backRepo *BackRepoStruct, lancelot *models.Lancelot) {

	// insertion point for checkout of pointer encoding
	// This loop redeem lancelot.GroupUse in the stage from the encode in the back repo
	// It parses all GroupUseDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	lancelot.GroupUse = lancelot.GroupUse[:0]
	for _, _GroupUseid := range lancelotDB.LancelotPointersEncoding.GroupUse {
		lancelot.GroupUse = append(lancelot.GroupUse, backRepo.BackRepoGroupUse.Map_GroupUseDBID_GroupUsePtr[uint(_GroupUseid)])
	}

	// This loop redeem lancelot.DocumentUse in the stage from the encode in the back repo
	// It parses all DocumentUseDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	lancelot.DocumentUse = lancelot.DocumentUse[:0]
	for _, _DocumentUseid := range lancelotDB.LancelotPointersEncoding.DocumentUse {
		lancelot.DocumentUse = append(lancelot.DocumentUse, backRepo.BackRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[uint(_DocumentUseid)])
	}

	// This loop redeem lancelot.GeoObjectUse in the stage from the encode in the back repo
	// It parses all GeoObjectUseDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	lancelot.GeoObjectUse = lancelot.GeoObjectUse[:0]
	for _, _GeoObjectUseid := range lancelotDB.LancelotPointersEncoding.GeoObjectUse {
		lancelot.GeoObjectUse = append(lancelot.GeoObjectUse, backRepo.BackRepoGeoObjectUse.Map_GeoObjectUseDBID_GeoObjectUsePtr[uint(_GeoObjectUseid)])
	}

	return
}

// CommitLancelot allows commit of a single lancelot (if already staged)
func (backRepo *BackRepoStruct) CommitLancelot(lancelot *models.Lancelot) {
	backRepo.BackRepoLancelot.CommitPhaseOneInstance(lancelot)
	if id, ok := backRepo.BackRepoLancelot.Map_LancelotPtr_LancelotDBID[lancelot]; ok {
		backRepo.BackRepoLancelot.CommitPhaseTwoInstance(backRepo, id, lancelot)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitLancelot allows checkout of a single lancelot (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLancelot(lancelot *models.Lancelot) {
	// check if the lancelot is staged
	if _, ok := backRepo.BackRepoLancelot.Map_LancelotPtr_LancelotDBID[lancelot]; ok {

		if id, ok := backRepo.BackRepoLancelot.Map_LancelotPtr_LancelotDBID[lancelot]; ok {
			var lancelotDB LancelotDB
			lancelotDB.ID = id

			if err := backRepo.BackRepoLancelot.db.First(&lancelotDB, id).Error; err != nil {
				log.Fatalln("CheckoutLancelot : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLancelot.CheckoutPhaseOneInstance(&lancelotDB)
			backRepo.BackRepoLancelot.CheckoutPhaseTwoInstance(backRepo, &lancelotDB)
		}
	}
}

// CopyBasicFieldsFromLancelot
func (lancelotDB *LancelotDB) CopyBasicFieldsFromLancelot(lancelot *models.Lancelot) {
	// insertion point for fields commit

	lancelotDB.Name_Data.String = lancelot.Name
	lancelotDB.Name_Data.Valid = true

	lancelotDB.Tag_Data.String = lancelot.Tag
	lancelotDB.Tag_Data.Valid = true

	lancelotDB.Description_Data.String = lancelot.Description
	lancelotDB.Description_Data.Valid = true
}

// CopyBasicFieldsFromLancelot_WOP
func (lancelotDB *LancelotDB) CopyBasicFieldsFromLancelot_WOP(lancelot *models.Lancelot_WOP) {
	// insertion point for fields commit

	lancelotDB.Name_Data.String = lancelot.Name
	lancelotDB.Name_Data.Valid = true

	lancelotDB.Tag_Data.String = lancelot.Tag
	lancelotDB.Tag_Data.Valid = true

	lancelotDB.Description_Data.String = lancelot.Description
	lancelotDB.Description_Data.Valid = true
}

// CopyBasicFieldsFromLancelotWOP
func (lancelotDB *LancelotDB) CopyBasicFieldsFromLancelotWOP(lancelot *LancelotWOP) {
	// insertion point for fields commit

	lancelotDB.Name_Data.String = lancelot.Name
	lancelotDB.Name_Data.Valid = true

	lancelotDB.Tag_Data.String = lancelot.Tag
	lancelotDB.Tag_Data.Valid = true

	lancelotDB.Description_Data.String = lancelot.Description
	lancelotDB.Description_Data.Valid = true
}

// CopyBasicFieldsToLancelot
func (lancelotDB *LancelotDB) CopyBasicFieldsToLancelot(lancelot *models.Lancelot) {
	// insertion point for checkout of basic fields (back repo to stage)
	lancelot.Name = lancelotDB.Name_Data.String
	lancelot.Tag = lancelotDB.Tag_Data.String
	lancelot.Description = lancelotDB.Description_Data.String
}

// CopyBasicFieldsToLancelot_WOP
func (lancelotDB *LancelotDB) CopyBasicFieldsToLancelot_WOP(lancelot *models.Lancelot_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	lancelot.Name = lancelotDB.Name_Data.String
	lancelot.Tag = lancelotDB.Tag_Data.String
	lancelot.Description = lancelotDB.Description_Data.String
}

// CopyBasicFieldsToLancelotWOP
func (lancelotDB *LancelotDB) CopyBasicFieldsToLancelotWOP(lancelot *LancelotWOP) {
	lancelot.ID = int(lancelotDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	lancelot.Name = lancelotDB.Name_Data.String
	lancelot.Tag = lancelotDB.Tag_Data.String
	lancelot.Description = lancelotDB.Description_Data.String
}

// Backup generates a json file from a slice of all LancelotDB instances in the backrepo
func (backRepoLancelot *BackRepoLancelotStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LancelotDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LancelotDB, 0)
	for _, lancelotDB := range backRepoLancelot.Map_LancelotDBID_LancelotDB {
		forBackup = append(forBackup, lancelotDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Lancelot ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Lancelot file", err.Error())
	}
}

// Backup generates a json file from a slice of all LancelotDB instances in the backrepo
func (backRepoLancelot *BackRepoLancelotStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LancelotDB, 0)
	for _, lancelotDB := range backRepoLancelot.Map_LancelotDBID_LancelotDB {
		forBackup = append(forBackup, lancelotDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Lancelot")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Lancelot_Fields, -1)
	for _, lancelotDB := range forBackup {

		var lancelotWOP LancelotWOP
		lancelotDB.CopyBasicFieldsToLancelotWOP(&lancelotWOP)

		row := sh.AddRow()
		row.WriteStruct(&lancelotWOP, -1)
	}
}

// RestoreXL from the "Lancelot" sheet all LancelotDB instances
func (backRepoLancelot *BackRepoLancelotStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLancelotid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Lancelot"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLancelot.rowVisitorLancelot)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoLancelot *BackRepoLancelotStruct) rowVisitorLancelot(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var lancelotWOP LancelotWOP
		row.ReadStruct(&lancelotWOP)

		// add the unmarshalled struct to the stage
		lancelotDB := new(LancelotDB)
		lancelotDB.CopyBasicFieldsFromLancelotWOP(&lancelotWOP)

		lancelotDB_ID_atBackupTime := lancelotDB.ID
		lancelotDB.ID = 0
		query := backRepoLancelot.db.Create(lancelotDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoLancelot.Map_LancelotDBID_LancelotDB[lancelotDB.ID] = lancelotDB
		BackRepoLancelotid_atBckpTime_newID[lancelotDB_ID_atBackupTime] = lancelotDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LancelotDB.json" in dirPath that stores an array
// of LancelotDB and stores it in the database
// the map BackRepoLancelotid_atBckpTime_newID is updated accordingly
func (backRepoLancelot *BackRepoLancelotStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLancelotid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LancelotDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Lancelot file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LancelotDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LancelotDBID_LancelotDB
	for _, lancelotDB := range forRestore {

		lancelotDB_ID_atBackupTime := lancelotDB.ID
		lancelotDB.ID = 0
		query := backRepoLancelot.db.Create(lancelotDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoLancelot.Map_LancelotDBID_LancelotDB[lancelotDB.ID] = lancelotDB
		BackRepoLancelotid_atBckpTime_newID[lancelotDB_ID_atBackupTime] = lancelotDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Lancelot file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Lancelot>id_atBckpTime_newID
// to compute new index
func (backRepoLancelot *BackRepoLancelotStruct) RestorePhaseTwo() {

	for _, lancelotDB := range backRepoLancelot.Map_LancelotDBID_LancelotDB {

		// next line of code is to avert unused variable compilation error
		_ = lancelotDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoLancelot.db.Model(lancelotDB).Updates(*lancelotDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoLancelot.ResetReversePointers commits all staged instances of Lancelot to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLancelot *BackRepoLancelotStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, lancelot := range backRepoLancelot.Map_LancelotDBID_LancelotPtr {
		backRepoLancelot.ResetReversePointersInstance(backRepo, idx, lancelot)
	}

	return
}

func (backRepoLancelot *BackRepoLancelotStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, lancelot *models.Lancelot) (Error error) {

	// fetch matching lancelotDB
	if lancelotDB, ok := backRepoLancelot.Map_LancelotDBID_LancelotDB[idx]; ok {
		_ = lancelotDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLancelotid_atBckpTime_newID map[uint]uint
