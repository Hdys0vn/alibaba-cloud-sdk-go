package vpc

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

// ModifyVpnConnectionAttribute invokes the vpc.ModifyVpnConnectionAttribute API synchronously
func (client *Client) ModifyVpnConnectionAttribute(request *ModifyVpnConnectionAttributeRequest) (response *ModifyVpnConnectionAttributeResponse, err error) {
	response = CreateModifyVpnConnectionAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpnConnectionAttributeWithChan invokes the vpc.ModifyVpnConnectionAttribute API asynchronously
func (client *Client) ModifyVpnConnectionAttributeWithChan(request *ModifyVpnConnectionAttributeRequest) (<-chan *ModifyVpnConnectionAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVpnConnectionAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpnConnectionAttribute(request)
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

// ModifyVpnConnectionAttributeWithCallback invokes the vpc.ModifyVpnConnectionAttribute API asynchronously
func (client *Client) ModifyVpnConnectionAttributeWithCallback(request *ModifyVpnConnectionAttributeRequest, callback func(response *ModifyVpnConnectionAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpnConnectionAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpnConnectionAttribute(request)
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

// ModifyVpnConnectionAttributeRequest is the request struct for api ModifyVpnConnectionAttribute
type ModifyVpnConnectionAttributeRequest struct {
	*requests.RpcRequest
	IkeConfig                  string                                                    `position:"Query" name:"IkeConfig"`
	ResourceOwnerId            requests.Integer                                          `position:"Query" name:"ResourceOwnerId"`
	AutoConfigRoute            requests.Boolean                                          `position:"Query" name:"AutoConfigRoute"`
	ClientToken                string                                                    `position:"Query" name:"ClientToken"`
	IpsecConfig                string                                                    `position:"Query" name:"IpsecConfig"`
	BgpConfig                  string                                                    `position:"Query" name:"BgpConfig"`
	HealthCheckConfig          string                                                    `position:"Query" name:"HealthCheckConfig"`
	LocalSubnet                string                                                    `position:"Query" name:"LocalSubnet"`
	EnableTunnelsBgp           requests.Boolean                                          `position:"Query" name:"EnableTunnelsBgp"`
	RemoteSubnet               string                                                    `position:"Query" name:"RemoteSubnet"`
	EffectImmediately          requests.Boolean                                          `position:"Query" name:"EffectImmediately"`
	ResourceOwnerAccount       string                                                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string                                                    `position:"Query" name:"OwnerAccount"`
	OwnerId                    requests.Integer                                          `position:"Query" name:"OwnerId"`
	EnableDpd                  requests.Boolean                                          `position:"Query" name:"EnableDpd"`
	TunnelOptionsSpecification *[]ModifyVpnConnectionAttributeTunnelOptionsSpecification `position:"Body" name:"TunnelOptionsSpecification"  type:"Repeated"`
	RemoteCaCertificate        string                                                    `position:"Query" name:"RemoteCaCertificate"`
	VpnConnectionId            string                                                    `position:"Query" name:"VpnConnectionId"`
	Name                       string                                                    `position:"Query" name:"Name"`
	EnableNatTraversal         requests.Boolean                                          `position:"Query" name:"EnableNatTraversal"`
}

// ModifyVpnConnectionAttributeTunnelOptionsSpecification is a repeated param struct in ModifyVpnConnectionAttributeRequest
type ModifyVpnConnectionAttributeTunnelOptionsSpecification struct {
	TunnelIpsecConfig   ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelIpsecConfig `name:"TunnelIpsecConfig" type:"Struct"`
	TunnelBgpConfig     ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelBgpConfig   `name:"TunnelBgpConfig" type:"Struct"`
	RemoteCaCertificate string                                                                  `name:"RemoteCaCertificate"`
	TunnelId            string                                                                  `name:"TunnelId"`
	TunnelIkeConfig     ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelIkeConfig   `name:"TunnelIkeConfig" type:"Struct"`
	EnableNatTraversal  string                                                                  `name:"EnableNatTraversal"`
	EnableDpd           string                                                                  `name:"EnableDpd"`
}

// ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelIpsecConfig is a repeated param struct in ModifyVpnConnectionAttributeRequest
type ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelIpsecConfig struct {
	IpsecPfs      string `name:"IpsecPfs"`
	IpsecLifetime string `name:"IpsecLifetime"`
	IpsecAuthAlg  string `name:"IpsecAuthAlg"`
	IpsecEncAlg   string `name:"IpsecEncAlg"`
}

// ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelBgpConfig is a repeated param struct in ModifyVpnConnectionAttributeRequest
type ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelBgpConfig struct {
	LocalAsn   string `name:"LocalAsn"`
	TunnelCidr string `name:"TunnelCidr"`
	LocalBgpIp string `name:"LocalBgpIp"`
}

// ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelIkeConfig is a repeated param struct in ModifyVpnConnectionAttributeRequest
type ModifyVpnConnectionAttributeTunnelOptionsSpecificationTunnelIkeConfig struct {
	IkeVersion  string `name:"IkeVersion"`
	IkeMode     string `name:"IkeMode"`
	IkeAuthAlg  string `name:"IkeAuthAlg"`
	Psk         string `name:"Psk"`
	IkePfs      string `name:"IkePfs"`
	IkeLifetime string `name:"IkeLifetime"`
	LocalId     string `name:"LocalId"`
	IkeEncAlg   string `name:"IkeEncAlg"`
	RemoteId    string `name:"RemoteId"`
}

// ModifyVpnConnectionAttributeResponse is the response struct for api ModifyVpnConnectionAttribute
type ModifyVpnConnectionAttributeResponse struct {
	*responses.BaseResponse
	EnableNatTraversal         bool                                                     `json:"EnableNatTraversal" xml:"EnableNatTraversal"`
	CreateTime                 int64                                                    `json:"CreateTime" xml:"CreateTime"`
	EffectImmediately          bool                                                     `json:"EffectImmediately" xml:"EffectImmediately"`
	VpnGatewayId               string                                                   `json:"VpnGatewayId" xml:"VpnGatewayId"`
	LocalSubnet                string                                                   `json:"LocalSubnet" xml:"LocalSubnet"`
	RequestId                  string                                                   `json:"RequestId" xml:"RequestId"`
	VpnConnectionId            string                                                   `json:"VpnConnectionId" xml:"VpnConnectionId"`
	Description                string                                                   `json:"Description" xml:"Description"`
	RemoteSubnet               string                                                   `json:"RemoteSubnet" xml:"RemoteSubnet"`
	CustomerGatewayId          string                                                   `json:"CustomerGatewayId" xml:"CustomerGatewayId"`
	Name                       string                                                   `json:"Name" xml:"Name"`
	EnableDpd                  bool                                                     `json:"EnableDpd" xml:"EnableDpd"`
	EnableTunnelsBgp           bool                                                     `json:"EnableTunnelsBgp" xml:"EnableTunnelsBgp"`
	IkeConfig                  IkeConfig                                                `json:"IkeConfig" xml:"IkeConfig"`
	IpsecConfig                IpsecConfig                                              `json:"IpsecConfig" xml:"IpsecConfig"`
	VcoHealthCheck             VcoHealthCheck                                           `json:"VcoHealthCheck" xml:"VcoHealthCheck"`
	VpnBgpConfig               VpnBgpConfigInModifyVpnConnectionAttribute               `json:"VpnBgpConfig" xml:"VpnBgpConfig"`
	TunnelOptionsSpecification TunnelOptionsSpecificationInModifyVpnConnectionAttribute `json:"TunnelOptionsSpecification" xml:"TunnelOptionsSpecification"`
}

// CreateModifyVpnConnectionAttributeRequest creates a request to invoke ModifyVpnConnectionAttribute API
func CreateModifyVpnConnectionAttributeRequest() (request *ModifyVpnConnectionAttributeRequest) {
	request = &ModifyVpnConnectionAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnConnectionAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpnConnectionAttributeResponse creates a response to parse from ModifyVpnConnectionAttribute response
func CreateModifyVpnConnectionAttributeResponse() (response *ModifyVpnConnectionAttributeResponse) {
	response = &ModifyVpnConnectionAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
