package dg

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

// AddDatabase invokes the dg.AddDatabase API synchronously
func (client *Client) AddDatabase(request *AddDatabaseRequest) (response *AddDatabaseResponse, err error) {
	response = CreateAddDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// AddDatabaseWithChan invokes the dg.AddDatabase API asynchronously
func (client *Client) AddDatabaseWithChan(request *AddDatabaseRequest) (<-chan *AddDatabaseResponse, <-chan error) {
	responseChan := make(chan *AddDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDatabase(request)
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

// AddDatabaseWithCallback invokes the dg.AddDatabase API asynchronously
func (client *Client) AddDatabaseWithCallback(request *AddDatabaseRequest, callback func(response *AddDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDatabaseResponse
		var err error
		defer close(result)
		response, err = client.AddDatabase(request)
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

// AddDatabaseRequest is the request struct for api AddDatabase
type AddDatabaseRequest struct {
	*requests.RpcRequest
	ClientToken   string           `position:"Body" name:"ClientToken"`
	Host          string           `position:"Body" name:"Host"`
	DbUserName    string           `position:"Body" name:"DbUserName"`
	DbDescription string           `position:"Body" name:"DbDescription"`
	GatewayId     string           `position:"Body" name:"GatewayId"`
	DbName        string           `position:"Body" name:"DbName"`
	Port          requests.Integer `position:"Body" name:"Port"`
	DbPassword    string           `position:"Body" name:"DbPassword"`
	DbType        string           `position:"Body" name:"DbType"`
}

// AddDatabaseResponse is the response struct for api AddDatabase
type AddDatabaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateAddDatabaseRequest creates a request to invoke AddDatabase API
func CreateAddDatabaseRequest() (request *AddDatabaseRequest) {
	request = &AddDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dg", "2019-03-27", "AddDatabase", "dg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDatabaseResponse creates a response to parse from AddDatabase response
func CreateAddDatabaseResponse() (response *AddDatabaseResponse) {
	response = &AddDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
