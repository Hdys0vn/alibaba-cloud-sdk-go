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

// GetBaselineKeyPath invokes the dataworks_public.GetBaselineKeyPath API synchronously
func (client *Client) GetBaselineKeyPath(request *GetBaselineKeyPathRequest) (response *GetBaselineKeyPathResponse, err error) {
	response = CreateGetBaselineKeyPathResponse()
	err = client.DoAction(request, response)
	return
}

// GetBaselineKeyPathWithChan invokes the dataworks_public.GetBaselineKeyPath API asynchronously
func (client *Client) GetBaselineKeyPathWithChan(request *GetBaselineKeyPathRequest) (<-chan *GetBaselineKeyPathResponse, <-chan error) {
	responseChan := make(chan *GetBaselineKeyPathResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBaselineKeyPath(request)
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

// GetBaselineKeyPathWithCallback invokes the dataworks_public.GetBaselineKeyPath API asynchronously
func (client *Client) GetBaselineKeyPathWithCallback(request *GetBaselineKeyPathRequest, callback func(response *GetBaselineKeyPathResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBaselineKeyPathResponse
		var err error
		defer close(result)
		response, err = client.GetBaselineKeyPath(request)
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

// GetBaselineKeyPathRequest is the request struct for api GetBaselineKeyPath
type GetBaselineKeyPathRequest struct {
	*requests.RpcRequest
	Bizdate    string           `position:"Body" name:"Bizdate"`
	InGroupId  requests.Integer `position:"Body" name:"InGroupId"`
	BaselineId requests.Integer `position:"Body" name:"BaselineId"`
}

// GetBaselineKeyPathResponse is the response struct for api GetBaselineKeyPath
type GetBaselineKeyPathResponse struct {
	*responses.BaseResponse
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string     `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string     `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool       `json:"Success" xml:"Success"`
	Data           []DataItem `json:"Data" xml:"Data"`
}

// CreateGetBaselineKeyPathRequest creates a request to invoke GetBaselineKeyPath API
func CreateGetBaselineKeyPathRequest() (request *GetBaselineKeyPathRequest) {
	request = &GetBaselineKeyPathRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetBaselineKeyPath", "", "")
	request.Method = requests.POST
	return
}

// CreateGetBaselineKeyPathResponse creates a response to parse from GetBaselineKeyPath response
func CreateGetBaselineKeyPathResponse() (response *GetBaselineKeyPathResponse) {
	response = &GetBaselineKeyPathResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
