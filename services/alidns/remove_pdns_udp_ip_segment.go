package alidns

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

// RemovePdnsUdpIpSegment invokes the alidns.RemovePdnsUdpIpSegment API synchronously
func (client *Client) RemovePdnsUdpIpSegment(request *RemovePdnsUdpIpSegmentRequest) (response *RemovePdnsUdpIpSegmentResponse, err error) {
	response = CreateRemovePdnsUdpIpSegmentResponse()
	err = client.DoAction(request, response)
	return
}

// RemovePdnsUdpIpSegmentWithChan invokes the alidns.RemovePdnsUdpIpSegment API asynchronously
func (client *Client) RemovePdnsUdpIpSegmentWithChan(request *RemovePdnsUdpIpSegmentRequest) (<-chan *RemovePdnsUdpIpSegmentResponse, <-chan error) {
	responseChan := make(chan *RemovePdnsUdpIpSegmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemovePdnsUdpIpSegment(request)
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

// RemovePdnsUdpIpSegmentWithCallback invokes the alidns.RemovePdnsUdpIpSegment API asynchronously
func (client *Client) RemovePdnsUdpIpSegmentWithCallback(request *RemovePdnsUdpIpSegmentRequest, callback func(response *RemovePdnsUdpIpSegmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemovePdnsUdpIpSegmentResponse
		var err error
		defer close(result)
		response, err = client.RemovePdnsUdpIpSegment(request)
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

// RemovePdnsUdpIpSegmentRequest is the request struct for api RemovePdnsUdpIpSegment
type RemovePdnsUdpIpSegmentRequest struct {
	*requests.RpcRequest
	Ip   string `position:"Query" name:"Ip"`
	Lang string `position:"Query" name:"Lang"`
}

// RemovePdnsUdpIpSegmentResponse is the response struct for api RemovePdnsUdpIpSegment
type RemovePdnsUdpIpSegmentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemovePdnsUdpIpSegmentRequest creates a request to invoke RemovePdnsUdpIpSegment API
func CreateRemovePdnsUdpIpSegmentRequest() (request *RemovePdnsUdpIpSegmentRequest) {
	request = &RemovePdnsUdpIpSegmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "RemovePdnsUdpIpSegment", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemovePdnsUdpIpSegmentResponse creates a response to parse from RemovePdnsUdpIpSegment response
func CreateRemovePdnsUdpIpSegmentResponse() (response *RemovePdnsUdpIpSegmentResponse) {
	response = &RemovePdnsUdpIpSegmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
