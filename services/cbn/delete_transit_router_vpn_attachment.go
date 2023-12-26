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

// DeleteTransitRouterVpnAttachment invokes the cbn.DeleteTransitRouterVpnAttachment API synchronously
func (client *Client) DeleteTransitRouterVpnAttachment(request *DeleteTransitRouterVpnAttachmentRequest) (response *DeleteTransitRouterVpnAttachmentResponse, err error) {
	response = CreateDeleteTransitRouterVpnAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTransitRouterVpnAttachmentWithChan invokes the cbn.DeleteTransitRouterVpnAttachment API asynchronously
func (client *Client) DeleteTransitRouterVpnAttachmentWithChan(request *DeleteTransitRouterVpnAttachmentRequest) (<-chan *DeleteTransitRouterVpnAttachmentResponse, <-chan error) {
	responseChan := make(chan *DeleteTransitRouterVpnAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTransitRouterVpnAttachment(request)
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

// DeleteTransitRouterVpnAttachmentWithCallback invokes the cbn.DeleteTransitRouterVpnAttachment API asynchronously
func (client *Client) DeleteTransitRouterVpnAttachmentWithCallback(request *DeleteTransitRouterVpnAttachmentRequest, callback func(response *DeleteTransitRouterVpnAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTransitRouterVpnAttachmentResponse
		var err error
		defer close(result)
		response, err = client.DeleteTransitRouterVpnAttachment(request)
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

// DeleteTransitRouterVpnAttachmentRequest is the request struct for api DeleteTransitRouterVpnAttachment
type DeleteTransitRouterVpnAttachmentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string           `position:"Query" name:"ClientToken"`
	DryRun                    requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType              string           `position:"Query" name:"ResourceType"`
	Version                   string           `position:"Query" name:"Version"`
	TransitRouterAttachmentId string           `position:"Query" name:"TransitRouterAttachmentId"`
	Force                     requests.Boolean `position:"Query" name:"Force"`
}

// DeleteTransitRouterVpnAttachmentResponse is the response struct for api DeleteTransitRouterVpnAttachment
type DeleteTransitRouterVpnAttachmentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTransitRouterVpnAttachmentRequest creates a request to invoke DeleteTransitRouterVpnAttachment API
func CreateDeleteTransitRouterVpnAttachmentRequest() (request *DeleteTransitRouterVpnAttachmentRequest) {
	request = &DeleteTransitRouterVpnAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteTransitRouterVpnAttachment", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTransitRouterVpnAttachmentResponse creates a response to parse from DeleteTransitRouterVpnAttachment response
func CreateDeleteTransitRouterVpnAttachmentResponse() (response *DeleteTransitRouterVpnAttachmentResponse) {
	response = &DeleteTransitRouterVpnAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}