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

// DescribeCheckWarningDetail invokes the aegis.DescribeCheckWarningDetail API synchronously
// api document: https://help.aliyun.com/api/aegis/describecheckwarningdetail.html
func (client *Client) DescribeCheckWarningDetail(request *DescribeCheckWarningDetailRequest) (response *DescribeCheckWarningDetailResponse, err error) {
	response = CreateDescribeCheckWarningDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCheckWarningDetailWithChan invokes the aegis.DescribeCheckWarningDetail API asynchronously
// api document: https://help.aliyun.com/api/aegis/describecheckwarningdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCheckWarningDetailWithChan(request *DescribeCheckWarningDetailRequest) (<-chan *DescribeCheckWarningDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeCheckWarningDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCheckWarningDetail(request)
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

// DescribeCheckWarningDetailWithCallback invokes the aegis.DescribeCheckWarningDetail API asynchronously
// api document: https://help.aliyun.com/api/aegis/describecheckwarningdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCheckWarningDetailWithCallback(request *DescribeCheckWarningDetailRequest, callback func(response *DescribeCheckWarningDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCheckWarningDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeCheckWarningDetail(request)
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

// DescribeCheckWarningDetailRequest is the request struct for api DescribeCheckWarningDetail
type DescribeCheckWarningDetailRequest struct {
	*requests.RpcRequest
	SourceIp       string           `position:"Query" name:"SourceIp"`
	Lang           string           `position:"Query" name:"Lang"`
	CheckWarningId requests.Integer `position:"Query" name:"CheckWarningId"`
}

// DescribeCheckWarningDetailResponse is the response struct for api DescribeCheckWarningDetail
type DescribeCheckWarningDetailResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	CheckId     int    `json:"CheckId" xml:"CheckId"`
	Level       string `json:"Level" xml:"Level"`
	Item        string `json:"Item" xml:"Item"`
	Prompt      string `json:"Prompt" xml:"Prompt"`
	Type        string `json:"Type" xml:"Type"`
	Advice      string `json:"Advice" xml:"Advice"`
	Description string `json:"Description" xml:"Description"`
}

// CreateDescribeCheckWarningDetailRequest creates a request to invoke DescribeCheckWarningDetail API
func CreateDescribeCheckWarningDetailRequest() (request *DescribeCheckWarningDetailRequest) {
	request = &DescribeCheckWarningDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeCheckWarningDetail", "vipaegis", "openAPI")
	return
}

// CreateDescribeCheckWarningDetailResponse creates a response to parse from DescribeCheckWarningDetail response
func CreateDescribeCheckWarningDetailResponse() (response *DescribeCheckWarningDetailResponse) {
	response = &DescribeCheckWarningDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
