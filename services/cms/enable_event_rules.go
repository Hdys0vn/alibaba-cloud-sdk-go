package cms

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

// EnableEventRules invokes the cms.EnableEventRules API synchronously
func (client *Client) EnableEventRules(request *EnableEventRulesRequest) (response *EnableEventRulesResponse, err error) {
	response = CreateEnableEventRulesResponse()
	err = client.DoAction(request, response)
	return
}

// EnableEventRulesWithChan invokes the cms.EnableEventRules API asynchronously
func (client *Client) EnableEventRulesWithChan(request *EnableEventRulesRequest) (<-chan *EnableEventRulesResponse, <-chan error) {
	responseChan := make(chan *EnableEventRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableEventRules(request)
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

// EnableEventRulesWithCallback invokes the cms.EnableEventRules API asynchronously
func (client *Client) EnableEventRulesWithCallback(request *EnableEventRulesRequest, callback func(response *EnableEventRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableEventRulesResponse
		var err error
		defer close(result)
		response, err = client.EnableEventRules(request)
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

// EnableEventRulesRequest is the request struct for api EnableEventRules
type EnableEventRulesRequest struct {
	*requests.RpcRequest
	RuleNames *[]string `position:"Query" name:"RuleNames"  type:"Repeated"`
}

// EnableEventRulesResponse is the response struct for api EnableEventRules
type EnableEventRulesResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateEnableEventRulesRequest creates a request to invoke EnableEventRules API
func CreateEnableEventRulesRequest() (request *EnableEventRulesRequest) {
	request = &EnableEventRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "EnableEventRules", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableEventRulesResponse creates a response to parse from EnableEventRules response
func CreateEnableEventRulesResponse() (response *EnableEventRulesResponse) {
	response = &EnableEventRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
