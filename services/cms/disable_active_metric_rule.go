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

// DisableActiveMetricRule invokes the cms.DisableActiveMetricRule API synchronously
func (client *Client) DisableActiveMetricRule(request *DisableActiveMetricRuleRequest) (response *DisableActiveMetricRuleResponse, err error) {
	response = CreateDisableActiveMetricRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DisableActiveMetricRuleWithChan invokes the cms.DisableActiveMetricRule API asynchronously
func (client *Client) DisableActiveMetricRuleWithChan(request *DisableActiveMetricRuleRequest) (<-chan *DisableActiveMetricRuleResponse, <-chan error) {
	responseChan := make(chan *DisableActiveMetricRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableActiveMetricRule(request)
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

// DisableActiveMetricRuleWithCallback invokes the cms.DisableActiveMetricRule API asynchronously
func (client *Client) DisableActiveMetricRuleWithCallback(request *DisableActiveMetricRuleRequest, callback func(response *DisableActiveMetricRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableActiveMetricRuleResponse
		var err error
		defer close(result)
		response, err = client.DisableActiveMetricRule(request)
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

// DisableActiveMetricRuleRequest is the request struct for api DisableActiveMetricRule
type DisableActiveMetricRuleRequest struct {
	*requests.RpcRequest
	Product string `position:"Query" name:"Product"`
}

// DisableActiveMetricRuleResponse is the response struct for api DisableActiveMetricRule
type DisableActiveMetricRuleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDisableActiveMetricRuleRequest creates a request to invoke DisableActiveMetricRule API
func CreateDisableActiveMetricRuleRequest() (request *DisableActiveMetricRuleRequest) {
	request = &DisableActiveMetricRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DisableActiveMetricRule", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableActiveMetricRuleResponse creates a response to parse from DisableActiveMetricRule response
func CreateDisableActiveMetricRuleResponse() (response *DisableActiveMetricRuleResponse) {
	response = &DisableActiveMetricRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
