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
var dummy_User_sql sql.NullBool
var dummy_User_time time.Duration
var dummy_User_sort sort.Float64Slice

// UserAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model userAPI
type UserAPI struct {
	gorm.Model

	models.User_WOP

	// encoding of pointers
	UserPointersEncoding UserPointersEncoding
}

// UserPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type UserPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// UserDB describes a user in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model userDB
type UserDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field userDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	UserPointersEncoding
}

// UserDBs arrays userDBs
// swagger:response userDBsResponse
type UserDBs []UserDB

// UserDBResponse provides response
// swagger:response userDBResponse
type UserDBResponse struct {
	UserDB
}

// UserWOP is a User without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type UserWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var User_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoUserStruct struct {
	// stores UserDB according to their gorm ID
	Map_UserDBID_UserDB map[uint]*UserDB

	// stores UserDB ID according to User address
	Map_UserPtr_UserDBID map[*models.User]uint

	// stores User according to their gorm ID
	Map_UserDBID_UserPtr map[uint]*models.User

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoUser *BackRepoUserStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoUser.stage
	return
}

func (backRepoUser *BackRepoUserStruct) GetDB() *gorm.DB {
	return backRepoUser.db
}

// GetUserDBFromUserPtr is a handy function to access the back repo instance from the stage instance
func (backRepoUser *BackRepoUserStruct) GetUserDBFromUserPtr(user *models.User) (userDB *UserDB) {
	id := backRepoUser.Map_UserPtr_UserDBID[user]
	userDB = backRepoUser.Map_UserDBID_UserDB[id]
	return
}

// BackRepoUser.CommitPhaseOne commits all staged instances of User to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUser *BackRepoUserStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for user := range stage.Users {
		backRepoUser.CommitPhaseOneInstance(user)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, user := range backRepoUser.Map_UserDBID_UserPtr {
		if _, ok := stage.Users[user]; !ok {
			backRepoUser.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoUser.CommitDeleteInstance commits deletion of User to the BackRepo
func (backRepoUser *BackRepoUserStruct) CommitDeleteInstance(id uint) (Error error) {

	user := backRepoUser.Map_UserDBID_UserPtr[id]

	// user is not staged anymore, remove userDB
	userDB := backRepoUser.Map_UserDBID_UserDB[id]
	query := backRepoUser.db.Unscoped().Delete(&userDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoUser.Map_UserPtr_UserDBID, user)
	delete(backRepoUser.Map_UserDBID_UserPtr, id)
	delete(backRepoUser.Map_UserDBID_UserDB, id)

	return
}

// BackRepoUser.CommitPhaseOneInstance commits user staged instances of User to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUser *BackRepoUserStruct) CommitPhaseOneInstance(user *models.User) (Error error) {

	// check if the user is not commited yet
	if _, ok := backRepoUser.Map_UserPtr_UserDBID[user]; ok {
		return
	}

	// initiate user
	var userDB UserDB
	userDB.CopyBasicFieldsFromUser(user)

	query := backRepoUser.db.Create(&userDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoUser.Map_UserPtr_UserDBID[user] = userDB.ID
	backRepoUser.Map_UserDBID_UserPtr[userDB.ID] = user
	backRepoUser.Map_UserDBID_UserDB[userDB.ID] = &userDB

	return
}

// BackRepoUser.CommitPhaseTwo commits all staged instances of User to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUser *BackRepoUserStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, user := range backRepoUser.Map_UserDBID_UserPtr {
		backRepoUser.CommitPhaseTwoInstance(backRepo, idx, user)
	}

	return
}

// BackRepoUser.CommitPhaseTwoInstance commits {{structname }} of models.User to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUser *BackRepoUserStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, user *models.User) (Error error) {

	// fetch matching userDB
	if userDB, ok := backRepoUser.Map_UserDBID_UserDB[idx]; ok {

		userDB.CopyBasicFieldsFromUser(user)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoUser.db.Save(&userDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown User intance %s", user.Name))
		return err
	}

	return
}

