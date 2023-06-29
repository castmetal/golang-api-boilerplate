package use_cases

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/castmetal/golang-api-boilerplate/src/domains/common"
	"github.com/castmetal/golang-api-boilerplate/src/domains/core/application/dtos"
	"github.com/castmetal/golang-api-boilerplate/src/domains/example"
	"github.com/castmetal/golang-api-boilerplate/src/infra/redis"
)

type (
	CreateExample interface {
		Execute(ctx context.Context, createExampleDTO *dtos.CreateExampleDTO) (dtos.CreateExampleResponseDTO, error)
		toResponse(example *example.Example) dtos.CreateExampleResponseDTO
	}
	CreateExampleRequest struct {
		CreateExample
		Repository  example.IExampleRepository
		RedisClient redis.IRedisClient
	}
)

func NewCreateExample(repository example.IExampleRepository, redisClient redis.IRedisClient) (CreateExample, error) {
	var uc CreateExample = &CreateExampleRequest{
		Repository:  repository,
		RedisClient: redisClient,
	}

	return uc, nil
}

func (uc *CreateExampleRequest) Execute(ctx context.Context, createExampleDTO *dtos.CreateExampleDTO) (dtos.CreateExampleResponseDTO, error) {
	var response = dtos.CreateExampleResponseDTO{}

	_, err := createExampleDTO.Validate()
	if err != nil {
		return response, common.InvalidParamsError(err.Error())
	}

	dtoBytes, err := createExampleDTO.ToBytes()
	if err != nil {
		return response, common.DefaultDomainError(err.Error())
	}

	dtoReader := bytes.NewReader(dtoBytes)
	exampleProps := getExampleProps(dtoReader)

	exmpl, err := example.NewExampleEntity(exampleProps)
	if err != nil {
		return response, common.DefaultDomainError(err.Error())
	}

	fByName, err := uc.Repository.FindOneByName(ctx, exmpl.Name)
	if err == nil && fByName.ID.String() != "" {
		return response, common.AlreadyExistsError("The example name: " + exmpl.Name)
	}

	err = uc.Repository.Create(ctx, exmpl)
	if err != nil {
		return response, common.DefaultDomainError(err.Error())
	}

	keys := example.GetRedisKeys()
	_ = uc.RedisClient.DelAllData(ctx, keys["REDIS_LIST_ALL_EXAMPLES_KEY"])

	return uc.toResponse(exmpl), nil
}

func (uc *CreateExampleRequest) toResponse(example *example.Example) dtos.CreateExampleResponseDTO {
	var response dtos.CreateExampleResponseDTO
	var exampleResponse dtos.ExampleResponseDTO

	examplesBytes, err := json.Marshal(example)
	if err != nil {
		return response
	}

	err = json.Unmarshal(examplesBytes, &exampleResponse)
	if err != nil {
		return response
	}

	return dtos.CreateExampleResponseDTO{
		Example: exampleResponse,
	}
}

func getExampleProps(message io.Reader) example.ExampleProps {
	var exampleProps example.ExampleProps
	messageBuffer := &bytes.Buffer{}
	messageBuffer.ReadFrom(message)

	if err := json.Unmarshal(messageBuffer.Bytes(), &exampleProps); err != nil {
		log.Fatal(err)
	}

	return exampleProps
}
