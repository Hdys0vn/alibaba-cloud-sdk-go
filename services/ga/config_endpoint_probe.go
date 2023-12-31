package ga

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

// ConfigEndpointProbe invokes the ga.ConfigEndpointProbe API synchronously
func (client *Client) ConfigEndpointProbe(request *ConfigEndpointProbeRequest) (response *ConfigEndpointProbeResponse, err error) {
	response = CreateConfigEndpointProbeResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigEndpointProbeWithChan invokes the ga.ConfigEndpointProbe API asynchronously
func (client *Client) ConfigEndpointProbeWithChan(request *ConfigEndpointProbeRequest) (<-chan *ConfigEndpointProbeResponse, <-chan error) {
	responseChan := make(chan *ConfigEndpointProbeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigEndpointProbe(request)
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

// ConfigEndpointProbeWithCallback invokes the ga.ConfigEndpointProbe API asynchronously
func (client *Client) ConfigEndpointProbeWithCallback(request *ConfigEndpointProbeRequest, callback func(response *ConfigEndpointProbeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigEndpointProbeResponse
		var err error
		defer close(result)
		response, err = client.ConfigEndpointProbe(request)
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

// ConfigEndpointProbeRequest is the request struct for api ConfigEndpointProbe
type ConfigEndpointProbeRequest struct {
	*requests.RpcRequest
	ClientToken     string `position:"Query" name:"ClientToken"`
	Endpoint        string `position:"Query" name:"Endpoint"`
	EndpointType    string `position:"Query" name:"EndpointType"`
	Enable          string `position:"Query" name:"Enable"`
	ProbeProtocol   string `position:"Query" name:"ProbeProtocol"`
	ProbePort       string `position:"Query" name:"ProbePort"`
	EndpointGroupId string `position:"Query" name:"EndpointGroupId"`
}

// ConfigEndpointProbeResponse is the response struct for api ConfigEndpointProbe
type ConfigEndpointProbeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigEndpointProbeRequest creates a request to invoke ConfigEndpointProbe API
func CreateConfigEndpointProbeRequest() (request *ConfigEndpointProbeRequest) {
	request = &ConfigEndpointProbeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "ConfigEndpointProbe", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigEndpointProbeResponse creates a response to parse from ConfigEndpointProbe response
func CreateConfigEndpointProbeResponse() (response *ConfigEndpointProbeResponse) {
	response = &ConfigEndpointProbeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
