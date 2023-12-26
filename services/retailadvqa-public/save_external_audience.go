package retailadvqa_public

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

// SaveExternalAudience invokes the retailadvqa_public.SaveExternalAudience API synchronously
func (client *Client) SaveExternalAudience(request *SaveExternalAudienceRequest) (response *SaveExternalAudienceResponse, err error) {
	response = CreateSaveExternalAudienceResponse()
	err = client.DoAction(request, response)
	return
}

// SaveExternalAudienceWithChan invokes the retailadvqa_public.SaveExternalAudience API asynchronously
func (client *Client) SaveExternalAudienceWithChan(request *SaveExternalAudienceRequest) (<-chan *SaveExternalAudienceResponse, <-chan error) {
	responseChan := make(chan *SaveExternalAudienceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveExternalAudience(request)
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

// SaveExternalAudienceWithCallback invokes the retailadvqa_public.SaveExternalAudience API asynchronously
func (client *Client) SaveExternalAudienceWithCallback(request *SaveExternalAudienceRequest, callback func(response *SaveExternalAudienceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveExternalAudienceResponse
		var err error
		defer close(result)
		response, err = client.SaveExternalAudience(request)
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

// SaveExternalAudienceRequest is the request struct for api SaveExternalAudience
type SaveExternalAudienceRequest struct {
	*requests.RpcRequest
	AccessId     string `position:"Query" name:"AccessId"`
	AudienceName string `position:"Query" name:"AudienceName"`
	MappingType  string `position:"Query" name:"MappingType"`
	DataSourceId string `position:"Query" name:"DataSourceId"`
	AudienceId   string `position:"Query" name:"AudienceId"`
	ParentId     string `position:"Query" name:"ParentId"`
	WorkspaceId  string `position:"Query" name:"WorkspaceId"`
}

// SaveExternalAudienceResponse is the response struct for api SaveExternalAudience
type SaveExternalAudienceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSaveExternalAudienceRequest creates a request to invoke SaveExternalAudience API
func CreateSaveExternalAudienceRequest() (request *SaveExternalAudienceRequest) {
	request = &SaveExternalAudienceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "SaveExternalAudience", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveExternalAudienceResponse creates a response to parse from SaveExternalAudience response
func CreateSaveExternalAudienceResponse() (response *SaveExternalAudienceResponse) {
	response = &SaveExternalAudienceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
