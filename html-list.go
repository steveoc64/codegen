package main

import (
	"fmt"
	"log"
	"os"
)

func generateHTML_list(listFile string) {

	f, err := os.Create(listFile)
	if err != nil {
		log.Fatalln("CreateFile:", listFile, err.Error())
	}
	defer f.Close()

	var str = fmt.Sprintf(`<lx-dialog class="dialog dialog__scrollable dialog--l bgc-light-gradient" id='%sLogDialog' escape-close="true">
    <div class="dialog__header">
        <div class="toolbar bgc-indigo-800 pl++">
            <span class="toolbar__label tc-white fs-title">
                Recent %ss Activity
            </span>       
            <span lx-dialog-close class="white">X</span>     
        </div>
    </div>

    <div class="data-table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th> Event</th>
            <th> Date</th>
            <th> Description</th>
            <th> IP Address</th>
            <th> By</th>
          </tr>
        </thead>
        <tbody>
          <tr class="data-table"
              ng-class="%ss.logClass(l)"        
              ng-repeat="l in %ss.getSelectedLogs()">
                <td>{{l.Type}}</td>
                <td>{{l.Logdate}}</td>
                <td>{{l.Descr}}</td>
                <td>{{l.IP}}</td>
                <td>{{l.Username}}</td>
          </tr>
        </tbody>
      </table>
    </div>
</lx-dialog>

<div class="data-table-container">
  <header class="data-table-header">
    <div class="data-table-header__label">
      <i class="icon icon--xs icon--blue-grey icon--circled mdi mdi-auto-fix"></i>   
      <span class="fs-title">%ss List</span>
    </div>

    <div class="data-table-header__actions">
      <lx-dropdown position="right" over-toggle="true">
        <button class="btn btn--l btn--blue btn--icon" lx-ripple lx-dropdown-toggle>
          <i class="mdi mdi-dots-vertical"></i>
        </button>

        <lx-dropdown-menu>
          <ul>
              <li><a class="dropdown-link" ui-sref="new%s">New %s</a></li>
              <li class="dropdown-divider"></li>
              <li><a class="dropdown-link" ng-click="%ss.showLogs()">Show Selected Logs</a></li>
              <li><a class="dropdown-link" ng-click="%ss.deleteSelected()">Delete Selected</a></li>
          </ul>
      	</lx-dropdown-menu>
  		</lx-dropdown>
     </div>
  </header>

  <table class="data-table">
    <thead>
      <tr>
				<!-- Generated List of Table Headers -->
`,
		tablename,
		Tablename,
		Tablename, Tablename,
		Tablename,
		tablename, Tablename,
		Tablename,
		Tablename)

	// Generate table headers for each column
	for i, col := range schema {
		if i > 0 {
			str += fmt.Sprintf(`        <th><a ng-click="%ss.setSort('%s')"> %s</a></th>
`,
				Tablename,
				col.UpColumn,
				col.UpColumn)

		}
	}

	str += fmt.Sprintf(`        <th> Edit </th>
      </tr>
    </thead>

    <tbody>
    	<!-- Generated List of TableData Cells -->
      <tr class="data-table__selectable-row"
      		ng-class="%ss.getClass(row)"
      		ng-repeat="row in %ss.%ss | orderBy:%ss.sortField:%ss.sortDir"
      		ng-click="%ss.clickedRow(row)">     
`,
		Tablename,
		Tablename, tablename,
		Tablename, Tablename,
		Tablename,
	)

	for _, col := range schema {
		str += fmt.Sprintf(`            <td><a ng-click="%ss.clickEdit(row)">{{row.%s}}</a></td>
`,
			Tablename,
			col.UpColumn)
	}

	str += fmt.Sprintf(`            <td>
            	<button class="btn btn--m btn--blue btn--icon" 
            					ng-click="%ss.clickEdit(row); $event.stopPropagation()"
            					lx-ripple>
            						<i class="mdi mdi-plus"></i>
            	</button>
            </td>
      </tr>		
    </tbody>
  </table>
</div>
`, Tablename)

	f.WriteString(str)
	f.Sync()
}
