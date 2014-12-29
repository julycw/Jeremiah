package note

import (
	"github.com/julycw/Jeremiah/controllers/manage"
	"github.com/julycw/Jeremiah/models"
	"log"
)

type DayNoteList struct {
	Day      string
	NoteList []models.Note
}

type NoteController struct {
	manage.BaseController
}

func (this *NoteController) Get() {
	this.CheckAvaliable("使用记录板")

	this.TplNames = "manage/note/index.tpl"
	this.LayoutSections["LayoutFooter"] = "manage/note/script.tpl"

	var noteList []models.Note = make([]models.Note, 0)

	if _, err := models.Orm.QueryTable("note").OrderBy("-create_on").Limit(50).All(&noteList); err != nil {
		log.Println(err.Error())
	} else {
		var dayNoteLists []DayNoteList = make([]DayNoteList, 0)
		for _, note := range noteList {
			dayStr := note.CreateOn.Format("01月02日")
			hasThisDay := false
			for i, day := range dayNoteLists {
				if dayStr == day.Day {
					dayNoteLists[i].NoteList = append(dayNoteLists[i].NoteList, note)
					hasThisDay = true
				}
			}

			if !hasThisDay {
				var dayNoteList DayNoteList
				dayNoteList.Day = dayStr
				dayNoteList.NoteList = make([]models.Note, 0)
				dayNoteList.NoteList = append(dayNoteList.NoteList, note)
				dayNoteLists = append(dayNoteLists, dayNoteList)
			}
		}
		this.Data["dayNoteLists"] = &dayNoteLists
	}
}

func (this *NoteController) Post() {
	this.Redirect("/manage/note", 302)
}

func (this *NoteController) Add() {
	this.CheckAvaliable("使用记录板")
	content := this.GetString("content")
	tags := this.GetString("tags")
	var note models.Note
	note.Content = content
	note.UserName = this.User.UserName
	note.Tags = tags
	_, err := models.Orm.Insert(&note)
	if err != nil {
		log.Println(err.Error())
	}
	this.Redirect("/manage/note", 302)
}
