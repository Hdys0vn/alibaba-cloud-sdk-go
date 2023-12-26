package retailadvqa_public

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

// CreateDefaultEvent invokes the retailadvqa_public.CreateDefaultEvent API synchronously
func (client *Client) CreateDefaultEvent(request *CreateDefaultEventRequest) (response *CreateDefaultEventResponse, err error) {
	response = CreateCreateDefaultEventResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDefaultEventWithChan invokes the retailadvqa_public.CreateDefaultEvent API asynchronously
func (client *Client) CreateDefaultEventWithChan(request *CreateDefaultEventRequest) (<-chan *CreateDefaultEventResponse, <-chan error) {
	responseChan := make(chan *CreateDefaultEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDefaultEvent(request)
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

// CreateDefaultEventWithCallback invokes the retailadvqa_public.CreateDefaultEvent API asynchronously
func (client *Client) CreateDefaultEventWithCallback(request *CreateDefaultEventRequest, callback func(response *CreateDefaultEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDefaultEventResponse
		var err error
		defer close(result)
		response, err = client.CreateDefaultEvent(request)
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

// CreateDefaultEventRequest is the request struct for api CreateDefaultEvent
type CreateDefaultEventRequest struct {
	*requests.RpcRequest
	AccessId      string `position:"Query" name:"AccessId"`
	TenantId      string `position:"Query" name:"TenantId"`
	EventModelStr string `position:"Body" name:"EventModelStr"`
	WorkspaceId   string `position:"Query" name:"WorkspaceId"`
}

// CreateDefaultEventResponse is the response struct for api CreateDefaultEvent
type CreateDefaultEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateDefaultEventRequest creates a request to invoke CreateDefaultEvent API
func CreateCreateDefaultEventRequest() (request *CreateDefaultEventRequest) {
	request = &CreateDefaultEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "CreateDefaultEvent", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDefaultEventResponse creates a response to parse from CreateDefaultEvent response
func CreateCreateDefaultEventResponse() (response *CreateDefaultEventResponse) {
	response = &CreateDefaultEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
