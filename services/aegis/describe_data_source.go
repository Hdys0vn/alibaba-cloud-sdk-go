package aegis

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

// DescribeDataSource invokes the aegis.DescribeDataSource API synchronously
// api document: https://help.aliyun.com/api/aegis/describedatasource.html
func (client *Client) DescribeDataSource(request *DescribeDataSourceRequest) (response *DescribeDataSourceResponse, err error) {
	response = CreateDescribeDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataSourceWithChan invokes the aegis.DescribeDataSource API asynchronously
// api document: https://help.aliyun.com/api/aegis/describedatasource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataSourceWithChan(request *DescribeDataSourceRequest) (<-chan *DescribeDataSourceResponse, <-chan error) {
	responseChan := make(chan *DescribeDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataSource(request)
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

// DescribeDataSourceWithCallback invokes the aegis.DescribeDataSource API asynchronously
// api document: https://help.aliyun.com/api/aegis/describedatasource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataSourceWithCallback(request *DescribeDataSourceRequest, callback func(response *DescribeDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataSourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataSource(request)
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

// DescribeDataSourceRequest is the request struct for api DescribeDataSource
type DescribeDataSourceRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Description string           `position:"Query" name:"Description"`
	Id          requests.Integer `position:"Query" name:"Id"`
	Lang        string           `position:"Query" name:"Lang"`
	GroupName   string           `position:"Query" name:"GroupName"`
}

// DescribeDataSourceResponse is the response struct for api DescribeDataSource
type DescribeDataSourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Count     int    `json:"Count" xml:"Count"`
	MetaDatas []Data `json:"MetaDatas" xml:"MetaDatas"`
}

// CreateDescribeDataSourceRequest creates a request to invoke DescribeDataSource API
func CreateDescribeDataSourceRequest() (request *DescribeDataSourceRequest) {
	request = &DescribeDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeDataSource", "vipaegis", "openAPI")
	return
}

// CreateDescribeDataSourceResponse creates a response to parse from DescribeDataSource response
func CreateDescribeDataSourceResponse() (response *DescribeDataSourceResponse) {
	response = &DescribeDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
