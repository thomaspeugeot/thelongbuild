// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/thelongbuild/go/models"
	"github.com/thomaspeugeot/thelongbuild/go/orm"
)

var _ orm.AWitchAPI
var _ gongtable.Cell

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	log.Println()
}

func fillUpTable[T models.Gongstruct](
	probe *Probe,
) {

	// fields := models.GetFields[T]()
	// _ = fields
	// reverseFields := models.GetReverseFields[T]()
	// _ = reverseFields

	// refresh the stage of interest
	probe.stageOfInterest.Checkout()

	// sliceOfGongStructsSorted := make([]*T, 0)

	// for _, structInstance := range sliceOfGongStructsSorted {
	// 	// row := new(gongtable.Row).Stage(probe.tableStage)
	// 	// row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

	// 	log.Println()

	// 	for _, reverseField := range reverseFields {

	// 		value := orm.GetReverseFieldOwnerName[T](
	// 			probe.stageOfInterest,
	// 			probe.backRepoOfInterest,
	// 			structInstance,
	// 			&reverseField)
	// 		_ = value
	// 		log.Println()
	// 	}
	// }
}
