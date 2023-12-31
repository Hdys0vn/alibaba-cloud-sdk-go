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

// GetImageInfo invokes the vod.GetImageInfo API synchronously
func (client *Client) GetImageInfo(request *GetImageInfoRequest) (response *GetImageInfoResponse, err error) {
	response = CreateGetImageInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetImageInfoWithChan invokes the vod.GetImageInfo API asynchronously
func (client *Client) GetImageInfoWithChan(request *GetImageInfoRequest) (<-chan *GetImageInfoResponse, <-chan error) {
	responseChan := make(chan *GetImageInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImageInfo(request)
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

// GetImageInfoWithCallback invokes the vod.GetImageInfo API asynchronously
func (client *Client) GetImageInfoWithCallback(request *GetImageInfoRequest, callback func(response *GetImageInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImageInfoResponse
		var err error
		defer close(result)
		response, err = client.GetImageInfo(request)
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

// GetImageInfoRequest is the request struct for api GetImageInfo
type GetImageInfoRequest struct {
	*requests.RpcRequest
	ImageId     string           `position:"Query" name:"ImageId"`
	OutputType  string           `position:"Query" name:"OutputType"`
	AuthTimeout requests.Integer `position:"Query" name:"AuthTimeout"`
}

// GetImageInfoResponse is the response struct for api GetImageInfo
type GetImageInfoResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	ImageInfo ImageInfo `json:"ImageInfo" xml:"ImageInfo"`
}

// CreateGetImageInfoRequest creates a request to invoke GetImageInfo API
func CreateGetImageInfoRequest() (request *GetImageInfoRequest) {
	request = &GetImageInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetImageInfo", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetImageInfoResponse creates a response to parse from GetImageInfo response
func CreateGetImageInfoResponse() (response *GetImageInfoResponse) {
	response = &GetImageInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
