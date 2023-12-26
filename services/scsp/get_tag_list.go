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

// GetTagList invokes the scsp.GetTagList API synchronously
func (client *Client) GetTagList(request *GetTagListRequest) (response *GetTagListResponse, err error) {
	response = CreateGetTagListResponse()
	err = client.DoAction(request, response)
	return
}

// GetTagListWithChan invokes the scsp.GetTagList API asynchronously
func (client *Client) GetTagListWithChan(request *GetTagListRequest) (<-chan *GetTagListResponse, <-chan error) {
	responseChan := make(chan *GetTagListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTagList(request)
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

// GetTagListWithCallback invokes the scsp.GetTagList API asynchronously
func (client *Client) GetTagListWithCallback(request *GetTagListRequest, callback func(response *GetTagListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTagListResponse
		var err error
		defer close(result)
		response, err = client.GetTagList(request)
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

// GetTagListRequest is the request struct for api GetTagList
type GetTagListRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Body"`
	EntityType string `position:"Body"`
	EntityId   string `position:"Body"`
}

// GetTagListResponse is the response struct for api GetTagList
type GetTagListResponse struct {
	*responses.BaseResponse
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Code      string     `json:"Code" xml:"Code"`
	Success   bool       `json:"Success" xml:"Success"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetTagListRequest creates a request to invoke GetTagList API
func CreateGetTagListRequest() (request *GetTagListRequest) {
	request = &GetTagListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "GetTagList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTagListResponse creates a response to parse from GetTagList response
func CreateGetTagListResponse() (response *GetTagListResponse) {
	response = &GetTagListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
