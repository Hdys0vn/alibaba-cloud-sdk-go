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

// DeleteAgent invokes the scsp.DeleteAgent API synchronously
func (client *Client) DeleteAgent(request *DeleteAgentRequest) (response *DeleteAgentResponse, err error) {
	response = CreateDeleteAgentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAgentWithChan invokes the scsp.DeleteAgent API asynchronously
func (client *Client) DeleteAgentWithChan(request *DeleteAgentRequest) (<-chan *DeleteAgentResponse, <-chan error) {
	responseChan := make(chan *DeleteAgentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAgent(request)
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

// DeleteAgentWithCallback invokes the scsp.DeleteAgent API asynchronously
func (client *Client) DeleteAgentWithCallback(request *DeleteAgentRequest, callback func(response *DeleteAgentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAgentResponse
		var err error
		defer close(result)
		response, err = client.DeleteAgent(request)
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

// DeleteAgentRequest is the request struct for api DeleteAgent
type DeleteAgentRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query"`
	InstanceId  string `position:"Query"`
	AccountName string `position:"Query"`
}

// DeleteAgentResponse is the response struct for api DeleteAgent
type DeleteAgentResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeleteAgentRequest creates a request to invoke DeleteAgent API
func CreateDeleteAgentRequest() (request *DeleteAgentRequest) {
	request = &DeleteAgentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "DeleteAgent", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteAgentResponse creates a response to parse from DeleteAgent response
func CreateDeleteAgentResponse() (response *DeleteAgentResponse) {
	response = &DeleteAgentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}