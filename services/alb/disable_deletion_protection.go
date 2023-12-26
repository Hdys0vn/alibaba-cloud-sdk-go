package alb

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

// DisableDeletionProtection invokes the alb.DisableDeletionProtection API synchronously
func (client *Client) DisableDeletionProtection(request *DisableDeletionProtectionRequest) (response *DisableDeletionProtectionResponse, err error) {
	response = CreateDisableDeletionProtectionResponse()
	err = client.DoAction(request, response)
	return
}

// DisableDeletionProtectionWithChan invokes the alb.DisableDeletionProtection API asynchronously
func (client *Client) DisableDeletionProtectionWithChan(request *DisableDeletionProtectionRequest) (<-chan *DisableDeletionProtectionResponse, <-chan error) {
	responseChan := make(chan *DisableDeletionProtectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableDeletionProtection(request)
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

// DisableDeletionProtectionWithCallback invokes the alb.DisableDeletionProtection API asynchronously
func (client *Client) DisableDeletionProtectionWithCallback(request *DisableDeletionProtectionRequest, callback func(response *DisableDeletionProtectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableDeletionProtectionResponse
		var err error
		defer close(result)
		response, err = client.DisableDeletionProtection(request)
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

// DisableDeletionProtectionRequest is the request struct for api DisableDeletionProtection
type DisableDeletionProtectionRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Query" name:"ClientToken"`
	ResourceId  string           `position:"Query" name:"ResourceId"`
	DryRun      requests.Boolean `position:"Query" name:"DryRun"`
}

// DisableDeletionProtectionResponse is the response struct for api DisableDeletionProtection
type DisableDeletionProtectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableDeletionProtectionRequest creates a request to invoke DisableDeletionProtection API
func CreateDisableDeletionProtectionRequest() (request *DisableDeletionProtectionRequest) {
	request = &DisableDeletionProtectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "DisableDeletionProtection", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableDeletionProtectionResponse creates a response to parse from DisableDeletionProtection response
func CreateDisableDeletionProtectionResponse() (response *DisableDeletionProtectionResponse) {
	response = &DisableDeletionProtectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
