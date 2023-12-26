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

// GetNodeChildren invokes the dataworks_public.GetNodeChildren API synchronously
func (client *Client) GetNodeChildren(request *GetNodeChildrenRequest) (response *GetNodeChildrenResponse, err error) {
	response = CreateGetNodeChildrenResponse()
	err = client.DoAction(request, response)
	return
}

// GetNodeChildrenWithChan invokes the dataworks_public.GetNodeChildren API asynchronously
func (client *Client) GetNodeChildrenWithChan(request *GetNodeChildrenRequest) (<-chan *GetNodeChildrenResponse, <-chan error) {
	responseChan := make(chan *GetNodeChildrenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNodeChildren(request)
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

// GetNodeChildrenWithCallback invokes the dataworks_public.GetNodeChildren API asynchronously
func (client *Client) GetNodeChildrenWithCallback(request *GetNodeChildrenRequest, callback func(response *GetNodeChildrenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNodeChildrenResponse
		var err error
		defer close(result)
		response, err = client.GetNodeChildren(request)
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

// GetNodeChildrenRequest is the request struct for api GetNodeChildren
type GetNodeChildrenRequest struct {
	*requests.RpcRequest
	ProjectEnv string           `position:"Body" name:"ProjectEnv"`
	NodeId     requests.Integer `position:"Body" name:"NodeId"`
}

// GetNodeChildrenResponse is the response struct for api GetNodeChildren
type GetNodeChildrenResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetNodeChildrenRequest creates a request to invoke GetNodeChildren API
func CreateGetNodeChildrenRequest() (request *GetNodeChildrenRequest) {
	request = &GetNodeChildrenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetNodeChildren", "", "")
	request.Method = requests.POST
	return
}

// CreateGetNodeChildrenResponse creates a response to parse from GetNodeChildren response
func CreateGetNodeChildrenResponse() (response *GetNodeChildrenResponse) {
	response = &GetNodeChildrenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
