package scsp

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

// GetAuthInfo invokes the scsp.GetAuthInfo API synchronously
func (client *Client) GetAuthInfo(request *GetAuthInfoRequest) (response *GetAuthInfoResponse, err error) {
	response = CreateGetAuthInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetAuthInfoWithChan invokes the scsp.GetAuthInfo API asynchronously
func (client *Client) GetAuthInfoWithChan(request *GetAuthInfoRequest) (<-chan *GetAuthInfoResponse, <-chan error) {
	responseChan := make(chan *GetAuthInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAuthInfo(request)
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

// GetAuthInfoWithCallback invokes the scsp.GetAuthInfo API asynchronously
func (client *Client) GetAuthInfoWithCallback(request *GetAuthInfoRequest, callback func(response *GetAuthInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAuthInfoResponse
		var err error
		defer close(result)
		response, err = client.GetAuthInfo(request)
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

// GetAuthInfoRequest is the request struct for api GetAuthInfo
type GetAuthInfoRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query"`
	ForeignId  string `position:"Query"`
	AppKey     string `position:"Query"`
}

// GetAuthInfoResponse is the response struct for api GetAuthInfo
type GetAuthInfoResponse struct {
	*responses.BaseResponse
	Message   string            `json:"Message" xml:"Message"`
	RequestId string            `json:"RequestId" xml:"RequestId"`
	Code      string            `json:"Code" xml:"Code"`
	Success   bool              `json:"Success" xml:"Success"`
	Data      DataInGetAuthInfo `json:"Data" xml:"Data"`
}

// CreateGetAuthInfoRequest creates a request to invoke GetAuthInfo API
func CreateGetAuthInfoRequest() (request *GetAuthInfoRequest) {
	request = &GetAuthInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "GetAuthInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAuthInfoResponse creates a response to parse from GetAuthInfo response
func CreateGetAuthInfoResponse() (response *GetAuthInfoResponse) {
	response = &GetAuthInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}