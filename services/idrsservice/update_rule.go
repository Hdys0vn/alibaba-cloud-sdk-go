package idrsservice

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

// UpdateRule invokes the idrsservice.UpdateRule API synchronously
func (client *Client) UpdateRule(request *UpdateRuleRequest) (response *UpdateRuleResponse, err error) {
	response = CreateUpdateRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRuleWithChan invokes the idrsservice.UpdateRule API asynchronously
func (client *Client) UpdateRuleWithChan(request *UpdateRuleRequest) (<-chan *UpdateRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRule(request)
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

// UpdateRuleWithCallback invokes the idrsservice.UpdateRule API asynchronously
func (client *Client) UpdateRuleWithCallback(request *UpdateRuleRequest, callback func(response *UpdateRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateRule(request)
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

// UpdateRuleRequest is the request struct for api UpdateRule
type UpdateRuleRequest struct {
	*requests.RpcRequest
	Content string `position:"Query" name:"Content"`
	Name    string `position:"Query" name:"Name"`
	Id      string `position:"Query" name:"Id"`
}

// UpdateRuleResponse is the response struct for api UpdateRule
type UpdateRuleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateRuleRequest creates a request to invoke UpdateRule API
func CreateUpdateRuleRequest() (request *UpdateRuleRequest) {
	request = &UpdateRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "UpdateRule", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateRuleResponse creates a response to parse from UpdateRule response
func CreateUpdateRuleResponse() (response *UpdateRuleResponse) {
	response = &UpdateRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
