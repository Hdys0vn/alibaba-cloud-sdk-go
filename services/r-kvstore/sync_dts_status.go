package r_kvstore

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

// SyncDtsStatus invokes the r_kvstore.SyncDtsStatus API synchronously
func (client *Client) SyncDtsStatus(request *SyncDtsStatusRequest) (response *SyncDtsStatusResponse, err error) {
	response = CreateSyncDtsStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SyncDtsStatusWithChan invokes the r_kvstore.SyncDtsStatus API asynchronously
func (client *Client) SyncDtsStatusWithChan(request *SyncDtsStatusRequest) (<-chan *SyncDtsStatusResponse, <-chan error) {
	responseChan := make(chan *SyncDtsStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncDtsStatus(request)
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

// SyncDtsStatusWithCallback invokes the r_kvstore.SyncDtsStatus API asynchronously
func (client *Client) SyncDtsStatusWithCallback(request *SyncDtsStatusRequest, callback func(response *SyncDtsStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncDtsStatusResponse
		var err error
		defer close(result)
		response, err = client.SyncDtsStatus(request)
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

// SyncDtsStatusRequest is the request struct for api SyncDtsStatus
type SyncDtsStatusRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Status               string           `position:"Query" name:"Status"`
}

// SyncDtsStatusResponse is the response struct for api SyncDtsStatus
type SyncDtsStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSyncDtsStatusRequest creates a request to invoke SyncDtsStatus API
func CreateSyncDtsStatusRequest() (request *SyncDtsStatusRequest) {
	request = &SyncDtsStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "SyncDtsStatus", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSyncDtsStatusResponse creates a response to parse from SyncDtsStatus response
func CreateSyncDtsStatusResponse() (response *SyncDtsStatusResponse) {
	response = &SyncDtsStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
