package main

import (
	"context"
	"encoding/json"

	"testing"
	"time"

	pb "github.com/YAWAL/GetMeConf/api"

	"errors"

	"github.com/YAWAL/GetMeConf/entitie"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type mockMongoDBConfigRepo struct {
}

func (m *mockMongoDBConfigRepo) Find(configName string) (*entitie.Mongodb, error) {
	return &entitie.Mongodb{Domain: "testName", Mongodb: true, Host: "testHost", Port: "testPort"}, nil
}

func (m *mockMongoDBConfigRepo) FindAll() ([]entitie.Mongodb, error) {
	return []entitie.Mongodb{{Domain: "testName", Mongodb: true, Host: "testHost", Port: "testPort"}}, nil
}

func (m *mockMongoDBConfigRepo) Update(config *entitie.Mongodb) (string, error) {
	return "OK", nil
}

func (m *mockMongoDBConfigRepo) Save(config *entitie.Mongodb) (string, error) {
	return "OK", nil
}

func (m *mockMongoDBConfigRepo) Delete(configName string) (string, error) {
	return "OK", nil
}

type mockErrorMongoDBConfigRepo struct {
}

func (m *mockErrorMongoDBConfigRepo) Find(configName string) (*entitie.Mongodb, error) {
	return nil, errors.New("error from database querying")
}

func (m *mockErrorMongoDBConfigRepo) FindAll() ([]entitie.Mongodb, error) {
	return nil, errors.New("error from database querying")
}

func (m *mockErrorMongoDBConfigRepo) Update(config *entitie.Mongodb) (string, error) {
	return "", errors.New("error from database querying")
}

func (m *mockErrorMongoDBConfigRepo) Save(config *entitie.Mongodb) (string, error) {
	return "", errors.New("error from database querying")
}
func (m *mockErrorMongoDBConfigRepo) Delete(configName string) (string, error) {
	return "", errors.New("error from database querying")
}

type mockTsConfigRepo struct {
}

func (m *mockTsConfigRepo) Find(configName string) (*entitie.Tsconfig, error) {
	return &entitie.Tsconfig{Module: "testModule", Target: "testTarget", SourceMap: true, Excluding: 1}, nil
}

func (m *mockTsConfigRepo) FindAll() ([]entitie.Tsconfig, error) {
	return []entitie.Tsconfig{{Module: "testModule", Target: "testTarget", SourceMap: true, Excluding: 1}}, nil
}

func (m *mockTsConfigRepo) Update(config *entitie.Tsconfig) (string, error) {
	return "OK", nil
}

func (m *mockTsConfigRepo) Save(config *entitie.Tsconfig) (string, error) {
	return "OK", nil
}

func (m *mockTsConfigRepo) Delete(configName string) (string, error) {
	return "OK", nil
}

type mockErrorTsConfigRepo struct {
}

func (m *mockErrorTsConfigRepo) Find(configName string) (*entitie.Tsconfig, error) {
	return nil, errors.New("error from database querying")
}

func (m *mockErrorTsConfigRepo) FindAll() ([]entitie.Tsconfig, error) {
	return nil, errors.New("error from database querying")
}

func (m *mockErrorTsConfigRepo) Update(config *entitie.Tsconfig) (string, error) {
	return "", errors.New("error from database querying")
}

func (m *mockErrorTsConfigRepo) Save(config *entitie.Tsconfig) (string, error) {
	return "", errors.New("error from database querying")
}
func (m *mockErrorTsConfigRepo) Delete(configName string) (string, error) {
	return "", errors.New("error from database querying")
}

type mockTempConfigRepo struct {
}

func (m *mockTempConfigRepo) Find(configName string) (*entitie.Tempconfig, error) {
	return &entitie.Tempconfig{RestApiRoot: "testApiRoot", Host: "testHost", Port: "testPort", Remoting: "testRemoting", LegasyExplorer: true}, nil
}

func (m *mockTempConfigRepo) FindAll() ([]entitie.Tempconfig, error) {
	return []entitie.Tempconfig{{RestApiRoot: "testApiRoot", Host: "testHost", Port: "testPort", Remoting: "testRemoting", LegasyExplorer: true}}, nil
}

func (m *mockTempConfigRepo) Update(config *entitie.Tempconfig) (string, error) {
	return "OK", nil
}

func (m *mockTempConfigRepo) Save(config *entitie.Tempconfig) (string, error) {
	return "OK", nil
}

func (m *mockTempConfigRepo) Delete(configName string) (string, error) {
	return "OK", nil
}

type mockErrorTempConfigRepo struct {
}

func (m *mockErrorTempConfigRepo) Find(configName string) (*entitie.Tempconfig, error) {
	return nil, errors.New("error from database querying")
}

func (m *mockErrorTempConfigRepo) FindAll() ([]entitie.Tempconfig, error) {
	return nil, errors.New("error from database querying")
}

func (m *mockErrorTempConfigRepo) Update(config *entitie.Tempconfig) (string, error) {
	return "", errors.New("error from database querying")
}

func (m *mockErrorTempConfigRepo) Save(config *entitie.Tempconfig) (string, error) {
	return "", errors.New("error from database querying")
}
func (m *mockErrorTempConfigRepo) Delete(configName string) (string, error) {
	return "", errors.New("error from database querying")
}

func TestGetConfigByName(t *testing.T) {

	configCache := cache.New(5*time.Minute, 10*time.Minute)
	mock := &mockConfigServer{}
	mock.configCache = configCache
	mock.mongoDBConfigRepo = &mockMongoDBConfigRepo{}
	mock.tsConfigRepo = &mockTsConfigRepo{}
	mock.tempConfigRepo = &mockTempConfigRepo{}
	res := &pb.GetConfigResponce{}
	err := mock.GetConfigByName(context.Background(), &pb.GetConfigByNameRequest{ConfigType: "mongodb", ConfigName: "testNameMongo"}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	var expectedConfig []byte
	expectedConfig, err = json.Marshal(entitie.Mongodb{Domain: "testName", Mongodb: true, Host: "testHost", Port: "testPort"})
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, expectedConfig, res.Config)

	err = mock.GetConfigByName(context.Background(), &pb.GetConfigByNameRequest{ConfigType: "tsconfig", ConfigName: "testNameTs"}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	expectedConfig, err = json.Marshal(entitie.Tsconfig{Module: "testModule", Target: "testTarget", SourceMap: true, Excluding: 1})
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, expectedConfig, res.Config)

	err = mock.GetConfigByName(context.Background(), &pb.GetConfigByNameRequest{ConfigType: "tempconfig", ConfigName: "testNameTemp"}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	expectedConfig, err = json.Marshal(entitie.Tempconfig{RestApiRoot: "testApiRoot", Host: "testHost", Port: "testPort", Remoting: "testRemoting", LegasyExplorer: true})
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, expectedConfig, res.Config)

	mock.configCache.Flush()

	mock.mongoDBConfigRepo = &mockErrorMongoDBConfigRepo{}
	expectedError := errors.New("error from database querying")
	err = mock.GetConfigByName(context.Background(), &pb.GetConfigByNameRequest{ConfigType: "mongodb", ConfigName: "testNameMongo"}, res)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
	mock.tsConfigRepo = &mockErrorTsConfigRepo{}
	err = mock.GetConfigByName(context.Background(), &pb.GetConfigByNameRequest{ConfigType: "tsconfig", ConfigName: "testNameTs"}, res)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
	mock.tempConfigRepo = &mockErrorTempConfigRepo{}
	err = mock.GetConfigByName(context.Background(), &pb.GetConfigByNameRequest{ConfigType: "tempconfig", ConfigName: "testNameTemp"}, res)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
	err = mock.GetConfigByName(context.Background(), &pb.GetConfigByNameRequest{ConfigType: "unexpectedConfigType", ConfigName: "testNameTemp"}, res)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("unexpected type"), err)
	}
}

