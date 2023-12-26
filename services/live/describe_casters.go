package live

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

// DescribeCasters invokes the live.DescribeCasters API synchronously
func (client *Client) DescribeCasters(request *DescribeCastersRequest) (response *DescribeCastersResponse, err error) {
	response = CreateDescribeCastersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCastersWithChan invokes the live.DescribeCasters API asynchronously
func (client *Client) DescribeCastersWithChan(request *DescribeCastersRequest) (<-chan *DescribeCastersResponse, <-chan error) {
	responseChan := make(chan *DescribeCastersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCasters(request)
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

// DescribeCastersWithCallback invokes the live.DescribeCasters API asynchronously
func (client *Client) DescribeCastersWithCallback(request *DescribeCastersRequest, callback func(response *DescribeCastersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCastersResponse
		var err error
		defer close(result)
		response, err = client.DescribeCasters(request)
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

// DescribeCastersRequest is the request struct for api DescribeCasters
type DescribeCastersRequest struct {
	*requests.RpcRequest
	StartTime        string           `position:"Query" name:"StartTime"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
	CasterName       string           `position:"Query" name:"CasterName"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	NormType         string           `position:"Query" name:"NormType"`
	CasterId         string           `position:"Query" name:"CasterId"`
	EndTime          string           `position:"Query" name:"EndTime"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	OrderByModifyAsc string           `position:"Query" name:"OrderByModifyAsc"`
	ChargeType       requests.Integer `position:"Query" name:"ChargeType"`
	Status           requests.Integer `position:"Query" name:"Status"`
}

// DescribeCastersResponse is the response struct for api DescribeCasters
type DescribeCastersResponse struct {
	*responses.BaseResponse
	Total      int        `json:"Total" xml:"Total"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	CasterList CasterList `json:"CasterList" xml:"CasterList"`
}

// CreateDescribeCastersRequest creates a request to invoke DescribeCasters API
func CreateDescribeCastersRequest() (request *DescribeCastersRequest) {
	request = &DescribeCastersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeCasters", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCastersResponse creates a response to parse from DescribeCasters response
func CreateDescribeCastersResponse() (response *DescribeCastersResponse) {
	response = &DescribeCastersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
