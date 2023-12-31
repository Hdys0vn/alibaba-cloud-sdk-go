package multimediaai

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

// DeleteFaceImage invokes the multimediaai.DeleteFaceImage API synchronously
// api document: https://help.aliyun.com/api/multimediaai/deletefaceimage.html
func (client *Client) DeleteFaceImage(request *DeleteFaceImageRequest) (response *DeleteFaceImageResponse, err error) {
	response = CreateDeleteFaceImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceImageWithChan invokes the multimediaai.DeleteFaceImage API asynchronously
// api document: https://help.aliyun.com/api/multimediaai/deletefaceimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFaceImageWithChan(request *DeleteFaceImageRequest) (<-chan *DeleteFaceImageResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceImage(request)
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

// DeleteFaceImageWithCallback invokes the multimediaai.DeleteFaceImage API asynchronously
// api document: https://help.aliyun.com/api/multimediaai/deletefaceimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFaceImageWithCallback(request *DeleteFaceImageRequest, callback func(response *DeleteFaceImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceImageResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceImage(request)
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

// DeleteFaceImageRequest is the request struct for api DeleteFaceImage
type DeleteFaceImageRequest struct {
	*requests.RpcRequest
	FaceGroupId  requests.Integer `position:"Query" name:"FaceGroupId"`
	FacePersonId requests.Integer `position:"Query" name:"FacePersonId"`
	FaceImageId  requests.Integer `position:"Query" name:"FaceImageId"`
}

// DeleteFaceImageResponse is the response struct for api DeleteFaceImage
type DeleteFaceImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFaceImageRequest creates a request to invoke DeleteFaceImage API
func CreateDeleteFaceImageRequest() (request *DeleteFaceImageRequest) {
	request = &DeleteFaceImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("multimediaai", "2019-08-10", "DeleteFaceImage", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteFaceImageResponse creates a response to parse from DeleteFaceImage response
func CreateDeleteFaceImageResponse() (response *DeleteFaceImageResponse) {
	response = &DeleteFaceImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
