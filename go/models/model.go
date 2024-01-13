package models

import "time"

type AbstractNode struct {
	IsNodeExpanded bool
}

func (abstractnode *AbstractNode) GetIsNodeExpanded() bool {
	return abstractnode.IsNodeExpanded
}

func (abstractnode *AbstractNode) SetIsNodeExpanded(isNodeExpanded bool) {
	abstractnode.IsNodeExpanded = isNodeExpanded
}

type KingArthur struct {

	//gong:text
	Name string

	// Probability Management
	IsWithProbaility bool
	Probability      ProbabilityEnum
}

// KingArthurShape is abstract syntax (this is part of a scenario/diagram)
//
// It might be considered oncrete syntax (here is the way it is represented in the diagram)
// but it is not. The position is part of the definition of the actor state in a scenario
//
//	(not like UML where the position of a shape has no effect on the code)
type KingArthurShape struct {
	Name       string
	ActorState *KingArthur

	RectShape
}

func (actorStateShape *KingArthurShape) GetImplName() string {
	return actorStateShape.ActorState.Name
}

func (actorStateShape *KingArthurShape) GetModel() *KingArthur {
	return actorStateShape.ActorState
}

// Setters
func (s *KingArthurShape) SetName(name string) {
	s.Name = name
}

// this getter function is used for casting the ActorStateShape into a Gongstruct
func (s *KingArthurShape) GetShape() (r *KingArthurShape) {
	r = s
	return
}

type TheNuteTransition struct {
	Name string

	StartState *KingArthur
	EndState   *KingArthur
}

type TheNuteShape struct {
	Name                 string
	ActorStateTransition *TheNuteTransition
	Start                *KingArthurShape
	End                  *KingArthurShape

	// if link type is floating orthogonal ratio, from 0 to 1,
	// where the anchor starts on the edge (horizontal / vertical)
	StartOrientation OrientationType
	StartRatio       float64
	EndOrientation   OrientationType
	EndRatio         float64

	// in case StartOrientation is the same as EndOrientation,
	// there is a perpendicular line that reach the corner at
	// CornerOffsetRatio
	CornerOffsetRatio float64
}

// TheBridge is a first rank element
type TheBridge struct {
	Name string

	//gong:text gong:width 500 gong:height 500
	Description string

	WhatIsYourPreferedColor []*WhatIsYourPreferedColor

	GroupUse []*GroupUse

	// GeoObjectUse provides link to geo objects
	GeoObjectUse []*GeoObjectUse

	// MapUse provides link to reference map for this analysis
	MapUse []*MapObjectUse

	AbstractNode
}

type SirRobin struct {
	Name string

	//gong:text, magic code to have the form editor have a text area instead of an input
	Description string

	Witches           []*AWitch
	Arthurs           []*KingArthurShape
	KnightWhoSayNis   []*KnightWhoSayNi
	BlackKnightShapes []*BlackKnightShape
	TheNuteShapes     []*TheNuteShape

	// origin of axis
	AxisOrign_X float64 // left of the arrow
	AxisOrign_Y float64

	// position of the vertical axis
	// in SVG, vertical coordinate go from 0 (top of the screen) to +infinite (bottom of the screen)
	VerticalAxis_Top_Y       float64
	VerticalAxis_Bottom_Y    float64
	VerticalAxis_StrokeWidth float64

	// HoryzontalAxis_Right_X is the X coordinate of the arrow at the end of the horizontal axis
	HorizontalAxis_Right_X float64

	// abscisse dates
	Start time.Time
	End   time.Time

	// IsInDrawMode indicates the the drawing can be edited (in development mode)
	// or not (in production mode)
	// IsInDrawMode is a transciant value. It can only be persisted as a false value
	// therefore, it is not persisted at all
	// gong:ignore
	IsInDrawMode bool
}

func (d *SirRobin) GetYearsWithJanuaryFirstBetween() (years []time.Time) {
	for year := d.Start.Year(); year <= d.End.Year(); year++ {
		janFirst := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		if janFirst.After(d.Start) && janFirst.Before(d.End) {
			years = append(years, janFirst)
		}
	}
	return
}

// Default generated models package docs
// (at least one file is necessary in a models package)

type Document struct {
	Name string

	// GeoObjectUse provides link to geo objects
	GeoObjectUse []*GeoObjectUse
}

type DocumentUse struct {
	Name     string
	Document *Document
}

type GalahadThePure struct {
	Name        string
	Description string
}

type AWitch struct {
	Name               string
	EvolutionDirection *GalahadThePure

	RectShape
}

func (aWitch *AWitch) GetImplName() string {
	return aWitch.EvolutionDirection.Name
}

func (aWitch *AWitch) GetModel() *GalahadThePure {
	return aWitch.EvolutionDirection
}

// Setters
func (eds *AWitch) SetName(name string) {
	eds.Name = name
}

// this getter function is used for casting the GenericRectShape into a Gongstruct
func (s *AWitch) GetShape() (r *AWitch) {
	r = s
	return
}

// FormNames
type FormNames string

