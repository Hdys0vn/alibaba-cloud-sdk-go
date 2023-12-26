package hbase

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

// DescribeRestoreTables invokes the hbase.DescribeRestoreTables API synchronously
func (client *Client) DescribeRestoreTables(request *DescribeRestoreTablesRequest) (response *DescribeRestoreTablesResponse, err error) {
	response = CreateDescribeRestoreTablesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRestoreTablesWithChan invokes the hbase.DescribeRestoreTables API asynchronously
func (client *Client) DescribeRestoreTablesWithChan(request *DescribeRestoreTablesRequest) (<-chan *DescribeRestoreTablesResponse, <-chan error) {
	responseChan := make(chan *DescribeRestoreTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRestoreTables(request)
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

// DescribeRestoreTablesWithCallback invokes the hbase.DescribeRestoreTables API asynchronously
func (client *Client) DescribeRestoreTablesWithCallback(request *DescribeRestoreTablesRequest, callback func(response *DescribeRestoreTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRestoreTablesResponse
		var err error
		defer close(result)
		response, err = client.DescribeRestoreTables(request)
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

// DescribeRestoreTablesRequest is the request struct for api DescribeRestoreTables
type DescribeRestoreTablesRequest struct {
	*requests.RpcRequest
	ClusterId       string `position:"Query" name:"ClusterId"`
	RestoreRecordId string `position:"Query" name:"RestoreRecordId"`
}

// DescribeRestoreTablesResponse is the response struct for api DescribeRestoreTables
type DescribeRestoreTablesResponse struct {
	*responses.BaseResponse
	RequestId         string                        `json:"RequestId" xml:"RequestId"`
	Tables            TablesInDescribeRestoreTables `json:"Tables" xml:"Tables"`
	RestoreSummary    RestoreSummary                `json:"RestoreSummary" xml:"RestoreSummary"`
	RestoreSchema     RestoreSchema                 `json:"RestoreSchema" xml:"RestoreSchema"`
	RestoreFull       RestoreFull                   `json:"RestoreFull" xml:"RestoreFull"`
	RestoreIncrDetail RestoreIncrDetail             `json:"RestoreIncrDetail" xml:"RestoreIncrDetail"`
}

// CreateDescribeRestoreTablesRequest creates a request to invoke DescribeRestoreTables API
func CreateDescribeRestoreTablesRequest() (request *DescribeRestoreTablesRequest) {
	request = &DescribeRestoreTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeRestoreTables", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRestoreTablesResponse creates a response to parse from DescribeRestoreTables response
func CreateDescribeRestoreTablesResponse() (response *DescribeRestoreTablesResponse) {
	response = &DescribeRestoreTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
