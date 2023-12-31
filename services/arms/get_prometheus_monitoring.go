package arms

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

// GetPrometheusMonitoring invokes the arms.GetPrometheusMonitoring API synchronously
func (client *Client) GetPrometheusMonitoring(request *GetPrometheusMonitoringRequest) (response *GetPrometheusMonitoringResponse, err error) {
	response = CreateGetPrometheusMonitoringResponse()
	err = client.DoAction(request, response)
	return
}

// GetPrometheusMonitoringWithChan invokes the arms.GetPrometheusMonitoring API asynchronously
func (client *Client) GetPrometheusMonitoringWithChan(request *GetPrometheusMonitoringRequest) (<-chan *GetPrometheusMonitoringResponse, <-chan error) {
	responseChan := make(chan *GetPrometheusMonitoringResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPrometheusMonitoring(request)
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

// GetPrometheusMonitoringWithCallback invokes the arms.GetPrometheusMonitoring API asynchronously
func (client *Client) GetPrometheusMonitoringWithCallback(request *GetPrometheusMonitoringRequest, callback func(response *GetPrometheusMonitoringResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPrometheusMonitoringResponse
		var err error
		defer close(result)
		response, err = client.GetPrometheusMonitoring(request)
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

// GetPrometheusMonitoringRequest is the request struct for api GetPrometheusMonitoring
type GetPrometheusMonitoringRequest struct {
	*requests.RpcRequest
	ClusterId      string `position:"Query" name:"ClusterId"`
	Type           string `position:"Query" name:"Type"`
	MonitoringName string `position:"Query" name:"MonitoringName"`
}

// GetPrometheusMonitoringResponse is the response struct for api GetPrometheusMonitoring
type GetPrometheusMonitoringResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetPrometheusMonitoringRequest creates a request to invoke GetPrometheusMonitoring API
func CreateGetPrometheusMonitoringRequest() (request *GetPrometheusMonitoringRequest) {
	request = &GetPrometheusMonitoringRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetPrometheusMonitoring", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPrometheusMonitoringResponse creates a response to parse from GetPrometheusMonitoring response
func CreateGetPrometheusMonitoringResponse() (response *GetPrometheusMonitoringResponse) {
	response = &GetPrometheusMonitoringResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
