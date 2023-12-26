package bpstudio

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

// ValuateApplication invokes the bpstudio.ValuateApplication API synchronously
func (client *Client) ValuateApplication(request *ValuateApplicationRequest) (response *ValuateApplicationResponse, err error) {
	response = CreateValuateApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// ValuateApplicationWithChan invokes the bpstudio.ValuateApplication API asynchronously
func (client *Client) ValuateApplicationWithChan(request *ValuateApplicationRequest) (<-chan *ValuateApplicationResponse, <-chan error) {
	responseChan := make(chan *ValuateApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ValuateApplication(request)
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

// ValuateApplicationWithCallback invokes the bpstudio.ValuateApplication API asynchronously
func (client *Client) ValuateApplicationWithCallback(request *ValuateApplicationRequest, callback func(response *ValuateApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ValuateApplicationResponse
		var err error
		defer close(result)
		response, err = client.ValuateApplication(request)
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

// ValuateApplicationRequest is the request struct for api ValuateApplication
type ValuateApplicationRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Body" name:"ResourceGroupId"`
	ApplicationId   string `position:"Body" name:"ApplicationId"`
}

// ValuateApplicationResponse is the response struct for api ValuateApplication
type ValuateApplicationResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      int64  `json:"Data" xml:"Data"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateValuateApplicationRequest creates a request to invoke ValuateApplication API
func CreateValuateApplicationRequest() (request *ValuateApplicationRequest) {
	request = &ValuateApplicationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BPStudio", "2021-09-31", "ValuateApplication", "bpstudio", "openAPI")
	request.Method = requests.POST
	return
}

// CreateValuateApplicationResponse creates a response to parse from ValuateApplication response
func CreateValuateApplicationResponse() (response *ValuateApplicationResponse) {
	response = &ValuateApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}