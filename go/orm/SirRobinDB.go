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
var dummy_SirRobin_sql sql.NullBool
var dummy_SirRobin_time time.Duration
var dummy_SirRobin_sort sort.Float64Slice

// SirRobinAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model sirrobinAPI
type SirRobinAPI struct {
	gorm.Model

	models.SirRobin_WOP

	// encoding of pointers
	SirRobinPointersEncoding SirRobinPointersEncoding
}

// SirRobinPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SirRobinPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Witches is a slice of pointers to another Struct (optional or 0..1)
	Witches IntSlice `gorm:"type:TEXT"`

	// field Arthurs is a slice of pointers to another Struct (optional or 0..1)
	Arthurs IntSlice `gorm:"type:TEXT"`

	// field KnightWhoSayNis is a slice of pointers to another Struct (optional or 0..1)
	KnightWhoSayNis IntSlice `gorm:"type:TEXT"`

	// field BlackKnightShapes is a slice of pointers to another Struct (optional or 0..1)
	BlackKnightShapes IntSlice `gorm:"type:TEXT"`

	// field TheNuteShapes is a slice of pointers to another Struct (optional or 0..1)
	TheNuteShapes IntSlice `gorm:"type:TEXT"`
}

// SirRobinDB describes a sirrobin in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model sirrobinDB
type SirRobinDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field sirrobinDB.Name
	Name_Data sql.NullString

	// Declation for basic field sirrobinDB.Description
	Description_Data sql.NullString

	// Declation for basic field sirrobinDB.AxisOrign_X
	AxisOrign_X_Data sql.NullFloat64

	// Declation for basic field sirrobinDB.AxisOrign_Y
	AxisOrign_Y_Data sql.NullFloat64

	// Declation for basic field sirrobinDB.VerticalAxis_Top_Y
	VerticalAxis_Top_Y_Data sql.NullFloat64

	// Declation for basic field sirrobinDB.VerticalAxis_Bottom_Y
	VerticalAxis_Bottom_Y_Data sql.NullFloat64

	// Declation for basic field sirrobinDB.VerticalAxis_StrokeWidth
	VerticalAxis_StrokeWidth_Data sql.NullFloat64

	// Declation for basic field sirrobinDB.HorizontalAxis_Right_X
	HorizontalAxis_Right_X_Data sql.NullFloat64

	// Declation for basic field sirrobinDB.Start
	Start_Data sql.NullTime

	// Declation for basic field sirrobinDB.End
	End_Data sql.NullTime
	// encoding of pointers
	SirRobinPointersEncoding
}

// SirRobinDBs arrays sirrobinDBs
// swagger:response sirrobinDBsResponse
type SirRobinDBs []SirRobinDB

// SirRobinDBResponse provides response
// swagger:response sirrobinDBResponse
type SirRobinDBResponse struct {
	SirRobinDB
}

// SirRobinWOP is a SirRobin without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SirRobinWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Description string `xlsx:"2"`

	AxisOrign_X float64 `xlsx:"3"`

	AxisOrign_Y float64 `xlsx:"4"`

	VerticalAxis_Top_Y float64 `xlsx:"5"`

	VerticalAxis_Bottom_Y float64 `xlsx:"6"`

	VerticalAxis_StrokeWidth float64 `xlsx:"7"`

	HorizontalAxis_Right_X float64 `xlsx:"8"`

	Start time.Time `xlsx:"9"`

	End time.Time `xlsx:"10"`
	// insertion for WOP pointer fields
}

var SirRobin_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Description",
	"AxisOrign_X",
	"AxisOrign_Y",
	"VerticalAxis_Top_Y",
	"VerticalAxis_Bottom_Y",
	"VerticalAxis_StrokeWidth",
	"HorizontalAxis_Right_X",
	"Start",
	"End",
}

