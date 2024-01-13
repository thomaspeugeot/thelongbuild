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
var __User__dummysDeclaration__ models.User
var __User_time__dummyDeclaration time.Duration

var mutexUser sync.Mutex

// An UserID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUser updateUser deleteUser
type UserID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UserInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUser updateUser
type UserInput struct {
	// The User to submit or modify
	// in: body
	User *orm.UserAPI
}

// GetUsers
//
// swagger:route GET /users users getUsers
//
// # Get all users
//
// Responses:
// default: genericError
//
//	200: userDBResponse
func (controller *Controller) GetUsers(c *gin.Context) {

	// source slice
	var userDBs []orm.UserDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUsers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUser.GetDB()

	query := db.Find(&userDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	userAPIs := make([]orm.UserAPI, 0)

	// for each user, update fields from the database nullable fields
	for idx := range userDBs {
		userDB := &userDBs[idx]
		_ = userDB
		var userAPI orm.UserAPI

		// insertion point for updating fields
		userAPI.ID = userDB.ID
		userDB.CopyBasicFieldsToUser_WOP(&userAPI.User_WOP)
		userAPI.UserPointersEncoding = userDB.UserPointersEncoding
		userAPIs = append(userAPIs, userAPI)
	}

	c.JSON(http.StatusOK, userAPIs)
}

// PostUser
//
// swagger:route POST /users users postUser
//
// Creates a user
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostUser(c *gin.Context) {

	mutexUser.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostUsers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUser.GetDB()

	// Validate input
	var input orm.UserAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create user
	userDB := orm.UserDB{}
	userDB.UserPointersEncoding = input.UserPointersEncoding
	userDB.CopyBasicFieldsFromUser_WOP(&input.User_WOP)

	query := db.Create(&userDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoUser.CheckoutPhaseOneInstance(&userDB)
	user := backRepo.BackRepoUser.Map_UserDBID_UserPtr[userDB.ID]

	if user != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), user)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, userDB)

	mutexUser.Unlock()
}

// GetUser
//
// swagger:route GET /users/{ID} users getUser
//
// Gets the details for a user.
//
// Responses:
// default: genericError
//
//	200: userDBResponse
func (controller *Controller) GetUser(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetUser", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUser.GetDB()

	// Get userDB in DB
	var userDB orm.UserDB
	if err := db.First(&userDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var userAPI orm.UserAPI
	userAPI.ID = userDB.ID
	userAPI.UserPointersEncoding = userDB.UserPointersEncoding
	userDB.CopyBasicFieldsToUser_WOP(&userAPI.User_WOP)

	c.JSON(http.StatusOK, userAPI)
}

// UpdateUser
//
// swagger:route PATCH /users/{ID} users updateUser
//
// # Update a user
//
// Responses:
// default: genericError
//
//	200: userDBResponse
func (controller *Controller) UpdateUser(c *gin.Context) {

	mutexUser.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateUser", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUser.GetDB()

	// Validate input
	var input orm.UserAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var userDB orm.UserDB

	// fetch the user
	query := db.First(&userDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	userDB.CopyBasicFieldsFromUser_WOP(&input.User_WOP)
	userDB.UserPointersEncoding = input.UserPointersEncoding

	query = db.Model(&userDB).Updates(userDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	userNew := new(models.User)
	userDB.CopyBasicFieldsToUser(userNew)

	// redeem pointers
	userDB.DecodePointers(backRepo, userNew)

	// get stage instance from DB instance, and call callback function
	userOld := backRepo.BackRepoUser.Map_UserDBID_UserPtr[userDB.ID]
	if userOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), userOld, userNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the userDB
	c.JSON(http.StatusOK, userDB)

	mutexUser.Unlock()
}

// DeleteUser
//
// swagger:route DELETE /users/{ID} users deleteUser
//
// # Delete a user
//
// default: genericError
//
//	200: userDBResponse
func (controller *Controller) DeleteUser(c *gin.Context) {

	mutexUser.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteUser", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoUser.GetDB()

	// Get model if exist
	var userDB orm.UserDB
	if err := db.First(&userDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&userDB)

	// get an instance (not staged) from DB instance, and call callback function
	userDeleted := new(models.User)
	userDB.CopyBasicFieldsToUser(userDeleted)

	// get stage instance from DB instance, and call callback function
	userStaged := backRepo.BackRepoUser.Map_UserDBID_UserPtr[userDB.ID]
	if userStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), userStaged, userDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexUser.Unlock()
}
