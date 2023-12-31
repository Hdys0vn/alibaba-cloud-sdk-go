package dataworks_public

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

// GetSuccessInstanceTrend invokes the dataworks_public.GetSuccessInstanceTrend API synchronously
func (client *Client) GetSuccessInstanceTrend(request *GetSuccessInstanceTrendRequest) (response *GetSuccessInstanceTrendResponse, err error) {
	response = CreateGetSuccessInstanceTrendResponse()
	err = client.DoAction(request, response)
	return
}

// GetSuccessInstanceTrendWithChan invokes the dataworks_public.GetSuccessInstanceTrend API asynchronously
func (client *Client) GetSuccessInstanceTrendWithChan(request *GetSuccessInstanceTrendRequest) (<-chan *GetSuccessInstanceTrendResponse, <-chan error) {
	responseChan := make(chan *GetSuccessInstanceTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSuccessInstanceTrend(request)
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

// GetSuccessInstanceTrendWithCallback invokes the dataworks_public.GetSuccessInstanceTrend API asynchronously
func (client *Client) GetSuccessInstanceTrendWithCallback(request *GetSuccessInstanceTrendRequest, callback func(response *GetSuccessInstanceTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSuccessInstanceTrendResponse
		var err error
		defer close(result)
		response, err = client.GetSuccessInstanceTrend(request)
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

// GetSuccessInstanceTrendRequest is the request struct for api GetSuccessInstanceTrend
type GetSuccessInstanceTrendRequest struct {
	*requests.RpcRequest
	ProjectId requests.Integer `position:"Body" name:"ProjectId"`
}

// GetSuccessInstanceTrendResponse is the response struct for api GetSuccessInstanceTrend
type GetSuccessInstanceTrendResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	InstanceStatusTrend InstanceStatusTrend `json:"InstanceStatusTrend" xml:"InstanceStatusTrend"`
}

// CreateGetSuccessInstanceTrendRequest creates a request to invoke GetSuccessInstanceTrend API
func CreateGetSuccessInstanceTrendRequest() (request *GetSuccessInstanceTrendRequest) {
	request = &GetSuccessInstanceTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetSuccessInstanceTrend", "", "")
	request.Method = requests.POST
	return
}

// CreateGetSuccessInstanceTrendResponse creates a response to parse from GetSuccessInstanceTrend response
func CreateGetSuccessInstanceTrendResponse() (response *GetSuccessInstanceTrendResponse) {
	response = &GetSuccessInstanceTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