type BackRepoSirRobinStruct struct {
	// stores SirRobinDB according to their gorm ID
	Map_SirRobinDBID_SirRobinDB map[uint]*SirRobinDB

	// stores SirRobinDB ID according to SirRobin address
	Map_SirRobinPtr_SirRobinDBID map[*models.SirRobin]uint

	// stores SirRobin according to their gorm ID
	Map_SirRobinDBID_SirRobinPtr map[uint]*models.SirRobin

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSirRobin *BackRepoSirRobinStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSirRobin.stage
	return
}

func (backRepoSirRobin *BackRepoSirRobinStruct) GetDB() *gorm.DB {
	return backRepoSirRobin.db
}

// GetSirRobinDBFromSirRobinPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSirRobin *BackRepoSirRobinStruct) GetSirRobinDBFromSirRobinPtr(sirrobin *models.SirRobin) (sirrobinDB *SirRobinDB) {
	id := backRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[sirrobin]
	sirrobinDB = backRepoSirRobin.Map_SirRobinDBID_SirRobinDB[id]
	return
}

// BackRepoSirRobin.CommitPhaseOne commits all staged instances of SirRobin to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSirRobin *BackRepoSirRobinStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for sirrobin := range stage.SirRobins {
		backRepoSirRobin.CommitPhaseOneInstance(sirrobin)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, sirrobin := range backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr {
		if _, ok := stage.SirRobins[sirrobin]; !ok {
			backRepoSirRobin.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSirRobin.CommitDeleteInstance commits deletion of SirRobin to the BackRepo
func (backRepoSirRobin *BackRepoSirRobinStruct) CommitDeleteInstance(id uint) (Error error) {

	sirrobin := backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[id]

	// sirrobin is not staged anymore, remove sirrobinDB
	sirrobinDB := backRepoSirRobin.Map_SirRobinDBID_SirRobinDB[id]
	query := backRepoSirRobin.db.Unscoped().Delete(&sirrobinDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSirRobin.Map_SirRobinPtr_SirRobinDBID, sirrobin)
	delete(backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr, id)
	delete(backRepoSirRobin.Map_SirRobinDBID_SirRobinDB, id)

	return
}

// BackRepoSirRobin.CommitPhaseOneInstance commits sirrobin staged instances of SirRobin to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSirRobin *BackRepoSirRobinStruct) CommitPhaseOneInstance(sirrobin *models.SirRobin) (Error error) {

	// check if the sirrobin is not commited yet
	if _, ok := backRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[sirrobin]; ok {
		return
	}

	// initiate sirrobin
	var sirrobinDB SirRobinDB
	sirrobinDB.CopyBasicFieldsFromSirRobin(sirrobin)

	query := backRepoSirRobin.db.Create(&sirrobinDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[sirrobin] = sirrobinDB.ID
	backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[sirrobinDB.ID] = sirrobin
	backRepoSirRobin.Map_SirRobinDBID_SirRobinDB[sirrobinDB.ID] = &sirrobinDB

	return
}

// BackRepoSirRobin.CommitPhaseTwo commits all staged instances of SirRobin to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSirRobin *BackRepoSirRobinStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, sirrobin := range backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr {
		backRepoSirRobin.CommitPhaseTwoInstance(backRepo, idx, sirrobin)
	}

	return
}

