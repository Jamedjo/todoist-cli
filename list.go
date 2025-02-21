package main

import (
	"fmt"
	"os"

	"github.com/acarl005/stripansi"
	todoist "github.com/sachaos/todoist/lib"
	"github.com/urfave/cli/v2"
)

func traverseItems(item *todoist.Item, f func(item *todoist.Item, depth int) error, depth int) error {
	err := f(item, depth)
	if err != nil {
		return err
	}

	if item.ChildItem != nil {
		err := traverseItems(item.ChildItem, f, depth+1)
		if err != nil {
			return err
		}
	}

	if item.BrotherItem != nil {
		err := traverseItems(item.BrotherItem, f, depth)
		if err != nil {
			return err
		}
	}

	return nil
}

func sortItems(itemListPtr *[][]string, byIndex int) {
	itemList := *itemListPtr
	length := len(itemList)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if stripansi.Strip(itemList[j][byIndex]) > stripansi.Strip(itemList[j+1][byIndex]) {
				tmp := itemList[j]
				itemList[j] = itemList[j+1]
				itemList[j+1] = tmp
			}
		}
	}
}

func List(c *cli.Context) error {
	client := GetClient(c)

	colorList := ColorList()
	projectsCount := len(client.Store.Projects)
	projectIds := make([]string, projectsCount)
	for i, project := range client.Store.Projects {
		projectIds[i] = project.GetID()
	}
	projectColorHash := GenerateColorHash(projectIds, colorList)

	rootItem := client.Store.RootItem

	if rootItem == nil {
		fmt.Fprintln(os.Stderr, "There is no task. You can fetch latest tasks by `todoist sync`.")
		return nil
	}

	for _, ex := range Filter(c.String("filter")) {
		itemList := [][]string{}
		err := traverseItems(rootItem, func(item *todoist.Item, depth int) error {
			r, err := Eval(ex, item, client.Store)
			if err != nil {
				return err
			}
			if !r || item.Checked {
				return nil
			}
			itemList = append(itemList, []string{
				IdFormat(item),
				PriorityFormat(item.Priority),
				DueDateFormat(item.DateTime(), item.AllDay),
				ProjectFormat(item.ProjectID, client.Store, projectColorHash, c) +
					SectionFormat(item.SectionID, client.Store, c),
				item.LabelsString(client.Store),
				ContentPrefix(client.Store, item, depth, c) + ContentFormat(item),
			})

			return nil
		}, 0)
		if err != nil {
			return err
		}

		if c.Bool("priority") == true {
			// sort output by priority
			// and no need to use "else block" as items returned by API are already sorted by task id
			sortItems(&itemList, 1)
		}

		defer writer.Flush()

		if c.Bool("header") {
			writer.Write([]string{"ID", "Priority", "DueDate", "Project", "Labels", "Content"})
		}

		for _, strings := range itemList {
			writer.Write(strings)
		}
	}
	return nil
}
