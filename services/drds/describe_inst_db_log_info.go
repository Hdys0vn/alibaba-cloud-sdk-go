package drds

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

// DescribeInstDbLogInfo invokes the drds.DescribeInstDbLogInfo API synchronously
func (client *Client) DescribeInstDbLogInfo(request *DescribeInstDbLogInfoRequest) (response *DescribeInstDbLogInfoResponse, err error) {
	response = CreateDescribeInstDbLogInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstDbLogInfoWithChan invokes the drds.DescribeInstDbLogInfo API asynchronously
func (client *Client) DescribeInstDbLogInfoWithChan(request *DescribeInstDbLogInfoRequest) (<-chan *DescribeInstDbLogInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeInstDbLogInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstDbLogInfo(request)
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

// DescribeInstDbLogInfoWithCallback invokes the drds.DescribeInstDbLogInfo API asynchronously
func (client *Client) DescribeInstDbLogInfoWithCallback(request *DescribeInstDbLogInfoRequest, callback func(response *DescribeInstDbLogInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstDbLogInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstDbLogInfo(request)
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

// DescribeInstDbLogInfoRequest is the request struct for api DescribeInstDbLogInfo
type DescribeInstDbLogInfoRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// DescribeInstDbLogInfoResponse is the response struct for api DescribeInstDbLogInfo
type DescribeInstDbLogInfoResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Success      bool         `json:"Success" xml:"Success"`
	LogTimeRange LogTimeRange `json:"LogTimeRange" xml:"LogTimeRange"`
}

// CreateDescribeInstDbLogInfoRequest creates a request to invoke DescribeInstDbLogInfo API
func CreateDescribeInstDbLogInfoRequest() (request *DescribeInstDbLogInfoRequest) {
	request = &DescribeInstDbLogInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeInstDbLogInfo", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstDbLogInfoResponse creates a response to parse from DescribeInstDbLogInfo response
func CreateDescribeInstDbLogInfoResponse() (response *DescribeInstDbLogInfoResponse) {
	response = &DescribeInstDbLogInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
