package retailcloud

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

// ResumeDeploy invokes the retailcloud.ResumeDeploy API synchronously
func (client *Client) ResumeDeploy(request *ResumeDeployRequest) (response *ResumeDeployResponse, err error) {
	response = CreateResumeDeployResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeDeployWithChan invokes the retailcloud.ResumeDeploy API asynchronously
func (client *Client) ResumeDeployWithChan(request *ResumeDeployRequest) (<-chan *ResumeDeployResponse, <-chan error) {
	responseChan := make(chan *ResumeDeployResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeDeploy(request)
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

// ResumeDeployWithCallback invokes the retailcloud.ResumeDeploy API asynchronously
func (client *Client) ResumeDeployWithCallback(request *ResumeDeployRequest, callback func(response *ResumeDeployResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeDeployResponse
		var err error
		defer close(result)
		response, err = client.ResumeDeploy(request)
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

// ResumeDeployRequest is the request struct for api ResumeDeploy
type ResumeDeployRequest struct {
	*requests.RpcRequest
	DeployOrderId requests.Integer `position:"Query" name:"DeployOrderId"`
}

// ResumeDeployResponse is the response struct for api ResumeDeploy
type ResumeDeployResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateResumeDeployRequest creates a request to invoke ResumeDeploy API
func CreateResumeDeployRequest() (request *ResumeDeployRequest) {
	request = &ResumeDeployRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ResumeDeploy", "", "")
	request.Method = requests.POST
	return
}

// CreateResumeDeployResponse creates a response to parse from ResumeDeploy response
func CreateResumeDeployResponse() (response *ResumeDeployResponse) {
	response = &ResumeDeployResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
