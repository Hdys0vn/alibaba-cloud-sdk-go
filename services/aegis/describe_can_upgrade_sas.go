package aegis

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

// DescribeCanUpgradeSas invokes the aegis.DescribeCanUpgradeSas API synchronously
// api document: https://help.aliyun.com/api/aegis/describecanupgradesas.html
func (client *Client) DescribeCanUpgradeSas(request *DescribeCanUpgradeSasRequest) (response *DescribeCanUpgradeSasResponse, err error) {
	response = CreateDescribeCanUpgradeSasResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCanUpgradeSasWithChan invokes the aegis.DescribeCanUpgradeSas API asynchronously
// api document: https://help.aliyun.com/api/aegis/describecanupgradesas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCanUpgradeSasWithChan(request *DescribeCanUpgradeSasRequest) (<-chan *DescribeCanUpgradeSasResponse, <-chan error) {
	responseChan := make(chan *DescribeCanUpgradeSasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCanUpgradeSas(request)
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

// DescribeCanUpgradeSasWithCallback invokes the aegis.DescribeCanUpgradeSas API asynchronously
// api document: https://help.aliyun.com/api/aegis/describecanupgradesas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCanUpgradeSasWithCallback(request *DescribeCanUpgradeSasRequest, callback func(response *DescribeCanUpgradeSasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCanUpgradeSasResponse
		var err error
		defer close(result)
		response, err = client.DescribeCanUpgradeSas(request)
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

// DescribeCanUpgradeSasRequest is the request struct for api DescribeCanUpgradeSas
type DescribeCanUpgradeSasRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeCanUpgradeSasResponse is the response struct for api DescribeCanUpgradeSas
type DescribeCanUpgradeSasResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDescribeCanUpgradeSasRequest creates a request to invoke DescribeCanUpgradeSas API
func CreateDescribeCanUpgradeSasRequest() (request *DescribeCanUpgradeSasRequest) {
	request = &DescribeCanUpgradeSasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeCanUpgradeSas", "vipaegis", "openAPI")
	return
}

// CreateDescribeCanUpgradeSasResponse creates a response to parse from DescribeCanUpgradeSas response
func CreateDescribeCanUpgradeSasResponse() (response *DescribeCanUpgradeSasResponse) {
	response = &DescribeCanUpgradeSasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
