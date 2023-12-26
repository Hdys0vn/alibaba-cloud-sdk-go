package nas

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

// StartDataFlow invokes the nas.StartDataFlow API synchronously
func (client *Client) StartDataFlow(request *StartDataFlowRequest) (response *StartDataFlowResponse, err error) {
	response = CreateStartDataFlowResponse()
	err = client.DoAction(request, response)
	return
}

// StartDataFlowWithChan invokes the nas.StartDataFlow API asynchronously
func (client *Client) StartDataFlowWithChan(request *StartDataFlowRequest) (<-chan *StartDataFlowResponse, <-chan error) {
	responseChan := make(chan *StartDataFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartDataFlow(request)
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

// StartDataFlowWithCallback invokes the nas.StartDataFlow API asynchronously
func (client *Client) StartDataFlowWithCallback(request *StartDataFlowRequest, callback func(response *StartDataFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartDataFlowResponse
		var err error
		defer close(result)
		response, err = client.StartDataFlow(request)
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

// StartDataFlowRequest is the request struct for api StartDataFlow
type StartDataFlowRequest struct {
	*requests.RpcRequest
	ClientToken  string           `position:"Query" name:"ClientToken"`
	FileSystemId string           `position:"Query" name:"FileSystemId"`
	DryRun       requests.Boolean `position:"Query" name:"DryRun"`
	DataFlowId   string           `position:"Query" name:"DataFlowId"`
}

// StartDataFlowResponse is the response struct for api StartDataFlow
type StartDataFlowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartDataFlowRequest creates a request to invoke StartDataFlow API
func CreateStartDataFlowRequest() (request *StartDataFlowRequest) {
	request = &StartDataFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "StartDataFlow", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartDataFlowResponse creates a response to parse from StartDataFlow response
func CreateStartDataFlowResponse() (response *StartDataFlowResponse) {
	response = &StartDataFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}