package pts

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

// ListJMeterReports invokes the pts.ListJMeterReports API synchronously
func (client *Client) ListJMeterReports(request *ListJMeterReportsRequest) (response *ListJMeterReportsResponse, err error) {
	response = CreateListJMeterReportsResponse()
	err = client.DoAction(request, response)
	return
}

// ListJMeterReportsWithChan invokes the pts.ListJMeterReports API asynchronously
func (client *Client) ListJMeterReportsWithChan(request *ListJMeterReportsRequest) (<-chan *ListJMeterReportsResponse, <-chan error) {
	responseChan := make(chan *ListJMeterReportsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJMeterReports(request)
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

// ListJMeterReportsWithCallback invokes the pts.ListJMeterReports API asynchronously
func (client *Client) ListJMeterReportsWithCallback(request *ListJMeterReportsRequest, callback func(response *ListJMeterReportsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJMeterReportsResponse
		var err error
		defer close(result)
		response, err = client.ListJMeterReports(request)
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

// ListJMeterReportsRequest is the request struct for api ListJMeterReports
type ListJMeterReportsRequest struct {
	*requests.RpcRequest
	ReportId   string           `position:"Query" name:"ReportId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	BeginTime  requests.Integer `position:"Query" name:"BeginTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	SceneId    string           `position:"Query" name:"SceneId"`
	Keyword    string           `position:"Query" name:"Keyword"`
}

// ListJMeterReportsResponse is the response struct for api ListJMeterReports
type ListJMeterReportsResponse struct {
	*responses.BaseResponse
	TotalCount     int64              `json:"TotalCount" xml:"TotalCount"`
	RequestId      string             `json:"RequestId" xml:"RequestId"`
	Message        string             `json:"Message" xml:"Message"`
	PageSize       int                `json:"PageSize" xml:"PageSize"`
	PageNumber     int                `json:"PageNumber" xml:"PageNumber"`
	HttpStatusCode int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string             `json:"Code" xml:"Code"`
	Success        bool               `json:"Success" xml:"Success"`
	Reports        []JMeterReportView `json:"Reports" xml:"Reports"`
}

// CreateListJMeterReportsRequest creates a request to invoke ListJMeterReports API
func CreateListJMeterReportsRequest() (request *ListJMeterReportsRequest) {
	request = &ListJMeterReportsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "ListJMeterReports", "", "")
	request.Method = requests.POST
	return
}

// CreateListJMeterReportsResponse creates a response to parse from ListJMeterReports response
func CreateListJMeterReportsResponse() (response *ListJMeterReportsResponse) {
	response = &ListJMeterReportsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
