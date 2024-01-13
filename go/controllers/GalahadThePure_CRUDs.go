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
var __GalahadThePure__dummysDeclaration__ models.GalahadThePure
var __GalahadThePure_time__dummyDeclaration time.Duration

var mutexGalahadThePure sync.Mutex

// An GalahadThePureID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGalahadThePure updateGalahadThePure deleteGalahadThePure
type GalahadThePureID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GalahadThePureInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGalahadThePure updateGalahadThePure
type GalahadThePureInput struct {
	// The GalahadThePure to submit or modify
	// in: body
	GalahadThePure *orm.GalahadThePureAPI
}

// GetGalahadThePures
//
// swagger:route GET /galahadthepures galahadthepures getGalahadThePures
//
// # Get all galahadthepures
//
// Responses:
// default: genericError
//
//	200: galahadthepureDBResponse
func (controller *Controller) GetGalahadThePures(c *gin.Context) {

	// source slice
	var galahadthepureDBs []orm.GalahadThePureDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGalahadThePures", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGalahadThePure.GetDB()

	query := db.Find(&galahadthepureDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	galahadthepureAPIs := make([]orm.GalahadThePureAPI, 0)

	// for each galahadthepure, update fields from the database nullable fields
	for idx := range galahadthepureDBs {
		galahadthepureDB := &galahadthepureDBs[idx]
		_ = galahadthepureDB
		var galahadthepureAPI orm.GalahadThePureAPI

		// insertion point for updating fields
		galahadthepureAPI.ID = galahadthepureDB.ID
		galahadthepureDB.CopyBasicFieldsToGalahadThePure_WOP(&galahadthepureAPI.GalahadThePure_WOP)
		galahadthepureAPI.GalahadThePurePointersEncoding = galahadthepureDB.GalahadThePurePointersEncoding
		galahadthepureAPIs = append(galahadthepureAPIs, galahadthepureAPI)
	}

	c.JSON(http.StatusOK, galahadthepureAPIs)
}

// PostGalahadThePure
//
// swagger:route POST /galahadthepures galahadthepures postGalahadThePure
//
// Creates a galahadthepure
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGalahadThePure(c *gin.Context) {

	mutexGalahadThePure.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGalahadThePures", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGalahadThePure.GetDB()

	// Validate input
	var input orm.GalahadThePureAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create galahadthepure
	galahadthepureDB := orm.GalahadThePureDB{}
	galahadthepureDB.GalahadThePurePointersEncoding = input.GalahadThePurePointersEncoding
	galahadthepureDB.CopyBasicFieldsFromGalahadThePure_WOP(&input.GalahadThePure_WOP)

	query := db.Create(&galahadthepureDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGalahadThePure.CheckoutPhaseOneInstance(&galahadthepureDB)
	galahadthepure := backRepo.BackRepoGalahadThePure.Map_GalahadThePureDBID_GalahadThePurePtr[galahadthepureDB.ID]

	if galahadthepure != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), galahadthepure)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, galahadthepureDB)

	mutexGalahadThePure.Unlock()
}

// GetGalahadThePure
//
// swagger:route GET /galahadthepures/{ID} galahadthepures getGalahadThePure
//
// Gets the details for a galahadthepure.
//
// Responses:
// default: genericError
//
//	200: galahadthepureDBResponse
func (controller *Controller) GetGalahadThePure(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGalahadThePure", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGalahadThePure.GetDB()

	// Get galahadthepureDB in DB
	var galahadthepureDB orm.GalahadThePureDB
	if err := db.First(&galahadthepureDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var galahadthepureAPI orm.GalahadThePureAPI
	galahadthepureAPI.ID = galahadthepureDB.ID
	galahadthepureAPI.GalahadThePurePointersEncoding = galahadthepureDB.GalahadThePurePointersEncoding
	galahadthepureDB.CopyBasicFieldsToGalahadThePure_WOP(&galahadthepureAPI.GalahadThePure_WOP)

	c.JSON(http.StatusOK, galahadthepureAPI)
}

// UpdateGalahadThePure
//
// swagger:route PATCH /galahadthepures/{ID} galahadthepures updateGalahadThePure
//
// # Update a galahadthepure
//
// Responses:
// default: genericError
//
//	200: galahadthepureDBResponse
func (controller *Controller) UpdateGalahadThePure(c *gin.Context) {

	mutexGalahadThePure.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGalahadThePure", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGalahadThePure.GetDB()

	// Validate input
	var input orm.GalahadThePureAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var galahadthepureDB orm.GalahadThePureDB

	// fetch the galahadthepure
	query := db.First(&galahadthepureDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	galahadthepureDB.CopyBasicFieldsFromGalahadThePure_WOP(&input.GalahadThePure_WOP)
	galahadthepureDB.GalahadThePurePointersEncoding = input.GalahadThePurePointersEncoding

	query = db.Model(&galahadthepureDB).Updates(galahadthepureDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	galahadthepureNew := new(models.GalahadThePure)
	galahadthepureDB.CopyBasicFieldsToGalahadThePure(galahadthepureNew)

	// redeem pointers
	galahadthepureDB.DecodePointers(backRepo, galahadthepureNew)

	// get stage instance from DB instance, and call callback function
	galahadthepureOld := backRepo.BackRepoGalahadThePure.Map_GalahadThePureDBID_GalahadThePurePtr[galahadthepureDB.ID]
	if galahadthepureOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), galahadthepureOld, galahadthepureNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the galahadthepureDB
	c.JSON(http.StatusOK, galahadthepureDB)

	mutexGalahadThePure.Unlock()
}

// DeleteGalahadThePure
//
// swagger:route DELETE /galahadthepures/{ID} galahadthepures deleteGalahadThePure
//
// # Delete a galahadthepure
//
// default: genericError
//
//	200: galahadthepureDBResponse
func (controller *Controller) DeleteGalahadThePure(c *gin.Context) {

	mutexGalahadThePure.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGalahadThePure", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGalahadThePure.GetDB()

	// Get model if exist
	var galahadthepureDB orm.GalahadThePureDB
	if err := db.First(&galahadthepureDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&galahadthepureDB)

	// get an instance (not staged) from DB instance, and call callback function
	galahadthepureDeleted := new(models.GalahadThePure)
	galahadthepureDB.CopyBasicFieldsToGalahadThePure(galahadthepureDeleted)

	// get stage instance from DB instance, and call callback function
	galahadthepureStaged := backRepo.BackRepoGalahadThePure.Map_GalahadThePureDBID_GalahadThePurePtr[galahadthepureDB.ID]
	if galahadthepureStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), galahadthepureStaged, galahadthepureDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexGalahadThePure.Unlock()
}
