package oos

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

// GetInventorySchema invokes the oos.GetInventorySchema API synchronously
func (client *Client) GetInventorySchema(request *GetInventorySchemaRequest) (response *GetInventorySchemaResponse, err error) {
	response = CreateGetInventorySchemaResponse()
	err = client.DoAction(request, response)
	return
}

// GetInventorySchemaWithChan invokes the oos.GetInventorySchema API asynchronously
func (client *Client) GetInventorySchemaWithChan(request *GetInventorySchemaRequest) (<-chan *GetInventorySchemaResponse, <-chan error) {
	responseChan := make(chan *GetInventorySchemaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInventorySchema(request)
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

// GetInventorySchemaWithCallback invokes the oos.GetInventorySchema API asynchronously
func (client *Client) GetInventorySchemaWithCallback(request *GetInventorySchemaRequest, callback func(response *GetInventorySchemaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInventorySchemaResponse
		var err error
		defer close(result)
		response, err = client.GetInventorySchema(request)
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

// GetInventorySchemaRequest is the request struct for api GetInventorySchema
type GetInventorySchemaRequest struct {
	*requests.RpcRequest
	Aggregator requests.Boolean `position:"Query" name:"Aggregator"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
	TypeName   string           `position:"Query" name:"TypeName"`
}

// GetInventorySchemaResponse is the response struct for api GetInventorySchema
type GetInventorySchemaResponse struct {
	*responses.BaseResponse
	NextToken  string   `json:"NextToken" xml:"NextToken"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	MaxResults string   `json:"MaxResults" xml:"MaxResults"`
	Schemas    []Schema `json:"Schemas" xml:"Schemas"`
}

// CreateGetInventorySchemaRequest creates a request to invoke GetInventorySchema API
func CreateGetInventorySchemaRequest() (request *GetInventorySchemaRequest) {
	request = &GetInventorySchemaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "GetInventorySchema", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetInventorySchemaResponse creates a response to parse from GetInventorySchema response
func CreateGetInventorySchemaResponse() (response *GetInventorySchemaResponse) {
	response = &GetInventorySchemaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}