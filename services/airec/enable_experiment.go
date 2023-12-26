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

// EnableExperiment invokes the airec.EnableExperiment API synchronously
func (client *Client) EnableExperiment(request *EnableExperimentRequest) (response *EnableExperimentResponse, err error) {
	response = CreateEnableExperimentResponse()
	err = client.DoAction(request, response)
	return
}

// EnableExperimentWithChan invokes the airec.EnableExperiment API asynchronously
func (client *Client) EnableExperimentWithChan(request *EnableExperimentRequest) (<-chan *EnableExperimentResponse, <-chan error) {
	responseChan := make(chan *EnableExperimentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableExperiment(request)
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

// EnableExperimentWithCallback invokes the airec.EnableExperiment API asynchronously
func (client *Client) EnableExperimentWithCallback(request *EnableExperimentRequest, callback func(response *EnableExperimentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableExperimentResponse
		var err error
		defer close(result)
		response, err = client.EnableExperiment(request)
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

// EnableExperimentRequest is the request struct for api EnableExperiment
type EnableExperimentRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
	SceneId    string `position:"Path" name:"sceneId"`
}

// EnableExperimentResponse is the response struct for api EnableExperiment
type EnableExperimentResponse struct {
	*responses.BaseResponse
	Result    bool   `json:"result" xml:"result"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateEnableExperimentRequest creates a request to invoke EnableExperiment API
func CreateEnableExperimentRequest() (request *EnableExperimentRequest) {
	request = &EnableExperimentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "EnableExperiment", "/v2/openapi/instances/[instanceId]/scenes/[sceneId]/actions/enable-experiment", "airec", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableExperimentResponse creates a response to parse from EnableExperiment response
func CreateEnableExperimentResponse() (response *EnableExperimentResponse) {
	response = &EnableExperimentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
