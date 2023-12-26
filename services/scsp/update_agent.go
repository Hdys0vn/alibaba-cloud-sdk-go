package scsp

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

// UpdateAgent invokes the scsp.UpdateAgent API synchronously
func (client *Client) UpdateAgent(request *UpdateAgentRequest) (response *UpdateAgentResponse, err error) {
	response = CreateUpdateAgentResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAgentWithChan invokes the scsp.UpdateAgent API asynchronously
func (client *Client) UpdateAgentWithChan(request *UpdateAgentRequest) (<-chan *UpdateAgentResponse, <-chan error) {
	responseChan := make(chan *UpdateAgentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAgent(request)
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

// UpdateAgentWithCallback invokes the scsp.UpdateAgent API asynchronously
func (client *Client) UpdateAgentWithCallback(request *UpdateAgentRequest, callback func(response *UpdateAgentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAgentResponse
		var err error
		defer close(result)
		response, err = client.UpdateAgent(request)
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

// UpdateAgentRequest is the request struct for api UpdateAgent
type UpdateAgentRequest struct {
	*requests.RpcRequest
	ClientToken      string    `position:"Body"`
	InstanceId       string    `position:"Body"`
	AccountName      string    `position:"Body"`
	DisplayName      string    `position:"Body"`
	SkillGroupId     *[]string `position:"Body" name:"SkillGroupId"  type:"Repeated"`
	SkillGroupIdList *[]string `position:"Body" name:"SkillGroupIdList"  type:"Repeated"`
}

// UpdateAgentResponse is the response struct for api UpdateAgent
type UpdateAgentResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateUpdateAgentRequest creates a request to invoke UpdateAgent API
func CreateUpdateAgentRequest() (request *UpdateAgentRequest) {
	request = &UpdateAgentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "UpdateAgent", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateAgentResponse creates a response to parse from UpdateAgent response
func CreateUpdateAgentResponse() (response *UpdateAgentResponse) {
	response = &UpdateAgentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
