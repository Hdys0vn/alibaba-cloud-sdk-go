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

// ListSceneParameters invokes the airec.ListSceneParameters API synchronously
func (client *Client) ListSceneParameters(request *ListSceneParametersRequest) (response *ListSceneParametersResponse, err error) {
	response = CreateListSceneParametersResponse()
	err = client.DoAction(request, response)
	return
}

// ListSceneParametersWithChan invokes the airec.ListSceneParameters API asynchronously
func (client *Client) ListSceneParametersWithChan(request *ListSceneParametersRequest) (<-chan *ListSceneParametersResponse, <-chan error) {
	responseChan := make(chan *ListSceneParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSceneParameters(request)
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

// ListSceneParametersWithCallback invokes the airec.ListSceneParameters API asynchronously
func (client *Client) ListSceneParametersWithCallback(request *ListSceneParametersRequest, callback func(response *ListSceneParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSceneParametersResponse
		var err error
		defer close(result)
		response, err = client.ListSceneParameters(request)
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

// ListSceneParametersRequest is the request struct for api ListSceneParameters
type ListSceneParametersRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
}

// ListSceneParametersResponse is the response struct for api ListSceneParameters
type ListSceneParametersResponse struct {
	*responses.BaseResponse
	Code      string                      `json:"code" xml:"code"`
	Message   string                      `json:"message" xml:"message"`
	RequestId string                      `json:"requestId" xml:"requestId"`
	Result    ResultInListSceneParameters `json:"result" xml:"result"`
}

// CreateListSceneParametersRequest creates a request to invoke ListSceneParameters API
func CreateListSceneParametersRequest() (request *ListSceneParametersRequest) {
	request = &ListSceneParametersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "ListSceneParameters", "/v2/openapi/instances/[instanceId]/dashboard/scene-parameters", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListSceneParametersResponse creates a response to parse from ListSceneParameters response
func CreateListSceneParametersResponse() (response *ListSceneParametersResponse) {
	response = &ListSceneParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
