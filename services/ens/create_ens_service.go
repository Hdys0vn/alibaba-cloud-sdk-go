package ens

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

// CreateEnsService invokes the ens.CreateEnsService API synchronously
func (client *Client) CreateEnsService(request *CreateEnsServiceRequest) (response *CreateEnsServiceResponse, err error) {
	response = CreateCreateEnsServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEnsServiceWithChan invokes the ens.CreateEnsService API asynchronously
func (client *Client) CreateEnsServiceWithChan(request *CreateEnsServiceRequest) (<-chan *CreateEnsServiceResponse, <-chan error) {
	responseChan := make(chan *CreateEnsServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEnsService(request)
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

// CreateEnsServiceWithCallback invokes the ens.CreateEnsService API asynchronously
func (client *Client) CreateEnsServiceWithCallback(request *CreateEnsServiceRequest, callback func(response *CreateEnsServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEnsServiceResponse
		var err error
		defer close(result)
		response, err = client.CreateEnsService(request)
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

// CreateEnsServiceRequest is the request struct for api CreateEnsService
type CreateEnsServiceRequest struct {
	*requests.RpcRequest
	EnsServiceId string `position:"Query" name:"EnsServiceId"`
	OrderType    string `position:"Query" name:"OrderType"`
}

// CreateEnsServiceResponse is the response struct for api CreateEnsService
type CreateEnsServiceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateEnsServiceRequest creates a request to invoke CreateEnsService API
func CreateCreateEnsServiceRequest() (request *CreateEnsServiceRequest) {
	request = &CreateEnsServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateEnsService", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateEnsServiceResponse creates a response to parse from CreateEnsService response
func CreateCreateEnsServiceResponse() (response *CreateEnsServiceResponse) {
	response = &CreateEnsServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
