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

// DeleteExporterRule invokes the cms.DeleteExporterRule API synchronously
func (client *Client) DeleteExporterRule(request *DeleteExporterRuleRequest) (response *DeleteExporterRuleResponse, err error) {
	response = CreateDeleteExporterRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteExporterRuleWithChan invokes the cms.DeleteExporterRule API asynchronously
func (client *Client) DeleteExporterRuleWithChan(request *DeleteExporterRuleRequest) (<-chan *DeleteExporterRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteExporterRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteExporterRule(request)
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

// DeleteExporterRuleWithCallback invokes the cms.DeleteExporterRule API asynchronously
func (client *Client) DeleteExporterRuleWithCallback(request *DeleteExporterRuleRequest, callback func(response *DeleteExporterRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteExporterRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteExporterRule(request)
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

// DeleteExporterRuleRequest is the request struct for api DeleteExporterRule
type DeleteExporterRuleRequest struct {
	*requests.RpcRequest
	RuleName string `position:"Query" name:"RuleName"`
}

// DeleteExporterRuleResponse is the response struct for api DeleteExporterRule
type DeleteExporterRuleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteExporterRuleRequest creates a request to invoke DeleteExporterRule API
func CreateDeleteExporterRuleRequest() (request *DeleteExporterRuleRequest) {
	request = &DeleteExporterRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteExporterRule", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteExporterRuleResponse creates a response to parse from DeleteExporterRule response
func CreateDeleteExporterRuleResponse() (response *DeleteExporterRuleResponse) {
	response = &DeleteExporterRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}