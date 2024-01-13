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
var __KingArthur__dummysDeclaration__ models.KingArthur
var __KingArthur_time__dummyDeclaration time.Duration

var mutexKingArthur sync.Mutex

// An KingArthurID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getKingArthur updateKingArthur deleteKingArthur
type KingArthurID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// KingArthurInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postKingArthur updateKingArthur
type KingArthurInput struct {
	// The KingArthur to submit or modify
	// in: body
	KingArthur *orm.KingArthurAPI
}

// GetKingArthurs
//
// swagger:route GET /kingarthurs kingarthurs getKingArthurs
//
// # Get all kingarthurs
//
// Responses:
// default: genericError
//
//	200: kingarthurDBResponse
func (controller *Controller) GetKingArthurs(c *gin.Context) {

	// source slice
	var kingarthurDBs []orm.KingArthurDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKingArthurs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthur.GetDB()

	query := db.Find(&kingarthurDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	kingarthurAPIs := make([]orm.KingArthurAPI, 0)

	// for each kingarthur, update fields from the database nullable fields
	for idx := range kingarthurDBs {
		kingarthurDB := &kingarthurDBs[idx]
		_ = kingarthurDB
		var kingarthurAPI orm.KingArthurAPI

		// insertion point for updating fields
		kingarthurAPI.ID = kingarthurDB.ID
		kingarthurDB.CopyBasicFieldsToKingArthur_WOP(&kingarthurAPI.KingArthur_WOP)
		kingarthurAPI.KingArthurPointersEncoding = kingarthurDB.KingArthurPointersEncoding
		kingarthurAPIs = append(kingarthurAPIs, kingarthurAPI)
	}

	c.JSON(http.StatusOK, kingarthurAPIs)
}

// PostKingArthur
//
// swagger:route POST /kingarthurs kingarthurs postKingArthur
//
// Creates a kingarthur
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostKingArthur(c *gin.Context) {

	mutexKingArthur.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostKingArthurs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthur.GetDB()

	// Validate input
	var input orm.KingArthurAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create kingarthur
	kingarthurDB := orm.KingArthurDB{}
	kingarthurDB.KingArthurPointersEncoding = input.KingArthurPointersEncoding
	kingarthurDB.CopyBasicFieldsFromKingArthur_WOP(&input.KingArthur_WOP)

	query := db.Create(&kingarthurDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoKingArthur.CheckoutPhaseOneInstance(&kingarthurDB)
	kingarthur := backRepo.BackRepoKingArthur.Map_KingArthurDBID_KingArthurPtr[kingarthurDB.ID]

	if kingarthur != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), kingarthur)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, kingarthurDB)

	mutexKingArthur.Unlock()
}

// GetKingArthur
//
// swagger:route GET /kingarthurs/{ID} kingarthurs getKingArthur
//
// Gets the details for a kingarthur.
//
// Responses:
// default: genericError
//
//	200: kingarthurDBResponse
func (controller *Controller) GetKingArthur(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKingArthur", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthur.GetDB()

	// Get kingarthurDB in DB
	var kingarthurDB orm.KingArthurDB
	if err := db.First(&kingarthurDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var kingarthurAPI orm.KingArthurAPI
	kingarthurAPI.ID = kingarthurDB.ID
	kingarthurAPI.KingArthurPointersEncoding = kingarthurDB.KingArthurPointersEncoding
	kingarthurDB.CopyBasicFieldsToKingArthur_WOP(&kingarthurAPI.KingArthur_WOP)

	c.JSON(http.StatusOK, kingarthurAPI)
}

// UpdateKingArthur
//
// swagger:route PATCH /kingarthurs/{ID} kingarthurs updateKingArthur
//
// # Update a kingarthur
//
// Responses:
// default: genericError
//
//	200: kingarthurDBResponse
func (controller *Controller) UpdateKingArthur(c *gin.Context) {

	mutexKingArthur.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateKingArthur", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthur.GetDB()

	// Validate input
	var input orm.KingArthurAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var kingarthurDB orm.KingArthurDB

	// fetch the kingarthur
	query := db.First(&kingarthurDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	kingarthurDB.CopyBasicFieldsFromKingArthur_WOP(&input.KingArthur_WOP)
	kingarthurDB.KingArthurPointersEncoding = input.KingArthurPointersEncoding

	query = db.Model(&kingarthurDB).Updates(kingarthurDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	kingarthurNew := new(models.KingArthur)
	kingarthurDB.CopyBasicFieldsToKingArthur(kingarthurNew)

	// redeem pointers
	kingarthurDB.DecodePointers(backRepo, kingarthurNew)

	// get stage instance from DB instance, and call callback function
	kingarthurOld := backRepo.BackRepoKingArthur.Map_KingArthurDBID_KingArthurPtr[kingarthurDB.ID]
	if kingarthurOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), kingarthurOld, kingarthurNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the kingarthurDB
	c.JSON(http.StatusOK, kingarthurDB)

	mutexKingArthur.Unlock()
}

// DeleteKingArthur
//
// swagger:route DELETE /kingarthurs/{ID} kingarthurs deleteKingArthur
//
// # Delete a kingarthur
//
// default: genericError
//
//	200: kingarthurDBResponse
func (controller *Controller) DeleteKingArthur(c *gin.Context) {

	mutexKingArthur.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteKingArthur", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKingArthur.GetDB()

	// Get model if exist
	var kingarthurDB orm.KingArthurDB
	if err := db.First(&kingarthurDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&kingarthurDB)

	// get an instance (not staged) from DB instance, and call callback function
	kingarthurDeleted := new(models.KingArthur)
	kingarthurDB.CopyBasicFieldsToKingArthur(kingarthurDeleted)

	// get stage instance from DB instance, and call callback function
	kingarthurStaged := backRepo.BackRepoKingArthur.Map_KingArthurDBID_KingArthurPtr[kingarthurDB.ID]
	if kingarthurStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), kingarthurStaged, kingarthurDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexKingArthur.Unlock()
}
