package main

import (
	"fmt"
	"log"
	"os"
)

func generateHTML_new(listFile string) {

	f, err := os.Create(listFile)
	if err != nil {
		log.Fatalln("CreateFile:", listFile, err.Error())
	}
	defer f.Close()

	var str = fmt.Sprintf(`<div class="main-section__title">
  <i class="icon icon--xs icon--blue icon--circled mdi mdi-auto-fix"></i>
  Add New %s
</div>

<form name="new%s.form" ng-submit="new%s.submit()">
<formly-form model="new%s.%s" fields="new%s.formFields">

<div flex-container="row" flex-align="end">
    <div flex-item>
      <button type="button" class="btn btn--l btn--blue-grey btn--raised" ng-click="new%s.abort()" lx-ripple>Cancel</button>
    </div>
    <div flex-item>
      <button type="submit" 
              class="btn btn--l btn--blue btn--raised" 
              lx-ripple>
                Add New %s
      </button>
    </div>
</div>

</formly-form>
</form>
`,
		Tablename,
		Tablename, Tablename,
		Tablename, tablename, Tablename,
		Tablename,
		Tablename)

	f.WriteString(str)
	f.Sync()
}
