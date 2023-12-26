package gpdb

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

// DescribeDownloadRecords invokes the gpdb.DescribeDownloadRecords API synchronously
func (client *Client) DescribeDownloadRecords(request *DescribeDownloadRecordsRequest) (response *DescribeDownloadRecordsResponse, err error) {
	response = CreateDescribeDownloadRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDownloadRecordsWithChan invokes the gpdb.DescribeDownloadRecords API asynchronously
func (client *Client) DescribeDownloadRecordsWithChan(request *DescribeDownloadRecordsRequest) (<-chan *DescribeDownloadRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeDownloadRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDownloadRecords(request)
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

// DescribeDownloadRecordsWithCallback invokes the gpdb.DescribeDownloadRecords API asynchronously
func (client *Client) DescribeDownloadRecordsWithCallback(request *DescribeDownloadRecordsRequest, callback func(response *DescribeDownloadRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDownloadRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDownloadRecords(request)
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

// DescribeDownloadRecordsRequest is the request struct for api DescribeDownloadRecords
type DescribeDownloadRecordsRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// DescribeDownloadRecordsResponse is the response struct for api DescribeDownloadRecords
type DescribeDownloadRecordsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Records   []Item `json:"Records" xml:"Records"`
}

// CreateDescribeDownloadRecordsRequest creates a request to invoke DescribeDownloadRecords API
func CreateDescribeDownloadRecordsRequest() (request *DescribeDownloadRecordsRequest) {
	request = &DescribeDownloadRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDownloadRecords", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDownloadRecordsResponse creates a response to parse from DescribeDownloadRecords response
func CreateDescribeDownloadRecordsResponse() (response *DescribeDownloadRecordsResponse) {
	response = &DescribeDownloadRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
