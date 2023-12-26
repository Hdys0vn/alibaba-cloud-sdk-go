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

// EditPhotos invokes the cloudphoto.EditPhotos API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/editphotos.html
func (client *Client) EditPhotos(request *EditPhotosRequest) (response *EditPhotosResponse, err error) {
	response = CreateEditPhotosResponse()
	err = client.DoAction(request, response)
	return
}

// EditPhotosWithChan invokes the cloudphoto.EditPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/editphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EditPhotosWithChan(request *EditPhotosRequest) (<-chan *EditPhotosResponse, <-chan error) {
	responseChan := make(chan *EditPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditPhotos(request)
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

// EditPhotosWithCallback invokes the cloudphoto.EditPhotos API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/editphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EditPhotosWithCallback(request *EditPhotosRequest, callback func(response *EditPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditPhotosResponse
		var err error
		defer close(result)
		response, err = client.EditPhotos(request)
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

// EditPhotosRequest is the request struct for api EditPhotos
type EditPhotosRequest struct {
	*requests.RpcRequest
	TakenAt         requests.Integer `position:"Query" name:"TakenAt"`
	LibraryId       string           `position:"Query" name:"LibraryId"`
	ShareExpireTime requests.Integer `position:"Query" name:"ShareExpireTime"`
	PhotoId         *[]string        `position:"Query" name:"PhotoId"  type:"Repeated"`
	StoreName       string           `position:"Query" name:"StoreName"`
	Remark          string           `position:"Query" name:"Remark"`
	Title           string           `position:"Query" name:"Title"`
}

// EditPhotosResponse is the response struct for api EditPhotos
type EditPhotosResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Action    string   `json:"Action" xml:"Action"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateEditPhotosRequest creates a request to invoke EditPhotos API
func CreateEditPhotosRequest() (request *EditPhotosRequest) {
	request = &EditPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "EditPhotos", "cloudphoto", "openAPI")
	return
}

// CreateEditPhotosResponse creates a response to parse from EditPhotos response
func CreateEditPhotosResponse() (response *EditPhotosResponse) {
	response = &EditPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
