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

// ModifyProtectionRuleCacheStatus invokes the waf_openapi.ModifyProtectionRuleCacheStatus API synchronously
func (client *Client) ModifyProtectionRuleCacheStatus(request *ModifyProtectionRuleCacheStatusRequest) (response *ModifyProtectionRuleCacheStatusResponse, err error) {
	response = CreateModifyProtectionRuleCacheStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyProtectionRuleCacheStatusWithChan invokes the waf_openapi.ModifyProtectionRuleCacheStatus API asynchronously
func (client *Client) ModifyProtectionRuleCacheStatusWithChan(request *ModifyProtectionRuleCacheStatusRequest) (<-chan *ModifyProtectionRuleCacheStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyProtectionRuleCacheStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyProtectionRuleCacheStatus(request)
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

// ModifyProtectionRuleCacheStatusWithCallback invokes the waf_openapi.ModifyProtectionRuleCacheStatus API asynchronously
func (client *Client) ModifyProtectionRuleCacheStatusWithCallback(request *ModifyProtectionRuleCacheStatusRequest, callback func(response *ModifyProtectionRuleCacheStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyProtectionRuleCacheStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyProtectionRuleCacheStatus(request)
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

// ModifyProtectionRuleCacheStatusRequest is the request struct for api ModifyProtectionRuleCacheStatus
type ModifyProtectionRuleCacheStatusRequest struct {
	*requests.RpcRequest
	DefenseType string           `position:"Query" name:"DefenseType"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Domain      string           `position:"Query" name:"Domain"`
	Lang        string           `position:"Query" name:"Lang"`
	RuleId      requests.Integer `position:"Query" name:"RuleId"`
}

// ModifyProtectionRuleCacheStatusResponse is the response struct for api ModifyProtectionRuleCacheStatus
type ModifyProtectionRuleCacheStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyProtectionRuleCacheStatusRequest creates a request to invoke ModifyProtectionRuleCacheStatus API
func CreateModifyProtectionRuleCacheStatusRequest() (request *ModifyProtectionRuleCacheStatusRequest) {
	request = &ModifyProtectionRuleCacheStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "ModifyProtectionRuleCacheStatus", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyProtectionRuleCacheStatusResponse creates a response to parse from ModifyProtectionRuleCacheStatus response
func CreateModifyProtectionRuleCacheStatusResponse() (response *ModifyProtectionRuleCacheStatusResponse) {
	response = &ModifyProtectionRuleCacheStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
