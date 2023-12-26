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

// StartPtsScene invokes the pts.StartPtsScene API synchronously
func (client *Client) StartPtsScene(request *StartPtsSceneRequest) (response *StartPtsSceneResponse, err error) {
	response = CreateStartPtsSceneResponse()
	err = client.DoAction(request, response)
	return
}

// StartPtsSceneWithChan invokes the pts.StartPtsScene API asynchronously
func (client *Client) StartPtsSceneWithChan(request *StartPtsSceneRequest) (<-chan *StartPtsSceneResponse, <-chan error) {
	responseChan := make(chan *StartPtsSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartPtsScene(request)
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

// StartPtsSceneWithCallback invokes the pts.StartPtsScene API asynchronously
func (client *Client) StartPtsSceneWithCallback(request *StartPtsSceneRequest, callback func(response *StartPtsSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartPtsSceneResponse
		var err error
		defer close(result)
		response, err = client.StartPtsScene(request)
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

// StartPtsSceneRequest is the request struct for api StartPtsScene
type StartPtsSceneRequest struct {
	*requests.RpcRequest
	SceneId string `position:"Query" name:"SceneId"`
}

// StartPtsSceneResponse is the response struct for api StartPtsScene
type StartPtsSceneResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	PlanId         string `json:"PlanId" xml:"PlanId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateStartPtsSceneRequest creates a request to invoke StartPtsScene API
func CreateStartPtsSceneRequest() (request *StartPtsSceneRequest) {
	request = &StartPtsSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "StartPtsScene", "", "")
	request.Method = requests.POST
	return
}

// CreateStartPtsSceneResponse creates a response to parse from StartPtsScene response
func CreateStartPtsSceneResponse() (response *StartPtsSceneResponse) {
	response = &StartPtsSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
