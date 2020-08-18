# Web-Service for Converting CSV to JSON

## Getting Started

0. Make sure golang dev package is available in your machine.

1. Run `go get -u ./...` to update this project's dependencies.

   Ignore any `undefined: n3cfg.***` errors.

2. Build.

   In 'build.sh', change 'password' in line "sudopwd='password'" to your real sudo/admin password & save.

   Run `build.sh`.

3. Release.

   Run `release.sh 'dest-platform' 'dest-path'`.

   e.g. run `./release.sh [linux64|win64|mac] ~/Desktop/csv2json/linux64/`

4. Docker Deploy (local, only for linux64 platform server).

   Make sure `Docker` installed.

   Jump into your release dest-path in above step.

   e.g. jump into `~/Desktop/csv2json/linux64/`

   Run `docker build --tag n3-csv2json .` to make docker image.

   Run `docker run --name csv2json --net host n3-csv2json:latest` to run docker image.

5. Test.

   Make sure `curl` installed.

   Simple curl test scripts in test.sh.

   Before running `./test.sh`, modify some params like URL (IP, PORT ...) if needed.

   OR write your own curl test like 'test.sh'.