func TestGetConfigByName_FromCache(t *testing.T) {
	testName := "testName"
	testConf := entitie.Mongodb{Domain: testName, Mongodb: true, Host: "testHost", Port: "testPort"}
	configCache := cache.New(5*time.Minute, 10*time.Minute)
	mock := &mockConfigServer{}
	mock.configCache = configCache

	byteRes, err := json.Marshal(testConf)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	configResponse := &pb.GetConfigResponce{Config: byteRes}
	mock.configCache.Set(testName, configResponse, 5*time.Minute)
	res := &pb.GetConfigResponce{}
	err = mock.GetConfigByName(context.Background(), &pb.GetConfigByNameRequest{ConfigType: "mongodb", ConfigName: "testName"}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	var expectedConfig []byte
	expectedConfig, err = json.Marshal(entitie.Mongodb{Domain: "testName", Mongodb: true, Host: "testHost", Port: "testPort"})
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, expectedConfig, res.Config)

}

type mockConfigServer struct {
	configServer
	grpc.ServerStream
	Results []*pb.GetConfigResponce
}

func (mcs *mockConfigServer) Send(response *pb.GetConfigResponce) error {
	mcs.Results = append(mcs.Results, response)
	return nil
}

func (mcs *mockConfigServer) Close() error {
	return nil
}

