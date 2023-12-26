package gpdb

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

// InitVectorDatabase invokes the gpdb.InitVectorDatabase API synchronously
func (client *Client) InitVectorDatabase(request *InitVectorDatabaseRequest) (response *InitVectorDatabaseResponse, err error) {
	response = CreateInitVectorDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// InitVectorDatabaseWithChan invokes the gpdb.InitVectorDatabase API asynchronously
func (client *Client) InitVectorDatabaseWithChan(request *InitVectorDatabaseRequest) (<-chan *InitVectorDatabaseResponse, <-chan error) {
	responseChan := make(chan *InitVectorDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitVectorDatabase(request)
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

// InitVectorDatabaseWithCallback invokes the gpdb.InitVectorDatabase API asynchronously
func (client *Client) InitVectorDatabaseWithCallback(request *InitVectorDatabaseRequest, callback func(response *InitVectorDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitVectorDatabaseResponse
		var err error
		defer close(result)
		response, err = client.InitVectorDatabase(request)
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

// InitVectorDatabaseRequest is the request struct for api InitVectorDatabase
type InitVectorDatabaseRequest struct {
	*requests.RpcRequest
	ManagerAccount         string           `position:"Query" name:"ManagerAccount"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ManagerAccountPassword string           `position:"Query" name:"ManagerAccountPassword"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

// InitVectorDatabaseResponse is the response struct for api InitVectorDatabase
type InitVectorDatabaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateInitVectorDatabaseRequest creates a request to invoke InitVectorDatabase API
func CreateInitVectorDatabaseRequest() (request *InitVectorDatabaseRequest) {
	request = &InitVectorDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "InitVectorDatabase", "", "")
	request.Method = requests.POST
	return
}

// CreateInitVectorDatabaseResponse creates a response to parse from InitVectorDatabase response
func CreateInitVectorDatabaseResponse() (response *InitVectorDatabaseResponse) {
	response = &InitVectorDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
