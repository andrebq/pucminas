# Directory structure

### cmd/...

Programs to simulate CRDT, Merkle and Authenticated messages.

### cmd/extractor

Program to read data in the format generated by the arduino measurement, and output it as 
a csv file.

### compose

docker compose files

### results

logs and CSV from executions

### scripts

bash scripts used to run the tests

### dist

final binary files

### arduino

Arduino related files (used to run measurements)

# Building

Just call `make`, to deploy call `make deploy_rpi`.

`make deploy_rpi` will build and copy the **ARM** binaries to `rpi:rpi\*`, update `Makefile` to
change how the copy is executed.