// BackRepoSirRobin.CommitPhaseTwoInstance commits {{structname }} of models.SirRobin to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSirRobin *BackRepoSirRobinStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, sirrobin *models.SirRobin) (Error error) {

	// fetch matching sirrobinDB
	if sirrobinDB, ok := backRepoSirRobin.Map_SirRobinDBID_SirRobinDB[idx]; ok {

		sirrobinDB.CopyBasicFieldsFromSirRobin(sirrobin)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		sirrobinDB.SirRobinPointersEncoding.Witches = make([]int, 0)
		// 2. encode
		for _, awitchAssocEnd := range sirrobin.Witches {
			awitchAssocEnd_DB :=
				backRepo.BackRepoAWitch.GetAWitchDBFromAWitchPtr(awitchAssocEnd)

			// the stage might be inconsistant, meaning that the awitchAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if awitchAssocEnd_DB == nil {
				continue
			}

			sirrobinDB.SirRobinPointersEncoding.Witches =
				append(sirrobinDB.SirRobinPointersEncoding.Witches, int(awitchAssocEnd_DB.ID))
		}

		// 1. reset
		sirrobinDB.SirRobinPointersEncoding.Arthurs = make([]int, 0)
		// 2. encode
		for _, kingarthurshapeAssocEnd := range sirrobin.Arthurs {
			kingarthurshapeAssocEnd_DB :=
				backRepo.BackRepoKingArthurShape.GetKingArthurShapeDBFromKingArthurShapePtr(kingarthurshapeAssocEnd)

			// the stage might be inconsistant, meaning that the kingarthurshapeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if kingarthurshapeAssocEnd_DB == nil {
				continue
			}

			sirrobinDB.SirRobinPointersEncoding.Arthurs =
				append(sirrobinDB.SirRobinPointersEncoding.Arthurs, int(kingarthurshapeAssocEnd_DB.ID))
		}

		// 1. reset
		sirrobinDB.SirRobinPointersEncoding.KnightWhoSayNis = make([]int, 0)
		// 2. encode
		for _, knightwhosayniAssocEnd := range sirrobin.KnightWhoSayNis {
			knightwhosayniAssocEnd_DB :=
				backRepo.BackRepoKnightWhoSayNi.GetKnightWhoSayNiDBFromKnightWhoSayNiPtr(knightwhosayniAssocEnd)

			// the stage might be inconsistant, meaning that the knightwhosayniAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if knightwhosayniAssocEnd_DB == nil {
				continue
			}

			sirrobinDB.SirRobinPointersEncoding.KnightWhoSayNis =
				append(sirrobinDB.SirRobinPointersEncoding.KnightWhoSayNis, int(knightwhosayniAssocEnd_DB.ID))
		}

		// 1. reset
		sirrobinDB.SirRobinPointersEncoding.BlackKnightShapes = make([]int, 0)
		// 2. encode
		for _, blackknightshapeAssocEnd := range sirrobin.BlackKnightShapes {
			blackknightshapeAssocEnd_DB :=
				backRepo.BackRepoBlackKnightShape.GetBlackKnightShapeDBFromBlackKnightShapePtr(blackknightshapeAssocEnd)

			// the stage might be inconsistant, meaning that the blackknightshapeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if blackknightshapeAssocEnd_DB == nil {
				continue
			}

			sirrobinDB.SirRobinPointersEncoding.BlackKnightShapes =
				append(sirrobinDB.SirRobinPointersEncoding.BlackKnightShapes, int(blackknightshapeAssocEnd_DB.ID))
		}

		// 1. reset
		sirrobinDB.SirRobinPointersEncoding.TheNuteShapes = make([]int, 0)
		// 2. encode
		for _, thenuteshapeAssocEnd := range sirrobin.TheNuteShapes {
			thenuteshapeAssocEnd_DB :=
				backRepo.BackRepoTheNuteShape.GetTheNuteShapeDBFromTheNuteShapePtr(thenuteshapeAssocEnd)

			// the stage might be inconsistant, meaning that the thenuteshapeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if thenuteshapeAssocEnd_DB == nil {
				continue
			}

			sirrobinDB.SirRobinPointersEncoding.TheNuteShapes =
				append(sirrobinDB.SirRobinPointersEncoding.TheNuteShapes, int(thenuteshapeAssocEnd_DB.ID))
		}

		query := backRepoSirRobin.db.Save(&sirrobinDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SirRobin intance %s", sirrobin.Name))
		return err
	}

	return
}

