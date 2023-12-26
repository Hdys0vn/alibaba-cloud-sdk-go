package pairecservice

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

// ListTableMetas invokes the pairecservice.ListTableMetas API synchronously
func (client *Client) ListTableMetas(request *ListTableMetasRequest) (response *ListTableMetasResponse, err error) {
	response = CreateListTableMetasResponse()
	err = client.DoAction(request, response)
	return
}

// ListTableMetasWithChan invokes the pairecservice.ListTableMetas API asynchronously
func (client *Client) ListTableMetasWithChan(request *ListTableMetasRequest) (<-chan *ListTableMetasResponse, <-chan error) {
	responseChan := make(chan *ListTableMetasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTableMetas(request)
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

// ListTableMetasWithCallback invokes the pairecservice.ListTableMetas API asynchronously
func (client *Client) ListTableMetasWithCallback(request *ListTableMetasRequest, callback func(response *ListTableMetasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTableMetasResponse
		var err error
		defer close(result)
		response, err = client.ListTableMetas(request)
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

// ListTableMetasRequest is the request struct for api ListTableMetas
type ListTableMetasRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Query" name:"InstanceId"`
	Module     string           `position:"Query" name:"Module"`
	Name       string           `position:"Query" name:"Name"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Type       string           `position:"Query" name:"Type"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListTableMetasResponse is the response struct for api ListTableMetas
type ListTableMetasResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	TotalCount int64            `json:"TotalCount" xml:"TotalCount"`
	TableMetas []TableMetasItem `json:"TableMetas" xml:"TableMetas"`
}

// CreateListTableMetasRequest creates a request to invoke ListTableMetas API
func CreateListTableMetasRequest() (request *ListTableMetasRequest) {
	request = &ListTableMetasRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "ListTableMetas", "/api/v1/tablemetas", "", "")
	request.Method = requests.GET
	return
}

// CreateListTableMetasResponse creates a response to parse from ListTableMetas response
func CreateListTableMetasResponse() (response *ListTableMetasResponse) {
	response = &ListTableMetasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}