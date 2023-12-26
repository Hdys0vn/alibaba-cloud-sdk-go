package sas

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

// OperationCancelIgnoreSuspEvent invokes the sas.OperationCancelIgnoreSuspEvent API synchronously
func (client *Client) OperationCancelIgnoreSuspEvent(request *OperationCancelIgnoreSuspEventRequest) (response *OperationCancelIgnoreSuspEventResponse, err error) {
	response = CreateOperationCancelIgnoreSuspEventResponse()
	err = client.DoAction(request, response)
	return
}

// OperationCancelIgnoreSuspEventWithChan invokes the sas.OperationCancelIgnoreSuspEvent API asynchronously
func (client *Client) OperationCancelIgnoreSuspEventWithChan(request *OperationCancelIgnoreSuspEventRequest) (<-chan *OperationCancelIgnoreSuspEventResponse, <-chan error) {
	responseChan := make(chan *OperationCancelIgnoreSuspEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperationCancelIgnoreSuspEvent(request)
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

// OperationCancelIgnoreSuspEventWithCallback invokes the sas.OperationCancelIgnoreSuspEvent API asynchronously
func (client *Client) OperationCancelIgnoreSuspEventWithCallback(request *OperationCancelIgnoreSuspEventRequest, callback func(response *OperationCancelIgnoreSuspEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperationCancelIgnoreSuspEventResponse
		var err error
		defer close(result)
		response, err = client.OperationCancelIgnoreSuspEvent(request)
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

// OperationCancelIgnoreSuspEventRequest is the request struct for api OperationCancelIgnoreSuspEvent
type OperationCancelIgnoreSuspEventRequest struct {
	*requests.RpcRequest
	SecurityEventIds *[]string `position:"Query" name:"SecurityEventIds"  type:"Repeated"`
}

// OperationCancelIgnoreSuspEventResponse is the response struct for api OperationCancelIgnoreSuspEvent
type OperationCancelIgnoreSuspEventResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	TimeCost       int64  `json:"TimeCost" xml:"TimeCost"`
}

// CreateOperationCancelIgnoreSuspEventRequest creates a request to invoke OperationCancelIgnoreSuspEvent API
func CreateOperationCancelIgnoreSuspEventRequest() (request *OperationCancelIgnoreSuspEventRequest) {
	request = &OperationCancelIgnoreSuspEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "OperationCancelIgnoreSuspEvent", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOperationCancelIgnoreSuspEventResponse creates a response to parse from OperationCancelIgnoreSuspEvent response
func CreateOperationCancelIgnoreSuspEventResponse() (response *OperationCancelIgnoreSuspEventResponse) {
	response = &OperationCancelIgnoreSuspEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
