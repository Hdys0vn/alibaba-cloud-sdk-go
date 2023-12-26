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

// DescribeVpcFirewallList invokes the cloudfw.DescribeVpcFirewallList API synchronously
func (client *Client) DescribeVpcFirewallList(request *DescribeVpcFirewallListRequest) (response *DescribeVpcFirewallListResponse, err error) {
	response = CreateDescribeVpcFirewallListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpcFirewallListWithChan invokes the cloudfw.DescribeVpcFirewallList API asynchronously
func (client *Client) DescribeVpcFirewallListWithChan(request *DescribeVpcFirewallListRequest) (<-chan *DescribeVpcFirewallListResponse, <-chan error) {
	responseChan := make(chan *DescribeVpcFirewallListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpcFirewallList(request)
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

// DescribeVpcFirewallListWithCallback invokes the cloudfw.DescribeVpcFirewallList API asynchronously
func (client *Client) DescribeVpcFirewallListWithCallback(request *DescribeVpcFirewallListRequest, callback func(response *DescribeVpcFirewallListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpcFirewallListResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpcFirewallList(request)
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

// DescribeVpcFirewallListRequest is the request struct for api DescribeVpcFirewallList
type DescribeVpcFirewallListRequest struct {
	*requests.RpcRequest
	ConnectSubType       string `position:"Query" name:"ConnectSubType"`
	VpcFirewallName      string `position:"Query" name:"VpcFirewallName"`
	SourceIp             string `position:"Query" name:"SourceIp"`
	PageSize             string `position:"Query" name:"PageSize"`
	Lang                 string `position:"Query" name:"Lang"`
	VpcFirewallId        string `position:"Query" name:"VpcFirewallId"`
	CurrentPage          string `position:"Query" name:"CurrentPage"`
	FirewallSwitchStatus string `position:"Query" name:"FirewallSwitchStatus"`
	RegionNo             string `position:"Query" name:"RegionNo"`
	MemberUid            string `position:"Query" name:"MemberUid"`
	PeerUid              string `position:"Query" name:"PeerUid"`
	VpcId                string `position:"Query" name:"VpcId"`
}

// DescribeVpcFirewallListResponse is the response struct for api DescribeVpcFirewallList
type DescribeVpcFirewallListResponse struct {
	*responses.BaseResponse
	TotalCount   int                             `json:"TotalCount" xml:"TotalCount"`
	RequestId    string                          `json:"RequestId" xml:"RequestId"`
	VpcFirewalls []DataInDescribeVpcFirewallList `json:"VpcFirewalls" xml:"VpcFirewalls"`
}

// CreateDescribeVpcFirewallListRequest creates a request to invoke DescribeVpcFirewallList API
func CreateDescribeVpcFirewallListRequest() (request *DescribeVpcFirewallListRequest) {
	request = &DescribeVpcFirewallListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeVpcFirewallList", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVpcFirewallListResponse creates a response to parse from DescribeVpcFirewallList response
func CreateDescribeVpcFirewallListResponse() (response *DescribeVpcFirewallListResponse) {
	response = &DescribeVpcFirewallListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}