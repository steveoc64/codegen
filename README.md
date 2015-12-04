# Code Generator

My own personal code generator for converting SQL structures into workable code
for CRUD based apps.

Backend generated code is in go

Frontend generated code is in LumX & ngFormly (angular 1 / bourbon / material design / formly forms)


## required:  ~/config.json 
	This is a JSON format file which must contain a field
	DataSourceName:  which includes access information for your SQL database

## Usage

```

$ codegen [-out dirname] -t tablename [-html]

-out dirname		Name of the directory to place generated code into (default = generated)

-t tablename		Name of the SQL table to base the generated code on
-as 				(Optional) Name of the Table in the generated code (default = same as tablename)

-html  				(Optional) Generate HTML

-formly 			(Optional) Generate Formly Form definitions 

-gotype				(Optional) Generate Go typedef for this table

-gorest				(Optional) Generate Go REST endpoints for this table

```


## HTML Generation

Run codegen with the -html flag to generate some basic HTML files

The following super basic HTML files will be generated into the output directrory :

- &lt;tablename&gt;.list.html .....	List the contents of the SQL table in a LumX data_table
- &lt;tablename&gt;.edit.html .....	Edit form for the SQL table
- &lt;tablename&gt;.new.html ......	New record form for the SQL table

## Formly Definitions

Run codegen with the -formly flag to generate a basic Formly layout in Javascript

The following super basic JS files will be generated into the output directory :

- &lt;tablename&gt;.form.js .....	Define field input types, and create a function that creates a form

## Go Code Generation

Run codegen with -gotype to generate a Go typedef to stdout

