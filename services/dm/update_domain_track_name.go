package dm

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

// UpdateDomainTrackName invokes the dm.UpdateDomainTrackName API synchronously
// api document: https://help.aliyun.com/api/dm/updatedomaintrackname.html
func (client *Client) UpdateDomainTrackName(request *UpdateDomainTrackNameRequest) (response *UpdateDomainTrackNameResponse, err error) {
	response = CreateUpdateDomainTrackNameResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDomainTrackNameWithChan invokes the dm.UpdateDomainTrackName API asynchronously
// api document: https://help.aliyun.com/api/dm/updatedomaintrackname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDomainTrackNameWithChan(request *UpdateDomainTrackNameRequest) (<-chan *UpdateDomainTrackNameResponse, <-chan error) {
	responseChan := make(chan *UpdateDomainTrackNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDomainTrackName(request)
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

// UpdateDomainTrackNameWithCallback invokes the dm.UpdateDomainTrackName API asynchronously
// api document: https://help.aliyun.com/api/dm/updatedomaintrackname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDomainTrackNameWithCallback(request *UpdateDomainTrackNameRequest, callback func(response *UpdateDomainTrackNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDomainTrackNameResponse
		var err error
		defer close(result)
		response, err = client.UpdateDomainTrackName(request)
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

// UpdateDomainTrackNameRequest is the request struct for api UpdateDomainTrackName
type UpdateDomainTrackNameRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DomainId             requests.Integer `position:"Query" name:"DomainId"`
	CnameTrackRecord     string           `position:"Query" name:"CnameTrackRecord"`
}

// UpdateDomainTrackNameResponse is the response struct for api UpdateDomainTrackName
type UpdateDomainTrackNameResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDomainTrackNameRequest creates a request to invoke UpdateDomainTrackName API
func CreateUpdateDomainTrackNameRequest() (request *UpdateDomainTrackNameRequest) {
	request = &UpdateDomainTrackNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "UpdateDomainTrackName", "", "")
	return
}

// CreateUpdateDomainTrackNameResponse creates a response to parse from UpdateDomainTrackName response
func CreateUpdateDomainTrackNameResponse() (response *UpdateDomainTrackNameResponse) {
	response = &UpdateDomainTrackNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
