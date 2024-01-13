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
var __BlackKnightShape__dummysDeclaration__ models.BlackKnightShape
var __BlackKnightShape_time__dummyDeclaration time.Duration

var mutexBlackKnightShape sync.Mutex

// An BlackKnightShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBlackKnightShape updateBlackKnightShape deleteBlackKnightShape
type BlackKnightShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BlackKnightShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBlackKnightShape updateBlackKnightShape
type BlackKnightShapeInput struct {
	// The BlackKnightShape to submit or modify
	// in: body
	BlackKnightShape *orm.BlackKnightShapeAPI
}

// GetBlackKnightShapes
//
// swagger:route GET /blackknightshapes blackknightshapes getBlackKnightShapes
//
// # Get all blackknightshapes
//
// Responses:
// default: genericError
//
//	200: blackknightshapeDBResponse
func (controller *Controller) GetBlackKnightShapes(c *gin.Context) {

	// source slice
	var blackknightshapeDBs []orm.BlackKnightShapeDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBlackKnightShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBlackKnightShape.GetDB()

	query := db.Find(&blackknightshapeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	blackknightshapeAPIs := make([]orm.BlackKnightShapeAPI, 0)

	// for each blackknightshape, update fields from the database nullable fields
	for idx := range blackknightshapeDBs {
		blackknightshapeDB := &blackknightshapeDBs[idx]
		_ = blackknightshapeDB
		var blackknightshapeAPI orm.BlackKnightShapeAPI

		// insertion point for updating fields
		blackknightshapeAPI.ID = blackknightshapeDB.ID
		blackknightshapeDB.CopyBasicFieldsToBlackKnightShape_WOP(&blackknightshapeAPI.BlackKnightShape_WOP)
		blackknightshapeAPI.BlackKnightShapePointersEncoding = blackknightshapeDB.BlackKnightShapePointersEncoding
		blackknightshapeAPIs = append(blackknightshapeAPIs, blackknightshapeAPI)
	}

	c.JSON(http.StatusOK, blackknightshapeAPIs)
}

// PostBlackKnightShape
//
// swagger:route POST /blackknightshapes blackknightshapes postBlackKnightShape
//
// Creates a blackknightshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBlackKnightShape(c *gin.Context) {

	mutexBlackKnightShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBlackKnightShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBlackKnightShape.GetDB()

	// Validate input
	var input orm.BlackKnightShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create blackknightshape
	blackknightshapeDB := orm.BlackKnightShapeDB{}
	blackknightshapeDB.BlackKnightShapePointersEncoding = input.BlackKnightShapePointersEncoding
	blackknightshapeDB.CopyBasicFieldsFromBlackKnightShape_WOP(&input.BlackKnightShape_WOP)

	query := db.Create(&blackknightshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBlackKnightShape.CheckoutPhaseOneInstance(&blackknightshapeDB)
	blackknightshape := backRepo.BackRepoBlackKnightShape.Map_BlackKnightShapeDBID_BlackKnightShapePtr[blackknightshapeDB.ID]

	if blackknightshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), blackknightshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, blackknightshapeDB)

	mutexBlackKnightShape.Unlock()
}

// GetBlackKnightShape
//
// swagger:route GET /blackknightshapes/{ID} blackknightshapes getBlackKnightShape
//
// Gets the details for a blackknightshape.
//
// Responses:
// default: genericError
//
//	200: blackknightshapeDBResponse
func (controller *Controller) GetBlackKnightShape(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBlackKnightShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBlackKnightShape.GetDB()

	// Get blackknightshapeDB in DB
	var blackknightshapeDB orm.BlackKnightShapeDB
	if err := db.First(&blackknightshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var blackknightshapeAPI orm.BlackKnightShapeAPI
	blackknightshapeAPI.ID = blackknightshapeDB.ID
	blackknightshapeAPI.BlackKnightShapePointersEncoding = blackknightshapeDB.BlackKnightShapePointersEncoding
	blackknightshapeDB.CopyBasicFieldsToBlackKnightShape_WOP(&blackknightshapeAPI.BlackKnightShape_WOP)

	c.JSON(http.StatusOK, blackknightshapeAPI)
}

// UpdateBlackKnightShape
//
// swagger:route PATCH /blackknightshapes/{ID} blackknightshapes updateBlackKnightShape
//
// # Update a blackknightshape
//
// Responses:
// default: genericError
//
//	200: blackknightshapeDBResponse
func (controller *Controller) UpdateBlackKnightShape(c *gin.Context) {

	mutexBlackKnightShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBlackKnightShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBlackKnightShape.GetDB()

	// Validate input
	var input orm.BlackKnightShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var blackknightshapeDB orm.BlackKnightShapeDB

	// fetch the blackknightshape
	query := db.First(&blackknightshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	blackknightshapeDB.CopyBasicFieldsFromBlackKnightShape_WOP(&input.BlackKnightShape_WOP)
	blackknightshapeDB.BlackKnightShapePointersEncoding = input.BlackKnightShapePointersEncoding

	query = db.Model(&blackknightshapeDB).Updates(blackknightshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	blackknightshapeNew := new(models.BlackKnightShape)
	blackknightshapeDB.CopyBasicFieldsToBlackKnightShape(blackknightshapeNew)

	// redeem pointers
	blackknightshapeDB.DecodePointers(backRepo, blackknightshapeNew)

	// get stage instance from DB instance, and call callback function
	blackknightshapeOld := backRepo.BackRepoBlackKnightShape.Map_BlackKnightShapeDBID_BlackKnightShapePtr[blackknightshapeDB.ID]
	if blackknightshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), blackknightshapeOld, blackknightshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the blackknightshapeDB
	c.JSON(http.StatusOK, blackknightshapeDB)

	mutexBlackKnightShape.Unlock()
}

// DeleteBlackKnightShape
//
// swagger:route DELETE /blackknightshapes/{ID} blackknightshapes deleteBlackKnightShape
//
// # Delete a blackknightshape
//
// default: genericError
//
//	200: blackknightshapeDBResponse
func (controller *Controller) DeleteBlackKnightShape(c *gin.Context) {

	mutexBlackKnightShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBlackKnightShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBlackKnightShape.GetDB()

	// Get model if exist
	var blackknightshapeDB orm.BlackKnightShapeDB
	if err := db.First(&blackknightshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&blackknightshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	blackknightshapeDeleted := new(models.BlackKnightShape)
	blackknightshapeDB.CopyBasicFieldsToBlackKnightShape(blackknightshapeDeleted)

	// get stage instance from DB instance, and call callback function
	blackknightshapeStaged := backRepo.BackRepoBlackKnightShape.Map_BlackKnightShapeDBID_BlackKnightShapePtr[blackknightshapeDB.ID]
	if blackknightshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), blackknightshapeStaged, blackknightshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexBlackKnightShape.Unlock()
}