// BackRepoSirRobin.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSirRobin *BackRepoSirRobinStruct) CheckoutPhaseOne() (Error error) {

	sirrobinDBArray := make([]SirRobinDB, 0)
	query := backRepoSirRobin.db.Find(&sirrobinDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	sirrobinInstancesToBeRemovedFromTheStage := make(map[*models.SirRobin]any)
	for key, value := range backRepoSirRobin.stage.SirRobins {
		sirrobinInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, sirrobinDB := range sirrobinDBArray {
		backRepoSirRobin.CheckoutPhaseOneInstance(&sirrobinDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		sirrobin, ok := backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[sirrobinDB.ID]
		if ok {
			delete(sirrobinInstancesToBeRemovedFromTheStage, sirrobin)
		}
	}

	// remove from stage and back repo's 3 maps all sirrobins that are not in the checkout
	for sirrobin := range sirrobinInstancesToBeRemovedFromTheStage {
		sirrobin.Unstage(backRepoSirRobin.GetStage())

		// remove instance from the back repo 3 maps
		sirrobinID := backRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[sirrobin]
		delete(backRepoSirRobin.Map_SirRobinPtr_SirRobinDBID, sirrobin)
		delete(backRepoSirRobin.Map_SirRobinDBID_SirRobinDB, sirrobinID)
		delete(backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr, sirrobinID)
	}

	return
}

// CheckoutPhaseOneInstance takes a sirrobinDB that has been found in the DB, updates the backRepo and stages the
// models version of the sirrobinDB
func (backRepoSirRobin *BackRepoSirRobinStruct) CheckoutPhaseOneInstance(sirrobinDB *SirRobinDB) (Error error) {

	sirrobin, ok := backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[sirrobinDB.ID]
	if !ok {
		sirrobin = new(models.SirRobin)

		backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[sirrobinDB.ID] = sirrobin
		backRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[sirrobin] = sirrobinDB.ID

		// append model store with the new element
		sirrobin.Name = sirrobinDB.Name_Data.String
		sirrobin.Stage(backRepoSirRobin.GetStage())
	}
	sirrobinDB.CopyBasicFieldsToSirRobin(sirrobin)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	sirrobin.Stage(backRepoSirRobin.GetStage())

	// preserve pointer to sirrobinDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SirRobinDBID_SirRobinDB)[sirrobinDB hold variable pointers
	sirrobinDB_Data := *sirrobinDB
	preservedPtrToSirRobin := &sirrobinDB_Data
	backRepoSirRobin.Map_SirRobinDBID_SirRobinDB[sirrobinDB.ID] = preservedPtrToSirRobin

	return
}

// BackRepoSirRobin.CheckoutPhaseTwo Checkouts all staged instances of SirRobin to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSirRobin *BackRepoSirRobinStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, sirrobinDB := range backRepoSirRobin.Map_SirRobinDBID_SirRobinDB {
		backRepoSirRobin.CheckoutPhaseTwoInstance(backRepo, sirrobinDB)
	}
	return
}

// BackRepoSirRobin.CheckoutPhaseTwoInstance Checkouts staged instances of SirRobin to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSirRobin *BackRepoSirRobinStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, sirrobinDB *SirRobinDB) (Error error) {

	sirrobin := backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[sirrobinDB.ID]

	sirrobinDB.DecodePointers(backRepo, sirrobin)

	return
}

