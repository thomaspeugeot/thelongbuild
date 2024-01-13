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
var __Workspace__dummysDeclaration__ models.Workspace
var __Workspace_time__dummyDeclaration time.Duration

var mutexWorkspace sync.Mutex

// An WorkspaceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getWorkspace updateWorkspace deleteWorkspace
type WorkspaceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// WorkspaceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postWorkspace updateWorkspace
type WorkspaceInput struct {
	// The Workspace to submit or modify
	// in: body
	Workspace *orm.WorkspaceAPI
}

// GetWorkspaces
//
// swagger:route GET /workspaces workspaces getWorkspaces
//
// # Get all workspaces
//
// Responses:
// default: genericError
//
//	200: workspaceDBResponse
func (controller *Controller) GetWorkspaces(c *gin.Context) {

	// source slice
	var workspaceDBs []orm.WorkspaceDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWorkspaces", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWorkspace.GetDB()

	query := db.Find(&workspaceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	workspaceAPIs := make([]orm.WorkspaceAPI, 0)

	// for each workspace, update fields from the database nullable fields
	for idx := range workspaceDBs {
		workspaceDB := &workspaceDBs[idx]
		_ = workspaceDB
		var workspaceAPI orm.WorkspaceAPI

		// insertion point for updating fields
		workspaceAPI.ID = workspaceDB.ID
		workspaceDB.CopyBasicFieldsToWorkspace_WOP(&workspaceAPI.Workspace_WOP)
		workspaceAPI.WorkspacePointersEncoding = workspaceDB.WorkspacePointersEncoding
		workspaceAPIs = append(workspaceAPIs, workspaceAPI)
	}

	c.JSON(http.StatusOK, workspaceAPIs)
}

// PostWorkspace
//
// swagger:route POST /workspaces workspaces postWorkspace
//
// Creates a workspace
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostWorkspace(c *gin.Context) {

	mutexWorkspace.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostWorkspaces", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWorkspace.GetDB()

	// Validate input
	var input orm.WorkspaceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create workspace
	workspaceDB := orm.WorkspaceDB{}
	workspaceDB.WorkspacePointersEncoding = input.WorkspacePointersEncoding
	workspaceDB.CopyBasicFieldsFromWorkspace_WOP(&input.Workspace_WOP)

	query := db.Create(&workspaceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoWorkspace.CheckoutPhaseOneInstance(&workspaceDB)
	workspace := backRepo.BackRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[workspaceDB.ID]

	if workspace != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), workspace)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, workspaceDB)

	mutexWorkspace.Unlock()
}

// GetWorkspace
//
// swagger:route GET /workspaces/{ID} workspaces getWorkspace
//
// Gets the details for a workspace.
//
// Responses:
// default: genericError
//
//	200: workspaceDBResponse
func (controller *Controller) GetWorkspace(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWorkspace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWorkspace.GetDB()

	// Get workspaceDB in DB
	var workspaceDB orm.WorkspaceDB
	if err := db.First(&workspaceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var workspaceAPI orm.WorkspaceAPI
	workspaceAPI.ID = workspaceDB.ID
	workspaceAPI.WorkspacePointersEncoding = workspaceDB.WorkspacePointersEncoding
	workspaceDB.CopyBasicFieldsToWorkspace_WOP(&workspaceAPI.Workspace_WOP)

	c.JSON(http.StatusOK, workspaceAPI)
}

// UpdateWorkspace
//
// swagger:route PATCH /workspaces/{ID} workspaces updateWorkspace
//
// # Update a workspace
//
// Responses:
// default: genericError
//
//	200: workspaceDBResponse
func (controller *Controller) UpdateWorkspace(c *gin.Context) {

	mutexWorkspace.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateWorkspace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWorkspace.GetDB()

	// Validate input
	var input orm.WorkspaceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var workspaceDB orm.WorkspaceDB

	// fetch the workspace
	query := db.First(&workspaceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	workspaceDB.CopyBasicFieldsFromWorkspace_WOP(&input.Workspace_WOP)
	workspaceDB.WorkspacePointersEncoding = input.WorkspacePointersEncoding

	query = db.Model(&workspaceDB).Updates(workspaceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	workspaceNew := new(models.Workspace)
	workspaceDB.CopyBasicFieldsToWorkspace(workspaceNew)

	// redeem pointers
	workspaceDB.DecodePointers(backRepo, workspaceNew)

	// get stage instance from DB instance, and call callback function
	workspaceOld := backRepo.BackRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[workspaceDB.ID]
	if workspaceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), workspaceOld, workspaceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the workspaceDB
	c.JSON(http.StatusOK, workspaceDB)

	mutexWorkspace.Unlock()
}

// DeleteWorkspace
//
// swagger:route DELETE /workspaces/{ID} workspaces deleteWorkspace
//
// # Delete a workspace
//
// default: genericError
//
//	200: workspaceDBResponse
func (controller *Controller) DeleteWorkspace(c *gin.Context) {

	mutexWorkspace.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteWorkspace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWorkspace.GetDB()

	// Get model if exist
	var workspaceDB orm.WorkspaceDB
	if err := db.First(&workspaceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&workspaceDB)

	// get an instance (not staged) from DB instance, and call callback function
	workspaceDeleted := new(models.Workspace)
	workspaceDB.CopyBasicFieldsToWorkspace(workspaceDeleted)

	// get stage instance from DB instance, and call callback function
	workspaceStaged := backRepo.BackRepoWorkspace.Map_WorkspaceDBID_WorkspacePtr[workspaceDB.ID]
	if workspaceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), workspaceStaged, workspaceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexWorkspace.Unlock()
}
