package cloudapi

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

// ImportSwagger invokes the cloudapi.ImportSwagger API synchronously
func (client *Client) ImportSwagger(request *ImportSwaggerRequest) (response *ImportSwaggerResponse, err error) {
	response = CreateImportSwaggerResponse()
	err = client.DoAction(request, response)
	return
}

// ImportSwaggerWithChan invokes the cloudapi.ImportSwagger API asynchronously
func (client *Client) ImportSwaggerWithChan(request *ImportSwaggerRequest) (<-chan *ImportSwaggerResponse, <-chan error) {
	responseChan := make(chan *ImportSwaggerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportSwagger(request)
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

// ImportSwaggerWithCallback invokes the cloudapi.ImportSwagger API asynchronously
func (client *Client) ImportSwaggerWithCallback(request *ImportSwaggerRequest, callback func(response *ImportSwaggerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportSwaggerResponse
		var err error
		defer close(result)
		response, err = client.ImportSwagger(request)
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

// ImportSwaggerRequest is the request struct for api ImportSwagger
type ImportSwaggerRequest struct {
	*requests.RpcRequest
	DataFormat      string           `position:"Query" name:"DataFormat"`
	DryRun          requests.Boolean `position:"Query" name:"DryRun"`
	Data            string           `position:"Body" name:"Data"`
	GroupId         string           `position:"Query" name:"GroupId"`
	GlobalCondition string           `position:"Query" name:"GlobalCondition"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	Overwrite       requests.Boolean `position:"Query" name:"Overwrite"`
}

// ImportSwaggerResponse is the response struct for api ImportSwagger
type ImportSwaggerResponse struct {
	*responses.BaseResponse
	RequestId    string                      `json:"RequestId" xml:"RequestId"`
	Success      SuccessInImportSwagger      `json:"Success" xml:"Success"`
	Failed       FailedInImportSwagger       `json:"Failed" xml:"Failed"`
	ModelFailed  ModelFailedInImportSwagger  `json:"ModelFailed" xml:"ModelFailed"`
	ModelSuccess ModelSuccessInImportSwagger `json:"ModelSuccess" xml:"ModelSuccess"`
}

// CreateImportSwaggerRequest creates a request to invoke ImportSwagger API
func CreateImportSwaggerRequest() (request *ImportSwaggerRequest) {
	request = &ImportSwaggerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ImportSwagger", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateImportSwaggerResponse creates a response to parse from ImportSwagger response
func CreateImportSwaggerResponse() (response *ImportSwaggerResponse) {
	response = &ImportSwaggerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}