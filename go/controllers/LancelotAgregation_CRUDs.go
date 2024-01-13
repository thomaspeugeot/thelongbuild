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
var __LancelotAgregation__dummysDeclaration__ models.LancelotAgregation
var __LancelotAgregation_time__dummyDeclaration time.Duration

var mutexLancelotAgregation sync.Mutex

// An LancelotAgregationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLancelotAgregation updateLancelotAgregation deleteLancelotAgregation
type LancelotAgregationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LancelotAgregationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLancelotAgregation updateLancelotAgregation
type LancelotAgregationInput struct {
	// The LancelotAgregation to submit or modify
	// in: body
	LancelotAgregation *orm.LancelotAgregationAPI
}

// GetLancelotAgregations
//
// swagger:route GET /lancelotagregations lancelotagregations getLancelotAgregations
//
// # Get all lancelotagregations
//
// Responses:
// default: genericError
//
//	200: lancelotagregationDBResponse
func (controller *Controller) GetLancelotAgregations(c *gin.Context) {

	// source slice
	var lancelotagregationDBs []orm.LancelotAgregationDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelotAgregations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregation.GetDB()

	query := db.Find(&lancelotagregationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lancelotagregationAPIs := make([]orm.LancelotAgregationAPI, 0)

	// for each lancelotagregation, update fields from the database nullable fields
	for idx := range lancelotagregationDBs {
		lancelotagregationDB := &lancelotagregationDBs[idx]
		_ = lancelotagregationDB
		var lancelotagregationAPI orm.LancelotAgregationAPI

		// insertion point for updating fields
		lancelotagregationAPI.ID = lancelotagregationDB.ID
		lancelotagregationDB.CopyBasicFieldsToLancelotAgregation_WOP(&lancelotagregationAPI.LancelotAgregation_WOP)
		lancelotagregationAPI.LancelotAgregationPointersEncoding = lancelotagregationDB.LancelotAgregationPointersEncoding
		lancelotagregationAPIs = append(lancelotagregationAPIs, lancelotagregationAPI)
	}

	c.JSON(http.StatusOK, lancelotagregationAPIs)
}

// PostLancelotAgregation
//
// swagger:route POST /lancelotagregations lancelotagregations postLancelotAgregation
//
// Creates a lancelotagregation
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLancelotAgregation(c *gin.Context) {

	mutexLancelotAgregation.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLancelotAgregations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregation.GetDB()

	// Validate input
	var input orm.LancelotAgregationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lancelotagregation
	lancelotagregationDB := orm.LancelotAgregationDB{}
	lancelotagregationDB.LancelotAgregationPointersEncoding = input.LancelotAgregationPointersEncoding
	lancelotagregationDB.CopyBasicFieldsFromLancelotAgregation_WOP(&input.LancelotAgregation_WOP)

	query := db.Create(&lancelotagregationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLancelotAgregation.CheckoutPhaseOneInstance(&lancelotagregationDB)
	lancelotagregation := backRepo.BackRepoLancelotAgregation.Map_LancelotAgregationDBID_LancelotAgregationPtr[lancelotagregationDB.ID]

	if lancelotagregation != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lancelotagregation)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lancelotagregationDB)

	mutexLancelotAgregation.Unlock()
}

// GetLancelotAgregation
//
// swagger:route GET /lancelotagregations/{ID} lancelotagregations getLancelotAgregation
//
// Gets the details for a lancelotagregation.
//
// Responses:
// default: genericError
//
//	200: lancelotagregationDBResponse
func (controller *Controller) GetLancelotAgregation(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelotAgregation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregation.GetDB()

	// Get lancelotagregationDB in DB
	var lancelotagregationDB orm.LancelotAgregationDB
	if err := db.First(&lancelotagregationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lancelotagregationAPI orm.LancelotAgregationAPI
	lancelotagregationAPI.ID = lancelotagregationDB.ID
	lancelotagregationAPI.LancelotAgregationPointersEncoding = lancelotagregationDB.LancelotAgregationPointersEncoding
	lancelotagregationDB.CopyBasicFieldsToLancelotAgregation_WOP(&lancelotagregationAPI.LancelotAgregation_WOP)

	c.JSON(http.StatusOK, lancelotagregationAPI)
}

// UpdateLancelotAgregation
//
// swagger:route PATCH /lancelotagregations/{ID} lancelotagregations updateLancelotAgregation
//
// # Update a lancelotagregation
//
// Responses:
// default: genericError
//
//	200: lancelotagregationDBResponse
func (controller *Controller) UpdateLancelotAgregation(c *gin.Context) {

	mutexLancelotAgregation.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLancelotAgregation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregation.GetDB()

	// Validate input
	var input orm.LancelotAgregationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lancelotagregationDB orm.LancelotAgregationDB

	// fetch the lancelotagregation
	query := db.First(&lancelotagregationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lancelotagregationDB.CopyBasicFieldsFromLancelotAgregation_WOP(&input.LancelotAgregation_WOP)
	lancelotagregationDB.LancelotAgregationPointersEncoding = input.LancelotAgregationPointersEncoding

	query = db.Model(&lancelotagregationDB).Updates(lancelotagregationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lancelotagregationNew := new(models.LancelotAgregation)
	lancelotagregationDB.CopyBasicFieldsToLancelotAgregation(lancelotagregationNew)

	// redeem pointers
	lancelotagregationDB.DecodePointers(backRepo, lancelotagregationNew)

	// get stage instance from DB instance, and call callback function
	lancelotagregationOld := backRepo.BackRepoLancelotAgregation.Map_LancelotAgregationDBID_LancelotAgregationPtr[lancelotagregationDB.ID]
	if lancelotagregationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lancelotagregationOld, lancelotagregationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lancelotagregationDB
	c.JSON(http.StatusOK, lancelotagregationDB)

	mutexLancelotAgregation.Unlock()
}

// DeleteLancelotAgregation
//
// swagger:route DELETE /lancelotagregations/{ID} lancelotagregations deleteLancelotAgregation
//
// # Delete a lancelotagregation
//
// default: genericError
//
//	200: lancelotagregationDBResponse
func (controller *Controller) DeleteLancelotAgregation(c *gin.Context) {

	mutexLancelotAgregation.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLancelotAgregation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregation.GetDB()

	// Get model if exist
	var lancelotagregationDB orm.LancelotAgregationDB
	if err := db.First(&lancelotagregationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&lancelotagregationDB)

	// get an instance (not staged) from DB instance, and call callback function
	lancelotagregationDeleted := new(models.LancelotAgregation)
	lancelotagregationDB.CopyBasicFieldsToLancelotAgregation(lancelotagregationDeleted)

	// get stage instance from DB instance, and call callback function
	lancelotagregationStaged := backRepo.BackRepoLancelotAgregation.Map_LancelotAgregationDBID_LancelotAgregationPtr[lancelotagregationDB.ID]
	if lancelotagregationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lancelotagregationStaged, lancelotagregationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLancelotAgregation.Unlock()
}
