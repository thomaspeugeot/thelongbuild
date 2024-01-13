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
var __LancelotCategory__dummysDeclaration__ models.LancelotCategory
var __LancelotCategory_time__dummyDeclaration time.Duration

var mutexLancelotCategory sync.Mutex

// An LancelotCategoryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLancelotCategory updateLancelotCategory deleteLancelotCategory
type LancelotCategoryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LancelotCategoryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLancelotCategory updateLancelotCategory
type LancelotCategoryInput struct {
	// The LancelotCategory to submit or modify
	// in: body
	LancelotCategory *orm.LancelotCategoryAPI
}

// GetLancelotCategorys
//
// swagger:route GET /lancelotcategorys lancelotcategorys getLancelotCategorys
//
// # Get all lancelotcategorys
//
// Responses:
// default: genericError
//
//	200: lancelotcategoryDBResponse
func (controller *Controller) GetLancelotCategorys(c *gin.Context) {

	// source slice
	var lancelotcategoryDBs []orm.LancelotCategoryDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelotCategorys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategory.GetDB()

	query := db.Find(&lancelotcategoryDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lancelotcategoryAPIs := make([]orm.LancelotCategoryAPI, 0)

	// for each lancelotcategory, update fields from the database nullable fields
	for idx := range lancelotcategoryDBs {
		lancelotcategoryDB := &lancelotcategoryDBs[idx]
		_ = lancelotcategoryDB
		var lancelotcategoryAPI orm.LancelotCategoryAPI

		// insertion point for updating fields
		lancelotcategoryAPI.ID = lancelotcategoryDB.ID
		lancelotcategoryDB.CopyBasicFieldsToLancelotCategory_WOP(&lancelotcategoryAPI.LancelotCategory_WOP)
		lancelotcategoryAPI.LancelotCategoryPointersEncoding = lancelotcategoryDB.LancelotCategoryPointersEncoding
		lancelotcategoryAPIs = append(lancelotcategoryAPIs, lancelotcategoryAPI)
	}

	c.JSON(http.StatusOK, lancelotcategoryAPIs)
}

// PostLancelotCategory
//
// swagger:route POST /lancelotcategorys lancelotcategorys postLancelotCategory
//
// Creates a lancelotcategory
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLancelotCategory(c *gin.Context) {

	mutexLancelotCategory.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLancelotCategorys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategory.GetDB()

	// Validate input
	var input orm.LancelotCategoryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create lancelotcategory
	lancelotcategoryDB := orm.LancelotCategoryDB{}
	lancelotcategoryDB.LancelotCategoryPointersEncoding = input.LancelotCategoryPointersEncoding
	lancelotcategoryDB.CopyBasicFieldsFromLancelotCategory_WOP(&input.LancelotCategory_WOP)

	query := db.Create(&lancelotcategoryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLancelotCategory.CheckoutPhaseOneInstance(&lancelotcategoryDB)
	lancelotcategory := backRepo.BackRepoLancelotCategory.Map_LancelotCategoryDBID_LancelotCategoryPtr[lancelotcategoryDB.ID]

	if lancelotcategory != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), lancelotcategory)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lancelotcategoryDB)

	mutexLancelotCategory.Unlock()
}

// GetLancelotCategory
//
// swagger:route GET /lancelotcategorys/{ID} lancelotcategorys getLancelotCategory
//
// Gets the details for a lancelotcategory.
//
// Responses:
// default: genericError
//
//	200: lancelotcategoryDBResponse
func (controller *Controller) GetLancelotCategory(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLancelotCategory", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategory.GetDB()

	// Get lancelotcategoryDB in DB
	var lancelotcategoryDB orm.LancelotCategoryDB
	if err := db.First(&lancelotcategoryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lancelotcategoryAPI orm.LancelotCategoryAPI
	lancelotcategoryAPI.ID = lancelotcategoryDB.ID
	lancelotcategoryAPI.LancelotCategoryPointersEncoding = lancelotcategoryDB.LancelotCategoryPointersEncoding
	lancelotcategoryDB.CopyBasicFieldsToLancelotCategory_WOP(&lancelotcategoryAPI.LancelotCategory_WOP)

	c.JSON(http.StatusOK, lancelotcategoryAPI)
}

// UpdateLancelotCategory
//
// swagger:route PATCH /lancelotcategorys/{ID} lancelotcategorys updateLancelotCategory
//
// # Update a lancelotcategory
//
// Responses:
// default: genericError
//
//	200: lancelotcategoryDBResponse
func (controller *Controller) UpdateLancelotCategory(c *gin.Context) {

	mutexLancelotCategory.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLancelotCategory", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategory.GetDB()

	// Validate input
	var input orm.LancelotCategoryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lancelotcategoryDB orm.LancelotCategoryDB

	// fetch the lancelotcategory
	query := db.First(&lancelotcategoryDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lancelotcategoryDB.CopyBasicFieldsFromLancelotCategory_WOP(&input.LancelotCategory_WOP)
	lancelotcategoryDB.LancelotCategoryPointersEncoding = input.LancelotCategoryPointersEncoding

	query = db.Model(&lancelotcategoryDB).Updates(lancelotcategoryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lancelotcategoryNew := new(models.LancelotCategory)
	lancelotcategoryDB.CopyBasicFieldsToLancelotCategory(lancelotcategoryNew)

	// redeem pointers
	lancelotcategoryDB.DecodePointers(backRepo, lancelotcategoryNew)

	// get stage instance from DB instance, and call callback function
	lancelotcategoryOld := backRepo.BackRepoLancelotCategory.Map_LancelotCategoryDBID_LancelotCategoryPtr[lancelotcategoryDB.ID]
	if lancelotcategoryOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lancelotcategoryOld, lancelotcategoryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lancelotcategoryDB
	c.JSON(http.StatusOK, lancelotcategoryDB)

	mutexLancelotCategory.Unlock()
}

// DeleteLancelotCategory
//
// swagger:route DELETE /lancelotcategorys/{ID} lancelotcategorys deleteLancelotCategory
//
// # Delete a lancelotcategory
//
// default: genericError
//
//	200: lancelotcategoryDBResponse
func (controller *Controller) DeleteLancelotCategory(c *gin.Context) {

	mutexLancelotCategory.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLancelotCategory", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLancelotCategory.GetDB()

	// Get model if exist
	var lancelotcategoryDB orm.LancelotCategoryDB
	if err := db.First(&lancelotcategoryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&lancelotcategoryDB)

	// get an instance (not staged) from DB instance, and call callback function
	lancelotcategoryDeleted := new(models.LancelotCategory)
	lancelotcategoryDB.CopyBasicFieldsToLancelotCategory(lancelotcategoryDeleted)

	// get stage instance from DB instance, and call callback function
	lancelotcategoryStaged := backRepo.BackRepoLancelotCategory.Map_LancelotCategoryDBID_LancelotCategoryPtr[lancelotcategoryDB.ID]
	if lancelotcategoryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lancelotcategoryStaged, lancelotcategoryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLancelotCategory.Unlock()
}
