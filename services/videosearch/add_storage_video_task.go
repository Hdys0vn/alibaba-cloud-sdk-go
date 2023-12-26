package videosearch

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

// AddStorageVideoTask invokes the videosearch.AddStorageVideoTask API synchronously
// api document: https://help.aliyun.com/api/videosearch/addstoragevideotask.html
func (client *Client) AddStorageVideoTask(request *AddStorageVideoTaskRequest) (response *AddStorageVideoTaskResponse, err error) {
	response = CreateAddStorageVideoTaskResponse()
	err = client.DoAction(request, response)
	return
}

// AddStorageVideoTaskWithChan invokes the videosearch.AddStorageVideoTask API asynchronously
// api document: https://help.aliyun.com/api/videosearch/addstoragevideotask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddStorageVideoTaskWithChan(request *AddStorageVideoTaskRequest) (<-chan *AddStorageVideoTaskResponse, <-chan error) {
	responseChan := make(chan *AddStorageVideoTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddStorageVideoTask(request)
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

// AddStorageVideoTaskWithCallback invokes the videosearch.AddStorageVideoTask API asynchronously
// api document: https://help.aliyun.com/api/videosearch/addstoragevideotask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddStorageVideoTaskWithCallback(request *AddStorageVideoTaskRequest, callback func(response *AddStorageVideoTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddStorageVideoTaskResponse
		var err error
		defer close(result)
		response, err = client.AddStorageVideoTask(request)
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

// AddStorageVideoTaskRequest is the request struct for api AddStorageVideoTask
type AddStorageVideoTaskRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	VideoTags   string `position:"Body" name:"VideoTags"`
	VideoId     string `position:"Body" name:"VideoId"`
	InstanceId  string `position:"Body" name:"InstanceId"`
	VideoUrl    string `position:"Body" name:"VideoUrl"`
	CallbackUrl string `position:"Body" name:"CallbackUrl"`
}

// AddStorageVideoTaskResponse is the response struct for api AddStorageVideoTask
type AddStorageVideoTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddStorageVideoTaskRequest creates a request to invoke AddStorageVideoTask API
func CreateAddStorageVideoTaskRequest() (request *AddStorageVideoTaskRequest) {
	request = &AddStorageVideoTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videosearch", "2020-02-25", "AddStorageVideoTask", "", "")
	request.Method = requests.POST
	return
}

// CreateAddStorageVideoTaskResponse creates a response to parse from AddStorageVideoTask response
func CreateAddStorageVideoTaskResponse() (response *AddStorageVideoTaskResponse) {
	response = &AddStorageVideoTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
