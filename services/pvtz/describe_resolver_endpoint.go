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

// DescribeResolverEndpoint invokes the pvtz.DescribeResolverEndpoint API synchronously
func (client *Client) DescribeResolverEndpoint(request *DescribeResolverEndpointRequest) (response *DescribeResolverEndpointResponse, err error) {
	response = CreateDescribeResolverEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResolverEndpointWithChan invokes the pvtz.DescribeResolverEndpoint API asynchronously
func (client *Client) DescribeResolverEndpointWithChan(request *DescribeResolverEndpointRequest) (<-chan *DescribeResolverEndpointResponse, <-chan error) {
	responseChan := make(chan *DescribeResolverEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResolverEndpoint(request)
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

// DescribeResolverEndpointWithCallback invokes the pvtz.DescribeResolverEndpoint API asynchronously
func (client *Client) DescribeResolverEndpointWithCallback(request *DescribeResolverEndpointRequest, callback func(response *DescribeResolverEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResolverEndpointResponse
		var err error
		defer close(result)
		response, err = client.DescribeResolverEndpoint(request)
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

// DescribeResolverEndpointRequest is the request struct for api DescribeResolverEndpoint
type DescribeResolverEndpointRequest struct {
	*requests.RpcRequest
	EndpointId   string `position:"Query" name:"EndpointId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeResolverEndpointResponse is the response struct for api DescribeResolverEndpoint
type DescribeResolverEndpointResponse struct {
	*responses.BaseResponse
	Status          string     `json:"Status" xml:"Status"`
	RequestId       string     `json:"RequestId" xml:"RequestId"`
	SecurityGroupId string     `json:"SecurityGroupId" xml:"SecurityGroupId"`
	CreateTime      string     `json:"CreateTime" xml:"CreateTime"`
	Name            string     `json:"Name" xml:"Name"`
	VpcRegionName   string     `json:"VpcRegionName" xml:"VpcRegionName"`
	VpcId           string     `json:"VpcId" xml:"VpcId"`
	UpdateTime      string     `json:"UpdateTime" xml:"UpdateTime"`
	VpcRegionId     string     `json:"VpcRegionId" xml:"VpcRegionId"`
	VpcName         string     `json:"VpcName" xml:"VpcName"`
	UpdateTimestamp int64      `json:"UpdateTimestamp" xml:"UpdateTimestamp"`
	Id              string     `json:"Id" xml:"Id"`
	CreateTimestamp int64      `json:"CreateTimestamp" xml:"CreateTimestamp"`
	IpConfigs       []IpConfig `json:"IpConfigs" xml:"IpConfigs"`
}

// CreateDescribeResolverEndpointRequest creates a request to invoke DescribeResolverEndpoint API
func CreateDescribeResolverEndpointRequest() (request *DescribeResolverEndpointRequest) {
	request = &DescribeResolverEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "DescribeResolverEndpoint", "pvtz", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeResolverEndpointResponse creates a response to parse from DescribeResolverEndpoint response
func CreateDescribeResolverEndpointResponse() (response *DescribeResolverEndpointResponse) {
	response = &DescribeResolverEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
