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

// DescribeApiBuySummary invokes the aegis.DescribeApiBuySummary API synchronously
// api document: https://help.aliyun.com/api/aegis/describeapibuysummary.html
func (client *Client) DescribeApiBuySummary(request *DescribeApiBuySummaryRequest) (response *DescribeApiBuySummaryResponse, err error) {
	response = CreateDescribeApiBuySummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiBuySummaryWithChan invokes the aegis.DescribeApiBuySummary API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeapibuysummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiBuySummaryWithChan(request *DescribeApiBuySummaryRequest) (<-chan *DescribeApiBuySummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeApiBuySummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiBuySummary(request)
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

// DescribeApiBuySummaryWithCallback invokes the aegis.DescribeApiBuySummary API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeapibuysummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiBuySummaryWithCallback(request *DescribeApiBuySummaryRequest, callback func(response *DescribeApiBuySummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiBuySummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiBuySummary(request)
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

// DescribeApiBuySummaryRequest is the request struct for api DescribeApiBuySummary
type DescribeApiBuySummaryRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeApiBuySummaryResponse is the response struct for api DescribeApiBuySummary
type DescribeApiBuySummaryResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	BuySummary BuySummary `json:"BuySummary" xml:"BuySummary"`
}

// CreateDescribeApiBuySummaryRequest creates a request to invoke DescribeApiBuySummary API
func CreateDescribeApiBuySummaryRequest() (request *DescribeApiBuySummaryRequest) {
	request = &DescribeApiBuySummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeApiBuySummary", "vipaegis", "openAPI")
	return
}

// CreateDescribeApiBuySummaryResponse creates a response to parse from DescribeApiBuySummary response
func CreateDescribeApiBuySummaryResponse() (response *DescribeApiBuySummaryResponse) {
	response = &DescribeApiBuySummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
