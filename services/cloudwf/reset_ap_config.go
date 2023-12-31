package cloudwf

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

// ResetApConfig invokes the cloudwf.ResetApConfig API synchronously
// api document: https://help.aliyun.com/api/cloudwf/resetapconfig.html
func (client *Client) ResetApConfig(request *ResetApConfigRequest) (response *ResetApConfigResponse, err error) {
	response = CreateResetApConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ResetApConfigWithChan invokes the cloudwf.ResetApConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/resetapconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetApConfigWithChan(request *ResetApConfigRequest) (<-chan *ResetApConfigResponse, <-chan error) {
	responseChan := make(chan *ResetApConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetApConfig(request)
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

// ResetApConfigWithCallback invokes the cloudwf.ResetApConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/resetapconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetApConfigWithCallback(request *ResetApConfigRequest, callback func(response *ResetApConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetApConfigResponse
		var err error
		defer close(result)
		response, err = client.ResetApConfig(request)
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

// ResetApConfigRequest is the request struct for api ResetApConfig
type ResetApConfigRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// ResetApConfigResponse is the response struct for api ResetApConfig
type ResetApConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateResetApConfigRequest creates a request to invoke ResetApConfig API
func CreateResetApConfigRequest() (request *ResetApConfigRequest) {
	request = &ResetApConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ResetApConfig", "cloudwf", "openAPI")
	return
}

// CreateResetApConfigResponse creates a response to parse from ResetApConfig response
func CreateResetApConfigResponse() (response *ResetApConfigResponse) {
	response = &ResetApConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
