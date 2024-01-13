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
var dummy_DocumentUse_sql sql.NullBool
var dummy_DocumentUse_time time.Duration
var dummy_DocumentUse_sort sort.Float64Slice

// DocumentUseAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model documentuseAPI
type DocumentUseAPI struct {
	gorm.Model

	models.DocumentUse_WOP

	// encoding of pointers
	DocumentUsePointersEncoding DocumentUsePointersEncoding
}

// DocumentUsePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type DocumentUsePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Document is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	DocumentID sql.NullInt64
}

// DocumentUseDB describes a documentuse in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model documentuseDB
type DocumentUseDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field documentuseDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	DocumentUsePointersEncoding
}

// DocumentUseDBs arrays documentuseDBs
// swagger:response documentuseDBsResponse
type DocumentUseDBs []DocumentUseDB

// DocumentUseDBResponse provides response
// swagger:response documentuseDBResponse
type DocumentUseDBResponse struct {
	DocumentUseDB
}

// DocumentUseWOP is a DocumentUse without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type DocumentUseWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var DocumentUse_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoDocumentUseStruct struct {
	// stores DocumentUseDB according to their gorm ID
	Map_DocumentUseDBID_DocumentUseDB map[uint]*DocumentUseDB

	// stores DocumentUseDB ID according to DocumentUse address
	Map_DocumentUsePtr_DocumentUseDBID map[*models.DocumentUse]uint

	// stores DocumentUse according to their gorm ID
	Map_DocumentUseDBID_DocumentUsePtr map[uint]*models.DocumentUse

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoDocumentUse *BackRepoDocumentUseStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoDocumentUse.stage
	return
}

func (backRepoDocumentUse *BackRepoDocumentUseStruct) GetDB() *gorm.DB {
	return backRepoDocumentUse.db
}

// GetDocumentUseDBFromDocumentUsePtr is a handy function to access the back repo instance from the stage instance
func (backRepoDocumentUse *BackRepoDocumentUseStruct) GetDocumentUseDBFromDocumentUsePtr(documentuse *models.DocumentUse) (documentuseDB *DocumentUseDB) {
	id := backRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID[documentuse]
	documentuseDB = backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB[id]
	return
}

// BackRepoDocumentUse.CommitPhaseOne commits all staged instances of DocumentUse to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for documentuse := range stage.DocumentUses {
		backRepoDocumentUse.CommitPhaseOneInstance(documentuse)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, documentuse := range backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr {
		if _, ok := stage.DocumentUses[documentuse]; !ok {
			backRepoDocumentUse.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoDocumentUse.CommitDeleteInstance commits deletion of DocumentUse to the BackRepo
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CommitDeleteInstance(id uint) (Error error) {

	documentuse := backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[id]

	// documentuse is not staged anymore, remove documentuseDB
	documentuseDB := backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB[id]
	query := backRepoDocumentUse.db.Unscoped().Delete(&documentuseDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID, documentuse)
	delete(backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr, id)
	delete(backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB, id)

	return
}

// BackRepoDocumentUse.CommitPhaseOneInstance commits documentuse staged instances of DocumentUse to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CommitPhaseOneInstance(documentuse *models.DocumentUse) (Error error) {

	// check if the documentuse is not commited yet
	if _, ok := backRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID[documentuse]; ok {
		return
	}

	// initiate documentuse
	var documentuseDB DocumentUseDB
	documentuseDB.CopyBasicFieldsFromDocumentUse(documentuse)

	query := backRepoDocumentUse.db.Create(&documentuseDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID[documentuse] = documentuseDB.ID
	backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[documentuseDB.ID] = documentuse
	backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB[documentuseDB.ID] = &documentuseDB

	return
}

// BackRepoDocumentUse.CommitPhaseTwo commits all staged instances of DocumentUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, documentuse := range backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr {
		backRepoDocumentUse.CommitPhaseTwoInstance(backRepo, idx, documentuse)
	}

	return
}

