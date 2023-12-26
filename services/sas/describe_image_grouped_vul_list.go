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

// DescribeImageGroupedVulList invokes the sas.DescribeImageGroupedVulList API synchronously
func (client *Client) DescribeImageGroupedVulList(request *DescribeImageGroupedVulListRequest) (response *DescribeImageGroupedVulListResponse, err error) {
	response = CreateDescribeImageGroupedVulListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageGroupedVulListWithChan invokes the sas.DescribeImageGroupedVulList API asynchronously
func (client *Client) DescribeImageGroupedVulListWithChan(request *DescribeImageGroupedVulListRequest) (<-chan *DescribeImageGroupedVulListResponse, <-chan error) {
	responseChan := make(chan *DescribeImageGroupedVulListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageGroupedVulList(request)
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

// DescribeImageGroupedVulListWithCallback invokes the sas.DescribeImageGroupedVulList API asynchronously
func (client *Client) DescribeImageGroupedVulListWithCallback(request *DescribeImageGroupedVulListRequest, callback func(response *DescribeImageGroupedVulListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageGroupedVulListResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageGroupedVulList(request)
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

// DescribeImageGroupedVulListRequest is the request struct for api DescribeImageGroupedVulList
type DescribeImageGroupedVulListRequest struct {
	*requests.RpcRequest
	Type           string           `position:"Query" name:"Type"`
	LastTsEnd      requests.Integer `position:"Query" name:"LastTsEnd"`
	CreateTsStart  requests.Integer `position:"Query" name:"CreateTsStart"`
	IsLatest       requests.Integer `position:"Query" name:"IsLatest"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	ImageTag       string           `position:"Query" name:"ImageTag"`
	Level          string           `position:"Query" name:"Level"`
	GroupId        string           `position:"Query" name:"GroupId"`
	OrderBy        string           `position:"Query" name:"OrderBy"`
	AliasName      string           `position:"Query" name:"AliasName"`
	PatchId        requests.Integer `position:"Query" name:"PatchId"`
	Name           string           `position:"Query" name:"Name"`
	CreateTsEnd    requests.Integer `position:"Query" name:"CreateTsEnd"`
	Necessity      string           `position:"Query" name:"Necessity"`
	Uuids          string           `position:"Query" name:"Uuids"`
	RepoId         string           `position:"Query" name:"RepoId"`
	StatusList     string           `position:"Query" name:"StatusList"`
	CveId          string           `position:"Query" name:"CveId"`
	Remark         string           `position:"Query" name:"Remark"`
	RepoNamespace  string           `position:"Query" name:"RepoNamespace"`
	ImageDigest    string           `position:"Query" name:"ImageDigest"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	Lang           string           `position:"Query" name:"Lang"`
	LastTsStart    requests.Integer `position:"Query" name:"LastTsStart"`
	Direction      string           `position:"Query" name:"Direction"`
	Dealed         string           `position:"Query" name:"Dealed"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	ClusterId      string           `position:"Query" name:"ClusterId"`
	SearchTags     string           `position:"Query" name:"SearchTags"`
	RepoName       string           `position:"Query" name:"RepoName"`
	RepoInstanceId string           `position:"Query" name:"RepoInstanceId"`
	ImageLayer     string           `position:"Query" name:"ImageLayer"`
	RepoRegionId   string           `position:"Query" name:"RepoRegionId"`
}

// DescribeImageGroupedVulListResponse is the response struct for api DescribeImageGroupedVulList
type DescribeImageGroupedVulListResponse struct {
	*responses.BaseResponse
	CurrentPage     int              `json:"CurrentPage" xml:"CurrentPage"`
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	PageSize        int              `json:"PageSize" xml:"PageSize"`
	TotalCount      int              `json:"TotalCount" xml:"TotalCount"`
	GroupedVulItems []GroupedVulItem `json:"GroupedVulItems" xml:"GroupedVulItems"`
}

// CreateDescribeImageGroupedVulListRequest creates a request to invoke DescribeImageGroupedVulList API
func CreateDescribeImageGroupedVulListRequest() (request *DescribeImageGroupedVulListRequest) {
	request = &DescribeImageGroupedVulListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeImageGroupedVulList", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeImageGroupedVulListResponse creates a response to parse from DescribeImageGroupedVulList response
func CreateDescribeImageGroupedVulListResponse() (response *DescribeImageGroupedVulListResponse) {
	response = &DescribeImageGroupedVulListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
