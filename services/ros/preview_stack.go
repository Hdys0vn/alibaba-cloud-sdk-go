package ros

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

// PreviewStack invokes the ros.PreviewStack API synchronously
// api document: https://help.aliyun.com/api/ros/previewstack.html
func (client *Client) PreviewStack(request *PreviewStackRequest) (response *PreviewStackResponse, err error) {
	response = CreatePreviewStackResponse()
	err = client.DoAction(request, response)
	return
}

// PreviewStackWithChan invokes the ros.PreviewStack API asynchronously
// api document: https://help.aliyun.com/api/ros/previewstack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreviewStackWithChan(request *PreviewStackRequest) (<-chan *PreviewStackResponse, <-chan error) {
	responseChan := make(chan *PreviewStackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreviewStack(request)
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

// PreviewStackWithCallback invokes the ros.PreviewStack API asynchronously
// api document: https://help.aliyun.com/api/ros/previewstack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreviewStackWithCallback(request *PreviewStackRequest, callback func(response *PreviewStackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreviewStackResponse
		var err error
		defer close(result)
		response, err = client.PreviewStack(request)
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

// PreviewStackRequest is the request struct for api PreviewStack
type PreviewStackRequest struct {
	*requests.RoaRequest
}

// PreviewStackResponse is the response struct for api PreviewStack
type PreviewStackResponse struct {
	*responses.BaseResponse
}

// CreatePreviewStackRequest creates a request to invoke PreviewStack API
func CreatePreviewStackRequest() (request *PreviewStackRequest) {
	request = &PreviewStackRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "PreviewStack", "/stacks/preview", "ROS", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePreviewStackResponse creates a response to parse from PreviewStack response
func CreatePreviewStackResponse() (response *PreviewStackResponse) {
	response = &PreviewStackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
