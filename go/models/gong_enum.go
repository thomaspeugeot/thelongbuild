// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for DirectionType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (directiontype DirectionType) ToString() (res string) {

	// migration of former implementation of enum
	switch directiontype {
	// insertion code per enum code
	case DIRECTION_UP:
		res = "DIRECTION_UP"
	case DIRECTION_DOWN:
		res = "DIRECTION_DOWN"
	}
	return
}

func (directiontype *DirectionType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DIRECTION_UP":
		*directiontype = DIRECTION_UP
	case "DIRECTION_DOWN":
		*directiontype = DIRECTION_DOWN
	default:
		return errUnkownEnum
	}
	return
}

func (directiontype *DirectionType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DIRECTION_UP":
		*directiontype = DIRECTION_UP
	case "DIRECTION_DOWN":
		*directiontype = DIRECTION_DOWN
	default:
		return errUnkownEnum
	}
	return
}

func (directiontype *DirectionType) ToCodeString() (res string) {

	switch *directiontype {
	// insertion code per enum code
	case DIRECTION_UP:
		res = "DIRECTION_UP"
	case DIRECTION_DOWN:
		res = "DIRECTION_DOWN"
	}
	return
}

func (directiontype DirectionType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DIRECTION_UP")
	res = append(res, "DIRECTION_DOWN")

	return
}

func (directiontype DirectionType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DIRECTION_UP")
	res = append(res, "DIRECTION_DOWN")

	return
}

// Utility function for FormNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (formnames FormNames) ToString() (res string) {

	// migration of former implementation of enum
	switch formnames {
	// insertion code per enum code
	case DefaultForm:
		res = "default form"
	case ModelForm:
		res = "Model Form"
	case ShapeForm:
		res = "Shape Form"
	}
	return
}

func (formnames *FormNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "default form":
		*formnames = DefaultForm
	case "Model Form":
		*formnames = ModelForm
	case "Shape Form":
		*formnames = ShapeForm
	default:
		return errUnkownEnum
	}
	return
}

func (formnames *FormNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "DefaultForm":
		*formnames = DefaultForm
	case "ModelForm":
		*formnames = ModelForm
	case "ShapeForm":
		*formnames = ShapeForm
	default:
		return errUnkownEnum
	}
	return
}

func (formnames *FormNames) ToCodeString() (res string) {

	switch *formnames {
	// insertion code per enum code
	case DefaultForm:
		res = "DefaultForm"
	case ModelForm:
		res = "ModelForm"
	case ShapeForm:
		res = "ShapeForm"
	}
	return
}

func (formnames FormNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "DefaultForm")
	res = append(res, "ModelForm")
	res = append(res, "ShapeForm")

	return
}

func (formnames FormNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "default form")
	res = append(res, "Model Form")
	res = append(res, "Shape Form")

	return
}

// Utility function for OrientationType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (orientationtype OrientationType) ToString() (res string) {

	// migration of former implementation of enum
	switch orientationtype {
	// insertion code per enum code
	case ORIENTATION_HORIZONTAL:
		res = "ORIENTATION_HORIZONTAL"
	case ORIENTATION_VERTICAL:
		res = "ORIENTATION_VERTICAL"
	}
	return
}

func (orientationtype *OrientationType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ORIENTATION_HORIZONTAL":
		*orientationtype = ORIENTATION_HORIZONTAL
	case "ORIENTATION_VERTICAL":
		*orientationtype = ORIENTATION_VERTICAL
	default:
		return errUnkownEnum
	}
	return
}

func (orientationtype *OrientationType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ORIENTATION_HORIZONTAL":
		*orientationtype = ORIENTATION_HORIZONTAL
	case "ORIENTATION_VERTICAL":
		*orientationtype = ORIENTATION_VERTICAL
	default:
		return errUnkownEnum
	}
	return
}

func (orientationtype *OrientationType) ToCodeString() (res string) {

	switch *orientationtype {
	// insertion code per enum code
	case ORIENTATION_HORIZONTAL:
		res = "ORIENTATION_HORIZONTAL"
	case ORIENTATION_VERTICAL:
		res = "ORIENTATION_VERTICAL"
	}
	return
}

func (orientationtype OrientationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ORIENTATION_HORIZONTAL")
	res = append(res, "ORIENTATION_VERTICAL")

	return
}

func (orientationtype OrientationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ORIENTATION_HORIZONTAL")
	res = append(res, "ORIENTATION_VERTICAL")

	return
}

