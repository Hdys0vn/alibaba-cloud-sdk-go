package ecd

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

// GetDesktopGroupDetail invokes the ecd.GetDesktopGroupDetail API synchronously
func (client *Client) GetDesktopGroupDetail(request *GetDesktopGroupDetailRequest) (response *GetDesktopGroupDetailResponse, err error) {
	response = CreateGetDesktopGroupDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetDesktopGroupDetailWithChan invokes the ecd.GetDesktopGroupDetail API asynchronously
func (client *Client) GetDesktopGroupDetailWithChan(request *GetDesktopGroupDetailRequest) (<-chan *GetDesktopGroupDetailResponse, <-chan error) {
	responseChan := make(chan *GetDesktopGroupDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDesktopGroupDetail(request)
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

// GetDesktopGroupDetailWithCallback invokes the ecd.GetDesktopGroupDetail API asynchronously
func (client *Client) GetDesktopGroupDetailWithCallback(request *GetDesktopGroupDetailRequest, callback func(response *GetDesktopGroupDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDesktopGroupDetailResponse
		var err error
		defer close(result)
		response, err = client.GetDesktopGroupDetail(request)
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

// GetDesktopGroupDetailRequest is the request struct for api GetDesktopGroupDetail
type GetDesktopGroupDetailRequest struct {
	*requests.RpcRequest
	DesktopGroupId string `position:"Query" name:"DesktopGroupId"`
}

// GetDesktopGroupDetailResponse is the response struct for api GetDesktopGroupDetail
type GetDesktopGroupDetailResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Desktops  Desktops `json:"Desktops" xml:"Desktops"`
}

// CreateGetDesktopGroupDetailRequest creates a request to invoke GetDesktopGroupDetail API
func CreateGetDesktopGroupDetailRequest() (request *GetDesktopGroupDetailRequest) {
	request = &GetDesktopGroupDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "GetDesktopGroupDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDesktopGroupDetailResponse creates a response to parse from GetDesktopGroupDetail response
func CreateGetDesktopGroupDetailResponse() (response *GetDesktopGroupDetailResponse) {
	response = &GetDesktopGroupDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}