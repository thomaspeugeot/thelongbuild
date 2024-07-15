// generated code - do not edit
package probe

import (
	"fmt"
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/thomaspeugeot/thelongbuild/go/models"
	"github.com/thomaspeugeot/thelongbuild/go/orm"
)

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.AWitch:
		fillUpTable[models.AWitch](probe)
	case *models.BlackKnightShape:
		fillUpTable[models.BlackKnightShape](probe)
	case *models.BringYourDead:
		fillUpTable[models.BringYourDead](probe)
	case *models.Document:
		fillUpTable[models.Document](probe)
	case *models.DocumentUse:
		fillUpTable[models.DocumentUse](probe)
	case *models.GalahadThePure:
		fillUpTable[models.GalahadThePure](probe)
	case *models.GeoObject:
		fillUpTable[models.GeoObject](probe)
	case *models.GeoObjectUse:
		fillUpTable[models.GeoObjectUse](probe)
	case *models.Group:
		fillUpTable[models.Group](probe)
	case *models.GroupUse:
		fillUpTable[models.GroupUse](probe)
	case *models.KingArthur:
		fillUpTable[models.KingArthur](probe)
	case *models.KingArthurShape:
		fillUpTable[models.KingArthurShape](probe)
	case *models.KnightWhoSayNi:
		fillUpTable[models.KnightWhoSayNi](probe)
	case *models.Lancelot:
		fillUpTable[models.Lancelot](probe)
	case *models.LancelotAgregation:
		fillUpTable[models.LancelotAgregation](probe)
	case *models.LancelotAgregationUse:
		fillUpTable[models.LancelotAgregationUse](probe)
	case *models.LancelotCategory:
		fillUpTable[models.LancelotCategory](probe)
	case *models.LancelotCategoryUse:
		fillUpTable[models.LancelotCategoryUse](probe)
	case *models.MapObject:
		fillUpTable[models.MapObject](probe)
	case *models.MapObjectUse:
		fillUpTable[models.MapObjectUse](probe)
	case *models.Repository:
		fillUpTable[models.Repository](probe)
	case *models.SirRobin:
		fillUpTable[models.SirRobin](probe)
	case *models.TheBridge:
		fillUpTable[models.TheBridge](probe)
	case *models.TheNuteShape:
		fillUpTable[models.TheNuteShape](probe)
	case *models.TheNuteTransition:
		fillUpTable[models.TheNuteTransition](probe)
	case *models.User:
		fillUpTable[models.User](probe)
	case *models.UserUse:
		fillUpTable[models.UserUse](probe)
	case *models.WhatIsYourPreferedColor:
		fillUpTable[models.WhatIsYourPreferedColor](probe)
	case *models.Workspace:
		fillUpTable[models.Workspace](probe)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	probe *Probe,
) {

	table := new(gongtable.Table).Stage(probe.tableStage)

	fields := models.GetFields[T]()
	reverseFields := models.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	probe.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](probe.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}

	column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(probe.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
			),
		}).Stage(probe.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(probe.tableStage)

		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
		for _, reverseField := range reverseFields {

			value := orm.GetReverseFieldOwnerName[T](
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
				&reverseField)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
	}
}
