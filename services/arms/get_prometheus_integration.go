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

// GetPrometheusIntegration invokes the arms.GetPrometheusIntegration API synchronously
func (client *Client) GetPrometheusIntegration(request *GetPrometheusIntegrationRequest) (response *GetPrometheusIntegrationResponse, err error) {
	response = CreateGetPrometheusIntegrationResponse()
	err = client.DoAction(request, response)
	return
}

// GetPrometheusIntegrationWithChan invokes the arms.GetPrometheusIntegration API asynchronously
func (client *Client) GetPrometheusIntegrationWithChan(request *GetPrometheusIntegrationRequest) (<-chan *GetPrometheusIntegrationResponse, <-chan error) {
	responseChan := make(chan *GetPrometheusIntegrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPrometheusIntegration(request)
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

// GetPrometheusIntegrationWithCallback invokes the arms.GetPrometheusIntegration API asynchronously
func (client *Client) GetPrometheusIntegrationWithCallback(request *GetPrometheusIntegrationRequest, callback func(response *GetPrometheusIntegrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPrometheusIntegrationResponse
		var err error
		defer close(result)
		response, err = client.GetPrometheusIntegration(request)
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

// GetPrometheusIntegrationRequest is the request struct for api GetPrometheusIntegration
type GetPrometheusIntegrationRequest struct {
	*requests.RpcRequest
	InstanceId      requests.Integer `position:"Query" name:"InstanceId"`
	IntegrationType string           `position:"Query" name:"IntegrationType"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
}

// GetPrometheusIntegrationResponse is the response struct for api GetPrometheusIntegration
type GetPrometheusIntegrationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetPrometheusIntegrationRequest creates a request to invoke GetPrometheusIntegration API
func CreateGetPrometheusIntegrationRequest() (request *GetPrometheusIntegrationRequest) {
	request = &GetPrometheusIntegrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetPrometheusIntegration", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPrometheusIntegrationResponse creates a response to parse from GetPrometheusIntegration response
func CreateGetPrometheusIntegrationResponse() (response *GetPrometheusIntegrationResponse) {
	response = &GetPrometheusIntegrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}