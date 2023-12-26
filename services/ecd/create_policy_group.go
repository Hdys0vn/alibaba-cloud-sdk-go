package ecd

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

// CreatePolicyGroup invokes the ecd.CreatePolicyGroup API synchronously
func (client *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (response *CreatePolicyGroupResponse, err error) {
	response = CreateCreatePolicyGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePolicyGroupWithChan invokes the ecd.CreatePolicyGroup API asynchronously
func (client *Client) CreatePolicyGroupWithChan(request *CreatePolicyGroupRequest) (<-chan *CreatePolicyGroupResponse, <-chan error) {
	responseChan := make(chan *CreatePolicyGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePolicyGroup(request)
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

// CreatePolicyGroupWithCallback invokes the ecd.CreatePolicyGroup API asynchronously
func (client *Client) CreatePolicyGroupWithCallback(request *CreatePolicyGroupRequest, callback func(response *CreatePolicyGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePolicyGroupResponse
		var err error
		defer close(result)
		response, err = client.CreatePolicyGroup(request)
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

// CreatePolicyGroupRequest is the request struct for api CreatePolicyGroup
type CreatePolicyGroupRequest struct {
	*requests.RpcRequest
	WatermarkCustomText         string                                          `position:"Query" name:"WatermarkCustomText"`
	PreemptLogin                string                                          `position:"Query" name:"PreemptLogin"`
	ClientType                  *[]CreatePolicyGroupClientType                  `position:"Query" name:"ClientType"  type:"Repeated"`
	UsbSupplyRedirectRule       *[]CreatePolicyGroupUsbSupplyRedirectRule       `position:"Query" name:"UsbSupplyRedirectRule"  type:"Repeated"`
	PrinterRedirection          string                                          `position:"Query" name:"PrinterRedirection"`
	PreemptLoginUser            *[]string                                       `position:"Query" name:"PreemptLoginUser"  type:"Repeated"`
	DomainList                  string                                          `position:"Query" name:"DomainList"`
	NetRedirect                 string                                          `position:"Query" name:"NetRedirect"`
	LocalDrive                  string                                          `position:"Query" name:"LocalDrive"`
	AuthorizeSecurityPolicyRule *[]CreatePolicyGroupAuthorizeSecurityPolicyRule `position:"Query" name:"AuthorizeSecurityPolicyRule"  type:"Repeated"`
	Recording                   string                                          `position:"Query" name:"Recording"`
	Clipboard                   string                                          `position:"Query" name:"Clipboard"`
	RecordingFps                requests.Integer                                `position:"Query" name:"RecordingFps"`
	UsbRedirect                 string                                          `position:"Query" name:"UsbRedirect"`
	WatermarkType               string                                          `position:"Query" name:"WatermarkType"`
	RecordingStartTime          string                                          `position:"Query" name:"RecordingStartTime"`
	Watermark                   string                                          `position:"Query" name:"Watermark"`
	CameraRedirect              string                                          `position:"Query" name:"CameraRedirect"`
	Html5Access                 string                                          `position:"Query" name:"Html5Access"`
	GpuAcceleration             string                                          `position:"Query" name:"GpuAcceleration"`
	Html5FileTransfer           string                                          `position:"Query" name:"Html5FileTransfer"`
	AuthorizeAccessPolicyRule   *[]CreatePolicyGroupAuthorizeAccessPolicyRule   `position:"Query" name:"AuthorizeAccessPolicyRule"  type:"Repeated"`
	VisualQuality               string                                          `position:"Query" name:"VisualQuality"`
	WatermarkTransparency       string                                          `position:"Query" name:"WatermarkTransparency"`
	Name                        string                                          `position:"Query" name:"Name"`
	RecordingEndTime            string                                          `position:"Query" name:"RecordingEndTime"`
}

// CreatePolicyGroupClientType is a repeated param struct in CreatePolicyGroupRequest
type CreatePolicyGroupClientType struct {
	ClientType string `name:"ClientType"`
	Status     string `name:"Status"`
}

// CreatePolicyGroupUsbSupplyRedirectRule is a repeated param struct in CreatePolicyGroupRequest
type CreatePolicyGroupUsbSupplyRedirectRule struct {
	ProductId       string `name:"ProductId"`
	DeviceSubclass  string `name:"DeviceSubclass"`
	UsbRedirectType string `name:"UsbRedirectType"`
	VendorId        string `name:"VendorId"`
	Description     string `name:"Description"`
	DeviceClass     string `name:"DeviceClass"`
	UsbRuleType     string `name:"UsbRuleType"`
}

// CreatePolicyGroupAuthorizeSecurityPolicyRule is a repeated param struct in CreatePolicyGroupRequest
type CreatePolicyGroupAuthorizeSecurityPolicyRule struct {
	PortRange   string `name:"PortRange"`
	IpProtocol  string `name:"IpProtocol"`
	Description string `name:"Description"`
	Type        string `name:"Type"`
	Priority    string `name:"Priority"`
	CidrIp      string `name:"CidrIp"`
	Policy      string `name:"Policy"`
}

// CreatePolicyGroupAuthorizeAccessPolicyRule is a repeated param struct in CreatePolicyGroupRequest
type CreatePolicyGroupAuthorizeAccessPolicyRule struct {
	Description string `name:"Description"`
	CidrIp      string `name:"CidrIp"`
}

// CreatePolicyGroupResponse is the response struct for api CreatePolicyGroup
type CreatePolicyGroupResponse struct {
	*responses.BaseResponse
	PolicyGroupId string `json:"PolicyGroupId" xml:"PolicyGroupId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatePolicyGroupRequest creates a request to invoke CreatePolicyGroup API
func CreateCreatePolicyGroupRequest() (request *CreatePolicyGroupRequest) {
	request = &CreatePolicyGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreatePolicyGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePolicyGroupResponse creates a response to parse from CreatePolicyGroup response
func CreateCreatePolicyGroupResponse() (response *CreatePolicyGroupResponse) {
	response = &CreatePolicyGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
