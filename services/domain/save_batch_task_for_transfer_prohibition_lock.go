package domain

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

// SaveBatchTaskForTransferProhibitionLock invokes the domain.SaveBatchTaskForTransferProhibitionLock API synchronously
func (client *Client) SaveBatchTaskForTransferProhibitionLock(request *SaveBatchTaskForTransferProhibitionLockRequest) (response *SaveBatchTaskForTransferProhibitionLockResponse, err error) {
	response = CreateSaveBatchTaskForTransferProhibitionLockResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForTransferProhibitionLockWithChan invokes the domain.SaveBatchTaskForTransferProhibitionLock API asynchronously
func (client *Client) SaveBatchTaskForTransferProhibitionLockWithChan(request *SaveBatchTaskForTransferProhibitionLockRequest) (<-chan *SaveBatchTaskForTransferProhibitionLockResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForTransferProhibitionLockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForTransferProhibitionLock(request)
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

// SaveBatchTaskForTransferProhibitionLockWithCallback invokes the domain.SaveBatchTaskForTransferProhibitionLock API asynchronously
func (client *Client) SaveBatchTaskForTransferProhibitionLockWithCallback(request *SaveBatchTaskForTransferProhibitionLockRequest, callback func(response *SaveBatchTaskForTransferProhibitionLockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForTransferProhibitionLockResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForTransferProhibitionLock(request)
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

// SaveBatchTaskForTransferProhibitionLockRequest is the request struct for api SaveBatchTaskForTransferProhibitionLock
type SaveBatchTaskForTransferProhibitionLockRequest struct {
	*requests.RpcRequest
	DomainName   *[]string        `position:"Query" name:"DomainName"  type:"Repeated"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
	Status       requests.Boolean `position:"Query" name:"Status"`
}

// SaveBatchTaskForTransferProhibitionLockResponse is the response struct for api SaveBatchTaskForTransferProhibitionLock
type SaveBatchTaskForTransferProhibitionLockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForTransferProhibitionLockRequest creates a request to invoke SaveBatchTaskForTransferProhibitionLock API
func CreateSaveBatchTaskForTransferProhibitionLockRequest() (request *SaveBatchTaskForTransferProhibitionLockRequest) {
	request = &SaveBatchTaskForTransferProhibitionLockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForTransferProhibitionLock", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveBatchTaskForTransferProhibitionLockResponse creates a response to parse from SaveBatchTaskForTransferProhibitionLock response
func CreateSaveBatchTaskForTransferProhibitionLockResponse() (response *SaveBatchTaskForTransferProhibitionLockResponse) {
	response = &SaveBatchTaskForTransferProhibitionLockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
