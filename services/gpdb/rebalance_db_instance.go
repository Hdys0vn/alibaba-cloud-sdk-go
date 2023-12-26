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

// RebalanceDBInstance invokes the gpdb.RebalanceDBInstance API synchronously
func (client *Client) RebalanceDBInstance(request *RebalanceDBInstanceRequest) (response *RebalanceDBInstanceResponse, err error) {
	response = CreateRebalanceDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RebalanceDBInstanceWithChan invokes the gpdb.RebalanceDBInstance API asynchronously
func (client *Client) RebalanceDBInstanceWithChan(request *RebalanceDBInstanceRequest) (<-chan *RebalanceDBInstanceResponse, <-chan error) {
	responseChan := make(chan *RebalanceDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RebalanceDBInstance(request)
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

// RebalanceDBInstanceWithCallback invokes the gpdb.RebalanceDBInstance API asynchronously
func (client *Client) RebalanceDBInstanceWithCallback(request *RebalanceDBInstanceRequest, callback func(response *RebalanceDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RebalanceDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.RebalanceDBInstance(request)
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

// RebalanceDBInstanceRequest is the request struct for api RebalanceDBInstance
type RebalanceDBInstanceRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Query" name:"ClientToken"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// RebalanceDBInstanceResponse is the response struct for api RebalanceDBInstance
type RebalanceDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRebalanceDBInstanceRequest creates a request to invoke RebalanceDBInstance API
func CreateRebalanceDBInstanceRequest() (request *RebalanceDBInstanceRequest) {
	request = &RebalanceDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "RebalanceDBInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateRebalanceDBInstanceResponse creates a response to parse from RebalanceDBInstance response
func CreateRebalanceDBInstanceResponse() (response *RebalanceDBInstanceResponse) {
	response = &RebalanceDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
