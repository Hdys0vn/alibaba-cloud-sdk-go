package bssopenapi

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

// CreateAgAccount invokes the bssopenapi.CreateAgAccount API synchronously
func (client *Client) CreateAgAccount(request *CreateAgAccountRequest) (response *CreateAgAccountResponse, err error) {
	response = CreateCreateAgAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAgAccountWithChan invokes the bssopenapi.CreateAgAccount API asynchronously
func (client *Client) CreateAgAccountWithChan(request *CreateAgAccountRequest) (<-chan *CreateAgAccountResponse, <-chan error) {
	responseChan := make(chan *CreateAgAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAgAccount(request)
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

// CreateAgAccountWithCallback invokes the bssopenapi.CreateAgAccount API asynchronously
func (client *Client) CreateAgAccountWithCallback(request *CreateAgAccountRequest, callback func(response *CreateAgAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAgAccountResponse
		var err error
		defer close(result)
		response, err = client.CreateAgAccount(request)
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

// CreateAgAccountRequest is the request struct for api CreateAgAccount
type CreateAgAccountRequest struct {
	*requests.RpcRequest
	FirstName      string `position:"Query" name:"FirstName"`
	CityName       string `position:"Query" name:"CityName"`
	Postcode       string `position:"Query" name:"Postcode"`
	EnterpriseName string `position:"Query" name:"EnterpriseName"`
	NationCode     string `position:"Query" name:"NationCode"`
	LastName       string `position:"Query" name:"LastName"`
	LoginEmail     string `position:"Query" name:"LoginEmail"`
	ProvinceName   string `position:"Query" name:"ProvinceName"`
	AccountAttr    string `position:"Query" name:"AccountAttr"`
}

// CreateAgAccountResponse is the response struct for api CreateAgAccount
type CreateAgAccountResponse struct {
	*responses.BaseResponse
	Code          string        `json:"Code" xml:"Code"`
	Message       string        `json:"Message" xml:"Message"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	Success       bool          `json:"Success" xml:"Success"`
	AgRelationDto AgRelationDto `json:"AgRelationDto" xml:"AgRelationDto"`
}

// CreateCreateAgAccountRequest creates a request to invoke CreateAgAccount API
func CreateCreateAgAccountRequest() (request *CreateAgAccountRequest) {
	request = &CreateAgAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "CreateAgAccount", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAgAccountResponse creates a response to parse from CreateAgAccount response
func CreateCreateAgAccountResponse() (response *CreateAgAccountResponse) {
	response = &CreateAgAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
