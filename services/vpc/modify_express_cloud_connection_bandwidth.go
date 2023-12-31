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

// ModifyExpressCloudConnectionBandwidth invokes the vpc.ModifyExpressCloudConnectionBandwidth API synchronously
func (client *Client) ModifyExpressCloudConnectionBandwidth(request *ModifyExpressCloudConnectionBandwidthRequest) (response *ModifyExpressCloudConnectionBandwidthResponse, err error) {
	response = CreateModifyExpressCloudConnectionBandwidthResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyExpressCloudConnectionBandwidthWithChan invokes the vpc.ModifyExpressCloudConnectionBandwidth API asynchronously
func (client *Client) ModifyExpressCloudConnectionBandwidthWithChan(request *ModifyExpressCloudConnectionBandwidthRequest) (<-chan *ModifyExpressCloudConnectionBandwidthResponse, <-chan error) {
	responseChan := make(chan *ModifyExpressCloudConnectionBandwidthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyExpressCloudConnectionBandwidth(request)
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

// ModifyExpressCloudConnectionBandwidthWithCallback invokes the vpc.ModifyExpressCloudConnectionBandwidth API asynchronously
func (client *Client) ModifyExpressCloudConnectionBandwidthWithCallback(request *ModifyExpressCloudConnectionBandwidthRequest, callback func(response *ModifyExpressCloudConnectionBandwidthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyExpressCloudConnectionBandwidthResponse
		var err error
		defer close(result)
		response, err = client.ModifyExpressCloudConnectionBandwidth(request)
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

// ModifyExpressCloudConnectionBandwidthRequest is the request struct for api ModifyExpressCloudConnectionBandwidth
type ModifyExpressCloudConnectionBandwidthRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EccId                string           `position:"Query" name:"EccId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string           `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyExpressCloudConnectionBandwidthResponse is the response struct for api ModifyExpressCloudConnectionBandwidth
type ModifyExpressCloudConnectionBandwidthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyExpressCloudConnectionBandwidthRequest creates a request to invoke ModifyExpressCloudConnectionBandwidth API
func CreateModifyExpressCloudConnectionBandwidthRequest() (request *ModifyExpressCloudConnectionBandwidthRequest) {
	request = &ModifyExpressCloudConnectionBandwidthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyExpressCloudConnectionBandwidth", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyExpressCloudConnectionBandwidthResponse creates a response to parse from ModifyExpressCloudConnectionBandwidth response
func CreateModifyExpressCloudConnectionBandwidthResponse() (response *ModifyExpressCloudConnectionBandwidthResponse) {
	response = &ModifyExpressCloudConnectionBandwidthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
