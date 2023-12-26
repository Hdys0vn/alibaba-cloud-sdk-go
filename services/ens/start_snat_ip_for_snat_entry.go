package ens

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

// StartSnatIpForSnatEntry invokes the ens.StartSnatIpForSnatEntry API synchronously
func (client *Client) StartSnatIpForSnatEntry(request *StartSnatIpForSnatEntryRequest) (response *StartSnatIpForSnatEntryResponse, err error) {
	response = CreateStartSnatIpForSnatEntryResponse()
	err = client.DoAction(request, response)
	return
}

// StartSnatIpForSnatEntryWithChan invokes the ens.StartSnatIpForSnatEntry API asynchronously
func (client *Client) StartSnatIpForSnatEntryWithChan(request *StartSnatIpForSnatEntryRequest) (<-chan *StartSnatIpForSnatEntryResponse, <-chan error) {
	responseChan := make(chan *StartSnatIpForSnatEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartSnatIpForSnatEntry(request)
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

// StartSnatIpForSnatEntryWithCallback invokes the ens.StartSnatIpForSnatEntry API asynchronously
func (client *Client) StartSnatIpForSnatEntryWithCallback(request *StartSnatIpForSnatEntryRequest, callback func(response *StartSnatIpForSnatEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartSnatIpForSnatEntryResponse
		var err error
		defer close(result)
		response, err = client.StartSnatIpForSnatEntry(request)
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

// StartSnatIpForSnatEntryRequest is the request struct for api StartSnatIpForSnatEntry
type StartSnatIpForSnatEntryRequest struct {
	*requests.RpcRequest
	SnatIp      string `position:"Query" name:"SnatIp"`
	SnatEntryId string `position:"Query" name:"SnatEntryId"`
}

// StartSnatIpForSnatEntryResponse is the response struct for api StartSnatIpForSnatEntry
type StartSnatIpForSnatEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartSnatIpForSnatEntryRequest creates a request to invoke StartSnatIpForSnatEntry API
func CreateStartSnatIpForSnatEntryRequest() (request *StartSnatIpForSnatEntryRequest) {
	request = &StartSnatIpForSnatEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "StartSnatIpForSnatEntry", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartSnatIpForSnatEntryResponse creates a response to parse from StartSnatIpForSnatEntry response
func CreateStartSnatIpForSnatEntryResponse() (response *StartSnatIpForSnatEntryResponse) {
	response = &StartSnatIpForSnatEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}