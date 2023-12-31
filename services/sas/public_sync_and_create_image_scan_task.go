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

// PublicSyncAndCreateImageScanTask invokes the sas.PublicSyncAndCreateImageScanTask API synchronously
func (client *Client) PublicSyncAndCreateImageScanTask(request *PublicSyncAndCreateImageScanTaskRequest) (response *PublicSyncAndCreateImageScanTaskResponse, err error) {
	response = CreatePublicSyncAndCreateImageScanTaskResponse()
	err = client.DoAction(request, response)
	return
}

// PublicSyncAndCreateImageScanTaskWithChan invokes the sas.PublicSyncAndCreateImageScanTask API asynchronously
func (client *Client) PublicSyncAndCreateImageScanTaskWithChan(request *PublicSyncAndCreateImageScanTaskRequest) (<-chan *PublicSyncAndCreateImageScanTaskResponse, <-chan error) {
	responseChan := make(chan *PublicSyncAndCreateImageScanTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublicSyncAndCreateImageScanTask(request)
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

// PublicSyncAndCreateImageScanTaskWithCallback invokes the sas.PublicSyncAndCreateImageScanTask API asynchronously
func (client *Client) PublicSyncAndCreateImageScanTaskWithCallback(request *PublicSyncAndCreateImageScanTaskRequest, callback func(response *PublicSyncAndCreateImageScanTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublicSyncAndCreateImageScanTaskResponse
		var err error
		defer close(result)
		response, err = client.PublicSyncAndCreateImageScanTask(request)
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

// PublicSyncAndCreateImageScanTaskRequest is the request struct for api PublicSyncAndCreateImageScanTask
type PublicSyncAndCreateImageScanTaskRequest struct {
	*requests.RpcRequest
	Images     string `position:"Query" name:"Images"`
	ImageItems string `position:"Query" name:"ImageItems"`
	SourceIp   string `position:"Query" name:"SourceIp"`
}

// PublicSyncAndCreateImageScanTaskResponse is the response struct for api PublicSyncAndCreateImageScanTask
type PublicSyncAndCreateImageScanTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreatePublicSyncAndCreateImageScanTaskRequest creates a request to invoke PublicSyncAndCreateImageScanTask API
func CreatePublicSyncAndCreateImageScanTaskRequest() (request *PublicSyncAndCreateImageScanTaskRequest) {
	request = &PublicSyncAndCreateImageScanTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "PublicSyncAndCreateImageScanTask", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePublicSyncAndCreateImageScanTaskResponse creates a response to parse from PublicSyncAndCreateImageScanTask response
func CreatePublicSyncAndCreateImageScanTaskResponse() (response *PublicSyncAndCreateImageScanTaskResponse) {
	response = &PublicSyncAndCreateImageScanTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
