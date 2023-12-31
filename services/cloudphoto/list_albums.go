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

// ListAlbums invokes the cloudphoto.ListAlbums API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/listalbums.html
func (client *Client) ListAlbums(request *ListAlbumsRequest) (response *ListAlbumsResponse, err error) {
	response = CreateListAlbumsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAlbumsWithChan invokes the cloudphoto.ListAlbums API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/listalbums.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAlbumsWithChan(request *ListAlbumsRequest) (<-chan *ListAlbumsResponse, <-chan error) {
	responseChan := make(chan *ListAlbumsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAlbums(request)
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

// ListAlbumsWithCallback invokes the cloudphoto.ListAlbums API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/listalbums.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAlbumsWithCallback(request *ListAlbumsRequest, callback func(response *ListAlbumsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAlbumsResponse
		var err error
		defer close(result)
		response, err = client.ListAlbums(request)
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

// ListAlbumsRequest is the request struct for api ListAlbums
type ListAlbumsRequest struct {
	*requests.RpcRequest
	Cursor    string           `position:"Query" name:"Cursor"`
	Size      requests.Integer `position:"Query" name:"Size"`
	LibraryId string           `position:"Query" name:"LibraryId"`
	StoreName string           `position:"Query" name:"StoreName"`
	State     string           `position:"Query" name:"State"`
	Direction string           `position:"Query" name:"Direction"`
}

// ListAlbumsResponse is the response struct for api ListAlbums
type ListAlbumsResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	NextCursor string  `json:"NextCursor" xml:"NextCursor"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	Action     string  `json:"Action" xml:"Action"`
	Albums     []Album `json:"Albums" xml:"Albums"`
}

// CreateListAlbumsRequest creates a request to invoke ListAlbums API
func CreateListAlbumsRequest() (request *ListAlbumsRequest) {
	request = &ListAlbumsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListAlbums", "cloudphoto", "openAPI")
	return
}

// CreateListAlbumsResponse creates a response to parse from ListAlbums response
func CreateListAlbumsResponse() (response *ListAlbumsResponse) {
	response = &ListAlbumsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
