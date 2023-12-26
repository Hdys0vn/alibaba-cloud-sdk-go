package aiworkspace

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

// RemoveImage invokes the aiworkspace.RemoveImage API synchronously
func (client *Client) RemoveImage(request *RemoveImageRequest) (response *RemoveImageResponse, err error) {
	response = CreateRemoveImageResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveImageWithChan invokes the aiworkspace.RemoveImage API asynchronously
func (client *Client) RemoveImageWithChan(request *RemoveImageRequest) (<-chan *RemoveImageResponse, <-chan error) {
	responseChan := make(chan *RemoveImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveImage(request)
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

// RemoveImageWithCallback invokes the aiworkspace.RemoveImage API asynchronously
func (client *Client) RemoveImageWithCallback(request *RemoveImageRequest, callback func(response *RemoveImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveImageResponse
		var err error
		defer close(result)
		response, err = client.RemoveImage(request)
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

// RemoveImageRequest is the request struct for api RemoveImage
type RemoveImageRequest struct {
	*requests.RoaRequest
	ImageId string `position:"Path" name:"ImageId"`
}

// RemoveImageResponse is the response struct for api RemoveImage
type RemoveImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveImageRequest creates a request to invoke RemoveImage API
func CreateRemoveImageRequest() (request *RemoveImageRequest) {
	request = &RemoveImageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AIWorkSpace", "2021-02-04", "RemoveImage", "/api/v1/images/[ImageId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateRemoveImageResponse creates a response to parse from RemoveImage response
func CreateRemoveImageResponse() (response *RemoveImageResponse) {
	response = &RemoveImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
