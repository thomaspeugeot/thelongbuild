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
var __GroupUse__dummysDeclaration__ models.GroupUse
var __GroupUse_time__dummyDeclaration time.Duration

var mutexGroupUse sync.Mutex

// An GroupUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGroupUse updateGroupUse deleteGroupUse
type GroupUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GroupUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGroupUse updateGroupUse
type GroupUseInput struct {
	// The GroupUse to submit or modify
	// in: body
	GroupUse *orm.GroupUseAPI
}

// GetGroupUses
//
// swagger:route GET /groupuses groupuses getGroupUses
//
// # Get all groupuses
//
// Responses:
// default: genericError
//
//	200: groupuseDBResponse
func (controller *Controller) GetGroupUses(c *gin.Context) {

	// source slice
	var groupuseDBs []orm.GroupUseDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroupUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroupUse.GetDB()

	query := db.Find(&groupuseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	groupuseAPIs := make([]orm.GroupUseAPI, 0)

	// for each groupuse, update fields from the database nullable fields
	for idx := range groupuseDBs {
		groupuseDB := &groupuseDBs[idx]
		_ = groupuseDB
		var groupuseAPI orm.GroupUseAPI

		// insertion point for updating fields
		groupuseAPI.ID = groupuseDB.ID
		groupuseDB.CopyBasicFieldsToGroupUse_WOP(&groupuseAPI.GroupUse_WOP)
		groupuseAPI.GroupUsePointersEncoding = groupuseDB.GroupUsePointersEncoding
		groupuseAPIs = append(groupuseAPIs, groupuseAPI)
	}

	c.JSON(http.StatusOK, groupuseAPIs)
}

// PostGroupUse
//
// swagger:route POST /groupuses groupuses postGroupUse
//
// Creates a groupuse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGroupUse(c *gin.Context) {

	mutexGroupUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGroupUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroupUse.GetDB()

	// Validate input
	var input orm.GroupUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create groupuse
	groupuseDB := orm.GroupUseDB{}
	groupuseDB.GroupUsePointersEncoding = input.GroupUsePointersEncoding
	groupuseDB.CopyBasicFieldsFromGroupUse_WOP(&input.GroupUse_WOP)

	query := db.Create(&groupuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGroupUse.CheckoutPhaseOneInstance(&groupuseDB)
	groupuse := backRepo.BackRepoGroupUse.Map_GroupUseDBID_GroupUsePtr[groupuseDB.ID]

	if groupuse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), groupuse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, groupuseDB)

	mutexGroupUse.Unlock()
}

// GetGroupUse
//
// swagger:route GET /groupuses/{ID} groupuses getGroupUse
//
// Gets the details for a groupuse.
//
// Responses:
// default: genericError
//
//	200: groupuseDBResponse
func (controller *Controller) GetGroupUse(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroupUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroupUse.GetDB()

	// Get groupuseDB in DB
	var groupuseDB orm.GroupUseDB
	if err := db.First(&groupuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var groupuseAPI orm.GroupUseAPI
	groupuseAPI.ID = groupuseDB.ID
	groupuseAPI.GroupUsePointersEncoding = groupuseDB.GroupUsePointersEncoding
	groupuseDB.CopyBasicFieldsToGroupUse_WOP(&groupuseAPI.GroupUse_WOP)

	c.JSON(http.StatusOK, groupuseAPI)
}

// UpdateGroupUse
//
// swagger:route PATCH /groupuses/{ID} groupuses updateGroupUse
//
// # Update a groupuse
//
// Responses:
// default: genericError
//
//	200: groupuseDBResponse
func (controller *Controller) UpdateGroupUse(c *gin.Context) {

	mutexGroupUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGroupUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroupUse.GetDB()

	// Validate input
	var input orm.GroupUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var groupuseDB orm.GroupUseDB

	// fetch the groupuse
	query := db.First(&groupuseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	groupuseDB.CopyBasicFieldsFromGroupUse_WOP(&input.GroupUse_WOP)
	groupuseDB.GroupUsePointersEncoding = input.GroupUsePointersEncoding

	query = db.Model(&groupuseDB).Updates(groupuseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	groupuseNew := new(models.GroupUse)
	groupuseDB.CopyBasicFieldsToGroupUse(groupuseNew)

	// redeem pointers
	groupuseDB.DecodePointers(backRepo, groupuseNew)

	// get stage instance from DB instance, and call callback function
	groupuseOld := backRepo.BackRepoGroupUse.Map_GroupUseDBID_GroupUsePtr[groupuseDB.ID]
	if groupuseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), groupuseOld, groupuseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the groupuseDB
	c.JSON(http.StatusOK, groupuseDB)

	mutexGroupUse.Unlock()
}

// DeleteGroupUse
//
// swagger:route DELETE /groupuses/{ID} groupuses deleteGroupUse
//
// # Delete a groupuse
//
// default: genericError
//
//	200: groupuseDBResponse
func (controller *Controller) DeleteGroupUse(c *gin.Context) {

	mutexGroupUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGroupUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroupUse.GetDB()

	// Get model if exist
	var groupuseDB orm.GroupUseDB
	if err := db.First(&groupuseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&groupuseDB)

	// get an instance (not staged) from DB instance, and call callback function
	groupuseDeleted := new(models.GroupUse)
	groupuseDB.CopyBasicFieldsToGroupUse(groupuseDeleted)

	// get stage instance from DB instance, and call callback function
	groupuseStaged := backRepo.BackRepoGroupUse.Map_GroupUseDBID_GroupUsePtr[groupuseDB.ID]
	if groupuseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), groupuseStaged, groupuseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexGroupUse.Unlock()
}
