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
var __AWitch__dummysDeclaration__ models.AWitch
var __AWitch_time__dummyDeclaration time.Duration

var mutexAWitch sync.Mutex

// An AWitchID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAWitch updateAWitch deleteAWitch
type AWitchID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AWitchInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAWitch updateAWitch
type AWitchInput struct {
	// The AWitch to submit or modify
	// in: body
	AWitch *orm.AWitchAPI
}

// GetAWitchs
//
// swagger:route GET /awitchs awitchs getAWitchs
//
// # Get all awitchs
//
// Responses:
// default: genericError
//
//	200: awitchDBResponse
func (controller *Controller) GetAWitchs(c *gin.Context) {

	// source slice
	var awitchDBs []orm.AWitchDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAWitchs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAWitch.GetDB()

	query := db.Find(&awitchDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	awitchAPIs := make([]orm.AWitchAPI, 0)

	// for each awitch, update fields from the database nullable fields
	for idx := range awitchDBs {
		awitchDB := &awitchDBs[idx]
		_ = awitchDB
		var awitchAPI orm.AWitchAPI

		// insertion point for updating fields
		awitchAPI.ID = awitchDB.ID
		awitchDB.CopyBasicFieldsToAWitch_WOP(&awitchAPI.AWitch_WOP)
		awitchAPI.AWitchPointersEncoding = awitchDB.AWitchPointersEncoding
		awitchAPIs = append(awitchAPIs, awitchAPI)
	}

	c.JSON(http.StatusOK, awitchAPIs)
}

// PostAWitch
//
// swagger:route POST /awitchs awitchs postAWitch
//
// Creates a awitch
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAWitch(c *gin.Context) {

	mutexAWitch.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAWitchs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAWitch.GetDB()

	// Validate input
	var input orm.AWitchAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create awitch
	awitchDB := orm.AWitchDB{}
	awitchDB.AWitchPointersEncoding = input.AWitchPointersEncoding
	awitchDB.CopyBasicFieldsFromAWitch_WOP(&input.AWitch_WOP)

	query := db.Create(&awitchDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAWitch.CheckoutPhaseOneInstance(&awitchDB)
	awitch := backRepo.BackRepoAWitch.Map_AWitchDBID_AWitchPtr[awitchDB.ID]

	if awitch != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), awitch)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, awitchDB)

	mutexAWitch.Unlock()
}

// GetAWitch
//
// swagger:route GET /awitchs/{ID} awitchs getAWitch
//
// Gets the details for a awitch.
//
// Responses:
// default: genericError
//
//	200: awitchDBResponse
func (controller *Controller) GetAWitch(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAWitch", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAWitch.GetDB()

	// Get awitchDB in DB
	var awitchDB orm.AWitchDB
	if err := db.First(&awitchDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var awitchAPI orm.AWitchAPI
	awitchAPI.ID = awitchDB.ID
	awitchAPI.AWitchPointersEncoding = awitchDB.AWitchPointersEncoding
	awitchDB.CopyBasicFieldsToAWitch_WOP(&awitchAPI.AWitch_WOP)

	c.JSON(http.StatusOK, awitchAPI)
}

// UpdateAWitch
//
// swagger:route PATCH /awitchs/{ID} awitchs updateAWitch
//
// # Update a awitch
//
// Responses:
// default: genericError
//
//	200: awitchDBResponse
func (controller *Controller) UpdateAWitch(c *gin.Context) {

	mutexAWitch.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAWitch", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAWitch.GetDB()

	// Validate input
	var input orm.AWitchAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var awitchDB orm.AWitchDB

	// fetch the awitch
	query := db.First(&awitchDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	awitchDB.CopyBasicFieldsFromAWitch_WOP(&input.AWitch_WOP)
	awitchDB.AWitchPointersEncoding = input.AWitchPointersEncoding

	query = db.Model(&awitchDB).Updates(awitchDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	awitchNew := new(models.AWitch)
	awitchDB.CopyBasicFieldsToAWitch(awitchNew)

	// redeem pointers
	awitchDB.DecodePointers(backRepo, awitchNew)

	// get stage instance from DB instance, and call callback function
	awitchOld := backRepo.BackRepoAWitch.Map_AWitchDBID_AWitchPtr[awitchDB.ID]
	if awitchOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), awitchOld, awitchNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the awitchDB
	c.JSON(http.StatusOK, awitchDB)

	mutexAWitch.Unlock()
}

// DeleteAWitch
//
// swagger:route DELETE /awitchs/{ID} awitchs deleteAWitch
//
// # Delete a awitch
//
// default: genericError
//
//	200: awitchDBResponse
func (controller *Controller) DeleteAWitch(c *gin.Context) {

	mutexAWitch.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAWitch", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAWitch.GetDB()

	// Get model if exist
	var awitchDB orm.AWitchDB
	if err := db.First(&awitchDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&awitchDB)

	// get an instance (not staged) from DB instance, and call callback function
	awitchDeleted := new(models.AWitch)
	awitchDB.CopyBasicFieldsToAWitch(awitchDeleted)

	// get stage instance from DB instance, and call callback function
	awitchStaged := backRepo.BackRepoAWitch.Map_AWitchDBID_AWitchPtr[awitchDB.ID]
	if awitchStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), awitchStaged, awitchDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexAWitch.Unlock()
}
