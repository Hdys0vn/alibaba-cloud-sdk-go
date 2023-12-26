package dbs

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

// DescribeDLAService invokes the dbs.DescribeDLAService API synchronously
func (client *Client) DescribeDLAService(request *DescribeDLAServiceRequest) (response *DescribeDLAServiceResponse, err error) {
	response = CreateDescribeDLAServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDLAServiceWithChan invokes the dbs.DescribeDLAService API asynchronously
func (client *Client) DescribeDLAServiceWithChan(request *DescribeDLAServiceRequest) (<-chan *DescribeDLAServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeDLAServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDLAService(request)
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

// DescribeDLAServiceWithCallback invokes the dbs.DescribeDLAService API asynchronously
func (client *Client) DescribeDLAServiceWithCallback(request *DescribeDLAServiceRequest, callback func(response *DescribeDLAServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDLAServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDLAService(request)
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

// DescribeDLAServiceRequest is the request struct for api DescribeDLAService
type DescribeDLAServiceRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Query" name:"ClientToken"`
	BackupPlanId string `position:"Query" name:"BackupPlanId"`
	OwnerId      string `position:"Query" name:"OwnerId"`
}

// DescribeDLAServiceResponse is the response struct for api DescribeDLAService
type DescribeDLAServiceResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	State          string `json:"State" xml:"State"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	AutoAdd        bool   `json:"AutoAdd" xml:"AutoAdd"`
	HaveJobFailed  bool   `json:"HaveJobFailed" xml:"HaveJobFailed"`
}

// CreateDescribeDLAServiceRequest creates a request to invoke DescribeDLAService API
func CreateDescribeDLAServiceRequest() (request *DescribeDLAServiceRequest) {
	request = &DescribeDLAServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribeDLAService", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDLAServiceResponse creates a response to parse from DescribeDLAService response
func CreateDescribeDLAServiceResponse() (response *DescribeDLAServiceResponse) {
	response = &DescribeDLAServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}