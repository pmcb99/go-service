# Go with Tests

Hey there! Welcome to my **Go with Tests** repo. This is where I'm diving into Go by writing unit tests. I'm using a test-driven development (TDD) approach to really get the hang of Go, and the plan is to eventually turn this into a full-blown project.

## Achievements So Far

- **Hello World**: Kicked things off by writing a simple function to greet recipients. If no name is provided, it defaults to "Hello stranger...".
- **Unit Testing**: Set up unit tests for the `HelloRecipient` function to ensure it returns the correct greeting. This includes:
  - Greeting a specific person by name.
  - Providing a generic greeting when no name is given.

## Getting Started

First things first, you'll need Go on your machine. If you haven't got it yet, grab it from the [official site](https://golang.org/dl/).

Once you're set up, clone this repo:

```bash
git clone https://github.com/yourusername/go-with-tests.git
cd go-with-tests
```

## Running Tests

To run the tests, just use the `go test` command. It'll run all the test files in the current directory and its subdirectories:

```bash
go test ./...
```

## Contributing

Got ideas or improvements? Awesome! Feel free to open an issue or submit a pull request. Just stick to the existing code style and make sure to include tests for any new stuff.

## License

This project is under the MIT License. Check out the [LICENSE](LICENSE) file for more info.

---

Thanks for checking out my journey. Happy coding and testing with Go!
