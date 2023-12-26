package cc5g

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

// DetachVpcFromNetLink invokes the cc5g.DetachVpcFromNetLink API synchronously
func (client *Client) DetachVpcFromNetLink(request *DetachVpcFromNetLinkRequest) (response *DetachVpcFromNetLinkResponse, err error) {
	response = CreateDetachVpcFromNetLinkResponse()
	err = client.DoAction(request, response)
	return
}

// DetachVpcFromNetLinkWithChan invokes the cc5g.DetachVpcFromNetLink API asynchronously
func (client *Client) DetachVpcFromNetLinkWithChan(request *DetachVpcFromNetLinkRequest) (<-chan *DetachVpcFromNetLinkResponse, <-chan error) {
	responseChan := make(chan *DetachVpcFromNetLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachVpcFromNetLink(request)
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

// DetachVpcFromNetLinkWithCallback invokes the cc5g.DetachVpcFromNetLink API asynchronously
func (client *Client) DetachVpcFromNetLinkWithCallback(request *DetachVpcFromNetLinkRequest, callback func(response *DetachVpcFromNetLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachVpcFromNetLinkResponse
		var err error
		defer close(result)
		response, err = client.DetachVpcFromNetLink(request)
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

// DetachVpcFromNetLinkRequest is the request struct for api DetachVpcFromNetLink
type DetachVpcFromNetLinkRequest struct {
	*requests.RpcRequest
	DryRun                   requests.Boolean `position:"Query" name:"DryRun"`
	NetLinkId                string           `position:"Query" name:"NetLinkId"`
	ClientToken              string           `position:"Query" name:"ClientToken"`
	WirelessCloudConnectorId string           `position:"Query" name:"WirelessCloudConnectorId"`
}

// DetachVpcFromNetLinkResponse is the response struct for api DetachVpcFromNetLink
type DetachVpcFromNetLinkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachVpcFromNetLinkRequest creates a request to invoke DetachVpcFromNetLink API
func CreateDetachVpcFromNetLinkRequest() (request *DetachVpcFromNetLinkRequest) {
	request = &DetachVpcFromNetLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "DetachVpcFromNetLink", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachVpcFromNetLinkResponse creates a response to parse from DetachVpcFromNetLink response
func CreateDetachVpcFromNetLinkResponse() (response *DetachVpcFromNetLinkResponse) {
	response = &DetachVpcFromNetLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}