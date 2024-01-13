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
var __Repository__dummysDeclaration__ models.Repository
var __Repository_time__dummyDeclaration time.Duration

var mutexRepository sync.Mutex

// An RepositoryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRepository updateRepository deleteRepository
type RepositoryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RepositoryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRepository updateRepository
type RepositoryInput struct {
	// The Repository to submit or modify
	// in: body
	Repository *orm.RepositoryAPI
}

// GetRepositorys
//
// swagger:route GET /repositorys repositorys getRepositorys
//
// # Get all repositorys
//
// Responses:
// default: genericError
//
//	200: repositoryDBResponse
func (controller *Controller) GetRepositorys(c *gin.Context) {

	// source slice
	var repositoryDBs []orm.RepositoryDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRepositorys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepository.GetDB()

	query := db.Find(&repositoryDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	repositoryAPIs := make([]orm.RepositoryAPI, 0)

	// for each repository, update fields from the database nullable fields
	for idx := range repositoryDBs {
		repositoryDB := &repositoryDBs[idx]
		_ = repositoryDB
		var repositoryAPI orm.RepositoryAPI

		// insertion point for updating fields
		repositoryAPI.ID = repositoryDB.ID
		repositoryDB.CopyBasicFieldsToRepository_WOP(&repositoryAPI.Repository_WOP)
		repositoryAPI.RepositoryPointersEncoding = repositoryDB.RepositoryPointersEncoding
		repositoryAPIs = append(repositoryAPIs, repositoryAPI)
	}

	c.JSON(http.StatusOK, repositoryAPIs)
}

// PostRepository
//
// swagger:route POST /repositorys repositorys postRepository
//
// Creates a repository
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRepository(c *gin.Context) {

	mutexRepository.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRepositorys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepository.GetDB()

	// Validate input
	var input orm.RepositoryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create repository
	repositoryDB := orm.RepositoryDB{}
	repositoryDB.RepositoryPointersEncoding = input.RepositoryPointersEncoding
	repositoryDB.CopyBasicFieldsFromRepository_WOP(&input.Repository_WOP)

	query := db.Create(&repositoryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRepository.CheckoutPhaseOneInstance(&repositoryDB)
	repository := backRepo.BackRepoRepository.Map_RepositoryDBID_RepositoryPtr[repositoryDB.ID]

	if repository != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), repository)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, repositoryDB)

	mutexRepository.Unlock()
}

// GetRepository
//
// swagger:route GET /repositorys/{ID} repositorys getRepository
//
// Gets the details for a repository.
//
// Responses:
// default: genericError
//
//	200: repositoryDBResponse
func (controller *Controller) GetRepository(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRepository", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepository.GetDB()

	// Get repositoryDB in DB
	var repositoryDB orm.RepositoryDB
	if err := db.First(&repositoryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var repositoryAPI orm.RepositoryAPI
	repositoryAPI.ID = repositoryDB.ID
	repositoryAPI.RepositoryPointersEncoding = repositoryDB.RepositoryPointersEncoding
	repositoryDB.CopyBasicFieldsToRepository_WOP(&repositoryAPI.Repository_WOP)

	c.JSON(http.StatusOK, repositoryAPI)
}

// UpdateRepository
//
// swagger:route PATCH /repositorys/{ID} repositorys updateRepository
//
// # Update a repository
//
// Responses:
// default: genericError
//
//	200: repositoryDBResponse
func (controller *Controller) UpdateRepository(c *gin.Context) {

	mutexRepository.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRepository", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepository.GetDB()

	// Validate input
	var input orm.RepositoryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var repositoryDB orm.RepositoryDB

	// fetch the repository
	query := db.First(&repositoryDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	repositoryDB.CopyBasicFieldsFromRepository_WOP(&input.Repository_WOP)
	repositoryDB.RepositoryPointersEncoding = input.RepositoryPointersEncoding

	query = db.Model(&repositoryDB).Updates(repositoryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	repositoryNew := new(models.Repository)
	repositoryDB.CopyBasicFieldsToRepository(repositoryNew)

	// redeem pointers
	repositoryDB.DecodePointers(backRepo, repositoryNew)

	// get stage instance from DB instance, and call callback function
	repositoryOld := backRepo.BackRepoRepository.Map_RepositoryDBID_RepositoryPtr[repositoryDB.ID]
	if repositoryOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), repositoryOld, repositoryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the repositoryDB
	c.JSON(http.StatusOK, repositoryDB)

	mutexRepository.Unlock()
}

// DeleteRepository
//
// swagger:route DELETE /repositorys/{ID} repositorys deleteRepository
//
// # Delete a repository
//
// default: genericError
//
//	200: repositoryDBResponse
func (controller *Controller) DeleteRepository(c *gin.Context) {

	mutexRepository.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRepository", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRepository.GetDB()

	// Get model if exist
	var repositoryDB orm.RepositoryDB
	if err := db.First(&repositoryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&repositoryDB)

	// get an instance (not staged) from DB instance, and call callback function
	repositoryDeleted := new(models.Repository)
	repositoryDB.CopyBasicFieldsToRepository(repositoryDeleted)

	// get stage instance from DB instance, and call callback function
	repositoryStaged := backRepo.BackRepoRepository.Map_RepositoryDBID_RepositoryPtr[repositoryDB.ID]
	if repositoryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), repositoryStaged, repositoryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexRepository.Unlock()
}
