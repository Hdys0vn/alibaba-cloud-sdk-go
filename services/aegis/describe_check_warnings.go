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

// DescribeCheckWarnings invokes the aegis.DescribeCheckWarnings API synchronously
// api document: https://help.aliyun.com/api/aegis/describecheckwarnings.html
func (client *Client) DescribeCheckWarnings(request *DescribeCheckWarningsRequest) (response *DescribeCheckWarningsResponse, err error) {
	response = CreateDescribeCheckWarningsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCheckWarningsWithChan invokes the aegis.DescribeCheckWarnings API asynchronously
// api document: https://help.aliyun.com/api/aegis/describecheckwarnings.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCheckWarningsWithChan(request *DescribeCheckWarningsRequest) (<-chan *DescribeCheckWarningsResponse, <-chan error) {
	responseChan := make(chan *DescribeCheckWarningsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCheckWarnings(request)
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

// DescribeCheckWarningsWithCallback invokes the aegis.DescribeCheckWarnings API asynchronously
// api document: https://help.aliyun.com/api/aegis/describecheckwarnings.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCheckWarningsWithCallback(request *DescribeCheckWarningsRequest, callback func(response *DescribeCheckWarningsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCheckWarningsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCheckWarnings(request)
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

// DescribeCheckWarningsRequest is the request struct for api DescribeCheckWarnings
type DescribeCheckWarningsRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Lang        string           `position:"Query" name:"Lang"`
	RiskId      requests.Integer `position:"Query" name:"RiskId"`
	Uuid        string           `position:"Query" name:"Uuid"`
}

// DescribeCheckWarningsResponse is the response struct for api DescribeCheckWarnings
type DescribeCheckWarningsResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	Count         int            `json:"Count" xml:"Count"`
	PageSize      int            `json:"PageSize" xml:"PageSize"`
	TotalCount    int            `json:"TotalCount" xml:"TotalCount"`
	CurrentPage   int            `json:"CurrentPage" xml:"CurrentPage"`
	CheckWarnings []CheckWarning `json:"CheckWarnings" xml:"CheckWarnings"`
}

// CreateDescribeCheckWarningsRequest creates a request to invoke DescribeCheckWarnings API
func CreateDescribeCheckWarningsRequest() (request *DescribeCheckWarningsRequest) {
	request = &DescribeCheckWarningsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeCheckWarnings", "vipaegis", "openAPI")
	return
}

// CreateDescribeCheckWarningsResponse creates a response to parse from DescribeCheckWarnings response
func CreateDescribeCheckWarningsResponse() (response *DescribeCheckWarningsResponse) {
	response = &DescribeCheckWarningsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
