package amp

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

// CoreEngineCopy invokes the amp.CoreEngineCopy API synchronously
func (client *Client) CoreEngineCopy(request *CoreEngineCopyRequest) (response *CoreEngineCopyResponse, err error) {
	response = CreateCoreEngineCopyResponse()
	err = client.DoAction(request, response)
	return
}

// CoreEngineCopyWithChan invokes the amp.CoreEngineCopy API asynchronously
func (client *Client) CoreEngineCopyWithChan(request *CoreEngineCopyRequest) (<-chan *CoreEngineCopyResponse, <-chan error) {
	responseChan := make(chan *CoreEngineCopyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CoreEngineCopy(request)
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

// CoreEngineCopyWithCallback invokes the amp.CoreEngineCopy API asynchronously
func (client *Client) CoreEngineCopyWithCallback(request *CoreEngineCopyRequest, callback func(response *CoreEngineCopyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CoreEngineCopyResponse
		var err error
		defer close(result)
		response, err = client.CoreEngineCopy(request)
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

// CoreEngineCopyRequest is the request struct for api CoreEngineCopy
type CoreEngineCopyRequest struct {
	*requests.RoaRequest
	Param5    string                  `position:"Query" name:"param5"`
	ParamList *[]string               `position:"Query" name:"ParamList"  type:"Json"`
	Param4    string                  `position:"Query" name:"param4"`
	Param1    *[]CoreEngineCopyParam1 `position:"Body" name:"Param1"  type:"Json"`
}

// CoreEngineCopyParam1 is a repeated param struct in CoreEngineCopyRequest
type CoreEngineCopyParam1 struct {
	Param3 string `name:"Param3"`
	Param1 string `name:"Param1"`
	Param2 string `name:"Param2"`
}

// CoreEngineCopyResponse is the response struct for api CoreEngineCopy
type CoreEngineCopyResponse struct {
	*responses.BaseResponse
}

// CreateCoreEngineCopyRequest creates a request to invoke CoreEngineCopy API
func CreateCoreEngineCopyRequest() (request *CoreEngineCopyRequest) {
	request = &CoreEngineCopyRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("amp", "2020-07-08", "CoreEngineCopy", "/api/new/demo", "ServiceCode", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCoreEngineCopyResponse creates a response to parse from CoreEngineCopy response
func CreateCoreEngineCopyResponse() (response *CoreEngineCopyResponse) {
	response = &CoreEngineCopyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}