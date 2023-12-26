package rds

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

// DescribeSQLLogReportList invokes the rds.DescribeSQLLogReportList API synchronously
func (client *Client) DescribeSQLLogReportList(request *DescribeSQLLogReportListRequest) (response *DescribeSQLLogReportListResponse, err error) {
	response = CreateDescribeSQLLogReportListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLLogReportListWithChan invokes the rds.DescribeSQLLogReportList API asynchronously
func (client *Client) DescribeSQLLogReportListWithChan(request *DescribeSQLLogReportListRequest) (<-chan *DescribeSQLLogReportListResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLLogReportListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLLogReportList(request)
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

// DescribeSQLLogReportListWithCallback invokes the rds.DescribeSQLLogReportList API asynchronously
func (client *Client) DescribeSQLLogReportListWithCallback(request *DescribeSQLLogReportListRequest, callback func(response *DescribeSQLLogReportListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLLogReportListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLLogReportList(request)
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

// DescribeSQLLogReportListRequest is the request struct for api DescribeSQLLogReportList
type DescribeSQLLogReportListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSQLLogReportListResponse is the response struct for api DescribeSQLLogReportList
type DescribeSQLLogReportListResponse struct {
	*responses.BaseResponse
	RequestId        string                          `json:"RequestId" xml:"RequestId"`
	PageNumber       int                             `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                             `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                             `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            ItemsInDescribeSQLLogReportList `json:"Items" xml:"Items"`
}

// CreateDescribeSQLLogReportListRequest creates a request to invoke DescribeSQLLogReportList API
func CreateDescribeSQLLogReportListRequest() (request *DescribeSQLLogReportListRequest) {
	request = &DescribeSQLLogReportListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogReportList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSQLLogReportListResponse creates a response to parse from DescribeSQLLogReportList response
func CreateDescribeSQLLogReportListResponse() (response *DescribeSQLLogReportListResponse) {
	response = &DescribeSQLLogReportListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
