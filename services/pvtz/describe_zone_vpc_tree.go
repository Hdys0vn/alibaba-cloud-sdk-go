package pvtz

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

// DescribeZoneVpcTree invokes the pvtz.DescribeZoneVpcTree API synchronously
func (client *Client) DescribeZoneVpcTree(request *DescribeZoneVpcTreeRequest) (response *DescribeZoneVpcTreeResponse, err error) {
	response = CreateDescribeZoneVpcTreeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeZoneVpcTreeWithChan invokes the pvtz.DescribeZoneVpcTree API asynchronously
func (client *Client) DescribeZoneVpcTreeWithChan(request *DescribeZoneVpcTreeRequest) (<-chan *DescribeZoneVpcTreeResponse, <-chan error) {
	responseChan := make(chan *DescribeZoneVpcTreeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeZoneVpcTree(request)
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

// DescribeZoneVpcTreeWithCallback invokes the pvtz.DescribeZoneVpcTree API asynchronously
func (client *Client) DescribeZoneVpcTreeWithCallback(request *DescribeZoneVpcTreeRequest, callback func(response *DescribeZoneVpcTreeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeZoneVpcTreeResponse
		var err error
		defer close(result)
		response, err = client.DescribeZoneVpcTree(request)
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

// DescribeZoneVpcTreeRequest is the request struct for api DescribeZoneVpcTree
type DescribeZoneVpcTreeRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeZoneVpcTreeResponse is the response struct for api DescribeZoneVpcTree
type DescribeZoneVpcTreeResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Zones     ZonesInDescribeZoneVpcTree `json:"Zones" xml:"Zones"`
}

// CreateDescribeZoneVpcTreeRequest creates a request to invoke DescribeZoneVpcTree API
func CreateDescribeZoneVpcTreeRequest() (request *DescribeZoneVpcTreeRequest) {
	request = &DescribeZoneVpcTreeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "DescribeZoneVpcTree", "pvtz", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeZoneVpcTreeResponse creates a response to parse from DescribeZoneVpcTree response
func CreateDescribeZoneVpcTreeResponse() (response *DescribeZoneVpcTreeResponse) {
	response = &DescribeZoneVpcTreeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
