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
var dummy_Document_sql sql.NullBool
var dummy_Document_time time.Duration
var dummy_Document_sort sort.Float64Slice

// DocumentAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model documentAPI
type DocumentAPI struct {
	gorm.Model

	models.Document_WOP

	// encoding of pointers
	DocumentPointersEncoding DocumentPointersEncoding
}

// DocumentPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type DocumentPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field GeoObjectUse is a slice of pointers to another Struct (optional or 0..1)
	GeoObjectUse IntSlice `gorm:"type:TEXT"`
}

// DocumentDB describes a document in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model documentDB
type DocumentDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field documentDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	DocumentPointersEncoding
}

// DocumentDBs arrays documentDBs
// swagger:response documentDBsResponse
type DocumentDBs []DocumentDB

// DocumentDBResponse provides response
// swagger:response documentDBResponse
type DocumentDBResponse struct {
	DocumentDB
}

// DocumentWOP is a Document without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type DocumentWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Document_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoDocumentStruct struct {
	// stores DocumentDB according to their gorm ID
	Map_DocumentDBID_DocumentDB map[uint]*DocumentDB

	// stores DocumentDB ID according to Document address
	Map_DocumentPtr_DocumentDBID map[*models.Document]uint

	// stores Document according to their gorm ID
	Map_DocumentDBID_DocumentPtr map[uint]*models.Document

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoDocument *BackRepoDocumentStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoDocument.stage
	return
}

func (backRepoDocument *BackRepoDocumentStruct) GetDB() *gorm.DB {
	return backRepoDocument.db
}

// GetDocumentDBFromDocumentPtr is a handy function to access the back repo instance from the stage instance
func (backRepoDocument *BackRepoDocumentStruct) GetDocumentDBFromDocumentPtr(document *models.Document) (documentDB *DocumentDB) {
	id := backRepoDocument.Map_DocumentPtr_DocumentDBID[document]
	documentDB = backRepoDocument.Map_DocumentDBID_DocumentDB[id]
	return
}

// BackRepoDocument.CommitPhaseOne commits all staged instances of Document to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDocument *BackRepoDocumentStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for document := range stage.Documents {
		backRepoDocument.CommitPhaseOneInstance(document)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, document := range backRepoDocument.Map_DocumentDBID_DocumentPtr {
		if _, ok := stage.Documents[document]; !ok {
			backRepoDocument.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoDocument.CommitDeleteInstance commits deletion of Document to the BackRepo
func (backRepoDocument *BackRepoDocumentStruct) CommitDeleteInstance(id uint) (Error error) {

	document := backRepoDocument.Map_DocumentDBID_DocumentPtr[id]

	// document is not staged anymore, remove documentDB
	documentDB := backRepoDocument.Map_DocumentDBID_DocumentDB[id]
	query := backRepoDocument.db.Unscoped().Delete(&documentDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoDocument.Map_DocumentPtr_DocumentDBID, document)
	delete(backRepoDocument.Map_DocumentDBID_DocumentPtr, id)
	delete(backRepoDocument.Map_DocumentDBID_DocumentDB, id)

	return
}

// BackRepoDocument.CommitPhaseOneInstance commits document staged instances of Document to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDocument *BackRepoDocumentStruct) CommitPhaseOneInstance(document *models.Document) (Error error) {

	// check if the document is not commited yet
	if _, ok := backRepoDocument.Map_DocumentPtr_DocumentDBID[document]; ok {
		return
	}

	// initiate document
	var documentDB DocumentDB
	documentDB.CopyBasicFieldsFromDocument(document)

	query := backRepoDocument.db.Create(&documentDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoDocument.Map_DocumentPtr_DocumentDBID[document] = documentDB.ID
	backRepoDocument.Map_DocumentDBID_DocumentPtr[documentDB.ID] = document
	backRepoDocument.Map_DocumentDBID_DocumentDB[documentDB.ID] = &documentDB

	return
}

// BackRepoDocument.CommitPhaseTwo commits all staged instances of Document to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocument *BackRepoDocumentStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, document := range backRepoDocument.Map_DocumentDBID_DocumentPtr {
		backRepoDocument.CommitPhaseTwoInstance(backRepo, idx, document)
	}

	return
}

// BackRepoDocument.CommitPhaseTwoInstance commits {{structname }} of models.Document to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocument *BackRepoDocumentStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, document *models.Document) (Error error) {

	// fetch matching documentDB
	if documentDB, ok := backRepoDocument.Map_DocumentDBID_DocumentDB[idx]; ok {

		documentDB.CopyBasicFieldsFromDocument(document)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		documentDB.DocumentPointersEncoding.GeoObjectUse = make([]int, 0)
		// 2. encode
		for _, geoobjectuseAssocEnd := range document.GeoObjectUse {
			geoobjectuseAssocEnd_DB :=
				backRepo.BackRepoGeoObjectUse.GetGeoObjectUseDBFromGeoObjectUsePtr(geoobjectuseAssocEnd)

			// the stage might be inconsistant, meaning that the geoobjectuseAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if geoobjectuseAssocEnd_DB == nil {
				continue
			}

			documentDB.DocumentPointersEncoding.GeoObjectUse =
				append(documentDB.DocumentPointersEncoding.GeoObjectUse, int(geoobjectuseAssocEnd_DB.ID))
		}

		query := backRepoDocument.db.Save(&documentDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Document intance %s", document.Name))
		return err
	}

	return
}

