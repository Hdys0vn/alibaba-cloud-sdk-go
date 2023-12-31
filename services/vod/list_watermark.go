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

// ListWatermark invokes the vod.ListWatermark API synchronously
func (client *Client) ListWatermark(request *ListWatermarkRequest) (response *ListWatermarkResponse, err error) {
	response = CreateListWatermarkResponse()
	err = client.DoAction(request, response)
	return
}

// ListWatermarkWithChan invokes the vod.ListWatermark API asynchronously
func (client *Client) ListWatermarkWithChan(request *ListWatermarkRequest) (<-chan *ListWatermarkResponse, <-chan error) {
	responseChan := make(chan *ListWatermarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListWatermark(request)
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

// ListWatermarkWithCallback invokes the vod.ListWatermark API asynchronously
func (client *Client) ListWatermarkWithCallback(request *ListWatermarkRequest, callback func(response *ListWatermarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListWatermarkResponse
		var err error
		defer close(result)
		response, err = client.ListWatermark(request)
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

// ListWatermarkRequest is the request struct for api ListWatermark
type ListWatermarkRequest struct {
	*requests.RpcRequest
	PageNo   requests.Integer `position:"Query" name:"PageNo"`
	AppId    string           `position:"Query" name:"AppId"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
}

// ListWatermarkResponse is the response struct for api ListWatermark
type ListWatermarkResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	WatermarkInfos []WatermarkInfo `json:"WatermarkInfos" xml:"WatermarkInfos"`
}

// CreateListWatermarkRequest creates a request to invoke ListWatermark API
func CreateListWatermarkRequest() (request *ListWatermarkRequest) {
	request = &ListWatermarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListWatermark", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListWatermarkResponse creates a response to parse from ListWatermark response
func CreateListWatermarkResponse() (response *ListWatermarkResponse) {
	response = &ListWatermarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
