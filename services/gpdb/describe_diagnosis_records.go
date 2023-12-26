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

// DescribeDiagnosisRecords invokes the gpdb.DescribeDiagnosisRecords API synchronously
func (client *Client) DescribeDiagnosisRecords(request *DescribeDiagnosisRecordsRequest) (response *DescribeDiagnosisRecordsResponse, err error) {
	response = CreateDescribeDiagnosisRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDiagnosisRecordsWithChan invokes the gpdb.DescribeDiagnosisRecords API asynchronously
func (client *Client) DescribeDiagnosisRecordsWithChan(request *DescribeDiagnosisRecordsRequest) (<-chan *DescribeDiagnosisRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeDiagnosisRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDiagnosisRecords(request)
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

// DescribeDiagnosisRecordsWithCallback invokes the gpdb.DescribeDiagnosisRecords API asynchronously
func (client *Client) DescribeDiagnosisRecordsWithCallback(request *DescribeDiagnosisRecordsRequest, callback func(response *DescribeDiagnosisRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDiagnosisRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDiagnosisRecords(request)
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

// DescribeDiagnosisRecordsRequest is the request struct for api DescribeDiagnosisRecords
type DescribeDiagnosisRecordsRequest struct {
	*requests.RpcRequest
	QueryCondition string           `position:"Query" name:"QueryCondition"`
	StartTime      string           `position:"Query" name:"StartTime"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	Database       string           `position:"Query" name:"Database"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId   string           `position:"Query" name:"DBInstanceId"`
	Keyword        string           `position:"Query" name:"Keyword"`
	Order          string           `position:"Query" name:"Order"`
	EndTime        string           `position:"Query" name:"EndTime"`
	User           string           `position:"Query" name:"User"`
}

// DescribeDiagnosisRecordsResponse is the response struct for api DescribeDiagnosisRecords
type DescribeDiagnosisRecordsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	Items      []Item `json:"Items" xml:"Items"`
}

// CreateDescribeDiagnosisRecordsRequest creates a request to invoke DescribeDiagnosisRecords API
func CreateDescribeDiagnosisRecordsRequest() (request *DescribeDiagnosisRecordsRequest) {
	request = &DescribeDiagnosisRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDiagnosisRecords", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDiagnosisRecordsResponse creates a response to parse from DescribeDiagnosisRecords response
func CreateDescribeDiagnosisRecordsResponse() (response *DescribeDiagnosisRecordsResponse) {
	response = &DescribeDiagnosisRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
