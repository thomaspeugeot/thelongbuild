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
var __TheNuteShape__dummysDeclaration__ models.TheNuteShape
var __TheNuteShape_time__dummyDeclaration time.Duration

var mutexTheNuteShape sync.Mutex

// An TheNuteShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTheNuteShape updateTheNuteShape deleteTheNuteShape
type TheNuteShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TheNuteShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTheNuteShape updateTheNuteShape
type TheNuteShapeInput struct {
	// The TheNuteShape to submit or modify
	// in: body
	TheNuteShape *orm.TheNuteShapeAPI
}

// GetTheNuteShapes
//
// swagger:route GET /thenuteshapes thenuteshapes getTheNuteShapes
//
// # Get all thenuteshapes
//
// Responses:
// default: genericError
//
//	200: thenuteshapeDBResponse
func (controller *Controller) GetTheNuteShapes(c *gin.Context) {

	// source slice
	var thenuteshapeDBs []orm.TheNuteShapeDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTheNuteShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteShape.GetDB()

	query := db.Find(&thenuteshapeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	thenuteshapeAPIs := make([]orm.TheNuteShapeAPI, 0)

	// for each thenuteshape, update fields from the database nullable fields
	for idx := range thenuteshapeDBs {
		thenuteshapeDB := &thenuteshapeDBs[idx]
		_ = thenuteshapeDB
		var thenuteshapeAPI orm.TheNuteShapeAPI

		// insertion point for updating fields
		thenuteshapeAPI.ID = thenuteshapeDB.ID
		thenuteshapeDB.CopyBasicFieldsToTheNuteShape_WOP(&thenuteshapeAPI.TheNuteShape_WOP)
		thenuteshapeAPI.TheNuteShapePointersEncoding = thenuteshapeDB.TheNuteShapePointersEncoding
		thenuteshapeAPIs = append(thenuteshapeAPIs, thenuteshapeAPI)
	}

	c.JSON(http.StatusOK, thenuteshapeAPIs)
}

// PostTheNuteShape
//
// swagger:route POST /thenuteshapes thenuteshapes postTheNuteShape
//
// Creates a thenuteshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTheNuteShape(c *gin.Context) {

	mutexTheNuteShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTheNuteShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteShape.GetDB()

	// Validate input
	var input orm.TheNuteShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create thenuteshape
	thenuteshapeDB := orm.TheNuteShapeDB{}
	thenuteshapeDB.TheNuteShapePointersEncoding = input.TheNuteShapePointersEncoding
	thenuteshapeDB.CopyBasicFieldsFromTheNuteShape_WOP(&input.TheNuteShape_WOP)

	query := db.Create(&thenuteshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTheNuteShape.CheckoutPhaseOneInstance(&thenuteshapeDB)
	thenuteshape := backRepo.BackRepoTheNuteShape.Map_TheNuteShapeDBID_TheNuteShapePtr[thenuteshapeDB.ID]

	if thenuteshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), thenuteshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, thenuteshapeDB)

	mutexTheNuteShape.Unlock()
}

// GetTheNuteShape
//
// swagger:route GET /thenuteshapes/{ID} thenuteshapes getTheNuteShape
//
// Gets the details for a thenuteshape.
//
// Responses:
// default: genericError
//
//	200: thenuteshapeDBResponse
func (controller *Controller) GetTheNuteShape(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTheNuteShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteShape.GetDB()

	// Get thenuteshapeDB in DB
	var thenuteshapeDB orm.TheNuteShapeDB
	if err := db.First(&thenuteshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var thenuteshapeAPI orm.TheNuteShapeAPI
	thenuteshapeAPI.ID = thenuteshapeDB.ID
	thenuteshapeAPI.TheNuteShapePointersEncoding = thenuteshapeDB.TheNuteShapePointersEncoding
	thenuteshapeDB.CopyBasicFieldsToTheNuteShape_WOP(&thenuteshapeAPI.TheNuteShape_WOP)

	c.JSON(http.StatusOK, thenuteshapeAPI)
}

// UpdateTheNuteShape
//
// swagger:route PATCH /thenuteshapes/{ID} thenuteshapes updateTheNuteShape
//
// # Update a thenuteshape
//
// Responses:
// default: genericError
//
//	200: thenuteshapeDBResponse
func (controller *Controller) UpdateTheNuteShape(c *gin.Context) {

	mutexTheNuteShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTheNuteShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteShape.GetDB()

	// Validate input
	var input orm.TheNuteShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var thenuteshapeDB orm.TheNuteShapeDB

	// fetch the thenuteshape
	query := db.First(&thenuteshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	thenuteshapeDB.CopyBasicFieldsFromTheNuteShape_WOP(&input.TheNuteShape_WOP)
	thenuteshapeDB.TheNuteShapePointersEncoding = input.TheNuteShapePointersEncoding

	query = db.Model(&thenuteshapeDB).Updates(thenuteshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	thenuteshapeNew := new(models.TheNuteShape)
	thenuteshapeDB.CopyBasicFieldsToTheNuteShape(thenuteshapeNew)

	// redeem pointers
	thenuteshapeDB.DecodePointers(backRepo, thenuteshapeNew)

	// get stage instance from DB instance, and call callback function
	thenuteshapeOld := backRepo.BackRepoTheNuteShape.Map_TheNuteShapeDBID_TheNuteShapePtr[thenuteshapeDB.ID]
	if thenuteshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), thenuteshapeOld, thenuteshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the thenuteshapeDB
	c.JSON(http.StatusOK, thenuteshapeDB)

	mutexTheNuteShape.Unlock()
}

// DeleteTheNuteShape
//
// swagger:route DELETE /thenuteshapes/{ID} thenuteshapes deleteTheNuteShape
//
// # Delete a thenuteshape
//
// default: genericError
//
//	200: thenuteshapeDBResponse
func (controller *Controller) DeleteTheNuteShape(c *gin.Context) {

	mutexTheNuteShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTheNuteShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteShape.GetDB()

	// Get model if exist
	var thenuteshapeDB orm.TheNuteShapeDB
	if err := db.First(&thenuteshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&thenuteshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	thenuteshapeDeleted := new(models.TheNuteShape)
	thenuteshapeDB.CopyBasicFieldsToTheNuteShape(thenuteshapeDeleted)

	// get stage instance from DB instance, and call callback function
	thenuteshapeStaged := backRepo.BackRepoTheNuteShape.Map_TheNuteShapeDBID_TheNuteShapePtr[thenuteshapeDB.ID]
	if thenuteshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), thenuteshapeStaged, thenuteshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexTheNuteShape.Unlock()
}
