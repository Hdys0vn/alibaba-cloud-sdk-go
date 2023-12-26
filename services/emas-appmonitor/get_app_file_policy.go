package emas_appmonitor

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

// GetAppFilePolicy invokes the emas_appmonitor.GetAppFilePolicy API synchronously
func (client *Client) GetAppFilePolicy(request *GetAppFilePolicyRequest) (response *GetAppFilePolicyResponse, err error) {
	response = CreateGetAppFilePolicyResponse()
	err = client.DoAction(request, response)
	return
}

// GetAppFilePolicyWithChan invokes the emas_appmonitor.GetAppFilePolicy API asynchronously
func (client *Client) GetAppFilePolicyWithChan(request *GetAppFilePolicyRequest) (<-chan *GetAppFilePolicyResponse, <-chan error) {
	responseChan := make(chan *GetAppFilePolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAppFilePolicy(request)
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

// GetAppFilePolicyWithCallback invokes the emas_appmonitor.GetAppFilePolicy API asynchronously
func (client *Client) GetAppFilePolicyWithCallback(request *GetAppFilePolicyRequest, callback func(response *GetAppFilePolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAppFilePolicyResponse
		var err error
		defer close(result)
		response, err = client.GetAppFilePolicy(request)
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

// GetAppFilePolicyRequest is the request struct for api GetAppFilePolicy
type GetAppFilePolicyRequest struct {
	*requests.RpcRequest
	FileType    string `position:"Body" name:"FileType"`
	UniqueAppId string `position:"Body" name:"UniqueAppId"`
}

// GetAppFilePolicyResponse is the response struct for api GetAppFilePolicy
type GetAppFilePolicyResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	OSSPolicy OSSPolicy `json:"OSSPolicy" xml:"OSSPolicy"`
}

// CreateGetAppFilePolicyRequest creates a request to invoke GetAppFilePolicy API
func CreateGetAppFilePolicyRequest() (request *GetAppFilePolicyRequest) {
	request = &GetAppFilePolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "GetAppFilePolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAppFilePolicyResponse creates a response to parse from GetAppFilePolicy response
func CreateGetAppFilePolicyResponse() (response *GetAppFilePolicyResponse) {
	response = &GetAppFilePolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}