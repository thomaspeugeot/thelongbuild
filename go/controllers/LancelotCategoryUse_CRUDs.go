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
var __LancelotCategoryUse__dummysDeclaration__ models.LancelotCategoryUse
var __LancelotCategoryUse_time__dummyDeclaration time.Duration

var mutexLancelotCategoryUse sync.Mutex

// An LancelotCategoryUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLancelotCategoryUse updateLancelotCategoryUse deleteLancelotCategoryUse
type LancelotCategoryUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LancelotCategoryUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLancelotCategoryUse updateLancelotCategoryUse
type LancelotCategoryUseInput struct {
	// The LancelotCategoryUse to submit or modify
	// in: body
	LancelotCategoryUse *orm.LancelotCategoryUseAPI
}

// GetLancelotCategoryUses
//
// swagger:route GET /lancelotcategoryuses lancelotcategoryuses getLancelotCategoryUses
//
// # Get all lancelotcategoryuses
//
// Responses:
// default: genericError
//
//	200: lancelotcategoryuseDBResponse
func (controller *Controller) GetLancelotCategoryUses(c *gin.Context) {

	// source slice
	var lancelotcategoryuseDBs []orm.LancelotCategoryUseDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelotCategoryUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategoryUse.GetDB()

	query := db.Find(&lancelotcategoryuseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lancelotcategoryuseAPIs := make([]orm.LancelotCategoryUseAPI, 0)

	// for each lancelotcategoryuse, update fields from the database nullable fields
	for idx := range lancelotcategoryuseDBs {
		lancelotcategoryuseDB := &lancelotcategoryuseDBs[idx]
		_ = lancelotcategoryuseDB
		var lancelotcategoryuseAPI orm.LancelotCategoryUseAPI

		// insertion point for updating fields
		lancelotcategoryuseAPI.ID = lancelotcategoryuseDB.ID
		lancelotcategoryuseDB.CopyBasicFieldsToLancelotCategoryUse_WOP(&lancelotcategoryuseAPI.LancelotCategoryUse_WOP)
		lancelotcategoryuseAPI.LancelotCategoryUsePointersEncoding = lancelotcategoryuseDB.LancelotCategoryUsePointersEncoding
		lancelotcategoryuseAPIs = append(lancelotcategoryuseAPIs, lancelotcategoryuseAPI)
	}

	c.JSON(http.StatusOK, lancelotcategoryuseAPIs)
}

// PostLancelotCategoryUse
//
// swagger:route POST /lancelotcategoryuses lancelotcategoryuses postLancelotCategoryUse
//
// Creates a lancelotcategoryuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLancelotCategoryUse(c *gin.Context) {

	mutexLancelotCategoryUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLancelotCategoryUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategoryUse.GetDB()

	// Validate input
	var input orm.LancelotCategoryUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lancelotcategoryuse
	lancelotcategoryuseDB := orm.LancelotCategoryUseDB{}
	lancelotcategoryuseDB.LancelotCategoryUsePointersEncoding = input.LancelotCategoryUsePointersEncoding
	lancelotcategoryuseDB.CopyBasicFieldsFromLancelotCategoryUse_WOP(&input.LancelotCategoryUse_WOP)

	query := db.Create(&lancelotcategoryuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLancelotCategoryUse.CheckoutPhaseOneInstance(&lancelotcategoryuseDB)
	lancelotcategoryuse := backRepo.BackRepoLancelotCategoryUse.Map_LancelotCategoryUseDBID_LancelotCategoryUsePtr[lancelotcategoryuseDB.ID]

	if lancelotcategoryuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lancelotcategoryuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lancelotcategoryuseDB)

	mutexLancelotCategoryUse.Unlock()
}

