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

// ListProxySQLExecAuditLog invokes the dms_enterprise.ListProxySQLExecAuditLog API synchronously
func (client *Client) ListProxySQLExecAuditLog(request *ListProxySQLExecAuditLogRequest) (response *ListProxySQLExecAuditLogResponse, err error) {
	response = CreateListProxySQLExecAuditLogResponse()
	err = client.DoAction(request, response)
	return
}

// ListProxySQLExecAuditLogWithChan invokes the dms_enterprise.ListProxySQLExecAuditLog API asynchronously
func (client *Client) ListProxySQLExecAuditLogWithChan(request *ListProxySQLExecAuditLogRequest) (<-chan *ListProxySQLExecAuditLogResponse, <-chan error) {
	responseChan := make(chan *ListProxySQLExecAuditLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProxySQLExecAuditLog(request)
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

// ListProxySQLExecAuditLogWithCallback invokes the dms_enterprise.ListProxySQLExecAuditLog API asynchronously
func (client *Client) ListProxySQLExecAuditLogWithCallback(request *ListProxySQLExecAuditLogRequest, callback func(response *ListProxySQLExecAuditLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProxySQLExecAuditLogResponse
		var err error
		defer close(result)
		response, err = client.ListProxySQLExecAuditLog(request)
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

// ListProxySQLExecAuditLogRequest is the request struct for api ListProxySQLExecAuditLog
type ListProxySQLExecAuditLogRequest struct {
	*requests.RpcRequest
	SearchName string           `position:"Query" name:"SearchName"`
	OpUserName string           `position:"Query" name:"OpUserName"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	SQLType    string           `position:"Query" name:"SQLType"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	ExecState  string           `position:"Query" name:"ExecState"`
}

// ListProxySQLExecAuditLogResponse is the response struct for api ListProxySQLExecAuditLog
type ListProxySQLExecAuditLogResponse struct {
	*responses.BaseResponse
	TotalCount               int64                    `json:"TotalCount" xml:"TotalCount"`
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	ErrorCode                string                   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage             string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                  bool                     `json:"Success" xml:"Success"`
	ProxySQLExecAuditLogList ProxySQLExecAuditLogList `json:"ProxySQLExecAuditLogList" xml:"ProxySQLExecAuditLogList"`
}

// CreateListProxySQLExecAuditLogRequest creates a request to invoke ListProxySQLExecAuditLog API
func CreateListProxySQLExecAuditLogRequest() (request *ListProxySQLExecAuditLogRequest) {
	request = &ListProxySQLExecAuditLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListProxySQLExecAuditLog", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListProxySQLExecAuditLogResponse creates a response to parse from ListProxySQLExecAuditLog response
func CreateListProxySQLExecAuditLogResponse() (response *ListProxySQLExecAuditLogResponse) {
	response = &ListProxySQLExecAuditLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