func TestGetConfigsByType(t *testing.T) {

	mock := &mockConfigServer{}
	mock.mongoDBConfigRepo = &mockMongoDBConfigRepo{}

	err := mock.GetConfigsByType(context.Background(), &pb.GetConfigsByTypeRequest{ConfigType: "mongodb"}, mock)
	assert.Equal(t, 1, len(mock.Results), "expected to contain 1 item")
	mock.tsConfigRepo = &mockTsConfigRepo{}
	err = mock.GetConfigsByType(context.Background(), &pb.GetConfigsByTypeRequest{ConfigType: "tsconfig"}, mock)
	assert.Equal(t, 2, len(mock.Results), "expected to contain 1 item")
	mock.tempConfigRepo = &mockTempConfigRepo{}
	err = mock.GetConfigsByType(context.Background(), &pb.GetConfigsByTypeRequest{ConfigType: "tempconfig"}, mock)
	assert.Equal(t, 3, len(mock.Results), "expected to contain 1 item")
	if err != nil {
		t.Error("error during unit testing of GetConfigsByType function: ", err)
	}
	err = mock.GetConfigsByType(context.Background(), &pb.GetConfigsByTypeRequest{ConfigType: "unexpectedConfigType"}, mock)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("unexpected type"), err)
	}

	expectedError := errors.New("error from database querying")
	err = nil
	mock.mongoDBConfigRepo = &mockErrorMongoDBConfigRepo{}
	err = mock.GetConfigsByType(context.Background(), &pb.GetConfigsByTypeRequest{ConfigType: "mongodb"}, mock)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}

	err = nil
	mock.tsConfigRepo = &mockErrorTsConfigRepo{}
	err = mock.GetConfigsByType(context.Background(), &pb.GetConfigsByTypeRequest{ConfigType: "tsconfig"}, mock)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("error from database querying"), err)
	}
	err = nil
	mock.tempConfigRepo = &mockErrorTempConfigRepo{}
	err = mock.GetConfigsByType(context.Background(), &pb.GetConfigsByTypeRequest{ConfigType: "tempconfig"}, mock)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
	err = nil
	mock.tempConfigRepo = &mockErrorTempConfigRepo{}
	err = mock.GetConfigsByType(context.Background(), &pb.GetConfigsByTypeRequest{ConfigType: "unexpectedType"}, mock)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("unexpected type"), err)
	}

}

