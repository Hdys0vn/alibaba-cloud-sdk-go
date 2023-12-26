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

// ListScriptVoiceConfigs invokes the outboundbot.ListScriptVoiceConfigs API synchronously
func (client *Client) ListScriptVoiceConfigs(request *ListScriptVoiceConfigsRequest) (response *ListScriptVoiceConfigsResponse, err error) {
	response = CreateListScriptVoiceConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// ListScriptVoiceConfigsWithChan invokes the outboundbot.ListScriptVoiceConfigs API asynchronously
func (client *Client) ListScriptVoiceConfigsWithChan(request *ListScriptVoiceConfigsRequest) (<-chan *ListScriptVoiceConfigsResponse, <-chan error) {
	responseChan := make(chan *ListScriptVoiceConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListScriptVoiceConfigs(request)
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

// ListScriptVoiceConfigsWithCallback invokes the outboundbot.ListScriptVoiceConfigs API asynchronously
func (client *Client) ListScriptVoiceConfigsWithCallback(request *ListScriptVoiceConfigsRequest, callback func(response *ListScriptVoiceConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListScriptVoiceConfigsResponse
		var err error
		defer close(result)
		response, err = client.ListScriptVoiceConfigs(request)
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

// ListScriptVoiceConfigsRequest is the request struct for api ListScriptVoiceConfigs
type ListScriptVoiceConfigsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	ScriptId   string           `position:"Query" name:"ScriptId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListScriptVoiceConfigsResponse is the response struct for api ListScriptVoiceConfigs
type ListScriptVoiceConfigsResponse struct {
	*responses.BaseResponse
	HttpStatusCode     int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code               string             `json:"Code" xml:"Code"`
	Message            string             `json:"Message" xml:"Message"`
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	Success            bool               `json:"Success" xml:"Success"`
	ScriptVoiceConfigs ScriptVoiceConfigs `json:"ScriptVoiceConfigs" xml:"ScriptVoiceConfigs"`
}

// CreateListScriptVoiceConfigsRequest creates a request to invoke ListScriptVoiceConfigs API
func CreateListScriptVoiceConfigsRequest() (request *ListScriptVoiceConfigsRequest) {
	request = &ListScriptVoiceConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListScriptVoiceConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateListScriptVoiceConfigsResponse creates a response to parse from ListScriptVoiceConfigs response
func CreateListScriptVoiceConfigsResponse() (response *ListScriptVoiceConfigsResponse) {
	response = &ListScriptVoiceConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}