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

// ResetDesktops invokes the ecd.ResetDesktops API synchronously
func (client *Client) ResetDesktops(request *ResetDesktopsRequest) (response *ResetDesktopsResponse, err error) {
	response = CreateResetDesktopsResponse()
	err = client.DoAction(request, response)
	return
}

// ResetDesktopsWithChan invokes the ecd.ResetDesktops API asynchronously
func (client *Client) ResetDesktopsWithChan(request *ResetDesktopsRequest) (<-chan *ResetDesktopsResponse, <-chan error) {
	responseChan := make(chan *ResetDesktopsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetDesktops(request)
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

// ResetDesktopsWithCallback invokes the ecd.ResetDesktops API asynchronously
func (client *Client) ResetDesktopsWithCallback(request *ResetDesktopsRequest, callback func(response *ResetDesktopsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetDesktopsResponse
		var err error
		defer close(result)
		response, err = client.ResetDesktops(request)
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

// ResetDesktopsRequest is the request struct for api ResetDesktops
type ResetDesktopsRequest struct {
	*requests.RpcRequest
	ImageId        string    `position:"Query" name:"ImageId"`
	ResetType      string    `position:"Query" name:"ResetType"`
	DesktopGroupId string    `position:"Query" name:"DesktopGroupId"`
	DesktopId      *[]string `position:"Query" name:"DesktopId"  type:"Repeated"`
	PayType        string    `position:"Query" name:"PayType"`
}

// ResetDesktopsResponse is the response struct for api ResetDesktops
type ResetDesktopsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetDesktopsRequest creates a request to invoke ResetDesktops API
func CreateResetDesktopsRequest() (request *ResetDesktopsRequest) {
	request = &ResetDesktopsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ResetDesktops", "", "")
	request.Method = requests.POST
	return
}

// CreateResetDesktopsResponse creates a response to parse from ResetDesktops response
func CreateResetDesktopsResponse() (response *ResetDesktopsResponse) {
	response = &ResetDesktopsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
