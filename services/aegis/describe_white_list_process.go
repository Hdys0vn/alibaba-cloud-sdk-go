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

// DescribeWhiteListProcess invokes the aegis.DescribeWhiteListProcess API synchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelistprocess.html
func (client *Client) DescribeWhiteListProcess(request *DescribeWhiteListProcessRequest) (response *DescribeWhiteListProcessResponse, err error) {
	response = CreateDescribeWhiteListProcessResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWhiteListProcessWithChan invokes the aegis.DescribeWhiteListProcess API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelistprocess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWhiteListProcessWithChan(request *DescribeWhiteListProcessRequest) (<-chan *DescribeWhiteListProcessResponse, <-chan error) {
	responseChan := make(chan *DescribeWhiteListProcessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWhiteListProcess(request)
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

// DescribeWhiteListProcessWithCallback invokes the aegis.DescribeWhiteListProcess API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelistprocess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWhiteListProcessWithCallback(request *DescribeWhiteListProcessRequest, callback func(response *DescribeWhiteListProcessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWhiteListProcessResponse
		var err error
		defer close(result)
		response, err = client.DescribeWhiteListProcess(request)
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

// DescribeWhiteListProcessRequest is the request struct for api DescribeWhiteListProcess
type DescribeWhiteListProcessRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	ProcessName string           `position:"Query" name:"ProcessName"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ProcessType requests.Integer `position:"Query" name:"ProcessType"`
	OrderBy     requests.Integer `position:"Query" name:"OrderBy"`
	StrategyId  requests.Integer `position:"Query" name:"StrategyId"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Lang        string           `position:"Query" name:"Lang"`
	Desc        requests.Integer `position:"Query" name:"Desc"`
}

// DescribeWhiteListProcessResponse is the response struct for api DescribeWhiteListProcess
type DescribeWhiteListProcessResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	Count       int       `json:"Count" xml:"Count"`
	PageSize    int       `json:"PageSize" xml:"PageSize"`
	TotalCount  int       `json:"TotalCount" xml:"TotalCount"`
	CurrentPage int       `json:"CurrentPage" xml:"CurrentPage"`
	Processes   []Process `json:"Processes" xml:"Processes"`
}

// CreateDescribeWhiteListProcessRequest creates a request to invoke DescribeWhiteListProcess API
func CreateDescribeWhiteListProcessRequest() (request *DescribeWhiteListProcessRequest) {
	request = &DescribeWhiteListProcessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeWhiteListProcess", "vipaegis", "openAPI")
	return
}

// CreateDescribeWhiteListProcessResponse creates a response to parse from DescribeWhiteListProcess response
func CreateDescribeWhiteListProcessResponse() (response *DescribeWhiteListProcessResponse) {
	response = &DescribeWhiteListProcessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