// BackRepoDocumentUse.CommitPhaseTwoInstance commits {{structname }} of models.DocumentUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, documentuse *models.DocumentUse) (Error error) {

	// fetch matching documentuseDB
	if documentuseDB, ok := backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB[idx]; ok {

		documentuseDB.CopyBasicFieldsFromDocumentUse(documentuse)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value documentuse.Document translates to updating the documentuse.DocumentID
		documentuseDB.DocumentID.Valid = true // allow for a 0 value (nil association)
		if documentuse.Document != nil {
			if DocumentId, ok := backRepo.BackRepoDocument.Map_DocumentPtr_DocumentDBID[documentuse.Document]; ok {
				documentuseDB.DocumentID.Int64 = int64(DocumentId)
				documentuseDB.DocumentID.Valid = true
			}
		} else {
			documentuseDB.DocumentID.Int64 = 0
			documentuseDB.DocumentID.Valid = true
		}

		query := backRepoDocumentUse.db.Save(&documentuseDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown DocumentUse intance %s", documentuse.Name))
		return err
	}

	return
}

// BackRepoDocumentUse.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CheckoutPhaseOne() (Error error) {

	documentuseDBArray := make([]DocumentUseDB, 0)
	query := backRepoDocumentUse.db.Find(&documentuseDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	documentuseInstancesToBeRemovedFromTheStage := make(map[*models.DocumentUse]any)
	for key, value := range backRepoDocumentUse.stage.DocumentUses {
		documentuseInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, documentuseDB := range documentuseDBArray {
		backRepoDocumentUse.CheckoutPhaseOneInstance(&documentuseDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		documentuse, ok := backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[documentuseDB.ID]
		if ok {
			delete(documentuseInstancesToBeRemovedFromTheStage, documentuse)
		}
	}

	// remove from stage and back repo's 3 maps all documentuses that are not in the checkout
	for documentuse := range documentuseInstancesToBeRemovedFromTheStage {
		documentuse.Unstage(backRepoDocumentUse.GetStage())

		// remove instance from the back repo 3 maps
		documentuseID := backRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID[documentuse]
		delete(backRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID, documentuse)
		delete(backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB, documentuseID)
		delete(backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr, documentuseID)
	}

	return
}

// CheckoutPhaseOneInstance takes a documentuseDB that has been found in the DB, updates the backRepo and stages the
// models version of the documentuseDB
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CheckoutPhaseOneInstance(documentuseDB *DocumentUseDB) (Error error) {

	documentuse, ok := backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[documentuseDB.ID]
	if !ok {
		documentuse = new(models.DocumentUse)

		backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[documentuseDB.ID] = documentuse
		backRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID[documentuse] = documentuseDB.ID

		// append model store with the new element
		documentuse.Name = documentuseDB.Name_Data.String
		documentuse.Stage(backRepoDocumentUse.GetStage())
	}
	documentuseDB.CopyBasicFieldsToDocumentUse(documentuse)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	documentuse.Stage(backRepoDocumentUse.GetStage())

	// preserve pointer to documentuseDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_DocumentUseDBID_DocumentUseDB)[documentuseDB hold variable pointers
	documentuseDB_Data := *documentuseDB
	preservedPtrToDocumentUse := &documentuseDB_Data
	backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB[documentuseDB.ID] = preservedPtrToDocumentUse

	return
}

// BackRepoDocumentUse.CheckoutPhaseTwo Checkouts all staged instances of DocumentUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, documentuseDB := range backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB {
		backRepoDocumentUse.CheckoutPhaseTwoInstance(backRepo, documentuseDB)
	}
	return
}

// BackRepoDocumentUse.CheckoutPhaseTwoInstance Checkouts staged instances of DocumentUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocumentUse *BackRepoDocumentUseStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, documentuseDB *DocumentUseDB) (Error error) {

	documentuse := backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[documentuseDB.ID]

	documentuseDB.DecodePointers(backRepo, documentuse)

	return
}

