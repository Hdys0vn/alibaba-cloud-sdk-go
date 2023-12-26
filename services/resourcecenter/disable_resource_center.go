package resourcecenter

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

// DisableResourceCenter invokes the resourcecenter.DisableResourceCenter API synchronously
func (client *Client) DisableResourceCenter(request *DisableResourceCenterRequest) (response *DisableResourceCenterResponse, err error) {
	response = CreateDisableResourceCenterResponse()
	err = client.DoAction(request, response)
	return
}

// DisableResourceCenterWithChan invokes the resourcecenter.DisableResourceCenter API asynchronously
func (client *Client) DisableResourceCenterWithChan(request *DisableResourceCenterRequest) (<-chan *DisableResourceCenterResponse, <-chan error) {
	responseChan := make(chan *DisableResourceCenterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableResourceCenter(request)
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

// DisableResourceCenterWithCallback invokes the resourcecenter.DisableResourceCenter API asynchronously
func (client *Client) DisableResourceCenterWithCallback(request *DisableResourceCenterRequest, callback func(response *DisableResourceCenterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableResourceCenterResponse
		var err error
		defer close(result)
		response, err = client.DisableResourceCenter(request)
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

// DisableResourceCenterRequest is the request struct for api DisableResourceCenter
type DisableResourceCenterRequest struct {
	*requests.RpcRequest
}

// DisableResourceCenterResponse is the response struct for api DisableResourceCenter
type DisableResourceCenterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableResourceCenterRequest creates a request to invoke DisableResourceCenter API
func CreateDisableResourceCenterRequest() (request *DisableResourceCenterRequest) {
	request = &DisableResourceCenterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "DisableResourceCenter", "", "")
	request.Method = requests.POST
	return
}

// CreateDisableResourceCenterResponse creates a response to parse from DisableResourceCenter response
func CreateDisableResourceCenterResponse() (response *DisableResourceCenterResponse) {
	response = &DisableResourceCenterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}