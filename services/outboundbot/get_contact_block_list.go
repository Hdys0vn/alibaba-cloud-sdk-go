package outboundbot

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

// GetContactBlockList invokes the outboundbot.GetContactBlockList API synchronously
func (client *Client) GetContactBlockList(request *GetContactBlockListRequest) (response *GetContactBlockListResponse, err error) {
	response = CreateGetContactBlockListResponse()
	err = client.DoAction(request, response)
	return
}

// GetContactBlockListWithChan invokes the outboundbot.GetContactBlockList API asynchronously
func (client *Client) GetContactBlockListWithChan(request *GetContactBlockListRequest) (<-chan *GetContactBlockListResponse, <-chan error) {
	responseChan := make(chan *GetContactBlockListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetContactBlockList(request)
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

// GetContactBlockListWithCallback invokes the outboundbot.GetContactBlockList API asynchronously
func (client *Client) GetContactBlockListWithCallback(request *GetContactBlockListRequest, callback func(response *GetContactBlockListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetContactBlockListResponse
		var err error
		defer close(result)
		response, err = client.GetContactBlockList(request)
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

// GetContactBlockListRequest is the request struct for api GetContactBlockList
type GetContactBlockListRequest struct {
	*requests.RpcRequest
	CountTotalRow requests.Boolean `position:"Query" name:"CountTotalRow"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// GetContactBlockListResponse is the response struct for api GetContactBlockList
type GetContactBlockListResponse struct {
	*responses.BaseResponse
	HttpStatusCode       int                  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code                 string               `json:"Code" xml:"Code"`
	Message              string               `json:"Message" xml:"Message"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Success              bool                 `json:"Success" xml:"Success"`
	ContactBlocklistList ContactBlocklistList `json:"ContactBlocklistList" xml:"ContactBlocklistList"`
}

// CreateGetContactBlockListRequest creates a request to invoke GetContactBlockList API
func CreateGetContactBlockListRequest() (request *GetContactBlockListRequest) {
	request = &GetContactBlockListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "GetContactBlockList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetContactBlockListResponse creates a response to parse from GetContactBlockList response
func CreateGetContactBlockListResponse() (response *GetContactBlockListResponse) {
	response = &GetContactBlockListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
