package oos

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

// ListInstancePatches invokes the oos.ListInstancePatches API synchronously
func (client *Client) ListInstancePatches(request *ListInstancePatchesRequest) (response *ListInstancePatchesResponse, err error) {
	response = CreateListInstancePatchesResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstancePatchesWithChan invokes the oos.ListInstancePatches API asynchronously
func (client *Client) ListInstancePatchesWithChan(request *ListInstancePatchesRequest) (<-chan *ListInstancePatchesResponse, <-chan error) {
	responseChan := make(chan *ListInstancePatchesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstancePatches(request)
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

// ListInstancePatchesWithCallback invokes the oos.ListInstancePatches API asynchronously
func (client *Client) ListInstancePatchesWithCallback(request *ListInstancePatchesRequest, callback func(response *ListInstancePatchesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstancePatchesResponse
		var err error
		defer close(result)
		response, err = client.ListInstancePatches(request)
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

// ListInstancePatchesRequest is the request struct for api ListInstancePatches
type ListInstancePatchesRequest struct {
	*requests.RpcRequest
	InstanceId    string           `position:"Query" name:"InstanceId"`
	NextToken     string           `position:"Query" name:"NextToken"`
	MaxResults    requests.Integer `position:"Query" name:"MaxResults"`
	PatchStatuses string           `position:"Query" name:"PatchStatuses"`
}

// ListInstancePatchesResponse is the response struct for api ListInstancePatches
type ListInstancePatchesResponse struct {
	*responses.BaseResponse
	NextToken  string  `json:"NextToken" xml:"NextToken"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	MaxResults int     `json:"MaxResults" xml:"MaxResults"`
	Patches    []Patch `json:"Patches" xml:"Patches"`
}

// CreateListInstancePatchesRequest creates a request to invoke ListInstancePatches API
func CreateListInstancePatchesRequest() (request *ListInstancePatchesRequest) {
	request = &ListInstancePatchesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ListInstancePatches", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstancePatchesResponse creates a response to parse from ListInstancePatches response
func CreateListInstancePatchesResponse() (response *ListInstancePatchesResponse) {
	response = &ListInstancePatchesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
