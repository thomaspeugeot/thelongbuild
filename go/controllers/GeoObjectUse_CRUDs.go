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
var __GeoObjectUse__dummysDeclaration__ models.GeoObjectUse
var __GeoObjectUse_time__dummyDeclaration time.Duration

var mutexGeoObjectUse sync.Mutex

// An GeoObjectUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGeoObjectUse updateGeoObjectUse deleteGeoObjectUse
type GeoObjectUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GeoObjectUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGeoObjectUse updateGeoObjectUse
type GeoObjectUseInput struct {
	// The GeoObjectUse to submit or modify
	// in: body
	GeoObjectUse *orm.GeoObjectUseAPI
}

// GetGeoObjectUses
//
// swagger:route GET /geoobjectuses geoobjectuses getGeoObjectUses
//
// # Get all geoobjectuses
//
// Responses:
// default: genericError
//
//	200: geoobjectuseDBResponse
func (controller *Controller) GetGeoObjectUses(c *gin.Context) {

	// source slice
	var geoobjectuseDBs []orm.GeoObjectUseDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGeoObjectUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObjectUse.GetDB()

	query := db.Find(&geoobjectuseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	geoobjectuseAPIs := make([]orm.GeoObjectUseAPI, 0)

	// for each geoobjectuse, update fields from the database nullable fields
	for idx := range geoobjectuseDBs {
		geoobjectuseDB := &geoobjectuseDBs[idx]
		_ = geoobjectuseDB
		var geoobjectuseAPI orm.GeoObjectUseAPI

		// insertion point for updating fields
		geoobjectuseAPI.ID = geoobjectuseDB.ID
		geoobjectuseDB.CopyBasicFieldsToGeoObjectUse_WOP(&geoobjectuseAPI.GeoObjectUse_WOP)
		geoobjectuseAPI.GeoObjectUsePointersEncoding = geoobjectuseDB.GeoObjectUsePointersEncoding
		geoobjectuseAPIs = append(geoobjectuseAPIs, geoobjectuseAPI)
	}

	c.JSON(http.StatusOK, geoobjectuseAPIs)
}

// PostGeoObjectUse
//
// swagger:route POST /geoobjectuses geoobjectuses postGeoObjectUse
//
// Creates a geoobjectuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGeoObjectUse(c *gin.Context) {

	mutexGeoObjectUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGeoObjectUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObjectUse.GetDB()

	// Validate input
	var input orm.GeoObjectUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create geoobjectuse
	geoobjectuseDB := orm.GeoObjectUseDB{}
	geoobjectuseDB.GeoObjectUsePointersEncoding = input.GeoObjectUsePointersEncoding
	geoobjectuseDB.CopyBasicFieldsFromGeoObjectUse_WOP(&input.GeoObjectUse_WOP)

	query := db.Create(&geoobjectuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGeoObjectUse.CheckoutPhaseOneInstance(&geoobjectuseDB)
	geoobjectuse := backRepo.BackRepoGeoObjectUse.Map_GeoObjectUseDBID_GeoObjectUsePtr[geoobjectuseDB.ID]

	if geoobjectuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), geoobjectuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, geoobjectuseDB)

	mutexGeoObjectUse.Unlock()
}

// GetGeoObjectUse
//
// swagger:route GET /geoobjectuses/{ID} geoobjectuses getGeoObjectUse
//
// Gets the details for a geoobjectuse.
//
// Responses:
// default: genericError
//
//	200: geoobjectuseDBResponse
func (controller *Controller) GetGeoObjectUse(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGeoObjectUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObjectUse.GetDB()

	// Get geoobjectuseDB in DB
	var geoobjectuseDB orm.GeoObjectUseDB
	if err := db.First(&geoobjectuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var geoobjectuseAPI orm.GeoObjectUseAPI
	geoobjectuseAPI.ID = geoobjectuseDB.ID
	geoobjectuseAPI.GeoObjectUsePointersEncoding = geoobjectuseDB.GeoObjectUsePointersEncoding
	geoobjectuseDB.CopyBasicFieldsToGeoObjectUse_WOP(&geoobjectuseAPI.GeoObjectUse_WOP)

	c.JSON(http.StatusOK, geoobjectuseAPI)
}

// UpdateGeoObjectUse
//
// swagger:route PATCH /geoobjectuses/{ID} geoobjectuses updateGeoObjectUse
//
// # Update a geoobjectuse
//
// Responses:
// default: genericError
//
//	200: geoobjectuseDBResponse
func (controller *Controller) UpdateGeoObjectUse(c *gin.Context) {

	mutexGeoObjectUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGeoObjectUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObjectUse.GetDB()

	// Validate input
	var input orm.GeoObjectUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var geoobjectuseDB orm.GeoObjectUseDB

	// fetch the geoobjectuse
	query := db.First(&geoobjectuseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	geoobjectuseDB.CopyBasicFieldsFromGeoObjectUse_WOP(&input.GeoObjectUse_WOP)
	geoobjectuseDB.GeoObjectUsePointersEncoding = input.GeoObjectUsePointersEncoding

	query = db.Model(&geoobjectuseDB).Updates(geoobjectuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	geoobjectuseNew := new(models.GeoObjectUse)
	geoobjectuseDB.CopyBasicFieldsToGeoObjectUse(geoobjectuseNew)

	// redeem pointers
	geoobjectuseDB.DecodePointers(backRepo, geoobjectuseNew)

	// get stage instance from DB instance, and call callback function
	geoobjectuseOld := backRepo.BackRepoGeoObjectUse.Map_GeoObjectUseDBID_GeoObjectUsePtr[geoobjectuseDB.ID]
	if geoobjectuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), geoobjectuseOld, geoobjectuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the geoobjectuseDB
	c.JSON(http.StatusOK, geoobjectuseDB)

	mutexGeoObjectUse.Unlock()
}

// DeleteGeoObjectUse
//
// swagger:route DELETE /geoobjectuses/{ID} geoobjectuses deleteGeoObjectUse
//
// # Delete a geoobjectuse
//
// default: genericError
//
//	200: geoobjectuseDBResponse
func (controller *Controller) DeleteGeoObjectUse(c *gin.Context) {

	mutexGeoObjectUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGeoObjectUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGeoObjectUse.GetDB()

	// Get model if exist
	var geoobjectuseDB orm.GeoObjectUseDB
	if err := db.First(&geoobjectuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&geoobjectuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	geoobjectuseDeleted := new(models.GeoObjectUse)
	geoobjectuseDB.CopyBasicFieldsToGeoObjectUse(geoobjectuseDeleted)

	// get stage instance from DB instance, and call callback function
	geoobjectuseStaged := backRepo.BackRepoGeoObjectUse.Map_GeoObjectUseDBID_GeoObjectUsePtr[geoobjectuseDB.ID]
	if geoobjectuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), geoobjectuseStaged, geoobjectuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexGeoObjectUse.Unlock()
}
