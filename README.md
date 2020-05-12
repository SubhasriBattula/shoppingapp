# shoppingapp
After writing a dockerfile build the docker using the below command
=>sudo docker build -t shoppingapp .


Then execute this command to run in the interactive mode 
    => sudo docker run --rm -it -v $(pwd):/go/src/app \
             shoppingapp bash\
             
root@ec0d2776bbf7:/go/src/app# 
Then in this root execute go get -u github.com/SubhasriBattula/shoppingapp/supermarket
then run main.go file => go run main.go
