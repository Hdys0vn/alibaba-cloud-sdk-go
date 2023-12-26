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

// DescribeDrdsTasks invokes the drds.DescribeDrdsTasks API synchronously
func (client *Client) DescribeDrdsTasks(request *DescribeDrdsTasksRequest) (response *DescribeDrdsTasksResponse, err error) {
	response = CreateDescribeDrdsTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsTasksWithChan invokes the drds.DescribeDrdsTasks API asynchronously
func (client *Client) DescribeDrdsTasksWithChan(request *DescribeDrdsTasksRequest) (<-chan *DescribeDrdsTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsTasks(request)
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

// DescribeDrdsTasksWithCallback invokes the drds.DescribeDrdsTasks API asynchronously
func (client *Client) DescribeDrdsTasksWithCallback(request *DescribeDrdsTasksRequest, callback func(response *DescribeDrdsTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsTasks(request)
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

// DescribeDrdsTasksRequest is the request struct for api DescribeDrdsTasks
type DescribeDrdsTasksRequest struct {
	*requests.RpcRequest
	TaskType       string `position:"Query" name:"TaskType"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// DescribeDrdsTasksResponse is the response struct for api DescribeDrdsTasks
type DescribeDrdsTasksResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Success   bool                     `json:"Success" xml:"Success"`
	Tasks     TasksInDescribeDrdsTasks `json:"Tasks" xml:"Tasks"`
}

// CreateDescribeDrdsTasksRequest creates a request to invoke DescribeDrdsTasks API
func CreateDescribeDrdsTasksRequest() (request *DescribeDrdsTasksRequest) {
	request = &DescribeDrdsTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsTasks", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDrdsTasksResponse creates a response to parse from DescribeDrdsTasks response
func CreateDescribeDrdsTasksResponse() (response *DescribeDrdsTasksResponse) {
	response = &DescribeDrdsTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}