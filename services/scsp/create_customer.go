package scsp

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

// CreateCustomer invokes the scsp.CreateCustomer API synchronously
func (client *Client) CreateCustomer(request *CreateCustomerRequest) (response *CreateCustomerResponse, err error) {
	response = CreateCreateCustomerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCustomerWithChan invokes the scsp.CreateCustomer API asynchronously
func (client *Client) CreateCustomerWithChan(request *CreateCustomerRequest) (<-chan *CreateCustomerResponse, <-chan error) {
	responseChan := make(chan *CreateCustomerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCustomer(request)
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

// CreateCustomerWithCallback invokes the scsp.CreateCustomer API asynchronously
func (client *Client) CreateCustomerWithCallback(request *CreateCustomerRequest, callback func(response *CreateCustomerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCustomerResponse
		var err error
		defer close(result)
		response, err = client.CreateCustomer(request)
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

// CreateCustomerRequest is the request struct for api CreateCustomer
type CreateCustomerRequest struct {
	*requests.RpcRequest
	Industry    string           `position:"Query"`
	OuterIdType requests.Integer `position:"Query"`
	Dingding    string           `position:"Query"`
	BizType     string           `position:"Query"`
	TypeCode    string           `position:"Query"`
	InstanceId  string           `position:"Query"`
	Contacter   string           `position:"Query"`
	ProdLineId  requests.Integer `position:"Query"`
	Phone       string           `position:"Query"`
	Name        string           `position:"Query"`
	ManagerName string           `position:"Query"`
	OuterId     string           `position:"Query"`
	Position    string           `position:"Query"`
	Email       string           `position:"Query"`
}

// CreateCustomerResponse is the response struct for api CreateCustomer
type CreateCustomerResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      int64  `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateCustomerRequest creates a request to invoke CreateCustomer API
func CreateCreateCustomerRequest() (request *CreateCustomerRequest) {
	request = &CreateCustomerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "CreateCustomer", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCustomerResponse creates a response to parse from CreateCustomer response
func CreateCreateCustomerResponse() (response *CreateCustomerResponse) {
	response = &CreateCustomerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
