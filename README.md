# Go API client for swagger

SendX is built on the simple tenet that users must have open access to their data. SendX API is the first step in that direction. To cite some examples:   - subscribe / unsubscribe a contact from a list   - Schedule campaign to a segment of users   - Trigger transactional emails   - Get / PUT / POST and DELETE operations on team, campaign, list, contact, report etc. and so on.  As companies grow big, custom use cases around email marketing also crop up. SendX API ensures   that SendX platform is able to satisfy such unforeseen use cases. They may range from building     custom reporting dashboard to tagging contacts with custom attributes or triggering emails based on recommendation algorithm.  We do our best to have all our URLs be [RESTful](http://en.wikipedia.org/wiki/Representational_state_transfer). Every endpoint (URL) may support one of four different http verbs. GET requests fetch information about an object, POST requests create objects, PUT requests update objects, and finally DELETE requests will delete objects.  Also all API calls besides:   - Subscribe / unsubscribe signup form  required **api_key** to be passed as **header**   ### The Envelope Every response is contained by an envelope. That is, each response has a predictable set of keys with which you can expect to interact: ```json {     \"status\": \"200\",      \"message\": \"OK\",     \"data\"\": [        {          ...        },        .        .        .     ] } ```  #### Status  The status key is used to communicate extra information about the response to the developer. If all goes well, you'll only ever see a code key with value 200. However, sometimes things go wrong, and in that case you might see a response like: ```json {     \"status\": \"404\" } ```  #### Data  The data key is the meat of the response. It may be a list containing single object or multiple objects  #### Message  This returns back human readable message. This is specially useful to make sense in case of error scenarios. 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build date: 2016-08-15T11:32:31.682Z
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:8080/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CampaignApi* | [**CampaignCampaignIdDelete**](docs/CampaignApi.md#campaigncampaigniddelete) | **Delete** /campaign/{campaignId} | Deletes a campaign
*CampaignApi* | [**CampaignCampaignIdGet**](docs/CampaignApi.md#campaigncampaignidget) | **Get** /campaign/{campaignId} | Find campaign by ID
*CampaignApi* | [**CampaignCampaignIdPut**](docs/CampaignApi.md#campaigncampaignidput) | **Put** /campaign/{campaignId} | Update a campaign by ID
*CampaignApi* | [**CampaignGet**](docs/CampaignApi.md#campaignget) | **Get** /campaign | Get information about all campaigns
*CampaignApi* | [**CampaignPost**](docs/CampaignApi.md#campaignpost) | **Post** /campaign | Add a new campaign
*ContactApi* | [**ContactContactIdDelete**](docs/ContactApi.md#contactcontactiddelete) | **Delete** /contact/{contactId} | Deletes a contact
*ContactApi* | [**ContactContactIdGet**](docs/ContactApi.md#contactcontactidget) | **Get** /contact/{contactId} | Find contact by ID
*ContactApi* | [**ContactContactIdPut**](docs/ContactApi.md#contactcontactidput) | **Put** /contact/{contactId} | Update a contact by ID
*ContactApi* | [**ContactGet**](docs/ContactApi.md#contactget) | **Get** /contact | Get information about all contacts
*ContactApi* | [**ContactPost**](docs/ContactApi.md#contactpost) | **Post** /contact | Add a new contact
*LinkApi* | [**LinkGet**](docs/LinkApi.md#linkget) | **Get** /link | Get information about all links
*LinkApi* | [**LinkLinkIdDelete**](docs/LinkApi.md#linklinkiddelete) | **Delete** /link/{linkId} | Deletes a link
*LinkApi* | [**LinkLinkIdGet**](docs/LinkApi.md#linklinkidget) | **Get** /link/{linkId} | Find link by ID
*LinkApi* | [**LinkLinkIdPut**](docs/LinkApi.md#linklinkidput) | **Put** /link/{linkId} | Update a link by ID
*LinkApi* | [**LinkPost**](docs/LinkApi.md#linkpost) | **Post** /link | Add a new link
*ListApi* | [**ListGet**](docs/ListApi.md#listget) | **Get** /list | Get information about all lists
*ListApi* | [**ListListIdContactsGet**](docs/ListApi.md#listlistidcontactsget) | **Get** /list/{listId}/contacts | Find contacts belonging to a list
*ListApi* | [**ListListIdDelete**](docs/ListApi.md#listlistiddelete) | **Delete** /list/{listId} | Deletes a list
*ListApi* | [**ListListIdGet**](docs/ListApi.md#listlistidget) | **Get** /list/{listId} | Find list by ID
*ListApi* | [**ListListIdPut**](docs/ListApi.md#listlistidput) | **Put** /list/{listId} | Update a list by ID
*ListApi* | [**ListPost**](docs/ListApi.md#listpost) | **Post** /list | Add a new list
*SendApi* | [**SendEmailPost**](docs/SendApi.md#sendemailpost) | **Post** /send/email | Send transactional email to user
*SubscribeApi* | [**SubscribeEncryptedListIdPost**](docs/SubscribeApi.md#subscribeencryptedlistidpost) | **Post** /subscribe/{encryptedListId} | Subscribe a new user a list
*SubscribeApi* | [**SubscribeEncryptedListIdPut**](docs/SubscribeApi.md#subscribeencryptedlistidput) | **Put** /subscribe/{encryptedListId} | Subscribe an existing user a list
*TagApi* | [**TagGet**](docs/TagApi.md#tagget) | **Get** /tag | Get information about all tags
*TagApi* | [**TagPost**](docs/TagApi.md#tagpost) | **Post** /tag | Add a new tag
*TagApi* | [**TagTagIdContactDelete**](docs/TagApi.md#tagtagidcontactdelete) | **Delete** /tag/{tagId}/contact | Remove a contact from a tag
*TagApi* | [**TagTagIdContactPost**](docs/TagApi.md#tagtagidcontactpost) | **Post** /tag/{tagId}/contact | Add a contact to a tag
*TagApi* | [**TagTagIdContactsGet**](docs/TagApi.md#tagtagidcontactsget) | **Get** /tag/{tagId}/contacts | Find contacts belonging to a tag
*TagApi* | [**TagTagIdDelete**](docs/TagApi.md#tagtagiddelete) | **Delete** /tag/{tagId} | Deletes a tag
*TagApi* | [**TagTagIdGet**](docs/TagApi.md#tagtagidget) | **Get** /tag/{tagId} | Find tag by ID
*TagApi* | [**TagTagIdPut**](docs/TagApi.md#tagtagidput) | **Put** /tag/{tagId} | Update a tag by ID
*TeamApi* | [**TeamGet**](docs/TeamApi.md#teamget) | **Get** /team | Get information about all teams
*TeamApi* | [**TeamPost**](docs/TeamApi.md#teampost) | **Post** /team | Add a new team
*TeamApi* | [**TeamTeamIdCampaignsGet**](docs/TeamApi.md#teamteamidcampaignsget) | **Get** /team/{teamId}/campaigns | Find campaigns of a team by ID
*TeamApi* | [**TeamTeamIdContactsGet**](docs/TeamApi.md#teamteamidcontactsget) | **Get** /team/{teamId}/contacts | Find contacts of a team by ID
*TeamApi* | [**TeamTeamIdDelete**](docs/TeamApi.md#teamteamiddelete) | **Delete** /team/{teamId} | Deletes a team
*TeamApi* | [**TeamTeamIdGet**](docs/TeamApi.md#teamteamidget) | **Get** /team/{teamId} | Find team by ID
*TeamApi* | [**TeamTeamIdListsGet**](docs/TeamApi.md#teamteamidlistsget) | **Get** /team/{teamId}/lists | Find lists of a team by ID
*TeamApi* | [**TeamTeamIdPut**](docs/TeamApi.md#teamteamidput) | **Put** /team/{teamId} | Update a team by ID
*TeamApi* | [**TeamTeamIdTagsGet**](docs/TeamApi.md#teamteamidtagsget) | **Get** /team/{teamId}/tags | Find tags of a team by ID
*UnsubscribeApi* | [**UnsubscribeEncryptedListIdPost**](docs/UnsubscribeApi.md#unsubscribeencryptedlistidpost) | **Post** /unsubscribe/{encryptedListId} | Unsubscribe a user from list based on email id


## Documentation For Models

 - [Campaign](docs/Campaign.md)
 - [CampaignAddUpdate](docs/CampaignAddUpdate.md)
 - [Contact](docs/Contact.md)
 - [ContactAddUpdate](docs/ContactAddUpdate.md)
 - [DeepListEmailContact](docs/DeepListEmailContact.md)
 - [DeepTeamEmailContact](docs/DeepTeamEmailContact.md)
 - [EContent](docs/EContent.md)
 - [EMessage](docs/EMessage.md)
 - [Email](docs/Email.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [Link](docs/Link.md)
 - [LinkAddUpdate](docs/LinkAddUpdate.md)
 - [List](docs/List.md)
 - [ListAddUpdate](docs/ListAddUpdate.md)
 - [Tag](docs/Tag.md)
 - [TagAddUpdate](docs/TagAddUpdate.md)
 - [TagContact](docs/TagContact.md)
 - [TagContactId](docs/TagContactId.md)
 - [Team](docs/Team.md)
 - [TeamAddUpdate](docs/TeamAddUpdate.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author


