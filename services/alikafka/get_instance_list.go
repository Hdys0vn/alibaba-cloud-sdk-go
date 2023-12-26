package alikafka

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

// GetInstanceList invokes the alikafka.GetInstanceList API synchronously
func (client *Client) GetInstanceList(request *GetInstanceListRequest) (response *GetInstanceListResponse, err error) {
	response = CreateGetInstanceListResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceListWithChan invokes the alikafka.GetInstanceList API asynchronously
func (client *Client) GetInstanceListWithChan(request *GetInstanceListRequest) (<-chan *GetInstanceListResponse, <-chan error) {
	responseChan := make(chan *GetInstanceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceList(request)
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

// GetInstanceListWithCallback invokes the alikafka.GetInstanceList API asynchronously
func (client *Client) GetInstanceListWithCallback(request *GetInstanceListRequest, callback func(response *GetInstanceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceListResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceList(request)
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

// GetInstanceListRequest is the request struct for api GetInstanceList
type GetInstanceListRequest struct {
	*requests.RpcRequest
	OrderId         string                `position:"Query" name:"OrderId"`
	ResourceGroupId string                `position:"Query" name:"ResourceGroupId"`
	InstanceId      *[]string             `position:"Query" name:"InstanceId"  type:"Repeated"`
	Tag             *[]GetInstanceListTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// GetInstanceListTag is a repeated param struct in GetInstanceListRequest
type GetInstanceListTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// GetInstanceListResponse is the response struct for api GetInstanceList
type GetInstanceListResponse struct {
	*responses.BaseResponse
	Code         int          `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Success      bool         `json:"Success" xml:"Success"`
	InstanceList InstanceList `json:"InstanceList" xml:"InstanceList"`
}

// CreateGetInstanceListRequest creates a request to invoke GetInstanceList API
func CreateGetInstanceListRequest() (request *GetInstanceListRequest) {
	request = &GetInstanceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "GetInstanceList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInstanceListResponse creates a response to parse from GetInstanceList response
func CreateGetInstanceListResponse() (response *GetInstanceListResponse) {
	response = &GetInstanceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
