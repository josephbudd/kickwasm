step=0

(( step += 1 ))
echo "STEP ${step}:"
echo "Build kickwasm."
echo "go install"
go install
if [ $? -gt 0 ]
then
    echo "Oops! Your kickwasm code is broken."
    exit
fi

(( step += 1 ))
echo ""
echo "STEP ${step}:"
echo "Build rekickwasm."
echo "cd ../rekickwasm"
echo "go install"
cd ../rekickwasm
go install
if [ $? -gt 0 ]
then
    echo "Oops! Your rekickwasm code is broken."
    exit
fi

(( step += 1 ))
echo ""
echo "STEP ${step}:"
echo "Build kicklpc."
echo "cd ../kicklpc"
echo "go install"
cd ../kicklpc
go install
if [ $? -gt 0 ]
then
    echo "Oops! Your kicklpc code is broken."
    exit
fi

(( step += 1 ))
echo ""
echo "STEP ${step}:"
echo "Build kickstore."
echo "cd ../kickstore"
echo "go install"
cd ../kickstore
go install
if [ $? -gt 0 ]
then
    echo "Oops! Your kickstore code is broken."
    exit
fi

echo ""
echo "EOJ:"
