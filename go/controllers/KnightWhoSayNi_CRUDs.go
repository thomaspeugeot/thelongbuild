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
var __KnightWhoSayNi__dummysDeclaration__ models.KnightWhoSayNi
var __KnightWhoSayNi_time__dummyDeclaration time.Duration

var mutexKnightWhoSayNi sync.Mutex

// An KnightWhoSayNiID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getKnightWhoSayNi updateKnightWhoSayNi deleteKnightWhoSayNi
type KnightWhoSayNiID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// KnightWhoSayNiInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postKnightWhoSayNi updateKnightWhoSayNi
type KnightWhoSayNiInput struct {
	// The KnightWhoSayNi to submit or modify
	// in: body
	KnightWhoSayNi *orm.KnightWhoSayNiAPI
}

// GetKnightWhoSayNis
//
// swagger:route GET /knightwhosaynis knightwhosaynis getKnightWhoSayNis
//
// # Get all knightwhosaynis
//
// Responses:
// default: genericError
//
//	200: knightwhosayniDBResponse
func (controller *Controller) GetKnightWhoSayNis(c *gin.Context) {

	// source slice
	var knightwhosayniDBs []orm.KnightWhoSayNiDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKnightWhoSayNis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKnightWhoSayNi.GetDB()

	query := db.Find(&knightwhosayniDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	knightwhosayniAPIs := make([]orm.KnightWhoSayNiAPI, 0)

	// for each knightwhosayni, update fields from the database nullable fields
	for idx := range knightwhosayniDBs {
		knightwhosayniDB := &knightwhosayniDBs[idx]
		_ = knightwhosayniDB
		var knightwhosayniAPI orm.KnightWhoSayNiAPI

		// insertion point for updating fields
		knightwhosayniAPI.ID = knightwhosayniDB.ID
		knightwhosayniDB.CopyBasicFieldsToKnightWhoSayNi_WOP(&knightwhosayniAPI.KnightWhoSayNi_WOP)
		knightwhosayniAPI.KnightWhoSayNiPointersEncoding = knightwhosayniDB.KnightWhoSayNiPointersEncoding
		knightwhosayniAPIs = append(knightwhosayniAPIs, knightwhosayniAPI)
	}

	c.JSON(http.StatusOK, knightwhosayniAPIs)
}

// PostKnightWhoSayNi
//
// swagger:route POST /knightwhosaynis knightwhosaynis postKnightWhoSayNi
//
// Creates a knightwhosayni
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostKnightWhoSayNi(c *gin.Context) {

	mutexKnightWhoSayNi.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostKnightWhoSayNis", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKnightWhoSayNi.GetDB()

	// Validate input
	var input orm.KnightWhoSayNiAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create knightwhosayni
	knightwhosayniDB := orm.KnightWhoSayNiDB{}
	knightwhosayniDB.KnightWhoSayNiPointersEncoding = input.KnightWhoSayNiPointersEncoding
	knightwhosayniDB.CopyBasicFieldsFromKnightWhoSayNi_WOP(&input.KnightWhoSayNi_WOP)

	query := db.Create(&knightwhosayniDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoKnightWhoSayNi.CheckoutPhaseOneInstance(&knightwhosayniDB)
	knightwhosayni := backRepo.BackRepoKnightWhoSayNi.Map_KnightWhoSayNiDBID_KnightWhoSayNiPtr[knightwhosayniDB.ID]

	if knightwhosayni != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), knightwhosayni)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, knightwhosayniDB)

	mutexKnightWhoSayNi.Unlock()
}

// GetKnightWhoSayNi
//
// swagger:route GET /knightwhosaynis/{ID} knightwhosaynis getKnightWhoSayNi
//
// Gets the details for a knightwhosayni.
//
// Responses:
// default: genericError
//
//	200: knightwhosayniDBResponse
func (controller *Controller) GetKnightWhoSayNi(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKnightWhoSayNi", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKnightWhoSayNi.GetDB()

	// Get knightwhosayniDB in DB
	var knightwhosayniDB orm.KnightWhoSayNiDB
	if err := db.First(&knightwhosayniDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var knightwhosayniAPI orm.KnightWhoSayNiAPI
	knightwhosayniAPI.ID = knightwhosayniDB.ID
	knightwhosayniAPI.KnightWhoSayNiPointersEncoding = knightwhosayniDB.KnightWhoSayNiPointersEncoding
	knightwhosayniDB.CopyBasicFieldsToKnightWhoSayNi_WOP(&knightwhosayniAPI.KnightWhoSayNi_WOP)

	c.JSON(http.StatusOK, knightwhosayniAPI)
}

// UpdateKnightWhoSayNi
//
// swagger:route PATCH /knightwhosaynis/{ID} knightwhosaynis updateKnightWhoSayNi
//
// # Update a knightwhosayni
//
// Responses:
// default: genericError
//
//	200: knightwhosayniDBResponse
func (controller *Controller) UpdateKnightWhoSayNi(c *gin.Context) {

	mutexKnightWhoSayNi.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateKnightWhoSayNi", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKnightWhoSayNi.GetDB()

	// Validate input
	var input orm.KnightWhoSayNiAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var knightwhosayniDB orm.KnightWhoSayNiDB

	// fetch the knightwhosayni
	query := db.First(&knightwhosayniDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	knightwhosayniDB.CopyBasicFieldsFromKnightWhoSayNi_WOP(&input.KnightWhoSayNi_WOP)
	knightwhosayniDB.KnightWhoSayNiPointersEncoding = input.KnightWhoSayNiPointersEncoding

	query = db.Model(&knightwhosayniDB).Updates(knightwhosayniDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	knightwhosayniNew := new(models.KnightWhoSayNi)
	knightwhosayniDB.CopyBasicFieldsToKnightWhoSayNi(knightwhosayniNew)

	// redeem pointers
	knightwhosayniDB.DecodePointers(backRepo, knightwhosayniNew)

	// get stage instance from DB instance, and call callback function
	knightwhosayniOld := backRepo.BackRepoKnightWhoSayNi.Map_KnightWhoSayNiDBID_KnightWhoSayNiPtr[knightwhosayniDB.ID]
	if knightwhosayniOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), knightwhosayniOld, knightwhosayniNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the knightwhosayniDB
	c.JSON(http.StatusOK, knightwhosayniDB)

	mutexKnightWhoSayNi.Unlock()
}

// DeleteKnightWhoSayNi
//
// swagger:route DELETE /knightwhosaynis/{ID} knightwhosaynis deleteKnightWhoSayNi
//
// # Delete a knightwhosayni
//
// default: genericError
//
//	200: knightwhosayniDBResponse
func (controller *Controller) DeleteKnightWhoSayNi(c *gin.Context) {

	mutexKnightWhoSayNi.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteKnightWhoSayNi", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKnightWhoSayNi.GetDB()

	// Get model if exist
	var knightwhosayniDB orm.KnightWhoSayNiDB
	if err := db.First(&knightwhosayniDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&knightwhosayniDB)

	// get an instance (not staged) from DB instance, and call callback function
	knightwhosayniDeleted := new(models.KnightWhoSayNi)
	knightwhosayniDB.CopyBasicFieldsToKnightWhoSayNi(knightwhosayniDeleted)

	// get stage instance from DB instance, and call callback function
	knightwhosayniStaged := backRepo.BackRepoKnightWhoSayNi.Map_KnightWhoSayNiDBID_KnightWhoSayNiPtr[knightwhosayniDB.ID]
	if knightwhosayniStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), knightwhosayniStaged, knightwhosayniDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexKnightWhoSayNi.Unlock()
}