func (sirrobinDB *SirRobinDB) DecodePointers(backRepo *BackRepoStruct, sirrobin *models.SirRobin) {

	// insertion point for checkout of pointer encoding
	// This loop redeem sirrobin.Witches in the stage from the encode in the back repo
	// It parses all AWitchDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sirrobin.Witches = sirrobin.Witches[:0]
	for _, _AWitchid := range sirrobinDB.SirRobinPointersEncoding.Witches {
		sirrobin.Witches = append(sirrobin.Witches, backRepo.BackRepoAWitch.Map_AWitchDBID_AWitchPtr[uint(_AWitchid)])
	}

	// This loop redeem sirrobin.Arthurs in the stage from the encode in the back repo
	// It parses all KingArthurShapeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sirrobin.Arthurs = sirrobin.Arthurs[:0]
	for _, _KingArthurShapeid := range sirrobinDB.SirRobinPointersEncoding.Arthurs {
		sirrobin.Arthurs = append(sirrobin.Arthurs, backRepo.BackRepoKingArthurShape.Map_KingArthurShapeDBID_KingArthurShapePtr[uint(_KingArthurShapeid)])
	}

	// This loop redeem sirrobin.KnightWhoSayNis in the stage from the encode in the back repo
	// It parses all KnightWhoSayNiDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sirrobin.KnightWhoSayNis = sirrobin.KnightWhoSayNis[:0]
	for _, _KnightWhoSayNiid := range sirrobinDB.SirRobinPointersEncoding.KnightWhoSayNis {
		sirrobin.KnightWhoSayNis = append(sirrobin.KnightWhoSayNis, backRepo.BackRepoKnightWhoSayNi.Map_KnightWhoSayNiDBID_KnightWhoSayNiPtr[uint(_KnightWhoSayNiid)])
	}

	// This loop redeem sirrobin.BlackKnightShapes in the stage from the encode in the back repo
	// It parses all BlackKnightShapeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sirrobin.BlackKnightShapes = sirrobin.BlackKnightShapes[:0]
	for _, _BlackKnightShapeid := range sirrobinDB.SirRobinPointersEncoding.BlackKnightShapes {
		sirrobin.BlackKnightShapes = append(sirrobin.BlackKnightShapes, backRepo.BackRepoBlackKnightShape.Map_BlackKnightShapeDBID_BlackKnightShapePtr[uint(_BlackKnightShapeid)])
	}

	// This loop redeem sirrobin.TheNuteShapes in the stage from the encode in the back repo
	// It parses all TheNuteShapeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sirrobin.TheNuteShapes = sirrobin.TheNuteShapes[:0]
	for _, _TheNuteShapeid := range sirrobinDB.SirRobinPointersEncoding.TheNuteShapes {
		sirrobin.TheNuteShapes = append(sirrobin.TheNuteShapes, backRepo.BackRepoTheNuteShape.Map_TheNuteShapeDBID_TheNuteShapePtr[uint(_TheNuteShapeid)])
	}

	return
}

