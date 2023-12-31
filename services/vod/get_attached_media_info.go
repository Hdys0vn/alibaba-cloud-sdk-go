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

// GetAttachedMediaInfo invokes the vod.GetAttachedMediaInfo API synchronously
func (client *Client) GetAttachedMediaInfo(request *GetAttachedMediaInfoRequest) (response *GetAttachedMediaInfoResponse, err error) {
	response = CreateGetAttachedMediaInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetAttachedMediaInfoWithChan invokes the vod.GetAttachedMediaInfo API asynchronously
func (client *Client) GetAttachedMediaInfoWithChan(request *GetAttachedMediaInfoRequest) (<-chan *GetAttachedMediaInfoResponse, <-chan error) {
	responseChan := make(chan *GetAttachedMediaInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAttachedMediaInfo(request)
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

// GetAttachedMediaInfoWithCallback invokes the vod.GetAttachedMediaInfo API asynchronously
func (client *Client) GetAttachedMediaInfoWithCallback(request *GetAttachedMediaInfoRequest, callback func(response *GetAttachedMediaInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAttachedMediaInfoResponse
		var err error
		defer close(result)
		response, err = client.GetAttachedMediaInfo(request)
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

// GetAttachedMediaInfoRequest is the request struct for api GetAttachedMediaInfo
type GetAttachedMediaInfoRequest struct {
	*requests.RpcRequest
	ResourceRealOwnerId requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	OutputType          string           `position:"Query" name:"OutputType"`
	MediaIds            string           `position:"Query" name:"MediaIds"`
	AuthTimeout         requests.Integer `position:"Query" name:"AuthTimeout"`
}

// GetAttachedMediaInfoResponse is the response struct for api GetAttachedMediaInfo
type GetAttachedMediaInfoResponse struct {
	*responses.BaseResponse
	RequestId         string          `json:"RequestId" xml:"RequestId"`
	NonExistMediaIds  []string        `json:"NonExistMediaIds" xml:"NonExistMediaIds"`
	AttachedMediaList []AttachedMedia `json:"AttachedMediaList" xml:"AttachedMediaList"`
}

// CreateGetAttachedMediaInfoRequest creates a request to invoke GetAttachedMediaInfo API
func CreateGetAttachedMediaInfoRequest() (request *GetAttachedMediaInfoRequest) {
	request = &GetAttachedMediaInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetAttachedMediaInfo", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAttachedMediaInfoResponse creates a response to parse from GetAttachedMediaInfo response
func CreateGetAttachedMediaInfoResponse() (response *GetAttachedMediaInfoResponse) {
	response = &GetAttachedMediaInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
