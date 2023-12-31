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

// DescribeExperimentEnvProgress invokes the airec.DescribeExperimentEnvProgress API synchronously
func (client *Client) DescribeExperimentEnvProgress(request *DescribeExperimentEnvProgressRequest) (response *DescribeExperimentEnvProgressResponse, err error) {
	response = CreateDescribeExperimentEnvProgressResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExperimentEnvProgressWithChan invokes the airec.DescribeExperimentEnvProgress API asynchronously
func (client *Client) DescribeExperimentEnvProgressWithChan(request *DescribeExperimentEnvProgressRequest) (<-chan *DescribeExperimentEnvProgressResponse, <-chan error) {
	responseChan := make(chan *DescribeExperimentEnvProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExperimentEnvProgress(request)
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

// DescribeExperimentEnvProgressWithCallback invokes the airec.DescribeExperimentEnvProgress API asynchronously
func (client *Client) DescribeExperimentEnvProgressWithCallback(request *DescribeExperimentEnvProgressRequest, callback func(response *DescribeExperimentEnvProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExperimentEnvProgressResponse
		var err error
		defer close(result)
		response, err = client.DescribeExperimentEnvProgress(request)
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

// DescribeExperimentEnvProgressRequest is the request struct for api DescribeExperimentEnvProgress
type DescribeExperimentEnvProgressRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
	SceneId    string `position:"Path" name:"sceneId"`
}

// DescribeExperimentEnvProgressResponse is the response struct for api DescribeExperimentEnvProgress
type DescribeExperimentEnvProgressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateDescribeExperimentEnvProgressRequest creates a request to invoke DescribeExperimentEnvProgress API
func CreateDescribeExperimentEnvProgressRequest() (request *DescribeExperimentEnvProgressRequest) {
	request = &DescribeExperimentEnvProgressRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "DescribeExperimentEnvProgress", "/v2/openapi/instances/[instanceId]/scenes/[sceneId]/experiment-progress", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeExperimentEnvProgressResponse creates a response to parse from DescribeExperimentEnvProgress response
func CreateDescribeExperimentEnvProgressResponse() (response *DescribeExperimentEnvProgressResponse) {
	response = &DescribeExperimentEnvProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
