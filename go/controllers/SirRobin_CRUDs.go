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
var __SirRobin__dummysDeclaration__ models.SirRobin
var __SirRobin_time__dummyDeclaration time.Duration

var mutexSirRobin sync.Mutex

// An SirRobinID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSirRobin updateSirRobin deleteSirRobin
type SirRobinID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SirRobinInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSirRobin updateSirRobin
type SirRobinInput struct {
	// The SirRobin to submit or modify
	// in: body
	SirRobin *orm.SirRobinAPI
}

// GetSirRobins
//
// swagger:route GET /sirrobins sirrobins getSirRobins
//
// # Get all sirrobins
//
// Responses:
// default: genericError
//
//	200: sirrobinDBResponse
func (controller *Controller) GetSirRobins(c *gin.Context) {

	// source slice
	var sirrobinDBs []orm.SirRobinDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSirRobins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSirRobin.GetDB()

	query := db.Find(&sirrobinDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	sirrobinAPIs := make([]orm.SirRobinAPI, 0)

	// for each sirrobin, update fields from the database nullable fields
	for idx := range sirrobinDBs {
		sirrobinDB := &sirrobinDBs[idx]
		_ = sirrobinDB
		var sirrobinAPI orm.SirRobinAPI

		// insertion point for updating fields
		sirrobinAPI.ID = sirrobinDB.ID
		sirrobinDB.CopyBasicFieldsToSirRobin_WOP(&sirrobinAPI.SirRobin_WOP)
		sirrobinAPI.SirRobinPointersEncoding = sirrobinDB.SirRobinPointersEncoding
		sirrobinAPIs = append(sirrobinAPIs, sirrobinAPI)
	}

	c.JSON(http.StatusOK, sirrobinAPIs)
}

// PostSirRobin
//
// swagger:route POST /sirrobins sirrobins postSirRobin
//
// Creates a sirrobin
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSirRobin(c *gin.Context) {

	mutexSirRobin.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSirRobins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSirRobin.GetDB()

	// Validate input
	var input orm.SirRobinAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create sirrobin
	sirrobinDB := orm.SirRobinDB{}
	sirrobinDB.SirRobinPointersEncoding = input.SirRobinPointersEncoding
	sirrobinDB.CopyBasicFieldsFromSirRobin_WOP(&input.SirRobin_WOP)

	query := db.Create(&sirrobinDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSirRobin.CheckoutPhaseOneInstance(&sirrobinDB)
	sirrobin := backRepo.BackRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[sirrobinDB.ID]

	if sirrobin != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), sirrobin)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, sirrobinDB)

	mutexSirRobin.Unlock()
}

// GetSirRobin
//
// swagger:route GET /sirrobins/{ID} sirrobins getSirRobin
//
// Gets the details for a sirrobin.
//
// Responses:
// default: genericError
//
//	200: sirrobinDBResponse
func (controller *Controller) GetSirRobin(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSirRobin", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSirRobin.GetDB()

	// Get sirrobinDB in DB
	var sirrobinDB orm.SirRobinDB
	if err := db.First(&sirrobinDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var sirrobinAPI orm.SirRobinAPI
	sirrobinAPI.ID = sirrobinDB.ID
	sirrobinAPI.SirRobinPointersEncoding = sirrobinDB.SirRobinPointersEncoding
	sirrobinDB.CopyBasicFieldsToSirRobin_WOP(&sirrobinAPI.SirRobin_WOP)

	c.JSON(http.StatusOK, sirrobinAPI)
}

// UpdateSirRobin
//
// swagger:route PATCH /sirrobins/{ID} sirrobins updateSirRobin
//
// # Update a sirrobin
//
// Responses:
// default: genericError
//
//	200: sirrobinDBResponse
func (controller *Controller) UpdateSirRobin(c *gin.Context) {

	mutexSirRobin.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSirRobin", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSirRobin.GetDB()

	// Validate input
	var input orm.SirRobinAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var sirrobinDB orm.SirRobinDB

	// fetch the sirrobin
	query := db.First(&sirrobinDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	sirrobinDB.CopyBasicFieldsFromSirRobin_WOP(&input.SirRobin_WOP)
	sirrobinDB.SirRobinPointersEncoding = input.SirRobinPointersEncoding

	query = db.Model(&sirrobinDB).Updates(sirrobinDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	sirrobinNew := new(models.SirRobin)
	sirrobinDB.CopyBasicFieldsToSirRobin(sirrobinNew)

	// redeem pointers
	sirrobinDB.DecodePointers(backRepo, sirrobinNew)

	// get stage instance from DB instance, and call callback function
	sirrobinOld := backRepo.BackRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[sirrobinDB.ID]
	if sirrobinOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), sirrobinOld, sirrobinNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the sirrobinDB
	c.JSON(http.StatusOK, sirrobinDB)

	mutexSirRobin.Unlock()
}

// DeleteSirRobin
//
// swagger:route DELETE /sirrobins/{ID} sirrobins deleteSirRobin
//
// # Delete a sirrobin
//
// default: genericError
//
//	200: sirrobinDBResponse
func (controller *Controller) DeleteSirRobin(c *gin.Context) {

	mutexSirRobin.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSirRobin", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSirRobin.GetDB()

	// Get model if exist
	var sirrobinDB orm.SirRobinDB
	if err := db.First(&sirrobinDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&sirrobinDB)

	// get an instance (not staged) from DB instance, and call callback function
	sirrobinDeleted := new(models.SirRobin)
	sirrobinDB.CopyBasicFieldsToSirRobin(sirrobinDeleted)

	// get stage instance from DB instance, and call callback function
	sirrobinStaged := backRepo.BackRepoSirRobin.Map_SirRobinDBID_SirRobinPtr[sirrobinDB.ID]
	if sirrobinStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), sirrobinStaged, sirrobinDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexSirRobin.Unlock()
}
