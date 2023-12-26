package sls

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

// CreateResourceOrchestration invokes the sls.CreateResourceOrchestration API synchronously
func (client *Client) CreateResourceOrchestration(request *CreateResourceOrchestrationRequest) (response *CreateResourceOrchestrationResponse, err error) {
	response = CreateCreateResourceOrchestrationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateResourceOrchestrationWithChan invokes the sls.CreateResourceOrchestration API asynchronously
func (client *Client) CreateResourceOrchestrationWithChan(request *CreateResourceOrchestrationRequest) (<-chan *CreateResourceOrchestrationResponse, <-chan error) {
	responseChan := make(chan *CreateResourceOrchestrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateResourceOrchestration(request)
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

// CreateResourceOrchestrationWithCallback invokes the sls.CreateResourceOrchestration API asynchronously
func (client *Client) CreateResourceOrchestrationWithCallback(request *CreateResourceOrchestrationRequest, callback func(response *CreateResourceOrchestrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateResourceOrchestrationResponse
		var err error
		defer close(result)
		response, err = client.CreateResourceOrchestration(request)
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

// CreateResourceOrchestrationRequest is the request struct for api CreateResourceOrchestration
type CreateResourceOrchestrationRequest struct {
	*requests.RpcRequest
	Language        string `position:"Query" name:"Language"`
	OperationPolicy string `position:"Query" name:"OperationPolicy"`
	Tokens          string `position:"Query" name:"Tokens"`
	AssetsId        string `position:"Query" name:"AssetsId"`
}

// CreateResourceOrchestrationResponse is the response struct for api CreateResourceOrchestration
type CreateResourceOrchestrationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCreateResourceOrchestrationRequest creates a request to invoke CreateResourceOrchestration API
func CreateCreateResourceOrchestrationRequest() (request *CreateResourceOrchestrationRequest) {
	request = &CreateResourceOrchestrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "CreateResourceOrchestration", "sls", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateResourceOrchestrationResponse creates a response to parse from CreateResourceOrchestration response
func CreateCreateResourceOrchestrationResponse() (response *CreateResourceOrchestrationResponse) {
	response = &CreateResourceOrchestrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