const (
	DefaultForm FormNames = "default form"

	// Name of the form for editing models objects
	ModelForm FormNames = "Model Form"

	// Name of the form for editing shape objects
	ShapeForm FormNames = "Shape Form"
)

// GenericNode is a Gongstruct that is also used as a node
// in the tree
type GenericNode[T Gongstruct] interface {
	*T
	GetIsNodeExpanded() bool
	SetIsNodeExpanded(isNodeExpanded bool)
}

// GenericRectShape
type GenericRectShape[ShapeType, ModelType Gongstruct] interface {
	*AWitch | *KingArthurShape | *KnightWhoSayNi | *BlackKnightShape

	// https://github.com/golang/go/issues/63940
	// waiting to avoid the below code

	GetImplName() string
	GetModel() *ModelType
	GetShape() *ShapeType

	GetName() string
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
	GetFillColor() string
	GetFillOpacity() float64
	GetStroke() string
	GetStrokeWidth() float64
	GetRX() float64

	SetName(name string)
	SetX(x float64)
	SetY(y float64)
	SetWidth(width float64)
	SetHeight(Height float64)
	SetFillColor(fillColor string)
	SetFillOpacity(fillOpacity float64)
	SetStroke(stroke string)
	SetStrokeWidth(strokeWidth float64)
	SetRX(rx float64)
}

type GeoObject struct {
	Name string
}

type GeoObjectUse struct {
	Name      string
	GeoObject *GeoObject
}

type Group struct {
	Name    string
	UserUse []*UserUse
}

type GroupUse struct {
	Name  string
	Group *Group
}

type MapObject struct {
	Name string
}

type MapObjectUse struct {
	Name string
	Map  *MapObject
}

type OrientationType string

// values for EnumType
const (
	ORIENTATION_HORIZONTAL OrientationType = "ORIENTATION_HORIZONTAL"
	ORIENTATION_VERTICAL   OrientationType = "ORIENTATION_VERTICAL"
)

type Lancelot struct {
	Name string

	// Tag is used for labelling a given version of a parameter
	Tag string

	// Description provides a longer account than the name of the parameter
	// This should allow the paramter to be understood
	// It also provides a reference to any input documents that
	// generated the parameter
	//gong:text gong:width 500 gong:height 500
	Description string

	// GroupUse provides the need to know capability
	GroupUse []*GroupUse

	// DocumentUse enables link to a document
	DocumentUse []*DocumentUse

	// GeoObjectUse provides link to geo objects
	GeoObjectUse []*GeoObjectUse
}

type LancelotAgregation struct {
	Name         string
	ParameterUse []*KnightWhoSayNi
}

type LancelotAgregationUse struct {
	Name                string
	ParameterAgregation *LancelotAgregation
}

type LancelotCategory struct {
	Name         string
	ParameterUse []*KnightWhoSayNi
}

type LancelotCategoryUse struct {
	Name              string
	ParameterCategory *LancelotCategory
}

type KnightWhoSayNi struct {
	Name      string
	Parameter *Lancelot

	// The shape might have arrowing showing the
	// influence on state actors
	Direction DirectionType

	RectShape
}

func (parameterShape *KnightWhoSayNi) GetImplName() string {
	return parameterShape.Parameter.Name
}

func (parameterShape *KnightWhoSayNi) GetModel() *Lancelot {
	return parameterShape.Parameter
}

// Setters
func (s *KnightWhoSayNi) SetName(name string) {
	s.Name = name
}

// this getter function is used for casting the GenericRectShape into a Gongstruct
func (s *KnightWhoSayNi) GetShape() (r *KnightWhoSayNi) {
	r = s
	return
}

func (s *KnightWhoSayNi) GetDirection() DirectionType {
	return s.Direction
}

func (s *KnightWhoSayNi) SetDirection(direction DirectionType) {
	s.Direction = direction
}

type DirectionType string

// values for EnumType
const (
	DIRECTION_UP   DirectionType = "DIRECTION_UP"
	DIRECTION_DOWN DirectionType = "DIRECTION_DOWN"
)

// ProbabilityEnum
type ProbabilityEnum string

const (
	PERCENT_100 ProbabilityEnum = "100 %"
	PERCENT_75  ProbabilityEnum = "75 %"
	PERCENT_50  ProbabilityEnum = "50 %"
	PERCENT_25  ProbabilityEnum = "25 %"
	PERCENT_00  ProbabilityEnum = "0 %"
)

type RectShape struct {
	X      float64
	Y      float64
	Width  float64
	Height float64

	FillColor   string
	FillOpacity float64
	Stroke      string
	StrokeWidth float64
	RX          float64
}

// GetX returns the X field.
func (shape *RectShape) GetX() float64 {
	return shape.X
}

// GetY returns the Y field.
func (shape *RectShape) GetY() float64 {
	return shape.Y
}

// GetWidth returns the Width field.
func (shape *RectShape) GetWidth() float64 {
	return shape.Width
}

// GetHeight returns the Height field.
func (shape *RectShape) GetHeight() float64 {
	return shape.Height
}

// GetFillColor returns the FillColor field.
func (shape *RectShape) GetFillColor() string {
	return shape.FillColor
}

