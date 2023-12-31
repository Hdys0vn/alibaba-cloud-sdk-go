package industry_brain

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

// AsyncResponsePost invokes the industry_brain.AsyncResponsePost API synchronously
func (client *Client) AsyncResponsePost(request *AsyncResponsePostRequest) (response *AsyncResponsePostResponse, err error) {
	response = CreateAsyncResponsePostResponse()
	err = client.DoAction(request, response)
	return
}

// AsyncResponsePostWithChan invokes the industry_brain.AsyncResponsePost API asynchronously
func (client *Client) AsyncResponsePostWithChan(request *AsyncResponsePostRequest) (<-chan *AsyncResponsePostResponse, <-chan error) {
	responseChan := make(chan *AsyncResponsePostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AsyncResponsePost(request)
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

// AsyncResponsePostWithCallback invokes the industry_brain.AsyncResponsePost API asynchronously
func (client *Client) AsyncResponsePostWithCallback(request *AsyncResponsePostRequest, callback func(response *AsyncResponsePostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AsyncResponsePostResponse
		var err error
		defer close(result)
		response, err = client.AsyncResponsePost(request)
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

// AsyncResponsePostRequest is the request struct for api AsyncResponsePost
type AsyncResponsePostRequest struct {
	*requests.RpcRequest
	Data     string         `position:"Query" name:"Data"`
	Context  string         `position:"Query" name:"Context"`
	Progress requests.Float `position:"Query" name:"Progress"`
	Message  string         `position:"Query" name:"Message"`
	TaskId   string         `position:"Query" name:"TaskId"`
	Status   string         `position:"Query" name:"Status"`
}

// AsyncResponsePostResponse is the response struct for api AsyncResponsePost
type AsyncResponsePostResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Status    string `json:"Status" xml:"Status"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateAsyncResponsePostRequest creates a request to invoke AsyncResponsePost API
func CreateAsyncResponsePostRequest() (request *AsyncResponsePostRequest) {
	request = &AsyncResponsePostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("industry-brain", "2019-06-29", "AsyncResponsePost", "", "")
	request.Method = requests.POST
	return
}

// CreateAsyncResponsePostResponse creates a response to parse from AsyncResponsePost response
func CreateAsyncResponsePostResponse() (response *AsyncResponsePostResponse) {
	response = &AsyncResponsePostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
