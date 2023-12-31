package hbase

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

// CreateHBaseSlbServer invokes the hbase.CreateHBaseSlbServer API synchronously
func (client *Client) CreateHBaseSlbServer(request *CreateHBaseSlbServerRequest) (response *CreateHBaseSlbServerResponse, err error) {
	response = CreateCreateHBaseSlbServerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateHBaseSlbServerWithChan invokes the hbase.CreateHBaseSlbServer API asynchronously
func (client *Client) CreateHBaseSlbServerWithChan(request *CreateHBaseSlbServerRequest) (<-chan *CreateHBaseSlbServerResponse, <-chan error) {
	responseChan := make(chan *CreateHBaseSlbServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHBaseSlbServer(request)
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

// CreateHBaseSlbServerWithCallback invokes the hbase.CreateHBaseSlbServer API asynchronously
func (client *Client) CreateHBaseSlbServerWithCallback(request *CreateHBaseSlbServerRequest, callback func(response *CreateHBaseSlbServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHBaseSlbServerResponse
		var err error
		defer close(result)
		response, err = client.CreateHBaseSlbServer(request)
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

// CreateHBaseSlbServerRequest is the request struct for api CreateHBaseSlbServer
type CreateHBaseSlbServerRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	ClusterId   string `position:"Query" name:"ClusterId"`
	SlbServer   string `position:"Query" name:"SlbServer"`
}

// CreateHBaseSlbServerResponse is the response struct for api CreateHBaseSlbServer
type CreateHBaseSlbServerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateHBaseSlbServerRequest creates a request to invoke CreateHBaseSlbServer API
func CreateCreateHBaseSlbServerRequest() (request *CreateHBaseSlbServerRequest) {
	request = &CreateHBaseSlbServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "CreateHBaseSlbServer", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateHBaseSlbServerResponse creates a response to parse from CreateHBaseSlbServer response
func CreateCreateHBaseSlbServerResponse() (response *CreateHBaseSlbServerResponse) {
	response = &CreateHBaseSlbServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
