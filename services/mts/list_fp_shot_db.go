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

// ListFpShotDB invokes the mts.ListFpShotDB API synchronously
func (client *Client) ListFpShotDB(request *ListFpShotDBRequest) (response *ListFpShotDBResponse, err error) {
	response = CreateListFpShotDBResponse()
	err = client.DoAction(request, response)
	return
}

// ListFpShotDBWithChan invokes the mts.ListFpShotDB API asynchronously
func (client *Client) ListFpShotDBWithChan(request *ListFpShotDBRequest) (<-chan *ListFpShotDBResponse, <-chan error) {
	responseChan := make(chan *ListFpShotDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFpShotDB(request)
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

// ListFpShotDBWithCallback invokes the mts.ListFpShotDB API asynchronously
func (client *Client) ListFpShotDBWithCallback(request *ListFpShotDBRequest, callback func(response *ListFpShotDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFpShotDBResponse
		var err error
		defer close(result)
		response, err = client.ListFpShotDB(request)
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

// ListFpShotDBRequest is the request struct for api ListFpShotDB
type ListFpShotDBRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FpDBIds              string           `position:"Query" name:"FpDBIds"`
}

// ListFpShotDBResponse is the response struct for api ListFpShotDB
type ListFpShotDBResponse struct {
	*responses.BaseResponse
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	NonExistIds  NonExistIdsInListFpShotDB `json:"NonExistIds" xml:"NonExistIds"`
	FpShotDBList FpShotDBList              `json:"FpShotDBList" xml:"FpShotDBList"`
}

// CreateListFpShotDBRequest creates a request to invoke ListFpShotDB API
func CreateListFpShotDBRequest() (request *ListFpShotDBRequest) {
	request = &ListFpShotDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListFpShotDB", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFpShotDBResponse creates a response to parse from ListFpShotDB response
func CreateListFpShotDBResponse() (response *ListFpShotDBResponse) {
	response = &ListFpShotDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
