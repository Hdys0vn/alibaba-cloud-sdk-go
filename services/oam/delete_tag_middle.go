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

// DeleteTagMiddle invokes the oam.DeleteTagMiddle API synchronously
// api document: https://help.aliyun.com/api/oam/deletetagmiddle.html
func (client *Client) DeleteTagMiddle(request *DeleteTagMiddleRequest) (response *DeleteTagMiddleResponse, err error) {
	response = CreateDeleteTagMiddleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTagMiddleWithChan invokes the oam.DeleteTagMiddle API asynchronously
// api document: https://help.aliyun.com/api/oam/deletetagmiddle.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTagMiddleWithChan(request *DeleteTagMiddleRequest) (<-chan *DeleteTagMiddleResponse, <-chan error) {
	responseChan := make(chan *DeleteTagMiddleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTagMiddle(request)
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

// DeleteTagMiddleWithCallback invokes the oam.DeleteTagMiddle API asynchronously
// api document: https://help.aliyun.com/api/oam/deletetagmiddle.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTagMiddleWithCallback(request *DeleteTagMiddleRequest, callback func(response *DeleteTagMiddleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTagMiddleResponse
		var err error
		defer close(result)
		response, err = client.DeleteTagMiddle(request)
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

// DeleteTagMiddleRequest is the request struct for api DeleteTagMiddle
type DeleteTagMiddleRequest struct {
	*requests.RpcRequest
	RoleCellId *[]string `position:"Query" name:"RoleCellId"  type:"Repeated"`
	TagId      string    `position:"Query" name:"TagId"`
}

// DeleteTagMiddleResponse is the response struct for api DeleteTagMiddle
type DeleteTagMiddleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTagMiddleRequest creates a request to invoke DeleteTagMiddle API
func CreateDeleteTagMiddleRequest() (request *DeleteTagMiddleRequest) {
	request = &DeleteTagMiddleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "DeleteTagMiddle", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTagMiddleResponse creates a response to parse from DeleteTagMiddle response
func CreateDeleteTagMiddleResponse() (response *DeleteTagMiddleResponse) {
	response = &DeleteTagMiddleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
