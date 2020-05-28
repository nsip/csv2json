#!/bin/bash

cd ./Server && ./clean.sh && cd -
cd ./Client && ./clean.sh && cd -

rm -f *.log

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
for f in $(find ./ -name '*.log' -or -name '*.doc'); do rm $f; done
