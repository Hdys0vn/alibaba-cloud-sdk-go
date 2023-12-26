package rtc

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

// DescribeRecordTemplates invokes the rtc.DescribeRecordTemplates API synchronously
func (client *Client) DescribeRecordTemplates(request *DescribeRecordTemplatesRequest) (response *DescribeRecordTemplatesResponse, err error) {
	response = CreateDescribeRecordTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordTemplatesWithChan invokes the rtc.DescribeRecordTemplates API asynchronously
func (client *Client) DescribeRecordTemplatesWithChan(request *DescribeRecordTemplatesRequest) (<-chan *DescribeRecordTemplatesResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordTemplates(request)
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

// DescribeRecordTemplatesWithCallback invokes the rtc.DescribeRecordTemplates API asynchronously
func (client *Client) DescribeRecordTemplatesWithCallback(request *DescribeRecordTemplatesRequest, callback func(response *DescribeRecordTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordTemplates(request)
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

// DescribeRecordTemplatesRequest is the request struct for api DescribeRecordTemplates
type DescribeRecordTemplatesRequest struct {
	*requests.RpcRequest
	TemplateIds *[]string        `position:"Query" name:"TemplateIds"  type:"Repeated"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	AppId       string           `position:"Query" name:"AppId"`
}

// DescribeRecordTemplatesResponse is the response struct for api DescribeRecordTemplates
type DescribeRecordTemplatesResponse struct {
	*responses.BaseResponse
	TotalPage int64      `json:"TotalPage" xml:"TotalPage"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	TotalNum  int64      `json:"TotalNum" xml:"TotalNum"`
	Templates []Template `json:"Templates" xml:"Templates"`
}

// CreateDescribeRecordTemplatesRequest creates a request to invoke DescribeRecordTemplates API
func CreateDescribeRecordTemplatesRequest() (request *DescribeRecordTemplatesRequest) {
	request = &DescribeRecordTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRecordTemplates", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRecordTemplatesResponse creates a response to parse from DescribeRecordTemplates response
func CreateDescribeRecordTemplatesResponse() (response *DescribeRecordTemplatesResponse) {
	response = &DescribeRecordTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