// CommitSirRobin allows commit of a single sirrobin (if already staged)
func (backRepo *BackRepoStruct) CommitSirRobin(sirrobin *models.SirRobin) {
	backRepo.BackRepoSirRobin.CommitPhaseOneInstance(sirrobin)
	if id, ok := backRepo.BackRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[sirrobin]; ok {
		backRepo.BackRepoSirRobin.CommitPhaseTwoInstance(backRepo, id, sirrobin)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSirRobin allows checkout of a single sirrobin (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSirRobin(sirrobin *models.SirRobin) {
	// check if the sirrobin is staged
	if _, ok := backRepo.BackRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[sirrobin]; ok {

		if id, ok := backRepo.BackRepoSirRobin.Map_SirRobinPtr_SirRobinDBID[sirrobin]; ok {
			var sirrobinDB SirRobinDB
			sirrobinDB.ID = id

			if err := backRepo.BackRepoSirRobin.db.First(&sirrobinDB, id).Error; err != nil {
				log.Fatalln("CheckoutSirRobin : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSirRobin.CheckoutPhaseOneInstance(&sirrobinDB)
			backRepo.BackRepoSirRobin.CheckoutPhaseTwoInstance(backRepo, &sirrobinDB)
		}
	}
}

// CopyBasicFieldsFromSirRobin
func (sirrobinDB *SirRobinDB) CopyBasicFieldsFromSirRobin(sirrobin *models.SirRobin) {
	// insertion point for fields commit

	sirrobinDB.Name_Data.String = sirrobin.Name
	sirrobinDB.Name_Data.Valid = true

	sirrobinDB.Description_Data.String = sirrobin.Description
	sirrobinDB.Description_Data.Valid = true

	sirrobinDB.AxisOrign_X_Data.Float64 = sirrobin.AxisOrign_X
	sirrobinDB.AxisOrign_X_Data.Valid = true

	sirrobinDB.AxisOrign_Y_Data.Float64 = sirrobin.AxisOrign_Y
	sirrobinDB.AxisOrign_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_Top_Y_Data.Float64 = sirrobin.VerticalAxis_Top_Y
	sirrobinDB.VerticalAxis_Top_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_Bottom_Y_Data.Float64 = sirrobin.VerticalAxis_Bottom_Y
	sirrobinDB.VerticalAxis_Bottom_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_StrokeWidth_Data.Float64 = sirrobin.VerticalAxis_StrokeWidth
	sirrobinDB.VerticalAxis_StrokeWidth_Data.Valid = true

	sirrobinDB.HorizontalAxis_Right_X_Data.Float64 = sirrobin.HorizontalAxis_Right_X
	sirrobinDB.HorizontalAxis_Right_X_Data.Valid = true

	sirrobinDB.Start_Data.Time = sirrobin.Start
	sirrobinDB.Start_Data.Valid = true

	sirrobinDB.End_Data.Time = sirrobin.End
	sirrobinDB.End_Data.Valid = true
}

// CopyBasicFieldsFromSirRobin_WOP
func (sirrobinDB *SirRobinDB) CopyBasicFieldsFromSirRobin_WOP(sirrobin *models.SirRobin_WOP) {
	// insertion point for fields commit

	sirrobinDB.Name_Data.String = sirrobin.Name
	sirrobinDB.Name_Data.Valid = true

	sirrobinDB.Description_Data.String = sirrobin.Description
	sirrobinDB.Description_Data.Valid = true

	sirrobinDB.AxisOrign_X_Data.Float64 = sirrobin.AxisOrign_X
	sirrobinDB.AxisOrign_X_Data.Valid = true

	sirrobinDB.AxisOrign_Y_Data.Float64 = sirrobin.AxisOrign_Y
	sirrobinDB.AxisOrign_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_Top_Y_Data.Float64 = sirrobin.VerticalAxis_Top_Y
	sirrobinDB.VerticalAxis_Top_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_Bottom_Y_Data.Float64 = sirrobin.VerticalAxis_Bottom_Y
	sirrobinDB.VerticalAxis_Bottom_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_StrokeWidth_Data.Float64 = sirrobin.VerticalAxis_StrokeWidth
	sirrobinDB.VerticalAxis_StrokeWidth_Data.Valid = true

	sirrobinDB.HorizontalAxis_Right_X_Data.Float64 = sirrobin.HorizontalAxis_Right_X
	sirrobinDB.HorizontalAxis_Right_X_Data.Valid = true

	sirrobinDB.Start_Data.Time = sirrobin.Start
	sirrobinDB.Start_Data.Valid = true

	sirrobinDB.End_Data.Time = sirrobin.End
	sirrobinDB.End_Data.Valid = true
}

// CopyBasicFieldsFromSirRobinWOP
func (sirrobinDB *SirRobinDB) CopyBasicFieldsFromSirRobinWOP(sirrobin *SirRobinWOP) {
	// insertion point for fields commit

	sirrobinDB.Name_Data.String = sirrobin.Name
	sirrobinDB.Name_Data.Valid = true

	sirrobinDB.Description_Data.String = sirrobin.Description
	sirrobinDB.Description_Data.Valid = true

	sirrobinDB.AxisOrign_X_Data.Float64 = sirrobin.AxisOrign_X
	sirrobinDB.AxisOrign_X_Data.Valid = true

	sirrobinDB.AxisOrign_Y_Data.Float64 = sirrobin.AxisOrign_Y
	sirrobinDB.AxisOrign_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_Top_Y_Data.Float64 = sirrobin.VerticalAxis_Top_Y
	sirrobinDB.VerticalAxis_Top_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_Bottom_Y_Data.Float64 = sirrobin.VerticalAxis_Bottom_Y
	sirrobinDB.VerticalAxis_Bottom_Y_Data.Valid = true

	sirrobinDB.VerticalAxis_StrokeWidth_Data.Float64 = sirrobin.VerticalAxis_StrokeWidth
	sirrobinDB.VerticalAxis_StrokeWidth_Data.Valid = true

	sirrobinDB.HorizontalAxis_Right_X_Data.Float64 = sirrobin.HorizontalAxis_Right_X
	sirrobinDB.HorizontalAxis_Right_X_Data.Valid = true

	sirrobinDB.Start_Data.Time = sirrobin.Start
	sirrobinDB.Start_Data.Valid = true

	sirrobinDB.End_Data.Time = sirrobin.End
	sirrobinDB.End_Data.Valid = true
}

// CopyBasicFieldsToSirRobin
func (sirrobinDB *SirRobinDB) CopyBasicFieldsToSirRobin(sirrobin *models.SirRobin) {
	// insertion point for checkout of basic fields (back repo to stage)
	sirrobin.Name = sirrobinDB.Name_Data.String
	sirrobin.Description = sirrobinDB.Description_Data.String
	sirrobin.AxisOrign_X = sirrobinDB.AxisOrign_X_Data.Float64
	sirrobin.AxisOrign_Y = sirrobinDB.AxisOrign_Y_Data.Float64
	sirrobin.VerticalAxis_Top_Y = sirrobinDB.VerticalAxis_Top_Y_Data.Float64
	sirrobin.VerticalAxis_Bottom_Y = sirrobinDB.VerticalAxis_Bottom_Y_Data.Float64
	sirrobin.VerticalAxis_StrokeWidth = sirrobinDB.VerticalAxis_StrokeWidth_Data.Float64
	sirrobin.HorizontalAxis_Right_X = sirrobinDB.HorizontalAxis_Right_X_Data.Float64
	sirrobin.Start = sirrobinDB.Start_Data.Time
	sirrobin.End = sirrobinDB.End_Data.Time
}

// CopyBasicFieldsToSirRobin_WOP
func (sirrobinDB *SirRobinDB) CopyBasicFieldsToSirRobin_WOP(sirrobin *models.SirRobin_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	sirrobin.Name = sirrobinDB.Name_Data.String
	sirrobin.Description = sirrobinDB.Description_Data.String
	sirrobin.AxisOrign_X = sirrobinDB.AxisOrign_X_Data.Float64
	sirrobin.AxisOrign_Y = sirrobinDB.AxisOrign_Y_Data.Float64
	sirrobin.VerticalAxis_Top_Y = sirrobinDB.VerticalAxis_Top_Y_Data.Float64
	sirrobin.VerticalAxis_Bottom_Y = sirrobinDB.VerticalAxis_Bottom_Y_Data.Float64
	sirrobin.VerticalAxis_StrokeWidth = sirrobinDB.VerticalAxis_StrokeWidth_Data.Float64
	sirrobin.HorizontalAxis_Right_X = sirrobinDB.HorizontalAxis_Right_X_Data.Float64
	sirrobin.Start = sirrobinDB.Start_Data.Time
	sirrobin.End = sirrobinDB.End_Data.Time
}

// CopyBasicFieldsToSirRobinWOP
func (sirrobinDB *SirRobinDB) CopyBasicFieldsToSirRobinWOP(sirrobin *SirRobinWOP) {
	sirrobin.ID = int(sirrobinDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	sirrobin.Name = sirrobinDB.Name_Data.String
	sirrobin.Description = sirrobinDB.Description_Data.String
	sirrobin.AxisOrign_X = sirrobinDB.AxisOrign_X_Data.Float64
	sirrobin.AxisOrign_Y = sirrobinDB.AxisOrign_Y_Data.Float64
	sirrobin.VerticalAxis_Top_Y = sirrobinDB.VerticalAxis_Top_Y_Data.Float64
	sirrobin.VerticalAxis_Bottom_Y = sirrobinDB.VerticalAxis_Bottom_Y_Data.Float64
	sirrobin.VerticalAxis_StrokeWidth = sirrobinDB.VerticalAxis_StrokeWidth_Data.Float64
	sirrobin.HorizontalAxis_Right_X = sirrobinDB.HorizontalAxis_Right_X_Data.Float64
	sirrobin.Start = sirrobinDB.Start_Data.Time
	sirrobin.End = sirrobinDB.End_Data.Time
}

// Backup generates a json file from a slice of all SirRobinDB instances in the backrepo
func (backRepoSirRobin *BackRepoSirRobinStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SirRobinDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SirRobinDB, 0)
	for _, sirrobinDB := range backRepoSirRobin.Map_SirRobinDBID_SirRobinDB {
		forBackup = append(forBackup, sirrobinDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SirRobin ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SirRobin file", err.Error())
	}
}

// Backup generates a json file from a slice of all SirRobinDB instances in the backrepo
func (backRepoSirRobin *BackRepoSirRobinStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SirRobinDB, 0)
	for _, sirrobinDB := range backRepoSirRobin.Map_SirRobinDBID_SirRobinDB {
		forBackup = append(forBackup, sirrobinDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SirRobin")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SirRobin_Fields, -1)
	for _, sirrobinDB := range forBackup {

		var sirrobinWOP SirRobinWOP
		sirrobinDB.CopyBasicFieldsToSirRobinWOP(&sirrobinWOP)

		row := sh.AddRow()
		row.WriteStruct(&sirrobinWOP, -1)
	}
}

// RestoreXL from the "SirRobin" sheet all SirRobinDB instances
func (backRepoSirRobin *BackRepoSirRobinStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSirRobinid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SirRobin"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSirRobin.rowVisitorSirRobin)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSirRobin *BackRepoSirRobinStruct) rowVisitorSirRobin(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var sirrobinWOP SirRobinWOP
		row.ReadStruct(&sirrobinWOP)

		// add the unmarshalled struct to the stage
		sirrobinDB := new(SirRobinDB)
		sirrobinDB.CopyBasicFieldsFromSirRobinWOP(&sirrobinWOP)

		sirrobinDB_ID_atBackupTime := sirrobinDB.ID
		sirrobinDB.ID = 0
		query := backRepoSirRobin.db.Create(sirrobinDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSirRobin.Map_SirRobinDBID_SirRobinDB[sirrobinDB.ID] = sirrobinDB
		BackRepoSirRobinid_atBckpTime_newID[sirrobinDB_ID_atBackupTime] = sirrobinDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SirRobinDB.json" in dirPath that stores an array
// of SirRobinDB and stores it in the database
// the map BackRepoSirRobinid_atBckpTime_newID is updated accordingly
func (backRepoSirRobin *BackRepoSirRobinStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSirRobinid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SirRobinDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SirRobin file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SirRobinDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SirRobinDBID_SirRobinDB
	for _, sirrobinDB := range forRestore {

		sirrobinDB_ID_atBackupTime := sirrobinDB.ID
		sirrobinDB.ID = 0
		query := backRepoSirRobin.db.Create(sirrobinDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSirRobin.Map_SirRobinDBID_SirRobinDB[sirrobinDB.ID] = sirrobinDB
		BackRepoSirRobinid_atBckpTime_newID[sirrobinDB_ID_atBackupTime] = sirrobinDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SirRobin file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SirRobin>id_atBckpTime_newID
// to compute new index
func (backRepoSirRobin *BackRepoSirRobinStruct) RestorePhaseTwo() {

	for _, sirrobinDB := range backRepoSirRobin.Map_SirRobinDBID_SirRobinDB {

		// next line of code is to avert unused variable compilation error
		_ = sirrobinDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoSirRobin.db.Model(sirrobinDB).Updates(*sirrobinDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSirRobin.ResetReversePointers commits all staged instances of SirRobin to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSirRobin *BackRepoSirRobinStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, sirrobin := range backRepoSirRobin.Map_SirRobinDBID_SirRobinPtr {
		backRepoSirRobin.ResetReversePointersInstance(backRepo, idx, sirrobin)
	}

	return
}

func (backRepoSirRobin *BackRepoSirRobinStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, sirrobin *models.SirRobin) (Error error) {

	// fetch matching sirrobinDB
	if sirrobinDB, ok := backRepoSirRobin.Map_SirRobinDBID_SirRobinDB[idx]; ok {
		_ = sirrobinDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSirRobinid_atBckpTime_newID map[uint]uint