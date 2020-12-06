#! /bin/bash

day=$1
if [[ -z $day ]]
then
    echo "Please enter day: ./newday.sh <day_num>"
    exit 1
fi

# Make the input file
touch inputs/day_$day

# Make the go files and replace the day number in the new file
cp -r days/template days/day_$day
sed -i '' "s/-1/$day/g" days/day_$day/main.go
sed -i '' "s/-1/$day/g" days/day_$day/part1_test.go
sed -i '' "s/-1/$day/g" days/day_$day/part2_test.go
sed -i '' "s/template/day_$day/g" days/day_$day/part1_test.go
sed -i '' "s/template/day_$day/g" days/day_$day/part2_test.go
sed -i '' "/t.Skip/d" days/day_$day/part1_test.go
sed -i '' "/t.Skip/d" days/day_$day/part2_test.go
