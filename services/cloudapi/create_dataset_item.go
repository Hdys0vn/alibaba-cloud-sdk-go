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

// CreateDatasetItem invokes the cloudapi.CreateDatasetItem API synchronously
func (client *Client) CreateDatasetItem(request *CreateDatasetItemRequest) (response *CreateDatasetItemResponse, err error) {
	response = CreateCreateDatasetItemResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDatasetItemWithChan invokes the cloudapi.CreateDatasetItem API asynchronously
func (client *Client) CreateDatasetItemWithChan(request *CreateDatasetItemRequest) (<-chan *CreateDatasetItemResponse, <-chan error) {
	responseChan := make(chan *CreateDatasetItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDatasetItem(request)
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

// CreateDatasetItemWithCallback invokes the cloudapi.CreateDatasetItem API asynchronously
func (client *Client) CreateDatasetItemWithCallback(request *CreateDatasetItemRequest, callback func(response *CreateDatasetItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDatasetItemResponse
		var err error
		defer close(result)
		response, err = client.CreateDatasetItem(request)
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

// CreateDatasetItemRequest is the request struct for api CreateDatasetItem
type CreateDatasetItemRequest struct {
	*requests.RpcRequest
	Description   string `position:"Query" name:"Description"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ExpiredTime   string `position:"Query" name:"ExpiredTime"`
	DatasetId     string `position:"Query" name:"DatasetId"`
	Value         string `position:"Query" name:"Value"`
}

// CreateDatasetItemResponse is the response struct for api CreateDatasetItem
type CreateDatasetItemResponse struct {
	*responses.BaseResponse
	DatasetItemId string `json:"DatasetItemId" xml:"DatasetItemId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDatasetItemRequest creates a request to invoke CreateDatasetItem API
func CreateCreateDatasetItemRequest() (request *CreateDatasetItemRequest) {
	request = &CreateDatasetItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateDatasetItem", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDatasetItemResponse creates a response to parse from CreateDatasetItem response
func CreateCreateDatasetItemResponse() (response *CreateDatasetItemResponse) {
	response = &CreateDatasetItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
