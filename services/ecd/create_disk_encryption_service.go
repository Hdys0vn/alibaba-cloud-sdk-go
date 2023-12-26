package ecd

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

// CreateDiskEncryptionService invokes the ecd.CreateDiskEncryptionService API synchronously
func (client *Client) CreateDiskEncryptionService(request *CreateDiskEncryptionServiceRequest) (response *CreateDiskEncryptionServiceResponse, err error) {
	response = CreateCreateDiskEncryptionServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDiskEncryptionServiceWithChan invokes the ecd.CreateDiskEncryptionService API asynchronously
func (client *Client) CreateDiskEncryptionServiceWithChan(request *CreateDiskEncryptionServiceRequest) (<-chan *CreateDiskEncryptionServiceResponse, <-chan error) {
	responseChan := make(chan *CreateDiskEncryptionServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDiskEncryptionService(request)
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

// CreateDiskEncryptionServiceWithCallback invokes the ecd.CreateDiskEncryptionService API asynchronously
func (client *Client) CreateDiskEncryptionServiceWithCallback(request *CreateDiskEncryptionServiceRequest, callback func(response *CreateDiskEncryptionServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDiskEncryptionServiceResponse
		var err error
		defer close(result)
		response, err = client.CreateDiskEncryptionService(request)
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

// CreateDiskEncryptionServiceRequest is the request struct for api CreateDiskEncryptionService
type CreateDiskEncryptionServiceRequest struct {
	*requests.RpcRequest
}

// CreateDiskEncryptionServiceResponse is the response struct for api CreateDiskEncryptionService
type CreateDiskEncryptionServiceResponse struct {
	*responses.BaseResponse
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateDiskEncryptionServiceRequest creates a request to invoke CreateDiskEncryptionService API
func CreateCreateDiskEncryptionServiceRequest() (request *CreateDiskEncryptionServiceRequest) {
	request = &CreateDiskEncryptionServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateDiskEncryptionService", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDiskEncryptionServiceResponse creates a response to parse from CreateDiskEncryptionService response
func CreateCreateDiskEncryptionServiceResponse() (response *CreateDiskEncryptionServiceResponse) {
	response = &CreateDiskEncryptionServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
