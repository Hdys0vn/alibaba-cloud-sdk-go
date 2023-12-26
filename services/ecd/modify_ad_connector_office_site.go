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

// ModifyADConnectorOfficeSite invokes the ecd.ModifyADConnectorOfficeSite API synchronously
func (client *Client) ModifyADConnectorOfficeSite(request *ModifyADConnectorOfficeSiteRequest) (response *ModifyADConnectorOfficeSiteResponse, err error) {
	response = CreateModifyADConnectorOfficeSiteResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyADConnectorOfficeSiteWithChan invokes the ecd.ModifyADConnectorOfficeSite API asynchronously
func (client *Client) ModifyADConnectorOfficeSiteWithChan(request *ModifyADConnectorOfficeSiteRequest) (<-chan *ModifyADConnectorOfficeSiteResponse, <-chan error) {
	responseChan := make(chan *ModifyADConnectorOfficeSiteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyADConnectorOfficeSite(request)
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

// ModifyADConnectorOfficeSiteWithCallback invokes the ecd.ModifyADConnectorOfficeSite API asynchronously
func (client *Client) ModifyADConnectorOfficeSiteWithCallback(request *ModifyADConnectorOfficeSiteRequest, callback func(response *ModifyADConnectorOfficeSiteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyADConnectorOfficeSiteResponse
		var err error
		defer close(result)
		response, err = client.ModifyADConnectorOfficeSite(request)
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

// ModifyADConnectorOfficeSiteRequest is the request struct for api ModifyADConnectorOfficeSite
type ModifyADConnectorOfficeSiteRequest struct {
	*requests.RpcRequest
	OfficeSiteId        string           `position:"Query" name:"OfficeSiteId"`
	SubDomainDnsAddress *[]string        `position:"Query" name:"SubDomainDnsAddress"  type:"Repeated"`
	SubDomainName       string           `position:"Query" name:"SubDomainName"`
	DomainPassword      string           `position:"Query" name:"DomainPassword"`
	AdHostname          string           `position:"Query" name:"AdHostname"`
	DomainName          string           `position:"Query" name:"DomainName"`
	OfficeSiteName      string           `position:"Query" name:"OfficeSiteName"`
	MfaEnabled          requests.Boolean `position:"Query" name:"MfaEnabled"`
	DomainUserName      string           `position:"Query" name:"DomainUserName"`
	DnsAddress          *[]string        `position:"Query" name:"DnsAddress"  type:"Repeated"`
	OUName              string           `position:"Query" name:"OUName"`
}

// ModifyADConnectorOfficeSiteResponse is the response struct for api ModifyADConnectorOfficeSite
type ModifyADConnectorOfficeSiteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyADConnectorOfficeSiteRequest creates a request to invoke ModifyADConnectorOfficeSite API
func CreateModifyADConnectorOfficeSiteRequest() (request *ModifyADConnectorOfficeSiteRequest) {
	request = &ModifyADConnectorOfficeSiteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ModifyADConnectorOfficeSite", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyADConnectorOfficeSiteResponse creates a response to parse from ModifyADConnectorOfficeSite response
func CreateModifyADConnectorOfficeSiteResponse() (response *ModifyADConnectorOfficeSiteResponse) {
	response = &ModifyADConnectorOfficeSiteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
