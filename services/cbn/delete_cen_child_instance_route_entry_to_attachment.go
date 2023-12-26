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

// DeleteCenChildInstanceRouteEntryToAttachment invokes the cbn.DeleteCenChildInstanceRouteEntryToAttachment API synchronously
func (client *Client) DeleteCenChildInstanceRouteEntryToAttachment(request *DeleteCenChildInstanceRouteEntryToAttachmentRequest) (response *DeleteCenChildInstanceRouteEntryToAttachmentResponse, err error) {
	response = CreateDeleteCenChildInstanceRouteEntryToAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCenChildInstanceRouteEntryToAttachmentWithChan invokes the cbn.DeleteCenChildInstanceRouteEntryToAttachment API asynchronously
func (client *Client) DeleteCenChildInstanceRouteEntryToAttachmentWithChan(request *DeleteCenChildInstanceRouteEntryToAttachmentRequest) (<-chan *DeleteCenChildInstanceRouteEntryToAttachmentResponse, <-chan error) {
	responseChan := make(chan *DeleteCenChildInstanceRouteEntryToAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCenChildInstanceRouteEntryToAttachment(request)
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

// DeleteCenChildInstanceRouteEntryToAttachmentWithCallback invokes the cbn.DeleteCenChildInstanceRouteEntryToAttachment API asynchronously
func (client *Client) DeleteCenChildInstanceRouteEntryToAttachmentWithCallback(request *DeleteCenChildInstanceRouteEntryToAttachmentRequest, callback func(response *DeleteCenChildInstanceRouteEntryToAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCenChildInstanceRouteEntryToAttachmentResponse
		var err error
		defer close(result)
		response, err = client.DeleteCenChildInstanceRouteEntryToAttachment(request)
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

// DeleteCenChildInstanceRouteEntryToAttachmentRequest is the request struct for api DeleteCenChildInstanceRouteEntryToAttachment
type DeleteCenChildInstanceRouteEntryToAttachmentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string           `position:"Query" name:"ClientToken"`
	CenId                     string           `position:"Query" name:"CenId"`
	RouteTableId              string           `position:"Query" name:"RouteTableId"`
	DryRun                    requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	DestinationCidrBlock      string           `position:"Query" name:"DestinationCidrBlock"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	Version                   string           `position:"Query" name:"Version"`
	TransitRouterAttachmentId string           `position:"Query" name:"TransitRouterAttachmentId"`
}

// DeleteCenChildInstanceRouteEntryToAttachmentResponse is the response struct for api DeleteCenChildInstanceRouteEntryToAttachment
type DeleteCenChildInstanceRouteEntryToAttachmentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCenChildInstanceRouteEntryToAttachmentRequest creates a request to invoke DeleteCenChildInstanceRouteEntryToAttachment API
func CreateDeleteCenChildInstanceRouteEntryToAttachmentRequest() (request *DeleteCenChildInstanceRouteEntryToAttachmentRequest) {
	request = &DeleteCenChildInstanceRouteEntryToAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteCenChildInstanceRouteEntryToAttachment", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCenChildInstanceRouteEntryToAttachmentResponse creates a response to parse from DeleteCenChildInstanceRouteEntryToAttachment response
func CreateDeleteCenChildInstanceRouteEntryToAttachmentResponse() (response *DeleteCenChildInstanceRouteEntryToAttachmentResponse) {
	response = &DeleteCenChildInstanceRouteEntryToAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}