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

// StartTestingJMeterScene invokes the pts.StartTestingJMeterScene API synchronously
func (client *Client) StartTestingJMeterScene(request *StartTestingJMeterSceneRequest) (response *StartTestingJMeterSceneResponse, err error) {
	response = CreateStartTestingJMeterSceneResponse()
	err = client.DoAction(request, response)
	return
}

// StartTestingJMeterSceneWithChan invokes the pts.StartTestingJMeterScene API asynchronously
func (client *Client) StartTestingJMeterSceneWithChan(request *StartTestingJMeterSceneRequest) (<-chan *StartTestingJMeterSceneResponse, <-chan error) {
	responseChan := make(chan *StartTestingJMeterSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartTestingJMeterScene(request)
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

// StartTestingJMeterSceneWithCallback invokes the pts.StartTestingJMeterScene API asynchronously
func (client *Client) StartTestingJMeterSceneWithCallback(request *StartTestingJMeterSceneRequest, callback func(response *StartTestingJMeterSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartTestingJMeterSceneResponse
		var err error
		defer close(result)
		response, err = client.StartTestingJMeterScene(request)
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

// StartTestingJMeterSceneRequest is the request struct for api StartTestingJMeterScene
type StartTestingJMeterSceneRequest struct {
	*requests.RpcRequest
	SceneId string `position:"Query" name:"SceneId"`
}

// StartTestingJMeterSceneResponse is the response struct for api StartTestingJMeterScene
type StartTestingJMeterSceneResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	ReportId       string `json:"ReportId" xml:"ReportId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateStartTestingJMeterSceneRequest creates a request to invoke StartTestingJMeterScene API
func CreateStartTestingJMeterSceneRequest() (request *StartTestingJMeterSceneRequest) {
	request = &StartTestingJMeterSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "StartTestingJMeterScene", "", "")
	request.Method = requests.POST
	return
}

// CreateStartTestingJMeterSceneResponse creates a response to parse from StartTestingJMeterScene response
func CreateStartTestingJMeterSceneResponse() (response *StartTestingJMeterSceneResponse) {
	response = &StartTestingJMeterSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
