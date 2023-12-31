package linkface

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

// UpdateFace invokes the linkface.UpdateFace API synchronously
// api document: https://help.aliyun.com/api/linkface/updateface.html
func (client *Client) UpdateFace(request *UpdateFaceRequest) (response *UpdateFaceResponse, err error) {
	response = CreateUpdateFaceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFaceWithChan invokes the linkface.UpdateFace API asynchronously
// api document: https://help.aliyun.com/api/linkface/updateface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateFaceWithChan(request *UpdateFaceRequest) (<-chan *UpdateFaceResponse, <-chan error) {
	responseChan := make(chan *UpdateFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFace(request)
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

// UpdateFaceWithCallback invokes the linkface.UpdateFace API asynchronously
// api document: https://help.aliyun.com/api/linkface/updateface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateFaceWithCallback(request *UpdateFaceRequest, callback func(response *UpdateFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFaceResponse
		var err error
		defer close(result)
		response, err = client.UpdateFace(request)
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

// UpdateFaceRequest is the request struct for api UpdateFace
type UpdateFaceRequest struct {
	*requests.RpcRequest
	Image    string `position:"Body" name:"Image"`
	UserId   string `position:"Body" name:"UserId"`
	UserInfo string `position:"Body" name:"UserInfo"`
}

// UpdateFaceResponse is the response struct for api UpdateFace
type UpdateFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateFaceRequest creates a request to invoke UpdateFace API
func CreateUpdateFaceRequest() (request *UpdateFaceRequest) {
	request = &UpdateFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkFace", "2018-07-20", "UpdateFace", "", "")
	return
}

// CreateUpdateFaceResponse creates a response to parse from UpdateFace response
func CreateUpdateFaceResponse() (response *UpdateFaceResponse) {
	response = &UpdateFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
