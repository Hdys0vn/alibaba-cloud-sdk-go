package domain

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

// QueryDomainAdminDivision invokes the domain.QueryDomainAdminDivision API synchronously
func (client *Client) QueryDomainAdminDivision(request *QueryDomainAdminDivisionRequest) (response *QueryDomainAdminDivisionResponse, err error) {
	response = CreateQueryDomainAdminDivisionResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainAdminDivisionWithChan invokes the domain.QueryDomainAdminDivision API asynchronously
func (client *Client) QueryDomainAdminDivisionWithChan(request *QueryDomainAdminDivisionRequest) (<-chan *QueryDomainAdminDivisionResponse, <-chan error) {
	responseChan := make(chan *QueryDomainAdminDivisionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainAdminDivision(request)
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

// QueryDomainAdminDivisionWithCallback invokes the domain.QueryDomainAdminDivision API asynchronously
func (client *Client) QueryDomainAdminDivisionWithCallback(request *QueryDomainAdminDivisionRequest, callback func(response *QueryDomainAdminDivisionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainAdminDivisionResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainAdminDivision(request)
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

// QueryDomainAdminDivisionRequest is the request struct for api QueryDomainAdminDivision
type QueryDomainAdminDivisionRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryDomainAdminDivisionResponse is the response struct for api QueryDomainAdminDivision
type QueryDomainAdminDivisionResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	AdminDivisions AdminDivisions `json:"AdminDivisions" xml:"AdminDivisions"`
}

// CreateQueryDomainAdminDivisionRequest creates a request to invoke QueryDomainAdminDivision API
func CreateQueryDomainAdminDivisionRequest() (request *QueryDomainAdminDivisionRequest) {
	request = &QueryDomainAdminDivisionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainAdminDivision", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDomainAdminDivisionResponse creates a response to parse from QueryDomainAdminDivision response
func CreateQueryDomainAdminDivisionResponse() (response *QueryDomainAdminDivisionResponse) {
	response = &QueryDomainAdminDivisionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
