package airec

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

// DescribeScene invokes the airec.DescribeScene API synchronously
func (client *Client) DescribeScene(request *DescribeSceneRequest) (response *DescribeSceneResponse, err error) {
	response = CreateDescribeSceneResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSceneWithChan invokes the airec.DescribeScene API asynchronously
func (client *Client) DescribeSceneWithChan(request *DescribeSceneRequest) (<-chan *DescribeSceneResponse, <-chan error) {
	responseChan := make(chan *DescribeSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScene(request)
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

// DescribeSceneWithCallback invokes the airec.DescribeScene API asynchronously
func (client *Client) DescribeSceneWithCallback(request *DescribeSceneRequest, callback func(response *DescribeSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSceneResponse
		var err error
		defer close(result)
		response, err = client.DescribeScene(request)
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

// DescribeSceneRequest is the request struct for api DescribeScene
type DescribeSceneRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
	SceneId    string `position:"Path" name:"sceneId"`
}

// DescribeSceneResponse is the response struct for api DescribeScene
type DescribeSceneResponse struct {
	*responses.BaseResponse
	RequestId string                `json:"requestId" xml:"requestId"`
	Code      string                `json:"code" xml:"code"`
	Message   string                `json:"message" xml:"message"`
	Result    ResultInDescribeScene `json:"result" xml:"result"`
}

// CreateDescribeSceneRequest creates a request to invoke DescribeScene API
func CreateDescribeSceneRequest() (request *DescribeSceneRequest) {
	request = &DescribeSceneRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "DescribeScene", "/v2/openapi/instances/[instanceId]/scenes/[sceneId]", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeSceneResponse creates a response to parse from DescribeScene response
func CreateDescribeSceneResponse() (response *DescribeSceneResponse) {
	response = &DescribeSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
