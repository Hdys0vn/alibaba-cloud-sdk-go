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

// ListLogicDatabases invokes the dms_enterprise.ListLogicDatabases API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listlogicdatabases.html
func (client *Client) ListLogicDatabases(request *ListLogicDatabasesRequest) (response *ListLogicDatabasesResponse, err error) {
	response = CreateListLogicDatabasesResponse()
	err = client.DoAction(request, response)
	return
}

// ListLogicDatabasesWithChan invokes the dms_enterprise.ListLogicDatabases API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listlogicdatabases.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListLogicDatabasesWithChan(request *ListLogicDatabasesRequest) (<-chan *ListLogicDatabasesResponse, <-chan error) {
	responseChan := make(chan *ListLogicDatabasesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLogicDatabases(request)
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

// ListLogicDatabasesWithCallback invokes the dms_enterprise.ListLogicDatabases API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listlogicdatabases.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListLogicDatabasesWithCallback(request *ListLogicDatabasesRequest, callback func(response *ListLogicDatabasesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLogicDatabasesResponse
		var err error
		defer close(result)
		response, err = client.ListLogicDatabases(request)
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

// ListLogicDatabasesRequest is the request struct for api ListLogicDatabases
type ListLogicDatabasesRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListLogicDatabasesResponse is the response struct for api ListLogicDatabases
type ListLogicDatabasesResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	ErrorMessage      string            `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode         string            `json:"ErrorCode" xml:"ErrorCode"`
	TotalCount        int64             `json:"TotalCount" xml:"TotalCount"`
	LogicDatabaseList LogicDatabaseList `json:"LogicDatabaseList" xml:"LogicDatabaseList"`
}

// CreateListLogicDatabasesRequest creates a request to invoke ListLogicDatabases API
func CreateListLogicDatabasesRequest() (request *ListLogicDatabasesRequest) {
	request = &ListLogicDatabasesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListLogicDatabases", "dmsenterprise", "openAPI")
	return
}

// CreateListLogicDatabasesResponse creates a response to parse from ListLogicDatabases response
func CreateListLogicDatabasesResponse() (response *ListLogicDatabasesResponse) {
	response = &ListLogicDatabasesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
