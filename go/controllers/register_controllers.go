// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/thomaspeugeot/thelongbuild/go")
	{ // insertion point for registrations
		v1.GET("/v1/awitchs", GetController().GetAWitchs)
		v1.GET("/v1/awitchs/:id", GetController().GetAWitch)
		v1.POST("/v1/awitchs", GetController().PostAWitch)
		v1.PATCH("/v1/awitchs/:id", GetController().UpdateAWitch)
		v1.PUT("/v1/awitchs/:id", GetController().UpdateAWitch)
		v1.DELETE("/v1/awitchs/:id", GetController().DeleteAWitch)

		v1.GET("/v1/blackknightshapes", GetController().GetBlackKnightShapes)
		v1.GET("/v1/blackknightshapes/:id", GetController().GetBlackKnightShape)
		v1.POST("/v1/blackknightshapes", GetController().PostBlackKnightShape)
		v1.PATCH("/v1/blackknightshapes/:id", GetController().UpdateBlackKnightShape)
		v1.PUT("/v1/blackknightshapes/:id", GetController().UpdateBlackKnightShape)
		v1.DELETE("/v1/blackknightshapes/:id", GetController().DeleteBlackKnightShape)

		v1.GET("/v1/bringyourdeads", GetController().GetBringYourDeads)
		v1.GET("/v1/bringyourdeads/:id", GetController().GetBringYourDead)
		v1.POST("/v1/bringyourdeads", GetController().PostBringYourDead)
		v1.PATCH("/v1/bringyourdeads/:id", GetController().UpdateBringYourDead)
		v1.PUT("/v1/bringyourdeads/:id", GetController().UpdateBringYourDead)
		v1.DELETE("/v1/bringyourdeads/:id", GetController().DeleteBringYourDead)

		v1.GET("/v1/documents", GetController().GetDocuments)
		v1.GET("/v1/documents/:id", GetController().GetDocument)
		v1.POST("/v1/documents", GetController().PostDocument)
		v1.PATCH("/v1/documents/:id", GetController().UpdateDocument)
		v1.PUT("/v1/documents/:id", GetController().UpdateDocument)
		v1.DELETE("/v1/documents/:id", GetController().DeleteDocument)

		v1.GET("/v1/documentuses", GetController().GetDocumentUses)
		v1.GET("/v1/documentuses/:id", GetController().GetDocumentUse)
		v1.POST("/v1/documentuses", GetController().PostDocumentUse)
		v1.PATCH("/v1/documentuses/:id", GetController().UpdateDocumentUse)
		v1.PUT("/v1/documentuses/:id", GetController().UpdateDocumentUse)
		v1.DELETE("/v1/documentuses/:id", GetController().DeleteDocumentUse)

		v1.GET("/v1/galahadthepures", GetController().GetGalahadThePures)
		v1.GET("/v1/galahadthepures/:id", GetController().GetGalahadThePure)
		v1.POST("/v1/galahadthepures", GetController().PostGalahadThePure)
		v1.PATCH("/v1/galahadthepures/:id", GetController().UpdateGalahadThePure)
		v1.PUT("/v1/galahadthepures/:id", GetController().UpdateGalahadThePure)
		v1.DELETE("/v1/galahadthepures/:id", GetController().DeleteGalahadThePure)

		v1.GET("/v1/geoobjects", GetController().GetGeoObjects)
		v1.GET("/v1/geoobjects/:id", GetController().GetGeoObject)
		v1.POST("/v1/geoobjects", GetController().PostGeoObject)
		v1.PATCH("/v1/geoobjects/:id", GetController().UpdateGeoObject)
		v1.PUT("/v1/geoobjects/:id", GetController().UpdateGeoObject)
		v1.DELETE("/v1/geoobjects/:id", GetController().DeleteGeoObject)

		v1.GET("/v1/geoobjectuses", GetController().GetGeoObjectUses)
		v1.GET("/v1/geoobjectuses/:id", GetController().GetGeoObjectUse)
		v1.POST("/v1/geoobjectuses", GetController().PostGeoObjectUse)
		v1.PATCH("/v1/geoobjectuses/:id", GetController().UpdateGeoObjectUse)
		v1.PUT("/v1/geoobjectuses/:id", GetController().UpdateGeoObjectUse)
		v1.DELETE("/v1/geoobjectuses/:id", GetController().DeleteGeoObjectUse)

		v1.GET("/v1/groups", GetController().GetGroups)
		v1.GET("/v1/groups/:id", GetController().GetGroup)
		v1.POST("/v1/groups", GetController().PostGroup)
		v1.PATCH("/v1/groups/:id", GetController().UpdateGroup)
		v1.PUT("/v1/groups/:id", GetController().UpdateGroup)
		v1.DELETE("/v1/groups/:id", GetController().DeleteGroup)

		v1.GET("/v1/groupuses", GetController().GetGroupUses)
		v1.GET("/v1/groupuses/:id", GetController().GetGroupUse)
		v1.POST("/v1/groupuses", GetController().PostGroupUse)
		v1.PATCH("/v1/groupuses/:id", GetController().UpdateGroupUse)
		v1.PUT("/v1/groupuses/:id", GetController().UpdateGroupUse)
		v1.DELETE("/v1/groupuses/:id", GetController().DeleteGroupUse)

		v1.GET("/v1/kingarthurs", GetController().GetKingArthurs)
		v1.GET("/v1/kingarthurs/:id", GetController().GetKingArthur)
		v1.POST("/v1/kingarthurs", GetController().PostKingArthur)
		v1.PATCH("/v1/kingarthurs/:id", GetController().UpdateKingArthur)
		v1.PUT("/v1/kingarthurs/:id", GetController().UpdateKingArthur)
		v1.DELETE("/v1/kingarthurs/:id", GetController().DeleteKingArthur)

		v1.GET("/v1/kingarthurshapes", GetController().GetKingArthurShapes)
		v1.GET("/v1/kingarthurshapes/:id", GetController().GetKingArthurShape)
		v1.POST("/v1/kingarthurshapes", GetController().PostKingArthurShape)
		v1.PATCH("/v1/kingarthurshapes/:id", GetController().UpdateKingArthurShape)
		v1.PUT("/v1/kingarthurshapes/:id", GetController().UpdateKingArthurShape)
		v1.DELETE("/v1/kingarthurshapes/:id", GetController().DeleteKingArthurShape)

		v1.GET("/v1/knightwhosaynis", GetController().GetKnightWhoSayNis)
		v1.GET("/v1/knightwhosaynis/:id", GetController().GetKnightWhoSayNi)
		v1.POST("/v1/knightwhosaynis", GetController().PostKnightWhoSayNi)
		v1.PATCH("/v1/knightwhosaynis/:id", GetController().UpdateKnightWhoSayNi)
		v1.PUT("/v1/knightwhosaynis/:id", GetController().UpdateKnightWhoSayNi)
		v1.DELETE("/v1/knightwhosaynis/:id", GetController().DeleteKnightWhoSayNi)

		v1.GET("/v1/lancelots", GetController().GetLancelots)
		v1.GET("/v1/lancelots/:id", GetController().GetLancelot)
		v1.POST("/v1/lancelots", GetController().PostLancelot)
		v1.PATCH("/v1/lancelots/:id", GetController().UpdateLancelot)
		v1.PUT("/v1/lancelots/:id", GetController().UpdateLancelot)
		v1.DELETE("/v1/lancelots/:id", GetController().DeleteLancelot)

		v1.GET("/v1/lancelotagregations", GetController().GetLancelotAgregations)
		v1.GET("/v1/lancelotagregations/:id", GetController().GetLancelotAgregation)
		v1.POST("/v1/lancelotagregations", GetController().PostLancelotAgregation)
		v1.PATCH("/v1/lancelotagregations/:id", GetController().UpdateLancelotAgregation)
		v1.PUT("/v1/lancelotagregations/:id", GetController().UpdateLancelotAgregation)
		v1.DELETE("/v1/lancelotagregations/:id", GetController().DeleteLancelotAgregation)

		v1.GET("/v1/lancelotagregationuses", GetController().GetLancelotAgregationUses)
		v1.GET("/v1/lancelotagregationuses/:id", GetController().GetLancelotAgregationUse)
		v1.POST("/v1/lancelotagregationuses", GetController().PostLancelotAgregationUse)
		v1.PATCH("/v1/lancelotagregationuses/:id", GetController().UpdateLancelotAgregationUse)
		v1.PUT("/v1/lancelotagregationuses/:id", GetController().UpdateLancelotAgregationUse)
		v1.DELETE("/v1/lancelotagregationuses/:id", GetController().DeleteLancelotAgregationUse)

		v1.GET("/v1/lancelotcategorys", GetController().GetLancelotCategorys)
		v1.GET("/v1/lancelotcategorys/:id", GetController().GetLancelotCategory)
		v1.POST("/v1/lancelotcategorys", GetController().PostLancelotCategory)
		v1.PATCH("/v1/lancelotcategorys/:id", GetController().UpdateLancelotCategory)
		v1.PUT("/v1/lancelotcategorys/:id", GetController().UpdateLancelotCategory)
		v1.DELETE("/v1/lancelotcategorys/:id", GetController().DeleteLancelotCategory)

		v1.GET("/v1/lancelotcategoryuses", GetController().GetLancelotCategoryUses)
		v1.GET("/v1/lancelotcategoryuses/:id", GetController().GetLancelotCategoryUse)
		v1.POST("/v1/lancelotcategoryuses", GetController().PostLancelotCategoryUse)
		v1.PATCH("/v1/lancelotcategoryuses/:id", GetController().UpdateLancelotCategoryUse)
		v1.PUT("/v1/lancelotcategoryuses/:id", GetController().UpdateLancelotCategoryUse)
		v1.DELETE("/v1/lancelotcategoryuses/:id", GetController().DeleteLancelotCategoryUse)

		v1.GET("/v1/mapobjects", GetController().GetMapObjects)
		v1.GET("/v1/mapobjects/:id", GetController().GetMapObject)
		v1.POST("/v1/mapobjects", GetController().PostMapObject)
		v1.PATCH("/v1/mapobjects/:id", GetController().UpdateMapObject)
		v1.PUT("/v1/mapobjects/:id", GetController().UpdateMapObject)
		v1.DELETE("/v1/mapobjects/:id", GetController().DeleteMapObject)

		v1.GET("/v1/mapobjectuses", GetController().GetMapObjectUses)
		v1.GET("/v1/mapobjectuses/:id", GetController().GetMapObjectUse)
		v1.POST("/v1/mapobjectuses", GetController().PostMapObjectUse)
		v1.PATCH("/v1/mapobjectuses/:id", GetController().UpdateMapObjectUse)
		v1.PUT("/v1/mapobjectuses/:id", GetController().UpdateMapObjectUse)
		v1.DELETE("/v1/mapobjectuses/:id", GetController().DeleteMapObjectUse)

		v1.GET("/v1/repositorys", GetController().GetRepositorys)
		v1.GET("/v1/repositorys/:id", GetController().GetRepository)
		v1.POST("/v1/repositorys", GetController().PostRepository)
		v1.PATCH("/v1/repositorys/:id", GetController().UpdateRepository)
		v1.PUT("/v1/repositorys/:id", GetController().UpdateRepository)
		v1.DELETE("/v1/repositorys/:id", GetController().DeleteRepository)

		v1.GET("/v1/sirrobins", GetController().GetSirRobins)
		v1.GET("/v1/sirrobins/:id", GetController().GetSirRobin)
		v1.POST("/v1/sirrobins", GetController().PostSirRobin)
		v1.PATCH("/v1/sirrobins/:id", GetController().UpdateSirRobin)
		v1.PUT("/v1/sirrobins/:id", GetController().UpdateSirRobin)
		v1.DELETE("/v1/sirrobins/:id", GetController().DeleteSirRobin)

		v1.GET("/v1/thebridges", GetController().GetTheBridges)
		v1.GET("/v1/thebridges/:id", GetController().GetTheBridge)
		v1.POST("/v1/thebridges", GetController().PostTheBridge)
		v1.PATCH("/v1/thebridges/:id", GetController().UpdateTheBridge)
		v1.PUT("/v1/thebridges/:id", GetController().UpdateTheBridge)
		v1.DELETE("/v1/thebridges/:id", GetController().DeleteTheBridge)

		v1.GET("/v1/thenuteshapes", GetController().GetTheNuteShapes)
		v1.GET("/v1/thenuteshapes/:id", GetController().GetTheNuteShape)
		v1.POST("/v1/thenuteshapes", GetController().PostTheNuteShape)
		v1.PATCH("/v1/thenuteshapes/:id", GetController().UpdateTheNuteShape)
		v1.PUT("/v1/thenuteshapes/:id", GetController().UpdateTheNuteShape)
		v1.DELETE("/v1/thenuteshapes/:id", GetController().DeleteTheNuteShape)

		v1.GET("/v1/thenutetransitions", GetController().GetTheNuteTransitions)
		v1.GET("/v1/thenutetransitions/:id", GetController().GetTheNuteTransition)
		v1.POST("/v1/thenutetransitions", GetController().PostTheNuteTransition)
		v1.PATCH("/v1/thenutetransitions/:id", GetController().UpdateTheNuteTransition)
		v1.PUT("/v1/thenutetransitions/:id", GetController().UpdateTheNuteTransition)
		v1.DELETE("/v1/thenutetransitions/:id", GetController().DeleteTheNuteTransition)

		v1.GET("/v1/users", GetController().GetUsers)
		v1.GET("/v1/users/:id", GetController().GetUser)
		v1.POST("/v1/users", GetController().PostUser)
		v1.PATCH("/v1/users/:id", GetController().UpdateUser)
		v1.PUT("/v1/users/:id", GetController().UpdateUser)
		v1.DELETE("/v1/users/:id", GetController().DeleteUser)

		v1.GET("/v1/useruses", GetController().GetUserUses)
		v1.GET("/v1/useruses/:id", GetController().GetUserUse)
		v1.POST("/v1/useruses", GetController().PostUserUse)
		v1.PATCH("/v1/useruses/:id", GetController().UpdateUserUse)
		v1.PUT("/v1/useruses/:id", GetController().UpdateUserUse)
		v1.DELETE("/v1/useruses/:id", GetController().DeleteUserUse)

		v1.GET("/v1/whatisyourpreferedcolors", GetController().GetWhatIsYourPreferedColors)
		v1.GET("/v1/whatisyourpreferedcolors/:id", GetController().GetWhatIsYourPreferedColor)
		v1.POST("/v1/whatisyourpreferedcolors", GetController().PostWhatIsYourPreferedColor)
		v1.PATCH("/v1/whatisyourpreferedcolors/:id", GetController().UpdateWhatIsYourPreferedColor)
		v1.PUT("/v1/whatisyourpreferedcolors/:id", GetController().UpdateWhatIsYourPreferedColor)
		v1.DELETE("/v1/whatisyourpreferedcolors/:id", GetController().DeleteWhatIsYourPreferedColor)

		v1.GET("/v1/workspaces", GetController().GetWorkspaces)
		v1.GET("/v1/workspaces/:id", GetController().GetWorkspace)
		v1.POST("/v1/workspaces", GetController().PostWorkspace)
		v1.PATCH("/v1/workspaces/:id", GetController().UpdateWorkspace)
		v1.PUT("/v1/workspaces/:id", GetController().UpdateWorkspace)
		v1.DELETE("/v1/workspaces/:id", GetController().DeleteWorkspace)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/thelongbuild/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
