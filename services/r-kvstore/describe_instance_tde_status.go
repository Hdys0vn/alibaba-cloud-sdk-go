package r_kvstore

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

// DescribeInstanceTDEStatus invokes the r_kvstore.DescribeInstanceTDEStatus API synchronously
func (client *Client) DescribeInstanceTDEStatus(request *DescribeInstanceTDEStatusRequest) (response *DescribeInstanceTDEStatusResponse, err error) {
	response = CreateDescribeInstanceTDEStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceTDEStatusWithChan invokes the r_kvstore.DescribeInstanceTDEStatus API asynchronously
func (client *Client) DescribeInstanceTDEStatusWithChan(request *DescribeInstanceTDEStatusRequest) (<-chan *DescribeInstanceTDEStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceTDEStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceTDEStatus(request)
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

// DescribeInstanceTDEStatusWithCallback invokes the r_kvstore.DescribeInstanceTDEStatus API asynchronously
func (client *Client) DescribeInstanceTDEStatusWithCallback(request *DescribeInstanceTDEStatusRequest, callback func(response *DescribeInstanceTDEStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceTDEStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceTDEStatus(request)
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

// DescribeInstanceTDEStatusRequest is the request struct for api DescribeInstanceTDEStatus
type DescribeInstanceTDEStatusRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeInstanceTDEStatusResponse is the response struct for api DescribeInstanceTDEStatus
type DescribeInstanceTDEStatusResponse struct {
	*responses.BaseResponse
	TDEStatus string `json:"TDEStatus" xml:"TDEStatus"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeInstanceTDEStatusRequest creates a request to invoke DescribeInstanceTDEStatus API
func CreateDescribeInstanceTDEStatusRequest() (request *DescribeInstanceTDEStatusRequest) {
	request = &DescribeInstanceTDEStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstanceTDEStatus", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceTDEStatusResponse creates a response to parse from DescribeInstanceTDEStatus response
func CreateDescribeInstanceTDEStatusResponse() (response *DescribeInstanceTDEStatusResponse) {
	response = &DescribeInstanceTDEStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
