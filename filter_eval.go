package main

import (
	"regexp"
	"strconv"
	"time"

	todoist "github.com/sachaos/todoist/lib"
)

var priorityRegex = regexp.MustCompile("^p([1-4])$")

func Eval(e Expression, item todoist.AbstractItem, projects todoist.Projects) (result bool, err error) {
	result = false
	switch e.(type) {
	case BoolInfixOpExpr:
		e := e.(BoolInfixOpExpr)
		lr, err := Eval(e.left, item, projects)
		rr, err := Eval(e.right, item, projects)
		if err != nil {
			return false, nil
		}
		switch e.operator {
		case '&':
			return lr && rr, nil
		case '|':
			return lr || rr, nil
		}
	case AssignedExpr:
		// todo tidy up pointer casting
	  return item.(*todoist.Item).ResponsibleUID != nil, nil
	case NoPriorityExpr:
		// ditto
		return item.(*todoist.Item).Priority == 1, nil
	case RecurringExpr:
		due := item.(*todoist.Item).Due
		if due == nil {
			return false, nil
		}
		return due.IsRecurring, nil
	case SubtaskExpr:
		parentId := item.(*todoist.Item).HaveParentID.ParentID
		return parentId != nil, nil
	case WeekdayExpr:
		due := item.DateTime()
		if (due == time.Time{}) {
			return false, nil
		}
		return due.Weekday() == e.(WeekdayExpr).day, nil
	// case SearchExpr:
	// case PersonExpr:
	// case ListExpr:
	// aaaa, I repeat, perhaps the parser should return collection
	// this makes no sense
	case ProjectExpr:
		e := e.(ProjectExpr)
		return EvalProject(e, item.GetProjectID(), projects), err
	case LabelExpr:
		e := e.(LabelExpr)
		return EvalLabel(e, item.GetLabelNames()), err
	case StringExpr:
		switch item.(type) {
		case *todoist.Item:
			item := item.(*todoist.Item)
			e := e.(StringExpr)
			return EvalAsPriority(e, item), err
		default:
			return false, nil
		}
	case DateExpr:
		e := e.(DateExpr)
		return EvalDate(e, item), err
	case ViewAllExpr:
		// we could fall through to the default, but let's not
		return true, nil
	case NotOpExpr:
		e := e.(NotOpExpr)
		r, err := Eval(e.expr, item, projects)
		if err != nil {
			return false, nil
		}
		return !r, nil
	default:
		return true, err
	}
	return
}

func EvalDate(e DateExpr, item todoist.AbstractItem) (result bool) {
	itemDate := item.DateTime()
	if e.operation == NO_TIME {
		i := item.(*todoist.Item)
		if i.Due == nil {
			// no date is also no time
			return true
		} else {
			_, err := time.Parse(time.DateOnly, i.Due.Date)
			return err == nil // only a date
		}
	}

	if (itemDate == time.Time{}) {
		return e.operation == NO_DUE_DATE
	}
	allDay := e.allDay
	dueDate := e.datetime
	switch e.operation {
	case DUE_ON:
		var startDate, endDate time.Time
		if allDay {
			startDate = dueDate
			endDate = dueDate.AddDate(0, 0, 1)
			if itemDate.Equal(startDate) || (itemDate.After(startDate) && itemDate.Before(endDate)) {
				return true
			}
		}
		return false
	case DUE_BEFORE:
		return itemDate.Before(dueDate)
	case DUE_AFTER:
		endDateTime := dueDate
		if allDay {
			endDateTime = dueDate.AddDate(0, 0, 1).Add(-time.Duration(time.Microsecond))
		}
		return itemDate.After(endDateTime)
	default:
		return false
	}
}

// todo maybe priority should have it's own expr
func EvalAsPriority(e StringExpr, item *todoist.Item) (result bool) {
	matched := priorityRegex.FindStringSubmatch(e.String())
	if len(matched) == 0 {
		return false
	} else {
		p, _ := strconv.Atoi(matched[1])
		if p == priorityMapping[item.Priority] {
			return true
		}
	}
	return false
}

func EvalProject(e ProjectExpr, projectID string, projects todoist.Projects) bool {
	for _, id := range projects.GetIDsByName(e.name, e.isAll) {
		if id == projectID {
			return true
		}
	}
	return false
}

func EvalLabel(e LabelExpr, labelNames []string) bool {
	if e.name == "" {
		return len(labelNames) == 0
	}

	for _, name := range labelNames {
		if name == e.name {
			return true
		}
	}

	return false
}
