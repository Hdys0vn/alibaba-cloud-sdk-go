package polardb

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

// EnableFirewallRules invokes the polardb.EnableFirewallRules API synchronously
func (client *Client) EnableFirewallRules(request *EnableFirewallRulesRequest) (response *EnableFirewallRulesResponse, err error) {
	response = CreateEnableFirewallRulesResponse()
	err = client.DoAction(request, response)
	return
}

// EnableFirewallRulesWithChan invokes the polardb.EnableFirewallRules API asynchronously
func (client *Client) EnableFirewallRulesWithChan(request *EnableFirewallRulesRequest) (<-chan *EnableFirewallRulesResponse, <-chan error) {
	responseChan := make(chan *EnableFirewallRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableFirewallRules(request)
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

// EnableFirewallRulesWithCallback invokes the polardb.EnableFirewallRules API asynchronously
func (client *Client) EnableFirewallRulesWithCallback(request *EnableFirewallRulesRequest, callback func(response *EnableFirewallRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableFirewallRulesResponse
		var err error
		defer close(result)
		response, err = client.EnableFirewallRules(request)
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

// EnableFirewallRulesRequest is the request struct for api EnableFirewallRules
type EnableFirewallRulesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Enable               requests.Boolean `position:"Query" name:"Enable"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RuleNameList         string           `position:"Query" name:"RuleNameList"`
}

// EnableFirewallRulesResponse is the response struct for api EnableFirewallRules
type EnableFirewallRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateEnableFirewallRulesRequest creates a request to invoke EnableFirewallRules API
func CreateEnableFirewallRulesRequest() (request *EnableFirewallRulesRequest) {
	request = &EnableFirewallRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "EnableFirewallRules", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableFirewallRulesResponse creates a response to parse from EnableFirewallRules response
func CreateEnableFirewallRulesResponse() (response *EnableFirewallRulesResponse) {
	response = &EnableFirewallRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
