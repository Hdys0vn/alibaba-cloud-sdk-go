package ecs

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

// DescribeDiagnosticReportAttributes invokes the ecs.DescribeDiagnosticReportAttributes API synchronously
func (client *Client) DescribeDiagnosticReportAttributes(request *DescribeDiagnosticReportAttributesRequest) (response *DescribeDiagnosticReportAttributesResponse, err error) {
	response = CreateDescribeDiagnosticReportAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDiagnosticReportAttributesWithChan invokes the ecs.DescribeDiagnosticReportAttributes API asynchronously
func (client *Client) DescribeDiagnosticReportAttributesWithChan(request *DescribeDiagnosticReportAttributesRequest) (<-chan *DescribeDiagnosticReportAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeDiagnosticReportAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDiagnosticReportAttributes(request)
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

// DescribeDiagnosticReportAttributesWithCallback invokes the ecs.DescribeDiagnosticReportAttributes API asynchronously
func (client *Client) DescribeDiagnosticReportAttributesWithCallback(request *DescribeDiagnosticReportAttributesRequest, callback func(response *DescribeDiagnosticReportAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDiagnosticReportAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDiagnosticReportAttributes(request)
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

// DescribeDiagnosticReportAttributesRequest is the request struct for api DescribeDiagnosticReportAttributes
type DescribeDiagnosticReportAttributesRequest struct {
	*requests.RpcRequest
	ReportId string `position:"Query" name:"ReportId"`
}

// DescribeDiagnosticReportAttributesResponse is the response struct for api DescribeDiagnosticReportAttributes
type DescribeDiagnosticReportAttributesResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ResourceId    string        `json:"ResourceId" xml:"ResourceId"`
	ResourceType  string        `json:"ResourceType" xml:"ResourceType"`
	ReportId      string        `json:"ReportId" xml:"ReportId"`
	Status        string        `json:"Status" xml:"Status"`
	CreationTime  string        `json:"CreationTime" xml:"CreationTime"`
	FinishedTime  string        `json:"FinishedTime" xml:"FinishedTime"`
	StartTime     string        `json:"StartTime" xml:"StartTime"`
	EndTime       string        `json:"EndTime" xml:"EndTime"`
	Severity      string        `json:"Severity" xml:"Severity"`
	MetricSetId   string        `json:"MetricSetId" xml:"MetricSetId"`
	Attributes    string        `json:"Attributes" xml:"Attributes"`
	MetricResults MetricResults `json:"MetricResults" xml:"MetricResults"`
}

// CreateDescribeDiagnosticReportAttributesRequest creates a request to invoke DescribeDiagnosticReportAttributes API
func CreateDescribeDiagnosticReportAttributesRequest() (request *DescribeDiagnosticReportAttributesRequest) {
	request = &DescribeDiagnosticReportAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDiagnosticReportAttributes", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDiagnosticReportAttributesResponse creates a response to parse from DescribeDiagnosticReportAttributes response
func CreateDescribeDiagnosticReportAttributesResponse() (response *DescribeDiagnosticReportAttributesResponse) {
	response = &DescribeDiagnosticReportAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
