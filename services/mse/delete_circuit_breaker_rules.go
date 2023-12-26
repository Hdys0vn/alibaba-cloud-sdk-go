package mse

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

// DeleteCircuitBreakerRules invokes the mse.DeleteCircuitBreakerRules API synchronously
func (client *Client) DeleteCircuitBreakerRules(request *DeleteCircuitBreakerRulesRequest) (response *DeleteCircuitBreakerRulesResponse, err error) {
	response = CreateDeleteCircuitBreakerRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCircuitBreakerRulesWithChan invokes the mse.DeleteCircuitBreakerRules API asynchronously
func (client *Client) DeleteCircuitBreakerRulesWithChan(request *DeleteCircuitBreakerRulesRequest) (<-chan *DeleteCircuitBreakerRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteCircuitBreakerRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCircuitBreakerRules(request)
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

// DeleteCircuitBreakerRulesWithCallback invokes the mse.DeleteCircuitBreakerRules API asynchronously
func (client *Client) DeleteCircuitBreakerRulesWithCallback(request *DeleteCircuitBreakerRulesRequest, callback func(response *DeleteCircuitBreakerRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCircuitBreakerRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteCircuitBreakerRules(request)
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

// DeleteCircuitBreakerRulesRequest is the request struct for api DeleteCircuitBreakerRules
type DeleteCircuitBreakerRulesRequest struct {
	*requests.RpcRequest
	MseSessionId   string    `position:"Query" name:"MseSessionId"`
	AppName        string    `position:"Query" name:"AppName"`
	Namespace      string    `position:"Query" name:"Namespace"`
	AcceptLanguage string    `position:"Query" name:"AcceptLanguage"`
	Ids            *[]string `position:"Query" name:"Ids"  type:"Json"`
}

// DeleteCircuitBreakerRulesResponse is the response struct for api DeleteCircuitBreakerRules
type DeleteCircuitBreakerRulesResponse struct {
	*responses.BaseResponse
	Code           int     `json:"Code" xml:"Code"`
	Message        string  `json:"Message" xml:"Message"`
	RequestId      string  `json:"RequestId" xml:"RequestId"`
	Success        bool    `json:"Success" xml:"Success"`
	HttpStatusCode int     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           []int64 `json:"Data" xml:"Data"`
}

// CreateDeleteCircuitBreakerRulesRequest creates a request to invoke DeleteCircuitBreakerRules API
func CreateDeleteCircuitBreakerRulesRequest() (request *DeleteCircuitBreakerRulesRequest) {
	request = &DeleteCircuitBreakerRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteCircuitBreakerRules", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCircuitBreakerRulesResponse creates a response to parse from DeleteCircuitBreakerRules response
func CreateDeleteCircuitBreakerRulesResponse() (response *DeleteCircuitBreakerRulesResponse) {
	response = &DeleteCircuitBreakerRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