// BackRepoUser.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoUser *BackRepoUserStruct) CheckoutPhaseOne() (Error error) {

	userDBArray := make([]UserDB, 0)
	query := backRepoUser.db.Find(&userDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	userInstancesToBeRemovedFromTheStage := make(map[*models.User]any)
	for key, value := range backRepoUser.stage.Users {
		userInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, userDB := range userDBArray {
		backRepoUser.CheckoutPhaseOneInstance(&userDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		user, ok := backRepoUser.Map_UserDBID_UserPtr[userDB.ID]
		if ok {
			delete(userInstancesToBeRemovedFromTheStage, user)
		}
	}

	// remove from stage and back repo's 3 maps all users that are not in the checkout
	for user := range userInstancesToBeRemovedFromTheStage {
		user.Unstage(backRepoUser.GetStage())

		// remove instance from the back repo 3 maps
		userID := backRepoUser.Map_UserPtr_UserDBID[user]
		delete(backRepoUser.Map_UserPtr_UserDBID, user)
		delete(backRepoUser.Map_UserDBID_UserDB, userID)
		delete(backRepoUser.Map_UserDBID_UserPtr, userID)
	}

	return
}

// CheckoutPhaseOneInstance takes a userDB that has been found in the DB, updates the backRepo and stages the
// models version of the userDB
func (backRepoUser *BackRepoUserStruct) CheckoutPhaseOneInstance(userDB *UserDB) (Error error) {

	user, ok := backRepoUser.Map_UserDBID_UserPtr[userDB.ID]
	if !ok {
		user = new(models.User)

		backRepoUser.Map_UserDBID_UserPtr[userDB.ID] = user
		backRepoUser.Map_UserPtr_UserDBID[user] = userDB.ID

		// append model store with the new element
		user.Name = userDB.Name_Data.String
		user.Stage(backRepoUser.GetStage())
	}
	userDB.CopyBasicFieldsToUser(user)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	user.Stage(backRepoUser.GetStage())

	// preserve pointer to userDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_UserDBID_UserDB)[userDB hold variable pointers
	userDB_Data := *userDB
	preservedPtrToUser := &userDB_Data
	backRepoUser.Map_UserDBID_UserDB[userDB.ID] = preservedPtrToUser

	return
}

// BackRepoUser.CheckoutPhaseTwo Checkouts all staged instances of User to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUser *BackRepoUserStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, userDB := range backRepoUser.Map_UserDBID_UserDB {
		backRepoUser.CheckoutPhaseTwoInstance(backRepo, userDB)
	}
	return
}

// BackRepoUser.CheckoutPhaseTwoInstance Checkouts staged instances of User to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUser *BackRepoUserStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, userDB *UserDB) (Error error) {

	user := backRepoUser.Map_UserDBID_UserPtr[userDB.ID]

	userDB.DecodePointers(backRepo, user)

	return
}

