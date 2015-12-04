# Code Generator

My own personal code generator for converting SQL structures into workable code
for CRUD based apps.

Backend generated code is for Go, using the DAT toolkit to access the database, and echo webserver

Frontend generated code is for LumX & ngFormly (angular 1 / bourbon / material design / formly forms)


## required:  ~/config.json 
	This is a JSON format file which must contain a field
	DataSourceName:  which includes access information for your SQL database

## Usage

```

$ codegen [-out dirname] -t tablename [-html]

-out dirname		Name of the directory to place generated code into (default = generated)

-t tablename		Name of the SQL table to base the generated code on
-as 				(Optional) Name of the Table in the generated code (default = same as tablename)

Partial Code Generators :

-html  				Generate HTML

-formly 			Generate Formly Form definitions 

-gotype				Generate Go typedef for this table

-go					Generate REST endpoints for this table in Go

-routes				Generate ui-router routing table for this SQL table

-controller			Generate angular controller for List / Edit / New actions

OR

-theworks			Generate ALL of the above

Link Table Generator :

-link table2name	Generate a complete set of HTML / Controller / Backend REST routes
					to provide inline links from the main SQL table to this child SQL table

```


## HTML Generation

Run codegen with the -html flag to generate some basic HTML files

The following HTML files will be generated into the output directrory :

- &lt;tablename&gt;.list.html .....	List the contents of the SQL table in a LumX data_table
- &lt;tablename&gt;.edit.html .....	Edit form for the SQL table
- &lt;tablename&gt;.new.html ......	New record form for the SQL table

```
	Example:  Generate HTML forms to manage List / Edit / New on the inventory table
	$ codegen -t inventory -html
	
	Generates the following new files :
	- generated/inventory.list.html
	- generated/inventory.edit.html
	- generated/inventory.new.html
```

## Formly Definitions

Run codegen with the -formly flag to generate a basic Formly layout in Javascript

The following JS files will be generated into the output directory :

- &lt;tablename&gt;.form.js .....	Define field input types, and create a function that creates a form

## Go Code Generation

Run codegen with -gotype to generate a Go typedef to stdout (useful for inlining a typedef from the editor)

Run codegen with -go to generate a set of REST functions to handle the usual CRUD ops

These functions are complete implementations which perform the required security checks with the JWT token,
exec SQL calls to read / write the data, return JSON results, and log the activites in the
sys_log database table (if needed)

```
	Example:  Generate Go backend code to handle table inventory
	$ codegen -t inventory -go
	
	Generates the following new file :
	- generated/inventory.go

	Which contains :
	- A go type declaration for a new structure 'DBinventory' 

	And the following new functions:
	- func queryInventory (c *echo.Context) error
	- func getInventory (c *echo.Context) error
	- func saveInventory (c *echo.Context) error
	- func deleteInventory (c *echo.Context) error

```

## UI-Router Route Generation

## AngularJS Controller Generation


## Child Table Link Generation

The purpose of this generator is to create a complete set of code to cross reference a child table from parent 
table.

```
	Example:  
	$ codegen -t football_team -link player

	Generates the following new file :
	- generated/football_team-player.link

	Which is a single file containing the following code snippets to cut-n-paste into your application :

	- (Go)  func queryFootballTeamPlayer(c *echo.Context) ... returns JSON result of all players for the given team

	- (JS)  ui-router .resolve statement for a 'player' object which should be pasted into the FootballTeam route state
		    You will need to then add 'player' as a DI variable on the 'FootballTeamCtrl' controller

	- (HTML) A HTML snippet that provides an lx-tab widget, containing a DataTable of all players for the
			 given football team.  You will need to paste this into the existing football_team.edit.html file

	- (Formly)	A new formly field definition  'football_team_id', which is an lx-select control + smarts
			that can be added to any forms for child records of football_team.  A 'player' record in this
			case might use the 'football_team_id' lx-select widget

```