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

// UpdateForwardingRules invokes the ga.UpdateForwardingRules API synchronously
func (client *Client) UpdateForwardingRules(request *UpdateForwardingRulesRequest) (response *UpdateForwardingRulesResponse, err error) {
	response = CreateUpdateForwardingRulesResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateForwardingRulesWithChan invokes the ga.UpdateForwardingRules API asynchronously
func (client *Client) UpdateForwardingRulesWithChan(request *UpdateForwardingRulesRequest) (<-chan *UpdateForwardingRulesResponse, <-chan error) {
	responseChan := make(chan *UpdateForwardingRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateForwardingRules(request)
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

// UpdateForwardingRulesWithCallback invokes the ga.UpdateForwardingRules API asynchronously
func (client *Client) UpdateForwardingRulesWithCallback(request *UpdateForwardingRulesRequest, callback func(response *UpdateForwardingRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateForwardingRulesResponse
		var err error
		defer close(result)
		response, err = client.UpdateForwardingRules(request)
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

// UpdateForwardingRulesRequest is the request struct for api UpdateForwardingRules
type UpdateForwardingRulesRequest struct {
	*requests.RpcRequest
	ClientToken     string `position:"Query" name:"ClientToken"`
	ListenerId      string `position:"Query" name:"ListenerId"`
	AcceleratorId   string `position:"Query" name:"AcceleratorId"`
	ForwardingRules string `position:"Query" name:"ForwardingRules"`
}

// UpdateForwardingRulesResponse is the response struct for api UpdateForwardingRules
type UpdateForwardingRulesResponse struct {
	*responses.BaseResponse
	RequestId       string                `json:"RequestId" xml:"RequestId"`
	ForwardingRules []ForwardingRulesItem `json:"ForwardingRules" xml:"ForwardingRules"`
}

// CreateUpdateForwardingRulesRequest creates a request to invoke UpdateForwardingRules API
func CreateUpdateForwardingRulesRequest() (request *UpdateForwardingRulesRequest) {
	request = &UpdateForwardingRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "UpdateForwardingRules", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateForwardingRulesResponse creates a response to parse from UpdateForwardingRules response
func CreateUpdateForwardingRulesResponse() (response *UpdateForwardingRulesResponse) {
	response = &UpdateForwardingRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
