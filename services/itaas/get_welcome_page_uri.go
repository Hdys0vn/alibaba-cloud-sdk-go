package itaas

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

// GetWelcomePageURI invokes the itaas.GetWelcomePageURI API synchronously
// api document: https://help.aliyun.com/api/itaas/getwelcomepageuri.html
func (client *Client) GetWelcomePageURI(request *GetWelcomePageURIRequest) (response *GetWelcomePageURIResponse, err error) {
	response = CreateGetWelcomePageURIResponse()
	err = client.DoAction(request, response)
	return
}

// GetWelcomePageURIWithChan invokes the itaas.GetWelcomePageURI API asynchronously
// api document: https://help.aliyun.com/api/itaas/getwelcomepageuri.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetWelcomePageURIWithChan(request *GetWelcomePageURIRequest) (<-chan *GetWelcomePageURIResponse, <-chan error) {
	responseChan := make(chan *GetWelcomePageURIResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWelcomePageURI(request)
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

// GetWelcomePageURIWithCallback invokes the itaas.GetWelcomePageURI API asynchronously
// api document: https://help.aliyun.com/api/itaas/getwelcomepageuri.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetWelcomePageURIWithCallback(request *GetWelcomePageURIRequest, callback func(response *GetWelcomePageURIResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWelcomePageURIResponse
		var err error
		defer close(result)
		response, err = client.GetWelcomePageURI(request)
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

// GetWelcomePageURIRequest is the request struct for api GetWelcomePageURI
type GetWelcomePageURIRequest struct {
	*requests.RpcRequest
	Clientappid string `position:"Query" name:"Clientappid"`
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
}

// GetWelcomePageURIResponse is the response struct for api GetWelcomePageURI
type GetWelcomePageURIResponse struct {
	*responses.BaseResponse
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	Data      string                       `json:"Data" xml:"Data"`
	ErrorCode int                          `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string                       `json:"ErrorMsg" xml:"ErrorMsg"`
	Success   bool                         `json:"Success" xml:"Success"`
	ErrorList ErrorListInGetWelcomePageURI `json:"ErrorList" xml:"ErrorList"`
}

// CreateGetWelcomePageURIRequest creates a request to invoke GetWelcomePageURI API
func CreateGetWelcomePageURIRequest() (request *GetWelcomePageURIRequest) {
	request = &GetWelcomePageURIRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-05", "GetWelcomePageURI", "itaas", "openAPI")
	return
}

// CreateGetWelcomePageURIResponse creates a response to parse from GetWelcomePageURI response
func CreateGetWelcomePageURIResponse() (response *GetWelcomePageURIResponse) {
	response = &GetWelcomePageURIResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
