# Characteristics of a Golang test function:

* The first and only parameter needs to be t *testing.T
* It begins with the word Test followed by a word or phrase starting with a capital letter.
* (usually the method under test i.e. TestValidateClient)
* Calls t.Error or t.Fail to indicate a failure (I called t.Errorf to provide more details)
* t.Log can be used to provide non-failing debug information
* Must be saved in a file named something_test.go such as: addition_test.go

# Statement coverage

The go test tool has built-in code-coverage for statements. To try it with out example above type in:

    $ go test -cover

## HTMLify

    go test -cover -coverprofile=c.out
    go tool cover -html=c.out -o coverage.html 