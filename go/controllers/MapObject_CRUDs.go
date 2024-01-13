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
var __MapObject__dummysDeclaration__ models.MapObject
var __MapObject_time__dummyDeclaration time.Duration

var mutexMapObject sync.Mutex

// An MapObjectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMapObject updateMapObject deleteMapObject
type MapObjectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MapObjectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMapObject updateMapObject
type MapObjectInput struct {
	// The MapObject to submit or modify
	// in: body
	MapObject *orm.MapObjectAPI
}

// GetMapObjects
//
// swagger:route GET /mapobjects mapobjects getMapObjects
//
// # Get all mapobjects
//
// Responses:
// default: genericError
//
//	200: mapobjectDBResponse
func (controller *Controller) GetMapObjects(c *gin.Context) {

	// source slice
	var mapobjectDBs []orm.MapObjectDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMapObjects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObject.GetDB()

	query := db.Find(&mapobjectDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	mapobjectAPIs := make([]orm.MapObjectAPI, 0)

	// for each mapobject, update fields from the database nullable fields
	for idx := range mapobjectDBs {
		mapobjectDB := &mapobjectDBs[idx]
		_ = mapobjectDB
		var mapobjectAPI orm.MapObjectAPI

		// insertion point for updating fields
		mapobjectAPI.ID = mapobjectDB.ID
		mapobjectDB.CopyBasicFieldsToMapObject_WOP(&mapobjectAPI.MapObject_WOP)
		mapobjectAPI.MapObjectPointersEncoding = mapobjectDB.MapObjectPointersEncoding
		mapobjectAPIs = append(mapobjectAPIs, mapobjectAPI)
	}

	c.JSON(http.StatusOK, mapobjectAPIs)
}

// PostMapObject
//
// swagger:route POST /mapobjects mapobjects postMapObject
//
// Creates a mapobject
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMapObject(c *gin.Context) {

	mutexMapObject.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMapObjects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObject.GetDB()

	// Validate input
	var input orm.MapObjectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create mapobject
	mapobjectDB := orm.MapObjectDB{}
	mapobjectDB.MapObjectPointersEncoding = input.MapObjectPointersEncoding
	mapobjectDB.CopyBasicFieldsFromMapObject_WOP(&input.MapObject_WOP)

	query := db.Create(&mapobjectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMapObject.CheckoutPhaseOneInstance(&mapobjectDB)
	mapobject := backRepo.BackRepoMapObject.Map_MapObjectDBID_MapObjectPtr[mapobjectDB.ID]

	if mapobject != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), mapobject)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, mapobjectDB)

	mutexMapObject.Unlock()
}

// GetMapObject
//
// swagger:route GET /mapobjects/{ID} mapobjects getMapObject
//
// Gets the details for a mapobject.
//
// Responses:
// default: genericError
//
//	200: mapobjectDBResponse
func (controller *Controller) GetMapObject(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMapObject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObject.GetDB()

	// Get mapobjectDB in DB
	var mapobjectDB orm.MapObjectDB
	if err := db.First(&mapobjectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var mapobjectAPI orm.MapObjectAPI
	mapobjectAPI.ID = mapobjectDB.ID
	mapobjectAPI.MapObjectPointersEncoding = mapobjectDB.MapObjectPointersEncoding
	mapobjectDB.CopyBasicFieldsToMapObject_WOP(&mapobjectAPI.MapObject_WOP)

	c.JSON(http.StatusOK, mapobjectAPI)
}

// UpdateMapObject
//
// swagger:route PATCH /mapobjects/{ID} mapobjects updateMapObject
//
// # Update a mapobject
//
// Responses:
// default: genericError
//
//	200: mapobjectDBResponse
func (controller *Controller) UpdateMapObject(c *gin.Context) {

	mutexMapObject.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMapObject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObject.GetDB()

	// Validate input
	var input orm.MapObjectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var mapobjectDB orm.MapObjectDB

	// fetch the mapobject
	query := db.First(&mapobjectDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	mapobjectDB.CopyBasicFieldsFromMapObject_WOP(&input.MapObject_WOP)
	mapobjectDB.MapObjectPointersEncoding = input.MapObjectPointersEncoding

	query = db.Model(&mapobjectDB).Updates(mapobjectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	mapobjectNew := new(models.MapObject)
	mapobjectDB.CopyBasicFieldsToMapObject(mapobjectNew)

	// redeem pointers
	mapobjectDB.DecodePointers(backRepo, mapobjectNew)

	// get stage instance from DB instance, and call callback function
	mapobjectOld := backRepo.BackRepoMapObject.Map_MapObjectDBID_MapObjectPtr[mapobjectDB.ID]
	if mapobjectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), mapobjectOld, mapobjectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the mapobjectDB
	c.JSON(http.StatusOK, mapobjectDB)

	mutexMapObject.Unlock()
}

// DeleteMapObject
//
// swagger:route DELETE /mapobjects/{ID} mapobjects deleteMapObject
//
// # Delete a mapobject
//
// default: genericError
//
//	200: mapobjectDBResponse
func (controller *Controller) DeleteMapObject(c *gin.Context) {

	mutexMapObject.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMapObject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMapObject.GetDB()

	// Get model if exist
	var mapobjectDB orm.MapObjectDB
	if err := db.First(&mapobjectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&mapobjectDB)

	// get an instance (not staged) from DB instance, and call callback function
	mapobjectDeleted := new(models.MapObject)
	mapobjectDB.CopyBasicFieldsToMapObject(mapobjectDeleted)

	// get stage instance from DB instance, and call callback function
	mapobjectStaged := backRepo.BackRepoMapObject.Map_MapObjectDBID_MapObjectPtr[mapobjectDB.ID]
	if mapobjectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), mapobjectStaged, mapobjectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexMapObject.Unlock()
}
