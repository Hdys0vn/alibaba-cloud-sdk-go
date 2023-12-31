package tesladam

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

// Action invokes the tesladam.Action API synchronously
// api document: https://help.aliyun.com/api/tesladam/action.html
func (client *Client) Action(request *ActionRequest) (response *ActionResponse, err error) {
	response = CreateActionResponse()
	err = client.DoAction(request, response)
	return
}

// ActionWithChan invokes the tesladam.Action API asynchronously
// api document: https://help.aliyun.com/api/tesladam/action.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ActionWithChan(request *ActionRequest) (<-chan *ActionResponse, <-chan error) {
	responseChan := make(chan *ActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Action(request)
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

// ActionWithCallback invokes the tesladam.Action API asynchronously
// api document: https://help.aliyun.com/api/tesladam/action.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ActionWithCallback(request *ActionRequest, callback func(response *ActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActionResponse
		var err error
		defer close(result)
		response, err = client.Action(request)
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

// ActionRequest is the request struct for api Action
type ActionRequest struct {
	*requests.RpcRequest
	OrderId  requests.Integer `position:"Query" name:"OrderId"`
	StepCode string           `position:"Query" name:"StepCode"`
}

// ActionResponse is the response struct for api Action
type ActionResponse struct {
	*responses.BaseResponse
	Status  bool   `json:"Status" xml:"Status"`
	Message string `json:"Message" xml:"Message"`
	Result  string `json:"Result" xml:"Result"`
}

// CreateActionRequest creates a request to invoke Action API
func CreateActionRequest() (request *ActionRequest) {
	request = &ActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaDam", "2018-01-18", "Action", "tesladam", "openAPI")
	return
}

// CreateActionResponse creates a response to parse from Action response
func CreateActionResponse() (response *ActionResponse) {
	response = &ActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
