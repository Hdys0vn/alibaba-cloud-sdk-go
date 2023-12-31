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

// SetFaceCover invokes the cloudphoto.SetFaceCover API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/setfacecover.html
func (client *Client) SetFaceCover(request *SetFaceCoverRequest) (response *SetFaceCoverResponse, err error) {
	response = CreateSetFaceCoverResponse()
	err = client.DoAction(request, response)
	return
}

// SetFaceCoverWithChan invokes the cloudphoto.SetFaceCover API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/setfacecover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetFaceCoverWithChan(request *SetFaceCoverRequest) (<-chan *SetFaceCoverResponse, <-chan error) {
	responseChan := make(chan *SetFaceCoverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetFaceCover(request)
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

// SetFaceCoverWithCallback invokes the cloudphoto.SetFaceCover API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/setfacecover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetFaceCoverWithCallback(request *SetFaceCoverRequest, callback func(response *SetFaceCoverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetFaceCoverResponse
		var err error
		defer close(result)
		response, err = client.SetFaceCover(request)
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

// SetFaceCoverRequest is the request struct for api SetFaceCover
type SetFaceCoverRequest struct {
	*requests.RpcRequest
	LibraryId string           `position:"Query" name:"LibraryId"`
	PhotoId   requests.Integer `position:"Query" name:"PhotoId"`
	StoreName string           `position:"Query" name:"StoreName"`
	FaceId    requests.Integer `position:"Query" name:"FaceId"`
}

// SetFaceCoverResponse is the response struct for api SetFaceCover
type SetFaceCoverResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
}

// CreateSetFaceCoverRequest creates a request to invoke SetFaceCover API
func CreateSetFaceCoverRequest() (request *SetFaceCoverRequest) {
	request = &SetFaceCoverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetFaceCover", "cloudphoto", "openAPI")
	return
}

// CreateSetFaceCoverResponse creates a response to parse from SetFaceCover response
func CreateSetFaceCoverResponse() (response *SetFaceCoverResponse) {
	response = &SetFaceCoverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