func TestCreateConfig(t *testing.T) {

	configCache := cache.New(5*time.Minute, 10*time.Minute)
	mock := &mockConfigServer{}
	mock.configCache = configCache
	mock.mongoDBConfigRepo = &mockMongoDBConfigRepo{}
	mock.tsConfigRepo = &mockTsConfigRepo{}
	mock.tempConfigRepo = &mockTempConfigRepo{}

	testConfMongo := entitie.Mongodb{Domain: "testName", Mongodb: true, Host: "testHost", Port: "testPort"}
	byteRes, err := json.Marshal(testConfMongo)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	res := &pb.Responce{}
	err = mock.CreateConfig(context.Background(), &pb.Config{ConfigType: "mongodb", Config: byteRes}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	expectedResponse := &pb.Responce{Status: "OK"}
	assert.Equal(t, expectedResponse, res)

	testConfTs := entitie.Tsconfig{Module: "testModule", Target: "testTarget", SourceMap: true, Excluding: 1}
	byteRes, err = json.Marshal(testConfTs)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}

	err = mock.CreateConfig(context.Background(), &pb.Config{ConfigType: "tsconfig", Config: byteRes}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, expectedResponse, res)

	testConfTemp := entitie.Tempconfig{RestApiRoot: "testApiRoot", Host: "testHost", Port: "testPort", Remoting: "testRemoting", LegasyExplorer: true}
	byteRes, err = json.Marshal(testConfTemp)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}

	err = mock.CreateConfig(context.Background(), &pb.Config{ConfigType: "tempconfig", Config: byteRes}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, expectedResponse, res)

	mock.mongoDBConfigRepo = &mockErrorMongoDBConfigRepo{}
	mock.tsConfigRepo = &mockErrorTsConfigRepo{}
	mock.tempConfigRepo = &mockErrorTempConfigRepo{}
	expectedError := errors.New("error from database querying")

	res = &pb.Responce{}
	resultingErr := mock.CreateConfig(context.Background(), &pb.Config{ConfigType: "mongodb", Config: byteRes}, res)
	if assert.Error(t, resultingErr) {
		assert.Equal(t, expectedError, resultingErr)
	}
	resultingErr = nil
	resultingErr = mock.CreateConfig(context.Background(), &pb.Config{ConfigType: "tsconfig", Config: byteRes}, res)
	if assert.Error(t, resultingErr) {
		assert.Equal(t, expectedError, resultingErr)
	}
	resultingErr = nil
	resultingErr = mock.CreateConfig(context.Background(), &pb.Config{ConfigType: "tempconfig", Config: byteRes}, res)
	if assert.Error(t, resultingErr) {
		assert.Equal(t, expectedError, resultingErr)
	}
	resultingErr = nil
	resultingErr = mock.CreateConfig(context.Background(), &pb.Config{ConfigType: "unexpectedType", Config: byteRes}, res)
	if assert.Error(t, resultingErr) {
		assert.Equal(t, errors.New("unexpected type"), resultingErr)
	}
}