// GetFillOpacity returns the FillOpacity field.
func (shape *RectShape) GetFillOpacity() float64 {
	return shape.FillOpacity
}

// GetStroke returns the Stroke field.
func (shape *RectShape) GetStroke() string {
	return shape.Stroke
}

// GetStrokeWidth returns the StrokeWidth field.
func (shape *RectShape) GetStrokeWidth() float64 {
	return shape.StrokeWidth
}

// GetRX returns the RX field.
func (shape *RectShape) GetRX() float64 {
	return shape.RX
}

func (shape *RectShape) SetX(x float64) {
	shape.X = x
}

func (shape *RectShape) SetY(y float64) {
	shape.Y = y
}

func (shape *RectShape) SetWidth(width float64) {
	shape.Width = width
}

func (shape *RectShape) SetHeight(Height float64) {
	shape.Height = Height
}

func (shape *RectShape) SetFillColor(fillColor string) {
	shape.FillColor = fillColor
}

func (shape *RectShape) SetFillOpacity(fillOpacity float64) {
	shape.FillOpacity = fillOpacity
}

func (shape *RectShape) SetStroke(stroke string) {
	shape.Stroke = stroke
}

func (shape *RectShape) SetStrokeWidth(strokeWidth float64) {
	shape.StrokeWidth = strokeWidth
}

func (shape *RectShape) SetRX(rx float64) {
	shape.RX = rx
}

// added a comment

// Repository is a first rank element
type Repository struct {
	Name         string
	ParameterUse []*KnightWhoSayNi
	GroupUse     []*GroupUse
}

type WhatIsYourPreferedColor struct {
	//gong:text, magic code to have the form editor have a text area instead of an input
	Name string

	//gong:text, magic code to have the form editor have a text area instead of an input
	Description string

	Diagrams           []*SirRobin
	DiagramsNodeFolded bool

	KingArthurs          []*KingArthur
	KingArthurNodeFolded bool

	Nutes []*TheNuteTransition
	RRRR  bool

	Galahard []*GalahadThePure
	IIUU     bool

	Lancelots          []*Lancelot
	LancelotsodeFolded bool

	BringYourDeadarameters []*BringYourDead
	RRRRT                  bool

	AbstractNode
}

type BringYourDead struct {
	Name string

	// Tag is used for labelling a given version of a parameter
	Tag string

	// Description provides a longer account than the name of the parameter
	// This should allow the paramter to be understood
	// It also provides a reference to any input documents that
	// generated the parameter
	//gong:text gong:width 500 gong:height 500
	Description string

	// Lancelots is the list of parameters this scenario
	// paramters aggregation
	Lancelots []*Lancelot
}

type BlackKnightShape struct {
	Name          string
	BringYourDead *BringYourDead

	// The shape might have arrowing showing the
	// influence on state actors
	Direction DirectionType

	RectShape
}

func (i *BlackKnightShape) GetImplName() string {
	return i.BringYourDead.Name
}

func (i *BlackKnightShape) GetModel() *BringYourDead {
	return i.BringYourDead
}

// Setters
func (s *BlackKnightShape) SetName(name string) {
	s.Name = name
}

// this getter function is used for casting the GenericRectShape into a Gongstruct
func (s *BlackKnightShape) GetShape() (r *BlackKnightShape) {
	r = s
	return
}

func (s *BlackKnightShape) GetDirection() DirectionType {
	return s.Direction
}

func (s *BlackKnightShape) SetDirection(direction DirectionType) {
	s.Direction = direction
}

type ShapeWithDirection interface {
	GetDirection() DirectionType
	SetDirection(direction DirectionType)
}

// StacksNames - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum StacksNames
type StacksNames string

// values for TableExtraNameEnum
const (
	// we have three application stacks. All have the same name
	WhatStackName StacksNames = "what"

	GongsvgStackName   StacksNames = "gongsvg"
	GongtreeStackName  StacksNames = "gongtree"
	GongtableStackName StacksNames = "gongtable"
)

// TreeNames
type TreeNames string

const (
	Sidebar TreeNames = "sidebar"
)

type User struct {
	Name string
}

type UserUse struct {
	Name string
	User *User
}

// Workspace is the singloton
type Workspace struct {
	Name            string
	SelectedDiagram *SirRobin

	// Default values for creation of shapes
	A *AWitch
	B *KnightWhoSayNi
	C *BlackKnightShape
	D *KingArthurShape
	E *TheNuteShape
}

func GetWorkspace(stage *StageStruct) (workspace *Workspace) {

	workspaces := *GetGongstructInstancesSet[Workspace](stage)

	for _workspace := range workspaces {
		workspace = _workspace
	}

	return
}

// GetPreferedColor parse all diagram and if one diagram
// matches the
func (workspace *Workspace) GetPreferedColor(stage *StageStruct) (scenario *WhatIsYourPreferedColor) {

	fieldName := GetAssociationName[WhatIsYourPreferedColor]().Diagrams[0].Name
	m := GetSliceOfPointersReverseMap[WhatIsYourPreferedColor, SirRobin](fieldName, stage)

	scenario = m[workspace.SelectedDiagram]

	return
}
