package r_kvstore

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

// DescribeUserClusterHost invokes the r_kvstore.DescribeUserClusterHost API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeuserclusterhost.html
func (client *Client) DescribeUserClusterHost(request *DescribeUserClusterHostRequest) (response *DescribeUserClusterHostResponse, err error) {
	response = CreateDescribeUserClusterHostResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserClusterHostWithChan invokes the r_kvstore.DescribeUserClusterHost API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeuserclusterhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserClusterHostWithChan(request *DescribeUserClusterHostRequest) (<-chan *DescribeUserClusterHostResponse, <-chan error) {
	responseChan := make(chan *DescribeUserClusterHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserClusterHost(request)
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

// DescribeUserClusterHostWithCallback invokes the r_kvstore.DescribeUserClusterHost API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeuserclusterhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserClusterHostWithCallback(request *DescribeUserClusterHostRequest, callback func(response *DescribeUserClusterHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserClusterHostResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserClusterHost(request)
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

// DescribeUserClusterHostRequest is the request struct for api DescribeUserClusterHost
type DescribeUserClusterHostRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           string           `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Engine               string           `position:"Query" name:"Engine"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MaxRecordsPerPage    string           `position:"Query" name:"MaxRecordsPerPage"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeUserClusterHostResponse is the response struct for api DescribeUserClusterHost
type DescribeUserClusterHostResponse struct {
	*responses.BaseResponse
	RequestId         string     `json:"RequestId" xml:"RequestId"`
	MaxRecordsPerPage string     `json:"MaxRecordsPerPage" xml:"MaxRecordsPerPage"`
	PageNumber        int        `json:"PageNumber" xml:"PageNumber"`
	TotalRecords      int        `json:"TotalRecords" xml:"TotalRecords"`
	ItemNumbers       int        `json:"ItemNumbers" xml:"ItemNumbers"`
	HostItems         []HostInfo `json:"HostItems" xml:"HostItems"`
}

// CreateDescribeUserClusterHostRequest creates a request to invoke DescribeUserClusterHost API
func CreateDescribeUserClusterHostRequest() (request *DescribeUserClusterHostRequest) {
	request = &DescribeUserClusterHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeUserClusterHost", "redisa", "openAPI")
	return
}

// CreateDescribeUserClusterHostResponse creates a response to parse from DescribeUserClusterHost response
func CreateDescribeUserClusterHostResponse() (response *DescribeUserClusterHostResponse) {
	response = &DescribeUserClusterHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
