package ga

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

// DescribeIpSet invokes the ga.DescribeIpSet API synchronously
func (client *Client) DescribeIpSet(request *DescribeIpSetRequest) (response *DescribeIpSetResponse, err error) {
	response = CreateDescribeIpSetResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpSetWithChan invokes the ga.DescribeIpSet API asynchronously
func (client *Client) DescribeIpSetWithChan(request *DescribeIpSetRequest) (<-chan *DescribeIpSetResponse, <-chan error) {
	responseChan := make(chan *DescribeIpSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpSet(request)
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

// DescribeIpSetWithCallback invokes the ga.DescribeIpSet API asynchronously
func (client *Client) DescribeIpSetWithCallback(request *DescribeIpSetRequest, callback func(response *DescribeIpSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpSetResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpSet(request)
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

// DescribeIpSetRequest is the request struct for api DescribeIpSet
type DescribeIpSetRequest struct {
	*requests.RpcRequest
	IpSetId string `position:"Query" name:"IpSetId"`
}

// DescribeIpSetResponse is the response struct for api DescribeIpSet
type DescribeIpSetResponse struct {
	*responses.BaseResponse
	IpSetId            string   `json:"IpSetId" xml:"IpSetId"`
	RequestId          string   `json:"RequestId" xml:"RequestId"`
	IpVersion          string   `json:"IpVersion" xml:"IpVersion"`
	State              string   `json:"State" xml:"State"`
	Bandwidth          int      `json:"Bandwidth" xml:"Bandwidth"`
	AccelerateRegionId string   `json:"AccelerateRegionId" xml:"AccelerateRegionId"`
	AcceleratorId      string   `json:"AcceleratorId" xml:"AcceleratorId"`
	IpAddressList      []string `json:"IpAddressList" xml:"IpAddressList"`
}

// CreateDescribeIpSetRequest creates a request to invoke DescribeIpSet API
func CreateDescribeIpSetRequest() (request *DescribeIpSetRequest) {
	request = &DescribeIpSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "DescribeIpSet", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIpSetResponse creates a response to parse from DescribeIpSet response
func CreateDescribeIpSetResponse() (response *DescribeIpSetResponse) {
	response = &DescribeIpSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
