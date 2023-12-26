package facebody

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

// UpdateFaceEntity invokes the facebody.UpdateFaceEntity API synchronously
func (client *Client) UpdateFaceEntity(request *UpdateFaceEntityRequest) (response *UpdateFaceEntityResponse, err error) {
	response = CreateUpdateFaceEntityResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFaceEntityWithChan invokes the facebody.UpdateFaceEntity API asynchronously
func (client *Client) UpdateFaceEntityWithChan(request *UpdateFaceEntityRequest) (<-chan *UpdateFaceEntityResponse, <-chan error) {
	responseChan := make(chan *UpdateFaceEntityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFaceEntity(request)
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

// UpdateFaceEntityWithCallback invokes the facebody.UpdateFaceEntity API asynchronously
func (client *Client) UpdateFaceEntityWithCallback(request *UpdateFaceEntityRequest, callback func(response *UpdateFaceEntityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFaceEntityResponse
		var err error
		defer close(result)
		response, err = client.UpdateFaceEntity(request)
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

// UpdateFaceEntityRequest is the request struct for api UpdateFaceEntity
type UpdateFaceEntityRequest struct {
	*requests.RpcRequest
	EntityId           string           `position:"Body" name:"EntityId"`
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	OssFile            string           `position:"Query" name:"OssFile"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
	Labels             string           `position:"Body" name:"Labels"`
	DbName             string           `position:"Body" name:"DbName"`
}

// UpdateFaceEntityResponse is the response struct for api UpdateFaceEntity
type UpdateFaceEntityResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateFaceEntityRequest creates a request to invoke UpdateFaceEntity API
func CreateUpdateFaceEntityRequest() (request *UpdateFaceEntityRequest) {
	request = &UpdateFaceEntityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "UpdateFaceEntity", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateFaceEntityResponse creates a response to parse from UpdateFaceEntity response
func CreateUpdateFaceEntityResponse() (response *UpdateFaceEntityResponse) {
	response = &UpdateFaceEntityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
