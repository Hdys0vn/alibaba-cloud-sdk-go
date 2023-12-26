package polardb

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

// CreateStoragePlan invokes the polardb.CreateStoragePlan API synchronously
func (client *Client) CreateStoragePlan(request *CreateStoragePlanRequest) (response *CreateStoragePlanResponse, err error) {
	response = CreateCreateStoragePlanResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStoragePlanWithChan invokes the polardb.CreateStoragePlan API asynchronously
func (client *Client) CreateStoragePlanWithChan(request *CreateStoragePlanRequest) (<-chan *CreateStoragePlanResponse, <-chan error) {
	responseChan := make(chan *CreateStoragePlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStoragePlan(request)
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

// CreateStoragePlanWithCallback invokes the polardb.CreateStoragePlan API asynchronously
func (client *Client) CreateStoragePlanWithCallback(request *CreateStoragePlanRequest, callback func(response *CreateStoragePlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStoragePlanResponse
		var err error
		defer close(result)
		response, err = client.CreateStoragePlan(request)
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

// CreateStoragePlanRequest is the request struct for api CreateStoragePlan
type CreateStoragePlanRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	StorageType          string           `position:"Query" name:"StorageType"`
	Period               string           `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime             string           `position:"Query" name:"UsedTime"`
	StorageClass         string           `position:"Query" name:"StorageClass"`
}

// CreateStoragePlanResponse is the response struct for api CreateStoragePlan
type CreateStoragePlanResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId      string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateStoragePlanRequest creates a request to invoke CreateStoragePlan API
func CreateCreateStoragePlanRequest() (request *CreateStoragePlanRequest) {
	request = &CreateStoragePlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CreateStoragePlan", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateStoragePlanResponse creates a response to parse from CreateStoragePlan response
func CreateCreateStoragePlanResponse() (response *CreateStoragePlanResponse) {
	response = &CreateStoragePlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}