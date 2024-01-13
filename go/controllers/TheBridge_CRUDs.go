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
var __TheBridge__dummysDeclaration__ models.TheBridge
var __TheBridge_time__dummyDeclaration time.Duration

var mutexTheBridge sync.Mutex

// An TheBridgeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTheBridge updateTheBridge deleteTheBridge
type TheBridgeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TheBridgeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTheBridge updateTheBridge
type TheBridgeInput struct {
	// The TheBridge to submit or modify
	// in: body
	TheBridge *orm.TheBridgeAPI
}

// GetTheBridges
//
// swagger:route GET /thebridges thebridges getTheBridges
//
// # Get all thebridges
//
// Responses:
// default: genericError
//
//	200: thebridgeDBResponse
func (controller *Controller) GetTheBridges(c *gin.Context) {

	// source slice
	var thebridgeDBs []orm.TheBridgeDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTheBridges", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheBridge.GetDB()

	query := db.Find(&thebridgeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	thebridgeAPIs := make([]orm.TheBridgeAPI, 0)

	// for each thebridge, update fields from the database nullable fields
	for idx := range thebridgeDBs {
		thebridgeDB := &thebridgeDBs[idx]
		_ = thebridgeDB
		var thebridgeAPI orm.TheBridgeAPI

		// insertion point for updating fields
		thebridgeAPI.ID = thebridgeDB.ID
		thebridgeDB.CopyBasicFieldsToTheBridge_WOP(&thebridgeAPI.TheBridge_WOP)
		thebridgeAPI.TheBridgePointersEncoding = thebridgeDB.TheBridgePointersEncoding
		thebridgeAPIs = append(thebridgeAPIs, thebridgeAPI)
	}

	c.JSON(http.StatusOK, thebridgeAPIs)
}

// PostTheBridge
//
// swagger:route POST /thebridges thebridges postTheBridge
//
// Creates a thebridge
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTheBridge(c *gin.Context) {

	mutexTheBridge.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTheBridges", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheBridge.GetDB()

	// Validate input
	var input orm.TheBridgeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create thebridge
	thebridgeDB := orm.TheBridgeDB{}
	thebridgeDB.TheBridgePointersEncoding = input.TheBridgePointersEncoding
	thebridgeDB.CopyBasicFieldsFromTheBridge_WOP(&input.TheBridge_WOP)

	query := db.Create(&thebridgeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTheBridge.CheckoutPhaseOneInstance(&thebridgeDB)
	thebridge := backRepo.BackRepoTheBridge.Map_TheBridgeDBID_TheBridgePtr[thebridgeDB.ID]

	if thebridge != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), thebridge)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, thebridgeDB)

	mutexTheBridge.Unlock()
}

// GetTheBridge
//
// swagger:route GET /thebridges/{ID} thebridges getTheBridge
//
// Gets the details for a thebridge.
//
// Responses:
// default: genericError
//
//	200: thebridgeDBResponse
func (controller *Controller) GetTheBridge(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTheBridge", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheBridge.GetDB()

	// Get thebridgeDB in DB
	var thebridgeDB orm.TheBridgeDB
	if err := db.First(&thebridgeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var thebridgeAPI orm.TheBridgeAPI
	thebridgeAPI.ID = thebridgeDB.ID
	thebridgeAPI.TheBridgePointersEncoding = thebridgeDB.TheBridgePointersEncoding
	thebridgeDB.CopyBasicFieldsToTheBridge_WOP(&thebridgeAPI.TheBridge_WOP)

	c.JSON(http.StatusOK, thebridgeAPI)
}

// UpdateTheBridge
//
// swagger:route PATCH /thebridges/{ID} thebridges updateTheBridge
//
// # Update a thebridge
//
// Responses:
// default: genericError
//
//	200: thebridgeDBResponse
func (controller *Controller) UpdateTheBridge(c *gin.Context) {

	mutexTheBridge.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTheBridge", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheBridge.GetDB()

	// Validate input
	var input orm.TheBridgeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var thebridgeDB orm.TheBridgeDB

	// fetch the thebridge
	query := db.First(&thebridgeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	thebridgeDB.CopyBasicFieldsFromTheBridge_WOP(&input.TheBridge_WOP)
	thebridgeDB.TheBridgePointersEncoding = input.TheBridgePointersEncoding

	query = db.Model(&thebridgeDB).Updates(thebridgeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	thebridgeNew := new(models.TheBridge)
	thebridgeDB.CopyBasicFieldsToTheBridge(thebridgeNew)

	// redeem pointers
	thebridgeDB.DecodePointers(backRepo, thebridgeNew)

	// get stage instance from DB instance, and call callback function
	thebridgeOld := backRepo.BackRepoTheBridge.Map_TheBridgeDBID_TheBridgePtr[thebridgeDB.ID]
	if thebridgeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), thebridgeOld, thebridgeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the thebridgeDB
	c.JSON(http.StatusOK, thebridgeDB)

	mutexTheBridge.Unlock()
}

// DeleteTheBridge
//
// swagger:route DELETE /thebridges/{ID} thebridges deleteTheBridge
//
// # Delete a thebridge
//
// default: genericError
//
//	200: thebridgeDBResponse
func (controller *Controller) DeleteTheBridge(c *gin.Context) {

	mutexTheBridge.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTheBridge", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheBridge.GetDB()

	// Get model if exist
	var thebridgeDB orm.TheBridgeDB
	if err := db.First(&thebridgeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&thebridgeDB)

	// get an instance (not staged) from DB instance, and call callback function
	thebridgeDeleted := new(models.TheBridge)
	thebridgeDB.CopyBasicFieldsToTheBridge(thebridgeDeleted)

	// get stage instance from DB instance, and call callback function
	thebridgeStaged := backRepo.BackRepoTheBridge.Map_TheBridgeDBID_TheBridgePtr[thebridgeDB.ID]
	if thebridgeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), thebridgeStaged, thebridgeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexTheBridge.Unlock()
}
