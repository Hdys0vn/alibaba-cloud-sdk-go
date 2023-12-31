package smartag

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

// ModifyDeviceAutoUpgradePolicy invokes the smartag.ModifyDeviceAutoUpgradePolicy API synchronously
func (client *Client) ModifyDeviceAutoUpgradePolicy(request *ModifyDeviceAutoUpgradePolicyRequest) (response *ModifyDeviceAutoUpgradePolicyResponse, err error) {
	response = CreateModifyDeviceAutoUpgradePolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDeviceAutoUpgradePolicyWithChan invokes the smartag.ModifyDeviceAutoUpgradePolicy API asynchronously
func (client *Client) ModifyDeviceAutoUpgradePolicyWithChan(request *ModifyDeviceAutoUpgradePolicyRequest) (<-chan *ModifyDeviceAutoUpgradePolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyDeviceAutoUpgradePolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDeviceAutoUpgradePolicy(request)
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

// ModifyDeviceAutoUpgradePolicyWithCallback invokes the smartag.ModifyDeviceAutoUpgradePolicy API asynchronously
func (client *Client) ModifyDeviceAutoUpgradePolicyWithCallback(request *ModifyDeviceAutoUpgradePolicyRequest, callback func(response *ModifyDeviceAutoUpgradePolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDeviceAutoUpgradePolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyDeviceAutoUpgradePolicy(request)
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

// ModifyDeviceAutoUpgradePolicyRequest is the request struct for api ModifyDeviceAutoUpgradePolicy
type ModifyDeviceAutoUpgradePolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CronExpression       string           `position:"Query" name:"CronExpression"`
	TimeZone             string           `position:"Query" name:"TimeZone"`
	UpgradeType          string           `position:"Query" name:"UpgradeType"`
	Duration             requests.Integer `position:"Query" name:"Duration"`
	SerialNumber         string           `position:"Query" name:"SerialNumber"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VersionType          string           `position:"Query" name:"VersionType"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// ModifyDeviceAutoUpgradePolicyResponse is the response struct for api ModifyDeviceAutoUpgradePolicy
type ModifyDeviceAutoUpgradePolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDeviceAutoUpgradePolicyRequest creates a request to invoke ModifyDeviceAutoUpgradePolicy API
func CreateModifyDeviceAutoUpgradePolicyRequest() (request *ModifyDeviceAutoUpgradePolicyRequest) {
	request = &ModifyDeviceAutoUpgradePolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifyDeviceAutoUpgradePolicy", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDeviceAutoUpgradePolicyResponse creates a response to parse from ModifyDeviceAutoUpgradePolicy response
func CreateModifyDeviceAutoUpgradePolicyResponse() (response *ModifyDeviceAutoUpgradePolicyResponse) {
	response = &ModifyDeviceAutoUpgradePolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
