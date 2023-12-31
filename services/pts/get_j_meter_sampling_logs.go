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

// GetJMeterSamplingLogs invokes the pts.GetJMeterSamplingLogs API synchronously
func (client *Client) GetJMeterSamplingLogs(request *GetJMeterSamplingLogsRequest) (response *GetJMeterSamplingLogsResponse, err error) {
	response = CreateGetJMeterSamplingLogsResponse()
	err = client.DoAction(request, response)
	return
}

// GetJMeterSamplingLogsWithChan invokes the pts.GetJMeterSamplingLogs API asynchronously
func (client *Client) GetJMeterSamplingLogsWithChan(request *GetJMeterSamplingLogsRequest) (<-chan *GetJMeterSamplingLogsResponse, <-chan error) {
	responseChan := make(chan *GetJMeterSamplingLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJMeterSamplingLogs(request)
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

// GetJMeterSamplingLogsWithCallback invokes the pts.GetJMeterSamplingLogs API asynchronously
func (client *Client) GetJMeterSamplingLogsWithCallback(request *GetJMeterSamplingLogsRequest, callback func(response *GetJMeterSamplingLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJMeterSamplingLogsResponse
		var err error
		defer close(result)
		response, err = client.GetJMeterSamplingLogs(request)
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

// GetJMeterSamplingLogsRequest is the request struct for api GetJMeterSamplingLogs
type GetJMeterSamplingLogsRequest struct {
	*requests.RpcRequest
	ResponseCode string           `position:"Query" name:"ResponseCode"`
	AgentId      requests.Integer `position:"Query" name:"AgentId"`
	ReportId     string           `position:"Query" name:"ReportId"`
	MinRT        requests.Integer `position:"Query" name:"MinRT"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	BeginTime    requests.Integer `position:"Query" name:"BeginTime"`
	Thread       string           `position:"Query" name:"Thread"`
	MaxRT        requests.Integer `position:"Query" name:"MaxRT"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	SamplerId    requests.Integer `position:"Query" name:"SamplerId"`
	Success      requests.Boolean `position:"Query" name:"Success"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Keyword      string           `position:"Query" name:"Keyword"`
}

// GetJMeterSamplingLogsResponse is the response struct for api GetJMeterSamplingLogs
type GetJMeterSamplingLogsResponse struct {
	*responses.BaseResponse
	TotalCount     int64    `json:"TotalCount" xml:"TotalCount"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Message        string   `json:"Message" xml:"Message"`
	PageSize       int      `json:"PageSize" xml:"PageSize"`
	PageNumber     int      `json:"PageNumber" xml:"PageNumber"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string   `json:"Code" xml:"Code"`
	Success        bool     `json:"Success" xml:"Success"`
	SampleResults  []string `json:"SampleResults" xml:"SampleResults"`
}

// CreateGetJMeterSamplingLogsRequest creates a request to invoke GetJMeterSamplingLogs API
func CreateGetJMeterSamplingLogsRequest() (request *GetJMeterSamplingLogsRequest) {
	request = &GetJMeterSamplingLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "GetJMeterSamplingLogs", "", "")
	request.Method = requests.POST
	return
}

// CreateGetJMeterSamplingLogsResponse creates a response to parse from GetJMeterSamplingLogs response
func CreateGetJMeterSamplingLogsResponse() (response *GetJMeterSamplingLogsResponse) {
	response = &GetJMeterSamplingLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
