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

// BindPrometheusGrafanaInstance invokes the arms.BindPrometheusGrafanaInstance API synchronously
func (client *Client) BindPrometheusGrafanaInstance(request *BindPrometheusGrafanaInstanceRequest) (response *BindPrometheusGrafanaInstanceResponse, err error) {
	response = CreateBindPrometheusGrafanaInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// BindPrometheusGrafanaInstanceWithChan invokes the arms.BindPrometheusGrafanaInstance API asynchronously
func (client *Client) BindPrometheusGrafanaInstanceWithChan(request *BindPrometheusGrafanaInstanceRequest) (<-chan *BindPrometheusGrafanaInstanceResponse, <-chan error) {
	responseChan := make(chan *BindPrometheusGrafanaInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindPrometheusGrafanaInstance(request)
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

// BindPrometheusGrafanaInstanceWithCallback invokes the arms.BindPrometheusGrafanaInstance API asynchronously
func (client *Client) BindPrometheusGrafanaInstanceWithCallback(request *BindPrometheusGrafanaInstanceRequest, callback func(response *BindPrometheusGrafanaInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindPrometheusGrafanaInstanceResponse
		var err error
		defer close(result)
		response, err = client.BindPrometheusGrafanaInstance(request)
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

// BindPrometheusGrafanaInstanceRequest is the request struct for api BindPrometheusGrafanaInstance
type BindPrometheusGrafanaInstanceRequest struct {
	*requests.RpcRequest
	GrafanaInstanceId string `position:"Query" name:"GrafanaInstanceId"`
	ClusterId         string `position:"Query" name:"ClusterId"`
	ResourceGroupId   string `position:"Query" name:"ResourceGroupId"`
}

// BindPrometheusGrafanaInstanceResponse is the response struct for api BindPrometheusGrafanaInstance
type BindPrometheusGrafanaInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateBindPrometheusGrafanaInstanceRequest creates a request to invoke BindPrometheusGrafanaInstance API
func CreateBindPrometheusGrafanaInstanceRequest() (request *BindPrometheusGrafanaInstanceRequest) {
	request = &BindPrometheusGrafanaInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "BindPrometheusGrafanaInstance", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBindPrometheusGrafanaInstanceResponse creates a response to parse from BindPrometheusGrafanaInstance response
func CreateBindPrometheusGrafanaInstanceResponse() (response *BindPrometheusGrafanaInstanceResponse) {
	response = &BindPrometheusGrafanaInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
