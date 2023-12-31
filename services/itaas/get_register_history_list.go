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

// GetRegisterHistoryList invokes the itaas.GetRegisterHistoryList API synchronously
// api document: https://help.aliyun.com/api/itaas/getregisterhistorylist.html
func (client *Client) GetRegisterHistoryList(request *GetRegisterHistoryListRequest) (response *GetRegisterHistoryListResponse, err error) {
	response = CreateGetRegisterHistoryListResponse()
	err = client.DoAction(request, response)
	return
}

// GetRegisterHistoryListWithChan invokes the itaas.GetRegisterHistoryList API asynchronously
// api document: https://help.aliyun.com/api/itaas/getregisterhistorylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRegisterHistoryListWithChan(request *GetRegisterHistoryListRequest) (<-chan *GetRegisterHistoryListResponse, <-chan error) {
	responseChan := make(chan *GetRegisterHistoryListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRegisterHistoryList(request)
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

// GetRegisterHistoryListWithCallback invokes the itaas.GetRegisterHistoryList API asynchronously
// api document: https://help.aliyun.com/api/itaas/getregisterhistorylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRegisterHistoryListWithCallback(request *GetRegisterHistoryListRequest, callback func(response *GetRegisterHistoryListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRegisterHistoryListResponse
		var err error
		defer close(result)
		response, err = client.GetRegisterHistoryList(request)
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

// GetRegisterHistoryListRequest is the request struct for api GetRegisterHistoryList
type GetRegisterHistoryListRequest struct {
	*requests.RpcRequest
	Clientappid string `position:"Query" name:"Clientappid"`
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
}

// GetRegisterHistoryListResponse is the response struct for api GetRegisterHistoryList
type GetRegisterHistoryListResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	ErrorCode int                               `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string                            `json:"ErrorMsg" xml:"ErrorMsg"`
	Success   bool                              `json:"Success" xml:"Success"`
	Data      DataInGetRegisterHistoryList      `json:"Data" xml:"Data"`
	ErrorList ErrorListInGetRegisterHistoryList `json:"ErrorList" xml:"ErrorList"`
}

// CreateGetRegisterHistoryListRequest creates a request to invoke GetRegisterHistoryList API
func CreateGetRegisterHistoryListRequest() (request *GetRegisterHistoryListRequest) {
	request = &GetRegisterHistoryListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-05", "GetRegisterHistoryList", "itaas", "openAPI")
	return
}

// CreateGetRegisterHistoryListResponse creates a response to parse from GetRegisterHistoryList response
func CreateGetRegisterHistoryListResponse() (response *GetRegisterHistoryListResponse) {
	response = &GetRegisterHistoryListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
