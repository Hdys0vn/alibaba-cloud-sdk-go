package airec

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

// DescribeSyncReportDetail invokes the airec.DescribeSyncReportDetail API synchronously
func (client *Client) DescribeSyncReportDetail(request *DescribeSyncReportDetailRequest) (response *DescribeSyncReportDetailResponse, err error) {
	response = CreateDescribeSyncReportDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSyncReportDetailWithChan invokes the airec.DescribeSyncReportDetail API asynchronously
func (client *Client) DescribeSyncReportDetailWithChan(request *DescribeSyncReportDetailRequest) (<-chan *DescribeSyncReportDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeSyncReportDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSyncReportDetail(request)
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

// DescribeSyncReportDetailWithCallback invokes the airec.DescribeSyncReportDetail API asynchronously
func (client *Client) DescribeSyncReportDetailWithCallback(request *DescribeSyncReportDetailRequest, callback func(response *DescribeSyncReportDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSyncReportDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeSyncReportDetail(request)
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

// DescribeSyncReportDetailRequest is the request struct for api DescribeSyncReportDetail
type DescribeSyncReportDetailRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"instanceId"`
	LevelType  string           `position:"Query" name:"levelType"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	StartTime  requests.Integer `position:"Query" name:"startTime"`
	Type       string           `position:"Query" name:"type"`
}

// DescribeSyncReportDetailResponse is the response struct for api DescribeSyncReportDetail
type DescribeSyncReportDetailResponse struct {
	*responses.BaseResponse
	Code      string       `json:"code" xml:"code"`
	Message   string       `json:"message" xml:"message"`
	RequestId string       `json:"requestId" xml:"requestId"`
	Result    []ResultItem `json:"result" xml:"result"`
}

// CreateDescribeSyncReportDetailRequest creates a request to invoke DescribeSyncReportDetail API
func CreateDescribeSyncReportDetailRequest() (request *DescribeSyncReportDetailRequest) {
	request = &DescribeSyncReportDetailRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "DescribeSyncReportDetail", "/v2/openapi/instances/[instanceId]/sync-reports/detail", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeSyncReportDetailResponse creates a response to parse from DescribeSyncReportDetail response
func CreateDescribeSyncReportDetailResponse() (response *DescribeSyncReportDetailResponse) {
	response = &DescribeSyncReportDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
