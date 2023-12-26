package outboundbot

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

// DeleteDialogueFlow invokes the outboundbot.DeleteDialogueFlow API synchronously
func (client *Client) DeleteDialogueFlow(request *DeleteDialogueFlowRequest) (response *DeleteDialogueFlowResponse, err error) {
	response = CreateDeleteDialogueFlowResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDialogueFlowWithChan invokes the outboundbot.DeleteDialogueFlow API asynchronously
func (client *Client) DeleteDialogueFlowWithChan(request *DeleteDialogueFlowRequest) (<-chan *DeleteDialogueFlowResponse, <-chan error) {
	responseChan := make(chan *DeleteDialogueFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDialogueFlow(request)
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

// DeleteDialogueFlowWithCallback invokes the outboundbot.DeleteDialogueFlow API asynchronously
func (client *Client) DeleteDialogueFlowWithCallback(request *DeleteDialogueFlowRequest, callback func(response *DeleteDialogueFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDialogueFlowResponse
		var err error
		defer close(result)
		response, err = client.DeleteDialogueFlow(request)
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

// DeleteDialogueFlowRequest is the request struct for api DeleteDialogueFlow
type DeleteDialogueFlowRequest struct {
	*requests.RpcRequest
	ScriptId       string `position:"Query" name:"ScriptId"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	DialogueFlowId string `position:"Query" name:"DialogueFlowId"`
}

// DeleteDialogueFlowResponse is the response struct for api DeleteDialogueFlow
type DeleteDialogueFlowResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeleteDialogueFlowRequest creates a request to invoke DeleteDialogueFlow API
func CreateDeleteDialogueFlowRequest() (request *DeleteDialogueFlowRequest) {
	request = &DeleteDialogueFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DeleteDialogueFlow", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDialogueFlowResponse creates a response to parse from DeleteDialogueFlow response
func CreateDeleteDialogueFlowResponse() (response *DeleteDialogueFlowResponse) {
	response = &DeleteDialogueFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}