package alidns

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

// DescribePdnsThreatLogs invokes the alidns.DescribePdnsThreatLogs API synchronously
func (client *Client) DescribePdnsThreatLogs(request *DescribePdnsThreatLogsRequest) (response *DescribePdnsThreatLogsResponse, err error) {
	response = CreateDescribePdnsThreatLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePdnsThreatLogsWithChan invokes the alidns.DescribePdnsThreatLogs API asynchronously
func (client *Client) DescribePdnsThreatLogsWithChan(request *DescribePdnsThreatLogsRequest) (<-chan *DescribePdnsThreatLogsResponse, <-chan error) {
	responseChan := make(chan *DescribePdnsThreatLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePdnsThreatLogs(request)
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

// DescribePdnsThreatLogsWithCallback invokes the alidns.DescribePdnsThreatLogs API asynchronously
func (client *Client) DescribePdnsThreatLogsWithCallback(request *DescribePdnsThreatLogsRequest, callback func(response *DescribePdnsThreatLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePdnsThreatLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribePdnsThreatLogs(request)
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

// DescribePdnsThreatLogsRequest is the request struct for api DescribePdnsThreatLogs
type DescribePdnsThreatLogsRequest struct {
	*requests.RpcRequest
	ThreatSourceIp string           `position:"Query" name:"ThreatSourceIp"`
	StartDate      string           `position:"Query" name:"StartDate"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	EndDate        string           `position:"Query" name:"EndDate"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	ThreatType     string           `position:"Query" name:"ThreatType"`
	Lang           string           `position:"Query" name:"Lang"`
	Keyword        string           `position:"Query" name:"Keyword"`
	ThreatLevel    string           `position:"Query" name:"ThreatLevel"`
}

// DescribePdnsThreatLogsResponse is the response struct for api DescribePdnsThreatLogs
type DescribePdnsThreatLogsResponse struct {
	*responses.BaseResponse
	TotalCount int64  `json:"TotalCount" xml:"TotalCount"`
	PageSize   int64  `json:"PageSize" xml:"PageSize"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PageNumber int64  `json:"PageNumber" xml:"PageNumber"`
	Logs       []Log  `json:"Logs" xml:"Logs"`
}

// CreateDescribePdnsThreatLogsRequest creates a request to invoke DescribePdnsThreatLogs API
func CreateDescribePdnsThreatLogsRequest() (request *DescribePdnsThreatLogsRequest) {
	request = &DescribePdnsThreatLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribePdnsThreatLogs", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePdnsThreatLogsResponse creates a response to parse from DescribePdnsThreatLogs response
func CreateDescribePdnsThreatLogsResponse() (response *DescribePdnsThreatLogsResponse) {
	response = &DescribePdnsThreatLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
