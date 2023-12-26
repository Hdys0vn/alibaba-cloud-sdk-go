package cdn

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

// BatchAddCdnDomain invokes the cdn.BatchAddCdnDomain API synchronously
func (client *Client) BatchAddCdnDomain(request *BatchAddCdnDomainRequest) (response *BatchAddCdnDomainResponse, err error) {
	response = CreateBatchAddCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// BatchAddCdnDomainWithChan invokes the cdn.BatchAddCdnDomain API asynchronously
func (client *Client) BatchAddCdnDomainWithChan(request *BatchAddCdnDomainRequest) (<-chan *BatchAddCdnDomainResponse, <-chan error) {
	responseChan := make(chan *BatchAddCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchAddCdnDomain(request)
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

// BatchAddCdnDomainWithCallback invokes the cdn.BatchAddCdnDomain API asynchronously
func (client *Client) BatchAddCdnDomainWithCallback(request *BatchAddCdnDomainRequest, callback func(response *BatchAddCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchAddCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.BatchAddCdnDomain(request)
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

// BatchAddCdnDomainRequest is the request struct for api BatchAddCdnDomain
type BatchAddCdnDomainRequest struct {
	*requests.RpcRequest
	Sources         string           `position:"Query" name:"Sources"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	CdnType         string           `position:"Query" name:"CdnType"`
	Scope           string           `position:"Query" name:"Scope"`
	TopLevelDomain  string           `position:"Query" name:"TopLevelDomain"`
	OwnerAccount    string           `position:"Query" name:"OwnerAccount"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	CheckUrl        string           `position:"Query" name:"CheckUrl"`
}

// BatchAddCdnDomainResponse is the response struct for api BatchAddCdnDomain
type BatchAddCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchAddCdnDomainRequest creates a request to invoke BatchAddCdnDomain API
func CreateBatchAddCdnDomainRequest() (request *BatchAddCdnDomainRequest) {
	request = &BatchAddCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "BatchAddCdnDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchAddCdnDomainResponse creates a response to parse from BatchAddCdnDomain response
func CreateBatchAddCdnDomainResponse() (response *BatchAddCdnDomainResponse) {
	response = &BatchAddCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
