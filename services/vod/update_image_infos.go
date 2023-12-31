package vod

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

// UpdateImageInfos invokes the vod.UpdateImageInfos API synchronously
func (client *Client) UpdateImageInfos(request *UpdateImageInfosRequest) (response *UpdateImageInfosResponse, err error) {
	response = CreateUpdateImageInfosResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateImageInfosWithChan invokes the vod.UpdateImageInfos API asynchronously
func (client *Client) UpdateImageInfosWithChan(request *UpdateImageInfosRequest) (<-chan *UpdateImageInfosResponse, <-chan error) {
	responseChan := make(chan *UpdateImageInfosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateImageInfos(request)
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

// UpdateImageInfosWithCallback invokes the vod.UpdateImageInfos API asynchronously
func (client *Client) UpdateImageInfosWithCallback(request *UpdateImageInfosRequest, callback func(response *UpdateImageInfosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateImageInfosResponse
		var err error
		defer close(result)
		response, err = client.UpdateImageInfos(request)
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

// UpdateImageInfosRequest is the request struct for api UpdateImageInfos
type UpdateImageInfosRequest struct {
	*requests.RpcRequest
	UpdateContent       string           `position:"Query" name:"UpdateContent"`
	ResourceRealOwnerId requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
}

// UpdateImageInfosResponse is the response struct for api UpdateImageInfos
type UpdateImageInfosResponse struct {
	*responses.BaseResponse
	RequestId        string                             `json:"RequestId" xml:"RequestId"`
	NonExistImageIds NonExistImageIdsInUpdateImageInfos `json:"NonExistImageIds" xml:"NonExistImageIds"`
}

// CreateUpdateImageInfosRequest creates a request to invoke UpdateImageInfos API
func CreateUpdateImageInfosRequest() (request *UpdateImageInfosRequest) {
	request = &UpdateImageInfosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "UpdateImageInfos", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateImageInfosResponse creates a response to parse from UpdateImageInfos response
func CreateUpdateImageInfosResponse() (response *UpdateImageInfosResponse) {
	response = &UpdateImageInfosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
