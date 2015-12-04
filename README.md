# Code Generator

My own personal code generator for converting SQL structures into workable code
for CRUD based apps.

Backend generated code is in go

Frontend generated code is in LumX  (angular 1 / bourbon / material design)

## Usage

required:  ~/config.json 
	This is a JSON format file which must contain a field
	DataSourceName:  which includes access information for your SQL database

```

$ codegen [-out dirname] -t tablename [-html]

-out dirname		Name of the directory to place generated code into (default = generated)

-t tablename		Name of the SQL table to base the generated code on

-html  					Generate HTML
```


## HTML Generation

The following HTML files will be generated into the output directrory :

- <tablename>s.html 	List the contents of the SQL table in a LumX data_table

