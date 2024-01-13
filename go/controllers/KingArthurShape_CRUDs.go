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
var __KingArthurShape__dummysDeclaration__ models.KingArthurShape
var __KingArthurShape_time__dummyDeclaration time.Duration

var mutexKingArthurShape sync.Mutex

// An KingArthurShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getKingArthurShape updateKingArthurShape deleteKingArthurShape
type KingArthurShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// KingArthurShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postKingArthurShape updateKingArthurShape
type KingArthurShapeInput struct {
	// The KingArthurShape to submit or modify
	// in: body
	KingArthurShape *orm.KingArthurShapeAPI
}

// GetKingArthurShapes
//
// swagger:route GET /kingarthurshapes kingarthurshapes getKingArthurShapes
//
// # Get all kingarthurshapes
//
// Responses:
// default: genericError
//
//	200: kingarthurshapeDBResponse
func (controller *Controller) GetKingArthurShapes(c *gin.Context) {

	// source slice
	var kingarthurshapeDBs []orm.KingArthurShapeDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKingArthurShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthurShape.GetDB()

	query := db.Find(&kingarthurshapeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	kingarthurshapeAPIs := make([]orm.KingArthurShapeAPI, 0)

	// for each kingarthurshape, update fields from the database nullable fields
	for idx := range kingarthurshapeDBs {
		kingarthurshapeDB := &kingarthurshapeDBs[idx]
		_ = kingarthurshapeDB
		var kingarthurshapeAPI orm.KingArthurShapeAPI

		// insertion point for updating fields
		kingarthurshapeAPI.ID = kingarthurshapeDB.ID
		kingarthurshapeDB.CopyBasicFieldsToKingArthurShape_WOP(&kingarthurshapeAPI.KingArthurShape_WOP)
		kingarthurshapeAPI.KingArthurShapePointersEncoding = kingarthurshapeDB.KingArthurShapePointersEncoding
		kingarthurshapeAPIs = append(kingarthurshapeAPIs, kingarthurshapeAPI)
	}

	c.JSON(http.StatusOK, kingarthurshapeAPIs)
}

// PostKingArthurShape
//
// swagger:route POST /kingarthurshapes kingarthurshapes postKingArthurShape
//
// Creates a kingarthurshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostKingArthurShape(c *gin.Context) {

	mutexKingArthurShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostKingArthurShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthurShape.GetDB()

	// Validate input
	var input orm.KingArthurShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create kingarthurshape
	kingarthurshapeDB := orm.KingArthurShapeDB{}
	kingarthurshapeDB.KingArthurShapePointersEncoding = input.KingArthurShapePointersEncoding
	kingarthurshapeDB.CopyBasicFieldsFromKingArthurShape_WOP(&input.KingArthurShape_WOP)

	query := db.Create(&kingarthurshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoKingArthurShape.CheckoutPhaseOneInstance(&kingarthurshapeDB)
	kingarthurshape := backRepo.BackRepoKingArthurShape.Map_KingArthurShapeDBID_KingArthurShapePtr[kingarthurshapeDB.ID]

	if kingarthurshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), kingarthurshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, kingarthurshapeDB)

	mutexKingArthurShape.Unlock()
}

// GetKingArthurShape
//
// swagger:route GET /kingarthurshapes/{ID} kingarthurshapes getKingArthurShape
//
// Gets the details for a kingarthurshape.
//
// Responses:
// default: genericError
//
//	200: kingarthurshapeDBResponse
func (controller *Controller) GetKingArthurShape(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKingArthurShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthurShape.GetDB()

	// Get kingarthurshapeDB in DB
	var kingarthurshapeDB orm.KingArthurShapeDB
	if err := db.First(&kingarthurshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var kingarthurshapeAPI orm.KingArthurShapeAPI
	kingarthurshapeAPI.ID = kingarthurshapeDB.ID
	kingarthurshapeAPI.KingArthurShapePointersEncoding = kingarthurshapeDB.KingArthurShapePointersEncoding
	kingarthurshapeDB.CopyBasicFieldsToKingArthurShape_WOP(&kingarthurshapeAPI.KingArthurShape_WOP)

	c.JSON(http.StatusOK, kingarthurshapeAPI)
}

// UpdateKingArthurShape
//
// swagger:route PATCH /kingarthurshapes/{ID} kingarthurshapes updateKingArthurShape
//
// # Update a kingarthurshape
//
// Responses:
// default: genericError
//
//	200: kingarthurshapeDBResponse
func (controller *Controller) UpdateKingArthurShape(c *gin.Context) {

	mutexKingArthurShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateKingArthurShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthurShape.GetDB()

	// Validate input
	var input orm.KingArthurShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var kingarthurshapeDB orm.KingArthurShapeDB

	// fetch the kingarthurshape
	query := db.First(&kingarthurshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	kingarthurshapeDB.CopyBasicFieldsFromKingArthurShape_WOP(&input.KingArthurShape_WOP)
	kingarthurshapeDB.KingArthurShapePointersEncoding = input.KingArthurShapePointersEncoding

	query = db.Model(&kingarthurshapeDB).Updates(kingarthurshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	kingarthurshapeNew := new(models.KingArthurShape)
	kingarthurshapeDB.CopyBasicFieldsToKingArthurShape(kingarthurshapeNew)

	// redeem pointers
	kingarthurshapeDB.DecodePointers(backRepo, kingarthurshapeNew)

	// get stage instance from DB instance, and call callback function
	kingarthurshapeOld := backRepo.BackRepoKingArthurShape.Map_KingArthurShapeDBID_KingArthurShapePtr[kingarthurshapeDB.ID]
	if kingarthurshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), kingarthurshapeOld, kingarthurshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the kingarthurshapeDB
	c.JSON(http.StatusOK, kingarthurshapeDB)

	mutexKingArthurShape.Unlock()
}

// DeleteKingArthurShape
//
// swagger:route DELETE /kingarthurshapes/{ID} kingarthurshapes deleteKingArthurShape
//
// # Delete a kingarthurshape
//
// default: genericError
//
//	200: kingarthurshapeDBResponse
func (controller *Controller) DeleteKingArthurShape(c *gin.Context) {

	mutexKingArthurShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteKingArthurShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthurShape.GetDB()

	// Get model if exist
	var kingarthurshapeDB orm.KingArthurShapeDB
	if err := db.First(&kingarthurshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&kingarthurshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	kingarthurshapeDeleted := new(models.KingArthurShape)
	kingarthurshapeDB.CopyBasicFieldsToKingArthurShape(kingarthurshapeDeleted)

	// get stage instance from DB instance, and call callback function
	kingarthurshapeStaged := backRepo.BackRepoKingArthurShape.Map_KingArthurShapeDBID_KingArthurShapePtr[kingarthurshapeDB.ID]
	if kingarthurshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), kingarthurshapeStaged, kingarthurshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexKingArthurShape.Unlock()
}
