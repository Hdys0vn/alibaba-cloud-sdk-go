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

// ListPersistentVolume invokes the retailcloud.ListPersistentVolume API synchronously
func (client *Client) ListPersistentVolume(request *ListPersistentVolumeRequest) (response *ListPersistentVolumeResponse, err error) {
	response = CreateListPersistentVolumeResponse()
	err = client.DoAction(request, response)
	return
}

// ListPersistentVolumeWithChan invokes the retailcloud.ListPersistentVolume API asynchronously
func (client *Client) ListPersistentVolumeWithChan(request *ListPersistentVolumeRequest) (<-chan *ListPersistentVolumeResponse, <-chan error) {
	responseChan := make(chan *ListPersistentVolumeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPersistentVolume(request)
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

// ListPersistentVolumeWithCallback invokes the retailcloud.ListPersistentVolume API asynchronously
func (client *Client) ListPersistentVolumeWithCallback(request *ListPersistentVolumeRequest, callback func(response *ListPersistentVolumeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPersistentVolumeResponse
		var err error
		defer close(result)
		response, err = client.ListPersistentVolume(request)
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

// ListPersistentVolumeRequest is the request struct for api ListPersistentVolume
type ListPersistentVolumeRequest struct {
	*requests.RpcRequest
	PageSize          requests.Integer `position:"Body" name:"PageSize"`
	PageNumber        requests.Integer `position:"Body" name:"PageNumber"`
	ClusterInstanceId string           `position:"Body" name:"ClusterInstanceId"`
}

// ListPersistentVolumeResponse is the response struct for api ListPersistentVolume
type ListPersistentVolumeResponse struct {
	*responses.BaseResponse
	RequestId  string                   `json:"RequestId" xml:"RequestId"`
	Code       int                      `json:"Code" xml:"Code"`
	PageSize   int                      `json:"PageSize" xml:"PageSize"`
	PageNumber int                      `json:"PageNumber" xml:"PageNumber"`
	TotalCount int64                    `json:"TotalCount" xml:"TotalCount"`
	ErrMsg     string                   `json:"ErrMsg" xml:"ErrMsg"`
	Data       []PersistentVolumeDetail `json:"Data" xml:"Data"`
}

// CreateListPersistentVolumeRequest creates a request to invoke ListPersistentVolume API
func CreateListPersistentVolumeRequest() (request *ListPersistentVolumeRequest) {
	request = &ListPersistentVolumeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListPersistentVolume", "", "")
	request.Method = requests.POST
	return
}

// CreateListPersistentVolumeResponse creates a response to parse from ListPersistentVolume response
func CreateListPersistentVolumeResponse() (response *ListPersistentVolumeResponse) {
	response = &ListPersistentVolumeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
