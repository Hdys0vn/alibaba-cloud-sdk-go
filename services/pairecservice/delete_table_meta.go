package pairecservice

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

// DeleteTableMeta invokes the pairecservice.DeleteTableMeta API synchronously
func (client *Client) DeleteTableMeta(request *DeleteTableMetaRequest) (response *DeleteTableMetaResponse, err error) {
	response = CreateDeleteTableMetaResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTableMetaWithChan invokes the pairecservice.DeleteTableMeta API asynchronously
func (client *Client) DeleteTableMetaWithChan(request *DeleteTableMetaRequest) (<-chan *DeleteTableMetaResponse, <-chan error) {
	responseChan := make(chan *DeleteTableMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTableMeta(request)
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

// DeleteTableMetaWithCallback invokes the pairecservice.DeleteTableMeta API asynchronously
func (client *Client) DeleteTableMetaWithCallback(request *DeleteTableMetaRequest, callback func(response *DeleteTableMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTableMetaResponse
		var err error
		defer close(result)
		response, err = client.DeleteTableMeta(request)
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

// DeleteTableMetaRequest is the request struct for api DeleteTableMeta
type DeleteTableMetaRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	TableMetaId string `position:"Path" name:"TableMetaId"`
}

// DeleteTableMetaResponse is the response struct for api DeleteTableMeta
type DeleteTableMetaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTableMetaRequest creates a request to invoke DeleteTableMeta API
func CreateDeleteTableMetaRequest() (request *DeleteTableMetaRequest) {
	request = &DeleteTableMetaRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "DeleteTableMeta", "/api/v1/tablemetas/[TableMetaId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteTableMetaResponse creates a response to parse from DeleteTableMeta response
func CreateDeleteTableMetaResponse() (response *DeleteTableMetaResponse) {
	response = &DeleteTableMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
