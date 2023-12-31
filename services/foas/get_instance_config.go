package foas

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

// GetInstanceConfig invokes the foas.GetInstanceConfig API synchronously
func (client *Client) GetInstanceConfig(request *GetInstanceConfigRequest) (response *GetInstanceConfigResponse, err error) {
	response = CreateGetInstanceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceConfigWithChan invokes the foas.GetInstanceConfig API asynchronously
func (client *Client) GetInstanceConfigWithChan(request *GetInstanceConfigRequest) (<-chan *GetInstanceConfigResponse, <-chan error) {
	responseChan := make(chan *GetInstanceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceConfig(request)
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

// GetInstanceConfigWithCallback invokes the foas.GetInstanceConfig API asynchronously
func (client *Client) GetInstanceConfigWithCallback(request *GetInstanceConfigRequest, callback func(response *GetInstanceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceConfigResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceConfig(request)
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

// GetInstanceConfigRequest is the request struct for api GetInstanceConfig
type GetInstanceConfigRequest struct {
	*requests.RoaRequest
	ProjectName string           `position:"Path" name:"projectName"`
	InstanceId  requests.Integer `position:"Path" name:"instanceId"`
	JobName     string           `position:"Path" name:"jobName"`
}

// GetInstanceConfigResponse is the response struct for api GetInstanceConfig
type GetInstanceConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Config    string `json:"Config" xml:"Config"`
}

// CreateGetInstanceConfigRequest creates a request to invoke GetInstanceConfig API
func CreateGetInstanceConfigRequest() (request *GetInstanceConfigRequest) {
	request = &GetInstanceConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetInstanceConfig", "/api/v2/projects/[projectName]/jobs/[jobName]/instances/[instanceId]/config", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetInstanceConfigResponse creates a response to parse from GetInstanceConfig response
func CreateGetInstanceConfigResponse() (response *GetInstanceConfigResponse) {
	response = &GetInstanceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
