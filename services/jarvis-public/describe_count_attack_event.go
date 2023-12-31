package jarvis_public

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

// DescribeCountAttackEvent invokes the jarvis_public.DescribeCountAttackEvent API synchronously
// api document: https://help.aliyun.com/api/jarvis-public/describecountattackevent.html
func (client *Client) DescribeCountAttackEvent(request *DescribeCountAttackEventRequest) (response *DescribeCountAttackEventResponse, err error) {
	response = CreateDescribeCountAttackEventResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCountAttackEventWithChan invokes the jarvis_public.DescribeCountAttackEvent API asynchronously
// api document: https://help.aliyun.com/api/jarvis-public/describecountattackevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCountAttackEventWithChan(request *DescribeCountAttackEventRequest) (<-chan *DescribeCountAttackEventResponse, <-chan error) {
	responseChan := make(chan *DescribeCountAttackEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCountAttackEvent(request)
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

// DescribeCountAttackEventWithCallback invokes the jarvis_public.DescribeCountAttackEvent API asynchronously
// api document: https://help.aliyun.com/api/jarvis-public/describecountattackevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCountAttackEventWithCallback(request *DescribeCountAttackEventRequest, callback func(response *DescribeCountAttackEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCountAttackEventResponse
		var err error
		defer close(result)
		response, err = client.DescribeCountAttackEvent(request)
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

// DescribeCountAttackEventRequest is the request struct for api DescribeCountAttackEvent
type DescribeCountAttackEventRequest struct {
	*requests.RpcRequest
	SourceIp     string           `position:"Query" name:"SourceIp"`
	ServerIpList string           `position:"Query" name:"ServerIpList"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	StartTime    requests.Integer `position:"Query" name:"StartTime"`
	Lang         string           `position:"Query" name:"Lang"`
	Region       string           `position:"Query" name:"Region"`
	ProductType  string           `position:"Query" name:"ProductType"`
}

// DescribeCountAttackEventResponse is the response struct for api DescribeCountAttackEvent
type DescribeCountAttackEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
	Count     int    `json:"Count" xml:"Count"`
}

// CreateDescribeCountAttackEventRequest creates a request to invoke DescribeCountAttackEvent API
func CreateDescribeCountAttackEventRequest() (request *DescribeCountAttackEventRequest) {
	request = &DescribeCountAttackEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis-public", "2018-06-21", "DescribeCountAttackEvent", "jarvis-public", "openAPI")
	return
}

// CreateDescribeCountAttackEventResponse creates a response to parse from DescribeCountAttackEvent response
func CreateDescribeCountAttackEventResponse() (response *DescribeCountAttackEventResponse) {
	response = &DescribeCountAttackEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
