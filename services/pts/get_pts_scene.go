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

// GetPtsScene invokes the pts.GetPtsScene API synchronously
func (client *Client) GetPtsScene(request *GetPtsSceneRequest) (response *GetPtsSceneResponse, err error) {
	response = CreateGetPtsSceneResponse()
	err = client.DoAction(request, response)
	return
}

// GetPtsSceneWithChan invokes the pts.GetPtsScene API asynchronously
func (client *Client) GetPtsSceneWithChan(request *GetPtsSceneRequest) (<-chan *GetPtsSceneResponse, <-chan error) {
	responseChan := make(chan *GetPtsSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPtsScene(request)
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

// GetPtsSceneWithCallback invokes the pts.GetPtsScene API asynchronously
func (client *Client) GetPtsSceneWithCallback(request *GetPtsSceneRequest, callback func(response *GetPtsSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPtsSceneResponse
		var err error
		defer close(result)
		response, err = client.GetPtsScene(request)
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

// GetPtsSceneRequest is the request struct for api GetPtsScene
type GetPtsSceneRequest struct {
	*requests.RpcRequest
	SceneId string `position:"Query" name:"SceneId"`
}

// GetPtsSceneResponse is the response struct for api GetPtsScene
type GetPtsSceneResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Scene          Scene  `json:"Scene" xml:"Scene"`
}

// CreateGetPtsSceneRequest creates a request to invoke GetPtsScene API
func CreateGetPtsSceneRequest() (request *GetPtsSceneRequest) {
	request = &GetPtsSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "GetPtsScene", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPtsSceneResponse creates a response to parse from GetPtsScene response
func CreateGetPtsSceneResponse() (response *GetPtsSceneResponse) {
	response = &GetPtsSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
