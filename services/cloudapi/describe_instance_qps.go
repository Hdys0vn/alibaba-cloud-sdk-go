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

// DescribeInstanceQps invokes the cloudapi.DescribeInstanceQps API synchronously
func (client *Client) DescribeInstanceQps(request *DescribeInstanceQpsRequest) (response *DescribeInstanceQpsResponse, err error) {
	response = CreateDescribeInstanceQpsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceQpsWithChan invokes the cloudapi.DescribeInstanceQps API asynchronously
func (client *Client) DescribeInstanceQpsWithChan(request *DescribeInstanceQpsRequest) (<-chan *DescribeInstanceQpsResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceQpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceQps(request)
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

// DescribeInstanceQpsWithCallback invokes the cloudapi.DescribeInstanceQps API asynchronously
func (client *Client) DescribeInstanceQpsWithCallback(request *DescribeInstanceQpsRequest, callback func(response *DescribeInstanceQpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceQpsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceQps(request)
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

// DescribeInstanceQpsRequest is the request struct for api DescribeInstanceQps
type DescribeInstanceQpsRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeInstanceQpsResponse is the response struct for api DescribeInstanceQps
type DescribeInstanceQpsResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	InstanceQps InstanceQps `json:"InstanceQps" xml:"InstanceQps"`
}

// CreateDescribeInstanceQpsRequest creates a request to invoke DescribeInstanceQps API
func CreateDescribeInstanceQpsRequest() (request *DescribeInstanceQpsRequest) {
	request = &DescribeInstanceQpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeInstanceQps", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceQpsResponse creates a response to parse from DescribeInstanceQps response
func CreateDescribeInstanceQpsResponse() (response *DescribeInstanceQpsResponse) {
	response = &DescribeInstanceQpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
