package waf_openapi

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

// DescribeProtectionModuleRules invokes the waf_openapi.DescribeProtectionModuleRules API synchronously
func (client *Client) DescribeProtectionModuleRules(request *DescribeProtectionModuleRulesRequest) (response *DescribeProtectionModuleRulesResponse, err error) {
	response = CreateDescribeProtectionModuleRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProtectionModuleRulesWithChan invokes the waf_openapi.DescribeProtectionModuleRules API asynchronously
func (client *Client) DescribeProtectionModuleRulesWithChan(request *DescribeProtectionModuleRulesRequest) (<-chan *DescribeProtectionModuleRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeProtectionModuleRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProtectionModuleRules(request)
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

// DescribeProtectionModuleRulesWithCallback invokes the waf_openapi.DescribeProtectionModuleRules API asynchronously
func (client *Client) DescribeProtectionModuleRulesWithCallback(request *DescribeProtectionModuleRulesRequest, callback func(response *DescribeProtectionModuleRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProtectionModuleRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeProtectionModuleRules(request)
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

// DescribeProtectionModuleRulesRequest is the request struct for api DescribeProtectionModuleRules
type DescribeProtectionModuleRulesRequest struct {
	*requests.RpcRequest
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
	DefenseType     string           `position:"Query" name:"DefenseType"`
	Query           string           `position:"Query" name:"Query"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	Domain          string           `position:"Query" name:"Domain"`
}

// DescribeProtectionModuleRulesResponse is the response struct for api DescribeProtectionModuleRules
type DescribeProtectionModuleRulesResponse struct {
	*responses.BaseResponse
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Rules      []Rule `json:"Rules" xml:"Rules"`
}

// CreateDescribeProtectionModuleRulesRequest creates a request to invoke DescribeProtectionModuleRules API
func CreateDescribeProtectionModuleRulesRequest() (request *DescribeProtectionModuleRulesRequest) {
	request = &DescribeProtectionModuleRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeProtectionModuleRules", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProtectionModuleRulesResponse creates a response to parse from DescribeProtectionModuleRules response
func CreateDescribeProtectionModuleRulesResponse() (response *DescribeProtectionModuleRulesResponse) {
	response = &DescribeProtectionModuleRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
