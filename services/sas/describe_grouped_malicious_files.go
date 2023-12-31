package sas

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

// DescribeGroupedMaliciousFiles invokes the sas.DescribeGroupedMaliciousFiles API synchronously
func (client *Client) DescribeGroupedMaliciousFiles(request *DescribeGroupedMaliciousFilesRequest) (response *DescribeGroupedMaliciousFilesResponse, err error) {
	response = CreateDescribeGroupedMaliciousFilesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGroupedMaliciousFilesWithChan invokes the sas.DescribeGroupedMaliciousFiles API asynchronously
func (client *Client) DescribeGroupedMaliciousFilesWithChan(request *DescribeGroupedMaliciousFilesRequest) (<-chan *DescribeGroupedMaliciousFilesResponse, <-chan error) {
	responseChan := make(chan *DescribeGroupedMaliciousFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGroupedMaliciousFiles(request)
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

// DescribeGroupedMaliciousFilesWithCallback invokes the sas.DescribeGroupedMaliciousFiles API asynchronously
func (client *Client) DescribeGroupedMaliciousFilesWithCallback(request *DescribeGroupedMaliciousFilesRequest, callback func(response *DescribeGroupedMaliciousFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGroupedMaliciousFilesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGroupedMaliciousFiles(request)
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

// DescribeGroupedMaliciousFilesRequest is the request struct for api DescribeGroupedMaliciousFiles
type DescribeGroupedMaliciousFilesRequest struct {
	*requests.RpcRequest
	RepoId             string           `position:"Query" name:"RepoId"`
	FuzzyMaliciousName string           `position:"Query" name:"FuzzyMaliciousName"`
	RepoNamespace      string           `position:"Query" name:"RepoNamespace"`
	SourceIp           string           `position:"Query" name:"SourceIp"`
	ImageDigest        string           `position:"Query" name:"ImageDigest"`
	PageSize           string           `position:"Query" name:"PageSize"`
	Lang               string           `position:"Query" name:"Lang"`
	ImageTag           string           `position:"Query" name:"ImageTag"`
	CurrentPage        requests.Integer `position:"Query" name:"CurrentPage"`
	ClusterId          string           `position:"Query" name:"ClusterId"`
	RepoName           string           `position:"Query" name:"RepoName"`
	RepoInstanceId     string           `position:"Query" name:"RepoInstanceId"`
	ImageLayer         string           `position:"Query" name:"ImageLayer"`
	Levels             string           `position:"Query" name:"Levels"`
	RepoRegionId       string           `position:"Query" name:"RepoRegionId"`
	Uuids              *[]string        `position:"Query" name:"Uuids"  type:"Repeated"`
}

// DescribeGroupedMaliciousFilesResponse is the response struct for api DescribeGroupedMaliciousFiles
type DescribeGroupedMaliciousFilesResponse struct {
	*responses.BaseResponse
	RequestId                    string                 `json:"RequestId" xml:"RequestId"`
	PageInfo                     PageInfo               `json:"PageInfo" xml:"PageInfo"`
	GroupedMaliciousFileResponse []GroupedMaliciousFile `json:"GroupedMaliciousFileResponse" xml:"GroupedMaliciousFileResponse"`
}

// CreateDescribeGroupedMaliciousFilesRequest creates a request to invoke DescribeGroupedMaliciousFiles API
func CreateDescribeGroupedMaliciousFilesRequest() (request *DescribeGroupedMaliciousFilesRequest) {
	request = &DescribeGroupedMaliciousFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeGroupedMaliciousFiles", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGroupedMaliciousFilesResponse creates a response to parse from DescribeGroupedMaliciousFiles response
func CreateDescribeGroupedMaliciousFilesResponse() (response *DescribeGroupedMaliciousFilesResponse) {
	response = &DescribeGroupedMaliciousFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