// Utility function for ProbabilityEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (probabilityenum ProbabilityEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch probabilityenum {
	// insertion code per enum code
	case PERCENT_100:
		res = "100 %"
	case PERCENT_75:
		res = "75 %"
	case PERCENT_50:
		res = "50 %"
	case PERCENT_25:
		res = "25 %"
	case PERCENT_00:
		res = "0 %"
	}
	return
}

func (probabilityenum *ProbabilityEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "100 %":
		*probabilityenum = PERCENT_100
	case "75 %":
		*probabilityenum = PERCENT_75
	case "50 %":
		*probabilityenum = PERCENT_50
	case "25 %":
		*probabilityenum = PERCENT_25
	case "0 %":
		*probabilityenum = PERCENT_00
	default:
		return errUnkownEnum
	}
	return
}

func (probabilityenum *ProbabilityEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PERCENT_100":
		*probabilityenum = PERCENT_100
	case "PERCENT_75":
		*probabilityenum = PERCENT_75
	case "PERCENT_50":
		*probabilityenum = PERCENT_50
	case "PERCENT_25":
		*probabilityenum = PERCENT_25
	case "PERCENT_00":
		*probabilityenum = PERCENT_00
	default:
		return errUnkownEnum
	}
	return
}

func (probabilityenum *ProbabilityEnum) ToCodeString() (res string) {

	switch *probabilityenum {
	// insertion code per enum code
	case PERCENT_100:
		res = "PERCENT_100"
	case PERCENT_75:
		res = "PERCENT_75"
	case PERCENT_50:
		res = "PERCENT_50"
	case PERCENT_25:
		res = "PERCENT_25"
	case PERCENT_00:
		res = "PERCENT_00"
	}
	return
}

func (probabilityenum ProbabilityEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PERCENT_100")
	res = append(res, "PERCENT_75")
	res = append(res, "PERCENT_50")
	res = append(res, "PERCENT_25")
	res = append(res, "PERCENT_00")

	return
}

func (probabilityenum ProbabilityEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "100 %")
	res = append(res, "75 %")
	res = append(res, "50 %")
	res = append(res, "25 %")
	res = append(res, "0 %")

	return
}

// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case WhatStackName:
		res = "what"
	case GongsvgStackName:
		res = "gongsvg"
	case GongtreeStackName:
		res = "gongtree"
	case GongtableStackName:
		res = "gongtable"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "what":
		*stacksnames = WhatStackName
	case "gongsvg":
		*stacksnames = GongsvgStackName
	case "gongtree":
		*stacksnames = GongtreeStackName
	case "gongtable":
		*stacksnames = GongtableStackName
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "WhatStackName":
		*stacksnames = WhatStackName
	case "GongsvgStackName":
		*stacksnames = GongsvgStackName
	case "GongtreeStackName":
		*stacksnames = GongtreeStackName
	case "GongtableStackName":
		*stacksnames = GongtableStackName
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case WhatStackName:
		res = "WhatStackName"
	case GongsvgStackName:
		res = "GongsvgStackName"
	case GongtreeStackName:
		res = "GongtreeStackName"
	case GongtableStackName:
		res = "GongtableStackName"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "WhatStackName")
	res = append(res, "GongsvgStackName")
	res = append(res, "GongtreeStackName")
	res = append(res, "GongtableStackName")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "what")
	res = append(res, "gongsvg")
	res = append(res, "gongtree")
	res = append(res, "gongtable")

	return
}

// Utility function for TreeNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (treenames TreeNames) ToString() (res string) {

	// migration of former implementation of enum
	switch treenames {
	// insertion code per enum code
	case Sidebar:
		res = "sidebar"
	}
	return
}

func (treenames *TreeNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "sidebar":
		*treenames = Sidebar
	default:
		return errUnkownEnum
	}
	return
}

func (treenames *TreeNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Sidebar":
		*treenames = Sidebar
	default:
		return errUnkownEnum
	}
	return
}

func (treenames *TreeNames) ToCodeString() (res string) {

	switch *treenames {
	// insertion code per enum code
	case Sidebar:
		res = "Sidebar"
	}
	return
}

func (treenames TreeNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Sidebar")

	return
}

func (treenames TreeNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "sidebar")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | DirectionType | FormNames | OrientationType | ProbabilityEnum | StacksNames | TreeNames
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*DirectionType | *FormNames | *OrientationType | *ProbabilityEnum | *StacksNames | *TreeNames
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
