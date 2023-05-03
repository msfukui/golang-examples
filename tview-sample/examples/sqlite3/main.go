package main

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	_ "github.com/mattn/go-sqlite3"
)

const (
	batchSize  = 80         // The number of rows loaded patch.
	finderPage = "*finder*" // The name of the Finder Page.
)

var (
	app         *tview.Application // The tview application.
	pages       *tview.Pages       // The application pages.
	finderFocus tview.Primitive    // The primitive in the Finder that last had focus.
)

type Database struct {
	id     int
	dbName string
	dbFile string
}

type Table struct {
	tableName string
}

type Column struct {
	id           int
	columnName   string
	dataType     string
	isNullable   int
	defaultValue sql.NullString
	primaryKey   sql.NullString
}

// Main entry point.
func main() {
	// Get connect string from the command line.
	if len(os.Args) < 2 {
		fmt.Println("Usage: Please provide a SQLite3 database file.")
		return
	}

	// Start the application.
	app = tview.NewApplication()
	finder(os.Args[1])
	if err := app.Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}
}

// Initialize database list, table list and column table.
func initFinder() (*tview.List, *tview.List, *tview.Table) {
	databases := tview.NewList().ShowSecondaryText(false)
	databases.SetBorder(true).SetTitle("Databases")
	columns := tview.NewTable().SetBorders(true)
	columns.SetBorder(true).SetTitle("Columns")
	tables := tview.NewList()
	tables.ShowSecondaryText(false).
		SetDoneFunc(func() {
			tables.Clear()
			columns.Clear()
			app.SetFocus(databases)
		})
	tables.SetBorder(true).SetTitle("Tables")
	return databases, tables, columns
}

// Get database name.
func getDatabaseName(connString string) []Database {
	con, err := sql.Open("sqlite3", connString)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	// get databases.
	rows, err := con.Query("pragma database_list")
	if err != nil {
		panic(err)
	}

	databases := []Database{}
	for rows.Next() {
		var v Database
		if err := rows.Scan(&v.id, &v.dbName, &v.dbFile); err != nil {
			panic(err)
		}
		databases = append(databases, v)
	}

	return databases
}

// Get table name.
func getTableName(connString string, databaseName string) []Table {
	con, err := sql.Open("sqlite3", connString)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	// get tables.
	rows, err := con.Query("select name from sqlite_master where type = 'table'")
	if err != nil {
		panic(err)
	}

	tables := []Table{}
	for rows.Next() {
		var v Table
		if err := rows.Scan(&v.tableName); err != nil {
			panic(err)
		}
		tables = append(tables, v)
	}

	return tables
}

// Get table's column name.
func getColumnName(connString string, databaseName string, tableName string) []Column {
	con, err := sql.Open("sqlite3", connString)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	// get table's column.
	rows, err := con.Query("pragma table_info('" + tableName + "')")
	if err != nil {
		panic(err)
	}

	columns := []Column{}
	for rows.Next() {
		var v Column
		if err := rows.Scan(&v.id, &v.columnName, &v.dataType, &v.isNullable, &v.defaultValue, &v.primaryKey); err != nil {
			panic(err)
		}
		columns = append(columns, v)
	}

	return columns
}