// BackRepoDocument.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoDocument *BackRepoDocumentStruct) CheckoutPhaseOne() (Error error) {

	documentDBArray := make([]DocumentDB, 0)
	query := backRepoDocument.db.Find(&documentDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	documentInstancesToBeRemovedFromTheStage := make(map[*models.Document]any)
	for key, value := range backRepoDocument.stage.Documents {
		documentInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, documentDB := range documentDBArray {
		backRepoDocument.CheckoutPhaseOneInstance(&documentDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		document, ok := backRepoDocument.Map_DocumentDBID_DocumentPtr[documentDB.ID]
		if ok {
			delete(documentInstancesToBeRemovedFromTheStage, document)
		}
	}

	// remove from stage and back repo's 3 maps all documents that are not in the checkout
	for document := range documentInstancesToBeRemovedFromTheStage {
		document.Unstage(backRepoDocument.GetStage())

		// remove instance from the back repo 3 maps
		documentID := backRepoDocument.Map_DocumentPtr_DocumentDBID[document]
		delete(backRepoDocument.Map_DocumentPtr_DocumentDBID, document)
		delete(backRepoDocument.Map_DocumentDBID_DocumentDB, documentID)
		delete(backRepoDocument.Map_DocumentDBID_DocumentPtr, documentID)
	}

	return
}

// CheckoutPhaseOneInstance takes a documentDB that has been found in the DB, updates the backRepo and stages the
// models version of the documentDB
func (backRepoDocument *BackRepoDocumentStruct) CheckoutPhaseOneInstance(documentDB *DocumentDB) (Error error) {

	document, ok := backRepoDocument.Map_DocumentDBID_DocumentPtr[documentDB.ID]
	if !ok {
		document = new(models.Document)

		backRepoDocument.Map_DocumentDBID_DocumentPtr[documentDB.ID] = document
		backRepoDocument.Map_DocumentPtr_DocumentDBID[document] = documentDB.ID

		// append model store with the new element
		document.Name = documentDB.Name_Data.String
		document.Stage(backRepoDocument.GetStage())
	}
	documentDB.CopyBasicFieldsToDocument(document)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	document.Stage(backRepoDocument.GetStage())

	// preserve pointer to documentDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_DocumentDBID_DocumentDB)[documentDB hold variable pointers
	documentDB_Data := *documentDB
	preservedPtrToDocument := &documentDB_Data
	backRepoDocument.Map_DocumentDBID_DocumentDB[documentDB.ID] = preservedPtrToDocument

	return
}

// BackRepoDocument.CheckoutPhaseTwo Checkouts all staged instances of Document to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocument *BackRepoDocumentStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, documentDB := range backRepoDocument.Map_DocumentDBID_DocumentDB {
		backRepoDocument.CheckoutPhaseTwoInstance(backRepo, documentDB)
	}
	return
}

// BackRepoDocument.CheckoutPhaseTwoInstance Checkouts staged instances of Document to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocument *BackRepoDocumentStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, documentDB *DocumentDB) (Error error) {

	document := backRepoDocument.Map_DocumentDBID_DocumentPtr[documentDB.ID]

	documentDB.DecodePointers(backRepo, document)

	return
}

func (documentDB *DocumentDB) DecodePointers(backRepo *BackRepoStruct, document *models.Document) {

	// insertion point for checkout of pointer encoding
	// This loop redeem document.GeoObjectUse in the stage from the encode in the back repo
	// It parses all GeoObjectUseDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	document.GeoObjectUse = document.GeoObjectUse[:0]
	for _, _GeoObjectUseid := range documentDB.DocumentPointersEncoding.GeoObjectUse {
		document.GeoObjectUse = append(document.GeoObjectUse, backRepo.BackRepoGeoObjectUse.Map_GeoObjectUseDBID_GeoObjectUsePtr[uint(_GeoObjectUseid)])
	}

	return
}

