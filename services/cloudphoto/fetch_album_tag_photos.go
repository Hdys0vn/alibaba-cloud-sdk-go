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

// FetchAlbumTagPhotos invokes the cloudphoto.FetchAlbumTagPhotos API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/fetchalbumtagphotos.html
func (client *Client) FetchAlbumTagPhotos(request *FetchAlbumTagPhotosRequest) (response *FetchAlbumTagPhotosResponse, err error) {
	response = CreateFetchAlbumTagPhotosResponse()
	err = client.DoAction(request, response)
	return
}

// FetchAlbumTagPhotosWithChan invokes the cloudphoto.FetchAlbumTagPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/fetchalbumtagphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FetchAlbumTagPhotosWithChan(request *FetchAlbumTagPhotosRequest) (<-chan *FetchAlbumTagPhotosResponse, <-chan error) {
	responseChan := make(chan *FetchAlbumTagPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FetchAlbumTagPhotos(request)
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

// FetchAlbumTagPhotosWithCallback invokes the cloudphoto.FetchAlbumTagPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/fetchalbumtagphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FetchAlbumTagPhotosWithCallback(request *FetchAlbumTagPhotosRequest, callback func(response *FetchAlbumTagPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FetchAlbumTagPhotosResponse
		var err error
		defer close(result)
		response, err = client.FetchAlbumTagPhotos(request)
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

// FetchAlbumTagPhotosRequest is the request struct for api FetchAlbumTagPhotos
type FetchAlbumTagPhotosRequest struct {
	*requests.RpcRequest
	Size      requests.Integer `position:"Query" name:"Size"`
	TagId     requests.Integer `position:"Query" name:"TagId"`
	LibraryId string           `position:"Query" name:"LibraryId"`
	AlbumId   requests.Integer `position:"Query" name:"AlbumId"`
	StoreName string           `position:"Query" name:"StoreName"`
	Page      requests.Integer `position:"Query" name:"Page"`
}

// FetchAlbumTagPhotosResponse is the response struct for api FetchAlbumTagPhotos
type FetchAlbumTagPhotosResponse struct {
	*responses.BaseResponse
	Code       string   `json:"Code" xml:"Code"`
	Message    string   `json:"Message" xml:"Message"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	Action     string   `json:"Action" xml:"Action"`
	Results    []Result `json:"Results" xml:"Results"`
}

// CreateFetchAlbumTagPhotosRequest creates a request to invoke FetchAlbumTagPhotos API
func CreateFetchAlbumTagPhotosRequest() (request *FetchAlbumTagPhotosRequest) {
	request = &FetchAlbumTagPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "FetchAlbumTagPhotos", "cloudphoto", "openAPI")
	return
}

// CreateFetchAlbumTagPhotosResponse creates a response to parse from FetchAlbumTagPhotos response
func CreateFetchAlbumTagPhotosResponse() (response *FetchAlbumTagPhotosResponse) {
	response = &FetchAlbumTagPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