func finder(connString string) {
	// Create the basic objects.
	databases, tables, columns := initFinder()

	dbs := getDatabaseName(connString)

	for _, db := range dbs {
		databases.AddItem(db.dbName, "", 0, func() {
			// A database was selected. Show all of its tables.
			columns.Clear()
			tables.Clear()
			tbs := getTableName(connString, db.dbName)
			for _, table := range tbs {
				tables.AddItem(table.tableName, "", 0, nil)
			}
			app.SetFocus(tables)

			// When the user navigates to a table, show its columns.
			tables.SetChangedFunc(func(i int, tableName string, t string, s rune) {
				// A table was selected. Show its columns.
				columns.Clear()
				cls := getColumnName(connString, db.dbName, tableName)
				columns.SetCell(0, 0, &tview.TableCell{Text: "name", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 1, &tview.TableCell{Text: "type", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 2, &tview.TableCell{Text: "notnull", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 3, &tview.TableCell{Text: "dflt_value", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 4, &tview.TableCell{Text: "pk", Align: tview.AlignCenter, Color: tcell.ColorYellow})
				for _, column := range cls {
					color := tcell.ColorWhite
					columns.SetCell(column.id+1, 0, &tview.TableCell{Text: column.columnName, Color: color}).
						SetCell(column.id+1, 1, &tview.TableCell{Text: column.dataType, Color: color}).
						SetCell(column.id+1, 2, &tview.TableCell{Text: strconv.Itoa(column.isNullable), Align: tview.AlignRight, Color: color}).
						SetCell(column.id+1, 3, &tview.TableCell{Text: column.defaultValue.String, Align: tview.AlignRight, Color: color}).
						SetCell(column.id+1, 4, &tview.TableCell{Text: column.primaryKey.String, Align: tview.AlignLeft, Color: color})
				}
			})
		})

		tables.SetCurrentItem(0) // Trigger the initial selection.
		// When the user selects a table, show its content.
		tables.SetSelectedFunc(func(i int, tableName string, t string, s rune) {
			content(connString, db.dbName, tableName)
		})
	}

	// Create the layout.
	flex := tview.NewFlex().
		AddItem(databases, 0, 1, true).
		AddItem(tables, 0, 1, false).
		AddItem(columns, 0, 3, false)

	// Set up the pages and show the Finder.
	pages = tview.NewPages().
		AddPage(finderPage, flex, true, true)
	app.SetRoot(pages, true)
}

func content(connString string, dbName string, tableName string) {
	finderFocus = app.GetFocus()

	// If this page already exists, just show it.
	if pages.HasPage(dbName + "." + tableName) {
		pages.SwitchToPage(dbName + "." + tableName)
		return
	}

	// We display the data in a table embedded in a frame.
	table := tview.NewTable().
		SetFixed(1, 0).
		SetSeparator(tview.BoxDrawingsLightHorizontal).
		SetBordersColor(tcell.ColorYellow)
	frame := tview.NewFrame(table).
		SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetBorder(true).
		SetTitle(fmt.Sprintf(`Contents of table "%s"`, tableName))

	rowCount := getTableCount(connString, dbName, tableName)

	loadRows := func(offset int) {
		columnNames, columns := getTableContents(connString, dbName, tableName, offset)
		for index, name := range columnNames {
			table.SetCell(0, index, &tview.TableCell{Text: name, Align: tview.AlignCenter, Color: tcell.ColorYellow})
		}
		// Transfer them to the table.
		row := table.GetRowCount()
		for index, column := range columns {
			switch value := column.(type) {
			case int64:
				table.SetCell(row, index, &tview.TableCell{Text: strconv.Itoa(int(value)), Align: tview.AlignRight, Color: tcell.ColorDarkCyan})
			case float64:
				table.SetCell(row, index, &tview.TableCell{Text: strconv.FormatFloat(value, 'f', 2, 64), Align: tview.AlignRight, Color: tcell.ColorDarkCyan})
			case string:
				table.SetCellSimple(row, index, value)
			case time.Time:
				t := value.Format("2006-01-02")
				table.SetCell(row, index, &tview.TableCell{Text: t, Align: tview.AlignRight, Color: tcell.ColorDarkMagenta})
			case []uint8:
				str := make([]byte, len(value))
				for index, num := range value {
					str[index] = byte(num)
				}
				table.SetCell(row, index, &tview.TableCell{Text: string(str), Align: tview.AlignRight, Color: tcell.ColorGreen})
			case nil:
				table.SetCell(row, index, &tview.TableCell{Text: "NULL", Align: tview.AlignCenter, Color: tcell.ColorRed})
			default:
				// We've encountered a type that we don't know yet.
				t := reflect.TypeOf(value)
				str := "?nil?"
				if t != nil {
					str = "?" + t.String() + "?"
				}
				table.SetCellSimple(row, index, str)
			}
		}
		// Show how much we've loaded.
		frame.Clear()
		loadMore := ""
		if table.GetRowCount()-1 < rowCount {
			loadMore = " - press Enter to load more"
		}
		loadMore = fmt.Sprintf("Loaded %d of %d rows%s", table.GetRowCount()-1, rowCount, loadMore)
		frame.AddText(loadMore, false, tview.AlignCenter, tcell.ColorYellow)
	}

	loadRows(0)

	// Handle key presses.
	table.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			// Go back to Finder.
			pages.SwitchToPage(finderPage)
			if finderFocus != nil {
				app.SetFocus(finderFocus)
			}
		case tcell.KeyEnter:
			// Load the next batch of rows.
			loadRows(table.GetRowCount() - 1)
			table.ScrollToEnd()
		}
	})

	// Add a new page and show it.
	pages.AddPage(dbName+"."+tableName, frame, true, true)
}

func getTableCount(connString string, databaseName string, tableName string) int {
	con, err := sql.Open("sqlite3", connString)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	var rowCount int
	err = con.QueryRow("select count(*) from " + tableName).Scan(&rowCount)
	if err != nil {
		panic(err)
	}

	return rowCount
}

// Load a batch of rows.
func getTableContents(connString string, databaseName string, tableName string, offset int) ([]string, []interface{}) {
	con, err := sql.Open("sqlite3", connString)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	rows, err := con.Query("select * from "+tableName+" limit ? offset ?", batchSize, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	columnNames, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	// Read the rows.
	columns := make([]interface{}, len(columnNames))
	columnPointers := make([]interface{}, len(columns))
	for index := range columnPointers {
		columnPointers[index] = &columns[index]
	}

	for rows.Next() {
		// Read the columns.
		err := rows.Scan(columnPointers...)
		if err != nil {
			panic(err)
		}
	}

	return columnNames, columns
}
