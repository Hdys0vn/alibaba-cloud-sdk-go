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

// ResizeDisk invokes the ens.ResizeDisk API synchronously
func (client *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
	response = CreateResizeDiskResponse()
	err = client.DoAction(request, response)
	return
}

// ResizeDiskWithChan invokes the ens.ResizeDisk API asynchronously
func (client *Client) ResizeDiskWithChan(request *ResizeDiskRequest) (<-chan *ResizeDiskResponse, <-chan error) {
	responseChan := make(chan *ResizeDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResizeDisk(request)
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

// ResizeDiskWithCallback invokes the ens.ResizeDisk API asynchronously
func (client *Client) ResizeDiskWithCallback(request *ResizeDiskRequest, callback func(response *ResizeDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResizeDiskResponse
		var err error
		defer close(result)
		response, err = client.ResizeDisk(request)
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

// ResizeDiskRequest is the request struct for api ResizeDisk
type ResizeDiskRequest struct {
	*requests.RpcRequest
	DiskId  string `position:"Query" name:"DiskId"`
	NewSize string `position:"Query" name:"NewSize"`
}

// ResizeDiskResponse is the response struct for api ResizeDisk
type ResizeDiskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateResizeDiskRequest creates a request to invoke ResizeDisk API
func CreateResizeDiskRequest() (request *ResizeDiskRequest) {
	request = &ResizeDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ResizeDisk", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResizeDiskResponse creates a response to parse from ResizeDisk response
func CreateResizeDiskResponse() (response *ResizeDiskResponse) {
	response = &ResizeDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