func TestDeleteConfig(t *testing.T) {

	configCache := cache.New(5*time.Minute, 10*time.Minute)
	mock := &mockConfigServer{}
	mock.configCache = configCache
	mock.mongoDBConfigRepo = &mockMongoDBConfigRepo{}
	mock.tsConfigRepo = &mockTsConfigRepo{}
	mock.tempConfigRepo = &mockTempConfigRepo{}
	res := &pb.Responce{}
	err := mock.DeleteConfig(context.Background(), &pb.DeleteConfigRequest{ConfigType: "mongodb", ConfigName: "testName"}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	expectedResponse := &pb.Responce{Status: "OK"}
	assert.Equal(t, expectedResponse, res)

	err = mock.DeleteConfig(context.Background(), &pb.DeleteConfigRequest{ConfigType: "tsconfig", ConfigName: "testName"}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}

	assert.Equal(t, expectedResponse, res)

	err = mock.DeleteConfig(context.Background(), &pb.DeleteConfigRequest{ConfigType: "tempconfig", ConfigName: "testName"}, res)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}

	assert.Equal(t, expectedResponse, res)

	mock.mongoDBConfigRepo = &mockErrorMongoDBConfigRepo{}
	mock.tsConfigRepo = &mockErrorTsConfigRepo{}
	mock.tempConfigRepo = &mockErrorTempConfigRepo{}
	expectedError := errors.New("error from database querying")
	resultingErr := mock.DeleteConfig(context.Background(), &pb.DeleteConfigRequest{ConfigType: "mongodb", ConfigName: "errorTestName"}, res)
	if assert.Error(t, resultingErr) {
		assert.Equal(t, expectedError, resultingErr)
	}
	resultingErr = nil
	resultingErr = mock.DeleteConfig(context.Background(), &pb.DeleteConfigRequest{ConfigType: "tsconfig", ConfigName: "errorTestName"}, res)
	if assert.Error(t, resultingErr) {
		assert.Equal(t, expectedError, resultingErr)
	}
	resultingErr = nil
	resultingErr = mock.DeleteConfig(context.Background(), &pb.DeleteConfigRequest{ConfigType: "tempconfig", ConfigName: "errorTestName"}, res)
	if assert.Error(t, resultingErr) {
		assert.Equal(t, expectedError, resultingErr)
	}
	resultingErr = nil
	resultingErr = mock.DeleteConfig(context.Background(), &pb.DeleteConfigRequest{ConfigType: "unexpectedType", ConfigName: "errorTestName"}, res)
	if assert.Error(t, resultingErr) {
		assert.Equal(t, errors.New("unexpected type"), resultingErr)
	}
}

func TestUpdateConfig(t *testing.T) {

	configCache := cache.New(5*time.Minute, 10*time.Minute)
	mock := &mockConfigServer{}
	mock.configCache = configCache
	mock.mongoDBConfigRepo = &mockMongoDBConfigRepo{}
	mock.tsConfigRepo = &mockTsConfigRepo{}
	mock.tempConfigRepo = &mockTempConfigRepo{}

	testConfMongo := entitie.Mongodb{Domain: "testName", Mongodb: true, Host: "testHost", Port: "testPort"}
	byteResMongo, err := json.Marshal(testConfMongo)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	testConfTs := entitie.Tsconfig{Module: "testModule", Target: "testTarget", SourceMap: true, Excluding: 1}
	byteResTs, err := json.Marshal(testConfTs)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	testConfTemp := entitie.Tempconfig{RestApiRoot: "testApiRoot", Host: "testHost", Port: "testPort", Remoting: "testRemoting", LegasyExplorer: true}
	byteResTemp, err := json.Marshal(testConfTemp)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	res := &pb.Responce{}
	err = mock.UpdateConfig(context.Background(), &pb.Config{ConfigType: "mongodb", Config: byteResMongo}, res)
	assert.Equal(t, &pb.Responce{Status: "OK"}, res)
	err = mock.UpdateConfig(context.Background(), &pb.Config{ConfigType: "tsconfig", Config: byteResTs}, res)
	assert.Equal(t, &pb.Responce{Status: "OK"}, res)
	err = mock.UpdateConfig(context.Background(), &pb.Config{ConfigType: "tempconfig", Config: byteResTemp}, res)
	assert.Equal(t, &pb.Responce{Status: "OK"}, res)
	err = mock.UpdateConfig(context.Background(), &pb.Config{ConfigType: "unexpectedConfigType"}, res)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("unexpected type"), err)
	}

	expectedError := errors.New("error from database querying")
	mock.mongoDBConfigRepo = &mockErrorMongoDBConfigRepo{}
	err = nil
	err = mock.UpdateConfig(context.Background(), &pb.Config{ConfigType: "mongodb", Config: byteResMongo}, res)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}

	err = nil
	mock.tsConfigRepo = &mockErrorTsConfigRepo{}
	err = mock.UpdateConfig(context.Background(), &pb.Config{ConfigType: "tsconfig", Config: byteResTs}, res)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}

	err = nil
	mock.tempConfigRepo = &mockErrorTempConfigRepo{}
	err = mock.UpdateConfig(context.Background(), &pb.Config{ConfigType: "tempconfig", Config: byteResTemp}, res)
	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
}
