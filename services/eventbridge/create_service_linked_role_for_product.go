package eventbridge

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

// CreateServiceLinkedRoleForProduct invokes the eventbridge.CreateServiceLinkedRoleForProduct API synchronously
func (client *Client) CreateServiceLinkedRoleForProduct(request *CreateServiceLinkedRoleForProductRequest) (response *CreateServiceLinkedRoleForProductResponse, err error) {
	response = CreateCreateServiceLinkedRoleForProductResponse()
	err = client.DoAction(request, response)
	return
}

// CreateServiceLinkedRoleForProductWithChan invokes the eventbridge.CreateServiceLinkedRoleForProduct API asynchronously
func (client *Client) CreateServiceLinkedRoleForProductWithChan(request *CreateServiceLinkedRoleForProductRequest) (<-chan *CreateServiceLinkedRoleForProductResponse, <-chan error) {
	responseChan := make(chan *CreateServiceLinkedRoleForProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateServiceLinkedRoleForProduct(request)
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

// CreateServiceLinkedRoleForProductWithCallback invokes the eventbridge.CreateServiceLinkedRoleForProduct API asynchronously
func (client *Client) CreateServiceLinkedRoleForProductWithCallback(request *CreateServiceLinkedRoleForProductRequest, callback func(response *CreateServiceLinkedRoleForProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateServiceLinkedRoleForProductResponse
		var err error
		defer close(result)
		response, err = client.CreateServiceLinkedRoleForProduct(request)
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

// CreateServiceLinkedRoleForProductRequest is the request struct for api CreateServiceLinkedRoleForProduct
type CreateServiceLinkedRoleForProductRequest struct {
	*requests.RpcRequest
	ProductName string `position:"Query" name:"ProductName"`
}

// CreateServiceLinkedRoleForProductResponse is the response struct for api CreateServiceLinkedRoleForProduct
type CreateServiceLinkedRoleForProductResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateServiceLinkedRoleForProductRequest creates a request to invoke CreateServiceLinkedRoleForProduct API
func CreateCreateServiceLinkedRoleForProductRequest() (request *CreateServiceLinkedRoleForProductRequest) {
	request = &CreateServiceLinkedRoleForProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "CreateServiceLinkedRoleForProduct", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateServiceLinkedRoleForProductResponse creates a response to parse from CreateServiceLinkedRoleForProduct response
func CreateCreateServiceLinkedRoleForProductResponse() (response *CreateServiceLinkedRoleForProductResponse) {
	response = &CreateServiceLinkedRoleForProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}