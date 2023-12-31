package opensearch

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

// GetScriptFileNames invokes the opensearch.GetScriptFileNames API synchronously
func (client *Client) GetScriptFileNames(request *GetScriptFileNamesRequest) (response *GetScriptFileNamesResponse, err error) {
	response = CreateGetScriptFileNamesResponse()
	err = client.DoAction(request, response)
	return
}

// GetScriptFileNamesWithChan invokes the opensearch.GetScriptFileNames API asynchronously
func (client *Client) GetScriptFileNamesWithChan(request *GetScriptFileNamesRequest) (<-chan *GetScriptFileNamesResponse, <-chan error) {
	responseChan := make(chan *GetScriptFileNamesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetScriptFileNames(request)
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

// GetScriptFileNamesWithCallback invokes the opensearch.GetScriptFileNames API asynchronously
func (client *Client) GetScriptFileNamesWithCallback(request *GetScriptFileNamesRequest, callback func(response *GetScriptFileNamesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetScriptFileNamesResponse
		var err error
		defer close(result)
		response, err = client.GetScriptFileNames(request)
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

// GetScriptFileNamesRequest is the request struct for api GetScriptFileNames
type GetScriptFileNamesRequest struct {
	*requests.RoaRequest
	AppVersionId     string `position:"Path" name:"appVersionId"`
	ScriptName       string `position:"Path" name:"scriptName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// GetScriptFileNamesResponse is the response struct for api GetScriptFileNames
type GetScriptFileNamesResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"requestId" xml:"requestId"`
	Result    []FileInfos `json:"result" xml:"result"`
}

// CreateGetScriptFileNamesRequest creates a request to invoke GetScriptFileNames API
func CreateGetScriptFileNamesRequest() (request *GetScriptFileNamesRequest) {
	request = &GetScriptFileNamesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "GetScriptFileNames", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appVersionId]/sort-scripts/[scriptName]/file-names", "", "")
	request.Method = requests.GET
	return
}

// CreateGetScriptFileNamesResponse creates a response to parse from GetScriptFileNames response
func CreateGetScriptFileNamesResponse() (response *GetScriptFileNamesResponse) {
	response = &GetScriptFileNamesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
