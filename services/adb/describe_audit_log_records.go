package adb

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

// DescribeAuditLogRecords invokes the adb.DescribeAuditLogRecords API synchronously
func (client *Client) DescribeAuditLogRecords(request *DescribeAuditLogRecordsRequest) (response *DescribeAuditLogRecordsResponse, err error) {
	response = CreateDescribeAuditLogRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAuditLogRecordsWithChan invokes the adb.DescribeAuditLogRecords API asynchronously
func (client *Client) DescribeAuditLogRecordsWithChan(request *DescribeAuditLogRecordsRequest) (<-chan *DescribeAuditLogRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeAuditLogRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAuditLogRecords(request)
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

// DescribeAuditLogRecordsWithCallback invokes the adb.DescribeAuditLogRecords API asynchronously
func (client *Client) DescribeAuditLogRecordsWithCallback(request *DescribeAuditLogRecordsRequest, callback func(response *DescribeAuditLogRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAuditLogRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAuditLogRecords(request)
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

// DescribeAuditLogRecordsRequest is the request struct for api DescribeAuditLogRecords
type DescribeAuditLogRecordsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	HostAddress          string           `position:"Query" name:"HostAddress"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	Order                string           `position:"Query" name:"Order"`
	SqlType              string           `position:"Query" name:"SqlType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	QueryKeyword         string           `position:"Query" name:"QueryKeyword"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBName               string           `position:"Query" name:"DBName"`
	Succeed              string           `position:"Query" name:"Succeed"`
	User                 string           `position:"Query" name:"User"`
	OrderType            string           `position:"Query" name:"OrderType"`
}

// DescribeAuditLogRecordsResponse is the response struct for api DescribeAuditLogRecords
type DescribeAuditLogRecordsResponse struct {
	*responses.BaseResponse
	TotalCount  string          `json:"TotalCount" xml:"TotalCount"`
	PageSize    string          `json:"PageSize" xml:"PageSize"`
	RequestId   string          `json:"RequestId" xml:"RequestId"`
	PageNumber  string          `json:"PageNumber" xml:"PageNumber"`
	DBClusterId string          `json:"DBClusterId" xml:"DBClusterId"`
	Items       []SlowLogRecord `json:"Items" xml:"Items"`
}

// CreateDescribeAuditLogRecordsRequest creates a request to invoke DescribeAuditLogRecords API
func CreateDescribeAuditLogRecordsRequest() (request *DescribeAuditLogRecordsRequest) {
	request = &DescribeAuditLogRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribeAuditLogRecords", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAuditLogRecordsResponse creates a response to parse from DescribeAuditLogRecords response
func CreateDescribeAuditLogRecordsResponse() (response *DescribeAuditLogRecordsResponse) {
	response = &DescribeAuditLogRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}