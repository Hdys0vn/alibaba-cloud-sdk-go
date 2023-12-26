package sas

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

// InstallCloudMonitor invokes the sas.InstallCloudMonitor API synchronously
func (client *Client) InstallCloudMonitor(request *InstallCloudMonitorRequest) (response *InstallCloudMonitorResponse, err error) {
	response = CreateInstallCloudMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// InstallCloudMonitorWithChan invokes the sas.InstallCloudMonitor API asynchronously
func (client *Client) InstallCloudMonitorWithChan(request *InstallCloudMonitorRequest) (<-chan *InstallCloudMonitorResponse, <-chan error) {
	responseChan := make(chan *InstallCloudMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallCloudMonitor(request)
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

// InstallCloudMonitorWithCallback invokes the sas.InstallCloudMonitor API asynchronously
func (client *Client) InstallCloudMonitorWithCallback(request *InstallCloudMonitorRequest, callback func(response *InstallCloudMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallCloudMonitorResponse
		var err error
		defer close(result)
		response, err = client.InstallCloudMonitor(request)
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

// InstallCloudMonitorRequest is the request struct for api InstallCloudMonitor
type InstallCloudMonitorRequest struct {
	*requests.RpcRequest
	AgentAccessKey string    `position:"Query" name:"AgentAccessKey"`
	AgentSecretKey string    `position:"Query" name:"AgentSecretKey"`
	SourceIp       string    `position:"Query" name:"SourceIp"`
	UuidList       *[]string `position:"Query" name:"UuidList"  type:"Repeated"`
	ArgusVersion   string    `position:"Query" name:"ArgusVersion"`
	InstanceIdList *[]string `position:"Query" name:"InstanceIdList"  type:"Repeated"`
}

// InstallCloudMonitorResponse is the response struct for api InstallCloudMonitor
type InstallCloudMonitorResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateInstallCloudMonitorRequest creates a request to invoke InstallCloudMonitor API
func CreateInstallCloudMonitorRequest() (request *InstallCloudMonitorRequest) {
	request = &InstallCloudMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "InstallCloudMonitor", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInstallCloudMonitorResponse creates a response to parse from InstallCloudMonitor response
func CreateInstallCloudMonitorResponse() (response *InstallCloudMonitorResponse) {
	response = &InstallCloudMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}