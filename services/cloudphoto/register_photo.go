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

// RegisterPhoto invokes the cloudphoto.RegisterPhoto API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/registerphoto.html
func (client *Client) RegisterPhoto(request *RegisterPhotoRequest) (response *RegisterPhotoResponse, err error) {
	response = CreateRegisterPhotoResponse()
	err = client.DoAction(request, response)
	return
}

// RegisterPhotoWithChan invokes the cloudphoto.RegisterPhoto API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/registerphoto.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegisterPhotoWithChan(request *RegisterPhotoRequest) (<-chan *RegisterPhotoResponse, <-chan error) {
	responseChan := make(chan *RegisterPhotoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterPhoto(request)
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

// RegisterPhotoWithCallback invokes the cloudphoto.RegisterPhoto API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/registerphoto.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegisterPhotoWithCallback(request *RegisterPhotoRequest, callback func(response *RegisterPhotoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterPhotoResponse
		var err error
		defer close(result)
		response, err = client.RegisterPhoto(request)
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

// RegisterPhotoRequest is the request struct for api RegisterPhoto
type RegisterPhotoRequest struct {
	*requests.RpcRequest
	LibraryId  string           `position:"Query" name:"LibraryId"`
	Latitude   requests.Float   `position:"Query" name:"Latitude"`
	PhotoTitle string           `position:"Query" name:"PhotoTitle"`
	StoreName  string           `position:"Query" name:"StoreName"`
	IsVideo    string           `position:"Query" name:"IsVideo"`
	Remark     string           `position:"Query" name:"Remark"`
	Size       requests.Integer `position:"Query" name:"Size"`
	TakenAt    requests.Integer `position:"Query" name:"TakenAt"`
	Width      requests.Integer `position:"Query" name:"Width"`
	Location   string           `position:"Query" name:"Location"`
	Longitude  requests.Float   `position:"Query" name:"Longitude"`
	Height     requests.Integer `position:"Query" name:"Height"`
	Md5        string           `position:"Query" name:"Md5"`
}

// RegisterPhotoResponse is the response struct for api RegisterPhoto
type RegisterPhotoResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Photo     Photo  `json:"Photo" xml:"Photo"`
}

// CreateRegisterPhotoRequest creates a request to invoke RegisterPhoto API
func CreateRegisterPhotoRequest() (request *RegisterPhotoRequest) {
	request = &RegisterPhotoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "RegisterPhoto", "cloudphoto", "openAPI")
	return
}

// CreateRegisterPhotoResponse creates a response to parse from RegisterPhoto response
func CreateRegisterPhotoResponse() (response *RegisterPhotoResponse) {
	response = &RegisterPhotoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
