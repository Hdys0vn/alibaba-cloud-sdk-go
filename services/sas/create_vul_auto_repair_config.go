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

// CreateVulAutoRepairConfig invokes the sas.CreateVulAutoRepairConfig API synchronously
func (client *Client) CreateVulAutoRepairConfig(request *CreateVulAutoRepairConfigRequest) (response *CreateVulAutoRepairConfigResponse, err error) {
	response = CreateCreateVulAutoRepairConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVulAutoRepairConfigWithChan invokes the sas.CreateVulAutoRepairConfig API asynchronously
func (client *Client) CreateVulAutoRepairConfigWithChan(request *CreateVulAutoRepairConfigRequest) (<-chan *CreateVulAutoRepairConfigResponse, <-chan error) {
	responseChan := make(chan *CreateVulAutoRepairConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVulAutoRepairConfig(request)
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

// CreateVulAutoRepairConfigWithCallback invokes the sas.CreateVulAutoRepairConfig API asynchronously
func (client *Client) CreateVulAutoRepairConfigWithCallback(request *CreateVulAutoRepairConfigRequest, callback func(response *CreateVulAutoRepairConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVulAutoRepairConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateVulAutoRepairConfig(request)
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

// CreateVulAutoRepairConfigRequest is the request struct for api CreateVulAutoRepairConfig
type CreateVulAutoRepairConfigRequest struct {
	*requests.RpcRequest
	Reason                  string                                              `position:"Query" name:"Reason"`
	Type                    string                                              `position:"Query" name:"Type"`
	SourceIp                string                                              `position:"Query" name:"SourceIp"`
	VulAutoRepairConfigList *[]CreateVulAutoRepairConfigVulAutoRepairConfigList `position:"Query" name:"VulAutoRepairConfigList"  type:"Repeated"`
}

// CreateVulAutoRepairConfigVulAutoRepairConfigList is a repeated param struct in CreateVulAutoRepairConfigRequest
type CreateVulAutoRepairConfigVulAutoRepairConfigList struct {
	AliasName string `name:"AliasName"`
	Name      string `name:"Name"`
}

// CreateVulAutoRepairConfigResponse is the response struct for api CreateVulAutoRepairConfig
type CreateVulAutoRepairConfigResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreateVulAutoRepairConfigRequest creates a request to invoke CreateVulAutoRepairConfig API
func CreateCreateVulAutoRepairConfigRequest() (request *CreateVulAutoRepairConfigRequest) {
	request = &CreateVulAutoRepairConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "CreateVulAutoRepairConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVulAutoRepairConfigResponse creates a response to parse from CreateVulAutoRepairConfig response
func CreateCreateVulAutoRepairConfigResponse() (response *CreateVulAutoRepairConfigResponse) {
	response = &CreateVulAutoRepairConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
