package retailcloud

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

// SyncPodInfo invokes the retailcloud.SyncPodInfo API synchronously
func (client *Client) SyncPodInfo(request *SyncPodInfoRequest) (response *SyncPodInfoResponse, err error) {
	response = CreateSyncPodInfoResponse()
	err = client.DoAction(request, response)
	return
}

// SyncPodInfoWithChan invokes the retailcloud.SyncPodInfo API asynchronously
func (client *Client) SyncPodInfoWithChan(request *SyncPodInfoRequest) (<-chan *SyncPodInfoResponse, <-chan error) {
	responseChan := make(chan *SyncPodInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncPodInfo(request)
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

// SyncPodInfoWithCallback invokes the retailcloud.SyncPodInfo API asynchronously
func (client *Client) SyncPodInfoWithCallback(request *SyncPodInfoRequest, callback func(response *SyncPodInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncPodInfoResponse
		var err error
		defer close(result)
		response, err = client.SyncPodInfo(request)
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

// SyncPodInfoRequest is the request struct for api SyncPodInfo
type SyncPodInfoRequest struct {
	*requests.RpcRequest
	Reason      string           `position:"Query" name:"Reason"`
	RequestId   string           `position:"Query" name:"RequestId"`
	PodName     string           `position:"Query" name:"PodName"`
	SideCarType string           `position:"Query" name:"SideCarType"`
	TaskId      requests.Integer `position:"Query" name:"TaskId"`
	Status      requests.Boolean `position:"Query" name:"Status"`
}

// SyncPodInfoResponse is the response struct for api SyncPodInfo
type SyncPodInfoResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateSyncPodInfoRequest creates a request to invoke SyncPodInfo API
func CreateSyncPodInfoRequest() (request *SyncPodInfoRequest) {
	request = &SyncPodInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "SyncPodInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateSyncPodInfoResponse creates a response to parse from SyncPodInfo response
func CreateSyncPodInfoResponse() (response *SyncPodInfoResponse) {
	response = &SyncPodInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
