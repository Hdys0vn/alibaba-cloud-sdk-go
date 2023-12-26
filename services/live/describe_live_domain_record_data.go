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

// DescribeLiveDomainRecordData invokes the live.DescribeLiveDomainRecordData API synchronously
func (client *Client) DescribeLiveDomainRecordData(request *DescribeLiveDomainRecordDataRequest) (response *DescribeLiveDomainRecordDataResponse, err error) {
	response = CreateDescribeLiveDomainRecordDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainRecordDataWithChan invokes the live.DescribeLiveDomainRecordData API asynchronously
func (client *Client) DescribeLiveDomainRecordDataWithChan(request *DescribeLiveDomainRecordDataRequest) (<-chan *DescribeLiveDomainRecordDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainRecordDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainRecordData(request)
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

// DescribeLiveDomainRecordDataWithCallback invokes the live.DescribeLiveDomainRecordData API asynchronously
func (client *Client) DescribeLiveDomainRecordDataWithCallback(request *DescribeLiveDomainRecordDataRequest, callback func(response *DescribeLiveDomainRecordDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainRecordDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainRecordData(request)
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

// DescribeLiveDomainRecordDataRequest is the request struct for api DescribeLiveDomainRecordData
type DescribeLiveDomainRecordDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	RecordType string           `position:"Query" name:"RecordType"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainRecordDataResponse is the response struct for api DescribeLiveDomainRecordData
type DescribeLiveDomainRecordDataResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	RecordDataInfos RecordDataInfos `json:"RecordDataInfos" xml:"RecordDataInfos"`
}

// CreateDescribeLiveDomainRecordDataRequest creates a request to invoke DescribeLiveDomainRecordData API
func CreateDescribeLiveDomainRecordDataRequest() (request *DescribeLiveDomainRecordDataRequest) {
	request = &DescribeLiveDomainRecordDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainRecordData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainRecordDataResponse creates a response to parse from DescribeLiveDomainRecordData response
func CreateDescribeLiveDomainRecordDataResponse() (response *DescribeLiveDomainRecordDataResponse) {
	response = &DescribeLiveDomainRecordDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
