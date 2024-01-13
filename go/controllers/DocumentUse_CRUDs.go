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
var __DocumentUse__dummysDeclaration__ models.DocumentUse
var __DocumentUse_time__dummyDeclaration time.Duration

var mutexDocumentUse sync.Mutex

// An DocumentUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDocumentUse updateDocumentUse deleteDocumentUse
type DocumentUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DocumentUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDocumentUse updateDocumentUse
type DocumentUseInput struct {
	// The DocumentUse to submit or modify
	// in: body
	DocumentUse *orm.DocumentUseAPI
}

// GetDocumentUses
//
// swagger:route GET /documentuses documentuses getDocumentUses
//
// # Get all documentuses
//
// Responses:
// default: genericError
//
//	200: documentuseDBResponse
func (controller *Controller) GetDocumentUses(c *gin.Context) {

	// source slice
	var documentuseDBs []orm.DocumentUseDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocumentUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentUse.GetDB()

	query := db.Find(&documentuseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	documentuseAPIs := make([]orm.DocumentUseAPI, 0)

	// for each documentuse, update fields from the database nullable fields
	for idx := range documentuseDBs {
		documentuseDB := &documentuseDBs[idx]
		_ = documentuseDB
		var documentuseAPI orm.DocumentUseAPI

		// insertion point for updating fields
		documentuseAPI.ID = documentuseDB.ID
		documentuseDB.CopyBasicFieldsToDocumentUse_WOP(&documentuseAPI.DocumentUse_WOP)
		documentuseAPI.DocumentUsePointersEncoding = documentuseDB.DocumentUsePointersEncoding
		documentuseAPIs = append(documentuseAPIs, documentuseAPI)
	}

	c.JSON(http.StatusOK, documentuseAPIs)
}

// PostDocumentUse
//
// swagger:route POST /documentuses documentuses postDocumentUse
//
// Creates a documentuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDocumentUse(c *gin.Context) {

	mutexDocumentUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDocumentUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentUse.GetDB()

	// Validate input
	var input orm.DocumentUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create documentuse
	documentuseDB := orm.DocumentUseDB{}
	documentuseDB.DocumentUsePointersEncoding = input.DocumentUsePointersEncoding
	documentuseDB.CopyBasicFieldsFromDocumentUse_WOP(&input.DocumentUse_WOP)

	query := db.Create(&documentuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDocumentUse.CheckoutPhaseOneInstance(&documentuseDB)
	documentuse := backRepo.BackRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[documentuseDB.ID]

	if documentuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), documentuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, documentuseDB)

	mutexDocumentUse.Unlock()
}

// GetDocumentUse
//
// swagger:route GET /documentuses/{ID} documentuses getDocumentUse
//
// Gets the details for a documentuse.
//
// Responses:
// default: genericError
//
//	200: documentuseDBResponse
func (controller *Controller) GetDocumentUse(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocumentUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentUse.GetDB()

	// Get documentuseDB in DB
	var documentuseDB orm.DocumentUseDB
	if err := db.First(&documentuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var documentuseAPI orm.DocumentUseAPI
	documentuseAPI.ID = documentuseDB.ID
	documentuseAPI.DocumentUsePointersEncoding = documentuseDB.DocumentUsePointersEncoding
	documentuseDB.CopyBasicFieldsToDocumentUse_WOP(&documentuseAPI.DocumentUse_WOP)

	c.JSON(http.StatusOK, documentuseAPI)
}

// UpdateDocumentUse
//
// swagger:route PATCH /documentuses/{ID} documentuses updateDocumentUse
//
// # Update a documentuse
//
// Responses:
// default: genericError
//
//	200: documentuseDBResponse
func (controller *Controller) UpdateDocumentUse(c *gin.Context) {

	mutexDocumentUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDocumentUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentUse.GetDB()

	// Validate input
	var input orm.DocumentUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var documentuseDB orm.DocumentUseDB

	// fetch the documentuse
	query := db.First(&documentuseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	documentuseDB.CopyBasicFieldsFromDocumentUse_WOP(&input.DocumentUse_WOP)
	documentuseDB.DocumentUsePointersEncoding = input.DocumentUsePointersEncoding

	query = db.Model(&documentuseDB).Updates(documentuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	documentuseNew := new(models.DocumentUse)
	documentuseDB.CopyBasicFieldsToDocumentUse(documentuseNew)

	// redeem pointers
	documentuseDB.DecodePointers(backRepo, documentuseNew)

	// get stage instance from DB instance, and call callback function
	documentuseOld := backRepo.BackRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[documentuseDB.ID]
	if documentuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), documentuseOld, documentuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the documentuseDB
	c.JSON(http.StatusOK, documentuseDB)

	mutexDocumentUse.Unlock()
}

// DeleteDocumentUse
//
// swagger:route DELETE /documentuses/{ID} documentuses deleteDocumentUse
//
// # Delete a documentuse
//
// default: genericError
//
//	200: documentuseDBResponse
func (controller *Controller) DeleteDocumentUse(c *gin.Context) {

	mutexDocumentUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDocumentUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocumentUse.GetDB()

	// Get model if exist
	var documentuseDB orm.DocumentUseDB
	if err := db.First(&documentuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&documentuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	documentuseDeleted := new(models.DocumentUse)
	documentuseDB.CopyBasicFieldsToDocumentUse(documentuseDeleted)

	// get stage instance from DB instance, and call callback function
	documentuseStaged := backRepo.BackRepoDocumentUse.Map_DocumentUseDBID_DocumentUsePtr[documentuseDB.ID]
	if documentuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), documentuseStaged, documentuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexDocumentUse.Unlock()
}
