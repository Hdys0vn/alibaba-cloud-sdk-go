package vs

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

// DescribeVsDomainUvData invokes the vs.DescribeVsDomainUvData API synchronously
func (client *Client) DescribeVsDomainUvData(request *DescribeVsDomainUvDataRequest) (response *DescribeVsDomainUvDataResponse, err error) {
	response = CreateDescribeVsDomainUvDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVsDomainUvDataWithChan invokes the vs.DescribeVsDomainUvData API asynchronously
func (client *Client) DescribeVsDomainUvDataWithChan(request *DescribeVsDomainUvDataRequest) (<-chan *DescribeVsDomainUvDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVsDomainUvDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVsDomainUvData(request)
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

// DescribeVsDomainUvDataWithCallback invokes the vs.DescribeVsDomainUvData API asynchronously
func (client *Client) DescribeVsDomainUvDataWithCallback(request *DescribeVsDomainUvDataRequest, callback func(response *DescribeVsDomainUvDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVsDomainUvDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVsDomainUvData(request)
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

// DescribeVsDomainUvDataRequest is the request struct for api DescribeVsDomainUvData
type DescribeVsDomainUvDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVsDomainUvDataResponse is the response struct for api DescribeVsDomainUvData
type DescribeVsDomainUvDataResponse struct {
	*responses.BaseResponse
	EndTime        string         `json:"EndTime" xml:"EndTime"`
	StartTime      string         `json:"StartTime" xml:"StartTime"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	DomainName     string         `json:"DomainName" xml:"DomainName"`
	DataInterval   string         `json:"DataInterval" xml:"DataInterval"`
	UvDataInterval UvDataInterval `json:"UvDataInterval" xml:"UvDataInterval"`
}

// CreateDescribeVsDomainUvDataRequest creates a request to invoke DescribeVsDomainUvData API
func CreateDescribeVsDomainUvDataRequest() (request *DescribeVsDomainUvDataRequest) {
	request = &DescribeVsDomainUvDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeVsDomainUvData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVsDomainUvDataResponse creates a response to parse from DescribeVsDomainUvData response
func CreateDescribeVsDomainUvDataResponse() (response *DescribeVsDomainUvDataResponse) {
	response = &DescribeVsDomainUvDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
