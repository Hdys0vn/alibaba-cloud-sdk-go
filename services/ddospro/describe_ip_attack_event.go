package ddospro

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

// DescribeIpAttackEvent invokes the ddospro.DescribeIpAttackEvent API synchronously
// api document: https://help.aliyun.com/api/ddospro/describeipattackevent.html
func (client *Client) DescribeIpAttackEvent(request *DescribeIpAttackEventRequest) (response *DescribeIpAttackEventResponse, err error) {
	response = CreateDescribeIpAttackEventResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpAttackEventWithChan invokes the ddospro.DescribeIpAttackEvent API asynchronously
// api document: https://help.aliyun.com/api/ddospro/describeipattackevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIpAttackEventWithChan(request *DescribeIpAttackEventRequest) (<-chan *DescribeIpAttackEventResponse, <-chan error) {
	responseChan := make(chan *DescribeIpAttackEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpAttackEvent(request)
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

// DescribeIpAttackEventWithCallback invokes the ddospro.DescribeIpAttackEvent API asynchronously
// api document: https://help.aliyun.com/api/ddospro/describeipattackevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIpAttackEventWithCallback(request *DescribeIpAttackEventRequest, callback func(response *DescribeIpAttackEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpAttackEventResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpAttackEvent(request)
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

// DescribeIpAttackEventRequest is the request struct for api DescribeIpAttackEvent
type DescribeIpAttackEventRequest struct {
	*requests.RpcRequest
	StartDateMillis requests.Integer `position:"Query" name:"StartDateMillis"`
	EndDateMillis   requests.Integer `position:"Query" name:"EndDateMillis"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Start           requests.Integer `position:"Query" name:"Start"`
	Ip              string           `position:"Query" name:"Ip"`
}

// DescribeIpAttackEventResponse is the response struct for api DescribeIpAttackEvent
type DescribeIpAttackEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeIpAttackEventRequest creates a request to invoke DescribeIpAttackEvent API
func CreateDescribeIpAttackEventRequest() (request *DescribeIpAttackEventRequest) {
	request = &DescribeIpAttackEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DescribeIpAttackEvent", "", "")
	return
}

// CreateDescribeIpAttackEventResponse creates a response to parse from DescribeIpAttackEvent response
func CreateDescribeIpAttackEventResponse() (response *DescribeIpAttackEventResponse) {
	response = &DescribeIpAttackEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
