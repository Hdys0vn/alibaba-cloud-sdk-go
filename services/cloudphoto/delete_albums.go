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

// DeleteAlbums invokes the cloudphoto.DeleteAlbums API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletealbums.html
func (client *Client) DeleteAlbums(request *DeleteAlbumsRequest) (response *DeleteAlbumsResponse, err error) {
	response = CreateDeleteAlbumsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAlbumsWithChan invokes the cloudphoto.DeleteAlbums API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletealbums.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAlbumsWithChan(request *DeleteAlbumsRequest) (<-chan *DeleteAlbumsResponse, <-chan error) {
	responseChan := make(chan *DeleteAlbumsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAlbums(request)
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

// DeleteAlbumsWithCallback invokes the cloudphoto.DeleteAlbums API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletealbums.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAlbumsWithCallback(request *DeleteAlbumsRequest, callback func(response *DeleteAlbumsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAlbumsResponse
		var err error
		defer close(result)
		response, err = client.DeleteAlbums(request)
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

// DeleteAlbumsRequest is the request struct for api DeleteAlbums
type DeleteAlbumsRequest struct {
	*requests.RpcRequest
	LibraryId string    `position:"Query" name:"LibraryId"`
	AlbumId   *[]string `position:"Query" name:"AlbumId"  type:"Repeated"`
	StoreName string    `position:"Query" name:"StoreName"`
}

// DeleteAlbumsResponse is the response struct for api DeleteAlbums
type DeleteAlbumsResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Action    string   `json:"Action" xml:"Action"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateDeleteAlbumsRequest creates a request to invoke DeleteAlbums API
func CreateDeleteAlbumsRequest() (request *DeleteAlbumsRequest) {
	request = &DeleteAlbumsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeleteAlbums", "cloudphoto", "openAPI")
	return
}

// CreateDeleteAlbumsResponse creates a response to parse from DeleteAlbums response
func CreateDeleteAlbumsResponse() (response *DeleteAlbumsResponse) {
	response = &DeleteAlbumsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
