package ga

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

// DeleteForwardingRules invokes the ga.DeleteForwardingRules API synchronously
func (client *Client) DeleteForwardingRules(request *DeleteForwardingRulesRequest) (response *DeleteForwardingRulesResponse, err error) {
	response = CreateDeleteForwardingRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteForwardingRulesWithChan invokes the ga.DeleteForwardingRules API asynchronously
func (client *Client) DeleteForwardingRulesWithChan(request *DeleteForwardingRulesRequest) (<-chan *DeleteForwardingRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteForwardingRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteForwardingRules(request)
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

// DeleteForwardingRulesWithCallback invokes the ga.DeleteForwardingRules API asynchronously
func (client *Client) DeleteForwardingRulesWithCallback(request *DeleteForwardingRulesRequest, callback func(response *DeleteForwardingRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteForwardingRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteForwardingRules(request)
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

// DeleteForwardingRulesRequest is the request struct for api DeleteForwardingRules
type DeleteForwardingRulesRequest struct {
	*requests.RpcRequest
	ClientToken       string `position:"Query" name:"ClientToken"`
	ListenerId        string `position:"Query" name:"ListenerId"`
	AcceleratorId     string `position:"Query" name:"AcceleratorId"`
	ForwardingRuleIds string `position:"Query" name:"ForwardingRuleIds"`
}

// DeleteForwardingRulesResponse is the response struct for api DeleteForwardingRules
type DeleteForwardingRulesResponse struct {
	*responses.BaseResponse
	RequestId       string                `json:"RequestId" xml:"RequestId"`
	ForwardingRules []ForwardingRulesItem `json:"ForwardingRules" xml:"ForwardingRules"`
}

// CreateDeleteForwardingRulesRequest creates a request to invoke DeleteForwardingRules API
func CreateDeleteForwardingRulesRequest() (request *DeleteForwardingRulesRequest) {
	request = &DeleteForwardingRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "DeleteForwardingRules", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteForwardingRulesResponse creates a response to parse from DeleteForwardingRules response
func CreateDeleteForwardingRulesResponse() (response *DeleteForwardingRulesResponse) {
	response = &DeleteForwardingRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
