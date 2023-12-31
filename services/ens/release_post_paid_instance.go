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

// ReleasePostPaidInstance invokes the ens.ReleasePostPaidInstance API synchronously
func (client *Client) ReleasePostPaidInstance(request *ReleasePostPaidInstanceRequest) (response *ReleasePostPaidInstanceResponse, err error) {
	response = CreateReleasePostPaidInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ReleasePostPaidInstanceWithChan invokes the ens.ReleasePostPaidInstance API asynchronously
func (client *Client) ReleasePostPaidInstanceWithChan(request *ReleasePostPaidInstanceRequest) (<-chan *ReleasePostPaidInstanceResponse, <-chan error) {
	responseChan := make(chan *ReleasePostPaidInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleasePostPaidInstance(request)
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

// ReleasePostPaidInstanceWithCallback invokes the ens.ReleasePostPaidInstance API asynchronously
func (client *Client) ReleasePostPaidInstanceWithCallback(request *ReleasePostPaidInstanceRequest, callback func(response *ReleasePostPaidInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleasePostPaidInstanceResponse
		var err error
		defer close(result)
		response, err = client.ReleasePostPaidInstance(request)
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

// ReleasePostPaidInstanceRequest is the request struct for api ReleasePostPaidInstance
type ReleasePostPaidInstanceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ReleasePostPaidInstanceResponse is the response struct for api ReleasePostPaidInstance
type ReleasePostPaidInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleasePostPaidInstanceRequest creates a request to invoke ReleasePostPaidInstance API
func CreateReleasePostPaidInstanceRequest() (request *ReleasePostPaidInstanceRequest) {
	request = &ReleasePostPaidInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ReleasePostPaidInstance", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleasePostPaidInstanceResponse creates a response to parse from ReleasePostPaidInstance response
func CreateReleasePostPaidInstanceResponse() (response *ReleasePostPaidInstanceResponse) {
	response = &ReleasePostPaidInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
