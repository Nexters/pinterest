package dto

import (
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/go-playground/validator"
)

func ToGroupDtoList(groupList []entities.Group) (groupDtoList []Group, err error) {
	for _, group := range groupList {
		groupInfo := Group{
			Type:      group.Type,
			Title:     group.Title,
			Text:      group.Text,
			Image:     group.Image,
			Order:     group.Order,
			ItemCount: group.ItemCount,
			Likes:     group.Likes,
			Link:      group.Link,
			UserID:    group.UserID,
		}
		validate := validator.New()
		err := validate.Struct(groupInfo)
		if err != nil {
			return groupDtoList, err
		}
		groupDtoList = append(groupDtoList, groupInfo)
	}
	return
}

func ToVisitLogDtoList(visitLogList []entities.VisitLog) (visitLogDtoList []VisitLog, err error) {
	for _, visitLog := range visitLogList {
		visitLogInfo := VisitLog{
			UserID: visitLog.UserID,
			Name:   visitLog.Name,
			Text:   visitLog.Text,
		}
		validate := validator.New()
		err := validate.Struct(visitLogInfo)
		if err != nil {
			return visitLogDtoList, err
		}
		visitLogDtoList = append(visitLogDtoList, visitLogInfo)
	}
	return
}
