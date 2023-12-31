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

// ListIpv4Gateways invokes the vpc.ListIpv4Gateways API synchronously
func (client *Client) ListIpv4Gateways(request *ListIpv4GatewaysRequest) (response *ListIpv4GatewaysResponse, err error) {
	response = CreateListIpv4GatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// ListIpv4GatewaysWithChan invokes the vpc.ListIpv4Gateways API asynchronously
func (client *Client) ListIpv4GatewaysWithChan(request *ListIpv4GatewaysRequest) (<-chan *ListIpv4GatewaysResponse, <-chan error) {
	responseChan := make(chan *ListIpv4GatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIpv4Gateways(request)
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

// ListIpv4GatewaysWithCallback invokes the vpc.ListIpv4Gateways API asynchronously
func (client *Client) ListIpv4GatewaysWithCallback(request *ListIpv4GatewaysRequest, callback func(response *ListIpv4GatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIpv4GatewaysResponse
		var err error
		defer close(result)
		response, err = client.ListIpv4Gateways(request)
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

// ListIpv4GatewaysRequest is the request struct for api ListIpv4Gateways
type ListIpv4GatewaysRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer        `position:"Query" name:"ResourceOwnerId"`
	Ipv4GatewayName      string                  `position:"Query" name:"Ipv4GatewayName"`
	ResourceGroupId      string                  `position:"Query" name:"ResourceGroupId"`
	NextToken            string                  `position:"Query" name:"NextToken"`
	Ipv4GatewayId        string                  `position:"Query" name:"Ipv4GatewayId"`
	ResourceOwnerAccount string                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                  `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer        `position:"Query" name:"OwnerId"`
	Tags                 *[]ListIpv4GatewaysTags `position:"Query" name:"Tags"  type:"Repeated"`
	VpcId                string                  `position:"Query" name:"VpcId"`
	MaxResults           requests.Integer        `position:"Query" name:"MaxResults"`
}

// ListIpv4GatewaysTags is a repeated param struct in ListIpv4GatewaysRequest
type ListIpv4GatewaysTags struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// ListIpv4GatewaysResponse is the response struct for api ListIpv4Gateways
type ListIpv4GatewaysResponse struct {
	*responses.BaseResponse
	NextToken         string                  `json:"NextToken" xml:"NextToken"`
	RequestId         string                  `json:"RequestId" xml:"RequestId"`
	TotalCount        string                  `json:"TotalCount" xml:"TotalCount"`
	Ipv4GatewayModels []Ipv4GatewayModelsItem `json:"Ipv4GatewayModels" xml:"Ipv4GatewayModels"`
}

// CreateListIpv4GatewaysRequest creates a request to invoke ListIpv4Gateways API
func CreateListIpv4GatewaysRequest() (request *ListIpv4GatewaysRequest) {
	request = &ListIpv4GatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ListIpv4Gateways", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListIpv4GatewaysResponse creates a response to parse from ListIpv4Gateways response
func CreateListIpv4GatewaysResponse() (response *ListIpv4GatewaysResponse) {
	response = &ListIpv4GatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
