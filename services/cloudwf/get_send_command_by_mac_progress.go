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

// GetSendCommandByMacProgress invokes the cloudwf.GetSendCommandByMacProgress API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getsendcommandbymacprogress.html
func (client *Client) GetSendCommandByMacProgress(request *GetSendCommandByMacProgressRequest) (response *GetSendCommandByMacProgressResponse, err error) {
	response = CreateGetSendCommandByMacProgressResponse()
	err = client.DoAction(request, response)
	return
}

// GetSendCommandByMacProgressWithChan invokes the cloudwf.GetSendCommandByMacProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getsendcommandbymacprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSendCommandByMacProgressWithChan(request *GetSendCommandByMacProgressRequest) (<-chan *GetSendCommandByMacProgressResponse, <-chan error) {
	responseChan := make(chan *GetSendCommandByMacProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSendCommandByMacProgress(request)
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

// GetSendCommandByMacProgressWithCallback invokes the cloudwf.GetSendCommandByMacProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getsendcommandbymacprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSendCommandByMacProgressWithCallback(request *GetSendCommandByMacProgressRequest, callback func(response *GetSendCommandByMacProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSendCommandByMacProgressResponse
		var err error
		defer close(result)
		response, err = client.GetSendCommandByMacProgress(request)
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

// GetSendCommandByMacProgressRequest is the request struct for api GetSendCommandByMacProgress
type GetSendCommandByMacProgressRequest struct {
	*requests.RpcRequest
}

// GetSendCommandByMacProgressResponse is the response struct for api GetSendCommandByMacProgress
type GetSendCommandByMacProgressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetSendCommandByMacProgressRequest creates a request to invoke GetSendCommandByMacProgress API
func CreateGetSendCommandByMacProgressRequest() (request *GetSendCommandByMacProgressRequest) {
	request = &GetSendCommandByMacProgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetSendCommandByMacProgress", "cloudwf", "openAPI")
	return
}

// CreateGetSendCommandByMacProgressResponse creates a response to parse from GetSendCommandByMacProgress response
func CreateGetSendCommandByMacProgressResponse() (response *GetSendCommandByMacProgressResponse) {
	response = &GetSendCommandByMacProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}