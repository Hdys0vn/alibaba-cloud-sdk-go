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

// ListBindAccount invokes the cloud_siem.ListBindAccount API synchronously
func (client *Client) ListBindAccount(request *ListBindAccountRequest) (response *ListBindAccountResponse, err error) {
	response = CreateListBindAccountResponse()
	err = client.DoAction(request, response)
	return
}

// ListBindAccountWithChan invokes the cloud_siem.ListBindAccount API asynchronously
func (client *Client) ListBindAccountWithChan(request *ListBindAccountRequest) (<-chan *ListBindAccountResponse, <-chan error) {
	responseChan := make(chan *ListBindAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBindAccount(request)
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

// ListBindAccountWithCallback invokes the cloud_siem.ListBindAccount API asynchronously
func (client *Client) ListBindAccountWithCallback(request *ListBindAccountRequest, callback func(response *ListBindAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBindAccountResponse
		var err error
		defer close(result)
		response, err = client.ListBindAccount(request)
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

// ListBindAccountRequest is the request struct for api ListBindAccount
type ListBindAccountRequest struct {
	*requests.RpcRequest
	CloudCode string `position:"Body" name:"CloudCode"`
}

// ListBindAccountResponse is the response struct for api ListBindAccount
type ListBindAccountResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListBindAccountRequest creates a request to invoke ListBindAccount API
func CreateListBindAccountRequest() (request *ListBindAccountRequest) {
	request = &ListBindAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListBindAccount", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListBindAccountResponse creates a response to parse from ListBindAccount response
func CreateListBindAccountResponse() (response *ListBindAccountResponse) {
	response = &ListBindAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
