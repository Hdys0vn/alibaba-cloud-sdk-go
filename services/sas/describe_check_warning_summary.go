package sas

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

// DescribeCheckWarningSummary invokes the sas.DescribeCheckWarningSummary API synchronously
func (client *Client) DescribeCheckWarningSummary(request *DescribeCheckWarningSummaryRequest) (response *DescribeCheckWarningSummaryResponse, err error) {
	response = CreateDescribeCheckWarningSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCheckWarningSummaryWithChan invokes the sas.DescribeCheckWarningSummary API asynchronously
func (client *Client) DescribeCheckWarningSummaryWithChan(request *DescribeCheckWarningSummaryRequest) (<-chan *DescribeCheckWarningSummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeCheckWarningSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCheckWarningSummary(request)
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

// DescribeCheckWarningSummaryWithCallback invokes the sas.DescribeCheckWarningSummary API asynchronously
func (client *Client) DescribeCheckWarningSummaryWithCallback(request *DescribeCheckWarningSummaryRequest, callback func(response *DescribeCheckWarningSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCheckWarningSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeCheckWarningSummary(request)
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

// DescribeCheckWarningSummaryRequest is the request struct for api DescribeCheckWarningSummary
type DescribeCheckWarningSummaryRequest struct {
	*requests.RpcRequest
	TargetType          string           `position:"Query" name:"TargetType"`
	ContainerFieldName  string           `position:"Query" name:"ContainerFieldName"`
	RiskName            string           `position:"Query" name:"RiskName"`
	SourceIp            string           `position:"Query" name:"SourceIp"`
	ContainerFieldValue string           `position:"Query" name:"ContainerFieldValue"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
	Lang                string           `position:"Query" name:"Lang"`
	CurrentPage         requests.Integer `position:"Query" name:"CurrentPage"`
	ClusterId           string           `position:"Query" name:"ClusterId"`
	RiskStatus          requests.Integer `position:"Query" name:"RiskStatus"`
	StrategyId          requests.Integer `position:"Query" name:"StrategyId"`
	TypeName            string           `position:"Query" name:"TypeName"`
	Status              string           `position:"Query" name:"Status"`
	Uuids               string           `position:"Query" name:"Uuids"`
}

// DescribeCheckWarningSummaryResponse is the response struct for api DescribeCheckWarningSummary
type DescribeCheckWarningSummaryResponse struct {
	*responses.BaseResponse
	CurrentPage     int              `json:"CurrentPage" xml:"CurrentPage"`
	PageSize        int              `json:"PageSize" xml:"PageSize"`
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	TotalCount      int              `json:"TotalCount" xml:"TotalCount"`
	Count           int              `json:"Count" xml:"Count"`
	WarningSummarys []WarningSummary `json:"WarningSummarys" xml:"WarningSummarys"`
}

// CreateDescribeCheckWarningSummaryRequest creates a request to invoke DescribeCheckWarningSummary API
func CreateDescribeCheckWarningSummaryRequest() (request *DescribeCheckWarningSummaryRequest) {
	request = &DescribeCheckWarningSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeCheckWarningSummary", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCheckWarningSummaryResponse creates a response to parse from DescribeCheckWarningSummary response
func CreateDescribeCheckWarningSummaryResponse() (response *DescribeCheckWarningSummaryResponse) {
	response = &DescribeCheckWarningSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
