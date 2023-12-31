package scdn

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

// SetScdnBotInfo invokes the scdn.SetScdnBotInfo API synchronously
func (client *Client) SetScdnBotInfo(request *SetScdnBotInfoRequest) (response *SetScdnBotInfoResponse, err error) {
	response = CreateSetScdnBotInfoResponse()
	err = client.DoAction(request, response)
	return
}

// SetScdnBotInfoWithChan invokes the scdn.SetScdnBotInfo API asynchronously
func (client *Client) SetScdnBotInfoWithChan(request *SetScdnBotInfoRequest) (<-chan *SetScdnBotInfoResponse, <-chan error) {
	responseChan := make(chan *SetScdnBotInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetScdnBotInfo(request)
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

// SetScdnBotInfoWithCallback invokes the scdn.SetScdnBotInfo API asynchronously
func (client *Client) SetScdnBotInfoWithCallback(request *SetScdnBotInfoRequest, callback func(response *SetScdnBotInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetScdnBotInfoResponse
		var err error
		defer close(result)
		response, err = client.SetScdnBotInfo(request)
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

// SetScdnBotInfoRequest is the request struct for api SetScdnBotInfo
type SetScdnBotInfoRequest struct {
	*requests.RpcRequest
	Enable     string `position:"Query" name:"Enable"`
	DomainName string `position:"Query" name:"DomainName"`
	Status     string `position:"Query" name:"Status"`
}

// SetScdnBotInfoResponse is the response struct for api SetScdnBotInfo
type SetScdnBotInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetScdnBotInfoRequest creates a request to invoke SetScdnBotInfo API
func CreateSetScdnBotInfoRequest() (request *SetScdnBotInfoRequest) {
	request = &SetScdnBotInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "SetScdnBotInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateSetScdnBotInfoResponse creates a response to parse from SetScdnBotInfo response
func CreateSetScdnBotInfoResponse() (response *SetScdnBotInfoResponse) {
	response = &SetScdnBotInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
