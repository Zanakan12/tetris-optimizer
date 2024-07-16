#! bin/bash

echo "⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/badexample00.txt
go run . examples/badexample00.txt | grep -q "ERROR"
if [ $? -eq 0 ]; then
    echo "\n✅"
else
    echo "\n❌"
fi
echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/badexample01.txt
go run . examples/badexample01.txt | grep -q "ERROR"
if [ $? -eq 0 ]; then
    echo "\n✅"
else
    echo "\n❌"
fi
echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/badexample02.txt
go run . examples/badexample02.txt | grep -q "ERROR"
if [ $? -eq 0 ]; then
    echo "\n✅"
else
    echo "\n❌"
fi
echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/badexample03.txt
go run . examples/badexample03.txt | grep -q "ERROR"
if [ $? -eq 0 ]; then
    echo "\n✅"
else
    echo "\n❌"
fi
echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/badexample04.txt
go run . examples/badexample04.txt | grep -q "ERROR"
if [ $? -eq 0 ]; then
    echo "\n✅"
else
    echo "\n❌"
fi
echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/badformat.txt
go run . examples/badformat.txt | grep -q "ERROR"
if [ $? -eq 0 ]; then
    echo "\n✅"
else
    echo "\n❌"
fi
echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/example00.txt
count=$(go run . examples/example00.txt | grep -o "⬛" | wc -l)
if [ $count -eq 0 ]; then
    echo "\nFound 0 occurrences of '⬛' ✅"
else
    echo "\nDid not find 0 occurrences of '⬛' ❌"
fi

echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/example01.txt
count=$(go run . examples/example01.txt | grep -o "⬛" | wc -l)
if [ $count -eq 9 ]; then
    echo "\nFound 9 occurrences of '⬛' ✅"
else
    echo "\nDid not find 9 occurrences of '⬛' ❌"
fi

echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/example02.txt
count=$(go run . examples/example02.txt | grep -o "⬛" | wc -l)
if [ $count -eq 4 ]; then
    echo "\nFound 4 occurrences of '⬛' ✅"
else
    echo "\nDid not find 4 occurrences of '⬛' ❌"
fi

echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/example03.txt
count=$(go run . examples/example03.txt | grep -o "⬛" | wc -l)
if [ $count -eq 5 ]; then
    echo "\nFound 5 occurrences of '⬛' ✅"
else
    echo "\nDid not find 5 occurrences of '⬛' ❌"
fi

echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"
go run . examples/hardexample.txt
count=$(go run . examples/hardexample.txt | grep -o "⬛" | wc -l)
if [ $count -eq 1 ]; then
    echo "\nFound 1 occurrences of '⬛' ✅"
else
    echo "\nDid not find 1 occurrences of '⬛' ❌"
fi

echo "\n⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️⬜️\n"