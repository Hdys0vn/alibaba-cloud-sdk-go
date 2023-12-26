package pts

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

// StopDebuggingJMeterScene invokes the pts.StopDebuggingJMeterScene API synchronously
func (client *Client) StopDebuggingJMeterScene(request *StopDebuggingJMeterSceneRequest) (response *StopDebuggingJMeterSceneResponse, err error) {
	response = CreateStopDebuggingJMeterSceneResponse()
	err = client.DoAction(request, response)
	return
}

// StopDebuggingJMeterSceneWithChan invokes the pts.StopDebuggingJMeterScene API asynchronously
func (client *Client) StopDebuggingJMeterSceneWithChan(request *StopDebuggingJMeterSceneRequest) (<-chan *StopDebuggingJMeterSceneResponse, <-chan error) {
	responseChan := make(chan *StopDebuggingJMeterSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDebuggingJMeterScene(request)
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

// StopDebuggingJMeterSceneWithCallback invokes the pts.StopDebuggingJMeterScene API asynchronously
func (client *Client) StopDebuggingJMeterSceneWithCallback(request *StopDebuggingJMeterSceneRequest, callback func(response *StopDebuggingJMeterSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDebuggingJMeterSceneResponse
		var err error
		defer close(result)
		response, err = client.StopDebuggingJMeterScene(request)
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

// StopDebuggingJMeterSceneRequest is the request struct for api StopDebuggingJMeterScene
type StopDebuggingJMeterSceneRequest struct {
	*requests.RpcRequest
	SceneId string `position:"Query" name:"SceneId"`
}

// StopDebuggingJMeterSceneResponse is the response struct for api StopDebuggingJMeterScene
type StopDebuggingJMeterSceneResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateStopDebuggingJMeterSceneRequest creates a request to invoke StopDebuggingJMeterScene API
func CreateStopDebuggingJMeterSceneRequest() (request *StopDebuggingJMeterSceneRequest) {
	request = &StopDebuggingJMeterSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "StopDebuggingJMeterScene", "", "")
	request.Method = requests.POST
	return
}

// CreateStopDebuggingJMeterSceneResponse creates a response to parse from StopDebuggingJMeterScene response
func CreateStopDebuggingJMeterSceneResponse() (response *StopDebuggingJMeterSceneResponse) {
	response = &StopDebuggingJMeterSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}