func (documentuseDB *DocumentUseDB) DecodePointers(backRepo *BackRepoStruct, documentuse *models.DocumentUse) {

	// insertion point for checkout of pointer encoding
	// Document field
	documentuse.Document = nil
	if documentuseDB.DocumentID.Int64 != 0 {
		documentuse.Document = backRepo.BackRepoDocument.Map_DocumentDBID_DocumentPtr[uint(documentuseDB.DocumentID.Int64)]
	}
	return
}

// CommitDocumentUse allows commit of a single documentuse (if already staged)
func (backRepo *BackRepoStruct) CommitDocumentUse(documentuse *models.DocumentUse) {
	backRepo.BackRepoDocumentUse.CommitPhaseOneInstance(documentuse)
	if id, ok := backRepo.BackRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID[documentuse]; ok {
		backRepo.BackRepoDocumentUse.CommitPhaseTwoInstance(backRepo, id, documentuse)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitDocumentUse allows checkout of a single documentuse (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutDocumentUse(documentuse *models.DocumentUse) {
	// check if the documentuse is staged
	if _, ok := backRepo.BackRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID[documentuse]; ok {

		if id, ok := backRepo.BackRepoDocumentUse.Map_DocumentUsePtr_DocumentUseDBID[documentuse]; ok {
			var documentuseDB DocumentUseDB
			documentuseDB.ID = id

			if err := backRepo.BackRepoDocumentUse.db.First(&documentuseDB, id).Error; err != nil {
				log.Fatalln("CheckoutDocumentUse : Problem with getting object with id:", id)
			}
			backRepo.BackRepoDocumentUse.CheckoutPhaseOneInstance(&documentuseDB)
			backRepo.BackRepoDocumentUse.CheckoutPhaseTwoInstance(backRepo, &documentuseDB)
		}
	}
}

// CopyBasicFieldsFromDocumentUse
func (documentuseDB *DocumentUseDB) CopyBasicFieldsFromDocumentUse(documentuse *models.DocumentUse) {
	// insertion point for fields commit

	documentuseDB.Name_Data.String = documentuse.Name
	documentuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromDocumentUse_WOP
func (documentuseDB *DocumentUseDB) CopyBasicFieldsFromDocumentUse_WOP(documentuse *models.DocumentUse_WOP) {
	// insertion point for fields commit

	documentuseDB.Name_Data.String = documentuse.Name
	documentuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromDocumentUseWOP
func (documentuseDB *DocumentUseDB) CopyBasicFieldsFromDocumentUseWOP(documentuse *DocumentUseWOP) {
	// insertion point for fields commit

	documentuseDB.Name_Data.String = documentuse.Name
	documentuseDB.Name_Data.Valid = true
}

// CopyBasicFieldsToDocumentUse
func (documentuseDB *DocumentUseDB) CopyBasicFieldsToDocumentUse(documentuse *models.DocumentUse) {
	// insertion point for checkout of basic fields (back repo to stage)
	documentuse.Name = documentuseDB.Name_Data.String
}

// CopyBasicFieldsToDocumentUse_WOP
func (documentuseDB *DocumentUseDB) CopyBasicFieldsToDocumentUse_WOP(documentuse *models.DocumentUse_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	documentuse.Name = documentuseDB.Name_Data.String
}

// CopyBasicFieldsToDocumentUseWOP
func (documentuseDB *DocumentUseDB) CopyBasicFieldsToDocumentUseWOP(documentuse *DocumentUseWOP) {
	documentuse.ID = int(documentuseDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	documentuse.Name = documentuseDB.Name_Data.String
}

// Backup generates a json file from a slice of all DocumentUseDB instances in the backrepo
func (backRepoDocumentUse *BackRepoDocumentUseStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "DocumentUseDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DocumentUseDB, 0)
	for _, documentuseDB := range backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB {
		forBackup = append(forBackup, documentuseDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json DocumentUse ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json DocumentUse file", err.Error())
	}
}

// Backup generates a json file from a slice of all DocumentUseDB instances in the backrepo
func (backRepoDocumentUse *BackRepoDocumentUseStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DocumentUseDB, 0)
	for _, documentuseDB := range backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB {
		forBackup = append(forBackup, documentuseDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("DocumentUse")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&DocumentUse_Fields, -1)
	for _, documentuseDB := range forBackup {

		var documentuseWOP DocumentUseWOP
		documentuseDB.CopyBasicFieldsToDocumentUseWOP(&documentuseWOP)

		row := sh.AddRow()
		row.WriteStruct(&documentuseWOP, -1)
	}
}

// RestoreXL from the "DocumentUse" sheet all DocumentUseDB instances
func (backRepoDocumentUse *BackRepoDocumentUseStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoDocumentUseid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["DocumentUse"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoDocumentUse.rowVisitorDocumentUse)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoDocumentUse *BackRepoDocumentUseStruct) rowVisitorDocumentUse(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var documentuseWOP DocumentUseWOP
		row.ReadStruct(&documentuseWOP)

		// add the unmarshalled struct to the stage
		documentuseDB := new(DocumentUseDB)
		documentuseDB.CopyBasicFieldsFromDocumentUseWOP(&documentuseWOP)

		documentuseDB_ID_atBackupTime := documentuseDB.ID
		documentuseDB.ID = 0
		query := backRepoDocumentUse.db.Create(documentuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB[documentuseDB.ID] = documentuseDB
		BackRepoDocumentUseid_atBckpTime_newID[documentuseDB_ID_atBackupTime] = documentuseDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "DocumentUseDB.json" in dirPath that stores an array
// of DocumentUseDB and stores it in the database
// the map BackRepoDocumentUseid_atBckpTime_newID is updated accordingly
func (backRepoDocumentUse *BackRepoDocumentUseStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoDocumentUseid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "DocumentUseDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json DocumentUse file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*DocumentUseDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_DocumentUseDBID_DocumentUseDB
	for _, documentuseDB := range forRestore {

		documentuseDB_ID_atBackupTime := documentuseDB.ID
		documentuseDB.ID = 0
		query := backRepoDocumentUse.db.Create(documentuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB[documentuseDB.ID] = documentuseDB
		BackRepoDocumentUseid_atBckpTime_newID[documentuseDB_ID_atBackupTime] = documentuseDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json DocumentUse file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<DocumentUse>id_atBckpTime_newID
// to compute new index
func (backRepoDocumentUse *BackRepoDocumentUseStruct) RestorePhaseTwo() {

	for _, documentuseDB := range backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB {

		// next line of code is to avert unused variable compilation error
		_ = documentuseDB

		// insertion point for reindexing pointers encoding
		// reindexing Document field
		if documentuseDB.DocumentID.Int64 != 0 {
			documentuseDB.DocumentID.Int64 = int64(BackRepoDocumentid_atBckpTime_newID[uint(documentuseDB.DocumentID.Int64)])
			documentuseDB.DocumentID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoDocumentUse.db.Model(documentuseDB).Updates(*documentuseDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoDocumentUse.ResetReversePointers commits all staged instances of DocumentUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocumentUse *BackRepoDocumentUseStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, documentuse := range backRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr {
		backRepoDocumentUse.ResetReversePointersInstance(backRepo, idx, documentuse)
	}

	return
}

func (backRepoDocumentUse *BackRepoDocumentUseStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, documentuse *models.DocumentUse) (Error error) {

	// fetch matching documentuseDB
	if documentuseDB, ok := backRepoDocumentUse.Map_DocumentUseDBID_DocumentUseDB[idx]; ok {
		_ = documentuseDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoDocumentUseid_atBckpTime_newID map[uint]uint