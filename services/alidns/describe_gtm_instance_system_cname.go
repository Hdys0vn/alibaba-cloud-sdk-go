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

// DescribeGtmInstanceSystemCname invokes the alidns.DescribeGtmInstanceSystemCname API synchronously
func (client *Client) DescribeGtmInstanceSystemCname(request *DescribeGtmInstanceSystemCnameRequest) (response *DescribeGtmInstanceSystemCnameResponse, err error) {
	response = CreateDescribeGtmInstanceSystemCnameResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGtmInstanceSystemCnameWithChan invokes the alidns.DescribeGtmInstanceSystemCname API asynchronously
func (client *Client) DescribeGtmInstanceSystemCnameWithChan(request *DescribeGtmInstanceSystemCnameRequest) (<-chan *DescribeGtmInstanceSystemCnameResponse, <-chan error) {
	responseChan := make(chan *DescribeGtmInstanceSystemCnameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGtmInstanceSystemCname(request)
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

// DescribeGtmInstanceSystemCnameWithCallback invokes the alidns.DescribeGtmInstanceSystemCname API asynchronously
func (client *Client) DescribeGtmInstanceSystemCnameWithCallback(request *DescribeGtmInstanceSystemCnameRequest, callback func(response *DescribeGtmInstanceSystemCnameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGtmInstanceSystemCnameResponse
		var err error
		defer close(result)
		response, err = client.DescribeGtmInstanceSystemCname(request)
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

// DescribeGtmInstanceSystemCnameRequest is the request struct for api DescribeGtmInstanceSystemCname
type DescribeGtmInstanceSystemCnameRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeGtmInstanceSystemCnameResponse is the response struct for api DescribeGtmInstanceSystemCname
type DescribeGtmInstanceSystemCnameResponse struct {
	*responses.BaseResponse
	SystemCname string `json:"SystemCname" xml:"SystemCname"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeGtmInstanceSystemCnameRequest creates a request to invoke DescribeGtmInstanceSystemCname API
func CreateDescribeGtmInstanceSystemCnameRequest() (request *DescribeGtmInstanceSystemCnameRequest) {
	request = &DescribeGtmInstanceSystemCnameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeGtmInstanceSystemCname", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGtmInstanceSystemCnameResponse creates a response to parse from DescribeGtmInstanceSystemCname response
func CreateDescribeGtmInstanceSystemCnameResponse() (response *DescribeGtmInstanceSystemCnameResponse) {
	response = &DescribeGtmInstanceSystemCnameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
