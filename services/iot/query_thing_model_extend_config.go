package iot

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryThingModelExtendConfig invokes the iot.QueryThingModelExtendConfig API synchronously
func (client *Client) QueryThingModelExtendConfig(request *QueryThingModelExtendConfigRequest) (response *QueryThingModelExtendConfigResponse, err error) {
	response = CreateQueryThingModelExtendConfigResponse()
	err = client.DoAction(request, response)
	return
}

// QueryThingModelExtendConfigWithChan invokes the iot.QueryThingModelExtendConfig API asynchronously
func (client *Client) QueryThingModelExtendConfigWithChan(request *QueryThingModelExtendConfigRequest) (<-chan *QueryThingModelExtendConfigResponse, <-chan error) {
	responseChan := make(chan *QueryThingModelExtendConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryThingModelExtendConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryThingModelExtendConfigWithCallback invokes the iot.QueryThingModelExtendConfig API asynchronously
func (client *Client) QueryThingModelExtendConfigWithCallback(request *QueryThingModelExtendConfigRequest, callback func(response *QueryThingModelExtendConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryThingModelExtendConfigResponse
		var err error
		defer close(result)
		response, err = client.QueryThingModelExtendConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QueryThingModelExtendConfigRequest is the request struct for api QueryThingModelExtendConfig
type QueryThingModelExtendConfigRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	ResourceGroupId   string `position:"Query" name:"ResourceGroupId"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
	ModelVersion      string `position:"Query" name:"ModelVersion"`
	FunctionBlockId   string `position:"Query" name:"FunctionBlockId"`
}

// QueryThingModelExtendConfigResponse is the response struct for api QueryThingModelExtendConfig
type QueryThingModelExtendConfigResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryThingModelExtendConfigRequest creates a request to invoke QueryThingModelExtendConfig API
func CreateQueryThingModelExtendConfigRequest() (request *QueryThingModelExtendConfigRequest) {
	request = &QueryThingModelExtendConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryThingModelExtendConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryThingModelExtendConfigResponse creates a response to parse from QueryThingModelExtendConfig response
func CreateQueryThingModelExtendConfigResponse() (response *QueryThingModelExtendConfigResponse) {
	response = &QueryThingModelExtendConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}