// do not modify, generated file
package fullstack

import (
	"github.com/thomaspeugeot/thelongbuild/go/controllers"
	"github.com/thomaspeugeot/thelongbuild/go/models"
	"github.com/thomaspeugeot/thelongbuild/go/orm"

	"github.com/gin-gonic/gin"
)

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.StageStruct,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	if stackPath != "" {
		controllers.GetController().AddBackRepo(backRepo, stackPath)
	}

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.AWitch](stage)
	models.SetOrchestratorOnAfterUpdate[models.BlackKnightShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.BringYourDead](stage)
	models.SetOrchestratorOnAfterUpdate[models.Document](stage)
	models.SetOrchestratorOnAfterUpdate[models.DocumentUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.GalahadThePure](stage)
	models.SetOrchestratorOnAfterUpdate[models.GeoObject](stage)
	models.SetOrchestratorOnAfterUpdate[models.GeoObjectUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.Group](stage)
	models.SetOrchestratorOnAfterUpdate[models.GroupUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.KingArthur](stage)
	models.SetOrchestratorOnAfterUpdate[models.KingArthurShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.KnightWhoSayNi](stage)
	models.SetOrchestratorOnAfterUpdate[models.Lancelot](stage)
	models.SetOrchestratorOnAfterUpdate[models.LancelotAgregation](stage)
	models.SetOrchestratorOnAfterUpdate[models.LancelotAgregationUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.LancelotCategory](stage)
	models.SetOrchestratorOnAfterUpdate[models.LancelotCategoryUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.MapObject](stage)
	models.SetOrchestratorOnAfterUpdate[models.MapObjectUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.Repository](stage)
	models.SetOrchestratorOnAfterUpdate[models.SirRobin](stage)
	models.SetOrchestratorOnAfterUpdate[models.TheBridge](stage)
	models.SetOrchestratorOnAfterUpdate[models.TheNuteShape](stage)
	models.SetOrchestratorOnAfterUpdate[models.TheNuteTransition](stage)
	models.SetOrchestratorOnAfterUpdate[models.User](stage)
	models.SetOrchestratorOnAfterUpdate[models.UserUse](stage)
	models.SetOrchestratorOnAfterUpdate[models.WhatIsYourPreferedColor](stage)
	models.SetOrchestratorOnAfterUpdate[models.Workspace](stage)

	return
}
