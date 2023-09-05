# ml-metadata-go-server
A go based server that implements a gRPC interface for [ml_metadata](https://github.com/google/ml-metadata/) library.
It adds other features on top of the functionality offered by the gRPC interface.
## Pre-requisites:
- go
- protoc
- python 3.9
## Building
Run the following command to build the server binary:
```
make build
```
The generated binary uses spf13 cmdline args. More information on using the server can be obtained by running the command:
```
./ml-metadata-go-server --help
```
## Creating/Migrating Server DB
The server uses a local SQLite DB file (`metadata.sqlite.db` by default), which can be configured using the `-d` cmdline option.
The following command creates the DB:
```
./ml-metadata-go-server migrate
```
### Loading metadata library
Run the following command to load a metadata library:
```
./ml-metadata-go-server migrate -m config/metadata-library
```
Note that currently no duplicate detection is done as the implementation is a WIP proof of concept. 
Running this command multiple times will create duplicate metadata types. 
To clear the DB simply delete the SQLite DB file `metadata.sqlite.db`. 

### Running Server
Run the following command to start the server:
```
make run/server &
```
### Running Python ml-metadata test client
Run the following command to run the ml-metadata Python test client:
```
make run/client
```
### Clean
Run the following command to clean the DB file, generated gRPC and GraphQL models, etc.:
```
make clean
```