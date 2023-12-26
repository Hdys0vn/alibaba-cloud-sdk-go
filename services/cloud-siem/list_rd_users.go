package cloud_siem

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

// ListRdUsers invokes the cloud_siem.ListRdUsers API synchronously
func (client *Client) ListRdUsers(request *ListRdUsersRequest) (response *ListRdUsersResponse, err error) {
	response = CreateListRdUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListRdUsersWithChan invokes the cloud_siem.ListRdUsers API asynchronously
func (client *Client) ListRdUsersWithChan(request *ListRdUsersRequest) (<-chan *ListRdUsersResponse, <-chan error) {
	responseChan := make(chan *ListRdUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRdUsers(request)
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

// ListRdUsersWithCallback invokes the cloud_siem.ListRdUsers API asynchronously
func (client *Client) ListRdUsersWithCallback(request *ListRdUsersRequest, callback func(response *ListRdUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRdUsersResponse
		var err error
		defer close(result)
		response, err = client.ListRdUsers(request)
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

// ListRdUsersRequest is the request struct for api ListRdUsers
type ListRdUsersRequest struct {
	*requests.RpcRequest
}

// ListRdUsersResponse is the response struct for api ListRdUsers
type ListRdUsersResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListRdUsersRequest creates a request to invoke ListRdUsers API
func CreateListRdUsersRequest() (request *ListRdUsersRequest) {
	request = &ListRdUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListRdUsers", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRdUsersResponse creates a response to parse from ListRdUsers response
func CreateListRdUsersResponse() (response *ListRdUsersResponse) {
	response = &ListRdUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
