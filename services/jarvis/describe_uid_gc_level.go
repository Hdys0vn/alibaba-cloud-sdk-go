package jarvis

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

// DescribeUidGcLevel invokes the jarvis.DescribeUidGcLevel API synchronously
// api document: https://help.aliyun.com/api/jarvis/describeuidgclevel.html
func (client *Client) DescribeUidGcLevel(request *DescribeUidGcLevelRequest) (response *DescribeUidGcLevelResponse, err error) {
	response = CreateDescribeUidGcLevelResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUidGcLevelWithChan invokes the jarvis.DescribeUidGcLevel API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describeuidgclevel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUidGcLevelWithChan(request *DescribeUidGcLevelRequest) (<-chan *DescribeUidGcLevelResponse, <-chan error) {
	responseChan := make(chan *DescribeUidGcLevelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUidGcLevel(request)
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

// DescribeUidGcLevelWithCallback invokes the jarvis.DescribeUidGcLevel API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describeuidgclevel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUidGcLevelWithCallback(request *DescribeUidGcLevelRequest, callback func(response *DescribeUidGcLevelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUidGcLevelResponse
		var err error
		defer close(result)
		response, err = client.DescribeUidGcLevel(request)
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

// DescribeUidGcLevelRequest is the request struct for api DescribeUidGcLevel
type DescribeUidGcLevelRequest struct {
	*requests.RpcRequest
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
	SourceCode string `position:"Query" name:"SourceCode"`
}

// DescribeUidGcLevelResponse is the response struct for api DescribeUidGcLevel
type DescribeUidGcLevelResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
	Gclevel   string `json:"Gclevel" xml:"Gclevel"`
}

// CreateDescribeUidGcLevelRequest creates a request to invoke DescribeUidGcLevel API
func CreateDescribeUidGcLevelRequest() (request *DescribeUidGcLevelRequest) {
	request = &DescribeUidGcLevelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DescribeUidGcLevel", "jarvis", "openAPI")
	return
}

// CreateDescribeUidGcLevelResponse creates a response to parse from DescribeUidGcLevel response
func CreateDescribeUidGcLevelResponse() (response *DescribeUidGcLevelResponse) {
	response = &DescribeUidGcLevelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
