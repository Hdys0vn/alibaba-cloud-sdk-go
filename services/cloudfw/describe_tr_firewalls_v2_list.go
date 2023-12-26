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

// DescribeTrFirewallsV2List invokes the cloudfw.DescribeTrFirewallsV2List API synchronously
func (client *Client) DescribeTrFirewallsV2List(request *DescribeTrFirewallsV2ListRequest) (response *DescribeTrFirewallsV2ListResponse, err error) {
	response = CreateDescribeTrFirewallsV2ListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTrFirewallsV2ListWithChan invokes the cloudfw.DescribeTrFirewallsV2List API asynchronously
func (client *Client) DescribeTrFirewallsV2ListWithChan(request *DescribeTrFirewallsV2ListRequest) (<-chan *DescribeTrFirewallsV2ListResponse, <-chan error) {
	responseChan := make(chan *DescribeTrFirewallsV2ListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTrFirewallsV2List(request)
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

// DescribeTrFirewallsV2ListWithCallback invokes the cloudfw.DescribeTrFirewallsV2List API asynchronously
func (client *Client) DescribeTrFirewallsV2ListWithCallback(request *DescribeTrFirewallsV2ListRequest, callback func(response *DescribeTrFirewallsV2ListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTrFirewallsV2ListResponse
		var err error
		defer close(result)
		response, err = client.DescribeTrFirewallsV2List(request)
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

// DescribeTrFirewallsV2ListRequest is the request struct for api DescribeTrFirewallsV2List
type DescribeTrFirewallsV2ListRequest struct {
	*requests.RpcRequest
	CenId                string           `position:"Query" name:"CenId"`
	FirewallId           string           `position:"Query" name:"FirewallId"`
	SourceIp             string           `position:"Query" name:"SourceIp"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	Lang                 string           `position:"Query" name:"Lang"`
	RouteMode            string           `position:"Query" name:"RouteMode"`
	FirewallName         string           `position:"Query" name:"FirewallName"`
	CurrentPage          requests.Integer `position:"Query" name:"CurrentPage"`
	FirewallSwitchStatus string           `position:"Query" name:"FirewallSwitchStatus"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	TransitRouterId      string           `position:"Query" name:"TransitRouterId"`
	RegionNo             string           `position:"Query" name:"RegionNo"`
}

// DescribeTrFirewallsV2ListResponse is the response struct for api DescribeTrFirewallsV2List
type DescribeTrFirewallsV2ListResponse struct {
	*responses.BaseResponse
	TotalCount     string               `json:"TotalCount" xml:"TotalCount"`
	RequestId      string               `json:"RequestId" xml:"RequestId"`
	VpcTrFirewalls []VpcTrFirewallsItem `json:"VpcTrFirewalls" xml:"VpcTrFirewalls"`
}

// CreateDescribeTrFirewallsV2ListRequest creates a request to invoke DescribeTrFirewallsV2List API
func CreateDescribeTrFirewallsV2ListRequest() (request *DescribeTrFirewallsV2ListRequest) {
	request = &DescribeTrFirewallsV2ListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeTrFirewallsV2List", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTrFirewallsV2ListResponse creates a response to parse from DescribeTrFirewallsV2List response
func CreateDescribeTrFirewallsV2ListResponse() (response *DescribeTrFirewallsV2ListResponse) {
	response = &DescribeTrFirewallsV2ListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
