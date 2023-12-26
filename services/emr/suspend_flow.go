package emr

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

// SuspendFlow invokes the emr.SuspendFlow API synchronously
func (client *Client) SuspendFlow(request *SuspendFlowRequest) (response *SuspendFlowResponse, err error) {
	response = CreateSuspendFlowResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendFlowWithChan invokes the emr.SuspendFlow API asynchronously
func (client *Client) SuspendFlowWithChan(request *SuspendFlowRequest) (<-chan *SuspendFlowResponse, <-chan error) {
	responseChan := make(chan *SuspendFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendFlow(request)
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

// SuspendFlowWithCallback invokes the emr.SuspendFlow API asynchronously
func (client *Client) SuspendFlowWithCallback(request *SuspendFlowRequest, callback func(response *SuspendFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendFlowResponse
		var err error
		defer close(result)
		response, err = client.SuspendFlow(request)
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

// SuspendFlowRequest is the request struct for api SuspendFlow
type SuspendFlowRequest struct {
	*requests.RpcRequest
	FlowInstanceId string `position:"Query" name:"FlowInstanceId"`
	ProjectId      string `position:"Query" name:"ProjectId"`
}

// SuspendFlowResponse is the response struct for api SuspendFlow
type SuspendFlowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateSuspendFlowRequest creates a request to invoke SuspendFlow API
func CreateSuspendFlowRequest() (request *SuspendFlowRequest) {
	request = &SuspendFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "SuspendFlow", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendFlowResponse creates a response to parse from SuspendFlow response
func CreateSuspendFlowResponse() (response *SuspendFlowResponse) {
	response = &SuspendFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}