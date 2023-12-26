package privatelink

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

// CreateVpcEndpointService invokes the privatelink.CreateVpcEndpointService API synchronously
func (client *Client) CreateVpcEndpointService(request *CreateVpcEndpointServiceRequest) (response *CreateVpcEndpointServiceResponse, err error) {
	response = CreateCreateVpcEndpointServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVpcEndpointServiceWithChan invokes the privatelink.CreateVpcEndpointService API asynchronously
func (client *Client) CreateVpcEndpointServiceWithChan(request *CreateVpcEndpointServiceRequest) (<-chan *CreateVpcEndpointServiceResponse, <-chan error) {
	responseChan := make(chan *CreateVpcEndpointServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpcEndpointService(request)
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

// CreateVpcEndpointServiceWithCallback invokes the privatelink.CreateVpcEndpointService API asynchronously
func (client *Client) CreateVpcEndpointServiceWithCallback(request *CreateVpcEndpointServiceRequest, callback func(response *CreateVpcEndpointServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpcEndpointServiceResponse
		var err error
		defer close(result)
		response, err = client.CreateVpcEndpointService(request)
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

// CreateVpcEndpointServiceRequest is the request struct for api CreateVpcEndpointService
type CreateVpcEndpointServiceRequest struct {
	*requests.RpcRequest
	PrivateServiceDomainEnabled requests.Boolean                     `position:"Query" name:"PrivateServiceDomainEnabled"`
	PrivateServiceDomain        string                               `position:"Query" name:"PrivateServiceDomain"`
	AutoAcceptEnabled           requests.Boolean                     `position:"Query" name:"AutoAcceptEnabled"`
	ClientToken                 string                               `position:"Query" name:"ClientToken"`
	SystemTag                   *[]CreateVpcEndpointServiceSystemTag `position:"Query" name:"SystemTag"  type:"Repeated"`
	Payer                       string                               `position:"Query" name:"Payer"`
	ResourceGroupId             string                               `position:"Query" name:"ResourceGroupId"`
	ZoneAffinityEnabled         requests.Boolean                     `position:"Query" name:"ZoneAffinityEnabled"`
	Tag                         *[]CreateVpcEndpointServiceTag       `position:"Query" name:"Tag"  type:"Repeated"`
	DryRun                      requests.Boolean                     `position:"Query" name:"DryRun"`
	Resource                    *[]CreateVpcEndpointServiceResource  `position:"Query" name:"Resource"  type:"Repeated"`
	ServiceResourceType         string                               `position:"Query" name:"ServiceResourceType"`
	ServiceSupportIPv6          requests.Boolean                     `position:"Query" name:"ServiceSupportIPv6"`
	ServiceDescription          string                               `position:"Query" name:"ServiceDescription"`
}

// CreateVpcEndpointServiceSystemTag is a repeated param struct in CreateVpcEndpointServiceRequest
type CreateVpcEndpointServiceSystemTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Scope string `name:"Scope"`
}

// CreateVpcEndpointServiceTag is a repeated param struct in CreateVpcEndpointServiceRequest
type CreateVpcEndpointServiceTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateVpcEndpointServiceResource is a repeated param struct in CreateVpcEndpointServiceRequest
type CreateVpcEndpointServiceResource struct {
	ResourceType string `name:"ResourceType"`
	ResourceId   string `name:"ResourceId"`
	ZoneId       string `name:"ZoneId"`
}

// CreateVpcEndpointServiceResponse is the response struct for api CreateVpcEndpointService
type CreateVpcEndpointServiceResponse struct {
	*responses.BaseResponse
	ServiceBusinessStatus string `json:"ServiceBusinessStatus" xml:"ServiceBusinessStatus"`
	RequestId             string `json:"RequestId" xml:"RequestId"`
	ServiceName           string `json:"ServiceName" xml:"ServiceName"`
	ServiceStatus         string `json:"ServiceStatus" xml:"ServiceStatus"`
	ServiceDescription    string `json:"ServiceDescription" xml:"ServiceDescription"`
	CreateTime            string `json:"CreateTime" xml:"CreateTime"`
	ServiceDomain         string `json:"ServiceDomain" xml:"ServiceDomain"`
	ZoneAffinityEnabled   bool   `json:"ZoneAffinityEnabled" xml:"ZoneAffinityEnabled"`
	AutoAcceptEnabled     bool   `json:"AutoAcceptEnabled" xml:"AutoAcceptEnabled"`
	ServiceId             string `json:"ServiceId" xml:"ServiceId"`
	ServiceSupportIPv6    bool   `json:"ServiceSupportIPv6" xml:"ServiceSupportIPv6"`
	ResourceGroupId       string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

// CreateCreateVpcEndpointServiceRequest creates a request to invoke CreateVpcEndpointService API
func CreateCreateVpcEndpointServiceRequest() (request *CreateVpcEndpointServiceRequest) {
	request = &CreateVpcEndpointServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "CreateVpcEndpointService", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVpcEndpointServiceResponse creates a response to parse from CreateVpcEndpointService response
func CreateCreateVpcEndpointServiceResponse() (response *CreateVpcEndpointServiceResponse) {
	response = &CreateVpcEndpointServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}