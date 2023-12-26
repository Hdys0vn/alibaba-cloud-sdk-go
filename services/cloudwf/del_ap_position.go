package cloudwf

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

// DelApPosition invokes the cloudwf.DelApPosition API synchronously
// api document: https://help.aliyun.com/api/cloudwf/delapposition.html
func (client *Client) DelApPosition(request *DelApPositionRequest) (response *DelApPositionResponse, err error) {
	response = CreateDelApPositionResponse()
	err = client.DoAction(request, response)
	return
}

// DelApPositionWithChan invokes the cloudwf.DelApPosition API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/delapposition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DelApPositionWithChan(request *DelApPositionRequest) (<-chan *DelApPositionResponse, <-chan error) {
	responseChan := make(chan *DelApPositionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DelApPosition(request)
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

// DelApPositionWithCallback invokes the cloudwf.DelApPosition API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/delapposition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DelApPositionWithCallback(request *DelApPositionRequest, callback func(response *DelApPositionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DelApPositionResponse
		var err error
		defer close(result)
		response, err = client.DelApPosition(request)
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

// DelApPositionRequest is the request struct for api DelApPosition
type DelApPositionRequest struct {
	*requests.RpcRequest
	ApAssetId requests.Integer `position:"Query" name:"ApAssetId"`
	MapId     requests.Integer `position:"Query" name:"MapId"`
}

// DelApPositionResponse is the response struct for api DelApPosition
type DelApPositionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateDelApPositionRequest creates a request to invoke DelApPosition API
func CreateDelApPositionRequest() (request *DelApPositionRequest) {
	request = &DelApPositionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "DelApPosition", "cloudwf", "openAPI")
	return
}

// CreateDelApPositionResponse creates a response to parse from DelApPosition response
func CreateDelApPositionResponse() (response *DelApPositionResponse) {
	response = &DelApPositionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
