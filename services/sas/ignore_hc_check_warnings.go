package sas

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

// IgnoreHcCheckWarnings invokes the sas.IgnoreHcCheckWarnings API synchronously
func (client *Client) IgnoreHcCheckWarnings(request *IgnoreHcCheckWarningsRequest) (response *IgnoreHcCheckWarningsResponse, err error) {
	response = CreateIgnoreHcCheckWarningsResponse()
	err = client.DoAction(request, response)
	return
}

// IgnoreHcCheckWarningsWithChan invokes the sas.IgnoreHcCheckWarnings API asynchronously
func (client *Client) IgnoreHcCheckWarningsWithChan(request *IgnoreHcCheckWarningsRequest) (<-chan *IgnoreHcCheckWarningsResponse, <-chan error) {
	responseChan := make(chan *IgnoreHcCheckWarningsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.IgnoreHcCheckWarnings(request)
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

// IgnoreHcCheckWarningsWithCallback invokes the sas.IgnoreHcCheckWarnings API asynchronously
func (client *Client) IgnoreHcCheckWarningsWithCallback(request *IgnoreHcCheckWarningsRequest, callback func(response *IgnoreHcCheckWarningsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *IgnoreHcCheckWarningsResponse
		var err error
		defer close(result)
		response, err = client.IgnoreHcCheckWarnings(request)
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

// IgnoreHcCheckWarningsRequest is the request struct for api IgnoreHcCheckWarnings
type IgnoreHcCheckWarningsRequest struct {
	*requests.RpcRequest
	Reason          string           `position:"Query" name:"Reason"`
	CheckIds        string           `position:"Query" name:"CheckIds"`
	RiskId          string           `position:"Query" name:"RiskId"`
	Type            requests.Integer `position:"Query" name:"Type"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	CheckWarningIds string           `position:"Query" name:"CheckWarningIds"`
}

// IgnoreHcCheckWarningsResponse is the response struct for api IgnoreHcCheckWarnings
type IgnoreHcCheckWarningsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateIgnoreHcCheckWarningsRequest creates a request to invoke IgnoreHcCheckWarnings API
func CreateIgnoreHcCheckWarningsRequest() (request *IgnoreHcCheckWarningsRequest) {
	request = &IgnoreHcCheckWarningsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "IgnoreHcCheckWarnings", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateIgnoreHcCheckWarningsResponse creates a response to parse from IgnoreHcCheckWarnings response
func CreateIgnoreHcCheckWarningsResponse() (response *IgnoreHcCheckWarningsResponse) {
	response = &IgnoreHcCheckWarningsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
