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
var __GeoObject__dummysDeclaration__ models.GeoObject
var __GeoObject_time__dummyDeclaration time.Duration

var mutexGeoObject sync.Mutex

// An GeoObjectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGeoObject updateGeoObject deleteGeoObject
type GeoObjectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GeoObjectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGeoObject updateGeoObject
type GeoObjectInput struct {
	// The GeoObject to submit or modify
	// in: body
	GeoObject *orm.GeoObjectAPI
}

// GetGeoObjects
//
// swagger:route GET /geoobjects geoobjects getGeoObjects
//
// # Get all geoobjects
//
// Responses:
// default: genericError
//
//	200: geoobjectDBResponse
func (controller *Controller) GetGeoObjects(c *gin.Context) {

	// source slice
	var geoobjectDBs []orm.GeoObjectDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGeoObjects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObject.GetDB()

	query := db.Find(&geoobjectDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	geoobjectAPIs := make([]orm.GeoObjectAPI, 0)

	// for each geoobject, update fields from the database nullable fields
	for idx := range geoobjectDBs {
		geoobjectDB := &geoobjectDBs[idx]
		_ = geoobjectDB
		var geoobjectAPI orm.GeoObjectAPI

		// insertion point for updating fields
		geoobjectAPI.ID = geoobjectDB.ID
		geoobjectDB.CopyBasicFieldsToGeoObject_WOP(&geoobjectAPI.GeoObject_WOP)
		geoobjectAPI.GeoObjectPointersEncoding = geoobjectDB.GeoObjectPointersEncoding
		geoobjectAPIs = append(geoobjectAPIs, geoobjectAPI)
	}

	c.JSON(http.StatusOK, geoobjectAPIs)
}

// PostGeoObject
//
// swagger:route POST /geoobjects geoobjects postGeoObject
//
// Creates a geoobject
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGeoObject(c *gin.Context) {

	mutexGeoObject.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGeoObjects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObject.GetDB()

	// Validate input
	var input orm.GeoObjectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create geoobject
	geoobjectDB := orm.GeoObjectDB{}
	geoobjectDB.GeoObjectPointersEncoding = input.GeoObjectPointersEncoding
	geoobjectDB.CopyBasicFieldsFromGeoObject_WOP(&input.GeoObject_WOP)

	query := db.Create(&geoobjectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGeoObject.CheckoutPhaseOneInstance(&geoobjectDB)
	geoobject := backRepo.BackRepoGeoObject.Map_GeoObjectDBID_GeoObjectPtr[geoobjectDB.ID]

	if geoobject != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), geoobject)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, geoobjectDB)

	mutexGeoObject.Unlock()
}

// GetGeoObject
//
// swagger:route GET /geoobjects/{ID} geoobjects getGeoObject
//
// Gets the details for a geoobject.
//
// Responses:
// default: genericError
//
//	200: geoobjectDBResponse
func (controller *Controller) GetGeoObject(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGeoObject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObject.GetDB()

	// Get geoobjectDB in DB
	var geoobjectDB orm.GeoObjectDB
	if err := db.First(&geoobjectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var geoobjectAPI orm.GeoObjectAPI
	geoobjectAPI.ID = geoobjectDB.ID
	geoobjectAPI.GeoObjectPointersEncoding = geoobjectDB.GeoObjectPointersEncoding
	geoobjectDB.CopyBasicFieldsToGeoObject_WOP(&geoobjectAPI.GeoObject_WOP)

	c.JSON(http.StatusOK, geoobjectAPI)
}

// UpdateGeoObject
//
// swagger:route PATCH /geoobjects/{ID} geoobjects updateGeoObject
//
// # Update a geoobject
//
// Responses:
// default: genericError
//
//	200: geoobjectDBResponse
func (controller *Controller) UpdateGeoObject(c *gin.Context) {

	mutexGeoObject.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGeoObject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObject.GetDB()

	// Validate input
	var input orm.GeoObjectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var geoobjectDB orm.GeoObjectDB

	// fetch the geoobject
	query := db.First(&geoobjectDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	geoobjectDB.CopyBasicFieldsFromGeoObject_WOP(&input.GeoObject_WOP)
	geoobjectDB.GeoObjectPointersEncoding = input.GeoObjectPointersEncoding

	query = db.Model(&geoobjectDB).Updates(geoobjectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	geoobjectNew := new(models.GeoObject)
	geoobjectDB.CopyBasicFieldsToGeoObject(geoobjectNew)

	// redeem pointers
	geoobjectDB.DecodePointers(backRepo, geoobjectNew)

	// get stage instance from DB instance, and call callback function
	geoobjectOld := backRepo.BackRepoGeoObject.Map_GeoObjectDBID_GeoObjectPtr[geoobjectDB.ID]
	if geoobjectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), geoobjectOld, geoobjectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the geoobjectDB
	c.JSON(http.StatusOK, geoobjectDB)

	mutexGeoObject.Unlock()
}

// DeleteGeoObject
//
// swagger:route DELETE /geoobjects/{ID} geoobjects deleteGeoObject
//
// # Delete a geoobject
//
// default: genericError
//
//	200: geoobjectDBResponse
func (controller *Controller) DeleteGeoObject(c *gin.Context) {

	mutexGeoObject.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGeoObject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObject.GetDB()

	// Get model if exist
	var geoobjectDB orm.GeoObjectDB
	if err := db.First(&geoobjectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&geoobjectDB)

	// get an instance (not staged) from DB instance, and call callback function
	geoobjectDeleted := new(models.GeoObject)
	geoobjectDB.CopyBasicFieldsToGeoObject(geoobjectDeleted)

	// get stage instance from DB instance, and call callback function
	geoobjectStaged := backRepo.BackRepoGeoObject.Map_GeoObjectDBID_GeoObjectPtr[geoobjectDB.ID]
	if geoobjectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), geoobjectStaged, geoobjectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexGeoObject.Unlock()
}
