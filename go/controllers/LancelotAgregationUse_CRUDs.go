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
var __LancelotAgregationUse__dummysDeclaration__ models.LancelotAgregationUse
var __LancelotAgregationUse_time__dummyDeclaration time.Duration

var mutexLancelotAgregationUse sync.Mutex

// An LancelotAgregationUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLancelotAgregationUse updateLancelotAgregationUse deleteLancelotAgregationUse
type LancelotAgregationUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LancelotAgregationUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLancelotAgregationUse updateLancelotAgregationUse
type LancelotAgregationUseInput struct {
	// The LancelotAgregationUse to submit or modify
	// in: body
	LancelotAgregationUse *orm.LancelotAgregationUseAPI
}

// GetLancelotAgregationUses
//
// swagger:route GET /lancelotagregationuses lancelotagregationuses getLancelotAgregationUses
//
// # Get all lancelotagregationuses
//
// Responses:
// default: genericError
//
//	200: lancelotagregationuseDBResponse
func (controller *Controller) GetLancelotAgregationUses(c *gin.Context) {

	// source slice
	var lancelotagregationuseDBs []orm.LancelotAgregationUseDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelotAgregationUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregationUse.GetDB()

	query := db.Find(&lancelotagregationuseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lancelotagregationuseAPIs := make([]orm.LancelotAgregationUseAPI, 0)

	// for each lancelotagregationuse, update fields from the database nullable fields
	for idx := range lancelotagregationuseDBs {
		lancelotagregationuseDB := &lancelotagregationuseDBs[idx]
		_ = lancelotagregationuseDB
		var lancelotagregationuseAPI orm.LancelotAgregationUseAPI

		// insertion point for updating fields
		lancelotagregationuseAPI.ID = lancelotagregationuseDB.ID
		lancelotagregationuseDB.CopyBasicFieldsToLancelotAgregationUse_WOP(&lancelotagregationuseAPI.LancelotAgregationUse_WOP)
		lancelotagregationuseAPI.LancelotAgregationUsePointersEncoding = lancelotagregationuseDB.LancelotAgregationUsePointersEncoding
		lancelotagregationuseAPIs = append(lancelotagregationuseAPIs, lancelotagregationuseAPI)
	}

	c.JSON(http.StatusOK, lancelotagregationuseAPIs)
}

// PostLancelotAgregationUse
//
// swagger:route POST /lancelotagregationuses lancelotagregationuses postLancelotAgregationUse
//
// Creates a lancelotagregationuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLancelotAgregationUse(c *gin.Context) {

	mutexLancelotAgregationUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLancelotAgregationUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregationUse.GetDB()

	// Validate input
	var input orm.LancelotAgregationUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lancelotagregationuse
	lancelotagregationuseDB := orm.LancelotAgregationUseDB{}
	lancelotagregationuseDB.LancelotAgregationUsePointersEncoding = input.LancelotAgregationUsePointersEncoding
	lancelotagregationuseDB.CopyBasicFieldsFromLancelotAgregationUse_WOP(&input.LancelotAgregationUse_WOP)

	query := db.Create(&lancelotagregationuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLancelotAgregationUse.CheckoutPhaseOneInstance(&lancelotagregationuseDB)
	lancelotagregationuse := backRepo.BackRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[lancelotagregationuseDB.ID]

	if lancelotagregationuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lancelotagregationuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lancelotagregationuseDB)

	mutexLancelotAgregationUse.Unlock()
}

// GetLancelotAgregationUse
//
// swagger:route GET /lancelotagregationuses/{ID} lancelotagregationuses getLancelotAgregationUse
//
// Gets the details for a lancelotagregationuse.
//
// Responses:
// default: genericError
//
//	200: lancelotagregationuseDBResponse
func (controller *Controller) GetLancelotAgregationUse(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelotAgregationUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregationUse.GetDB()

	// Get lancelotagregationuseDB in DB
	var lancelotagregationuseDB orm.LancelotAgregationUseDB
	if err := db.First(&lancelotagregationuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lancelotagregationuseAPI orm.LancelotAgregationUseAPI
	lancelotagregationuseAPI.ID = lancelotagregationuseDB.ID
	lancelotagregationuseAPI.LancelotAgregationUsePointersEncoding = lancelotagregationuseDB.LancelotAgregationUsePointersEncoding
	lancelotagregationuseDB.CopyBasicFieldsToLancelotAgregationUse_WOP(&lancelotagregationuseAPI.LancelotAgregationUse_WOP)

	c.JSON(http.StatusOK, lancelotagregationuseAPI)
}

// UpdateLancelotAgregationUse
//
// swagger:route PATCH /lancelotagregationuses/{ID} lancelotagregationuses updateLancelotAgregationUse
//
// # Update a lancelotagregationuse
//
// Responses:
// default: genericError
//
//	200: lancelotagregationuseDBResponse
func (controller *Controller) UpdateLancelotAgregationUse(c *gin.Context) {

	mutexLancelotAgregationUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLancelotAgregationUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregationUse.GetDB()

	// Validate input
	var input orm.LancelotAgregationUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lancelotagregationuseDB orm.LancelotAgregationUseDB

	// fetch the lancelotagregationuse
	query := db.First(&lancelotagregationuseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lancelotagregationuseDB.CopyBasicFieldsFromLancelotAgregationUse_WOP(&input.LancelotAgregationUse_WOP)
	lancelotagregationuseDB.LancelotAgregationUsePointersEncoding = input.LancelotAgregationUsePointersEncoding

	query = db.Model(&lancelotagregationuseDB).Updates(lancelotagregationuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lancelotagregationuseNew := new(models.LancelotAgregationUse)
	lancelotagregationuseDB.CopyBasicFieldsToLancelotAgregationUse(lancelotagregationuseNew)

	// redeem pointers
	lancelotagregationuseDB.DecodePointers(backRepo, lancelotagregationuseNew)

	// get stage instance from DB instance, and call callback function
	lancelotagregationuseOld := backRepo.BackRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[lancelotagregationuseDB.ID]
	if lancelotagregationuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lancelotagregationuseOld, lancelotagregationuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lancelotagregationuseDB
	c.JSON(http.StatusOK, lancelotagregationuseDB)

	mutexLancelotAgregationUse.Unlock()
}

// DeleteLancelotAgregationUse
//
// swagger:route DELETE /lancelotagregationuses/{ID} lancelotagregationuses deleteLancelotAgregationUse
//
// # Delete a lancelotagregationuse
//
// default: genericError
//
//	200: lancelotagregationuseDBResponse
func (controller *Controller) DeleteLancelotAgregationUse(c *gin.Context) {

	mutexLancelotAgregationUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLancelotAgregationUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotAgregationUse.GetDB()

	// Get model if exist
	var lancelotagregationuseDB orm.LancelotAgregationUseDB
	if err := db.First(&lancelotagregationuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&lancelotagregationuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	lancelotagregationuseDeleted := new(models.LancelotAgregationUse)
	lancelotagregationuseDB.CopyBasicFieldsToLancelotAgregationUse(lancelotagregationuseDeleted)

	// get stage instance from DB instance, and call callback function
	lancelotagregationuseStaged := backRepo.BackRepoLancelotAgregationUse.Map_LancelotAgregationUseDBID_LancelotAgregationUsePtr[lancelotagregationuseDB.ID]
	if lancelotagregationuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lancelotagregationuseStaged, lancelotagregationuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLancelotAgregationUse.Unlock()
}
