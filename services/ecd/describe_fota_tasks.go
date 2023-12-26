package ecd

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

// DescribeFotaTasks invokes the ecd.DescribeFotaTasks API synchronously
func (client *Client) DescribeFotaTasks(request *DescribeFotaTasksRequest) (response *DescribeFotaTasksResponse, err error) {
	response = CreateDescribeFotaTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFotaTasksWithChan invokes the ecd.DescribeFotaTasks API asynchronously
func (client *Client) DescribeFotaTasksWithChan(request *DescribeFotaTasksRequest) (<-chan *DescribeFotaTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeFotaTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFotaTasks(request)
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

// DescribeFotaTasksWithCallback invokes the ecd.DescribeFotaTasks API asynchronously
func (client *Client) DescribeFotaTasksWithCallback(request *DescribeFotaTasksRequest, callback func(response *DescribeFotaTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFotaTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeFotaTasks(request)
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

// DescribeFotaTasksRequest is the request struct for api DescribeFotaTasks
type DescribeFotaTasksRequest struct {
	*requests.RpcRequest
	UserStatus string           `position:"Query" name:"UserStatus"`
	FotaStatus string           `position:"Query" name:"FotaStatus"`
	TaskUid    *[]string        `position:"Query" name:"TaskUid"  type:"Repeated"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// DescribeFotaTasksResponse is the response struct for api DescribeFotaTasks
type DescribeFotaTasksResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	FotaTasks []FotaTask `json:"FotaTasks" xml:"FotaTasks"`
}

// CreateDescribeFotaTasksRequest creates a request to invoke DescribeFotaTasks API
func CreateDescribeFotaTasksRequest() (request *DescribeFotaTasksRequest) {
	request = &DescribeFotaTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeFotaTasks", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeFotaTasksResponse creates a response to parse from DescribeFotaTasks response
func CreateDescribeFotaTasksResponse() (response *DescribeFotaTasksResponse) {
	response = &DescribeFotaTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}