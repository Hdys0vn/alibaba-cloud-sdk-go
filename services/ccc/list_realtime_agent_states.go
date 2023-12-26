package ccc

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

// ListRealtimeAgentStates invokes the ccc.ListRealtimeAgentStates API synchronously
func (client *Client) ListRealtimeAgentStates(request *ListRealtimeAgentStatesRequest) (response *ListRealtimeAgentStatesResponse, err error) {
	response = CreateListRealtimeAgentStatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListRealtimeAgentStatesWithChan invokes the ccc.ListRealtimeAgentStates API asynchronously
func (client *Client) ListRealtimeAgentStatesWithChan(request *ListRealtimeAgentStatesRequest) (<-chan *ListRealtimeAgentStatesResponse, <-chan error) {
	responseChan := make(chan *ListRealtimeAgentStatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRealtimeAgentStates(request)
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

// ListRealtimeAgentStatesWithCallback invokes the ccc.ListRealtimeAgentStates API asynchronously
func (client *Client) ListRealtimeAgentStatesWithCallback(request *ListRealtimeAgentStatesRequest, callback func(response *ListRealtimeAgentStatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRealtimeAgentStatesResponse
		var err error
		defer close(result)
		response, err = client.ListRealtimeAgentStates(request)
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

// ListRealtimeAgentStatesRequest is the request struct for api ListRealtimeAgentStates
type ListRealtimeAgentStatesRequest struct {
	*requests.RpcRequest
	CallTypeList     string           `position:"Query" name:"CallTypeList"`
	Query            string           `position:"Query" name:"Query"`
	OutboundScenario requests.Boolean `position:"Query" name:"OutboundScenario"`
	PageNumber       requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
	WorkModeList     string           `position:"Query" name:"WorkModeList"`
	AgentIdList      string           `position:"Body" name:"AgentIdList"`
	SkillGroupId     string           `position:"Query" name:"SkillGroupId"`
	AgentName        string           `position:"Query" name:"AgentName"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	StateList        string           `position:"Body" name:"StateList"`
}

// ListRealtimeAgentStatesResponse is the response struct for api ListRealtimeAgentStates
type ListRealtimeAgentStatesResponse struct {
	*responses.BaseResponse
	Code           string                        `json:"Code" xml:"Code"`
	HttpStatusCode int                           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                        `json:"Message" xml:"Message"`
	RequestId      string                        `json:"RequestId" xml:"RequestId"`
	Data           DataInListRealtimeAgentStates `json:"Data" xml:"Data"`
}

// CreateListRealtimeAgentStatesRequest creates a request to invoke ListRealtimeAgentStates API
func CreateListRealtimeAgentStatesRequest() (request *ListRealtimeAgentStatesRequest) {
	request = &ListRealtimeAgentStatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListRealtimeAgentStates", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRealtimeAgentStatesResponse creates a response to parse from ListRealtimeAgentStates response
func CreateListRealtimeAgentStatesResponse() (response *ListRealtimeAgentStatesResponse) {
	response = &ListRealtimeAgentStatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
