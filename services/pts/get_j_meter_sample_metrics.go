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

// GetJMeterSampleMetrics invokes the pts.GetJMeterSampleMetrics API synchronously
func (client *Client) GetJMeterSampleMetrics(request *GetJMeterSampleMetricsRequest) (response *GetJMeterSampleMetricsResponse, err error) {
	response = CreateGetJMeterSampleMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// GetJMeterSampleMetricsWithChan invokes the pts.GetJMeterSampleMetrics API asynchronously
func (client *Client) GetJMeterSampleMetricsWithChan(request *GetJMeterSampleMetricsRequest) (<-chan *GetJMeterSampleMetricsResponse, <-chan error) {
	responseChan := make(chan *GetJMeterSampleMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJMeterSampleMetrics(request)
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

// GetJMeterSampleMetricsWithCallback invokes the pts.GetJMeterSampleMetrics API asynchronously
func (client *Client) GetJMeterSampleMetricsWithCallback(request *GetJMeterSampleMetricsRequest, callback func(response *GetJMeterSampleMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJMeterSampleMetricsResponse
		var err error
		defer close(result)
		response, err = client.GetJMeterSampleMetrics(request)
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

// GetJMeterSampleMetricsRequest is the request struct for api GetJMeterSampleMetrics
type GetJMeterSampleMetricsRequest struct {
	*requests.RpcRequest
	ReportId  string           `position:"Query" name:"ReportId"`
	SamplerId requests.Integer `position:"Query" name:"SamplerId"`
	EndTime   requests.Integer `position:"Query" name:"EndTime"`
	BeginTime requests.Integer `position:"Query" name:"BeginTime"`
}

// GetJMeterSampleMetricsResponse is the response struct for api GetJMeterSampleMetrics
type GetJMeterSampleMetricsResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	Message          string                 `json:"Message" xml:"Message"`
	SamplerMap       map[string]interface{} `json:"SamplerMap" xml:"SamplerMap"`
	Code             string                 `json:"Code" xml:"Code"`
	Success          bool                   `json:"Success" xml:"Success"`
	SampleMetricList []string               `json:"SampleMetricList" xml:"SampleMetricList"`
}

// CreateGetJMeterSampleMetricsRequest creates a request to invoke GetJMeterSampleMetrics API
func CreateGetJMeterSampleMetricsRequest() (request *GetJMeterSampleMetricsRequest) {
	request = &GetJMeterSampleMetricsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "GetJMeterSampleMetrics", "", "")
	request.Method = requests.POST
	return
}

// CreateGetJMeterSampleMetricsResponse creates a response to parse from GetJMeterSampleMetrics response
func CreateGetJMeterSampleMetricsResponse() (response *GetJMeterSampleMetricsResponse) {
	response = &GetJMeterSampleMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
