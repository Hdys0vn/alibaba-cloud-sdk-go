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

// CreateADConnectorOfficeSite invokes the ecd.CreateADConnectorOfficeSite API synchronously
func (client *Client) CreateADConnectorOfficeSite(request *CreateADConnectorOfficeSiteRequest) (response *CreateADConnectorOfficeSiteResponse, err error) {
	response = CreateCreateADConnectorOfficeSiteResponse()
	err = client.DoAction(request, response)
	return
}

// CreateADConnectorOfficeSiteWithChan invokes the ecd.CreateADConnectorOfficeSite API asynchronously
func (client *Client) CreateADConnectorOfficeSiteWithChan(request *CreateADConnectorOfficeSiteRequest) (<-chan *CreateADConnectorOfficeSiteResponse, <-chan error) {
	responseChan := make(chan *CreateADConnectorOfficeSiteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateADConnectorOfficeSite(request)
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

// CreateADConnectorOfficeSiteWithCallback invokes the ecd.CreateADConnectorOfficeSite API asynchronously
func (client *Client) CreateADConnectorOfficeSiteWithCallback(request *CreateADConnectorOfficeSiteRequest, callback func(response *CreateADConnectorOfficeSiteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateADConnectorOfficeSiteResponse
		var err error
		defer close(result)
		response, err = client.CreateADConnectorOfficeSite(request)
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

// CreateADConnectorOfficeSiteRequest is the request struct for api CreateADConnectorOfficeSite
type CreateADConnectorOfficeSiteRequest struct {
	*requests.RpcRequest
	CenId                string           `position:"Query" name:"CenId"`
	SubDomainDnsAddress  *[]string        `position:"Query" name:"SubDomainDnsAddress"  type:"Repeated"`
	CenOwnerId           requests.Integer `position:"Query" name:"CenOwnerId"`
	EnableInternetAccess requests.Boolean `position:"Query" name:"EnableInternetAccess"`
	SubDomainName        string           `position:"Query" name:"SubDomainName"`
	DomainPassword       string           `position:"Query" name:"DomainPassword"`
	VerifyCode           string           `position:"Query" name:"VerifyCode"`
	EnableAdminAccess    requests.Boolean `position:"Query" name:"EnableAdminAccess"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	DesktopAccessType    string           `position:"Query" name:"DesktopAccessType"`
	AdHostname           string           `position:"Query" name:"AdHostname"`
	DomainName           string           `position:"Query" name:"DomainName"`
	Specification        requests.Integer `position:"Query" name:"Specification"`
	OfficeSiteName       string           `position:"Query" name:"OfficeSiteName"`
	MfaEnabled           requests.Boolean `position:"Query" name:"MfaEnabled"`
	DomainUserName       string           `position:"Query" name:"DomainUserName"`
	CidrBlock            string           `position:"Query" name:"CidrBlock"`
	ProtocolType         string           `position:"Query" name:"ProtocolType"`
	DnsAddress           *[]string        `position:"Query" name:"DnsAddress"  type:"Repeated"`
}

// CreateADConnectorOfficeSiteResponse is the response struct for api CreateADConnectorOfficeSite
type CreateADConnectorOfficeSiteResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	OfficeSiteId string `json:"OfficeSiteId" xml:"OfficeSiteId"`
}

// CreateCreateADConnectorOfficeSiteRequest creates a request to invoke CreateADConnectorOfficeSite API
func CreateCreateADConnectorOfficeSiteRequest() (request *CreateADConnectorOfficeSiteRequest) {
	request = &CreateADConnectorOfficeSiteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateADConnectorOfficeSite", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateADConnectorOfficeSiteResponse creates a response to parse from CreateADConnectorOfficeSite response
func CreateCreateADConnectorOfficeSiteResponse() (response *CreateADConnectorOfficeSiteResponse) {
	response = &CreateADConnectorOfficeSiteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}