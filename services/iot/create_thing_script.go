package iot

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

// CreateThingScript invokes the iot.CreateThingScript API synchronously
func (client *Client) CreateThingScript(request *CreateThingScriptRequest) (response *CreateThingScriptResponse, err error) {
	response = CreateCreateThingScriptResponse()
	err = client.DoAction(request, response)
	return
}

// CreateThingScriptWithChan invokes the iot.CreateThingScript API asynchronously
func (client *Client) CreateThingScriptWithChan(request *CreateThingScriptRequest) (<-chan *CreateThingScriptResponse, <-chan error) {
	responseChan := make(chan *CreateThingScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateThingScript(request)
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

// CreateThingScriptWithCallback invokes the iot.CreateThingScript API asynchronously
func (client *Client) CreateThingScriptWithCallback(request *CreateThingScriptRequest, callback func(response *CreateThingScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateThingScriptResponse
		var err error
		defer close(result)
		response, err = client.CreateThingScript(request)
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

// CreateThingScriptRequest is the request struct for api CreateThingScript
type CreateThingScriptRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ScriptType        string `position:"Query" name:"ScriptType"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
	ScriptContent     string `position:"Query" name:"ScriptContent"`
}

// CreateThingScriptResponse is the response struct for api CreateThingScript
type CreateThingScriptResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateCreateThingScriptRequest creates a request to invoke CreateThingScript API
func CreateCreateThingScriptRequest() (request *CreateThingScriptRequest) {
	request = &CreateThingScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateThingScript", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateThingScriptResponse creates a response to parse from CreateThingScript response
func CreateCreateThingScriptResponse() (response *CreateThingScriptResponse) {
	response = &CreateThingScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