func (userDB *UserDB) DecodePointers(backRepo *BackRepoStruct, user *models.User) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitUser allows commit of a single user (if already staged)
func (backRepo *BackRepoStruct) CommitUser(user *models.User) {
	backRepo.BackRepoUser.CommitPhaseOneInstance(user)
	if id, ok := backRepo.BackRepoUser.Map_UserPtr_UserDBID[user]; ok {
		backRepo.BackRepoUser.CommitPhaseTwoInstance(backRepo, id, user)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitUser allows checkout of a single user (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutUser(user *models.User) {
	// check if the user is staged
	if _, ok := backRepo.BackRepoUser.Map_UserPtr_UserDBID[user]; ok {

		if id, ok := backRepo.BackRepoUser.Map_UserPtr_UserDBID[user]; ok {
			var userDB UserDB
			userDB.ID = id

			if err := backRepo.BackRepoUser.db.First(&userDB, id).Error; err != nil {
				log.Fatalln("CheckoutUser : Problem with getting object with id:", id)
			}
			backRepo.BackRepoUser.CheckoutPhaseOneInstance(&userDB)
			backRepo.BackRepoUser.CheckoutPhaseTwoInstance(backRepo, &userDB)
		}
	}
}

// CopyBasicFieldsFromUser
func (userDB *UserDB) CopyBasicFieldsFromUser(user *models.User) {
	// insertion point for fields commit

	userDB.Name_Data.String = user.Name
	userDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromUser_WOP
func (userDB *UserDB) CopyBasicFieldsFromUser_WOP(user *models.User_WOP) {
	// insertion point for fields commit

	userDB.Name_Data.String = user.Name
	userDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromUserWOP
func (userDB *UserDB) CopyBasicFieldsFromUserWOP(user *UserWOP) {
	// insertion point for fields commit

	userDB.Name_Data.String = user.Name
	userDB.Name_Data.Valid = true
}

// CopyBasicFieldsToUser
func (userDB *UserDB) CopyBasicFieldsToUser(user *models.User) {
	// insertion point for checkout of basic fields (back repo to stage)
	user.Name = userDB.Name_Data.String
}

// CopyBasicFieldsToUser_WOP
func (userDB *UserDB) CopyBasicFieldsToUser_WOP(user *models.User_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	user.Name = userDB.Name_Data.String
}

// CopyBasicFieldsToUserWOP
func (userDB *UserDB) CopyBasicFieldsToUserWOP(user *UserWOP) {
	user.ID = int(userDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	user.Name = userDB.Name_Data.String
}

// Backup generates a json file from a slice of all UserDB instances in the backrepo
func (backRepoUser *BackRepoUserStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "UserDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UserDB, 0)
	for _, userDB := range backRepoUser.Map_UserDBID_UserDB {
		forBackup = append(forBackup, userDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json User ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json User file", err.Error())
	}
}

// Backup generates a json file from a slice of all UserDB instances in the backrepo
func (backRepoUser *BackRepoUserStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UserDB, 0)
	for _, userDB := range backRepoUser.Map_UserDBID_UserDB {
		forBackup = append(forBackup, userDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("User")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&User_Fields, -1)
	for _, userDB := range forBackup {

		var userWOP UserWOP
		userDB.CopyBasicFieldsToUserWOP(&userWOP)

		row := sh.AddRow()
		row.WriteStruct(&userWOP, -1)
	}
}

// RestoreXL from the "User" sheet all UserDB instances
func (backRepoUser *BackRepoUserStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoUserid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["User"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoUser.rowVisitorUser)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoUser *BackRepoUserStruct) rowVisitorUser(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var userWOP UserWOP
		row.ReadStruct(&userWOP)

		// add the unmarshalled struct to the stage
		userDB := new(UserDB)
		userDB.CopyBasicFieldsFromUserWOP(&userWOP)

		userDB_ID_atBackupTime := userDB.ID
		userDB.ID = 0
		query := backRepoUser.db.Create(userDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoUser.Map_UserDBID_UserDB[userDB.ID] = userDB
		BackRepoUserid_atBckpTime_newID[userDB_ID_atBackupTime] = userDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "UserDB.json" in dirPath that stores an array
// of UserDB and stores it in the database
// the map BackRepoUserid_atBckpTime_newID is updated accordingly
func (backRepoUser *BackRepoUserStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoUserid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "UserDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json User file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*UserDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_UserDBID_UserDB
	for _, userDB := range forRestore {

		userDB_ID_atBackupTime := userDB.ID
		userDB.ID = 0
		query := backRepoUser.db.Create(userDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoUser.Map_UserDBID_UserDB[userDB.ID] = userDB
		BackRepoUserid_atBckpTime_newID[userDB_ID_atBackupTime] = userDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json User file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<User>id_atBckpTime_newID
// to compute new index
func (backRepoUser *BackRepoUserStruct) RestorePhaseTwo() {

	for _, userDB := range backRepoUser.Map_UserDBID_UserDB {

		// next line of code is to avert unused variable compilation error
		_ = userDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoUser.db.Model(userDB).Updates(*userDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoUser.ResetReversePointers commits all staged instances of User to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUser *BackRepoUserStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, user := range backRepoUser.Map_UserDBID_UserPtr {
		backRepoUser.ResetReversePointersInstance(backRepo, idx, user)
	}

	return
}

func (backRepoUser *BackRepoUserStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, user *models.User) (Error error) {

	// fetch matching userDB
	if userDB, ok := backRepoUser.Map_UserDBID_UserDB[idx]; ok {
		_ = userDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoUserid_atBckpTime_newID map[uint]uint
