package use_cases

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/castmetal/golang-api-boilerplate/src/domains/common"
	"github.com/castmetal/golang-api-boilerplate/src/domains/core/application/dtos"
	"github.com/castmetal/golang-api-boilerplate/src/domains/example"
	"github.com/castmetal/golang-api-boilerplate/src/infra/redis"
)

const CACHE_DURATION = 24
const CACHE_TYPE = "HOUR"

type (
	ListAllExample interface {
		Execute(ctx context.Context, listAllExampleDTO *dtos.ListAllExampleDTO) (dtos.ListAllExampleResponseDTO, error)
		toResponse(example *example.Example) dtos.ListAllExampleResponseDTO
	}
	ListAllExampleRequest struct {
		ListAllExample
		Repository  example.IExampleRepository
		RedisClient redis.IRedisClient
	}
)

func NewListAllExample(repository example.IExampleRepository, redisClient redis.IRedisClient) (*ListAllExampleRequest, error) {
	var uc *ListAllExampleRequest = &ListAllExampleRequest{
		Repository:  repository,
		RedisClient: redisClient,
	}

	return uc, nil
}

func (uc *ListAllExampleRequest) Execute(ctx context.Context, listAllExampleDTO *dtos.ListAllExampleDTO) (dtos.ListAllExampleResponseDTO, error) {
	var response = dtos.ListAllExampleResponseDTO{}

	_, err := listAllExampleDTO.Validate()
	if err != nil {
		return response, common.InvalidParamsError(err.Error())
	}

	keys := example.GetRedisKeys()
	hashKey, _ := listAllExampleDTO.ToBytes()
	hash := md5.Sum(hashKey)
	hashStr := fmt.Sprintf("%x", string(hash[:]))
	key := keys["REDIS_LIST_ALL_EXAMPLES_KEY"] + "/" + hashStr

	redisResult, _ := uc.RedisClient.GetData(ctx, key)
	if redisResult != "" {
		json.Unmarshal([]byte(redisResult), &response)

		return response, nil
	}

	allData, err := uc.Repository.ListAll(ctx, listAllExampleDTO.Limit, listAllExampleDTO.Offset)
	if err != nil {
		return response, common.DefaultDomainError(err.Error())
	}

	response, err = uc.toResponse(allData)
	if err != nil {
		return response, common.DefaultDomainError(err.Error())
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		return response, common.DefaultDomainError(err.Error())
	}

	var cacheDuration time.Duration
	if CACHE_TYPE == "HOUR" {
		cacheDuration = time.Duration(CACHE_DURATION * time.Hour)
	} else {
		cacheDuration = time.Duration(CACHE_DURATION * time.Minute)
	}

	_ = uc.RedisClient.SetData(ctx, key, string(responseBytes), cacheDuration)

	return response, nil
}

func (uc *ListAllExampleRequest) toResponse(examples []*example.Example) (dtos.ListAllExampleResponseDTO, error) {
	examplesBytes, err := json.Marshal(examples)
	if err != nil {
		return dtos.ListAllExampleResponseDTO{}, err
	}
	var exampleResponse []dtos.ExampleResponseDTO

	err = json.Unmarshal(examplesBytes, &exampleResponse)
	if err != nil {
		return dtos.ListAllExampleResponseDTO{}, err
	}

	return dtos.ListAllExampleResponseDTO{
		Examples: exampleResponse,
	}, nil
}