// GetLancelotCategoryUse
//
// swagger:route GET /lancelotcategoryuses/{ID} lancelotcategoryuses getLancelotCategoryUse
//
// Gets the details for a lancelotcategoryuse.
//
// Responses:
// default: genericError
//
//	200: lancelotcategoryuseDBResponse
func (controller *Controller) GetLancelotCategoryUse(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelotCategoryUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategoryUse.GetDB()

	// Get lancelotcategoryuseDB in DB
	var lancelotcategoryuseDB orm.LancelotCategoryUseDB
	if err := db.First(&lancelotcategoryuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lancelotcategoryuseAPI orm.LancelotCategoryUseAPI
	lancelotcategoryuseAPI.ID = lancelotcategoryuseDB.ID
	lancelotcategoryuseAPI.LancelotCategoryUsePointersEncoding = lancelotcategoryuseDB.LancelotCategoryUsePointersEncoding
	lancelotcategoryuseDB.CopyBasicFieldsToLancelotCategoryUse_WOP(&lancelotcategoryuseAPI.LancelotCategoryUse_WOP)

	c.JSON(http.StatusOK, lancelotcategoryuseAPI)
}

// UpdateLancelotCategoryUse
//
// swagger:route PATCH /lancelotcategoryuses/{ID} lancelotcategoryuses updateLancelotCategoryUse
//
// # Update a lancelotcategoryuse
//
// Responses:
// default: genericError
//
//	200: lancelotcategoryuseDBResponse
func (controller *Controller) UpdateLancelotCategoryUse(c *gin.Context) {

	mutexLancelotCategoryUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLancelotCategoryUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategoryUse.GetDB()

	// Validate input
	var input orm.LancelotCategoryUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lancelotcategoryuseDB orm.LancelotCategoryUseDB

	// fetch the lancelotcategoryuse
	query := db.First(&lancelotcategoryuseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lancelotcategoryuseDB.CopyBasicFieldsFromLancelotCategoryUse_WOP(&input.LancelotCategoryUse_WOP)
	lancelotcategoryuseDB.LancelotCategoryUsePointersEncoding = input.LancelotCategoryUsePointersEncoding

	query = db.Model(&lancelotcategoryuseDB).Updates(lancelotcategoryuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lancelotcategoryuseNew := new(models.LancelotCategoryUse)
	lancelotcategoryuseDB.CopyBasicFieldsToLancelotCategoryUse(lancelotcategoryuseNew)

	// redeem pointers
	lancelotcategoryuseDB.DecodePointers(backRepo, lancelotcategoryuseNew)

	// get stage instance from DB instance, and call callback function
	lancelotcategoryuseOld := backRepo.BackRepoLancelotCategoryUse.Map_LancelotCategoryUseDBID_LancelotCategoryUsePtr[lancelotcategoryuseDB.ID]
	if lancelotcategoryuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lancelotcategoryuseOld, lancelotcategoryuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lancelotcategoryuseDB
	c.JSON(http.StatusOK, lancelotcategoryuseDB)

	mutexLancelotCategoryUse.Unlock()
}

// DeleteLancelotCategoryUse
//
// swagger:route DELETE /lancelotcategoryuses/{ID} lancelotcategoryuses deleteLancelotCategoryUse
//
// # Delete a lancelotcategoryuse
//
// default: genericError
//
//	200: lancelotcategoryuseDBResponse
func (controller *Controller) DeleteLancelotCategoryUse(c *gin.Context) {

	mutexLancelotCategoryUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLancelotCategoryUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategoryUse.GetDB()

	// Get model if exist
	var lancelotcategoryuseDB orm.LancelotCategoryUseDB
	if err := db.First(&lancelotcategoryuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&lancelotcategoryuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	lancelotcategoryuseDeleted := new(models.LancelotCategoryUse)
	lancelotcategoryuseDB.CopyBasicFieldsToLancelotCategoryUse(lancelotcategoryuseDeleted)

	// get stage instance from DB instance, and call callback function
	lancelotcategoryuseStaged := backRepo.BackRepoLancelotCategoryUse.Map_LancelotCategoryUseDBID_LancelotCategoryUsePtr[lancelotcategoryuseDB.ID]
	if lancelotcategoryuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lancelotcategoryuseStaged, lancelotcategoryuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLancelotCategoryUse.Unlock()
}
