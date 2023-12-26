package cloudfw

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

// DescribeTrFirewallPolicyBackUpAssociationList invokes the cloudfw.DescribeTrFirewallPolicyBackUpAssociationList API synchronously
func (client *Client) DescribeTrFirewallPolicyBackUpAssociationList(request *DescribeTrFirewallPolicyBackUpAssociationListRequest) (response *DescribeTrFirewallPolicyBackUpAssociationListResponse, err error) {
	response = CreateDescribeTrFirewallPolicyBackUpAssociationListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTrFirewallPolicyBackUpAssociationListWithChan invokes the cloudfw.DescribeTrFirewallPolicyBackUpAssociationList API asynchronously
func (client *Client) DescribeTrFirewallPolicyBackUpAssociationListWithChan(request *DescribeTrFirewallPolicyBackUpAssociationListRequest) (<-chan *DescribeTrFirewallPolicyBackUpAssociationListResponse, <-chan error) {
	responseChan := make(chan *DescribeTrFirewallPolicyBackUpAssociationListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTrFirewallPolicyBackUpAssociationList(request)
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

// DescribeTrFirewallPolicyBackUpAssociationListWithCallback invokes the cloudfw.DescribeTrFirewallPolicyBackUpAssociationList API asynchronously
func (client *Client) DescribeTrFirewallPolicyBackUpAssociationListWithCallback(request *DescribeTrFirewallPolicyBackUpAssociationListRequest, callback func(response *DescribeTrFirewallPolicyBackUpAssociationListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTrFirewallPolicyBackUpAssociationListResponse
		var err error
		defer close(result)
		response, err = client.DescribeTrFirewallPolicyBackUpAssociationList(request)
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

// DescribeTrFirewallPolicyBackUpAssociationListRequest is the request struct for api DescribeTrFirewallPolicyBackUpAssociationList
type DescribeTrFirewallPolicyBackUpAssociationListRequest struct {
	*requests.RpcRequest
	FirewallId              string `position:"Query" name:"FirewallId"`
	SourceIp                string `position:"Query" name:"SourceIp"`
	Lang                    string `position:"Query" name:"Lang"`
	TrFirewallRoutePolicyId string `position:"Query" name:"TrFirewallRoutePolicyId"`
}

// DescribeTrFirewallPolicyBackUpAssociationListResponse is the response struct for api DescribeTrFirewallPolicyBackUpAssociationList
type DescribeTrFirewallPolicyBackUpAssociationListResponse struct {
	*responses.BaseResponse
	RequestId                      string                               `json:"RequestId" xml:"RequestId"`
	PolicyAssociationBackupConfigs []PolicyAssociationBackupConfigsItem `json:"PolicyAssociationBackupConfigs" xml:"PolicyAssociationBackupConfigs"`
}

// CreateDescribeTrFirewallPolicyBackUpAssociationListRequest creates a request to invoke DescribeTrFirewallPolicyBackUpAssociationList API
func CreateDescribeTrFirewallPolicyBackUpAssociationListRequest() (request *DescribeTrFirewallPolicyBackUpAssociationListRequest) {
	request = &DescribeTrFirewallPolicyBackUpAssociationListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeTrFirewallPolicyBackUpAssociationList", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTrFirewallPolicyBackUpAssociationListResponse creates a response to parse from DescribeTrFirewallPolicyBackUpAssociationList response
func CreateDescribeTrFirewallPolicyBackUpAssociationListResponse() (response *DescribeTrFirewallPolicyBackUpAssociationListResponse) {
	response = &DescribeTrFirewallPolicyBackUpAssociationListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
