package visionai_poc

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

// CheckMultiagent invokes the visionai_poc.CheckMultiagent API synchronously
// api document: https://help.aliyun.com/api/visionai-poc/checkmultiagent.html
func (client *Client) CheckMultiagent(request *CheckMultiagentRequest) (response *CheckMultiagentResponse, err error) {
	response = CreateCheckMultiagentResponse()
	err = client.DoAction(request, response)
	return
}

// CheckMultiagentWithChan invokes the visionai_poc.CheckMultiagent API asynchronously
// api document: https://help.aliyun.com/api/visionai-poc/checkmultiagent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckMultiagentWithChan(request *CheckMultiagentRequest) (<-chan *CheckMultiagentResponse, <-chan error) {
	responseChan := make(chan *CheckMultiagentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckMultiagent(request)
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

// CheckMultiagentWithCallback invokes the visionai_poc.CheckMultiagent API asynchronously
// api document: https://help.aliyun.com/api/visionai-poc/checkmultiagent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckMultiagentWithCallback(request *CheckMultiagentRequest, callback func(response *CheckMultiagentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckMultiagentResponse
		var err error
		defer close(result)
		response, err = client.CheckMultiagent(request)
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

// CheckMultiagentRequest is the request struct for api CheckMultiagent
type CheckMultiagentRequest struct {
	*requests.RpcRequest
	Method   string `position:"Body" name:"Method"`
	ImageUrl string `position:"Body" name:"ImageUrl"`
	Url      string `position:"Body" name:"Url"`
}

// CheckMultiagentResponse is the response struct for api CheckMultiagent
type CheckMultiagentResponse struct {
	*responses.BaseResponse
	Code     string   `json:"Code" xml:"Code"`
	Success  bool     `json:"Success" xml:"Success"`
	Message  string   `json:"Message" xml:"Message"`
	Response Response `json:"Response" xml:"Response"`
}

// CreateCheckMultiagentRequest creates a request to invoke CheckMultiagent API
func CreateCheckMultiagentRequest() (request *CheckMultiagentRequest) {
	request = &CheckMultiagentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("visionai-poc", "2020-04-08", "CheckMultiagent", "", "")
	return
}

// CreateCheckMultiagentResponse creates a response to parse from CheckMultiagent response
func CreateCheckMultiagentResponse() (response *CheckMultiagentResponse) {
	response = &CheckMultiagentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
