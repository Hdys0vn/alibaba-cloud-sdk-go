package resourcemanager

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

// CancelPromoteResourceAccount invokes the resourcemanager.CancelPromoteResourceAccount API synchronously
func (client *Client) CancelPromoteResourceAccount(request *CancelPromoteResourceAccountRequest) (response *CancelPromoteResourceAccountResponse, err error) {
	response = CreateCancelPromoteResourceAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CancelPromoteResourceAccountWithChan invokes the resourcemanager.CancelPromoteResourceAccount API asynchronously
func (client *Client) CancelPromoteResourceAccountWithChan(request *CancelPromoteResourceAccountRequest) (<-chan *CancelPromoteResourceAccountResponse, <-chan error) {
	responseChan := make(chan *CancelPromoteResourceAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelPromoteResourceAccount(request)
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

// CancelPromoteResourceAccountWithCallback invokes the resourcemanager.CancelPromoteResourceAccount API asynchronously
func (client *Client) CancelPromoteResourceAccountWithCallback(request *CancelPromoteResourceAccountRequest, callback func(response *CancelPromoteResourceAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelPromoteResourceAccountResponse
		var err error
		defer close(result)
		response, err = client.CancelPromoteResourceAccount(request)
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

// CancelPromoteResourceAccountRequest is the request struct for api CancelPromoteResourceAccount
type CancelPromoteResourceAccountRequest struct {
	*requests.RpcRequest
	RecordId string `position:"Query" name:"RecordId"`
}

// CancelPromoteResourceAccountResponse is the response struct for api CancelPromoteResourceAccount
type CancelPromoteResourceAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelPromoteResourceAccountRequest creates a request to invoke CancelPromoteResourceAccount API
func CreateCancelPromoteResourceAccountRequest() (request *CancelPromoteResourceAccountRequest) {
	request = &CancelPromoteResourceAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "CancelPromoteResourceAccount", "", "")
	request.Method = requests.POST
	return
}

// CreateCancelPromoteResourceAccountResponse creates a response to parse from CancelPromoteResourceAccount response
func CreateCancelPromoteResourceAccountResponse() (response *CancelPromoteResourceAccountResponse) {
	response = &CancelPromoteResourceAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
