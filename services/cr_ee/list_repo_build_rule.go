package cr_ee

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

// ListRepoBuildRule invokes the cr.ListRepoBuildRule API synchronously
// api document: https://help.aliyun.com/api/cr/listrepobuildrule.html
func (client *Client) ListRepoBuildRule(request *ListRepoBuildRuleRequest) (response *ListRepoBuildRuleResponse, err error) {
	response = CreateListRepoBuildRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepoBuildRuleWithChan invokes the cr.ListRepoBuildRule API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepobuildrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoBuildRuleWithChan(request *ListRepoBuildRuleRequest) (<-chan *ListRepoBuildRuleResponse, <-chan error) {
	responseChan := make(chan *ListRepoBuildRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepoBuildRule(request)
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

// ListRepoBuildRuleWithCallback invokes the cr.ListRepoBuildRule API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepobuildrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoBuildRuleWithCallback(request *ListRepoBuildRuleRequest, callback func(response *ListRepoBuildRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepoBuildRuleResponse
		var err error
		defer close(result)
		response, err = client.ListRepoBuildRule(request)
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

// ListRepoBuildRuleRequest is the request struct for api ListRepoBuildRule
type ListRepoBuildRuleRequest struct {
	*requests.RpcRequest
	RepoId     string           `position:"Query" name:"RepoId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageNo     requests.Integer `position:"Query" name:"PageNo"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListRepoBuildRuleResponse is the response struct for api ListRepoBuildRule
type ListRepoBuildRuleResponse struct {
	*responses.BaseResponse
	ListRepoBuildRuleIsSuccess bool             `json:"IsSuccess" xml:"IsSuccess"`
	Code                       string           `json:"Code" xml:"Code"`
	RequestId                  string           `json:"RequestId" xml:"RequestId"`
	PageNo                     int              `json:"PageNo" xml:"PageNo"`
	PageSize                   int              `json:"PageSize" xml:"PageSize"`
	TotalCount                 string           `json:"TotalCount" xml:"TotalCount"`
	BuildRules                 []BuildRulesItem `json:"BuildRules" xml:"BuildRules"`
}

// CreateListRepoBuildRuleRequest creates a request to invoke ListRepoBuildRule API
func CreateListRepoBuildRuleRequest() (request *ListRepoBuildRuleRequest) {
	request = &ListRepoBuildRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "ListRepoBuildRule", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRepoBuildRuleResponse creates a response to parse from ListRepoBuildRule response
func CreateListRepoBuildRuleResponse() (response *ListRepoBuildRuleResponse) {
	response = &ListRepoBuildRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
