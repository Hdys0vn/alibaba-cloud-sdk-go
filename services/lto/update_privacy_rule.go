package lto

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

// UpdatePrivacyRule invokes the lto.UpdatePrivacyRule API synchronously
func (client *Client) UpdatePrivacyRule(request *UpdatePrivacyRuleRequest) (response *UpdatePrivacyRuleResponse, err error) {
	response = CreateUpdatePrivacyRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePrivacyRuleWithChan invokes the lto.UpdatePrivacyRule API asynchronously
func (client *Client) UpdatePrivacyRuleWithChan(request *UpdatePrivacyRuleRequest) (<-chan *UpdatePrivacyRuleResponse, <-chan error) {
	responseChan := make(chan *UpdatePrivacyRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePrivacyRule(request)
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

// UpdatePrivacyRuleWithCallback invokes the lto.UpdatePrivacyRule API asynchronously
func (client *Client) UpdatePrivacyRuleWithCallback(request *UpdatePrivacyRuleRequest, callback func(response *UpdatePrivacyRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePrivacyRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdatePrivacyRule(request)
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

// UpdatePrivacyRuleRequest is the request struct for api UpdatePrivacyRule
type UpdatePrivacyRuleRequest struct {
	*requests.RpcRequest
	PrivacyRuleId string `position:"Query" name:"PrivacyRuleId"`
	Remark        string `position:"Query" name:"Remark"`
	Name          string `position:"Query" name:"Name"`
	AlgImpl       string `position:"Query" name:"AlgImpl"`
	AlgType       string `position:"Query" name:"AlgType"`
}

// UpdatePrivacyRuleResponse is the response struct for api UpdatePrivacyRule
type UpdatePrivacyRuleResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateUpdatePrivacyRuleRequest creates a request to invoke UpdatePrivacyRule API
func CreateUpdatePrivacyRuleRequest() (request *UpdatePrivacyRuleRequest) {
	request = &UpdatePrivacyRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "UpdatePrivacyRule", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdatePrivacyRuleResponse creates a response to parse from UpdatePrivacyRule response
func CreateUpdatePrivacyRuleResponse() (response *UpdatePrivacyRuleResponse) {
	response = &UpdatePrivacyRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
