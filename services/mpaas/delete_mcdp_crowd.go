package mpaas

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

// DeleteMcdpCrowd invokes the mpaas.DeleteMcdpCrowd API synchronously
func (client *Client) DeleteMcdpCrowd(request *DeleteMcdpCrowdRequest) (response *DeleteMcdpCrowdResponse, err error) {
	response = CreateDeleteMcdpCrowdResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMcdpCrowdWithChan invokes the mpaas.DeleteMcdpCrowd API asynchronously
func (client *Client) DeleteMcdpCrowdWithChan(request *DeleteMcdpCrowdRequest) (<-chan *DeleteMcdpCrowdResponse, <-chan error) {
	responseChan := make(chan *DeleteMcdpCrowdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMcdpCrowd(request)
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

// DeleteMcdpCrowdWithCallback invokes the mpaas.DeleteMcdpCrowd API asynchronously
func (client *Client) DeleteMcdpCrowdWithCallback(request *DeleteMcdpCrowdRequest, callback func(response *DeleteMcdpCrowdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMcdpCrowdResponse
		var err error
		defer close(result)
		response, err = client.DeleteMcdpCrowd(request)
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

// DeleteMcdpCrowdRequest is the request struct for api DeleteMcdpCrowd
type DeleteMcdpCrowdRequest struct {
	*requests.RpcRequest
	TenantId                              string `position:"Body" name:"TenantId"`
	AppId                                 string `position:"Body" name:"AppId"`
	MpaasMappcenterMcdpCrowdDeleteJsonStr string `position:"Body" name:"MpaasMappcenterMcdpCrowdDeleteJsonStr"`
	WorkspaceId                           string `position:"Body" name:"WorkspaceId"`
}

// DeleteMcdpCrowdResponse is the response struct for api DeleteMcdpCrowd
type DeleteMcdpCrowdResponse struct {
	*responses.BaseResponse
	ResultMessage string                         `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                         `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                         `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInDeleteMcdpCrowd `json:"ResultContent" xml:"ResultContent"`
}

// CreateDeleteMcdpCrowdRequest creates a request to invoke DeleteMcdpCrowd API
func CreateDeleteMcdpCrowdRequest() (request *DeleteMcdpCrowdRequest) {
	request = &DeleteMcdpCrowdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DeleteMcdpCrowd", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMcdpCrowdResponse creates a response to parse from DeleteMcdpCrowd response
func CreateDeleteMcdpCrowdResponse() (response *DeleteMcdpCrowdResponse) {
	response = &DeleteMcdpCrowdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
