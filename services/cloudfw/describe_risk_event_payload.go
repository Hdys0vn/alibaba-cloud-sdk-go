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

// DescribeRiskEventPayload invokes the cloudfw.DescribeRiskEventPayload API synchronously
func (client *Client) DescribeRiskEventPayload(request *DescribeRiskEventPayloadRequest) (response *DescribeRiskEventPayloadResponse, err error) {
	response = CreateDescribeRiskEventPayloadResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRiskEventPayloadWithChan invokes the cloudfw.DescribeRiskEventPayload API asynchronously
func (client *Client) DescribeRiskEventPayloadWithChan(request *DescribeRiskEventPayloadRequest) (<-chan *DescribeRiskEventPayloadResponse, <-chan error) {
	responseChan := make(chan *DescribeRiskEventPayloadResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRiskEventPayload(request)
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

// DescribeRiskEventPayloadWithCallback invokes the cloudfw.DescribeRiskEventPayload API asynchronously
func (client *Client) DescribeRiskEventPayloadWithCallback(request *DescribeRiskEventPayloadRequest, callback func(response *DescribeRiskEventPayloadResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRiskEventPayloadResponse
		var err error
		defer close(result)
		response, err = client.DescribeRiskEventPayload(request)
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

// DescribeRiskEventPayloadRequest is the request struct for api DescribeRiskEventPayload
type DescribeRiskEventPayloadRequest struct {
	*requests.RpcRequest
	SrcIP        string `position:"Query" name:"SrcIP"`
	PublicIP     string `position:"Query" name:"PublicIP"`
	StartTime    string `position:"Query" name:"StartTime"`
	UUID         string `position:"Query" name:"UUID"`
	DstVpcId     string `position:"Query" name:"DstVpcId"`
	DstIP        string `position:"Query" name:"DstIP"`
	FirewallType string `position:"Query" name:"FirewallType"`
	EndTime      string `position:"Query" name:"EndTime"`
	SrcVpcId     string `position:"Query" name:"SrcVpcId"`
}

// DescribeRiskEventPayloadResponse is the response struct for api DescribeRiskEventPayload
type DescribeRiskEventPayloadResponse struct {
	*responses.BaseResponse
	HighlightEnd   int    `json:"HighlightEnd" xml:"HighlightEnd"`
	DstIP          string `json:"DstIP" xml:"DstIP"`
	DstPort        int    `json:"DstPort" xml:"DstPort"`
	Payload        string `json:"Payload" xml:"Payload"`
	SrcVpcId       string `json:"SrcVpcId" xml:"SrcVpcId"`
	SrcPort        int    `json:"SrcPort" xml:"SrcPort"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	PayloadLen     int    `json:"PayloadLen" xml:"PayloadLen"`
	DstVpcId       string `json:"DstVpcId" xml:"DstVpcId"`
	SrcIP          string `json:"SrcIP" xml:"SrcIP"`
	HighlightStart int    `json:"HighlightStart" xml:"HighlightStart"`
	Proto          string `json:"Proto" xml:"Proto"`
	XForwardFor    string `json:"XForwardFor" xml:"XForwardFor"`
	RealIp         string `json:"RealIp" xml:"RealIp"`
}

// CreateDescribeRiskEventPayloadRequest creates a request to invoke DescribeRiskEventPayload API
func CreateDescribeRiskEventPayloadRequest() (request *DescribeRiskEventPayloadRequest) {
	request = &DescribeRiskEventPayloadRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeRiskEventPayload", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRiskEventPayloadResponse creates a response to parse from DescribeRiskEventPayload response
func CreateDescribeRiskEventPayloadResponse() (response *DescribeRiskEventPayloadResponse) {
	response = &DescribeRiskEventPayloadResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}