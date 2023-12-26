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

// DescribePlayTopVideos invokes the vod.DescribePlayTopVideos API synchronously
func (client *Client) DescribePlayTopVideos(request *DescribePlayTopVideosRequest) (response *DescribePlayTopVideosResponse, err error) {
	response = CreateDescribePlayTopVideosResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePlayTopVideosWithChan invokes the vod.DescribePlayTopVideos API asynchronously
func (client *Client) DescribePlayTopVideosWithChan(request *DescribePlayTopVideosRequest) (<-chan *DescribePlayTopVideosResponse, <-chan error) {
	responseChan := make(chan *DescribePlayTopVideosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePlayTopVideos(request)
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

// DescribePlayTopVideosWithCallback invokes the vod.DescribePlayTopVideos API asynchronously
func (client *Client) DescribePlayTopVideosWithCallback(request *DescribePlayTopVideosRequest, callback func(response *DescribePlayTopVideosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePlayTopVideosResponse
		var err error
		defer close(result)
		response, err = client.DescribePlayTopVideos(request)
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

// DescribePlayTopVideosRequest is the request struct for api DescribePlayTopVideos
type DescribePlayTopVideosRequest struct {
	*requests.RpcRequest
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	BizDate  string           `position:"Query" name:"BizDate"`
	PageNo   requests.Integer `position:"Query" name:"PageNo"`
}

// DescribePlayTopVideosResponse is the response struct for api DescribePlayTopVideos
type DescribePlayTopVideosResponse struct {
	*responses.BaseResponse
	PageNo        int64         `json:"PageNo" xml:"PageNo"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	PageSize      int64         `json:"PageSize" xml:"PageSize"`
	TotalNum      int64         `json:"TotalNum" xml:"TotalNum"`
	TopPlayVideos TopPlayVideos `json:"TopPlayVideos" xml:"TopPlayVideos"`
}

// CreateDescribePlayTopVideosRequest creates a request to invoke DescribePlayTopVideos API
func CreateDescribePlayTopVideosRequest() (request *DescribePlayTopVideosRequest) {
	request = &DescribePlayTopVideosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribePlayTopVideos", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePlayTopVideosResponse creates a response to parse from DescribePlayTopVideos response
func CreateDescribePlayTopVideosResponse() (response *DescribePlayTopVideosResponse) {
	response = &DescribePlayTopVideosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
