package cloudapi

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

// DescribeApiGroupVpcWhitelist invokes the cloudapi.DescribeApiGroupVpcWhitelist API synchronously
func (client *Client) DescribeApiGroupVpcWhitelist(request *DescribeApiGroupVpcWhitelistRequest) (response *DescribeApiGroupVpcWhitelistResponse, err error) {
	response = CreateDescribeApiGroupVpcWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiGroupVpcWhitelistWithChan invokes the cloudapi.DescribeApiGroupVpcWhitelist API asynchronously
func (client *Client) DescribeApiGroupVpcWhitelistWithChan(request *DescribeApiGroupVpcWhitelistRequest) (<-chan *DescribeApiGroupVpcWhitelistResponse, <-chan error) {
	responseChan := make(chan *DescribeApiGroupVpcWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiGroupVpcWhitelist(request)
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

// DescribeApiGroupVpcWhitelistWithCallback invokes the cloudapi.DescribeApiGroupVpcWhitelist API asynchronously
func (client *Client) DescribeApiGroupVpcWhitelistWithCallback(request *DescribeApiGroupVpcWhitelistRequest, callback func(response *DescribeApiGroupVpcWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiGroupVpcWhitelistResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiGroupVpcWhitelist(request)
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

// DescribeApiGroupVpcWhitelistRequest is the request struct for api DescribeApiGroupVpcWhitelist
type DescribeApiGroupVpcWhitelistRequest struct {
	*requests.RpcRequest
	GroupId       string `position:"Query" name:"GroupId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeApiGroupVpcWhitelistResponse is the response struct for api DescribeApiGroupVpcWhitelist
type DescribeApiGroupVpcWhitelistResponse struct {
	*responses.BaseResponse
	VpcIds    string `json:"VpcIds" xml:"VpcIds"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeApiGroupVpcWhitelistRequest creates a request to invoke DescribeApiGroupVpcWhitelist API
func CreateDescribeApiGroupVpcWhitelistRequest() (request *DescribeApiGroupVpcWhitelistRequest) {
	request = &DescribeApiGroupVpcWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiGroupVpcWhitelist", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeApiGroupVpcWhitelistResponse creates a response to parse from DescribeApiGroupVpcWhitelist response
func CreateDescribeApiGroupVpcWhitelistResponse() (response *DescribeApiGroupVpcWhitelistResponse) {
	response = &DescribeApiGroupVpcWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}