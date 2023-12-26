package dms_enterprise

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

// ListLogicTables invokes the dms_enterprise.ListLogicTables API synchronously
func (client *Client) ListLogicTables(request *ListLogicTablesRequest) (response *ListLogicTablesResponse, err error) {
	response = CreateListLogicTablesResponse()
	err = client.DoAction(request, response)
	return
}

// ListLogicTablesWithChan invokes the dms_enterprise.ListLogicTables API asynchronously
func (client *Client) ListLogicTablesWithChan(request *ListLogicTablesRequest) (<-chan *ListLogicTablesResponse, <-chan error) {
	responseChan := make(chan *ListLogicTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLogicTables(request)
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

// ListLogicTablesWithCallback invokes the dms_enterprise.ListLogicTables API asynchronously
func (client *Client) ListLogicTablesWithCallback(request *ListLogicTablesRequest, callback func(response *ListLogicTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLogicTablesResponse
		var err error
		defer close(result)
		response, err = client.ListLogicTables(request)
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

// ListLogicTablesRequest is the request struct for api ListLogicTables
type ListLogicTablesRequest struct {
	*requests.RpcRequest
	SearchName string           `position:"Query" name:"SearchName"`
	ReturnGuid requests.Boolean `position:"Query" name:"ReturnGuid"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	DatabaseId string           `position:"Query" name:"DatabaseId"`
}

// ListLogicTablesResponse is the response struct for api ListLogicTables
type ListLogicTablesResponse struct {
	*responses.BaseResponse
	TotalCount     int64          `json:"TotalCount" xml:"TotalCount"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	ErrorCode      string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool           `json:"Success" xml:"Success"`
	LogicTableList LogicTableList `json:"LogicTableList" xml:"LogicTableList"`
}

// CreateListLogicTablesRequest creates a request to invoke ListLogicTables API
func CreateListLogicTablesRequest() (request *ListLogicTablesRequest) {
	request = &ListLogicTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListLogicTables", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLogicTablesResponse creates a response to parse from ListLogicTables response
func CreateListLogicTablesResponse() (response *ListLogicTablesResponse) {
	response = &ListLogicTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
