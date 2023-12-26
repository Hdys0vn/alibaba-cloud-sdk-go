package cloudcallcenter

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

// DescribeVnRecording invokes the cloudcallcenter.DescribeVnRecording API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnrecording.html
func (client *Client) DescribeVnRecording(request *DescribeVnRecordingRequest) (response *DescribeVnRecordingResponse, err error) {
	response = CreateDescribeVnRecordingResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVnRecordingWithChan invokes the cloudcallcenter.DescribeVnRecording API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnrecording.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVnRecordingWithChan(request *DescribeVnRecordingRequest) (<-chan *DescribeVnRecordingResponse, <-chan error) {
	responseChan := make(chan *DescribeVnRecordingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVnRecording(request)
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

// DescribeVnRecordingWithCallback invokes the cloudcallcenter.DescribeVnRecording API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnrecording.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVnRecordingWithCallback(request *DescribeVnRecordingRequest, callback func(response *DescribeVnRecordingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVnRecordingResponse
		var err error
		defer close(result)
		response, err = client.DescribeVnRecording(request)
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

// DescribeVnRecordingRequest is the request struct for api DescribeVnRecording
type DescribeVnRecordingRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// DescribeVnRecordingResponse is the response struct for api DescribeVnRecording
type DescribeVnRecordingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	FileName  string `json:"FileName" xml:"FileName"`
	FilePath  string `json:"FilePath" xml:"FilePath"`
}

// CreateDescribeVnRecordingRequest creates a request to invoke DescribeVnRecording API
func CreateDescribeVnRecordingRequest() (request *DescribeVnRecordingRequest) {
	request = &DescribeVnRecordingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DescribeVnRecording", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeVnRecordingResponse creates a response to parse from DescribeVnRecording response
func CreateDescribeVnRecordingResponse() (response *DescribeVnRecordingResponse) {
	response = &DescribeVnRecordingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
