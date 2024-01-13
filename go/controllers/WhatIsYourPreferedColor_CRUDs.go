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
var __WhatIsYourPreferedColor__dummysDeclaration__ models.WhatIsYourPreferedColor
var __WhatIsYourPreferedColor_time__dummyDeclaration time.Duration

var mutexWhatIsYourPreferedColor sync.Mutex

// An WhatIsYourPreferedColorID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getWhatIsYourPreferedColor updateWhatIsYourPreferedColor deleteWhatIsYourPreferedColor
type WhatIsYourPreferedColorID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// WhatIsYourPreferedColorInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postWhatIsYourPreferedColor updateWhatIsYourPreferedColor
type WhatIsYourPreferedColorInput struct {
	// The WhatIsYourPreferedColor to submit or modify
	// in: body
	WhatIsYourPreferedColor *orm.WhatIsYourPreferedColorAPI
}

// GetWhatIsYourPreferedColors
//
// swagger:route GET /whatisyourpreferedcolors whatisyourpreferedcolors getWhatIsYourPreferedColors
//
// # Get all whatisyourpreferedcolors
//
// Responses:
// default: genericError
//
//	200: whatisyourpreferedcolorDBResponse
func (controller *Controller) GetWhatIsYourPreferedColors(c *gin.Context) {

	// source slice
	var whatisyourpreferedcolorDBs []orm.WhatIsYourPreferedColorDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWhatIsYourPreferedColors", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhatIsYourPreferedColor.GetDB()

	query := db.Find(&whatisyourpreferedcolorDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	whatisyourpreferedcolorAPIs := make([]orm.WhatIsYourPreferedColorAPI, 0)

	// for each whatisyourpreferedcolor, update fields from the database nullable fields
	for idx := range whatisyourpreferedcolorDBs {
		whatisyourpreferedcolorDB := &whatisyourpreferedcolorDBs[idx]
		_ = whatisyourpreferedcolorDB
		var whatisyourpreferedcolorAPI orm.WhatIsYourPreferedColorAPI

		// insertion point for updating fields
		whatisyourpreferedcolorAPI.ID = whatisyourpreferedcolorDB.ID
		whatisyourpreferedcolorDB.CopyBasicFieldsToWhatIsYourPreferedColor_WOP(&whatisyourpreferedcolorAPI.WhatIsYourPreferedColor_WOP)
		whatisyourpreferedcolorAPI.WhatIsYourPreferedColorPointersEncoding = whatisyourpreferedcolorDB.WhatIsYourPreferedColorPointersEncoding
		whatisyourpreferedcolorAPIs = append(whatisyourpreferedcolorAPIs, whatisyourpreferedcolorAPI)
	}

	c.JSON(http.StatusOK, whatisyourpreferedcolorAPIs)
}

// PostWhatIsYourPreferedColor
//
// swagger:route POST /whatisyourpreferedcolors whatisyourpreferedcolors postWhatIsYourPreferedColor
//
// Creates a whatisyourpreferedcolor
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostWhatIsYourPreferedColor(c *gin.Context) {

	mutexWhatIsYourPreferedColor.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostWhatIsYourPreferedColors", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhatIsYourPreferedColor.GetDB()

	// Validate input
	var input orm.WhatIsYourPreferedColorAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create whatisyourpreferedcolor
	whatisyourpreferedcolorDB := orm.WhatIsYourPreferedColorDB{}
	whatisyourpreferedcolorDB.WhatIsYourPreferedColorPointersEncoding = input.WhatIsYourPreferedColorPointersEncoding
	whatisyourpreferedcolorDB.CopyBasicFieldsFromWhatIsYourPreferedColor_WOP(&input.WhatIsYourPreferedColor_WOP)

	query := db.Create(&whatisyourpreferedcolorDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoWhatIsYourPreferedColor.CheckoutPhaseOneInstance(&whatisyourpreferedcolorDB)
	whatisyourpreferedcolor := backRepo.BackRepoWhatIsYourPreferedColor.Map_WhatIsYourPreferedColorDBID_WhatIsYourPreferedColorPtr[whatisyourpreferedcolorDB.ID]

	if whatisyourpreferedcolor != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), whatisyourpreferedcolor)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, whatisyourpreferedcolorDB)

	mutexWhatIsYourPreferedColor.Unlock()
}

