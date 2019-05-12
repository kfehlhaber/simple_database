#!/bin/bash
FAILED=0

fail() {
	echo '!! Test FAIL - ' $1
	let "FAILED += 1"
}

finished() {
  echo ''
  if [[ $FAILED -ne 0 ]]; then
    echo 'There were failed integration tests'
  else
    echo 'All integration tests passed'
  fi
  exit $FAILED
}

testValidSet() {
	echo "Set simple value"
	RESULT=$(curl http://localhost/set?1=value1)
	if [[ $RESULT != 'SUCCESS' ]];
	then
		fail "testValidSet"
	fi
}

testValidGet() {
	echo "Valid get"
	RESULT=$(curl http://localhost/get?key=1)
	if [[ $RESULT != 'value1' ]];
	then
		fail "testValidGet"
	fi
}

testInvalidGet() {
	echo "Invalid get"
	RESULT=$(curl http://localhost/get?wrong-key=1)
	if [[ $RESULT = 'value1' ]];
	then
		fail "testInvalidGet"
	fi
}

doReadWrite() {
	curl http://localhost/set?1=value$1
	RESULT=$(curl http://localhost/get?key=1)
	if [[ -n $RESULT ]];
	then
		fail "massReadWrite"
	fi

}
massReadWrite() {
	echo "mass readwrite"
	for i in {1..100}
	do
		doReadWrite $i &
		pids[${i}]=$!
	done
	for pid in ${pids[*]}; do
    	wait $pid
	done
}
##############################################################
#
# START TEST
#
##############################################################

./build/simple_db

echo Running integration tests...
echo ''

testValidSet
testValidGet
testInvalidGet
massReadWrite

finished



