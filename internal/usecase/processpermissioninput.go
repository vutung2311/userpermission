package usecase

import (
	"strconv"
	"strings"

	"userpermission/internal/model"
)

func ProcessPermissionInput(input []string) ([]string, error) {
	n, err := strconv.Atoi(input[0])
	if err != nil {
		return nil, err
	}
	userCount := n + 1
	userPermissionStorage := make([]*model.UserPermission, userCount)

	for i, j := 0, 1; i < userCount; i, j = i+1, j+1 {
		userPermissionStorage[i] = new(model.UserPermission)
		userPermissionStorage[i].AddPermissions(strings.Split(input[j], " ")...)
	}

	for i, j := 1, userCount+1; i < userCount; i, j = i+1, j+1 {
		if input[j] == "CEO" {
			userPermissionStorage[i].SetParent(userPermissionStorage[0])
			continue
		}
		index, err := strconv.Atoi(input[j])
		if err != nil {
			return nil, err
		}
		userPermissionStorage[i].SetParent(userPermissionStorage[index])
	}

	output := make([]string, 0, userCount)
	for i := range userPermissionStorage {
		output = append(output, strings.Join(userPermissionStorage[i].GetPermissions(), ", "))
	}

	for i := 2 * userCount; i < len(input); i++ {
		cmd := strings.Split(input[i], " ")
		userNumber, err := strconv.Atoi(cmd[1])
		if err != nil && cmd[1] != "CEO" {
			return nil, err
		}
		if cmd[1] == "CEO" {
			userNumber = 0
		}

		switch cmd[0] {
		case "ADD":
			userPermissionStorage[userNumber].AddPermission(cmd[2])
		case "REMOVE":
			userPermissionStorage[userNumber].RemovePermission(cmd[2])
		case "QUERY":
			output = append(output, strings.Join(userPermissionStorage[userNumber].GetPermissions(), ", "))
		default:
			break
		}
	}

	return output, nil
}
