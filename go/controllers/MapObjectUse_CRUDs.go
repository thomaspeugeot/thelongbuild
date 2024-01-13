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
var __MapObjectUse__dummysDeclaration__ models.MapObjectUse
var __MapObjectUse_time__dummyDeclaration time.Duration

var mutexMapObjectUse sync.Mutex

// An MapObjectUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMapObjectUse updateMapObjectUse deleteMapObjectUse
type MapObjectUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MapObjectUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMapObjectUse updateMapObjectUse
type MapObjectUseInput struct {
	// The MapObjectUse to submit or modify
	// in: body
	MapObjectUse *orm.MapObjectUseAPI
}

// GetMapObjectUses
//
// swagger:route GET /mapobjectuses mapobjectuses getMapObjectUses
//
// # Get all mapobjectuses
//
// Responses:
// default: genericError
//
//	200: mapobjectuseDBResponse
func (controller *Controller) GetMapObjectUses(c *gin.Context) {

	// source slice
	var mapobjectuseDBs []orm.MapObjectUseDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMapObjectUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObjectUse.GetDB()

	query := db.Find(&mapobjectuseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	mapobjectuseAPIs := make([]orm.MapObjectUseAPI, 0)

	// for each mapobjectuse, update fields from the database nullable fields
	for idx := range mapobjectuseDBs {
		mapobjectuseDB := &mapobjectuseDBs[idx]
		_ = mapobjectuseDB
		var mapobjectuseAPI orm.MapObjectUseAPI

		// insertion point for updating fields
		mapobjectuseAPI.ID = mapobjectuseDB.ID
		mapobjectuseDB.CopyBasicFieldsToMapObjectUse_WOP(&mapobjectuseAPI.MapObjectUse_WOP)
		mapobjectuseAPI.MapObjectUsePointersEncoding = mapobjectuseDB.MapObjectUsePointersEncoding
		mapobjectuseAPIs = append(mapobjectuseAPIs, mapobjectuseAPI)
	}

	c.JSON(http.StatusOK, mapobjectuseAPIs)
}

// PostMapObjectUse
//
// swagger:route POST /mapobjectuses mapobjectuses postMapObjectUse
//
// Creates a mapobjectuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMapObjectUse(c *gin.Context) {

	mutexMapObjectUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMapObjectUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObjectUse.GetDB()

	// Validate input
	var input orm.MapObjectUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create mapobjectuse
	mapobjectuseDB := orm.MapObjectUseDB{}
	mapobjectuseDB.MapObjectUsePointersEncoding = input.MapObjectUsePointersEncoding
	mapobjectuseDB.CopyBasicFieldsFromMapObjectUse_WOP(&input.MapObjectUse_WOP)

	query := db.Create(&mapobjectuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMapObjectUse.CheckoutPhaseOneInstance(&mapobjectuseDB)
	mapobjectuse := backRepo.BackRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[mapobjectuseDB.ID]

	if mapobjectuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), mapobjectuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, mapobjectuseDB)

	mutexMapObjectUse.Unlock()
}

// GetMapObjectUse
//
// swagger:route GET /mapobjectuses/{ID} mapobjectuses getMapObjectUse
//
// Gets the details for a mapobjectuse.
//
// Responses:
// default: genericError
//
//	200: mapobjectuseDBResponse
func (controller *Controller) GetMapObjectUse(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMapObjectUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObjectUse.GetDB()

	// Get mapobjectuseDB in DB
	var mapobjectuseDB orm.MapObjectUseDB
	if err := db.First(&mapobjectuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var mapobjectuseAPI orm.MapObjectUseAPI
	mapobjectuseAPI.ID = mapobjectuseDB.ID
	mapobjectuseAPI.MapObjectUsePointersEncoding = mapobjectuseDB.MapObjectUsePointersEncoding
	mapobjectuseDB.CopyBasicFieldsToMapObjectUse_WOP(&mapobjectuseAPI.MapObjectUse_WOP)

	c.JSON(http.StatusOK, mapobjectuseAPI)
}

// UpdateMapObjectUse
//
// swagger:route PATCH /mapobjectuses/{ID} mapobjectuses updateMapObjectUse
//
// # Update a mapobjectuse
//
// Responses:
// default: genericError
//
//	200: mapobjectuseDBResponse
func (controller *Controller) UpdateMapObjectUse(c *gin.Context) {

	mutexMapObjectUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMapObjectUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObjectUse.GetDB()

	// Validate input
	var input orm.MapObjectUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var mapobjectuseDB orm.MapObjectUseDB

	// fetch the mapobjectuse
	query := db.First(&mapobjectuseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	mapobjectuseDB.CopyBasicFieldsFromMapObjectUse_WOP(&input.MapObjectUse_WOP)
	mapobjectuseDB.MapObjectUsePointersEncoding = input.MapObjectUsePointersEncoding

	query = db.Model(&mapobjectuseDB).Updates(mapobjectuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	mapobjectuseNew := new(models.MapObjectUse)
	mapobjectuseDB.CopyBasicFieldsToMapObjectUse(mapobjectuseNew)

	// redeem pointers
	mapobjectuseDB.DecodePointers(backRepo, mapobjectuseNew)

	// get stage instance from DB instance, and call callback function
	mapobjectuseOld := backRepo.BackRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[mapobjectuseDB.ID]
	if mapobjectuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), mapobjectuseOld, mapobjectuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the mapobjectuseDB
	c.JSON(http.StatusOK, mapobjectuseDB)

	mutexMapObjectUse.Unlock()
}

// DeleteMapObjectUse
//
// swagger:route DELETE /mapobjectuses/{ID} mapobjectuses deleteMapObjectUse
//
// # Delete a mapobjectuse
//
// default: genericError
//
//	200: mapobjectuseDBResponse
func (controller *Controller) DeleteMapObjectUse(c *gin.Context) {

	mutexMapObjectUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMapObjectUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObjectUse.GetDB()

	// Get model if exist
	var mapobjectuseDB orm.MapObjectUseDB
	if err := db.First(&mapobjectuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&mapobjectuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	mapobjectuseDeleted := new(models.MapObjectUse)
	mapobjectuseDB.CopyBasicFieldsToMapObjectUse(mapobjectuseDeleted)

	// get stage instance from DB instance, and call callback function
	mapobjectuseStaged := backRepo.BackRepoMapObjectUse.Map_MapObjectUseDBID_MapObjectUsePtr[mapobjectuseDB.ID]
	if mapobjectuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), mapobjectuseStaged, mapobjectuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexMapObjectUse.Unlock()
}
