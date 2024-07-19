package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/mrumyantsev/swagmod/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()

	c := NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	r.Run(":8080")
}

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  ResponseOk
// @Failure      201  {object}  ResponseError1
// @Failure      202  {object}  ResponseError2
// @Failure      203  {object}  ResponseError3
// @Router       /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

// ListAccounts godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {object}  ResponseOk
// @Failure      201  {object}  ResponseError1
// @Failure      202  {object}  ResponseError2
// @Failure      203  {object}  ResponseError3
// @Router       /accounts [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

type ResponseOk struct {
	Errno   int    `json:"errno" example:"0"`
	Message string `json:"message" example:""`
	Extra   string `json:"extra" example:""`
	Payload string `json:"payload" example:"{\"id\": 4298009, \"createDatetime\": \"2006-06-26T09:37Z08:00\"}"`
}

type ResponseError1 struct {
	Errno   int    `json:"errno" example:"11055"`
	Message string `json:"message" example:"error 1"`
	Extra   string `json:"extra" example:"extra 1"`
	Payload string `json:"payload" example:""`
}

type ResponseError2 struct {
	Errno   int    `json:"errno" example:"11068"`
	Message string `json:"message" example:"error 2"`
	Extra   string `json:"extra" example:"extra 2"`
	Payload string `json:"payload" example:""`
}

type ResponseError3 struct {
	Errno   int    `json:"errno" example:"11070"`
	Message string `json:"message" example:"error 3"`
	Extra   string `json:"extra" example:"extra 3"`
	Payload string `json:"payload" example:""`
}
