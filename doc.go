/*
Primitives for cascading error handling w/o stacktrace.

In-place replacement for errors.New(), analogue of fmt.Errorf().
Based on "pkg/errors" without a stack trace for simple logging.

Usage

Make a simple error:

	if token == "" {
		return errors.New("token is not provided")
	}

Make an error with a formatted message:

	if _, ok := dict[key]; !ok {
		return errors.Errorf("key %q is not set", key)
	}

Add context to an error:

	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		return errors.Wrap(err, "failed to get records")
	}

Add a formatted annotation to an error:

	db, err := gorm.Open("postgres", url)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to %s", url)
	}

Print a wrapped error message as usual, for example:

	fmt.Println(err)
	// failed to get records: open data.txt: no such file or directory

Retrieve the cause of an error:

	switch errors.Cause(err).(type) {
	case *os.PathError:
		// handle specifically
	default:
		// general error
	}
*/
package errors
