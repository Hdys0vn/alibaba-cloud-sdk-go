package vod

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

// ListLiveRecordVideo invokes the vod.ListLiveRecordVideo API synchronously
func (client *Client) ListLiveRecordVideo(request *ListLiveRecordVideoRequest) (response *ListLiveRecordVideoResponse, err error) {
	response = CreateListLiveRecordVideoResponse()
	err = client.DoAction(request, response)
	return
}

// ListLiveRecordVideoWithChan invokes the vod.ListLiveRecordVideo API asynchronously
func (client *Client) ListLiveRecordVideoWithChan(request *ListLiveRecordVideoRequest) (<-chan *ListLiveRecordVideoResponse, <-chan error) {
	responseChan := make(chan *ListLiveRecordVideoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLiveRecordVideo(request)
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

// ListLiveRecordVideoWithCallback invokes the vod.ListLiveRecordVideo API asynchronously
func (client *Client) ListLiveRecordVideoWithCallback(request *ListLiveRecordVideoRequest, callback func(response *ListLiveRecordVideoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLiveRecordVideoResponse
		var err error
		defer close(result)
		response, err = client.ListLiveRecordVideo(request)
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

// ListLiveRecordVideoRequest is the request struct for api ListLiveRecordVideo
type ListLiveRecordVideoRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	AppName    string           `position:"Query" name:"AppName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	StreamName string           `position:"Query" name:"StreamName"`
	QueryType  string           `position:"Query" name:"QueryType"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	PageNo     requests.Integer `position:"Query" name:"PageNo"`
	SortBy     string           `position:"Query" name:"SortBy"`
}

// ListLiveRecordVideoResponse is the response struct for api ListLiveRecordVideo
type ListLiveRecordVideoResponse struct {
	*responses.BaseResponse
	Total               int                 `json:"Total" xml:"Total"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	LiveRecordVideoList LiveRecordVideoList `json:"LiveRecordVideoList" xml:"LiveRecordVideoList"`
}

// CreateListLiveRecordVideoRequest creates a request to invoke ListLiveRecordVideo API
func CreateListLiveRecordVideoRequest() (request *ListLiveRecordVideoRequest) {
	request = &ListLiveRecordVideoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListLiveRecordVideo", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLiveRecordVideoResponse creates a response to parse from ListLiveRecordVideo response
func CreateListLiveRecordVideoResponse() (response *ListLiveRecordVideoResponse) {
	response = &ListLiveRecordVideoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