// CommitDocument allows commit of a single document (if already staged)
func (backRepo *BackRepoStruct) CommitDocument(document *models.Document) {
	backRepo.BackRepoDocument.CommitPhaseOneInstance(document)
	if id, ok := backRepo.BackRepoDocument.Map_DocumentPtr_DocumentDBID[document]; ok {
		backRepo.BackRepoDocument.CommitPhaseTwoInstance(backRepo, id, document)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitDocument allows checkout of a single document (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutDocument(document *models.Document) {
	// check if the document is staged
	if _, ok := backRepo.BackRepoDocument.Map_DocumentPtr_DocumentDBID[document]; ok {

		if id, ok := backRepo.BackRepoDocument.Map_DocumentPtr_DocumentDBID[document]; ok {
			var documentDB DocumentDB
			documentDB.ID = id

			if err := backRepo.BackRepoDocument.db.First(&documentDB, id).Error; err != nil {
				log.Fatalln("CheckoutDocument : Problem with getting object with id:", id)
			}
			backRepo.BackRepoDocument.CheckoutPhaseOneInstance(&documentDB)
			backRepo.BackRepoDocument.CheckoutPhaseTwoInstance(backRepo, &documentDB)
		}
	}
}

// CopyBasicFieldsFromDocument
func (documentDB *DocumentDB) CopyBasicFieldsFromDocument(document *models.Document) {
	// insertion point for fields commit

	documentDB.Name_Data.String = document.Name
	documentDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromDocument_WOP
func (documentDB *DocumentDB) CopyBasicFieldsFromDocument_WOP(document *models.Document_WOP) {
	// insertion point for fields commit

	documentDB.Name_Data.String = document.Name
	documentDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromDocumentWOP
func (documentDB *DocumentDB) CopyBasicFieldsFromDocumentWOP(document *DocumentWOP) {
	// insertion point for fields commit

	documentDB.Name_Data.String = document.Name
	documentDB.Name_Data.Valid = true
}

// CopyBasicFieldsToDocument
func (documentDB *DocumentDB) CopyBasicFieldsToDocument(document *models.Document) {
	// insertion point for checkout of basic fields (back repo to stage)
	document.Name = documentDB.Name_Data.String
}

// CopyBasicFieldsToDocument_WOP
func (documentDB *DocumentDB) CopyBasicFieldsToDocument_WOP(document *models.Document_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	document.Name = documentDB.Name_Data.String
}

// CopyBasicFieldsToDocumentWOP
func (documentDB *DocumentDB) CopyBasicFieldsToDocumentWOP(document *DocumentWOP) {
	document.ID = int(documentDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	document.Name = documentDB.Name_Data.String
}

// Backup generates a json file from a slice of all DocumentDB instances in the backrepo
func (backRepoDocument *BackRepoDocumentStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "DocumentDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DocumentDB, 0)
	for _, documentDB := range backRepoDocument.Map_DocumentDBID_DocumentDB {
		forBackup = append(forBackup, documentDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Document ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Document file", err.Error())
	}
}

// Backup generates a json file from a slice of all DocumentDB instances in the backrepo
func (backRepoDocument *BackRepoDocumentStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DocumentDB, 0)
	for _, documentDB := range backRepoDocument.Map_DocumentDBID_DocumentDB {
		forBackup = append(forBackup, documentDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Document")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Document_Fields, -1)
	for _, documentDB := range forBackup {

		var documentWOP DocumentWOP
		documentDB.CopyBasicFieldsToDocumentWOP(&documentWOP)

		row := sh.AddRow()
		row.WriteStruct(&documentWOP, -1)
	}
}

// RestoreXL from the "Document" sheet all DocumentDB instances
func (backRepoDocument *BackRepoDocumentStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoDocumentid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Document"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoDocument.rowVisitorDocument)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoDocument *BackRepoDocumentStruct) rowVisitorDocument(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var documentWOP DocumentWOP
		row.ReadStruct(&documentWOP)

		// add the unmarshalled struct to the stage
		documentDB := new(DocumentDB)
		documentDB.CopyBasicFieldsFromDocumentWOP(&documentWOP)

		documentDB_ID_atBackupTime := documentDB.ID
		documentDB.ID = 0
		query := backRepoDocument.db.Create(documentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDocument.Map_DocumentDBID_DocumentDB[documentDB.ID] = documentDB
		BackRepoDocumentid_atBckpTime_newID[documentDB_ID_atBackupTime] = documentDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "DocumentDB.json" in dirPath that stores an array
// of DocumentDB and stores it in the database
// the map BackRepoDocumentid_atBckpTime_newID is updated accordingly
func (backRepoDocument *BackRepoDocumentStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoDocumentid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "DocumentDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Document file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*DocumentDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_DocumentDBID_DocumentDB
	for _, documentDB := range forRestore {

		documentDB_ID_atBackupTime := documentDB.ID
		documentDB.ID = 0
		query := backRepoDocument.db.Create(documentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDocument.Map_DocumentDBID_DocumentDB[documentDB.ID] = documentDB
		BackRepoDocumentid_atBckpTime_newID[documentDB_ID_atBackupTime] = documentDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Document file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Document>id_atBckpTime_newID
// to compute new index
func (backRepoDocument *BackRepoDocumentStruct) RestorePhaseTwo() {

	for _, documentDB := range backRepoDocument.Map_DocumentDBID_DocumentDB {

		// next line of code is to avert unused variable compilation error
		_ = documentDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoDocument.db.Model(documentDB).Updates(*documentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoDocument.ResetReversePointers commits all staged instances of Document to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDocument *BackRepoDocumentStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, document := range backRepoDocument.Map_DocumentDBID_DocumentPtr {
		backRepoDocument.ResetReversePointersInstance(backRepo, idx, document)
	}

	return
}

func (backRepoDocument *BackRepoDocumentStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, document *models.Document) (Error error) {

	// fetch matching documentDB
	if documentDB, ok := backRepoDocument.Map_DocumentDBID_DocumentDB[idx]; ok {
		_ = documentDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoDocumentid_atBckpTime_newID map[uint]uint