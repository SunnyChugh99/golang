package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	originalData := `
				{
					"boardInfo": {
						"_id": "63c0f9c322c1f3f0b0e71e7d",
						"state": "ACTIVE",
						"priority": "HIGH",
						"postStartDate": "",
						"postEndDate": "",
						"postSchedule": "",
						"proprietaryRegisterDate": "",
						"story": {
							"publicStartDate": "0001-01-01T00:00:00Z",
							"publicEndDate": "0001-01-01T00:00:00Z"
						},
						"ppv": 0,
						"visible": "",
						"sharable": false,
						"searchable": "PUBLIC",
						"saving": 0,
						"reaction": false,
						"viewCount": 0,
						"ownerInfo": {
							"id": 58,
							"firstName": "Yash",
							"lastName": "Shah",
							"photo": "",
							"screenName": "58",
							"userName": ""
						},
						"taggedPeople": null,
						"lockedDate": "",
						"parentID": "",
						"display": null,
						"coverImage": "",
						"chat": false,
						"isComments": false,
						"title": "New Board",
						"type": "BOARD",
						"description": "my description",
						"tags": [
							"a",
							"b",
							"c"
						],
						"owner": "58",
						"admins": [
							"51"
						],
						"authors": null,
						"viewers": null,
						"subscribers": null,
						"blocked": null,
						"sort": 0,
						"comments": null,
						"createDate": "2023-01-13T06:27:15.113Z",
						"modifiedDate": "2023-01-13T06:27:15.113Z",
						"sortDate": "",
						"deletedDate": "0001-01-01T00:00:00Z",
						"expiryDate": "0001-01-01T00:00:00Z",
						"isPassword": true,
						"location": "New Board"
					},
					"things": [
						{
							"_id": "63d7acf8a9e7effe9f788ccb",
							"boardID": "63c0f9c322c1f3f0b0e71e7d",
							"title": "Task for board 6 P1",
							"type": "TASK",
							"description": "This is the task note. This is the task note. This is the task note. This is the task note. This is the task note. This is the task note.",
							"tags": null,
							"owner": "58",
							"state": "ACTIVE",
							"priority": "high",
							"postStartDate": "",
							"proprietaryRegisterDate": "",
							"viewCount": 0,
							"createDate": "2023-01-30T11:41:44.995Z",
							"modifiedDate": "2023-01-30T11:41:44.995Z",
							"sortDate": "",
							"deletedDate": "0001-01-01T00:00:00Z",
							"expiryDate": "0001-01-01T00:00:00Z",
							"sort": 0,
							"story": {
								"publicStartDate": "0001-01-01T00:00:00Z",
								"publicEndDate": "0001-01-01T00:00:00Z"
							},
							"ppv": 0,
							"visible": "PUBLIC",
							"sharable": false,
							"searchable": "PUBLIC",
							"saving": 0,
							"reaction": false,
							"comments": null,
							"thingLocation": "",
							"ownerInfo": {
								"id": 58,
								"firstName": "Yash",
								"lastName": "Shah",
								"photo": "",
								"screenName": "58",
								"userName": ""
							},
							"assignedMemberInfo": null,
							"postEndDate": "",
							"postSchedule": "",
							"taggedPeople": null,
							"lockedDate": "",
							"taskPriority": "",
							"assignedToID": "",
							"taskStatus": "open",
							"reminderTimer": "0001-01-01T00:00:00Z",
							"dueDate": "2023-02-02",
							"dueTime": "21:14",
							"completeDate": "0001-01-01T00:00:00Z",
							"archiveDate": "0001-01-01T00:00:00Z"
						},
						{
							"_id": "63d7a704a9e7effe9f788c3d",
							"boardID": "63c0f9c322c1f3f0b0e71e7d",
							"title": "Task for board 6",
							"type": "TASK",
							"description": "Required data",
							"tags": null,
							"owner": "58",
							"state": "ACTIVE",
							"priority": "",
							"postStartDate": "",
							"proprietaryRegisterDate": "",
							"viewCount": 0,
							"createDate": "2023-01-30T11:16:20.377Z",
							"modifiedDate": "2023-01-30T11:16:20.377Z",
							"sortDate": "",
							"deletedDate": "0001-01-01T00:00:00Z",
							"expiryDate": "0001-01-01T00:00:00Z",
							"sort": 0,
							"story": {
								"publicStartDate": "0001-01-01T00:00:00Z",
								"publicEndDate": "0001-01-01T00:00:00Z"
							},
							"ppv": 0,
							"visible": "PUBLIC",
							"sharable": false,
							"searchable": "PUBLIC",
							"saving": 0,
							"reaction": false,
							"comments": null,
							"thingLocation": "",
							"ownerInfo": {
								"id": 58,
								"firstName": "Yash",
								"lastName": "Shah",
								"photo": "",
								"screenName": "58",
								"userName": ""
							},
							"assignedMemberInfo": null,
							"postEndDate": "",
							"postSchedule": "",
							"taggedPeople": null,
							"lockedDate": "",
							"taskPriority": "",
							"assignedToID": "",
							"taskStatus": "",
							"reminderTimer": "0001-01-01T00:00:00Z",
							"dueDate": "",
							"dueTime": "",
							"completeDate": "0001-01-01T00:00:00Z",
							"archiveDate": "0001-01-01T00:00:00Z"
						},
						{
							"_id": "63d7a6f6a9e7effe9f788c3a",
							"boardID": "63c0f9c322c1f3f0b0e71e7d",
							"title": "Task for board 6",
							"type": "TASK",
							"description": "Required data",
							"tags": null,
							"owner": "58",
							"state": "ACTIVE",
							"priority": "",
							"postStartDate": "",
							"proprietaryRegisterDate": "",
							"viewCount": 0,
							"createDate": "2023-01-30T11:16:06.052Z",
							"modifiedDate": "2023-01-30T11:16:06.052Z",
							"sortDate": "",
							"deletedDate": "0001-01-01T00:00:00Z",
							"expiryDate": "0001-01-01T00:00:00Z",
							"sort": 0,
							"story": {
								"publicStartDate": "0001-01-01T00:00:00Z",
								"publicEndDate": "0001-01-01T00:00:00Z"
							},
							"ppv": 0,
							"visible": "PUBLIC",
							"sharable": false,
							"searchable": "PUBLIC",
							"saving": 0,
							"reaction": false,
							"comments": null,
							"thingLocation": "",
							"ownerInfo": {
								"id": 58,
								"firstName": "Yash",
								"lastName": "Shah",
								"photo": "",
								"screenName": "58",
								"userName": ""
							},
							"assignedMemberInfo": null,
							"postEndDate": "",
							"postSchedule": "",
							"taggedPeople": null,
							"lockedDate": "",
							"taskPriority": "",
							"assignedToID": "",
							"taskStatus": "",
							"reminderTimer": "0001-01-01T00:00:00Z",
							"dueDate": "",
							"dueTime": "",
							"completeDate": "0001-01-01T00:00:00Z",
							"archiveDate": "0001-01-01T00:00:00Z"
						},
						{
							"_id": "63bc21d1fbaa99caf77b846e",
							"boardID": "63c0f9c322c1f3f0b0e71e7d",
							"title": "task title",
							"type": "TASK",
							"description": "Just a task description.",
							"tags": [
								"task1",
								"task2",
								"strings"
							],
							"owner": "13",
							"state": "ACTIVE",
							"priority": "",
							"postStartDate": "",
							"proprietaryRegisterDate": "",
							"viewCount": 0,
							"createDate": "2023-01-09T14:16:49.004Z",
							"modifiedDate": "2023-01-09T14:16:49.004Z",
							"sortDate": "",
							"deletedDate": "0001-01-01T00:00:00Z",
							"expiryDate": "0001-01-01T00:00:00Z",
							"sort": 0,
							"story": {
								"publicStartDate": "0001-01-01T00:00:00Z",
								"publicEndDate": "0001-01-01T00:00:00Z"
							},
							"ppv": 0,
							"visible": "PUBLIC",
							"sharable": false,
							"searchable": "PUBLIC",
							"saving": 0,
							"reaction": false,
							"comments": null,
							"thingLocation": "",
							"ownerInfo": {
								"id": 13,
								"firstName": "XYZ",
								"lastName": "ABC",
								"photo": "",
								"screenName": "Beta",
								"userName": ""
							},
							"assignedMemberInfo": null,
							"postEndDate": "",
							"postSchedule": "",
							"taggedPeople": null,
							"lockedDate": "",
							"taskPriority": "HIGH",
							"assignedToID": "14",
							"taskStatus": "",
							"reminderTimer": "0001-01-01T00:00:00Z",
							"dueDate": "2023-05-05",
							"dueTime": "",
							"completeDate": "0001-01-01T00:00:00Z",
							"archiveDate": "0001-01-01T00:00:00Z"
						},
						{
							"_id": "63c91357d60f9738f59ac79a",
							"state": "ACTIVE",
							"priority": "",
							"postStartDate": "",
							"postEndDate": "",
							"postSchedule": "",
							"proprietaryRegisterDate": "",
							"story": {
								"publicStartDate": "0001-01-01T00:00:00Z",
								"publicEndDate": "0001-01-01T00:00:00Z"
							},
							"ppv": 0,
							"visible": "PUBLIC",
							"sharable": false,
							"searchable": "PUBLIC",
							"saving": 0,
							"reaction": false,
							"viewCount": 0,
							"ownerInfo": {
								"id": 58,
								"firstName": "Yash",
								"lastName": "Shah",
								"photo": "",
								"screenName": "58",
								"userName": ""
							},
							"taggedPeople": null,
							"lockedDate": "",
							"parentID": "63c0f9c322c1f3f0b0e71e7d",
							"display": null,
							"coverImage": "",
							"chat": false,
							"isComments": false,
							"title": "Child board 1",
							"type": "BOARD",
							"description": "",
							"tags": [
								"a",
								"b",
								"c",
								"d",
								"e",
								"f",
								"g"
							],
							"owner": "58",
							"admins": null,
							"authors": null,
							"viewers": null,
							"subscribers": null,
							"blocked": null,
							"sort": 0,
							"comments": null,
							"createDate": "2023-01-19T09:54:31.29Z",
							"modifiedDate": "2023-01-19T09:54:31.29Z",
							"sortDate": "",
							"deletedDate": "0001-01-01T00:00:00Z",
							"expiryDate": "0001-01-01T00:00:00Z",
							"isPassword": false,
							"location": ""
						}
					]
				}
				`
	newData := `
				{
						"_id": "63c91357d60f9738f59ac79a",
						"state": "ACTIVE",
						"priority": "",
						"postStartDate": "",
						"postEndDate": "",
						"postSchedule": "",
						"proprietaryRegisterDate": "",
						"story": {
							"publicStartDate": "0001-01-01T00:00:00Z",
							"publicEndDate": "0001-01-01T00:00:00Z"
						},
						"ppv": 0,
						"visible": "PUBLIC",
						"sharable": false,
						"searchable": "PUBLIC",
						"saving": 0,
						"reaction": false,
						"viewCount": 0,
						"ownerInfo": {
							"id": 58,
							"firstName": "Yash",
							"lastName": "Shah",
							"photo": "",
							"screenName": "58",
							"userName": ""
						},
						"taggedPeople": null,
						"lockedDate": "",
						"parentID": "63c0f9c322c1f3f0b0e71e7d",
						"display": null,
						"coverImage": "",
						"chat": false,
						"isComments": false,
						"title": "Child board 1",
						"type": "BOARD",
						"description": "",
						"tags": [
							"a",
							"b",
							"c",
							"d",
							"e",
							"f",
							"g"
						],
						"owner": "58",
						"admins": null,
						"authors": null,
						"viewers": null,
						"subscribers": null,
						"blocked": null,
						"sort": 0,
						"comments": null,
						"createDate": "2023-01-19T09:54:31.29Z",
						"modifiedDate": "2023-01-19T09:54:31.29Z",
						"sortDate": "",
						"deletedDate": "0001-01-01T00:00:00Z",
						"expiryDate": "0001-01-01T00:00:00Z",
						"isPassword": false,
						"location": ""
				}`


	/*
		Objective 1: Check the length of "things"
		Objective 2: Append the "newData" in "things" of originalData
		Objective 3: Verify by checking the length of "things"

	*/
	var dataMap map[string]interface{}
	if err := json.Unmarshal([]byte(originalData), &dataMap); err != nil {
		fmt.Println("Error unmarshalling originalData:", err)
		return
	}

	things := dataMap["things"].([]interface{})
	// if !ok {
	// 	fmt.Println("Invalid data format: 'things' is not an array")
	// 	return
	// }

	fmt.Println(newData)
	fmt.Printf("v is of type %T\n", things)

	fmt.Println("Length of 'things' before appending newData:", len(things))
	fmt.Printf("v is of type %T\n", things)

	fmt.Println("type of things", reflect.TypeOf(things))

	newDataMap := make(map[string]interface{})
	
	err := json.Unmarshal([]byte(newData), &newDataMap)

	if err != nil{
		fmt.Println("problem in parsing new data")
	}
	fmt.Println("new data print")
	fmt.Println(len(newDataMap))

	things = append(things, newDataMap)	
	fmt.Println(len(things))
	dataMap["things"] = things

	updatedData, err := json.Marshal(dataMap)
	fmt.Println("after update")
	fmt.Println(string(updatedData))

}
