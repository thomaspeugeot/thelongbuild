// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
	"github.com/thomaspeugeot/thelongbuild/go/orm"
)

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	log.Println()
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

	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(probe.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		table.Rows = append(table.Rows, row)

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			_ = value
		}
		for _, reverseField := range reverseFields {

			value := orm.GetReverseFieldOwnerName[T](
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
				&reverseField)
			_ = value
		}
	}
}
