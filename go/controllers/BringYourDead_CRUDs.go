// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/thomaspeugeot/thelongbuild/go/models"
	"github.com/thomaspeugeot/thelongbuild/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __BringYourDead__dummysDeclaration__ models.BringYourDead
var __BringYourDead_time__dummyDeclaration time.Duration

var mutexBringYourDead sync.Mutex

// An BringYourDeadID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBringYourDead updateBringYourDead deleteBringYourDead
type BringYourDeadID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BringYourDeadInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBringYourDead updateBringYourDead
type BringYourDeadInput struct {
	// The BringYourDead to submit or modify
	// in: body
	BringYourDead *orm.BringYourDeadAPI
}

// GetBringYourDeads
//
// swagger:route GET /bringyourdeads bringyourdeads getBringYourDeads
//
// # Get all bringyourdeads
//
// Responses:
// default: genericError
//
//	200: bringyourdeadDBResponse
func (controller *Controller) GetBringYourDeads(c *gin.Context) {

	// source slice
	var bringyourdeadDBs []orm.BringYourDeadDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBringYourDeads", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBringYourDead.GetDB()

	query := db.Find(&bringyourdeadDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bringyourdeadAPIs := make([]orm.BringYourDeadAPI, 0)

	// for each bringyourdead, update fields from the database nullable fields
	for idx := range bringyourdeadDBs {
		bringyourdeadDB := &bringyourdeadDBs[idx]
		_ = bringyourdeadDB
		var bringyourdeadAPI orm.BringYourDeadAPI

		// insertion point for updating fields
		bringyourdeadAPI.ID = bringyourdeadDB.ID
		bringyourdeadDB.CopyBasicFieldsToBringYourDead_WOP(&bringyourdeadAPI.BringYourDead_WOP)
		bringyourdeadAPI.BringYourDeadPointersEncoding = bringyourdeadDB.BringYourDeadPointersEncoding
		bringyourdeadAPIs = append(bringyourdeadAPIs, bringyourdeadAPI)
	}

	c.JSON(http.StatusOK, bringyourdeadAPIs)
}

// PostBringYourDead
//
// swagger:route POST /bringyourdeads bringyourdeads postBringYourDead
//
// Creates a bringyourdead
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBringYourDead(c *gin.Context) {

	mutexBringYourDead.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBringYourDeads", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBringYourDead.GetDB()

	// Validate input
	var input orm.BringYourDeadAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bringyourdead
	bringyourdeadDB := orm.BringYourDeadDB{}
	bringyourdeadDB.BringYourDeadPointersEncoding = input.BringYourDeadPointersEncoding
	bringyourdeadDB.CopyBasicFieldsFromBringYourDead_WOP(&input.BringYourDead_WOP)

	query := db.Create(&bringyourdeadDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBringYourDead.CheckoutPhaseOneInstance(&bringyourdeadDB)
	bringyourdead := backRepo.BackRepoBringYourDead.Map_BringYourDeadDBID_BringYourDeadPtr[bringyourdeadDB.ID]

	if bringyourdead != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bringyourdead)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bringyourdeadDB)

	mutexBringYourDead.Unlock()
}

// GetBringYourDead
//
// swagger:route GET /bringyourdeads/{ID} bringyourdeads getBringYourDead
//
// Gets the details for a bringyourdead.
//
// Responses:
// default: genericError
//
//	200: bringyourdeadDBResponse
func (controller *Controller) GetBringYourDead(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBringYourDead", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBringYourDead.GetDB()

	// Get bringyourdeadDB in DB
	var bringyourdeadDB orm.BringYourDeadDB
	if err := db.First(&bringyourdeadDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bringyourdeadAPI orm.BringYourDeadAPI
	bringyourdeadAPI.ID = bringyourdeadDB.ID
	bringyourdeadAPI.BringYourDeadPointersEncoding = bringyourdeadDB.BringYourDeadPointersEncoding
	bringyourdeadDB.CopyBasicFieldsToBringYourDead_WOP(&bringyourdeadAPI.BringYourDead_WOP)

	c.JSON(http.StatusOK, bringyourdeadAPI)
}

// UpdateBringYourDead
//
// swagger:route PATCH /bringyourdeads/{ID} bringyourdeads updateBringYourDead
//
// # Update a bringyourdead
//
// Responses:
// default: genericError
//
//	200: bringyourdeadDBResponse
func (controller *Controller) UpdateBringYourDead(c *gin.Context) {

	mutexBringYourDead.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBringYourDead", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBringYourDead.GetDB()

	// Validate input
	var input orm.BringYourDeadAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bringyourdeadDB orm.BringYourDeadDB

	// fetch the bringyourdead
	query := db.First(&bringyourdeadDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bringyourdeadDB.CopyBasicFieldsFromBringYourDead_WOP(&input.BringYourDead_WOP)
	bringyourdeadDB.BringYourDeadPointersEncoding = input.BringYourDeadPointersEncoding

	query = db.Model(&bringyourdeadDB).Updates(bringyourdeadDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bringyourdeadNew := new(models.BringYourDead)
	bringyourdeadDB.CopyBasicFieldsToBringYourDead(bringyourdeadNew)

	// redeem pointers
	bringyourdeadDB.DecodePointers(backRepo, bringyourdeadNew)

	// get stage instance from DB instance, and call callback function
	bringyourdeadOld := backRepo.BackRepoBringYourDead.Map_BringYourDeadDBID_BringYourDeadPtr[bringyourdeadDB.ID]
	if bringyourdeadOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bringyourdeadOld, bringyourdeadNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bringyourdeadDB
	c.JSON(http.StatusOK, bringyourdeadDB)

	mutexBringYourDead.Unlock()
}

// DeleteBringYourDead
//
// swagger:route DELETE /bringyourdeads/{ID} bringyourdeads deleteBringYourDead
//
// # Delete a bringyourdead
//
// default: genericError
//
//	200: bringyourdeadDBResponse
func (controller *Controller) DeleteBringYourDead(c *gin.Context) {

	mutexBringYourDead.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBringYourDead", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBringYourDead.GetDB()

	// Get model if exist
	var bringyourdeadDB orm.BringYourDeadDB
	if err := db.First(&bringyourdeadDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bringyourdeadDB)

	// get an instance (not staged) from DB instance, and call callback function
	bringyourdeadDeleted := new(models.BringYourDead)
	bringyourdeadDB.CopyBasicFieldsToBringYourDead(bringyourdeadDeleted)

	// get stage instance from DB instance, and call callback function
	bringyourdeadStaged := backRepo.BackRepoBringYourDead.Map_BringYourDeadDBID_BringYourDeadPtr[bringyourdeadDB.ID]
	if bringyourdeadStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bringyourdeadStaged, bringyourdeadDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexBringYourDead.Unlock()
}
