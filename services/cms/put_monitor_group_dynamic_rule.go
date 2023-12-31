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

// PutMonitorGroupDynamicRule invokes the cms.PutMonitorGroupDynamicRule API synchronously
func (client *Client) PutMonitorGroupDynamicRule(request *PutMonitorGroupDynamicRuleRequest) (response *PutMonitorGroupDynamicRuleResponse, err error) {
	response = CreatePutMonitorGroupDynamicRuleResponse()
	err = client.DoAction(request, response)
	return
}

// PutMonitorGroupDynamicRuleWithChan invokes the cms.PutMonitorGroupDynamicRule API asynchronously
func (client *Client) PutMonitorGroupDynamicRuleWithChan(request *PutMonitorGroupDynamicRuleRequest) (<-chan *PutMonitorGroupDynamicRuleResponse, <-chan error) {
	responseChan := make(chan *PutMonitorGroupDynamicRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutMonitorGroupDynamicRule(request)
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

// PutMonitorGroupDynamicRuleWithCallback invokes the cms.PutMonitorGroupDynamicRule API asynchronously
func (client *Client) PutMonitorGroupDynamicRuleWithCallback(request *PutMonitorGroupDynamicRuleRequest, callback func(response *PutMonitorGroupDynamicRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutMonitorGroupDynamicRuleResponse
		var err error
		defer close(result)
		response, err = client.PutMonitorGroupDynamicRule(request)
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

// PutMonitorGroupDynamicRuleRequest is the request struct for api PutMonitorGroupDynamicRule
type PutMonitorGroupDynamicRuleRequest struct {
	*requests.RpcRequest
	GroupRules *[]PutMonitorGroupDynamicRuleGroupRules `position:"Query" name:"GroupRules"  type:"Repeated"`
	GroupId    requests.Integer                        `position:"Query" name:"GroupId"`
	IsAsync    requests.Boolean                        `position:"Query" name:"IsAsync"`
}

// PutMonitorGroupDynamicRuleGroupRules is a repeated param struct in PutMonitorGroupDynamicRuleRequest
type PutMonitorGroupDynamicRuleGroupRules struct {
	FilterRelation string                                         `name:"FilterRelation"`
	Filters        *[]PutMonitorGroupDynamicRuleGroupRulesFilters `name:"Filters" type:"Repeated"`
	Category       string                                         `name:"Category"`
}

// PutMonitorGroupDynamicRuleGroupRulesFilters is a repeated param struct in PutMonitorGroupDynamicRuleRequest
type PutMonitorGroupDynamicRuleGroupRulesFilters struct {
	Function string `name:"Function"`
	Name     string `name:"Name"`
	Value    string `name:"Value"`
}

// PutMonitorGroupDynamicRuleResponse is the response struct for api PutMonitorGroupDynamicRule
type PutMonitorGroupDynamicRuleResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreatePutMonitorGroupDynamicRuleRequest creates a request to invoke PutMonitorGroupDynamicRule API
func CreatePutMonitorGroupDynamicRuleRequest() (request *PutMonitorGroupDynamicRuleRequest) {
	request = &PutMonitorGroupDynamicRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "PutMonitorGroupDynamicRule", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutMonitorGroupDynamicRuleResponse creates a response to parse from PutMonitorGroupDynamicRule response
func CreatePutMonitorGroupDynamicRuleResponse() (response *PutMonitorGroupDynamicRuleResponse) {
	response = &PutMonitorGroupDynamicRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
