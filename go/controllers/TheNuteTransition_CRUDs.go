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
var __TheNuteTransition__dummysDeclaration__ models.TheNuteTransition
var __TheNuteTransition_time__dummyDeclaration time.Duration

var mutexTheNuteTransition sync.Mutex

// An TheNuteTransitionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTheNuteTransition updateTheNuteTransition deleteTheNuteTransition
type TheNuteTransitionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TheNuteTransitionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTheNuteTransition updateTheNuteTransition
type TheNuteTransitionInput struct {
	// The TheNuteTransition to submit or modify
	// in: body
	TheNuteTransition *orm.TheNuteTransitionAPI
}

// GetTheNuteTransitions
//
// swagger:route GET /thenutetransitions thenutetransitions getTheNuteTransitions
//
// # Get all thenutetransitions
//
// Responses:
// default: genericError
//
//	200: thenutetransitionDBResponse
func (controller *Controller) GetTheNuteTransitions(c *gin.Context) {

	// source slice
	var thenutetransitionDBs []orm.TheNuteTransitionDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTheNuteTransitions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteTransition.GetDB()

	query := db.Find(&thenutetransitionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	thenutetransitionAPIs := make([]orm.TheNuteTransitionAPI, 0)

	// for each thenutetransition, update fields from the database nullable fields
	for idx := range thenutetransitionDBs {
		thenutetransitionDB := &thenutetransitionDBs[idx]
		_ = thenutetransitionDB
		var thenutetransitionAPI orm.TheNuteTransitionAPI

		// insertion point for updating fields
		thenutetransitionAPI.ID = thenutetransitionDB.ID
		thenutetransitionDB.CopyBasicFieldsToTheNuteTransition_WOP(&thenutetransitionAPI.TheNuteTransition_WOP)
		thenutetransitionAPI.TheNuteTransitionPointersEncoding = thenutetransitionDB.TheNuteTransitionPointersEncoding
		thenutetransitionAPIs = append(thenutetransitionAPIs, thenutetransitionAPI)
	}

	c.JSON(http.StatusOK, thenutetransitionAPIs)
}

// PostTheNuteTransition
//
// swagger:route POST /thenutetransitions thenutetransitions postTheNuteTransition
//
// Creates a thenutetransition
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTheNuteTransition(c *gin.Context) {

	mutexTheNuteTransition.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTheNuteTransitions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteTransition.GetDB()

	// Validate input
	var input orm.TheNuteTransitionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create thenutetransition
	thenutetransitionDB := orm.TheNuteTransitionDB{}
	thenutetransitionDB.TheNuteTransitionPointersEncoding = input.TheNuteTransitionPointersEncoding
	thenutetransitionDB.CopyBasicFieldsFromTheNuteTransition_WOP(&input.TheNuteTransition_WOP)

	query := db.Create(&thenutetransitionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTheNuteTransition.CheckoutPhaseOneInstance(&thenutetransitionDB)
	thenutetransition := backRepo.BackRepoTheNuteTransition.Map_TheNuteTransitionDBID_TheNuteTransitionPtr[thenutetransitionDB.ID]

	if thenutetransition != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), thenutetransition)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, thenutetransitionDB)

	mutexTheNuteTransition.Unlock()
}

// GetTheNuteTransition
//
// swagger:route GET /thenutetransitions/{ID} thenutetransitions getTheNuteTransition
//
// Gets the details for a thenutetransition.
//
// Responses:
// default: genericError
//
//	200: thenutetransitionDBResponse
func (controller *Controller) GetTheNuteTransition(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTheNuteTransition", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteTransition.GetDB()

	// Get thenutetransitionDB in DB
	var thenutetransitionDB orm.TheNuteTransitionDB
	if err := db.First(&thenutetransitionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var thenutetransitionAPI orm.TheNuteTransitionAPI
	thenutetransitionAPI.ID = thenutetransitionDB.ID
	thenutetransitionAPI.TheNuteTransitionPointersEncoding = thenutetransitionDB.TheNuteTransitionPointersEncoding
	thenutetransitionDB.CopyBasicFieldsToTheNuteTransition_WOP(&thenutetransitionAPI.TheNuteTransition_WOP)

	c.JSON(http.StatusOK, thenutetransitionAPI)
}

// UpdateTheNuteTransition
//
// swagger:route PATCH /thenutetransitions/{ID} thenutetransitions updateTheNuteTransition
//
// # Update a thenutetransition
//
// Responses:
// default: genericError
//
//	200: thenutetransitionDBResponse
func (controller *Controller) UpdateTheNuteTransition(c *gin.Context) {

	mutexTheNuteTransition.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTheNuteTransition", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteTransition.GetDB()

	// Validate input
	var input orm.TheNuteTransitionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var thenutetransitionDB orm.TheNuteTransitionDB

	// fetch the thenutetransition
	query := db.First(&thenutetransitionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	thenutetransitionDB.CopyBasicFieldsFromTheNuteTransition_WOP(&input.TheNuteTransition_WOP)
	thenutetransitionDB.TheNuteTransitionPointersEncoding = input.TheNuteTransitionPointersEncoding

	query = db.Model(&thenutetransitionDB).Updates(thenutetransitionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	thenutetransitionNew := new(models.TheNuteTransition)
	thenutetransitionDB.CopyBasicFieldsToTheNuteTransition(thenutetransitionNew)

	// redeem pointers
	thenutetransitionDB.DecodePointers(backRepo, thenutetransitionNew)

	// get stage instance from DB instance, and call callback function
	thenutetransitionOld := backRepo.BackRepoTheNuteTransition.Map_TheNuteTransitionDBID_TheNuteTransitionPtr[thenutetransitionDB.ID]
	if thenutetransitionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), thenutetransitionOld, thenutetransitionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the thenutetransitionDB
	c.JSON(http.StatusOK, thenutetransitionDB)

	mutexTheNuteTransition.Unlock()
}

// DeleteTheNuteTransition
//
// swagger:route DELETE /thenutetransitions/{ID} thenutetransitions deleteTheNuteTransition
//
// # Delete a thenutetransition
//
// default: genericError
//
//	200: thenutetransitionDBResponse
func (controller *Controller) DeleteTheNuteTransition(c *gin.Context) {

	mutexTheNuteTransition.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTheNuteTransition", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTheNuteTransition.GetDB()

	// Get model if exist
	var thenutetransitionDB orm.TheNuteTransitionDB
	if err := db.First(&thenutetransitionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&thenutetransitionDB)

	// get an instance (not staged) from DB instance, and call callback function
	thenutetransitionDeleted := new(models.TheNuteTransition)
	thenutetransitionDB.CopyBasicFieldsToTheNuteTransition(thenutetransitionDeleted)

	// get stage instance from DB instance, and call callback function
	thenutetransitionStaged := backRepo.BackRepoTheNuteTransition.Map_TheNuteTransitionDBID_TheNuteTransitionPtr[thenutetransitionDB.ID]
	if thenutetransitionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), thenutetransitionStaged, thenutetransitionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexTheNuteTransition.Unlock()
}
