package oam

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

// AddTag invokes the oam.AddTag API synchronously
// api document: https://help.aliyun.com/api/oam/addtag.html
func (client *Client) AddTag(request *AddTagRequest) (response *AddTagResponse, err error) {
	response = CreateAddTagResponse()
	err = client.DoAction(request, response)
	return
}

// AddTagWithChan invokes the oam.AddTag API asynchronously
// api document: https://help.aliyun.com/api/oam/addtag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTagWithChan(request *AddTagRequest) (<-chan *AddTagResponse, <-chan error) {
	responseChan := make(chan *AddTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTag(request)
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

// AddTagWithCallback invokes the oam.AddTag API asynchronously
// api document: https://help.aliyun.com/api/oam/addtag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTagWithCallback(request *AddTagRequest, callback func(response *AddTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTagResponse
		var err error
		defer close(result)
		response, err = client.AddTag(request)
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

// AddTagRequest is the request struct for api AddTag
type AddTagRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
}

// AddTagResponse is the response struct for api AddTag
type AddTagResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddTagRequest creates a request to invoke AddTag API
func CreateAddTagRequest() (request *AddTagRequest) {
	request = &AddTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "AddTag", "", "")
	request.Method = requests.POST
	return
}

// CreateAddTagResponse creates a response to parse from AddTag response
func CreateAddTagResponse() (response *AddTagResponse) {
	response = &AddTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