// GetWhatIsYourPreferedColor
//
// swagger:route GET /whatisyourpreferedcolors/{ID} whatisyourpreferedcolors getWhatIsYourPreferedColor
//
// Gets the details for a whatisyourpreferedcolor.
//
// Responses:
// default: genericError
//
//	200: whatisyourpreferedcolorDBResponse
func (controller *Controller) GetWhatIsYourPreferedColor(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWhatIsYourPreferedColor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhatIsYourPreferedColor.GetDB()

	// Get whatisyourpreferedcolorDB in DB
	var whatisyourpreferedcolorDB orm.WhatIsYourPreferedColorDB
	if err := db.First(&whatisyourpreferedcolorDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var whatisyourpreferedcolorAPI orm.WhatIsYourPreferedColorAPI
	whatisyourpreferedcolorAPI.ID = whatisyourpreferedcolorDB.ID
	whatisyourpreferedcolorAPI.WhatIsYourPreferedColorPointersEncoding = whatisyourpreferedcolorDB.WhatIsYourPreferedColorPointersEncoding
	whatisyourpreferedcolorDB.CopyBasicFieldsToWhatIsYourPreferedColor_WOP(&whatisyourpreferedcolorAPI.WhatIsYourPreferedColor_WOP)

	c.JSON(http.StatusOK, whatisyourpreferedcolorAPI)
}

// UpdateWhatIsYourPreferedColor
//
// swagger:route PATCH /whatisyourpreferedcolors/{ID} whatisyourpreferedcolors updateWhatIsYourPreferedColor
//
// # Update a whatisyourpreferedcolor
//
// Responses:
// default: genericError
//
//	200: whatisyourpreferedcolorDBResponse
func (controller *Controller) UpdateWhatIsYourPreferedColor(c *gin.Context) {

	mutexWhatIsYourPreferedColor.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateWhatIsYourPreferedColor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhatIsYourPreferedColor.GetDB()

	// Validate input
	var input orm.WhatIsYourPreferedColorAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var whatisyourpreferedcolorDB orm.WhatIsYourPreferedColorDB

	// fetch the whatisyourpreferedcolor
	query := db.First(&whatisyourpreferedcolorDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	whatisyourpreferedcolorDB.CopyBasicFieldsFromWhatIsYourPreferedColor_WOP(&input.WhatIsYourPreferedColor_WOP)
	whatisyourpreferedcolorDB.WhatIsYourPreferedColorPointersEncoding = input.WhatIsYourPreferedColorPointersEncoding

	query = db.Model(&whatisyourpreferedcolorDB).Updates(whatisyourpreferedcolorDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	whatisyourpreferedcolorNew := new(models.WhatIsYourPreferedColor)
	whatisyourpreferedcolorDB.CopyBasicFieldsToWhatIsYourPreferedColor(whatisyourpreferedcolorNew)

	// redeem pointers
	whatisyourpreferedcolorDB.DecodePointers(backRepo, whatisyourpreferedcolorNew)

	// get stage instance from DB instance, and call callback function
	whatisyourpreferedcolorOld := backRepo.BackRepoWhatIsYourPreferedColor.Map_WhatIsYourPreferedColorDBID_WhatIsYourPreferedColorPtr[whatisyourpreferedcolorDB.ID]
	if whatisyourpreferedcolorOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), whatisyourpreferedcolorOld, whatisyourpreferedcolorNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the whatisyourpreferedcolorDB
	c.JSON(http.StatusOK, whatisyourpreferedcolorDB)

	mutexWhatIsYourPreferedColor.Unlock()
}

// DeleteWhatIsYourPreferedColor
//
// swagger:route DELETE /whatisyourpreferedcolors/{ID} whatisyourpreferedcolors deleteWhatIsYourPreferedColor
//
// # Delete a whatisyourpreferedcolor
//
// default: genericError
//
//	200: whatisyourpreferedcolorDBResponse
func (controller *Controller) DeleteWhatIsYourPreferedColor(c *gin.Context) {

	mutexWhatIsYourPreferedColor.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteWhatIsYourPreferedColor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWhatIsYourPreferedColor.GetDB()

	// Get model if exist
	var whatisyourpreferedcolorDB orm.WhatIsYourPreferedColorDB
	if err := db.First(&whatisyourpreferedcolorDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&whatisyourpreferedcolorDB)

	// get an instance (not staged) from DB instance, and call callback function
	whatisyourpreferedcolorDeleted := new(models.WhatIsYourPreferedColor)
	whatisyourpreferedcolorDB.CopyBasicFieldsToWhatIsYourPreferedColor(whatisyourpreferedcolorDeleted)

	// get stage instance from DB instance, and call callback function
	whatisyourpreferedcolorStaged := backRepo.BackRepoWhatIsYourPreferedColor.Map_WhatIsYourPreferedColorDBID_WhatIsYourPreferedColorPtr[whatisyourpreferedcolorDB.ID]
	if whatisyourpreferedcolorStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), whatisyourpreferedcolorStaged, whatisyourpreferedcolorDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexWhatIsYourPreferedColor.Unlock()
}
