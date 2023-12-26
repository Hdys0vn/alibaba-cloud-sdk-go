package swas_open

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

// DeleteFirewallRule invokes the swas_open.DeleteFirewallRule API synchronously
func (client *Client) DeleteFirewallRule(request *DeleteFirewallRuleRequest) (response *DeleteFirewallRuleResponse, err error) {
	response = CreateDeleteFirewallRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFirewallRuleWithChan invokes the swas_open.DeleteFirewallRule API asynchronously
func (client *Client) DeleteFirewallRuleWithChan(request *DeleteFirewallRuleRequest) (<-chan *DeleteFirewallRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteFirewallRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFirewallRule(request)
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

// DeleteFirewallRuleWithCallback invokes the swas_open.DeleteFirewallRule API asynchronously
func (client *Client) DeleteFirewallRuleWithCallback(request *DeleteFirewallRuleRequest, callback func(response *DeleteFirewallRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFirewallRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteFirewallRule(request)
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

// DeleteFirewallRuleRequest is the request struct for api DeleteFirewallRule
type DeleteFirewallRuleRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	RuleId      string `position:"Query" name:"RuleId"`
}

// DeleteFirewallRuleResponse is the response struct for api DeleteFirewallRule
type DeleteFirewallRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFirewallRuleRequest creates a request to invoke DeleteFirewallRule API
func CreateDeleteFirewallRuleRequest() (request *DeleteFirewallRuleRequest) {
	request = &DeleteFirewallRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DeleteFirewallRule", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFirewallRuleResponse creates a response to parse from DeleteFirewallRule response
func CreateDeleteFirewallRuleResponse() (response *DeleteFirewallRuleResponse) {
	response = &DeleteFirewallRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
