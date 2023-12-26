package sddp

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

// DescribeTemplateAllRules invokes the sddp.DescribeTemplateAllRules API synchronously
func (client *Client) DescribeTemplateAllRules(request *DescribeTemplateAllRulesRequest) (response *DescribeTemplateAllRulesResponse, err error) {
	response = CreateDescribeTemplateAllRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTemplateAllRulesWithChan invokes the sddp.DescribeTemplateAllRules API asynchronously
func (client *Client) DescribeTemplateAllRulesWithChan(request *DescribeTemplateAllRulesRequest) (<-chan *DescribeTemplateAllRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeTemplateAllRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTemplateAllRules(request)
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

// DescribeTemplateAllRulesWithCallback invokes the sddp.DescribeTemplateAllRules API asynchronously
func (client *Client) DescribeTemplateAllRulesWithCallback(request *DescribeTemplateAllRulesRequest, callback func(response *DescribeTemplateAllRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTemplateAllRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeTemplateAllRules(request)
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

// DescribeTemplateAllRulesRequest is the request struct for api DescribeTemplateAllRules
type DescribeTemplateAllRulesRequest struct {
	*requests.RpcRequest
	TemplateId requests.Integer `position:"Query" name:"TemplateId"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
	Lang       string           `position:"Query" name:"Lang"`
}

// DescribeTemplateAllRulesResponse is the response struct for api DescribeTemplateAllRules
type DescribeTemplateAllRulesResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	RuleList  []RuleInDescribeTemplateAllRules `json:"RuleList" xml:"RuleList"`
}

// CreateDescribeTemplateAllRulesRequest creates a request to invoke DescribeTemplateAllRules API
func CreateDescribeTemplateAllRulesRequest() (request *DescribeTemplateAllRulesRequest) {
	request = &DescribeTemplateAllRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeTemplateAllRules", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTemplateAllRulesResponse creates a response to parse from DescribeTemplateAllRules response
func CreateDescribeTemplateAllRulesResponse() (response *DescribeTemplateAllRulesResponse) {
	response = &DescribeTemplateAllRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
