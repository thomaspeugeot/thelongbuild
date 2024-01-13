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
var __Lancelot__dummysDeclaration__ models.Lancelot
var __Lancelot_time__dummyDeclaration time.Duration

var mutexLancelot sync.Mutex

// An LancelotID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLancelot updateLancelot deleteLancelot
type LancelotID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LancelotInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLancelot updateLancelot
type LancelotInput struct {
	// The Lancelot to submit or modify
	// in: body
	Lancelot *orm.LancelotAPI
}

// GetLancelots
//
// swagger:route GET /lancelots lancelots getLancelots
//
// # Get all lancelots
//
// Responses:
// default: genericError
//
//	200: lancelotDBResponse
func (controller *Controller) GetLancelots(c *gin.Context) {

	// source slice
	var lancelotDBs []orm.LancelotDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelots", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelot.GetDB()

	query := db.Find(&lancelotDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lancelotAPIs := make([]orm.LancelotAPI, 0)

	// for each lancelot, update fields from the database nullable fields
	for idx := range lancelotDBs {
		lancelotDB := &lancelotDBs[idx]
		_ = lancelotDB
		var lancelotAPI orm.LancelotAPI

		// insertion point for updating fields
		lancelotAPI.ID = lancelotDB.ID
		lancelotDB.CopyBasicFieldsToLancelot_WOP(&lancelotAPI.Lancelot_WOP)
		lancelotAPI.LancelotPointersEncoding = lancelotDB.LancelotPointersEncoding
		lancelotAPIs = append(lancelotAPIs, lancelotAPI)
	}

	c.JSON(http.StatusOK, lancelotAPIs)
}

// PostLancelot
//
// swagger:route POST /lancelots lancelots postLancelot
//
// Creates a lancelot
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLancelot(c *gin.Context) {

	mutexLancelot.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLancelots", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelot.GetDB()

	// Validate input
	var input orm.LancelotAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lancelot
	lancelotDB := orm.LancelotDB{}
	lancelotDB.LancelotPointersEncoding = input.LancelotPointersEncoding
	lancelotDB.CopyBasicFieldsFromLancelot_WOP(&input.Lancelot_WOP)

	query := db.Create(&lancelotDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLancelot.CheckoutPhaseOneInstance(&lancelotDB)
	lancelot := backRepo.BackRepoLancelot.Map_LancelotDBID_LancelotPtr[lancelotDB.ID]

	if lancelot != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lancelot)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lancelotDB)

	mutexLancelot.Unlock()
}

// GetLancelot
//
// swagger:route GET /lancelots/{ID} lancelots getLancelot
//
// Gets the details for a lancelot.
//
// Responses:
// default: genericError
//
//	200: lancelotDBResponse
func (controller *Controller) GetLancelot(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelot.GetDB()

	// Get lancelotDB in DB
	var lancelotDB orm.LancelotDB
	if err := db.First(&lancelotDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lancelotAPI orm.LancelotAPI
	lancelotAPI.ID = lancelotDB.ID
	lancelotAPI.LancelotPointersEncoding = lancelotDB.LancelotPointersEncoding
	lancelotDB.CopyBasicFieldsToLancelot_WOP(&lancelotAPI.Lancelot_WOP)

	c.JSON(http.StatusOK, lancelotAPI)
}

// UpdateLancelot
//
// swagger:route PATCH /lancelots/{ID} lancelots updateLancelot
//
// # Update a lancelot
//
// Responses:
// default: genericError
//
//	200: lancelotDBResponse
func (controller *Controller) UpdateLancelot(c *gin.Context) {

	mutexLancelot.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLancelot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelot.GetDB()

	// Validate input
	var input orm.LancelotAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lancelotDB orm.LancelotDB

	// fetch the lancelot
	query := db.First(&lancelotDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lancelotDB.CopyBasicFieldsFromLancelot_WOP(&input.Lancelot_WOP)
	lancelotDB.LancelotPointersEncoding = input.LancelotPointersEncoding

	query = db.Model(&lancelotDB).Updates(lancelotDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lancelotNew := new(models.Lancelot)
	lancelotDB.CopyBasicFieldsToLancelot(lancelotNew)

	// redeem pointers
	lancelotDB.DecodePointers(backRepo, lancelotNew)

	// get stage instance from DB instance, and call callback function
	lancelotOld := backRepo.BackRepoLancelot.Map_LancelotDBID_LancelotPtr[lancelotDB.ID]
	if lancelotOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lancelotOld, lancelotNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lancelotDB
	c.JSON(http.StatusOK, lancelotDB)

	mutexLancelot.Unlock()
}

// DeleteLancelot
//
// swagger:route DELETE /lancelots/{ID} lancelots deleteLancelot
//
// # Delete a lancelot
//
// default: genericError
//
//	200: lancelotDBResponse
func (controller *Controller) DeleteLancelot(c *gin.Context) {

	mutexLancelot.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLancelot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelot.GetDB()

	// Get model if exist
	var lancelotDB orm.LancelotDB
	if err := db.First(&lancelotDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&lancelotDB)

	// get an instance (not staged) from DB instance, and call callback function
	lancelotDeleted := new(models.Lancelot)
	lancelotDB.CopyBasicFieldsToLancelot(lancelotDeleted)

	// get stage instance from DB instance, and call callback function
	lancelotStaged := backRepo.BackRepoLancelot.Map_LancelotDBID_LancelotPtr[lancelotDB.ID]
	if lancelotStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lancelotStaged, lancelotDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLancelot.Unlock()
}
