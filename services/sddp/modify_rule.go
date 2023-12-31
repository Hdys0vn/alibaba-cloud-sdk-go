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

// ModifyRule invokes the sddp.ModifyRule API synchronously
func (client *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
	response = CreateModifyRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRuleWithChan invokes the sddp.ModifyRule API asynchronously
func (client *Client) ModifyRuleWithChan(request *ModifyRuleRequest) (<-chan *ModifyRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRule(request)
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

// ModifyRuleWithCallback invokes the sddp.ModifyRule API asynchronously
func (client *Client) ModifyRuleWithCallback(request *ModifyRuleRequest, callback func(response *ModifyRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyRule(request)
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

// ModifyRuleRequest is the request struct for api ModifyRule
type ModifyRuleRequest struct {
	*requests.RpcRequest
	WarnLevel       requests.Integer `position:"Query" name:"WarnLevel"`
	ProductCode     string           `position:"Query" name:"ProductCode"`
	ProductId       requests.Integer `position:"Query" name:"ProductId"`
	Description     string           `position:"Query" name:"Description"`
	RiskLevelId     requests.Integer `position:"Query" name:"RiskLevelId"`
	Content         string           `position:"Query" name:"Content"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	MatchType       requests.Integer `position:"Query" name:"MatchType"`
	Id              requests.Integer `position:"Query" name:"Id"`
	Lang            string           `position:"Query" name:"Lang"`
	SupportForm     requests.Integer `position:"Query" name:"SupportForm"`
	FeatureType     requests.Integer `position:"Query" name:"FeatureType"`
	RuleType        requests.Integer `position:"Query" name:"RuleType"`
	StatExpress     string           `position:"Query" name:"StatExpress"`
	ContentCategory requests.Integer `position:"Query" name:"ContentCategory"`
	CustomType      requests.Integer `position:"Query" name:"CustomType"`
	Target          string           `position:"Query" name:"Target"`
	Name            string           `position:"Query" name:"Name"`
	Category        requests.Integer `position:"Query" name:"Category"`
}

// ModifyRuleResponse is the response struct for api ModifyRule
type ModifyRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyRuleRequest creates a request to invoke ModifyRule API
func CreateModifyRuleRequest() (request *ModifyRuleRequest) {
	request = &ModifyRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "ModifyRule", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyRuleResponse creates a response to parse from ModifyRule response
func CreateModifyRuleResponse() (response *ModifyRuleResponse) {
	response = &ModifyRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
