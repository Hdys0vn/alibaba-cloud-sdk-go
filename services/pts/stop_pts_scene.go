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

// StopPtsScene invokes the pts.StopPtsScene API synchronously
func (client *Client) StopPtsScene(request *StopPtsSceneRequest) (response *StopPtsSceneResponse, err error) {
	response = CreateStopPtsSceneResponse()
	err = client.DoAction(request, response)
	return
}

// StopPtsSceneWithChan invokes the pts.StopPtsScene API asynchronously
func (client *Client) StopPtsSceneWithChan(request *StopPtsSceneRequest) (<-chan *StopPtsSceneResponse, <-chan error) {
	responseChan := make(chan *StopPtsSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopPtsScene(request)
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

// StopPtsSceneWithCallback invokes the pts.StopPtsScene API asynchronously
func (client *Client) StopPtsSceneWithCallback(request *StopPtsSceneRequest, callback func(response *StopPtsSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopPtsSceneResponse
		var err error
		defer close(result)
		response, err = client.StopPtsScene(request)
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

// StopPtsSceneRequest is the request struct for api StopPtsScene
type StopPtsSceneRequest struct {
	*requests.RpcRequest
	SceneId string `position:"Query" name:"SceneId"`
}

// StopPtsSceneResponse is the response struct for api StopPtsScene
type StopPtsSceneResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateStopPtsSceneRequest creates a request to invoke StopPtsScene API
func CreateStopPtsSceneRequest() (request *StopPtsSceneRequest) {
	request = &StopPtsSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "StopPtsScene", "", "")
	request.Method = requests.POST
	return
}

// CreateStopPtsSceneResponse creates a response to parse from StopPtsScene response
func CreateStopPtsSceneResponse() (response *StopPtsSceneResponse) {
	response = &StopPtsSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
