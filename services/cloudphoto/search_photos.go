package cloudphoto

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

// SearchPhotos invokes the cloudphoto.SearchPhotos API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/searchphotos.html
func (client *Client) SearchPhotos(request *SearchPhotosRequest) (response *SearchPhotosResponse, err error) {
	response = CreateSearchPhotosResponse()
	err = client.DoAction(request, response)
	return
}

// SearchPhotosWithChan invokes the cloudphoto.SearchPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/searchphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchPhotosWithChan(request *SearchPhotosRequest) (<-chan *SearchPhotosResponse, <-chan error) {
	responseChan := make(chan *SearchPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchPhotos(request)
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

// SearchPhotosWithCallback invokes the cloudphoto.SearchPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/searchphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchPhotosWithCallback(request *SearchPhotosRequest, callback func(response *SearchPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchPhotosResponse
		var err error
		defer close(result)
		response, err = client.SearchPhotos(request)
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

// SearchPhotosRequest is the request struct for api SearchPhotos
type SearchPhotosRequest struct {
	*requests.RpcRequest
	Size      requests.Integer `position:"Query" name:"Size"`
	LibraryId string           `position:"Query" name:"LibraryId"`
	StoreName string           `position:"Query" name:"StoreName"`
	Page      requests.Integer `position:"Query" name:"Page"`
	Keyword   string           `position:"Query" name:"Keyword"`
}

// SearchPhotosResponse is the response struct for api SearchPhotos
type SearchPhotosResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	Action     string  `json:"Action" xml:"Action"`
	Photos     []Photo `json:"Photos" xml:"Photos"`
}

// CreateSearchPhotosRequest creates a request to invoke SearchPhotos API
func CreateSearchPhotosRequest() (request *SearchPhotosRequest) {
	request = &SearchPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "SearchPhotos", "cloudphoto", "openAPI")
	return
}

// CreateSearchPhotosResponse creates a response to parse from SearchPhotos response
func CreateSearchPhotosResponse() (response *SearchPhotosResponse) {
	response = &SearchPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
