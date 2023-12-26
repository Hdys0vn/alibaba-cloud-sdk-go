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

// DescribeDataReDistributeInfo invokes the gpdb.DescribeDataReDistributeInfo API synchronously
func (client *Client) DescribeDataReDistributeInfo(request *DescribeDataReDistributeInfoRequest) (response *DescribeDataReDistributeInfoResponse, err error) {
	response = CreateDescribeDataReDistributeInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataReDistributeInfoWithChan invokes the gpdb.DescribeDataReDistributeInfo API asynchronously
func (client *Client) DescribeDataReDistributeInfoWithChan(request *DescribeDataReDistributeInfoRequest) (<-chan *DescribeDataReDistributeInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDataReDistributeInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataReDistributeInfo(request)
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

// DescribeDataReDistributeInfoWithCallback invokes the gpdb.DescribeDataReDistributeInfo API asynchronously
func (client *Client) DescribeDataReDistributeInfoWithCallback(request *DescribeDataReDistributeInfoRequest, callback func(response *DescribeDataReDistributeInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataReDistributeInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataReDistributeInfo(request)
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

// DescribeDataReDistributeInfoRequest is the request struct for api DescribeDataReDistributeInfo
type DescribeDataReDistributeInfoRequest struct {
	*requests.RpcRequest
	DBInstanceId string           `position:"Query" name:"DBInstanceId"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDataReDistributeInfoResponse is the response struct for api DescribeDataReDistributeInfo
type DescribeDataReDistributeInfoResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	DataReDistributeInfo DataReDistributeInfo `json:"DataReDistributeInfo" xml:"DataReDistributeInfo"`
}

// CreateDescribeDataReDistributeInfoRequest creates a request to invoke DescribeDataReDistributeInfo API
func CreateDescribeDataReDistributeInfoRequest() (request *DescribeDataReDistributeInfoRequest) {
	request = &DescribeDataReDistributeInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDataReDistributeInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDataReDistributeInfoResponse creates a response to parse from DescribeDataReDistributeInfo response
func CreateDescribeDataReDistributeInfoResponse() (response *DescribeDataReDistributeInfoResponse) {
	response = &DescribeDataReDistributeInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}