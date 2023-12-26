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

// DescribeEipGatewayInfo invokes the vpc.DescribeEipGatewayInfo API synchronously
func (client *Client) DescribeEipGatewayInfo(request *DescribeEipGatewayInfoRequest) (response *DescribeEipGatewayInfoResponse, err error) {
	response = CreateDescribeEipGatewayInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEipGatewayInfoWithChan invokes the vpc.DescribeEipGatewayInfo API asynchronously
func (client *Client) DescribeEipGatewayInfoWithChan(request *DescribeEipGatewayInfoRequest) (<-chan *DescribeEipGatewayInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeEipGatewayInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEipGatewayInfo(request)
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

// DescribeEipGatewayInfoWithCallback invokes the vpc.DescribeEipGatewayInfo API asynchronously
func (client *Client) DescribeEipGatewayInfoWithCallback(request *DescribeEipGatewayInfoRequest, callback func(response *DescribeEipGatewayInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEipGatewayInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeEipGatewayInfo(request)
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

// DescribeEipGatewayInfoRequest is the request struct for api DescribeEipGatewayInfo
type DescribeEipGatewayInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeEipGatewayInfoResponse is the response struct for api DescribeEipGatewayInfo
type DescribeEipGatewayInfoResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	EipInfos  EipInfos `json:"EipInfos" xml:"EipInfos"`
}

// CreateDescribeEipGatewayInfoRequest creates a request to invoke DescribeEipGatewayInfo API
func CreateDescribeEipGatewayInfoRequest() (request *DescribeEipGatewayInfoRequest) {
	request = &DescribeEipGatewayInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeEipGatewayInfo", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEipGatewayInfoResponse creates a response to parse from DescribeEipGatewayInfo response
func CreateDescribeEipGatewayInfoResponse() (response *DescribeEipGatewayInfoResponse) {
	response = &DescribeEipGatewayInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
