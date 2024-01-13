// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.probe = probe
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance *T
	probe    *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.AWitch:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.BlackKnightShape:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.BringYourDead:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Document:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DocumentUse:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GalahadThePure:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GeoObject:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GeoObjectUse:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Group:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GroupUse:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.KingArthur:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.KingArthurShape:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.KnightWhoSayNi:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Lancelot:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.LancelotAgregation:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.LancelotAgregationUse:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.LancelotCategory:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.LancelotCategoryUse:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.MapObject:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.MapObjectUse:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Repository:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SirRobin:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TheBridge:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TheNuteShape:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TheNuteTransition:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.User:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.UserUse:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.WhatIsYourPreferedColor:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Workspace:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}
