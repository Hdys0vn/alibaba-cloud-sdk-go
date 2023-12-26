package dataworks_public

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

// GetQualityEntity invokes the dataworks_public.GetQualityEntity API synchronously
func (client *Client) GetQualityEntity(request *GetQualityEntityRequest) (response *GetQualityEntityResponse, err error) {
	response = CreateGetQualityEntityResponse()
	err = client.DoAction(request, response)
	return
}

// GetQualityEntityWithChan invokes the dataworks_public.GetQualityEntity API asynchronously
func (client *Client) GetQualityEntityWithChan(request *GetQualityEntityRequest) (<-chan *GetQualityEntityResponse, <-chan error) {
	responseChan := make(chan *GetQualityEntityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQualityEntity(request)
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

// GetQualityEntityWithCallback invokes the dataworks_public.GetQualityEntity API asynchronously
func (client *Client) GetQualityEntityWithCallback(request *GetQualityEntityRequest, callback func(response *GetQualityEntityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQualityEntityResponse
		var err error
		defer close(result)
		response, err = client.GetQualityEntity(request)
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

// GetQualityEntityRequest is the request struct for api GetQualityEntity
type GetQualityEntityRequest struct {
	*requests.RpcRequest
	ProjectName     string           `position:"Body" name:"ProjectName"`
	MatchExpression string           `position:"Body" name:"MatchExpression"`
	EnvType         string           `position:"Body" name:"EnvType"`
	TableName       string           `position:"Body" name:"TableName"`
	ProjectId       requests.Integer `position:"Body" name:"ProjectId"`
}

// GetQualityEntityResponse is the response struct for api GetQualityEntity
type GetQualityEntityResponse struct {
	*responses.BaseResponse
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string      `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string      `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool        `json:"Success" xml:"Success"`
	Data           []EntityDto `json:"Data" xml:"Data"`
}

// CreateGetQualityEntityRequest creates a request to invoke GetQualityEntity API
func CreateGetQualityEntityRequest() (request *GetQualityEntityRequest) {
	request = &GetQualityEntityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetQualityEntity", "", "")
	request.Method = requests.POST
	return
}

// CreateGetQualityEntityResponse creates a response to parse from GetQualityEntity response
func CreateGetQualityEntityResponse() (response *GetQualityEntityResponse) {
	response = &GetQualityEntityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}