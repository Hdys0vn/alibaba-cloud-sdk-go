package edas

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

// InstallAgent invokes the edas.InstallAgent API synchronously
func (client *Client) InstallAgent(request *InstallAgentRequest) (response *InstallAgentResponse, err error) {
	response = CreateInstallAgentResponse()
	err = client.DoAction(request, response)
	return
}

// InstallAgentWithChan invokes the edas.InstallAgent API asynchronously
func (client *Client) InstallAgentWithChan(request *InstallAgentRequest) (<-chan *InstallAgentResponse, <-chan error) {
	responseChan := make(chan *InstallAgentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallAgent(request)
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

// InstallAgentWithCallback invokes the edas.InstallAgent API asynchronously
func (client *Client) InstallAgentWithCallback(request *InstallAgentRequest, callback func(response *InstallAgentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallAgentResponse
		var err error
		defer close(result)
		response, err = client.InstallAgent(request)
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

// InstallAgentRequest is the request struct for api InstallAgent
type InstallAgentRequest struct {
	*requests.RoaRequest
	InstanceIds string           `position:"Query" name:"InstanceIds"`
	DoAsync     requests.Boolean `position:"Query" name:"DoAsync"`
	ClusterId   string           `position:"Query" name:"ClusterId"`
}

// InstallAgentResponse is the response struct for api InstallAgent
type InstallAgentResponse struct {
	*responses.BaseResponse
	Code                int                 `json:"Code" xml:"Code"`
	Message             string              `json:"Message" xml:"Message"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	ExecutionResultList ExecutionResultList `json:"ExecutionResultList" xml:"ExecutionResultList"`
}

// CreateInstallAgentRequest creates a request to invoke InstallAgent API
func CreateInstallAgentRequest() (request *InstallAgentRequest) {
	request = &InstallAgentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "InstallAgent", "/pop/v5/ecss/install_agent", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInstallAgentResponse creates a response to parse from InstallAgent response
func CreateInstallAgentResponse() (response *InstallAgentResponse) {
	response = &InstallAgentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
