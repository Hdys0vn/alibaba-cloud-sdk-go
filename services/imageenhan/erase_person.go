package imageenhan

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

// ErasePerson invokes the imageenhan.ErasePerson API synchronously
func (client *Client) ErasePerson(request *ErasePersonRequest) (response *ErasePersonResponse, err error) {
	response = CreateErasePersonResponse()
	err = client.DoAction(request, response)
	return
}

// ErasePersonWithChan invokes the imageenhan.ErasePerson API asynchronously
func (client *Client) ErasePersonWithChan(request *ErasePersonRequest) (<-chan *ErasePersonResponse, <-chan error) {
	responseChan := make(chan *ErasePersonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ErasePerson(request)
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

// ErasePersonWithCallback invokes the imageenhan.ErasePerson API asynchronously
func (client *Client) ErasePersonWithCallback(request *ErasePersonRequest, callback func(response *ErasePersonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ErasePersonResponse
		var err error
		defer close(result)
		response, err = client.ErasePerson(request)
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

// ErasePersonRequest is the request struct for api ErasePerson
type ErasePersonRequest struct {
	*requests.RpcRequest
	UserMask string `position:"Body" name:"UserMask"`
	ImageURL string `position:"Body" name:"ImageURL"`
}

// ErasePersonResponse is the response struct for api ErasePerson
type ErasePersonResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateErasePersonRequest creates a request to invoke ErasePerson API
func CreateErasePersonRequest() (request *ErasePersonRequest) {
	request = &ErasePersonRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageenhan", "2019-09-30", "ErasePerson", "imageenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateErasePersonResponse creates a response to parse from ErasePerson response
func CreateErasePersonResponse() (response *ErasePersonResponse) {
	response = &ErasePersonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
