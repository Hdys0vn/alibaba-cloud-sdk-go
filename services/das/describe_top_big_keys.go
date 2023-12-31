package das

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

// DescribeTopBigKeys invokes the das.DescribeTopBigKeys API synchronously
func (client *Client) DescribeTopBigKeys(request *DescribeTopBigKeysRequest) (response *DescribeTopBigKeysResponse, err error) {
	response = CreateDescribeTopBigKeysResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTopBigKeysWithChan invokes the das.DescribeTopBigKeys API asynchronously
func (client *Client) DescribeTopBigKeysWithChan(request *DescribeTopBigKeysRequest) (<-chan *DescribeTopBigKeysResponse, <-chan error) {
	responseChan := make(chan *DescribeTopBigKeysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTopBigKeys(request)
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

// DescribeTopBigKeysWithCallback invokes the das.DescribeTopBigKeys API asynchronously
func (client *Client) DescribeTopBigKeysWithCallback(request *DescribeTopBigKeysRequest, callback func(response *DescribeTopBigKeysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTopBigKeysResponse
		var err error
		defer close(result)
		response, err = client.DescribeTopBigKeys(request)
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

// DescribeTopBigKeysRequest is the request struct for api DescribeTopBigKeys
type DescribeTopBigKeysRequest struct {
	*requests.RpcRequest
	EndTime        string `position:"Query" name:"EndTime"`
	StartTime      string `position:"Query" name:"StartTime"`
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	NodeId         string `position:"Query" name:"NodeId"`
}

// DescribeTopBigKeysResponse is the response struct for api DescribeTopBigKeys
type DescribeTopBigKeysResponse struct {
	*responses.BaseResponse
	Message   string                   `json:"Message" xml:"Message"`
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Code      string                   `json:"Code" xml:"Code"`
	Success   string                   `json:"Success" xml:"Success"`
	Data      DataInDescribeTopBigKeys `json:"Data" xml:"Data"`
}

// CreateDescribeTopBigKeysRequest creates a request to invoke DescribeTopBigKeys API
func CreateDescribeTopBigKeysRequest() (request *DescribeTopBigKeysRequest) {
	request = &DescribeTopBigKeysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "DescribeTopBigKeys", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeTopBigKeysResponse creates a response to parse from DescribeTopBigKeys response
func CreateDescribeTopBigKeysResponse() (response *DescribeTopBigKeysResponse) {
	response = &DescribeTopBigKeysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
