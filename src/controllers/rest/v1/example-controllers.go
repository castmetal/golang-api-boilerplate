package rest_controllers_v1

import (
	"net/http"
	"strconv"

	"github.com/castmetal/golang-api-boilerplate/src/config"
	"github.com/castmetal/golang-api-boilerplate/src/domains/common"
	"github.com/castmetal/golang-api-boilerplate/src/domains/core/application/dtos"
	use_cases "github.com/castmetal/golang-api-boilerplate/src/domains/core/use-cases"
	"github.com/castmetal/golang-api-boilerplate/src/domains/example"
	"github.com/castmetal/golang-api-boilerplate/src/infra/redis"
	"github.com/castmetal/golang-api-boilerplate/src/infra/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

// CreateExample  godoc
//
// @Summary  Create an example based on the name input
// @Description Creating an example
// @Tags   Create Example
// @Accept   json
// @Produce  json
// @Param   createExample body  dtos.CreateExampleDTO true "CreateExample Data"
// @Success  200   {object} dtos.CreateExampleResponseDTO
// @Router   /v1/example [post]
func CreateExampleControllerV1(c *gin.Context, redisConn redis.IRedisClient, exampleRepository example.IExampleRepository) {
	var createExampleDTO dtos.CreateExampleDTO
	if err := c.ShouldBindJSON(&createExampleDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Request.Body != nil {
		defer c.Request.Body.Close()
	}

	ucCreateExample, err := use_cases.NewCreateExample(exampleRepository, redisConn)
	if err != nil {
		errMessage := common.InvalidConnectionError(err.Error())
		common.HandleHttpErrors(errMessage, c)
		return
	}

	response, err := ucCreateExample.Execute(c, &createExampleDTO)
	if err != nil {
		common.HandleHttpErrors(err, c)
		return
	}

	c.IndentedJSON(http.StatusCreated, response)
}

// ListAllExamples  godoc
//
// @Summary  List all examples in database
// @Description Listing all examples that was stored in the database
// @Tags   ListAll Example
// @Accept   json
// @Produce  json
// @Param    limit query string false "Limit" Format(numeric)
// @Param    offset query string false "Offset" Format(numeric)
// @Success  200   {object} dtos.ListAllExampleResponseDTO
// @Router   /v1/example [get]
func ListAllExamplesV1(c *gin.Context, redisConn redis.IRedisClient, exampleRepository example.IExampleRepository) {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limit = 10
	}

	listAllExampleDTO := &dtos.ListAllExampleDTO{
		Limit:  limit,
		Offset: offset,
	}

	if c.Request.Body != nil {
		defer c.Request.Body.Close()
	}

	ucListAllExample, err := use_cases.NewListAllExample(exampleRepository, redisConn)
	if err != nil {
		errMessage := common.InvalidConnectionError(err.Error())
		common.HandleHttpErrors(errMessage, c)
		return
	}

	response, err := ucListAllExample.Execute(c, listAllExampleDTO)
	if err != nil {
		common.HandleHttpErrors(err, c)
		return
	}

	c.IndentedJSON(http.StatusOK, response)

}

func SetExampleControllers(routerEngine *gin.Engine, config config.EnvStruct, pgConn *pgxpool.Pool, redisConn redis.IRedisClient) {
	exampleRepository := repositories.NewExampleRepository(pgConn)

	v1 := routerEngine.Group("/v1")
	{
		v1.POST("/example", func(c *gin.Context) {
			CreateExampleControllerV1(c, redisConn, exampleRepository)
		})

		v1.GET("/example", func(c *gin.Context) {
			ListAllExamplesV1(c, redisConn, exampleRepository)
		})
	}

}
