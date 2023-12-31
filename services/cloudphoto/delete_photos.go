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

// DeletePhotos invokes the cloudphoto.DeletePhotos API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletephotos.html
func (client *Client) DeletePhotos(request *DeletePhotosRequest) (response *DeletePhotosResponse, err error) {
	response = CreateDeletePhotosResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePhotosWithChan invokes the cloudphoto.DeletePhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletephotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePhotosWithChan(request *DeletePhotosRequest) (<-chan *DeletePhotosResponse, <-chan error) {
	responseChan := make(chan *DeletePhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePhotos(request)
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

// DeletePhotosWithCallback invokes the cloudphoto.DeletePhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletephotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePhotosWithCallback(request *DeletePhotosRequest, callback func(response *DeletePhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePhotosResponse
		var err error
		defer close(result)
		response, err = client.DeletePhotos(request)
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

// DeletePhotosRequest is the request struct for api DeletePhotos
type DeletePhotosRequest struct {
	*requests.RpcRequest
	LibraryId string    `position:"Query" name:"LibraryId"`
	StoreName string    `position:"Query" name:"StoreName"`
	PhotoId   *[]string `position:"Query" name:"PhotoId"  type:"Repeated"`
}

// DeletePhotosResponse is the response struct for api DeletePhotos
type DeletePhotosResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Action    string   `json:"Action" xml:"Action"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateDeletePhotosRequest creates a request to invoke DeletePhotos API
func CreateDeletePhotosRequest() (request *DeletePhotosRequest) {
	request = &DeletePhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeletePhotos", "cloudphoto", "openAPI")
	return
}

// CreateDeletePhotosResponse creates a response to parse from DeletePhotos response
func CreateDeletePhotosResponse() (response *DeletePhotosResponse) {
	response = &DeletePhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
