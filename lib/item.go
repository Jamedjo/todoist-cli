package todoist

import (
	"context"
	"regexp"
	"strings"
	"time"
)

var linkRegex = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

var PriorityMapping = map[int]int{
	1: 4,
	2: 3,
	3: 2,
	4: 1,
}

const (
	RFC3339Date                 = "2006-01-02"
	RFC3339DateTime             = "2006-01-02T15:04:05"
	RFC3339DateTimeWithTimeZone = "2006-01-02T15:04:05Z07:00"
)

type Due struct {
	Date        string `json:"date"`
	TimeZone    string `json:"timezone"`
	IsRecurring bool   `json:"is_recurring"`
	String      string `json:"string"`
	Lang        string `json:"lang"`
}

type BaseItem struct {
	HaveID
	HaveProjectID
	Content     string `json:"content"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

func (bitem BaseItem) GetContent() string {
	return bitem.Content
}

func (bitem BaseItem) GetDescription() string {
	return bitem.Description
}

type CompletedItem struct {
	BaseItem
	CompletedData string      `json:"completed_at"`
	MetaData      interface{} `json:"meta_data"`
	TaskID        string      `json:"task_id"`
}

func (item CompletedItem) DateTime() time.Time {
	t, _ := time.Parse(time.RFC3339, item.CompletedData)
	return t
}

func (item CompletedItem) CreatedTime() time.Time {
	return time.Time{}
}

func (item CompletedItem) GetProjectID() string {
	return item.ProjectID
}

func (item CompletedItem) GetLabelNames() []string {
	return []string{}
}

func (item CompletedItem) GetAssignedByUID() string {
	return ""
}

func (item CompletedItem) GetResponsibleUID() string {
	return ""
}

func (item CompletedItem) GetUserID() string {
	return item.UserID
}

func (item CompletedItem) HasParent() bool {
	return false
}

func (item CompletedItem) HasPriority() bool {
	return false
}

func (item CompletedItem) HasTime() bool {
	return false
}

func (item CompletedItem) IsRecurring() bool {
	return false
}

func (item BaseItem) SearchMatches(re string) bool {
	match, _ := regexp.MatchString(strings.ToLower(re), strings.ToLower(item.GetContent()))
	return match
}

func (item CompletedItem) SearchMatches(re string) bool {
	return item.BaseItem.SearchMatches(re)
}

type CompletedItems []CompletedItem

type Item struct {
	BaseItem
	HaveParentID
	HaveIndent
	HaveSectionID
	ChildItem      *Item       `json:"-"`
	BrotherItem    *Item       `json:"-"`
	AllDay         bool        `json:"all_day"`
	AssignedByUID  string      `json:"assigned_by_uid"`
	Checked        bool        `json:"checked"`
	Collapsed      bool        `json:"collapsed"`
	CompletedAt    string      `json:"completed_at"`
	DateAdded      string      `json:"added_at"`
	DateLang       string      `json:"date_lang"`
	DateString     string      `json:"date_string"`
	DayOrder       int         `json:"day_order"`
	Due            *Due        `json:"due"`
	HasMoreNotes   bool        `json:"has_more_notes"`
	IsArchived     int         `json:"is_archived"`
	IsDeleted      bool        `json:"is_deleted"`
	ChildOrder     int         `json:"child_order"`
	LabelNames     []string    `json:"labels"`
	Priority       int         `json:"priority"`
	AutoReminder   bool        `json:"auto_reminder"`
	ResponsibleUID interface{} `json:"responsible_uid"`
	SyncID         interface{} `json:"sync_id"`
}

type Items []Item

func (a Items) Len() int           { return len(a) }
func (a Items) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Items) Less(i, j int) bool { return a[i].ID < a[j].ID }

func (a Items) At(i int) IDCarrier { return a[i] }

func (item Item) DateTime() time.Time {
	var date string
	// TODO: it would be more correct to get the timezone from Store.User
	location := time.Local

	if item.Due == nil {
		date = ""
	} else {
		if item.Due.TimeZone != "" {
			dueTz, err := time.LoadLocation(item.Due.TimeZone)
			if err == nil {
				location = dueTz
			}
		}
		date = item.Due.Date
	}
	// 2020-03-03T14:00:00
	// 2020-01-17T23:00:00Z
	t, err := time.ParseInLocation(RFC3339DateTimeWithTimeZone, date, location)
	if err != nil {
		t, err = time.ParseInLocation(RFC3339DateTime, date, location)
	}
	if err != nil {
		t, _ = time.ParseInLocation(RFC3339Date, date, location)
	}
	return t
}

func (item Item) GetProjectID() string {
	return item.ProjectID
}

func (item Item) GetLabelNames() []string {
	return item.LabelNames
}

func (item Item) CreatedTime() time.Time {
	t, _ := time.Parse(time.RFC3339, item.DateAdded)
	return t
}

func (item Item) HasTime() bool {
	if item.Due == nil {
		// no time doesn't apply if no due date
		return false
	} else {
		_, err := time.Parse(time.DateOnly, item.Due.Date)
		return err == nil // only a date
	}
}

func (item Item) SearchMatches(re string) bool {
	return item.BaseItem.SearchMatches(re)
}

func (item Item) GetResponsibleUID() string {
	id, ok := item.ResponsibleUID.(string)
	if !ok {
		return ""
	}
	return id
}

func (item Item) HasPriority() bool {
	return item.Priority == 1
}

func (item Item) IsRecurring() bool {
	due := item.Due
	if due == nil {
		return false
	}
	return due.IsRecurring
}

func (item Item) GetContent() string {
	return item.Content
}

func (item Item) HasParent() bool {
	if item.HaveParentID.ParentID == nil {
		return false
	}
	return *item.HaveParentID.ParentID != ""
}

func (item Item) GetAssignedByUID() string {
	return item.AssignedByUID
}

func (item Item) GetUserID() string {
	return item.UserID
}

func (item Item) Removable() bool {
	return item.IsDeleted || item.CompletedAt != ""
}

// interface for Eval actions
type AbstractItem interface {
	CreatedTime() time.Time
	DateTime() time.Time
	GetProjectID() string
	GetLabelNames() []string
	HasTime() bool
	GetResponsibleUID() string
	SearchMatches(string) bool
	HasPriority() bool
	IsRecurring() bool
	GetContent() string
	HasParent() bool
	GetAssignedByUID() string
	GetUserID() string
}

func GetContentTitle(item ContentCarrier) string {
	return linkRegex.ReplaceAllString(item.GetContent(), "$1")
}

func GetDescription(item DescriptionCarrier) string {
	return linkRegex.ReplaceAllString(item.GetDescription(), "$1")
}

func GetContentURL(item ContentCarrier) []string {
	if HasURL(item) {
		matches := linkRegex.FindAllStringSubmatch(item.GetContent(), -1)
		if matches != nil {
			urls := make([]string, len(matches))
			for i, match := range matches {
				urls[i] = match[2]
			}
			return urls
		}
	}
	return []string{}
}

func HasURL(item interface{}) bool {
	switch item.(type) {
	case ContentCarrier:
		return linkRegex.MatchString(item.(ContentCarrier).GetContent())
	case DescriptionCarrier:
		return linkRegex.MatchString(item.(DescriptionCarrier).GetDescription())
	default:
		return false
	}
}

func (item Item) AddParam() interface{} {
	param := map[string]interface{}{}
	if item.Content != "" {
		param["content"] = item.Content
	}
	if item.Description != "" {
		param["description"] = item.Description
	}
	if item.DateString != "" {
		param["date_string"] = item.DateString
	}
	if len(item.LabelNames) != 0 {
		param["labels"] = item.LabelNames
	}
	if item.Priority != 0 {
		param["priority"] = item.Priority
	}
	if item.ProjectID != "" {
		param["project_id"] = item.ProjectID
	}
	if item.Due != nil {
		param["due"] = item.Due
	}
	param["auto_reminder"] = item.AutoReminder

	return param
}

func (item Item) UpdateParam() interface{} {
	param := map[string]interface{}{}
	if item.ID != "" {
		param["id"] = item.ID
	}
	if item.Content != "" {
		param["content"] = item.Content
	}
	if item.Description != "" {
		param["description"] = item.Description
	}
	if item.DateString != "" {
		param["date_string"] = item.DateString
	}
	// TODO: more cool
	if item.DateString == "null" {
		param["date_string"] = ""
	}
	if len(item.LabelNames) != 0 {
		param["labels"] = item.LabelNames
	}
	if item.Priority != 0 {
		param["priority"] = item.Priority
	}
	if item.Due != nil {
		param["due"] = item.Due
	}
	return param
}

func (item *Item) MoveParam(projectId string) interface{} {
	param := map[string]interface{}{
		"id":         item.ID,
		"project_id": projectId,
	}
	return param
}

func (item Item) LabelsString(store *Store) string {
	var b strings.Builder
	labelIDs := []string{}
	for _, labelName := range item.LabelNames {
		labelIDs = append(labelIDs, store.Labels.GetIDByName(labelName))
	}
	for i, labelId := range labelIDs {
		label := store.FindLabel(labelId)
		b.WriteString("@" + label.Name)
		if i < len(labelIDs)-1 {
			b.WriteString(",")
		}
	}
	return b.String()
}

func (c *Client) AddItem(ctx context.Context, item Item) error {
	commands := Commands{
		NewCommand("item_add", item.AddParam()),
	}
	_, err := c.ExecCommands(ctx, commands)
	return err
}

func (c *Client) UpdateItem(ctx context.Context, item Item) error {
	commands := Commands{
		NewCommand("item_update", item.UpdateParam()),
	}
	_, err := c.ExecCommands(ctx, commands)
	return err
}

func (c *Client) CloseItem(ctx context.Context, ids []string) error {
	var commands Commands
	for _, id := range ids {
		command := NewCommand("item_close", map[string]interface{}{"id": id})
		commands = append(commands, command)
	}
	_, err := c.ExecCommands(ctx, commands)
	return err
}

func (c *Client) DeleteItem(ctx context.Context, ids []string) error {
	var commands Commands
	for _, id := range ids {
		command := NewCommand("item_delete", map[string]interface{}{"id": id})
		commands = append(commands, command)
	}
	_, err := c.ExecCommands(ctx, commands)
	return err
}

func (c *Client) MoveItem(ctx context.Context, item *Item, projectId string) error {
	commands := Commands{
		NewCommand("item_move", item.MoveParam(projectId)),
	}
	_, err := c.ExecCommands(ctx, commands)
	return err
}
