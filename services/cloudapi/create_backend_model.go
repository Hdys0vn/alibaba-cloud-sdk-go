package cloudapi

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

// CreateBackendModel invokes the cloudapi.CreateBackendModel API synchronously
func (client *Client) CreateBackendModel(request *CreateBackendModelRequest) (response *CreateBackendModelResponse, err error) {
	response = CreateCreateBackendModelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBackendModelWithChan invokes the cloudapi.CreateBackendModel API asynchronously
func (client *Client) CreateBackendModelWithChan(request *CreateBackendModelRequest) (<-chan *CreateBackendModelResponse, <-chan error) {
	responseChan := make(chan *CreateBackendModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBackendModel(request)
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

// CreateBackendModelWithCallback invokes the cloudapi.CreateBackendModel API asynchronously
func (client *Client) CreateBackendModelWithCallback(request *CreateBackendModelRequest, callback func(response *CreateBackendModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBackendModelResponse
		var err error
		defer close(result)
		response, err = client.CreateBackendModel(request)
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

// CreateBackendModelRequest is the request struct for api CreateBackendModel
type CreateBackendModelRequest struct {
	*requests.RpcRequest
	StageName        string `position:"Query" name:"StageName"`
	BackendModelData string `position:"Query" name:"BackendModelData"`
	BackendId        string `position:"Query" name:"BackendId"`
	Description      string `position:"Query" name:"Description"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	BackendType      string `position:"Query" name:"BackendType"`
}

// CreateBackendModelResponse is the response struct for api CreateBackendModel
type CreateBackendModelResponse struct {
	*responses.BaseResponse
	BackendModelId string `json:"BackendModelId" xml:"BackendModelId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateBackendModelRequest creates a request to invoke CreateBackendModel API
func CreateCreateBackendModelRequest() (request *CreateBackendModelRequest) {
	request = &CreateBackendModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateBackendModel", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBackendModelResponse creates a response to parse from CreateBackendModel response
func CreateCreateBackendModelResponse() (response *CreateBackendModelResponse) {
	response = &CreateBackendModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}