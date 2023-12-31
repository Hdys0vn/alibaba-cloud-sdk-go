package aegis

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

// CreateOrUpdateRule invokes the aegis.CreateOrUpdateRule API synchronously
// api document: https://help.aliyun.com/api/aegis/createorupdaterule.html
func (client *Client) CreateOrUpdateRule(request *CreateOrUpdateRuleRequest) (response *CreateOrUpdateRuleResponse, err error) {
	response = CreateCreateOrUpdateRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateRuleWithChan invokes the aegis.CreateOrUpdateRule API asynchronously
// api document: https://help.aliyun.com/api/aegis/createorupdaterule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrUpdateRuleWithChan(request *CreateOrUpdateRuleRequest) (<-chan *CreateOrUpdateRuleResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateRule(request)
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

// CreateOrUpdateRuleWithCallback invokes the aegis.CreateOrUpdateRule API asynchronously
// api document: https://help.aliyun.com/api/aegis/createorupdaterule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrUpdateRuleWithCallback(request *CreateOrUpdateRuleRequest, callback func(response *CreateOrUpdateRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateRule(request)
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

// CreateOrUpdateRuleRequest is the request struct for api CreateOrUpdateRule
type CreateOrUpdateRuleRequest struct {
	*requests.RpcRequest
	WarnLevel       string           `position:"Query" name:"WarnLevel"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	StatisticsRules string           `position:"Query" name:"StatisticsRules"`
	DataSourceId    requests.Integer `position:"Query" name:"DataSourceId"`
	Description     string           `position:"Query" name:"Description"`
	RuleName        string           `position:"Query" name:"RuleName"`
	Id              requests.Integer `position:"Query" name:"Id"`
	Lang            string           `position:"Query" name:"Lang"`
	Expressions     string           `position:"Query" name:"Expressions"`
	Actions         string           `position:"Query" name:"Actions"`
	RuleGroupIds    string           `position:"Query" name:"RuleGroupIds"`
}

// CreateOrUpdateRuleResponse is the response struct for api CreateOrUpdateRule
type CreateOrUpdateRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateOrUpdateRuleRequest creates a request to invoke CreateOrUpdateRule API
func CreateCreateOrUpdateRuleRequest() (request *CreateOrUpdateRuleRequest) {
	request = &CreateOrUpdateRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "CreateOrUpdateRule", "vipaegis", "openAPI")
	return
}

// CreateCreateOrUpdateRuleResponse creates a response to parse from CreateOrUpdateRule response
func CreateCreateOrUpdateRuleResponse() (response *CreateOrUpdateRuleResponse) {
	response = &CreateOrUpdateRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
