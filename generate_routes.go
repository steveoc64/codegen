package main

import (
	"fmt"
	"log"
	"os"
)

func generate_Routes(listFile string) {

	f, err := os.Create(listFile)
	if err != nil {
		log.Fatalln("CreateFile:", listFile, err.Error())
	}
	defer f.Close()

	str := ""

	str += fmt.Sprintf(`          .state('%s',{
            url: '/%s',
            acl: 'ACL', // TODO ... declare ACLs for this route
            cache: false,
            templateUrl: 'html/%s.list.html',
            controller: '%sCtrl as %s',
            resolve: {
              %ss: function(DB%s) {
                return DB%s.query()
              },
              logs: function(DBSysLog) {
                return DBSysLog.query({RefType: 'X', Limit: 100})  // TODO ... set the SysLog type for this table
              }
            }
          })
            .state('edit%s',{
              url: '/%s/edit/:id',
              acl: 'ACL',
              templateUrl: 'html/%s.edit.html',
              controller: 'Edit%sCtrl as edit%s',
              resolve: {
                %s: function(DB%s,$stateParams) {
                  return DB%s.get({id: $stateParams.id})
                },
                logs: function(DBSysLog,$stateParams) {
                  return DBSysLog.query({
                    RefType: 'X',   // TODO ... set the SysLog ref type
                    RefID: $stateParams.id,
                    Limit: 100})
                }
              }
            })
            .state('new%s',{
              url: '/new%s',
              acl: 'ACL',
              templateUrl: 'html/%s.new.html',
              controller: 'adminNew%sCtrl as new%s',
            })
`,
		tablename,
		tablename,
		tablename,
		Tablename, Tablename,
		tablename, Tablename,
		Tablename,
		tablename,
		tablename,
		tablename,
		Tablename, Tablename,
		tablename, Tablename, Tablename,
		tablename, tablename, tablename,
		Tablename, Tablename)

	f.WriteString(str)
	f.Sync()
}
