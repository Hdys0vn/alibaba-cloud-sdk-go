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

// UpdatePrometheusMonitoring invokes the arms.UpdatePrometheusMonitoring API synchronously
func (client *Client) UpdatePrometheusMonitoring(request *UpdatePrometheusMonitoringRequest) (response *UpdatePrometheusMonitoringResponse, err error) {
	response = CreateUpdatePrometheusMonitoringResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePrometheusMonitoringWithChan invokes the arms.UpdatePrometheusMonitoring API asynchronously
func (client *Client) UpdatePrometheusMonitoringWithChan(request *UpdatePrometheusMonitoringRequest) (<-chan *UpdatePrometheusMonitoringResponse, <-chan error) {
	responseChan := make(chan *UpdatePrometheusMonitoringResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePrometheusMonitoring(request)
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

// UpdatePrometheusMonitoringWithCallback invokes the arms.UpdatePrometheusMonitoring API asynchronously
func (client *Client) UpdatePrometheusMonitoringWithCallback(request *UpdatePrometheusMonitoringRequest, callback func(response *UpdatePrometheusMonitoringResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePrometheusMonitoringResponse
		var err error
		defer close(result)
		response, err = client.UpdatePrometheusMonitoring(request)
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

// UpdatePrometheusMonitoringRequest is the request struct for api UpdatePrometheusMonitoring
type UpdatePrometheusMonitoringRequest struct {
	*requests.RpcRequest
	ConfigYaml     string `position:"Body" name:"ConfigYaml"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	Type           string `position:"Query" name:"Type"`
	MonitoringName string `position:"Query" name:"MonitoringName"`
}

// UpdatePrometheusMonitoringResponse is the response struct for api UpdatePrometheusMonitoring
type UpdatePrometheusMonitoringResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdatePrometheusMonitoringRequest creates a request to invoke UpdatePrometheusMonitoring API
func CreateUpdatePrometheusMonitoringRequest() (request *UpdatePrometheusMonitoringRequest) {
	request = &UpdatePrometheusMonitoringRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdatePrometheusMonitoring", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePrometheusMonitoringResponse creates a response to parse from UpdatePrometheusMonitoring response
func CreateUpdatePrometheusMonitoringResponse() (response *UpdatePrometheusMonitoringResponse) {
	response = &UpdatePrometheusMonitoringResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}