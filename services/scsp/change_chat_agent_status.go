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

// ChangeChatAgentStatus invokes the scsp.ChangeChatAgentStatus API synchronously
func (client *Client) ChangeChatAgentStatus(request *ChangeChatAgentStatusRequest) (response *ChangeChatAgentStatusResponse, err error) {
	response = CreateChangeChatAgentStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeChatAgentStatusWithChan invokes the scsp.ChangeChatAgentStatus API asynchronously
func (client *Client) ChangeChatAgentStatusWithChan(request *ChangeChatAgentStatusRequest) (<-chan *ChangeChatAgentStatusResponse, <-chan error) {
	responseChan := make(chan *ChangeChatAgentStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeChatAgentStatus(request)
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

// ChangeChatAgentStatusWithCallback invokes the scsp.ChangeChatAgentStatus API asynchronously
func (client *Client) ChangeChatAgentStatusWithCallback(request *ChangeChatAgentStatusRequest, callback func(response *ChangeChatAgentStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeChatAgentStatusResponse
		var err error
		defer close(result)
		response, err = client.ChangeChatAgentStatus(request)
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

// ChangeChatAgentStatusRequest is the request struct for api ChangeChatAgentStatus
type ChangeChatAgentStatusRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Body"`
	InstanceId  string `position:"Body"`
	AccountName string `position:"Body"`
	Method      string `position:"Body"`
}

// ChangeChatAgentStatusResponse is the response struct for api ChangeChatAgentStatus
type ChangeChatAgentStatusResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           string `json:"Data" xml:"Data"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateChangeChatAgentStatusRequest creates a request to invoke ChangeChatAgentStatus API
func CreateChangeChatAgentStatusRequest() (request *ChangeChatAgentStatusRequest) {
	request = &ChangeChatAgentStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "ChangeChatAgentStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateChangeChatAgentStatusResponse creates a response to parse from ChangeChatAgentStatus response
func CreateChangeChatAgentStatusResponse() (response *ChangeChatAgentStatusResponse) {
	response = &ChangeChatAgentStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
