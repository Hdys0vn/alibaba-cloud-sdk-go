package mts

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

// UnregisterCustomFace invokes the mts.UnregisterCustomFace API synchronously
func (client *Client) UnregisterCustomFace(request *UnregisterCustomFaceRequest) (response *UnregisterCustomFaceResponse, err error) {
	response = CreateUnregisterCustomFaceResponse()
	err = client.DoAction(request, response)
	return
}

// UnregisterCustomFaceWithChan invokes the mts.UnregisterCustomFace API asynchronously
func (client *Client) UnregisterCustomFaceWithChan(request *UnregisterCustomFaceRequest) (<-chan *UnregisterCustomFaceResponse, <-chan error) {
	responseChan := make(chan *UnregisterCustomFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnregisterCustomFace(request)
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

// UnregisterCustomFaceWithCallback invokes the mts.UnregisterCustomFace API asynchronously
func (client *Client) UnregisterCustomFaceWithCallback(request *UnregisterCustomFaceRequest, callback func(response *UnregisterCustomFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnregisterCustomFaceResponse
		var err error
		defer close(result)
		response, err = client.UnregisterCustomFace(request)
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

// UnregisterCustomFaceRequest is the request struct for api UnregisterCustomFace
type UnregisterCustomFaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PersonId             string           `position:"Query" name:"PersonId"`
	CategoryId           string           `position:"Query" name:"CategoryId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	FaceId               string           `position:"Query" name:"FaceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// UnregisterCustomFaceResponse is the response struct for api UnregisterCustomFace
type UnregisterCustomFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnregisterCustomFaceRequest creates a request to invoke UnregisterCustomFace API
func CreateUnregisterCustomFaceRequest() (request *UnregisterCustomFaceRequest) {
	request = &UnregisterCustomFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UnregisterCustomFace", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnregisterCustomFaceResponse creates a response to parse from UnregisterCustomFace response
func CreateUnregisterCustomFaceResponse() (response *UnregisterCustomFaceResponse) {
	response = &UnregisterCustomFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
