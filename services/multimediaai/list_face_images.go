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

// ListFaceImages invokes the multimediaai.ListFaceImages API synchronously
// api document: https://help.aliyun.com/api/multimediaai/listfaceimages.html
func (client *Client) ListFaceImages(request *ListFaceImagesRequest) (response *ListFaceImagesResponse, err error) {
	response = CreateListFaceImagesResponse()
	err = client.DoAction(request, response)
	return
}

// ListFaceImagesWithChan invokes the multimediaai.ListFaceImages API asynchronously
// api document: https://help.aliyun.com/api/multimediaai/listfaceimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFaceImagesWithChan(request *ListFaceImagesRequest) (<-chan *ListFaceImagesResponse, <-chan error) {
	responseChan := make(chan *ListFaceImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFaceImages(request)
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

// ListFaceImagesWithCallback invokes the multimediaai.ListFaceImages API asynchronously
// api document: https://help.aliyun.com/api/multimediaai/listfaceimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFaceImagesWithCallback(request *ListFaceImagesRequest, callback func(response *ListFaceImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFaceImagesResponse
		var err error
		defer close(result)
		response, err = client.ListFaceImages(request)
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

// ListFaceImagesRequest is the request struct for api ListFaceImages
type ListFaceImagesRequest struct {
	*requests.RpcRequest
	FaceGroupId  requests.Integer `position:"Query" name:"FaceGroupId"`
	FacePersonId requests.Integer `position:"Query" name:"FacePersonId"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListFaceImagesResponse is the response struct for api ListFaceImages
type ListFaceImagesResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	PageNumber int         `json:"PageNumber" xml:"PageNumber"`
	PageSize   int         `json:"PageSize" xml:"PageSize"`
	TotalCount int64       `json:"TotalCount" xml:"TotalCount"`
	FaceImages []FaceImage `json:"FaceImages" xml:"FaceImages"`
}

// CreateListFaceImagesRequest creates a request to invoke ListFaceImages API
func CreateListFaceImagesRequest() (request *ListFaceImagesRequest) {
	request = &ListFaceImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("multimediaai", "2019-08-10", "ListFaceImages", "", "")
	request.Method = requests.POST
	return
}

// CreateListFaceImagesResponse creates a response to parse from ListFaceImages response
func CreateListFaceImagesResponse() (response *ListFaceImagesResponse) {
	response = &ListFaceImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
