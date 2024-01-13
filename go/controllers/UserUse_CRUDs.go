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
var __UserUse__dummysDeclaration__ models.UserUse
var __UserUse_time__dummyDeclaration time.Duration

var mutexUserUse sync.Mutex

// An UserUseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUserUse updateUserUse deleteUserUse
type UserUseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UserUseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUserUse updateUserUse
type UserUseInput struct {
	// The UserUse to submit or modify
	// in: body
	UserUse *orm.UserUseAPI
}

// GetUserUses
//
// swagger:route GET /useruses useruses getUserUses
//
// # Get all useruses
//
// Responses:
// default: genericError
//
//	200: useruseDBResponse
func (controller *Controller) GetUserUses(c *gin.Context) {

	// source slice
	var useruseDBs []orm.UserUseDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUserUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserUse.GetDB()

	query := db.Find(&useruseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	useruseAPIs := make([]orm.UserUseAPI, 0)

	// for each useruse, update fields from the database nullable fields
	for idx := range useruseDBs {
		useruseDB := &useruseDBs[idx]
		_ = useruseDB
		var useruseAPI orm.UserUseAPI

		// insertion point for updating fields
		useruseAPI.ID = useruseDB.ID
		useruseDB.CopyBasicFieldsToUserUse_WOP(&useruseAPI.UserUse_WOP)
		useruseAPI.UserUsePointersEncoding = useruseDB.UserUsePointersEncoding
		useruseAPIs = append(useruseAPIs, useruseAPI)
	}

	c.JSON(http.StatusOK, useruseAPIs)
}

// PostUserUse
//
// swagger:route POST /useruses useruses postUserUse
//
// Creates a useruse
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostUserUse(c *gin.Context) {

	mutexUserUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostUserUses", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserUse.GetDB()

	// Validate input
	var input orm.UserUseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create useruse
	useruseDB := orm.UserUseDB{}
	useruseDB.UserUsePointersEncoding = input.UserUsePointersEncoding
	useruseDB.CopyBasicFieldsFromUserUse_WOP(&input.UserUse_WOP)

	query := db.Create(&useruseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoUserUse.CheckoutPhaseOneInstance(&useruseDB)
	useruse := backRepo.BackRepoUserUse.Map_UserUseDBID_UserUsePtr[useruseDB.ID]

	if useruse != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), useruse)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, useruseDB)

	mutexUserUse.Unlock()
}

// GetUserUse
//
// swagger:route GET /useruses/{ID} useruses getUserUse
//
// Gets the details for a useruse.
//
// Responses:
// default: genericError
//
//	200: useruseDBResponse
func (controller *Controller) GetUserUse(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUserUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserUse.GetDB()

	// Get useruseDB in DB
	var useruseDB orm.UserUseDB
	if err := db.First(&useruseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var useruseAPI orm.UserUseAPI
	useruseAPI.ID = useruseDB.ID
	useruseAPI.UserUsePointersEncoding = useruseDB.UserUsePointersEncoding
	useruseDB.CopyBasicFieldsToUserUse_WOP(&useruseAPI.UserUse_WOP)

	c.JSON(http.StatusOK, useruseAPI)
}

// UpdateUserUse
//
// swagger:route PATCH /useruses/{ID} useruses updateUserUse
//
// # Update a useruse
//
// Responses:
// default: genericError
//
//	200: useruseDBResponse
func (controller *Controller) UpdateUserUse(c *gin.Context) {

	mutexUserUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateUserUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserUse.GetDB()

	// Validate input
	var input orm.UserUseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var useruseDB orm.UserUseDB

	// fetch the useruse
	query := db.First(&useruseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	useruseDB.CopyBasicFieldsFromUserUse_WOP(&input.UserUse_WOP)
	useruseDB.UserUsePointersEncoding = input.UserUsePointersEncoding

	query = db.Model(&useruseDB).Updates(useruseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	useruseNew := new(models.UserUse)
	useruseDB.CopyBasicFieldsToUserUse(useruseNew)

	// redeem pointers
	useruseDB.DecodePointers(backRepo, useruseNew)

	// get stage instance from DB instance, and call callback function
	useruseOld := backRepo.BackRepoUserUse.Map_UserUseDBID_UserUsePtr[useruseDB.ID]
	if useruseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), useruseOld, useruseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the useruseDB
	c.JSON(http.StatusOK, useruseDB)

	mutexUserUse.Unlock()
}

// DeleteUserUse
//
// swagger:route DELETE /useruses/{ID} useruses deleteUserUse
//
// # Delete a useruse
//
// default: genericError
//
//	200: useruseDBResponse
func (controller *Controller) DeleteUserUse(c *gin.Context) {

	mutexUserUse.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteUserUse", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUserUse.GetDB()

	// Get model if exist
	var useruseDB orm.UserUseDB
	if err := db.First(&useruseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&useruseDB)

	// get an instance (not staged) from DB instance, and call callback function
	useruseDeleted := new(models.UserUse)
	useruseDB.CopyBasicFieldsToUserUse(useruseDeleted)

	// get stage instance from DB instance, and call callback function
	useruseStaged := backRepo.BackRepoUserUse.Map_UserUseDBID_UserUsePtr[useruseDB.ID]
	if useruseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), useruseStaged, useruseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexUserUse.Unlock()
}
