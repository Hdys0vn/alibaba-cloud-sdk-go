package drds

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

// DescribeDrdsDB invokes the drds.DescribeDrdsDB API synchronously
func (client *Client) DescribeDrdsDB(request *DescribeDrdsDBRequest) (response *DescribeDrdsDBResponse, err error) {
	response = CreateDescribeDrdsDBResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsDBWithChan invokes the drds.DescribeDrdsDB API asynchronously
func (client *Client) DescribeDrdsDBWithChan(request *DescribeDrdsDBRequest) (<-chan *DescribeDrdsDBResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsDB(request)
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

// DescribeDrdsDBWithCallback invokes the drds.DescribeDrdsDB API asynchronously
func (client *Client) DescribeDrdsDBWithCallback(request *DescribeDrdsDBRequest, callback func(response *DescribeDrdsDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsDBResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsDB(request)
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

// DescribeDrdsDBRequest is the request struct for api DescribeDrdsDB
type DescribeDrdsDBRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// DescribeDrdsDBResponse is the response struct for api DescribeDrdsDB
type DescribeDrdsDBResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeDrdsDBRequest creates a request to invoke DescribeDrdsDB API
func CreateDescribeDrdsDBRequest() (request *DescribeDrdsDBRequest) {
	request = &DescribeDrdsDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsDB", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDrdsDBResponse creates a response to parse from DescribeDrdsDB response
func CreateDescribeDrdsDBResponse() (response *DescribeDrdsDBResponse) {
	response = &DescribeDrdsDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
