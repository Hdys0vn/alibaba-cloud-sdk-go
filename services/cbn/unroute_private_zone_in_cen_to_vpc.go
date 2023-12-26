package cbn

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

// UnroutePrivateZoneInCenToVpc invokes the cbn.UnroutePrivateZoneInCenToVpc API synchronously
func (client *Client) UnroutePrivateZoneInCenToVpc(request *UnroutePrivateZoneInCenToVpcRequest) (response *UnroutePrivateZoneInCenToVpcResponse, err error) {
	response = CreateUnroutePrivateZoneInCenToVpcResponse()
	err = client.DoAction(request, response)
	return
}

// UnroutePrivateZoneInCenToVpcWithChan invokes the cbn.UnroutePrivateZoneInCenToVpc API asynchronously
func (client *Client) UnroutePrivateZoneInCenToVpcWithChan(request *UnroutePrivateZoneInCenToVpcRequest) (<-chan *UnroutePrivateZoneInCenToVpcResponse, <-chan error) {
	responseChan := make(chan *UnroutePrivateZoneInCenToVpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnroutePrivateZoneInCenToVpc(request)
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

// UnroutePrivateZoneInCenToVpcWithCallback invokes the cbn.UnroutePrivateZoneInCenToVpc API asynchronously
func (client *Client) UnroutePrivateZoneInCenToVpcWithCallback(request *UnroutePrivateZoneInCenToVpcRequest, callback func(response *UnroutePrivateZoneInCenToVpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnroutePrivateZoneInCenToVpcResponse
		var err error
		defer close(result)
		response, err = client.UnroutePrivateZoneInCenToVpc(request)
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

// UnroutePrivateZoneInCenToVpcRequest is the request struct for api UnroutePrivateZoneInCenToVpc
type UnroutePrivateZoneInCenToVpcRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	AccessRegionId       string           `position:"Query" name:"AccessRegionId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
}

// UnroutePrivateZoneInCenToVpcResponse is the response struct for api UnroutePrivateZoneInCenToVpc
type UnroutePrivateZoneInCenToVpcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnroutePrivateZoneInCenToVpcRequest creates a request to invoke UnroutePrivateZoneInCenToVpc API
func CreateUnroutePrivateZoneInCenToVpcRequest() (request *UnroutePrivateZoneInCenToVpcRequest) {
	request = &UnroutePrivateZoneInCenToVpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "UnroutePrivateZoneInCenToVpc", "", "")
	request.Method = requests.POST
	return
}

// CreateUnroutePrivateZoneInCenToVpcResponse creates a response to parse from UnroutePrivateZoneInCenToVpc response
func CreateUnroutePrivateZoneInCenToVpcResponse() (response *UnroutePrivateZoneInCenToVpcResponse) {
	response = &UnroutePrivateZoneInCenToVpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
