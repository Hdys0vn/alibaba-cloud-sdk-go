package cr_ee

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

// CreateRepoBuildRule invokes the cr.CreateRepoBuildRule API synchronously
// api document: https://help.aliyun.com/api/cr/createrepobuildrule.html
func (client *Client) CreateRepoBuildRule(request *CreateRepoBuildRuleRequest) (response *CreateRepoBuildRuleResponse, err error) {
	response = CreateCreateRepoBuildRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRepoBuildRuleWithChan invokes the cr.CreateRepoBuildRule API asynchronously
// api document: https://help.aliyun.com/api/cr/createrepobuildrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepoBuildRuleWithChan(request *CreateRepoBuildRuleRequest) (<-chan *CreateRepoBuildRuleResponse, <-chan error) {
	responseChan := make(chan *CreateRepoBuildRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRepoBuildRule(request)
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

// CreateRepoBuildRuleWithCallback invokes the cr.CreateRepoBuildRule API asynchronously
// api document: https://help.aliyun.com/api/cr/createrepobuildrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepoBuildRuleWithCallback(request *CreateRepoBuildRuleRequest, callback func(response *CreateRepoBuildRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRepoBuildRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateRepoBuildRule(request)
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

// CreateRepoBuildRuleRequest is the request struct for api CreateRepoBuildRule
type CreateRepoBuildRuleRequest struct {
	*requests.RpcRequest
	RepoId             string `position:"Query" name:"RepoId"`
	PushName           string `position:"Query" name:"PushName"`
	DockerfileName     string `position:"Query" name:"DockerfileName"`
	DockerfileLocation string `position:"Query" name:"DockerfileLocation"`
	InstanceId         string `position:"Query" name:"InstanceId"`
	ImageTag           string `position:"Query" name:"ImageTag"`
	PushType           string `position:"Query" name:"PushType"`
}

// CreateRepoBuildRuleResponse is the response struct for api CreateRepoBuildRule
type CreateRepoBuildRuleResponse struct {
	*responses.BaseResponse
	CreateRepoBuildRuleIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                         string `json:"Code" xml:"Code"`
	RequestId                    string `json:"RequestId" xml:"RequestId"`
	BuildRuleId                  string `json:"BuildRuleId" xml:"BuildRuleId"`
}

// CreateCreateRepoBuildRuleRequest creates a request to invoke CreateRepoBuildRule API
func CreateCreateRepoBuildRuleRequest() (request *CreateRepoBuildRuleRequest) {
	request = &CreateRepoBuildRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "CreateRepoBuildRule", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRepoBuildRuleResponse creates a response to parse from CreateRepoBuildRule response
func CreateCreateRepoBuildRuleResponse() (response *CreateRepoBuildRuleResponse) {
	response = &CreateRepoBuildRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
