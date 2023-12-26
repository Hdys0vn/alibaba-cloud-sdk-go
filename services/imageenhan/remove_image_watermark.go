package imageenhan

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

// RemoveImageWatermark invokes the imageenhan.RemoveImageWatermark API synchronously
func (client *Client) RemoveImageWatermark(request *RemoveImageWatermarkRequest) (response *RemoveImageWatermarkResponse, err error) {
	response = CreateRemoveImageWatermarkResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveImageWatermarkWithChan invokes the imageenhan.RemoveImageWatermark API asynchronously
func (client *Client) RemoveImageWatermarkWithChan(request *RemoveImageWatermarkRequest) (<-chan *RemoveImageWatermarkResponse, <-chan error) {
	responseChan := make(chan *RemoveImageWatermarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveImageWatermark(request)
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

// RemoveImageWatermarkWithCallback invokes the imageenhan.RemoveImageWatermark API asynchronously
func (client *Client) RemoveImageWatermarkWithCallback(request *RemoveImageWatermarkRequest, callback func(response *RemoveImageWatermarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveImageWatermarkResponse
		var err error
		defer close(result)
		response, err = client.RemoveImageWatermark(request)
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

// RemoveImageWatermarkRequest is the request struct for api RemoveImageWatermark
type RemoveImageWatermarkRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// RemoveImageWatermarkResponse is the response struct for api RemoveImageWatermark
type RemoveImageWatermarkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRemoveImageWatermarkRequest creates a request to invoke RemoveImageWatermark API
func CreateRemoveImageWatermarkRequest() (request *RemoveImageWatermarkRequest) {
	request = &RemoveImageWatermarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageenhan", "2019-09-30", "RemoveImageWatermark", "imageenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveImageWatermarkResponse creates a response to parse from RemoveImageWatermark response
func CreateRemoveImageWatermarkResponse() (response *RemoveImageWatermarkResponse) {
	response = &RemoveImageWatermarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
