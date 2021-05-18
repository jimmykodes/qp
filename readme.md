# QP - a query string parser

A golang cli for taking query strings and outputting json objects.

## Install
```shell
# install directly from github
go install github.com/jimmykodes/qp

# or clone first and install
git clone git@github.com:jimmykodes/qp.git
cd qp
go install qp.go
```

## Running
```shell
# basic usage
qp [in_file] [out_file]
# take in a query string from a text file and put the json on stdout
qp query.txt
# save output json to file
qp query.txt output.json
# suggestion: pipe to jq for nice json formatting
qp query.txt | jq
# suggestion (mac specific): go from clipboard to jq
pbpaste | qp | jq
```

## Example:
```shell
qp << EOL
test=1&data=other&bool=false&id=1&id=2
EOL
```
will yield
```json
{"bool":false,"data":"other","id":[1,2],"test":1}
```