package ddoscoo

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

// DescribeDefenseRecords invokes the ddoscoo.DescribeDefenseRecords API synchronously
func (client *Client) DescribeDefenseRecords(request *DescribeDefenseRecordsRequest) (response *DescribeDefenseRecordsResponse, err error) {
	response = CreateDescribeDefenseRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDefenseRecordsWithChan invokes the ddoscoo.DescribeDefenseRecords API asynchronously
func (client *Client) DescribeDefenseRecordsWithChan(request *DescribeDefenseRecordsRequest) (<-chan *DescribeDefenseRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeDefenseRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDefenseRecords(request)
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

// DescribeDefenseRecordsWithCallback invokes the ddoscoo.DescribeDefenseRecords API asynchronously
func (client *Client) DescribeDefenseRecordsWithCallback(request *DescribeDefenseRecordsRequest, callback func(response *DescribeDefenseRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDefenseRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDefenseRecords(request)
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

// DescribeDefenseRecordsRequest is the request struct for api DescribeDefenseRecords
type DescribeDefenseRecordsRequest struct {
	*requests.RpcRequest
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
}

// DescribeDefenseRecordsResponse is the response struct for api DescribeDefenseRecords
type DescribeDefenseRecordsResponse struct {
	*responses.BaseResponse
	TotalCount     int64           `json:"TotalCount" xml:"TotalCount"`
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	DefenseRecords []DefenseRecord `json:"DefenseRecords" xml:"DefenseRecords"`
}

// CreateDescribeDefenseRecordsRequest creates a request to invoke DescribeDefenseRecords API
func CreateDescribeDefenseRecordsRequest() (request *DescribeDefenseRecordsRequest) {
	request = &DescribeDefenseRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeDefenseRecords", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDefenseRecordsResponse creates a response to parse from DescribeDefenseRecords response
func CreateDescribeDefenseRecordsResponse() (response *DescribeDefenseRecordsResponse) {
	response = &DescribeDefenseRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
