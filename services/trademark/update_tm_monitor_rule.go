package trademark

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

// UpdateTmMonitorRule invokes the trademark.UpdateTmMonitorRule API synchronously
// api document: https://help.aliyun.com/api/trademark/updatetmmonitorrule.html
func (client *Client) UpdateTmMonitorRule(request *UpdateTmMonitorRuleRequest) (response *UpdateTmMonitorRuleResponse, err error) {
	response = CreateUpdateTmMonitorRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTmMonitorRuleWithChan invokes the trademark.UpdateTmMonitorRule API asynchronously
// api document: https://help.aliyun.com/api/trademark/updatetmmonitorrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateTmMonitorRuleWithChan(request *UpdateTmMonitorRuleRequest) (<-chan *UpdateTmMonitorRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateTmMonitorRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTmMonitorRule(request)
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

// UpdateTmMonitorRuleWithCallback invokes the trademark.UpdateTmMonitorRule API asynchronously
// api document: https://help.aliyun.com/api/trademark/updatetmmonitorrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateTmMonitorRuleWithCallback(request *UpdateTmMonitorRuleRequest, callback func(response *UpdateTmMonitorRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTmMonitorRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateTmMonitorRule(request)
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

// UpdateTmMonitorRuleRequest is the request struct for api UpdateTmMonitorRule
type UpdateTmMonitorRuleRequest struct {
	*requests.RpcRequest
	NotifyStatus *[]string        `position:"Query" name:"NotifyStatus"  type:"Repeated"`
	RuleName     string           `position:"Query" name:"RuleName"`
	Id           requests.Integer `position:"Query" name:"Id"`
}

// UpdateTmMonitorRuleResponse is the response struct for api UpdateTmMonitorRule
type UpdateTmMonitorRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateUpdateTmMonitorRuleRequest creates a request to invoke UpdateTmMonitorRule API
func CreateUpdateTmMonitorRuleRequest() (request *UpdateTmMonitorRuleRequest) {
	request = &UpdateTmMonitorRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "UpdateTmMonitorRule", "trademark", "openAPI")
	return
}

// CreateUpdateTmMonitorRuleResponse creates a response to parse from UpdateTmMonitorRule response
func CreateUpdateTmMonitorRuleResponse() (response *UpdateTmMonitorRuleResponse) {
	response = &UpdateTmMonitorRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
