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

// UpdatePrometheusMonitoringStatus invokes the arms.UpdatePrometheusMonitoringStatus API synchronously
func (client *Client) UpdatePrometheusMonitoringStatus(request *UpdatePrometheusMonitoringStatusRequest) (response *UpdatePrometheusMonitoringStatusResponse, err error) {
	response = CreateUpdatePrometheusMonitoringStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePrometheusMonitoringStatusWithChan invokes the arms.UpdatePrometheusMonitoringStatus API asynchronously
func (client *Client) UpdatePrometheusMonitoringStatusWithChan(request *UpdatePrometheusMonitoringStatusRequest) (<-chan *UpdatePrometheusMonitoringStatusResponse, <-chan error) {
	responseChan := make(chan *UpdatePrometheusMonitoringStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePrometheusMonitoringStatus(request)
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

// UpdatePrometheusMonitoringStatusWithCallback invokes the arms.UpdatePrometheusMonitoringStatus API asynchronously
func (client *Client) UpdatePrometheusMonitoringStatusWithCallback(request *UpdatePrometheusMonitoringStatusRequest, callback func(response *UpdatePrometheusMonitoringStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePrometheusMonitoringStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdatePrometheusMonitoringStatus(request)
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

// UpdatePrometheusMonitoringStatusRequest is the request struct for api UpdatePrometheusMonitoringStatus
type UpdatePrometheusMonitoringStatusRequest struct {
	*requests.RpcRequest
	ClusterId      string `position:"Query" name:"ClusterId"`
	Type           string `position:"Query" name:"Type"`
	MonitoringName string `position:"Query" name:"MonitoringName"`
	Status         string `position:"Query" name:"Status"`
}

// UpdatePrometheusMonitoringStatusResponse is the response struct for api UpdatePrometheusMonitoringStatus
type UpdatePrometheusMonitoringStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdatePrometheusMonitoringStatusRequest creates a request to invoke UpdatePrometheusMonitoringStatus API
func CreateUpdatePrometheusMonitoringStatusRequest() (request *UpdatePrometheusMonitoringStatusRequest) {
	request = &UpdatePrometheusMonitoringStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdatePrometheusMonitoringStatus", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePrometheusMonitoringStatusResponse creates a response to parse from UpdatePrometheusMonitoringStatus response
func CreateUpdatePrometheusMonitoringStatusResponse() (response *UpdatePrometheusMonitoringStatusResponse) {
	response = &UpdatePrometheusMonitoringStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
