/* 
 * SendX API
 *
 * SendX is built on the simple tenet that users must have open access to their data. SendX API is the first step in that direction. To cite some examples:   - subscribe / unsubscribe a contact from a list   - Schedule campaign to a segment of users   - Trigger transactional emails   - Get / PUT / POST and DELETE operations on team, campaign, list, contact, report etc. and so on.  As companies grow big, custom use cases around email marketing also crop up. SendX API ensures that SendX platform is able to satisfy such unforeseen use cases. They may range from building custom reporting dashboard to tagging contacts with custom attributes or triggering emails based on recommendation algorithm.  We do our best to have all our URLs be [RESTful](http://en.wikipedia.org/wiki/Representational_state_transfer). Every endpoint (URL) may support one of four different http verbs. GET requests fetch information about an object, POST requests create objects, PUT requests update objects, and finally DELETE requests will delete objects.  Also all API calls besides:   - Subscribe / unsubscribe signup form required **api_key** to be passed as **header**   ### The Envelope Every response is contained by an envelope. That is, each response has a predictable set of keys with which you can expect to interact: ```json {     \"status\": \"200\",     \"message\": \"OK\",     \"data\"\": [        {          ...        },        .        .        .     ] } ```  #### Status The status key is used to communicate extra information about the response to the developer. If all goes well, you'll only ever see a code key with value 200. However, sometimes things go wrong, and in that case you might see a response like: ```json {     \"status\": \"404\" } ```  #### Data The data key is the meat of the response. It may be a list containing single object or multiple objects  #### Message This returns back human readable message. This is specially useful to make sense in case of error scenarios. 
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"net/url"
	"github.com/go-resty/resty"
)

type APIClient struct {
}

func (c *APIClient) SelectHeaderContentType(contentTypes []string) string {

	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

func (c *APIClient) SelectHeaderAccept(accepts []string) string {

	if len(accepts) == 0 {
		return ""
	}
	if contains(accepts, "application/json") {
		return "application/json"
	}
	return strings.Join(accepts, ",")
}

func contains(source []string, containvalue string) bool {
	for _, a := range source {
		if strings.ToLower(a) == strings.ToLower(containvalue) {
			return true
		}
	}
	return false
}

func (c *APIClient) CallAPI(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) (*resty.Response, error) {

	//set debug flag
	configuration := NewConfiguration()
	resty.SetDebug(configuration.GetDebug())

	request := prepareRequest(postBody, headerParams, queryParams, formParams, fileName, fileBytes)

	switch strings.ToUpper(method) {
	case "GET":
		response, err := request.Get(path)
		return response, err
	case "POST":
		response, err := request.Post(path)
		return response, err
	case "PUT":
		response, err := request.Put(path)
		return response, err
	case "PATCH":
		response, err := request.Patch(path)
		return response, err
	case "DELETE":
		response, err := request.Delete(path)
		return response, err
	}

	return nil, fmt.Errorf("invalid method %v", method)
}

func (c *APIClient) ParameterToString(obj interface{},collectionFormat string) string {
	if reflect.TypeOf(obj).String() == "[]string" {
		switch	collectionFormat {
		case "pipes":
			return strings.Join(obj.([]string), "|")
		case "ssv":
			return strings.Join(obj.([]string), " ")
		case "tsv":
			return strings.Join(obj.([]string), "\t")	
		case "csv" :
			return strings.Join(obj.([]string), ",")
		}
	}

	return fmt.Sprintf("%v", obj)
}

func prepareRequest(postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) *resty.Request {

	request := resty.R()
	request.SetBody(postBody)

	// add header parameter, if any
	if len(headerParams) > 0 {
		request.SetHeaders(headerParams)
	}

	// add query parameter, if any
	if len(queryParams) > 0 {
		request.SetMultiValueQueryParams(queryParams)
	}

	// add form parameter, if any
	if len(formParams) > 0 {
		request.SetFormData(formParams)
	}

	if len(fileBytes) > 0 && fileName != "" {
		_, fileNm := filepath.Split(fileName)
		request.SetFileReader("file", fileNm, bytes.NewReader(fileBytes))
	}
	return request
}
