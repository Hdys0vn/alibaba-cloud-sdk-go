package ahas_openapi

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

// DeleteDegradeRule invokes the ahas_openapi.DeleteDegradeRule API synchronously
func (client *Client) DeleteDegradeRule(request *DeleteDegradeRuleRequest) (response *DeleteDegradeRuleResponse, err error) {
	response = CreateDeleteDegradeRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDegradeRuleWithChan invokes the ahas_openapi.DeleteDegradeRule API asynchronously
func (client *Client) DeleteDegradeRuleWithChan(request *DeleteDegradeRuleRequest) (<-chan *DeleteDegradeRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteDegradeRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDegradeRule(request)
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

// DeleteDegradeRuleWithCallback invokes the ahas_openapi.DeleteDegradeRule API asynchronously
func (client *Client) DeleteDegradeRuleWithCallback(request *DeleteDegradeRuleRequest, callback func(response *DeleteDegradeRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDegradeRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteDegradeRule(request)
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

// DeleteDegradeRuleRequest is the request struct for api DeleteDegradeRule
type DeleteDegradeRuleRequest struct {
	*requests.RpcRequest
	AhasRegionId string           `position:"Query" name:"AhasRegionId"`
	RuleId       requests.Integer `position:"Query" name:"RuleId"`
}

// DeleteDegradeRuleResponse is the response struct for api DeleteDegradeRule
type DeleteDegradeRuleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteDegradeRuleRequest creates a request to invoke DeleteDegradeRule API
func CreateDeleteDegradeRuleRequest() (request *DeleteDegradeRuleRequest) {
	request = &DeleteDegradeRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "DeleteDegradeRule", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDegradeRuleResponse creates a response to parse from DeleteDegradeRule response
func CreateDeleteDegradeRuleResponse() (response *DeleteDegradeRuleResponse) {
	response = &DeleteDegradeRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
