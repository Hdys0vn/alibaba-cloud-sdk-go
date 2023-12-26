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

// GetMezzanineInfo invokes the vod.GetMezzanineInfo API synchronously
func (client *Client) GetMezzanineInfo(request *GetMezzanineInfoRequest) (response *GetMezzanineInfoResponse, err error) {
	response = CreateGetMezzanineInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetMezzanineInfoWithChan invokes the vod.GetMezzanineInfo API asynchronously
func (client *Client) GetMezzanineInfoWithChan(request *GetMezzanineInfoRequest) (<-chan *GetMezzanineInfoResponse, <-chan error) {
	responseChan := make(chan *GetMezzanineInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMezzanineInfo(request)
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

// GetMezzanineInfoWithCallback invokes the vod.GetMezzanineInfo API asynchronously
func (client *Client) GetMezzanineInfoWithCallback(request *GetMezzanineInfoRequest, callback func(response *GetMezzanineInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMezzanineInfoResponse
		var err error
		defer close(result)
		response, err = client.GetMezzanineInfo(request)
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

// GetMezzanineInfoRequest is the request struct for api GetMezzanineInfo
type GetMezzanineInfoRequest struct {
	*requests.RpcRequest
	OutputType     string           `position:"Query" name:"OutputType"`
	AuthTimeout    requests.Integer `position:"Query" name:"AuthTimeout"`
	VideoId        string           `position:"Query" name:"VideoId"`
	PreviewSegment requests.Boolean `position:"Query" name:"PreviewSegment"`
	AdditionType   string           `position:"Query" name:"AdditionType"`
}

// GetMezzanineInfoResponse is the response struct for api GetMezzanineInfo
type GetMezzanineInfoResponse struct {
	*responses.BaseResponse
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Mezzanine MezzanineInGetMezzanineInfo `json:"Mezzanine" xml:"Mezzanine"`
}

// CreateGetMezzanineInfoRequest creates a request to invoke GetMezzanineInfo API
func CreateGetMezzanineInfoRequest() (request *GetMezzanineInfoRequest) {
	request = &GetMezzanineInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetMezzanineInfo", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMezzanineInfoResponse creates a response to parse from GetMezzanineInfo response
func CreateGetMezzanineInfoResponse() (response *GetMezzanineInfoResponse) {
	response = &GetMezzanineInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
