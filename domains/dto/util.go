package dto

import (
	"github.com/Nexters/pinterest/domains/entities"
)

func ToGroupDtoList(groupList []entities.Group) (groupDtoList []Group) {
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
		groupDtoList = append(groupDtoList, groupInfo)
	}
	return
}

func ToVisitLogDtoList(visitLogList []entities.VisitLog) (visitLogDtoList []VisitLog) {
	for _, visitLog := range visitLogList {
		visitLogInfo := VisitLog{
			UserID: visitLog.UserID,
			Name:   visitLog.Name,
			Text:   visitLog.Text,
		}
		visitLogDtoList = append(visitLogDtoList, visitLogInfo)
	}
	return
}
