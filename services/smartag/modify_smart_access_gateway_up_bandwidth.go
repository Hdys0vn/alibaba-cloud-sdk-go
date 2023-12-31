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

// ModifySmartAccessGatewayUpBandwidth invokes the smartag.ModifySmartAccessGatewayUpBandwidth API synchronously
func (client *Client) ModifySmartAccessGatewayUpBandwidth(request *ModifySmartAccessGatewayUpBandwidthRequest) (response *ModifySmartAccessGatewayUpBandwidthResponse, err error) {
	response = CreateModifySmartAccessGatewayUpBandwidthResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySmartAccessGatewayUpBandwidthWithChan invokes the smartag.ModifySmartAccessGatewayUpBandwidth API asynchronously
func (client *Client) ModifySmartAccessGatewayUpBandwidthWithChan(request *ModifySmartAccessGatewayUpBandwidthRequest) (<-chan *ModifySmartAccessGatewayUpBandwidthResponse, <-chan error) {
	responseChan := make(chan *ModifySmartAccessGatewayUpBandwidthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySmartAccessGatewayUpBandwidth(request)
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

// ModifySmartAccessGatewayUpBandwidthWithCallback invokes the smartag.ModifySmartAccessGatewayUpBandwidth API asynchronously
func (client *Client) ModifySmartAccessGatewayUpBandwidthWithCallback(request *ModifySmartAccessGatewayUpBandwidthRequest, callback func(response *ModifySmartAccessGatewayUpBandwidthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySmartAccessGatewayUpBandwidthResponse
		var err error
		defer close(result)
		response, err = client.ModifySmartAccessGatewayUpBandwidth(request)
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

// ModifySmartAccessGatewayUpBandwidthRequest is the request struct for api ModifySmartAccessGatewayUpBandwidth
type ModifySmartAccessGatewayUpBandwidthRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UpBandwidthWan       requests.Integer `position:"Query" name:"UpBandwidthWan"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UpBandwidth4G        requests.Integer `position:"Query" name:"UpBandwidth4G"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// ModifySmartAccessGatewayUpBandwidthResponse is the response struct for api ModifySmartAccessGatewayUpBandwidth
type ModifySmartAccessGatewayUpBandwidthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySmartAccessGatewayUpBandwidthRequest creates a request to invoke ModifySmartAccessGatewayUpBandwidth API
func CreateModifySmartAccessGatewayUpBandwidthRequest() (request *ModifySmartAccessGatewayUpBandwidthRequest) {
	request = &ModifySmartAccessGatewayUpBandwidthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifySmartAccessGatewayUpBandwidth", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySmartAccessGatewayUpBandwidthResponse creates a response to parse from ModifySmartAccessGatewayUpBandwidth response
func CreateModifySmartAccessGatewayUpBandwidthResponse() (response *ModifySmartAccessGatewayUpBandwidthResponse) {
	response = &ModifySmartAccessGatewayUpBandwidthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
